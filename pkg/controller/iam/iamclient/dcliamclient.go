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
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/kcclite"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/externalonlygvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	dcliam "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	dclunstructiam "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DCLIAMClient struct {
	dclClient  *dcliam.Client
	kubeClient client.Client
	converter  *conversion.Converter
	smLoader   *servicemappingloader.ServiceMappingLoader
}

func (d *DCLIAMClient) SetPolicyMember(ctx context.Context, tfIAMClient *TFIAMClient, policyMember *v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error) {
	dclResource, nn, err := d.getDCLResourceAndNamespacedNameFromPolicyMember(ctx, policyMember)
	if err != nil {
		return nil, err
	}

	dclPolicyMember, err := newDCLPolicyMemberFromKRMPolicyMember(ctx, policyMember, dclResource, nn.Namespace, tfIAMClient)
	if err != nil {
		return nil, fmt.Errorf("error converting DCL Policy Member from KRM Policy Member: %w", err)
	}
	dclPolicyMemberResource := dclunstructiam.MemberToUnstructured(dclPolicyMember)
	// DCL's SetPolicyMember returns a Policy not a Member, which is not what this function should return
	_, err = dclunstruct.SetPolicyMember(ctx, d.dclClient.Config, dclResource, dclPolicyMemberResource)
	if err != nil {
		return nil, fmt.Errorf("error setting IAMPolicyMember for resource %v: %w", nn, err)
	}
	return policyMember, nil
}

func (d *DCLIAMClient) GetPolicyMember(ctx context.Context, tfIAMClient *TFIAMClient, policyMember *v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error) {
	dclResource, nn, err := d.getDCLResourceAndNamespacedNameFromPolicyMember(ctx, policyMember)
	if err != nil {
		return nil, err
	}
	id, err := ResolveMemberIdentity(ctx, policyMember.Spec.Member, policyMember.Spec.MemberFrom, nn.Namespace, tfIAMClient)
	if err != nil {
		return nil, err
	}
	dclPolicyResource, err := dclunstruct.GetPolicyMember(ctx, d.dclClient.Config,
		dclResource, policyMember.Spec.Role, id)
	if err != nil {
		return nil, fmt.Errorf("error getting IAMPolicyMember for resource %v: %w", nn, err)
	}
	return newIAMPolicyMemberFromDCLResource(dclPolicyResource, policyMember)
}

func (d *DCLIAMClient) DeletePolicyMember(ctx context.Context, tfIAMClient *TFIAMClient, policyMember *v1beta1.IAMPolicyMember) error {
	dclResource, nn, err := d.getDCLResourceAndNamespacedNameFromPolicyMember(ctx, policyMember)
	if err != nil {
		return err
	}

	dclPolicyMember, err := newDCLPolicyMemberFromKRMPolicyMember(ctx, policyMember, dclResource, nn.Namespace, tfIAMClient)
	if err != nil {
		return fmt.Errorf("error converting DCL Policy Member from KRM Policy Member: %w", err)
	}
	dclPolicyMemberResource := dclunstructiam.MemberToUnstructured(dclPolicyMember)

	err = dclunstruct.DeletePolicyMember(ctx, d.dclClient.Config, dclResource, dclPolicyMemberResource)
	if err != nil {
		return fmt.Errorf("error deleting IAMPolicyMember for resource: %w", err)
	}
	return nil
}

func (d *DCLIAMClient) SetPolicy(ctx context.Context, policy *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	namespace, resourceRef := extractNamespaceAndResourceReference(policy)
	// TODO: Support the handling of external field for policy.
	nn := types.NamespacedName{
		Namespace: useIfNonEmptyElseDefaultTo(resourceRef.Namespace, namespace),
		Name:      resourceRef.Name,
	}
	dclSchema, err := d.getSchemaFromResourceReference(resourceRef)
	if err != nil {
		return nil, err
	}

	// Check if the resource supports IAM on DCL.
	supportsIAM, err := extension.HasIam(dclSchema)
	if err != nil {
		return nil, err
	}
	if !supportsIAM {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM policies", resourceRef.Kind)
	}

	dclResource, err := d.getDCLResource(ctx, resourceRef, dclSchema, namespace)
	if err != nil {
		return nil, fmt.Errorf("error getting referenced DCL resource, with reference %v: %w", nn, err)
	}

	dclPolicy, err := newDCLPolicyFromKRMPolicy(policy, dclResource)
	if err != nil {
		return nil, fmt.Errorf("error converting DCLPolicy from KRM Policy: %w", err)
	}
	dclPolicyResource := dclunstructiam.PolicyToUnstructured(dclPolicy)
	liveDCLResource, err := dclunstruct.GetPolicy(ctx, d.dclClient.Config, dclResource)
	if err != nil {
		return nil, fmt.Errorf("error getting IAMPolicy for resource %v: %w", nn, err)
	}
	// Check if resource's live state matches the desired state.
	if reflect.DeepEqual(dclPolicyResource, liveDCLResource) {
		logger.Info("underlying resource is already up to date", "resource", k8s.GetNamespacedName(policy))
		return newIAMPolicyFromDCLResource(liveDCLResource, policy)
	}
	dclPolicyResource, err = dclunstruct.SetPolicyWithEtag(ctx, d.dclClient.Config, dclResource, dclPolicyResource)
	if err != nil {
		return nil, fmt.Errorf("error setting IAMPolicy for resource %v: %w", nn, err)
	}

	return newIAMPolicyFromDCLResource(dclPolicyResource, policy)
}

