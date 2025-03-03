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
	storagepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/storage/storage_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage"
)

// Server implements the gRPC interface for HmacKey.
type HmacKeyServer struct{}

// ProtoToHmacKeyStateEnum converts a HmacKeyStateEnum enum from its proto representation.
func ProtoToStorageHmacKeyStateEnum(e storagepb.StorageHmacKeyStateEnum) *storage.HmacKeyStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := storagepb.StorageHmacKeyStateEnum_name[int32(e)]; ok {
		e := storage.HmacKeyStateEnum(n[len("StorageHmacKeyStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToHmacKey converts a HmacKey resource from its proto representation.
func ProtoToHmacKey(p *storagepb.StorageHmacKey) *storage.HmacKey {
	obj := &storage.HmacKey{
		Name:                dcl.StringOrNil(p.Name),
		TimeCreated:         dcl.StringOrNil(p.TimeCreated),
		Updated:             dcl.StringOrNil(p.Updated),
		Secret:              dcl.StringOrNil(p.Secret),
		State:               ProtoToStorageHmacKeyStateEnum(p.GetState()),
		Project:             dcl.StringOrNil(p.Project),
		ServiceAccountEmail: dcl.StringOrNil(p.ServiceAccountEmail),
	}
	return obj
}

// HmacKeyStateEnumToProto converts a HmacKeyStateEnum enum to its proto representation.
func StorageHmacKeyStateEnumToProto(e *storage.HmacKeyStateEnum) storagepb.StorageHmacKeyStateEnum {
	if e == nil {
		return storagepb.StorageHmacKeyStateEnum(0)
	}
	if v, ok := storagepb.StorageHmacKeyStateEnum_value["HmacKeyStateEnum"+string(*e)]; ok {
		return storagepb.StorageHmacKeyStateEnum(v)
	}
	return storagepb.StorageHmacKeyStateEnum(0)
}

// HmacKeyToProto converts a HmacKey resource to its proto representation.
func HmacKeyToProto(resource *storage.HmacKey) *storagepb.StorageHmacKey {
	p := &storagepb.StorageHmacKey{
		Name:                dcl.ValueOrEmptyString(resource.Name),
		TimeCreated:         dcl.ValueOrEmptyString(resource.TimeCreated),
		Updated:             dcl.ValueOrEmptyString(resource.Updated),
		Secret:              dcl.ValueOrEmptyString(resource.Secret),
		State:               StorageHmacKeyStateEnumToProto(resource.State),
		Project:             dcl.ValueOrEmptyString(resource.Project),
		ServiceAccountEmail: dcl.ValueOrEmptyString(resource.ServiceAccountEmail),
	}

	return p
}

// ApplyHmacKey handles the gRPC request by passing it to the underlying HmacKey Apply() method.
func (s *HmacKeyServer) applyHmacKey(ctx context.Context, c *storage.Client, request *storagepb.ApplyStorageHmacKeyRequest) (*storagepb.StorageHmacKey, error) {
	p := ProtoToHmacKey(request.GetResource())
	res, err := c.ApplyHmacKey(ctx, p)
	if err != nil {
		return nil, err
	}
	r := HmacKeyToProto(res)
	return r, nil
}

// ApplyHmacKey handles the gRPC request by passing it to the underlying HmacKey Apply() method.
func (s *HmacKeyServer) ApplyStorageHmacKey(ctx context.Context, request *storagepb.ApplyStorageHmacKeyRequest) (*storagepb.StorageHmacKey, error) {
	cl, err := createConfigHmacKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyHmacKey(ctx, cl, request)
}

// DeleteHmacKey handles the gRPC request by passing it to the underlying HmacKey Delete() method.
func (s *HmacKeyServer) DeleteStorageHmacKey(ctx context.Context, request *storagepb.DeleteStorageHmacKeyRequest) (*emptypb.Empty, error) {

	cl, err := createConfigHmacKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteHmacKey(ctx, ProtoToHmacKey(request.GetResource()))

}

// ListStorageHmacKey handles the gRPC request by passing it to the underlying HmacKeyList() method.
func (s *HmacKeyServer) ListStorageHmacKey(ctx context.Context, request *storagepb.ListStorageHmacKeyRequest) (*storagepb.ListStorageHmacKeyResponse, error) {
	cl, err := createConfigHmacKey(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListHmacKey(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*storagepb.StorageHmacKey
	for _, r := range resources.Items {
		rp := HmacKeyToProto(r)
		protos = append(protos, rp)
	}
	return &storagepb.ListStorageHmacKeyResponse{Items: protos}, nil
}

func createConfigHmacKey(ctx context.Context, service_account_file string) (*storage.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return storage.NewClient(conf), nil
}
