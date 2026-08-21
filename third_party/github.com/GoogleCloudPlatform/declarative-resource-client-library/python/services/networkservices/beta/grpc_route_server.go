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

// GrpcRouteServer implements the gRPC interface for GrpcRoute.
type GrpcRouteServer struct{}

// ProtoToGrpcRouteRulesMatchesMethodTypeEnum converts a GrpcRouteRulesMatchesMethodTypeEnum enum from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum(e betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum) *beta.GrpcRouteRulesMatchesMethodTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum_name[int32(e)]; ok {
		e := beta.GrpcRouteRulesMatchesMethodTypeEnum(n[len("NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGrpcRouteRulesMatchesHeadersTypeEnum converts a GrpcRouteRulesMatchesHeadersTypeEnum enum from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum(e betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum) *beta.GrpcRouteRulesMatchesHeadersTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum_name[int32(e)]; ok {
		e := beta.GrpcRouteRulesMatchesHeadersTypeEnum(n[len("NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGrpcRouteRules converts a GrpcRouteRules object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRules(p *betapb.NetworkservicesBetaGrpcRouteRules) *beta.GrpcRouteRules {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRules{
		Action: ProtoToNetworkservicesBetaGrpcRouteRulesAction(p.GetAction()),
	}
	for _, r := range p.GetMatches() {
		obj.Matches = append(obj.Matches, *ProtoToNetworkservicesBetaGrpcRouteRulesMatches(r))
	}
	return obj
}

// ProtoToGrpcRouteRulesMatches converts a GrpcRouteRulesMatches object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesMatches(p *betapb.NetworkservicesBetaGrpcRouteRulesMatches) *beta.GrpcRouteRulesMatches {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRulesMatches{
		Method: ProtoToNetworkservicesBetaGrpcRouteRulesMatchesMethod(p.GetMethod()),
	}
	for _, r := range p.GetHeaders() {
		obj.Headers = append(obj.Headers, *ProtoToNetworkservicesBetaGrpcRouteRulesMatchesHeaders(r))
	}
	return obj
}

// ProtoToGrpcRouteRulesMatchesMethod converts a GrpcRouteRulesMatchesMethod object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesMatchesMethod(p *betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethod) *beta.GrpcRouteRulesMatchesMethod {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRulesMatchesMethod{
		Type:          ProtoToNetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum(p.GetType()),
		GrpcService:   dcl.StringOrNil(p.GetGrpcService()),
		GrpcMethod:    dcl.StringOrNil(p.GetGrpcMethod()),
		CaseSensitive: dcl.Bool(p.GetCaseSensitive()),
	}
	return obj
}

// ProtoToGrpcRouteRulesMatchesHeaders converts a GrpcRouteRulesMatchesHeaders object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesMatchesHeaders(p *betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeaders) *beta.GrpcRouteRulesMatchesHeaders {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRulesMatchesHeaders{
		Type:  ProtoToNetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum(p.GetType()),
		Key:   dcl.StringOrNil(p.GetKey()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToGrpcRouteRulesAction converts a GrpcRouteRulesAction object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesAction(p *betapb.NetworkservicesBetaGrpcRouteRulesAction) *beta.GrpcRouteRulesAction {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRulesAction{
		FaultInjectionPolicy: ProtoToNetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
		Timeout:              dcl.StringOrNil(p.GetTimeout()),
		RetryPolicy:          ProtoToNetworkservicesBetaGrpcRouteRulesActionRetryPolicy(p.GetRetryPolicy()),
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworkservicesBetaGrpcRouteRulesActionDestinations(r))
	}
	return obj
}

