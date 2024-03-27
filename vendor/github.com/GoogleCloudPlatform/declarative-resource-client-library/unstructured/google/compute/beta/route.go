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
package compute

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Route struct{}

func RouteToUnstructured(r *dclService.Route) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "beta",
			Type:    "Route",
		},
		Object: make(map[string]interface{}),
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DestRange != nil {
		u.Object["destRange"] = *r.DestRange
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Network != nil {
		u.Object["network"] = *r.Network
	}
	if r.NextHopGateway != nil {
		u.Object["nextHopGateway"] = *r.NextHopGateway
	}
	if r.NextHopIP != nil {
		u.Object["nextHopIP"] = *r.NextHopIP
	}
	if r.NextHopIlb != nil {
		u.Object["nextHopIlb"] = *r.NextHopIlb
	}
	if r.NextHopInstance != nil {
		u.Object["nextHopInstance"] = *r.NextHopInstance
	}
	if r.NextHopNetwork != nil {
		u.Object["nextHopNetwork"] = *r.NextHopNetwork
	}
	if r.NextHopPeering != nil {
		u.Object["nextHopPeering"] = *r.NextHopPeering
	}
	if r.NextHopVpnTunnel != nil {
		u.Object["nextHopVpnTunnel"] = *r.NextHopVpnTunnel
	}
	if r.Priority != nil {
		u.Object["priority"] = *r.Priority
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	var rTag []interface{}
	for _, rTagVal := range r.Tag {
		rTag = append(rTag, rTagVal)
	}
	u.Object["tag"] = rTag
	var rWarning []interface{}
	for _, rWarningVal := range r.Warning {
		rWarningObject := make(map[string]interface{})
		if rWarningVal.Code != nil {
			rWarningObject["code"] = string(*rWarningVal.Code)
		}
		if rWarningVal.Data != nil {
			rWarningValData := make(map[string]interface{})
			for k, v := range rWarningVal.Data {
				rWarningValData[k] = v
			}
			rWarningObject["data"] = rWarningValData
		}
		if rWarningVal.Message != nil {
			rWarningObject["message"] = *rWarningVal.Message
		}
		rWarning = append(rWarning, rWarningObject)
	}
	u.Object["warning"] = rWarning
	return u
}

func UnstructuredToRoute(u *unstructured.Resource) (*dclService.Route, error) {
	r := &dclService.Route{}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["destRange"]; ok {
		if s, ok := u.Object["destRange"].(string); ok {
			r.DestRange = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DestRange: expected string")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if i, ok := u.Object["id"].(int64); ok {
			r.Id = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Id: expected int64")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["network"]; ok {
		if s, ok := u.Object["network"].(string); ok {
			r.Network = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Network: expected string")
		}
	}
	if _, ok := u.Object["nextHopGateway"]; ok {
		if s, ok := u.Object["nextHopGateway"].(string); ok {
			r.NextHopGateway = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NextHopGateway: expected string")
		}
	}
	if _, ok := u.Object["nextHopIP"]; ok {
		if s, ok := u.Object["nextHopIP"].(string); ok {
			r.NextHopIP = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NextHopIP: expected string")
		}
	}
	if _, ok := u.Object["nextHopIlb"]; ok {
		if s, ok := u.Object["nextHopIlb"].(string); ok {
			r.NextHopIlb = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NextHopIlb: expected string")
		}
	}
	if _, ok := u.Object["nextHopInstance"]; ok {
		if s, ok := u.Object["nextHopInstance"].(string); ok {
			r.NextHopInstance = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NextHopInstance: expected string")
		}
	}
	if _, ok := u.Object["nextHopNetwork"]; ok {
		if s, ok := u.Object["nextHopNetwork"].(string); ok {
			r.NextHopNetwork = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NextHopNetwork: expected string")
		}
	}
	if _, ok := u.Object["nextHopPeering"]; ok {
		if s, ok := u.Object["nextHopPeering"].(string); ok {
			r.NextHopPeering = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NextHopPeering: expected string")
		}
	}
	if _, ok := u.Object["nextHopVpnTunnel"]; ok {
		if s, ok := u.Object["nextHopVpnTunnel"].(string); ok {
			r.NextHopVpnTunnel = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.NextHopVpnTunnel: expected string")
		}
	}
	if _, ok := u.Object["priority"]; ok {
		if i, ok := u.Object["priority"].(int64); ok {
			r.Priority = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Priority: expected int64")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	if _, ok := u.Object["tag"]; ok {
		if s, ok := u.Object["tag"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.Tag = append(r.Tag, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Tag: expected []interface{}")
		}
	}
	if _, ok := u.Object["warning"]; ok {
		if s, ok := u.Object["warning"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rWarning dclService.RouteWarning
					if _, ok := objval["code"]; ok {
						if s, ok := objval["code"].(string); ok {
							rWarning.Code = dclService.RouteWarningCodeEnumRef(s)
						} else {
							return nil, fmt.Errorf("rWarning.Code: expected string")
						}
					}
					if _, ok := objval["data"]; ok {
						if rWarningData, ok := objval["data"].(map[string]interface{}); ok {
							m := make(map[string]string)
							for k, v := range rWarningData {
								if s, ok := v.(string); ok {
									m[k] = s
								}
							}
							rWarning.Data = m
						} else {
							return nil, fmt.Errorf("rWarning.Data: expected map[string]interface{}")
						}
					}
					if _, ok := objval["message"]; ok {
						if s, ok := objval["message"].(string); ok {
							rWarning.Message = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rWarning.Message: expected string")
						}
					}
					r.Warning = append(r.Warning, rWarning)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Warning: expected []interface{}")
		}
	}
	return r, nil
}

func GetRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRoute(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetRoute(ctx, r)
	if err != nil {
		return nil, err
	}
	return RouteToUnstructured(r), nil
}

func ListRoute(ctx context.Context, config *dcl.Config, project string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListRoute(ctx, project)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, RouteToUnstructured(r))
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

func ApplyRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRoute(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRoute(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyRoute(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return RouteToUnstructured(r), nil
}

func RouteHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRoute(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToRoute(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyRoute(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToRoute(u)
	if err != nil {
		return err
	}
	return c.DeleteRoute(ctx, r)
}

func RouteID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToRoute(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Route) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"Route",
		"beta",
	}
}

func (r *Route) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Route) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Route) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Route) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Route) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Route) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Route) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetRoute(ctx, config, resource)
}

func (r *Route) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyRoute(ctx, config, resource, opts...)
}

func (r *Route) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return RouteHasDiff(ctx, config, resource, opts...)
}

func (r *Route) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteRoute(ctx, config, resource)
}

func (r *Route) ID(resource *unstructured.Resource) (string, error) {
	return RouteID(resource)
}

func init() {
	unstructured.Register(&Route{})
}
