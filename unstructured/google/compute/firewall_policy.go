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
package compute

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type FirewallPolicy struct{}

func FirewallPolicyToUnstructured(r *dclService.FirewallPolicy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "ga",
			Type:    "FirewallPolicy",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreationTimestamp != nil {
		u.Object["creationTimestamp"] = *r.CreationTimestamp
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Fingerprint != nil {
		u.Object["fingerprint"] = *r.Fingerprint
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Parent != nil {
		u.Object["parent"] = *r.Parent
	}
	if r.RuleTupleCount != nil {
		u.Object["ruleTupleCount"] = *r.RuleTupleCount
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.SelfLinkWithId != nil {
		u.Object["selfLinkWithId"] = *r.SelfLinkWithId
	}
	if r.ShortName != nil {
		u.Object["shortName"] = *r.ShortName
	}
	return u
}

func UnstructuredToFirewallPolicy(u *unstructured.Resource) (*dclService.FirewallPolicy, error) {
	r := &dclService.FirewallPolicy{}
	if _, ok := u.Object["creationTimestamp"]; ok {
		if s, ok := u.Object["creationTimestamp"].(string); ok {
			r.CreationTimestamp = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreationTimestamp: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["fingerprint"]; ok {
		if s, ok := u.Object["fingerprint"].(string); ok {
			r.Fingerprint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Fingerprint: expected string")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if s, ok := u.Object["id"].(string); ok {
			r.Id = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Id: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["parent"]; ok {
		if s, ok := u.Object["parent"].(string); ok {
			r.Parent = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Parent: expected string")
		}
	}
	if _, ok := u.Object["ruleTupleCount"]; ok {
		if i, ok := u.Object["ruleTupleCount"].(int64); ok {
			r.RuleTupleCount = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.RuleTupleCount: expected int64")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["selfLinkWithId"]; ok {
		if s, ok := u.Object["selfLinkWithId"].(string); ok {
			r.SelfLinkWithId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLinkWithId: expected string")
		}
	}
	if _, ok := u.Object["shortName"]; ok {
		if s, ok := u.Object["shortName"].(string); ok {
			r.ShortName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ShortName: expected string")
		}
	}
	return r, nil
}

func GetFirewallPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetFirewallPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return FirewallPolicyToUnstructured(r), nil
}

func ListFirewallPolicy(ctx context.Context, config *dcl.Config, parent string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListFirewallPolicy(ctx, parent)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, FirewallPolicyToUnstructured(r))
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

func ApplyFirewallPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFirewallPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyFirewallPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return FirewallPolicyToUnstructured(r), nil
}

func FirewallPolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFirewallPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyFirewallPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteFirewallPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicy(u)
	if err != nil {
		return err
	}
	return c.DeleteFirewallPolicy(ctx, r)
}

func FirewallPolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToFirewallPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *FirewallPolicy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"FirewallPolicy",
		"ga",
	}
}

func (r *FirewallPolicy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetFirewallPolicy(ctx, config, resource)
}

func (r *FirewallPolicy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyFirewallPolicy(ctx, config, resource, opts...)
}

func (r *FirewallPolicy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return FirewallPolicyHasDiff(ctx, config, resource, opts...)
}

func (r *FirewallPolicy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteFirewallPolicy(ctx, config, resource)
}

func (r *FirewallPolicy) ID(resource *unstructured.Resource) (string, error) {
	return FirewallPolicyID(resource)
}

func init() {
	unstructured.Register(&FirewallPolicy{})
}
