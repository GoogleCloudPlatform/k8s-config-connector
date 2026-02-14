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

package iamclient

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/externalonlygvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type TFIAMClient struct {
	kubeClient client.Client
	provider   *tfschema.Provider
	smLoader   *servicemappingloader.ServiceMappingLoader
}

func (t *TFIAMClient) SetPolicyMember(ctx context.Context, policyMember *v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error) {
	rc, err := t.getResourceConfigForReferencedResource(ctx, policyMember)
	if err != nil {
		return nil, fmt.Errorf("error getting resource config for referenced resource: %w", err)
	}
	if !resourceSupportsIAMPolicyMember(rc) {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM policies", rc.Kind)
	}
	resource, err := t.newResource(ctx, policyMember, rc)
	if err != nil {
		return nil, err
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, t.provider, t.kubeClient, t.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error fetching live state for resource: %w", err)
	}
	cfg, _, err := krmtotf.KRMResourceToTFResourceConfig(resource, t.kubeClient, t.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error creating resource config: %w", err)
	}
	diff, err := resource.TFResource.Diff(ctx, liveState, cfg, t.provider.Meta())
	if err != nil {
		return nil, fmt.Errorf("error calculating diff: %w", err)
	}
	if !liveState.Empty() && diff.RequiresNew() {
		return nil, k8s.NewImmutableFieldsMutationError(tfresource.ImmutableFieldsFromDiff(diff))
	}
	if diff.Empty() {
		logger.V(2).Info("underlying resource is already up to date", "resource", k8s.GetNamespacedName(policyMember))
		return policyMember, nil
	}

	// Report diff to structured-reporting subsystem
	{
		report := &structuredreporting.Diff{}
		u, err := resource.MarshalAsUnstructured()
		if err != nil {
			log := log.FromContext(ctx)
			log.Error(err, "error reporting diff")
		}
		report.Object = u
		if diff != nil {
			for k, attr := range diff.Attributes {
				report.Fields = append(report.Fields, structuredreporting.DiffField{
					ID:  k,
					Old: attr.Old,
					New: attr.New,
				})
			}
		}
		report.IsNewObject = liveState.Empty()
		structuredreporting.ReportDiff(ctx, report)
	}

	newState, diagnostics := resource.TFResource.Apply(ctx, liveState, diff, t.provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return nil, fmt.Errorf("error applying changes: %w", err)
	}
	return newIAMPolicyMemberFromTFState(resource, newState, policyMember)
}

func (t *TFIAMClient) GetPolicyMember(ctx context.Context, policyMember *v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error) {
	rc, err := t.getResourceConfigForReferencedResource(ctx, policyMember)
	if err != nil {
		return nil, fmt.Errorf("error getting resource config for referenced resource: %w", err)
	}
	if !resourceSupportsIAMPolicyMember(rc) {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM policies", rc.Kind)
	}
	resource, err := t.newResourceSkeleton(ctx, policyMember, rc)
	if err != nil {
		return nil, fmt.Errorf("error building resource skeleton for getting IAM resource: %w", err)
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, t.provider, t.kubeClient, t.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error fetching live state for resource: %w", err)
	}
	if liveState.Empty() {
		return nil, ErrNotFound
	}
	return newIAMPolicyMemberFromTFState(resource, liveState, policyMember)
}

func (t *TFIAMClient) DeletePolicyMember(ctx context.Context, policyMember *v1beta1.IAMPolicyMember) error {
	rc, err := t.getResourceConfigForReferencedResource(ctx, policyMember)
	if err != nil {
		return fmt.Errorf("error getting resource config for referenced resource: %w", err)
	}
	if !resourceSupportsIAMPolicyMember(rc) {
		return fmt.Errorf("invalid resource reference: kind %v does not support IAM policies", rc.Kind)
	}
	resource, err := t.newResource(ctx, policyMember, rc)
	if err != nil {
		return err
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, t.provider, t.kubeClient, t.smLoader)
	if err != nil {
		return fmt.Errorf("error fetching live state for resource: %w", err)
	}
	if liveState.Empty() {
		return ErrNotFound
	}
	_, diagnostics := resource.TFResource.Apply(ctx, liveState, &terraform.InstanceDiff{Destroy: true}, t.provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return fmt.Errorf("error deleting IAMPolicyMember: %w", err)
	}
	return nil
}

