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
	networkservicespb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/networkservices_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices"
)

// HttpRouteServer implements the gRPC interface for HttpRoute.
type HttpRouteServer struct{}

// ProtoToHttpRouteRulesActionRedirectResponseCodeEnum converts a HttpRouteRulesActionRedirectResponseCodeEnum enum from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum(e networkservicespb.NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum) *networkservices.HttpRouteRulesActionRedirectResponseCodeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := networkservicespb.NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum_name[int32(e)]; ok {
		e := networkservices.HttpRouteRulesActionRedirectResponseCodeEnum(n[len("NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum"):])
		return &e
	}
	return nil
}

// ProtoToHttpRouteRules converts a HttpRouteRules object from its proto representation.
func ProtoToNetworkservicesHttpRouteRules(p *networkservicespb.NetworkservicesHttpRouteRules) *networkservices.HttpRouteRules {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRules{
		Action: ProtoToNetworkservicesHttpRouteRulesAction(p.GetAction()),
	}
	for _, r := range p.GetMatches() {
		obj.Matches = append(obj.Matches, *ProtoToNetworkservicesHttpRouteRulesMatches(r))
	}
	return obj
}

// ProtoToHttpRouteRulesMatches converts a HttpRouteRulesMatches object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesMatches(p *networkservicespb.NetworkservicesHttpRouteRulesMatches) *networkservices.HttpRouteRulesMatches {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesMatches{
		FullPathMatch: dcl.StringOrNil(p.GetFullPathMatch()),
		PrefixMatch:   dcl.StringOrNil(p.GetPrefixMatch()),
		RegexMatch:    dcl.StringOrNil(p.GetRegexMatch()),
		IgnoreCase:    dcl.Bool(p.GetIgnoreCase()),
	}
	for _, r := range p.GetHeaders() {
		obj.Headers = append(obj.Headers, *ProtoToNetworkservicesHttpRouteRulesMatchesHeaders(r))
	}
	for _, r := range p.GetQueryParameters() {
		obj.QueryParameters = append(obj.QueryParameters, *ProtoToNetworkservicesHttpRouteRulesMatchesQueryParameters(r))
	}
	return obj
}

// ProtoToHttpRouteRulesMatchesHeaders converts a HttpRouteRulesMatchesHeaders object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesMatchesHeaders(p *networkservicespb.NetworkservicesHttpRouteRulesMatchesHeaders) *networkservices.HttpRouteRulesMatchesHeaders {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesMatchesHeaders{
		Header:       dcl.StringOrNil(p.GetHeader()),
		ExactMatch:   dcl.StringOrNil(p.GetExactMatch()),
		RegexMatch:   dcl.StringOrNil(p.GetRegexMatch()),
		PrefixMatch:  dcl.StringOrNil(p.GetPrefixMatch()),
		PresentMatch: dcl.Bool(p.GetPresentMatch()),
		SuffixMatch:  dcl.StringOrNil(p.GetSuffixMatch()),
		RangeMatch:   ProtoToNetworkservicesHttpRouteRulesMatchesHeadersRangeMatch(p.GetRangeMatch()),
		InvertMatch:  dcl.Bool(p.GetInvertMatch()),
	}
	return obj
}

// ProtoToHttpRouteRulesMatchesHeadersRangeMatch converts a HttpRouteRulesMatchesHeadersRangeMatch object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesMatchesHeadersRangeMatch(p *networkservicespb.NetworkservicesHttpRouteRulesMatchesHeadersRangeMatch) *networkservices.HttpRouteRulesMatchesHeadersRangeMatch {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesMatchesHeadersRangeMatch{
		Start: dcl.Int64OrNil(p.GetStart()),
		End:   dcl.Int64OrNil(p.GetEnd()),
	}
	return obj
}

