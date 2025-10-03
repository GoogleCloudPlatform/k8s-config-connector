// Copyright 2021 Google LLC. All Rights Reserved.
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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for UrlMap.
type UrlMapServer struct{}

// ProtoToUrlMapDefaultUrlRedirectRedirectResponseCodeEnum converts a UrlMapDefaultUrlRedirectRedirectResponseCodeEnum enum from its proto representation.
func ProtoToComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(e betapb.ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum) *beta.UrlMapDefaultUrlRedirectRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := beta.UrlMapDefaultUrlRedirectRedirectResponseCodeEnum(n[len("ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum converts a UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum enum from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(e betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum) *beta.UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := beta.UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(n[len("ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum converts a UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum enum from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(e betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum) *beta.UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := beta.UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(n[len("ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum enum from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(e betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum) *beta.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum_name[int32(e)]; ok {
		e := beta.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(n[len("ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum converts a UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum enum from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(e betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum) *beta.UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := beta.UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(n[len("ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapDefaultRouteAction converts a UrlMapDefaultRouteAction resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteAction(p *betapb.ComputeBetaUrlMapDefaultRouteAction) *beta.UrlMapDefaultRouteAction {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteAction{
		UrlRewrite:           ProtoToComputeBetaUrlMapDefaultRouteActionUrlRewrite(p.GetUrlRewrite()),
		Timeout:              ProtoToComputeBetaUrlMapDefaultRouteActionTimeout(p.GetTimeout()),
		RetryPolicy:          ProtoToComputeBetaUrlMapDefaultRouteActionRetryPolicy(p.GetRetryPolicy()),
		RequestMirrorPolicy:  ProtoToComputeBetaUrlMapDefaultRouteActionRequestMirrorPolicy(p.GetRequestMirrorPolicy()),
		CorsPolicy:           ProtoToComputeBetaUrlMapDefaultRouteActionCorsPolicy(p.GetCorsPolicy()),
		FaultInjectionPolicy: ProtoToComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
	}
	for _, r := range p.GetWeightedBackendService() {
		obj.WeightedBackendService = append(obj.WeightedBackendService, *ProtoToComputeBetaUrlMapDefaultRouteActionWeightedBackendService(r))
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionWeightedBackendService converts a UrlMapDefaultRouteActionWeightedBackendService resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionWeightedBackendService(p *betapb.ComputeBetaUrlMapDefaultRouteActionWeightedBackendService) *beta.UrlMapDefaultRouteActionWeightedBackendService {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionWeightedBackendService{
		BackendService: dcl.StringOrNil(p.BackendService),
		Weight:         dcl.Int64OrNil(p.Weight),
		HeaderAction:   ProtoToComputeBetaUrlMapHeaderAction(p.GetHeaderAction()),
	}
	return obj
}

// ProtoToUrlMapHeaderAction converts a UrlMapHeaderAction resource from its proto representation.
func ProtoToComputeBetaUrlMapHeaderAction(p *betapb.ComputeBetaUrlMapHeaderAction) *beta.UrlMapHeaderAction {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapHeaderAction{}
	for _, r := range p.GetRequestHeadersToRemove() {
		obj.RequestHeadersToRemove = append(obj.RequestHeadersToRemove, r)
	}
	for _, r := range p.GetRequestHeadersToAdd() {
		obj.RequestHeadersToAdd = append(obj.RequestHeadersToAdd, *ProtoToComputeBetaUrlMapHeaderActionRequestHeadersToAdd(r))
	}
	for _, r := range p.GetResponseHeadersToRemove() {
		obj.ResponseHeadersToRemove = append(obj.ResponseHeadersToRemove, r)
	}
	for _, r := range p.GetResponseHeadersToAdd() {
		obj.ResponseHeadersToAdd = append(obj.ResponseHeadersToAdd, *ProtoToComputeBetaUrlMapHeaderActionResponseHeadersToAdd(r))
	}
	return obj
}

// ProtoToUrlMapHeaderActionRequestHeadersToAdd converts a UrlMapHeaderActionRequestHeadersToAdd resource from its proto representation.
func ProtoToComputeBetaUrlMapHeaderActionRequestHeadersToAdd(p *betapb.ComputeBetaUrlMapHeaderActionRequestHeadersToAdd) *beta.UrlMapHeaderActionRequestHeadersToAdd {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapHeaderActionRequestHeadersToAdd{
		HeaderName:  dcl.StringOrNil(p.HeaderName),
		HeaderValue: dcl.StringOrNil(p.HeaderValue),
		Replace:     dcl.Bool(p.Replace),
	}
	return obj
}

// ProtoToUrlMapHeaderActionResponseHeadersToAdd converts a UrlMapHeaderActionResponseHeadersToAdd resource from its proto representation.
func ProtoToComputeBetaUrlMapHeaderActionResponseHeadersToAdd(p *betapb.ComputeBetaUrlMapHeaderActionResponseHeadersToAdd) *beta.UrlMapHeaderActionResponseHeadersToAdd {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapHeaderActionResponseHeadersToAdd{
		HeaderName:  dcl.StringOrNil(p.HeaderName),
		HeaderValue: dcl.StringOrNil(p.HeaderValue),
		Replace:     dcl.Bool(p.Replace),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionUrlRewrite converts a UrlMapDefaultRouteActionUrlRewrite resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionUrlRewrite(p *betapb.ComputeBetaUrlMapDefaultRouteActionUrlRewrite) *beta.UrlMapDefaultRouteActionUrlRewrite {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.StringOrNil(p.PathPrefixRewrite),
		HostRewrite:       dcl.StringOrNil(p.HostRewrite),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionTimeout converts a UrlMapDefaultRouteActionTimeout resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionTimeout(p *betapb.ComputeBetaUrlMapDefaultRouteActionTimeout) *beta.UrlMapDefaultRouteActionTimeout {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionRetryPolicy converts a UrlMapDefaultRouteActionRetryPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionRetryPolicy(p *betapb.ComputeBetaUrlMapDefaultRouteActionRetryPolicy) *beta.UrlMapDefaultRouteActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionRetryPolicy{
		NumRetries:    dcl.Int64OrNil(p.NumRetries),
		PerTryTimeout: ProtoToComputeBetaUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(p.GetPerTryTimeout()),
	}
	for _, r := range p.GetRetryCondition() {
		obj.RetryCondition = append(obj.RetryCondition, r)
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionRetryPolicyPerTryTimeout converts a UrlMapDefaultRouteActionRetryPolicyPerTryTimeout resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(p *betapb.ComputeBetaUrlMapDefaultRouteActionRetryPolicyPerTryTimeout) *beta.UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionRequestMirrorPolicy converts a UrlMapDefaultRouteActionRequestMirrorPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionRequestMirrorPolicy(p *betapb.ComputeBetaUrlMapDefaultRouteActionRequestMirrorPolicy) *beta.UrlMapDefaultRouteActionRequestMirrorPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionRequestMirrorPolicy{
		BackendService: dcl.StringOrNil(p.BackendService),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionCorsPolicy converts a UrlMapDefaultRouteActionCorsPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionCorsPolicy(p *betapb.ComputeBetaUrlMapDefaultRouteActionCorsPolicy) *beta.UrlMapDefaultRouteActionCorsPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionCorsPolicy{
		MaxAge:           dcl.Int64OrNil(p.MaxAge),
		AllowCredentials: dcl.Bool(p.AllowCredentials),
		Disabled:         dcl.Bool(p.Disabled),
	}
	for _, r := range p.GetAllowOrigin() {
		obj.AllowOrigin = append(obj.AllowOrigin, r)
	}
	for _, r := range p.GetAllowOriginRegex() {
		obj.AllowOriginRegex = append(obj.AllowOriginRegex, r)
	}
	for _, r := range p.GetAllowMethod() {
		obj.AllowMethod = append(obj.AllowMethod, r)
	}
	for _, r := range p.GetAllowHeader() {
		obj.AllowHeader = append(obj.AllowHeader, r)
	}
	for _, r := range p.GetExposeHeader() {
		obj.ExposeHeader = append(obj.ExposeHeader, r)
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionFaultInjectionPolicy converts a UrlMapDefaultRouteActionFaultInjectionPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicy(p *betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicy) *beta.UrlMapDefaultRouteActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionFaultInjectionPolicy{
		Delay: ProtoToComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionFaultInjectionPolicyDelay converts a UrlMapDefaultRouteActionFaultInjectionPolicyDelay resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelay(p *betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelay) *beta.UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ProtoToComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(p.GetFixedDelay()),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay converts a UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(p *betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) *beta.UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionFaultInjectionPolicyAbort converts a UrlMapDefaultRouteActionFaultInjectionPolicyAbort resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyAbort(p *betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyAbort) *beta.UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.HttpStatus),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapDefaultUrlRedirect converts a UrlMapDefaultUrlRedirect resource from its proto representation.
func ProtoToComputeBetaUrlMapDefaultUrlRedirect(p *betapb.ComputeBetaUrlMapDefaultUrlRedirect) *beta.UrlMapDefaultUrlRedirect {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapDefaultUrlRedirect{
		HostRedirect:         dcl.StringOrNil(p.HostRedirect),
		PathRedirect:         dcl.StringOrNil(p.PathRedirect),
		PrefixRedirect:       dcl.StringOrNil(p.PrefixRedirect),
		RedirectResponseCode: ProtoToComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(p.GetRedirectResponseCode()),
		HttpsRedirect:        dcl.Bool(p.HttpsRedirect),
		StripQuery:           dcl.Bool(p.StripQuery),
	}
	return obj
}

// ProtoToUrlMapHostRule converts a UrlMapHostRule resource from its proto representation.
func ProtoToComputeBetaUrlMapHostRule(p *betapb.ComputeBetaUrlMapHostRule) *beta.UrlMapHostRule {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapHostRule{
		Description: dcl.StringOrNil(p.Description),
		PathMatcher: dcl.StringOrNil(p.PathMatcher),
	}
	for _, r := range p.GetHost() {
		obj.Host = append(obj.Host, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcher converts a UrlMapPathMatcher resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcher(p *betapb.ComputeBetaUrlMapPathMatcher) *beta.UrlMapPathMatcher {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcher{
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		DefaultService:     dcl.StringOrNil(p.DefaultService),
		DefaultRouteAction: ProtoToComputeBetaUrlMapDefaultRouteAction(p.GetDefaultRouteAction()),
		DefaultUrlRedirect: ProtoToComputeBetaUrlMapPathMatcherDefaultUrlRedirect(p.GetDefaultUrlRedirect()),
		HeaderAction:       ProtoToComputeBetaUrlMapHeaderAction(p.GetHeaderAction()),
	}
	for _, r := range p.GetPathRule() {
		obj.PathRule = append(obj.PathRule, *ProtoToComputeBetaUrlMapPathMatcherPathRule(r))
	}
	for _, r := range p.GetRouteRule() {
		obj.RouteRule = append(obj.RouteRule, *ProtoToComputeBetaUrlMapPathMatcherRouteRule(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherDefaultUrlRedirect converts a UrlMapPathMatcherDefaultUrlRedirect resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherDefaultUrlRedirect(p *betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirect) *beta.UrlMapPathMatcherDefaultUrlRedirect {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherDefaultUrlRedirect{
		HostRedirect:         dcl.StringOrNil(p.HostRedirect),
		PathRedirect:         dcl.StringOrNil(p.PathRedirect),
		PrefixRedirect:       dcl.StringOrNil(p.PrefixRedirect),
		RedirectResponseCode: ProtoToComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(p.GetRedirectResponseCode()),
		HttpsRedirect:        dcl.Bool(p.HttpsRedirect),
		StripQuery:           dcl.Bool(p.StripQuery),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRule converts a UrlMapPathMatcherPathRule resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRule(p *betapb.ComputeBetaUrlMapPathMatcherPathRule) *beta.UrlMapPathMatcherPathRule {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRule{
		BackendService: dcl.StringOrNil(p.BackendService),
		RouteAction:    ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteAction(p.GetRouteAction()),
		UrlRedirect:    ProtoToComputeBetaUrlMapPathMatcherPathRuleUrlRedirect(p.GetUrlRedirect()),
	}
	for _, r := range p.GetPath() {
		obj.Path = append(obj.Path, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteAction converts a UrlMapPathMatcherPathRuleRouteAction resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteAction(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteAction) *beta.UrlMapPathMatcherPathRuleRouteAction {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteAction{
		UrlRewrite:           ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionUrlRewrite(p.GetUrlRewrite()),
		Timeout:              ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionTimeout(p.GetTimeout()),
		RetryPolicy:          ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicy(p.GetRetryPolicy()),
		RequestMirrorPolicy:  ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(p.GetRequestMirrorPolicy()),
		CorsPolicy:           ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionCorsPolicy(p.GetCorsPolicy()),
		FaultInjectionPolicy: ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
	}
	for _, r := range p.GetWeightedBackendService() {
		obj.WeightedBackendService = append(obj.WeightedBackendService, *ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionWeightedBackendService converts a UrlMapPathMatcherPathRuleRouteActionWeightedBackendService resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionWeightedBackendService) *beta.UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{
		BackendService: dcl.StringOrNil(p.BackendService),
		Weight:         dcl.Int64OrNil(p.Weight),
		HeaderAction:   ProtoToComputeBetaUrlMapHeaderAction(p.GetHeaderAction()),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionUrlRewrite converts a UrlMapPathMatcherPathRuleRouteActionUrlRewrite resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionUrlRewrite(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionUrlRewrite) *beta.UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.StringOrNil(p.PathPrefixRewrite),
		HostRewrite:       dcl.StringOrNil(p.HostRewrite),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionTimeout converts a UrlMapPathMatcherPathRuleRouteActionTimeout resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionTimeout(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionTimeout) *beta.UrlMapPathMatcherPathRuleRouteActionTimeout {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionRetryPolicy converts a UrlMapPathMatcherPathRuleRouteActionRetryPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicy(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicy) *beta.UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionRetryPolicy{
		NumRetries:    dcl.Int64OrNil(p.NumRetries),
		PerTryTimeout: ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(p.GetPerTryTimeout()),
	}
	for _, r := range p.GetRetryCondition() {
		obj.RetryCondition = append(obj.RetryCondition, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout converts a UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) *beta.UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy converts a UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) *beta.UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{
		BackendService: dcl.StringOrNil(p.BackendService),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionCorsPolicy converts a UrlMapPathMatcherPathRuleRouteActionCorsPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionCorsPolicy(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionCorsPolicy) *beta.UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionCorsPolicy{
		MaxAge:           dcl.Int64OrNil(p.MaxAge),
		AllowCredentials: dcl.Bool(p.AllowCredentials),
		Disabled:         dcl.Bool(p.Disabled),
	}
	for _, r := range p.GetAllowOrigin() {
		obj.AllowOrigin = append(obj.AllowOrigin, r)
	}
	for _, r := range p.GetAllowOriginRegex() {
		obj.AllowOriginRegex = append(obj.AllowOriginRegex, r)
	}
	for _, r := range p.GetAllowMethod() {
		obj.AllowMethod = append(obj.AllowMethod, r)
	}
	for _, r := range p.GetAllowHeader() {
		obj.AllowHeader = append(obj.AllowHeader, r)
	}
	for _, r := range p.GetExposeHeader() {
		obj.ExposeHeader = append(obj.ExposeHeader, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) *beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{
		Delay: ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) *beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(p.GetFixedDelay()),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) *beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.HttpStatus),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleUrlRedirect converts a UrlMapPathMatcherPathRuleUrlRedirect resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherPathRuleUrlRedirect(p *betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirect) *beta.UrlMapPathMatcherPathRuleUrlRedirect {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherPathRuleUrlRedirect{
		HostRedirect:         dcl.StringOrNil(p.HostRedirect),
		PathRedirect:         dcl.StringOrNil(p.PathRedirect),
		PrefixRedirect:       dcl.StringOrNil(p.PrefixRedirect),
		RedirectResponseCode: ProtoToComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(p.GetRedirectResponseCode()),
		HttpsRedirect:        dcl.Bool(p.HttpsRedirect),
		StripQuery:           dcl.Bool(p.StripQuery),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRule converts a UrlMapPathMatcherRouteRule resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRule(p *betapb.ComputeBetaUrlMapPathMatcherRouteRule) *beta.UrlMapPathMatcherRouteRule {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRule{
		Priority:       dcl.Int64OrNil(p.Priority),
		Description:    dcl.StringOrNil(p.Description),
		BackendService: dcl.StringOrNil(p.BackendService),
		RouteAction:    ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteAction(p.GetRouteAction()),
		UrlRedirect:    ProtoToComputeBetaUrlMapPathMatcherRouteRuleUrlRedirect(p.GetUrlRedirect()),
		HeaderAction:   ProtoToComputeBetaUrlMapHeaderAction(p.GetHeaderAction()),
	}
	for _, r := range p.GetMatchRule() {
		obj.MatchRule = append(obj.MatchRule, *ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRule(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRule converts a UrlMapPathMatcherRouteRuleMatchRule resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRule(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRule) *beta.UrlMapPathMatcherRouteRuleMatchRule {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleMatchRule{
		PrefixMatch:   dcl.StringOrNil(p.PrefixMatch),
		FullPathMatch: dcl.StringOrNil(p.FullPathMatch),
		RegexMatch:    dcl.StringOrNil(p.RegexMatch),
		IgnoreCase:    dcl.Bool(p.IgnoreCase),
	}
	for _, r := range p.GetHeaderMatch() {
		obj.HeaderMatch = append(obj.HeaderMatch, *ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(r))
	}
	for _, r := range p.GetQueryParameterMatch() {
		obj.QueryParameterMatch = append(obj.QueryParameterMatch, *ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(r))
	}
	for _, r := range p.GetMetadataFilter() {
		obj.MetadataFilter = append(obj.MetadataFilter, *ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch converts a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) *beta.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{
		HeaderName:   dcl.StringOrNil(p.HeaderName),
		ExactMatch:   dcl.StringOrNil(p.ExactMatch),
		RegexMatch:   dcl.StringOrNil(p.RegexMatch),
		RangeMatch:   ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(p.GetRangeMatch()),
		PresentMatch: dcl.Bool(p.PresentMatch),
		PrefixMatch:  dcl.StringOrNil(p.PrefixMatch),
		SuffixMatch:  dcl.StringOrNil(p.SuffixMatch),
		InvertMatch:  dcl.Bool(p.InvertMatch),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch converts a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) *beta.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{
		RangeStart: dcl.Int64OrNil(p.RangeStart),
		RangeEnd:   dcl.Int64OrNil(p.RangeEnd),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch converts a UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) *beta.UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{
		Name:         dcl.StringOrNil(p.Name),
		PresentMatch: dcl.Bool(p.PresentMatch),
		ExactMatch:   dcl.StringOrNil(p.ExactMatch),
		RegexMatch:   dcl.StringOrNil(p.RegexMatch),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) *beta.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{
		FilterMatchCriteria: ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(p.GetFilterMatchCriteria()),
	}
	for _, r := range p.GetFilterLabel() {
		obj.FilterLabel = append(obj.FilterLabel, *ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) *beta.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteAction converts a UrlMapPathMatcherRouteRuleRouteAction resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteAction(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteAction) *beta.UrlMapPathMatcherRouteRuleRouteAction {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteAction{
		UrlRewrite:           ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(p.GetUrlRewrite()),
		Timeout:              ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionTimeout(p.GetTimeout()),
		RetryPolicy:          ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(p.GetRetryPolicy()),
		RequestMirrorPolicy:  ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(p.GetRequestMirrorPolicy()),
		CorsPolicy:           ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(p.GetCorsPolicy()),
		FaultInjectionPolicy: ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
	}
	for _, r := range p.GetWeightedBackendService() {
		obj.WeightedBackendService = append(obj.WeightedBackendService, *ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService converts a UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) *beta.UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{
		BackendService: dcl.StringOrNil(p.BackendService),
		Weight:         dcl.Int64OrNil(p.Weight),
		HeaderAction:   ProtoToComputeBetaUrlMapHeaderAction(p.GetHeaderAction()),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionUrlRewrite converts a UrlMapPathMatcherRouteRuleRouteActionUrlRewrite resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionUrlRewrite) *beta.UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.StringOrNil(p.PathPrefixRewrite),
		HostRewrite:       dcl.StringOrNil(p.HostRewrite),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionTimeout converts a UrlMapPathMatcherRouteRuleRouteActionTimeout resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionTimeout(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionTimeout) *beta.UrlMapPathMatcherRouteRuleRouteActionTimeout {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionRetryPolicy converts a UrlMapPathMatcherRouteRuleRouteActionRetryPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicy) *beta.UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{
		NumRetries:    dcl.Int64OrNil(p.NumRetries),
		PerTryTimeout: ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(p.GetPerTryTimeout()),
	}
	for _, r := range p.GetRetryCondition() {
		obj.RetryCondition = append(obj.RetryCondition, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout converts a UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) *beta.UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy converts a UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) *beta.UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{
		BackendService: dcl.StringOrNil(p.BackendService),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionCorsPolicy converts a UrlMapPathMatcherRouteRuleRouteActionCorsPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionCorsPolicy) *beta.UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{
		MaxAge:           dcl.Int64OrNil(p.MaxAge),
		AllowCredentials: dcl.Bool(p.AllowCredentials),
		Disabled:         dcl.Bool(p.Disabled),
	}
	for _, r := range p.GetAllowOrigin() {
		obj.AllowOrigin = append(obj.AllowOrigin, r)
	}
	for _, r := range p.GetAllowOriginRegex() {
		obj.AllowOriginRegex = append(obj.AllowOriginRegex, r)
	}
	for _, r := range p.GetAllowMethod() {
		obj.AllowMethod = append(obj.AllowMethod, r)
	}
	for _, r := range p.GetAllowHeader() {
		obj.AllowHeader = append(obj.AllowHeader, r)
	}
	for _, r := range p.GetExposeHeader() {
		obj.ExposeHeader = append(obj.ExposeHeader, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) *beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{
		Delay: ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) *beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(p.GetFixedDelay()),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) *beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.HttpStatus),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleUrlRedirect converts a UrlMapPathMatcherRouteRuleUrlRedirect resource from its proto representation.
func ProtoToComputeBetaUrlMapPathMatcherRouteRuleUrlRedirect(p *betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirect) *beta.UrlMapPathMatcherRouteRuleUrlRedirect {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapPathMatcherRouteRuleUrlRedirect{
		HostRedirect:         dcl.StringOrNil(p.HostRedirect),
		PathRedirect:         dcl.StringOrNil(p.PathRedirect),
		PrefixRedirect:       dcl.StringOrNil(p.PrefixRedirect),
		RedirectResponseCode: ProtoToComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(p.GetRedirectResponseCode()),
		HttpsRedirect:        dcl.Bool(p.HttpsRedirect),
		StripQuery:           dcl.Bool(p.StripQuery),
	}
	return obj
}

// ProtoToUrlMapTest converts a UrlMapTest resource from its proto representation.
func ProtoToComputeBetaUrlMapTest(p *betapb.ComputeBetaUrlMapTest) *beta.UrlMapTest {
	if p == nil {
		return nil
	}
	obj := &beta.UrlMapTest{
		Description:            dcl.StringOrNil(p.Description),
		Host:                   dcl.StringOrNil(p.Host),
		Path:                   dcl.StringOrNil(p.Path),
		ExpectedBackendService: dcl.StringOrNil(p.ExpectedBackendService),
	}
	return obj
}

// ProtoToUrlMap converts a UrlMap resource from its proto representation.
func ProtoToUrlMap(p *betapb.ComputeBetaUrlMap) *beta.UrlMap {
	obj := &beta.UrlMap{
		DefaultRouteAction: ProtoToComputeBetaUrlMapDefaultRouteAction(p.GetDefaultRouteAction()),
		DefaultService:     dcl.StringOrNil(p.DefaultService),
		DefaultUrlRedirect: ProtoToComputeBetaUrlMapDefaultUrlRedirect(p.GetDefaultUrlRedirect()),
		Description:        dcl.StringOrNil(p.Description),
		SelfLink:           dcl.StringOrNil(p.SelfLink),
		HeaderAction:       ProtoToComputeBetaUrlMapHeaderAction(p.GetHeaderAction()),
		Name:               dcl.StringOrNil(p.Name),
		Region:             dcl.StringOrNil(p.Region),
		Project:            dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetHostRule() {
		obj.HostRule = append(obj.HostRule, *ProtoToComputeBetaUrlMapHostRule(r))
	}
	for _, r := range p.GetPathMatcher() {
		obj.PathMatcher = append(obj.PathMatcher, *ProtoToComputeBetaUrlMapPathMatcher(r))
	}
	for _, r := range p.GetTest() {
		obj.Test = append(obj.Test, *ProtoToComputeBetaUrlMapTest(r))
	}
	return obj
}

// UrlMapDefaultUrlRedirectRedirectResponseCodeEnumToProto converts a UrlMapDefaultUrlRedirectRedirectResponseCodeEnum enum to its proto representation.
func ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnumToProto(e *beta.UrlMapDefaultUrlRedirectRedirectResponseCodeEnum) betapb.ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum {
	if e == nil {
		return betapb.ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(0)
	}
	if v, ok := betapb.ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum_value["UrlMapDefaultUrlRedirectRedirectResponseCodeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(v)
	}
	return betapb.ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(0)
}

// UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumToProto converts a UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum enum to its proto representation.
func ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumToProto(e *beta.UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum) betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum {
	if e == nil {
		return betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(0)
	}
	if v, ok := betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum_value["UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(v)
	}
	return betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(0)
}

// UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumToProto converts a UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum enum to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumToProto(e *beta.UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum) betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum {
	if e == nil {
		return betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(0)
	}
	if v, ok := betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum_value["UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(v)
	}
	return betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(0)
}

// UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumToProto converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum enum to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumToProto(e *beta.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum) betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == nil {
		return betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(0)
	}
	if v, ok := betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum_value["UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum"+string(*e)]; ok {
		return betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(v)
	}
	return betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(0)
}

// UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumToProto converts a UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum enum to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumToProto(e *beta.UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum) betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum {
	if e == nil {
		return betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(0)
	}
	if v, ok := betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum_value["UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(v)
	}
	return betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(0)
}

// UrlMapDefaultRouteActionToProto converts a UrlMapDefaultRouteAction resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionToProto(o *beta.UrlMapDefaultRouteAction) *betapb.ComputeBetaUrlMapDefaultRouteAction {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteAction{
		UrlRewrite:           ComputeBetaUrlMapDefaultRouteActionUrlRewriteToProto(o.UrlRewrite),
		Timeout:              ComputeBetaUrlMapDefaultRouteActionTimeoutToProto(o.Timeout),
		RetryPolicy:          ComputeBetaUrlMapDefaultRouteActionRetryPolicyToProto(o.RetryPolicy),
		RequestMirrorPolicy:  ComputeBetaUrlMapDefaultRouteActionRequestMirrorPolicyToProto(o.RequestMirrorPolicy),
		CorsPolicy:           ComputeBetaUrlMapDefaultRouteActionCorsPolicyToProto(o.CorsPolicy),
		FaultInjectionPolicy: ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy),
	}
	for _, r := range o.WeightedBackendService {
		p.WeightedBackendService = append(p.WeightedBackendService, ComputeBetaUrlMapDefaultRouteActionWeightedBackendServiceToProto(&r))
	}
	return p
}

// UrlMapDefaultRouteActionWeightedBackendServiceToProto converts a UrlMapDefaultRouteActionWeightedBackendService resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionWeightedBackendServiceToProto(o *beta.UrlMapDefaultRouteActionWeightedBackendService) *betapb.ComputeBetaUrlMapDefaultRouteActionWeightedBackendService {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionWeightedBackendService{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		Weight:         dcl.ValueOrEmptyInt64(o.Weight),
		HeaderAction:   ComputeBetaUrlMapHeaderActionToProto(o.HeaderAction),
	}
	return p
}

// UrlMapHeaderActionToProto converts a UrlMapHeaderAction resource to its proto representation.
func ComputeBetaUrlMapHeaderActionToProto(o *beta.UrlMapHeaderAction) *betapb.ComputeBetaUrlMapHeaderAction {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapHeaderAction{}
	for _, r := range o.RequestHeadersToRemove {
		p.RequestHeadersToRemove = append(p.RequestHeadersToRemove, r)
	}
	for _, r := range o.RequestHeadersToAdd {
		p.RequestHeadersToAdd = append(p.RequestHeadersToAdd, ComputeBetaUrlMapHeaderActionRequestHeadersToAddToProto(&r))
	}
	for _, r := range o.ResponseHeadersToRemove {
		p.ResponseHeadersToRemove = append(p.ResponseHeadersToRemove, r)
	}
	for _, r := range o.ResponseHeadersToAdd {
		p.ResponseHeadersToAdd = append(p.ResponseHeadersToAdd, ComputeBetaUrlMapHeaderActionResponseHeadersToAddToProto(&r))
	}
	return p
}

// UrlMapHeaderActionRequestHeadersToAddToProto converts a UrlMapHeaderActionRequestHeadersToAdd resource to its proto representation.
func ComputeBetaUrlMapHeaderActionRequestHeadersToAddToProto(o *beta.UrlMapHeaderActionRequestHeadersToAdd) *betapb.ComputeBetaUrlMapHeaderActionRequestHeadersToAdd {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapHeaderActionRequestHeadersToAdd{
		HeaderName:  dcl.ValueOrEmptyString(o.HeaderName),
		HeaderValue: dcl.ValueOrEmptyString(o.HeaderValue),
		Replace:     dcl.ValueOrEmptyBool(o.Replace),
	}
	return p
}

// UrlMapHeaderActionResponseHeadersToAddToProto converts a UrlMapHeaderActionResponseHeadersToAdd resource to its proto representation.
func ComputeBetaUrlMapHeaderActionResponseHeadersToAddToProto(o *beta.UrlMapHeaderActionResponseHeadersToAdd) *betapb.ComputeBetaUrlMapHeaderActionResponseHeadersToAdd {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapHeaderActionResponseHeadersToAdd{
		HeaderName:  dcl.ValueOrEmptyString(o.HeaderName),
		HeaderValue: dcl.ValueOrEmptyString(o.HeaderValue),
		Replace:     dcl.ValueOrEmptyBool(o.Replace),
	}
	return p
}

// UrlMapDefaultRouteActionUrlRewriteToProto converts a UrlMapDefaultRouteActionUrlRewrite resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionUrlRewriteToProto(o *beta.UrlMapDefaultRouteActionUrlRewrite) *betapb.ComputeBetaUrlMapDefaultRouteActionUrlRewrite {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.ValueOrEmptyString(o.PathPrefixRewrite),
		HostRewrite:       dcl.ValueOrEmptyString(o.HostRewrite),
	}
	return p
}

// UrlMapDefaultRouteActionTimeoutToProto converts a UrlMapDefaultRouteActionTimeout resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionTimeoutToProto(o *beta.UrlMapDefaultRouteActionTimeout) *betapb.ComputeBetaUrlMapDefaultRouteActionTimeout {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapDefaultRouteActionRetryPolicyToProto converts a UrlMapDefaultRouteActionRetryPolicy resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionRetryPolicyToProto(o *beta.UrlMapDefaultRouteActionRetryPolicy) *betapb.ComputeBetaUrlMapDefaultRouteActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionRetryPolicy{
		NumRetries:    dcl.ValueOrEmptyInt64(o.NumRetries),
		PerTryTimeout: ComputeBetaUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutToProto(o.PerTryTimeout),
	}
	for _, r := range o.RetryCondition {
		p.RetryCondition = append(p.RetryCondition, r)
	}
	return p
}

// UrlMapDefaultRouteActionRetryPolicyPerTryTimeoutToProto converts a UrlMapDefaultRouteActionRetryPolicyPerTryTimeout resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutToProto(o *beta.UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) *betapb.ComputeBetaUrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapDefaultRouteActionRequestMirrorPolicyToProto converts a UrlMapDefaultRouteActionRequestMirrorPolicy resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionRequestMirrorPolicyToProto(o *beta.UrlMapDefaultRouteActionRequestMirrorPolicy) *betapb.ComputeBetaUrlMapDefaultRouteActionRequestMirrorPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionRequestMirrorPolicy{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
	}
	return p
}

// UrlMapDefaultRouteActionCorsPolicyToProto converts a UrlMapDefaultRouteActionCorsPolicy resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionCorsPolicyToProto(o *beta.UrlMapDefaultRouteActionCorsPolicy) *betapb.ComputeBetaUrlMapDefaultRouteActionCorsPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionCorsPolicy{
		MaxAge:           dcl.ValueOrEmptyInt64(o.MaxAge),
		AllowCredentials: dcl.ValueOrEmptyBool(o.AllowCredentials),
		Disabled:         dcl.ValueOrEmptyBool(o.Disabled),
	}
	for _, r := range o.AllowOrigin {
		p.AllowOrigin = append(p.AllowOrigin, r)
	}
	for _, r := range o.AllowOriginRegex {
		p.AllowOriginRegex = append(p.AllowOriginRegex, r)
	}
	for _, r := range o.AllowMethod {
		p.AllowMethod = append(p.AllowMethod, r)
	}
	for _, r := range o.AllowHeader {
		p.AllowHeader = append(p.AllowHeader, r)
	}
	for _, r := range o.ExposeHeader {
		p.ExposeHeader = append(p.ExposeHeader, r)
	}
	return p
}

// UrlMapDefaultRouteActionFaultInjectionPolicyToProto converts a UrlMapDefaultRouteActionFaultInjectionPolicy resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyToProto(o *beta.UrlMapDefaultRouteActionFaultInjectionPolicy) *betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicy{
		Delay: ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelayToProto(o.Delay),
		Abort: ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyAbortToProto(o.Abort),
	}
	return p
}

// UrlMapDefaultRouteActionFaultInjectionPolicyDelayToProto converts a UrlMapDefaultRouteActionFaultInjectionPolicyDelay resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelayToProto(o *beta.UrlMapDefaultRouteActionFaultInjectionPolicyDelay) *betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o.FixedDelay),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayToProto converts a UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o *beta.UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) *betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapDefaultRouteActionFaultInjectionPolicyAbortToProto converts a UrlMapDefaultRouteActionFaultInjectionPolicyAbort resource to its proto representation.
func ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyAbortToProto(o *beta.UrlMapDefaultRouteActionFaultInjectionPolicyAbort) *betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.ValueOrEmptyInt64(o.HttpStatus),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapDefaultUrlRedirectToProto converts a UrlMapDefaultUrlRedirect resource to its proto representation.
func ComputeBetaUrlMapDefaultUrlRedirectToProto(o *beta.UrlMapDefaultUrlRedirect) *betapb.ComputeBetaUrlMapDefaultUrlRedirect {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapDefaultUrlRedirect{
		HostRedirect:         dcl.ValueOrEmptyString(o.HostRedirect),
		PathRedirect:         dcl.ValueOrEmptyString(o.PathRedirect),
		PrefixRedirect:       dcl.ValueOrEmptyString(o.PrefixRedirect),
		RedirectResponseCode: ComputeBetaUrlMapDefaultUrlRedirectRedirectResponseCodeEnumToProto(o.RedirectResponseCode),
		HttpsRedirect:        dcl.ValueOrEmptyBool(o.HttpsRedirect),
		StripQuery:           dcl.ValueOrEmptyBool(o.StripQuery),
	}
	return p
}

// UrlMapHostRuleToProto converts a UrlMapHostRule resource to its proto representation.
func ComputeBetaUrlMapHostRuleToProto(o *beta.UrlMapHostRule) *betapb.ComputeBetaUrlMapHostRule {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapHostRule{
		Description: dcl.ValueOrEmptyString(o.Description),
		PathMatcher: dcl.ValueOrEmptyString(o.PathMatcher),
	}
	for _, r := range o.Host {
		p.Host = append(p.Host, r)
	}
	return p
}

// UrlMapPathMatcherToProto converts a UrlMapPathMatcher resource to its proto representation.
func ComputeBetaUrlMapPathMatcherToProto(o *beta.UrlMapPathMatcher) *betapb.ComputeBetaUrlMapPathMatcher {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcher{
		Name:               dcl.ValueOrEmptyString(o.Name),
		Description:        dcl.ValueOrEmptyString(o.Description),
		DefaultService:     dcl.ValueOrEmptyString(o.DefaultService),
		DefaultRouteAction: ComputeBetaUrlMapDefaultRouteActionToProto(o.DefaultRouteAction),
		DefaultUrlRedirect: ComputeBetaUrlMapPathMatcherDefaultUrlRedirectToProto(o.DefaultUrlRedirect),
		HeaderAction:       ComputeBetaUrlMapHeaderActionToProto(o.HeaderAction),
	}
	for _, r := range o.PathRule {
		p.PathRule = append(p.PathRule, ComputeBetaUrlMapPathMatcherPathRuleToProto(&r))
	}
	for _, r := range o.RouteRule {
		p.RouteRule = append(p.RouteRule, ComputeBetaUrlMapPathMatcherRouteRuleToProto(&r))
	}
	return p
}

// UrlMapPathMatcherDefaultUrlRedirectToProto converts a UrlMapPathMatcherDefaultUrlRedirect resource to its proto representation.
func ComputeBetaUrlMapPathMatcherDefaultUrlRedirectToProto(o *beta.UrlMapPathMatcherDefaultUrlRedirect) *betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirect {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherDefaultUrlRedirect{
		HostRedirect:         dcl.ValueOrEmptyString(o.HostRedirect),
		PathRedirect:         dcl.ValueOrEmptyString(o.PathRedirect),
		PrefixRedirect:       dcl.ValueOrEmptyString(o.PrefixRedirect),
		RedirectResponseCode: ComputeBetaUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumToProto(o.RedirectResponseCode),
		HttpsRedirect:        dcl.ValueOrEmptyBool(o.HttpsRedirect),
		StripQuery:           dcl.ValueOrEmptyBool(o.StripQuery),
	}
	return p
}

// UrlMapPathMatcherPathRuleToProto converts a UrlMapPathMatcherPathRule resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleToProto(o *beta.UrlMapPathMatcherPathRule) *betapb.ComputeBetaUrlMapPathMatcherPathRule {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRule{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		RouteAction:    ComputeBetaUrlMapPathMatcherPathRuleRouteActionToProto(o.RouteAction),
		UrlRedirect:    ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectToProto(o.UrlRedirect),
	}
	for _, r := range o.Path {
		p.Path = append(p.Path, r)
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionToProto converts a UrlMapPathMatcherPathRuleRouteAction resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionToProto(o *beta.UrlMapPathMatcherPathRuleRouteAction) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteAction {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteAction{
		UrlRewrite:           ComputeBetaUrlMapPathMatcherPathRuleRouteActionUrlRewriteToProto(o.UrlRewrite),
		Timeout:              ComputeBetaUrlMapPathMatcherPathRuleRouteActionTimeoutToProto(o.Timeout),
		RetryPolicy:          ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicyToProto(o.RetryPolicy),
		RequestMirrorPolicy:  ComputeBetaUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyToProto(o.RequestMirrorPolicy),
		CorsPolicy:           ComputeBetaUrlMapPathMatcherPathRuleRouteActionCorsPolicyToProto(o.CorsPolicy),
		FaultInjectionPolicy: ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy),
	}
	for _, r := range o.WeightedBackendService {
		p.WeightedBackendService = append(p.WeightedBackendService, ComputeBetaUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceToProto(&r))
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceToProto converts a UrlMapPathMatcherPathRuleRouteActionWeightedBackendService resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionWeightedBackendService{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		Weight:         dcl.ValueOrEmptyInt64(o.Weight),
		HeaderAction:   ComputeBetaUrlMapHeaderActionToProto(o.HeaderAction),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionUrlRewriteToProto converts a UrlMapPathMatcherPathRuleRouteActionUrlRewrite resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionUrlRewriteToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionUrlRewrite) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.ValueOrEmptyString(o.PathPrefixRewrite),
		HostRewrite:       dcl.ValueOrEmptyString(o.HostRewrite),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionTimeoutToProto converts a UrlMapPathMatcherPathRuleRouteActionTimeout resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionTimeoutToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionTimeout) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionTimeout {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionRetryPolicyToProto converts a UrlMapPathMatcherPathRuleRouteActionRetryPolicy resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicyToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionRetryPolicy) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicy{
		NumRetries:    dcl.ValueOrEmptyInt64(o.NumRetries),
		PerTryTimeout: ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutToProto(o.PerTryTimeout),
	}
	for _, r := range o.RetryCondition {
		p.RetryCondition = append(p.RetryCondition, r)
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutToProto converts a UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyToProto converts a UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionCorsPolicyToProto converts a UrlMapPathMatcherPathRuleRouteActionCorsPolicy resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionCorsPolicyToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionCorsPolicy) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionCorsPolicy{
		MaxAge:           dcl.ValueOrEmptyInt64(o.MaxAge),
		AllowCredentials: dcl.ValueOrEmptyBool(o.AllowCredentials),
		Disabled:         dcl.ValueOrEmptyBool(o.Disabled),
	}
	for _, r := range o.AllowOrigin {
		p.AllowOrigin = append(p.AllowOrigin, r)
	}
	for _, r := range o.AllowOriginRegex {
		p.AllowOriginRegex = append(p.AllowOriginRegex, r)
	}
	for _, r := range o.AllowMethod {
		p.AllowMethod = append(p.AllowMethod, r)
	}
	for _, r := range o.AllowHeader {
		p.AllowHeader = append(p.AllowHeader, r)
	}
	for _, r := range o.ExposeHeader {
		p.ExposeHeader = append(p.ExposeHeader, r)
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyToProto converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{
		Delay: ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayToProto(o.Delay),
		Abort: ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortToProto(o.Abort),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayToProto converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o.FixedDelay),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortToProto converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortToProto(o *beta.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) *betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.ValueOrEmptyInt64(o.HttpStatus),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapPathMatcherPathRuleUrlRedirectToProto converts a UrlMapPathMatcherPathRuleUrlRedirect resource to its proto representation.
func ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectToProto(o *beta.UrlMapPathMatcherPathRuleUrlRedirect) *betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirect {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherPathRuleUrlRedirect{
		HostRedirect:         dcl.ValueOrEmptyString(o.HostRedirect),
		PathRedirect:         dcl.ValueOrEmptyString(o.PathRedirect),
		PrefixRedirect:       dcl.ValueOrEmptyString(o.PrefixRedirect),
		RedirectResponseCode: ComputeBetaUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumToProto(o.RedirectResponseCode),
		HttpsRedirect:        dcl.ValueOrEmptyBool(o.HttpsRedirect),
		StripQuery:           dcl.ValueOrEmptyBool(o.StripQuery),
	}
	return p
}

// UrlMapPathMatcherRouteRuleToProto converts a UrlMapPathMatcherRouteRule resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleToProto(o *beta.UrlMapPathMatcherRouteRule) *betapb.ComputeBetaUrlMapPathMatcherRouteRule {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRule{
		Priority:       dcl.ValueOrEmptyInt64(o.Priority),
		Description:    dcl.ValueOrEmptyString(o.Description),
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		RouteAction:    ComputeBetaUrlMapPathMatcherRouteRuleRouteActionToProto(o.RouteAction),
		UrlRedirect:    ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectToProto(o.UrlRedirect),
		HeaderAction:   ComputeBetaUrlMapHeaderActionToProto(o.HeaderAction),
	}
	for _, r := range o.MatchRule {
		p.MatchRule = append(p.MatchRule, ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleToProto(&r))
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleToProto converts a UrlMapPathMatcherRouteRuleMatchRule resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleToProto(o *beta.UrlMapPathMatcherRouteRuleMatchRule) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRule {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRule{
		PrefixMatch:   dcl.ValueOrEmptyString(o.PrefixMatch),
		FullPathMatch: dcl.ValueOrEmptyString(o.FullPathMatch),
		RegexMatch:    dcl.ValueOrEmptyString(o.RegexMatch),
		IgnoreCase:    dcl.ValueOrEmptyBool(o.IgnoreCase),
	}
	for _, r := range o.HeaderMatch {
		p.HeaderMatch = append(p.HeaderMatch, ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchToProto(&r))
	}
	for _, r := range o.QueryParameterMatch {
		p.QueryParameterMatch = append(p.QueryParameterMatch, ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchToProto(&r))
	}
	for _, r := range o.MetadataFilter {
		p.MetadataFilter = append(p.MetadataFilter, ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterToProto(&r))
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchToProto converts a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchToProto(o *beta.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{
		HeaderName:   dcl.ValueOrEmptyString(o.HeaderName),
		ExactMatch:   dcl.ValueOrEmptyString(o.ExactMatch),
		RegexMatch:   dcl.ValueOrEmptyString(o.RegexMatch),
		RangeMatch:   ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchToProto(o.RangeMatch),
		PresentMatch: dcl.ValueOrEmptyBool(o.PresentMatch),
		PrefixMatch:  dcl.ValueOrEmptyString(o.PrefixMatch),
		SuffixMatch:  dcl.ValueOrEmptyString(o.SuffixMatch),
		InvertMatch:  dcl.ValueOrEmptyBool(o.InvertMatch),
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchToProto converts a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchToProto(o *beta.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{
		RangeStart: dcl.ValueOrEmptyInt64(o.RangeStart),
		RangeEnd:   dcl.ValueOrEmptyInt64(o.RangeEnd),
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchToProto converts a UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchToProto(o *beta.UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{
		Name:         dcl.ValueOrEmptyString(o.Name),
		PresentMatch: dcl.ValueOrEmptyBool(o.PresentMatch),
		ExactMatch:   dcl.ValueOrEmptyString(o.ExactMatch),
		RegexMatch:   dcl.ValueOrEmptyString(o.RegexMatch),
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterToProto converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterToProto(o *beta.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{
		FilterMatchCriteria: ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumToProto(o.FilterMatchCriteria),
	}
	for _, r := range o.FilterLabel {
		p.FilterLabel = append(p.FilterLabel, ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelToProto(&r))
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelToProto converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelToProto(o *beta.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionToProto converts a UrlMapPathMatcherRouteRuleRouteAction resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionToProto(o *beta.UrlMapPathMatcherRouteRuleRouteAction) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteAction {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteAction{
		UrlRewrite:           ComputeBetaUrlMapPathMatcherRouteRuleRouteActionUrlRewriteToProto(o.UrlRewrite),
		Timeout:              ComputeBetaUrlMapPathMatcherRouteRuleRouteActionTimeoutToProto(o.Timeout),
		RetryPolicy:          ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicyToProto(o.RetryPolicy),
		RequestMirrorPolicy:  ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyToProto(o.RequestMirrorPolicy),
		CorsPolicy:           ComputeBetaUrlMapPathMatcherRouteRuleRouteActionCorsPolicyToProto(o.CorsPolicy),
		FaultInjectionPolicy: ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy),
	}
	for _, r := range o.WeightedBackendService {
		p.WeightedBackendService = append(p.WeightedBackendService, ComputeBetaUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceToProto(&r))
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceToProto converts a UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		Weight:         dcl.ValueOrEmptyInt64(o.Weight),
		HeaderAction:   ComputeBetaUrlMapHeaderActionToProto(o.HeaderAction),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionUrlRewriteToProto converts a UrlMapPathMatcherRouteRuleRouteActionUrlRewrite resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionUrlRewriteToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.ValueOrEmptyString(o.PathPrefixRewrite),
		HostRewrite:       dcl.ValueOrEmptyString(o.HostRewrite),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionTimeoutToProto converts a UrlMapPathMatcherRouteRuleRouteActionTimeout resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionTimeoutToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionTimeout) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionTimeout {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionRetryPolicyToProto converts a UrlMapPathMatcherRouteRuleRouteActionRetryPolicy resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicyToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicy{
		NumRetries:    dcl.ValueOrEmptyInt64(o.NumRetries),
		PerTryTimeout: ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutToProto(o.PerTryTimeout),
	}
	for _, r := range o.RetryCondition {
		p.RetryCondition = append(p.RetryCondition, r)
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutToProto converts a UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyToProto converts a UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionCorsPolicyToProto converts a UrlMapPathMatcherRouteRuleRouteActionCorsPolicy resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionCorsPolicyToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionCorsPolicy{
		MaxAge:           dcl.ValueOrEmptyInt64(o.MaxAge),
		AllowCredentials: dcl.ValueOrEmptyBool(o.AllowCredentials),
		Disabled:         dcl.ValueOrEmptyBool(o.Disabled),
	}
	for _, r := range o.AllowOrigin {
		p.AllowOrigin = append(p.AllowOrigin, r)
	}
	for _, r := range o.AllowOriginRegex {
		p.AllowOriginRegex = append(p.AllowOriginRegex, r)
	}
	for _, r := range o.AllowMethod {
		p.AllowMethod = append(p.AllowMethod, r)
	}
	for _, r := range o.AllowHeader {
		p.AllowHeader = append(p.AllowHeader, r)
	}
	for _, r := range o.ExposeHeader {
		p.ExposeHeader = append(p.ExposeHeader, r)
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyToProto converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{
		Delay: ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayToProto(o.Delay),
		Abort: ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortToProto(o.Abort),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayToProto converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o.FixedDelay),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortToProto converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortToProto(o *beta.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.ValueOrEmptyInt64(o.HttpStatus),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapPathMatcherRouteRuleUrlRedirectToProto converts a UrlMapPathMatcherRouteRuleUrlRedirect resource to its proto representation.
func ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectToProto(o *beta.UrlMapPathMatcherRouteRuleUrlRedirect) *betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirect {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirect{
		HostRedirect:         dcl.ValueOrEmptyString(o.HostRedirect),
		PathRedirect:         dcl.ValueOrEmptyString(o.PathRedirect),
		PrefixRedirect:       dcl.ValueOrEmptyString(o.PrefixRedirect),
		RedirectResponseCode: ComputeBetaUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumToProto(o.RedirectResponseCode),
		HttpsRedirect:        dcl.ValueOrEmptyBool(o.HttpsRedirect),
		StripQuery:           dcl.ValueOrEmptyBool(o.StripQuery),
	}
	return p
}

// UrlMapTestToProto converts a UrlMapTest resource to its proto representation.
func ComputeBetaUrlMapTestToProto(o *beta.UrlMapTest) *betapb.ComputeBetaUrlMapTest {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaUrlMapTest{
		Description:            dcl.ValueOrEmptyString(o.Description),
		Host:                   dcl.ValueOrEmptyString(o.Host),
		Path:                   dcl.ValueOrEmptyString(o.Path),
		ExpectedBackendService: dcl.ValueOrEmptyString(o.ExpectedBackendService),
	}
	return p
}

// UrlMapToProto converts a UrlMap resource to its proto representation.
func UrlMapToProto(resource *beta.UrlMap) *betapb.ComputeBetaUrlMap {
	p := &betapb.ComputeBetaUrlMap{
		DefaultRouteAction: ComputeBetaUrlMapDefaultRouteActionToProto(resource.DefaultRouteAction),
		DefaultService:     dcl.ValueOrEmptyString(resource.DefaultService),
		DefaultUrlRedirect: ComputeBetaUrlMapDefaultUrlRedirectToProto(resource.DefaultUrlRedirect),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		SelfLink:           dcl.ValueOrEmptyString(resource.SelfLink),
		HeaderAction:       ComputeBetaUrlMapHeaderActionToProto(resource.HeaderAction),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Region:             dcl.ValueOrEmptyString(resource.Region),
		Project:            dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.HostRule {
		p.HostRule = append(p.HostRule, ComputeBetaUrlMapHostRuleToProto(&r))
	}
	for _, r := range resource.PathMatcher {
		p.PathMatcher = append(p.PathMatcher, ComputeBetaUrlMapPathMatcherToProto(&r))
	}
	for _, r := range resource.Test {
		p.Test = append(p.Test, ComputeBetaUrlMapTestToProto(&r))
	}

	return p
}

// ApplyUrlMap handles the gRPC request by passing it to the underlying UrlMap Apply() method.
func (s *UrlMapServer) applyUrlMap(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaUrlMapRequest) (*betapb.ComputeBetaUrlMap, error) {
	p := ProtoToUrlMap(request.GetResource())
	res, err := c.ApplyUrlMap(ctx, p)
	if err != nil {
		return nil, err
	}
	r := UrlMapToProto(res)
	return r, nil
}

// ApplyUrlMap handles the gRPC request by passing it to the underlying UrlMap Apply() method.
func (s *UrlMapServer) ApplyComputeBetaUrlMap(ctx context.Context, request *betapb.ApplyComputeBetaUrlMapRequest) (*betapb.ComputeBetaUrlMap, error) {
	cl, err := createConfigUrlMap(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyUrlMap(ctx, cl, request)
}

// DeleteUrlMap handles the gRPC request by passing it to the underlying UrlMap Delete() method.
func (s *UrlMapServer) DeleteComputeBetaUrlMap(ctx context.Context, request *betapb.DeleteComputeBetaUrlMapRequest) (*emptypb.Empty, error) {

	cl, err := createConfigUrlMap(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteUrlMap(ctx, ProtoToUrlMap(request.GetResource()))

}

// ListComputeBetaUrlMap handles the gRPC request by passing it to the underlying UrlMapList() method.
func (s *UrlMapServer) ListComputeBetaUrlMap(ctx context.Context, request *betapb.ListComputeBetaUrlMapRequest) (*betapb.ListComputeBetaUrlMapResponse, error) {
	cl, err := createConfigUrlMap(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListUrlMap(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaUrlMap
	for _, r := range resources.Items {
		rp := UrlMapToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaUrlMapResponse{Items: protos}, nil
}

func createConfigUrlMap(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
