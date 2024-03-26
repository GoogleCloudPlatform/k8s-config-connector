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
package filestore

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Instance struct{}

func InstanceToUnstructured(r *dclService.Instance) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "filestore",
			Version: "beta",
			Type:    "Instance",
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
	var rFileShares []interface{}
	for _, rFileSharesVal := range r.FileShares {
		rFileSharesObject := make(map[string]interface{})
		if rFileSharesVal.CapacityGb != nil {
			rFileSharesObject["capacityGb"] = *rFileSharesVal.CapacityGb
		}
		if rFileSharesVal.Name != nil {
			rFileSharesObject["name"] = *rFileSharesVal.Name
		}
		var rFileSharesValNfsExportOptions []interface{}
		for _, rFileSharesValNfsExportOptionsVal := range rFileSharesVal.NfsExportOptions {
			rFileSharesValNfsExportOptionsObject := make(map[string]interface{})
			if rFileSharesValNfsExportOptionsVal.AccessMode != nil {
				rFileSharesValNfsExportOptionsObject["accessMode"] = string(*rFileSharesValNfsExportOptionsVal.AccessMode)
			}
			if rFileSharesValNfsExportOptionsVal.AnonGid != nil {
				rFileSharesValNfsExportOptionsObject["anonGid"] = *rFileSharesValNfsExportOptionsVal.AnonGid
			}
			if rFileSharesValNfsExportOptionsVal.AnonUid != nil {
				rFileSharesValNfsExportOptionsObject["anonUid"] = *rFileSharesValNfsExportOptionsVal.AnonUid
			}
			var rFileSharesValNfsExportOptionsValIPRanges []interface{}
			for _, rFileSharesValNfsExportOptionsValIPRangesVal := range rFileSharesValNfsExportOptionsVal.IPRanges {
				rFileSharesValNfsExportOptionsValIPRanges = append(rFileSharesValNfsExportOptionsValIPRanges, rFileSharesValNfsExportOptionsValIPRangesVal)
			}
			rFileSharesValNfsExportOptionsObject["ipRanges"] = rFileSharesValNfsExportOptionsValIPRanges
			if rFileSharesValNfsExportOptionsVal.SquashMode != nil {
				rFileSharesValNfsExportOptionsObject["squashMode"] = string(*rFileSharesValNfsExportOptionsVal.SquashMode)
			}
			rFileSharesValNfsExportOptions = append(rFileSharesValNfsExportOptions, rFileSharesValNfsExportOptionsObject)
		}
		rFileSharesObject["nfsExportOptions"] = rFileSharesValNfsExportOptions
		if rFileSharesVal.SourceBackup != nil {
			rFileSharesObject["sourceBackup"] = *rFileSharesVal.SourceBackup
		}
		rFileShares = append(rFileShares, rFileSharesObject)
	}
	u.Object["fileShares"] = rFileShares
	if r.Labels != nil {
		rLabels := make(map[string]interface{})
		for k, v := range r.Labels {
			rLabels[k] = v
		}
		u.Object["labels"] = rLabels
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	var rNetworks []interface{}
	for _, rNetworksVal := range r.Networks {
		rNetworksObject := make(map[string]interface{})
		var rNetworksValIPAddresses []interface{}
		for _, rNetworksValIPAddressesVal := range rNetworksVal.IPAddresses {
			rNetworksValIPAddresses = append(rNetworksValIPAddresses, rNetworksValIPAddressesVal)
		}
		rNetworksObject["ipAddresses"] = rNetworksValIPAddresses
		var rNetworksValModes []interface{}
		for _, rNetworksValModesVal := range rNetworksVal.Modes {
			rNetworksValModes = append(rNetworksValModes, string(rNetworksValModesVal))
		}
		rNetworksObject["modes"] = rNetworksValModes
		if rNetworksVal.Network != nil {
			rNetworksObject["network"] = *rNetworksVal.Network
		}
		if rNetworksVal.ReservedIPRange != nil {
			rNetworksObject["reservedIPRange"] = *rNetworksVal.ReservedIPRange
		}
		rNetworks = append(rNetworks, rNetworksObject)
	}
	u.Object["networks"] = rNetworks
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.StatusMessage != nil {
		u.Object["statusMessage"] = *r.StatusMessage
	}
	if r.Tier != nil {
		u.Object["tier"] = string(*r.Tier)
	}
	return u
}