func (d *DCLIAMClient) GetPolicy(ctx context.Context, policy *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	namespace, resourceRef := extractNamespaceAndResourceReference(policy)
	// TODO: Support the handling of external field for policy.
	nn := types.NamespacedName{
		Namespace: useIfNonEmptyElseDefaultTo(resourceRef.Namespace, namespace),
		Name:      resourceRef.Name,
	}
	dclSchema, err := d.getSchemaFromResourceReference(resourceRef)
	if err != nil {
		return nil, err
	}

	// Check if the resource supports IAM on DCL.
	supportsIAM, err := extension.HasIam(dclSchema)
	if err != nil {
		return nil, err
	}
	if !supportsIAM {
		return nil, fmt.Errorf("invalid resource reference: kind %v does not support IAM policies", resourceRef.Kind)
	}

	dclResource, err := d.getDCLResource(ctx, resourceRef, dclSchema, namespace)
	if err != nil {
		return nil, fmt.Errorf("error getting referenced DCL resource, with reference %v: %w", nn, err)
	}

	dclPolicyResource, err := dclunstruct.GetPolicy(ctx, d.dclClient.Config, dclResource)
	if err != nil {
		return nil, fmt.Errorf("error getting IAMPolicy for resource %v: %w", nn, err)
	}
	return newIAMPolicyFromDCLResource(dclPolicyResource, policy)
}

func (d *DCLIAMClient) DeletePolicy(ctx context.Context, policy *v1beta1.IAMPolicy) error {
	namespace, resourceRef := extractNamespaceAndResourceReference(policy)
	// TODO: Support the handling of external field for policy.
	nn := types.NamespacedName{
		Namespace: useIfNonEmptyElseDefaultTo(resourceRef.Namespace, namespace),
		Name:      resourceRef.Name,
	}
	dclSchema, err := d.getSchemaFromResourceReference(resourceRef)
	if err != nil {
		return err
	}

	// Check if the resource supports IAM on DCL.
	supportsIAM, err := extension.HasIam(dclSchema)
	if err != nil {
		return err
	}
	if !supportsIAM {
		return fmt.Errorf("invalid resource reference: kind %v does not support IAM policies", resourceRef.Kind)
	}

	dclResource, err := d.getDCLResource(ctx, resourceRef, dclSchema, namespace)
	if err != nil {
		return fmt.Errorf("error getting referenced DCL resource, with reference %v: %w", nn, err)
	}

	dclPolicyResource, err := dclunstruct.GetPolicy(ctx, d.dclClient.Config, dclResource)
	if err != nil {
		return fmt.Errorf("error getting IAMPolicy for resource%v: %w", nn, err)
	}

	// Setting an empty policy is technically deleting the policy as DCL only
	// mirrors the get/set calls from the API for now.
	emptyPolicy := setEmptyPolicyForDeletion(dclPolicyResource)
	_, err = dclunstruct.SetPolicy(ctx, d.dclClient.Config, dclResource, emptyPolicy)
	if err != nil {
		return fmt.Errorf("error deleting IAMPolicy for resource: %w", err)
	}
	return nil
}

func (d *DCLIAMClient) getSchemaFromResourceReference(resourceRef v1beta1.ResourceReference) (*openapi.Schema, error) {
	gvk := resourceRef.GroupVersionKind()
	// ExternalOnlyGVKs are only supported by TF.
	if externalonlygvks.IsExternalOnlyGVK(gvk) {
		return nil, fmt.Errorf("invalid DCL resource reference type: kind %v is not supported", resourceRef.Kind)
	}

	dclSchema, err := dclschemaloader.GetDCLSchemaForGVK(gvk,
		d.converter.MetadataLoader, d.converter.SchemaLoader)
	if err != nil {
		return nil, err
	}
	return dclSchema, nil
}

