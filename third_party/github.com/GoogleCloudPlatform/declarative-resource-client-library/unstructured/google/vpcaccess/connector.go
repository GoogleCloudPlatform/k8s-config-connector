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
package vpcaccess

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vpcaccess"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Connector struct{}

func ConnectorToUnstructured(r *dclService.Connector) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "vpcaccess",
			Version: "ga",
			Type:    "Connector",
		},
		Object: make(map[string]interface{}),
	}
	var rConnectedProjects []interface{}
	for _, rConnectedProjectsVal := range r.ConnectedProjects {
		rConnectedProjects = append(rConnectedProjects, rConnectedProjectsVal)
	}
	u.Object["connectedProjects"] = rConnectedProjects
	if r.IPCidrRange != nil {
		u.Object["ipCidrRange"] = *r.IPCidrRange
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.MachineType != nil {
		u.Object["machineType"] = *r.MachineType
	}
	if r.MaxInstances != nil {
		u.Object["maxInstances"] = *r.MaxInstances
	}
	if r.MaxThroughput != nil {
		u.Object["maxThroughput"] = *r.MaxThroughput
	}
	if r.MinInstances != nil {
		u.Object["minInstances"] = *r.MinInstances
	}
	if r.MinThroughput != nil {
		u.Object["minThroughput"] = *r.MinThroughput
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Network != nil {
		u.Object["network"] = *r.Network
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.Subnet != nil && r.Subnet != dclService.EmptyConnectorSubnet {
		rSubnet := make(map[string]interface{})
		if r.Subnet.Name != nil {
			rSubnet["name"] = *r.Subnet.Name
		}
		if r.Subnet.ProjectId != nil {
			rSubnet["projectId"] = *r.Subnet.ProjectId
		}
		u.Object["subnet"] = rSubnet
	}
	return u
}

func UnstructuredToConnector(u *unstructured.Resource) (*dclService.Connector, error) {
	r := &dclService.Connector{}
	if _, ok := u.Object["connectedProjects"]; ok {
		if s, ok := u.Object["connectedProjects"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.ConnectedProjects = append(r.ConnectedProjects, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.ConnectedProjects: expected []interface{}")
		}
	}
	if _, ok := u.Object["ipCidrRange"]; ok {
		if s, ok := u.Object["ipCidrRange"].(string); ok {
			r.IPCidrRange = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.IPCidrRange: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["machineType"]; ok {
		if s, ok := u.Object["machineType"].(string); ok {
			r.MachineType = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.MachineType: expected string")
		}
	}
	if _, ok := u.Object["maxInstances"]; ok {
		if i, ok := u.Object["maxInstances"].(int64); ok {
			r.MaxInstances = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.MaxInstances: expected int64")
		}
	}
	if _, ok := u.Object["maxThroughput"]; ok {
		if i, ok := u.Object["maxThroughput"].(int64); ok {
			r.MaxThroughput = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.MaxThroughput: expected int64")
		}
	}
	if _, ok := u.Object["minInstances"]; ok {
		if i, ok := u.Object["minInstances"].(int64); ok {
			r.MinInstances = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.MinInstances: expected int64")
		}
	}
	if _, ok := u.Object["minThroughput"]; ok {
		if i, ok := u.Object["minThroughput"].(int64); ok {
			r.MinThroughput = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.MinThroughput: expected int64")
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
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.ConnectorStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["subnet"]; ok {
		if rSubnet, ok := u.Object["subnet"].(map[string]interface{}); ok {
			r.Subnet = &dclService.ConnectorSubnet{}
			if _, ok := rSubnet["name"]; ok {
				if s, ok := rSubnet["name"].(string); ok {
					r.Subnet.Name = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Subnet.Name: expected string")
				}
			}
			if _, ok := rSubnet["projectId"]; ok {
				if s, ok := rSubnet["projectId"].(string); ok {
					r.Subnet.ProjectId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Subnet.ProjectId: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Subnet: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetConnector(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConnector(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetConnector(ctx, r)
	if err != nil {
		return nil, err
	}
	return ConnectorToUnstructured(r), nil
}

func ListConnector(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListConnector(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, ConnectorToUnstructured(r))
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

func ApplyConnector(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConnector(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToConnector(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyConnector(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return ConnectorToUnstructured(r), nil
}

func ConnectorHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConnector(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToConnector(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyConnector(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteConnector(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToConnector(u)
	if err != nil {
		return err
	}
	return c.DeleteConnector(ctx, r)
}

func ConnectorID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToConnector(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Connector) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"vpcaccess",
		"Connector",
		"ga",
	}
}

func (r *Connector) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connector) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connector) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Connector) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connector) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connector) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Connector) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetConnector(ctx, config, resource)
}

func (r *Connector) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyConnector(ctx, config, resource, opts...)
}

func (r *Connector) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return ConnectorHasDiff(ctx, config, resource, opts...)
}

func (r *Connector) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteConnector(ctx, config, resource)
}

func (r *Connector) ID(resource *unstructured.Resource) (string, error) {
	return ConnectorID(resource)
}

func init() {
	unstructured.Register(&Connector{})
}
