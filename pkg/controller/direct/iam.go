// Copyright 2024 Google LLC
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

package direct

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	expr "google.golang.org/genproto/googleapis/type/expr"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	LogTypeUnspecified = "LOG_TYPE_UNSPECIFIED"
	AdminRead          = "ADMIN_READ"
	DataWrite          = "DATA_WRITE"
	DataRead           = "DATA_READ"
)

type IAMAdapter interface {
	GetIAMPolicy(ctx context.Context) (*iampb.Policy, error)
	SetIAMPolicy(ctx context.Context, policy *iampb.Policy) (*iampb.Policy, error)
}

// GetIAMPolicyMember returns the actual IAMPolicyMember for the specified member and referenced resource.
func GetIAMPolicyMember(ctx context.Context, reader client.Reader, want *v1beta1.IAMPolicyMember, memberID v1beta1.Member) (*v1beta1.IAMPolicyMember, error) {
	adapter, err := registry.AdapterForReference(ctx, reader, want.GetNamespace(), want.Spec.ResourceReference)
	if err != nil {
		return nil, fmt.Errorf("building adapter: %w", err)
	}
	iamAdapter, ok := adapter.(IAMAdapter)
	if !ok {
		return nil, fmt.Errorf("adapter does not implement IAMAdapter")
	}

	policy, err := iamAdapter.GetIAMPolicy(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting IAM policy: %w", err)
	}

	actual := &v1beta1.IAMPolicyMember{}
	actual.ObjectMeta = want.ObjectMeta
	actual.Spec = v1beta1.IAMPolicyMemberSpec{
		ResourceReference: want.Spec.ResourceReference,
	}

	actual.Spec.Member = memberID

	for _, binding := range policy.Bindings {
		if binding.Role != want.Spec.Role {
			continue
		}
		for _, member := range binding.Members {
			if memberMatches(member, string(memberID)) {
				actual.Spec.Role = want.Spec.Role
			}
		}
	}
	return actual, nil
}

// SetIAMPolicyMember will update the IAM policy for the specified member
func SetIAMPolicyMember(ctx context.Context, reader client.Reader, want *v1beta1.IAMPolicyMember, memberID v1beta1.Member) (*v1beta1.IAMPolicyMember, error) {
	adapter, err := registry.AdapterForReference(ctx, reader, want.GetNamespace(), want.Spec.ResourceReference)
	if err != nil {
		return nil, fmt.Errorf("building adapter: %w", err)
	}
	iamAdapter, ok := adapter.(IAMAdapter)
	if !ok {
		return nil, fmt.Errorf("adapter does not implement IAMAdapter")
	}

	policy, err := iamAdapter.GetIAMPolicy(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting IAM policy: %w", err)
	}

	var binding *iampb.Binding
	for _, b := range policy.Bindings {
		if b.Role != want.Spec.Role {
			continue
		}
		binding = b
	}

	if binding == nil {
		binding = &iampb.Binding{
			Role: want.Spec.Role,
		}
		policy.Bindings = append(policy.Bindings, binding)
	}

	hasMember := false
	for _, member := range binding.Members {
		if memberMatches(member, string(memberID)) {
			hasMember = true
		}
	}
	latest := policy
	if !hasMember {
		binding.Members = append(binding.Members, string(memberID))
		newPolicy, err := iamAdapter.SetIAMPolicy(ctx, policy)
		if err != nil {
			return nil, fmt.Errorf("setting IAM policy: %w", err)
		}
		latest = newPolicy
	}

	actual := &v1beta1.IAMPolicyMember{}
	actual.ObjectMeta = want.ObjectMeta
	actual.Spec = v1beta1.IAMPolicyMemberSpec{
		ResourceReference: want.Spec.ResourceReference,
	}

	actual.Spec.Member = memberID

	for _, binding := range latest.Bindings {
		if binding.Role != want.Spec.Role {
			continue
		}
		for _, member := range binding.Members {
			if memberMatches(member, string(memberID)) {
				actual.Spec.Role = want.Spec.Role
			}
		}
	}
	return actual, nil
}

