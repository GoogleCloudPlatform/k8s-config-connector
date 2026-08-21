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
	vertexaipb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/vertexai_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai"
)

// MetadataStoreServer implements the gRPC interface for MetadataStore.
type MetadataStoreServer struct{}

// ProtoToMetadataStoreEncryptionSpec converts a MetadataStoreEncryptionSpec object from its proto representation.
func ProtoToVertexaiMetadataStoreEncryptionSpec(p *vertexaipb.VertexaiMetadataStoreEncryptionSpec) *vertexai.MetadataStoreEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &vertexai.MetadataStoreEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToMetadataStoreState converts a MetadataStoreState object from its proto representation.
func ProtoToVertexaiMetadataStoreState(p *vertexaipb.VertexaiMetadataStoreState) *vertexai.MetadataStoreState {
	if p == nil {
		return nil
	}
	obj := &vertexai.MetadataStoreState{
		DiskUtilizationBytes: dcl.Int64OrNil(p.GetDiskUtilizationBytes()),
	}
	return obj
}

// ProtoToMetadataStore converts a MetadataStore resource from its proto representation.
func ProtoToMetadataStore(p *vertexaipb.VertexaiMetadataStore) *vertexai.MetadataStore {
	obj := &vertexai.MetadataStore{
		Name:           dcl.StringOrNil(p.GetName()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		EncryptionSpec: ProtoToVertexaiMetadataStoreEncryptionSpec(p.GetEncryptionSpec()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		State:          ProtoToVertexaiMetadataStoreState(p.GetState()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// MetadataStoreEncryptionSpecToProto converts a MetadataStoreEncryptionSpec object to its proto representation.
func VertexaiMetadataStoreEncryptionSpecToProto(o *vertexai.MetadataStoreEncryptionSpec) *vertexaipb.VertexaiMetadataStoreEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiMetadataStoreEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// MetadataStoreStateToProto converts a MetadataStoreState object to its proto representation.
func VertexaiMetadataStoreStateToProto(o *vertexai.MetadataStoreState) *vertexaipb.VertexaiMetadataStoreState {
	if o == nil {
		return nil
	}
	p := &vertexaipb.VertexaiMetadataStoreState{}
	p.SetDiskUtilizationBytes(dcl.ValueOrEmptyInt64(o.DiskUtilizationBytes))
	return p
}

// MetadataStoreToProto converts a MetadataStore resource to its proto representation.
func MetadataStoreToProto(resource *vertexai.MetadataStore) *vertexaipb.VertexaiMetadataStore {
	p := &vertexaipb.VertexaiMetadataStore{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEncryptionSpec(VertexaiMetadataStoreEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(VertexaiMetadataStoreStateToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Apply() method.
func (s *MetadataStoreServer) applyMetadataStore(ctx context.Context, c *vertexai.Client, request *vertexaipb.ApplyVertexaiMetadataStoreRequest) (*vertexaipb.VertexaiMetadataStore, error) {
	p := ProtoToMetadataStore(request.GetResource())
	res, err := c.ApplyMetadataStore(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetadataStoreToProto(res)
	return r, nil
}

// applyVertexaiMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Apply() method.
func (s *MetadataStoreServer) ApplyVertexaiMetadataStore(ctx context.Context, request *vertexaipb.ApplyVertexaiMetadataStoreRequest) (*vertexaipb.VertexaiMetadataStore, error) {
	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetadataStore(ctx, cl, request)
}

// DeleteMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Delete() method.
func (s *MetadataStoreServer) DeleteVertexaiMetadataStore(ctx context.Context, request *vertexaipb.DeleteVertexaiMetadataStoreRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMetadataStore(ctx, ProtoToMetadataStore(request.GetResource()))

}

// ListVertexaiMetadataStore handles the gRPC request by passing it to the underlying MetadataStoreList() method.
func (s *MetadataStoreServer) ListVertexaiMetadataStore(ctx context.Context, request *vertexaipb.ListVertexaiMetadataStoreRequest) (*vertexaipb.ListVertexaiMetadataStoreResponse, error) {
	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetadataStore(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*vertexaipb.VertexaiMetadataStore
	for _, r := range resources.Items {
		rp := MetadataStoreToProto(r)
		protos = append(protos, rp)
	}
	p := &vertexaipb.ListVertexaiMetadataStoreResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMetadataStore(ctx context.Context, service_account_file string) (*vertexai.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return vertexai.NewClient(conf), nil
}
