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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertex/alpha/vertex_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertex/alpha"
)

// MetadataSchemaServer implements the gRPC interface for MetadataSchema.
type MetadataSchemaServer struct{}

// ProtoToMetadataSchemaSchemaTypeEnum converts a MetadataSchemaSchemaTypeEnum enum from its proto representation.
func ProtoToVertexAlphaMetadataSchemaSchemaTypeEnum(e alphapb.VertexAlphaMetadataSchemaSchemaTypeEnum) *alpha.MetadataSchemaSchemaTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.VertexAlphaMetadataSchemaSchemaTypeEnum_name[int32(e)]; ok {
		e := alpha.MetadataSchemaSchemaTypeEnum(n[len("VertexAlphaMetadataSchemaSchemaTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToMetadataSchema converts a MetadataSchema resource from its proto representation.
func ProtoToMetadataSchema(p *alphapb.VertexAlphaMetadataSchema) *alpha.MetadataSchema {
	obj := &alpha.MetadataSchema{
		Name:          dcl.StringOrNil(p.GetName()),
		SchemaVersion: dcl.StringOrNil(p.GetSchemaVersion()),
		Schema:        dcl.StringOrNil(p.GetSchema()),
		SchemaType:    ProtoToVertexAlphaMetadataSchemaSchemaTypeEnum(p.GetSchemaType()),
		CreateTime:    dcl.StringOrNil(p.GetCreateTime()),
		Project:       dcl.StringOrNil(p.GetProject()),
		Location:      dcl.StringOrNil(p.GetLocation()),
		MetadataStore: dcl.StringOrNil(p.GetMetadataStore()),
	}
	return obj
}

// MetadataSchemaSchemaTypeEnumToProto converts a MetadataSchemaSchemaTypeEnum enum to its proto representation.
func VertexAlphaMetadataSchemaSchemaTypeEnumToProto(e *alpha.MetadataSchemaSchemaTypeEnum) alphapb.VertexAlphaMetadataSchemaSchemaTypeEnum {
	if e == nil {
		return alphapb.VertexAlphaMetadataSchemaSchemaTypeEnum(0)
	}
	if v, ok := alphapb.VertexAlphaMetadataSchemaSchemaTypeEnum_value["MetadataSchemaSchemaTypeEnum"+string(*e)]; ok {
		return alphapb.VertexAlphaMetadataSchemaSchemaTypeEnum(v)
	}
	return alphapb.VertexAlphaMetadataSchemaSchemaTypeEnum(0)
}

// MetadataSchemaToProto converts a MetadataSchema resource to its proto representation.
func MetadataSchemaToProto(resource *alpha.MetadataSchema) *alphapb.VertexAlphaMetadataSchema {
	p := &alphapb.VertexAlphaMetadataSchema{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetSchemaVersion(dcl.ValueOrEmptyString(resource.SchemaVersion))
	p.SetSchema(dcl.ValueOrEmptyString(resource.Schema))
	p.SetSchemaType(VertexAlphaMetadataSchemaSchemaTypeEnumToProto(resource.SchemaType))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetMetadataStore(dcl.ValueOrEmptyString(resource.MetadataStore))

	return p
}

// applyMetadataSchema handles the gRPC request by passing it to the underlying MetadataSchema Apply() method.
func (s *MetadataSchemaServer) applyMetadataSchema(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVertexAlphaMetadataSchemaRequest) (*alphapb.VertexAlphaMetadataSchema, error) {
	p := ProtoToMetadataSchema(request.GetResource())
	res, err := c.ApplyMetadataSchema(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetadataSchemaToProto(res)
	return r, nil
}

// applyVertexAlphaMetadataSchema handles the gRPC request by passing it to the underlying MetadataSchema Apply() method.
func (s *MetadataSchemaServer) ApplyVertexAlphaMetadataSchema(ctx context.Context, request *alphapb.ApplyVertexAlphaMetadataSchemaRequest) (*alphapb.VertexAlphaMetadataSchema, error) {
	cl, err := createConfigMetadataSchema(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetadataSchema(ctx, cl, request)
}

// DeleteMetadataSchema handles the gRPC request by passing it to the underlying MetadataSchema Delete() method.
func (s *MetadataSchemaServer) DeleteVertexAlphaMetadataSchema(ctx context.Context, request *alphapb.DeleteVertexAlphaMetadataSchemaRequest) (*emptypb.Empty, error) {

	return nil, errors.New("no delete endpoint for MetadataSchema")

}

// ListVertexAlphaMetadataSchema handles the gRPC request by passing it to the underlying MetadataSchemaList() method.
func (s *MetadataSchemaServer) ListVertexAlphaMetadataSchema(ctx context.Context, request *alphapb.ListVertexAlphaMetadataSchemaRequest) (*alphapb.ListVertexAlphaMetadataSchemaResponse, error) {
	cl, err := createConfigMetadataSchema(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetadataSchema(ctx, request.GetProject(), request.GetLocation(), request.GetMetadataStore())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VertexAlphaMetadataSchema
	for _, r := range resources.Items {
		rp := MetadataSchemaToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVertexAlphaMetadataSchemaResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMetadataSchema(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
