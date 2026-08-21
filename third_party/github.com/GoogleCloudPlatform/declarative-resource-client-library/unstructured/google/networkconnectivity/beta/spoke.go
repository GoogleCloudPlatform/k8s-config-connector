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
package networkconnectivity

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkconnectivity/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Spoke struct{}

func SpokeToUnstructured(r *dclService.Spoke) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networkconnectivity",
			Version: "beta",
			Type:    "Spoke",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Hub != nil {
		u.Object["hub"] = *r.Hub
	}
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.LinkedInterconnectAttachments != nil && r.LinkedInterconnectAttachments != dclService.EmptySpokeLinkedInterconnectAttachments {
		rLinkedInterconnectAttachments := make(map[string]interface{})
		if r.LinkedInterconnectAttachments.SiteToSiteDataTransfer != nil {
			rLinkedInterconnectAttachments["siteToSiteDataTransfer"] = *r.LinkedInterconnectAttachments.SiteToSiteDataTransfer
		}
		var rLinkedInterconnectAttachmentsUris []interface{}
		for _, rLinkedInterconnectAttachmentsUrisVal := range r.LinkedInterconnectAttachments.Uris {
			rLinkedInterconnectAttachmentsUris = append(rLinkedInterconnectAttachmentsUris, rLinkedInterconnectAttachmentsUrisVal)
		}
		rLinkedInterconnectAttachments["uris"] = rLinkedInterconnectAttachmentsUris
		u.Object["linkedInterconnectAttachments"] = rLinkedInterconnectAttachments
	}
	if r.LinkedRouterApplianceInstances != nil && r.LinkedRouterApplianceInstances != dclService.EmptySpokeLinkedRouterApplianceInstances {
		rLinkedRouterApplianceInstances := make(map[string]interface{})
		var rLinkedRouterApplianceInstancesInstances []interface{}
		for _, rLinkedRouterApplianceInstancesInstancesVal := range r.LinkedRouterApplianceInstances.Instances {
			rLinkedRouterApplianceInstancesInstancesObject := make(map[string]interface{})
			if rLinkedRouterApplianceInstancesInstancesVal.IPAddress != nil {
				rLinkedRouterApplianceInstancesInstancesObject["ipAddress"] = *rLinkedRouterApplianceInstancesInstancesVal.IPAddress
			}
			if rLinkedRouterApplianceInstancesInstancesVal.VirtualMachine != nil {
				rLinkedRouterApplianceInstancesInstancesObject["virtualMachine"] = *rLinkedRouterApplianceInstancesInstancesVal.VirtualMachine
			}
			rLinkedRouterApplianceInstancesInstances = append(rLinkedRouterApplianceInstancesInstances, rLinkedRouterApplianceInstancesInstancesObject)
		}
		rLinkedRouterApplianceInstances["instances"] = rLinkedRouterApplianceInstancesInstances
		if r.LinkedRouterApplianceInstances.SiteToSiteDataTransfer != nil {
			rLinkedRouterApplianceInstances["siteToSiteDataTransfer"] = *r.LinkedRouterApplianceInstances.SiteToSiteDataTransfer
		}
		u.Object["linkedRouterApplianceInstances"] = rLinkedRouterApplianceInstances
	}
	if r.LinkedVPCNetwork != nil && r.LinkedVPCNetwork != dclService.EmptySpokeLinkedVPCNetwork {
		rLinkedVPCNetwork := make(map[string]interface{})
		var rLinkedVPCNetworkExcludeExportRanges []interface{}
		for _, rLinkedVPCNetworkExcludeExportRangesVal := range r.LinkedVPCNetwork.ExcludeExportRanges {
			rLinkedVPCNetworkExcludeExportRanges = append(rLinkedVPCNetworkExcludeExportRanges, rLinkedVPCNetworkExcludeExportRangesVal)
		}
		rLinkedVPCNetwork["excludeExportRanges"] = rLinkedVPCNetworkExcludeExportRanges
		if r.LinkedVPCNetwork.Uri != nil {
			rLinkedVPCNetwork["uri"] = *r.LinkedVPCNetwork.Uri
		}
		u.Object["linkedVPCNetwork"] = rLinkedVPCNetwork
	}
	if r.LinkedVpnTunnels != nil && r.LinkedVpnTunnels != dclService.EmptySpokeLinkedVpnTunnels {
		rLinkedVpnTunnels := make(map[string]interface{})
		if r.LinkedVpnTunnels.SiteToSiteDataTransfer != nil {
			rLinkedVpnTunnels["siteToSiteDataTransfer"] = *r.LinkedVpnTunnels.SiteToSiteDataTransfer
		}
		var rLinkedVpnTunnelsUris []interface{}
		for _, rLinkedVpnTunnelsUrisVal := range r.LinkedVpnTunnels.Uris {
			rLinkedVpnTunnelsUris = append(rLinkedVpnTunnelsUris, rLinkedVpnTunnelsUrisVal)
		}
		rLinkedVpnTunnels["uris"] = rLinkedVpnTunnelsUris
		u.Object["linkedVpnTunnels"] = rLinkedVpnTunnels
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
	if r.UniqueId != nil {
		u.Object["uniqueId"] = *r.UniqueId
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToSpoke(u *unstructured.Resource) (*dclService.Spoke, error) {
	r := &dclService.Spoke{}
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
	if _, ok := u.Object["hub"]; ok {
		if s, ok := u.Object["hub"].(string); ok {
			r.Hub = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Hub: expected string")
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
	if _, ok := u.Object["linkedInterconnectAttachments"]; ok {
		if rLinkedInterconnectAttachments, ok := u.Object["linkedInterconnectAttachments"].(map[string]interface{}); ok {
			r.LinkedInterconnectAttachments = &dclService.SpokeLinkedInterconnectAttachments{}
			if _, ok := rLinkedInterconnectAttachments["siteToSiteDataTransfer"]; ok {
				if b, ok := rLinkedInterconnectAttachments["siteToSiteDataTransfer"].(bool); ok {
					r.LinkedInterconnectAttachments.SiteToSiteDataTransfer = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.LinkedInterconnectAttachments.SiteToSiteDataTransfer: expected bool")
				}
			}
			if _, ok := rLinkedInterconnectAttachments["uris"]; ok {
				if s, ok := rLinkedInterconnectAttachments["uris"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.LinkedInterconnectAttachments.Uris = append(r.LinkedInterconnectAttachments.Uris, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.LinkedInterconnectAttachments.Uris: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LinkedInterconnectAttachments: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["linkedRouterApplianceInstances"]; ok {
		if rLinkedRouterApplianceInstances, ok := u.Object["linkedRouterApplianceInstances"].(map[string]interface{}); ok {
			r.LinkedRouterApplianceInstances = &dclService.SpokeLinkedRouterApplianceInstances{}
			if _, ok := rLinkedRouterApplianceInstances["instances"]; ok {
				if s, ok := rLinkedRouterApplianceInstances["instances"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rLinkedRouterApplianceInstancesInstances dclService.SpokeLinkedRouterApplianceInstancesInstances
							if _, ok := objval["ipAddress"]; ok {
								if s, ok := objval["ipAddress"].(string); ok {
									rLinkedRouterApplianceInstancesInstances.IPAddress = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rLinkedRouterApplianceInstancesInstances.IPAddress: expected string")
								}
							}
							if _, ok := objval["virtualMachine"]; ok {
								if s, ok := objval["virtualMachine"].(string); ok {
									rLinkedRouterApplianceInstancesInstances.VirtualMachine = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rLinkedRouterApplianceInstancesInstances.VirtualMachine: expected string")
								}
							}
							r.LinkedRouterApplianceInstances.Instances = append(r.LinkedRouterApplianceInstances.Instances, rLinkedRouterApplianceInstancesInstances)
						}
					}
				} else {
					return nil, fmt.Errorf("r.LinkedRouterApplianceInstances.Instances: expected []interface{}")
				}
			}
			if _, ok := rLinkedRouterApplianceInstances["siteToSiteDataTransfer"]; ok {
				if b, ok := rLinkedRouterApplianceInstances["siteToSiteDataTransfer"].(bool); ok {
					r.LinkedRouterApplianceInstances.SiteToSiteDataTransfer = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.LinkedRouterApplianceInstances.SiteToSiteDataTransfer: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LinkedRouterApplianceInstances: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["linkedVPCNetwork"]; ok {
		if rLinkedVPCNetwork, ok := u.Object["linkedVPCNetwork"].(map[string]interface{}); ok {
			r.LinkedVPCNetwork = &dclService.SpokeLinkedVPCNetwork{}
			if _, ok := rLinkedVPCNetwork["excludeExportRanges"]; ok {
				if s, ok := rLinkedVPCNetwork["excludeExportRanges"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.LinkedVPCNetwork.ExcludeExportRanges = append(r.LinkedVPCNetwork.ExcludeExportRanges, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.LinkedVPCNetwork.ExcludeExportRanges: expected []interface{}")
				}
			}
			if _, ok := rLinkedVPCNetwork["uri"]; ok {
				if s, ok := rLinkedVPCNetwork["uri"].(string); ok {
					r.LinkedVPCNetwork.Uri = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.LinkedVPCNetwork.Uri: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LinkedVPCNetwork: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["linkedVpnTunnels"]; ok {
		if rLinkedVpnTunnels, ok := u.Object["linkedVpnTunnels"].(map[string]interface{}); ok {
			r.LinkedVpnTunnels = &dclService.SpokeLinkedVpnTunnels{}
			if _, ok := rLinkedVpnTunnels["siteToSiteDataTransfer"]; ok {
				if b, ok := rLinkedVpnTunnels["siteToSiteDataTransfer"].(bool); ok {
					r.LinkedVpnTunnels.SiteToSiteDataTransfer = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.LinkedVpnTunnels.SiteToSiteDataTransfer: expected bool")
				}
			}
			if _, ok := rLinkedVpnTunnels["uris"]; ok {
				if s, ok := rLinkedVpnTunnels["uris"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.LinkedVpnTunnels.Uris = append(r.LinkedVpnTunnels.Uris, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.LinkedVpnTunnels.Uris: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.LinkedVpnTunnels: expected map[string]interface{}")
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
			r.State = dclService.SpokeStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["uniqueId"]; ok {
		if s, ok := u.Object["uniqueId"].(string); ok {
			r.UniqueId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UniqueId: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	return r, nil
}

func GetSpoke(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToSpoke(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetSpoke(ctx, r)
	if err != nil {
		return nil, err
	}
	return SpokeToUnstructured(r), nil
}

func ListSpoke(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListSpoke(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, SpokeToUnstructured(r))
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

func ApplySpoke(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToSpoke(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToSpoke(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplySpoke(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return SpokeToUnstructured(r), nil
}

func SpokeHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToSpoke(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToSpoke(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplySpoke(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteSpoke(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToSpoke(u)
	if err != nil {
		return err
	}
	return c.DeleteSpoke(ctx, r)
}

func SpokeID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToSpoke(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Spoke) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networkconnectivity",
		"Spoke",
		"beta",
	}
}

func (r *Spoke) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Spoke) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Spoke) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Spoke) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Spoke) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Spoke) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Spoke) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetSpoke(ctx, config, resource)
}

func (r *Spoke) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplySpoke(ctx, config, resource, opts...)
}

func (r *Spoke) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return SpokeHasDiff(ctx, config, resource, opts...)
}

func (r *Spoke) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteSpoke(ctx, config, resource)
}

func (r *Spoke) ID(resource *unstructured.Resource) (string, error) {
	return SpokeID(resource)
}

func init() {
	unstructured.Register(&Spoke{})
}