// ProtoToGrpcRouteRulesActionDestinations converts a GrpcRouteRulesActionDestinations object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesActionDestinations(p *betapb.NetworkservicesBetaGrpcRouteRulesActionDestinations) *beta.GrpcRouteRulesActionDestinations {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRulesActionDestinations{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionFaultInjectionPolicy converts a GrpcRouteRulesActionFaultInjectionPolicy object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicy(p *betapb.NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicy) *beta.GrpcRouteRulesActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRulesActionFaultInjectionPolicy{
		Delay: ProtoToNetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToNetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionFaultInjectionPolicyDelay converts a GrpcRouteRulesActionFaultInjectionPolicyDelay object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyDelay(p *betapb.NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyDelay) *beta.GrpcRouteRulesActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRulesActionFaultInjectionPolicyDelay{
		FixedDelay: dcl.StringOrNil(p.GetFixedDelay()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionFaultInjectionPolicyAbort converts a GrpcRouteRulesActionFaultInjectionPolicyAbort object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyAbort(p *betapb.NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyAbort) *beta.GrpcRouteRulesActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRulesActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.GetHttpStatus()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionRetryPolicy converts a GrpcRouteRulesActionRetryPolicy object from its proto representation.
func ProtoToNetworkservicesBetaGrpcRouteRulesActionRetryPolicy(p *betapb.NetworkservicesBetaGrpcRouteRulesActionRetryPolicy) *beta.GrpcRouteRulesActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &beta.GrpcRouteRulesActionRetryPolicy{
		NumRetries: dcl.Int64OrNil(p.GetNumRetries()),
	}
	for _, r := range p.GetRetryConditions() {
		obj.RetryConditions = append(obj.RetryConditions, r)
	}
	return obj
}

// ProtoToGrpcRoute converts a GrpcRoute resource from its proto representation.
func ProtoToGrpcRoute(p *betapb.NetworkservicesBetaGrpcRoute) *beta.GrpcRoute {
	obj := &beta.GrpcRoute{
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
		obj.Rules = append(obj.Rules, *ProtoToNetworkservicesBetaGrpcRouteRules(r))
	}
	return obj
}

// GrpcRouteRulesMatchesMethodTypeEnumToProto converts a GrpcRouteRulesMatchesMethodTypeEnum enum to its proto representation.
func NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnumToProto(e *beta.GrpcRouteRulesMatchesMethodTypeEnum) betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum {
	if e == nil {
		return betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum(0)
	}
	if v, ok := betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum_value["GrpcRouteRulesMatchesMethodTypeEnum"+string(*e)]; ok {
		return betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum(v)
	}
	return betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnum(0)
}

// GrpcRouteRulesMatchesHeadersTypeEnumToProto converts a GrpcRouteRulesMatchesHeadersTypeEnum enum to its proto representation.
func NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnumToProto(e *beta.GrpcRouteRulesMatchesHeadersTypeEnum) betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum {
	if e == nil {
		return betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum(0)
	}
	if v, ok := betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum_value["GrpcRouteRulesMatchesHeadersTypeEnum"+string(*e)]; ok {
		return betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum(v)
	}
	return betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnum(0)
}

// GrpcRouteRulesToProto converts a GrpcRouteRules object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesToProto(o *beta.GrpcRouteRules) *betapb.NetworkservicesBetaGrpcRouteRules {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRules{}
	p.SetAction(NetworkservicesBetaGrpcRouteRulesActionToProto(o.Action))
	sMatches := make([]*betapb.NetworkservicesBetaGrpcRouteRulesMatches, len(o.Matches))
	for i, r := range o.Matches {
		sMatches[i] = NetworkservicesBetaGrpcRouteRulesMatchesToProto(&r)
	}
	p.SetMatches(sMatches)
	return p
}

// GrpcRouteRulesMatchesToProto converts a GrpcRouteRulesMatches object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesMatchesToProto(o *beta.GrpcRouteRulesMatches) *betapb.NetworkservicesBetaGrpcRouteRulesMatches {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRulesMatches{}
	p.SetMethod(NetworkservicesBetaGrpcRouteRulesMatchesMethodToProto(o.Method))
	sHeaders := make([]*betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeaders, len(o.Headers))
	for i, r := range o.Headers {
		sHeaders[i] = NetworkservicesBetaGrpcRouteRulesMatchesHeadersToProto(&r)
	}
	p.SetHeaders(sHeaders)
	return p
}

// GrpcRouteRulesMatchesMethodToProto converts a GrpcRouteRulesMatchesMethod object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesMatchesMethodToProto(o *beta.GrpcRouteRulesMatchesMethod) *betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethod {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRulesMatchesMethod{}
	p.SetType(NetworkservicesBetaGrpcRouteRulesMatchesMethodTypeEnumToProto(o.Type))
	p.SetGrpcService(dcl.ValueOrEmptyString(o.GrpcService))
	p.SetGrpcMethod(dcl.ValueOrEmptyString(o.GrpcMethod))
	p.SetCaseSensitive(dcl.ValueOrEmptyBool(o.CaseSensitive))
	return p
}

