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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/iam/beta/iam_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/iam/beta"
)

// RoleServer implements the gRPC interface for Role.
type RoleServer struct{}

// ProtoToRoleStageEnum converts a RoleStageEnum enum from its proto representation.
func ProtoToIamBetaRoleStageEnum(e betapb.IamBetaRoleStageEnum) *beta.RoleStageEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.IamBetaRoleStageEnum_name[int32(e)]; ok {
		e := beta.RoleStageEnum(n[len("IamBetaRoleStageEnum"):])
		return &e
	}
	return nil
}

// ProtoToRoleLocalizedValues converts a RoleLocalizedValues object from its proto representation.
func ProtoToIamBetaRoleLocalizedValues(p *betapb.IamBetaRoleLocalizedValues) *beta.RoleLocalizedValues {
	if p == nil {
		return nil
	}
	obj := &beta.RoleLocalizedValues{
		LocalizedTitle:       dcl.StringOrNil(p.GetLocalizedTitle()),
		LocalizedDescription: dcl.StringOrNil(p.GetLocalizedDescription()),
	}
	return obj
}

// ProtoToRole converts a Role resource from its proto representation.
func ProtoToRole(p *betapb.IamBetaRole) *beta.Role {
	obj := &beta.Role{
		Name:            dcl.StringOrNil(p.GetName()),
		Title:           dcl.StringOrNil(p.GetTitle()),
		Description:     dcl.StringOrNil(p.GetDescription()),
		LocalizedValues: ProtoToIamBetaRoleLocalizedValues(p.GetLocalizedValues()),
		LifecyclePhase:  dcl.StringOrNil(p.GetLifecyclePhase()),
		GroupName:       dcl.StringOrNil(p.GetGroupName()),
		GroupTitle:      dcl.StringOrNil(p.GetGroupTitle()),
		Stage:           ProtoToIamBetaRoleStageEnum(p.GetStage()),
		Etag:            dcl.StringOrNil(p.GetEtag()),
		Deleted:         dcl.Bool(p.GetDeleted()),
		Parent:          dcl.StringOrNil(p.GetParent()),
	}
	for _, r := range p.GetIncludedPermissions() {
		obj.IncludedPermissions = append(obj.IncludedPermissions, r)
	}
	for _, r := range p.GetIncludedRoles() {
		obj.IncludedRoles = append(obj.IncludedRoles, r)
	}
	return obj
}

// RoleStageEnumToProto converts a RoleStageEnum enum to its proto representation.
func IamBetaRoleStageEnumToProto(e *beta.RoleStageEnum) betapb.IamBetaRoleStageEnum {
	if e == nil {
		return betapb.IamBetaRoleStageEnum(0)
	}
	if v, ok := betapb.IamBetaRoleStageEnum_value["RoleStageEnum"+string(*e)]; ok {
		return betapb.IamBetaRoleStageEnum(v)
	}
	return betapb.IamBetaRoleStageEnum(0)
}

// RoleLocalizedValuesToProto converts a RoleLocalizedValues object to its proto representation.
func IamBetaRoleLocalizedValuesToProto(o *beta.RoleLocalizedValues) *betapb.IamBetaRoleLocalizedValues {
	if o == nil {
		return nil
	}
	p := &betapb.IamBetaRoleLocalizedValues{}
	p.SetLocalizedTitle(dcl.ValueOrEmptyString(o.LocalizedTitle))
	p.SetLocalizedDescription(dcl.ValueOrEmptyString(o.LocalizedDescription))
	return p
}

// RoleToProto converts a Role resource to its proto representation.
func RoleToProto(resource *beta.Role) *betapb.IamBetaRole {
	p := &betapb.IamBetaRole{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetTitle(dcl.ValueOrEmptyString(resource.Title))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetLocalizedValues(IamBetaRoleLocalizedValuesToProto(resource.LocalizedValues))
	p.SetLifecyclePhase(dcl.ValueOrEmptyString(resource.LifecyclePhase))
	p.SetGroupName(dcl.ValueOrEmptyString(resource.GroupName))
	p.SetGroupTitle(dcl.ValueOrEmptyString(resource.GroupTitle))
	p.SetStage(IamBetaRoleStageEnumToProto(resource.Stage))
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetDeleted(dcl.ValueOrEmptyBool(resource.Deleted))
	p.SetParent(dcl.ValueOrEmptyString(resource.Parent))
	sIncludedPermissions := make([]string, len(resource.IncludedPermissions))
	for i, r := range resource.IncludedPermissions {
		sIncludedPermissions[i] = r
	}
	p.SetIncludedPermissions(sIncludedPermissions)
	sIncludedRoles := make([]string, len(resource.IncludedRoles))
	for i, r := range resource.IncludedRoles {
		sIncludedRoles[i] = r
	}
	p.SetIncludedRoles(sIncludedRoles)

	return p
}

// applyRole handles the gRPC request by passing it to the underlying Role Apply() method.
func (s *RoleServer) applyRole(ctx context.Context, c *beta.Client, request *betapb.ApplyIamBetaRoleRequest) (*betapb.IamBetaRole, error) {
	p := ProtoToRole(request.GetResource())
	res, err := c.ApplyRole(ctx, p)
	if err != nil {
		return nil, err
	}
	r := RoleToProto(res)
	return r, nil
}

// applyIamBetaRole handles the gRPC request by passing it to the underlying Role Apply() method.
func (s *RoleServer) ApplyIamBetaRole(ctx context.Context, request *betapb.ApplyIamBetaRoleRequest) (*betapb.IamBetaRole, error) {
	cl, err := createConfigRole(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyRole(ctx, cl, request)
}

// DeleteRole handles the gRPC request by passing it to the underlying Role Delete() method.
func (s *RoleServer) DeleteIamBetaRole(ctx context.Context, request *betapb.DeleteIamBetaRoleRequest) (*emptypb.Empty, error) {

	cl, err := createConfigRole(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteRole(ctx, ProtoToRole(request.GetResource()))

}

// ListIamBetaRole handles the gRPC request by passing it to the underlying RoleList() method.
func (s *RoleServer) ListIamBetaRole(ctx context.Context, request *betapb.ListIamBetaRoleRequest) (*betapb.ListIamBetaRoleResponse, error) {
	cl, err := createConfigRole(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListRole(ctx, request.GetParent())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.IamBetaRole
	for _, r := range resources.Items {
		rp := RoleToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListIamBetaRoleResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigRole(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