func (t *TFIAMClient) SetPolicy(ctx context.Context, policy *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	rc, err := t.getResourceConfigForReferencedResource(ctx, policy)
	if err != nil {
		return nil, fmt.Errorf("error getting resource config for referenced resource: %w", err)
	}
	if !resourceSupportsIAMPolicy(rc) {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM policies", rc.Kind)
	}
	if len(policy.Spec.AuditConfigs) > 0 && !resourceSupportsIAMAuditConfigs(rc) {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM audit configs", rc.Kind)
	}
	resource, err := t.newResource(ctx, policy, rc)
	if err != nil {
		return nil, err
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, t.provider, t.kubeClient, t.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error fetching live state for resource")
	}
	cfg, _, err := krmtotf.KRMResourceToTFResourceConfig(resource, t.kubeClient, t.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error creating resource config: %w", err)
	}
	diff, err := resource.TFResource.Diff(ctx, liveState, cfg, t.provider.Meta())
	if err != nil {
		return nil, fmt.Errorf("error calculating diff: %w", err)
	}
	if !liveState.Empty() && diff.RequiresNew() {
		return nil, k8s.NewImmutableFieldsMutationError(tfresource.ImmutableFieldsFromDiff(diff))
	}
	if diff.Empty() {
		logger.V(2).Info("underlying resource is already up to date", "resource", k8s.GetNamespacedName(policy))
		return policy, nil
	}
	newState, diagnostics := resource.TFResource.Apply(ctx, liveState, diff, t.provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return nil, fmt.Errorf("error applying changes: %w", err)
	}
	return newIAMPolicyFromTFState(resource, newState, policy)
}

func (t *TFIAMClient) GetPolicy(ctx context.Context, policy *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	rc, err := t.getResourceConfigForReferencedResource(ctx, policy)
	if err != nil {
		return nil, fmt.Errorf("error getting resource config for referenced resource: %w", err)
	}
	if !resourceSupportsIAMPolicy(rc) {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM policies", rc.Kind)
	}
	if len(policy.Spec.AuditConfigs) > 0 && !resourceSupportsIAMAuditConfigs(rc) {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM audit configs", rc.Kind)
	}
	resource, err := t.newResourceSkeleton(ctx, policy, rc)
	if err != nil {
		return nil, fmt.Errorf("error building resource skeleton for getting IAM resource: %w", err)
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, t.provider, t.kubeClient, t.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error fetching live state for resource: %w", err)
	}
	if liveState.Empty() {
		return nil, ErrNotFound
	}
	return newIAMPolicyFromTFState(resource, liveState, policy)
}

func (t *TFIAMClient) DeletePolicy(ctx context.Context, policy *v1beta1.IAMPolicy) error {
	rc, err := t.getResourceConfigForReferencedResource(ctx, policy)
	if err != nil {
		return fmt.Errorf("error getting resource config for referenced resource: %w", err)
	}
	if !resourceSupportsIAMPolicy(rc) {
		return fmt.Errorf("invalid resource reference: kind %v does not support IAM policies", rc.Kind)
	}
	if len(policy.Spec.AuditConfigs) > 0 && !resourceSupportsIAMAuditConfigs(rc) {
		return fmt.Errorf("invalid resource reference: kind %v does not support IAM audit configs", rc.Kind)
	}
	resource, err := t.newResource(ctx, policy, rc)
	if err != nil {
		return err
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, t.provider, t.kubeClient, t.smLoader)
	if err != nil {
		return fmt.Errorf("error fetching live state for resource: %w", err)
	}
	if liveState.Empty() {
		return ErrNotFound
	}
	_, diagnostics := resource.TFResource.Apply(ctx, liveState, &terraform.InstanceDiff{Destroy: true}, t.provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return fmt.Errorf("error deleting IAMPolicy: %w", err)
	}
	return nil
}

