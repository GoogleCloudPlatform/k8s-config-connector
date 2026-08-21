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
package dataproc

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type AutoscalingPolicy struct{}

func AutoscalingPolicyToUnstructured(r *dclService.AutoscalingPolicy) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "dataproc",
			Version: "alpha",
			Type:    "AutoscalingPolicy",
		},
		Object: make(map[string]interface{}),
	}
	if r.BasicAlgorithm != nil && r.BasicAlgorithm != dclService.EmptyAutoscalingPolicyBasicAlgorithm {
		rBasicAlgorithm := make(map[string]interface{})
		if r.BasicAlgorithm.CooldownPeriod != nil {
			rBasicAlgorithm["cooldownPeriod"] = *r.BasicAlgorithm.CooldownPeriod
		}
		if r.BasicAlgorithm.YarnConfig != nil && r.BasicAlgorithm.YarnConfig != dclService.EmptyAutoscalingPolicyBasicAlgorithmYarnConfig {
			rBasicAlgorithmYarnConfig := make(map[string]interface{})
			if r.BasicAlgorithm.YarnConfig.GracefulDecommissionTimeout != nil {
				rBasicAlgorithmYarnConfig["gracefulDecommissionTimeout"] = *r.BasicAlgorithm.YarnConfig.GracefulDecommissionTimeout
			}
			if r.BasicAlgorithm.YarnConfig.ScaleDownFactor != nil {
				rBasicAlgorithmYarnConfig["scaleDownFactor"] = *r.BasicAlgorithm.YarnConfig.ScaleDownFactor
			}
			if r.BasicAlgorithm.YarnConfig.ScaleDownMinWorkerFraction != nil {
				rBasicAlgorithmYarnConfig["scaleDownMinWorkerFraction"] = *r.BasicAlgorithm.YarnConfig.ScaleDownMinWorkerFraction
			}
			if r.BasicAlgorithm.YarnConfig.ScaleUpFactor != nil {
				rBasicAlgorithmYarnConfig["scaleUpFactor"] = *r.BasicAlgorithm.YarnConfig.ScaleUpFactor
			}
			if r.BasicAlgorithm.YarnConfig.ScaleUpMinWorkerFraction != nil {
				rBasicAlgorithmYarnConfig["scaleUpMinWorkerFraction"] = *r.BasicAlgorithm.YarnConfig.ScaleUpMinWorkerFraction
			}
			rBasicAlgorithm["yarnConfig"] = rBasicAlgorithmYarnConfig
		}
		u.Object["basicAlgorithm"] = rBasicAlgorithm
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
	if r.SecondaryWorkerConfig != nil && r.SecondaryWorkerConfig != dclService.EmptyAutoscalingPolicySecondaryWorkerConfig {
		rSecondaryWorkerConfig := make(map[string]interface{})
		if r.SecondaryWorkerConfig.MaxInstances != nil {
			rSecondaryWorkerConfig["maxInstances"] = *r.SecondaryWorkerConfig.MaxInstances
		}
		if r.SecondaryWorkerConfig.MinInstances != nil {
			rSecondaryWorkerConfig["minInstances"] = *r.SecondaryWorkerConfig.MinInstances
		}
		if r.SecondaryWorkerConfig.Weight != nil {
			rSecondaryWorkerConfig["weight"] = *r.SecondaryWorkerConfig.Weight
		}
		u.Object["secondaryWorkerConfig"] = rSecondaryWorkerConfig
	}
	if r.WorkerConfig != nil && r.WorkerConfig != dclService.EmptyAutoscalingPolicyWorkerConfig {
		rWorkerConfig := make(map[string]interface{})
		if r.WorkerConfig.MaxInstances != nil {
			rWorkerConfig["maxInstances"] = *r.WorkerConfig.MaxInstances
		}
		if r.WorkerConfig.MinInstances != nil {
			rWorkerConfig["minInstances"] = *r.WorkerConfig.MinInstances
		}
		if r.WorkerConfig.Weight != nil {
			rWorkerConfig["weight"] = *r.WorkerConfig.Weight
		}
		u.Object["workerConfig"] = rWorkerConfig
	}
	return u
}

