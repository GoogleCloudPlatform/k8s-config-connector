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

// EndpointPolicyServer implements the gRPC interface for EndpointPolicy.
type EndpointPolicyServer struct{}

// ProtoToEndpointPolicyTypeEnum converts a EndpointPolicyTypeEnum enum from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyTypeEnum(e betapb.NetworkservicesBetaEndpointPolicyTypeEnum) *beta.EndpointPolicyTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkservicesBetaEndpointPolicyTypeEnum_name[int32(e)]; ok {
		e := beta.EndpointPolicyTypeEnum(n[len("NetworkservicesBetaEndpointPolicyTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(e betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_name[int32(e)]; ok {
		e := beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(n[len("NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"):])
		return &e
	}
	return nil
}

// ProtoToEndpointPolicyEndpointMatcher converts a EndpointPolicyEndpointMatcher object from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcher(p *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcher) *beta.EndpointPolicyEndpointMatcher {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointPolicyEndpointMatcher{
		MetadataLabelMatcher: ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher(p.GetMetadataLabelMatcher()),
	}
	return obj
}

// ProtoToEndpointPolicyEndpointMatcherMetadataLabelMatcher converts a EndpointPolicyEndpointMatcherMetadataLabelMatcher object from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher(p *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher) *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointPolicyEndpointMatcherMetadataLabelMatcher{
		MetadataLabelMatchCriteria: ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(p.GetMetadataLabelMatchCriteria()),
	}
	for _, r := range p.GetMetadataLabels() {
		obj.MetadataLabels = append(obj.MetadataLabels, *ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(r))
	}
	return obj
}

// ProtoToEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels object from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels(p *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{
		LabelName:  dcl.StringOrNil(p.GetLabelName()),
		LabelValue: dcl.StringOrNil(p.GetLabelValue()),
	}
	return obj
}

// ProtoToEndpointPolicyTrafficPortSelector converts a EndpointPolicyTrafficPortSelector object from its proto representation.
func ProtoToNetworkservicesBetaEndpointPolicyTrafficPortSelector(p *betapb.NetworkservicesBetaEndpointPolicyTrafficPortSelector) *beta.EndpointPolicyTrafficPortSelector {
	if p == nil {
		return nil
	}
	obj := &beta.EndpointPolicyTrafficPortSelector{}
	for _, r := range p.GetPorts() {
		obj.Ports = append(obj.Ports, r)
	}
	return obj
}

// ProtoToEndpointPolicy converts a EndpointPolicy resource from its proto representation.
func ProtoToEndpointPolicy(p *betapb.NetworkservicesBetaEndpointPolicy) *beta.EndpointPolicy {
	obj := &beta.EndpointPolicy{
		Name:                dcl.StringOrNil(p.GetName()),
		CreateTime:          dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:          dcl.StringOrNil(p.GetUpdateTime()),
		Type:                ProtoToNetworkservicesBetaEndpointPolicyTypeEnum(p.GetType()),
		AuthorizationPolicy: dcl.StringOrNil(p.GetAuthorizationPolicy()),
		EndpointMatcher:     ProtoToNetworkservicesBetaEndpointPolicyEndpointMatcher(p.GetEndpointMatcher()),
		TrafficPortSelector: ProtoToNetworkservicesBetaEndpointPolicyTrafficPortSelector(p.GetTrafficPortSelector()),
		Description:         dcl.StringOrNil(p.GetDescription()),
		ServerTlsPolicy:     dcl.StringOrNil(p.GetServerTlsPolicy()),
		ClientTlsPolicy:     dcl.StringOrNil(p.GetClientTlsPolicy()),
		Project:             dcl.StringOrNil(p.GetProject()),
		Location:            dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// EndpointPolicyTypeEnumToProto converts a EndpointPolicyTypeEnum enum to its proto representation.
func NetworkservicesBetaEndpointPolicyTypeEnumToProto(e *beta.EndpointPolicyTypeEnum) betapb.NetworkservicesBetaEndpointPolicyTypeEnum {
	if e == nil {
		return betapb.NetworkservicesBetaEndpointPolicyTypeEnum(0)
	}
	if v, ok := betapb.NetworkservicesBetaEndpointPolicyTypeEnum_value["EndpointPolicyTypeEnum"+string(*e)]; ok {
		return betapb.NetworkservicesBetaEndpointPolicyTypeEnum(v)
	}
	return betapb.NetworkservicesBetaEndpointPolicyTypeEnum(0)
}

// EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum enum to its proto representation.
func NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(e *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum) betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum {
	if e == nil {
		return betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
	}
	if v, ok := betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum_value["EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum"+string(*e)]; ok {
		return betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(v)
	}
	return betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnum(0)
}

// EndpointPolicyEndpointMatcherToProto converts a EndpointPolicyEndpointMatcher object to its proto representation.
func NetworkservicesBetaEndpointPolicyEndpointMatcherToProto(o *beta.EndpointPolicyEndpointMatcher) *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcher {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointPolicyEndpointMatcher{}
	p.SetMetadataLabelMatcher(NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherToProto(o.MetadataLabelMatcher))
	return p
}

// EndpointPolicyEndpointMatcherMetadataLabelMatcherToProto converts a EndpointPolicyEndpointMatcherMetadataLabelMatcher object to its proto representation.
func NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherToProto(o *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcher) *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcher{}
	p.SetMetadataLabelMatchCriteria(NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelMatchCriteriaEnumToProto(o.MetadataLabelMatchCriteria))
	sMetadataLabels := make([]*betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels, len(o.MetadataLabels))
	for i, r := range o.MetadataLabels {
		sMetadataLabels[i] = NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(&r)
	}
	p.SetMetadataLabels(sMetadataLabels)
	return p
}

// EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto converts a EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels object to its proto representation.
func NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabelsToProto(o *beta.EndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels) *betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels{}
	p.SetLabelName(dcl.ValueOrEmptyString(o.LabelName))
	p.SetLabelValue(dcl.ValueOrEmptyString(o.LabelValue))
	return p
}

