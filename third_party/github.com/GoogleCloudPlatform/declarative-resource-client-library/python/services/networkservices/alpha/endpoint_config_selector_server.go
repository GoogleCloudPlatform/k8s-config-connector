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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/networkservices/alpha/networkservices_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/networkservices/alpha"
)

// Server implements the gRPC interface for EndpointConfigSelector.
type EndpointConfigSelectorServer struct{}

// ProtoToEndpointConfigSelectorTypeEnum converts a EndpointConfigSelectorTypeEnum enum from its proto representation.
func ProtoToNetworkservicesAlphaEndpointConfigSelectorTypeEnum(e alphapb.NetworkservicesAlphaEndpointConfigSelectorTypeEnum) *alpha.EndpointConfigSelectorTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkservicesAlphaEndpointConfigSelectorTypeEnum_name[int32(e)]; ok {
		e := alpha.EndpointConfigSelectorTypeEnum(n[len("NetworkservicesAlphaEndpointConfigSelectorTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum from its proto representation.
func ProtoToNetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(e alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) *alpha.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_name[int32(e)]; ok {
		e := alpha.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(n[len("NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointConfigSelectorHttpFilters converts a EndpointConfigSelectorHttpFilters resource from its proto representation.
func ProtoToNetworkservicesAlphaEndpointConfigSelectorHttpFilters(p *alphapb.NetworkservicesAlphaEndpointConfigSelectorHttpFilters) *alpha.EndpointConfigSelectorHttpFilters {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointConfigSelectorHttpFilters{}
	for _, r := range p.GetHttpFilters() {
		obj.HttpFilters = append(obj.HttpFilters, r)
	}
	return obj
}

// ProtoToEndpointConfigSelectorEndpointMatcher converts a EndpointConfigSelectorEndpointMatcher resource from its proto representation.
func ProtoToNetworkservicesAlphaEndpointConfigSelectorEndpointMatcher(p *alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcher) *alpha.EndpointConfigSelectorEndpointMatcher {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointConfigSelectorEndpointMatcher{
		MetadataLabelMatcher: ProtoToNetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(p.GetMetadataLabelMatcher()),
	}
	return obj
}

// ProtoToEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher resource from its proto representation.
func ProtoToNetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher(p *alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) *alpha.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{
		MetadataLabelMatchCriteria: ProtoToNetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(p.GetMetadataLabelMatchCriteria()),
	}
	for _, r := range p.GetMetadataLabels() {
		obj.MetadataLabels = append(obj.MetadataLabels, *ProtoToNetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(r))
	}
	return obj
}

// ProtoToEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels resource from its proto representation.
func ProtoToNetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels(p *alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) *alpha.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{
		LabelName:  dcl.StringOrNil(p.LabelName),
		LabelValue: dcl.StringOrNil(p.LabelValue),
	}
	return obj
}

// ProtoToEndpointConfigSelectorTrafficPortSelector converts a EndpointConfigSelectorTrafficPortSelector resource from its proto representation.
func ProtoToNetworkservicesAlphaEndpointConfigSelectorTrafficPortSelector(p *alphapb.NetworkservicesAlphaEndpointConfigSelectorTrafficPortSelector) *alpha.EndpointConfigSelectorTrafficPortSelector {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointConfigSelectorTrafficPortSelector{}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToEndpointConfigSelector converts a EndpointConfigSelector resource from its proto representation.
func ProtoToEndpointConfigSelector(p *alphapb.NetworkservicesAlphaEndpointConfigSelector) *alpha.EndpointConfigSelector {
	obj := &alpha.EndpointConfigSelector{
		Name:                dcl.StringOrNil(p.Name),
		CreateTime:          dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:          dcl.StringOrNil(p.GetUpdateTime()),
		Type:                ProtoToNetworkservicesAlphaEndpointConfigSelectorTypeEnum(p.GetType()),
		AuthorizationPolicy: dcl.StringOrNil(p.AuthorizationPolicy),
		HttpFilters:         ProtoToNetworkservicesAlphaEndpointConfigSelectorHttpFilters(p.GetHttpFilters()),
		EndpointMatcher:     ProtoToNetworkservicesAlphaEndpointConfigSelectorEndpointMatcher(p.GetEndpointMatcher()),
		TrafficPortSelector: ProtoToNetworkservicesAlphaEndpointConfigSelectorTrafficPortSelector(p.GetTrafficPortSelector()),
		Description:         dcl.StringOrNil(p.Description),
		ServerTlsPolicy:     dcl.StringOrNil(p.ServerTlsPolicy),
		ClientTlsPolicy:     dcl.StringOrNil(p.ClientTlsPolicy),
		Project:             dcl.StringOrNil(p.Project),
		Location:            dcl.StringOrNil(p.Location),
	}
	return obj
}

// EndpointConfigSelectorTypeEnumToProto converts a EndpointConfigSelectorTypeEnum enum to its proto representation.
func NetworkservicesAlphaEndpointConfigSelectorTypeEnumToProto(e *alpha.EndpointConfigSelectorTypeEnum) alphapb.NetworkservicesAlphaEndpointConfigSelectorTypeEnum {
	if e == nil {
		return alphapb.NetworkservicesAlphaEndpointConfigSelectorTypeEnum(0)
	}
	if v, ok := alphapb.NetworkservicesAlphaEndpointConfigSelectorTypeEnum_value["EndpointConfigSelectorTypeEnum"+string(*e)]; ok {
		return alphapb.NetworkservicesAlphaEndpointConfigSelectorTypeEnum(v)
	}
	return alphapb.NetworkservicesAlphaEndpointConfigSelectorTypeEnum(0)
}

// EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum to its proto representation.
func NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(e *alpha.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == nil {
		return alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
	}
	if v, ok := alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_value["EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"+string(*e)]; ok {
		return alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(v)
	}
	return alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
}

// EndpointConfigSelectorHttpFiltersToProto converts a EndpointConfigSelectorHttpFilters resource to its proto representation.
func NetworkservicesAlphaEndpointConfigSelectorHttpFiltersToProto(o *alpha.EndpointConfigSelectorHttpFilters) *alphapb.NetworkservicesAlphaEndpointConfigSelectorHttpFilters {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaEndpointConfigSelectorHttpFilters{}
	for _, r := range o.HttpFilters {
		p.HttpFilters = append(p.HttpFilters, r)
	}
	return p
}

// EndpointConfigSelectorEndpointMatcherToProto converts a EndpointConfigSelectorEndpointMatcher resource to its proto representation.
func NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherToProto(o *alpha.EndpointConfigSelectorEndpointMatcher) *alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcher {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcher{
		MetadataLabelMatcher: NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherToProto(o.MetadataLabelMatcher),
	}
	return p
}

// EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherToProto converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher resource to its proto representation.
func NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherToProto(o *alpha.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcher) *alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcher{
		MetadataLabelMatchCriteria: NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(o.MetadataLabelMatchCriteria),
	}
	for _, r := range o.MetadataLabels {
		p.MetadataLabels = append(p.MetadataLabels, NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(&r))
	}
	return p
}

// EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto converts a EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels resource to its proto representation.
func NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(o *alpha.EndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels) *alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherMetadataLabelMatcherMetadataLabels{
		LabelName:  dcl.ValueOrEmptyString(o.LabelName),
		LabelValue: dcl.ValueOrEmptyString(o.LabelValue),
	}
	return p
}

// EndpointConfigSelectorTrafficPortSelectorToProto converts a EndpointConfigSelectorTrafficPortSelector resource to its proto representation.
func NetworkservicesAlphaEndpointConfigSelectorTrafficPortSelectorToProto(o *alpha.EndpointConfigSelectorTrafficPortSelector) *alphapb.NetworkservicesAlphaEndpointConfigSelectorTrafficPortSelector {
	if o == nil {
		return nil
	}
	p := &alphapb.NetworkservicesAlphaEndpointConfigSelectorTrafficPortSelector{}
	for _, r := range o.Ports {
		p.Ports = append(p.Ports, r)
	}
	return p
}

// EndpointConfigSelectorToProto converts a EndpointConfigSelector resource to its proto representation.
func EndpointConfigSelectorToProto(resource *alpha.EndpointConfigSelector) *alphapb.NetworkservicesAlphaEndpointConfigSelector {
	p := &alphapb.NetworkservicesAlphaEndpointConfigSelector{
		Name:                dcl.ValueOrEmptyString(resource.Name),
		CreateTime:          dcl.ValueOrEmptyString(resource.CreateTime),
		UpdateTime:          dcl.ValueOrEmptyString(resource.UpdateTime),
		Type:                NetworkservicesAlphaEndpointConfigSelectorTypeEnumToProto(resource.Type),
		AuthorizationPolicy: dcl.ValueOrEmptyString(resource.AuthorizationPolicy),
		HttpFilters:         NetworkservicesAlphaEndpointConfigSelectorHttpFiltersToProto(resource.HttpFilters),
		EndpointMatcher:     NetworkservicesAlphaEndpointConfigSelectorEndpointMatcherToProto(resource.EndpointMatcher),
		TrafficPortSelector: NetworkservicesAlphaEndpointConfigSelectorTrafficPortSelectorToProto(resource.TrafficPortSelector),
		Description:         dcl.ValueOrEmptyString(resource.Description),
		ServerTlsPolicy:     dcl.ValueOrEmptyString(resource.ServerTlsPolicy),
		ClientTlsPolicy:     dcl.ValueOrEmptyString(resource.ClientTlsPolicy),
		Project:             dcl.ValueOrEmptyString(resource.Project),
		Location:            dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyEndpointConfigSelector handles the gRPC request by passing it to the underlying EndpointConfigSelector Apply() method.
func (s *EndpointConfigSelectorServer) applyEndpointConfigSelector(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkservicesAlphaEndpointConfigSelectorRequest) (*alphapb.NetworkservicesAlphaEndpointConfigSelector, error) {
	p := ProtoToEndpointConfigSelector(request.GetResource())
	res, err := c.ApplyEndpointConfigSelector(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointConfigSelectorToProto(res)
	return r, nil
}

// ApplyEndpointConfigSelector handles the gRPC request by passing it to the underlying EndpointConfigSelector Apply() method.
func (s *EndpointConfigSelectorServer) ApplyNetworkservicesAlphaEndpointConfigSelector(ctx context.Context, request *alphapb.ApplyNetworkservicesAlphaEndpointConfigSelectorRequest) (*alphapb.NetworkservicesAlphaEndpointConfigSelector, error) {
	cl, err := createConfigEndpointConfigSelector(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyEndpointConfigSelector(ctx, cl, request)
}

// DeleteEndpointConfigSelector handles the gRPC request by passing it to the underlying EndpointConfigSelector Delete() method.
func (s *EndpointConfigSelectorServer) DeleteNetworkservicesAlphaEndpointConfigSelector(ctx context.Context, request *alphapb.DeleteNetworkservicesAlphaEndpointConfigSelectorRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpointConfigSelector(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpointConfigSelector(ctx, ProtoToEndpointConfigSelector(request.GetResource()))

}

// ListNetworkservicesAlphaEndpointConfigSelector handles the gRPC request by passing it to the underlying EndpointConfigSelectorList() method.
func (s *EndpointConfigSelectorServer) ListNetworkservicesAlphaEndpointConfigSelector(ctx context.Context, request *alphapb.ListNetworkservicesAlphaEndpointConfigSelectorRequest) (*alphapb.ListNetworkservicesAlphaEndpointConfigSelectorResponse, error) {
	cl, err := createConfigEndpointConfigSelector(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEndpointConfigSelector(ctx, ProtoToEndpointConfigSelector(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkservicesAlphaEndpointConfigSelector
	for _, r := range resources.Items {
		rp := EndpointConfigSelectorToProto(r)
		protos = append(protos, rp)
	}
	return &alphapb.ListNetworkservicesAlphaEndpointConfigSelectorResponse{Items: protos}, nil
}

func createConfigEndpointConfigSelector(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
