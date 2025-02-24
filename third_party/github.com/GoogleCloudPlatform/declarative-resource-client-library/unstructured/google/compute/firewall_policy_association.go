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

type FirewallPolicyAssociation struct{}

func FirewallPolicyAssociationToUnstructured(r *dclService.FirewallPolicyAssociation) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "ga",
			Type:    "FirewallPolicyAssociation",
		},
		Object: make(map[string]interface{}),
	}
	if r.AttachmentTarget != nil {
		u.Object["attachmentTarget"] = *r.AttachmentTarget
	}
	if r.FirewallPolicy != nil {
		u.Object["firewallPolicy"] = *r.FirewallPolicy
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.ShortName != nil {
		u.Object["shortName"] = *r.ShortName
	}
	return u
}

func UnstructuredToFirewallPolicyAssociation(u *unstructured.Resource) (*dclService.FirewallPolicyAssociation, error) {
	r := &dclService.FirewallPolicyAssociation{}
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
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
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

func GetFirewallPolicyAssociation(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicyAssociation(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetFirewallPolicyAssociation(ctx, r)
	if err != nil {
		return nil, err
	}
	return FirewallPolicyAssociationToUnstructured(r), nil
}

func ListFirewallPolicyAssociation(ctx context.Context, config *dcl.Config, firewallPolicy string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListFirewallPolicyAssociation(ctx, firewallPolicy)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, FirewallPolicyAssociationToUnstructured(r))
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

func ApplyFirewallPolicyAssociation(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicyAssociation(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFirewallPolicyAssociation(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyFirewallPolicyAssociation(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return FirewallPolicyAssociationToUnstructured(r), nil
}

func FirewallPolicyAssociationHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicyAssociation(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFirewallPolicyAssociation(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyFirewallPolicyAssociation(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteFirewallPolicyAssociation(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFirewallPolicyAssociation(u)
	if err != nil {
		return err
	}
	return c.DeleteFirewallPolicyAssociation(ctx, r)
}

func FirewallPolicyAssociationID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToFirewallPolicyAssociation(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *FirewallPolicyAssociation) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"FirewallPolicyAssociation",
		"ga",
	}
}

func (r *FirewallPolicyAssociation) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyAssociation) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyAssociation) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyAssociation) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyAssociation) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyAssociation) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *FirewallPolicyAssociation) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetFirewallPolicyAssociation(ctx, config, resource)
}

func (r *FirewallPolicyAssociation) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyFirewallPolicyAssociation(ctx, config, resource, opts...)
}

func (r *FirewallPolicyAssociation) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return FirewallPolicyAssociationHasDiff(ctx, config, resource, opts...)
}

func (r *FirewallPolicyAssociation) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteFirewallPolicyAssociation(ctx, config, resource)
}

func (r *FirewallPolicyAssociation) ID(resource *unstructured.Resource) (string, error) {
	return FirewallPolicyAssociationID(resource)
}

func init() {
	unstructured.Register(&FirewallPolicyAssociation{})
}
