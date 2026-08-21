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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/sql/beta/sql_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sql/beta"
)

// Server implements the gRPC interface for User.
type UserServer struct{}

// ProtoToUserTypeEnum converts a UserTypeEnum enum from its proto representation.
func ProtoToSqlBetaUserTypeEnum(e betapb.SqlBetaUserTypeEnum) *beta.UserTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.SqlBetaUserTypeEnum_name[int32(e)]; ok {
		e := beta.UserTypeEnum(n[len("SqlBetaUserTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUserSqlserverUserDetails converts a UserSqlserverUserDetails resource from its proto representation.
func ProtoToSqlBetaUserSqlserverUserDetails(p *betapb.SqlBetaUserSqlserverUserDetails) *beta.UserSqlserverUserDetails {
	if p == nil {
		return nil
	}
	obj := &beta.UserSqlserverUserDetails{
		Disabled: dcl.Bool(p.Disabled),
	}
	for _, r := range p.GetServerRoles() {
		obj.ServerRoles = append(obj.ServerRoles, r)
	}
	return obj
}

// ProtoToUser converts a User resource from its proto representation.
func ProtoToUser(p *betapb.SqlBetaUser) *beta.User {
	obj := &beta.User{
		Name:                 dcl.StringOrNil(p.Name),
		Password:             dcl.StringOrNil(p.Password),
		Project:              dcl.StringOrNil(p.Project),
		Instance:             dcl.StringOrNil(p.Instance),
		SqlserverUserDetails: ProtoToSqlBetaUserSqlserverUserDetails(p.GetSqlserverUserDetails()),
		Type:                 ProtoToSqlBetaUserTypeEnum(p.GetType()),
		Etag:                 dcl.StringOrNil(p.Etag),
		Host:                 dcl.StringOrNil(p.Host),
	}
	return obj
}

// UserTypeEnumToProto converts a UserTypeEnum enum to its proto representation.
func SqlBetaUserTypeEnumToProto(e *beta.UserTypeEnum) betapb.SqlBetaUserTypeEnum {
	if e == nil {
		return betapb.SqlBetaUserTypeEnum(0)
	}
	if v, ok := betapb.SqlBetaUserTypeEnum_value["UserTypeEnum"+string(*e)]; ok {
		return betapb.SqlBetaUserTypeEnum(v)
	}
	return betapb.SqlBetaUserTypeEnum(0)
}

// UserSqlserverUserDetailsToProto converts a UserSqlserverUserDetails resource to its proto representation.
func SqlBetaUserSqlserverUserDetailsToProto(o *beta.UserSqlserverUserDetails) *betapb.SqlBetaUserSqlserverUserDetails {
	if o == nil {
		return nil
	}
	p := &betapb.SqlBetaUserSqlserverUserDetails{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	for _, r := range o.ServerRoles {
		p.ServerRoles = append(p.ServerRoles, r)
	}
	return p
}

// UserToProto converts a User resource to its proto representation.
func UserToProto(resource *beta.User) *betapb.SqlBetaUser {
	p := &betapb.SqlBetaUser{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Password:             dcl.ValueOrEmptyString(resource.Password),
		Project:              dcl.ValueOrEmptyString(resource.Project),
		Instance:             dcl.ValueOrEmptyString(resource.Instance),
		SqlserverUserDetails: SqlBetaUserSqlserverUserDetailsToProto(resource.SqlserverUserDetails),
		Type:                 SqlBetaUserTypeEnumToProto(resource.Type),
		Etag:                 dcl.ValueOrEmptyString(resource.Etag),
		Host:                 dcl.ValueOrEmptyString(resource.Host),
	}

	return p
}

// ApplyUser handles the gRPC request by passing it to the underlying User Apply() method.
func (s *UserServer) applyUser(ctx context.Context, c *beta.Client, request *betapb.ApplySqlBetaUserRequest) (*betapb.SqlBetaUser, error) {
	p := ProtoToUser(request.GetResource())
	res, err := c.ApplyUser(ctx, p)
	if err != nil {
		return nil, err
	}
	r := UserToProto(res)
	return r, nil
}

// ApplyUser handles the gRPC request by passing it to the underlying User Apply() method.
func (s *UserServer) ApplySqlBetaUser(ctx context.Context, request *betapb.ApplySqlBetaUserRequest) (*betapb.SqlBetaUser, error) {
	cl, err := createConfigUser(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyUser(ctx, cl, request)
}

// DeleteUser handles the gRPC request by passing it to the underlying User Delete() method.
func (s *UserServer) DeleteSqlBetaUser(ctx context.Context, request *betapb.DeleteSqlBetaUserRequest) (*emptypb.Empty, error) {

	cl, err := createConfigUser(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteUser(ctx, ProtoToUser(request.GetResource()))

}

// ListSqlBetaUser handles the gRPC request by passing it to the underlying UserList() method.
func (s *UserServer) ListSqlBetaUser(ctx context.Context, request *betapb.ListSqlBetaUserRequest) (*betapb.ListSqlBetaUserResponse, error) {
	cl, err := createConfigUser(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListUser(ctx, request.Project, request.Instance)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.SqlBetaUser
	for _, r := range resources.Items {
		rp := UserToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListSqlBetaUserResponse{Items: protos}, nil
}

func createConfigUser(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
