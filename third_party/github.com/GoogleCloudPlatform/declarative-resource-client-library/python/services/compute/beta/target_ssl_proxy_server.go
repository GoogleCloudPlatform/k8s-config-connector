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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for TargetSslProxy.
type TargetSslProxyServer struct{}

// ProtoToTargetSslProxyProxyHeaderEnum converts a TargetSslProxyProxyHeaderEnum enum from its proto representation.
func ProtoToComputeBetaTargetSslProxyProxyHeaderEnum(e betapb.ComputeBetaTargetSslProxyProxyHeaderEnum) *beta.TargetSslProxyProxyHeaderEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaTargetSslProxyProxyHeaderEnum_name[int32(e)]; ok {
		e := beta.TargetSslProxyProxyHeaderEnum(n[len("ComputeBetaTargetSslProxyProxyHeaderEnum"):])
		return &e
	}
	return nil
}

// ProtoToTargetSslProxy converts a TargetSslProxy resource from its proto representation.
func ProtoToTargetSslProxy(p *betapb.ComputeBetaTargetSslProxy) *beta.TargetSslProxy {
	obj := &beta.TargetSslProxy{
		Id:          dcl.Int64OrNil(p.Id),
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		SelfLink:    dcl.StringOrNil(p.SelfLink),
		Service:     dcl.StringOrNil(p.Service),
		ProxyHeader: ProtoToComputeBetaTargetSslProxyProxyHeaderEnum(p.GetProxyHeader()),
		SslPolicy:   dcl.StringOrNil(p.SslPolicy),
		Project:     dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetSslCertificates() {
		obj.SslCertificates = append(obj.SslCertificates, r)
	}
	return obj
}

// TargetSslProxyProxyHeaderEnumToProto converts a TargetSslProxyProxyHeaderEnum enum to its proto representation.
func ComputeBetaTargetSslProxyProxyHeaderEnumToProto(e *beta.TargetSslProxyProxyHeaderEnum) betapb.ComputeBetaTargetSslProxyProxyHeaderEnum {
	if e == nil {
		return betapb.ComputeBetaTargetSslProxyProxyHeaderEnum(0)
	}
	if v, ok := betapb.ComputeBetaTargetSslProxyProxyHeaderEnum_value["TargetSslProxyProxyHeaderEnum"+string(*e)]; ok {
		return betapb.ComputeBetaTargetSslProxyProxyHeaderEnum(v)
	}
	return betapb.ComputeBetaTargetSslProxyProxyHeaderEnum(0)
}

// TargetSslProxyToProto converts a TargetSslProxy resource to its proto representation.
func TargetSslProxyToProto(resource *beta.TargetSslProxy) *betapb.ComputeBetaTargetSslProxy {
	p := &betapb.ComputeBetaTargetSslProxy{
		Id:          dcl.ValueOrEmptyInt64(resource.Id),
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		SelfLink:    dcl.ValueOrEmptyString(resource.SelfLink),
		Service:     dcl.ValueOrEmptyString(resource.Service),
		ProxyHeader: ComputeBetaTargetSslProxyProxyHeaderEnumToProto(resource.ProxyHeader),
		SslPolicy:   dcl.ValueOrEmptyString(resource.SslPolicy),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.SslCertificates {
		p.SslCertificates = append(p.SslCertificates, r)
	}

	return p
}

// ApplyTargetSslProxy handles the gRPC request by passing it to the underlying TargetSslProxy Apply() method.
func (s *TargetSslProxyServer) applyTargetSslProxy(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaTargetSslProxyRequest) (*betapb.ComputeBetaTargetSslProxy, error) {
	p := ProtoToTargetSslProxy(request.GetResource())
	res, err := c.ApplyTargetSslProxy(ctx, p)
	if err != nil {
		return nil, err
	}
	r := TargetSslProxyToProto(res)
	return r, nil
}

// ApplyTargetSslProxy handles the gRPC request by passing it to the underlying TargetSslProxy Apply() method.
func (s *TargetSslProxyServer) ApplyComputeBetaTargetSslProxy(ctx context.Context, request *betapb.ApplyComputeBetaTargetSslProxyRequest) (*betapb.ComputeBetaTargetSslProxy, error) {
	cl, err := createConfigTargetSslProxy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyTargetSslProxy(ctx, cl, request)
}

// DeleteTargetSslProxy handles the gRPC request by passing it to the underlying TargetSslProxy Delete() method.
func (s *TargetSslProxyServer) DeleteComputeBetaTargetSslProxy(ctx context.Context, request *betapb.DeleteComputeBetaTargetSslProxyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigTargetSslProxy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteTargetSslProxy(ctx, ProtoToTargetSslProxy(request.GetResource()))

}

// ListComputeBetaTargetSslProxy handles the gRPC request by passing it to the underlying TargetSslProxyList() method.
func (s *TargetSslProxyServer) ListComputeBetaTargetSslProxy(ctx context.Context, request *betapb.ListComputeBetaTargetSslProxyRequest) (*betapb.ListComputeBetaTargetSslProxyResponse, error) {
	cl, err := createConfigTargetSslProxy(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListTargetSslProxy(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaTargetSslProxy
	for _, r := range resources.Items {
		rp := TargetSslProxyToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaTargetSslProxyResponse{Items: protos}, nil
}

func createConfigTargetSslProxy(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
