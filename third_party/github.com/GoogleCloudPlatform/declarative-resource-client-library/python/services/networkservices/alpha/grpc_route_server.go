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

// GrpcRouteServer implements the gRPC interface for GrpcRoute.
type GrpcRouteServer struct{}

// ProtoToGrpcRouteRulesMatchesMethodTypeEnum converts a GrpcRouteRulesMatchesMethodTypeEnum enum from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum(e alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum) *alpha.GrpcRouteRulesMatchesMethodTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum_name[int32(e)]; ok {
		e := alpha.GrpcRouteRulesMatchesMethodTypeEnum(n[len("NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGrpcRouteRulesMatchesHeadersTypeEnum converts a GrpcRouteRulesMatchesHeadersTypeEnum enum from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum(e alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum) *alpha.GrpcRouteRulesMatchesHeadersTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum_name[int32(e)]; ok {
		e := alpha.GrpcRouteRulesMatchesHeadersTypeEnum(n[len("NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToGrpcRouteRules converts a GrpcRouteRules object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRules(p *alphapb.NetworkservicesAlphaGrpcRouteRules) *alpha.GrpcRouteRules {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRules{
		Action: ProtoToNetworkservicesAlphaGrpcRouteRulesAction(p.GetAction()),
	}
	for _, r := range p.GetMatches() {
		obj.Matches = append(obj.Matches, *ProtoToNetworkservicesAlphaGrpcRouteRulesMatches(r))
	}
	return obj
}

// ProtoToGrpcRouteRulesMatches converts a GrpcRouteRulesMatches object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesMatches(p *alphapb.NetworkservicesAlphaGrpcRouteRulesMatches) *alpha.GrpcRouteRulesMatches {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRulesMatches{
		Method: ProtoToNetworkservicesAlphaGrpcRouteRulesMatchesMethod(p.GetMethod()),
	}
	for _, r := range p.GetHeaders() {
		obj.Headers = append(obj.Headers, *ProtoToNetworkservicesAlphaGrpcRouteRulesMatchesHeaders(r))
	}
	return obj
}

// ProtoToGrpcRouteRulesMatchesMethod converts a GrpcRouteRulesMatchesMethod object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesMatchesMethod(p *alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethod) *alpha.GrpcRouteRulesMatchesMethod {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRulesMatchesMethod{
		Type:          ProtoToNetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum(p.GetType()),
		GrpcService:   dcl.StringOrNil(p.GetGrpcService()),
		GrpcMethod:    dcl.StringOrNil(p.GetGrpcMethod()),
		CaseSensitive: dcl.Bool(p.GetCaseSensitive()),
	}
	return obj
}

// ProtoToGrpcRouteRulesMatchesHeaders converts a GrpcRouteRulesMatchesHeaders object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesMatchesHeaders(p *alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeaders) *alpha.GrpcRouteRulesMatchesHeaders {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRulesMatchesHeaders{
		Type:  ProtoToNetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum(p.GetType()),
		Key:   dcl.StringOrNil(p.GetKey()),
		Value: dcl.StringOrNil(p.GetValue()),
	}
	return obj
}

// ProtoToGrpcRouteRulesAction converts a GrpcRouteRulesAction object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesAction(p *alphapb.NetworkservicesAlphaGrpcRouteRulesAction) *alpha.GrpcRouteRulesAction {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRulesAction{
		FaultInjectionPolicy: ProtoToNetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicy(p.GetFaultInjectionPolicy()),
		Timeout:              dcl.StringOrNil(p.GetTimeout()),
		RetryPolicy:          ProtoToNetworkservicesAlphaGrpcRouteRulesActionRetryPolicy(p.GetRetryPolicy()),
	}
	for _, r := range p.GetDestinations() {
		obj.Destinations = append(obj.Destinations, *ProtoToNetworkservicesAlphaGrpcRouteRulesActionDestinations(r))
	}
	return obj
}

