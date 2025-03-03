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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/beta/apigee_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/beta"
)

// EnvgroupServer implements the gRPC interface for Envgroup.
type EnvgroupServer struct{}

// ProtoToEnvgroupStateEnum converts a EnvgroupStateEnum enum from its proto representation.
func ProtoToApigeeBetaEnvgroupStateEnum(e betapb.ApigeeBetaEnvgroupStateEnum) *beta.EnvgroupStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ApigeeBetaEnvgroupStateEnum_name[int32(e)]; ok {
		e := beta.EnvgroupStateEnum(n[len("ApigeeBetaEnvgroupStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvgroup converts a Envgroup resource from its proto representation.
func ProtoToEnvgroup(p *betapb.ApigeeBetaEnvgroup) *beta.Envgroup {
	obj := &beta.Envgroup{
		Name:               dcl.StringOrNil(p.GetName()),
		CreatedAt:          dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:     dcl.Int64OrNil(p.GetLastModifiedAt()),
		State:              ProtoToApigeeBetaEnvgroupStateEnum(p.GetState()),
		ApigeeOrganization: dcl.StringOrNil(p.GetApigeeOrganization()),
	}
	for _, r := range p.GetHostnames() {
		obj.Hostnames = append(obj.Hostnames, r)
	}
	return obj
}

// EnvgroupStateEnumToProto converts a EnvgroupStateEnum enum to its proto representation.
func ApigeeBetaEnvgroupStateEnumToProto(e *beta.EnvgroupStateEnum) betapb.ApigeeBetaEnvgroupStateEnum {
	if e == nil {
		return betapb.ApigeeBetaEnvgroupStateEnum(0)
	}
	if v, ok := betapb.ApigeeBetaEnvgroupStateEnum_value["EnvgroupStateEnum"+string(*e)]; ok {
		return betapb.ApigeeBetaEnvgroupStateEnum(v)
	}
	return betapb.ApigeeBetaEnvgroupStateEnum(0)
}

// EnvgroupToProto converts a Envgroup resource to its proto representation.
func EnvgroupToProto(resource *beta.Envgroup) *betapb.ApigeeBetaEnvgroup {
	p := &betapb.ApigeeBetaEnvgroup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetState(ApigeeBetaEnvgroupStateEnumToProto(resource.State))
	p.SetApigeeOrganization(dcl.ValueOrEmptyString(resource.ApigeeOrganization))
	sHostnames := make([]string, len(resource.Hostnames))
	for i, r := range resource.Hostnames {
		sHostnames[i] = r
	}
	p.SetHostnames(sHostnames)

	return p
}

// applyEnvgroup handles the gRPC request by passing it to the underlying Envgroup Apply() method.
func (s *EnvgroupServer) applyEnvgroup(ctx context.Context, c *beta.Client, request *betapb.ApplyApigeeBetaEnvgroupRequest) (*betapb.ApigeeBetaEnvgroup, error) {
	p := ProtoToEnvgroup(request.GetResource())
	res, err := c.ApplyEnvgroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvgroupToProto(res)
	return r, nil
}

// applyApigeeBetaEnvgroup handles the gRPC request by passing it to the underlying Envgroup Apply() method.
func (s *EnvgroupServer) ApplyApigeeBetaEnvgroup(ctx context.Context, request *betapb.ApplyApigeeBetaEnvgroupRequest) (*betapb.ApigeeBetaEnvgroup, error) {
	cl, err := createConfigEnvgroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEnvgroup(ctx, cl, request)
}

// DeleteEnvgroup handles the gRPC request by passing it to the underlying Envgroup Delete() method.
func (s *EnvgroupServer) DeleteApigeeBetaEnvgroup(ctx context.Context, request *betapb.DeleteApigeeBetaEnvgroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvgroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvgroup(ctx, ProtoToEnvgroup(request.GetResource()))

}

// ListApigeeBetaEnvgroup handles the gRPC request by passing it to the underlying EnvgroupList() method.
func (s *EnvgroupServer) ListApigeeBetaEnvgroup(ctx context.Context, request *betapb.ListApigeeBetaEnvgroupRequest) (*betapb.ListApigeeBetaEnvgroupResponse, error) {
	cl, err := createConfigEnvgroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvgroup(ctx, request.GetApigeeOrganization())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ApigeeBetaEnvgroup
	for _, r := range resources.Items {
		rp := EnvgroupToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListApigeeBetaEnvgroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEnvgroup(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
