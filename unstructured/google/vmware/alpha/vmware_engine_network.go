// Copyright 2023 Google LLC. All Rights Reserved.
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
package vmware

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vmware/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type VmwareEngineNetwork struct{}

func VmwareEngineNetworkToUnstructured(r *dclService.VmwareEngineNetwork) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "vmware",
			Version: "alpha",
			Type:    "VmwareEngineNetwork",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
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
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.Type != nil {
		u.Object["type"] = string(*r.Type)
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	var rVPCNetworks []interface{}
	for _, rVPCNetworksVal := range r.VPCNetworks {
		rVPCNetworksObject := make(map[string]interface{})
		if rVPCNetworksVal.Network != nil {
			rVPCNetworksObject["network"] = *rVPCNetworksVal.Network
		}
		if rVPCNetworksVal.Type != nil {
			rVPCNetworksObject["type"] = string(*rVPCNetworksVal.Type)
		}
		rVPCNetworks = append(rVPCNetworks, rVPCNetworksObject)
	}
	u.Object["vpcNetworks"] = rVPCNetworks
	return u
}

func UnstructuredToVmwareEngineNetwork(u *unstructured.Resource) (*dclService.VmwareEngineNetwork, error) {
	r := &dclService.VmwareEngineNetwork{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
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
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.VmwareEngineNetworkStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dclService.VmwareEngineNetworkTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
		}
	}
	if _, ok := u.Object["uid"]; ok {
		if s, ok := u.Object["uid"].(string); ok {
			r.Uid = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Uid: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["vpcNetworks"]; ok {
		if s, ok := u.Object["vpcNetworks"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rVPCNetworks dclService.VmwareEngineNetworkVPCNetworks
					if _, ok := objval["network"]; ok {
						if s, ok := objval["network"].(string); ok {
							rVPCNetworks.Network = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rVPCNetworks.Network: expected string")
						}
					}
					if _, ok := objval["type"]; ok {
						if s, ok := objval["type"].(string); ok {
							rVPCNetworks.Type = dclService.VmwareEngineNetworkVPCNetworksTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rVPCNetworks.Type: expected string")
						}
					}
					r.VPCNetworks = append(r.VPCNetworks, rVPCNetworks)
				}
			}
		} else {
			return nil, fmt.Errorf("r.VPCNetworks: expected []interface{}")
		}
	}
	return r, nil
}

func GetVmwareEngineNetwork(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToVmwareEngineNetwork(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetVmwareEngineNetwork(ctx, r)
	if err != nil {
		return nil, err
	}
	return VmwareEngineNetworkToUnstructured(r), nil
}

func ListVmwareEngineNetwork(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListVmwareEngineNetwork(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, VmwareEngineNetworkToUnstructured(r))
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

func ApplyVmwareEngineNetwork(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToVmwareEngineNetwork(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToVmwareEngineNetwork(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyVmwareEngineNetwork(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return VmwareEngineNetworkToUnstructured(r), nil
}

func VmwareEngineNetworkHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToVmwareEngineNetwork(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToVmwareEngineNetwork(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyVmwareEngineNetwork(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteVmwareEngineNetwork(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToVmwareEngineNetwork(u)
	if err != nil {
		return err
	}
	return c.DeleteVmwareEngineNetwork(ctx, r)
}

func VmwareEngineNetworkID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToVmwareEngineNetwork(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *VmwareEngineNetwork) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"vmware",
		"VmwareEngineNetwork",
		"alpha",
	}
}

func (r *VmwareEngineNetwork) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VmwareEngineNetwork) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VmwareEngineNetwork) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *VmwareEngineNetwork) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VmwareEngineNetwork) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VmwareEngineNetwork) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VmwareEngineNetwork) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetVmwareEngineNetwork(ctx, config, resource)
}

func (r *VmwareEngineNetwork) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyVmwareEngineNetwork(ctx, config, resource, opts...)
}

func (r *VmwareEngineNetwork) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return VmwareEngineNetworkHasDiff(ctx, config, resource, opts...)
}

func (r *VmwareEngineNetwork) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteVmwareEngineNetwork(ctx, config, resource)
}

func (r *VmwareEngineNetwork) ID(resource *unstructured.Resource) (string, error) {
	return VmwareEngineNetworkID(resource)
}

func init() {
	unstructured.Register(&VmwareEngineNetwork{})
}
