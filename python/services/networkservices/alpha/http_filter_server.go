// Copyright 2022 Google LLC. All Rights Reserved.
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

// HttpFilterServer implements the gRPC interface for HttpFilter.
type HttpFilterServer struct{}

// ProtoToHttpFilter converts a HttpFilter resource from its proto representation.
func ProtoToHttpFilter(p *alphapb.NetworkservicesAlphaHttpFilter) *alpha.HttpFilter {
	obj := &alpha.HttpFilter{
		Name:          dcl.StringOrNil(p.GetName()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:    dcl.StringOrNil(p.GetUpdateTime()),
		FilterName:    dcl.StringOrNil(p.GetFilterName()),
		ConfigTypeUrl: dcl.StringOrNil(p.GetConfigTypeUrl()),
		Config:        dcl.StringOrNil(p.GetConfig()),
		Description:   dcl.StringOrNil(p.GetDescription()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// HttpFilterToProto converts a HttpFilter resource to its proto representation.
func HttpFilterToProto(resource *alpha.HttpFilter) *alphapb.NetworkservicesAlphaHttpFilter {
	p := &alphapb.NetworkservicesAlphaHttpFilter{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetFilterName(dcl.ValueOrEmptyString(resource.FilterName))
	p.SetConfigTypeUrl(dcl.ValueOrEmptyString(resource.ConfigTypeUrl))
	p.SetConfig(dcl.ValueOrEmptyString(resource.Config))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyHttpFilter handles the gRPC request by passing it to the underlying HttpFilter Apply() method.
func (s *HttpFilterServer) applyHttpFilter(ctx context.Context, c *alpha.Client, request *alphapb.ApplyNetworkservicesAlphaHttpFilterRequest) (*alphapb.NetworkservicesAlphaHttpFilter, error) {
	p := ProtoToHttpFilter(request.GetResource())
	res, err := c.ApplyHttpFilter(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HttpFilterToProto(res)
	return r, nil
}

// applyNetworkservicesAlphaHttpFilter handles the gRPC request by passing it to the underlying HttpFilter Apply() method.
func (s *HttpFilterServer) ApplyNetworkservicesAlphaHttpFilter(ctx context.Context, request *alphapb.ApplyNetworkservicesAlphaHttpFilterRequest) (*alphapb.NetworkservicesAlphaHttpFilter, error) {
	cl, err := createConfigHttpFilter(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyHttpFilter(ctx, cl, request)
}

// DeleteHttpFilter handles the gRPC request by passing it to the underlying HttpFilter Delete() method.
func (s *HttpFilterServer) DeleteNetworkservicesAlphaHttpFilter(ctx context.Context, request *alphapb.DeleteNetworkservicesAlphaHttpFilterRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHttpFilter(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHttpFilter(ctx, ProtoToHttpFilter(request.GetResource()))

}

// ListNetworkservicesAlphaHttpFilter handles the gRPC request by passing it to the underlying HttpFilterList() method.
func (s *HttpFilterServer) ListNetworkservicesAlphaHttpFilter(ctx context.Context, request *alphapb.ListNetworkservicesAlphaHttpFilterRequest) (*alphapb.ListNetworkservicesAlphaHttpFilterResponse, error) {
	cl, err := createConfigHttpFilter(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHttpFilter(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.NetworkservicesAlphaHttpFilter
	for _, r := range resources.Items {
		rp := HttpFilterToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListNetworkservicesAlphaHttpFilterResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigHttpFilter(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