func (d *DCLIAMClient) getDCLResourceAndNamespacedNameFromPolicyMember(ctx context.Context, policyMember *v1beta1.IAMPolicyMember) (*dclunstruct.Resource, types.NamespacedName, error) {
	namespace, resourceRef := extractNamespaceAndResourceReference(policyMember)
	// TODO: Support the handling of external field for policy member.
	//if resourceRef.External != "" {
	//	return Resource constructed based on external
	//}
	nn := types.NamespacedName{
		Namespace: useIfNonEmptyElseDefaultTo(resourceRef.Namespace, namespace),
		Name:      resourceRef.Name,
	}
	dclSchema, err := d.getSchemaFromResourceReference(resourceRef)
	if err != nil {
		return nil, nn, err
	}

	// Check if the resource supports IAM on DCL.
	supportsIAM, err := extension.HasIam(dclSchema)
	if err != nil {
		return nil, nn, err
	}
	if !supportsIAM {
		return nil, nn, fmt.Errorf("invalid resource reference: kind %v does not support IAM policy member", resourceRef.Kind)
	}

	dclResource, err := d.getDCLResource(ctx, resourceRef, dclSchema, namespace)
	if err != nil {
		return nil, nn, fmt.Errorf("error getting referenced DCL resource, with reference %v: %w", nn, err)
	}
	return dclResource, nn, nil
}

func (d *DCLIAMClient) getDCLResource(ctx context.Context, resourceRef v1beta1.ResourceReference,
	dclSchema *openapi.Schema, namespace string) (*dclunstruct.Resource, error) {
	gvk := resourceRef.GroupVersionKind()
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	nn := types.NamespacedName{
		Namespace: useIfNonEmptyElseDefaultTo(resourceRef.Namespace, namespace),
		Name:      resourceRef.Name,
	}
	if err := d.kubeClient.Get(ctx, nn, u); err != nil {
		if errors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(gvk, nn)
		}
		return nil, fmt.Errorf("error retrieving resource '%v' with GroupVersionKind '%v': %w", nn, gvk, err)
	}
	resource, err := dcl.NewResource(u, dclSchema)
	if err != nil {
		return nil, err
	}
	lite, err := kcclite.ToKCCLiteBestEffort(resource, d.converter.MetadataLoader, d.converter.SchemaLoader, d.smLoader, d.kubeClient)
	if err != nil {
		return nil, fmt.Errorf("error converting KCC full to KCC lite: %w", err)
	}
	dclResource, err := d.converter.KRMObjectToDCLObject(lite)
	if err != nil {
		return nil, err
	}
	return dclResource, nil
}

func setEmptyPolicyForDeletion(dclPolicyResource *dclunstruct.Resource) *dclunstruct.Resource {
	dclPolicyResource.Object["bindings"] = []interface{}{}
	dclPolicyResource.Object["etag"] = ""
	return dclPolicyResource
}

// This function is only used in SetPolicy to convert the newly requested KRMPolicy to a
// DCL policy resource that can interact with DCL's IAMClient.
func newDCLPolicyFromKRMPolicy(policy *v1beta1.IAMPolicy, dclResource *dclunstruct.Resource) (*dcliam.Policy, error) {
	if policy.Spec.AuditConfigs != nil {
		return nil, fmt.Errorf("policy resource contains AuditConfigs which are not currently supported by DCL-based resources")
	}
	return &dcliam.Policy{
		Bindings: kccToDCLBindings(policy.Spec.Bindings),
		Etag:     &policy.Spec.Etag,
		Resource: dclResource,
	}, nil
}

// Converts a DCL IAMPolicy resource into a KRM IAMPolicy resource.
func newIAMPolicyFromDCLResource(dclResource *dclunstruct.Resource, origPolicy *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	if dclResource.STV.Service != "iam" || dclResource.STV.Type != "Policy" {
		return nil, fmt.Errorf("dclResource was not IAMPolicy instead was service %s and type %s", dclResource.STV.Service, dclResource.STV.Type)
	}
	iamPolicy := v1beta1.IAMPolicy{}
	iamPolicy.ObjectMeta = origPolicy.ObjectMeta
	iamPolicy.Spec.ResourceReference = origPolicy.Spec.ResourceReference
	if dclResource.Object["bindings"] != nil {
		binds, err := dclToKCCBindings(dclResource.Object["bindings"].([]interface{}))
		if err != nil {
			return nil, fmt.Errorf("error converting DCL Bindings to KCC Bindings: %w", err)
		}
		iamPolicy.Spec.Bindings = binds
	}
	etag := dclResource.Object["etag"].(string)
	if etag == "" {
		return nil, fmt.Errorf("unexpected empty etag read from the IAMPolicy resource")
	}
	iamPolicy.Spec.Etag = etag
	return &iamPolicy, nil
}

