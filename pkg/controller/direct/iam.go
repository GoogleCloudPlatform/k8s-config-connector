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

	"cloud.google.com/go/iam/apiv1/iampb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	exprpb "google.golang.org/genproto/googleapis/type/expr"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
			if member == string(memberID) {
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
		if member == string(memberID) {
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
			if member == string(memberID) {
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
		if member == string(removeMember) {
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
	actual.Spec = *IAMPolicySpec_FromProto(policy)

	return actual, nil
}

// SetIAMPolicy will update the IAM policy for the referenced resource
func SetIAMPolicy(ctx context.Context, reader client.Reader, want *v1beta1.IAMPolicy) (*v1beta1.IAMPolicy, error) {
	adapter, err := registry.AdapterForReference(ctx, reader, want.GetNamespace(), want.Spec.ResourceReference)
	if err != nil {
		return nil, fmt.Errorf("building adapter: %w", err)
	}
	iamAdapter, ok := adapter.(IAMAdapter)
	if !ok {
		return nil, fmt.Errorf("adapter does not implement IAMAdapter")
	}

	protoPolicy := IAMPolicySpec_ToProto(&want.Spec)
	newPolicy, err := iamAdapter.SetIAMPolicy(ctx, protoPolicy)
	if err != nil {
		return nil, fmt.Errorf("setting IAM policy: %w", err)
	}

	actual := &v1beta1.IAMPolicy{}
	actual.ObjectMeta = want.ObjectMeta
	actual.Spec = *IAMPolicySpec_FromProto(newPolicy)

	return actual, nil
}

func IAMPolicySpec_ToProto(in *v1beta1.IAMPolicySpec) *iampb.Policy {
	if in == nil {
		return nil
	}

	protoPolicy := &iampb.Policy{
		Version: 3,
	}

	// Map Bindings
	if len(in.Bindings) > 0 {
		protoPolicy.Bindings = make([]*iampb.Binding, 0, len(in.Bindings))
		for _, b := range in.Bindings {
			pbBinding := &iampb.Binding{
				Role:    b.Role,
				Members: make([]string, len(b.Members)),
			}
			for i, member := range b.Members {
				pbBinding.Members[i] = string(member)
			}

			if b.Condition != nil {
				pbBinding.Condition = &exprpb.Expr{
					Expression:  b.Condition.Expression,
					Title:       b.Condition.Title,
					Description: b.Condition.Description,
				}
			}
			protoPolicy.Bindings = append(protoPolicy.Bindings, pbBinding)
		}
	}

	// Map AuditConfigs
	if len(in.AuditConfigs) > 0 {
		protoPolicy.AuditConfigs = make([]*iampb.AuditConfig, 0, len(in.AuditConfigs))
		for _, ac := range in.AuditConfigs {
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
			protoPolicy.AuditConfigs = append(protoPolicy.AuditConfigs, pbAuditConfig)
		}
	}

	return protoPolicy
}

func IAMPolicySpec_FromProto(in *iampb.Policy) *v1beta1.IAMPolicySpec {
	if in == nil {
		return nil
	}

	out := &v1beta1.IAMPolicySpec{}

	// Map Bindings from Proto to KRM
	if len(in.Bindings) > 0 {
		out.Bindings = make([]v1beta1.IAMPolicyBinding, 0, len(in.Bindings))
		for _, pbBinding := range in.Bindings {
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
			out.Bindings = append(out.Bindings, binding)
		}
	}

	// Map AuditConfigs from Proto to KRM
	if len(in.AuditConfigs) > 0 {
		out.AuditConfigs = make([]v1beta1.IAMPolicyAuditConfig, 0, len(in.AuditConfigs))
		for _, pbAuditConfig := range in.AuditConfigs {
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
			out.AuditConfigs = append(out.AuditConfigs, ac)
		}
	}

	return out
}

func mapV1Beta1LogTypeToProto(logTypeString string) iampb.AuditLogConfig_LogType {
	switch logTypeString {
	case "LOG_TYPE_UNSPECIFIED":
		return iampb.AuditLogConfig_LOG_TYPE_UNSPECIFIED
	case "ADMIN_READ":
		return iampb.AuditLogConfig_ADMIN_READ
	case "DATA_WRITE":
		return iampb.AuditLogConfig_DATA_WRITE
	case "DATA_READ":
		return iampb.AuditLogConfig_DATA_READ
	default:
		return iampb.AuditLogConfig_LOG_TYPE_UNSPECIFIED
	}
}

func mapProtoLogTypeToKRM(logTypeEnum iampb.AuditLogConfig_LogType) string {
	switch logTypeEnum {
	case iampb.AuditLogConfig_LOG_TYPE_UNSPECIFIED:
		return "LOG_TYPE_UNSPECIFIED"
	case iampb.AuditLogConfig_ADMIN_READ:
		return "ADMIN_READ"
	case iampb.AuditLogConfig_DATA_WRITE:
		return "DATA_WRITE"
	case iampb.AuditLogConfig_DATA_READ:
		return "DATA_READ"
	default:
		return "LOG_TYPE_UNSPECIFIED"
	}
}
