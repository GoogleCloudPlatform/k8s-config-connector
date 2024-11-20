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

	bigqueryconnection "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryconnection/v1beta1"
	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (c *IAMClient) isDCLBasedIAMResource(iamInterface interface{}) bool {
	_, resourceRef := extractNamespaceAndResourceReference(iamInterface)
	return c.isDCLBasedResource(resourceRef.GroupVersionKind())
}

func (c *IAMClient) isDCLBasedResource(gvk schema.GroupVersionKind) bool {
	return metadata.IsDCLBasedResourceKind(gvk, c.DCLIAMClient.converter.MetadataLoader)
}

// ResolveMemberIdentity checks only one of Member/MemberFrom is provided, and then tries to resolve identity.
// MemberFrom can only have oneOf a ServiceAccountRef, a LogSinkRef, a SQLInstanceRef, so to resolve these
// values, it is necessary to call on the TFIAMClient
func ResolveMemberIdentity(ctx context.Context, member v1beta1.Member,
	memberFrom *v1beta1.MemberSource, namespace string, tfIAMClient *TFIAMClient) (id string, err error) {
	if member != "" && memberFrom != nil {
		return id, fmt.Errorf("both 'member' and 'memberFrom' are used. Exactly one of them must be used")
	}

	if member == "" && memberFrom == nil {
		return id, fmt.Errorf("both 'member' and 'memberFrom' are empty. Exactly one of them must be used")
	}

	if member != "" {
		return string(member), nil
	}

	if err := memberFrom.Validate(); err != nil {
		return id, err
	}

	if id, err := tryResolveTFMemberReference(ctx, memberFrom, namespace, tfIAMClient); err != nil {
		return "", err
	} else if id != "" { // successfully resolved member reference from TF-based resources
		return id, nil
	}

	if id, err = tryResolveDirectMemberReference(ctx, memberFrom, namespace, tfIAMClient.kubeClient); err != nil {
		return "", err
	} else if id != "" { // successfully resolved member reference from direct resources
		return id, nil
	}

	return "", fmt.Errorf("unable to resolve member identity from %v", memberFrom)
}

func tryResolveTFMemberReference(ctx context.Context, memberFrom *v1beta1.MemberSource, namespace string, tfIAMClient *TFIAMClient) (string, error) {
	var refs []*v1beta1.MemberReference
	var gvks []schema.GroupVersionKind

	if memberFrom.ServiceAccountRef != nil {
		refs = append(refs, memberFrom.ServiceAccountRef)
		gvks = append(gvks, IAMServiceAccountGVK)
	}

	if memberFrom.LogSinkRef != nil {
		refs = append(refs, memberFrom.LogSinkRef)
		gvks = append(gvks, LoggingLogSinkGVK)
	}

	if memberFrom.SQLInstanceRef != nil {
		refs = append(refs, memberFrom.SQLInstanceRef)
		gvks = append(gvks, SQLInstanceGVK)
	}

	if memberFrom.ServiceIdentityRef != nil {
		refs = append(refs, memberFrom.ServiceIdentityRef)
		gvks = append(gvks, ServiceIdentityGVK)
	}

	if len(refs) == 1 {
		return tfIAMClient.resolveMemberReference(ctx, refs[0], gvks[0], namespace)
	}
	return "", nil
}

func tryResolveDirectMemberReference(ctx context.Context, memberFrom *v1beta1.MemberSource, namespace string, reader client.Reader) (string, error) {
	if memberFrom.BigQueryConnectionConnectionRef != nil {
		return bigqueryconnection.ResolveServiceAccountID(ctx, reader, namespace, memberFrom.BigQueryConnectionConnectionRef)
	}
	// TODO: handle more direct resource reference
	return "", nil
}

