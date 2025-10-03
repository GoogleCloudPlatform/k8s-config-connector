// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package iam

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	iamDCL "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

// Member describes a single IAMPolicyMember.
type Member struct{}

// MemberToUnstructured converts a DCL IAM PolicyMember to an unstructured.Resource.
func MemberToUnstructured(r *iamDCL.Member) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "iam",
			Version: "ga",
			Type:    "PolicyMember",
		},
		Object: make(map[string]interface{}),
	}
	if r.Role != nil {
		u.Object["role"] = string(*r.Role)
	}
	if r.Member != nil {
		u.Object["member"] = string(*r.Member)
	}
	return u
}

// UnstructuredToMember converts an unstructured.Resource to a DCL IAM PolicyMember.
func UnstructuredToMember(u *unstructured.Resource) (*iamDCL.Member, error) {
	r := &iamDCL.Member{}
	if _, ok := u.Object["role"]; ok {
		s, ok := u.Object["role"].(string)
		if !ok {
			return nil, fmt.Errorf("r.role: expected string, got %v", s)
		}
		r.Role = dcl.String(s)
	}
	if _, ok := u.Object["member"]; ok {
		s, ok := u.Object["member"].(string)
		if !ok {
			return nil, fmt.Errorf("r.member: expected string, got %v", s)
		}
		r.Member = dcl.String(s)
	}
	return r, nil
}

// GetMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func GetMember(_ context.Context, _ *dcl.Config, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// ApplyMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func ApplyMember(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// MemberHasDiff is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func MemberHasDiff(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ ...dcl.ApplyOption) (bool, error) {
	return false, unstructured.ErrNoSuchMethod
}

// DeleteMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func DeleteMember(_ context.Context, _ *dcl.Config, _ *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

// IDMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func IDMember(_ *unstructured.Resource) (string, error) {
	return "", unstructured.ErrNoSuchMethod
}

// MemberGetPolicy is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func MemberGetPolicy(_ context.Context, _ *dcl.Config, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// MemberSetPolicy is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func MemberSetPolicy(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// MemberSetPolicyWithEtag is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func MemberSetPolicyWithEtag(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// MemberGetPolicyMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func MemberGetPolicyMember(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _, _ string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// MemberSetPolicyMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func MemberSetPolicyMember(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// MemberDeletePolicyMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func MemberDeletePolicyMember(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

// STV returns the ServiceTypeVersion of the IAM PolicyMember resource.
func (r *Member) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		Service: "iam",
		Type:    "PolicyMember",
		Version: "ga",
	}
}

// Get calls the empty GetMember function.
func (r *Member) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetMember(ctx, config, resource)
}

// Apply calls the empty ApplyMember function.
func (r *Member) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyMember(ctx, config, resource, opts...)
}

// HasDiff calls the empty MemberHasDiff function.
func (r *Member) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return MemberHasDiff(ctx, config, resource, opts...)
}

// Delete calls the empty DeleteMember function.
func (r *Member) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteMember(ctx, config, resource)
}

// ID calls the empty IDMember function.
func (r *Member) ID(resource *unstructured.Resource) (string, error) {
	return IDMember(resource)
}

// GetPolicy calls the empty MemberGetPolicy function.
func (r *Member) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return MemberGetPolicy(ctx, config, resource)
}

// SetPolicy calls the empty MemberSetPolicy function.
func (r *Member) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return MemberSetPolicy(ctx, config, resource, policy)
}

// SetPolicyWithEtag calls the empty MemberSetPolicyWithEtag function.
func (r *Member) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return MemberSetPolicyWithEtag(ctx, config, resource, policy)
}

// GetPolicyMember calls the empty MemberGetPolicyMember function.
func (r *Member) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return MemberGetPolicyMember(ctx, config, resource, role, member)
}

// SetPolicyMember calls the empty MemberSetPolicyMember function.
func (r *Member) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return MemberSetPolicyMember(ctx, config, resource, member)
}

// DeletePolicyMember calls the empty MemberDeletePolicyMember function.
func (r *Member) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return MemberDeletePolicyMember(ctx, config, resource, member)
}

func init() {
	unstructured.Register(&Member{})
}
