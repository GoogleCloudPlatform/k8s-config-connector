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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/alpha/apigee_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/alpha"
)

// EnvironmentServer implements the gRPC interface for Environment.
type EnvironmentServer struct{}

// ProtoToEnvironmentStateEnum converts a EnvironmentStateEnum enum from its proto representation.
func ProtoToApigeeAlphaEnvironmentStateEnum(e alphapb.ApigeeAlphaEnvironmentStateEnum) *alpha.EnvironmentStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ApigeeAlphaEnvironmentStateEnum_name[int32(e)]; ok {
		e := alpha.EnvironmentStateEnum(n[len("ApigeeAlphaEnvironmentStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvironment converts a Environment resource from its proto representation.
func ProtoToEnvironment(p *alphapb.ApigeeAlphaEnvironment) *alpha.Environment {
	obj := &alpha.Environment{
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		CreatedAt:          dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:     dcl.Int64OrNil(p.GetLastModifiedAt()),
		DisplayName:        dcl.StringOrNil(p.GetDisplayName()),
		State:              ProtoToApigeeAlphaEnvironmentStateEnum(p.GetState()),
		ApigeeOrganization: dcl.StringOrNil(p.GetApigeeOrganization()),
	}
	return obj
}

// EnvironmentStateEnumToProto converts a EnvironmentStateEnum enum to its proto representation.
func ApigeeAlphaEnvironmentStateEnumToProto(e *alpha.EnvironmentStateEnum) alphapb.ApigeeAlphaEnvironmentStateEnum {
	if e == nil {
		return alphapb.ApigeeAlphaEnvironmentStateEnum(0)
	}
	if v, ok := alphapb.ApigeeAlphaEnvironmentStateEnum_value["EnvironmentStateEnum"+string(*e)]; ok {
		return alphapb.ApigeeAlphaEnvironmentStateEnum(v)
	}
	return alphapb.ApigeeAlphaEnvironmentStateEnum(0)
}

// EnvironmentToProto converts a Environment resource to its proto representation.
func EnvironmentToProto(resource *alpha.Environment) *alphapb.ApigeeAlphaEnvironment {
	p := &alphapb.ApigeeAlphaEnvironment{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetDisplayName(dcl.ValueOrEmptyString(resource.DisplayName))
	p.SetState(ApigeeAlphaEnvironmentStateEnumToProto(resource.State))
	p.SetApigeeOrganization(dcl.ValueOrEmptyString(resource.ApigeeOrganization))
	mProperties := make(map[string]string, len(resource.Properties))
	for k, r := range resource.Properties {
		mProperties[k] = r
	}
	p.SetProperties(mProperties)

	return p
}

// applyEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) applyEnvironment(ctx context.Context, c *alpha.Client, request *alphapb.ApplyApigeeAlphaEnvironmentRequest) (*alphapb.ApigeeAlphaEnvironment, error) {
	p := ProtoToEnvironment(request.GetResource())
	res, err := c.ApplyEnvironment(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvironmentToProto(res)
	return r, nil
}

// applyApigeeAlphaEnvironment handles the gRPC request by passing it to the underlying Environment Apply() method.
func (s *EnvironmentServer) ApplyApigeeAlphaEnvironment(ctx context.Context, request *alphapb.ApplyApigeeAlphaEnvironmentRequest) (*alphapb.ApigeeAlphaEnvironment, error) {
	cl, err := createConfigEnvironment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEnvironment(ctx, cl, request)
}

// DeleteEnvironment handles the gRPC request by passing it to the underlying Environment Delete() method.
func (s *EnvironmentServer) DeleteApigeeAlphaEnvironment(ctx context.Context, request *alphapb.DeleteApigeeAlphaEnvironmentRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvironment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvironment(ctx, ProtoToEnvironment(request.GetResource()))

}

// ListApigeeAlphaEnvironment handles the gRPC request by passing it to the underlying EnvironmentList() method.
func (s *EnvironmentServer) ListApigeeAlphaEnvironment(ctx context.Context, request *alphapb.ListApigeeAlphaEnvironmentRequest) (*alphapb.ListApigeeAlphaEnvironmentResponse, error) {
	cl, err := createConfigEnvironment(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvironment(ctx, request.GetApigeeOrganization())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ApigeeAlphaEnvironment
	for _, r := range resources.Items {
		rp := EnvironmentToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListApigeeAlphaEnvironmentResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEnvironment(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