func extractNamespaceAndResourceReference(iamInterface interface{}) (string, v1beta1.ResourceReference) {
	switch iamObject := iamInterface.(type) {
	case *v1beta1.IAMPolicy:
		return iamObject.Namespace, iamObject.Spec.ResourceReference
	case *v1beta1.IAMPolicyMember:
		return iamObject.Namespace, iamObject.Spec.ResourceReference
	case *v1beta1.IAMAuditConfig:
		return iamObject.Namespace, iamObject.Spec.ResourceReference
	}
	panic(fmt.Errorf("unknown type: %v", reflect.TypeOf(iamInterface).Name()))
}

func embedPolicyData(spec map[string]interface{}) error {
	policyData, ok := spec["policyData"].(string)
	if !ok {
		return nil
	}
	delete(spec, "policyData")
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(policyData), &m); err != nil {
		return fmt.Errorf("error converting policyData '%v' to map: %w", policyData, err)
	}
	for k, v := range m {
		spec[k] = v
	}
	return nil
}

// An unfortunate reality is that the GVK is not always properly filled in when
// reading a resource from the K8s API server, and there are functions that
// need the Kind to be filled to work (e.g. krmtotf.NewResource,
// k8s.MarshalAsUnstructured, etc.). The Kind is not set because the TypeMeta
// is empty. The reason why the TypeMeta is empty is because in
// k8s.io/apimachinery/pkg/runtime/serializer/versioning/versioning.go the
// GVK is cleared inside of Decode(...)
func SetGVK(iamInterface interface{}) {
	switch iamObject := iamInterface.(type) {
	case *v1beta1.IAMPolicy:
		setPolicyGVK(iamObject)
	case *v1beta1.IAMPartialPolicy:
		setPartialPolicyGVK(iamObject)
	case *v1beta1.IAMPolicyMember:
		setPolicyMemberGVK(iamObject)
	case *v1beta1.IAMAuditConfig:
		setAuditConfigGVK(iamObject)
	default:
		panic(fmt.Errorf("unknown type: %v", reflect.TypeOf(iamInterface).Name()))
	}
}

func setPolicyGVK(policy *v1beta1.IAMPolicy) {
	policy.SetGroupVersionKind(v1beta1.IAMPolicyGVK)
}

func setPartialPolicyGVK(partialPolicy *v1beta1.IAMPartialPolicy) {
	partialPolicy.SetGroupVersionKind(v1beta1.IAMPartialPolicyGVK)
}

func setPolicyMemberGVK(policyMember *v1beta1.IAMPolicyMember) {
	policyMember.SetGroupVersionKind(v1beta1.IAMPolicyMemberGVK)
}

func setAuditConfigGVK(auditConfig *v1beta1.IAMAuditConfig) {
	auditConfig.SetGroupVersionKind(v1beta1.IAMAuditConfigGVK)
}

func resourceSupportsIAMPolicy(rc *corekccv1alpha1.ResourceConfig) bool {
	return rc.IAMConfig.PolicyName != ""
}

func resourceSupportsIAMPolicyMember(rc *corekccv1alpha1.ResourceConfig) bool {
	return rc.IAMConfig.PolicyMemberName != ""
}

func resourceSupportsIAMAuditConfigs(rc *corekccv1alpha1.ResourceConfig) bool {
	return rc.IAMConfig.AuditConfigName != ""
}

func useIfNonEmptyElseDefaultTo(str, backup string) string {
	if str != "" {
		return str
	}
	return backup
}

func parseNameFromID(id string) (string, error) {
	if strings.TrimSpace(id) == "" {
		return "", fmt.Errorf("error parsing name from id: id is empty")
	}
	parts := strings.Split(id, "/")
	return parts[len(parts)-1], nil
}

func idMatchesTemplate(id, idTemplate string) bool {
	idTokens := strings.Split(id, "/")
	idTemplateTokens := strings.Split(idTemplate, "/")
	if len(idTokens) != len(idTemplateTokens) {
		return false
	}
	for i := range idTokens {
		if idTokens[i] != idTemplateTokens[i] &&
			!idTemplateVarsRegex.MatchString(idTemplateTokens[i]) {
			return false
		}
	}
	return true
}