func (t *TFIAMClient) SetAuditConfig(ctx context.Context, auditConfig *v1beta1.IAMAuditConfig) (*v1beta1.IAMAuditConfig, error) {
	rc, err := t.getResourceConfigForReferencedResource(ctx, auditConfig)
	if err != nil {
		return nil, fmt.Errorf("error getting resource config for referenced resource: %w", err)
	}
	if !resourceSupportsIAMAuditConfigs(rc) {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM audit configs", rc.Kind)
	}
	resource, err := t.newResource(ctx, auditConfig, rc)
	if err != nil {
		return nil, err
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, t.provider, t.kubeClient, t.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error fetching live state for resource: %w", err)
	}
	cfg, _, err := krmtotf.KRMResourceToTFResourceConfig(resource, t.kubeClient, t.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error creating resource config: %w", err)
	}
	diff, err := resource.TFResource.Diff(ctx, liveState, cfg, t.provider.Meta())
	if err != nil {
		return nil, fmt.Errorf("error calculating diff: %w", err)
	}
	if !liveState.Empty() && diff.RequiresNew() {
		return nil, k8s.NewImmutableFieldsMutationError(tfresource.ImmutableFieldsFromDiff(diff))
	}
	if diff.Empty() {
		logger.V(2).Info("underlying resource is already up to date", "resource", k8s.GetNamespacedName(auditConfig))
		return auditConfig, nil
	}
	newState, diagnostics := resource.TFResource.Apply(ctx, liveState, diff, t.provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return nil, fmt.Errorf("error applying changes: %w", err)
	}
	return newIAMAuditConfigFromTFState(resource, newState, auditConfig)
}

func (t *TFIAMClient) GetAuditConfig(ctx context.Context, auditConfig *v1beta1.IAMAuditConfig) (*v1beta1.IAMAuditConfig, error) {
	rc, err := t.getResourceConfigForReferencedResource(ctx, auditConfig)
	if err != nil {
		return nil, fmt.Errorf("error getting resource config for referenced resource: %w", err)
	}
	if !resourceSupportsIAMAuditConfigs(rc) {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM audit configs", rc.Kind)
	}
	resource, err := t.newResourceSkeleton(ctx, auditConfig, rc)
	if err != nil {
		return nil, fmt.Errorf("error building resource skeleton for getting IAM resource: %w", err)
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, t.provider, t.kubeClient, t.smLoader)
	if err != nil {
		return nil, fmt.Errorf("error fetching live state for resource: %w", err)
	}
	if liveState.Empty() {
		return nil, ErrNotFound
	}
	return newIAMAuditConfigFromTFState(resource, liveState, auditConfig)
}

func (t *TFIAMClient) DeleteAuditConfig(ctx context.Context, auditConfig *v1beta1.IAMAuditConfig) error {
	rc, err := t.getResourceConfigForReferencedResource(ctx, auditConfig)
	if err != nil {
		return fmt.Errorf("error getting resource config for referenced resource: %w", err)
	}
	if !resourceSupportsIAMAuditConfigs(rc) {
		return fmt.Errorf("invalid resource reference: kind %v does not support IAM audit configs", rc.Kind)
	}
	resource, err := t.newResource(ctx, auditConfig, rc)
	if err != nil {
		return err
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, t.provider, t.kubeClient, t.smLoader)
	if err != nil {
		return fmt.Errorf("error fetching live state for resource: %w", err)
	}
	if liveState.Empty() {
		return ErrNotFound
	}
	_, diagnostics := resource.TFResource.Apply(ctx, liveState, &terraform.InstanceDiff{Destroy: true}, t.provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return fmt.Errorf("error deleting IAMAuditConfig: %w", err)
	}
	return nil
}

