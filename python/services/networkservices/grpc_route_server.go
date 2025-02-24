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

// GrpcRouteServer implements the gRPC interface for GrpcRoute.
type GrpcRouteServer struct{}

// ProtoToGrpcRouteRulesMatchesMethodTypeEnum converts a GrpcRouteRulesMatchesMethodTypeEnum enum from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesMatchesMethodTypeEnum(e networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethodTypeEnum) *networkservices.GrpcRouteRulesMatchesMethodTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethodTypeEnum_name[int32(e)]; ok {
		e := networkservices.GrpcRouteRulesMatchesMethodTypeEnum(n[len("NetworkservicesGrpcRouteRulesMatchesMethodTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGrpcRouteRulesMatchesHeadersTypeEnum converts a GrpcRouteRulesMatchesHeadersTypeEnum enum from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum(e networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum) *networkservices.GrpcRouteRulesMatchesHeadersTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum_name[int32(e)]; ok {
		e := networkservices.GrpcRouteRulesMatchesHeadersTypeEnum(n[len("NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGrpcRouteRules converts a GrpcRouteRules object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRules(p *networkservicespb.NetworkservicesGrpcRouteRules) *networkservices.GrpcRouteRules {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRules{
		Action: ProtoToNetworkservicesGrpcRouteRulesAction(p.GetAction()),
	}
	for _, r := range p.GetMatches() {
		obj.Matches = append(obj.Matches, *ProtoToNetworkservicesGrpcRouteRulesMatches(r))
	}
	return obj
}

// ProtoToGrpcRouteRulesMatches converts a GrpcRouteRulesMatches object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesMatches(p *networkservicespb.NetworkservicesGrpcRouteRulesMatches) *networkservices.GrpcRouteRulesMatches {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRulesMatches{
		Method: ProtoToNetworkservicesGrpcRouteRulesMatchesMethod(p.GetMethod()),
	}
	for _, r := range p.GetHeaders() {
		obj.Headers = append(obj.Headers, *ProtoToNetworkservicesGrpcRouteRulesMatchesHeaders(r))
	}
	return obj
}

// ProtoToGrpcRouteRulesMatchesMethod converts a GrpcRouteRulesMatchesMethod object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesMatchesMethod(p *networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethod) *networkservices.GrpcRouteRulesMatchesMethod {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRulesMatchesMethod{
		Type:          ProtoToNetworkservicesGrpcRouteRulesMatchesMethodTypeEnum(p.GetType()),
		GrpcService:   dcl.StringOrNil(p.GetGrpcService()),
		GrpcMethod:    dcl.StringOrNil(p.GetGrpcMethod()),
		CaseSensitive: dcl.Bool(p.GetCaseSensitive()),
	}
	return obj
}

// ProtoToGrpcRouteRulesMatchesHeaders converts a GrpcRouteRulesMatchesHeaders object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesMatchesHeaders(p *networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeaders) *networkservices.GrpcRouteRulesMatchesHeaders {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRulesMatchesHeaders{
		Type:  ProtoToNetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum(p.GetType()),
		Key:   dcl.StringOrNil(p.GetKey()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToGrpcRouteRulesAction converts a GrpcRouteRulesAction object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesAction(p *networkservicespb.NetworkservicesGrpcRouteRulesAction) *networkservices.GrpcRouteRulesAction {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRulesAction{
		FaultInjectionPolicy: ProtoToNetworkservicesGrpcRouteRulesActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
		Timeout:              dcl.StringOrNil(p.GetTimeout()),
		RetryPolicy:          ProtoToNetworkservicesGrpcRouteRulesActionRetryPolicy(p.GetRetryPolicy()),
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworkservicesGrpcRouteRulesActionDestinations(r))
	}
	return obj
}

