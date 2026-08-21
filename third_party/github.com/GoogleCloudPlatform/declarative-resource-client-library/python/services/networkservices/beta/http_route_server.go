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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/beta/networkservices_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta"
)

// HttpRouteServer implements the gRPC interface for HttpRoute.
type HttpRouteServer struct{}

// ProtoToHttpRouteRulesActionRedirectResponseCodeEnum converts a HttpRouteRulesActionRedirectResponseCodeEnum enum from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum(e betapb.NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum) *beta.HttpRouteRulesActionRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := beta.HttpRouteRulesActionRedirectResponseCodeEnum(n[len("NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToHttpRouteRules converts a HttpRouteRules object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRules(p *betapb.NetworkservicesBetaHttpRouteRules) *beta.HttpRouteRules {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRules{
		Action: ProtoToNetworkservicesBetaHttpRouteRulesAction(p.GetAction()),
	}
	for _, r := range p.GetMatches() {
		obj.Matches = append(obj.Matches, *ProtoToNetworkservicesBetaHttpRouteRulesMatches(r))
	}
	return obj
}

// ProtoToHttpRouteRulesMatches converts a HttpRouteRulesMatches object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesMatches(p *betapb.NetworkservicesBetaHttpRouteRulesMatches) *beta.HttpRouteRulesMatches {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesMatches{
		FullPathMatch: dcl.StringOrNil(p.GetFullPathMatch()),
		PrefixMatch:   dcl.StringOrNil(p.GetPrefixMatch()),
		RegexMatch:    dcl.StringOrNil(p.GetRegexMatch()),
		IgnoreCase:    dcl.Bool(p.GetIgnoreCase()),
	}
	for _, r := range p.GetHeaders() {
		obj.Headers = append(obj.Headers, *ProtoToNetworkservicesBetaHttpRouteRulesMatchesHeaders(r))
	}
	for _, r := range p.GetQueryParameters() {
		obj.QueryParameters = append(obj.QueryParameters, *ProtoToNetworkservicesBetaHttpRouteRulesMatchesQueryParameters(r))
	}
	return obj
}

// ProtoToHttpRouteRulesMatchesHeaders converts a HttpRouteRulesMatchesHeaders object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesMatchesHeaders(p *betapb.NetworkservicesBetaHttpRouteRulesMatchesHeaders) *beta.HttpRouteRulesMatchesHeaders {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesMatchesHeaders{
		Header:       dcl.StringOrNil(p.GetHeader()),
		ExactMatch:   dcl.StringOrNil(p.GetExactMatch()),
		RegexMatch:   dcl.StringOrNil(p.GetRegexMatch()),
		PrefixMatch:  dcl.StringOrNil(p.GetPrefixMatch()),
		PresentMatch: dcl.Bool(p.GetPresentMatch()),
		SuffixMatch:  dcl.StringOrNil(p.GetSuffixMatch()),
		RangeMatch:   ProtoToNetworkservicesBetaHttpRouteRulesMatchesHeadersRangeMatch(p.GetRangeMatch()),
		InvertMatch:  dcl.Bool(p.GetInvertMatch()),
	}
	return obj
}

// ProtoToHttpRouteRulesMatchesHeadersRangeMatch converts a HttpRouteRulesMatchesHeadersRangeMatch object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesMatchesHeadersRangeMatch(p *betapb.NetworkservicesBetaHttpRouteRulesMatchesHeadersRangeMatch) *beta.HttpRouteRulesMatchesHeadersRangeMatch {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesMatchesHeadersRangeMatch{
		Start: dcl.Int64OrNil(p.GetStart()),
		End:   dcl.Int64OrNil(p.GetEnd()),
	}
	return obj
}

// ProtoToHttpRouteRulesMatchesQueryParameters converts a HttpRouteRulesMatchesQueryParameters object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesMatchesQueryParameters(p *betapb.NetworkservicesBetaHttpRouteRulesMatchesQueryParameters) *beta.HttpRouteRulesMatchesQueryParameters {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesMatchesQueryParameters{
		QueryParameter: dcl.StringOrNil(p.GetQueryParameter()),
		ExactMatch:     dcl.StringOrNil(p.GetExactMatch()),
		RegexMatch:     dcl.StringOrNil(p.GetRegexMatch()),
		PresentMatch:   dcl.Bool(p.GetPresentMatch()),
	}
	return obj
}

