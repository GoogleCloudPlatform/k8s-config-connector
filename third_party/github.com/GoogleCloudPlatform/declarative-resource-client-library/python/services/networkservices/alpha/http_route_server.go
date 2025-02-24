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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/alpha/networkservices_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
)

// HttpRouteServer implements the gRPC interface for HttpRoute.
type HttpRouteServer struct{}

// ProtoToHttpRouteRulesActionRedirectResponseCodeEnum converts a HttpRouteRulesActionRedirectResponseCodeEnum enum from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum(e alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum) *alpha.HttpRouteRulesActionRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := alpha.HttpRouteRulesActionRedirectResponseCodeEnum(n[len("NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToHttpRouteRules converts a HttpRouteRules object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRules(p *alphapb.NetworkservicesAlphaHttpRouteRules) *alpha.HttpRouteRules {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRules{
		Action: ProtoToNetworkservicesAlphaHttpRouteRulesAction(p.GetAction()),
	}
	for _, r := range p.GetMatches() {
		obj.Matches = append(obj.Matches, *ProtoToNetworkservicesAlphaHttpRouteRulesMatches(r))
	}
	return obj
}

// ProtoToHttpRouteRulesMatches converts a HttpRouteRulesMatches object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesMatches(p *alphapb.NetworkservicesAlphaHttpRouteRulesMatches) *alpha.HttpRouteRulesMatches {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesMatches{
		FullPathMatch: dcl.StringOrNil(p.GetFullPathMatch()),
		PrefixMatch:   dcl.StringOrNil(p.GetPrefixMatch()),
		RegexMatch:    dcl.StringOrNil(p.GetRegexMatch()),
		IgnoreCase:    dcl.Bool(p.GetIgnoreCase()),
	}
	for _, r := range p.GetHeaders() {
		obj.Headers = append(obj.Headers, *ProtoToNetworkservicesAlphaHttpRouteRulesMatchesHeaders(r))
	}
	for _, r := range p.GetQueryParameters() {
		obj.QueryParameters = append(obj.QueryParameters, *ProtoToNetworkservicesAlphaHttpRouteRulesMatchesQueryParameters(r))
	}
	return obj
}

// ProtoToHttpRouteRulesMatchesHeaders converts a HttpRouteRulesMatchesHeaders object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesMatchesHeaders(p *alphapb.NetworkservicesAlphaHttpRouteRulesMatchesHeaders) *alpha.HttpRouteRulesMatchesHeaders {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesMatchesHeaders{
		Header:       dcl.StringOrNil(p.GetHeader()),
		ExactMatch:   dcl.StringOrNil(p.GetExactMatch()),
		RegexMatch:   dcl.StringOrNil(p.GetRegexMatch()),
		PrefixMatch:  dcl.StringOrNil(p.GetPrefixMatch()),
		PresentMatch: dcl.Bool(p.GetPresentMatch()),
		SuffixMatch:  dcl.StringOrNil(p.GetSuffixMatch()),
		RangeMatch:   ProtoToNetworkservicesAlphaHttpRouteRulesMatchesHeadersRangeMatch(p.GetRangeMatch()),
		InvertMatch:  dcl.Bool(p.GetInvertMatch()),
	}
	return obj
}

// ProtoToHttpRouteRulesMatchesHeadersRangeMatch converts a HttpRouteRulesMatchesHeadersRangeMatch object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesMatchesHeadersRangeMatch(p *alphapb.NetworkservicesAlphaHttpRouteRulesMatchesHeadersRangeMatch) *alpha.HttpRouteRulesMatchesHeadersRangeMatch {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesMatchesHeadersRangeMatch{
		Start: dcl.Int64OrNil(p.GetStart()),
		End:   dcl.Int64OrNil(p.GetEnd()),
	}
	return obj
}

// ProtoToHttpRouteRulesMatchesQueryParameters converts a HttpRouteRulesMatchesQueryParameters object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesMatchesQueryParameters(p *alphapb.NetworkservicesAlphaHttpRouteRulesMatchesQueryParameters) *alpha.HttpRouteRulesMatchesQueryParameters {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesMatchesQueryParameters{
		QueryParameter: dcl.StringOrNil(p.GetQueryParameter()),
		ExactMatch:     dcl.StringOrNil(p.GetExactMatch()),
		RegexMatch:     dcl.StringOrNil(p.GetRegexMatch()),
		PresentMatch:   dcl.Bool(p.GetPresentMatch()),
	}
	return obj
}

