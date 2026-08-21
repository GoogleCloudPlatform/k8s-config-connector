// Copyright 2025 Google LLC
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

package iam

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
)

type MemberIdentityResolver interface {
	Resolve(v1beta1.Member, *v1beta1.MemberSource, string) (string, error)
}

type iamBindingKey struct {
	Role      string
	Condition v1beta1.IAMCondition
}

// ComputePartialPolicyWithMergedBindings returns the IAMPartialPolicy that results after the user's intent (as specified by the input
// IAMPartialPolicy) is merged with the underlying IAM policy (as specified by the input IAMPolicy). This function also resolves all memberFrom
// fields to member fields and ensures the returned IAMPartialPolicy only contains member fields.

// The status.AllBindings in the returned IAMPartialPolicy reflects a mix of user specified bindings and the existing bindings associated with the GCP resource.
// The merge strategy takes effect on the member level with {role, condition} tuples as keys.
// The status.LastAppliedBindings in the returned IAMPartialPolicy reflects a list of canonical bindings that specified by users.
func ComputePartialPolicyWithMergedBindings(partialPolicy *v1beta1.IAMPartialPolicy, livePolicy *v1beta1.IAMPolicy, resolver MemberIdentityResolver) (*v1beta1.IAMPartialPolicy, error) {
	desiredPartialPolicy := partialPolicy.DeepCopy()
	specifiedBindings, err := ConvertIAMPartialBindingsToIAMPolicyBindings(partialPolicy, resolver)
	if err != nil {
		return nil, fmt.Errorf("error converting IAMPartialPolicy bindings to IAMPolicy bindings: %w", err)
	}

	// merge live bindings with user specified bindings
	mergeBindings := mergeBindingSlices(specifiedBindings, livePolicy.Spec.Bindings)
	// compute members that users intend to delete per binding
	toRemove := computeDeletedMembersPerBinding(specifiedBindings, partialPolicy.Status.LastAppliedBindings)
	// remove deleted members per binding
	desiredAllBindings := removeMembersPerBinding(mergeBindings, toRemove)

	// record lastAppliedBinding as user specified bindings
	sortBindingSlice(specifiedBindings)
	desiredPartialPolicy.Status.LastAppliedBindings = specifiedBindings
	sortBindingSlice(desiredAllBindings)
	desiredPartialPolicy.Status.AllBindings = desiredAllBindings
	return desiredPartialPolicy, nil
}

// ComputePartialPolicyWithRemainingBindings returns the IAMPartialPolicy that results after the user's last applied bindings (as specified by the input
// IAMPartialPolicy) are deleted from the underlying IAM Policy (as specified by the input IAMPolicy). This function is intended to be called on IAMPartialPolicy
// resources deletion.
//
// The status.AllBindings in the returned IAMPartialPolicy reflects the remaining bindings that are computed by pruning last applied bindings (bindings managed by KCC)
// from all the existing bindings from the underlying IAM Policy.
// The status.LastAppliedBindings in the returned IAMPartialPolicy will be cleared.
func ComputePartialPolicyWithRemainingBindings(partialPolicy *v1beta1.IAMPartialPolicy, livePolicy *v1beta1.IAMPolicy) *v1beta1.IAMPartialPolicy {
	desiredPartialPolicy := partialPolicy.DeepCopy()
	remainingBindings := removeMembersPerBinding(livePolicy.Spec.Bindings, partialPolicy.Status.LastAppliedBindings)
	// record the remaining bindings as (new) all bindings.
	sortBindingSlice(remainingBindings)
	desiredPartialPolicy.Status.AllBindings = remainingBindings
	// clear last applied bindings
	desiredPartialPolicy.Status.LastAppliedBindings = make([]v1beta1.IAMPolicyBinding, 0)
	return desiredPartialPolicy
}

func ConvertIAMPartialBindingsToIAMPolicyBindings(partialPolicy *v1beta1.IAMPartialPolicy, resolver MemberIdentityResolver) (bindings []v1beta1.IAMPolicyBinding, err error) {
	res := make([]v1beta1.IAMPolicyBinding, 0)
	for _, binding := range partialPolicy.Spec.Bindings {
		convertedBinding, err := toIAMPolicyBinding(binding, resolver, partialPolicy.Namespace)
		if err != nil {
			return bindings, fmt.Errorf("error converting IAMPartialPolicy binding to IAMPolicy binding: %w", err)
		}
		res = append(res, convertedBinding)
	}
	return mergeBindingsWithSameRoleAndCondition(res), nil
}

func toIAMPolicyBinding(b v1beta1.IAMPartialPolicyBinding, resolver MemberIdentityResolver, defaultNamespace string) (binding v1beta1.IAMPolicyBinding, err error) {
	members := make([]v1beta1.Member, 0)
	for _, member := range b.Members {
		resolvedMember, err := resolver.Resolve(member.Member, member.MemberFrom, defaultNamespace)
		if err != nil {
			return binding, fmt.Errorf("error resolving member identity of IAMPartialPolicy binding: %w", err)
		}
		members = append(members, v1beta1.Member(resolvedMember))
	}

	return v1beta1.IAMPolicyBinding{
		Role:      b.Role,
		Condition: b.Condition,
		Members:   members,
	}, nil
}

