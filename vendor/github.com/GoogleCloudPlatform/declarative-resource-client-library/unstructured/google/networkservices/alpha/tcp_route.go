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
package networkservices

import (
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type TcpRoute struct{}

func TcpRouteToUnstructured(r *dclService.TcpRoute) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networkservices",
			Version: "alpha",
			Type:    "TcpRoute",
		},
		Object: make(map[string]interface{}),
	}
	if r.CreateTime != nil {
		u.Object["createTime"] = *r.CreateTime
	}
	if r.Description != nil {
		u.Object["description"] = *r.Description
	}
	var rGateways []interface{}
	for _, rGatewaysVal := range r.Gateways {
		rGateways = append(rGateways, rGatewaysVal)
	}
	u.Object["gateways"] = rGateways
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
	var rMeshes []interface{}
	for _, rMeshesVal := range r.Meshes {
		rMeshes = append(rMeshes, rMeshesVal)
	}
	u.Object["meshes"] = rMeshes
	if r.Name != nil {
		u.Object["name"] = *r.Name
	}
	if r.Project != nil {
		u.Object["project"] = *r.Project
	}
	var rRules []interface{}
	for _, rRulesVal := range r.Rules {
		rRulesObject := make(map[string]interface{})
		if rRulesVal.Action != nil && rRulesVal.Action != dclService.EmptyTcpRouteRulesAction {
			rRulesValAction := make(map[string]interface{})
			var rRulesValActionDestinations []interface{}
			for _, rRulesValActionDestinationsVal := range rRulesVal.Action.Destinations {
				rRulesValActionDestinationsObject := make(map[string]interface{})
				if rRulesValActionDestinationsVal.ServiceName != nil {
					rRulesValActionDestinationsObject["serviceName"] = *rRulesValActionDestinationsVal.ServiceName
				}
				if rRulesValActionDestinationsVal.Weight != nil {
					rRulesValActionDestinationsObject["weight"] = *rRulesValActionDestinationsVal.Weight
				}
				rRulesValActionDestinations = append(rRulesValActionDestinations, rRulesValActionDestinationsObject)
			}
			rRulesValAction["destinations"] = rRulesValActionDestinations
			if rRulesVal.Action.OriginalDestination != nil {
				rRulesValAction["originalDestination"] = *rRulesVal.Action.OriginalDestination
			}
			rRulesObject["action"] = rRulesValAction
		}
		var rRulesValMatches []interface{}
		for _, rRulesValMatchesVal := range rRulesVal.Matches {
			rRulesValMatchesObject := make(map[string]interface{})
			if rRulesValMatchesVal.Address != nil {
				rRulesValMatchesObject["address"] = *rRulesValMatchesVal.Address
			}
			if rRulesValMatchesVal.Port != nil {
				rRulesValMatchesObject["port"] = *rRulesValMatchesVal.Port
			}
			rRulesValMatches = append(rRulesValMatches, rRulesValMatchesObject)
		}
		rRulesObject["matches"] = rRulesValMatches
		rRules = append(rRules, rRulesObject)
	}
	u.Object["rules"] = rRules
	if r.SelfLink != nil {
		u.Object["selfLink"] = *r.SelfLink
	}
	if r.UpdateTime != nil {
		u.Object["updateTime"] = *r.UpdateTime
	}
	return u
}

