// Copyright 2020 Google LLC. All Rights Reserved.
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
	sqlpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/sql/sql_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/sql"
)

// Server implements the gRPC interface for User.
type UserServer struct{}

// ProtoToUserTypeEnum converts a UserTypeEnum enum from its proto representation.
func ProtoToSqlUserTypeEnum(e sqlpb.SqlUserTypeEnum) *sql.UserTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := sqlpb.SqlUserTypeEnum_name[int32(e)]; ok {
		e := sql.UserTypeEnum(n[len("UserTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToUserSqlserverUserDetails converts a UserSqlserverUserDetails resource from its proto representation.
func ProtoToSqlUserSqlserverUserDetails(p *sqlpb.SqlUserSqlserverUserDetails) *sql.UserSqlserverUserDetails {
	if p == nil {
		return nil
	}
	obj := &sql.UserSqlserverUserDetails{
		Disabled: dcl.Bool(p.Disabled),
	}
	for _, r := range p.GetServerRoles() {
		obj.ServerRoles = append(obj.ServerRoles, r)
	}
	return obj
}

// ProtoToUser converts a User resource from its proto representation.
func ProtoToUser(p *sqlpb.SqlUser) *sql.User {
	obj := &sql.User{
		Name:                 dcl.StringOrNil(p.Name),
		Password:             dcl.StringOrNil(p.Password),
		Project:              dcl.StringOrNil(p.Project),
		Instance:             dcl.StringOrNil(p.Instance),
		SqlserverUserDetails: ProtoToSqlUserSqlserverUserDetails(p.GetSqlserverUserDetails()),
		Type:                 ProtoToSqlUserTypeEnum(p.GetType()),
		Etag:                 dcl.StringOrNil(p.Etag),
		Host:                 dcl.StringOrNil(p.Host),
	}
	return obj
}

// UserTypeEnumToProto converts a UserTypeEnum enum to its proto representation.
func SqlUserTypeEnumToProto(e *sql.UserTypeEnum) sqlpb.SqlUserTypeEnum {
	if e == nil {
		return sqlpb.SqlUserTypeEnum(0)
	}
	if v, ok := sqlpb.SqlUserTypeEnum_value["UserTypeEnum"+string(*e)]; ok {
		return sqlpb.SqlUserTypeEnum(v)
	}
	return sqlpb.SqlUserTypeEnum(0)
}

// UserSqlserverUserDetailsToProto converts a UserSqlserverUserDetails resource to its proto representation.
func SqlUserSqlserverUserDetailsToProto(o *sql.UserSqlserverUserDetails) *sqlpb.SqlUserSqlserverUserDetails {
	if o == nil {
		return nil
	}
	p := &sqlpb.SqlUserSqlserverUserDetails{
		Disabled: dcl.ValueOrEmptyBool(o.Disabled),
	}
	for _, r := range o.ServerRoles {
		p.ServerRoles = append(p.ServerRoles, r)
	}
	return p
}

// UserToProto converts a User resource to its proto representation.
func UserToProto(resource *sql.User) *sqlpb.SqlUser {
	p := &sqlpb.SqlUser{
		Name:                 dcl.ValueOrEmptyString(resource.Name),
		Password:             dcl.ValueOrEmptyString(resource.Password),
		Project:              dcl.ValueOrEmptyString(resource.Project),
		Instance:             dcl.ValueOrEmptyString(resource.Instance),
		SqlserverUserDetails: SqlUserSqlserverUserDetailsToProto(resource.SqlserverUserDetails),
		Type:                 SqlUserTypeEnumToProto(resource.Type),
		Etag:                 dcl.ValueOrEmptyString(resource.Etag),
		Host:                 dcl.ValueOrEmptyString(resource.Host),
	}

	return p
}

// ApplyUser handles the gRPC request by passing it to the underlying User Apply() method.
func (s *UserServer) applyUser(ctx context.Context, c *sql.Client, request *sqlpb.ApplySqlUserRequest) (*sqlpb.SqlUser, error) {
	p := ProtoToUser(request.GetResource())
	res, err := c.ApplyUser(ctx, p)
	if err != nil {
		return nil, err
	}
	r := UserToProto(res)
	return r, nil
}

// ApplyUser handles the gRPC request by passing it to the underlying User Apply() method.
func (s *UserServer) ApplySqlUser(ctx context.Context, request *sqlpb.ApplySqlUserRequest) (*sqlpb.SqlUser, error) {
	cl, err := createConfigUser(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyUser(ctx, cl, request)
}

// DeleteUser handles the gRPC request by passing it to the underlying User Delete() method.
func (s *UserServer) DeleteSqlUser(ctx context.Context, request *sqlpb.DeleteSqlUserRequest) (*emptypb.Empty, error) {
	cl, err := createConfigUser(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteUser(ctx, ProtoToUser(request.GetResource()))
}

// ListUser handles the gRPC request by passing it to the underlying UserList() method.
func (s *UserServer) ListSqlUser(ctx context.Context, request *sqlpb.ListSqlUserRequest) (*sqlpb.ListSqlUserResponse, error) {
	cl, err := createConfigUser(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListUser(ctx, request.Project, request.Instance)
	if err != nil {
		return nil, err
	}
	var protos []*sqlpb.SqlUser
	for _, r := range resources.Items {
		rp := UserToProto(r)
		protos = append(protos, rp)
	}
	return &sqlpb.ListSqlUserResponse{Items: protos}, nil
}

func createConfigUser(ctx context.Context, service_account_file string) (*sql.Client, error) {

	client, err := dcl.FromCredentialsFile(ctx, service_account_file)
	if err != nil {
		return nil, err
	}

	conf := dcl.NewConfig(client, dcl.WithUserAgent("dcl-test"))
	if err != nil {
		return nil, err
	}
	return sql.NewClient(conf), nil
}
