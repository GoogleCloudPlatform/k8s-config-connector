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

type NetworkFirewallPolicyAssociation struct{}

func NetworkFirewallPolicyAssociationToUnstructured(r *dclService.NetworkFirewallPolicyAssociation) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "ga",
			Type:    "NetworkFirewallPolicyAssociation",
		},
		Object: make(map[string]interface{}),
	}
	if r.AttachmentTarget != nil {
		u.Object["attachmentTarget"] = *r.AttachmentTarget
	}
	if r.FirewallPolicy != nil {
		u.Object["firewallPolicy"] = *r.FirewallPolicy
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
	if r.ShortName != nil {
		u.Object["shortName"] = *r.ShortName
	}
	return u
}

func UnstructuredToNetworkFirewallPolicyAssociation(u *unstructured.Resource) (*dclService.NetworkFirewallPolicyAssociation, error) {
	r := &dclService.NetworkFirewallPolicyAssociation{}
	if _, ok := u.Object["attachmentTarget"]; ok {
		if s, ok := u.Object["attachmentTarget"].(string); ok {
			r.AttachmentTarget = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AttachmentTarget: expected string")
		}
	}
	if _, ok := u.Object["firewallPolicy"]; ok {
		if s, ok := u.Object["firewallPolicy"].(string); ok {
			r.FirewallPolicy = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.FirewallPolicy: expected string")
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
	if _, ok := u.Object["shortName"]; ok {
		if s, ok := u.Object["shortName"].(string); ok {
			r.ShortName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ShortName: expected string")
		}
	}
	return r, nil
}

func GetNetworkFirewallPolicyAssociation(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicyAssociation(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetNetworkFirewallPolicyAssociation(ctx, r)
	if err != nil {
		return nil, err
	}
	return NetworkFirewallPolicyAssociationToUnstructured(r), nil
}

func ListNetworkFirewallPolicyAssociation(ctx context.Context, config *dcl.Config, project string, location string, firewallPolicy string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListNetworkFirewallPolicyAssociation(ctx, project, location, firewallPolicy)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, NetworkFirewallPolicyAssociationToUnstructured(r))
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

func ApplyNetworkFirewallPolicyAssociation(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicyAssociation(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNetworkFirewallPolicyAssociation(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyNetworkFirewallPolicyAssociation(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return NetworkFirewallPolicyAssociationToUnstructured(r), nil
}

func NetworkFirewallPolicyAssociationHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicyAssociation(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNetworkFirewallPolicyAssociation(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyNetworkFirewallPolicyAssociation(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteNetworkFirewallPolicyAssociation(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetworkFirewallPolicyAssociation(u)
	if err != nil {
		return err
	}
	return c.DeleteNetworkFirewallPolicyAssociation(ctx, r)
}

func NetworkFirewallPolicyAssociationID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToNetworkFirewallPolicyAssociation(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *NetworkFirewallPolicyAssociation) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"NetworkFirewallPolicyAssociation",
		"ga",
	}
}

func (r *NetworkFirewallPolicyAssociation) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyAssociation) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyAssociation) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyAssociation) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyAssociation) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyAssociation) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NetworkFirewallPolicyAssociation) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetNetworkFirewallPolicyAssociation(ctx, config, resource)
}

func (r *NetworkFirewallPolicyAssociation) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyNetworkFirewallPolicyAssociation(ctx, config, resource, opts...)
}

func (r *NetworkFirewallPolicyAssociation) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return NetworkFirewallPolicyAssociationHasDiff(ctx, config, resource, opts...)
}

func (r *NetworkFirewallPolicyAssociation) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteNetworkFirewallPolicyAssociation(ctx, config, resource)
}

func (r *NetworkFirewallPolicyAssociation) ID(resource *unstructured.Resource) (string, error) {
	return NetworkFirewallPolicyAssociationID(resource)
}

func init() {
	unstructured.Register(&NetworkFirewallPolicyAssociation{})
}
