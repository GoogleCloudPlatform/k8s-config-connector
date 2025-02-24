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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/alpha/apigee_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee/alpha"
)

// EnvironmentGroupServer implements the gRPC interface for EnvironmentGroup.
type EnvironmentGroupServer struct{}

// ProtoToEnvironmentGroupStateEnum converts a EnvironmentGroupStateEnum enum from its proto representation.
func ProtoToApigeeAlphaEnvironmentGroupStateEnum(e alphapb.ApigeeAlphaEnvironmentGroupStateEnum) *alpha.EnvironmentGroupStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.ApigeeAlphaEnvironmentGroupStateEnum_name[int32(e)]; ok {
		e := alpha.EnvironmentGroupStateEnum(n[len("ApigeeAlphaEnvironmentGroupStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvironmentGroup converts a EnvironmentGroup resource from its proto representation.
func ProtoToEnvironmentGroup(p *alphapb.ApigeeAlphaEnvironmentGroup) *alpha.EnvironmentGroup {
	obj := &alpha.EnvironmentGroup{
		Name:               dcl.StringOrNil(p.GetName()),
		CreatedAt:          dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:     dcl.Int64OrNil(p.GetLastModifiedAt()),
		State:              ProtoToApigeeAlphaEnvironmentGroupStateEnum(p.GetState()),
		ApigeeOrganization: dcl.StringOrNil(p.GetApigeeOrganization()),
	}
	for _, r := range p.GetHostnames() {
		obj.Hostnames = append(obj.Hostnames, r)
	}
	return obj
}

// EnvironmentGroupStateEnumToProto converts a EnvironmentGroupStateEnum enum to its proto representation.
func ApigeeAlphaEnvironmentGroupStateEnumToProto(e *alpha.EnvironmentGroupStateEnum) alphapb.ApigeeAlphaEnvironmentGroupStateEnum {
	if e == nil {
		return alphapb.ApigeeAlphaEnvironmentGroupStateEnum(0)
	}
	if v, ok := alphapb.ApigeeAlphaEnvironmentGroupStateEnum_value["EnvironmentGroupStateEnum"+string(*e)]; ok {
		return alphapb.ApigeeAlphaEnvironmentGroupStateEnum(v)
	}
	return alphapb.ApigeeAlphaEnvironmentGroupStateEnum(0)
}

// EnvironmentGroupToProto converts a EnvironmentGroup resource to its proto representation.
func EnvironmentGroupToProto(resource *alpha.EnvironmentGroup) *alphapb.ApigeeAlphaEnvironmentGroup {
	p := &alphapb.ApigeeAlphaEnvironmentGroup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetState(ApigeeAlphaEnvironmentGroupStateEnumToProto(resource.State))
	p.SetApigeeOrganization(dcl.ValueOrEmptyString(resource.ApigeeOrganization))
	sHostnames := make([]string, len(resource.Hostnames))
	for i, r := range resource.Hostnames {
		sHostnames[i] = r
	}
	p.SetHostnames(sHostnames)

	return p
}

// applyEnvironmentGroup handles the gRPC request by passing it to the underlying EnvironmentGroup Apply() method.
func (s *EnvironmentGroupServer) applyEnvironmentGroup(ctx context.Context, c *alpha.Client, request *alphapb.ApplyApigeeAlphaEnvironmentGroupRequest) (*alphapb.ApigeeAlphaEnvironmentGroup, error) {
	p := ProtoToEnvironmentGroup(request.GetResource())
	res, err := c.ApplyEnvironmentGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvironmentGroupToProto(res)
	return r, nil
}

// applyApigeeAlphaEnvironmentGroup handles the gRPC request by passing it to the underlying EnvironmentGroup Apply() method.
func (s *EnvironmentGroupServer) ApplyApigeeAlphaEnvironmentGroup(ctx context.Context, request *alphapb.ApplyApigeeAlphaEnvironmentGroupRequest) (*alphapb.ApigeeAlphaEnvironmentGroup, error) {
	cl, err := createConfigEnvironmentGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEnvironmentGroup(ctx, cl, request)
}

// DeleteEnvironmentGroup handles the gRPC request by passing it to the underlying EnvironmentGroup Delete() method.
func (s *EnvironmentGroupServer) DeleteApigeeAlphaEnvironmentGroup(ctx context.Context, request *alphapb.DeleteApigeeAlphaEnvironmentGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvironmentGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvironmentGroup(ctx, ProtoToEnvironmentGroup(request.GetResource()))

}

// ListApigeeAlphaEnvironmentGroup handles the gRPC request by passing it to the underlying EnvironmentGroupList() method.
func (s *EnvironmentGroupServer) ListApigeeAlphaEnvironmentGroup(ctx context.Context, request *alphapb.ListApigeeAlphaEnvironmentGroupRequest) (*alphapb.ListApigeeAlphaEnvironmentGroupResponse, error) {
	cl, err := createConfigEnvironmentGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvironmentGroup(ctx, request.GetApigeeOrganization())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.ApigeeAlphaEnvironmentGroup
	for _, r := range resources.Items {
		rp := EnvironmentGroupToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListApigeeAlphaEnvironmentGroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEnvironmentGroup(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
