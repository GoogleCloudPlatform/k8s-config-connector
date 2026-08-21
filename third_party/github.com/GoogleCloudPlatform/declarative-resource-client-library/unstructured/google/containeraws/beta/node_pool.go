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
package containeraws

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/containeraws/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type NodePool struct{}

func NodePoolToUnstructured(r *dclService.NodePool) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "containeraws",
			Version: "beta",
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
	if r.Cluster != nil {
		u.Object["cluster"] = *r.Cluster
	}
	if r.Config != nil && r.Config != dclService.EmptyNodePoolConfig {
		rConfig := make(map[string]interface{})
		if r.Config.AutoscalingMetricsCollection != nil && r.Config.AutoscalingMetricsCollection != dclService.EmptyNodePoolConfigAutoscalingMetricsCollection {
			rConfigAutoscalingMetricsCollection := make(map[string]interface{})
			if r.Config.AutoscalingMetricsCollection.Granularity != nil {
				rConfigAutoscalingMetricsCollection["granularity"] = *r.Config.AutoscalingMetricsCollection.Granularity
			}
			var rConfigAutoscalingMetricsCollectionMetrics []interface{}
			for _, rConfigAutoscalingMetricsCollectionMetricsVal := range r.Config.AutoscalingMetricsCollection.Metrics {
				rConfigAutoscalingMetricsCollectionMetrics = append(rConfigAutoscalingMetricsCollectionMetrics, rConfigAutoscalingMetricsCollectionMetricsVal)
			}
			rConfigAutoscalingMetricsCollection["metrics"] = rConfigAutoscalingMetricsCollectionMetrics
			rConfig["autoscalingMetricsCollection"] = rConfigAutoscalingMetricsCollection
		}
		if r.Config.ConfigEncryption != nil && r.Config.ConfigEncryption != dclService.EmptyNodePoolConfigConfigEncryption {
			rConfigConfigEncryption := make(map[string]interface{})
			if r.Config.ConfigEncryption.KmsKeyArn != nil {
				rConfigConfigEncryption["kmsKeyArn"] = *r.Config.ConfigEncryption.KmsKeyArn
			}
			rConfig["configEncryption"] = rConfigConfigEncryption
		}
		if r.Config.IamInstanceProfile != nil {
			rConfig["iamInstanceProfile"] = *r.Config.IamInstanceProfile
		}
		if r.Config.ImageType != nil {
			rConfig["imageType"] = *r.Config.ImageType
		}
		if r.Config.InstancePlacement != nil && r.Config.InstancePlacement != dclService.EmptyNodePoolConfigInstancePlacement {
			rConfigInstancePlacement := make(map[string]interface{})
			if r.Config.InstancePlacement.Tenancy != nil {
				rConfigInstancePlacement["tenancy"] = string(*r.Config.InstancePlacement.Tenancy)
			}
			rConfig["instancePlacement"] = rConfigInstancePlacement
		}
		if r.Config.InstanceType != nil {
			rConfig["instanceType"] = *r.Config.InstanceType
		}
		if r.Config.Labels != nil {
			rConfigLabels := make(map[string]interface{})
			for k, v := range r.Config.Labels {
				rConfigLabels[k] = v
			}
			rConfig["labels"] = rConfigLabels
		}
		if r.Config.ProxyConfig != nil && r.Config.ProxyConfig != dclService.EmptyNodePoolConfigProxyConfig {
			rConfigProxyConfig := make(map[string]interface{})
			if r.Config.ProxyConfig.SecretArn != nil {
				rConfigProxyConfig["secretArn"] = *r.Config.ProxyConfig.SecretArn
			}
			if r.Config.ProxyConfig.SecretVersion != nil {
				rConfigProxyConfig["secretVersion"] = *r.Config.ProxyConfig.SecretVersion
			}
			rConfig["proxyConfig"] = rConfigProxyConfig
		}
		if r.Config.RootVolume != nil && r.Config.RootVolume != dclService.EmptyNodePoolConfigRootVolume {
			rConfigRootVolume := make(map[string]interface{})
			if r.Config.RootVolume.Iops != nil {
				rConfigRootVolume["iops"] = *r.Config.RootVolume.Iops
			}
			if r.Config.RootVolume.KmsKeyArn != nil {
				rConfigRootVolume["kmsKeyArn"] = *r.Config.RootVolume.KmsKeyArn
			}
			if r.Config.RootVolume.SizeGib != nil {
				rConfigRootVolume["sizeGib"] = *r.Config.RootVolume.SizeGib
			}
			if r.Config.RootVolume.Throughput != nil {
				rConfigRootVolume["throughput"] = *r.Config.RootVolume.Throughput
			}
			if r.Config.RootVolume.VolumeType != nil {
				rConfigRootVolume["volumeType"] = string(*r.Config.RootVolume.VolumeType)
			}
			rConfig["rootVolume"] = rConfigRootVolume
		}
		var rConfigSecurityGroupIds []interface{}
		for _, rConfigSecurityGroupIdsVal := range r.Config.SecurityGroupIds {
			rConfigSecurityGroupIds = append(rConfigSecurityGroupIds, rConfigSecurityGroupIdsVal)
		}
		rConfig["securityGroupIds"] = rConfigSecurityGroupIds
		if r.Config.SpotConfig != nil && r.Config.SpotConfig != dclService.EmptyNodePoolConfigSpotConfig {
			rConfigSpotConfig := make(map[string]interface{})
			var rConfigSpotConfigInstanceTypes []interface{}
			for _, rConfigSpotConfigInstanceTypesVal := range r.Config.SpotConfig.InstanceTypes {
				rConfigSpotConfigInstanceTypes = append(rConfigSpotConfigInstanceTypes, rConfigSpotConfigInstanceTypesVal)
			}
			rConfigSpotConfig["instanceTypes"] = rConfigSpotConfigInstanceTypes
			rConfig["spotConfig"] = rConfigSpotConfig
		}
		if r.Config.SshConfig != nil && r.Config.SshConfig != dclService.EmptyNodePoolConfigSshConfig {
			rConfigSshConfig := make(map[string]interface{})
			if r.Config.SshConfig.Ec2KeyPair != nil {
				rConfigSshConfig["ec2KeyPair"] = *r.Config.SshConfig.Ec2KeyPair
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
		var rConfigTaints []interface{}
		for _, rConfigTaintsVal := range r.Config.Taints {
			rConfigTaintsObject := make(map[string]interface{})
			if rConfigTaintsVal.Effect != nil {
				rConfigTaintsObject["effect"] = string(*rConfigTaintsVal.Effect)
			}
			if rConfigTaintsVal.Key != nil {
				rConfigTaintsObject["key"] = *rConfigTaintsVal.Key
			}
			if rConfigTaintsVal.Value != nil {
				rConfigTaintsObject["value"] = *rConfigTaintsVal.Value
			}
			rConfigTaints = append(rConfigTaints, rConfigTaintsObject)
		}
		rConfig["taints"] = rConfigTaints
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
	if r.UpdateSettings != nil && r.UpdateSettings != dclService.EmptyNodePoolUpdateSettings {
		rUpdateSettings := make(map[string]interface{})
		if r.UpdateSettings.SurgeSettings != nil && r.UpdateSettings.SurgeSettings != dclService.EmptyNodePoolUpdateSettingsSurgeSettings {
			rUpdateSettingsSurgeSettings := make(map[string]interface{})
			if r.UpdateSettings.SurgeSettings.MaxSurge != nil {
				rUpdateSettingsSurgeSettings["maxSurge"] = *r.UpdateSettings.SurgeSettings.MaxSurge
			}
			if r.UpdateSettings.SurgeSettings.MaxUnavailable != nil {
				rUpdateSettingsSurgeSettings["maxUnavailable"] = *r.UpdateSettings.SurgeSettings.MaxUnavailable
			}
			rUpdateSettings["surgeSettings"] = rUpdateSettingsSurgeSettings
		}
		u.Object["updateSettings"] = rUpdateSettings
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
			if _, ok := rConfig["autoscalingMetricsCollection"]; ok {
				if rConfigAutoscalingMetricsCollection, ok := rConfig["autoscalingMetricsCollection"].(map[string]interface{}); ok {
					r.Config.AutoscalingMetricsCollection = &dclService.NodePoolConfigAutoscalingMetricsCollection{}
					if _, ok := rConfigAutoscalingMetricsCollection["granularity"]; ok {
						if s, ok := rConfigAutoscalingMetricsCollection["granularity"].(string); ok {
							r.Config.AutoscalingMetricsCollection.Granularity = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.AutoscalingMetricsCollection.Granularity: expected string")
						}
					}
					if _, ok := rConfigAutoscalingMetricsCollection["metrics"]; ok {
						if s, ok := rConfigAutoscalingMetricsCollection["metrics"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Config.AutoscalingMetricsCollection.Metrics = append(r.Config.AutoscalingMetricsCollection.Metrics, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.AutoscalingMetricsCollection.Metrics: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.AutoscalingMetricsCollection: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["configEncryption"]; ok {
				if rConfigConfigEncryption, ok := rConfig["configEncryption"].(map[string]interface{}); ok {
					r.Config.ConfigEncryption = &dclService.NodePoolConfigConfigEncryption{}
					if _, ok := rConfigConfigEncryption["kmsKeyArn"]; ok {
						if s, ok := rConfigConfigEncryption["kmsKeyArn"].(string); ok {
							r.Config.ConfigEncryption.KmsKeyArn = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.ConfigEncryption.KmsKeyArn: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.ConfigEncryption: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["iamInstanceProfile"]; ok {
				if s, ok := rConfig["iamInstanceProfile"].(string); ok {
					r.Config.IamInstanceProfile = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Config.IamInstanceProfile: expected string")
				}
			}
			if _, ok := rConfig["imageType"]; ok {
				if s, ok := rConfig["imageType"].(string); ok {
					r.Config.ImageType = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Config.ImageType: expected string")
				}
			}
			if _, ok := rConfig["instancePlacement"]; ok {
				if rConfigInstancePlacement, ok := rConfig["instancePlacement"].(map[string]interface{}); ok {
					r.Config.InstancePlacement = &dclService.NodePoolConfigInstancePlacement{}
					if _, ok := rConfigInstancePlacement["tenancy"]; ok {
						if s, ok := rConfigInstancePlacement["tenancy"].(string); ok {
							r.Config.InstancePlacement.Tenancy = dclService.NodePoolConfigInstancePlacementTenancyEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Config.InstancePlacement.Tenancy: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.InstancePlacement: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["instanceType"]; ok {
				if s, ok := rConfig["instanceType"].(string); ok {
					r.Config.InstanceType = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.Config.InstanceType: expected string")
				}
			}
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
					if _, ok := rConfigProxyConfig["secretArn"]; ok {
						if s, ok := rConfigProxyConfig["secretArn"].(string); ok {
							r.Config.ProxyConfig.SecretArn = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.ProxyConfig.SecretArn: expected string")
						}
					}
					if _, ok := rConfigProxyConfig["secretVersion"]; ok {
						if s, ok := rConfigProxyConfig["secretVersion"].(string); ok {
							r.Config.ProxyConfig.SecretVersion = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.ProxyConfig.SecretVersion: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.ProxyConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["rootVolume"]; ok {
				if rConfigRootVolume, ok := rConfig["rootVolume"].(map[string]interface{}); ok {
					r.Config.RootVolume = &dclService.NodePoolConfigRootVolume{}
					if _, ok := rConfigRootVolume["iops"]; ok {
						if i, ok := rConfigRootVolume["iops"].(int64); ok {
							r.Config.RootVolume.Iops = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Config.RootVolume.Iops: expected int64")
						}
					}
					if _, ok := rConfigRootVolume["kmsKeyArn"]; ok {
						if s, ok := rConfigRootVolume["kmsKeyArn"].(string); ok {
							r.Config.RootVolume.KmsKeyArn = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.RootVolume.KmsKeyArn: expected string")
						}
					}
					if _, ok := rConfigRootVolume["sizeGib"]; ok {
						if i, ok := rConfigRootVolume["sizeGib"].(int64); ok {
							r.Config.RootVolume.SizeGib = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Config.RootVolume.SizeGib: expected int64")
						}
					}
					if _, ok := rConfigRootVolume["throughput"]; ok {
						if i, ok := rConfigRootVolume["throughput"].(int64); ok {
							r.Config.RootVolume.Throughput = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.Config.RootVolume.Throughput: expected int64")
						}
					}
					if _, ok := rConfigRootVolume["volumeType"]; ok {
						if s, ok := rConfigRootVolume["volumeType"].(string); ok {
							r.Config.RootVolume.VolumeType = dclService.NodePoolConfigRootVolumeVolumeTypeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.Config.RootVolume.VolumeType: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.RootVolume: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["securityGroupIds"]; ok {
				if s, ok := rConfig["securityGroupIds"].([]interface{}); ok {
					for _, ss := range s {
						if strval, ok := ss.(string); ok {
							r.Config.SecurityGroupIds = append(r.Config.SecurityGroupIds, strval)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.SecurityGroupIds: expected []interface{}")
				}
			}
			if _, ok := rConfig["spotConfig"]; ok {
				if rConfigSpotConfig, ok := rConfig["spotConfig"].(map[string]interface{}); ok {
					r.Config.SpotConfig = &dclService.NodePoolConfigSpotConfig{}
					if _, ok := rConfigSpotConfig["instanceTypes"]; ok {
						if s, ok := rConfigSpotConfig["instanceTypes"].([]interface{}); ok {
							for _, ss := range s {
								if strval, ok := ss.(string); ok {
									r.Config.SpotConfig.InstanceTypes = append(r.Config.SpotConfig.InstanceTypes, strval)
								}
							}
						} else {
							return nil, fmt.Errorf("r.Config.SpotConfig.InstanceTypes: expected []interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.SpotConfig: expected map[string]interface{}")
				}
			}
			if _, ok := rConfig["sshConfig"]; ok {
				if rConfigSshConfig, ok := rConfig["sshConfig"].(map[string]interface{}); ok {
					r.Config.SshConfig = &dclService.NodePoolConfigSshConfig{}
					if _, ok := rConfigSshConfig["ec2KeyPair"]; ok {
						if s, ok := rConfigSshConfig["ec2KeyPair"].(string); ok {
							r.Config.SshConfig.Ec2KeyPair = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Config.SshConfig.Ec2KeyPair: expected string")
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
			if _, ok := rConfig["taints"]; ok {
				if s, ok := rConfig["taints"].([]interface{}); ok {
					for _, o := range s {
						if objval, ok := o.(map[string]interface{}); ok {
							var rConfigTaints dclService.NodePoolConfigTaints
							if _, ok := objval["effect"]; ok {
								if s, ok := objval["effect"].(string); ok {
									rConfigTaints.Effect = dclService.NodePoolConfigTaintsEffectEnumRef(s)
								} else {
									return nil, fmt.Errorf("rConfigTaints.Effect: expected string")
								}
							}
							if _, ok := objval["key"]; ok {
								if s, ok := objval["key"].(string); ok {
									rConfigTaints.Key = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rConfigTaints.Key: expected string")
								}
							}
							if _, ok := objval["value"]; ok {
								if s, ok := objval["value"].(string); ok {
									rConfigTaints.Value = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rConfigTaints.Value: expected string")
								}
							}
							r.Config.Taints = append(r.Config.Taints, rConfigTaints)
						}
					}
				} else {
					return nil, fmt.Errorf("r.Config.Taints: expected []interface{}")
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
	if _, ok := u.Object["updateSettings"]; ok {
		if rUpdateSettings, ok := u.Object["updateSettings"].(map[string]interface{}); ok {
			r.UpdateSettings = &dclService.NodePoolUpdateSettings{}
			if _, ok := rUpdateSettings["surgeSettings"]; ok {
				if rUpdateSettingsSurgeSettings, ok := rUpdateSettings["surgeSettings"].(map[string]interface{}); ok {
					r.UpdateSettings.SurgeSettings = &dclService.NodePoolUpdateSettingsSurgeSettings{}
					if _, ok := rUpdateSettingsSurgeSettings["maxSurge"]; ok {
						if i, ok := rUpdateSettingsSurgeSettings["maxSurge"].(int64); ok {
							r.UpdateSettings.SurgeSettings.MaxSurge = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.UpdateSettings.SurgeSettings.MaxSurge: expected int64")
						}
					}
					if _, ok := rUpdateSettingsSurgeSettings["maxUnavailable"]; ok {
						if i, ok := rUpdateSettingsSurgeSettings["maxUnavailable"].(int64); ok {
							r.UpdateSettings.SurgeSettings.MaxUnavailable = dcl.Int64(i)
						} else {
							return nil, fmt.Errorf("r.UpdateSettings.SurgeSettings.MaxUnavailable: expected int64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.UpdateSettings.SurgeSettings: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.UpdateSettings: expected map[string]interface{}")
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
		"containeraws",
		"NodePool",
		"beta",
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