// EndpointPolicyTrafficPortSelectorToProto converts a EndpointPolicyTrafficPortSelector object to its proto representation.
func NetworkservicesBetaEndpointPolicyTrafficPortSelectorToProto(o *beta.EndpointPolicyTrafficPortSelector) *betapb.NetworkservicesBetaEndpointPolicyTrafficPortSelector {
	if o == nil {
		return nil
	}
	p := &betapb.NetworkservicesBetaEndpointPolicyTrafficPortSelector{}
	sPorts := make([]string, len(o.Ports))
	for i, r := range o.Ports {
		sPorts[i] = r
	}
	p.SetPorts(sPorts)
	return p
}

// EndpointPolicyToProto converts a EndpointPolicy resource to its proto representation.
func EndpointPolicyToProto(resource *beta.EndpointPolicy) *betapb.NetworkservicesBetaEndpointPolicy {
	p := &betapb.NetworkservicesBetaEndpointPolicy{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetType(NetworkservicesBetaEndpointPolicyTypeEnumToProto(resource.Type))
	p.SetAuthorizationPolicy(dcl.ValueOrEmptyString(resource.AuthorizationPolicy))
	p.SetEndpointMatcher(NetworkservicesBetaEndpointPolicyEndpointMatcherToProto(resource.EndpointMatcher))
	p.SetTrafficPortSelector(NetworkservicesBetaEndpointPolicyTrafficPortSelectorToProto(resource.TrafficPortSelector))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetServerTlsPolicy(dcl.ValueOrEmptyString(resource.ServerTlsPolicy))
	p.SetClientTlsPolicy(dcl.ValueOrEmptyString(resource.ClientTlsPolicy))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicy Apply() method.
func (s *EndpointPolicyServer) applyEndpointPolicy(ctx context.Context, c *beta.Client, request *betapb.ApplyNetworkservicesBetaEndpointPolicyRequest) (*betapb.NetworkservicesBetaEndpointPolicy, error) {
	p := ProtoToEndpointPolicy(request.GetResource())
	res, err := c.ApplyEndpointPolicy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointPolicyToProto(res)
	return r, nil
}

// applyNetworkservicesBetaEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicy Apply() method.
func (s *EndpointPolicyServer) ApplyNetworkservicesBetaEndpointPolicy(ctx context.Context, request *betapb.ApplyNetworkservicesBetaEndpointPolicyRequest) (*betapb.NetworkservicesBetaEndpointPolicy, error) {
	cl, err := createConfigEndpointPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEndpointPolicy(ctx, cl, request)
}

// DeleteEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicy Delete() method.
func (s *EndpointPolicyServer) DeleteNetworkservicesBetaEndpointPolicy(ctx context.Context, request *betapb.DeleteNetworkservicesBetaEndpointPolicyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpointPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpointPolicy(ctx, ProtoToEndpointPolicy(request.GetResource()))

}

// ListNetworkservicesBetaEndpointPolicy handles the gRPC request by passing it to the underlying EndpointPolicyList() method.
func (s *EndpointPolicyServer) ListNetworkservicesBetaEndpointPolicy(ctx context.Context, request *betapb.ListNetworkservicesBetaEndpointPolicyRequest) (*betapb.ListNetworkservicesBetaEndpointPolicyResponse, error) {
	cl, err := createConfigEndpointPolicy(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEndpointPolicy(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.NetworkservicesBetaEndpointPolicy
	for _, r := range resources.Items {
		rp := EndpointPolicyToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListNetworkservicesBetaEndpointPolicyResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEndpointPolicy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