// ProtoToHttpRouteRulesMatchesQueryParameters converts a HttpRouteRulesMatchesQueryParameters object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesMatchesQueryParameters(p *networkservicespb.NetworkservicesHttpRouteRulesMatchesQueryParameters) *networkservices.HttpRouteRulesMatchesQueryParameters {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesMatchesQueryParameters{
		QueryParameter: dcl.StringOrNil(p.GetQueryParameter()),
		ExactMatch:     dcl.StringOrNil(p.GetExactMatch()),
		RegexMatch:     dcl.StringOrNil(p.GetRegexMatch()),
		PresentMatch:   dcl.Bool(p.GetPresentMatch()),
	}
	return obj
}

// ProtoToHttpRouteRulesAction converts a HttpRouteRulesAction object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesAction(p *networkservicespb.NetworkservicesHttpRouteRulesAction) *networkservices.HttpRouteRulesAction {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesAction{
		Redirect:               ProtoToNetworkservicesHttpRouteRulesActionRedirect(p.GetRedirect()),
		FaultInjectionPolicy:   ProtoToNetworkservicesHttpRouteRulesActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
		RequestHeaderModifier:  ProtoToNetworkservicesHttpRouteRulesActionRequestHeaderModifier(p.GetRequestHeaderModifier()),
		ResponseHeaderModifier: ProtoToNetworkservicesHttpRouteRulesActionResponseHeaderModifier(p.GetResponseHeaderModifier()),
		UrlRewrite:             ProtoToNetworkservicesHttpRouteRulesActionUrlRewrite(p.GetUrlRewrite()),
		Timeout:                dcl.StringOrNil(p.GetTimeout()),
		RetryPolicy:            ProtoToNetworkservicesHttpRouteRulesActionRetryPolicy(p.GetRetryPolicy()),
		RequestMirrorPolicy:    ProtoToNetworkservicesHttpRouteRulesActionRequestMirrorPolicy(p.GetRequestMirrorPolicy()),
		CorsPolicy:             ProtoToNetworkservicesHttpRouteRulesActionCorsPolicy(p.GetCorsPolicy()),
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworkservicesHttpRouteRulesActionDestinations(r))
	}
	return obj
}

