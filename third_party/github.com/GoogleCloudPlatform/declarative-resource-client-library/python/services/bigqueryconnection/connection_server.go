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
	bigqueryconnectionpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigqueryconnection/bigqueryconnection_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigqueryconnection"
)

// Server implements the gRPC interface for Connection.
type ConnectionServer struct{}

// ProtoToConnectionCloudSqlTypeEnum converts a ConnectionCloudSqlTypeEnum enum from its proto representation.
func ProtoToBigqueryconnectionConnectionCloudSqlTypeEnum(e bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlTypeEnum) *bigqueryconnection.ConnectionCloudSqlTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlTypeEnum_name[int32(e)]; ok {
		e := bigqueryconnection.ConnectionCloudSqlTypeEnum(n[len("BigqueryconnectionConnectionCloudSqlTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToConnectionCloudSql converts a ConnectionCloudSql resource from its proto representation.
func ProtoToBigqueryconnectionConnectionCloudSql(p *bigqueryconnectionpb.BigqueryconnectionConnectionCloudSql) *bigqueryconnection.ConnectionCloudSql {
	if p == nil {
		return nil
	}
	obj := &bigqueryconnection.ConnectionCloudSql{
		InstanceId: dcl.StringOrNil(p.InstanceId),
		Database:   dcl.StringOrNil(p.Database),
		Type:       ProtoToBigqueryconnectionConnectionCloudSqlTypeEnum(p.GetType()),
		Credential: ProtoToBigqueryconnectionConnectionCloudSqlCredential(p.GetCredential()),
	}
	return obj
}

// ProtoToConnectionCloudSqlCredential converts a ConnectionCloudSqlCredential resource from its proto representation.
func ProtoToBigqueryconnectionConnectionCloudSqlCredential(p *bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlCredential) *bigqueryconnection.ConnectionCloudSqlCredential {
	if p == nil {
		return nil
	}
	obj := &bigqueryconnection.ConnectionCloudSqlCredential{
		Username: dcl.StringOrNil(p.Username),
		Password: dcl.StringOrNil(p.Password),
	}
	return obj
}

// ProtoToConnection converts a Connection resource from its proto representation.
func ProtoToConnection(p *bigqueryconnectionpb.BigqueryconnectionConnection) *bigqueryconnection.Connection {
	obj := &bigqueryconnection.Connection{
		Name:             dcl.StringOrNil(p.Name),
		FriendlyName:     dcl.StringOrNil(p.FriendlyName),
		Description:      dcl.StringOrNil(p.Description),
		CloudSql:         ProtoToBigqueryconnectionConnectionCloudSql(p.GetCloudSql()),
		CreationTime:     dcl.Int64OrNil(p.CreationTime),
		LastModifiedTime: dcl.Int64OrNil(p.LastModifiedTime),
		HasCredential:    dcl.Bool(p.HasCredential),
		Project:          dcl.StringOrNil(p.Project),
		Location:         dcl.StringOrNil(p.Location),
	}
	return obj
}

// ConnectionCloudSqlTypeEnumToProto converts a ConnectionCloudSqlTypeEnum enum to its proto representation.
func BigqueryconnectionConnectionCloudSqlTypeEnumToProto(e *bigqueryconnection.ConnectionCloudSqlTypeEnum) bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlTypeEnum {
	if e == nil {
		return bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlTypeEnum(0)
	}
	if v, ok := bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlTypeEnum_value["ConnectionCloudSqlTypeEnum"+string(*e)]; ok {
		return bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlTypeEnum(v)
	}
	return bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlTypeEnum(0)
}

// ConnectionCloudSqlToProto converts a ConnectionCloudSql resource to its proto representation.
func BigqueryconnectionConnectionCloudSqlToProto(o *bigqueryconnection.ConnectionCloudSql) *bigqueryconnectionpb.BigqueryconnectionConnectionCloudSql {
	if o == nil {
		return nil
	}
	p := &bigqueryconnectionpb.BigqueryconnectionConnectionCloudSql{
		InstanceId: dcl.ValueOrEmptyString(o.InstanceId),
		Database:   dcl.ValueOrEmptyString(o.Database),
		Type:       BigqueryconnectionConnectionCloudSqlTypeEnumToProto(o.Type),
		Credential: BigqueryconnectionConnectionCloudSqlCredentialToProto(o.Credential),
	}
	return p
}

// ConnectionCloudSqlCredentialToProto converts a ConnectionCloudSqlCredential resource to its proto representation.
func BigqueryconnectionConnectionCloudSqlCredentialToProto(o *bigqueryconnection.ConnectionCloudSqlCredential) *bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlCredential {
	if o == nil {
		return nil
	}
	p := &bigqueryconnectionpb.BigqueryconnectionConnectionCloudSqlCredential{
		Username: dcl.ValueOrEmptyString(o.Username),
		Password: dcl.ValueOrEmptyString(o.Password),
	}
	return p
}

// ConnectionToProto converts a Connection resource to its proto representation.
func ConnectionToProto(resource *bigqueryconnection.Connection) *bigqueryconnectionpb.BigqueryconnectionConnection {
	p := &bigqueryconnectionpb.BigqueryconnectionConnection{
		Name:             dcl.ValueOrEmptyString(resource.Name),
		FriendlyName:     dcl.ValueOrEmptyString(resource.FriendlyName),
		Description:      dcl.ValueOrEmptyString(resource.Description),
		CloudSql:         BigqueryconnectionConnectionCloudSqlToProto(resource.CloudSql),
		CreationTime:     dcl.ValueOrEmptyInt64(resource.CreationTime),
		LastModifiedTime: dcl.ValueOrEmptyInt64(resource.LastModifiedTime),
		HasCredential:    dcl.ValueOrEmptyBool(resource.HasCredential),
		Project:          dcl.ValueOrEmptyString(resource.Project),
		Location:         dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyConnection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) applyConnection(ctx context.Context, c *bigqueryconnection.Client, request *bigqueryconnectionpb.ApplyBigqueryconnectionConnectionRequest) (*bigqueryconnectionpb.BigqueryconnectionConnection, error) {
	p := ProtoToConnection(request.GetResource())
	res, err := c.ApplyConnection(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ConnectionToProto(res)
	return r, nil
}

// ApplyConnection handles the gRPC request by passing it to the underlying Connection Apply() method.
func (s *ConnectionServer) ApplyBigqueryconnectionConnection(ctx context.Context, request *bigqueryconnectionpb.ApplyBigqueryconnectionConnectionRequest) (*bigqueryconnectionpb.BigqueryconnectionConnection, error) {
	cl, err := createConfigConnection(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyConnection(ctx, cl, request)
}

// DeleteConnection handles the gRPC request by passing it to the underlying Connection Delete() method.
func (s *ConnectionServer) DeleteBigqueryconnectionConnection(ctx context.Context, request *bigqueryconnectionpb.DeleteBigqueryconnectionConnectionRequest) (*emptypb.Empty, error) {

	cl, err := createConfigConnection(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteConnection(ctx, ProtoToConnection(request.GetResource()))

}

// ListBigqueryconnectionConnection handles the gRPC request by passing it to the underlying ConnectionList() method.
func (s *ConnectionServer) ListBigqueryconnectionConnection(ctx context.Context, request *bigqueryconnectionpb.ListBigqueryconnectionConnectionRequest) (*bigqueryconnectionpb.ListBigqueryconnectionConnectionResponse, error) {
	cl, err := createConfigConnection(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListConnection(ctx, request.Project, request.Location)
	if err != nil {
		return nil, err
	}
	var protos []*bigqueryconnectionpb.BigqueryconnectionConnection
	for _, r := range resources.Items {
		rp := ConnectionToProto(r)
		protos = append(protos, rp)
	}
	return &bigqueryconnectionpb.ListBigqueryconnectionConnectionResponse{Items: protos}, nil
}

func createConfigConnection(ctx context.Context, service_account_file string) (*bigqueryconnection.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigqueryconnection.NewClient(conf), nil
}
