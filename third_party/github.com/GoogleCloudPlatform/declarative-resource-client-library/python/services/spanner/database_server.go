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
	spannerpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/spanner/spanner_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/spanner"
)

// Server implements the gRPC interface for Database.
type DatabaseServer struct{}

// ProtoToDatabaseStateEnum converts a DatabaseStateEnum enum from its proto representation.
func ProtoToSpannerDatabaseStateEnum(e spannerpb.SpannerDatabaseStateEnum) *spanner.DatabaseStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := spannerpb.SpannerDatabaseStateEnum_name[int32(e)]; ok {
		e := spanner.DatabaseStateEnum(n[len("SpannerDatabaseStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToDatabase converts a Database resource from its proto representation.
func ProtoToDatabase(p *spannerpb.SpannerDatabase) *spanner.Database {
	obj := &spanner.Database{
		Name:     dcl.StringOrNil(p.Name),
		Instance: dcl.StringOrNil(p.Instance),
		State:    ProtoToSpannerDatabaseStateEnum(p.GetState()),
		Project:  dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetDdl() {
		obj.Ddl = append(obj.Ddl, r)
	}
	return obj
}

// DatabaseStateEnumToProto converts a DatabaseStateEnum enum to its proto representation.
func SpannerDatabaseStateEnumToProto(e *spanner.DatabaseStateEnum) spannerpb.SpannerDatabaseStateEnum {
	if e == nil {
		return spannerpb.SpannerDatabaseStateEnum(0)
	}
	if v, ok := spannerpb.SpannerDatabaseStateEnum_value["DatabaseStateEnum"+string(*e)]; ok {
		return spannerpb.SpannerDatabaseStateEnum(v)
	}
	return spannerpb.SpannerDatabaseStateEnum(0)
}

// DatabaseToProto converts a Database resource to its proto representation.
func DatabaseToProto(resource *spanner.Database) *spannerpb.SpannerDatabase {
	p := &spannerpb.SpannerDatabase{
		Name:     dcl.ValueOrEmptyString(resource.Name),
		Instance: dcl.ValueOrEmptyString(resource.Instance),
		State:    SpannerDatabaseStateEnumToProto(resource.State),
		Project:  dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.Ddl {
		p.Ddl = append(p.Ddl, r)
	}

	return p
}

// ApplyDatabase handles the gRPC request by passing it to the underlying Database Apply() method.
func (s *DatabaseServer) applyDatabase(ctx context.Context, c *spanner.Client, request *spannerpb.ApplySpannerDatabaseRequest) (*spannerpb.SpannerDatabase, error) {
	p := ProtoToDatabase(request.GetResource())
	res, err := c.ApplyDatabase(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DatabaseToProto(res)
	return r, nil
}

// ApplyDatabase handles the gRPC request by passing it to the underlying Database Apply() method.
func (s *DatabaseServer) ApplySpannerDatabase(ctx context.Context, request *spannerpb.ApplySpannerDatabaseRequest) (*spannerpb.SpannerDatabase, error) {
	cl, err := createConfigDatabase(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyDatabase(ctx, cl, request)
}

// DeleteDatabase handles the gRPC request by passing it to the underlying Database Delete() method.
func (s *DatabaseServer) DeleteSpannerDatabase(ctx context.Context, request *spannerpb.DeleteSpannerDatabaseRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDatabase(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDatabase(ctx, ProtoToDatabase(request.GetResource()))

}

// ListSpannerDatabase handles the gRPC request by passing it to the underlying DatabaseList() method.
func (s *DatabaseServer) ListSpannerDatabase(ctx context.Context, request *spannerpb.ListSpannerDatabaseRequest) (*spannerpb.ListSpannerDatabaseResponse, error) {
	cl, err := createConfigDatabase(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDatabase(ctx, request.Project, request.Instance)
	if err != nil {
		return nil, err
	}
	var protos []*spannerpb.SpannerDatabase
	for _, r := range resources.Items {
		rp := DatabaseToProto(r)
		protos = append(protos, rp)
	}
	return &spannerpb.ListSpannerDatabaseResponse{Items: protos}, nil
}

func createConfigDatabase(ctx context.Context, service_account_file string) (*spanner.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return spanner.NewClient(conf), nil
}