// ProtoToGrpcRouteRulesActionDestinations converts a GrpcRouteRulesActionDestinations object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesActionDestinations(p *alphapb.NetworkservicesAlphaGrpcRouteRulesActionDestinations) *alpha.GrpcRouteRulesActionDestinations {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRulesActionDestinations{
		Weight:      dcl.Int64OrNil(p.GetWeight()),
		ServiceName: dcl.StringOrNil(p.GetServiceName()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionFaultInjectionPolicy converts a GrpcRouteRulesActionFaultInjectionPolicy object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicy(p *alphapb.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicy) *alpha.GrpcRouteRulesActionFaultInjectionPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRulesActionFaultInjectionPolicy{
		Delay: ProtoToNetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyDelay(p.GetDelay()),
		Abort: ProtoToNetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyAbort(p.GetAbort()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionFaultInjectionPolicyDelay converts a GrpcRouteRulesActionFaultInjectionPolicyDelay object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyDelay(p *alphapb.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyDelay) *alpha.GrpcRouteRulesActionFaultInjectionPolicyDelay {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRulesActionFaultInjectionPolicyDelay{
		FixedDelay: dcl.StringOrNil(p.GetFixedDelay()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionFaultInjectionPolicyAbort converts a GrpcRouteRulesActionFaultInjectionPolicyAbort object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyAbort(p *alphapb.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyAbort) *alpha.GrpcRouteRulesActionFaultInjectionPolicyAbort {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRulesActionFaultInjectionPolicyAbort{
		HttpStatus: dcl.Int64OrNil(p.GetHttpStatus()),
		Percentage: dcl.Int64OrNil(p.GetPercentage()),
	}
	return obj
}

// ProtoToGrpcRouteRulesActionRetryPolicy converts a GrpcRouteRulesActionRetryPolicy object from its proto representation.
func ProtoToNetworkservicesAlphaGrpcRouteRulesActionRetryPolicy(p *alphapb.NetworkservicesAlphaGrpcRouteRulesActionRetryPolicy) *alpha.GrpcRouteRulesActionRetryPolicy {
	if p == nil {
		return nil
	}
	obj := &alpha.GrpcRouteRulesActionRetryPolicy{
		NumRetries: dcl.Int64OrNil(p.GetNumRetries()),
	}
	for _, r := range p.GetRetryConditions() {
		obj.RetryConditions = append(obj.RetryConditions, r)
	}
	return obj
}

// ProtoToGrpcRoute converts a GrpcRoute resource from its proto representation.
func ProtoToGrpcRoute(p *alphapb.NetworkservicesAlphaGrpcRoute) *alpha.GrpcRoute {
	obj := &alpha.GrpcRoute{
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
		obj.Rules = append(obj.Rules, *ProtoToNetworkservicesAlphaGrpcRouteRules(r))
	}
	return obj
}

// GrpcRouteRulesMatchesMethodTypeEnumToProto converts a GrpcRouteRulesMatchesMethodTypeEnum enum to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnumToProto(e *alpha.GrpcRouteRulesMatchesMethodTypeEnum) alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum {
	if e == nil {
		return alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum(0)
	}
	if v, ok := alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum_value["GrpcRouteRulesMatchesMethodTypeEnum"+string(*e)]; ok {
		return alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum(v)
	}
	return alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnum(0)
}

// GrpcRouteRulesMatchesHeadersTypeEnumToProto converts a GrpcRouteRulesMatchesHeadersTypeEnum enum to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnumToProto(e *alpha.GrpcRouteRulesMatchesHeadersTypeEnum) alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum {
	if e == nil {
		return alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum(0)
	}
	if v, ok := alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum_value["GrpcRouteRulesMatchesHeadersTypeEnum"+string(*e)]; ok {
		return alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum(v)
	}
	return alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnum(0)
}

// GrpcRouteRulesToProto converts a GrpcRouteRules object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesToProto(o *alpha.GrpcRouteRules) *alphapb.NetworkservicesAlphaGrpcRouteRules {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRules{}
	p.SetAction(NetworkservicesAlphaGrpcRouteRulesActionToProto(o.Action))
	sMatches := make([]*alphapb.NetworkservicesAlphaGrpcRouteRulesMatches, len(o.Matches))
	for i, r := range o.Matches {
		sMatches[i] = NetworkservicesAlphaGrpcRouteRulesMatchesToProto(&r)
	}
	p.SetMatches(sMatches)
	return p
}

// GrpcRouteRulesMatchesToProto converts a GrpcRouteRulesMatches object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesMatchesToProto(o *alpha.GrpcRouteRulesMatches) *alphapb.NetworkservicesAlphaGrpcRouteRulesMatches {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRulesMatches{}
	p.SetMethod(NetworkservicesAlphaGrpcRouteRulesMatchesMethodToProto(o.Method))
	sHeaders := make([]*alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeaders, len(o.Headers))
	for i, r := range o.Headers {
		sHeaders[i] = NetworkservicesAlphaGrpcRouteRulesMatchesHeadersToProto(&r)
	}
	p.SetHeaders(sHeaders)
	return p
}

// GrpcRouteRulesMatchesMethodToProto converts a GrpcRouteRulesMatchesMethod object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesMatchesMethodToProto(o *alpha.GrpcRouteRulesMatchesMethod) *alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethod {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesMethod{}
	p.SetType(NetworkservicesAlphaGrpcRouteRulesMatchesMethodTypeEnumToProto(o.Type))
	p.SetGrpcService(dcl.ValueOrEmptyString(o.GrpcService))
	p.SetGrpcMethod(dcl.ValueOrEmptyString(o.GrpcMethod))
	p.SetCaseSensitive(dcl.ValueOrEmptyBool(o.CaseSensitive))
	return p
}

// GrpcRouteRulesMatchesHeadersToProto converts a GrpcRouteRulesMatchesHeaders object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesMatchesHeadersToProto(o *alpha.GrpcRouteRulesMatchesHeaders) *alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeaders {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRulesMatchesHeaders{}
	p.SetType(NetworkservicesAlphaGrpcRouteRulesMatchesHeadersTypeEnumToProto(o.Type))
	p.SetKey(dcl.ValueOrEmptyString(o.Key))
	p.SetValue(dcl.ValueOrEmptyString(o.Value))
	return p
}

