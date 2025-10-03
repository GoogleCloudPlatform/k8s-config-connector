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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type GrpcRoute struct{}

func GrpcRouteToUnstructured(r *dclService.GrpcRoute) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networkservices",
			Version: "beta",
			Type:    "GrpcRoute",
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
	var rHostnames []interface{}
	for _, rHostnamesVal := range r.Hostnames {
		rHostnames = append(rHostnames, rHostnamesVal)
	}
	u.Object["hostnames"] = rHostnames
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
		if rRulesVal.Action != nil && rRulesVal.Action != dclService.EmptyGrpcRouteRulesAction {
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
			if rRulesVal.Action.FaultInjectionPolicy != nil && rRulesVal.Action.FaultInjectionPolicy != dclService.EmptyGrpcRouteRulesActionFaultInjectionPolicy {
				rRulesValActionFaultInjectionPolicy := make(map[string]interface{})
				if rRulesVal.Action.FaultInjectionPolicy.Abort != nil && rRulesVal.Action.FaultInjectionPolicy.Abort != dclService.EmptyGrpcRouteRulesActionFaultInjectionPolicyAbort {
					rRulesValActionFaultInjectionPolicyAbort := make(map[string]interface{})
					if rRulesVal.Action.FaultInjectionPolicy.Abort.HttpStatus != nil {
						rRulesValActionFaultInjectionPolicyAbort["httpStatus"] = *rRulesVal.Action.FaultInjectionPolicy.Abort.HttpStatus
					}
					if rRulesVal.Action.FaultInjectionPolicy.Abort.Percentage != nil {
						rRulesValActionFaultInjectionPolicyAbort["percentage"] = *rRulesVal.Action.FaultInjectionPolicy.Abort.Percentage
					}
					rRulesValActionFaultInjectionPolicy["abort"] = rRulesValActionFaultInjectionPolicyAbort
				}
				if rRulesVal.Action.FaultInjectionPolicy.Delay != nil && rRulesVal.Action.FaultInjectionPolicy.Delay != dclService.EmptyGrpcRouteRulesActionFaultInjectionPolicyDelay {
					rRulesValActionFaultInjectionPolicyDelay := make(map[string]interface{})
					if rRulesVal.Action.FaultInjectionPolicy.Delay.FixedDelay != nil {
						rRulesValActionFaultInjectionPolicyDelay["fixedDelay"] = *rRulesVal.Action.FaultInjectionPolicy.Delay.FixedDelay
					}
					if rRulesVal.Action.FaultInjectionPolicy.Delay.Percentage != nil {
						rRulesValActionFaultInjectionPolicyDelay["percentage"] = *rRulesVal.Action.FaultInjectionPolicy.Delay.Percentage
					}
					rRulesValActionFaultInjectionPolicy["delay"] = rRulesValActionFaultInjectionPolicyDelay
				}
				rRulesValAction["faultInjectionPolicy"] = rRulesValActionFaultInjectionPolicy
			}
			if rRulesVal.Action.RetryPolicy != nil && rRulesVal.Action.RetryPolicy != dclService.EmptyGrpcRouteRulesActionRetryPolicy {
				rRulesValActionRetryPolicy := make(map[string]interface{})
				if rRulesVal.Action.RetryPolicy.NumRetries != nil {
					rRulesValActionRetryPolicy["numRetries"] = *rRulesVal.Action.RetryPolicy.NumRetries
				}
				var rRulesValActionRetryPolicyRetryConditions []interface{}
				for _, rRulesValActionRetryPolicyRetryConditionsVal := range rRulesVal.Action.RetryPolicy.RetryConditions {
					rRulesValActionRetryPolicyRetryConditions = append(rRulesValActionRetryPolicyRetryConditions, rRulesValActionRetryPolicyRetryConditionsVal)
				}
				rRulesValActionRetryPolicy["retryConditions"] = rRulesValActionRetryPolicyRetryConditions
				rRulesValAction["retryPolicy"] = rRulesValActionRetryPolicy
			}
			if rRulesVal.Action.Timeout != nil {
				rRulesValAction["timeout"] = *rRulesVal.Action.Timeout
			}
			rRulesObject["action"] = rRulesValAction
		}
		var rRulesValMatches []interface{}
		for _, rRulesValMatchesVal := range rRulesVal.Matches {
			rRulesValMatchesObject := make(map[string]interface{})
			var rRulesValMatchesValHeaders []interface{}
			for _, rRulesValMatchesValHeadersVal := range rRulesValMatchesVal.Headers {
				rRulesValMatchesValHeadersObject := make(map[string]interface{})
				if rRulesValMatchesValHeadersVal.Key != nil {
					rRulesValMatchesValHeadersObject["key"] = *rRulesValMatchesValHeadersVal.Key
				}
				if rRulesValMatchesValHeadersVal.Type != nil {
					rRulesValMatchesValHeadersObject["type"] = string(*rRulesValMatchesValHeadersVal.Type)
				}
				if rRulesValMatchesValHeadersVal.Value != nil {
					rRulesValMatchesValHeadersObject["value"] = *rRulesValMatchesValHeadersVal.Value
				}
				rRulesValMatchesValHeaders = append(rRulesValMatchesValHeaders, rRulesValMatchesValHeadersObject)
			}
			rRulesValMatchesObject["headers"] = rRulesValMatchesValHeaders
			if rRulesValMatchesVal.Method != nil && rRulesValMatchesVal.Method != dclService.EmptyGrpcRouteRulesMatchesMethod {
				rRulesValMatchesValMethod := make(map[string]interface{})
				if rRulesValMatchesVal.Method.CaseSensitive != nil {
					rRulesValMatchesValMethod["caseSensitive"] = *rRulesValMatchesVal.Method.CaseSensitive
				}
				if rRulesValMatchesVal.Method.GrpcMethod != nil {
					rRulesValMatchesValMethod["grpcMethod"] = *rRulesValMatchesVal.Method.GrpcMethod
				}
				if rRulesValMatchesVal.Method.GrpcService != nil {
					rRulesValMatchesValMethod["grpcService"] = *rRulesValMatchesVal.Method.GrpcService
				}
				if rRulesValMatchesVal.Method.Type != nil {
					rRulesValMatchesValMethod["type"] = string(*rRulesValMatchesVal.Method.Type)
				}
				rRulesValMatchesObject["method"] = rRulesValMatchesValMethod
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

func UnstructuredToGrpcRoute(u *unstructured.Resource) (*dclService.GrpcRoute, error) {
	r := &dclService.GrpcRoute{}
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
	if _, ok := u.Object["hostnames"]; ok {
		if s, ok := u.Object["hostnames"].([]interface{}); ok {
			for _, ss := range s {
				if strval, ok := ss.(string); ok {
					r.Hostnames = append(r.Hostnames, strval)
				}
			}
		} else {
			return nil, fmt.Errorf("r.Hostnames: expected []interface{}")
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
					var rRules dclService.GrpcRouteRules
					if _, ok := objval["action"]; ok {
						if rRulesAction, ok := objval["action"].(map[string]interface{}); ok {
							rRules.Action = &dclService.GrpcRouteRulesAction{}
							if _, ok := rRulesAction["destinations"]; ok {
								if s, ok := rRulesAction["destinations"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rRulesActionDestinations dclService.GrpcRouteRulesActionDestinations
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
							if _, ok := rRulesAction["faultInjectionPolicy"]; ok {
								if rRulesActionFaultInjectionPolicy, ok := rRulesAction["faultInjectionPolicy"].(map[string]interface{}); ok {
									rRules.Action.FaultInjectionPolicy = &dclService.GrpcRouteRulesActionFaultInjectionPolicy{}
									if _, ok := rRulesActionFaultInjectionPolicy["abort"]; ok {
										if rRulesActionFaultInjectionPolicyAbort, ok := rRulesActionFaultInjectionPolicy["abort"].(map[string]interface{}); ok {
											rRules.Action.FaultInjectionPolicy.Abort = &dclService.GrpcRouteRulesActionFaultInjectionPolicyAbort{}
											if _, ok := rRulesActionFaultInjectionPolicyAbort["httpStatus"]; ok {
												if i, ok := rRulesActionFaultInjectionPolicyAbort["httpStatus"].(int64); ok {
													rRules.Action.FaultInjectionPolicy.Abort.HttpStatus = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rRules.Action.FaultInjectionPolicy.Abort.HttpStatus: expected int64")
												}
											}
											if _, ok := rRulesActionFaultInjectionPolicyAbort["percentage"]; ok {
												if i, ok := rRulesActionFaultInjectionPolicyAbort["percentage"].(int64); ok {
													rRules.Action.FaultInjectionPolicy.Abort.Percentage = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rRules.Action.FaultInjectionPolicy.Abort.Percentage: expected int64")
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.FaultInjectionPolicy.Abort: expected map[string]interface{}")
										}
									}
									if _, ok := rRulesActionFaultInjectionPolicy["delay"]; ok {
										if rRulesActionFaultInjectionPolicyDelay, ok := rRulesActionFaultInjectionPolicy["delay"].(map[string]interface{}); ok {
											rRules.Action.FaultInjectionPolicy.Delay = &dclService.GrpcRouteRulesActionFaultInjectionPolicyDelay{}
											if _, ok := rRulesActionFaultInjectionPolicyDelay["fixedDelay"]; ok {
												if s, ok := rRulesActionFaultInjectionPolicyDelay["fixedDelay"].(string); ok {
													rRules.Action.FaultInjectionPolicy.Delay.FixedDelay = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRules.Action.FaultInjectionPolicy.Delay.FixedDelay: expected string")
												}
											}
											if _, ok := rRulesActionFaultInjectionPolicyDelay["percentage"]; ok {
												if i, ok := rRulesActionFaultInjectionPolicyDelay["percentage"].(int64); ok {
													rRules.Action.FaultInjectionPolicy.Delay.Percentage = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rRules.Action.FaultInjectionPolicy.Delay.Percentage: expected int64")
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.FaultInjectionPolicy.Delay: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rRules.Action.FaultInjectionPolicy: expected map[string]interface{}")
								}
							}
							if _, ok := rRulesAction["retryPolicy"]; ok {
								if rRulesActionRetryPolicy, ok := rRulesAction["retryPolicy"].(map[string]interface{}); ok {
									rRules.Action.RetryPolicy = &dclService.GrpcRouteRulesActionRetryPolicy{}
									if _, ok := rRulesActionRetryPolicy["numRetries"]; ok {
										if i, ok := rRulesActionRetryPolicy["numRetries"].(int64); ok {
											rRules.Action.RetryPolicy.NumRetries = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rRules.Action.RetryPolicy.NumRetries: expected int64")
										}
									}
									if _, ok := rRulesActionRetryPolicy["retryConditions"]; ok {
										if s, ok := rRulesActionRetryPolicy["retryConditions"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRules.Action.RetryPolicy.RetryConditions = append(rRules.Action.RetryPolicy.RetryConditions, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.RetryPolicy.RetryConditions: expected []interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rRules.Action.RetryPolicy: expected map[string]interface{}")
								}
							}
							if _, ok := rRulesAction["timeout"]; ok {
								if s, ok := rRulesAction["timeout"].(string); ok {
									rRules.Action.Timeout = dcl.String(s)
								} else {
									return nil, fmt.Errorf("rRules.Action.Timeout: expected string")
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
									var rRulesMatches dclService.GrpcRouteRulesMatches
									if _, ok := objval["headers"]; ok {
										if s, ok := objval["headers"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rRulesMatchesHeaders dclService.GrpcRouteRulesMatchesHeaders
													if _, ok := objval["key"]; ok {
														if s, ok := objval["key"].(string); ok {
															rRulesMatchesHeaders.Key = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.Key: expected string")
														}
													}
													if _, ok := objval["type"]; ok {
														if s, ok := objval["type"].(string); ok {
															rRulesMatchesHeaders.Type = dclService.GrpcRouteRulesMatchesHeadersTypeEnumRef(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.Type: expected string")
														}
													}
													if _, ok := objval["value"]; ok {
														if s, ok := objval["value"].(string); ok {
															rRulesMatchesHeaders.Value = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.Value: expected string")
														}
													}
													rRulesMatches.Headers = append(rRulesMatches.Headers, rRulesMatchesHeaders)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesMatches.Headers: expected []interface{}")
										}
									}
									if _, ok := objval["method"]; ok {
										if rRulesMatchesMethod, ok := objval["method"].(map[string]interface{}); ok {
											rRulesMatches.Method = &dclService.GrpcRouteRulesMatchesMethod{}
											if _, ok := rRulesMatchesMethod["caseSensitive"]; ok {
												if b, ok := rRulesMatchesMethod["caseSensitive"].(bool); ok {
													rRulesMatches.Method.CaseSensitive = dcl.Bool(b)
												} else {
													return nil, fmt.Errorf("rRulesMatches.Method.CaseSensitive: expected bool")
												}
											}
											if _, ok := rRulesMatchesMethod["grpcMethod"]; ok {
												if s, ok := rRulesMatchesMethod["grpcMethod"].(string); ok {
													rRulesMatches.Method.GrpcMethod = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRulesMatches.Method.GrpcMethod: expected string")
												}
											}
											if _, ok := rRulesMatchesMethod["grpcService"]; ok {
												if s, ok := rRulesMatchesMethod["grpcService"].(string); ok {
													rRulesMatches.Method.GrpcService = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRulesMatches.Method.GrpcService: expected string")
												}
											}
											if _, ok := rRulesMatchesMethod["type"]; ok {
												if s, ok := rRulesMatchesMethod["type"].(string); ok {
													rRulesMatches.Method.Type = dclService.GrpcRouteRulesMatchesMethodTypeEnumRef(s)
												} else {
													return nil, fmt.Errorf("rRulesMatches.Method.Type: expected string")
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesMatches.Method: expected map[string]interface{}")
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

func GetGrpcRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGrpcRoute(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetGrpcRoute(ctx, r)
	if err != nil {
		return nil, err
	}
	return GrpcRouteToUnstructured(r), nil
}

func ListGrpcRoute(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListGrpcRoute(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, GrpcRouteToUnstructured(r))
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

func ApplyGrpcRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGrpcRoute(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGrpcRoute(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyGrpcRoute(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return GrpcRouteToUnstructured(r), nil
}

func GrpcRouteHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGrpcRoute(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToGrpcRoute(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyGrpcRoute(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteGrpcRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToGrpcRoute(u)
	if err != nil {
		return err
	}
	return c.DeleteGrpcRoute(ctx, r)
}

func GrpcRouteID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToGrpcRoute(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *GrpcRoute) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networkservices",
		"GrpcRoute",
		"beta",
	}
}

func (r *GrpcRoute) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GrpcRoute) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GrpcRoute) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *GrpcRoute) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GrpcRoute) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GrpcRoute) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *GrpcRoute) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetGrpcRoute(ctx, config, resource)
}

func (r *GrpcRoute) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyGrpcRoute(ctx, config, resource, opts...)
}

func (r *GrpcRoute) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return GrpcRouteHasDiff(ctx, config, resource, opts...)
}

func (r *GrpcRoute) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteGrpcRoute(ctx, config, resource)
}

func (r *GrpcRoute) ID(resource *unstructured.Resource) (string, error) {
	return GrpcRouteID(resource)
}

func init() {
	unstructured.Register(&GrpcRoute{})
}
