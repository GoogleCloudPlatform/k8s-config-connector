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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/beta/networkservices_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/beta"
)

// Server implements the gRPC interface for EndpointConfigSelector.
type EndpointConfigSelectorServer struct{}

// ProtoToEndpointConfigSelectorTypeEnum converts a EndpointConfigSelectorTypeEnum enum from its proto representation.
func ProtoToNetworkservicesBetaEndpointConfigSelectorTypeEnum(e betapb.NetworkservicesBetaEndpointConfigSelectorTypeEnum) *beta.EndpointConfigSelectorTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkservicesBetaEndpointConfigSelectorTypeEnum_name[int32(e)]; ok {
		e := beta.EndpointConfigSelectorTypeEnum(n[len("NetworkservicesBetaEndpointConfigSelectorTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum from its proto representation.
func ProtoToNetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(e betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) *beta.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_name[int32(e)]; ok {
		e := beta.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(n[len("NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointConfigSelectorHttpFilters converts a EndpointConfigSelectorHttpFilters resource from its proto representation.
func ProtoToNetworkservicesBetaEndpointConfigSelectorHttpFilters(p *betapb.NetworkservicesBetaEndpointConfigSelectorHttpFilters) *beta.EndpointConfigSelectorHttpFilters {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointConfigSelectorHttpFilters{}
	for _, r := range p.GetHttpFilters() {
		obj.HttpFilters = append(obj.HttpFilters, r)
	}
	return obj
}

// ProtoToEndpointConfigSelectorEndpointMatcher converts a EndpointConfigSelectorEndpointMatcher resource from its proto representation.
func ProtoToNetworkservicesBetaEndpointConfigSelectorEndpointMatcher(p *betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcher) *beta.EndpointConfigSelectorEndpointMatcher {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointConfigSelectorEndpointMatcher{
		MetadataLabelMatcher: ProtoToNetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(p.GetMetadataLabelMatcher()),
	}
	return obj
}

// ProtoToEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher resource from its proto representation.
func ProtoToNetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(p *betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) *beta.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{
		MetadataLabelMatchCriteria: ProtoToNetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(p.GetMetadataLabelMatchCriteria()),
	}
	for _, r := range p.GetMetadataLabels() {
		obj.MetadataLabels = append(obj.MetadataLabels, *ProtoToNetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(r))
	}
	return obj
}

// ProtoToEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels resource from its proto representation.
func ProtoToNetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(p *betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) *beta.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{
		LabelName:  dcl.StringOrNil(p.LabelName),
		LabelValue: dcl.StringOrNil(p.LabelValue),
	}
	return obj
}

// ProtoToEndpointConfigSelectorTrafficPortSelector converts a EndpointConfigSelectorTrafficPortSelector resource from its proto representation.
func ProtoToNetworkservicesBetaEndpointConfigSelectorTrafficPortSelector(p *betapb.NetworkservicesBetaEndpointConfigSelectorTrafficPortSelector) *beta.EndpointConfigSelectorTrafficPortSelector {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointConfigSelectorTrafficPortSelector{}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToEndpointConfigSelector converts a EndpointConfigSelector resource from its proto representation.
func ProtoToEndpointConfigSelector(p *betapb.NetworkservicesBetaEndpointConfigSelector) *beta.EndpointConfigSelector {
	obj := &beta.EndpointConfigSelector{
		Name:                dcl.StringOrNil(p.Name),
		CreateTime:          dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:          dcl.StringOrNil(p.GetUpdateTime()),
		Type:                ProtoToNetworkservicesBetaEndpointConfigSelectorTypeEnum(p.GetType()),
		AuthorizationPolicy: dcl.StringOrNil(p.AuthorizationPolicy),
		HttpFilters:         ProtoToNetworkservicesBetaEndpointConfigSelectorHttpFilters(p.GetHttpFilters()),
		EndpointMatcher:     ProtoToNetworkservicesBetaEndpointConfigSelectorEndpointMatcher(p.GetEndpointMatcher()),
		TrafficPortSelector: ProtoToNetworkservicesBetaEndpointConfigSelectorTrafficPortSelector(p.GetTrafficPortSelector()),
		Description:         dcl.StringOrNil(p.Description),
		ServerTlsPolicy:     dcl.StringOrNil(p.ServerTlsPolicy),
		ClientTlsPolicy:     dcl.StringOrNil(p.ClientTlsPolicy),
		Project:             dcl.StringOrNil(p.Project),
		Location:            dcl.StringOrNil(p.Location),
	}
	return obj
}

// EndpointConfigSelectorTypeEnumToProto converts a EndpointConfigSelectorTypeEnum enum to its proto representation.
func NetworkservicesBetaEndpointConfigSelectorTypeEnumToProto(e *beta.EndpointConfigSelectorTypeEnum) betapb.NetworkservicesBetaEndpointConfigSelectorTypeEnum {
	if e == nil {
		return betapb.NetworkservicesBetaEndpointConfigSelectorTypeEnum(0)
	}
	if v, ok := betapb.NetworkservicesBetaEndpointConfigSelectorTypeEnum_value["EndpointConfigSelectorTypeEnum"+string(*e)]; ok {
		return betapb.NetworkservicesBetaEndpointConfigSelectorTypeEnum(v)
	}
	return betapb.NetworkservicesBetaEndpointConfigSelectorTypeEnum(0)
}

// EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum to its proto representation.
func NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(e *beta.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == nil {
		return betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
	}
	if v, ok := betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_value["EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"+string(*e)]; ok {
		return betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(v)
	}
	return betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
}

// EndpointConfigSelectorHttpFiltersToProto converts a EndpointConfigSelectorHttpFilters resource to its proto representation.
func NetworkservicesBetaEndpointConfigSelectorHttpFiltersToProto(o *beta.EndpointConfigSelectorHttpFilters) *betapb.NetworkservicesBetaEndpointConfigSelectorHttpFilters {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointConfigSelectorHttpFilters{}
	for _, r := range o.HttpFilters {
		p.HttpFilters = append(p.HttpFilters, r)
	}
	return p
}

// EndpointConfigSelectorEndpointMatcherToProto converts a EndpointConfigSelectorEndpointMatcher resource to its proto representation.
func NetworkservicesBetaEndpointConfigSelectorEndpointMatcherToProto(o *beta.EndpointConfigSelectorEndpointMatcher) *betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcher {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcher{
		MetadataLabelMatcher: NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherToProto(o.MetadataLabelMatcher),
	}
	return p
}

// EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherToProto converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher resource to its proto representation.
func NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherToProto(o *beta.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) *betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{
		MetadataLabelMatchCriteria: NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(o.MetadataLabelMatchCriteria),
	}
	for _, r := range o.MetadataLabels {
		p.MetadataLabels = append(p.MetadataLabels, NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(&r))
	}
	return p
}

// EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels resource to its proto representation.
func NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(o *beta.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) *betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{
		LabelName:  dcl.ValueOrEmptyString(o.LabelName),
		LabelValue: dcl.ValueOrEmptyString(o.LabelValue),
	}
	return p
}

// EndpointConfigSelectorTrafficPortSelectorToProto converts a EndpointConfigSelectorTrafficPortSelector resource to its proto representation.
func NetworkservicesBetaEndpointConfigSelectorTrafficPortSelectorToProto(o *beta.EndpointConfigSelectorTrafficPortSelector) *betapb.NetworkservicesBetaEndpointConfigSelectorTrafficPortSelector {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointConfigSelectorTrafficPortSelector{}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	return p
}

// EndpointConfigSelectorToProto converts a EndpointConfigSelector resource to its proto representation.
func EndpointConfigSelectorToProto(resource *beta.EndpointConfigSelector) *betapb.NetworkservicesBetaEndpointConfigSelector {
	p := &betapb.NetworkservicesBetaEndpointConfigSelector{
		Name:                dcl.ValueOrEmptyString(resource.Name),
		CreateTime:          dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:          dcl.ValueOrEmptyString(resource.UpdateTime),
		Type:                NetworkservicesBetaEndpointConfigSelectorTypeEnumToProto(resource.Type),
		AuthorizationPolicy: dcl.ValueOrEmptyString(resource.AuthorizationPolicy),
		HttpFilters:         NetworkservicesBetaEndpointConfigSelectorHttpFiltersToProto(resource.HttpFilters),
		EndpointMatcher:     NetworkservicesBetaEndpointConfigSelectorEndpointMatcherToProto(resource.EndpointMatcher),
		TrafficPortSelector: NetworkservicesBetaEndpointConfigSelectorTrafficPortSelectorToProto(resource.TrafficPortSelector),
		Description:         dcl.ValueOrEmptyString(resource.Description),
		ServerTlsPolicy:     dcl.ValueOrEmptyString(resource.ServerTlsPolicy),
		ClientTlsPolicy:     dcl.ValueOrEmptyString(resource.ClientTlsPolicy),
		Project:             dcl.ValueOrEmptyString(resource.Project),
		Location:            dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyEndpointConfigSelector handles the gRPC request by passing it to the underlying EndpointConfigSelector Apply() method.
func (s *EndpointConfigSelectorServer) applyEndpointConfigSelector(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkservicesBetaEndpointConfigSelectorRequest) (*betapb.NetworkservicesBetaEndpointConfigSelector, error) {
	p := ProtoToEndpointConfigSelector(request.GetResource())
	res, err := c.ApplyEndpointConfigSelector(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointConfigSelectorToProto(res)
	return r, nil
}

// ApplyEndpointConfigSelector handles the gRPC request by passing it to the underlying EndpointConfigSelector Apply() method.
func (s *EndpointConfigSelectorServer) ApplyNetworkservicesBetaEndpointConfigSelector(ctx context.Context, request *betapb.ApplyNetworkservicesBetaEndpointConfigSelectorRequest) (*betapb.NetworkservicesBetaEndpointConfigSelector, error) {
	cl, err := createConfigEndpointConfigSelector(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyEndpointConfigSelector(ctx, cl, request)
}

// DeleteEndpointConfigSelector handles the gRPC request by passing it to the underlying EndpointConfigSelector Delete() method.
func (s *EndpointConfigSelectorServer) DeleteNetworkservicesBetaEndpointConfigSelector(ctx context.Context, request *betapb.DeleteNetworkservicesBetaEndpointConfigSelectorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpointConfigSelector(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpointConfigSelector(ctx, ProtoToEndpointConfigSelector(request.GetResource()))

}

// ListNetworkservicesBetaEndpointConfigSelector handles the gRPC request by passing it to the underlying EndpointConfigSelectorList() method.
func (s *EndpointConfigSelectorServer) ListNetworkservicesBetaEndpointConfigSelector(ctx context.Context, request *betapb.ListNetworkservicesBetaEndpointConfigSelectorRequest) (*betapb.ListNetworkservicesBetaEndpointConfigSelectorResponse, error) {
	cl, err := createConfigEndpointConfigSelector(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEndpointConfigSelector(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkservicesBetaEndpointConfigSelector
	for _, r := range resources.Items {
		rp := EndpointConfigSelectorToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListNetworkservicesBetaEndpointConfigSelectorResponse{Items: protos}, nil
}

func createConfigEndpointConfigSelector(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
