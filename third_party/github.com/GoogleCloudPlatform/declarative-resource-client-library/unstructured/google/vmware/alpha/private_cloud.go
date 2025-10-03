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
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vmware/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	iamUnstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured/google/iam"
)

type PrivateCloud struct{}

func PrivateCloudToUnstructured(r *dclService.PrivateCloud) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "vmware",
			Version: "alpha",
			Type:    "PrivateCloud",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DeleteTime != nil {
		u.Object["deleteTime"] = *r.DeleteTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.ExpireTime != nil {
		u.Object["expireTime"] = *r.ExpireTime
	}
	if r.Hcx != nil && r.Hcx != dclService.EmptyPrivateCloudHcx {
		rHcx := make(map[string]interface{})
		if r.Hcx.Fqdn != nil {
			rHcx["fqdn"] = *r.Hcx.Fqdn
		}
		if r.Hcx.InternalIP != nil {
			rHcx["internalIP"] = *r.Hcx.InternalIP
		}
		if r.Hcx.State != nil {
			rHcx["state"] = string(*r.Hcx.State)
		}
		if r.Hcx.Version != nil {
			rHcx["version"] = *r.Hcx.Version
		}
		u.Object["hcx"] = rHcx
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.ManagementCluster != nil && r.ManagementCluster != dclService.EmptyPrivateCloudManagementCluster {
		rManagementCluster := make(map[string]interface{})
		if r.ManagementCluster.ClusterId != nil {
			rManagementCluster["clusterId"] = *r.ManagementCluster.ClusterId
		}
		u.Object["managementCluster"] = rManagementCluster
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.NetworkConfig != nil && r.NetworkConfig != dclService.EmptyPrivateCloudNetworkConfig {
		rNetworkConfig := make(map[string]interface{})
		if r.NetworkConfig.ManagementCidr != nil {
			rNetworkConfig["managementCidr"] = *r.NetworkConfig.ManagementCidr
		}
		if r.NetworkConfig.ManagementIPAddressLayoutVersion != nil {
			rNetworkConfig["managementIPAddressLayoutVersion"] = *r.NetworkConfig.ManagementIPAddressLayoutVersion
		}
		if r.NetworkConfig.VmwareEngineNetwork != nil {
			rNetworkConfig["vmwareEngineNetwork"] = *r.NetworkConfig.VmwareEngineNetwork
		}
		if r.NetworkConfig.VmwareEngineNetworkCanonical != nil {
			rNetworkConfig["vmwareEngineNetworkCanonical"] = *r.NetworkConfig.VmwareEngineNetworkCanonical
		}
		u.Object["networkConfig"] = rNetworkConfig
	}
	if r.Nsx != nil && r.Nsx != dclService.EmptyPrivateCloudNsx {
		rNsx := make(map[string]interface{})
		if r.Nsx.Fqdn != nil {
			rNsx["fqdn"] = *r.Nsx.Fqdn
		}
		if r.Nsx.InternalIP != nil {
			rNsx["internalIP"] = *r.Nsx.InternalIP
		}
		if r.Nsx.State != nil {
			rNsx["state"] = string(*r.Nsx.State)
		}
		if r.Nsx.Version != nil {
			rNsx["version"] = *r.Nsx.Version
		}
		u.Object["nsx"] = rNsx
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.Vcenter != nil && r.Vcenter != dclService.EmptyPrivateCloudVcenter {
		rVcenter := make(map[string]interface{})
		if r.Vcenter.Fqdn != nil {
			rVcenter["fqdn"] = *r.Vcenter.Fqdn
		}
		if r.Vcenter.InternalIP != nil {
			rVcenter["internalIP"] = *r.Vcenter.InternalIP
		}
		if r.Vcenter.State != nil {
			rVcenter["state"] = string(*r.Vcenter.State)
		}
		if r.Vcenter.Version != nil {
			rVcenter["version"] = *r.Vcenter.Version
		}
		u.Object["vcenter"] = rVcenter
	}
	return u
}

func UnstructuredToPrivateCloud(u *unstructured.Resource) (*dclService.PrivateCloud, error) {
	r := &dclService.PrivateCloud{}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["deleteTime"]; ok {
		if s, ok := u.Object["deleteTime"].(string); ok {
			r.DeleteTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DeleteTime: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["expireTime"]; ok {
		if s, ok := u.Object["expireTime"].(string); ok {
			r.ExpireTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ExpireTime: expected string")
		}
	}
	if _, ok := u.Object["hcx"]; ok {
		if rHcx, ok := u.Object["hcx"].(map[string]interface{}); ok {
			r.Hcx = &dclService.PrivateCloudHcx{}
			if _, ok := rHcx["fqdn"]; ok {
				if s, ok := rHcx["fqdn"].(string); ok {
					r.Hcx.Fqdn = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Hcx.Fqdn: expected string")
				}
			}
			if _, ok := rHcx["internalIP"]; ok {
				if s, ok := rHcx["internalIP"].(string); ok {
					r.Hcx.InternalIP = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Hcx.InternalIP: expected string")
				}
			}
			if _, ok := rHcx["state"]; ok {
				if s, ok := rHcx["state"].(string); ok {
					r.Hcx.State = dclService.PrivateCloudHcxStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Hcx.State: expected string")
				}
			}
			if _, ok := rHcx["version"]; ok {
				if s, ok := rHcx["version"].(string); ok {
					r.Hcx.Version = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Hcx.Version: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Hcx: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["managementCluster"]; ok {
		if rManagementCluster, ok := u.Object["managementCluster"].(map[string]interface{}); ok {
			r.ManagementCluster = &dclService.PrivateCloudManagementCluster{}
			if _, ok := rManagementCluster["clusterId"]; ok {
				if s, ok := rManagementCluster["clusterId"].(string); ok {
					r.ManagementCluster.ClusterId = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.ManagementCluster.ClusterId: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ManagementCluster: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["name"]; ok {
		if s, ok := u.Object["name"].(string); ok {
			r.Name = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Name: expected string")
		}
	}
	if _, ok := u.Object["networkConfig"]; ok {
		if rNetworkConfig, ok := u.Object["networkConfig"].(map[string]interface{}); ok {
			r.NetworkConfig = &dclService.PrivateCloudNetworkConfig{}
			if _, ok := rNetworkConfig["managementCidr"]; ok {
				if s, ok := rNetworkConfig["managementCidr"].(string); ok {
					r.NetworkConfig.ManagementCidr = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.NetworkConfig.ManagementCidr: expected string")
				}
			}
			if _, ok := rNetworkConfig["managementIPAddressLayoutVersion"]; ok {
				if i, ok := rNetworkConfig["managementIPAddressLayoutVersion"].(int64); ok {
					r.NetworkConfig.ManagementIPAddressLayoutVersion = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.NetworkConfig.ManagementIPAddressLayoutVersion: expected int64")
				}
			}
			if _, ok := rNetworkConfig["vmwareEngineNetwork"]; ok {
				if s, ok := rNetworkConfig["vmwareEngineNetwork"].(string); ok {
					r.NetworkConfig.VmwareEngineNetwork = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.NetworkConfig.VmwareEngineNetwork: expected string")
				}
			}
			if _, ok := rNetworkConfig["vmwareEngineNetworkCanonical"]; ok {
				if s, ok := rNetworkConfig["vmwareEngineNetworkCanonical"].(string); ok {
					r.NetworkConfig.VmwareEngineNetworkCanonical = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.NetworkConfig.VmwareEngineNetworkCanonical: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.NetworkConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["nsx"]; ok {
		if rNsx, ok := u.Object["nsx"].(map[string]interface{}); ok {
			r.Nsx = &dclService.PrivateCloudNsx{}
			if _, ok := rNsx["fqdn"]; ok {
				if s, ok := rNsx["fqdn"].(string); ok {
					r.Nsx.Fqdn = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Nsx.Fqdn: expected string")
				}
			}
			if _, ok := rNsx["internalIP"]; ok {
				if s, ok := rNsx["internalIP"].(string); ok {
					r.Nsx.InternalIP = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Nsx.InternalIP: expected string")
				}
			}
			if _, ok := rNsx["state"]; ok {
				if s, ok := rNsx["state"].(string); ok {
					r.Nsx.State = dclService.PrivateCloudNsxStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Nsx.State: expected string")
				}
			}
			if _, ok := rNsx["version"]; ok {
				if s, ok := rNsx["version"].(string); ok {
					r.Nsx.Version = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Nsx.Version: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Nsx: expected map[string]interface{}")
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
			r.State = dclService.PrivateCloudStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
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
	if _, ok := u.Object["vcenter"]; ok {
		if rVcenter, ok := u.Object["vcenter"].(map[string]interface{}); ok {
			r.Vcenter = &dclService.PrivateCloudVcenter{}
			if _, ok := rVcenter["fqdn"]; ok {
				if s, ok := rVcenter["fqdn"].(string); ok {
					r.Vcenter.Fqdn = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Vcenter.Fqdn: expected string")
				}
			}
			if _, ok := rVcenter["internalIP"]; ok {
				if s, ok := rVcenter["internalIP"].(string); ok {
					r.Vcenter.InternalIP = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Vcenter.InternalIP: expected string")
				}
			}
			if _, ok := rVcenter["state"]; ok {
				if s, ok := rVcenter["state"].(string); ok {
					r.Vcenter.State = dclService.PrivateCloudVcenterStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.Vcenter.State: expected string")
				}
			}
			if _, ok := rVcenter["version"]; ok {
				if s, ok := rVcenter["version"].(string); ok {
					r.Vcenter.Version = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Vcenter.Version: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Vcenter: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetPrivateCloud(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetPrivateCloud(ctx, r)
	if err != nil {
		return nil, err
	}
	return PrivateCloudToUnstructured(r), nil
}

func ListPrivateCloud(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListPrivateCloud(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, PrivateCloudToUnstructured(r))
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

func ApplyPrivateCloud(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPrivateCloud(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyPrivateCloud(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return PrivateCloudToUnstructured(r), nil
}

func PrivateCloudHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToPrivateCloud(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyPrivateCloud(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeletePrivateCloud(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return err
	}
	return c.DeletePrivateCloud(ctx, r)
}

func PrivateCloudID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *PrivateCloud) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"vmware",
		"PrivateCloud",
		"alpha",
	}
}

func SetPolicyPrivateCloud(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicy(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func SetPolicyWithEtagPrivateCloud(ctx context.Context, config *dcl.Config, u *unstructured.Resource, p *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return nil, err
	}
	policy, err := iamUnstruct.UnstructuredToPolicy(p)
	if err != nil {
		return nil, err
	}
	policy.Resource = r
	iamClient := iam.NewClient(config)
	newPolicy, err := iamClient.SetPolicyWithEtag(ctx, policy)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(newPolicy), nil
}

func GetPolicyPrivateCloud(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policy, err := iamClient.GetPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func SetPolicyMemberPrivateCloud(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) (*unstructured.Resource, error) {
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return nil, err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return nil, err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	policy, err := iamClient.SetMember(ctx, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.PolicyToUnstructured(policy), nil
}

func GetPolicyMemberPrivateCloud(ctx context.Context, config *dcl.Config, u *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return nil, err
	}
	iamClient := iam.NewClient(config)
	policyMember, err := iamClient.GetMember(ctx, r, role, member)
	if err != nil {
		return nil, err
	}
	return iamUnstruct.MemberToUnstructured(policyMember), nil
}

func DeletePolicyMemberPrivateCloud(ctx context.Context, config *dcl.Config, u *unstructured.Resource, m *unstructured.Resource) error {
	r, err := UnstructuredToPrivateCloud(u)
	if err != nil {
		return err
	}
	member, err := iamUnstruct.UnstructuredToMember(m)
	if err != nil {
		return err
	}
	member.Resource = r
	iamClient := iam.NewClient(config)
	if err := iamClient.DeleteMember(ctx, member); err != nil {
		return err
	}
	return nil
}

func (r *PrivateCloud) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyMemberPrivateCloud(ctx, config, resource, member)
}

func (r *PrivateCloud) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return GetPolicyMemberPrivateCloud(ctx, config, resource, role, member)
}

func (r *PrivateCloud) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return DeletePolicyMemberPrivateCloud(ctx, config, resource, member)
}

func (r *PrivateCloud) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyPrivateCloud(ctx, config, resource, policy)
}

func (r *PrivateCloud) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return SetPolicyWithEtagPrivateCloud(ctx, config, resource, policy)
}

func (r *PrivateCloud) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPolicyPrivateCloud(ctx, config, resource)
}

func (r *PrivateCloud) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetPrivateCloud(ctx, config, resource)
}

func (r *PrivateCloud) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyPrivateCloud(ctx, config, resource, opts...)
}

func (r *PrivateCloud) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return PrivateCloudHasDiff(ctx, config, resource, opts...)
}

func (r *PrivateCloud) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeletePrivateCloud(ctx, config, resource)
}

func (r *PrivateCloud) ID(resource *unstructured.Resource) (string, error) {
	return PrivateCloudID(resource)
}

func init() {
	unstructured.Register(&PrivateCloud{})
}
