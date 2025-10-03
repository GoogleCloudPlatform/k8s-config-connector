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
	apigeepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/apigee/apigee_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/apigee"
)

// EnvironmentGroupServer implements the gRPC interface for EnvironmentGroup.
type EnvironmentGroupServer struct{}

// ProtoToEnvironmentGroupStateEnum converts a EnvironmentGroupStateEnum enum from its proto representation.
func ProtoToApigeeEnvironmentGroupStateEnum(e apigeepb.ApigeeEnvironmentGroupStateEnum) *apigee.EnvironmentGroupStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := apigeepb.ApigeeEnvironmentGroupStateEnum_name[int32(e)]; ok {
		e := apigee.EnvironmentGroupStateEnum(n[len("ApigeeEnvironmentGroupStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToEnvironmentGroup converts a EnvironmentGroup resource from its proto representation.
func ProtoToEnvironmentGroup(p *apigeepb.ApigeeEnvironmentGroup) *apigee.EnvironmentGroup {
	obj := &apigee.EnvironmentGroup{
		Name:               dcl.StringOrNil(p.GetName()),
		CreatedAt:          dcl.Int64OrNil(p.GetCreatedAt()),
		LastModifiedAt:     dcl.Int64OrNil(p.GetLastModifiedAt()),
		State:              ProtoToApigeeEnvironmentGroupStateEnum(p.GetState()),
		ApigeeOrganization: dcl.StringOrNil(p.GetApigeeOrganization()),
	}
	for _, r := range p.GetHostnames() {
		obj.Hostnames = append(obj.Hostnames, r)
	}
	return obj
}

// EnvironmentGroupStateEnumToProto converts a EnvironmentGroupStateEnum enum to its proto representation.
func ApigeeEnvironmentGroupStateEnumToProto(e *apigee.EnvironmentGroupStateEnum) apigeepb.ApigeeEnvironmentGroupStateEnum {
	if e == nil {
		return apigeepb.ApigeeEnvironmentGroupStateEnum(0)
	}
	if v, ok := apigeepb.ApigeeEnvironmentGroupStateEnum_value["EnvironmentGroupStateEnum"+string(*e)]; ok {
		return apigeepb.ApigeeEnvironmentGroupStateEnum(v)
	}
	return apigeepb.ApigeeEnvironmentGroupStateEnum(0)
}

// EnvironmentGroupToProto converts a EnvironmentGroup resource to its proto representation.
func EnvironmentGroupToProto(resource *apigee.EnvironmentGroup) *apigeepb.ApigeeEnvironmentGroup {
	p := &apigeepb.ApigeeEnvironmentGroup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreatedAt(dcl.ValueOrEmptyInt64(resource.CreatedAt))
	p.SetLastModifiedAt(dcl.ValueOrEmptyInt64(resource.LastModifiedAt))
	p.SetState(ApigeeEnvironmentGroupStateEnumToProto(resource.State))
	p.SetApigeeOrganization(dcl.ValueOrEmptyString(resource.ApigeeOrganization))
	sHostnames := make([]string, len(resource.Hostnames))
	for i, r := range resource.Hostnames {
		sHostnames[i] = r
	}
	p.SetHostnames(sHostnames)

	return p
}

// applyEnvironmentGroup handles the gRPC request by passing it to the underlying EnvironmentGroup Apply() method.
func (s *EnvironmentGroupServer) applyEnvironmentGroup(ctx context.Context, c *apigee.Client, request *apigeepb.ApplyApigeeEnvironmentGroupRequest) (*apigeepb.ApigeeEnvironmentGroup, error) {
	p := ProtoToEnvironmentGroup(request.GetResource())
	res, err := c.ApplyEnvironmentGroup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := EnvironmentGroupToProto(res)
	return r, nil
}

// applyApigeeEnvironmentGroup handles the gRPC request by passing it to the underlying EnvironmentGroup Apply() method.
func (s *EnvironmentGroupServer) ApplyApigeeEnvironmentGroup(ctx context.Context, request *apigeepb.ApplyApigeeEnvironmentGroupRequest) (*apigeepb.ApigeeEnvironmentGroup, error) {
	cl, err := createConfigEnvironmentGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyEnvironmentGroup(ctx, cl, request)
}

// DeleteEnvironmentGroup handles the gRPC request by passing it to the underlying EnvironmentGroup Delete() method.
func (s *EnvironmentGroupServer) DeleteApigeeEnvironmentGroup(ctx context.Context, request *apigeepb.DeleteApigeeEnvironmentGroupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigEnvironmentGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteEnvironmentGroup(ctx, ProtoToEnvironmentGroup(request.GetResource()))

}

// ListApigeeEnvironmentGroup handles the gRPC request by passing it to the underlying EnvironmentGroupList() method.
func (s *EnvironmentGroupServer) ListApigeeEnvironmentGroup(ctx context.Context, request *apigeepb.ListApigeeEnvironmentGroupRequest) (*apigeepb.ListApigeeEnvironmentGroupResponse, error) {
	cl, err := createConfigEnvironmentGroup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListEnvironmentGroup(ctx, request.GetApigeeOrganization())
	if err != nil {
		return nil, err
	}
	var protos []*apigeepb.ApigeeEnvironmentGroup
	for _, r := range resources.Items {
		rp := EnvironmentGroupToProto(r)
		protos = append(protos, rp)
	}
	p := &apigeepb.ListApigeeEnvironmentGroupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigEnvironmentGroup(ctx context.Context, service_account_file string) (*apigee.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return apigee.NewClient(conf), nil
}