func UnstructuredToAutoscalingPolicy(u *unstructured.Resource) (*dclService.AutoscalingPolicy, error) {
	r := &dclService.AutoscalingPolicy{}
	if _, ok := u.Object["basicAlgorithm"]; ok {
		if rBasicAlgorithm, ok := u.Object["basicAlgorithm"].(map[string]interface{}); ok {
			r.BasicAlgorithm = &dclService.AutoscalingPolicyBasicAlgorithm{}
			if _, ok := rBasicAlgorithm["cooldownPeriod"]; ok {
				if s, ok := rBasicAlgorithm["cooldownPeriod"].(string); ok {
					r.BasicAlgorithm.CooldownPeriod = dcl.String(s)
				} else {
					return nil, fmt.Errorf("r.BasicAlgorithm.CooldownPeriod: expected string")
				}
			}
			if _, ok := rBasicAlgorithm["yarnConfig"]; ok {
				if rBasicAlgorithmYarnConfig, ok := rBasicAlgorithm["yarnConfig"].(map[string]interface{}); ok {
					r.BasicAlgorithm.YarnConfig = &dclService.AutoscalingPolicyBasicAlgorithmYarnConfig{}
					if _, ok := rBasicAlgorithmYarnConfig["gracefulDecommissionTimeout"]; ok {
						if s, ok := rBasicAlgorithmYarnConfig["gracefulDecommissionTimeout"].(string); ok {
							r.BasicAlgorithm.YarnConfig.GracefulDecommissionTimeout = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.BasicAlgorithm.YarnConfig.GracefulDecommissionTimeout: expected string")
						}
					}
					if _, ok := rBasicAlgorithmYarnConfig["scaleDownFactor"]; ok {
						if f, ok := rBasicAlgorithmYarnConfig["scaleDownFactor"].(float64); ok {
							r.BasicAlgorithm.YarnConfig.ScaleDownFactor = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.BasicAlgorithm.YarnConfig.ScaleDownFactor: expected float64")
						}
					}
					if _, ok := rBasicAlgorithmYarnConfig["scaleDownMinWorkerFraction"]; ok {
						if f, ok := rBasicAlgorithmYarnConfig["scaleDownMinWorkerFraction"].(float64); ok {
							r.BasicAlgorithm.YarnConfig.ScaleDownMinWorkerFraction = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.BasicAlgorithm.YarnConfig.ScaleDownMinWorkerFraction: expected float64")
						}
					}
					if _, ok := rBasicAlgorithmYarnConfig["scaleUpFactor"]; ok {
						if f, ok := rBasicAlgorithmYarnConfig["scaleUpFactor"].(float64); ok {
							r.BasicAlgorithm.YarnConfig.ScaleUpFactor = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.BasicAlgorithm.YarnConfig.ScaleUpFactor: expected float64")
						}
					}
					if _, ok := rBasicAlgorithmYarnConfig["scaleUpMinWorkerFraction"]; ok {
						if f, ok := rBasicAlgorithmYarnConfig["scaleUpMinWorkerFraction"].(float64); ok {
							r.BasicAlgorithm.YarnConfig.ScaleUpMinWorkerFraction = dcl.Float64(f)
						} else {
							return nil, fmt.Errorf("r.BasicAlgorithm.YarnConfig.ScaleUpMinWorkerFraction: expected float64")
						}
					}
				} else {
					return nil, fmt.Errorf("r.BasicAlgorithm.YarnConfig: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.BasicAlgorithm: expected map[string]interface{}")
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
	if _, ok := u.Object["secondaryWorkerConfig"]; ok {
		if rSecondaryWorkerConfig, ok := u.Object["secondaryWorkerConfig"].(map[string]interface{}); ok {
			r.SecondaryWorkerConfig = &dclService.AutoscalingPolicySecondaryWorkerConfig{}
			if _, ok := rSecondaryWorkerConfig["maxInstances"]; ok {
				if i, ok := rSecondaryWorkerConfig["maxInstances"].(int64); ok {
					r.SecondaryWorkerConfig.MaxInstances = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.SecondaryWorkerConfig.MaxInstances: expected int64")
				}
			}
			if _, ok := rSecondaryWorkerConfig["minInstances"]; ok {
				if i, ok := rSecondaryWorkerConfig["minInstances"].(int64); ok {
					r.SecondaryWorkerConfig.MinInstances = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.SecondaryWorkerConfig.MinInstances: expected int64")
				}
			}
			if _, ok := rSecondaryWorkerConfig["weight"]; ok {
				if i, ok := rSecondaryWorkerConfig["weight"].(int64); ok {
					r.SecondaryWorkerConfig.Weight = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.SecondaryWorkerConfig.Weight: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.SecondaryWorkerConfig: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["workerConfig"]; ok {
		if rWorkerConfig, ok := u.Object["workerConfig"].(map[string]interface{}); ok {
			r.WorkerConfig = &dclService.AutoscalingPolicyWorkerConfig{}
			if _, ok := rWorkerConfig["maxInstances"]; ok {
				if i, ok := rWorkerConfig["maxInstances"].(int64); ok {
					r.WorkerConfig.MaxInstances = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.WorkerConfig.MaxInstances: expected int64")
				}
			}
			if _, ok := rWorkerConfig["minInstances"]; ok {
				if i, ok := rWorkerConfig["minInstances"].(int64); ok {
					r.WorkerConfig.MinInstances = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.WorkerConfig.MinInstances: expected int64")
				}
			}
			if _, ok := rWorkerConfig["weight"]; ok {
				if i, ok := rWorkerConfig["weight"].(int64); ok {
					r.WorkerConfig.Weight = dcl.Int64(i)
				} else {
					return nil, fmt.Errorf("r.WorkerConfig.Weight: expected int64")
				}
			}
		} else {
			return nil, fmt.Errorf("r.WorkerConfig: expected map[string]interface{}")
		}
	}
	return r, nil
}

func GetAutoscalingPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAutoscalingPolicy(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetAutoscalingPolicy(ctx, r)
	if err != nil {
		return nil, err
	}
	return AutoscalingPolicyToUnstructured(r), nil
}

func ListAutoscalingPolicy(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListAutoscalingPolicy(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, AutoscalingPolicyToUnstructured(r))
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

func ApplyAutoscalingPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAutoscalingPolicy(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAutoscalingPolicy(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyAutoscalingPolicy(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return AutoscalingPolicyToUnstructured(r), nil
}

func AutoscalingPolicyHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAutoscalingPolicy(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToAutoscalingPolicy(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyAutoscalingPolicy(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteAutoscalingPolicy(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToAutoscalingPolicy(u)
	if err != nil {
		return err
	}
	return c.DeleteAutoscalingPolicy(ctx, r)
}

func AutoscalingPolicyID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToAutoscalingPolicy(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *AutoscalingPolicy) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"dataproc",
		"AutoscalingPolicy",
		"alpha",
	}
}

func (r *AutoscalingPolicy) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AutoscalingPolicy) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AutoscalingPolicy) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *AutoscalingPolicy) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AutoscalingPolicy) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AutoscalingPolicy) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *AutoscalingPolicy) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetAutoscalingPolicy(ctx, config, resource)
}

func (r *AutoscalingPolicy) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyAutoscalingPolicy(ctx, config, resource, opts...)
}

func (r *AutoscalingPolicy) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return AutoscalingPolicyHasDiff(ctx, config, resource, opts...)
}

func (r *AutoscalingPolicy) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteAutoscalingPolicy(ctx, config, resource)
}

func (r *AutoscalingPolicy) ID(resource *unstructured.Resource) (string, error) {
	return AutoscalingPolicyID(resource)
}

func init() {
	unstructured.Register(&AutoscalingPolicy{})
}