// ProtoToHttpRouteRulesAction converts a HttpRouteRulesAction object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesAction(p *betapb.NetworkservicesBetaHttpRouteRulesAction) *beta.HttpRouteRulesAction {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesAction{
		Redirect:               ProtoToNetworkservicesBetaHttpRouteRulesActionRedirect(p.GetRedirect()),
		FaultInjectionPolicy:   ProtoToNetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
		RequestHeaderModifier:  ProtoToNetworkservicesBetaHttpRouteRulesActionRequestHeaderModifier(p.GetRequestHeaderModifier()),
		ResponseHeaderModifier: ProtoToNetworkservicesBetaHttpRouteRulesActionResponseHeaderModifier(p.GetResponseHeaderModifier()),
		UrlRewrite:             ProtoToNetworkservicesBetaHttpRouteRulesActionUrlRewrite(p.GetUrlRewrite()),
		Timeout:                dcl.StringOrNil(p.GetTimeout()),
		RetryPolicy:            ProtoToNetworkservicesBetaHttpRouteRulesActionRetryPolicy(p.GetRetryPolicy()),
		RequestMirrorPolicy:    ProtoToNetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicy(p.GetRequestMirrorPolicy()),
		CorsPolicy:             ProtoToNetworkservicesBetaHttpRouteRulesActionCorsPolicy(p.GetCorsPolicy()),
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworkservicesBetaHttpRouteRulesActionDestinations(r))
	}
	return obj
}