// newResource creates a Resource which represents the given IAM object.
func (t *TFIAMClient) newResource(ctx context.Context, iamInterface interface{}, rc *corekccv1alpha1.ResourceConfig) (*krmtotf.Resource, error) {
	SetGVK(iamInterface)
	unstructSkeleton, err := t.newUnstructuredSkeleton(ctx, iamInterface, rc)
	if err != nil {
		return nil, fmt.Errorf("error building a minimal unstructured skeleton for the IAM resource: %w", err)
	}

	// Add any extra fields from the IAM object not included in the unstructured skeleton
	switch iamObject := iamInterface.(type) {
	case *v1beta1.IAMPolicy:
		krmPolicyUnstruct, err := k8s.MarshalObjectAsUnstructured(iamObject)
		if err != nil {
			return nil, err
		}
		spec := krmPolicyUnstruct.Object["spec"].(map[string]interface{})
		policyDataMap := map[string]interface{}{
			"bindings":     spec["bindings"],
			"auditConfigs": spec["auditConfigs"],
		}
		if iamObject.Spec.Etag != "" {
			policyDataMap["etag"] = iamObject.Spec.Etag
		}
		b, err := json.Marshal(policyDataMap)
		if err != nil {
			return nil, fmt.Errorf("error marshalling policy data to JSON: %w", err)
		}
		if err := unstructured.SetNestedField(unstructSkeleton.Object, string(b), "spec", "policyData"); err != nil {
			return nil, err
		}
	case *v1beta1.IAMPolicyMember:
		// An unstructured skeleton for getting an IAM policy member already
		// fully represents the IAM policy member.
		break
	case *v1beta1.IAMAuditConfig:
		auditLogConfigs := iamObject.Spec.AuditLogConfigs
		var auditLogConfigSlice []interface{}
		if err := util.Marshal(auditLogConfigs, &auditLogConfigSlice); err != nil {
			return nil, fmt.Errorf("unable to marshal %v to slice: %w", reflect.TypeOf(auditLogConfigs).Name(), err)
		}
		if err := unstructured.SetNestedSlice(unstructSkeleton.Object, auditLogConfigSlice, "spec", "auditLogConfig"); err != nil {
			return nil, err
		}
	default:
		panic(fmt.Errorf("unknown type: %v", reflect.TypeOf(iamInterface).Name()))
	}

	iamSM, err := t.newServiceMappingForAssociatedIAMInterface(ctx, iamInterface, rc.IAMConfig)
	if err != nil {
		return nil, fmt.Errorf("error building service mapping for IAM resource: %w", err)
	}
	return krmtotf.NewResource(unstructSkeleton, iamSM, t.provider)
}

// newResourceSkeleton creates a Resource for the given IAM object which
// contains just enough fields to perform a Terraform read.
func (t *TFIAMClient) newResourceSkeleton(ctx context.Context, iamInterface interface{}, rc *corekccv1alpha1.ResourceConfig) (*krmtotf.Resource, error) {
	SetGVK(iamInterface)
	unstructSkeleton, err := t.newUnstructuredSkeleton(ctx, iamInterface, rc)
	if err != nil {
		return nil, fmt.Errorf("error building unstructured skeleton for getting IAM resource: %w", err)
	}
	iamSM, err := t.newServiceMappingForAssociatedIAMInterface(ctx, iamInterface, rc.IAMConfig)
	if err != nil {
		return nil, fmt.Errorf("error building service mapping for IAM resource: %w", err)
	}
	return krmtotf.NewResource(unstructSkeleton, iamSM, t.provider)
}

// newUnstructuredSkeleton creates an Unstructured for the given IAM
// object which contains just enough fields to perform a Terraform read.
func (t *TFIAMClient) newUnstructuredSkeleton(ctx context.Context, iamInterface interface{}, rc *corekccv1alpha1.ResourceConfig) (*unstructured.Unstructured, error) {
	namespace, resourceRef := extractNamespaceAndResourceReference(iamInterface)
	unstructSkeleton, err := t.buildUnstructuredIAMSkeletonFromReference(ctx, iamInterface, namespace, resourceRef, rc)
	if err != nil {
		return nil, fmt.Errorf("error building unstructured skeleton from referenced resource: %w", err)
	}
	switch iamObject := iamInterface.(type) {
	case *v1beta1.IAMPolicy:
		// IAM policies are uniquely identified by the resource they are referencing.
		// No extra information is required to be able to get a policy from GCP.
		return unstructSkeleton, nil
	case *v1beta1.IAMPolicyMember:
		// IAM policy members are uniquely identified by the tuple of (member, role, conditions).
		member, err := ResolveMemberIdentity(ctx, iamObject.Spec.Member, iamObject.Spec.MemberFrom, iamObject.Namespace, t)
		if err != nil {
			return nil, fmt.Errorf("could not resolve member identity for IAMPolicyMember: %w", err)
		}
		if err := unstructured.SetNestedField(unstructSkeleton.Object, member, "spec", "member"); err != nil {
			return nil, err
		}
		if err := unstructured.SetNestedField(unstructSkeleton.Object, iamObject.Spec.Role, "spec", "role"); err != nil {
			return nil, err
		}
		if iamObject.Spec.Condition != nil {
			condition := iamObject.Spec.Condition
			var conditionMap map[string]interface{}
			if err := util.Marshal(condition, &conditionMap); err != nil {
				return nil, fmt.Errorf("unable to marshal %v to map: %w", reflect.TypeOf(condition).Name(), err)
			}
			if err := unstructured.SetNestedMap(unstructSkeleton.Object, conditionMap, "spec", "condition"); err != nil {
				return nil, err
			}
		}
		return unstructSkeleton, nil
	case *v1beta1.IAMAuditConfig:
		// IAM audit configs are uniquely identified by their service field.
		if err := unstructured.SetNestedField(unstructSkeleton.Object, iamObject.Spec.Service, "spec", "service"); err != nil {
			return nil, err
		}
		return unstructSkeleton, nil
	default:
		panic(fmt.Errorf("unknown type: %v", reflect.TypeOf(iamInterface).Name()))
	}
}

