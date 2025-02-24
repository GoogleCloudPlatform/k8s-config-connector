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

type PacketMirroring struct{}

func PacketMirroringToUnstructured(r *dclService.PacketMirroring) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "compute",
			Version: "alpha",
			Type:    "PacketMirroring",
		},
		Object: make(map[string]interface{}),
	}
	if r.CollectorIlb != nil && r.CollectorIlb != dclService.EmptyPacketMirroringCollectorIlb {
		rCollectorIlb := make(map[string]interface{})
		if r.CollectorIlb.CanonicalUrl != nil {
			rCollectorIlb["canonicalUrl"] = *r.CollectorIlb.CanonicalUrl
		}
		if r.CollectorIlb.Url != nil {
			rCollectorIlb["url"] = *r.CollectorIlb.Url
		}
		u.Object["collectorIlb"] = rCollectorIlb
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.Enable != nil {
		u.Object["enable"] = string(*r.Enable)
	}
	if r.Filter != nil && r.Filter != dclService.EmptyPacketMirroringFilter {
		rFilter := make(map[string]interface{})
		var rFilterCidrRanges []interface{}
		for _, rFilterCidrRangesVal := range r.Filter.CidrRanges {
			rFilterCidrRanges = append(rFilterCidrRanges, rFilterCidrRangesVal)
		}
		rFilter["cidrRanges"] = rFilterCidrRanges
		if r.Filter.Direction != nil {
			rFilter["direction"] = string(*r.Filter.Direction)
		}
		var rFilterIPProtocols []interface{}
		for _, rFilterIPProtocolsVal := range r.Filter.IPProtocols {
			rFilterIPProtocols = append(rFilterIPProtocols, rFilterIPProtocolsVal)
		}
		rFilter["ipProtocols"] = rFilterIPProtocols
		u.Object["filter"] = rFilter
	}
	if r.Id != nil {
		u.Object["id"] = *r.Id
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.MirroredResources != nil && r.MirroredResources != dclService.EmptyPacketMirroringMirroredResources {
		rMirroredResources := make(map[string]interface{})
		var rMirroredResourcesInstances []interface{}
		for _, rMirroredResourcesInstancesVal := range r.MirroredResources.Instances {
			rMirroredResourcesInstancesObject := make(map[string]interface{})
			if rMirroredResourcesInstancesVal.CanonicalUrl != nil {
				rMirroredResourcesInstancesObject["canonicalUrl"] = *rMirroredResourcesInstancesVal.CanonicalUrl
			}
			if rMirroredResourcesInstancesVal.Url != nil {
				rMirroredResourcesInstancesObject["url"] = *rMirroredResourcesInstancesVal.Url
			}
			rMirroredResourcesInstances = append(rMirroredResourcesInstances, rMirroredResourcesInstancesObject)
		}
		rMirroredResources["instances"] = rMirroredResourcesInstances
		var rMirroredResourcesSubnetworks []interface{}
		for _, rMirroredResourcesSubnetworksVal := range r.MirroredResources.Subnetworks {
			rMirroredResourcesSubnetworksObject := make(map[string]interface{})
			if rMirroredResourcesSubnetworksVal.CanonicalUrl != nil {
				rMirroredResourcesSubnetworksObject["canonicalUrl"] = *rMirroredResourcesSubnetworksVal.CanonicalUrl
			}
			if rMirroredResourcesSubnetworksVal.Url != nil {
				rMirroredResourcesSubnetworksObject["url"] = *rMirroredResourcesSubnetworksVal.Url
			}
			rMirroredResourcesSubnetworks = append(rMirroredResourcesSubnetworks, rMirroredResourcesSubnetworksObject)
		}
		rMirroredResources["subnetworks"] = rMirroredResourcesSubnetworks
		var rMirroredResourcesTags []interface{}
		for _, rMirroredResourcesTagsVal := range r.MirroredResources.Tags {
			rMirroredResourcesTags = append(rMirroredResourcesTags, rMirroredResourcesTagsVal)
		}
		rMirroredResources["tags"] = rMirroredResourcesTags
		u.Object["mirroredResources"] = rMirroredResources
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Network != nil && r.Network != dclService.EmptyPacketMirroringNetwork {
		rNetwork := make(map[string]interface{})
		if r.Network.CanonicalUrl != nil {
			rNetwork["canonicalUrl"] = *r.Network.CanonicalUrl
		}
		if r.Network.Url != nil {
			rNetwork["url"] = *r.Network.Url
		}
		u.Object["network"] = rNetwork
	}
	if r.Priority != nil {
		u.Object["priority"] = *r.Priority
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Region != nil {
		u.Object["region"] = *r.Region
	}
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	return u
}

func UnstructuredToPacketMirroring(u *unstructured.Resource) (*dclService.PacketMirroring, error) {
	r := &dclService.PacketMirroring{}
	if _, ok := u.Object["collectorIlb"]; ok {
		if rCollectorIlb, ok := u.Object["collectorIlb"].(map[string]interface{}); ok {
			r.CollectorIlb = &dclService.PacketMirroringCollectorIlb{}
			if _, ok := rCollectorIlb["canonicalUrl"]; ok {
				if s, ok := rCollectorIlb["canonicalUrl"].(string); ok {
					r.CollectorIlb.CanonicalUrl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.CollectorIlb.CanonicalUrl: expected string")
				}
			}
			if _, ok := rCollectorIlb["url"]; ok {
				if s, ok := rCollectorIlb["url"].(string); ok {
					r.CollectorIlb.Url = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.CollectorIlb.Url: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.CollectorIlb: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["enable"]; ok {
		if s, ok := u.Object["enable"].(string); ok {
			r.Enable = dclService.PacketMirroringEnableEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Enable: expected string")
		}
	}
	if _, ok := u.Object["filter"]; ok {
		if rFilter, ok := u.Object["filter"].(map[string]interface{}); ok {
			r.Filter = &dclService.PacketMirroringFilter{}
			if _, ok := rFilter["cidrRanges"]; ok {
				if s, ok := rFilter["cidrRanges"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Filter.CidrRanges = append(r.Filter.CidrRanges, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Filter.CidrRanges: expected []interface{}")
				}
			}
			if _, ok := rFilter["direction"]; ok {
				if s, ok := rFilter["direction"].(string); ok {
					r.Filter.Direction = dclService.PacketMirroringFilterDirectionEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Filter.Direction: expected string")
				}
			}
			if _, ok := rFilter["ipProtocols"]; ok {
				if s, ok := rFilter["ipProtocols"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Filter.IPProtocols = append(r.Filter.IPProtocols, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Filter.IPProtocols: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Filter: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["id"]; ok {
		if i, ok := u.Object["id"].(int64); ok {
			r.Id = dcl.Int64(i)
		} else {
			return nil, fmt.Errorf("r.Id: expected int64")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["mirroredResources"]; ok {
		if rMirroredResources, ok := u.Object["mirroredResources"].(map[string]interface{}); ok {
			r.MirroredResources = &dclService.PacketMirroringMirroredResources{}
			if _, ok := rMirroredResources["instances"]; ok {
				if s, ok := rMirroredResources["instances"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rMirroredResourcesInstances dclService.PacketMirroringMirroredResourcesInstances
							if _, ok := objval["canonicalUrl"]; ok {
								if s, ok := objval["canonicalUrl"].(string); ok {
									rMirroredResourcesInstances.CanonicalUrl = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rMirroredResourcesInstances.CanonicalUrl: expected string")
								}
							}
							if _, ok := objval["url"]; ok {
								if s, ok := objval["url"].(string); ok {
									rMirroredResourcesInstances.Url = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rMirroredResourcesInstances.Url: expected string")
								}
							}
							r.MirroredResources.Instances = append(r.MirroredResources.Instances, rMirroredResourcesInstances)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MirroredResources.Instances: expected []interface{}")
				}
			}
			if _, ok := rMirroredResources["subnetworks"]; ok {
				if s, ok := rMirroredResources["subnetworks"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rMirroredResourcesSubnetworks dclService.PacketMirroringMirroredResourcesSubnetworks
							if _, ok := objval["canonicalUrl"]; ok {
								if s, ok := objval["canonicalUrl"].(string); ok {
									rMirroredResourcesSubnetworks.CanonicalUrl = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rMirroredResourcesSubnetworks.CanonicalUrl: expected string")
								}
							}
							if _, ok := objval["url"]; ok {
								if s, ok := objval["url"].(string); ok {
									rMirroredResourcesSubnetworks.Url = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rMirroredResourcesSubnetworks.Url: expected string")
								}
							}
							r.MirroredResources.Subnetworks = append(r.MirroredResources.Subnetworks, rMirroredResourcesSubnetworks)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MirroredResources.Subnetworks: expected []interface{}")
				}
			}
			if _, ok := rMirroredResources["tags"]; ok {
				if s, ok := rMirroredResources["tags"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.MirroredResources.Tags = append(r.MirroredResources.Tags, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.MirroredResources.Tags: expected []interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MirroredResources: expected map[string]interface{}")
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
		if rNetwork, ok := u.Object["network"].(map[string]interface{}); ok {
			r.Network = &dclService.PacketMirroringNetwork{}
			if _, ok := rNetwork["canonicalUrl"]; ok {
				if s, ok := rNetwork["canonicalUrl"].(string); ok {
					r.Network.CanonicalUrl = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Network.CanonicalUrl: expected string")
				}
			}
			if _, ok := rNetwork["url"]; ok {
				if s, ok := rNetwork["url"].(string); ok {
					r.Network.Url = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Network.Url: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Network: expected map[string]interface{}")
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
	if _, ok := u.Object["region"]; ok {
		if s, ok := u.Object["region"].(string); ok {
			r.Region = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Region: expected string")
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

func GetPacketMirroring(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPacketMirroring(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetPacketMirroring(ctx, r)
	if err != nil {
		return nil, err
	}
	return PacketMirroringToUnstructured(r), nil
}

func ListPacketMirroring(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListPacketMirroring(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, PacketMirroringToUnstructured(r))
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

func ApplyPacketMirroring(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPacketMirroring(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPacketMirroring(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyPacketMirroring(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return PacketMirroringToUnstructured(r), nil
}

func PacketMirroringHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPacketMirroring(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPacketMirroring(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyPacketMirroring(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeletePacketMirroring(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPacketMirroring(u)
	if err != nil {
		return err
	}
	return c.DeletePacketMirroring(ctx, r)
}

func PacketMirroringID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToPacketMirroring(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *PacketMirroring) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"compute",
		"PacketMirroring",
		"alpha",
	}
}

func (r *PacketMirroring) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PacketMirroring) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PacketMirroring) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *PacketMirroring) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PacketMirroring) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PacketMirroring) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *PacketMirroring) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPacketMirroring(ctx, config, resource)
}

func (r *PacketMirroring) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyPacketMirroring(ctx, config, resource, opts...)
}

func (r *PacketMirroring) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return PacketMirroringHasDiff(ctx, config, resource, opts...)
}

func (r *PacketMirroring) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeletePacketMirroring(ctx, config, resource)
}

func (r *PacketMirroring) ID(resource *unstructured.Resource) (string, error) {
	return PacketMirroringID(resource)
}

func init() {
	unstructured.Register(&PacketMirroring{})
}
