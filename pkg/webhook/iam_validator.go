// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/externalonlygvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type iamValidatorHandler struct {
	smLoader              *servicemappingloader.ServiceMappingLoader
	serviceMetadataLoader metadata.ServiceMetadataLoader
	schemaLoader          dclschemaloader.DCLSchemaLoader
}

func NewIAMValidatorHandler(smLoader *servicemappingloader.ServiceMappingLoader,
	serviceMetadataLoader metadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader) HandlerFunc {
	return func(mgr manager.Manager) admission.Handler {
		return &iamValidatorHandler{
			smLoader:              smLoader,
			serviceMetadataLoader: serviceMetadataLoader,
			schemaLoader:          schemaLoader,
		}
	}
}

func (a *iamValidatorHandler) Handle(_ context.Context, req admission.Request) admission.Response {
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.Object.Raw, nil, obj); err != nil {
		klog.Error(err)
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error decoding object: %w", err))
	}
	switch {
	case isIAMPolicy(obj):
		policy, err := toIAMPolicy(obj)
		if err != nil {
			return admission.Errored(http.StatusInternalServerError, err)
		}
		refResourceGVK := policy.Spec.ResourceReference.GroupVersionKind()
		isDCLResource := metadata.IsDCLBasedResourceKind(refResourceGVK, a.serviceMetadataLoader)
		return a.validateIAMPolicy(policy, isDCLResource)

	case isIAMPartialPolicy(obj):
		partialPolicy, err := toIAMPartialPolicy(obj)
		if err != nil {
			return admission.Errored(http.StatusInternalServerError, err)
		}
		refResourceGVK := partialPolicy.Spec.ResourceReference.GroupVersionKind()
		isDCLResource := metadata.IsDCLBasedResourceKind(refResourceGVK, a.serviceMetadataLoader)
		return a.validateIAMPartialPolicy(partialPolicy, isDCLResource)

	case isIAMPolicyMember(obj):
		policyMember, err := toIAMPolicyMember(obj)
		if err != nil {
			return admission.Errored(http.StatusInternalServerError, err)
		}
		refResourceGVK := policyMember.Spec.ResourceReference.GroupVersionKind()
		isDCLResource := metadata.IsDCLBasedResourceKind(refResourceGVK, a.serviceMetadataLoader)
		return a.validateIAMPolicyMember(policyMember, isDCLResource)
	case isIAMAuditConfig(obj):
		auditConfig, err := toIAMAuditConfig(obj)
		if err != nil {
			return admission.Errored(http.StatusInternalServerError, err)
		}
		refResourceGVK := auditConfig.Spec.ResourceReference.GroupVersionKind()
		isDCLResource := metadata.IsDCLBasedResourceKind(refResourceGVK, a.serviceMetadataLoader)
		if isDCLResource {
			return admission.Errored(http.StatusForbidden,
				fmt.Errorf("object of GroupVersionKind %v does not have IAM Audit Config support", obj.GroupVersionKind()))
		}
		rcs, err := getResourceConfigs(a.smLoader, refResourceGVK)
		if err != nil {
			return admission.Errored(http.StatusBadRequest, err)
		}
		return validateIAMAuditConfig(auditConfig, rcs)
	default:
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("object of GroupVersionKind %v is not a supported IAM resource", obj.GroupVersionKind()))
	}
}

func toIAMPolicy(obj *unstructured.Unstructured) (*v1beta1.IAMPolicy, error) {
	policy := &v1beta1.IAMPolicy{}
	if err := util.Marshal(obj, policy); err != nil {
		return nil, fmt.Errorf("error parsing %v into IAM Policy object: %w", obj.GetName(), err)
	}
	return policy, nil
}

func toIAMPartialPolicy(obj *unstructured.Unstructured) (*v1beta1.IAMPartialPolicy, error) {
	partialPolicy := &v1beta1.IAMPartialPolicy{}
	if err := util.Marshal(obj, partialPolicy); err != nil {
		return nil, fmt.Errorf("error parsing %v into IAMPartialPolicy object: %w", obj.GetName(), err)
	}
	return partialPolicy, nil
}

func toIAMPolicyMember(obj *unstructured.Unstructured) (*v1beta1.IAMPolicyMember, error) {
	policyMember := &v1beta1.IAMPolicyMember{}
	if err := util.Marshal(obj, policyMember); err != nil {
		return nil, fmt.Errorf("error parsing %v into IAM Policy Member object: %w", obj.GetName(), err)
	}
	return policyMember, nil
}

func toIAMAuditConfig(obj *unstructured.Unstructured) (*v1beta1.IAMAuditConfig, error) {
	auditConfig := &v1beta1.IAMAuditConfig{}
	if err := util.Marshal(obj, auditConfig); err != nil {
		return nil, fmt.Errorf("error parsing %v into IAMAuditConfig object: %w", obj.GetName(), err)
	}
	return auditConfig, nil
}