// DeleteIAMPolicyMember will remove the specified member for the IAM policy for a resource
func DeleteIAMPolicyMember(ctx context.Context, reader client.Reader, want *v1beta1.IAMPolicyMember, removeMember v1beta1.Member) error {
	log := klog.FromContext(ctx)

	adapter, err := registry.AdapterForReference(ctx, reader, want.GetNamespace(), want.Spec.ResourceReference)
	if err != nil {
		return fmt.Errorf("building adapter: %w", err)
	}
	iamAdapter, ok := adapter.(IAMAdapter)
	if !ok {
		return fmt.Errorf("adapter does not implement IAMAdapter")
	}

	policy, err := iamAdapter.GetIAMPolicy(ctx)
	if err != nil {
		return fmt.Errorf("getting IAM policy: %w", err)
	}

	var binding *iampb.Binding
	for _, b := range policy.Bindings {
		if b.Role != want.Spec.Role {
			continue
		}
		binding = b
	}

	if binding == nil {
		return nil
	}

	var newMembers []string
	removedMember := false
	for _, member := range binding.Members {
		if memberMatches(member, string(removeMember)) {
			removedMember = true
			continue
		}
		newMembers = append(newMembers, member)
	}
	binding.Members = newMembers

	if !removedMember {
		return nil
	}
	newPolicy, err := iamAdapter.SetIAMPolicy(ctx, policy)
	if err != nil {
		return fmt.Errorf("setting IAM policy: %w", err)
	}

	log.Info("updated iam policy to remove member", "updatedPolicy", newPolicy, "member", removeMember)
	return nil
}

func memberMatches(member, wantMember string) bool {
	return strings.EqualFold(stripTombstone(member), stripTombstone(wantMember))
}

func stripTombstone(member string) string {
	if strings.HasPrefix(strings.ToLower(member), "deleted:") {
		member = member[len("deleted:"):]
	}
	if i := strings.Index(member, "?"); i != -1 {
		member = member[:i]
	}
	return member
}

// GetIAMPolicy returns the actual IAMPolicy for the referenced resource.
func GetIAMPolicy(ctx context.Context, reader client.Reader, want *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	adapter, err := registry.AdapterForReference(ctx, reader, want.GetNamespace(), want.Spec.ResourceReference)
	if err != nil {
		return nil, fmt.Errorf("building adapter: %w", err)
	}
	iamAdapter, ok := adapter.(IAMAdapter)
	if !ok {
		return nil, fmt.Errorf("adapter does not implement IAMAdapter")
	}

	policy, err := iamAdapter.GetIAMPolicy(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting IAM policy: %w", err)
	}

	actual := &v1beta1.IAMPolicy{}
	actual.ObjectMeta = want.ObjectMeta
	actual.Spec.ResourceReference = want.Spec.ResourceReference
	actual.Spec.Etag = string(policy.Etag)
	actual.Spec.Bindings = IAMBindingsFromProto(policy.Bindings)
	actual.Spec.AuditConfigs = IAMAuditConfigsFromProto(policy.AuditConfigs)

	return actual, nil
}

