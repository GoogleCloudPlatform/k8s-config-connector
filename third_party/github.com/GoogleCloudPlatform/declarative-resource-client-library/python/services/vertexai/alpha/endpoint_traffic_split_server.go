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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/alpha/vertexai_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/alpha"
)

// EndpointTrafficSplitServer implements the gRPC interface for EndpointTrafficSplit.
type EndpointTrafficSplitServer struct{}

// ProtoToEndpointTrafficSplitTrafficSplit converts a EndpointTrafficSplitTrafficSplit object from its proto representation.
func ProtoToVertexaiAlphaEndpointTrafficSplitTrafficSplit(p *alphapb.VertexaiAlphaEndpointTrafficSplitTrafficSplit) *alpha.EndpointTrafficSplitTrafficSplit {
	if p == nil {
		return nil
	}
	obj := &alpha.EndpointTrafficSplitTrafficSplit{
		DeployedModelId:   dcl.StringOrNil(p.GetDeployedModelId()),
		TrafficPercentage: dcl.Int64OrNil(p.GetTrafficPercentage()),
	}
	return obj
}

// ProtoToEndpointTrafficSplit converts a EndpointTrafficSplit resource from its proto representation.
func ProtoToEndpointTrafficSplit(p *alphapb.VertexaiAlphaEndpointTrafficSplit) *alpha.EndpointTrafficSplit {
	obj := &alpha.EndpointTrafficSplit{
		Endpoint: dcl.StringOrNil(p.GetEndpoint()),
		Project:  dcl.StringOrNil(p.GetProject()),
		Location: dcl.StringOrNil(p.GetLocation()),
		Etag:     dcl.StringOrNil(p.GetEtag()),
	}
	for _, r := range p.GetTrafficSplit() {
		obj.TrafficSplit = append(obj.TrafficSplit, *ProtoToVertexaiAlphaEndpointTrafficSplitTrafficSplit(r))
	}
	return obj
}

// EndpointTrafficSplitTrafficSplitToProto converts a EndpointTrafficSplitTrafficSplit object to its proto representation.
func VertexaiAlphaEndpointTrafficSplitTrafficSplitToProto(o *alpha.EndpointTrafficSplitTrafficSplit) *alphapb.VertexaiAlphaEndpointTrafficSplitTrafficSplit {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexaiAlphaEndpointTrafficSplitTrafficSplit{}
	p.SetDeployedModelId(dcl.ValueOrEmptyString(o.DeployedModelId))
	p.SetTrafficPercentage(dcl.ValueOrEmptyInt64(o.TrafficPercentage))
	return p
}

// EndpointTrafficSplitToProto converts a EndpointTrafficSplit resource to its proto representation.
func EndpointTrafficSplitToProto(resource *alpha.EndpointTrafficSplit) *alphapb.VertexaiAlphaEndpointTrafficSplit {
	p := &alphapb.VertexaiAlphaEndpointTrafficSplit{}
	p.SetEndpoint(dcl.ValueOrEmptyString(resource.Endpoint))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	sTrafficSplit := make([]*alphapb.VertexaiAlphaEndpointTrafficSplitTrafficSplit, len(resource.TrafficSplit))
	for i, r := range resource.TrafficSplit {
		sTrafficSplit[i] = VertexaiAlphaEndpointTrafficSplitTrafficSplitToProto(&r)
	}
	p.SetTrafficSplit(sTrafficSplit)

	return p
}

// applyEndpointTrafficSplit handles the gRPC request by passing it to the underlying EndpointTrafficSplit Apply() method.
func (s *EndpointTrafficSplitServer) applyEndpointTrafficSplit(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVertexaiAlphaEndpointTrafficSplitRequest) (*alphapb.VertexaiAlphaEndpointTrafficSplit, error) {
	p := ProtoToEndpointTrafficSplit(request.GetResource())
	res, err := c.ApplyEndpointTrafficSplit(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EndpointTrafficSplitToProto(res)
	return r, nil
}

// applyVertexaiAlphaEndpointTrafficSplit handles the gRPC request by passing it to the underlying EndpointTrafficSplit Apply() method.
func (s *EndpointTrafficSplitServer) ApplyVertexaiAlphaEndpointTrafficSplit(ctx context.Context, request *alphapb.ApplyVertexaiAlphaEndpointTrafficSplitRequest) (*alphapb.VertexaiAlphaEndpointTrafficSplit, error) {
	cl, err := createConfigEndpointTrafficSplit(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEndpointTrafficSplit(ctx, cl, request)
}

// DeleteEndpointTrafficSplit handles the gRPC request by passing it to the underlying EndpointTrafficSplit Delete() method.
func (s *EndpointTrafficSplitServer) DeleteVertexaiAlphaEndpointTrafficSplit(ctx context.Context, request *alphapb.DeleteVertexaiAlphaEndpointTrafficSplitRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEndpointTrafficSplit(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEndpointTrafficSplit(ctx, ProtoToEndpointTrafficSplit(request.GetResource()))

}

// ListVertexaiAlphaEndpointTrafficSplit is a no-op method because EndpointTrafficSplit has no list method.
func (s *EndpointTrafficSplitServer) ListVertexaiAlphaEndpointTrafficSplit(_ context.Context, _ *alphapb.ListVertexaiAlphaEndpointTrafficSplitRequest) (*alphapb.ListVertexaiAlphaEndpointTrafficSplitResponse, error) {
	return nil, nil
}

func createConfigEndpointTrafficSplit(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