// ProtoToHttpRouteRulesAction converts a HttpRouteRulesAction object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesAction(p *alphapb.NetworkservicesAlphaHttpRouteRulesAction) *alpha.HttpRouteRulesAction {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesAction{
		Redirect:               ProtoToNetworkservicesAlphaHttpRouteRulesActionRedirect(p.GetRedirect()),
		FaultInjectionPolicy:   ProtoToNetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
		RequestHeaderModifier:  ProtoToNetworkservicesAlphaHttpRouteRulesActionRequestHeaderModifier(p.GetRequestHeaderModifier()),
		ResponseHeaderModifier: ProtoToNetworkservicesAlphaHttpRouteRulesActionResponseHeaderModifier(p.GetResponseHeaderModifier()),
		UrlRewrite:             ProtoToNetworkservicesAlphaHttpRouteRulesActionUrlRewrite(p.GetUrlRewrite()),
		Timeout:                dcl.StringOrNil(p.GetTimeout()),
		RetryPolicy:            ProtoToNetworkservicesAlphaHttpRouteRulesActionRetryPolicy(p.GetRetryPolicy()),
		RequestMirrorPolicy:    ProtoToNetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicy(p.GetRequestMirrorPolicy()),
		CorsPolicy:             ProtoToNetworkservicesAlphaHttpRouteRulesActionCorsPolicy(p.GetCorsPolicy()),
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworkservicesAlphaHttpRouteRulesActionDestinations(r))
	}
	return obj
}

