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
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	vertexpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertex/vertex_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertex"
)

// MetadataStoreServer implements the gRPC interface for MetadataStore.
type MetadataStoreServer struct{}

// ProtoToMetadataStoreEncryptionSpec converts a MetadataStoreEncryptionSpec object from its proto representation.
func ProtoToVertexMetadataStoreEncryptionSpec(p *vertexpb.VertexMetadataStoreEncryptionSpec) *vertex.MetadataStoreEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &vertex.MetadataStoreEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToMetadataStoreState converts a MetadataStoreState object from its proto representation.
func ProtoToVertexMetadataStoreState(p *vertexpb.VertexMetadataStoreState) *vertex.MetadataStoreState {
	if p == nil {
		return nil
	}
	obj := &vertex.MetadataStoreState{
		DiskUtilizationBytes: dcl.Int64OrNil(p.GetDiskUtilizationBytes()),
	}
	return obj
}

// ProtoToMetadataStore converts a MetadataStore resource from its proto representation.
func ProtoToMetadataStore(p *vertexpb.VertexMetadataStore) *vertex.MetadataStore {
	obj := &vertex.MetadataStore{
		Name:           dcl.StringOrNil(p.GetName()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		EncryptionSpec: ProtoToVertexMetadataStoreEncryptionSpec(p.GetEncryptionSpec()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		State:          ProtoToVertexMetadataStoreState(p.GetState()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// MetadataStoreEncryptionSpecToProto converts a MetadataStoreEncryptionSpec object to its proto representation.
func VertexMetadataStoreEncryptionSpecToProto(o *vertex.MetadataStoreEncryptionSpec) *vertexpb.VertexMetadataStoreEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &vertexpb.VertexMetadataStoreEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// MetadataStoreStateToProto converts a MetadataStoreState object to its proto representation.
func VertexMetadataStoreStateToProto(o *vertex.MetadataStoreState) *vertexpb.VertexMetadataStoreState {
	if o == nil {
		return nil
	}
	p := &vertexpb.VertexMetadataStoreState{}
	p.SetDiskUtilizationBytes(dcl.ValueOrEmptyInt64(o.DiskUtilizationBytes))
	return p
}

// MetadataStoreToProto converts a MetadataStore resource to its proto representation.
func MetadataStoreToProto(resource *vertex.MetadataStore) *vertexpb.VertexMetadataStore {
	p := &vertexpb.VertexMetadataStore{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEncryptionSpec(VertexMetadataStoreEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(VertexMetadataStoreStateToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Apply() method.
func (s *MetadataStoreServer) applyMetadataStore(ctx context.Context, c *vertex.Client, request *vertexpb.ApplyVertexMetadataStoreRequest) (*vertexpb.VertexMetadataStore, error) {
	p := ProtoToMetadataStore(request.GetResource())
	res, err := c.ApplyMetadataStore(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetadataStoreToProto(res)
	return r, nil
}

// applyVertexMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Apply() method.
func (s *MetadataStoreServer) ApplyVertexMetadataStore(ctx context.Context, request *vertexpb.ApplyVertexMetadataStoreRequest) (*vertexpb.VertexMetadataStore, error) {
	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetadataStore(ctx, cl, request)
}

// DeleteMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Delete() method.
func (s *MetadataStoreServer) DeleteVertexMetadataStore(ctx context.Context, request *vertexpb.DeleteVertexMetadataStoreRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMetadataStore(ctx, ProtoToMetadataStore(request.GetResource()))

}

// ListVertexMetadataStore handles the gRPC request by passing it to the underlying MetadataStoreList() method.
func (s *MetadataStoreServer) ListVertexMetadataStore(ctx context.Context, request *vertexpb.ListVertexMetadataStoreRequest) (*vertexpb.ListVertexMetadataStoreResponse, error) {
	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetadataStore(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*vertexpb.VertexMetadataStore
	for _, r := range resources.Items {
		rp := MetadataStoreToProto(r)
		protos = append(protos, rp)
	}
	p := &vertexpb.ListVertexMetadataStoreResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMetadataStore(ctx context.Context, service_account_file string) (*vertex.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return vertex.NewClient(conf), nil
}