func UnstructuredToInstance(u *unstructured.Resource) (*dclService.Instance, error) {
	r := &dclService.Instance{}
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
	if _, ok := u.Object["fileShares"]; ok {
		if s, ok := u.Object["fileShares"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rFileShares dclService.InstanceFileShares
					if _, ok := objval["capacityGb"]; ok {
						if i, ok := objval["capacityGb"].(int64); ok {
							rFileShares.CapacityGb = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("rFileShares.CapacityGb: expected int64")
						}
					}
					if _, ok := objval["name"]; ok {
						if s, ok := objval["name"].(string); ok {
							rFileShares.Name = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rFileShares.Name: expected string")
						}
					}
					if _, ok := objval["nfsExportOptions"]; ok {
						if s, ok := objval["nfsExportOptions"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rFileSharesNfsExportOptions dclService.InstanceFileSharesNfsExportOptions
									if _, ok := objval["accessMode"]; ok {
										if s, ok := objval["accessMode"].(string); ok {
											rFileSharesNfsExportOptions.AccessMode = dclService.InstanceFileSharesNfsExportOptionsAccessModeEnumRef(s)
										} else {
											return nil, fmt.Errorf("rFileSharesNfsExportOptions.AccessMode: expected string")
										}
									}
									if _, ok := objval["anonGid"]; ok {
										if i, ok := objval["anonGid"].(int64); ok {
											rFileSharesNfsExportOptions.AnonGid = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rFileSharesNfsExportOptions.AnonGid: expected int64")
										}
									}
									if _, ok := objval["anonUid"]; ok {
										if i, ok := objval["anonUid"].(int64); ok {
											rFileSharesNfsExportOptions.AnonUid = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rFileSharesNfsExportOptions.AnonUid: expected int64")
										}
									}
									if _, ok := objval["ipRanges"]; ok {
										if s, ok := objval["ipRanges"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rFileSharesNfsExportOptions.IPRanges = append(rFileSharesNfsExportOptions.IPRanges, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rFileSharesNfsExportOptions.IPRanges: expected []interface{}")
										}
									}
									if _, ok := objval["squashMode"]; ok {
										if s, ok := objval["squashMode"].(string); ok {
											rFileSharesNfsExportOptions.SquashMode = dclService.InstanceFileSharesNfsExportOptionsSquashModeEnumRef(s)
										} else {
											return nil, fmt.Errorf("rFileSharesNfsExportOptions.SquashMode: expected string")
										}
									}
									rFileShares.NfsExportOptions = append(rFileShares.NfsExportOptions, rFileSharesNfsExportOptions)
								}
							}
						} else {
							return nil, fmt.Errorf("rFileShares.NfsExportOptions: expected []interface{}")
						}
					}
					if _, ok := objval["sourceBackup"]; ok {
						if s, ok := objval["sourceBackup"].(string); ok {
							rFileShares.SourceBackup = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rFileShares.SourceBackup: expected string")
						}
					}
					r.FileShares = append(r.FileShares, rFileShares)
				}
			}
		} else {
			return nil, fmt.Errorf("r.FileShares: expected []interface{}")
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
	if _, ok := u.Object["networks"]; ok {
		if s, ok := u.Object["networks"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rNetworks dclService.InstanceNetworks
					if _, ok := objval["ipAddresses"]; ok {
						if s, ok := objval["ipAddresses"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rNetworks.IPAddresses = append(rNetworks.IPAddresses, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rNetworks.IPAddresses: expected []interface{}")
						}
					}
					if _, ok := objval["modes"]; ok {
						if s, ok := objval["modes"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rNetworks.Modes = append(rNetworks.Modes, dclService.InstanceNetworksModesEnum(strval))
								}
							}
						} else {
							return nil, fmt.Errorf("rNetworks.Modes: expected []interface{}")
						}
					}
					if _, ok := objval["network"]; ok {
						if s, ok := objval["network"].(string); ok {
							rNetworks.Network = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rNetworks.Network: expected string")
						}
					}
					if _, ok := objval["reservedIPRange"]; ok {
						if s, ok := objval["reservedIPRange"].(string); ok {
							rNetworks.ReservedIPRange = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rNetworks.ReservedIPRange: expected string")
						}
					}
					r.Networks = append(r.Networks, rNetworks)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Networks: expected []interface{}")
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
			r.State = dclService.InstanceStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["statusMessage"]; ok {
		if s, ok := u.Object["statusMessage"].(string); ok {
			r.StatusMessage = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.StatusMessage: expected string")
		}
	}
	if _, ok := u.Object["tier"]; ok {
		if s, ok := u.Object["tier"].(string); ok {
			r.Tier = dclService.InstanceTierEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Tier: expected string")
		}
	}
	return r, nil
}

func GetInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetInstance(ctx, r)
	if err != nil {
		return nil, err
	}
	return InstanceToUnstructured(r), nil
}

func ListInstance(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListInstance(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, InstanceToUnstructured(r))
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

func ApplyInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInstance(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyInstance(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return InstanceToUnstructured(r), nil
}

func InstanceHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToInstance(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyInstance(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteInstance(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return err
	}
	return c.DeleteInstance(ctx, r)
}

func InstanceID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToInstance(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Instance) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"filestore",
		"Instance",
		"beta",
	}
}

func (r *Instance) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Instance) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Instance) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetInstance(ctx, config, resource)
}

func (r *Instance) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyInstance(ctx, config, resource, opts...)
}

func (r *Instance) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return InstanceHasDiff(ctx, config, resource, opts...)
}

func (r *Instance) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteInstance(ctx, config, resource)
}

func (r *Instance) ID(resource *unstructured.Resource) (string, error) {
	return InstanceID(resource)
}

func init() {
	unstructured.Register(&Instance{})
}
