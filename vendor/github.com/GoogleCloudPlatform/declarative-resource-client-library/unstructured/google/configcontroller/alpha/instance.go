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
package configcontroller

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/configcontroller/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Instance struct{}

func InstanceToUnstructured(r *dclService.Instance) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "configcontroller",
			Version: "alpha",
			Type:    "Instance",
		},
		Object: make(map[string]interface{}),
	}
	if r.GkeResourceLink != nil {
		u.Object["gkeResourceLink"] = *r.GkeResourceLink
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
	if r.ManagementConfig != nil && r.ManagementConfig != dclService.EmptyInstanceManagementConfig {
		rManagementConfig := make(map[string]interface{})
		if r.ManagementConfig.FullManagementConfig != nil && r.ManagementConfig.FullManagementConfig != dclService.EmptyInstanceManagementConfigFullManagementConfig {
			rManagementConfigFullManagementConfig := make(map[string]interface{})
			if r.ManagementConfig.FullManagementConfig.ClusterCidrBlock != nil {
				rManagementConfigFullManagementConfig["clusterCidrBlock"] = *r.ManagementConfig.FullManagementConfig.ClusterCidrBlock
			}
			if r.ManagementConfig.FullManagementConfig.ClusterNamedRange != nil {
				rManagementConfigFullManagementConfig["clusterNamedRange"] = *r.ManagementConfig.FullManagementConfig.ClusterNamedRange
			}
			if r.ManagementConfig.FullManagementConfig.ManBlock != nil {
				rManagementConfigFullManagementConfig["manBlock"] = *r.ManagementConfig.FullManagementConfig.ManBlock
			}
			if r.ManagementConfig.FullManagementConfig.MasterIPv4CidrBlock != nil {
				rManagementConfigFullManagementConfig["masterIPv4CidrBlock"] = *r.ManagementConfig.FullManagementConfig.MasterIPv4CidrBlock
			}
			if r.ManagementConfig.FullManagementConfig.Network != nil {
				rManagementConfigFullManagementConfig["network"] = *r.ManagementConfig.FullManagementConfig.Network
			}
			if r.ManagementConfig.FullManagementConfig.ServicesCidrBlock != nil {
				rManagementConfigFullManagementConfig["servicesCidrBlock"] = *r.ManagementConfig.FullManagementConfig.ServicesCidrBlock
			}
			if r.ManagementConfig.FullManagementConfig.ServicesNamedRange != nil {
				rManagementConfigFullManagementConfig["servicesNamedRange"] = *r.ManagementConfig.FullManagementConfig.ServicesNamedRange
			}
			rManagementConfig["fullManagementConfig"] = rManagementConfigFullManagementConfig
		}
		if r.ManagementConfig.StandardManagementConfig != nil && r.ManagementConfig.StandardManagementConfig != dclService.EmptyInstanceManagementConfigStandardManagementConfig {
			rManagementConfigStandardManagementConfig := make(map[string]interface{})
			if r.ManagementConfig.StandardManagementConfig.ClusterCidrBlock != nil {
				rManagementConfigStandardManagementConfig["clusterCidrBlock"] = *r.ManagementConfig.StandardManagementConfig.ClusterCidrBlock
			}
			if r.ManagementConfig.StandardManagementConfig.ClusterNamedRange != nil {
				rManagementConfigStandardManagementConfig["clusterNamedRange"] = *r.ManagementConfig.StandardManagementConfig.ClusterNamedRange
			}
			if r.ManagementConfig.StandardManagementConfig.ManBlock != nil {
				rManagementConfigStandardManagementConfig["manBlock"] = *r.ManagementConfig.StandardManagementConfig.ManBlock
			}
			if r.ManagementConfig.StandardManagementConfig.MasterIPv4CidrBlock != nil {
				rManagementConfigStandardManagementConfig["masterIPv4CidrBlock"] = *r.ManagementConfig.StandardManagementConfig.MasterIPv4CidrBlock
			}
			if r.ManagementConfig.StandardManagementConfig.Network != nil {
				rManagementConfigStandardManagementConfig["network"] = *r.ManagementConfig.StandardManagementConfig.Network
			}
			if r.ManagementConfig.StandardManagementConfig.ServicesCidrBlock != nil {
				rManagementConfigStandardManagementConfig["servicesCidrBlock"] = *r.ManagementConfig.StandardManagementConfig.ServicesCidrBlock
			}
			if r.ManagementConfig.StandardManagementConfig.ServicesNamedRange != nil {
				rManagementConfigStandardManagementConfig["servicesNamedRange"] = *r.ManagementConfig.StandardManagementConfig.ServicesNamedRange
			}
			rManagementConfig["standardManagementConfig"] = rManagementConfigStandardManagementConfig
		}
		u.Object["managementConfig"] = rManagementConfig
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
	if r.UsePrivateEndpoint != nil {
		u.Object["usePrivateEndpoint"] = *r.UsePrivateEndpoint
	}
	return u
}

func UnstructuredToInstance(u *unstructured.Resource) (*dclService.Instance, error) {
	r := &dclService.Instance{}
	if _, ok := u.Object["gkeResourceLink"]; ok {
		if s, ok := u.Object["gkeResourceLink"].(string); ok {
			r.GkeResourceLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.GkeResourceLink: expected string")
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
	if _, ok := u.Object["managementConfig"]; ok {
		if rManagementConfig, ok := u.Object["managementConfig"].(map[string]interface{}); ok {
			r.ManagementConfig = &dclService.InstanceManagementConfig{}
			if _, ok := rManagementConfig["fullManagementConfig"]; ok {
				if rManagementConfigFullManagementConfig, ok := rManagementConfig["fullManagementConfig"].(map[string]interface{}); ok {
					r.ManagementConfig.FullManagementConfig = &dclService.InstanceManagementConfigFullManagementConfig{}
					if _, ok := rManagementConfigFullManagementConfig["clusterCidrBlock"]; ok {
						if s, ok := rManagementConfigFullManagementConfig["clusterCidrBlock"].(string); ok {
							r.ManagementConfig.FullManagementConfig.ClusterCidrBlock = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.FullManagementConfig.ClusterCidrBlock: expected string")
						}
					}
					if _, ok := rManagementConfigFullManagementConfig["clusterNamedRange"]; ok {
						if s, ok := rManagementConfigFullManagementConfig["clusterNamedRange"].(string); ok {
							r.ManagementConfig.FullManagementConfig.ClusterNamedRange = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.FullManagementConfig.ClusterNamedRange: expected string")
						}
					}
					if _, ok := rManagementConfigFullManagementConfig["manBlock"]; ok {
						if s, ok := rManagementConfigFullManagementConfig["manBlock"].(string); ok {
							r.ManagementConfig.FullManagementConfig.ManBlock = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.FullManagementConfig.ManBlock: expected string")
						}
					}
					if _, ok := rManagementConfigFullManagementConfig["masterIPv4CidrBlock"]; ok {
						if s, ok := rManagementConfigFullManagementConfig["masterIPv4CidrBlock"].(string); ok {
							r.ManagementConfig.FullManagementConfig.MasterIPv4CidrBlock = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.FullManagementConfig.MasterIPv4CidrBlock: expected string")
						}
					}
					if _, ok := rManagementConfigFullManagementConfig["network"]; ok {
						if s, ok := rManagementConfigFullManagementConfig["network"].(string); ok {
							r.ManagementConfig.FullManagementConfig.Network = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.FullManagementConfig.Network: expected string")
						}
					}
					if _, ok := rManagementConfigFullManagementConfig["servicesCidrBlock"]; ok {
						if s, ok := rManagementConfigFullManagementConfig["servicesCidrBlock"].(string); ok {
							r.ManagementConfig.FullManagementConfig.ServicesCidrBlock = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.FullManagementConfig.ServicesCidrBlock: expected string")
						}
					}
					if _, ok := rManagementConfigFullManagementConfig["servicesNamedRange"]; ok {
						if s, ok := rManagementConfigFullManagementConfig["servicesNamedRange"].(string); ok {
							r.ManagementConfig.FullManagementConfig.ServicesNamedRange = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.FullManagementConfig.ServicesNamedRange: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ManagementConfig.FullManagementConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rManagementConfig["standardManagementConfig"]; ok {
				if rManagementConfigStandardManagementConfig, ok := rManagementConfig["standardManagementConfig"].(map[string]interface{}); ok {
					r.ManagementConfig.StandardManagementConfig = &dclService.InstanceManagementConfigStandardManagementConfig{}
					if _, ok := rManagementConfigStandardManagementConfig["clusterCidrBlock"]; ok {
						if s, ok := rManagementConfigStandardManagementConfig["clusterCidrBlock"].(string); ok {
							r.ManagementConfig.StandardManagementConfig.ClusterCidrBlock = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.StandardManagementConfig.ClusterCidrBlock: expected string")
						}
					}
					if _, ok := rManagementConfigStandardManagementConfig["clusterNamedRange"]; ok {
						if s, ok := rManagementConfigStandardManagementConfig["clusterNamedRange"].(string); ok {
							r.ManagementConfig.StandardManagementConfig.ClusterNamedRange = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.StandardManagementConfig.ClusterNamedRange: expected string")
						}
					}
					if _, ok := rManagementConfigStandardManagementConfig["manBlock"]; ok {
						if s, ok := rManagementConfigStandardManagementConfig["manBlock"].(string); ok {
							r.ManagementConfig.StandardManagementConfig.ManBlock = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.StandardManagementConfig.ManBlock: expected string")
						}
					}
					if _, ok := rManagementConfigStandardManagementConfig["masterIPv4CidrBlock"]; ok {
						if s, ok := rManagementConfigStandardManagementConfig["masterIPv4CidrBlock"].(string); ok {
							r.ManagementConfig.StandardManagementConfig.MasterIPv4CidrBlock = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.StandardManagementConfig.MasterIPv4CidrBlock: expected string")
						}
					}
					if _, ok := rManagementConfigStandardManagementConfig["network"]; ok {
						if s, ok := rManagementConfigStandardManagementConfig["network"].(string); ok {
							r.ManagementConfig.StandardManagementConfig.Network = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.StandardManagementConfig.Network: expected string")
						}
					}
					if _, ok := rManagementConfigStandardManagementConfig["servicesCidrBlock"]; ok {
						if s, ok := rManagementConfigStandardManagementConfig["servicesCidrBlock"].(string); ok {
							r.ManagementConfig.StandardManagementConfig.ServicesCidrBlock = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.StandardManagementConfig.ServicesCidrBlock: expected string")
						}
					}
					if _, ok := rManagementConfigStandardManagementConfig["servicesNamedRange"]; ok {
						if s, ok := rManagementConfigStandardManagementConfig["servicesNamedRange"].(string); ok {
							r.ManagementConfig.StandardManagementConfig.ServicesNamedRange = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.ManagementConfig.StandardManagementConfig.ServicesNamedRange: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.ManagementConfig.StandardManagementConfig: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ManagementConfig: expected map[string]interface{}")
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
			r.State = dclService.InstanceStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["usePrivateEndpoint"]; ok {
		if b, ok := u.Object["usePrivateEndpoint"].(bool); ok {
			r.UsePrivateEndpoint = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.UsePrivateEndpoint: expected bool")
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
		"configcontroller",
		"Instance",
		"alpha",
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
