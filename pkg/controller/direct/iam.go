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
