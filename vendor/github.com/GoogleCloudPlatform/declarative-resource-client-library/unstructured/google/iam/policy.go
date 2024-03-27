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
// Package iam contains handwritten IAMPolicy-to-unstructured functions.
package iam

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	iamDCL "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

// Policy describes a single IAMPolicy.
type Policy struct{}

// PolicyToUnstructured converts a DCL IAM Policy to an unstructured.Resource.
func PolicyToUnstructured(r *iamDCL.Policy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "iam",
			Version: "ga",
			Type:    "Policy",
		},
		Object: make(map[string]interface{}),
	}
	var rBindings []interface{}
	for _, rBindingsVal := range r.Bindings {
		rBindingsObject := make(map[string]interface{})
		if rBindingsVal.Role != nil {
			rBindingsObject["role"] = string(*rBindingsVal.Role)
		}
		if rBindingsVal.Members != nil {
			var rBindingsValMembers []interface{}
			for _, rBindingsValMembersVal := range rBindingsVal.Members {
				rBindingsValMembers = append(rBindingsValMembers, rBindingsValMembersVal)
			}
			rBindingsObject["members"] = rBindingsValMembers
		}
		if rBindingsVal.Condition != nil {
			rBindingsValCondition := make(map[string]interface{})
			if rBindingsVal.Condition.Title != nil {
				rBindingsValCondition["title"] = string(*rBindingsVal.Condition.Title)
			}
			if rBindingsVal.Condition.Description != nil {
				rBindingsValCondition["description"] = string(*rBindingsVal.Condition.Description)
			}
			if rBindingsVal.Condition.Expression != nil {
				rBindingsValCondition["expression"] = string(*rBindingsVal.Condition.Expression)
			}
			rBindingsObject["condition"] = rBindingsValCondition
		}
		rBindings = append(rBindings, rBindingsObject)
	}
	u.Object["bindings"] = rBindings
	if r.Etag != nil {
		u.Object["etag"] = string(*r.Etag)
	}
	if r.Version != nil {
		u.Object["version"] = int(*r.Version)
	}
	return u
}

// UnstructuredToPolicy converts an unstructured.Resource to a DCL IAM Policy.
func UnstructuredToPolicy(u *unstructured.Resource) (*iamDCL.Policy, error) {
	r := &iamDCL.Policy{}
	if _, ok := u.Object["bindings"]; ok {
		s, ok := u.Object["bindings"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("r.Bindings: expected []interface{}")
		}
		for _, o := range s {
			objval, ok := o.(map[string]interface{})
			if !ok {
				continue
			}
			var rBinding iamDCL.Binding
			if _, ok := objval["role"]; ok {
				s, ok := objval["role"].(string)
				if !ok {
					return nil, fmt.Errorf("rBinding.Role: expected string")
				}
				rBinding.Role = dcl.String(s)
			}
			if _, ok := objval["members"]; ok {
				if s, ok := objval["members"].([]interface{}); ok {
					for _, ss := range s {
						strval, ok := ss.(string)
						if !ok {
							return nil, fmt.Errorf("rBinding.Members: expected []interface{}")
						}
						rBinding.Members = append(rBinding.Members, strval)
					}
				}
			}
			if _, ok := objval["condition"]; ok {
				if rBindingCondition, ok := objval["condition"].(map[string]interface{}); ok {
					rBinding.Condition = &iamDCL.Condition{}
					if s, ok := rBindingCondition["title"].(string); ok {
						rBinding.Condition.Title = dcl.String(s)
					}
					if s, ok := rBindingCondition["description"].(string); ok {
						rBinding.Condition.Description = dcl.String(s)
					}
					if s, ok := rBindingCondition["expression"].(string); ok {
						rBinding.Condition.Expression = dcl.String(s)
					}
				}
			}
			r.Bindings = append(r.Bindings, rBinding)
		}
	}
	if _, ok := u.Object["etag"]; ok {
		s, ok := u.Object["etag"].(string)
		if !ok {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
		r.Etag = dcl.String(s)
	}
	if _, ok := u.Object["version"]; ok {
		i, ok := u.Object["version"].(int)
		if !ok {
			return nil, fmt.Errorf("r.Version: expected int64")
		}
		r.Version = &i
	}
	return r, nil
}

// GetPolicy is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func GetPolicy(_ context.Context, _ *dcl.Config, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// ApplyPolicy is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func ApplyPolicy(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// PolicyHasDiff is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func PolicyHasDiff(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ ...dcl.ApplyOption) (bool, error) {
	return false, unstructured.ErrNoSuchMethod
}

// DeletePolicy is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func DeletePolicy(_ context.Context, _ *dcl.Config, _ *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

// IDPolicy is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func IDPolicy(_ *unstructured.Resource) (string, error) {
	return "", unstructured.ErrNoSuchMethod
}

// PolicyGetPolicy is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func PolicyGetPolicy(_ context.Context, _ *dcl.Config, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// PolicySetPolicy is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func PolicySetPolicy(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// PolicySetPolicyWithEtag is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func PolicySetPolicyWithEtag(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// PolicyGetPolicyMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func PolicyGetPolicyMember(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _, _ string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// PolicySetPolicyMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func PolicySetPolicyMember(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

// PolicyDeletePolicyMember is an empty function that will return an error because these types are meant
// to be handled by the resources that the policies apply to rather than by explicit policy resources.
func PolicyDeletePolicyMember(_ context.Context, _ *dcl.Config, _ *unstructured.Resource, _ *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

// STV returns the ServiceTypeVersion of the IAM Policy resource.
func (r *Policy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		Service: "iam",
		Type:    "Policy",
		Version: "ga",
	}
}

// Get calls the empty GetPolicy function.
func (r *Policy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicy(ctx, config, resource)
}

// Apply calls the empty ApplyPolicy function.
func (r *Policy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyPolicy(ctx, config, resource, opts...)
}

// HasDiff calls the empty PolicyHasDiff function.
func (r *Policy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return PolicyHasDiff(ctx, config, resource, opts...)
}

// Delete calls the empty DeletePolicy function.
func (r *Policy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeletePolicy(ctx, config, resource)
}

// ID calls the empty IDPolicy function.
func (r *Policy) ID(resource *unstructured.Resource) (string, error) {
	return IDPolicy(resource)
}

// GetPolicy calls the empty PolicyGetPolicy function.
func (r *Policy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return PolicyGetPolicy(ctx, config, resource)
}

// SetPolicy calls the empty PolicySetPolicy function.
func (r *Policy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return PolicySetPolicy(ctx, config, resource, policy)
}

// SetPolicyWithEtag calls the empty PolicySetPolicyWithEtag function.
func (r *Policy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return PolicySetPolicyWithEtag(ctx, config, resource, policy)
}

// GetPolicyMember calls the empty PolicyGetPolicyMember function.
func (r *Policy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return PolicyGetPolicyMember(ctx, config, resource, role, member)
}

// SetPolicyMember calls the empty PolicySetPolicyMember function.
func (r *Policy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return PolicySetPolicyMember(ctx, config, resource, member)
}

// DeletePolicyMember calls the empty PolicyDeletePolicyMember function.
func (r *Policy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return PolicyDeletePolicyMember(ctx, config, resource, member)
}

func init() {
	unstructured.Register(&Policy{})
}
