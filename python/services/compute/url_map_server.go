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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for UrlMap.
type UrlMapServer struct{}

// ProtoToUrlMapDefaultUrlRedirectRedirectResponseCodeEnum converts a UrlMapDefaultUrlRedirectRedirectResponseCodeEnum enum from its proto representation.
func ProtoToComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(e computepb.ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum) *compute.UrlMapDefaultUrlRedirectRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := compute.UrlMapDefaultUrlRedirectRedirectResponseCodeEnum(n[len("ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum converts a UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum enum from its proto representation.
func ProtoToComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(e computepb.ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum) *compute.UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := compute.UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(n[len("ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum converts a UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum enum from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(e computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum) *compute.UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := compute.UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(n[len("ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum enum from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(e computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum) *compute.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum_name[int32(e)]; ok {
		e := compute.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(n[len("ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum converts a UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum enum from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(e computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum) *compute.UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := compute.UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(n[len("ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUrlMapDefaultRouteAction converts a UrlMapDefaultRouteAction resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteAction(p *computepb.ComputeUrlMapDefaultRouteAction) *compute.UrlMapDefaultRouteAction {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteAction{
		UrlRewrite:           ProtoToComputeUrlMapDefaultRouteActionUrlRewrite(p.GetUrlRewrite()),
		Timeout:              ProtoToComputeUrlMapDefaultRouteActionTimeout(p.GetTimeout()),
		RetryPolicy:          ProtoToComputeUrlMapDefaultRouteActionRetryPolicy(p.GetRetryPolicy()),
		RequestMirrorPolicy:  ProtoToComputeUrlMapDefaultRouteActionRequestMirrorPolicy(p.GetRequestMirrorPolicy()),
		CorsPolicy:           ProtoToComputeUrlMapDefaultRouteActionCorsPolicy(p.GetCorsPolicy()),
		FaultInjectionPolicy: ProtoToComputeUrlMapDefaultRouteActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
	}
	for _, r := range p.GetWeightedBackendService() {
		obj.WeightedBackendService = append(obj.WeightedBackendService, *ProtoToComputeUrlMapDefaultRouteActionWeightedBackendService(r))
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionWeightedBackendService converts a UrlMapDefaultRouteActionWeightedBackendService resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionWeightedBackendService(p *computepb.ComputeUrlMapDefaultRouteActionWeightedBackendService) *compute.UrlMapDefaultRouteActionWeightedBackendService {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionWeightedBackendService{
		BackendService: dcl.StringOrNil(p.BackendService),
		Weight:         dcl.Int64OrNil(p.Weight),
		HeaderAction:   ProtoToComputeUrlMapHeaderAction(p.GetHeaderAction()),
	}
	return obj
}

// ProtoToUrlMapHeaderAction converts a UrlMapHeaderAction resource from its proto representation.
func ProtoToComputeUrlMapHeaderAction(p *computepb.ComputeUrlMapHeaderAction) *compute.UrlMapHeaderAction {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapHeaderAction{}
	for _, r := range p.GetRequestHeadersToRemove() {
		obj.RequestHeadersToRemove = append(obj.RequestHeadersToRemove, r)
	}
	for _, r := range p.GetRequestHeadersToAdd() {
		obj.RequestHeadersToAdd = append(obj.RequestHeadersToAdd, *ProtoToComputeUrlMapHeaderActionRequestHeadersToAdd(r))
	}
	for _, r := range p.GetResponseHeadersToRemove() {
		obj.ResponseHeadersToRemove = append(obj.ResponseHeadersToRemove, r)
	}
	for _, r := range p.GetResponseHeadersToAdd() {
		obj.ResponseHeadersToAdd = append(obj.ResponseHeadersToAdd, *ProtoToComputeUrlMapHeaderActionResponseHeadersToAdd(r))
	}
	return obj
}

// ProtoToUrlMapHeaderActionRequestHeadersToAdd converts a UrlMapHeaderActionRequestHeadersToAdd resource from its proto representation.
func ProtoToComputeUrlMapHeaderActionRequestHeadersToAdd(p *computepb.ComputeUrlMapHeaderActionRequestHeadersToAdd) *compute.UrlMapHeaderActionRequestHeadersToAdd {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapHeaderActionRequestHeadersToAdd{
		HeaderName:  dcl.StringOrNil(p.HeaderName),
		HeaderValue: dcl.StringOrNil(p.HeaderValue),
		Replace:     dcl.Bool(p.Replace),
	}
	return obj
}

// ProtoToUrlMapHeaderActionResponseHeadersToAdd converts a UrlMapHeaderActionResponseHeadersToAdd resource from its proto representation.
func ProtoToComputeUrlMapHeaderActionResponseHeadersToAdd(p *computepb.ComputeUrlMapHeaderActionResponseHeadersToAdd) *compute.UrlMapHeaderActionResponseHeadersToAdd {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapHeaderActionResponseHeadersToAdd{
		HeaderName:  dcl.StringOrNil(p.HeaderName),
		HeaderValue: dcl.StringOrNil(p.HeaderValue),
		Replace:     dcl.Bool(p.Replace),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionUrlRewrite converts a UrlMapDefaultRouteActionUrlRewrite resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionUrlRewrite(p *computepb.ComputeUrlMapDefaultRouteActionUrlRewrite) *compute.UrlMapDefaultRouteActionUrlRewrite {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.StringOrNil(p.PathPrefixRewrite),
		HostRewrite:       dcl.StringOrNil(p.HostRewrite),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionTimeout converts a UrlMapDefaultRouteActionTimeout resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionTimeout(p *computepb.ComputeUrlMapDefaultRouteActionTimeout) *compute.UrlMapDefaultRouteActionTimeout {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionRetryPolicy converts a UrlMapDefaultRouteActionRetryPolicy resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionRetryPolicy(p *computepb.ComputeUrlMapDefaultRouteActionRetryPolicy) *compute.UrlMapDefaultRouteActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionRetryPolicy{
		NumRetries:    dcl.Int64OrNil(p.NumRetries),
		PerTryTimeout: ProtoToComputeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(p.GetPerTryTimeout()),
	}
	for _, r := range p.GetRetryCondition() {
		obj.RetryCondition = append(obj.RetryCondition, r)
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionRetryPolicyPerTryTimeout converts a UrlMapDefaultRouteActionRetryPolicyPerTryTimeout resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout(p *computepb.ComputeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout) *compute.UrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionRequestMirrorPolicy converts a UrlMapDefaultRouteActionRequestMirrorPolicy resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionRequestMirrorPolicy(p *computepb.ComputeUrlMapDefaultRouteActionRequestMirrorPolicy) *compute.UrlMapDefaultRouteActionRequestMirrorPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionRequestMirrorPolicy{
		BackendService: dcl.StringOrNil(p.BackendService),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionCorsPolicy converts a UrlMapDefaultRouteActionCorsPolicy resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionCorsPolicy(p *computepb.ComputeUrlMapDefaultRouteActionCorsPolicy) *compute.UrlMapDefaultRouteActionCorsPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionCorsPolicy{
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
func ProtoToComputeUrlMapDefaultRouteActionFaultInjectionPolicy(p *computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicy) *compute.UrlMapDefaultRouteActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionFaultInjectionPolicy{
		Delay: ProtoToComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToComputeUrlMapDefaultRouteActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionFaultInjectionPolicyDelay converts a UrlMapDefaultRouteActionFaultInjectionPolicyDelay resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelay(p *computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelay) *compute.UrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ProtoToComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(p.GetFixedDelay()),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay converts a UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay(p *computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) *compute.UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapDefaultRouteActionFaultInjectionPolicyAbort converts a UrlMapDefaultRouteActionFaultInjectionPolicyAbort resource from its proto representation.
func ProtoToComputeUrlMapDefaultRouteActionFaultInjectionPolicyAbort(p *computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyAbort) *compute.UrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.HttpStatus),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapDefaultUrlRedirect converts a UrlMapDefaultUrlRedirect resource from its proto representation.
func ProtoToComputeUrlMapDefaultUrlRedirect(p *computepb.ComputeUrlMapDefaultUrlRedirect) *compute.UrlMapDefaultUrlRedirect {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapDefaultUrlRedirect{
		HostRedirect:         dcl.StringOrNil(p.HostRedirect),
		PathRedirect:         dcl.StringOrNil(p.PathRedirect),
		PrefixRedirect:       dcl.StringOrNil(p.PrefixRedirect),
		RedirectResponseCode: ProtoToComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(p.GetRedirectResponseCode()),
		HttpsRedirect:        dcl.Bool(p.HttpsRedirect),
		StripQuery:           dcl.Bool(p.StripQuery),
	}
	return obj
}

// ProtoToUrlMapHostRule converts a UrlMapHostRule resource from its proto representation.
func ProtoToComputeUrlMapHostRule(p *computepb.ComputeUrlMapHostRule) *compute.UrlMapHostRule {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapHostRule{
		Description: dcl.StringOrNil(p.Description),
		PathMatcher: dcl.StringOrNil(p.PathMatcher),
	}
	for _, r := range p.GetHost() {
		obj.Host = append(obj.Host, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcher converts a UrlMapPathMatcher resource from its proto representation.
func ProtoToComputeUrlMapPathMatcher(p *computepb.ComputeUrlMapPathMatcher) *compute.UrlMapPathMatcher {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcher{
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		DefaultService:     dcl.StringOrNil(p.DefaultService),
		DefaultRouteAction: ProtoToComputeUrlMapDefaultRouteAction(p.GetDefaultRouteAction()),
		DefaultUrlRedirect: ProtoToComputeUrlMapPathMatcherDefaultUrlRedirect(p.GetDefaultUrlRedirect()),
		HeaderAction:       ProtoToComputeUrlMapHeaderAction(p.GetHeaderAction()),
	}
	for _, r := range p.GetPathRule() {
		obj.PathRule = append(obj.PathRule, *ProtoToComputeUrlMapPathMatcherPathRule(r))
	}
	for _, r := range p.GetRouteRule() {
		obj.RouteRule = append(obj.RouteRule, *ProtoToComputeUrlMapPathMatcherRouteRule(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherDefaultUrlRedirect converts a UrlMapPathMatcherDefaultUrlRedirect resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherDefaultUrlRedirect(p *computepb.ComputeUrlMapPathMatcherDefaultUrlRedirect) *compute.UrlMapPathMatcherDefaultUrlRedirect {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherDefaultUrlRedirect{
		HostRedirect:         dcl.StringOrNil(p.HostRedirect),
		PathRedirect:         dcl.StringOrNil(p.PathRedirect),
		PrefixRedirect:       dcl.StringOrNil(p.PrefixRedirect),
		RedirectResponseCode: ProtoToComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(p.GetRedirectResponseCode()),
		HttpsRedirect:        dcl.Bool(p.HttpsRedirect),
		StripQuery:           dcl.Bool(p.StripQuery),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRule converts a UrlMapPathMatcherPathRule resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRule(p *computepb.ComputeUrlMapPathMatcherPathRule) *compute.UrlMapPathMatcherPathRule {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRule{
		BackendService: dcl.StringOrNil(p.BackendService),
		RouteAction:    ProtoToComputeUrlMapPathMatcherPathRuleRouteAction(p.GetRouteAction()),
		UrlRedirect:    ProtoToComputeUrlMapPathMatcherPathRuleUrlRedirect(p.GetUrlRedirect()),
	}
	for _, r := range p.GetPath() {
		obj.Path = append(obj.Path, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteAction converts a UrlMapPathMatcherPathRuleRouteAction resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteAction(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteAction) *compute.UrlMapPathMatcherPathRuleRouteAction {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteAction{
		UrlRewrite:           ProtoToComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite(p.GetUrlRewrite()),
		Timeout:              ProtoToComputeUrlMapPathMatcherPathRuleRouteActionTimeout(p.GetTimeout()),
		RetryPolicy:          ProtoToComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy(p.GetRetryPolicy()),
		RequestMirrorPolicy:  ProtoToComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(p.GetRequestMirrorPolicy()),
		CorsPolicy:           ProtoToComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy(p.GetCorsPolicy()),
		FaultInjectionPolicy: ProtoToComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
	}
	for _, r := range p.GetWeightedBackendService() {
		obj.WeightedBackendService = append(obj.WeightedBackendService, *ProtoToComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionWeightedBackendService converts a UrlMapPathMatcherPathRuleRouteActionWeightedBackendService resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendService(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendService) *compute.UrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionWeightedBackendService{
		BackendService: dcl.StringOrNil(p.BackendService),
		Weight:         dcl.Int64OrNil(p.Weight),
		HeaderAction:   ProtoToComputeUrlMapHeaderAction(p.GetHeaderAction()),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionUrlRewrite converts a UrlMapPathMatcherPathRuleRouteActionUrlRewrite resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite) *compute.UrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.StringOrNil(p.PathPrefixRewrite),
		HostRewrite:       dcl.StringOrNil(p.HostRewrite),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionTimeout converts a UrlMapPathMatcherPathRuleRouteActionTimeout resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionTimeout(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionTimeout) *compute.UrlMapPathMatcherPathRuleRouteActionTimeout {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionRetryPolicy converts a UrlMapPathMatcherPathRuleRouteActionRetryPolicy resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy) *compute.UrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionRetryPolicy{
		NumRetries:    dcl.Int64OrNil(p.NumRetries),
		PerTryTimeout: ProtoToComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(p.GetPerTryTimeout()),
	}
	for _, r := range p.GetRetryCondition() {
		obj.RetryCondition = append(obj.RetryCondition, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout converts a UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) *compute.UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy converts a UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) *compute.UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{
		BackendService: dcl.StringOrNil(p.BackendService),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionCorsPolicy converts a UrlMapPathMatcherPathRuleRouteActionCorsPolicy resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy) *compute.UrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionCorsPolicy{
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
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) *compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{
		Delay: ProtoToComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) *compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ProtoToComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(p.GetFixedDelay()),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort(p *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) *compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.HttpStatus),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapPathMatcherPathRuleUrlRedirect converts a UrlMapPathMatcherPathRuleUrlRedirect resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherPathRuleUrlRedirect(p *computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirect) *compute.UrlMapPathMatcherPathRuleUrlRedirect {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherPathRuleUrlRedirect{
		HostRedirect:         dcl.StringOrNil(p.HostRedirect),
		PathRedirect:         dcl.StringOrNil(p.PathRedirect),
		PrefixRedirect:       dcl.StringOrNil(p.PrefixRedirect),
		RedirectResponseCode: ProtoToComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(p.GetRedirectResponseCode()),
		HttpsRedirect:        dcl.Bool(p.HttpsRedirect),
		StripQuery:           dcl.Bool(p.StripQuery),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRule converts a UrlMapPathMatcherRouteRule resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRule(p *computepb.ComputeUrlMapPathMatcherRouteRule) *compute.UrlMapPathMatcherRouteRule {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRule{
		Priority:       dcl.Int64OrNil(p.Priority),
		Description:    dcl.StringOrNil(p.Description),
		BackendService: dcl.StringOrNil(p.BackendService),
		RouteAction:    ProtoToComputeUrlMapPathMatcherRouteRuleRouteAction(p.GetRouteAction()),
		UrlRedirect:    ProtoToComputeUrlMapPathMatcherRouteRuleUrlRedirect(p.GetUrlRedirect()),
		HeaderAction:   ProtoToComputeUrlMapHeaderAction(p.GetHeaderAction()),
	}
	for _, r := range p.GetMatchRule() {
		obj.MatchRule = append(obj.MatchRule, *ProtoToComputeUrlMapPathMatcherRouteRuleMatchRule(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRule converts a UrlMapPathMatcherRouteRuleMatchRule resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleMatchRule(p *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRule) *compute.UrlMapPathMatcherRouteRuleMatchRule {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleMatchRule{
		PrefixMatch:   dcl.StringOrNil(p.PrefixMatch),
		FullPathMatch: dcl.StringOrNil(p.FullPathMatch),
		RegexMatch:    dcl.StringOrNil(p.RegexMatch),
		IgnoreCase:    dcl.Bool(p.IgnoreCase),
	}
	for _, r := range p.GetHeaderMatch() {
		obj.HeaderMatch = append(obj.HeaderMatch, *ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(r))
	}
	for _, r := range p.GetQueryParameterMatch() {
		obj.QueryParameterMatch = append(obj.QueryParameterMatch, *ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(r))
	}
	for _, r := range p.GetMetadataFilter() {
		obj.MetadataFilter = append(obj.MetadataFilter, *ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch converts a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch(p *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) *compute.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{
		HeaderName:   dcl.StringOrNil(p.HeaderName),
		ExactMatch:   dcl.StringOrNil(p.ExactMatch),
		RegexMatch:   dcl.StringOrNil(p.RegexMatch),
		RangeMatch:   ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(p.GetRangeMatch()),
		PresentMatch: dcl.Bool(p.PresentMatch),
		PrefixMatch:  dcl.StringOrNil(p.PrefixMatch),
		SuffixMatch:  dcl.StringOrNil(p.SuffixMatch),
		InvertMatch:  dcl.Bool(p.InvertMatch),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch converts a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch(p *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) *compute.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{
		RangeStart: dcl.Int64OrNil(p.RangeStart),
		RangeEnd:   dcl.Int64OrNil(p.RangeEnd),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch converts a UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch(p *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) *compute.UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{
		Name:         dcl.StringOrNil(p.Name),
		PresentMatch: dcl.Bool(p.PresentMatch),
		ExactMatch:   dcl.StringOrNil(p.ExactMatch),
		RegexMatch:   dcl.StringOrNil(p.RegexMatch),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter(p *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) *compute.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{
		FilterMatchCriteria: ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(p.GetFilterMatchCriteria()),
	}
	for _, r := range p.GetFilterLabel() {
		obj.FilterLabel = append(obj.FilterLabel, *ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel(p *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) *compute.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{
		Name:  dcl.StringOrNil(p.Name),
		Value: dcl.StringOrNil(p.Value),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteAction converts a UrlMapPathMatcherRouteRuleRouteAction resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteAction(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteAction) *compute.UrlMapPathMatcherRouteRuleRouteAction {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteAction{
		UrlRewrite:           ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(p.GetUrlRewrite()),
		Timeout:              ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionTimeout(p.GetTimeout()),
		RetryPolicy:          ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(p.GetRetryPolicy()),
		RequestMirrorPolicy:  ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(p.GetRequestMirrorPolicy()),
		CorsPolicy:           ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(p.GetCorsPolicy()),
		FaultInjectionPolicy: ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
	}
	for _, r := range p.GetWeightedBackendService() {
		obj.WeightedBackendService = append(obj.WeightedBackendService, *ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(r))
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService converts a UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) *compute.UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{
		BackendService: dcl.StringOrNil(p.BackendService),
		Weight:         dcl.Int64OrNil(p.Weight),
		HeaderAction:   ProtoToComputeUrlMapHeaderAction(p.GetHeaderAction()),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionUrlRewrite converts a UrlMapPathMatcherRouteRuleRouteActionUrlRewrite resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite) *compute.UrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.StringOrNil(p.PathPrefixRewrite),
		HostRewrite:       dcl.StringOrNil(p.HostRewrite),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionTimeout converts a UrlMapPathMatcherRouteRuleRouteActionTimeout resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionTimeout(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionTimeout) *compute.UrlMapPathMatcherRouteRuleRouteActionTimeout {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionRetryPolicy converts a UrlMapPathMatcherRouteRuleRouteActionRetryPolicy resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy) *compute.UrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionRetryPolicy{
		NumRetries:    dcl.Int64OrNil(p.NumRetries),
		PerTryTimeout: ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(p.GetPerTryTimeout()),
	}
	for _, r := range p.GetRetryCondition() {
		obj.RetryCondition = append(obj.RetryCondition, r)
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout converts a UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) *compute.UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy converts a UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) *compute.UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{
		BackendService: dcl.StringOrNil(p.BackendService),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionCorsPolicy converts a UrlMapPathMatcherRouteRuleRouteActionCorsPolicy resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy) *compute.UrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionCorsPolicy{
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
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) *compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{
		Delay: ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) *compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(p.GetFixedDelay()),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.Int64OrNil(p.Seconds),
		Nanos:   dcl.Int64OrNil(p.Nanos),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort(p *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) *compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.HttpStatus),
		Percentage: dcl.Float64OrNil(p.Percentage),
	}
	return obj
}

// ProtoToUrlMapPathMatcherRouteRuleUrlRedirect converts a UrlMapPathMatcherRouteRuleUrlRedirect resource from its proto representation.
func ProtoToComputeUrlMapPathMatcherRouteRuleUrlRedirect(p *computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirect) *compute.UrlMapPathMatcherRouteRuleUrlRedirect {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapPathMatcherRouteRuleUrlRedirect{
		HostRedirect:         dcl.StringOrNil(p.HostRedirect),
		PathRedirect:         dcl.StringOrNil(p.PathRedirect),
		PrefixRedirect:       dcl.StringOrNil(p.PrefixRedirect),
		RedirectResponseCode: ProtoToComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(p.GetRedirectResponseCode()),
		HttpsRedirect:        dcl.Bool(p.HttpsRedirect),
		StripQuery:           dcl.Bool(p.StripQuery),
	}
	return obj
}

// ProtoToUrlMapTest converts a UrlMapTest resource from its proto representation.
func ProtoToComputeUrlMapTest(p *computepb.ComputeUrlMapTest) *compute.UrlMapTest {
	if p == nil {
		return nil
	}
	obj := &compute.UrlMapTest{
		Description:            dcl.StringOrNil(p.Description),
		Host:                   dcl.StringOrNil(p.Host),
		Path:                   dcl.StringOrNil(p.Path),
		ExpectedBackendService: dcl.StringOrNil(p.ExpectedBackendService),
	}
	return obj
}

// ProtoToUrlMap converts a UrlMap resource from its proto representation.
func ProtoToUrlMap(p *computepb.ComputeUrlMap) *compute.UrlMap {
	obj := &compute.UrlMap{
		DefaultRouteAction: ProtoToComputeUrlMapDefaultRouteAction(p.GetDefaultRouteAction()),
		DefaultService:     dcl.StringOrNil(p.DefaultService),
		DefaultUrlRedirect: ProtoToComputeUrlMapDefaultUrlRedirect(p.GetDefaultUrlRedirect()),
		Description:        dcl.StringOrNil(p.Description),
		SelfLink:           dcl.StringOrNil(p.SelfLink),
		HeaderAction:       ProtoToComputeUrlMapHeaderAction(p.GetHeaderAction()),
		Name:               dcl.StringOrNil(p.Name),
		Region:             dcl.StringOrNil(p.Region),
		Project:            dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetHostRule() {
		obj.HostRule = append(obj.HostRule, *ProtoToComputeUrlMapHostRule(r))
	}
	for _, r := range p.GetPathMatcher() {
		obj.PathMatcher = append(obj.PathMatcher, *ProtoToComputeUrlMapPathMatcher(r))
	}
	for _, r := range p.GetTest() {
		obj.Test = append(obj.Test, *ProtoToComputeUrlMapTest(r))
	}
	return obj
}

// UrlMapDefaultUrlRedirectRedirectResponseCodeEnumToProto converts a UrlMapDefaultUrlRedirectRedirectResponseCodeEnum enum to its proto representation.
func ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnumToProto(e *compute.UrlMapDefaultUrlRedirectRedirectResponseCodeEnum) computepb.ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum {
	if e == nil {
		return computepb.ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(0)
	}
	if v, ok := computepb.ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum_value["UrlMapDefaultUrlRedirectRedirectResponseCodeEnum"+string(*e)]; ok {
		return computepb.ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(v)
	}
	return computepb.ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnum(0)
}

// UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumToProto converts a UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum enum to its proto representation.
func ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumToProto(e *compute.UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum) computepb.ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum {
	if e == nil {
		return computepb.ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(0)
	}
	if v, ok := computepb.ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum_value["UrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum"+string(*e)]; ok {
		return computepb.ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(v)
	}
	return computepb.ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnum(0)
}

// UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumToProto converts a UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum enum to its proto representation.
func ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumToProto(e *compute.UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum) computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum {
	if e == nil {
		return computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(0)
	}
	if v, ok := computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum_value["UrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum"+string(*e)]; ok {
		return computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(v)
	}
	return computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnum(0)
}

// UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumToProto converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum enum to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumToProto(e *compute.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum) computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum {
	if e == nil {
		return computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(0)
	}
	if v, ok := computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum_value["UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum"+string(*e)]; ok {
		return computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(v)
	}
	return computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnum(0)
}

// UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumToProto converts a UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum enum to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumToProto(e *compute.UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum) computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum {
	if e == nil {
		return computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(0)
	}
	if v, ok := computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum_value["UrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum"+string(*e)]; ok {
		return computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(v)
	}
	return computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnum(0)
}

// UrlMapDefaultRouteActionToProto converts a UrlMapDefaultRouteAction resource to its proto representation.
func ComputeUrlMapDefaultRouteActionToProto(o *compute.UrlMapDefaultRouteAction) *computepb.ComputeUrlMapDefaultRouteAction {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteAction{
		UrlRewrite:           ComputeUrlMapDefaultRouteActionUrlRewriteToProto(o.UrlRewrite),
		Timeout:              ComputeUrlMapDefaultRouteActionTimeoutToProto(o.Timeout),
		RetryPolicy:          ComputeUrlMapDefaultRouteActionRetryPolicyToProto(o.RetryPolicy),
		RequestMirrorPolicy:  ComputeUrlMapDefaultRouteActionRequestMirrorPolicyToProto(o.RequestMirrorPolicy),
		CorsPolicy:           ComputeUrlMapDefaultRouteActionCorsPolicyToProto(o.CorsPolicy),
		FaultInjectionPolicy: ComputeUrlMapDefaultRouteActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy),
	}
	for _, r := range o.WeightedBackendService {
		p.WeightedBackendService = append(p.WeightedBackendService, ComputeUrlMapDefaultRouteActionWeightedBackendServiceToProto(&r))
	}
	return p
}

// UrlMapDefaultRouteActionWeightedBackendServiceToProto converts a UrlMapDefaultRouteActionWeightedBackendService resource to its proto representation.
func ComputeUrlMapDefaultRouteActionWeightedBackendServiceToProto(o *compute.UrlMapDefaultRouteActionWeightedBackendService) *computepb.ComputeUrlMapDefaultRouteActionWeightedBackendService {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionWeightedBackendService{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		Weight:         dcl.ValueOrEmptyInt64(o.Weight),
		HeaderAction:   ComputeUrlMapHeaderActionToProto(o.HeaderAction),
	}
	return p
}

// UrlMapHeaderActionToProto converts a UrlMapHeaderAction resource to its proto representation.
func ComputeUrlMapHeaderActionToProto(o *compute.UrlMapHeaderAction) *computepb.ComputeUrlMapHeaderAction {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapHeaderAction{}
	for _, r := range o.RequestHeadersToRemove {
		p.RequestHeadersToRemove = append(p.RequestHeadersToRemove, r)
	}
	for _, r := range o.RequestHeadersToAdd {
		p.RequestHeadersToAdd = append(p.RequestHeadersToAdd, ComputeUrlMapHeaderActionRequestHeadersToAddToProto(&r))
	}
	for _, r := range o.ResponseHeadersToRemove {
		p.ResponseHeadersToRemove = append(p.ResponseHeadersToRemove, r)
	}
	for _, r := range o.ResponseHeadersToAdd {
		p.ResponseHeadersToAdd = append(p.ResponseHeadersToAdd, ComputeUrlMapHeaderActionResponseHeadersToAddToProto(&r))
	}
	return p
}

// UrlMapHeaderActionRequestHeadersToAddToProto converts a UrlMapHeaderActionRequestHeadersToAdd resource to its proto representation.
func ComputeUrlMapHeaderActionRequestHeadersToAddToProto(o *compute.UrlMapHeaderActionRequestHeadersToAdd) *computepb.ComputeUrlMapHeaderActionRequestHeadersToAdd {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapHeaderActionRequestHeadersToAdd{
		HeaderName:  dcl.ValueOrEmptyString(o.HeaderName),
		HeaderValue: dcl.ValueOrEmptyString(o.HeaderValue),
		Replace:     dcl.ValueOrEmptyBool(o.Replace),
	}
	return p
}

// UrlMapHeaderActionResponseHeadersToAddToProto converts a UrlMapHeaderActionResponseHeadersToAdd resource to its proto representation.
func ComputeUrlMapHeaderActionResponseHeadersToAddToProto(o *compute.UrlMapHeaderActionResponseHeadersToAdd) *computepb.ComputeUrlMapHeaderActionResponseHeadersToAdd {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapHeaderActionResponseHeadersToAdd{
		HeaderName:  dcl.ValueOrEmptyString(o.HeaderName),
		HeaderValue: dcl.ValueOrEmptyString(o.HeaderValue),
		Replace:     dcl.ValueOrEmptyBool(o.Replace),
	}
	return p
}

// UrlMapDefaultRouteActionUrlRewriteToProto converts a UrlMapDefaultRouteActionUrlRewrite resource to its proto representation.
func ComputeUrlMapDefaultRouteActionUrlRewriteToProto(o *compute.UrlMapDefaultRouteActionUrlRewrite) *computepb.ComputeUrlMapDefaultRouteActionUrlRewrite {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.ValueOrEmptyString(o.PathPrefixRewrite),
		HostRewrite:       dcl.ValueOrEmptyString(o.HostRewrite),
	}
	return p
}

// UrlMapDefaultRouteActionTimeoutToProto converts a UrlMapDefaultRouteActionTimeout resource to its proto representation.
func ComputeUrlMapDefaultRouteActionTimeoutToProto(o *compute.UrlMapDefaultRouteActionTimeout) *computepb.ComputeUrlMapDefaultRouteActionTimeout {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapDefaultRouteActionRetryPolicyToProto converts a UrlMapDefaultRouteActionRetryPolicy resource to its proto representation.
func ComputeUrlMapDefaultRouteActionRetryPolicyToProto(o *compute.UrlMapDefaultRouteActionRetryPolicy) *computepb.ComputeUrlMapDefaultRouteActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionRetryPolicy{
		NumRetries:    dcl.ValueOrEmptyInt64(o.NumRetries),
		PerTryTimeout: ComputeUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutToProto(o.PerTryTimeout),
	}
	for _, r := range o.RetryCondition {
		p.RetryCondition = append(p.RetryCondition, r)
	}
	return p
}

// UrlMapDefaultRouteActionRetryPolicyPerTryTimeoutToProto converts a UrlMapDefaultRouteActionRetryPolicyPerTryTimeout resource to its proto representation.
func ComputeUrlMapDefaultRouteActionRetryPolicyPerTryTimeoutToProto(o *compute.UrlMapDefaultRouteActionRetryPolicyPerTryTimeout) *computepb.ComputeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapDefaultRouteActionRequestMirrorPolicyToProto converts a UrlMapDefaultRouteActionRequestMirrorPolicy resource to its proto representation.
func ComputeUrlMapDefaultRouteActionRequestMirrorPolicyToProto(o *compute.UrlMapDefaultRouteActionRequestMirrorPolicy) *computepb.ComputeUrlMapDefaultRouteActionRequestMirrorPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionRequestMirrorPolicy{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
	}
	return p
}

// UrlMapDefaultRouteActionCorsPolicyToProto converts a UrlMapDefaultRouteActionCorsPolicy resource to its proto representation.
func ComputeUrlMapDefaultRouteActionCorsPolicyToProto(o *compute.UrlMapDefaultRouteActionCorsPolicy) *computepb.ComputeUrlMapDefaultRouteActionCorsPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionCorsPolicy{
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
func ComputeUrlMapDefaultRouteActionFaultInjectionPolicyToProto(o *compute.UrlMapDefaultRouteActionFaultInjectionPolicy) *computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicy{
		Delay: ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayToProto(o.Delay),
		Abort: ComputeUrlMapDefaultRouteActionFaultInjectionPolicyAbortToProto(o.Abort),
	}
	return p
}

// UrlMapDefaultRouteActionFaultInjectionPolicyDelayToProto converts a UrlMapDefaultRouteActionFaultInjectionPolicyDelay resource to its proto representation.
func ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayToProto(o *compute.UrlMapDefaultRouteActionFaultInjectionPolicyDelay) *computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o.FixedDelay),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayToProto converts a UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay resource to its proto representation.
func ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o *compute.UrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay) *computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapDefaultRouteActionFaultInjectionPolicyAbortToProto converts a UrlMapDefaultRouteActionFaultInjectionPolicyAbort resource to its proto representation.
func ComputeUrlMapDefaultRouteActionFaultInjectionPolicyAbortToProto(o *compute.UrlMapDefaultRouteActionFaultInjectionPolicyAbort) *computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.ValueOrEmptyInt64(o.HttpStatus),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapDefaultUrlRedirectToProto converts a UrlMapDefaultUrlRedirect resource to its proto representation.
func ComputeUrlMapDefaultUrlRedirectToProto(o *compute.UrlMapDefaultUrlRedirect) *computepb.ComputeUrlMapDefaultUrlRedirect {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapDefaultUrlRedirect{
		HostRedirect:         dcl.ValueOrEmptyString(o.HostRedirect),
		PathRedirect:         dcl.ValueOrEmptyString(o.PathRedirect),
		PrefixRedirect:       dcl.ValueOrEmptyString(o.PrefixRedirect),
		RedirectResponseCode: ComputeUrlMapDefaultUrlRedirectRedirectResponseCodeEnumToProto(o.RedirectResponseCode),
		HttpsRedirect:        dcl.ValueOrEmptyBool(o.HttpsRedirect),
		StripQuery:           dcl.ValueOrEmptyBool(o.StripQuery),
	}
	return p
}

// UrlMapHostRuleToProto converts a UrlMapHostRule resource to its proto representation.
func ComputeUrlMapHostRuleToProto(o *compute.UrlMapHostRule) *computepb.ComputeUrlMapHostRule {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapHostRule{
		Description: dcl.ValueOrEmptyString(o.Description),
		PathMatcher: dcl.ValueOrEmptyString(o.PathMatcher),
	}
	for _, r := range o.Host {
		p.Host = append(p.Host, r)
	}
	return p
}

// UrlMapPathMatcherToProto converts a UrlMapPathMatcher resource to its proto representation.
func ComputeUrlMapPathMatcherToProto(o *compute.UrlMapPathMatcher) *computepb.ComputeUrlMapPathMatcher {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcher{
		Name:               dcl.ValueOrEmptyString(o.Name),
		Description:        dcl.ValueOrEmptyString(o.Description),
		DefaultService:     dcl.ValueOrEmptyString(o.DefaultService),
		DefaultRouteAction: ComputeUrlMapDefaultRouteActionToProto(o.DefaultRouteAction),
		DefaultUrlRedirect: ComputeUrlMapPathMatcherDefaultUrlRedirectToProto(o.DefaultUrlRedirect),
		HeaderAction:       ComputeUrlMapHeaderActionToProto(o.HeaderAction),
	}
	for _, r := range o.PathRule {
		p.PathRule = append(p.PathRule, ComputeUrlMapPathMatcherPathRuleToProto(&r))
	}
	for _, r := range o.RouteRule {
		p.RouteRule = append(p.RouteRule, ComputeUrlMapPathMatcherRouteRuleToProto(&r))
	}
	return p
}

// UrlMapPathMatcherDefaultUrlRedirectToProto converts a UrlMapPathMatcherDefaultUrlRedirect resource to its proto representation.
func ComputeUrlMapPathMatcherDefaultUrlRedirectToProto(o *compute.UrlMapPathMatcherDefaultUrlRedirect) *computepb.ComputeUrlMapPathMatcherDefaultUrlRedirect {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherDefaultUrlRedirect{
		HostRedirect:         dcl.ValueOrEmptyString(o.HostRedirect),
		PathRedirect:         dcl.ValueOrEmptyString(o.PathRedirect),
		PrefixRedirect:       dcl.ValueOrEmptyString(o.PrefixRedirect),
		RedirectResponseCode: ComputeUrlMapPathMatcherDefaultUrlRedirectRedirectResponseCodeEnumToProto(o.RedirectResponseCode),
		HttpsRedirect:        dcl.ValueOrEmptyBool(o.HttpsRedirect),
		StripQuery:           dcl.ValueOrEmptyBool(o.StripQuery),
	}
	return p
}

// UrlMapPathMatcherPathRuleToProto converts a UrlMapPathMatcherPathRule resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleToProto(o *compute.UrlMapPathMatcherPathRule) *computepb.ComputeUrlMapPathMatcherPathRule {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRule{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		RouteAction:    ComputeUrlMapPathMatcherPathRuleRouteActionToProto(o.RouteAction),
		UrlRedirect:    ComputeUrlMapPathMatcherPathRuleUrlRedirectToProto(o.UrlRedirect),
	}
	for _, r := range o.Path {
		p.Path = append(p.Path, r)
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionToProto converts a UrlMapPathMatcherPathRuleRouteAction resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionToProto(o *compute.UrlMapPathMatcherPathRuleRouteAction) *computepb.ComputeUrlMapPathMatcherPathRuleRouteAction {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteAction{
		UrlRewrite:           ComputeUrlMapPathMatcherPathRuleRouteActionUrlRewriteToProto(o.UrlRewrite),
		Timeout:              ComputeUrlMapPathMatcherPathRuleRouteActionTimeoutToProto(o.Timeout),
		RetryPolicy:          ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyToProto(o.RetryPolicy),
		RequestMirrorPolicy:  ComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyToProto(o.RequestMirrorPolicy),
		CorsPolicy:           ComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicyToProto(o.CorsPolicy),
		FaultInjectionPolicy: ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy),
	}
	for _, r := range o.WeightedBackendService {
		p.WeightedBackendService = append(p.WeightedBackendService, ComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceToProto(&r))
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceToProto converts a UrlMapPathMatcherPathRuleRouteActionWeightedBackendService resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionWeightedBackendService) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendService {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionWeightedBackendService{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		Weight:         dcl.ValueOrEmptyInt64(o.Weight),
		HeaderAction:   ComputeUrlMapHeaderActionToProto(o.HeaderAction),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionUrlRewriteToProto converts a UrlMapPathMatcherPathRuleRouteActionUrlRewrite resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionUrlRewriteToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionUrlRewrite) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.ValueOrEmptyString(o.PathPrefixRewrite),
		HostRewrite:       dcl.ValueOrEmptyString(o.HostRewrite),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionTimeoutToProto converts a UrlMapPathMatcherPathRuleRouteActionTimeout resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionTimeoutToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionTimeout) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionTimeout {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionRetryPolicyToProto converts a UrlMapPathMatcherPathRuleRouteActionRetryPolicy resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionRetryPolicy) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicy{
		NumRetries:    dcl.ValueOrEmptyInt64(o.NumRetries),
		PerTryTimeout: ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutToProto(o.PerTryTimeout),
	}
	for _, r := range o.RetryCondition {
		p.RetryCondition = append(p.RetryCondition, r)
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutToProto converts a UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeoutToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyToProto converts a UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionCorsPolicyToProto converts a UrlMapPathMatcherPathRuleRouteActionCorsPolicy resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicyToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionCorsPolicy) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionCorsPolicy{
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
func ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy{
		Delay: ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayToProto(o.Delay),
		Abort: ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortToProto(o.Abort),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayToProto converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o.FixedDelay),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortToProto converts a UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbortToProto(o *compute.UrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort) *computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.ValueOrEmptyInt64(o.HttpStatus),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapPathMatcherPathRuleUrlRedirectToProto converts a UrlMapPathMatcherPathRuleUrlRedirect resource to its proto representation.
func ComputeUrlMapPathMatcherPathRuleUrlRedirectToProto(o *compute.UrlMapPathMatcherPathRuleUrlRedirect) *computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirect {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherPathRuleUrlRedirect{
		HostRedirect:         dcl.ValueOrEmptyString(o.HostRedirect),
		PathRedirect:         dcl.ValueOrEmptyString(o.PathRedirect),
		PrefixRedirect:       dcl.ValueOrEmptyString(o.PrefixRedirect),
		RedirectResponseCode: ComputeUrlMapPathMatcherPathRuleUrlRedirectRedirectResponseCodeEnumToProto(o.RedirectResponseCode),
		HttpsRedirect:        dcl.ValueOrEmptyBool(o.HttpsRedirect),
		StripQuery:           dcl.ValueOrEmptyBool(o.StripQuery),
	}
	return p
}

// UrlMapPathMatcherRouteRuleToProto converts a UrlMapPathMatcherRouteRule resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleToProto(o *compute.UrlMapPathMatcherRouteRule) *computepb.ComputeUrlMapPathMatcherRouteRule {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRule{
		Priority:       dcl.ValueOrEmptyInt64(o.Priority),
		Description:    dcl.ValueOrEmptyString(o.Description),
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		RouteAction:    ComputeUrlMapPathMatcherRouteRuleRouteActionToProto(o.RouteAction),
		UrlRedirect:    ComputeUrlMapPathMatcherRouteRuleUrlRedirectToProto(o.UrlRedirect),
		HeaderAction:   ComputeUrlMapHeaderActionToProto(o.HeaderAction),
	}
	for _, r := range o.MatchRule {
		p.MatchRule = append(p.MatchRule, ComputeUrlMapPathMatcherRouteRuleMatchRuleToProto(&r))
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleToProto converts a UrlMapPathMatcherRouteRuleMatchRule resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleMatchRuleToProto(o *compute.UrlMapPathMatcherRouteRuleMatchRule) *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRule {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleMatchRule{
		PrefixMatch:   dcl.ValueOrEmptyString(o.PrefixMatch),
		FullPathMatch: dcl.ValueOrEmptyString(o.FullPathMatch),
		RegexMatch:    dcl.ValueOrEmptyString(o.RegexMatch),
		IgnoreCase:    dcl.ValueOrEmptyBool(o.IgnoreCase),
	}
	for _, r := range o.HeaderMatch {
		p.HeaderMatch = append(p.HeaderMatch, ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchToProto(&r))
	}
	for _, r := range o.QueryParameterMatch {
		p.QueryParameterMatch = append(p.QueryParameterMatch, ComputeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchToProto(&r))
	}
	for _, r := range o.MetadataFilter {
		p.MetadataFilter = append(p.MetadataFilter, ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterToProto(&r))
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchToProto converts a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchToProto(o *compute.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatch) *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatch{
		HeaderName:   dcl.ValueOrEmptyString(o.HeaderName),
		ExactMatch:   dcl.ValueOrEmptyString(o.ExactMatch),
		RegexMatch:   dcl.ValueOrEmptyString(o.RegexMatch),
		RangeMatch:   ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchToProto(o.RangeMatch),
		PresentMatch: dcl.ValueOrEmptyBool(o.PresentMatch),
		PrefixMatch:  dcl.ValueOrEmptyString(o.PrefixMatch),
		SuffixMatch:  dcl.ValueOrEmptyString(o.SuffixMatch),
		InvertMatch:  dcl.ValueOrEmptyBool(o.InvertMatch),
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchToProto converts a UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchToProto(o *compute.UrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch) *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch{
		RangeStart: dcl.ValueOrEmptyInt64(o.RangeStart),
		RangeEnd:   dcl.ValueOrEmptyInt64(o.RangeEnd),
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchToProto converts a UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatchToProto(o *compute.UrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch) *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleQueryParameterMatch{
		Name:         dcl.ValueOrEmptyString(o.Name),
		PresentMatch: dcl.ValueOrEmptyBool(o.PresentMatch),
		ExactMatch:   dcl.ValueOrEmptyString(o.ExactMatch),
		RegexMatch:   dcl.ValueOrEmptyString(o.RegexMatch),
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterToProto converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterToProto(o *compute.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilter) *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilter{
		FilterMatchCriteria: ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterMatchCriteriaEnumToProto(o.FilterMatchCriteria),
	}
	for _, r := range o.FilterLabel {
		p.FilterLabel = append(p.FilterLabel, ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelToProto(&r))
	}
	return p
}

// UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelToProto converts a UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabelToProto(o *compute.UrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel) *computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel{
		Name:  dcl.ValueOrEmptyString(o.Name),
		Value: dcl.ValueOrEmptyString(o.Value),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionToProto converts a UrlMapPathMatcherRouteRuleRouteAction resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionToProto(o *compute.UrlMapPathMatcherRouteRuleRouteAction) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteAction {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteAction{
		UrlRewrite:           ComputeUrlMapPathMatcherRouteRuleRouteActionUrlRewriteToProto(o.UrlRewrite),
		Timeout:              ComputeUrlMapPathMatcherRouteRuleRouteActionTimeoutToProto(o.Timeout),
		RetryPolicy:          ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyToProto(o.RetryPolicy),
		RequestMirrorPolicy:  ComputeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyToProto(o.RequestMirrorPolicy),
		CorsPolicy:           ComputeUrlMapPathMatcherRouteRuleRouteActionCorsPolicyToProto(o.CorsPolicy),
		FaultInjectionPolicy: ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy),
	}
	for _, r := range o.WeightedBackendService {
		p.WeightedBackendService = append(p.WeightedBackendService, ComputeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceToProto(&r))
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceToProto converts a UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionWeightedBackendService) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionWeightedBackendService{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
		Weight:         dcl.ValueOrEmptyInt64(o.Weight),
		HeaderAction:   ComputeUrlMapHeaderActionToProto(o.HeaderAction),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionUrlRewriteToProto converts a UrlMapPathMatcherRouteRuleRouteActionUrlRewrite resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionUrlRewriteToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionUrlRewrite) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionUrlRewrite{
		PathPrefixRewrite: dcl.ValueOrEmptyString(o.PathPrefixRewrite),
		HostRewrite:       dcl.ValueOrEmptyString(o.HostRewrite),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionTimeoutToProto converts a UrlMapPathMatcherRouteRuleRouteActionTimeout resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionTimeoutToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionTimeout) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionTimeout {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionRetryPolicyToProto converts a UrlMapPathMatcherRouteRuleRouteActionRetryPolicy resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionRetryPolicy) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicy{
		NumRetries:    dcl.ValueOrEmptyInt64(o.NumRetries),
		PerTryTimeout: ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutToProto(o.PerTryTimeout),
	}
	for _, r := range o.RetryCondition {
		p.RetryCondition = append(p.RetryCondition, r)
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutToProto converts a UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeoutToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionRetryPolicyPerTryTimeout{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyToProto converts a UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicy{
		BackendService: dcl.ValueOrEmptyString(o.BackendService),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionCorsPolicyToProto converts a UrlMapPathMatcherRouteRuleRouteActionCorsPolicy resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionCorsPolicyToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionCorsPolicy) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionCorsPolicy{
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
func ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicy{
		Delay: ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayToProto(o.Delay),
		Abort: ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortToProto(o.Abort),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayToProto converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelay{
		FixedDelay: ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o.FixedDelay),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelayToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyDelayFixedDelay{
		Seconds: dcl.ValueOrEmptyInt64(o.Seconds),
		Nanos:   dcl.ValueOrEmptyInt64(o.Nanos),
	}
	return p
}

// UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortToProto converts a UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbortToProto(o *compute.UrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort) *computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.ValueOrEmptyInt64(o.HttpStatus),
		Percentage: dcl.ValueOrEmptyDouble(o.Percentage),
	}
	return p
}

// UrlMapPathMatcherRouteRuleUrlRedirectToProto converts a UrlMapPathMatcherRouteRuleUrlRedirect resource to its proto representation.
func ComputeUrlMapPathMatcherRouteRuleUrlRedirectToProto(o *compute.UrlMapPathMatcherRouteRuleUrlRedirect) *computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirect {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapPathMatcherRouteRuleUrlRedirect{
		HostRedirect:         dcl.ValueOrEmptyString(o.HostRedirect),
		PathRedirect:         dcl.ValueOrEmptyString(o.PathRedirect),
		PrefixRedirect:       dcl.ValueOrEmptyString(o.PrefixRedirect),
		RedirectResponseCode: ComputeUrlMapPathMatcherRouteRuleUrlRedirectRedirectResponseCodeEnumToProto(o.RedirectResponseCode),
		HttpsRedirect:        dcl.ValueOrEmptyBool(o.HttpsRedirect),
		StripQuery:           dcl.ValueOrEmptyBool(o.StripQuery),
	}
	return p
}

// UrlMapTestToProto converts a UrlMapTest resource to its proto representation.
func ComputeUrlMapTestToProto(o *compute.UrlMapTest) *computepb.ComputeUrlMapTest {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeUrlMapTest{
		Description:            dcl.ValueOrEmptyString(o.Description),
		Host:                   dcl.ValueOrEmptyString(o.Host),
		Path:                   dcl.ValueOrEmptyString(o.Path),
		ExpectedBackendService: dcl.ValueOrEmptyString(o.ExpectedBackendService),
	}
	return p
}

// UrlMapToProto converts a UrlMap resource to its proto representation.
func UrlMapToProto(resource *compute.UrlMap) *computepb.ComputeUrlMap {
	p := &computepb.ComputeUrlMap{
		DefaultRouteAction: ComputeUrlMapDefaultRouteActionToProto(resource.DefaultRouteAction),
		DefaultService:     dcl.ValueOrEmptyString(resource.DefaultService),
		DefaultUrlRedirect: ComputeUrlMapDefaultUrlRedirectToProto(resource.DefaultUrlRedirect),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		SelfLink:           dcl.ValueOrEmptyString(resource.SelfLink),
		HeaderAction:       ComputeUrlMapHeaderActionToProto(resource.HeaderAction),
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Region:             dcl.ValueOrEmptyString(resource.Region),
		Project:            dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.HostRule {
		p.HostRule = append(p.HostRule, ComputeUrlMapHostRuleToProto(&r))
	}
	for _, r := range resource.PathMatcher {
		p.PathMatcher = append(p.PathMatcher, ComputeUrlMapPathMatcherToProto(&r))
	}
	for _, r := range resource.Test {
		p.Test = append(p.Test, ComputeUrlMapTestToProto(&r))
	}

	return p
}

// ApplyUrlMap handles the gRPC request by passing it to the underlying UrlMap Apply() method.
func (s *UrlMapServer) applyUrlMap(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeUrlMapRequest) (*computepb.ComputeUrlMap, error) {
	p := ProtoToUrlMap(request.GetResource())
	res, err := c.ApplyUrlMap(ctx, p)
	if err != nil {
		return nil, err
	}
	r := UrlMapToProto(res)
	return r, nil
}

// ApplyUrlMap handles the gRPC request by passing it to the underlying UrlMap Apply() method.
func (s *UrlMapServer) ApplyComputeUrlMap(ctx context.Context, request *computepb.ApplyComputeUrlMapRequest) (*computepb.ComputeUrlMap, error) {
	cl, err := createConfigUrlMap(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyUrlMap(ctx, cl, request)
}

// DeleteUrlMap handles the gRPC request by passing it to the underlying UrlMap Delete() method.
func (s *UrlMapServer) DeleteComputeUrlMap(ctx context.Context, request *computepb.DeleteComputeUrlMapRequest) (*emptypb.Empty, error) {

	cl, err := createConfigUrlMap(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteUrlMap(ctx, ProtoToUrlMap(request.GetResource()))

}

// ListComputeUrlMap handles the gRPC request by passing it to the underlying UrlMapList() method.
func (s *UrlMapServer) ListComputeUrlMap(ctx context.Context, request *computepb.ListComputeUrlMapRequest) (*computepb.ListComputeUrlMapResponse, error) {
	cl, err := createConfigUrlMap(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListUrlMap(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeUrlMap
	for _, r := range resources.Items {
		rp := UrlMapToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeUrlMapResponse{Items: protos}, nil
}

func createConfigUrlMap(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