func mergeBindingsWithSameRoleAndCondition(bindings []v1beta1.IAMPolicyBinding) []v1beta1.IAMPolicyBinding {
	bindingMap := mergeBindings(bindings)
	mergedBindings := make([]v1beta1.IAMPolicyBinding, 0)
	for _, binding := range bindingMap {
		if len(binding.Members) > 0 {
			mergedBindings = append(mergedBindings, binding)
		}
	}
	return mergedBindings
}

func mergeBindingSlices(bindingSlice1, bindingSlice2 []v1beta1.IAMPolicyBinding) []v1beta1.IAMPolicyBinding {
	mergedBindings := make([]v1beta1.IAMPolicyBinding, 0)
	mergedBindings = append(mergedBindings, bindingSlice1...)
	mergedBindings = append(mergedBindings, bindingSlice2...)
	return mergeBindingsWithSameRoleAndCondition(mergedBindings)
}

func mergeBindings(bindings []v1beta1.IAMPolicyBinding) map[iamBindingKey]v1beta1.IAMPolicyBinding {
	bindingMap := make(map[iamBindingKey]v1beta1.IAMPolicyBinding)
	for _, binding := range bindings {
		key := getIamBindingKey(binding)
		b, ok := bindingMap[key]
		if !ok {
			bindingMap[key] = *binding.DeepCopy()
			continue
		}
		b.Members = mergeMembers(b.Members, binding.Members)
		bindingMap[key] = b
	}
	return bindingMap
}

func computeDeletedMembersPerBinding(bindings, lastAppliedBindings []v1beta1.IAMPolicyBinding) []v1beta1.IAMPolicyBinding {
	res := make([]v1beta1.IAMPolicyBinding, 0)
	bindingMap := mergeBindings(bindings)
	lastAppliedBindingMap := mergeBindings(lastAppliedBindings)
	for k, a := range lastAppliedBindingMap {
		b, ok := bindingMap[k]
		if !ok {
			res = append(res, *a.DeepCopy())
			continue
		}
		removedMembers := computeDeletedMembers(b.Members, a.Members)
		if len(removedMembers) > 0 {
			b.Members = removedMembers
			res = append(res, b)
		}
	}
	return res
}

func getIamBindingKey(binding v1beta1.IAMPolicyBinding) iamBindingKey {
	key := iamBindingKey{}
	key.Role = binding.Role
	if binding.Condition != nil {
		key.Condition = *binding.Condition
	}
	return key
}

func removeMembersPerBinding(bindings, deletedBindings []v1beta1.IAMPolicyBinding) []v1beta1.IAMPolicyBinding {
	bindingMap := mergeBindings(bindings)
	for _, binding := range deletedBindings {
		key := getIamBindingKey(binding)
		if mergedBinding, ok := bindingMap[key]; ok {
			mergedBinding.Members = removeDeletedMembers(mergedBinding.Members, binding.Members)
			bindingMap[key] = mergedBinding
		}
	}
	res := make([]v1beta1.IAMPolicyBinding, 0)
	for _, binding := range bindingMap {
		if len(binding.Members) > 0 {
			res = append(res, binding)
		}
	}
	return res
}

func removeDeletedMembers(members, deletedMembers []v1beta1.Member) []v1beta1.Member {
	memberMap := make(map[v1beta1.Member]bool)
	for _, deletedMember := range deletedMembers {
		memberMap[deletedMember] = true
	}
	res := make([]v1beta1.Member, 0)
	for _, m := range members {
		if _, ok := memberMap[m]; !ok {
			res = append(res, m)
		}
	}
	return res
}

func computeDeletedMembers(members, lastAppliedMembers []v1beta1.Member) []v1beta1.Member {
	memberMap := make(map[v1beta1.Member]bool)
	res := make([]v1beta1.Member, 0)
	for _, member := range members {
		memberMap[member] = true
	}
	for _, appliedMember := range lastAppliedMembers {
		if _, ok := memberMap[appliedMember]; !ok {
			res = append(res, appliedMember)
		}
	}
	return res
}

func mergeMembers(memberSlice1, memberSlice2 []v1beta1.Member) []v1beta1.Member {
	memberMap := make(map[v1beta1.Member]bool)
	for _, m := range memberSlice1 {
		memberMap[m] = true
	}
	for _, m := range memberSlice2 {
		memberMap[m] = true
	}
	res := make([]v1beta1.Member, 0)
	for k := range memberMap {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

func sortBindingSlice(bindings []v1beta1.IAMPolicyBinding) {
	slices.SortFunc(bindings, func(a, b v1beta1.IAMPolicyBinding) int {
		k1 := getIamBindingKey(a)
		k2 := getIamBindingKey(b)
		if k1.Role != k2.Role {
			return strings.Compare(k1.Role, k2.Role)
		}
		if k1.Condition.Title != k2.Condition.Title {
			return strings.Compare(k1.Condition.Title, k2.Condition.Title)
		}
		if k1.Condition.Description != k2.Condition.Description {
			return strings.Compare(k1.Condition.Description, k2.Condition.Description)
		}
		if k1.Condition.Expression != k2.Condition.Expression {
			return strings.Compare(k1.Condition.Expression, k2.Condition.Expression)
		}
		// return 0 // returns 0 means incomparable aka k1 == k2; https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings
		// so we force an ordering in this case of k1 < k2 for determinism in other areas of the stack (like status displays)
		return -1
	})
}