// Converts a DCL IAMPolicyMember resource into a KRM IAMPolicyMember resource.
func newIAMPolicyMemberFromDCLResource(dclMemberResource *dclunstruct.Resource, origPolicyMember *v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error) {
	if dclMemberResource.STV.Service != "iam" || dclMemberResource.STV.Type != "PolicyMember" {
		return nil, fmt.Errorf("dclResource was not IAM Member instead was service %s and type %s", dclMemberResource.STV.Service, dclMemberResource.STV.Type)
	}
	iamPolicyMember := v1beta1.IAMPolicyMember{}
	iamPolicyMember.ObjectMeta = origPolicyMember.ObjectMeta
	iamPolicyMember.Spec = v1beta1.IAMPolicyMemberSpec{
		ResourceReference: origPolicyMember.Spec.ResourceReference,
		Role:              dclMemberResource.Object["role"].(string),
	}
	if dclMemberResource.Object["member"] != nil {
		member := dclMemberResource.Object["member"].(string)
		iamPolicyMember.Spec.Member = v1beta1.Member(member)
	}
	if origPolicyMember.Spec.Member == "" {
		iamPolicyMember.Spec.Member = ""
		iamPolicyMember.Spec.MemberFrom = origPolicyMember.Spec.MemberFrom
	}
	return &iamPolicyMember, nil
}

// This function is only used in SetPolicyMember to convert the newly requested KRMPolicyMember to a
// DCL policy resource that can interact with DCL's IAMClient.
func newDCLPolicyMemberFromKRMPolicyMember(ctx context.Context, policyMember *v1beta1.IAMPolicyMember,
	dclResource *dclunstruct.Resource, namespace string,
	tfIAMClient *TFIAMClient) (*dcliam.Member, error) {
	member, err := ResolveMemberIdentity(ctx, policyMember.Spec.Member, policyMember.Spec.MemberFrom, namespace, tfIAMClient)
	if err != nil {
		return nil, err
	}
	// TODO(b/228226694): DCL-based resources don't currently support Member conditions.
	if policyMember.Spec.Condition != nil {
		return nil, fmt.Errorf("the IAMPolicyMember for this resource of gvk %v does not currently support conditions", policyMember.Spec.ResourceReference.GroupVersionKind())
	}
	return &dcliam.Member{
		Role:     &policyMember.Spec.Role,
		Member:   &member,
		Resource: dclResource,
	}, nil
}

func kccToDCLBindings(kccBindings []v1beta1.IAMPolicyBinding) []dcliam.Binding {
	ret := make([]dcliam.Binding, len(kccBindings))
	for i, bind := range kccBindings {
		role := bind.Role
		ret[i] = dcliam.Binding{
			Role:      &role,
			Members:   convertMembersToStringArr(bind.Members),
			Condition: kccToDCLCondition(bind.Condition),
		}
	}
	return ret
}

func kccToDCLCondition(condition *v1beta1.IAMCondition) *dcliam.Condition {
	if condition == nil {
		return nil
	}
	return &dcliam.Condition{
		Title:       &condition.Title,
		Description: &condition.Description,
		Expression:  &condition.Expression,
	}
}

func dclToKCCBindings(dclBindings []interface{}) ([]v1beta1.IAMPolicyBinding, error) {
	if dclBindings == nil {
		return nil, nil
	}
	ret := make([]v1beta1.IAMPolicyBinding, len(dclBindings))
	for i, bind := range dclBindings {
		binding := bind.(map[string]interface{})
		var cond *v1beta1.IAMCondition
		if binding["condition"] != nil {
			cond = &v1beta1.IAMCondition{}
			if err := util.Marshal(binding["condition"], cond); err != nil {
				return nil, err
			}
		}
		ret[i] = v1beta1.IAMPolicyBinding{
			Members:   convertMapToMemberArr(binding),
			Role:      binding["role"].(string), // no nil check needed because required value
			Condition: cond,
		}
	}
	return ret, nil
}

func convertMapToMemberArr(binding map[string]interface{}) []v1beta1.Member {
	if binding["members"] == nil {
		return nil
	}
	members := binding["members"].([]interface{})
	var ret []v1beta1.Member
	for _, val := range members {
		memberStr := val.(string)
		ret = append(ret, v1beta1.Member(memberStr))
	}
	return ret
}

func convertMembersToStringArr(members []v1beta1.Member) []string {
	var ret []string
	for _, val := range members {
		ret = append(ret, string(val))
	}
	return ret
}
