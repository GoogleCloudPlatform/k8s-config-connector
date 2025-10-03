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

// Server implements the gRPC interface for Database.
type DatabaseServer struct{}

// ProtoToDatabase converts a Database resource from its proto representation.
func ProtoToDatabase(p *sqlpb.SqlDatabase) *sql.Database {
	obj := &sql.Database{
		Charset:   dcl.StringOrNil(p.Charset),
		Collation: dcl.StringOrNil(p.Collation),
		Instance:  dcl.StringOrNil(p.Instance),
		Name:      dcl.StringOrNil(p.Name),
		Project:   dcl.StringOrNil(p.Project),
		SelfLink:  dcl.StringOrNil(p.SelfLink),
	}
	return obj
}

// DatabaseToProto converts a Database resource to its proto representation.
func DatabaseToProto(resource *sql.Database) *sqlpb.SqlDatabase {
	p := &sqlpb.SqlDatabase{
		Charset:   dcl.ValueOrEmptyString(resource.Charset),
		Collation: dcl.ValueOrEmptyString(resource.Collation),
		Instance:  dcl.ValueOrEmptyString(resource.Instance),
		Name:      dcl.ValueOrEmptyString(resource.Name),
		Project:   dcl.ValueOrEmptyString(resource.Project),
		SelfLink:  dcl.ValueOrEmptyString(resource.SelfLink),
	}

	return p
}

// ApplyDatabase handles the gRPC request by passing it to the underlying Database Apply() method.
func (s *DatabaseServer) applyDatabase(ctx context.Context, c *sql.Client, request *sqlpb.ApplySqlDatabaseRequest) (*sqlpb.SqlDatabase, error) {
	p := ProtoToDatabase(request.GetResource())
	res, err := c.ApplyDatabase(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DatabaseToProto(res)
	return r, nil
}

// ApplyDatabase handles the gRPC request by passing it to the underlying Database Apply() method.
func (s *DatabaseServer) ApplySqlDatabase(ctx context.Context, request *sqlpb.ApplySqlDatabaseRequest) (*sqlpb.SqlDatabase, error) {
	cl, err := createConfigDatabase(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyDatabase(ctx, cl, request)
}

// DeleteDatabase handles the gRPC request by passing it to the underlying Database Delete() method.
func (s *DatabaseServer) DeleteSqlDatabase(ctx context.Context, request *sqlpb.DeleteSqlDatabaseRequest) (*emptypb.Empty, error) {
	cl, err := createConfigDatabase(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDatabase(ctx, ProtoToDatabase(request.GetResource()))
}

// ListDatabase handles the gRPC request by passing it to the underlying DatabaseList() method.
func (s *DatabaseServer) ListSqlDatabase(ctx context.Context, request *sqlpb.ListSqlDatabaseRequest) (*sqlpb.ListSqlDatabaseResponse, error) {
	cl, err := createConfigDatabase(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDatabase(ctx, request.Project, request.Instance)
	if err != nil {
		return nil, err
	}
	var protos []*sqlpb.SqlDatabase
	for _, r := range resources.Items {
		rp := DatabaseToProto(r)
		protos = append(protos, rp)
	}
	return &sqlpb.ListSqlDatabaseResponse{Items: protos}, nil
}

func createConfigDatabase(ctx context.Context, service_account_file string) (*sql.Client, error) {

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
