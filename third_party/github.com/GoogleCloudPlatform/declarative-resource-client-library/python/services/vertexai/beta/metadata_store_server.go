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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/vertexai/beta/vertexai_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai/beta"
)

// MetadataStoreServer implements the gRPC interface for MetadataStore.
type MetadataStoreServer struct{}

// ProtoToMetadataStoreEncryptionSpec converts a MetadataStoreEncryptionSpec object from its proto representation.
func ProtoToVertexaiBetaMetadataStoreEncryptionSpec(p *betapb.VertexaiBetaMetadataStoreEncryptionSpec) *beta.MetadataStoreEncryptionSpec {
	if p == nil {
		return nil
	}
	obj := &beta.MetadataStoreEncryptionSpec{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToMetadataStoreState converts a MetadataStoreState object from its proto representation.
func ProtoToVertexaiBetaMetadataStoreState(p *betapb.VertexaiBetaMetadataStoreState) *beta.MetadataStoreState {
	if p == nil {
		return nil
	}
	obj := &beta.MetadataStoreState{
		DiskUtilizationBytes: dcl.Int64OrNil(p.GetDiskUtilizationBytes()),
	}
	return obj
}

// ProtoToMetadataStore converts a MetadataStore resource from its proto representation.
func ProtoToMetadataStore(p *betapb.VertexaiBetaMetadataStore) *beta.MetadataStore {
	obj := &beta.MetadataStore{
		Name:           dcl.StringOrNil(p.GetName()),
		CreateTime:     dcl.StringOrNil(p.GetCreateTime()),
		UpdateTime:     dcl.StringOrNil(p.GetUpdateTime()),
		EncryptionSpec: ProtoToVertexaiBetaMetadataStoreEncryptionSpec(p.GetEncryptionSpec()),
		Description:    dcl.StringOrNil(p.GetDescription()),
		State:          ProtoToVertexaiBetaMetadataStoreState(p.GetState()),
		Project:        dcl.StringOrNil(p.GetProject()),
		Location:       dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// MetadataStoreEncryptionSpecToProto converts a MetadataStoreEncryptionSpec object to its proto representation.
func VertexaiBetaMetadataStoreEncryptionSpecToProto(o *beta.MetadataStoreEncryptionSpec) *betapb.VertexaiBetaMetadataStoreEncryptionSpec {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaMetadataStoreEncryptionSpec{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// MetadataStoreStateToProto converts a MetadataStoreState object to its proto representation.
func VertexaiBetaMetadataStoreStateToProto(o *beta.MetadataStoreState) *betapb.VertexaiBetaMetadataStoreState {
	if o == nil {
		return nil
	}
	p := &betapb.VertexaiBetaMetadataStoreState{}
	p.SetDiskUtilizationBytes(dcl.ValueOrEmptyInt64(o.DiskUtilizationBytes))
	return p
}

// MetadataStoreToProto converts a MetadataStore resource to its proto representation.
func MetadataStoreToProto(resource *beta.MetadataStore) *betapb.VertexaiBetaMetadataStore {
	p := &betapb.VertexaiBetaMetadataStore{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetUpdateTime(dcl.ValueOrEmptyString(resource.UpdateTime))
	p.SetEncryptionSpec(VertexaiBetaMetadataStoreEncryptionSpecToProto(resource.EncryptionSpec))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(VertexaiBetaMetadataStoreStateToProto(resource.State))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))

	return p
}

// applyMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Apply() method.
func (s *MetadataStoreServer) applyMetadataStore(ctx context.Context, c *beta.Client, request *betapb.ApplyVertexaiBetaMetadataStoreRequest) (*betapb.VertexaiBetaMetadataStore, error) {
	p := ProtoToMetadataStore(request.GetResource())
	res, err := c.ApplyMetadataStore(ctx, p)
	if err != nil {
		return nil, err
	}
	r := MetadataStoreToProto(res)
	return r, nil
}

// applyVertexaiBetaMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Apply() method.
func (s *MetadataStoreServer) ApplyVertexaiBetaMetadataStore(ctx context.Context, request *betapb.ApplyVertexaiBetaMetadataStoreRequest) (*betapb.VertexaiBetaMetadataStore, error) {
	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyMetadataStore(ctx, cl, request)
}

// DeleteMetadataStore handles the gRPC request by passing it to the underlying MetadataStore Delete() method.
func (s *MetadataStoreServer) DeleteVertexaiBetaMetadataStore(ctx context.Context, request *betapb.DeleteVertexaiBetaMetadataStoreRequest) (*emptypb.Empty, error) {

	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteMetadataStore(ctx, ProtoToMetadataStore(request.GetResource()))

}

// ListVertexaiBetaMetadataStore handles the gRPC request by passing it to the underlying MetadataStoreList() method.
func (s *MetadataStoreServer) ListVertexaiBetaMetadataStore(ctx context.Context, request *betapb.ListVertexaiBetaMetadataStoreRequest) (*betapb.ListVertexaiBetaMetadataStoreResponse, error) {
	cl, err := createConfigMetadataStore(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListMetadataStore(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.VertexaiBetaMetadataStore
	for _, r := range resources.Items {
		rp := MetadataStoreToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListVertexaiBetaMetadataStoreResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigMetadataStore(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