// ProtoToHttpRouteRulesActionDestinations converts a HttpRouteRulesActionDestinations object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionDestinations(p *betapb.NetworkservicesBetaHttpRouteRulesActionDestinations) *beta.HttpRouteRulesActionDestinations {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionDestinations{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRedirect converts a HttpRouteRulesActionRedirect object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionRedirect(p *betapb.NetworkservicesBetaHttpRouteRulesActionRedirect) *beta.HttpRouteRulesActionRedirect {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionRedirect{
		HostRedirect:  dcl.StringOrNil(p.GetHostRedirect()),
		PathRedirect:  dcl.StringOrNil(p.GetPathRedirect()),
		PrefixRewrite: dcl.StringOrNil(p.GetPrefixRewrite()),
		ResponseCode:  ProtoToNetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum(p.GetResponseCode()),
		HttpsRedirect: dcl.Bool(p.GetHttpsRedirect()),
		StripQuery:    dcl.Bool(p.GetStripQuery()),
		PortRedirect:  dcl.Int64OrNil(p.GetPortRedirect()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionFaultInjectionPolicy converts a HttpRouteRulesActionFaultInjectionPolicy object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicy(p *betapb.NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicy) *beta.HttpRouteRulesActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionFaultInjectionPolicy{
		Delay: ProtoToNetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToNetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionFaultInjectionPolicyDelay converts a HttpRouteRulesActionFaultInjectionPolicyDelay object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyDelay(p *betapb.NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyDelay) *beta.HttpRouteRulesActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionFaultInjectionPolicyDelay{
		FixedDelay: dcl.StringOrNil(p.GetFixedDelay()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionFaultInjectionPolicyAbort converts a HttpRouteRulesActionFaultInjectionPolicyAbort object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyAbort(p *betapb.NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyAbort) *beta.HttpRouteRulesActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.GetHttpStatus()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRequestHeaderModifier converts a HttpRouteRulesActionRequestHeaderModifier object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionRequestHeaderModifier(p *betapb.NetworkservicesBetaHttpRouteRulesActionRequestHeaderModifier) *beta.HttpRouteRulesActionRequestHeaderModifier {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionRequestHeaderModifier{}
	for _, r := range p.GetRemove() {
		obj.Remove = append(obj.Remove, r)
	}
	return obj
}

// ProtoToHttpRouteRulesActionResponseHeaderModifier converts a HttpRouteRulesActionResponseHeaderModifier object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionResponseHeaderModifier(p *betapb.NetworkservicesBetaHttpRouteRulesActionResponseHeaderModifier) *beta.HttpRouteRulesActionResponseHeaderModifier {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionResponseHeaderModifier{}
	for _, r := range p.GetRemove() {
		obj.Remove = append(obj.Remove, r)
	}
	return obj
}

// ProtoToHttpRouteRulesActionUrlRewrite converts a HttpRouteRulesActionUrlRewrite object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionUrlRewrite(p *betapb.NetworkservicesBetaHttpRouteRulesActionUrlRewrite) *beta.HttpRouteRulesActionUrlRewrite {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionUrlRewrite{
		PathPrefixRewrite: dcl.StringOrNil(p.GetPathPrefixRewrite()),
		HostRewrite:       dcl.StringOrNil(p.GetHostRewrite()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRetryPolicy converts a HttpRouteRulesActionRetryPolicy object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionRetryPolicy(p *betapb.NetworkservicesBetaHttpRouteRulesActionRetryPolicy) *beta.HttpRouteRulesActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionRetryPolicy{
		NumRetries:    dcl.Int64OrNil(p.GetNumRetries()),
		PerTryTimeout: dcl.StringOrNil(p.GetPerTryTimeout()),
	}
	for _, r := range p.GetRetryConditions() {
		obj.RetryConditions = append(obj.RetryConditions, r)
	}
	return obj
}

// ProtoToHttpRouteRulesActionRequestMirrorPolicy converts a HttpRouteRulesActionRequestMirrorPolicy object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicy(p *betapb.NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicy) *beta.HttpRouteRulesActionRequestMirrorPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionRequestMirrorPolicy{
		Destination: ProtoToNetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicyDestination(p.GetDestination()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRequestMirrorPolicyDestination converts a HttpRouteRulesActionRequestMirrorPolicyDestination object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicyDestination(p *betapb.NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicyDestination) *beta.HttpRouteRulesActionRequestMirrorPolicyDestination {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionRequestMirrorPolicyDestination{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionCorsPolicy converts a HttpRouteRulesActionCorsPolicy object from its proto representation.
func ProtoToNetworkservicesBetaHttpRouteRulesActionCorsPolicy(p *betapb.NetworkservicesBetaHttpRouteRulesActionCorsPolicy) *beta.HttpRouteRulesActionCorsPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.HttpRouteRulesActionCorsPolicy{
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
func ProtoToHttpRoute(p *betapb.NetworkservicesBetaHttpRoute) *beta.HttpRoute {
	obj := &beta.HttpRoute{
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
		obj.Rules = append(obj.Rules, *ProtoToNetworkservicesBetaHttpRouteRules(r))
	}
	return obj
}

// HttpRouteRulesActionRedirectResponseCodeEnumToProto converts a HttpRouteRulesActionRedirectResponseCodeEnum enum to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnumToProto(e *beta.HttpRouteRulesActionRedirectResponseCodeEnum) betapb.NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum {
	if e == nil {
		return betapb.NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum(0)
	}
	if v, ok := betapb.NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum_value["HttpRouteRulesActionRedirectResponseCodeEnum"+string(*e)]; ok {
		return betapb.NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum(v)
	}
	return betapb.NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnum(0)
}

// HttpRouteRulesToProto converts a HttpRouteRules object to its proto representation.
func NetworkservicesBetaHttpRouteRulesToProto(o *beta.HttpRouteRules) *betapb.NetworkservicesBetaHttpRouteRules {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRules{}
	p.SetAction(NetworkservicesBetaHttpRouteRulesActionToProto(o.Action))
	sMatches := make([]*betapb.NetworkservicesBetaHttpRouteRulesMatches, len(o.Matches))
	for i, r := range o.Matches {
		sMatches[i] = NetworkservicesBetaHttpRouteRulesMatchesToProto(&r)
	}
	p.SetMatches(sMatches)
	return p
}

// HttpRouteRulesMatchesToProto converts a HttpRouteRulesMatches object to its proto representation.
func NetworkservicesBetaHttpRouteRulesMatchesToProto(o *beta.HttpRouteRulesMatches) *betapb.NetworkservicesBetaHttpRouteRulesMatches {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesMatches{}
	p.SetFullPathMatch(dcl.ValueOrEmptyString(o.FullPathMatch))
	p.SetPrefixMatch(dcl.ValueOrEmptyString(o.PrefixMatch))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	p.SetIgnoreCase(dcl.ValueOrEmptyBool(o.IgnoreCase))
	sHeaders := make([]*betapb.NetworkservicesBetaHttpRouteRulesMatchesHeaders, len(o.Headers))
	for i, r := range o.Headers {
		sHeaders[i] = NetworkservicesBetaHttpRouteRulesMatchesHeadersToProto(&r)
	}
	p.SetHeaders(sHeaders)
	sQueryParameters := make([]*betapb.NetworkservicesBetaHttpRouteRulesMatchesQueryParameters, len(o.QueryParameters))
	for i, r := range o.QueryParameters {
		sQueryParameters[i] = NetworkservicesBetaHttpRouteRulesMatchesQueryParametersToProto(&r)
	}
	p.SetQueryParameters(sQueryParameters)
	return p
}

// HttpRouteRulesMatchesHeadersToProto converts a HttpRouteRulesMatchesHeaders object to its proto representation.
func NetworkservicesBetaHttpRouteRulesMatchesHeadersToProto(o *beta.HttpRouteRulesMatchesHeaders) *betapb.NetworkservicesBetaHttpRouteRulesMatchesHeaders {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesMatchesHeaders{}
	p.SetHeader(dcl.ValueOrEmptyString(o.Header))
	p.SetExactMatch(dcl.ValueOrEmptyString(o.ExactMatch))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	p.SetPrefixMatch(dcl.ValueOrEmptyString(o.PrefixMatch))
	p.SetPresentMatch(dcl.ValueOrEmptyBool(o.PresentMatch))
	p.SetSuffixMatch(dcl.ValueOrEmptyString(o.SuffixMatch))
	p.SetRangeMatch(NetworkservicesBetaHttpRouteRulesMatchesHeadersRangeMatchToProto(o.RangeMatch))
	p.SetInvertMatch(dcl.ValueOrEmptyBool(o.InvertMatch))
	return p
}

// HttpRouteRulesMatchesHeadersRangeMatchToProto converts a HttpRouteRulesMatchesHeadersRangeMatch object to its proto representation.
func NetworkservicesBetaHttpRouteRulesMatchesHeadersRangeMatchToProto(o *beta.HttpRouteRulesMatchesHeadersRangeMatch) *betapb.NetworkservicesBetaHttpRouteRulesMatchesHeadersRangeMatch {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesMatchesHeadersRangeMatch{}
	p.SetStart(dcl.ValueOrEmptyInt64(o.Start))
	p.SetEnd(dcl.ValueOrEmptyInt64(o.End))
	return p
}

// HttpRouteRulesMatchesQueryParametersToProto converts a HttpRouteRulesMatchesQueryParameters object to its proto representation.
func NetworkservicesBetaHttpRouteRulesMatchesQueryParametersToProto(o *beta.HttpRouteRulesMatchesQueryParameters) *betapb.NetworkservicesBetaHttpRouteRulesMatchesQueryParameters {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesMatchesQueryParameters{}
	p.SetQueryParameter(dcl.ValueOrEmptyString(o.QueryParameter))
	p.SetExactMatch(dcl.ValueOrEmptyString(o.ExactMatch))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	p.SetPresentMatch(dcl.ValueOrEmptyBool(o.PresentMatch))
	return p
}

// HttpRouteRulesActionToProto converts a HttpRouteRulesAction object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionToProto(o *beta.HttpRouteRulesAction) *betapb.NetworkservicesBetaHttpRouteRulesAction {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesAction{}
	p.SetRedirect(NetworkservicesBetaHttpRouteRulesActionRedirectToProto(o.Redirect))
	p.SetFaultInjectionPolicy(NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy))
	p.SetRequestHeaderModifier(NetworkservicesBetaHttpRouteRulesActionRequestHeaderModifierToProto(o.RequestHeaderModifier))
	p.SetResponseHeaderModifier(NetworkservicesBetaHttpRouteRulesActionResponseHeaderModifierToProto(o.ResponseHeaderModifier))
	p.SetUrlRewrite(NetworkservicesBetaHttpRouteRulesActionUrlRewriteToProto(o.UrlRewrite))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetRetryPolicy(NetworkservicesBetaHttpRouteRulesActionRetryPolicyToProto(o.RetryPolicy))
	p.SetRequestMirrorPolicy(NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicyToProto(o.RequestMirrorPolicy))
	p.SetCorsPolicy(NetworkservicesBetaHttpRouteRulesActionCorsPolicyToProto(o.CorsPolicy))
	sDestinations := make([]*betapb.NetworkservicesBetaHttpRouteRulesActionDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworkservicesBetaHttpRouteRulesActionDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// HttpRouteRulesActionDestinationsToProto converts a HttpRouteRulesActionDestinations object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionDestinationsToProto(o *beta.HttpRouteRulesActionDestinations) *betapb.NetworkservicesBetaHttpRouteRulesActionDestinations {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionDestinations{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// HttpRouteRulesActionRedirectToProto converts a HttpRouteRulesActionRedirect object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionRedirectToProto(o *beta.HttpRouteRulesActionRedirect) *betapb.NetworkservicesBetaHttpRouteRulesActionRedirect {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionRedirect{}
	p.SetHostRedirect(dcl.ValueOrEmptyString(o.HostRedirect))
	p.SetPathRedirect(dcl.ValueOrEmptyString(o.PathRedirect))
	p.SetPrefixRewrite(dcl.ValueOrEmptyString(o.PrefixRewrite))
	p.SetResponseCode(NetworkservicesBetaHttpRouteRulesActionRedirectResponseCodeEnumToProto(o.ResponseCode))
	p.SetHttpsRedirect(dcl.ValueOrEmptyBool(o.HttpsRedirect))
	p.SetStripQuery(dcl.ValueOrEmptyBool(o.StripQuery))
	p.SetPortRedirect(dcl.ValueOrEmptyInt64(o.PortRedirect))
	return p
}

// HttpRouteRulesActionFaultInjectionPolicyToProto converts a HttpRouteRulesActionFaultInjectionPolicy object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyToProto(o *beta.HttpRouteRulesActionFaultInjectionPolicy) *betapb.NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicy{}
	p.SetDelay(NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyDelayToProto(o.Delay))
	p.SetAbort(NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyAbortToProto(o.Abort))
	return p
}

// HttpRouteRulesActionFaultInjectionPolicyDelayToProto converts a HttpRouteRulesActionFaultInjectionPolicyDelay object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyDelayToProto(o *beta.HttpRouteRulesActionFaultInjectionPolicyDelay) *betapb.NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyDelay{}
	p.SetFixedDelay(dcl.ValueOrEmptyString(o.FixedDelay))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// HttpRouteRulesActionFaultInjectionPolicyAbortToProto converts a HttpRouteRulesActionFaultInjectionPolicyAbort object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyAbortToProto(o *beta.HttpRouteRulesActionFaultInjectionPolicyAbort) *betapb.NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionFaultInjectionPolicyAbort{}
	p.SetHttpStatus(dcl.ValueOrEmptyInt64(o.HttpStatus))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// HttpRouteRulesActionRequestHeaderModifierToProto converts a HttpRouteRulesActionRequestHeaderModifier object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionRequestHeaderModifierToProto(o *beta.HttpRouteRulesActionRequestHeaderModifier) *betapb.NetworkservicesBetaHttpRouteRulesActionRequestHeaderModifier {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionRequestHeaderModifier{}
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
func NetworkservicesBetaHttpRouteRulesActionResponseHeaderModifierToProto(o *beta.HttpRouteRulesActionResponseHeaderModifier) *betapb.NetworkservicesBetaHttpRouteRulesActionResponseHeaderModifier {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionResponseHeaderModifier{}
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
func NetworkservicesBetaHttpRouteRulesActionUrlRewriteToProto(o *beta.HttpRouteRulesActionUrlRewrite) *betapb.NetworkservicesBetaHttpRouteRulesActionUrlRewrite {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionUrlRewrite{}
	p.SetPathPrefixRewrite(dcl.ValueOrEmptyString(o.PathPrefixRewrite))
	p.SetHostRewrite(dcl.ValueOrEmptyString(o.HostRewrite))
	return p
}

// HttpRouteRulesActionRetryPolicyToProto converts a HttpRouteRulesActionRetryPolicy object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionRetryPolicyToProto(o *beta.HttpRouteRulesActionRetryPolicy) *betapb.NetworkservicesBetaHttpRouteRulesActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionRetryPolicy{}
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
func NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicyToProto(o *beta.HttpRouteRulesActionRequestMirrorPolicy) *betapb.NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicy{}
	p.SetDestination(NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicyDestinationToProto(o.Destination))
	return p
}

// HttpRouteRulesActionRequestMirrorPolicyDestinationToProto converts a HttpRouteRulesActionRequestMirrorPolicyDestination object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicyDestinationToProto(o *beta.HttpRouteRulesActionRequestMirrorPolicyDestination) *betapb.NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicyDestination {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionRequestMirrorPolicyDestination{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// HttpRouteRulesActionCorsPolicyToProto converts a HttpRouteRulesActionCorsPolicy object to its proto representation.
func NetworkservicesBetaHttpRouteRulesActionCorsPolicyToProto(o *beta.HttpRouteRulesActionCorsPolicy) *betapb.NetworkservicesBetaHttpRouteRulesActionCorsPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaHttpRouteRulesActionCorsPolicy{}
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
func HttpRouteToProto(resource *beta.HttpRoute) *betapb.NetworkservicesBetaHttpRoute {
	p := &betapb.NetworkservicesBetaHttpRoute{}
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
	sRules := make([]*betapb.NetworkservicesBetaHttpRouteRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworkservicesBetaHttpRouteRulesToProto(&r)
	}
	p.SetRules(sRules)

	return p
}

// applyHttpRoute handles the gRPC request by passing it to the underlying HttpRoute Apply() method.
func (s *HttpRouteServer) applyHttpRoute(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkservicesBetaHttpRouteRequest) (*betapb.NetworkservicesBetaHttpRoute, error) {
	p := ProtoToHttpRoute(request.GetResource())
	res, err := c.ApplyHttpRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HttpRouteToProto(res)
	return r, nil
}

// applyNetworkservicesBetaHttpRoute handles the gRPC request by passing it to the underlying HttpRoute Apply() method.
func (s *HttpRouteServer) ApplyNetworkservicesBetaHttpRoute(ctx context.Context, request *betapb.ApplyNetworkservicesBetaHttpRouteRequest) (*betapb.NetworkservicesBetaHttpRoute, error) {
	cl, err := createConfigHttpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyHttpRoute(ctx, cl, request)
}

// DeleteHttpRoute handles the gRPC request by passing it to the underlying HttpRoute Delete() method.
func (s *HttpRouteServer) DeleteNetworkservicesBetaHttpRoute(ctx context.Context, request *betapb.DeleteNetworkservicesBetaHttpRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHttpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHttpRoute(ctx, ProtoToHttpRoute(request.GetResource()))

}

// ListNetworkservicesBetaHttpRoute handles the gRPC request by passing it to the underlying HttpRouteList() method.
func (s *HttpRouteServer) ListNetworkservicesBetaHttpRoute(ctx context.Context, request *betapb.ListNetworkservicesBetaHttpRouteRequest) (*betapb.ListNetworkservicesBetaHttpRouteResponse, error) {
	cl, err := createConfigHttpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHttpRoute(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkservicesBetaHttpRoute
	for _, r := range resources.Items {
		rp := HttpRouteToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworkservicesBetaHttpRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigHttpRoute(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
