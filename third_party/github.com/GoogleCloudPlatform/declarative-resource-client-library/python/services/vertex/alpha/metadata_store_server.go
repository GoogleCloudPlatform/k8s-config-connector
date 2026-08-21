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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertex/alpha/vertex_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertex/alpha"
)

// MetadataStoreServer implements the gRPC interface for MetadataStore.
type MetadataStoreServer struct{}

// ProtoToMetadataStoreEncryptionSpec converts a MetadataStoreEncryptionSpec object from its proto representation.
func ProtoToVertexAlphaMetadataStoreEncryptionSpec(p *alphapb.VertexAlphaMetadataStoreEncryptionSpec) *alpha.MetadataStoreEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &alpha.MetadataStoreEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToMetadataStoreState converts a MetadataStoreState object from its proto representation.
func ProtoToVertexAlphaMetadataStoreState(p *alphapb.VertexAlphaMetadataStoreState) *alpha.MetadataStoreState {
	if p == nil {
		return nil
	}
	obj := &alpha.MetadataStoreState{
		DiskUtilizationBytes: dcl.Int64OrNil(p.GetDiskUtilizationBytes()),
	}
	return obj
}

// ProtoToMetadataStore converts a MetadataStore resource from its proto representation.
func ProtoToMetadataStore(p *alphapb.VertexAlphaMetadataStore) *alpha.MetadataStore {
	obj := &alpha.MetadataStore{
		Name:           dcl.StringOrNil(p.GetName()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		EncryptionSpec: ProtoToVertexAlphaMetadataStoreEncryptionSpec(p.GetEncryptionSpec()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		State:          ProtoToVertexAlphaMetadataStoreState(p.GetState()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// MetadataStoreEncryptionSpecToProto converts a MetadataStoreEncryptionSpec object to its proto representation.
func VertexAlphaMetadataStoreEncryptionSpecToProto(o *alpha.MetadataStoreEncryptionSpec) *alphapb.VertexAlphaMetadataStoreEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaMetadataStoreEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// MetadataStoreStateToProto converts a MetadataStoreState object to its proto representation.
func VertexAlphaMetadataStoreStateToProto(o *alpha.MetadataStoreState) *alphapb.VertexAlphaMetadataStoreState {
	if o == nil {
		return nil
	}
	p := &alphapb.VertexAlphaMetadataStoreState{}
	p.SetDiskUtilizationBytes(dcl.ValueOrEmptyInt64(o.DiskUtilizationBytes))
	return p
}

// MetadataStoreToProto converts a MetadataStore resource to its proto representation.
func MetadataStoreToProto(resource *alpha.MetadataStore) *alphapb.VertexAlphaMetadataStore {
	p := &alphapb.VertexAlphaMetadataStore{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEncryptionSpec(VertexAlphaMetadataStoreEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(VertexAlphaMetadataStoreStateToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Apply() method.
func (s *MetadataStoreServer) applyMetadataStore(ctx context.Context, c *alpha.Client, request *alphapb.ApplyVertexAlphaMetadataStoreRequest) (*alphapb.VertexAlphaMetadataStore, error) {
	p := ProtoToMetadataStore(request.GetResource())
	res, err := c.ApplyMetadataStore(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetadataStoreToProto(res)
	return r, nil
}

// applyVertexAlphaMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Apply() method.
func (s *MetadataStoreServer) ApplyVertexAlphaMetadataStore(ctx context.Context, request *alphapb.ApplyVertexAlphaMetadataStoreRequest) (*alphapb.VertexAlphaMetadataStore, error) {
	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetadataStore(ctx, cl, request)
}

// DeleteMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Delete() method.
func (s *MetadataStoreServer) DeleteVertexAlphaMetadataStore(ctx context.Context, request *alphapb.DeleteVertexAlphaMetadataStoreRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMetadataStore(ctx, ProtoToMetadataStore(request.GetResource()))

}

// ListVertexAlphaMetadataStore handles the gRPC request by passing it to the underlying MetadataStoreList() method.
func (s *MetadataStoreServer) ListVertexAlphaMetadataStore(ctx context.Context, request *alphapb.ListVertexAlphaMetadataStoreRequest) (*alphapb.ListVertexAlphaMetadataStoreResponse, error) {
	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetadataStore(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.VertexAlphaMetadataStore
	for _, r := range resources.Items {
		rp := MetadataStoreToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListVertexAlphaMetadataStoreResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMetadataStore(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