// GrpcRouteRulesActionToProto converts a GrpcRouteRulesAction object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesActionToProto(o *alpha.GrpcRouteRulesAction) *alphapb.NetworkservicesAlphaGrpcRouteRulesAction {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRulesAction{}
	p.SetFaultInjectionPolicy(NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyToProto(o.FaultInjectionPolicy))
	p.SetTimeout(dcl.ValueOrEmptyString(o.Timeout))
	p.SetRetryPolicy(NetworkservicesAlphaGrpcRouteRulesActionRetryPolicyToProto(o.RetryPolicy))
	sDestinations := make([]*alphapb.NetworkservicesAlphaGrpcRouteRulesActionDestinations, len(o.Destinations))
	for i, r := range o.Destinations {
		sDestinations[i] = NetworkservicesAlphaGrpcRouteRulesActionDestinationsToProto(&r)
	}
	p.SetDestinations(sDestinations)
	return p
}

// GrpcRouteRulesActionDestinationsToProto converts a GrpcRouteRulesActionDestinations object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesActionDestinationsToProto(o *alpha.GrpcRouteRulesActionDestinations) *alphapb.NetworkservicesAlphaGrpcRouteRulesActionDestinations {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRulesActionDestinations{}
	p.SetWeight(dcl.ValueOrEmptyInt64(o.Weight))
	p.SetServiceName(dcl.ValueOrEmptyString(o.ServiceName))
	return p
}

// GrpcRouteRulesActionFaultInjectionPolicyToProto converts a GrpcRouteRulesActionFaultInjectionPolicy object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyToProto(o *alpha.GrpcRouteRulesActionFaultInjectionPolicy) *alphapb.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicy{}
	p.SetDelay(NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyDelayToProto(o.Delay))
	p.SetAbort(NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyAbortToProto(o.Abort))
	return p
}