func (t *TFIAMClient) buildUnstructuredIAMSkeletonFromReference(ctx context.Context, iamInterface interface{}, namespace string,
	resourceRef v1beta1.ResourceReference, rc *corekccv1alpha1.ResourceConfig) (*unstructured.Unstructured, error) {
	id, err := t.getResourceID(ctx, resourceRef, namespace)
	if err != nil {
		return nil, fmt.Errorf("couldn't get resource id for resource reference: %w", err)
	}

	iamObject := t.newIAMObjectFromInterface(iamInterface)
	unstruct, err := k8s.MarshalObjectAsUnstructured(iamObject)
	if err != nil {
		return nil, err
	}

	if externalonlygvks.IsExternalOnlyGVK(resourceRef.GroupVersionKind()) {
		return unstructuredIAMSkeletonForExternalOnlyRef(resourceRef, unstruct)
	}

	tfResourceName := rc.Name
	tfInfo := &terraform.InstanceInfo{
		Type: tfResourceName,
	}
	tfStateParsedFromID, err := krmtotf.ImportState(ctx, id, tfInfo, t.provider)
	if err != nil {
		return nil, fmt.Errorf("failed to import state given id %v: %w", id, err)
	}

	refFieldName := text.SnakeCaseToLowerCamelCase(rc.IAMConfig.ReferenceField.Name)
	refFieldVal, err := getReferenceFieldValue(id, rc)
	if err != nil {
		return nil, fmt.Errorf("failed to get reference field value given id %v: %w", id, err)
	}
	spec := map[string]interface{}{
		refFieldName: refFieldVal,
	}
	for k, v := range tfStateParsedFromID.Attributes {
		spec[k] = v
	}

	unstruct.Object["spec"] = spec
	return unstruct, nil
}

func (t *TFIAMClient) getResourceID(ctx context.Context, resourceRef v1beta1.ResourceReference, namespace string) (string, error) {
	if resourceRef.External != "" {
		return resourceRef.External, nil
	}

	nn := types.NamespacedName{
		Namespace: useIfNonEmptyElseDefaultTo(resourceRef.Namespace, namespace),
		Name:      resourceRef.Name,
	}
	refResource, err := t.getResource(ctx, resourceRef.GroupVersionKind(), nn)
	if err != nil {
		return "", err
	}
	id, err := refResource.GetImportID(t.kubeClient, t.smLoader)
	if err != nil {
		return "", fmt.Errorf("error getting import ID for referenced resource: %w", err)
	}
	if id != "" {
		return id, nil
	}
	return "", fmt.Errorf("couldn't construct id for referenced resource")
}

func (t *TFIAMClient) getResourceConfigForReferencedResource(ctx context.Context, iamInterface interface{}) (*corekccv1alpha1.ResourceConfig, error) {
	namespace, resourceRef := extractNamespaceAndResourceReference(iamInterface)
	if resourceRef.External != "" {
		return t.getResourceConfigForExternalRef(resourceRef)
	}
	nn := types.NamespacedName{
		Namespace: useIfNonEmptyElseDefaultTo(resourceRef.Namespace, namespace),
		Name:      resourceRef.Name,
	}
	refResource, err := t.getResource(ctx, resourceRef.GroupVersionKind(), nn)
	if err != nil {
		return nil, err
	}
	return &refResource.ResourceConfig, nil
}

