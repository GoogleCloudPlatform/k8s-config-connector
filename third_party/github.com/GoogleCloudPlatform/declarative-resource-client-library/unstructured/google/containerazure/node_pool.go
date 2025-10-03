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
package containerazure

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containerazure"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type NodePool struct{}

func NodePoolToUnstructured(r *dclService.NodePool) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "containerazure",
			Version: "ga",
			Type:    "NodePool",
		},
		Object: make(map[string]interface{}),
	}
	if r.Annotations != nil {
		rAnnotations := make(map[string]interface{})
		for k, v := range r.Annotations {
			rAnnotations[k] = v
		}
		u.Object["annotations"] = rAnnotations
	}
	if r.Autoscaling != nil && r.Autoscaling != dclService.EmptyNodePoolAutoscaling {
		rAutoscaling := make(map[string]interface{})
		if r.Autoscaling.MaxNodeCount != nil {
			rAutoscaling["maxNodeCount"] = *r.Autoscaling.MaxNodeCount
		}
		if r.Autoscaling.MinNodeCount != nil {
			rAutoscaling["minNodeCount"] = *r.Autoscaling.MinNodeCount
		}
		u.Object["autoscaling"] = rAutoscaling
	}
	if r.AzureAvailabilityZone != nil {
		u.Object["azureAvailabilityZone"] = *r.AzureAvailabilityZone
	}
	if r.Cluster != nil {
		u.Object["cluster"] = *r.Cluster
	}
	if r.Config != nil && r.Config != dclService.EmptyNodePoolConfig {
		rConfig := make(map[string]interface{})
		if r.Config.Labels != nil {
			rConfigLabels := make(map[string]interface{})
			for k, v := range r.Config.Labels {
				rConfigLabels[k] = v
			}
			rConfig["labels"] = rConfigLabels
		}
		if r.Config.ProxyConfig != nil && r.Config.ProxyConfig != dclService.EmptyNodePoolConfigProxyConfig {
			rConfigProxyConfig := make(map[string]interface{})
			if r.Config.ProxyConfig.ResourceGroupId != nil {
				rConfigProxyConfig["resourceGroupId"] = *r.Config.ProxyConfig.ResourceGroupId
			}
			if r.Config.ProxyConfig.SecretId != nil {
				rConfigProxyConfig["secretId"] = *r.Config.ProxyConfig.SecretId
			}
			rConfig["proxyConfig"] = rConfigProxyConfig
		}
		if r.Config.RootVolume != nil && r.Config.RootVolume != dclService.EmptyNodePoolConfigRootVolume {
			rConfigRootVolume := make(map[string]interface{})
			if r.Config.RootVolume.SizeGib != nil {
				rConfigRootVolume["sizeGib"] = *r.Config.RootVolume.SizeGib
			}
			rConfig["rootVolume"] = rConfigRootVolume
		}
		if r.Config.SshConfig != nil && r.Config.SshConfig != dclService.EmptyNodePoolConfigSshConfig {
			rConfigSshConfig := make(map[string]interface{})
			if r.Config.SshConfig.AuthorizedKey != nil {
				rConfigSshConfig["authorizedKey"] = *r.Config.SshConfig.AuthorizedKey
			}
			rConfig["sshConfig"] = rConfigSshConfig
		}
		if r.Config.Tags != nil {
			rConfigTags := make(map[string]interface{})
			for k, v := range r.Config.Tags {
				rConfigTags[k] = v
			}
			rConfig["tags"] = rConfigTags
		}
		if r.Config.VmSize != nil {
			rConfig["vmSize"] = *r.Config.VmSize
		}
		u.Object["config"] = rConfig
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Etag != nil {
		u.Object["etag"] = *r.Etag
	}
	if r.Location != nil {
		u.Object["location"] = *r.Location
	}
	if r.Management != nil && r.Management != dclService.EmptyNodePoolManagement {
		rManagement := make(map[string]interface{})
		if r.Management.AutoRepair != nil {
			rManagement["autoRepair"] = *r.Management.AutoRepair
		}
		u.Object["management"] = rManagement
	}
	if r.MaxPodsConstraint != nil && r.MaxPodsConstraint != dclService.EmptyNodePoolMaxPodsConstraint {
		rMaxPodsConstraint := make(map[string]interface{})
		if r.MaxPodsConstraint.MaxPodsPerNode != nil {
			rMaxPodsConstraint["maxPodsPerNode"] = *r.MaxPodsConstraint.MaxPodsPerNode
		}
		u.Object["maxPodsConstraint"] = rMaxPodsConstraint
	}
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.Reconciling != nil {
		u.Object["reconciling"] = *r.Reconciling
	}
	if r.State != nil {
		u.Object["state"] = string(*r.State)
	}
	if r.SubnetId != nil {
		u.Object["subnetId"] = *r.SubnetId
	}
	if r.Uid != nil {
		u.Object["uid"] = *r.Uid
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	if r.Version != nil {
		u.Object["version"] = *r.Version
	}
	return u
}

func UnstructuredToNodePool(u *unstructured.Resource) (*dclService.NodePool, error) {
	r := &dclService.NodePool{}
	if _, ok := u.Object["annotations"]; ok {
		if rAnnotations, ok := u.Object["annotations"].(map[string]interface{}); ok {
			m := make(map[string]string)
			for k, v := range rAnnotations {
				if s, ok := v.(string); ok {
					m[k] = s
				}
			}
			r.Annotations = m
		} else {
			return nil, fmt.Errorf("r.Annotations: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["autoscaling"]; ok {
		if rAutoscaling, ok := u.Object["autoscaling"].(map[string]interface{}); ok {
			r.Autoscaling = &dclService.NodePoolAutoscaling{}
			if _, ok := rAutoscaling["maxNodeCount"]; ok {
				if i, ok := rAutoscaling["maxNodeCount"].(int64); ok {
					r.Autoscaling.MaxNodeCount = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.Autoscaling.MaxNodeCount: expected int64")
				}
			}
			if _, ok := rAutoscaling["minNodeCount"]; ok {
				if i, ok := rAutoscaling["minNodeCount"].(int64); ok {
					r.Autoscaling.MinNodeCount = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.Autoscaling.MinNodeCount: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Autoscaling: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["azureAvailabilityZone"]; ok {
		if s, ok := u.Object["azureAvailabilityZone"].(string); ok {
			r.AzureAvailabilityZone = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.AzureAvailabilityZone: expected string")
		}
	}
	if _, ok := u.Object["cluster"]; ok {
		if s, ok := u.Object["cluster"].(string); ok {
			r.Cluster = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Cluster: expected string")
		}
	}
	if _, ok := u.Object["config"]; ok {
		if rConfig, ok := u.Object["config"].(map[string]interface{}); ok {
			r.Config = &dclService.NodePoolConfig{}
			if _, ok := rConfig["labels"]; ok {
				if rConfigLabels, ok := rConfig["labels"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rConfigLabels {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.Config.Labels = m
				} else {
					return nil, fmt.Errorf("r.Config.Labels: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["proxyConfig"]; ok {
				if rConfigProxyConfig, ok := rConfig["proxyConfig"].(map[string]interface{}); ok {
					r.Config.ProxyConfig = &dclService.NodePoolConfigProxyConfig{}
					if _, ok := rConfigProxyConfig["resourceGroupId"]; ok {
						if s, ok := rConfigProxyConfig["resourceGroupId"].(string); ok {
							r.Config.ProxyConfig.ResourceGroupId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.ProxyConfig.ResourceGroupId: expected string")
						}
					}
					if _, ok := rConfigProxyConfig["secretId"]; ok {
						if s, ok := rConfigProxyConfig["secretId"].(string); ok {
							r.Config.ProxyConfig.SecretId = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.ProxyConfig.SecretId: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.ProxyConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["rootVolume"]; ok {
				if rConfigRootVolume, ok := rConfig["rootVolume"].(map[string]interface{}); ok {
					r.Config.RootVolume = &dclService.NodePoolConfigRootVolume{}
					if _, ok := rConfigRootVolume["sizeGib"]; ok {
						if i, ok := rConfigRootVolume["sizeGib"].(int64); ok {
							r.Config.RootVolume.SizeGib = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Config.RootVolume.SizeGib: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.RootVolume: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["sshConfig"]; ok {
				if rConfigSshConfig, ok := rConfig["sshConfig"].(map[string]interface{}); ok {
					r.Config.SshConfig = &dclService.NodePoolConfigSshConfig{}
					if _, ok := rConfigSshConfig["authorizedKey"]; ok {
						if s, ok := rConfigSshConfig["authorizedKey"].(string); ok {
							r.Config.SshConfig.AuthorizedKey = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.SshConfig.AuthorizedKey: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.SshConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["tags"]; ok {
				if rConfigTags, ok := rConfig["tags"].(map[string]interface{}); ok {
					m := make(map[string]string)
					for k, v := range rConfigTags {
						if s, ok := v.(string); ok {
							m[k] = s
						}
					}
					r.Config.Tags = m
				} else {
					return nil, fmt.Errorf("r.Config.Tags: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["vmSize"]; ok {
				if s, ok := rConfig["vmSize"].(string); ok {
					r.Config.VmSize = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Config.VmSize: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Config: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["createTime"]; ok {
		if s, ok := u.Object["createTime"].(string); ok {
			r.CreateTime = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.CreateTime: expected string")
		}
	}
	if _, ok := u.Object["etag"]; ok {
		if s, ok := u.Object["etag"].(string); ok {
			r.Etag = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Etag: expected string")
		}
	}
	if _, ok := u.Object["location"]; ok {
		if s, ok := u.Object["location"].(string); ok {
			r.Location = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Location: expected string")
		}
	}
	if _, ok := u.Object["management"]; ok {
		if rManagement, ok := u.Object["management"].(map[string]interface{}); ok {
			r.Management = &dclService.NodePoolManagement{}
			if _, ok := rManagement["autoRepair"]; ok {
				if b, ok := rManagement["autoRepair"].(bool); ok {
					r.Management.AutoRepair = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.Management.AutoRepair: expected bool")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Management: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["maxPodsConstraint"]; ok {
		if rMaxPodsConstraint, ok := u.Object["maxPodsConstraint"].(map[string]interface{}); ok {
			r.MaxPodsConstraint = &dclService.NodePoolMaxPodsConstraint{}
			if _, ok := rMaxPodsConstraint["maxPodsPerNode"]; ok {
				if i, ok := rMaxPodsConstraint["maxPodsPerNode"].(int64); ok {
					r.MaxPodsConstraint.MaxPodsPerNode = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.MaxPodsConstraint.MaxPodsPerNode: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.MaxPodsConstraint: expected map[string]interface{}")
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
	if _, ok := u.Object["reconciling"]; ok {
		if b, ok := u.Object["reconciling"].(bool); ok {
			r.Reconciling = dcl.Bool(b)
		} else {
			return nil, fmt.Errorf("r.Reconciling: expected bool")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if s, ok := u.Object["state"].(string); ok {
			r.State = dclService.NodePoolStateEnumRef(s)
		} else {
			return nil, fmt.Errorf("r.State: expected string")
		}
	}
	if _, ok := u.Object["subnetId"]; ok {
		if s, ok := u.Object["subnetId"].(string); ok {
			r.SubnetId = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SubnetId: expected string")
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
	if _, ok := u.Object["version"]; ok {
		if s, ok := u.Object["version"].(string); ok {
			r.Version = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Version: expected string")
		}
	}
	return r, nil
}

func GetNodePool(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNodePool(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetNodePool(ctx, r)
	if err != nil {
		return nil, err
	}
	return NodePoolToUnstructured(r), nil
}

func ListNodePool(ctx context.Context, config *dcl.Config, project string, location string, cluster string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListNodePool(ctx, project, location, cluster)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, NodePoolToUnstructured(r))
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

func ApplyNodePool(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNodePool(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNodePool(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyNodePool(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return NodePoolToUnstructured(r), nil
}

func NodePoolHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNodePool(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToNodePool(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyNodePool(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteNodePool(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToNodePool(u)
	if err != nil {
		return err
	}
	return c.DeleteNodePool(ctx, r)
}

func NodePoolID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToNodePool(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *NodePool) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"containerazure",
		"NodePool",
		"ga",
	}
}

func (r *NodePool) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NodePool) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NodePool) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *NodePool) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NodePool) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NodePool) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *NodePool) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetNodePool(ctx, config, resource)
}

func (r *NodePool) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyNodePool(ctx, config, resource, opts...)
}

func (r *NodePool) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return NodePoolHasDiff(ctx, config, resource, opts...)
}

func (r *NodePool) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteNodePool(ctx, config, resource)
}

func (r *NodePool) ID(resource *unstructured.Resource) (string, error) {
	return NodePoolID(resource)
}

func init() {
	unstructured.Register(&NodePool{})
}