// ProtoToGrpcRouteRulesActionDestinations converts a GrpcRouteRulesActionDestinations object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesActionDestinations(p *networkservicespb.NetworkservicesGrpcRouteRulesActionDestinations) *networkservices.GrpcRouteRulesActionDestinations {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRulesActionDestinations{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionFaultInjectionPolicy converts a GrpcRouteRulesActionFaultInjectionPolicy object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesActionFaultInjectionPolicy(p *networkservicespb.NetworkservicesGrpcRouteRulesActionFaultInjectionPolicy) *networkservices.GrpcRouteRulesActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRulesActionFaultInjectionPolicy{
		Delay: ProtoToNetworkservicesGrpcRouteRulesActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToNetworkservicesGrpcRouteRulesActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionFaultInjectionPolicyDelay converts a GrpcRouteRulesActionFaultInjectionPolicyDelay object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesActionFaultInjectionPolicyDelay(p *networkservicespb.NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyDelay) *networkservices.GrpcRouteRulesActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRulesActionFaultInjectionPolicyDelay{
		FixedDelay: dcl.StringOrNil(p.GetFixedDelay()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionFaultInjectionPolicyAbort converts a GrpcRouteRulesActionFaultInjectionPolicyAbort object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesActionFaultInjectionPolicyAbort(p *networkservicespb.NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyAbort) *networkservices.GrpcRouteRulesActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRulesActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.GetHttpStatus()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionRetryPolicy converts a GrpcRouteRulesActionRetryPolicy object from its proto representation.
func ProtoToNetworkservicesGrpcRouteRulesActionRetryPolicy(p *networkservicespb.NetworkservicesGrpcRouteRulesActionRetryPolicy) *networkservices.GrpcRouteRulesActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &networkservices.GrpcRouteRulesActionRetryPolicy{
		NumRetries: dcl.Int64OrNil(p.GetNumRetries()),
	}
	for _, r := range p.GetRetryConditions() {
		obj.RetryConditions = append(obj.RetryConditions, r)
	}
	return obj
}

// ProtoToGrpcRoute converts a GrpcRoute resource from its proto representation.
func ProtoToGrpcRoute(p *networkservicespb.NetworkservicesGrpcRoute) *networkservices.GrpcRoute {
	obj := &networkservices.GrpcRoute{
		Name:        dcl.StringOrNil(p.GetName()),
		CreateTime:  dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:  dcl.StringOrNil(p.GetUpdateTime()),
		Description: dcl.StringOrNil(p.GetDescription()),
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
		obj.Rules = append(obj.Rules, *ProtoToNetworkservicesGrpcRouteRules(r))
	}
	return obj
}

// GrpcRouteRulesMatchesMethodTypeEnumToProto converts a GrpcRouteRulesMatchesMethodTypeEnum enum to its proto representation.
func NetworkservicesGrpcRouteRulesMatchesMethodTypeEnumToProto(e *networkservices.GrpcRouteRulesMatchesMethodTypeEnum) networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethodTypeEnum {
	if e == nil {
		return networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethodTypeEnum(0)
	}
	if v, ok := networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethodTypeEnum_value["GrpcRouteRulesMatchesMethodTypeEnum"+string(*e)]; ok {
		return networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethodTypeEnum(v)
	}
	return networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethodTypeEnum(0)
}

// GrpcRouteRulesMatchesHeadersTypeEnumToProto converts a GrpcRouteRulesMatchesHeadersTypeEnum enum to its proto representation.
func NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnumToProto(e *networkservices.GrpcRouteRulesMatchesHeadersTypeEnum) networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum {
	if e == nil {
		return networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum(0)
	}
	if v, ok := networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum_value["GrpcRouteRulesMatchesHeadersTypeEnum"+string(*e)]; ok {
		return networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum(v)
	}
	return networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnum(0)
}

// GrpcRouteRulesToProto converts a GrpcRouteRules object to its proto representation.
func NetworkservicesGrpcRouteRulesToProto(o *networkservices.GrpcRouteRules) *networkservicespb.NetworkservicesGrpcRouteRules {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRules{}
	p.SetAction(NetworkservicesGrpcRouteRulesActionToProto(o.Action))
	sMatches := make([]*networkservicespb.NetworkservicesGrpcRouteRulesMatches, len(o.Matches))
	for i, r := range o.Matches {
		sMatches[i] = NetworkservicesGrpcRouteRulesMatchesToProto(&r)
	}
	p.SetMatches(sMatches)
	return p
}

// GrpcRouteRulesMatchesToProto converts a GrpcRouteRulesMatches object to its proto representation.
func NetworkservicesGrpcRouteRulesMatchesToProto(o *networkservices.GrpcRouteRulesMatches) *networkservicespb.NetworkservicesGrpcRouteRulesMatches {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRulesMatches{}
	p.SetMethod(NetworkservicesGrpcRouteRulesMatchesMethodToProto(o.Method))
	sHeaders := make([]*networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeaders, len(o.Headers))
	for i, r := range o.Headers {
		sHeaders[i] = NetworkservicesGrpcRouteRulesMatchesHeadersToProto(&r)
	}
	p.SetHeaders(sHeaders)
	return p
}

// GrpcRouteRulesMatchesMethodToProto converts a GrpcRouteRulesMatchesMethod object to its proto representation.
func NetworkservicesGrpcRouteRulesMatchesMethodToProto(o *networkservices.GrpcRouteRulesMatchesMethod) *networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethod {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRulesMatchesMethod{}
	p.SetType(NetworkservicesGrpcRouteRulesMatchesMethodTypeEnumToProto(o.Type))
	p.SetGrpcService(dcl.ValueOrEmptyString(o.GrpcService))
	p.SetGrpcMethod(dcl.ValueOrEmptyString(o.GrpcMethod))
	p.SetCaseSensitive(dcl.ValueOrEmptyBool(o.CaseSensitive))
	return p
}

// GrpcRouteRulesMatchesHeadersToProto converts a GrpcRouteRulesMatchesHeaders object to its proto representation.
func NetworkservicesGrpcRouteRulesMatchesHeadersToProto(o *networkservices.GrpcRouteRulesMatchesHeaders) *networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeaders {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRulesMatchesHeaders{}
	p.SetType(NetworkservicesGrpcRouteRulesMatchesHeadersTypeEnumToProto(o.Type))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// GrpcRouteRulesActionToProto converts a GrpcRouteRulesAction object to its proto representation.
func NetworkservicesGrpcRouteRulesActionToProto(o *networkservices.GrpcRouteRulesAction) *networkservicespb.NetworkservicesGrpcRouteRulesAction {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRulesAction{}
	p.SetFaultInjectionPolicy(NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetRetryPolicy(NetworkservicesGrpcRouteRulesActionRetryPolicyToProto(o.RetryPolicy))
	sDestinations := make([]*networkservicespb.NetworkservicesGrpcRouteRulesActionDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworkservicesGrpcRouteRulesActionDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// GrpcRouteRulesActionDestinationsToProto converts a GrpcRouteRulesActionDestinations object to its proto representation.
func NetworkservicesGrpcRouteRulesActionDestinationsToProto(o *networkservices.GrpcRouteRulesActionDestinations) *networkservicespb.NetworkservicesGrpcRouteRulesActionDestinations {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRulesActionDestinations{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// GrpcRouteRulesActionFaultInjectionPolicyToProto converts a GrpcRouteRulesActionFaultInjectionPolicy object to its proto representation.
func NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyToProto(o *networkservices.GrpcRouteRulesActionFaultInjectionPolicy) *networkservicespb.NetworkservicesGrpcRouteRulesActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRulesActionFaultInjectionPolicy{}
	p.SetDelay(NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyDelayToProto(o.Delay))
	p.SetAbort(NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyAbortToProto(o.Abort))
	return p
}

// GrpcRouteRulesActionFaultInjectionPolicyDelayToProto converts a GrpcRouteRulesActionFaultInjectionPolicyDelay object to its proto representation.
func NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyDelayToProto(o *networkservices.GrpcRouteRulesActionFaultInjectionPolicyDelay) *networkservicespb.NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyDelay{}
	p.SetFixedDelay(dcl.ValueOrEmptyString(o.FixedDelay))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// GrpcRouteRulesActionFaultInjectionPolicyAbortToProto converts a GrpcRouteRulesActionFaultInjectionPolicyAbort object to its proto representation.
func NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyAbortToProto(o *networkservices.GrpcRouteRulesActionFaultInjectionPolicyAbort) *networkservicespb.NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRulesActionFaultInjectionPolicyAbort{}
	p.SetHttpStatus(dcl.ValueOrEmptyInt64(o.HttpStatus))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// GrpcRouteRulesActionRetryPolicyToProto converts a GrpcRouteRulesActionRetryPolicy object to its proto representation.
func NetworkservicesGrpcRouteRulesActionRetryPolicyToProto(o *networkservices.GrpcRouteRulesActionRetryPolicy) *networkservicespb.NetworkservicesGrpcRouteRulesActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &networkservicespb.NetworkservicesGrpcRouteRulesActionRetryPolicy{}
	p.SetNumRetries(dcl.ValueOrEmptyInt64(o.NumRetries))
	sRetryConditions := make([]string, len(o.RetryConditions))
	for i, r := range o.RetryConditions {
		sRetryConditions[i] = r
	}
	p.SetRetryConditions(sRetryConditions)
	return p
}

// GrpcRouteToProto converts a GrpcRoute resource to its proto representation.
func GrpcRouteToProto(resource *networkservices.GrpcRoute) *networkservicespb.NetworkservicesGrpcRoute {
	p := &networkservicespb.NetworkservicesGrpcRoute{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
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
	sRules := make([]*networkservicespb.NetworkservicesGrpcRouteRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworkservicesGrpcRouteRulesToProto(&r)
	}
	p.SetRules(sRules)

	return p
}

// applyGrpcRoute handles the gRPC request by passing it to the underlying GrpcRoute Apply() method.
func (s *GrpcRouteServer) applyGrpcRoute(ctx context.Context, c *networkservices.Client, request *networkservicespb.ApplyNetworkservicesGrpcRouteRequest) (*networkservicespb.NetworkservicesGrpcRoute, error) {
	p := ProtoToGrpcRoute(request.GetResource())
	res, err := c.ApplyGrpcRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GrpcRouteToProto(res)
	return r, nil
}

// applyNetworkservicesGrpcRoute handles the gRPC request by passing it to the underlying GrpcRoute Apply() method.
func (s *GrpcRouteServer) ApplyNetworkservicesGrpcRoute(ctx context.Context, request *networkservicespb.ApplyNetworkservicesGrpcRouteRequest) (*networkservicespb.NetworkservicesGrpcRoute, error) {
	cl, err := createConfigGrpcRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGrpcRoute(ctx, cl, request)
}

// DeleteGrpcRoute handles the gRPC request by passing it to the underlying GrpcRoute Delete() method.
func (s *GrpcRouteServer) DeleteNetworkservicesGrpcRoute(ctx context.Context, request *networkservicespb.DeleteNetworkservicesGrpcRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGrpcRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGrpcRoute(ctx, ProtoToGrpcRoute(request.GetResource()))

}

// ListNetworkservicesGrpcRoute handles the gRPC request by passing it to the underlying GrpcRouteList() method.
func (s *GrpcRouteServer) ListNetworkservicesGrpcRoute(ctx context.Context, request *networkservicespb.ListNetworkservicesGrpcRouteRequest) (*networkservicespb.ListNetworkservicesGrpcRouteResponse, error) {
	cl, err := createConfigGrpcRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGrpcRoute(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*networkservicespb.NetworkservicesGrpcRoute
	for _, r := range resources.Items {
		rp := GrpcRouteToProto(r)
		protos = append(protos, rp)
	}
	p := &networkservicespb.ListNetworkservicesGrpcRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGrpcRoute(ctx context.Context, service_account_file string) (*networkservices.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return networkservices.NewClient(conf), nil
}
