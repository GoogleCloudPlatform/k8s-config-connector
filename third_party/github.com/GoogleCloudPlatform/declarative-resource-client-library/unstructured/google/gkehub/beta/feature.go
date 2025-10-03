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
package gkehub

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type Feature struct{}

func FeatureToUnstructured(r *dclService.Feature) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "gkehub",
			Version: "beta",
			Type:    "Feature",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.DeleteTime != nil {
		u.Object["deleteTime"] = *r.DeleteTime
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
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	if r.ResourceState != nil && r.ResourceState != dclService.EmptyFeatureResourceState {
		rResourceState := make(map[string]interface{})
		if r.ResourceState.HasResources != nil {
			rResourceState["hasResources"] = *r.ResourceState.HasResources
		}
		if r.ResourceState.State != nil {
			rResourceState["state"] = string(*r.ResourceState.State)
		}
		u.Object["resourceState"] = rResourceState
	}
	if r.Spec != nil && r.Spec != dclService.EmptyFeatureSpec {
		rSpec := make(map[string]interface{})
		if r.Spec.Fleetobservability != nil && r.Spec.Fleetobservability != dclService.EmptyFeatureSpecFleetobservability {
			rSpecFleetobservability := make(map[string]interface{})
			if r.Spec.Fleetobservability.LoggingConfig != nil && r.Spec.Fleetobservability.LoggingConfig != dclService.EmptyFeatureSpecFleetobservabilityLoggingConfig {
				rSpecFleetobservabilityLoggingConfig := make(map[string]interface{})
				if r.Spec.Fleetobservability.LoggingConfig.DefaultConfig != nil && r.Spec.Fleetobservability.LoggingConfig.DefaultConfig != dclService.EmptyFeatureSpecFleetobservabilityLoggingConfigDefaultConfig {
					rSpecFleetobservabilityLoggingConfigDefaultConfig := make(map[string]interface{})
					if r.Spec.Fleetobservability.LoggingConfig.DefaultConfig.Mode != nil {
						rSpecFleetobservabilityLoggingConfigDefaultConfig["mode"] = string(*r.Spec.Fleetobservability.LoggingConfig.DefaultConfig.Mode)
					}
					rSpecFleetobservabilityLoggingConfig["defaultConfig"] = rSpecFleetobservabilityLoggingConfigDefaultConfig
				}
				if r.Spec.Fleetobservability.LoggingConfig.FleetScopeLogsConfig != nil && r.Spec.Fleetobservability.LoggingConfig.FleetScopeLogsConfig != dclService.EmptyFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig {
					rSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig := make(map[string]interface{})
					if r.Spec.Fleetobservability.LoggingConfig.FleetScopeLogsConfig.Mode != nil {
						rSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig["mode"] = string(*r.Spec.Fleetobservability.LoggingConfig.FleetScopeLogsConfig.Mode)
					}
					rSpecFleetobservabilityLoggingConfig["fleetScopeLogsConfig"] = rSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig
				}
				rSpecFleetobservability["loggingConfig"] = rSpecFleetobservabilityLoggingConfig
			}
			rSpec["fleetobservability"] = rSpecFleetobservability
		}
		if r.Spec.Multiclusteringress != nil && r.Spec.Multiclusteringress != dclService.EmptyFeatureSpecMulticlusteringress {
			rSpecMulticlusteringress := make(map[string]interface{})
			if r.Spec.Multiclusteringress.ConfigMembership != nil {
				rSpecMulticlusteringress["configMembership"] = *r.Spec.Multiclusteringress.ConfigMembership
			}
			rSpec["multiclusteringress"] = rSpecMulticlusteringress
		}
		u.Object["spec"] = rSpec
	}
	if r.State != nil && r.State != dclService.EmptyFeatureState {
		rState := make(map[string]interface{})
		if r.State.State != nil && r.State.State != dclService.EmptyFeatureStateState {
			rStateState := make(map[string]interface{})
			if r.State.State.Code != nil {
				rStateState["code"] = string(*r.State.State.Code)
			}
			if r.State.State.Description != nil {
				rStateState["description"] = *r.State.State.Description
			}
			if r.State.State.UpdateTime != nil {
				rStateState["updateTime"] = *r.State.State.UpdateTime
			}
			rState["state"] = rStateState
		}
		u.Object["state"] = rState
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToFeature(u *unstructured.Resource) (*dclService.Feature, error) {
	r := &dclService.Feature{}
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
	if _, ok := u.Object["project"]; ok {
		if s, ok := u.Object["project"].(string); ok {
			r.Project = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.Project: expected string")
		}
	}
	if _, ok := u.Object["resourceState"]; ok {
		if rResourceState, ok := u.Object["resourceState"].(map[string]interface{}); ok {
			r.ResourceState = &dclService.FeatureResourceState{}
			if _, ok := rResourceState["hasResources"]; ok {
				if b, ok := rResourceState["hasResources"].(bool); ok {
					r.ResourceState.HasResources = dcl.Bool(b)
				} else {
					return nil, fmt.Errorf("r.ResourceState.HasResources: expected bool")
				}
			}
			if _, ok := rResourceState["state"]; ok {
				if s, ok := rResourceState["state"].(string); ok {
					r.ResourceState.State = dclService.FeatureResourceStateStateEnumRef(s)
				} else {
					return nil, fmt.Errorf("r.ResourceState.State: expected string")
				}
			}
		} else {
			return nil, fmt.Errorf("r.ResourceState: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["spec"]; ok {
		if rSpec, ok := u.Object["spec"].(map[string]interface{}); ok {
			r.Spec = &dclService.FeatureSpec{}
			if _, ok := rSpec["fleetobservability"]; ok {
				if rSpecFleetobservability, ok := rSpec["fleetobservability"].(map[string]interface{}); ok {
					r.Spec.Fleetobservability = &dclService.FeatureSpecFleetobservability{}
					if _, ok := rSpecFleetobservability["loggingConfig"]; ok {
						if rSpecFleetobservabilityLoggingConfig, ok := rSpecFleetobservability["loggingConfig"].(map[string]interface{}); ok {
							r.Spec.Fleetobservability.LoggingConfig = &dclService.FeatureSpecFleetobservabilityLoggingConfig{}
							if _, ok := rSpecFleetobservabilityLoggingConfig["defaultConfig"]; ok {
								if rSpecFleetobservabilityLoggingConfigDefaultConfig, ok := rSpecFleetobservabilityLoggingConfig["defaultConfig"].(map[string]interface{}); ok {
									r.Spec.Fleetobservability.LoggingConfig.DefaultConfig = &dclService.FeatureSpecFleetobservabilityLoggingConfigDefaultConfig{}
									if _, ok := rSpecFleetobservabilityLoggingConfigDefaultConfig["mode"]; ok {
										if s, ok := rSpecFleetobservabilityLoggingConfigDefaultConfig["mode"].(string); ok {
											r.Spec.Fleetobservability.LoggingConfig.DefaultConfig.Mode = dclService.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigModeEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Spec.Fleetobservability.LoggingConfig.DefaultConfig.Mode: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Spec.Fleetobservability.LoggingConfig.DefaultConfig: expected map[string]interface{}")
								}
							}
							if _, ok := rSpecFleetobservabilityLoggingConfig["fleetScopeLogsConfig"]; ok {
								if rSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig, ok := rSpecFleetobservabilityLoggingConfig["fleetScopeLogsConfig"].(map[string]interface{}); ok {
									r.Spec.Fleetobservability.LoggingConfig.FleetScopeLogsConfig = &dclService.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig{}
									if _, ok := rSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig["mode"]; ok {
										if s, ok := rSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig["mode"].(string); ok {
											r.Spec.Fleetobservability.LoggingConfig.FleetScopeLogsConfig.Mode = dclService.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigModeEnumRef(s)
										} else {
											return nil, fmt.Errorf("r.Spec.Fleetobservability.LoggingConfig.FleetScopeLogsConfig.Mode: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("r.Spec.Fleetobservability.LoggingConfig.FleetScopeLogsConfig: expected map[string]interface{}")
								}
							}
						} else {
							return nil, fmt.Errorf("r.Spec.Fleetobservability.LoggingConfig: expected map[string]interface{}")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Spec.Fleetobservability: expected map[string]interface{}")
				}
			}
			if _, ok := rSpec["multiclusteringress"]; ok {
				if rSpecMulticlusteringress, ok := rSpec["multiclusteringress"].(map[string]interface{}); ok {
					r.Spec.Multiclusteringress = &dclService.FeatureSpecMulticlusteringress{}
					if _, ok := rSpecMulticlusteringress["configMembership"]; ok {
						if s, ok := rSpecMulticlusteringress["configMembership"].(string); ok {
							r.Spec.Multiclusteringress.ConfigMembership = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.Spec.Multiclusteringress.ConfigMembership: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.Spec.Multiclusteringress: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.Spec: expected map[string]interface{}")
		}
	}
	if _, ok := u.Object["state"]; ok {
		if rState, ok := u.Object["state"].(map[string]interface{}); ok {
			r.State = &dclService.FeatureState{}
			if _, ok := rState["state"]; ok {
				if rStateState, ok := rState["state"].(map[string]interface{}); ok {
					r.State.State = &dclService.FeatureStateState{}
					if _, ok := rStateState["code"]; ok {
						if s, ok := rStateState["code"].(string); ok {
							r.State.State.Code = dclService.FeatureStateStateCodeEnumRef(s)
						} else {
							return nil, fmt.Errorf("r.State.State.Code: expected string")
						}
					}
					if _, ok := rStateState["description"]; ok {
						if s, ok := rStateState["description"].(string); ok {
							r.State.State.Description = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.State.State.Description: expected string")
						}
					}
					if _, ok := rStateState["updateTime"]; ok {
						if s, ok := rStateState["updateTime"].(string); ok {
							r.State.State.UpdateTime = dcl.String(s)
						} else {
							return nil, fmt.Errorf("r.State.State.UpdateTime: expected string")
						}
					}
				} else {
					return nil, fmt.Errorf("r.State.State: expected map[string]interface{}")
				}
			}
		} else {
			return nil, fmt.Errorf("r.State: expected map[string]interface{}")
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

func GetFeature(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFeature(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetFeature(ctx, r)
	if err != nil {
		return nil, err
	}
	return FeatureToUnstructured(r), nil
}

func ListFeature(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListFeature(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, FeatureToUnstructured(r))
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

func ApplyFeature(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFeature(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFeature(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyFeature(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return FeatureToUnstructured(r), nil
}

func FeatureHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFeature(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToFeature(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyFeature(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteFeature(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToFeature(u)
	if err != nil {
		return err
	}
	return c.DeleteFeature(ctx, r)
}

func FeatureID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToFeature(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *Feature) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"gkehub",
		"Feature",
		"beta",
	}
}

func (r *Feature) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Feature) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Feature) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *Feature) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Feature) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Feature) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *Feature) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetFeature(ctx, config, resource)
}

func (r *Feature) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyFeature(ctx, config, resource, opts...)
}

func (r *Feature) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return FeatureHasDiff(ctx, config, resource, opts...)
}

func (r *Feature) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteFeature(ctx, config, resource)
}

func (r *Feature) ID(resource *unstructured.Resource) (string, error) {
	return FeatureID(resource)
}

func init() {
	unstructured.Register(&Feature{})
}