func (t *TFIAMClient) getResourceConfigForExternalRef(resourceRef v1beta1.ResourceReference) (*corekccv1alpha1.ResourceConfig, error) {
	if resourceRef.External == "" {
		return nil, fmt.Errorf("external field is empty")
	}
	gvk := resourceRef.GroupVersionKind()
	if externalonlygvks.IsExternalOnlyGVK(gvk) {
		return GetResourceConfigForExternalOnlyGVK(gvk)
	}
	sm, err := t.smLoader.GetServiceMapping(gvk.Group)
	if err != nil {
		return nil, fmt.Errorf("failed to load ServiceMapping for GroupVersionKind %v: %w", gvk, err)
	}
	rcs := servicemappingloader.GetResourceConfigsForKind(sm, gvk.Kind)
	switch len(rcs) {
	case 0:
		return nil, fmt.Errorf("couldn't find any ResourceConfig defined for GroupVersionKind %v", gvk)
	case 1:
		return rcs[0], nil
	// It is possible for a kind to have multiple ResourceConfigs (e.g.
	// ResourceConfigs with different locationalities). In this case, use the
	// resource ID specified by the external reference to distinguish between
	// different ResourceConfigs since different ResourceConfigs will have
	// different ID templates.
	default:
		id := resourceRef.External
		for _, rc := range rcs {
			if idMatchesTemplate(id, rc.IDTemplate) {
				return rc, nil
			}
		}
		return nil, fmt.Errorf("no ResourceConfig found for GroupVersionKind %v with an IDTemplate that matches the id %v", gvk, id)
	}
}

func (t *TFIAMClient) getResource(ctx context.Context, gvk schema.GroupVersionKind, nn types.NamespacedName) (*krmtotf.Resource, error) {
	if nn.Name == "" {
		return t.getResourceForHeadlessKind(ctx, gvk, nn)
	}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	if err := t.kubeClient.Get(ctx, nn, u); err != nil {
		if errors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(gvk, nn)
		}
		return nil, fmt.Errorf("error retrieving resource '%v' with GroupVersionKind '%v': %w", nn, gvk, err)
	}
	sm, err := t.smLoader.GetServiceMapping(gvk.Group)
	if err != nil {
		return nil, err
	}
	resource, err := krmtotf.NewResource(u, sm, t.provider)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling '%v' to resource: %w", nn, err)
	}
	if !k8s.IsResourceReady(&resource.Resource) {
		return nil, k8s.NewReferenceNotReadyErrorForResource(&resource.Resource)
	}
	return resource, nil
}

func (t *TFIAMClient) getResourceForHeadlessKind(ctx context.Context, gvk schema.GroupVersionKind, nn types.NamespacedName) (*krmtotf.Resource, error) {
	switch gvk.Kind {
	case ProjectKind:
		return t.getProjectResource(ctx, nn)
	default:
		return nil, fmt.Errorf("unrecognized IAM kind '%v'", gvk.Kind)
	}
}

func (t *TFIAMClient) getProjectResource(ctx context.Context, nn types.NamespacedName) (*krmtotf.Resource, error) {
	projectID, err := k8s.GetProjectIDForNamespace(ctx, t.kubeClient, nn.Namespace)
	if err != nil {
		return nil, fmt.Errorf("error getting project ID for namespace: %w", err)
	}
	u := unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind": ProjectKind,
			"metadata": map[string]interface{}{
				"name": projectID,
			},
		},
	}
	sm, err := t.smLoader.GetServiceMapping(ResourceManagerGroup)
	if err != nil {
		return nil, fmt.Errorf("error getting service mapping for kind 'Project': %w", err)
	}
	resource, err := krmtotf.NewResource(&u, sm, t.provider)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling '%v' to resource: %w", nn, err)
	}
	return resource, nil
}

