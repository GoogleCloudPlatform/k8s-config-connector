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
	dclService "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
)

type HttpRoute struct{}

func HttpRouteToUnstructured(r *dclService.HttpRoute) *unstructured.Resource {
	u := &unstructured.Resource{
		STV: unstructured.ServiceTypeVersion{
			Service: "networkservices",
			Version: "ga",
			Type:    "HttpRoute",
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
		if rRulesVal.Action != nil && rRulesVal.Action != dclService.EmptyHttpRouteRulesAction {
			rRulesValAction := make(map[string]interface{})
			if rRulesVal.Action.CorsPolicy != nil && rRulesVal.Action.CorsPolicy != dclService.EmptyHttpRouteRulesActionCorsPolicy {
				rRulesValActionCorsPolicy := make(map[string]interface{})
				if rRulesVal.Action.CorsPolicy.AllowCredentials != nil {
					rRulesValActionCorsPolicy["allowCredentials"] = *rRulesVal.Action.CorsPolicy.AllowCredentials
				}
				var rRulesValActionCorsPolicyAllowHeaders []interface{}
				for _, rRulesValActionCorsPolicyAllowHeadersVal := range rRulesVal.Action.CorsPolicy.AllowHeaders {
					rRulesValActionCorsPolicyAllowHeaders = append(rRulesValActionCorsPolicyAllowHeaders, rRulesValActionCorsPolicyAllowHeadersVal)
				}
				rRulesValActionCorsPolicy["allowHeaders"] = rRulesValActionCorsPolicyAllowHeaders
				var rRulesValActionCorsPolicyAllowMethods []interface{}
				for _, rRulesValActionCorsPolicyAllowMethodsVal := range rRulesVal.Action.CorsPolicy.AllowMethods {
					rRulesValActionCorsPolicyAllowMethods = append(rRulesValActionCorsPolicyAllowMethods, rRulesValActionCorsPolicyAllowMethodsVal)
				}
				rRulesValActionCorsPolicy["allowMethods"] = rRulesValActionCorsPolicyAllowMethods
				var rRulesValActionCorsPolicyAllowOriginRegexes []interface{}
				for _, rRulesValActionCorsPolicyAllowOriginRegexesVal := range rRulesVal.Action.CorsPolicy.AllowOriginRegexes {
					rRulesValActionCorsPolicyAllowOriginRegexes = append(rRulesValActionCorsPolicyAllowOriginRegexes, rRulesValActionCorsPolicyAllowOriginRegexesVal)
				}
				rRulesValActionCorsPolicy["allowOriginRegexes"] = rRulesValActionCorsPolicyAllowOriginRegexes
				var rRulesValActionCorsPolicyAllowOrigins []interface{}
				for _, rRulesValActionCorsPolicyAllowOriginsVal := range rRulesVal.Action.CorsPolicy.AllowOrigins {
					rRulesValActionCorsPolicyAllowOrigins = append(rRulesValActionCorsPolicyAllowOrigins, rRulesValActionCorsPolicyAllowOriginsVal)
				}
				rRulesValActionCorsPolicy["allowOrigins"] = rRulesValActionCorsPolicyAllowOrigins
				if rRulesVal.Action.CorsPolicy.Disabled != nil {
					rRulesValActionCorsPolicy["disabled"] = *rRulesVal.Action.CorsPolicy.Disabled
				}
				var rRulesValActionCorsPolicyExposeHeaders []interface{}
				for _, rRulesValActionCorsPolicyExposeHeadersVal := range rRulesVal.Action.CorsPolicy.ExposeHeaders {
					rRulesValActionCorsPolicyExposeHeaders = append(rRulesValActionCorsPolicyExposeHeaders, rRulesValActionCorsPolicyExposeHeadersVal)
				}
				rRulesValActionCorsPolicy["exposeHeaders"] = rRulesValActionCorsPolicyExposeHeaders
				if rRulesVal.Action.CorsPolicy.MaxAge != nil {
					rRulesValActionCorsPolicy["maxAge"] = *rRulesVal.Action.CorsPolicy.MaxAge
				}
				rRulesValAction["corsPolicy"] = rRulesValActionCorsPolicy
			}
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
			if rRulesVal.Action.FaultInjectionPolicy != nil && rRulesVal.Action.FaultInjectionPolicy != dclService.EmptyHttpRouteRulesActionFaultInjectionPolicy {
				rRulesValActionFaultInjectionPolicy := make(map[string]interface{})
				if rRulesVal.Action.FaultInjectionPolicy.Abort != nil && rRulesVal.Action.FaultInjectionPolicy.Abort != dclService.EmptyHttpRouteRulesActionFaultInjectionPolicyAbort {
					rRulesValActionFaultInjectionPolicyAbort := make(map[string]interface{})
					if rRulesVal.Action.FaultInjectionPolicy.Abort.HttpStatus != nil {
						rRulesValActionFaultInjectionPolicyAbort["httpStatus"] = *rRulesVal.Action.FaultInjectionPolicy.Abort.HttpStatus
					}
					if rRulesVal.Action.FaultInjectionPolicy.Abort.Percentage != nil {
						rRulesValActionFaultInjectionPolicyAbort["percentage"] = *rRulesVal.Action.FaultInjectionPolicy.Abort.Percentage
					}
					rRulesValActionFaultInjectionPolicy["abort"] = rRulesValActionFaultInjectionPolicyAbort
				}
				if rRulesVal.Action.FaultInjectionPolicy.Delay != nil && rRulesVal.Action.FaultInjectionPolicy.Delay != dclService.EmptyHttpRouteRulesActionFaultInjectionPolicyDelay {
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
			if rRulesVal.Action.Redirect != nil && rRulesVal.Action.Redirect != dclService.EmptyHttpRouteRulesActionRedirect {
				rRulesValActionRedirect := make(map[string]interface{})
				if rRulesVal.Action.Redirect.HostRedirect != nil {
					rRulesValActionRedirect["hostRedirect"] = *rRulesVal.Action.Redirect.HostRedirect
				}
				if rRulesVal.Action.Redirect.HttpsRedirect != nil {
					rRulesValActionRedirect["httpsRedirect"] = *rRulesVal.Action.Redirect.HttpsRedirect
				}
				if rRulesVal.Action.Redirect.PathRedirect != nil {
					rRulesValActionRedirect["pathRedirect"] = *rRulesVal.Action.Redirect.PathRedirect
				}
				if rRulesVal.Action.Redirect.PortRedirect != nil {
					rRulesValActionRedirect["portRedirect"] = *rRulesVal.Action.Redirect.PortRedirect
				}
				if rRulesVal.Action.Redirect.PrefixRewrite != nil {
					rRulesValActionRedirect["prefixRewrite"] = *rRulesVal.Action.Redirect.PrefixRewrite
				}
				if rRulesVal.Action.Redirect.ResponseCode != nil {
					rRulesValActionRedirect["responseCode"] = string(*rRulesVal.Action.Redirect.ResponseCode)
				}
				if rRulesVal.Action.Redirect.StripQuery != nil {
					rRulesValActionRedirect["stripQuery"] = *rRulesVal.Action.Redirect.StripQuery
				}
				rRulesValAction["redirect"] = rRulesValActionRedirect
			}
			if rRulesVal.Action.RequestHeaderModifier != nil && rRulesVal.Action.RequestHeaderModifier != dclService.EmptyHttpRouteRulesActionRequestHeaderModifier {
				rRulesValActionRequestHeaderModifier := make(map[string]interface{})
				if rRulesVal.Action.RequestHeaderModifier.Add != nil {
					rRulesValActionRequestHeaderModifierAdd := make(map[string]interface{})
					for k, v := range rRulesVal.Action.RequestHeaderModifier.Add {
						rRulesValActionRequestHeaderModifierAdd[k] = v
					}
					rRulesValActionRequestHeaderModifier["add"] = rRulesValActionRequestHeaderModifierAdd
				}
				var rRulesValActionRequestHeaderModifierRemove []interface{}
				for _, rRulesValActionRequestHeaderModifierRemoveVal := range rRulesVal.Action.RequestHeaderModifier.Remove {
					rRulesValActionRequestHeaderModifierRemove = append(rRulesValActionRequestHeaderModifierRemove, rRulesValActionRequestHeaderModifierRemoveVal)
				}
				rRulesValActionRequestHeaderModifier["remove"] = rRulesValActionRequestHeaderModifierRemove
				if rRulesVal.Action.RequestHeaderModifier.Set != nil {
					rRulesValActionRequestHeaderModifierSet := make(map[string]interface{})
					for k, v := range rRulesVal.Action.RequestHeaderModifier.Set {
						rRulesValActionRequestHeaderModifierSet[k] = v
					}
					rRulesValActionRequestHeaderModifier["set"] = rRulesValActionRequestHeaderModifierSet
				}
				rRulesValAction["requestHeaderModifier"] = rRulesValActionRequestHeaderModifier
			}
			if rRulesVal.Action.RequestMirrorPolicy != nil && rRulesVal.Action.RequestMirrorPolicy != dclService.EmptyHttpRouteRulesActionRequestMirrorPolicy {
				rRulesValActionRequestMirrorPolicy := make(map[string]interface{})
				if rRulesVal.Action.RequestMirrorPolicy.Destination != nil && rRulesVal.Action.RequestMirrorPolicy.Destination != dclService.EmptyHttpRouteRulesActionRequestMirrorPolicyDestination {
					rRulesValActionRequestMirrorPolicyDestination := make(map[string]interface{})
					if rRulesVal.Action.RequestMirrorPolicy.Destination.ServiceName != nil {
						rRulesValActionRequestMirrorPolicyDestination["serviceName"] = *rRulesVal.Action.RequestMirrorPolicy.Destination.ServiceName
					}
					if rRulesVal.Action.RequestMirrorPolicy.Destination.Weight != nil {
						rRulesValActionRequestMirrorPolicyDestination["weight"] = *rRulesVal.Action.RequestMirrorPolicy.Destination.Weight
					}
					rRulesValActionRequestMirrorPolicy["destination"] = rRulesValActionRequestMirrorPolicyDestination
				}
				rRulesValAction["requestMirrorPolicy"] = rRulesValActionRequestMirrorPolicy
			}
			if rRulesVal.Action.ResponseHeaderModifier != nil && rRulesVal.Action.ResponseHeaderModifier != dclService.EmptyHttpRouteRulesActionResponseHeaderModifier {
				rRulesValActionResponseHeaderModifier := make(map[string]interface{})
				if rRulesVal.Action.ResponseHeaderModifier.Add != nil {
					rRulesValActionResponseHeaderModifierAdd := make(map[string]interface{})
					for k, v := range rRulesVal.Action.ResponseHeaderModifier.Add {
						rRulesValActionResponseHeaderModifierAdd[k] = v
					}
					rRulesValActionResponseHeaderModifier["add"] = rRulesValActionResponseHeaderModifierAdd
				}
				var rRulesValActionResponseHeaderModifierRemove []interface{}
				for _, rRulesValActionResponseHeaderModifierRemoveVal := range rRulesVal.Action.ResponseHeaderModifier.Remove {
					rRulesValActionResponseHeaderModifierRemove = append(rRulesValActionResponseHeaderModifierRemove, rRulesValActionResponseHeaderModifierRemoveVal)
				}
				rRulesValActionResponseHeaderModifier["remove"] = rRulesValActionResponseHeaderModifierRemove
				if rRulesVal.Action.ResponseHeaderModifier.Set != nil {
					rRulesValActionResponseHeaderModifierSet := make(map[string]interface{})
					for k, v := range rRulesVal.Action.ResponseHeaderModifier.Set {
						rRulesValActionResponseHeaderModifierSet[k] = v
					}
					rRulesValActionResponseHeaderModifier["set"] = rRulesValActionResponseHeaderModifierSet
				}
				rRulesValAction["responseHeaderModifier"] = rRulesValActionResponseHeaderModifier
			}
			if rRulesVal.Action.RetryPolicy != nil && rRulesVal.Action.RetryPolicy != dclService.EmptyHttpRouteRulesActionRetryPolicy {
				rRulesValActionRetryPolicy := make(map[string]interface{})
				if rRulesVal.Action.RetryPolicy.NumRetries != nil {
					rRulesValActionRetryPolicy["numRetries"] = *rRulesVal.Action.RetryPolicy.NumRetries
				}
				if rRulesVal.Action.RetryPolicy.PerTryTimeout != nil {
					rRulesValActionRetryPolicy["perTryTimeout"] = *rRulesVal.Action.RetryPolicy.PerTryTimeout
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
			if rRulesVal.Action.UrlRewrite != nil && rRulesVal.Action.UrlRewrite != dclService.EmptyHttpRouteRulesActionUrlRewrite {
				rRulesValActionUrlRewrite := make(map[string]interface{})
				if rRulesVal.Action.UrlRewrite.HostRewrite != nil {
					rRulesValActionUrlRewrite["hostRewrite"] = *rRulesVal.Action.UrlRewrite.HostRewrite
				}
				if rRulesVal.Action.UrlRewrite.PathPrefixRewrite != nil {
					rRulesValActionUrlRewrite["pathPrefixRewrite"] = *rRulesVal.Action.UrlRewrite.PathPrefixRewrite
				}
				rRulesValAction["urlRewrite"] = rRulesValActionUrlRewrite
			}
			rRulesObject["action"] = rRulesValAction
		}
		var rRulesValMatches []interface{}
		for _, rRulesValMatchesVal := range rRulesVal.Matches {
			rRulesValMatchesObject := make(map[string]interface{})
			if rRulesValMatchesVal.FullPathMatch != nil {
				rRulesValMatchesObject["fullPathMatch"] = *rRulesValMatchesVal.FullPathMatch
			}
			var rRulesValMatchesValHeaders []interface{}
			for _, rRulesValMatchesValHeadersVal := range rRulesValMatchesVal.Headers {
				rRulesValMatchesValHeadersObject := make(map[string]interface{})
				if rRulesValMatchesValHeadersVal.ExactMatch != nil {
					rRulesValMatchesValHeadersObject["exactMatch"] = *rRulesValMatchesValHeadersVal.ExactMatch
				}
				if rRulesValMatchesValHeadersVal.Header != nil {
					rRulesValMatchesValHeadersObject["header"] = *rRulesValMatchesValHeadersVal.Header
				}
				if rRulesValMatchesValHeadersVal.InvertMatch != nil {
					rRulesValMatchesValHeadersObject["invertMatch"] = *rRulesValMatchesValHeadersVal.InvertMatch
				}
				if rRulesValMatchesValHeadersVal.PrefixMatch != nil {
					rRulesValMatchesValHeadersObject["prefixMatch"] = *rRulesValMatchesValHeadersVal.PrefixMatch
				}
				if rRulesValMatchesValHeadersVal.PresentMatch != nil {
					rRulesValMatchesValHeadersObject["presentMatch"] = *rRulesValMatchesValHeadersVal.PresentMatch
				}
				if rRulesValMatchesValHeadersVal.RangeMatch != nil && rRulesValMatchesValHeadersVal.RangeMatch != dclService.EmptyHttpRouteRulesMatchesHeadersRangeMatch {
					rRulesValMatchesValHeadersValRangeMatch := make(map[string]interface{})
					if rRulesValMatchesValHeadersVal.RangeMatch.End != nil {
						rRulesValMatchesValHeadersValRangeMatch["end"] = *rRulesValMatchesValHeadersVal.RangeMatch.End
					}
					if rRulesValMatchesValHeadersVal.RangeMatch.Start != nil {
						rRulesValMatchesValHeadersValRangeMatch["start"] = *rRulesValMatchesValHeadersVal.RangeMatch.Start
					}
					rRulesValMatchesValHeadersObject["rangeMatch"] = rRulesValMatchesValHeadersValRangeMatch
				}
				if rRulesValMatchesValHeadersVal.RegexMatch != nil {
					rRulesValMatchesValHeadersObject["regexMatch"] = *rRulesValMatchesValHeadersVal.RegexMatch
				}
				if rRulesValMatchesValHeadersVal.SuffixMatch != nil {
					rRulesValMatchesValHeadersObject["suffixMatch"] = *rRulesValMatchesValHeadersVal.SuffixMatch
				}
				rRulesValMatchesValHeaders = append(rRulesValMatchesValHeaders, rRulesValMatchesValHeadersObject)
			}
			rRulesValMatchesObject["headers"] = rRulesValMatchesValHeaders
			if rRulesValMatchesVal.IgnoreCase != nil {
				rRulesValMatchesObject["ignoreCase"] = *rRulesValMatchesVal.IgnoreCase
			}
			if rRulesValMatchesVal.PrefixMatch != nil {
				rRulesValMatchesObject["prefixMatch"] = *rRulesValMatchesVal.PrefixMatch
			}
			var rRulesValMatchesValQueryParameters []interface{}
			for _, rRulesValMatchesValQueryParametersVal := range rRulesValMatchesVal.QueryParameters {
				rRulesValMatchesValQueryParametersObject := make(map[string]interface{})
				if rRulesValMatchesValQueryParametersVal.ExactMatch != nil {
					rRulesValMatchesValQueryParametersObject["exactMatch"] = *rRulesValMatchesValQueryParametersVal.ExactMatch
				}
				if rRulesValMatchesValQueryParametersVal.PresentMatch != nil {
					rRulesValMatchesValQueryParametersObject["presentMatch"] = *rRulesValMatchesValQueryParametersVal.PresentMatch
				}
				if rRulesValMatchesValQueryParametersVal.QueryParameter != nil {
					rRulesValMatchesValQueryParametersObject["queryParameter"] = *rRulesValMatchesValQueryParametersVal.QueryParameter
				}
				if rRulesValMatchesValQueryParametersVal.RegexMatch != nil {
					rRulesValMatchesValQueryParametersObject["regexMatch"] = *rRulesValMatchesValQueryParametersVal.RegexMatch
				}
				rRulesValMatchesValQueryParameters = append(rRulesValMatchesValQueryParameters, rRulesValMatchesValQueryParametersObject)
			}
			rRulesValMatchesObject["queryParameters"] = rRulesValMatchesValQueryParameters
			if rRulesValMatchesVal.RegexMatch != nil {
				rRulesValMatchesObject["regexMatch"] = *rRulesValMatchesVal.RegexMatch
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

func UnstructuredToHttpRoute(u *unstructured.Resource) (*dclService.HttpRoute, error) {
	r := &dclService.HttpRoute{}
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
					var rRules dclService.HttpRouteRules
					if _, ok := objval["action"]; ok {
						if rRulesAction, ok := objval["action"].(map[string]interface{}); ok {
							rRules.Action = &dclService.HttpRouteRulesAction{}
							if _, ok := rRulesAction["corsPolicy"]; ok {
								if rRulesActionCorsPolicy, ok := rRulesAction["corsPolicy"].(map[string]interface{}); ok {
									rRules.Action.CorsPolicy = &dclService.HttpRouteRulesActionCorsPolicy{}
									if _, ok := rRulesActionCorsPolicy["allowCredentials"]; ok {
										if b, ok := rRulesActionCorsPolicy["allowCredentials"].(bool); ok {
											rRules.Action.CorsPolicy.AllowCredentials = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rRules.Action.CorsPolicy.AllowCredentials: expected bool")
										}
									}
									if _, ok := rRulesActionCorsPolicy["allowHeaders"]; ok {
										if s, ok := rRulesActionCorsPolicy["allowHeaders"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRules.Action.CorsPolicy.AllowHeaders = append(rRules.Action.CorsPolicy.AllowHeaders, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.CorsPolicy.AllowHeaders: expected []interface{}")
										}
									}
									if _, ok := rRulesActionCorsPolicy["allowMethods"]; ok {
										if s, ok := rRulesActionCorsPolicy["allowMethods"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRules.Action.CorsPolicy.AllowMethods = append(rRules.Action.CorsPolicy.AllowMethods, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.CorsPolicy.AllowMethods: expected []interface{}")
										}
									}
									if _, ok := rRulesActionCorsPolicy["allowOriginRegexes"]; ok {
										if s, ok := rRulesActionCorsPolicy["allowOriginRegexes"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRules.Action.CorsPolicy.AllowOriginRegexes = append(rRules.Action.CorsPolicy.AllowOriginRegexes, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.CorsPolicy.AllowOriginRegexes: expected []interface{}")
										}
									}
									if _, ok := rRulesActionCorsPolicy["allowOrigins"]; ok {
										if s, ok := rRulesActionCorsPolicy["allowOrigins"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRules.Action.CorsPolicy.AllowOrigins = append(rRules.Action.CorsPolicy.AllowOrigins, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.CorsPolicy.AllowOrigins: expected []interface{}")
										}
									}
									if _, ok := rRulesActionCorsPolicy["disabled"]; ok {
										if b, ok := rRulesActionCorsPolicy["disabled"].(bool); ok {
											rRules.Action.CorsPolicy.Disabled = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rRules.Action.CorsPolicy.Disabled: expected bool")
										}
									}
									if _, ok := rRulesActionCorsPolicy["exposeHeaders"]; ok {
										if s, ok := rRulesActionCorsPolicy["exposeHeaders"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRules.Action.CorsPolicy.ExposeHeaders = append(rRules.Action.CorsPolicy.ExposeHeaders, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.CorsPolicy.ExposeHeaders: expected []interface{}")
										}
									}
									if _, ok := rRulesActionCorsPolicy["maxAge"]; ok {
										if s, ok := rRulesActionCorsPolicy["maxAge"].(string); ok {
											rRules.Action.CorsPolicy.MaxAge = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRules.Action.CorsPolicy.MaxAge: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rRules.Action.CorsPolicy: expected map[string]interface{}")
								}
							}
							if _, ok := rRulesAction["destinations"]; ok {
								if s, ok := rRulesAction["destinations"].([]interface{}); ok {
									for _, o := range s {
										if objval, ok := o.(map[string]interface{}); ok {
											var rRulesActionDestinations dclService.HttpRouteRulesActionDestinations
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
									rRules.Action.FaultInjectionPolicy = &dclService.HttpRouteRulesActionFaultInjectionPolicy{}
									if _, ok := rRulesActionFaultInjectionPolicy["abort"]; ok {
										if rRulesActionFaultInjectionPolicyAbort, ok := rRulesActionFaultInjectionPolicy["abort"].(map[string]interface{}); ok {
											rRules.Action.FaultInjectionPolicy.Abort = &dclService.HttpRouteRulesActionFaultInjectionPolicyAbort{}
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
											rRules.Action.FaultInjectionPolicy.Delay = &dclService.HttpRouteRulesActionFaultInjectionPolicyDelay{}
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
							if _, ok := rRulesAction["redirect"]; ok {
								if rRulesActionRedirect, ok := rRulesAction["redirect"].(map[string]interface{}); ok {
									rRules.Action.Redirect = &dclService.HttpRouteRulesActionRedirect{}
									if _, ok := rRulesActionRedirect["hostRedirect"]; ok {
										if s, ok := rRulesActionRedirect["hostRedirect"].(string); ok {
											rRules.Action.Redirect.HostRedirect = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRules.Action.Redirect.HostRedirect: expected string")
										}
									}
									if _, ok := rRulesActionRedirect["httpsRedirect"]; ok {
										if b, ok := rRulesActionRedirect["httpsRedirect"].(bool); ok {
											rRules.Action.Redirect.HttpsRedirect = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rRules.Action.Redirect.HttpsRedirect: expected bool")
										}
									}
									if _, ok := rRulesActionRedirect["pathRedirect"]; ok {
										if s, ok := rRulesActionRedirect["pathRedirect"].(string); ok {
											rRules.Action.Redirect.PathRedirect = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRules.Action.Redirect.PathRedirect: expected string")
										}
									}
									if _, ok := rRulesActionRedirect["portRedirect"]; ok {
										if i, ok := rRulesActionRedirect["portRedirect"].(int64); ok {
											rRules.Action.Redirect.PortRedirect = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rRules.Action.Redirect.PortRedirect: expected int64")
										}
									}
									if _, ok := rRulesActionRedirect["prefixRewrite"]; ok {
										if s, ok := rRulesActionRedirect["prefixRewrite"].(string); ok {
											rRules.Action.Redirect.PrefixRewrite = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRules.Action.Redirect.PrefixRewrite: expected string")
										}
									}
									if _, ok := rRulesActionRedirect["responseCode"]; ok {
										if s, ok := rRulesActionRedirect["responseCode"].(string); ok {
											rRules.Action.Redirect.ResponseCode = dclService.HttpRouteRulesActionRedirectResponseCodeEnumRef(s)
										} else {
											return nil, fmt.Errorf("rRules.Action.Redirect.ResponseCode: expected string")
										}
									}
									if _, ok := rRulesActionRedirect["stripQuery"]; ok {
										if b, ok := rRulesActionRedirect["stripQuery"].(bool); ok {
											rRules.Action.Redirect.StripQuery = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rRules.Action.Redirect.StripQuery: expected bool")
										}
									}
								} else {
									return nil, fmt.Errorf("rRules.Action.Redirect: expected map[string]interface{}")
								}
							}
							if _, ok := rRulesAction["requestHeaderModifier"]; ok {
								if rRulesActionRequestHeaderModifier, ok := rRulesAction["requestHeaderModifier"].(map[string]interface{}); ok {
									rRules.Action.RequestHeaderModifier = &dclService.HttpRouteRulesActionRequestHeaderModifier{}
									if _, ok := rRulesActionRequestHeaderModifier["add"]; ok {
										if rRulesActionRequestHeaderModifierAdd, ok := rRulesActionRequestHeaderModifier["add"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rRulesActionRequestHeaderModifierAdd {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rRules.Action.RequestHeaderModifier.Add = m
										} else {
											return nil, fmt.Errorf("rRules.Action.RequestHeaderModifier.Add: expected map[string]interface{}")
										}
									}
									if _, ok := rRulesActionRequestHeaderModifier["remove"]; ok {
										if s, ok := rRulesActionRequestHeaderModifier["remove"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRules.Action.RequestHeaderModifier.Remove = append(rRules.Action.RequestHeaderModifier.Remove, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.RequestHeaderModifier.Remove: expected []interface{}")
										}
									}
									if _, ok := rRulesActionRequestHeaderModifier["set"]; ok {
										if rRulesActionRequestHeaderModifierSet, ok := rRulesActionRequestHeaderModifier["set"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rRulesActionRequestHeaderModifierSet {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rRules.Action.RequestHeaderModifier.Set = m
										} else {
											return nil, fmt.Errorf("rRules.Action.RequestHeaderModifier.Set: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rRules.Action.RequestHeaderModifier: expected map[string]interface{}")
								}
							}
							if _, ok := rRulesAction["requestMirrorPolicy"]; ok {
								if rRulesActionRequestMirrorPolicy, ok := rRulesAction["requestMirrorPolicy"].(map[string]interface{}); ok {
									rRules.Action.RequestMirrorPolicy = &dclService.HttpRouteRulesActionRequestMirrorPolicy{}
									if _, ok := rRulesActionRequestMirrorPolicy["destination"]; ok {
										if rRulesActionRequestMirrorPolicyDestination, ok := rRulesActionRequestMirrorPolicy["destination"].(map[string]interface{}); ok {
											rRules.Action.RequestMirrorPolicy.Destination = &dclService.HttpRouteRulesActionRequestMirrorPolicyDestination{}
											if _, ok := rRulesActionRequestMirrorPolicyDestination["serviceName"]; ok {
												if s, ok := rRulesActionRequestMirrorPolicyDestination["serviceName"].(string); ok {
													rRules.Action.RequestMirrorPolicy.Destination.ServiceName = dcl.String(s)
												} else {
													return nil, fmt.Errorf("rRules.Action.RequestMirrorPolicy.Destination.ServiceName: expected string")
												}
											}
											if _, ok := rRulesActionRequestMirrorPolicyDestination["weight"]; ok {
												if i, ok := rRulesActionRequestMirrorPolicyDestination["weight"].(int64); ok {
													rRules.Action.RequestMirrorPolicy.Destination.Weight = dcl.Int64(i)
												} else {
													return nil, fmt.Errorf("rRules.Action.RequestMirrorPolicy.Destination.Weight: expected int64")
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.RequestMirrorPolicy.Destination: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rRules.Action.RequestMirrorPolicy: expected map[string]interface{}")
								}
							}
							if _, ok := rRulesAction["responseHeaderModifier"]; ok {
								if rRulesActionResponseHeaderModifier, ok := rRulesAction["responseHeaderModifier"].(map[string]interface{}); ok {
									rRules.Action.ResponseHeaderModifier = &dclService.HttpRouteRulesActionResponseHeaderModifier{}
									if _, ok := rRulesActionResponseHeaderModifier["add"]; ok {
										if rRulesActionResponseHeaderModifierAdd, ok := rRulesActionResponseHeaderModifier["add"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rRulesActionResponseHeaderModifierAdd {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rRules.Action.ResponseHeaderModifier.Add = m
										} else {
											return nil, fmt.Errorf("rRules.Action.ResponseHeaderModifier.Add: expected map[string]interface{}")
										}
									}
									if _, ok := rRulesActionResponseHeaderModifier["remove"]; ok {
										if s, ok := rRulesActionResponseHeaderModifier["remove"].([]interface{}); ok {
											for _, ss := range s {
												if strval, ok := ss.(string); ok {
													rRules.Action.ResponseHeaderModifier.Remove = append(rRules.Action.ResponseHeaderModifier.Remove, strval)
												}
											}
										} else {
											return nil, fmt.Errorf("rRules.Action.ResponseHeaderModifier.Remove: expected []interface{}")
										}
									}
									if _, ok := rRulesActionResponseHeaderModifier["set"]; ok {
										if rRulesActionResponseHeaderModifierSet, ok := rRulesActionResponseHeaderModifier["set"].(map[string]interface{}); ok {
											m := make(map[string]string)
											for k, v := range rRulesActionResponseHeaderModifierSet {
												if s, ok := v.(string); ok {
													m[k] = s
												}
											}
											rRules.Action.ResponseHeaderModifier.Set = m
										} else {
											return nil, fmt.Errorf("rRules.Action.ResponseHeaderModifier.Set: expected map[string]interface{}")
										}
									}
								} else {
									return nil, fmt.Errorf("rRules.Action.ResponseHeaderModifier: expected map[string]interface{}")
								}
							}
							if _, ok := rRulesAction["retryPolicy"]; ok {
								if rRulesActionRetryPolicy, ok := rRulesAction["retryPolicy"].(map[string]interface{}); ok {
									rRules.Action.RetryPolicy = &dclService.HttpRouteRulesActionRetryPolicy{}
									if _, ok := rRulesActionRetryPolicy["numRetries"]; ok {
										if i, ok := rRulesActionRetryPolicy["numRetries"].(int64); ok {
											rRules.Action.RetryPolicy.NumRetries = dcl.Int64(i)
										} else {
											return nil, fmt.Errorf("rRules.Action.RetryPolicy.NumRetries: expected int64")
										}
									}
									if _, ok := rRulesActionRetryPolicy["perTryTimeout"]; ok {
										if s, ok := rRulesActionRetryPolicy["perTryTimeout"].(string); ok {
											rRules.Action.RetryPolicy.PerTryTimeout = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRules.Action.RetryPolicy.PerTryTimeout: expected string")
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
							if _, ok := rRulesAction["urlRewrite"]; ok {
								if rRulesActionUrlRewrite, ok := rRulesAction["urlRewrite"].(map[string]interface{}); ok {
									rRules.Action.UrlRewrite = &dclService.HttpRouteRulesActionUrlRewrite{}
									if _, ok := rRulesActionUrlRewrite["hostRewrite"]; ok {
										if s, ok := rRulesActionUrlRewrite["hostRewrite"].(string); ok {
											rRules.Action.UrlRewrite.HostRewrite = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRules.Action.UrlRewrite.HostRewrite: expected string")
										}
									}
									if _, ok := rRulesActionUrlRewrite["pathPrefixRewrite"]; ok {
										if s, ok := rRulesActionUrlRewrite["pathPrefixRewrite"].(string); ok {
											rRules.Action.UrlRewrite.PathPrefixRewrite = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRules.Action.UrlRewrite.PathPrefixRewrite: expected string")
										}
									}
								} else {
									return nil, fmt.Errorf("rRules.Action.UrlRewrite: expected map[string]interface{}")
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
									var rRulesMatches dclService.HttpRouteRulesMatches
									if _, ok := objval["fullPathMatch"]; ok {
										if s, ok := objval["fullPathMatch"].(string); ok {
											rRulesMatches.FullPathMatch = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRulesMatches.FullPathMatch: expected string")
										}
									}
									if _, ok := objval["headers"]; ok {
										if s, ok := objval["headers"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rRulesMatchesHeaders dclService.HttpRouteRulesMatchesHeaders
													if _, ok := objval["exactMatch"]; ok {
														if s, ok := objval["exactMatch"].(string); ok {
															rRulesMatchesHeaders.ExactMatch = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.ExactMatch: expected string")
														}
													}
													if _, ok := objval["header"]; ok {
														if s, ok := objval["header"].(string); ok {
															rRulesMatchesHeaders.Header = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.Header: expected string")
														}
													}
													if _, ok := objval["invertMatch"]; ok {
														if b, ok := objval["invertMatch"].(bool); ok {
															rRulesMatchesHeaders.InvertMatch = dcl.Bool(b)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.InvertMatch: expected bool")
														}
													}
													if _, ok := objval["prefixMatch"]; ok {
														if s, ok := objval["prefixMatch"].(string); ok {
															rRulesMatchesHeaders.PrefixMatch = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.PrefixMatch: expected string")
														}
													}
													if _, ok := objval["presentMatch"]; ok {
														if b, ok := objval["presentMatch"].(bool); ok {
															rRulesMatchesHeaders.PresentMatch = dcl.Bool(b)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.PresentMatch: expected bool")
														}
													}
													if _, ok := objval["rangeMatch"]; ok {
														if rRulesMatchesHeadersRangeMatch, ok := objval["rangeMatch"].(map[string]interface{}); ok {
															rRulesMatchesHeaders.RangeMatch = &dclService.HttpRouteRulesMatchesHeadersRangeMatch{}
															if _, ok := rRulesMatchesHeadersRangeMatch["end"]; ok {
																if i, ok := rRulesMatchesHeadersRangeMatch["end"].(int64); ok {
																	rRulesMatchesHeaders.RangeMatch.End = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rRulesMatchesHeaders.RangeMatch.End: expected int64")
																}
															}
															if _, ok := rRulesMatchesHeadersRangeMatch["start"]; ok {
																if i, ok := rRulesMatchesHeadersRangeMatch["start"].(int64); ok {
																	rRulesMatchesHeaders.RangeMatch.Start = dcl.Int64(i)
																} else {
																	return nil, fmt.Errorf("rRulesMatchesHeaders.RangeMatch.Start: expected int64")
																}
															}
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.RangeMatch: expected map[string]interface{}")
														}
													}
													if _, ok := objval["regexMatch"]; ok {
														if s, ok := objval["regexMatch"].(string); ok {
															rRulesMatchesHeaders.RegexMatch = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.RegexMatch: expected string")
														}
													}
													if _, ok := objval["suffixMatch"]; ok {
														if s, ok := objval["suffixMatch"].(string); ok {
															rRulesMatchesHeaders.SuffixMatch = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesHeaders.SuffixMatch: expected string")
														}
													}
													rRulesMatches.Headers = append(rRulesMatches.Headers, rRulesMatchesHeaders)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesMatches.Headers: expected []interface{}")
										}
									}
									if _, ok := objval["ignoreCase"]; ok {
										if b, ok := objval["ignoreCase"].(bool); ok {
											rRulesMatches.IgnoreCase = dcl.Bool(b)
										} else {
											return nil, fmt.Errorf("rRulesMatches.IgnoreCase: expected bool")
										}
									}
									if _, ok := objval["prefixMatch"]; ok {
										if s, ok := objval["prefixMatch"].(string); ok {
											rRulesMatches.PrefixMatch = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRulesMatches.PrefixMatch: expected string")
										}
									}
									if _, ok := objval["queryParameters"]; ok {
										if s, ok := objval["queryParameters"].([]interface{}); ok {
											for _, o := range s {
												if objval, ok := o.(map[string]interface{}); ok {
													var rRulesMatchesQueryParameters dclService.HttpRouteRulesMatchesQueryParameters
													if _, ok := objval["exactMatch"]; ok {
														if s, ok := objval["exactMatch"].(string); ok {
															rRulesMatchesQueryParameters.ExactMatch = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesQueryParameters.ExactMatch: expected string")
														}
													}
													if _, ok := objval["presentMatch"]; ok {
														if b, ok := objval["presentMatch"].(bool); ok {
															rRulesMatchesQueryParameters.PresentMatch = dcl.Bool(b)
														} else {
															return nil, fmt.Errorf("rRulesMatchesQueryParameters.PresentMatch: expected bool")
														}
													}
													if _, ok := objval["queryParameter"]; ok {
														if s, ok := objval["queryParameter"].(string); ok {
															rRulesMatchesQueryParameters.QueryParameter = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesQueryParameters.QueryParameter: expected string")
														}
													}
													if _, ok := objval["regexMatch"]; ok {
														if s, ok := objval["regexMatch"].(string); ok {
															rRulesMatchesQueryParameters.RegexMatch = dcl.String(s)
														} else {
															return nil, fmt.Errorf("rRulesMatchesQueryParameters.RegexMatch: expected string")
														}
													}
													rRulesMatches.QueryParameters = append(rRulesMatches.QueryParameters, rRulesMatchesQueryParameters)
												}
											}
										} else {
											return nil, fmt.Errorf("rRulesMatches.QueryParameters: expected []interface{}")
										}
									}
									if _, ok := objval["regexMatch"]; ok {
										if s, ok := objval["regexMatch"].(string); ok {
											rRulesMatches.RegexMatch = dcl.String(s)
										} else {
											return nil, fmt.Errorf("rRulesMatches.RegexMatch: expected string")
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

func GetHttpRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToHttpRoute(u)
	if err != nil {
		return nil, err
	}
	r, err = c.GetHttpRoute(ctx, r)
	if err != nil {
		return nil, err
	}
	return HttpRouteToUnstructured(r), nil
}

func ListHttpRoute(ctx context.Context, config *dcl.Config, project string, location string) ([]*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	l, err := c.ListHttpRoute(ctx, project, location)
	if err != nil {
		return nil, err
	}
	var resources []*unstructured.Resource
	for {
		for _, r := range l.Items {
			resources = append(resources, HttpRouteToUnstructured(r))
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

func ApplyHttpRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToHttpRoute(u)
	if err != nil {
		return nil, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToHttpRoute(ush)
		if err != nil {
			return nil, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	r, err = c.ApplyHttpRoute(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return HttpRouteToUnstructured(r), nil
}

func HttpRouteHasDiff(ctx context.Context, config *dcl.Config, u *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	c := dclService.NewClient(config)
	r, err := UnstructuredToHttpRoute(u)
	if err != nil {
		return false, err
	}
	if ush := unstructured.FetchStateHint(opts); ush != nil {
		sh, err := UnstructuredToHttpRoute(ush)
		if err != nil {
			return false, err
		}
		opts = append(opts, dcl.WithStateHint(sh))
	}
	opts = append(opts, dcl.WithLifecycleParam(dcl.BlockDestruction), dcl.WithLifecycleParam(dcl.BlockCreation), dcl.WithLifecycleParam(dcl.BlockModification))
	_, err = c.ApplyHttpRoute(ctx, r, opts...)
	if err != nil {
		if _, ok := err.(dcl.ApplyInfeasibleError); ok {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func DeleteHttpRoute(ctx context.Context, config *dcl.Config, u *unstructured.Resource) error {
	c := dclService.NewClient(config)
	r, err := UnstructuredToHttpRoute(u)
	if err != nil {
		return err
	}
	return c.DeleteHttpRoute(ctx, r)
}

func HttpRouteID(u *unstructured.Resource) (string, error) {
	r, err := UnstructuredToHttpRoute(u)
	if err != nil {
		return "", err
	}
	return r.ID()
}

func (r *HttpRoute) STV() unstructured.ServiceTypeVersion {
	return unstructured.ServiceTypeVersion{
		"networkservices",
		"HttpRoute",
		"ga",
	}
}

func (r *HttpRoute) SetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *HttpRoute) GetPolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, role, member string) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *HttpRoute) DeletePolicyMember(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, member *unstructured.Resource) error {
	return unstructured.ErrNoSuchMethod
}

func (r *HttpRoute) SetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *HttpRoute) SetPolicyWithEtag(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, policy *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *HttpRoute) GetPolicy(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return nil, unstructured.ErrNoSuchMethod
}

func (r *HttpRoute) Get(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) (*unstructured.Resource, error) {
	return GetHttpRoute(ctx, config, resource)
}

func (r *HttpRoute) Apply(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (*unstructured.Resource, error) {
	return ApplyHttpRoute(ctx, config, resource, opts...)
}

func (r *HttpRoute) HasDiff(ctx context.Context, config *dcl.Config, resource *unstructured.Resource, opts ...dcl.ApplyOption) (bool, error) {
	return HttpRouteHasDiff(ctx, config, resource, opts...)
}

func (r *HttpRoute) Delete(ctx context.Context, config *dcl.Config, resource *unstructured.Resource) error {
	return DeleteHttpRoute(ctx, config, resource)
}

func (r *HttpRoute) ID(resource *unstructured.Resource) (string, error) {
	return HttpRouteID(resource)
}

func init() {
	unstructured.Register(&HttpRoute{})
}
