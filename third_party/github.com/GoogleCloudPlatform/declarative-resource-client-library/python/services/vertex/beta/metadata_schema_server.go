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
	"errors"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertex/beta/vertex_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertex/beta"
)

// MetadataSchemaServer implements the gRPC interface for MetadataSchema.
type MetadataSchemaServer struct{}

// ProtoToMetadataSchemaSchemaTypeEnum converts a MetadataSchemaSchemaTypeEnum enum from its proto representation.
func ProtoToVertexBetaMetadataSchemaSchemaTypeEnum(e betapb.VertexBetaMetadataSchemaSchemaTypeEnum) *beta.MetadataSchemaSchemaTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.VertexBetaMetadataSchemaSchemaTypeEnum_name[int32(e)]; ok {
		e := beta.MetadataSchemaSchemaTypeEnum(n[len("VertexBetaMetadataSchemaSchemaTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetadataSchema converts a MetadataSchema resource from its proto representation.
func ProtoToMetadataSchema(p *betapb.VertexBetaMetadataSchema) *beta.MetadataSchema {
	obj := &beta.MetadataSchema{
		Name:          dcl.StringOrNil(p.GetName()),
		SchemaVersion: dcl.StringOrNil(p.GetSchemaVersion()),
		Schema:        dcl.StringOrNil(p.GetSchema()),
		SchemaType:    ProtoToVertexBetaMetadataSchemaSchemaTypeEnum(p.GetSchemaType()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		MetadataStore: dcl.StringOrNil(p.GetMetadataStore()),
	}
	return obj
}

// MetadataSchemaSchemaTypeEnumToProto converts a MetadataSchemaSchemaTypeEnum enum to its proto representation.
func VertexBetaMetadataSchemaSchemaTypeEnumToProto(e *beta.MetadataSchemaSchemaTypeEnum) betapb.VertexBetaMetadataSchemaSchemaTypeEnum {
	if e == nil {
		return betapb.VertexBetaMetadataSchemaSchemaTypeEnum(0)
	}
	if v, ok := betapb.VertexBetaMetadataSchemaSchemaTypeEnum_value["MetadataSchemaSchemaTypeEnum"+string(*e)]; ok {
		return betapb.VertexBetaMetadataSchemaSchemaTypeEnum(v)
	}
	return betapb.VertexBetaMetadataSchemaSchemaTypeEnum(0)
}

// MetadataSchemaToProto converts a MetadataSchema resource to its proto representation.
func MetadataSchemaToProto(resource *beta.MetadataSchema) *betapb.VertexBetaMetadataSchema {
	p := &betapb.VertexBetaMetadataSchema{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSchemaVersion(dcl.ValueOrEmptyString(resource.SchemaVersion))
	p.SetSchema(dcl.ValueOrEmptyString(resource.Schema))
	p.SetSchemaType(VertexBetaMetadataSchemaSchemaTypeEnumToProto(resource.SchemaType))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetMetadataStore(dcl.ValueOrEmptyString(resource.MetadataStore))

	return p
}

// applyMetadataSchema handles the gRPC request by passing it to the underlying MetadataSchema Apply() method.
func (s *MetadataSchemaServer) applyMetadataSchema(ctx context.Context, c *beta.Client, request *betapb.ApplyVertexBetaMetadataSchemaRequest) (*betapb.VertexBetaMetadataSchema, error) {
	p := ProtoToMetadataSchema(request.GetResource())
	res, err := c.ApplyMetadataSchema(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetadataSchemaToProto(res)
	return r, nil
}

// applyVertexBetaMetadataSchema handles the gRPC request by passing it to the underlying MetadataSchema Apply() method.
func (s *MetadataSchemaServer) ApplyVertexBetaMetadataSchema(ctx context.Context, request *betapb.ApplyVertexBetaMetadataSchemaRequest) (*betapb.VertexBetaMetadataSchema, error) {
	cl, err := createConfigMetadataSchema(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetadataSchema(ctx, cl, request)
}

// DeleteMetadataSchema handles the gRPC request by passing it to the underlying MetadataSchema Delete() method.
func (s *MetadataSchemaServer) DeleteVertexBetaMetadataSchema(ctx context.Context, request *betapb.DeleteVertexBetaMetadataSchemaRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for MetadataSchema")

}

// ListVertexBetaMetadataSchema handles the gRPC request by passing it to the underlying MetadataSchemaList() method.
func (s *MetadataSchemaServer) ListVertexBetaMetadataSchema(ctx context.Context, request *betapb.ListVertexBetaMetadataSchemaRequest) (*betapb.ListVertexBetaMetadataSchemaResponse, error) {
	cl, err := createConfigMetadataSchema(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetadataSchema(ctx, request.GetProject(), request.GetLocation(), request.GetMetadataStore())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.VertexBetaMetadataSchema
	for _, r := range resources.Items {
		rp := MetadataSchemaToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListVertexBetaMetadataSchemaResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMetadataSchema(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
