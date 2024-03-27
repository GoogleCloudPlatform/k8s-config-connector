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
package datafusion

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datafusion/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Instance struct{}

func InstanceToUnstructured(r *dclService.Instance) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "datafusion",
			Version: "beta",
			Type:    "Instance",
		},
		Object: make(map[string]interface{}),
	}
	if r.ApiEndpoint != nil {
		u.Object["apiEndpoint"] = *r.ApiEndpoint
	}
	var rAvailableVersion []interface{}
	for _, rAvailableVersionVal := range r.AvailableVersion {
		rAvailableVersionObject := make(map[string]interface{})
		var rAvailableVersionValAvailableFeatures []interface{}
		for _, rAvailableVersionValAvailableFeaturesVal := range rAvailableVersionVal.AvailableFeatures {
			rAvailableVersionValAvailableFeatures = append(rAvailableVersionValAvailableFeatures, rAvailableVersionValAvailableFeaturesVal)
		}
		rAvailableVersionObject["availableFeatures"] = rAvailableVersionValAvailableFeatures
		if rAvailableVersionVal.DefaultVersion != nil {
			rAvailableVersionObject["defaultVersion"] = *rAvailableVersionVal.DefaultVersion
		}
		if rAvailableVersionVal.VersionNumber != nil {
			rAvailableVersionObject["versionNumber"] = *rAvailableVersionVal.VersionNumber
		}
		rAvailableVersion = append(rAvailableVersion, rAvailableVersionObject)
	}
	u.Object["availableVersion"] = rAvailableVersion
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DataprocServiceAccount != nil {
		u.Object["dataprocServiceAccount"] = *r.DataprocServiceAccount
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	if r.DisplayName != nil {
		u.Object["displayName"] = *r.DisplayName
	}
	if r.EnableStackdriverLogging != nil {
		u.Object["enableStackdriverLogging"] = *r.EnableStackdriverLogging
	}
	if r.EnableStackdriverMonitoring != nil {
		u.Object["enableStackdriverMonitoring"] = *r.EnableStackdriverMonitoring
	}
	if r.GcsBucket != nil {
		u.Object["gcsBucket"] = *r.GcsBucket
	}
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
	if r.NetworkConfig != nil && r.NetworkConfig != dclService.EmptyInstanceNetworkConfig {
		rNetworkConfig := make(map[string]interface{})
		if r.NetworkConfig.IPAllocation != nil {
			rNetworkConfig["ipAllocation"] = *r.NetworkConfig.IPAllocation
		}
		if r.NetworkConfig.Network != nil {
			rNetworkConfig["network"] = *r.NetworkConfig.Network
		}
		u.Object["networkConfig"] = rNetworkConfig
	}
	if r.Options != nil {
		rOptions := make(map[string]interface{})
		for k, v := range r.Options {
			rOptions[k] = v
		}
		u.Object["options"] = rOptions
	}
	if r.P4ServiceAccount != nil {
		u.Object["p4ServiceAccount"] = *r.P4ServiceAccount
	}
	if r.PrivateInstance != nil {
		u.Object["privateInstance"] = *r.PrivateInstance
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ServiceEndpoint != nil {
		u.Object["serviceEndpoint"] = *r.ServiceEndpoint
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.StateMessage != nil {
		u.Object["stateMessage"] = *r.StateMessage
	}
	if r.TenantProjectId != nil {
		u.Object["tenantProjectId"] = *r.TenantProjectId
	}
	if r.Type != nil {
		u.Object["type"] = string(*r.Type)
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.Version != nil {
		u.Object["version"] = *r.Version
	}
	if r.Zone != nil {
		u.Object["zone"] = *r.Zone
	}
	return u
}

func UnstructuredToInstance(u *unstructured.Resource) (*dclService.Instance, error) {
	r := &dclService.Instance{}
	if _, ok := u.Object["apiEndpoint"]; ok {
		if s, ok := u.Object["apiEndpoint"].(string); ok {
			r.ApiEndpoint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ApiEndpoint: expected string")
		}
	}
	if _, ok := u.Object["availableVersion"]; ok {
		if s, ok := u.Object["availableVersion"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rAvailableVersion dclService.InstanceAvailableVersion
					if _, ok := objval["availableFeatures"]; ok {
						if s, ok := objval["availableFeatures"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									rAvailableVersion.AvailableFeatures = append(rAvailableVersion.AvailableFeatures, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("rAvailableVersion.AvailableFeatures: expected []interface{}")
						}
					}
					if _, ok := objval["defaultVersion"]; ok {
						if b, ok := objval["defaultVersion"].(bool); ok {
							rAvailableVersion.DefaultVersion = dcl.Bool(b)
						} else {
							return nil, fmt.Errorf("rAvailableVersion.DefaultVersion: expected bool")
						}
					}
					if _, ok := objval["versionNumber"]; ok {
						if s, ok := objval["versionNumber"].(string); ok {
							rAvailableVersion.VersionNumber = dcl.String(s)
						} else {
							return nil, fmt.Errorf("rAvailableVersion.VersionNumber: expected string")
						}
					}
					r.AvailableVersion = append(r.AvailableVersion, rAvailableVersion)
				}
			}
		} else {
			return nil, fmt.Errorf("r.AvailableVersion: expected []interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["dataprocServiceAccount"]; ok {
		if s, ok := u.Object["dataprocServiceAccount"].(string); ok {
			r.DataprocServiceAccount = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DataprocServiceAccount: expected string")
		}
	}
	if _, ok := u.Object["description"]; ok {
		if s, ok := u.Object["description"].(string); ok {
			r.Description = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Description: expected string")
		}
	}
	if _, ok := u.Object["displayName"]; ok {
		if s, ok := u.Object["displayName"].(string); ok {
			r.DisplayName = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.DisplayName: expected string")
		}
	}
	if _, ok := u.Object["enableStackdriverLogging"]; ok {
		if b, ok := u.Object["enableStackdriverLogging"].(bool); ok {
			r.EnableStackdriverLogging = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableStackdriverLogging: expected bool")
		}
	}
	if _, ok := u.Object["enableStackdriverMonitoring"]; ok {
		if b, ok := u.Object["enableStackdriverMonitoring"].(bool); ok {
			r.EnableStackdriverMonitoring = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.EnableStackdriverMonitoring: expected bool")
		}
	}
	if _, ok := u.Object["gcsBucket"]; ok {
		if s, ok := u.Object["gcsBucket"].(string); ok {
			r.GcsBucket = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.GcsBucket: expected string")
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
	if _, ok := u.Object["networkConfig"]; ok {
		if rNetworkConfig, ok := u.Object["networkConfig"].(map[string]interface{}); ok {
			r.NetworkConfig = &dclService.InstanceNetworkConfig{}
			if _, ok := rNetworkConfig["ipAllocation"]; ok {
				if s, ok := rNetworkConfig["ipAllocation"].(string); ok {
					r.NetworkConfig.IPAllocation = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.NetworkConfig.IPAllocation: expected string")
				}
			}
			if _, ok := rNetworkConfig["network"]; ok {
				if s, ok := rNetworkConfig["network"].(string); ok {
					r.NetworkConfig.Network = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.NetworkConfig.Network: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.NetworkConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["options"]; ok {
		if rOptions, ok := u.Object["options"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rOptions {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Options = m
		} else {
			return nil, fmt.Errorf("r.Options: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["p4ServiceAccount"]; ok {
		if s, ok := u.Object["p4ServiceAccount"].(string); ok {
			r.P4ServiceAccount = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.P4ServiceAccount: expected string")
		}
	}
	if _, ok := u.Object["privateInstance"]; ok {
		if b, ok := u.Object["privateInstance"].(bool); ok {
			r.PrivateInstance = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.PrivateInstance: expected bool")
		}
	}
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["serviceEndpoint"]; ok {
		if s, ok := u.Object["serviceEndpoint"].(string); ok {
			r.ServiceEndpoint = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.ServiceEndpoint: expected string")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.InstanceStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["stateMessage"]; ok {
		if s, ok := u.Object["stateMessage"].(string); ok {
			r.StateMessage = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.StateMessage: expected string")
		}
	}
	if _, ok := u.Object["tenantProjectId"]; ok {
		if s, ok := u.Object["tenantProjectId"].(string); ok {
			r.TenantProjectId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.TenantProjectId: expected string")
		}
	}
	if _, ok := u.Object["type"]; ok {
		if s, ok := u.Object["type"].(string); ok {
			r.Type = dclService.InstanceTypeEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.Type: expected string")
		}
	}
	if _, ok := u.Object["updateTime"]; ok {
		if s, ok := u.Object["updateTime"].(string); ok {
			r.UpdateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.UpdateTime: expected string")
		}
	}
	if _, ok := u.Object["version"]; ok {
		if s, ok := u.Object["version"].(string); ok {
			r.Version = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Version: expected string")
		}
	}
	if _, ok := u.Object["zone"]; ok {
		if s, ok := u.Object["zone"].(string); ok {
			r.Zone = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Zone: expected string")
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
		"datafusion",
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
