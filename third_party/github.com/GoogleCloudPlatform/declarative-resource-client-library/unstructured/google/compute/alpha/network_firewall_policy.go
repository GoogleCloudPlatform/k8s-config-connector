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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type NetworkFirewallPolicy struct{}

func NetworkFirewallPolicyToUnstructured(r *dclService.NetworkFirewallPolicy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "alpha",
			Type:    "NetworkFirewallPolicy",
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
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Region != nil {
		u.Object["region"] = *r.Region
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
	return u
}

func UnstructuredToNetworkFirewallPolicy(u *unstructured.Resource) (*dclService.NetworkFirewallPolicy, error) {
	r := &dclService.NetworkFirewallPolicy{}
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
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["region"]; ok {
		if s, ok := u.Object["region"].(string); ok {
			r.Region = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Region: expected string")
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
	return r, nil
}

func GetNetworkFirewallPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetNetworkFirewallPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return NetworkFirewallPolicyToUnstructured(r), nil
}

func ListNetworkFirewallPolicy(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListNetworkFirewallPolicy(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, NetworkFirewallPolicyToUnstructured(r))
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

func ApplyNetworkFirewallPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNetworkFirewallPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyNetworkFirewallPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return NetworkFirewallPolicyToUnstructured(r), nil
}

func NetworkFirewallPolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNetworkFirewallPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyNetworkFirewallPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteNetworkFirewallPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicy(u)
	if err != nil {
		return err
	}
	return c.DeleteNetworkFirewallPolicy(ctx, r)
}

func NetworkFirewallPolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToNetworkFirewallPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *NetworkFirewallPolicy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"NetworkFirewallPolicy",
		"alpha",
	}
}

func (r *NetworkFirewallPolicy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetNetworkFirewallPolicy(ctx, config, resource)
}

func (r *NetworkFirewallPolicy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyNetworkFirewallPolicy(ctx, config, resource, opts...)
}

func (r *NetworkFirewallPolicy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return NetworkFirewallPolicyHasDiff(ctx, config, resource, opts...)
}

func (r *NetworkFirewallPolicy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteNetworkFirewallPolicy(ctx, config, resource)
}

func (r *NetworkFirewallPolicy) ID(resource *unstructured.Resource) (string, error) {
	return NetworkFirewallPolicyID(resource)
}

func init() {
	unstructured.Register(&NetworkFirewallPolicy{})
}