func UnstructuredToTcpRoute(u *unstructured.Resource) (*dclService.TcpRoute, error) {
	r := &dclService.TcpRoute{}
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
	if _, ok := u.Object["gateways"]; ok {
		if s, ok := u.Object["gateways"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.Gateways = append(r.Gateways, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Gateways: expected []interface{}")
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
	if _, ok := u.Object["meshes"]; ok {
		if s, ok := u.Object["meshes"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.Meshes = append(r.Meshes, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Meshes: expected []interface{}")
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
	if _, ok := u.Object["rules"]; ok {
		if s, ok := u.Object["rules"].([]interface{}); ok {
			for _, o := range s {
				if objval, ok := o.(map[string]interface{}); ok {
					var rRules dclService.TcpRouteRules
					if _, ok := objval["action"]; ok {
						if rRulesAction, ok := objval["action"].(map[string]interface{}); ok {
							rRules.Action = &dclService.TcpRouteRulesAction{}
							if _, ok := rRulesAction["destinations"]; ok {
								if s, ok := rRulesAction["destinations"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rRulesActionDestinations dclService.TcpRouteRulesActionDestinations
											if _, ok := objval["serviceName"]; ok {
												if s, ok := objval["serviceName"].(string); ok {
													rRulesActionDestinations.ServiceName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRulesActionDestinations.ServiceName: expected string")
												}
											}
											if _, ok := objval["weight"]; ok {
												if i, ok := objval["weight"].(int64); ok {
													rRulesActionDestinations.Weight = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rRulesActionDestinations.Weight: expected int64")
												}
											}
											rRules.Action.Destinations = append(rRules.Action.Destinations, rRulesActionDestinations)
										}
									}
								} else {
									return nil, fmt.Errorf("rRules.Action.Destinations: expected []interface{}")
								}
							}
							if _, ok := rRulesAction["originalDestination"]; ok {
								if b, ok := rRulesAction["originalDestination"].(bool); ok {
									rRules.Action.OriginalDestination = dcl.Bool(b)
								} else {
									return nil, fmt.Errorf("rRules.Action.OriginalDestination: expected bool")
								}
							}
						} else {
							return nil, fmt.Errorf("rRules.Action: expected map[string]interface{}")
						}
					}
					if _, ok := objval["matches"]; ok {
						if s, ok := objval["matches"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rRulesMatches dclService.TcpRouteRulesMatches
									if _, ok := objval["address"]; ok {
										if s, ok := objval["address"].(string); ok {
											rRulesMatches.Address = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRulesMatches.Address: expected string")
										}
									}
									if _, ok := objval["port"]; ok {
										if s, ok := objval["port"].(string); ok {
											rRulesMatches.Port = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRulesMatches.Port: expected string")
										}
									}
									rRules.Matches = append(rRules.Matches, rRulesMatches)
								}
							}
						} else {
							return nil, fmt.Errorf("rRules.Matches: expected []interface{}")
						}
					}
					r.Rules = append(r.Rules, rRules)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Rules: expected []interface{}")
		}
	}
	if _, ok := u.Object["selfLink"]; ok {
		if s, ok := u.Object["selfLink"].(string); ok {
			r.SelfLink = dcl.String(s)
		} else {
			return nil, fmt.Errorf("r.SelfLink: expected string")
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

func GetTcpRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTcpRoute(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetTcpRoute(ctx, r)
	if err != nil {
		return nil, err
	}
	return TcpRouteToUnstructured(r), nil
}

func ListTcpRoute(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListTcpRoute(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, TcpRouteToUnstructured(r))
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

func ApplyTcpRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTcpRoute(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTcpRoute(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyTcpRoute(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return TcpRouteToUnstructured(r), nil
}

func TcpRouteHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTcpRoute(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTcpRoute(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyTcpRoute(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteTcpRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTcpRoute(u)
	if err != nil {
		return err
	}
	return c.DeleteTcpRoute(ctx, r)
}

func TcpRouteID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToTcpRoute(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *TcpRoute) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networkservices",
		"TcpRoute",
		"alpha",
	}
}

func (r *TcpRoute) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TcpRoute) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TcpRoute) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *TcpRoute) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TcpRoute) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TcpRoute) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TcpRoute) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetTcpRoute(ctx, config, resource)
}

func (r *TcpRoute) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyTcpRoute(ctx, config, resource, opts...)
}

func (r *TcpRoute) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return TcpRouteHasDiff(ctx, config, resource, opts...)
}

func (r *TcpRoute) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteTcpRoute(ctx, config, resource)
}

func (r *TcpRoute) ID(resource *unstructured.Resource) (string, error) {
	return TcpRouteID(resource)
}

func init() {
	unstructured.Register(&TcpRoute{})
}