// ProtoToHttpRouteRulesActionDestinations converts a HttpRouteRulesActionDestinations object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionDestinations(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionDestinations) *alpha.HttpRouteRulesActionDestinations {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionDestinations{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRedirect converts a HttpRouteRulesActionRedirect object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionRedirect(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirect) *alpha.HttpRouteRulesActionRedirect {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionRedirect{
		HostRedirect:  dcl.StringOrNil(p.GetHostRedirect()),
		PathRedirect:  dcl.StringOrNil(p.GetPathRedirect()),
		PrefixRewrite: dcl.StringOrNil(p.GetPrefixRewrite()),
		ResponseCode:  ProtoToNetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum(p.GetResponseCode()),
		HttpsRedirect: dcl.Bool(p.GetHttpsRedirect()),
		StripQuery:    dcl.Bool(p.GetStripQuery()),
		PortRedirect:  dcl.Int64OrNil(p.GetPortRedirect()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionFaultInjectionPolicy converts a HttpRouteRulesActionFaultInjectionPolicy object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicy(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicy) *alpha.HttpRouteRulesActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionFaultInjectionPolicy{
		Delay: ProtoToNetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToNetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionFaultInjectionPolicyDelay converts a HttpRouteRulesActionFaultInjectionPolicyDelay object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyDelay(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyDelay) *alpha.HttpRouteRulesActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionFaultInjectionPolicyDelay{
		FixedDelay: dcl.StringOrNil(p.GetFixedDelay()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionFaultInjectionPolicyAbort converts a HttpRouteRulesActionFaultInjectionPolicyAbort object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyAbort(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyAbort) *alpha.HttpRouteRulesActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.GetHttpStatus()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRequestHeaderModifier converts a HttpRouteRulesActionRequestHeaderModifier object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionRequestHeaderModifier(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionRequestHeaderModifier) *alpha.HttpRouteRulesActionRequestHeaderModifier {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionRequestHeaderModifier{}
	for _, r := range p.GetRemove() {
		obj.Remove = append(obj.Remove, r)
	}
	return obj
}

// ProtoToHttpRouteRulesActionResponseHeaderModifier converts a HttpRouteRulesActionResponseHeaderModifier object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionResponseHeaderModifier(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionResponseHeaderModifier) *alpha.HttpRouteRulesActionResponseHeaderModifier {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionResponseHeaderModifier{}
	for _, r := range p.GetRemove() {
		obj.Remove = append(obj.Remove, r)
	}
	return obj
}

// ProtoToHttpRouteRulesActionUrlRewrite converts a HttpRouteRulesActionUrlRewrite object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionUrlRewrite(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionUrlRewrite) *alpha.HttpRouteRulesActionUrlRewrite {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionUrlRewrite{
		PathPrefixRewrite: dcl.StringOrNil(p.GetPathPrefixRewrite()),
		HostRewrite:       dcl.StringOrNil(p.GetHostRewrite()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRetryPolicy converts a HttpRouteRulesActionRetryPolicy object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionRetryPolicy(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionRetryPolicy) *alpha.HttpRouteRulesActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionRetryPolicy{
		NumRetries:    dcl.Int64OrNil(p.GetNumRetries()),
		PerTryTimeout: dcl.StringOrNil(p.GetPerTryTimeout()),
	}
	for _, r := range p.GetRetryConditions() {
		obj.RetryConditions = append(obj.RetryConditions, r)
	}
	return obj
}

// ProtoToHttpRouteRulesActionRequestMirrorPolicy converts a HttpRouteRulesActionRequestMirrorPolicy object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicy(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicy) *alpha.HttpRouteRulesActionRequestMirrorPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionRequestMirrorPolicy{
		Destination: ProtoToNetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyDestination(p.GetDestination()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRequestMirrorPolicyDestination converts a HttpRouteRulesActionRequestMirrorPolicyDestination object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyDestination(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyDestination) *alpha.HttpRouteRulesActionRequestMirrorPolicyDestination {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionRequestMirrorPolicyDestination{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionCorsPolicy converts a HttpRouteRulesActionCorsPolicy object from its proto representation.
func ProtoToNetworkservicesAlphaHttpRouteRulesActionCorsPolicy(p *alphapb.NetworkservicesAlphaHttpRouteRulesActionCorsPolicy) *alpha.HttpRouteRulesActionCorsPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.HttpRouteRulesActionCorsPolicy{
		MaxAge:           dcl.StringOrNil(p.GetMaxAge()),
		AllowCredentials: dcl.Bool(p.GetAllowCredentials()),
		Disabled:         dcl.Bool(p.GetDisabled()),
	}
	for _, r := range p.GetAllowOrigins() {
		obj.AllowOrigins = append(obj.AllowOrigins, r)
	}
	for _, r := range p.GetAllowOriginRegexes() {
		obj.AllowOriginRegexes = append(obj.AllowOriginRegexes, r)
	}
	for _, r := range p.GetAllowMethods() {
		obj.AllowMethods = append(obj.AllowMethods, r)
	}
	for _, r := range p.GetAllowHeaders() {
		obj.AllowHeaders = append(obj.AllowHeaders, r)
	}
	for _, r := range p.GetExposeHeaders() {
		obj.ExposeHeaders = append(obj.ExposeHeaders, r)
	}
	return obj
}

// ProtoToHttpRoute converts a HttpRoute resource from its proto representation.
func ProtoToHttpRoute(p *alphapb.NetworkservicesAlphaHttpRoute) *alpha.HttpRoute {
	obj := &alpha.HttpRoute{
		Name:        dcl.StringOrNil(p.GetName()),
		Description: dcl.StringOrNil(p.GetDescription()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Project:     dcl.StringOrNil(p.GetProject()),
		Location:    dcl.StringOrNil(p.GetLocation()),
		SelfLink:    dcl.StringOrNil(p.GetSelfLink()),
	}
	for _, r := range p.GetHostnames() {
		obj.Hostnames = append(obj.Hostnames, r)
	}
	for _, r := range p.GetMeshes() {
		obj.Meshes = append(obj.Meshes, r)
	}
	for _, r := range p.GetGateways() {
		obj.Gateways = append(obj.Gateways, r)
	}
	for _, r := range p.GetRules() {
		obj.Rules = append(obj.Rules, *ProtoToNetworkservicesAlphaHttpRouteRules(r))
	}
	return obj
}

// HttpRouteRulesActionRedirectResponseCodeEnumToProto converts a HttpRouteRulesActionRedirectResponseCodeEnum enum to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnumToProto(e *alpha.HttpRouteRulesActionRedirectResponseCodeEnum) alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum {
	if e == nil {
		return alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum(0)
	}
	if v, ok := alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum_value["HttpRouteRulesActionRedirectResponseCodeEnum"+string(*e)]; ok {
		return alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum(v)
	}
	return alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnum(0)
}

// HttpRouteRulesToProto converts a HttpRouteRules object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesToProto(o *alpha.HttpRouteRules) *alphapb.NetworkservicesAlphaHttpRouteRules {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRules{}
	p.SetAction(NetworkservicesAlphaHttpRouteRulesActionToProto(o.Action))
	sMatches := make([]*alphapb.NetworkservicesAlphaHttpRouteRulesMatches, len(o.Matches))
	for i, r := range o.Matches {
		sMatches[i] = NetworkservicesAlphaHttpRouteRulesMatchesToProto(&r)
	}
	p.SetMatches(sMatches)
	return p
}

// HttpRouteRulesMatchesToProto converts a HttpRouteRulesMatches object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesMatchesToProto(o *alpha.HttpRouteRulesMatches) *alphapb.NetworkservicesAlphaHttpRouteRulesMatches {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesMatches{}
	p.SetFullPathMatch(dcl.ValueOrEmptyString(o.FullPathMatch))
	p.SetPrefixMatch(dcl.ValueOrEmptyString(o.PrefixMatch))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	p.SetIgnoreCase(dcl.ValueOrEmptyBool(o.IgnoreCase))
	sHeaders := make([]*alphapb.NetworkservicesAlphaHttpRouteRulesMatchesHeaders, len(o.Headers))
	for i, r := range o.Headers {
		sHeaders[i] = NetworkservicesAlphaHttpRouteRulesMatchesHeadersToProto(&r)
	}
	p.SetHeaders(sHeaders)
	sQueryParameters := make([]*alphapb.NetworkservicesAlphaHttpRouteRulesMatchesQueryParameters, len(o.QueryParameters))
	for i, r := range o.QueryParameters {
		sQueryParameters[i] = NetworkservicesAlphaHttpRouteRulesMatchesQueryParametersToProto(&r)
	}
	p.SetQueryParameters(sQueryParameters)
	return p
}

// HttpRouteRulesMatchesHeadersToProto converts a HttpRouteRulesMatchesHeaders object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesMatchesHeadersToProto(o *alpha.HttpRouteRulesMatchesHeaders) *alphapb.NetworkservicesAlphaHttpRouteRulesMatchesHeaders {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesMatchesHeaders{}
	p.SetHeader(dcl.ValueOrEmptyString(o.Header))
	p.SetExactMatch(dcl.ValueOrEmptyString(o.ExactMatch))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	p.SetPrefixMatch(dcl.ValueOrEmptyString(o.PrefixMatch))
	p.SetPresentMatch(dcl.ValueOrEmptyBool(o.PresentMatch))
	p.SetSuffixMatch(dcl.ValueOrEmptyString(o.SuffixMatch))
	p.SetRangeMatch(NetworkservicesAlphaHttpRouteRulesMatchesHeadersRangeMatchToProto(o.RangeMatch))
	p.SetInvertMatch(dcl.ValueOrEmptyBool(o.InvertMatch))
	return p
}

// HttpRouteRulesMatchesHeadersRangeMatchToProto converts a HttpRouteRulesMatchesHeadersRangeMatch object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesMatchesHeadersRangeMatchToProto(o *alpha.HttpRouteRulesMatchesHeadersRangeMatch) *alphapb.NetworkservicesAlphaHttpRouteRulesMatchesHeadersRangeMatch {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesMatchesHeadersRangeMatch{}
	p.SetStart(dcl.ValueOrEmptyInt64(o.Start))
	p.SetEnd(dcl.ValueOrEmptyInt64(o.End))
	return p
}

// HttpRouteRulesMatchesQueryParametersToProto converts a HttpRouteRulesMatchesQueryParameters object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesMatchesQueryParametersToProto(o *alpha.HttpRouteRulesMatchesQueryParameters) *alphapb.NetworkservicesAlphaHttpRouteRulesMatchesQueryParameters {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesMatchesQueryParameters{}
	p.SetQueryParameter(dcl.ValueOrEmptyString(o.QueryParameter))
	p.SetExactMatch(dcl.ValueOrEmptyString(o.ExactMatch))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	p.SetPresentMatch(dcl.ValueOrEmptyBool(o.PresentMatch))
	return p
}