// GrpcRouteRulesMatchesHeadersToProto converts a GrpcRouteRulesMatchesHeaders object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesMatchesHeadersToProto(o *beta.GrpcRouteRulesMatchesHeaders) *betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeaders {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRulesMatchesHeaders{}
	p.SetType(NetworkservicesBetaGrpcRouteRulesMatchesHeadersTypeEnumToProto(o.Type))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// GrpcRouteRulesActionToProto converts a GrpcRouteRulesAction object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesActionToProto(o *beta.GrpcRouteRulesAction) *betapb.NetworkservicesBetaGrpcRouteRulesAction {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRulesAction{}
	p.SetFaultInjectionPolicy(NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetRetryPolicy(NetworkservicesBetaGrpcRouteRulesActionRetryPolicyToProto(o.RetryPolicy))
	sDestinations := make([]*betapb.NetworkservicesBetaGrpcRouteRulesActionDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworkservicesBetaGrpcRouteRulesActionDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// GrpcRouteRulesActionDestinationsToProto converts a GrpcRouteRulesActionDestinations object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesActionDestinationsToProto(o *beta.GrpcRouteRulesActionDestinations) *betapb.NetworkservicesBetaGrpcRouteRulesActionDestinations {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRulesActionDestinations{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// GrpcRouteRulesActionFaultInjectionPolicyToProto converts a GrpcRouteRulesActionFaultInjectionPolicy object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyToProto(o *beta.GrpcRouteRulesActionFaultInjectionPolicy) *betapb.NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicy{}
	p.SetDelay(NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyDelayToProto(o.Delay))
	p.SetAbort(NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyAbortToProto(o.Abort))
	return p
}

// GrpcRouteRulesActionFaultInjectionPolicyDelayToProto converts a GrpcRouteRulesActionFaultInjectionPolicyDelay object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyDelayToProto(o *beta.GrpcRouteRulesActionFaultInjectionPolicyDelay) *betapb.NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyDelay{}
	p.SetFixedDelay(dcl.ValueOrEmptyString(o.FixedDelay))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// GrpcRouteRulesActionFaultInjectionPolicyAbortToProto converts a GrpcRouteRulesActionFaultInjectionPolicyAbort object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyAbortToProto(o *beta.GrpcRouteRulesActionFaultInjectionPolicyAbort) *betapb.NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRulesActionFaultInjectionPolicyAbort{}
	p.SetHttpStatus(dcl.ValueOrEmptyInt64(o.HttpStatus))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// GrpcRouteRulesActionRetryPolicyToProto converts a GrpcRouteRulesActionRetryPolicy object to its proto representation.
func NetworkservicesBetaGrpcRouteRulesActionRetryPolicyToProto(o *beta.GrpcRouteRulesActionRetryPolicy) *betapb.NetworkservicesBetaGrpcRouteRulesActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaGrpcRouteRulesActionRetryPolicy{}
	p.SetNumRetries(dcl.ValueOrEmptyInt64(o.NumRetries))
	sRetryConditions := make([]string, len(o.RetryConditions))
	for i, r := range o.RetryConditions {
		sRetryConditions[i] = r
	}
	p.SetRetryConditions(sRetryConditions)
	return p
}

// GrpcRouteToProto converts a GrpcRoute resource to its proto representation.
func GrpcRouteToProto(resource *beta.GrpcRoute) *betapb.NetworkservicesBetaGrpcRoute {
	p := &betapb.NetworkservicesBetaGrpcRoute{}
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
	sRules := make([]*betapb.NetworkservicesBetaGrpcRouteRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworkservicesBetaGrpcRouteRulesToProto(&r)
	}
	p.SetRules(sRules)

	return p
}

// applyGrpcRoute handles the gRPC request by passing it to the underlying GrpcRoute Apply() method.
func (s *GrpcRouteServer) applyGrpcRoute(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkservicesBetaGrpcRouteRequest) (*betapb.NetworkservicesBetaGrpcRoute, error) {
	p := ProtoToGrpcRoute(request.GetResource())
	res, err := c.ApplyGrpcRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GrpcRouteToProto(res)
	return r, nil
}

// applyNetworkservicesBetaGrpcRoute handles the gRPC request by passing it to the underlying GrpcRoute Apply() method.
func (s *GrpcRouteServer) ApplyNetworkservicesBetaGrpcRoute(ctx context.Context, request *betapb.ApplyNetworkservicesBetaGrpcRouteRequest) (*betapb.NetworkservicesBetaGrpcRoute, error) {
	cl, err := createConfigGrpcRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGrpcRoute(ctx, cl, request)
}

// DeleteGrpcRoute handles the gRPC request by passing it to the underlying GrpcRoute Delete() method.
func (s *GrpcRouteServer) DeleteNetworkservicesBetaGrpcRoute(ctx context.Context, request *betapb.DeleteNetworkservicesBetaGrpcRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGrpcRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGrpcRoute(ctx, ProtoToGrpcRoute(request.GetResource()))

}

// ListNetworkservicesBetaGrpcRoute handles the gRPC request by passing it to the underlying GrpcRouteList() method.
func (s *GrpcRouteServer) ListNetworkservicesBetaGrpcRoute(ctx context.Context, request *betapb.ListNetworkservicesBetaGrpcRouteRequest) (*betapb.ListNetworkservicesBetaGrpcRouteResponse, error) {
	cl, err := createConfigGrpcRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGrpcRoute(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkservicesBetaGrpcRoute
	for _, r := range resources.Items {
		rp := GrpcRouteToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworkservicesBetaGrpcRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGrpcRoute(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
