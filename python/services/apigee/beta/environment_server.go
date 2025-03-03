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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/beta/apigee_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/beta"
)

// EnvironmentServer implements the gRPC interface for Environment.
type EnvironmentServer struct{}

// ProtoToEnvironmentStateEnum converts a EnvironmentStateEnum enum from its proto representation.
func ProtoToApigeeBetaEnvironmentStateEnum(e betapb.ApigeeBetaEnvironmentStateEnum) *beta.EnvironmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ApigeeBetaEnvironmentStateEnum_name[int32(e)]; ok {
		e := beta.EnvironmentStateEnum(n[len("ApigeeBetaEnvironmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvironment converts a Environment resource from its proto representation.
func ProtoToEnvironment(p *betapb.ApigeeBetaEnvironment) *beta.Environment {
	obj := &beta.Environment{
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		CreatedAt:          dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:     dcl.Int64OrNil(p.GetLastModifiedAt()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		State:              ProtoToApigeeBetaEnvironmentStateEnum(p.GetState()),
		ApigeeOrganization: dcl.StringOrNil(p.GetApigeeOrganization()),
	}
	return obj
}

// EnvironmentStateEnumToProto converts a EnvironmentStateEnum enum to its proto representation.
func ApigeeBetaEnvironmentStateEnumToProto(e *beta.EnvironmentStateEnum) betapb.ApigeeBetaEnvironmentStateEnum {
	if e == nil {
		return betapb.ApigeeBetaEnvironmentStateEnum(0)
	}
	if v, ok := betapb.ApigeeBetaEnvironmentStateEnum_value["EnvironmentStateEnum"+string(*e)]; ok {
		return betapb.ApigeeBetaEnvironmentStateEnum(v)
	}
	return betapb.ApigeeBetaEnvironmentStateEnum(0)
}

// EnvironmentToProto converts a Environment resource to its proto representation.
func EnvironmentToProto(resource *beta.Environment) *betapb.ApigeeBetaEnvironment {
	p := &betapb.ApigeeBetaEnvironment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetState(ApigeeBetaEnvironmentStateEnumToProto(resource.State))
	p.SetApigeeOrganization(dcl.ValueOrEmptyString(resource.ApigeeOrganization))
	mProperties := make(map[string]string, len(resource.Properties))
	for k, r := range resource.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)

	return p
}

// applyEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) applyEnvironment(ctx context.Context, c *beta.Client, request *betapb.ApplyApigeeBetaEnvironmentRequest) (*betapb.ApigeeBetaEnvironment, error) {
	p := ProtoToEnvironment(request.GetResource())
	res, err := c.ApplyEnvironment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvironmentToProto(res)
	return r, nil
}

// applyApigeeBetaEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) ApplyApigeeBetaEnvironment(ctx context.Context, request *betapb.ApplyApigeeBetaEnvironmentRequest) (*betapb.ApigeeBetaEnvironment, error) {
	cl, err := createConfigEnvironment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEnvironment(ctx, cl, request)
}

// DeleteEnvironment handles the gRPC request by passing it to the underlying Environment Delete() method.
func (s *EnvironmentServer) DeleteApigeeBetaEnvironment(ctx context.Context, request *betapb.DeleteApigeeBetaEnvironmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvironment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvironment(ctx, ProtoToEnvironment(request.GetResource()))

}

// ListApigeeBetaEnvironment handles the gRPC request by passing it to the underlying EnvironmentList() method.
func (s *EnvironmentServer) ListApigeeBetaEnvironment(ctx context.Context, request *betapb.ListApigeeBetaEnvironmentRequest) (*betapb.ListApigeeBetaEnvironmentResponse, error) {
	cl, err := createConfigEnvironment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvironment(ctx, request.GetApigeeOrganization())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ApigeeBetaEnvironment
	for _, r := range resources.Items {
		rp := EnvironmentToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListApigeeBetaEnvironmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEnvironment(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