// HttpRouteRulesActionToProto converts a HttpRouteRulesAction object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionToProto(o *alpha.HttpRouteRulesAction) *alphapb.NetworkservicesAlphaHttpRouteRulesAction {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesAction{}
	p.SetRedirect(NetworkservicesAlphaHttpRouteRulesActionRedirectToProto(o.Redirect))
	p.SetFaultInjectionPolicy(NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy))
	p.SetRequestHeaderModifier(NetworkservicesAlphaHttpRouteRulesActionRequestHeaderModifierToProto(o.RequestHeaderModifier))
	p.SetResponseHeaderModifier(NetworkservicesAlphaHttpRouteRulesActionResponseHeaderModifierToProto(o.ResponseHeaderModifier))
	p.SetUrlRewrite(NetworkservicesAlphaHttpRouteRulesActionUrlRewriteToProto(o.UrlRewrite))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetRetryPolicy(NetworkservicesAlphaHttpRouteRulesActionRetryPolicyToProto(o.RetryPolicy))
	p.SetRequestMirrorPolicy(NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyToProto(o.RequestMirrorPolicy))
	p.SetCorsPolicy(NetworkservicesAlphaHttpRouteRulesActionCorsPolicyToProto(o.CorsPolicy))
	sDestinations := make([]*alphapb.NetworkservicesAlphaHttpRouteRulesActionDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworkservicesAlphaHttpRouteRulesActionDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// HttpRouteRulesActionDestinationsToProto converts a HttpRouteRulesActionDestinations object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionDestinationsToProto(o *alpha.HttpRouteRulesActionDestinations) *alphapb.NetworkservicesAlphaHttpRouteRulesActionDestinations {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionDestinations{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// HttpRouteRulesActionRedirectToProto converts a HttpRouteRulesActionRedirect object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionRedirectToProto(o *alpha.HttpRouteRulesActionRedirect) *alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirect {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionRedirect{}
	p.SetHostRedirect(dcl.ValueOrEmptyString(o.HostRedirect))
	p.SetPathRedirect(dcl.ValueOrEmptyString(o.PathRedirect))
	p.SetPrefixRewrite(dcl.ValueOrEmptyString(o.PrefixRewrite))
	p.SetResponseCode(NetworkservicesAlphaHttpRouteRulesActionRedirectResponseCodeEnumToProto(o.ResponseCode))
	p.SetHttpsRedirect(dcl.ValueOrEmptyBool(o.HttpsRedirect))
	p.SetStripQuery(dcl.ValueOrEmptyBool(o.StripQuery))
	p.SetPortRedirect(dcl.ValueOrEmptyInt64(o.PortRedirect))
	return p
}