func getDCLSchema(gvk schema.GroupVersionKind, serviceMetadataLoader metadata.ServiceMetadataLoader, schemaLoader dclschemaloader.DCLSchemaLoader) (*openapi.Schema, admission.Response) {
	dclSchema, err := dclschemaloader.GetDCLSchemaForGVK(gvk, serviceMetadataLoader, schemaLoader)
	if err != nil {
		return nil, admission.Errored(http.StatusBadRequest, err)
	}
	return dclSchema, allowedResponse
}

func getResourceConfigs(smLoader *servicemappingloader.ServiceMappingLoader, gvk schema.GroupVersionKind) ([]*v1alpha1.ResourceConfig, error) {
	// Support the case where the user specified the kind as "Project" and nothing else.
	// TODO(kcc-eng): Remove once we drop headless IAM support.
	if gvk.Group == "" {
		if gvk.Kind == kcciamclient.ProjectKind {
			gvk = kcciamclient.ProjectGVK
		} else {
			return []*v1alpha1.ResourceConfig{}, fmt.Errorf("resource reference for kind '%v' must include API group", gvk.Kind)
		}
	}
	if externalonlygvks.IsExternalOnlyGVK(gvk) {
		rc, err := kcciamclient.GetResourceConfigForExternalOnlyGVK(gvk)
		if err != nil {
			return []*v1alpha1.ResourceConfig{}, fmt.Errorf("error getting ResourceConfig for external GroupVersionKind %v: %w", gvk, err)
		}
		return []*v1alpha1.ResourceConfig{rc}, nil
	}
	rcs, err := smLoader.GetResourceConfigs(gvk)
	if err != nil {
		return []*v1alpha1.ResourceConfig{}, fmt.Errorf("error getting ResourceConfig for GroupVersionKind %v: %w", gvk, err)
	}
	if len(rcs) == 0 {
		return []*v1alpha1.ResourceConfig{}, fmt.Errorf("couldn't find any ResourceConfig defined for GroupVersionKind %v", gvk)
	}
	return rcs, nil
}

func (a *iamValidatorHandler) validateIAMPolicy(policy *v1beta1.IAMPolicy, isDCLResource bool) admission.Response {
	resourceRef := policy.Spec.ResourceReference
	if isDCLResource {
		return a.dclValidateIAMPolicy(policy)
	}

	// TF-based resource.
	rcs, err := getResourceConfigs(a.smLoader, resourceRef.GroupVersionKind())
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	return a.tfValidateIAMPolicy(policy, rcs)
}

func (a *iamValidatorHandler) validateIAMPartialPolicy(partialPolicy *v1beta1.IAMPartialPolicy, isDCLResource bool) admission.Response {
	resourceRef := partialPolicy.Spec.ResourceReference
	if isDCLResource {
		return a.dclValidateIAMPartialPolicy(partialPolicy)
	}
	// TF-based resource.
	rcs, err := getResourceConfigs(a.smLoader, resourceRef.GroupVersionKind())
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	return a.tfValidateIAMPartialPolicy(partialPolicy, rcs)
}

func (a *iamValidatorHandler) validateIAMPolicyMember(policyMember *v1beta1.IAMPolicyMember, isDCLResource bool) admission.Response {
	resourceRef := policyMember.Spec.ResourceReference
	if isDCLResource {
		return a.dclValidateIAMPolicyMember(policyMember)
	}
	// TF-based resource.
	rcs, err := getResourceConfigs(a.smLoader, resourceRef.GroupVersionKind())
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	return a.tfValidateIAMPolicyMember(policyMember, rcs)
}

func validateIAMAuditConfig(auditConfig *v1beta1.IAMAuditConfig, refResourceRCs []*v1alpha1.ResourceConfig) admission.Response {
	resourceRef := auditConfig.Spec.ResourceReference
	if !doesTFResourceSupportAuditConfigs(refResourceRCs) {
		return admission.Errored(http.StatusForbidden,
			fmt.Errorf("GroupVersionKind %v does not support IAM Audit Configs", resourceRef.GroupVersionKind()))
	}
	return allowedResponse
}

func (a *iamValidatorHandler) dclValidateIAMPolicy(policy *v1beta1.IAMPolicy) admission.Response {
	resourceRef := policy.Spec.ResourceReference
	// Check that DCL-based resource supports IAMPolicy
	dclSchema, resp := getDCLSchema(resourceRef.GroupVersionKind(), a.serviceMetadataLoader, a.schemaLoader)
	if !resp.Allowed {
		return resp
	}
	supportsIAM, err := extension.HasIam(dclSchema)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}
	if !supportsIAM {
		return admission.Errored(http.StatusForbidden, fmt.Errorf("GroupVersionKind %v does not support IAM Policy", resourceRef.GroupVersionKind()))
	}
	// Currently, DCL-based resources that have IAMPolicy also support IAMConditions
	// and we don't need to check for conditions.
	// TODO: (b/182505291) DCL-based resources do not currently support IAMAuditConfigs
	if len(policy.Spec.AuditConfigs) > 0 {
		return admission.Errored(http.StatusForbidden, fmt.Errorf("GroupVersionKind %v does not support IAM Audit Configs", resourceRef.GroupVersionKind()))
	}
	return allowedResponse
}