// SetIAMPolicy will update the IAM policy for the referenced resource.
func SetIAMPolicy(ctx context.Context, reader client.Reader, want *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	log := klog.FromContext(ctx)

	adapter, err := registry.AdapterForReference(ctx, reader, want.GetNamespace(), want.Spec.ResourceReference)
	if err != nil {
		return nil, fmt.Errorf("building adapter: %w", err)
	}
	iamAdapter, ok := adapter.(IAMAdapter)
	if !ok {
		return nil, fmt.Errorf("adapter does not implement IAMAdapter")
	}

	livePolicy, err := iamAdapter.GetIAMPolicy(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting IAM policy: %w", err)
	}

	actual := &v1beta1.IAMPolicy{}
	actual.ObjectMeta = want.ObjectMeta
	actual.Spec.ResourceReference = want.Spec.ResourceReference
	actual.Spec.Etag = string(livePolicy.Etag)
	actual.Spec.Bindings = IAMBindingsFromProto(livePolicy.Bindings)
	actual.Spec.AuditConfigs = IAMAuditConfigsFromProto(livePolicy.AuditConfigs)

	diff := v1beta1.IAMPolicySpecDiffers(&want.Spec, &actual.Spec)
	if !diff.HasDiff() {
		log.Info("underlying resource is already up to date", "resource", k8s.GetNamespacedName(want))
		return actual, nil
	}

	newPolicy := &iampb.Policy{
		Version:      3,
		Etag:         livePolicy.Etag,
		Bindings:     IAMBindingsToProto(want.Spec.Bindings),
		AuditConfigs: IAMAuditConfigsToProto(want.Spec.AuditConfigs),
	}
	if want.Spec.Etag != "" {
		newPolicy.Etag = []byte(want.Spec.Etag)
	}

	setResult, err := iamAdapter.SetIAMPolicy(ctx, newPolicy)
	if err != nil {
		return nil, fmt.Errorf("setting IAM policy: %w", err)
	}

	res := &v1beta1.IAMPolicy{}
	res.ObjectMeta = want.ObjectMeta
	res.Spec.ResourceReference = want.Spec.ResourceReference
	res.Spec.Etag = string(setResult.Etag)
	res.Spec.Bindings = IAMBindingsFromProto(setResult.Bindings)
	res.Spec.AuditConfigs = IAMAuditConfigsFromProto(setResult.AuditConfigs)

	return res, nil
}

// DeleteIAMPolicy will remove the IAM policy for the referenced resource.
func DeleteIAMPolicy(ctx context.Context, reader client.Reader, want *v1beta1.IAMPolicy) error {
	log := klog.FromContext(ctx)

	adapter, err := registry.AdapterForReference(ctx, reader, want.GetNamespace(), want.Spec.ResourceReference)
	if err != nil {
		return fmt.Errorf("building adapter: %w", err)
	}
	iamAdapter, ok := adapter.(IAMAdapter)
	if !ok {
		return fmt.Errorf("adapter does not implement IAMAdapter")
	}

	policy, err := iamAdapter.GetIAMPolicy(ctx)
	if err != nil {
		return fmt.Errorf("getting IAM policy: %w", err)
	}

	if len(policy.Bindings) == 0 && len(policy.AuditConfigs) == 0 {
		return nil
	}

	policy.Bindings = nil
	if len(want.Spec.AuditConfigs) > 0 {
		policy.AuditConfigs = nil
	}

	newPolicy, err := iamAdapter.SetIAMPolicy(ctx, policy)
	if err != nil {
		return fmt.Errorf("setting IAM policy: %w", err)
	}

	log.Info("updated iam policy to empty for deletion", "updatedPolicy", newPolicy)
	return nil
}

func IAMBindingsToProto(bindings []v1beta1.IAMPolicyBinding) []*iampb.Binding {
	if len(bindings) == 0 {
		return nil
	}
	protoBindings := make([]*iampb.Binding, 0, len(bindings))
	for _, b := range bindings {
		pbBinding := &iampb.Binding{
			Role:    b.Role,
			Members: make([]string, len(b.Members)),
		}
		for i, member := range b.Members {
			pbBinding.Members[i] = string(member)
		}
		if b.Condition != nil {
			pbBinding.Condition = &expr.Expr{
				Expression:  b.Condition.Expression,
				Title:       b.Condition.Title,
				Description: b.Condition.Description,
			}
		}
		protoBindings = append(protoBindings, pbBinding)
	}
	return protoBindings
}

func IAMBindingsFromProto(bindings []*iampb.Binding) []v1beta1.IAMPolicyBinding {
	if len(bindings) == 0 {
		return nil
	}
	out := make([]v1beta1.IAMPolicyBinding, 0, len(bindings))
	for _, pbBinding := range bindings {
		binding := v1beta1.IAMPolicyBinding{
			Role:    pbBinding.Role,
			Members: make([]v1beta1.Member, len(pbBinding.Members)),
		}
		for i, member := range pbBinding.Members {
			binding.Members[i] = v1beta1.Member(member)
		}
		if pbBinding.Condition != nil {
			binding.Condition = &v1beta1.IAMCondition{
				Expression:  pbBinding.Condition.Expression,
				Title:       pbBinding.Condition.Title,
				Description: pbBinding.Condition.Description,
			}
		}
		out = append(out, binding)
	}
	return out
}

