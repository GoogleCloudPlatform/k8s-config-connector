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

package iamresource

import (
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SplitPolicy(policy *v1beta1.IAMPolicy) []v1beta1.IAMPolicyMember {
	policyMembers := make([]v1beta1.IAMPolicyMember, 0)
	for _, binding := range policy.Spec.Bindings {
		for _, member := range binding.Members {
			policyMember := v1beta1.IAMPolicyMember{
				TypeMeta: v1.TypeMeta{
					APIVersion: v1beta1.IAMPolicyMemberGVK.GroupVersion().String(),
					Kind:       v1beta1.IAMPolicyMemberGVK.Kind,
				},
				ObjectMeta: v1.ObjectMeta{
					Annotations: deepcopy.StringStringMap(policy.Annotations),
					Labels:      deepcopy.StringStringMap(policy.Labels),
					Name:        fmt.Sprintf("%v-%v", policy.Name, len(policyMembers)+1),
					Namespace:   policy.Namespace,
				},
				Spec: v1beta1.IAMPolicyMemberSpec{
					Member:            member,
					Role:              binding.Role,
					Condition:         binding.Condition,
					ResourceReference: policy.Spec.ResourceReference,
				},
			}
			policyMembers = append(policyMembers, policyMember)
		}
	}
	return policyMembers
}

func ConvertIAMPolicyToIAMPartialPolicy(policy *v1beta1.IAMPolicy) *v1beta1.IAMPartialPolicy {
	partialPolicy := &v1beta1.IAMPartialPolicy{
		TypeMeta: v1.TypeMeta{
			APIVersion: v1beta1.IAMPartialPolicyGVK.GroupVersion().String(),
			Kind:       v1beta1.IAMPartialPolicyGVK.Kind,
		},
		ObjectMeta: v1.ObjectMeta{
			Annotations: deepcopy.StringStringMap(policy.Annotations),
			Labels:      deepcopy.StringStringMap(policy.Labels),
			Name:        policy.Name,
			Namespace:   policy.Namespace,
		},
		Spec: v1beta1.IAMPartialPolicySpec{
			ResourceReference: policy.Spec.ResourceReference,
		},
	}

	partialPolicyBindings := make([]v1beta1.IAMPartialPolicyBinding, 0, len(policy.Spec.Bindings))
	for _, binding := range policy.Spec.Bindings {
		partialPolicyBinding := v1beta1.IAMPartialPolicyBinding{
			Role:      binding.Role,
			Condition: binding.Condition,
			Members:   convertToIAMPartialPolicyMembers(binding.Members),
		}
		partialPolicyBindings = append(partialPolicyBindings, partialPolicyBinding)
	}
	partialPolicy.Spec.Bindings = partialPolicyBindings
	return partialPolicy
}

func convertToIAMPartialPolicyMembers(members []v1beta1.Member) []v1beta1.IAMPartialPolicyMember {
	res := make([]v1beta1.IAMPartialPolicyMember, 0, len(members))
	for _, m := range members {
		res = append(res, v1beta1.IAMPartialPolicyMember{
			Member: m,
		})
	}
	return res
}
