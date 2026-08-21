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

type VpnTunnel struct{}

func VpnTunnelToUnstructured(r *dclService.VpnTunnel) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "alpha",
			Type:    "VpnTunnel",
		},
		Object: make(map[string]interface{}),
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DetailedStatus != nil {
		u.Object["detailedStatus"] = *r.DetailedStatus
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.IkeVersion != nil {
		u.Object["ikeVersion"] = *r.IkeVersion
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	var rLocalTrafficSelector []interface{}
	for _, rLocalTrafficSelectorVal := range r.LocalTrafficSelector {
		rLocalTrafficSelector = append(rLocalTrafficSelector, rLocalTrafficSelectorVal)
	}
	u.Object["localTrafficSelector"] = rLocalTrafficSelector
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.PeerExternalGateway != nil {
		u.Object["peerExternalGateway"] = *r.PeerExternalGateway
	}
	if r.PeerExternalGatewayInterface != nil {
		u.Object["peerExternalGatewayInterface"] = *r.PeerExternalGatewayInterface
	}
	if r.PeerGcpGateway != nil {
		u.Object["peerGcpGateway"] = *r.PeerGcpGateway
	}
	if r.PeerIP != nil {
		u.Object["peerIP"] = *r.PeerIP
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	var rRemoteTrafficSelector []interface{}
	for _, rRemoteTrafficSelectorVal := range r.RemoteTrafficSelector {
		rRemoteTrafficSelector = append(rRemoteTrafficSelector, rRemoteTrafficSelectorVal)
	}
	u.Object["remoteTrafficSelector"] = rRemoteTrafficSelector
	if r.Router != nil {
		u.Object["router"] = *r.Router
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.SharedSecret != nil {
		u.Object["sharedSecret"] = *r.SharedSecret
	}
	if r.SharedSecretHash != nil {
		u.Object["sharedSecretHash"] = *r.SharedSecretHash
	}
	if r.Status != nil {
		u.Object["status"] = string(*r.Status)
	}
	if r.TargetVpnGateway != nil {
		u.Object["targetVpnGateway"] = *r.TargetVpnGateway
	}
	if r.VpnGateway != nil {
		u.Object["vpnGateway"] = *r.VpnGateway
	}
	if r.VpnGatewayInterface != nil {
		u.Object["vpnGatewayInterface"] = *r.VpnGatewayInterface
	}
	return u
}

func UnstructuredToVpnTunnel(u *unstructured.Resource) (*dclService.VpnTunnel, error) {
	r := &dclService.VpnTunnel{}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["detailedStatus"]; ok {
		if s, ok := u.Object["detailedStatus"].(string); ok {
			r.DetailedStatus = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DetailedStatus: expected string")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if i, ok := u.Object["id"].(int64); ok {
			r.Id = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Id: expected int64")
		}
	}
	if _, ok := u.Object["ikeVersion"]; ok {
		if i, ok := u.Object["ikeVersion"].(int64); ok {
			r.IkeVersion = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.IkeVersion: expected int64")
		}
	}
	if _, ok := u.Object["labels"]; ok {
		if rLabels, ok := u.Object["labels"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rLabels {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Labels = m
		} else {
			return nil, fmt.Errorf("r.Labels: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["localTrafficSelector"]; ok {
		if s, ok := u.Object["localTrafficSelector"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.LocalTrafficSelector = append(r.LocalTrafficSelector, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.LocalTrafficSelector: expected []interface{}")
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
	if _, ok := u.Object["peerExternalGateway"]; ok {
		if s, ok := u.Object["peerExternalGateway"].(string); ok {
			r.PeerExternalGateway = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.PeerExternalGateway: expected string")
		}
	}
	if _, ok := u.Object["peerExternalGatewayInterface"]; ok {
		if i, ok := u.Object["peerExternalGatewayInterface"].(int64); ok {
			r.PeerExternalGatewayInterface = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.PeerExternalGatewayInterface: expected int64")
		}
	}
	if _, ok := u.Object["peerGcpGateway"]; ok {
		if s, ok := u.Object["peerGcpGateway"].(string); ok {
			r.PeerGcpGateway = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.PeerGcpGateway: expected string")
		}
	}
	if _, ok := u.Object["peerIP"]; ok {
		if s, ok := u.Object["peerIP"].(string); ok {
			r.PeerIP = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.PeerIP: expected string")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["remoteTrafficSelector"]; ok {
		if s, ok := u.Object["remoteTrafficSelector"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.RemoteTrafficSelector = append(r.RemoteTrafficSelector, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.RemoteTrafficSelector: expected []interface{}")
		}
	}
	if _, ok := u.Object["router"]; ok {
		if s, ok := u.Object["router"].(string); ok {
			r.Router = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Router: expected string")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["sharedSecret"]; ok {
		if s, ok := u.Object["sharedSecret"].(string); ok {
			r.SharedSecret = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SharedSecret: expected string")
		}
	}
	if _, ok := u.Object["sharedSecretHash"]; ok {
		if s, ok := u.Object["sharedSecretHash"].(string); ok {
			r.SharedSecretHash = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SharedSecretHash: expected string")
		}
	}
	if _, ok := u.Object["status"]; ok {
		if s, ok := u.Object["status"].(string); ok {
			r.Status = dclService.VpnTunnelStatusEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Status: expected string")
		}
	}
	if _, ok := u.Object["targetVpnGateway"]; ok {
		if s, ok := u.Object["targetVpnGateway"].(string); ok {
			r.TargetVpnGateway = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.TargetVpnGateway: expected string")
		}
	}
	if _, ok := u.Object["vpnGateway"]; ok {
		if s, ok := u.Object["vpnGateway"].(string); ok {
			r.VpnGateway = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.VpnGateway: expected string")
		}
	}
	if _, ok := u.Object["vpnGatewayInterface"]; ok {
		if i, ok := u.Object["vpnGatewayInterface"].(int64); ok {
			r.VpnGatewayInterface = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.VpnGatewayInterface: expected int64")
		}
	}
	return r, nil
}

func GetVpnTunnel(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToVpnTunnel(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetVpnTunnel(ctx, r)
	if err != nil {
		return nil, err
	}
	return VpnTunnelToUnstructured(r), nil
}

func ListVpnTunnel(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListVpnTunnel(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, VpnTunnelToUnstructured(r))
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

func ApplyVpnTunnel(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToVpnTunnel(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToVpnTunnel(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyVpnTunnel(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return VpnTunnelToUnstructured(r), nil
}

func VpnTunnelHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToVpnTunnel(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToVpnTunnel(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyVpnTunnel(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteVpnTunnel(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToVpnTunnel(u)
	if err != nil {
		return err
	}
	return c.DeleteVpnTunnel(ctx, r)
}

func VpnTunnelID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToVpnTunnel(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *VpnTunnel) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"VpnTunnel",
		"alpha",
	}
}

func (r *VpnTunnel) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VpnTunnel) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VpnTunnel) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *VpnTunnel) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VpnTunnel) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VpnTunnel) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *VpnTunnel) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetVpnTunnel(ctx, config, resource)
}

func (r *VpnTunnel) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyVpnTunnel(ctx, config, resource, opts...)
}

func (r *VpnTunnel) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return VpnTunnelHasDiff(ctx, config, resource, opts...)
}

func (r *VpnTunnel) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteVpnTunnel(ctx, config, resource)
}

func (r *VpnTunnel) ID(resource *unstructured.Resource) (string, error) {
	return VpnTunnelID(resource)
}

func init() {
	unstructured.Register(&VpnTunnel{})
}
