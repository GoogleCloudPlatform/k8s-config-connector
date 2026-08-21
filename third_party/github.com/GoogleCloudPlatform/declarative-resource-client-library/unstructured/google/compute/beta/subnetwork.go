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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Subnetwork struct{}

func SubnetworkToUnstructured(r *dclService.Subnetwork) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "beta",
			Type:    "Subnetwork",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreationTimestamp != nil {
		u.Object["creationTimestamp"] = *r.CreationTimestamp
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.EnableFlowLogs != nil {
		u.Object["enableFlowLogs"] = *r.EnableFlowLogs
	}
	if r.Fingerprint != nil {
		u.Object["fingerprint"] = *r.Fingerprint
	}
	if r.GatewayAddress != nil {
		u.Object["gatewayAddress"] = *r.GatewayAddress
	}
	if r.IPCidrRange != nil {
		u.Object["ipCidrRange"] = *r.IPCidrRange
	}
	if r.LogConfig != nil && r.LogConfig != dclService.EmptySubnetworkLogConfig {
		rLogConfig := make(map[string]interface{})
		if r.LogConfig.AggregationInterval != nil {
			rLogConfig["aggregationInterval"] = string(*r.LogConfig.AggregationInterval)
		}
		if r.LogConfig.FlowSampling != nil {
			rLogConfig["flowSampling"] = *r.LogConfig.FlowSampling
		}
		if r.LogConfig.Metadata != nil {
			rLogConfig["metadata"] = string(*r.LogConfig.Metadata)
		}
		u.Object["logConfig"] = rLogConfig
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Network != nil {
		u.Object["network"] = *r.Network
	}
	if r.PrivateIPGoogleAccess != nil {
		u.Object["privateIPGoogleAccess"] = *r.PrivateIPGoogleAccess
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Purpose != nil {
		u.Object["purpose"] = string(*r.Purpose)
	}
	if r.Region != nil {
		u.Object["region"] = *r.Region
	}
	if r.Role != nil {
		u.Object["role"] = string(*r.Role)
	}
	var rSecondaryIPRanges []interface{}
	for _, rSecondaryIPRangesVal := range r.SecondaryIPRanges {
		rSecondaryIPRangesObject := make(map[string]interface{})
		if rSecondaryIPRangesVal.IPCidrRange != nil {
			rSecondaryIPRangesObject["ipCidrRange"] = *rSecondaryIPRangesVal.IPCidrRange
		}
		if rSecondaryIPRangesVal.RangeName != nil {
			rSecondaryIPRangesObject["rangeName"] = *rSecondaryIPRangesVal.RangeName
		}
		rSecondaryIPRanges = append(rSecondaryIPRanges, rSecondaryIPRangesObject)
	}
	u.Object["secondaryIPRanges"] = rSecondaryIPRanges
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	return u
}

func UnstructuredToSubnetwork(u *unstructured.Resource) (*dclService.Subnetwork, error) {
	r := &dclService.Subnetwork{}
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
	if _, ok := u.Object["enableFlowLogs"]; ok {
		if b, ok := u.Object["enableFlowLogs"].(bool); ok {
			r.EnableFlowLogs = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableFlowLogs: expected bool")
		}
	}
	if _, ok := u.Object["fingerprint"]; ok {
		if s, ok := u.Object["fingerprint"].(string); ok {
			r.Fingerprint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Fingerprint: expected string")
		}
	}
	if _, ok := u.Object["gatewayAddress"]; ok {
		if s, ok := u.Object["gatewayAddress"].(string); ok {
			r.GatewayAddress = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.GatewayAddress: expected string")
		}
	}
	if _, ok := u.Object["ipCidrRange"]; ok {
		if s, ok := u.Object["ipCidrRange"].(string); ok {
			r.IPCidrRange = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.IPCidrRange: expected string")
		}
	}
	if _, ok := u.Object["logConfig"]; ok {
		if rLogConfig, ok := u.Object["logConfig"].(map[string]interface{}); ok {
			r.LogConfig = &dclService.SubnetworkLogConfig{}
			if _, ok := rLogConfig["aggregationInterval"]; ok {
				if s, ok := rLogConfig["aggregationInterval"].(string); ok {
					r.LogConfig.AggregationInterval = dclService.SubnetworkLogConfigAggregationIntervalEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.LogConfig.AggregationInterval: expected string")
				}
			}
			if _, ok := rLogConfig["flowSampling"]; ok {
				if f, ok := rLogConfig["flowSampling"].(float64); ok {
					r.LogConfig.FlowSampling = dcl.Float64(f)
				} else {
					return nil, fmt.Errorf("r.LogConfig.FlowSampling: expected float64")
				}
			}
			if _, ok := rLogConfig["metadata"]; ok {
				if s, ok := rLogConfig["metadata"].(string); ok {
					r.LogConfig.Metadata = dclService.SubnetworkLogConfigMetadataEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.LogConfig.Metadata: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LogConfig: expected map[string]interface{}")
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
	if _, ok := u.Object["privateIPGoogleAccess"]; ok {
		if b, ok := u.Object["privateIPGoogleAccess"].(bool); ok {
			r.PrivateIPGoogleAccess = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.PrivateIPGoogleAccess: expected bool")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["purpose"]; ok {
		if s, ok := u.Object["purpose"].(string); ok {
			r.Purpose = dclService.SubnetworkPurposeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Purpose: expected string")
		}
	}
	if _, ok := u.Object["region"]; ok {
		if s, ok := u.Object["region"].(string); ok {
			r.Region = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Region: expected string")
		}
	}
	if _, ok := u.Object["role"]; ok {
		if s, ok := u.Object["role"].(string); ok {
			r.Role = dclService.SubnetworkRoleEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Role: expected string")
		}
	}
	if _, ok := u.Object["secondaryIPRanges"]; ok {
		if s, ok := u.Object["secondaryIPRanges"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rSecondaryIPRanges dclService.SubnetworkSecondaryIPRanges
					if _, ok := objval["ipCidrRange"]; ok {
						if s, ok := objval["ipCidrRange"].(string); ok {
							rSecondaryIPRanges.IPCidrRange = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rSecondaryIPRanges.IPCidrRange: expected string")
						}
					}
					if _, ok := objval["rangeName"]; ok {
						if s, ok := objval["rangeName"].(string); ok {
							rSecondaryIPRanges.RangeName = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rSecondaryIPRanges.RangeName: expected string")
						}
					}
					r.SecondaryIPRanges = append(r.SecondaryIPRanges, rSecondaryIPRanges)
				}
			}
		} else {
			return nil, fmt.Errorf("r.SecondaryIPRanges: expected []interface{}")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
		}
	}
	return r, nil
}

func GetSubnetwork(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToSubnetwork(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetSubnetwork(ctx, r)
	if err != nil {
		return nil, err
	}
	return SubnetworkToUnstructured(r), nil
}

func ListSubnetwork(ctx context.Context, config *dcl.Config, project string, region string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListSubnetwork(ctx, project, region)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, SubnetworkToUnstructured(r))
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

func ApplySubnetwork(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToSubnetwork(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToSubnetwork(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplySubnetwork(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return SubnetworkToUnstructured(r), nil
}

func SubnetworkHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToSubnetwork(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToSubnetwork(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplySubnetwork(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteSubnetwork(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToSubnetwork(u)
	if err != nil {
		return err
	}
	return c.DeleteSubnetwork(ctx, r)
}

func SubnetworkID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToSubnetwork(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Subnetwork) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"Subnetwork",
		"beta",
	}
}

func (r *Subnetwork) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Subnetwork) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Subnetwork) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Subnetwork) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Subnetwork) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Subnetwork) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Subnetwork) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetSubnetwork(ctx, config, resource)
}

func (r *Subnetwork) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplySubnetwork(ctx, config, resource, opts...)
}

func (r *Subnetwork) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return SubnetworkHasDiff(ctx, config, resource, opts...)
}

func (r *Subnetwork) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteSubnetwork(ctx, config, resource)
}

func (r *Subnetwork) ID(resource *unstructured.Resource) (string, error) {
	return SubnetworkID(resource)
}

func init() {
	unstructured.Register(&Subnetwork{})
}
