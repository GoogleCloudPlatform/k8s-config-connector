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

type TlsRoute struct{}

func TlsRouteToUnstructured(r *dclService.TlsRoute) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networkservices",
			Version: "alpha",
			Type:    "TlsRoute",
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
		if rRulesVal.Action != nil && rRulesVal.Action != dclService.EmptyTlsRouteRulesAction {
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
			rRulesObject["action"] = rRulesValAction
		}
		var rRulesValMatches []interface{}
		for _, rRulesValMatchesVal := range rRulesVal.Matches {
			rRulesValMatchesObject := make(map[string]interface{})
			var rRulesValMatchesValAlpn []interface{}
			for _, rRulesValMatchesValAlpnVal := range rRulesValMatchesVal.Alpn {
				rRulesValMatchesValAlpn = append(rRulesValMatchesValAlpn, rRulesValMatchesValAlpnVal)
			}
			rRulesValMatchesObject["alpn"] = rRulesValMatchesValAlpn
			var rRulesValMatchesValSniHost []interface{}
			for _, rRulesValMatchesValSniHostVal := range rRulesValMatchesVal.SniHost {
				rRulesValMatchesValSniHost = append(rRulesValMatchesValSniHost, rRulesValMatchesValSniHostVal)
			}
			rRulesValMatchesObject["sniHost"] = rRulesValMatchesValSniHost
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

func UnstructuredToTlsRoute(u *unstructured.Resource) (*dclService.TlsRoute, error) {
	r := &dclService.TlsRoute{}
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
					var rRules dclService.TlsRouteRules
					if _, ok := objval["action"]; ok {
						if rRulesAction, ok := objval["action"].(map[string]interface{}); ok {
							rRules.Action = &dclService.TlsRouteRulesAction{}
							if _, ok := rRulesAction["destinations"]; ok {
								if s, ok := rRulesAction["destinations"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rRulesActionDestinations dclService.TlsRouteRulesActionDestinations
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
						} else {
							return nil, fmt.Errorf("rRules.Action: expected map[string]interface{}")
						}
					}
					if _, ok := objval["matches"]; ok {
						if s, ok := objval["matches"].([]interface{}); ok {
							for _, o := range s {
								if objval, ok := o.(map[string]interface{}); ok {
									var rRulesMatches dclService.TlsRouteRulesMatches
									if _, ok := objval["alpn"]; ok {
										if s, ok := objval["alpn"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRulesMatches.Alpn = append(rRulesMatches.Alpn, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesMatches.Alpn: expected []interface{}")
										}
									}
									if _, ok := objval["sniHost"]; ok {
										if s, ok := objval["sniHost"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRulesMatches.SniHost = append(rRulesMatches.SniHost, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesMatches.SniHost: expected []interface{}")
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

func GetTlsRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTlsRoute(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetTlsRoute(ctx, r)
	if err != nil {
		return nil, err
	}
	return TlsRouteToUnstructured(r), nil
}

func ListTlsRoute(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListTlsRoute(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, TlsRouteToUnstructured(r))
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

func ApplyTlsRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTlsRoute(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTlsRoute(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyTlsRoute(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return TlsRouteToUnstructured(r), nil
}

func TlsRouteHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTlsRoute(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToTlsRoute(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyTlsRoute(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteTlsRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToTlsRoute(u)
	if err != nil {
		return err
	}
	return c.DeleteTlsRoute(ctx, r)
}

func TlsRouteID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToTlsRoute(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *TlsRoute) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networkservices",
		"TlsRoute",
		"alpha",
	}
}

func (r *TlsRoute) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TlsRoute) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TlsRoute) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *TlsRoute) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TlsRoute) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TlsRoute) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *TlsRoute) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetTlsRoute(ctx, config, resource)
}

func (r *TlsRoute) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyTlsRoute(ctx, config, resource, opts...)
}

func (r *TlsRoute) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return TlsRouteHasDiff(ctx, config, resource, opts...)
}

func (r *TlsRoute) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteTlsRoute(ctx, config, resource)
}

func (r *TlsRoute) ID(resource *unstructured.Resource) (string, error) {
	return TlsRouteID(resource)
}

func init() {
	unstructured.Register(&TlsRoute{})
}