// HttpRouteRulesActionFaultInjectionPolicyToProto converts a HttpRouteRulesActionFaultInjectionPolicy object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyToProto(o *alpha.HttpRouteRulesActionFaultInjectionPolicy) *alphapb.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicy{}
	p.SetDelay(NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyDelayToProto(o.Delay))
	p.SetAbort(NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyAbortToProto(o.Abort))
	return p
}

// HttpRouteRulesActionFaultInjectionPolicyDelayToProto converts a HttpRouteRulesActionFaultInjectionPolicyDelay object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyDelayToProto(o *alpha.HttpRouteRulesActionFaultInjectionPolicyDelay) *alphapb.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyDelay{}
	p.SetFixedDelay(dcl.ValueOrEmptyString(o.FixedDelay))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// HttpRouteRulesActionFaultInjectionPolicyAbortToProto converts a HttpRouteRulesActionFaultInjectionPolicyAbort object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyAbortToProto(o *alpha.HttpRouteRulesActionFaultInjectionPolicyAbort) *alphapb.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionFaultInjectionPolicyAbort{}
	p.SetHttpStatus(dcl.ValueOrEmptyInt64(o.HttpStatus))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// HttpRouteRulesActionRequestHeaderModifierToProto converts a HttpRouteRulesActionRequestHeaderModifier object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionRequestHeaderModifierToProto(o *alpha.HttpRouteRulesActionRequestHeaderModifier) *alphapb.NetworkservicesAlphaHttpRouteRulesActionRequestHeaderModifier {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionRequestHeaderModifier{}
	mSet := make(map[string]string, len(o.Set))
	for k, r := range o.Set {
		mSet[k] = r
	}
	p.SetSet(mSet)
	mAdd := make(map[string]string, len(o.Add))
	for k, r := range o.Add {
		mAdd[k] = r
	}
	p.SetAdd(mAdd)
	sRemove := make([]string, len(o.Remove))
	for i, r := range o.Remove {
		sRemove[i] = r
	}
	p.SetRemove(sRemove)
	return p
}