// GrpcRouteRulesActionFaultInjectionPolicyDelayToProto converts a GrpcRouteRulesActionFaultInjectionPolicyDelay object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyDelayToProto(o *alpha.GrpcRouteRulesActionFaultInjectionPolicyDelay) *alphapb.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyDelay {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyDelay{}
	p.SetFixedDelay(dcl.ValueOrEmptyString(o.FixedDelay))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// GrpcRouteRulesActionFaultInjectionPolicyAbortToProto converts a GrpcRouteRulesActionFaultInjectionPolicyAbort object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyAbortToProto(o *alpha.GrpcRouteRulesActionFaultInjectionPolicyAbort) *alphapb.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyAbort {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRulesActionFaultInjectionPolicyAbort{}
	p.SetHttpStatus(dcl.ValueOrEmptyInt64(o.HttpStatus))
	p.SetPercentage(dcl.ValueOrEmptyInt64(o.Percentage))
	return p
}

// GrpcRouteRulesActionRetryPolicyToProto converts a GrpcRouteRulesActionRetryPolicy object to its proto representation.
func NetworkservicesAlphaGrpcRouteRulesActionRetryPolicyToProto(o *alpha.GrpcRouteRulesActionRetryPolicy) *alphapb.NetworkservicesAlphaGrpcRouteRulesActionRetryPolicy {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaGrpcRouteRulesActionRetryPolicy{}
	p.SetNumRetries(dcl.ValueOrEmptyInt64(o.NumRetries))
	sRetryConditions := make([]string, len(o.RetryConditions))
	for i, r := range o.RetryConditions {
		sRetryConditions[i] = r
	}
	p.SetRetryConditions(sRetryConditions)
	return p
}

// GrpcRouteToProto converts a GrpcRoute resource to its proto representation.
func GrpcRouteToProto(resource *alpha.GrpcRoute) *alphapb.NetworkservicesAlphaGrpcRoute {
	p := &alphapb.NetworkservicesAlphaGrpcRoute{}
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
	sRules := make([]*alphapb.NetworkservicesAlphaGrpcRouteRules, len(resource.Rules))
	for i, r := range resource.Rules {
		sRules[i] = NetworkservicesAlphaGrpcRouteRulesToProto(&r)
	}
	p.SetRules(sRules)

	return p
}

// applyGrpcRoute handles the gRPC request by passing it to the underlying GrpcRoute Apply() method.
func (s *GrpcRouteServer) applyGrpcRoute(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkservicesAlphaGrpcRouteRequest) (*alphapb.NetworkservicesAlphaGrpcRoute, error) {
	p := ProtoToGrpcRoute(request.GetResource())
	res, err := c.ApplyGrpcRoute(ctx, p)
	if err != nil {
		return nil, err
	}
	r := GrpcRouteToProto(res)
	return r, nil
}

// applyNetworkservicesAlphaGrpcRoute handles the gRPC request by passing it to the underlying GrpcRoute Apply() method.
func (s *GrpcRouteServer) ApplyNetworkservicesAlphaGrpcRoute(ctx context.Context, request *alphapb.ApplyNetworkservicesAlphaGrpcRouteRequest) (*alphapb.NetworkservicesAlphaGrpcRoute, error) {
	cl, err := createConfigGrpcRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyGrpcRoute(ctx, cl, request)
}

// DeleteGrpcRoute handles the gRPC request by passing it to the underlying GrpcRoute Delete() method.
func (s *GrpcRouteServer) DeleteNetworkservicesAlphaGrpcRoute(ctx context.Context, request *alphapb.DeleteNetworkservicesAlphaGrpcRouteRequest) (*emptypb.Empty, error) {

	cl, err := createConfigGrpcRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteGrpcRoute(ctx, ProtoToGrpcRoute(request.GetResource()))

}

// ListNetworkservicesAlphaGrpcRoute handles the gRPC request by passing it to the underlying GrpcRouteList() method.
func (s *GrpcRouteServer) ListNetworkservicesAlphaGrpcRoute(ctx context.Context, request *alphapb.ListNetworkservicesAlphaGrpcRouteRequest) (*alphapb.ListNetworkservicesAlphaGrpcRouteResponse, error) {
	cl, err := createConfigGrpcRoute(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListGrpcRoute(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkservicesAlphaGrpcRoute
	for _, r := range resources.Items {
		rp := GrpcRouteToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworkservicesAlphaGrpcRouteResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigGrpcRoute(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