func IAMAuditConfigsToProto(auditConfigs []v1beta1.IAMPolicyAuditConfig) []*iampb.AuditConfig {
	if len(auditConfigs) == 0 {
		return nil
	}
	protoAuditConfigs := make([]*iampb.AuditConfig, 0, len(auditConfigs))
	for _, ac := range auditConfigs {
		pbAuditConfig := &iampb.AuditConfig{
			Service: ac.Service,
		}
		if len(ac.AuditLogConfigs) > 0 {
			pbAuditConfig.AuditLogConfigs = make([]*iampb.AuditLogConfig, 0, len(ac.AuditLogConfigs))
			for _, alc := range ac.AuditLogConfigs {
				pbAlc := &iampb.AuditLogConfig{
					LogType: mapV1Beta1LogTypeToProto(alc.LogType),
				}
				if len(alc.ExemptedMembers) > 0 {
					pbAlc.ExemptedMembers = make([]string, len(alc.ExemptedMembers))
					for i, em := range alc.ExemptedMembers {
						pbAlc.ExemptedMembers[i] = string(em)
					}
				}
				pbAuditConfig.AuditLogConfigs = append(pbAuditConfig.AuditLogConfigs, pbAlc)
			}
		}
		protoAuditConfigs = append(protoAuditConfigs, pbAuditConfig)
	}
	return protoAuditConfigs
}

func IAMAuditConfigsFromProto(auditConfigs []*iampb.AuditConfig) []v1beta1.IAMPolicyAuditConfig {
	if len(auditConfigs) == 0 {
		return nil
	}
	out := make([]v1beta1.IAMPolicyAuditConfig, 0, len(auditConfigs))
	for _, pbAuditConfig := range auditConfigs {
		ac := v1beta1.IAMPolicyAuditConfig{
			Service: pbAuditConfig.Service,
		}
		if len(pbAuditConfig.AuditLogConfigs) > 0 {
			ac.AuditLogConfigs = make([]v1beta1.AuditLogConfig, 0, len(pbAuditConfig.AuditLogConfigs))
			for _, pbAlc := range pbAuditConfig.AuditLogConfigs {
				alc := v1beta1.AuditLogConfig{
					LogType: mapProtoLogTypeToKRM(pbAlc.LogType),
				}
				if len(pbAlc.ExemptedMembers) > 0 {
					alc.ExemptedMembers = make([]v1beta1.Member, len(pbAlc.ExemptedMembers))
					for i, em := range pbAlc.ExemptedMembers {
						alc.ExemptedMembers[i] = v1beta1.Member(em)
					}
				}
				ac.AuditLogConfigs = append(ac.AuditLogConfigs, alc)
			}
		}
		out = append(out, ac)
	}
	return out
}

func mapV1Beta1LogTypeToProto(logTypeString string) iampb.AuditLogConfig_LogType {
	switch strings.ToUpper(logTypeString) {
	case LogTypeUnspecified:
		return iampb.AuditLogConfig_LOG_TYPE_UNSPECIFIED
	case AdminRead:
		return iampb.AuditLogConfig_ADMIN_READ
	case DataWrite:
		return iampb.AuditLogConfig_DATA_WRITE
	case DataRead:
		return iampb.AuditLogConfig_DATA_READ
	default:
		return iampb.AuditLogConfig_LOG_TYPE_UNSPECIFIED
	}
}

func mapProtoLogTypeToKRM(logTypeEnum iampb.AuditLogConfig_LogType) string {
	switch logTypeEnum {
	case iampb.AuditLogConfig_LOG_TYPE_UNSPECIFIED:
		return LogTypeUnspecified
	case iampb.AuditLogConfig_ADMIN_READ:
		return AdminRead
	case iampb.AuditLogConfig_DATA_WRITE:
		return DataWrite
	case iampb.AuditLogConfig_DATA_READ:
		return DataRead
	default:
		return LogTypeUnspecified
	}
}