// HttpRouteRulesActionResponseHeaderModifierToProto converts a HttpRouteRulesActionResponseHeaderModifier object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionResponseHeaderModifierToProto(o *alpha.HttpRouteRulesActionResponseHeaderModifier) *alphapb.NetworkservicesAlphaHttpRouteRulesActionResponseHeaderModifier {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionResponseHeaderModifier{}
	mSet := make(map[string]string, len(o.Set))
	for k, r := range o.Set {
		mSet[k] = r
	}
	p.SetSet(mSet)
	mAdd := make(map[string]string, len(o.Add))
	for k, r := range o.Add {
		mAdd[k] = r
	}
	p.SetAdd(mAdd)
	sRemove := make([]string, len(o.Remove))
	for i, r := range o.Remove {
		sRemove[i] = r
	}
	p.SetRemove(sRemove)
	return p
}

// HttpRouteRulesActionUrlRewriteToProto converts a HttpRouteRulesActionUrlRewrite object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionUrlRewriteToProto(o *alpha.HttpRouteRulesActionUrlRewrite) *alphapb.NetworkservicesAlphaHttpRouteRulesActionUrlRewrite {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionUrlRewrite{}
	p.SetPathPrefixRewrite(dcl.ValueOrEmptyString(o.PathPrefixRewrite))
	p.SetHostRewrite(dcl.ValueOrEmptyString(o.HostRewrite))
	return p
}

// HttpRouteRulesActionRetryPolicyToProto converts a HttpRouteRulesActionRetryPolicy object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionRetryPolicyToProto(o *alpha.HttpRouteRulesActionRetryPolicy) *alphapb.NetworkservicesAlphaHttpRouteRulesActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionRetryPolicy{}
	p.SetNumRetries(dcl.ValueOrEmptyInt64(o.NumRetries))
	p.SetPerTryTimeout(dcl.ValueOrEmptyString(o.PerTryTimeout))
	sRetryConditions := make([]string, len(o.RetryConditions))
	for i, r := range o.RetryConditions {
		sRetryConditions[i] = r
	}
	p.SetRetryConditions(sRetryConditions)
	return p
}

// HttpRouteRulesActionRequestMirrorPolicyToProto converts a HttpRouteRulesActionRequestMirrorPolicy object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyToProto(o *alpha.HttpRouteRulesActionRequestMirrorPolicy) *alphapb.NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicy{}
	p.SetDestination(NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyDestinationToProto(o.Destination))
	return p
}

