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
package apigee

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type Environment struct{}

func EnvironmentToUnstructured(r *dclService.Environment) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "apigee",
			Version: "beta",
			Type:    "Environment",
		},
		Object: make(map[string]interface{}),
	}
	if r.ApigeeOrganization != nil {
		u.Object["apigeeOrganization"] = *r.ApigeeOrganization
	}
	if r.CreatedAt != nil {
		u.Object["createdAt"] = *r.CreatedAt
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.LastModifiedAt != nil {
		u.Object["lastModifiedAt"] = *r.LastModifiedAt
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Properties != nil {
		rProperties := make(map[string]interface{})
		for k, v := range r.Properties {
			rProperties[k] = v
		}
		u.Object["properties"] = rProperties
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	return u
}

func UnstructuredToEnvironment(u *unstructured.Resource) (*dclService.Environment, error) {
	r := &dclService.Environment{}
	if _, ok := u.Object["apigeeOrganization"]; ok {
		if s, ok := u.Object["apigeeOrganization"].(string); ok {
			r.ApigeeOrganization = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ApigeeOrganization: expected string")
		}
	}
	if _, ok := u.Object["createdAt"]; ok {
		if i, ok := u.Object["createdAt"].(int64); ok {
			r.CreatedAt = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.CreatedAt: expected int64")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["lastModifiedAt"]; ok {
		if i, ok := u.Object["lastModifiedAt"].(int64); ok {
			r.LastModifiedAt = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.LastModifiedAt: expected int64")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["properties"]; ok {
		if rProperties, ok := u.Object["properties"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rProperties {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Properties = m
		} else {
			return nil, fmt.Errorf("r.Properties: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.EnvironmentStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	return r, nil
}

func GetEnvironment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetEnvironment(ctx, r)
	if err != nil {
		return nil, err
	}
	return EnvironmentToUnstructured(r), nil
}

func ListEnvironment(ctx context.Context, config *dcl.Config, apigeeOrganization string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListEnvironment(ctx, apigeeOrganization)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, EnvironmentToUnstructured(r))
		}
		if !l.HasNext() {
			break
		}
		if err := l.Next(ctx, c); err != nil {
			return nil, err
		}
	}
	return resources, nil
}

func ApplyEnvironment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEnvironment(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyEnvironment(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return EnvironmentToUnstructured(r), nil
}

func EnvironmentHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToEnvironment(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyEnvironment(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteEnvironment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return err
	}
	return c.DeleteEnvironment(ctx, r)
}

func EnvironmentID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Environment) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"apigee",
		"Environment",
		"beta",
	}
}

func SetPolicyEnvironment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func SetPolicyWithEtagEnvironment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicyWithEtag(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func GetPolicyEnvironment(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policy, err := iamClient.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func SetPolicyMemberEnvironment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return nil, err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return nil, err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	policy, err := iamClient.SetMember(ctx, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func GetPolicyMemberEnvironment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policyMember, err := iamClient.GetMember(ctx, r, role, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.MemberToUnstructured(policyMember), nil
}

func DeletePolicyMemberEnvironment(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToEnvironment(u)
	if err != nil {
		return err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	if err := iamClient.DeleteMember(ctx, member); err != nil {
		return err
	}
	return nil
}

func (r *Environment) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberEnvironment(ctx, config, resource, member)
}

func (r *Environment) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberEnvironment(ctx, config, resource, role, member)
}

func (r *Environment) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberEnvironment(ctx, config, resource, member)
}

func (r *Environment) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyEnvironment(ctx, config, resource, policy)
}

func (r *Environment) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagEnvironment(ctx, config, resource, policy)
}

func (r *Environment) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyEnvironment(ctx, config, resource)
}

func (r *Environment) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetEnvironment(ctx, config, resource)
}

func (r *Environment) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyEnvironment(ctx, config, resource, opts...)
}

func (r *Environment) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return EnvironmentHasDiff(ctx, config, resource, opts...)
}

func (r *Environment) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteEnvironment(ctx, config, resource)
}

func (r *Environment) ID(resource *unstructured.Resource) (string, error) {
	return EnvironmentID(resource)
}

func init() {
	unstructured.Register(&Environment{})
}
