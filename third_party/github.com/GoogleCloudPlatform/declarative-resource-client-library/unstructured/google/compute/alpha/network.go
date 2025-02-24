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

type Network struct{}

func NetworkToUnstructured(r *dclService.Network) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "alpha",
			Type:    "Network",
		},
		Object: make(map[string]interface{}),
	}
	if r.AutoCreateSubnetworks != nil {
		u.Object["autoCreateSubnetworks"] = *r.AutoCreateSubnetworks
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.GatewayIPv4 != nil {
		u.Object["gatewayIPv4"] = *r.GatewayIPv4
	}
	if r.Mtu != nil {
		u.Object["mtu"] = *r.Mtu
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.RoutingConfig != nil && r.RoutingConfig != dclService.EmptyNetworkRoutingConfig {
		rRoutingConfig := make(map[string]interface{})
		if r.RoutingConfig.RoutingMode != nil {
			rRoutingConfig["routingMode"] = string(*r.RoutingConfig.RoutingMode)
		}
		u.Object["routingConfig"] = rRoutingConfig
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.SelfLinkWithId != nil {
		u.Object["selfLinkWithId"] = *r.SelfLinkWithId
	}
	return u
}

func UnstructuredToNetwork(u *unstructured.Resource) (*dclService.Network, error) {
	r := &dclService.Network{}
	if _, ok := u.Object["autoCreateSubnetworks"]; ok {
		if b, ok := u.Object["autoCreateSubnetworks"].(bool); ok {
			r.AutoCreateSubnetworks = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.AutoCreateSubnetworks: expected bool")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["gatewayIPv4"]; ok {
		if s, ok := u.Object["gatewayIPv4"].(string); ok {
			r.GatewayIPv4 = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.GatewayIPv4: expected string")
		}
	}
	if _, ok := u.Object["mtu"]; ok {
		if i, ok := u.Object["mtu"].(int64); ok {
			r.Mtu = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Mtu: expected int64")
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
	if _, ok := u.Object["routingConfig"]; ok {
		if rRoutingConfig, ok := u.Object["routingConfig"].(map[string]interface{}); ok {
			r.RoutingConfig = &dclService.NetworkRoutingConfig{}
			if _, ok := rRoutingConfig["routingMode"]; ok {
				if s, ok := rRoutingConfig["routingMode"].(string); ok {
					r.RoutingConfig.RoutingMode = dclService.NetworkRoutingConfigRoutingModeEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.RoutingConfig.RoutingMode: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.RoutingConfig: expected map[string]interface{}")
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

func GetNetwork(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetwork(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetNetwork(ctx, r)
	if err != nil {
		return nil, err
	}
	return NetworkToUnstructured(r), nil
}

func ListNetwork(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListNetwork(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, NetworkToUnstructured(r))
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

func ApplyNetwork(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetwork(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNetwork(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyNetwork(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return NetworkToUnstructured(r), nil
}

func NetworkHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetwork(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNetwork(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyNetwork(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteNetwork(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNetwork(u)
	if err != nil {
		return err
	}
	return c.DeleteNetwork(ctx, r)
}

func NetworkID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToNetwork(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Network) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"Network",
		"alpha",
	}
}

func (r *Network) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Network) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Network) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Network) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Network) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Network) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Network) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetNetwork(ctx, config, resource)
}

func (r *Network) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyNetwork(ctx, config, resource, opts...)
}

func (r *Network) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return NetworkHasDiff(ctx, config, resource, opts...)
}

func (r *Network) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteNetwork(ctx, config, resource)
}

func (r *Network) ID(resource *unstructured.Resource) (string, error) {
	return NetworkID(resource)
}

func init() {
	unstructured.Register(&Network{})
}