// HttpRouteRulesActionRequestMirrorPolicyDestinationToProto converts a HttpRouteRulesActionRequestMirrorPolicyDestination object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyDestinationToProto(o *alpha.HttpRouteRulesActionRequestMirrorPolicyDestination) *alphapb.NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyDestination {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionRequestMirrorPolicyDestination{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// HttpRouteRulesActionCorsPolicyToProto converts a HttpRouteRulesActionCorsPolicy object to its proto representation.
func NetworkservicesAlphaHttpRouteRulesActionCorsPolicyToProto(o *alpha.HttpRouteRulesActionCorsPolicy) *alphapb.NetworkservicesAlphaHttpRouteRulesActionCorsPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaHttpRouteRulesActionCorsPolicy{}
	p.SetMaxAge(dcl.ValueOrEmptyString(o.MaxAge))
	p.SetAllowCredentials(dcl.ValueOrEmptyBool(o.AllowCredentials))
	p.SetDisabled(dcl.ValueOrEmptyBool(o.Disabled))
	sAllowOrigins := make([]string, len(o.AllowOrigins))
	for i, r := range o.AllowOrigins {
		sAllowOrigins[i] = r
	}
	p.SetAllowOrigins(sAllowOrigins)
	sAllowOriginRegexes := make([]string, len(o.AllowOriginRegexes))
	for i, r := range o.AllowOriginRegexes {
		sAllowOriginRegexes[i] = r
	}
	p.SetAllowOriginRegexes(sAllowOriginRegexes)
	sAllowMethods := make([]string, len(o.AllowMethods))
	for i, r := range o.AllowMethods {
		sAllowMethods[i] = r
	}
	p.SetAllowMethods(sAllowMethods)
	sAllowHeaders := make([]string, len(o.AllowHeaders))
	for i, r := range o.AllowHeaders {
		sAllowHeaders[i] = r
	}
	p.SetAllowHeaders(sAllowHeaders)
	sExposeHeaders := make([]string, len(o.ExposeHeaders))
	for i, r := range o.ExposeHeaders {
		sExposeHeaders[i] = r
	}
	p.SetExposeHeaders(sExposeHeaders)
	return p
}

// HttpRouteToProto converts a HttpRoute resource to its proto representation.
func HttpRouteToProto(resource *alpha.HttpRoute) *alphapb.NetworkservicesAlphaHttpRoute {
	p := &alphapb.NetworkservicesAlphaHttpRoute{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	sHostnames := make([]string, len(resource.Hostnames))
	for i, r := range resource.Hostnames {
		sHostnames[i] = r
	}
	p.SetHostnames(sHostnames)
	sMeshes := make([]string, len(resource.Meshes))
	for i, r := range resource.Meshes {
		sMeshes[i] = r
	}
	p.SetMeshes(sMeshes)
	sGateways := make([]string, len(resource.Gateways))
	for i, r := range resource.Gateways {
		sGateways[i] = r
	}
	p.SetGateways(sGateways)
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sRules := make([]*alphapb.NetworkservicesAlphaHttpRouteRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworkservicesAlphaHttpRouteRulesToProto(&r)
	}
	p.SetRules(sRules)

	return p
}

// applyHttpRoute handles the gRPC request by passing it to the underlying HttpRoute Apply() method.
func (s *HttpRouteServer) applyHttpRoute(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkservicesAlphaHttpRouteRequest) (*alphapb.NetworkservicesAlphaHttpRoute, error) {
	p := ProtoToHttpRoute(request.GetResource())
	res, err := c.ApplyHttpRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HttpRouteToProto(res)
	return r, nil
}

// applyNetworkservicesAlphaHttpRoute handles the gRPC request by passing it to the underlying HttpRoute Apply() method.
func (s *HttpRouteServer) ApplyNetworkservicesAlphaHttpRoute(ctx context.Context, request *alphapb.ApplyNetworkservicesAlphaHttpRouteRequest) (*alphapb.NetworkservicesAlphaHttpRoute, error) {
	cl, err := createConfigHttpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyHttpRoute(ctx, cl, request)
}

// DeleteHttpRoute handles the gRPC request by passing it to the underlying HttpRoute Delete() method.
func (s *HttpRouteServer) DeleteNetworkservicesAlphaHttpRoute(ctx context.Context, request *alphapb.DeleteNetworkservicesAlphaHttpRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHttpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHttpRoute(ctx, ProtoToHttpRoute(request.GetResource()))

}

// ListNetworkservicesAlphaHttpRoute handles the gRPC request by passing it to the underlying HttpRouteList() method.
func (s *HttpRouteServer) ListNetworkservicesAlphaHttpRoute(ctx context.Context, request *alphapb.ListNetworkservicesAlphaHttpRouteRequest) (*alphapb.ListNetworkservicesAlphaHttpRouteResponse, error) {
	cl, err := createConfigHttpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHttpRoute(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkservicesAlphaHttpRoute
	for _, r := range resources.Items {
		rp := HttpRouteToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworkservicesAlphaHttpRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigHttpRoute(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
