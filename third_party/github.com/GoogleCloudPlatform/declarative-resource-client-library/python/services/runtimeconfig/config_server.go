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
	runtimeconfigpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/runtimeconfig/runtimeconfig_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/runtimeconfig"
)

// Server implements the gRPC interface for Config.
type ConfigServer struct{}

// ProtoToConfig converts a Config resource from its proto representation.
func ProtoToConfig(p *runtimeconfigpb.RuntimeconfigConfig) *runtimeconfig.Config {
	obj := &runtimeconfig.Config{
		Name:        dcl.StringOrNil(p.Name),
		Description: dcl.StringOrNil(p.Description),
		Project:     dcl.StringOrNil(p.Project),
	}
	return obj
}

// ConfigToProto converts a Config resource to its proto representation.
func ConfigToProto(resource *runtimeconfig.Config) *runtimeconfigpb.RuntimeconfigConfig {
	p := &runtimeconfigpb.RuntimeconfigConfig{
		Name:        dcl.ValueOrEmptyString(resource.Name),
		Description: dcl.ValueOrEmptyString(resource.Description),
		Project:     dcl.ValueOrEmptyString(resource.Project),
	}

	return p
}

// ApplyConfig handles the gRPC request by passing it to the underlying Config Apply() method.
func (s *ConfigServer) applyConfig(ctx context.Context, c *runtimeconfig.Client, request *runtimeconfigpb.ApplyRuntimeconfigConfigRequest) (*runtimeconfigpb.RuntimeconfigConfig, error) {
	p := ProtoToConfig(request.GetResource())
	res, err := c.ApplyConfig(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConfigToProto(res)
	return r, nil
}

// ApplyConfig handles the gRPC request by passing it to the underlying Config Apply() method.
func (s *ConfigServer) ApplyRuntimeconfigConfig(ctx context.Context, request *runtimeconfigpb.ApplyRuntimeconfigConfigRequest) (*runtimeconfigpb.RuntimeconfigConfig, error) {
	cl, err := createConfigConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyConfig(ctx, cl, request)
}

// DeleteConfig handles the gRPC request by passing it to the underlying Config Delete() method.
func (s *ConfigServer) DeleteRuntimeconfigConfig(ctx context.Context, request *runtimeconfigpb.DeleteRuntimeconfigConfigRequest) (*emptypb.Empty, error) {

	cl, err := createConfigConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteConfig(ctx, ProtoToConfig(request.GetResource()))

}

// ListConfig handles the gRPC request by passing it to the underlying ConfigList() method.
func (s *ConfigServer) ListRuntimeconfigConfig(ctx context.Context, request *runtimeconfigpb.ListRuntimeconfigConfigRequest) (*runtimeconfigpb.ListRuntimeconfigConfigResponse, error) {
	cl, err := createConfigConfig(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListConfig(ctx, request.Project, request.Name)
	if err != nil {
		return nil, err
	}
	var protos []*runtimeconfigpb.RuntimeconfigConfig
	for _, r := range resources.Items {
		rp := ConfigToProto(r)
		protos = append(protos, rp)
	}
	return &runtimeconfigpb.ListRuntimeconfigConfigResponse{Items: protos}, nil
}

func createConfigConfig(ctx context.Context, service_account_file string) (*runtimeconfig.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return runtimeconfig.NewClient(conf), nil
}