func (a *iamValidatorHandler) tfValidateIAMPolicy(policy *v1beta1.IAMPolicy, rcs []*v1alpha1.ResourceConfig) admission.Response {
	resourceRef := policy.Spec.ResourceReference
	if doesIAMPolicyHaveConditions(policy) && !doesTFResourceSupportConditions(rcs) {
		return admission.Errored(http.StatusForbidden,
			fmt.Errorf("GroupVersionKind %v does not support IAM Conditions", resourceRef.GroupVersionKind()))
	}
	if len(policy.Spec.AuditConfigs) > 0 && !doesTFResourceSupportAuditConfigs(rcs) {
		return admission.Errored(http.StatusForbidden,
			fmt.Errorf("GroupVersionKind %v does not support IAM Audit Configs", resourceRef.GroupVersionKind()))
	}
	return allowedResponse
}

func (a *iamValidatorHandler) dclValidateIAMPartialPolicy(partialPolicy *v1beta1.IAMPartialPolicy) admission.Response {
	resourceRef := partialPolicy.Spec.ResourceReference
	// Check that DCL-based resource supports IAMPolicy
	dclSchema, resp := getDCLSchema(resourceRef.GroupVersionKind(), a.serviceMetadataLoader, a.schemaLoader)
	if !resp.Allowed {
		return resp
	}
	supportsIAM, err := extension.HasIam(dclSchema)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}
	if !supportsIAM {
		return admission.Errored(http.StatusForbidden, fmt.Errorf("GroupVersionKind %v does not support IAM Partial Policy", resourceRef.GroupVersionKind()))
	}
	return allowedResponse
}

func (a *iamValidatorHandler) tfValidateIAMPartialPolicy(partialPolicy *v1beta1.IAMPartialPolicy, rcs []*v1alpha1.ResourceConfig) admission.Response {
	resourceRef := partialPolicy.Spec.ResourceReference
	if doesIAMPartialPolicyHaveConditions(partialPolicy) && !doesTFResourceSupportConditions(rcs) {
		return admission.Errored(http.StatusForbidden,
			fmt.Errorf("GroupVersionKind %v does not support IAM Conditions", resourceRef.GroupVersionKind()))
	}
	return allowedResponse
}

func (a *iamValidatorHandler) dclValidateIAMPolicyMember(policyMember *v1beta1.IAMPolicyMember) admission.Response {
	resourceRef := policyMember.Spec.ResourceReference

	// Check that DCL-based resource supports IAMPolicy
	dclSchema, resp := getDCLSchema(resourceRef.GroupVersionKind(), a.serviceMetadataLoader, a.schemaLoader)
	if !resp.Allowed {
		return resp
	}
	supportsIAM, err := extension.HasIam(dclSchema)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}
	// Beginnings of direct IAM support: direct-IAM added to existing DCL resource
	if registry.IsIAMDirect(resourceRef.GroupVersionKind().GroupKind()) {
		supportsIAM = true
	}
	if !supportsIAM {
		return admission.Errored(http.StatusForbidden, fmt.Errorf("GroupVersionKind %v does not support IAM Policy Member", resourceRef.GroupVersionKind()))
	}
	// TODO (b/228226694): IAMPolicyMember does not currently support conditions.
	if doesIAMPolicyMemberHaveCondition(policyMember) {
		return admission.Errored(http.StatusForbidden,
			fmt.Errorf("GroupVersionKind %v does not support IAM Conditions in IAM Policy Member", resourceRef.GroupVersionKind()))
	}
	return allowedResponse
}

func (a *iamValidatorHandler) tfValidateIAMPolicyMember(policyMember *v1beta1.IAMPolicyMember, rcs []*v1alpha1.ResourceConfig) admission.Response {
	resourceRef := policyMember.Spec.ResourceReference
	if doesIAMPolicyMemberHaveCondition(policyMember) && !doesTFResourceSupportConditions(rcs) {
		return admission.Errored(http.StatusForbidden,
			fmt.Errorf("GroupVersionKind %v does not support IAM Conditions", resourceRef.GroupVersionKind()))
	}
	return allowedResponse
}

func doesIAMPolicyHaveConditions(policy *v1beta1.IAMPolicy) bool {
	for _, binding := range policy.Spec.Bindings {
		if binding.Condition != nil {
			return true
		}
	}
	return false
}

func doesIAMPartialPolicyHaveConditions(partialPolicy *v1beta1.IAMPartialPolicy) bool {
	for _, binding := range partialPolicy.Spec.Bindings {
		if binding.Condition != nil {
			return true
		}
	}
	return false
}

func doesIAMPolicyMemberHaveCondition(policyMember *v1beta1.IAMPolicyMember) bool {
	return policyMember.Spec.Condition != nil
}

func doesTFResourceSupportConditions(rcs []*v1alpha1.ResourceConfig) bool {
	// All ResourceConfigs for a given kind have the same value for IAMConfig.SupportsConditions.
	return rcs[0].IAMConfig.SupportsConditions
}

func doesTFResourceSupportAuditConfigs(rcs []*v1alpha1.ResourceConfig) bool {
	// All ResourceConfigs for a given kind support or don't support IAM audit configs.
	return rcs[0].IAMConfig.AuditConfigName != ""
}