func (t *TFIAMClient) resolveMemberReference(ctx context.Context, ref *v1beta1.MemberReference,
	gvk schema.GroupVersionKind, resourceNamespace string) (string, error) {

	nn := types.NamespacedName{
		Namespace: useIfNonEmptyElseDefaultTo(ref.Namespace, resourceNamespace),
		Name:      ref.Name,
	}
	refResource, err := t.getResource(ctx, gvk, nn)
	if err != nil {
		return "", err
	}
	memberRefConfig := refResource.ResourceConfig.IAMMemberReferenceConfig
	val, err := resolveTargetFieldValue(refResource, memberRefConfig.TargetField)
	if err != nil {
		return "", err
	}
	if memberRefConfig.ValueTemplate == "" {
		return val, nil
	}
	return krmtotf.ResolveValueTemplate(memberRefConfig.ValueTemplate, val, refResource, t.kubeClient, t.smLoader)
}

func (t *TFIAMClient) newIAMObjectFromInterface(iamInterface interface{}) metav1.Object {
	switch iamInterface.(type) {
	case *v1beta1.IAMPolicy:
		return &v1beta1.IAMPolicy{
			TypeMeta: metav1.TypeMeta{
				APIVersion: v1beta1.IAMPolicyGVK.GroupVersion().String(),
				Kind:       v1beta1.IAMPolicyGVK.Kind,
			},
		}
	case *v1beta1.IAMPolicyMember:
		return &v1beta1.IAMPolicyMember{
			TypeMeta: metav1.TypeMeta{
				APIVersion: v1beta1.IAMPolicyMemberGVK.GroupVersion().String(),
				Kind:       v1beta1.IAMPolicyMemberGVK.Kind,
			},
		}
	case *v1beta1.IAMAuditConfig:
		return &v1beta1.IAMAuditConfig{
			TypeMeta: metav1.TypeMeta{
				APIVersion: v1beta1.IAMAuditConfigGVK.GroupVersion().String(),
				Kind:       v1beta1.IAMAuditConfigGVK.Kind,
			},
		}
	}
	panic(fmt.Errorf("unknown type: %v", reflect.TypeOf(iamInterface).Name()))
}

func (t *TFIAMClient) newServiceMappingForAssociatedIAMInterface(_ context.Context, iamInterface interface{}, iamConfig corekccv1alpha1.IAMConfig) (*corekccv1alpha1.ServiceMapping, error) {
	switch iamInterface.(type) {
	case *v1beta1.IAMPolicy:
		return newServiceMappingForGVKAndTFResourceName(v1beta1.IAMPolicyGVK, iamConfig.PolicyName), nil
	case *v1beta1.IAMPolicyMember:
		return newServiceMappingForGVKAndTFResourceName(v1beta1.IAMPolicyMemberGVK, iamConfig.PolicyMemberName), nil
	case *v1beta1.IAMAuditConfig:
		return newServiceMappingForGVKAndTFResourceName(v1beta1.IAMAuditConfigGVK, iamConfig.AuditConfigName), nil
	}
	panic(fmt.Errorf("unknown type: %v", reflect.TypeOf(iamInterface).Name()))
}

func getReferenceFieldValue(id string, rc *corekccv1alpha1.ResourceConfig) (string, error) {
	switch rc.IAMConfig.ReferenceField.Type {
	case corekccv1alpha1.IAMReferenceTypeId:
		return id, nil
	case corekccv1alpha1.IAMReferenceTypeName:
		return parseNameFromID(id)
	default:
		panic(fmt.Errorf("unknown value type: %v", rc.IAMConfig.ReferenceField.Type))
	}
}

func resolveTargetFieldValue(r *krmtotf.Resource, targetField string) (string, error) {
	key := text.SnakeCaseToLowerCamelCase(targetField)
	switch key {
	case "":
		panic(fmt.Errorf("empty target field specified"))
	default:
		if val, exists, _ := unstructured.NestedString(r.Spec, strings.Split(key, ".")...); exists {
			return val, nil
		}
		if val, exists, _ := unstructured.NestedString(r.GetStatusOrObservedState(), strings.Split(key, ".")...); exists {
			return val, nil
		}
		return "", fmt.Errorf("couldn't resolve the value for target field %v from the referenced resource %v", targetField, r.GetNamespacedName())
	}
}