// ProtoToHttpRouteRulesActionDestinations converts a HttpRouteRulesActionDestinations object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionDestinations(p *networkservicespb.NetworkservicesHttpRouteRulesActionDestinations) *networkservices.HttpRouteRulesActionDestinations {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionDestinations{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRedirect converts a HttpRouteRulesActionRedirect object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionRedirect(p *networkservicespb.NetworkservicesHttpRouteRulesActionRedirect) *networkservices.HttpRouteRulesActionRedirect {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionRedirect{
		HostRedirect:  dcl.StringOrNil(p.GetHostRedirect()),
		PathRedirect:  dcl.StringOrNil(p.GetPathRedirect()),
		PrefixRewrite: dcl.StringOrNil(p.GetPrefixRewrite()),
		ResponseCode:  ProtoToNetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum(p.GetResponseCode()),
		HttpsRedirect: dcl.Bool(p.GetHttpsRedirect()),
		StripQuery:    dcl.Bool(p.GetStripQuery()),
		PortRedirect:  dcl.Int64OrNil(p.GetPortRedirect()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionFaultInjectionPolicy converts a HttpRouteRulesActionFaultInjectionPolicy object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionFaultInjectionPolicy(p *networkservicespb.NetworkservicesHttpRouteRulesActionFaultInjectionPolicy) *networkservices.HttpRouteRulesActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionFaultInjectionPolicy{
		Delay: ProtoToNetworkservicesHttpRouteRulesActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToNetworkservicesHttpRouteRulesActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionFaultInjectionPolicyDelay converts a HttpRouteRulesActionFaultInjectionPolicyDelay object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionFaultInjectionPolicyDelay(p *networkservicespb.NetworkservicesHttpRouteRulesActionFaultInjectionPolicyDelay) *networkservices.HttpRouteRulesActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionFaultInjectionPolicyDelay{
		FixedDelay: dcl.StringOrNil(p.GetFixedDelay()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionFaultInjectionPolicyAbort converts a HttpRouteRulesActionFaultInjectionPolicyAbort object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionFaultInjectionPolicyAbort(p *networkservicespb.NetworkservicesHttpRouteRulesActionFaultInjectionPolicyAbort) *networkservices.HttpRouteRulesActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.GetHttpStatus()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRequestHeaderModifier converts a HttpRouteRulesActionRequestHeaderModifier object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionRequestHeaderModifier(p *networkservicespb.NetworkservicesHttpRouteRulesActionRequestHeaderModifier) *networkservices.HttpRouteRulesActionRequestHeaderModifier {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionRequestHeaderModifier{}
	for _, r := range p.GetRemove() {
		obj.Remove = append(obj.Remove, r)
	}
	return obj
}

// ProtoToHttpRouteRulesActionResponseHeaderModifier converts a HttpRouteRulesActionResponseHeaderModifier object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionResponseHeaderModifier(p *networkservicespb.NetworkservicesHttpRouteRulesActionResponseHeaderModifier) *networkservices.HttpRouteRulesActionResponseHeaderModifier {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionResponseHeaderModifier{}
	for _, r := range p.GetRemove() {
		obj.Remove = append(obj.Remove, r)
	}
	return obj
}

// ProtoToHttpRouteRulesActionUrlRewrite converts a HttpRouteRulesActionUrlRewrite object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionUrlRewrite(p *networkservicespb.NetworkservicesHttpRouteRulesActionUrlRewrite) *networkservices.HttpRouteRulesActionUrlRewrite {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionUrlRewrite{
		PathPrefixRewrite: dcl.StringOrNil(p.GetPathPrefixRewrite()),
		HostRewrite:       dcl.StringOrNil(p.GetHostRewrite()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRetryPolicy converts a HttpRouteRulesActionRetryPolicy object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionRetryPolicy(p *networkservicespb.NetworkservicesHttpRouteRulesActionRetryPolicy) *networkservices.HttpRouteRulesActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionRetryPolicy{
		NumRetries:    dcl.Int64OrNil(p.GetNumRetries()),
		PerTryTimeout: dcl.StringOrNil(p.GetPerTryTimeout()),
	}
	for _, r := range p.GetRetryConditions() {
		obj.RetryConditions = append(obj.RetryConditions, r)
	}
	return obj
}

// ProtoToHttpRouteRulesActionRequestMirrorPolicy converts a HttpRouteRulesActionRequestMirrorPolicy object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionRequestMirrorPolicy(p *networkservicespb.NetworkservicesHttpRouteRulesActionRequestMirrorPolicy) *networkservices.HttpRouteRulesActionRequestMirrorPolicy {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionRequestMirrorPolicy{
		Destination: ProtoToNetworkservicesHttpRouteRulesActionRequestMirrorPolicyDestination(p.GetDestination()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionRequestMirrorPolicyDestination converts a HttpRouteRulesActionRequestMirrorPolicyDestination object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionRequestMirrorPolicyDestination(p *networkservicespb.NetworkservicesHttpRouteRulesActionRequestMirrorPolicyDestination) *networkservices.HttpRouteRulesActionRequestMirrorPolicyDestination {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionRequestMirrorPolicyDestination{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToHttpRouteRulesActionCorsPolicy converts a HttpRouteRulesActionCorsPolicy object from its proto representation.
func ProtoToNetworkservicesHttpRouteRulesActionCorsPolicy(p *networkservicespb.NetworkservicesHttpRouteRulesActionCorsPolicy) *networkservices.HttpRouteRulesActionCorsPolicy {
	if p == nil {
		return nil
	}
	obj := &networkservices.HttpRouteRulesActionCorsPolicy{
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
func ProtoToHttpRoute(p *networkservicespb.NetworkservicesHttpRoute) *networkservices.HttpRoute {
	obj := &networkservices.HttpRoute{
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
		obj.Rules = append(obj.Rules, *ProtoToNetworkservicesHttpRouteRules(r))
	}
	return obj
}

// HttpRouteRulesActionRedirectResponseCodeEnumToProto converts a HttpRouteRulesActionRedirectResponseCodeEnum enum to its proto representation.
func NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnumToProto(e *networkservices.HttpRouteRulesActionRedirectResponseCodeEnum) networkservicespb.NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum {
	if e == nil {
		return networkservicespb.NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum(0)
	}
	if v, ok := networkservicespb.NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum_value["HttpRouteRulesActionRedirectResponseCodeEnum"+string(*e)]; ok {
		return networkservicespb.NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum(v)
	}
	return networkservicespb.NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnum(0)
}

// HttpRouteRulesToProto converts a HttpRouteRules object to its proto representation.
func NetworkservicesHttpRouteRulesToProto(o *networkservices.HttpRouteRules) *networkservicespb.NetworkservicesHttpRouteRules {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRules{}
	p.SetAction(NetworkservicesHttpRouteRulesActionToProto(o.Action))
	sMatches := make([]*networkservicespb.NetworkservicesHttpRouteRulesMatches, len(o.Matches))
	for i, r := range o.Matches {
		sMatches[i] = NetworkservicesHttpRouteRulesMatchesToProto(&r)
	}
	p.SetMatches(sMatches)
	return p
}

// HttpRouteRulesMatchesToProto converts a HttpRouteRulesMatches object to its proto representation.
func NetworkservicesHttpRouteRulesMatchesToProto(o *networkservices.HttpRouteRulesMatches) *networkservicespb.NetworkservicesHttpRouteRulesMatches {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesMatches{}
	p.SetFullPathMatch(dcl.ValueOrEmptyString(o.FullPathMatch))
	p.SetPrefixMatch(dcl.ValueOrEmptyString(o.PrefixMatch))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	p.SetIgnoreCase(dcl.ValueOrEmptyBool(o.IgnoreCase))
	sHeaders := make([]*networkservicespb.NetworkservicesHttpRouteRulesMatchesHeaders, len(o.Headers))
	for i, r := range o.Headers {
		sHeaders[i] = NetworkservicesHttpRouteRulesMatchesHeadersToProto(&r)
	}
	p.SetHeaders(sHeaders)
	sQueryParameters := make([]*networkservicespb.NetworkservicesHttpRouteRulesMatchesQueryParameters, len(o.QueryParameters))
	for i, r := range o.QueryParameters {
		sQueryParameters[i] = NetworkservicesHttpRouteRulesMatchesQueryParametersToProto(&r)
	}
	p.SetQueryParameters(sQueryParameters)
	return p
}

// HttpRouteRulesMatchesHeadersToProto converts a HttpRouteRulesMatchesHeaders object to its proto representation.
func NetworkservicesHttpRouteRulesMatchesHeadersToProto(o *networkservices.HttpRouteRulesMatchesHeaders) *networkservicespb.NetworkservicesHttpRouteRulesMatchesHeaders {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesMatchesHeaders{}
	p.SetHeader(dcl.ValueOrEmptyString(o.Header))
	p.SetExactMatch(dcl.ValueOrEmptyString(o.ExactMatch))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	p.SetPrefixMatch(dcl.ValueOrEmptyString(o.PrefixMatch))
	p.SetPresentMatch(dcl.ValueOrEmptyBool(o.PresentMatch))
	p.SetSuffixMatch(dcl.ValueOrEmptyString(o.SuffixMatch))
	p.SetRangeMatch(NetworkservicesHttpRouteRulesMatchesHeadersRangeMatchToProto(o.RangeMatch))
	p.SetInvertMatch(dcl.ValueOrEmptyBool(o.InvertMatch))
	return p
}

// HttpRouteRulesMatchesHeadersRangeMatchToProto converts a HttpRouteRulesMatchesHeadersRangeMatch object to its proto representation.
func NetworkservicesHttpRouteRulesMatchesHeadersRangeMatchToProto(o *networkservices.HttpRouteRulesMatchesHeadersRangeMatch) *networkservicespb.NetworkservicesHttpRouteRulesMatchesHeadersRangeMatch {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesMatchesHeadersRangeMatch{}
	p.SetStart(dcl.ValueOrEmptyInt64(o.Start))
	p.SetEnd(dcl.ValueOrEmptyInt64(o.End))
	return p
}

// HttpRouteRulesMatchesQueryParametersToProto converts a HttpRouteRulesMatchesQueryParameters object to its proto representation.
func NetworkservicesHttpRouteRulesMatchesQueryParametersToProto(o *networkservices.HttpRouteRulesMatchesQueryParameters) *networkservicespb.NetworkservicesHttpRouteRulesMatchesQueryParameters {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesMatchesQueryParameters{}
	p.SetQueryParameter(dcl.ValueOrEmptyString(o.QueryParameter))
	p.SetExactMatch(dcl.ValueOrEmptyString(o.ExactMatch))
	p.SetRegexMatch(dcl.ValueOrEmptyString(o.RegexMatch))
	p.SetPresentMatch(dcl.ValueOrEmptyBool(o.PresentMatch))
	return p
}

// HttpRouteRulesActionToProto converts a HttpRouteRulesAction object to its proto representation.
func NetworkservicesHttpRouteRulesActionToProto(o *networkservices.HttpRouteRulesAction) *networkservicespb.NetworkservicesHttpRouteRulesAction {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesAction{}
	p.SetRedirect(NetworkservicesHttpRouteRulesActionRedirectToProto(o.Redirect))
	p.SetFaultInjectionPolicy(NetworkservicesHttpRouteRulesActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy))
	p.SetRequestHeaderModifier(NetworkservicesHttpRouteRulesActionRequestHeaderModifierToProto(o.RequestHeaderModifier))
	p.SetResponseHeaderModifier(NetworkservicesHttpRouteRulesActionResponseHeaderModifierToProto(o.ResponseHeaderModifier))
	p.SetUrlRewrite(NetworkservicesHttpRouteRulesActionUrlRewriteToProto(o.UrlRewrite))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetRetryPolicy(NetworkservicesHttpRouteRulesActionRetryPolicyToProto(o.RetryPolicy))
	p.SetRequestMirrorPolicy(NetworkservicesHttpRouteRulesActionRequestMirrorPolicyToProto(o.RequestMirrorPolicy))
	p.SetCorsPolicy(NetworkservicesHttpRouteRulesActionCorsPolicyToProto(o.CorsPolicy))
	sDestinations := make([]*networkservicespb.NetworkservicesHttpRouteRulesActionDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworkservicesHttpRouteRulesActionDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// HttpRouteRulesActionDestinationsToProto converts a HttpRouteRulesActionDestinations object to its proto representation.
func NetworkservicesHttpRouteRulesActionDestinationsToProto(o *networkservices.HttpRouteRulesActionDestinations) *networkservicespb.NetworkservicesHttpRouteRulesActionDestinations {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionDestinations{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// HttpRouteRulesActionRedirectToProto converts a HttpRouteRulesActionRedirect object to its proto representation.
func NetworkservicesHttpRouteRulesActionRedirectToProto(o *networkservices.HttpRouteRulesActionRedirect) *networkservicespb.NetworkservicesHttpRouteRulesActionRedirect {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionRedirect{}
	p.SetHostRedirect(dcl.ValueOrEmptyString(o.HostRedirect))
	p.SetPathRedirect(dcl.ValueOrEmptyString(o.PathRedirect))
	p.SetPrefixRewrite(dcl.ValueOrEmptyString(o.PrefixRewrite))
	p.SetResponseCode(NetworkservicesHttpRouteRulesActionRedirectResponseCodeEnumToProto(o.ResponseCode))
	p.SetHttpsRedirect(dcl.ValueOrEmptyBool(o.HttpsRedirect))
	p.SetStripQuery(dcl.ValueOrEmptyBool(o.StripQuery))
	p.SetPortRedirect(dcl.ValueOrEmptyInt64(o.PortRedirect))
	return p
}

// HttpRouteRulesActionFaultInjectionPolicyToProto converts a HttpRouteRulesActionFaultInjectionPolicy object to its proto representation.
func NetworkservicesHttpRouteRulesActionFaultInjectionPolicyToProto(o *networkservices.HttpRouteRulesActionFaultInjectionPolicy) *networkservicespb.NetworkservicesHttpRouteRulesActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionFaultInjectionPolicy{}
	p.SetDelay(NetworkservicesHttpRouteRulesActionFaultInjectionPolicyDelayToProto(o.Delay))
	p.SetAbort(NetworkservicesHttpRouteRulesActionFaultInjectionPolicyAbortToProto(o.Abort))
	return p
}

// HttpRouteRulesActionFaultInjectionPolicyDelayToProto converts a HttpRouteRulesActionFaultInjectionPolicyDelay object to its proto representation.
func NetworkservicesHttpRouteRulesActionFaultInjectionPolicyDelayToProto(o *networkservices.HttpRouteRulesActionFaultInjectionPolicyDelay) *networkservicespb.NetworkservicesHttpRouteRulesActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionFaultInjectionPolicyDelay{}
	p.SetFixedDelay(dcl.ValueOrEmptyString(o.FixedDelay))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// HttpRouteRulesActionFaultInjectionPolicyAbortToProto converts a HttpRouteRulesActionFaultInjectionPolicyAbort object to its proto representation.
func NetworkservicesHttpRouteRulesActionFaultInjectionPolicyAbortToProto(o *networkservices.HttpRouteRulesActionFaultInjectionPolicyAbort) *networkservicespb.NetworkservicesHttpRouteRulesActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionFaultInjectionPolicyAbort{}
	p.SetHttpStatus(dcl.ValueOrEmptyInt64(o.HttpStatus))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// HttpRouteRulesActionRequestHeaderModifierToProto converts a HttpRouteRulesActionRequestHeaderModifier object to its proto representation.
func NetworkservicesHttpRouteRulesActionRequestHeaderModifierToProto(o *networkservices.HttpRouteRulesActionRequestHeaderModifier) *networkservicespb.NetworkservicesHttpRouteRulesActionRequestHeaderModifier {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionRequestHeaderModifier{}
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
func NetworkservicesHttpRouteRulesActionResponseHeaderModifierToProto(o *networkservices.HttpRouteRulesActionResponseHeaderModifier) *networkservicespb.NetworkservicesHttpRouteRulesActionResponseHeaderModifier {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionResponseHeaderModifier{}
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
func NetworkservicesHttpRouteRulesActionUrlRewriteToProto(o *networkservices.HttpRouteRulesActionUrlRewrite) *networkservicespb.NetworkservicesHttpRouteRulesActionUrlRewrite {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionUrlRewrite{}
	p.SetPathPrefixRewrite(dcl.ValueOrEmptyString(o.PathPrefixRewrite))
	p.SetHostRewrite(dcl.ValueOrEmptyString(o.HostRewrite))
	return p
}

// HttpRouteRulesActionRetryPolicyToProto converts a HttpRouteRulesActionRetryPolicy object to its proto representation.
func NetworkservicesHttpRouteRulesActionRetryPolicyToProto(o *networkservices.HttpRouteRulesActionRetryPolicy) *networkservicespb.NetworkservicesHttpRouteRulesActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionRetryPolicy{}
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
func NetworkservicesHttpRouteRulesActionRequestMirrorPolicyToProto(o *networkservices.HttpRouteRulesActionRequestMirrorPolicy) *networkservicespb.NetworkservicesHttpRouteRulesActionRequestMirrorPolicy {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionRequestMirrorPolicy{}
	p.SetDestination(NetworkservicesHttpRouteRulesActionRequestMirrorPolicyDestinationToProto(o.Destination))
	return p
}

// HttpRouteRulesActionRequestMirrorPolicyDestinationToProto converts a HttpRouteRulesActionRequestMirrorPolicyDestination object to its proto representation.
func NetworkservicesHttpRouteRulesActionRequestMirrorPolicyDestinationToProto(o *networkservices.HttpRouteRulesActionRequestMirrorPolicyDestination) *networkservicespb.NetworkservicesHttpRouteRulesActionRequestMirrorPolicyDestination {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionRequestMirrorPolicyDestination{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// HttpRouteRulesActionCorsPolicyToProto converts a HttpRouteRulesActionCorsPolicy object to its proto representation.
func NetworkservicesHttpRouteRulesActionCorsPolicyToProto(o *networkservices.HttpRouteRulesActionCorsPolicy) *networkservicespb.NetworkservicesHttpRouteRulesActionCorsPolicy {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesHttpRouteRulesActionCorsPolicy{}
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
func HttpRouteToProto(resource *networkservices.HttpRoute) *networkservicespb.NetworkservicesHttpRoute {
	p := &networkservicespb.NetworkservicesHttpRoute{}
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
	sRules := make([]*networkservicespb.NetworkservicesHttpRouteRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworkservicesHttpRouteRulesToProto(&r)
	}
	p.SetRules(sRules)

	return p
}

// applyHttpRoute handles the gRPC request by passing it to the underlying HttpRoute Apply() method.
func (s *HttpRouteServer) applyHttpRoute(ctx context.Context, c *networkservices.Client, request *networkservicespb.ApplyNetworkservicesHttpRouteRequest) (*networkservicespb.NetworkservicesHttpRoute, error) {
	p := ProtoToHttpRoute(request.GetResource())
	res, err := c.ApplyHttpRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HttpRouteToProto(res)
	return r, nil
}

// applyNetworkservicesHttpRoute handles the gRPC request by passing it to the underlying HttpRoute Apply() method.
func (s *HttpRouteServer) ApplyNetworkservicesHttpRoute(ctx context.Context, request *networkservicespb.ApplyNetworkservicesHttpRouteRequest) (*networkservicespb.NetworkservicesHttpRoute, error) {
	cl, err := createConfigHttpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyHttpRoute(ctx, cl, request)
}

// DeleteHttpRoute handles the gRPC request by passing it to the underlying HttpRoute Delete() method.
func (s *HttpRouteServer) DeleteNetworkservicesHttpRoute(ctx context.Context, request *networkservicespb.DeleteNetworkservicesHttpRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHttpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHttpRoute(ctx, ProtoToHttpRoute(request.GetResource()))

}

// ListNetworkservicesHttpRoute handles the gRPC request by passing it to the underlying HttpRouteList() method.
func (s *HttpRouteServer) ListNetworkservicesHttpRoute(ctx context.Context, request *networkservicespb.ListNetworkservicesHttpRouteRequest) (*networkservicespb.ListNetworkservicesHttpRouteResponse, error) {
	cl, err := createConfigHttpRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHttpRoute(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*networkservicespb.NetworkservicesHttpRoute
	for _, r := range resources.Items {
		rp := HttpRouteToProto(r)
		protos = append(protos, rp)
	}
	p := &networkservicespb.ListNetworkservicesHttpRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigHttpRoute(ctx context.Context, service_account_file string) (*networkservices.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return networkservices.NewClient(conf), nil
}
