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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type ServiceAccount struct{}

func ServiceAccountToUnstructured(r *dclService.ServiceAccount) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "iam",
			Version: "alpha",
			Type:    "ServiceAccount",
		},
		Object: make(map[string]interface{}),
	}
	if r.ActasResources != nil && r.ActasResources != dclService.EmptyServiceAccountActasResources {
		rActasResources := make(map[string]interface{})
		var rActasResourcesResources []interface{}
		for _, rActasResourcesResourcesVal := range r.ActasResources.Resources {
			rActasResourcesResourcesObject := make(map[string]interface{})
			if rActasResourcesResourcesVal.FullResourceName != nil {
				rActasResourcesResourcesObject["fullResourceName"] = *rActasResourcesResourcesVal.FullResourceName
			}
			rActasResourcesResources = append(rActasResourcesResources, rActasResourcesResourcesObject)
		}
		rActasResources["resources"] = rActasResourcesResources
		u.Object["actasResources"] = rActasResources
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Disabled != nil {
		u.Object["disabled"] = *r.Disabled
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.Email != nil {
		u.Object["email"] = *r.Email
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.OAuth2ClientId != nil {
		u.Object["oauth2ClientId"] = *r.OAuth2ClientId
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.UniqueId != nil {
		u.Object["uniqueId"] = *r.UniqueId
	}
	return u
}

func UnstructuredToServiceAccount(u *unstructured.Resource) (*dclService.ServiceAccount, error) {
	r := &dclService.ServiceAccount{}
	if _, ok := u.Object["actasResources"]; ok {
		if rActasResources, ok := u.Object["actasResources"].(map[string]interface{}); ok {
			r.ActasResources = &dclService.ServiceAccountActasResources{}
			if _, ok := rActasResources["resources"]; ok {
				if s, ok := rActasResources["resources"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rActasResourcesResources dclService.ServiceAccountActasResourcesResources
							if _, ok := objval["fullResourceName"]; ok {
								if s, ok := objval["fullResourceName"].(string); ok {
									rActasResourcesResources.FullResourceName = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rActasResourcesResources.FullResourceName: expected string")
								}
							}
							r.ActasResources.Resources = append(r.ActasResources.Resources, rActasResourcesResources)
						}
					}
				} else {
					return nil, fmt.Errorf("r.ActasResources.Resources: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ActasResources: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["disabled"]; ok {
		if b, ok := u.Object["disabled"].(bool); ok {
			r.Disabled = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Disabled: expected bool")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["email"]; ok {
		if s, ok := u.Object["email"].(string); ok {
			r.Email = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Email: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["oauth2ClientId"]; ok {
		if s, ok := u.Object["oauth2ClientId"].(string); ok {
			r.OAuth2ClientId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.OAuth2ClientId: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["uniqueId"]; ok {
		if s, ok := u.Object["uniqueId"].(string); ok {
			r.UniqueId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UniqueId: expected string")
		}
	}
	return r, nil
}

func GetServiceAccount(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceAccount(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetServiceAccount(ctx, r)
	if err != nil {
		return nil, err
	}
	return ServiceAccountToUnstructured(r), nil
}

func ListServiceAccount(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListServiceAccount(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ServiceAccountToUnstructured(r))
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

func ApplyServiceAccount(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceAccount(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToServiceAccount(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyServiceAccount(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ServiceAccountToUnstructured(r), nil
}

func ServiceAccountHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceAccount(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToServiceAccount(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyServiceAccount(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteServiceAccount(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToServiceAccount(u)
	if err != nil {
		return err
	}
	return c.DeleteServiceAccount(ctx, r)
}

func ServiceAccountID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToServiceAccount(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *ServiceAccount) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"iam",
		"ServiceAccount",
		"alpha",
	}
}

func SetPolicyServiceAccount(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToServiceAccount(u)
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

func SetPolicyWithEtagServiceAccount(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToServiceAccount(u)
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

func GetPolicyServiceAccount(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToServiceAccount(u)
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

func SetPolicyMemberServiceAccount(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToServiceAccount(u)
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

func GetPolicyMemberServiceAccount(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToServiceAccount(u)
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

func DeletePolicyMemberServiceAccount(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToServiceAccount(u)
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

func (r *ServiceAccount) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberServiceAccount(ctx, config, resource, member)
}

func (r *ServiceAccount) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberServiceAccount(ctx, config, resource, role, member)
}

func (r *ServiceAccount) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberServiceAccount(ctx, config, resource, member)
}

func (r *ServiceAccount) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyServiceAccount(ctx, config, resource, policy)
}

func (r *ServiceAccount) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagServiceAccount(ctx, config, resource, policy)
}

func (r *ServiceAccount) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyServiceAccount(ctx, config, resource)
}

func (r *ServiceAccount) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetServiceAccount(ctx, config, resource)
}

func (r *ServiceAccount) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyServiceAccount(ctx, config, resource, opts...)
}

func (r *ServiceAccount) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ServiceAccountHasDiff(ctx, config, resource, opts...)
}

func (r *ServiceAccount) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteServiceAccount(ctx, config, resource)
}

func (r *ServiceAccount) ID(resource *unstructured.Resource) (string, error) {
	return ServiceAccountID(resource)
}

func init() {
	unstructured.Register(&ServiceAccount{})
}