func newIAMPolicyFromTFState(resource *krmtotf.Resource, state *terraform.InstanceState, origPolicy *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	resource.Spec, resource.Status = krmtotf.GetSpecAndStatusFromState(resource, state)
	if err := embedPolicyData(resource.Spec); err != nil {
		return nil, err
	}
	u, err := resource.MarshalAsUnstructured()
	if err != nil {
		return nil, fmt.Errorf("error marshalling resource to unstructured: %w", err)
	}
	iamPolicy := v1beta1.IAMPolicy{}
	if err := util.Marshal(u, &iamPolicy); err != nil {
		return nil, fmt.Errorf("error marshalling unstructured to iampolicy: %w", err)
	}
	iamPolicy.Spec.ResourceReference = origPolicy.Spec.ResourceReference
	iamPolicy.ObjectMeta = origPolicy.ObjectMeta
	etag := krmtotf.GetEtagFromState(resource, state)
	if etag == "" {
		return nil, fmt.Errorf("unexpected empty etag read from the IAM Policy resource")
	}
	iamPolicy.Spec.Etag = etag
	return &iamPolicy, nil
}

func newIAMPolicyMemberFromTFState(resource *krmtotf.Resource, state *terraform.InstanceState, origPolicyMember *v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error) {
	resource.Spec, resource.Status = krmtotf.GetSpecAndStatusFromState(resource, state)
	u, err := resource.MarshalAsUnstructured()
	if err != nil {
		return nil, fmt.Errorf("error marshalling resource to unstructured: %w", err)
	}
	iamPolicyMember := v1beta1.IAMPolicyMember{}
	if err := util.Marshal(u, &iamPolicyMember); err != nil {
		return nil, fmt.Errorf("error marshalling unstructured to IAMPolicyMember: %w", err)
	}
	if origPolicyMember.Spec.Member == "" {
		iamPolicyMember.Spec.Member = ""
		iamPolicyMember.Spec.MemberFrom = origPolicyMember.Spec.MemberFrom
	}
	iamPolicyMember.Spec.ResourceReference = origPolicyMember.Spec.ResourceReference
	iamPolicyMember.ObjectMeta = origPolicyMember.ObjectMeta
	return &iamPolicyMember, nil
}

func newIAMAuditConfigFromTFState(resource *krmtotf.Resource, state *terraform.InstanceState, origAuditConfig *v1beta1.IAMAuditConfig) (*v1beta1.IAMAuditConfig, error) {
	resource.Spec, resource.Status = krmtotf.GetSpecAndStatusFromState(resource, state)
	if auditLogConfigs, ok := resource.Spec["auditLogConfig"]; ok {
		resource.Spec["auditLogConfigs"] = auditLogConfigs
		delete(resource.Spec, "auditLogConfig")
	}
	u, err := resource.MarshalAsUnstructured()
	if err != nil {
		return nil, fmt.Errorf("error marshalling resource to unstructured: %w", err)
	}
	iamAuditConfig := v1beta1.IAMAuditConfig{}
	if err := util.Marshal(u, &iamAuditConfig); err != nil {
		return nil, fmt.Errorf("error marshalling unstructured to IAMAuditConfig: %w", err)
	}
	iamAuditConfig.Spec.ResourceReference = origAuditConfig.Spec.ResourceReference
	iamAuditConfig.ObjectMeta = origAuditConfig.ObjectMeta
	return &iamAuditConfig, nil
}

func newServiceMappingForGVKAndTFResourceName(iamGVK schema.GroupVersionKind, tfResourceName string) *corekccv1alpha1.
	ServiceMapping {
	return &corekccv1alpha1.ServiceMapping{
		TypeMeta: metav1.TypeMeta{
			APIVersion: iamGVK.GroupVersion().String(),
			Kind:       reflect.TypeOf(corekccv1alpha1.ServiceMapping{}).Name(),
		},
		ObjectMeta: metav1.ObjectMeta{},
		Spec: corekccv1alpha1.ServiceMappingSpec{
			Name:    iamGVK.Kind,
			Version: iamGVK.Version,
			Resources: []corekccv1alpha1.ResourceConfig{
				{
					Name:       tfResourceName,
					Kind:       iamGVK.Kind,
					SkipImport: true,
					ResourceReferences: []corekccv1alpha1.ReferenceConfig{
						{
							TypeConfig: corekccv1alpha1.TypeConfig{
								Key: "resourceRef",
							},
						},
					},
				},
			},
		},
	}
}
