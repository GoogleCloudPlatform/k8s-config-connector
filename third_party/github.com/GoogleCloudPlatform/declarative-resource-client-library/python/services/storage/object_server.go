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

// Server implements the gRPC interface for Object.
type ObjectServer struct{}

// ProtoToObjectOwner converts a ObjectOwner resource from its proto representation.
func ProtoToStorageObjectOwner(p *storagepb.StorageObjectOwner) *storage.ObjectOwner {
	if p == nil {
		return nil
	}
	obj := &storage.ObjectOwner{
		Entity:   dcl.StringOrNil(p.Entity),
		EntityId: dcl.StringOrNil(p.EntityId),
	}
	return obj
}

// ProtoToObjectCustomerEncryption converts a ObjectCustomerEncryption resource from its proto representation.
func ProtoToStorageObjectCustomerEncryption(p *storagepb.StorageObjectCustomerEncryption) *storage.ObjectCustomerEncryption {
	if p == nil {
		return nil
	}
	obj := &storage.ObjectCustomerEncryption{
		EncryptionAlgorithm: dcl.StringOrNil(p.EncryptionAlgorithm),
		KeySha256:           dcl.StringOrNil(p.KeySha256),
		Key:                 dcl.StringOrNil(p.Key),
	}
	return obj
}

// ProtoToObject converts a Object resource from its proto representation.
func ProtoToObject(p *storagepb.StorageObject) *storage.Object {
	obj := &storage.Object{
		Name:                    dcl.StringOrNil(p.Name),
		Bucket:                  dcl.StringOrNil(p.Bucket),
		Generation:              dcl.Int64OrNil(p.Generation),
		Metageneration:          dcl.Int64OrNil(p.Metageneration),
		Id:                      dcl.StringOrNil(p.Id),
		SelfLink:                dcl.StringOrNil(p.SelfLink),
		ContentType:             dcl.StringOrNil(p.ContentType),
		TimeCreated:             dcl.StringOrNil(p.GetTimeCreated()),
		Updated:                 dcl.StringOrNil(p.GetUpdated()),
		CustomTime:              dcl.StringOrNil(p.GetCustomTime()),
		TimeDeleted:             dcl.StringOrNil(p.GetTimeDeleted()),
		TemporaryHold:           dcl.Bool(p.TemporaryHold),
		EventBasedHold:          dcl.Bool(p.EventBasedHold),
		RetentionExpirationTime: dcl.StringOrNil(p.GetRetentionExpirationTime()),
		StorageClass:            dcl.StringOrNil(p.StorageClass),
		TimeStorageClassUpdated: dcl.StringOrNil(p.GetTimeStorageClassUpdated()),
		Size:                    dcl.Int64OrNil(p.Size),
		Md5Hash:                 dcl.StringOrNil(p.Md5Hash),
		MediaLink:               dcl.StringOrNil(p.MediaLink),
		Owner:                   ProtoToStorageObjectOwner(p.GetOwner()),
		Crc32c:                  dcl.StringOrNil(p.Crc32C),
		ComponentCount:          dcl.Int64OrNil(p.ComponentCount),
		Etag:                    dcl.StringOrNil(p.Etag),
		CustomerEncryption:      ProtoToStorageObjectCustomerEncryption(p.GetCustomerEncryption()),
		KmsKeyName:              dcl.StringOrNil(p.KmsKeyName),
		Content:                 dcl.StringOrNil(p.Content),
	}
	return obj
}

// ObjectOwnerToProto converts a ObjectOwner resource to its proto representation.
func StorageObjectOwnerToProto(o *storage.ObjectOwner) *storagepb.StorageObjectOwner {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageObjectOwner{
		Entity:   dcl.ValueOrEmptyString(o.Entity),
		EntityId: dcl.ValueOrEmptyString(o.EntityId),
	}
	return p
}

// ObjectCustomerEncryptionToProto converts a ObjectCustomerEncryption resource to its proto representation.
func StorageObjectCustomerEncryptionToProto(o *storage.ObjectCustomerEncryption) *storagepb.StorageObjectCustomerEncryption {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageObjectCustomerEncryption{
		EncryptionAlgorithm: dcl.ValueOrEmptyString(o.EncryptionAlgorithm),
		KeySha256:           dcl.ValueOrEmptyString(o.KeySha256),
		Key:                 dcl.ValueOrEmptyString(o.Key),
	}
	return p
}

// ObjectToProto converts a Object resource to its proto representation.
func ObjectToProto(resource *storage.Object) *storagepb.StorageObject {
	p := &storagepb.StorageObject{
		Name:                    dcl.ValueOrEmptyString(resource.Name),
		Bucket:                  dcl.ValueOrEmptyString(resource.Bucket),
		Generation:              dcl.ValueOrEmptyInt64(resource.Generation),
		Metageneration:          dcl.ValueOrEmptyInt64(resource.Metageneration),
		Id:                      dcl.ValueOrEmptyString(resource.Id),
		SelfLink:                dcl.ValueOrEmptyString(resource.SelfLink),
		ContentType:             dcl.ValueOrEmptyString(resource.ContentType),
		TimeCreated:             dcl.ValueOrEmptyString(resource.TimeCreated),
		Updated:                 dcl.ValueOrEmptyString(resource.Updated),
		CustomTime:              dcl.ValueOrEmptyString(resource.CustomTime),
		TimeDeleted:             dcl.ValueOrEmptyString(resource.TimeDeleted),
		TemporaryHold:           dcl.ValueOrEmptyBool(resource.TemporaryHold),
		EventBasedHold:          dcl.ValueOrEmptyBool(resource.EventBasedHold),
		RetentionExpirationTime: dcl.ValueOrEmptyString(resource.RetentionExpirationTime),
		StorageClass:            dcl.ValueOrEmptyString(resource.StorageClass),
		TimeStorageClassUpdated: dcl.ValueOrEmptyString(resource.TimeStorageClassUpdated),
		Size:                    dcl.ValueOrEmptyInt64(resource.Size),
		Md5Hash:                 dcl.ValueOrEmptyString(resource.Md5Hash),
		MediaLink:               dcl.ValueOrEmptyString(resource.MediaLink),
		Owner:                   StorageObjectOwnerToProto(resource.Owner),
		Crc32C:                  dcl.ValueOrEmptyString(resource.Crc32c),
		ComponentCount:          dcl.ValueOrEmptyInt64(resource.ComponentCount),
		Etag:                    dcl.ValueOrEmptyString(resource.Etag),
		CustomerEncryption:      StorageObjectCustomerEncryptionToProto(resource.CustomerEncryption),
		KmsKeyName:              dcl.ValueOrEmptyString(resource.KmsKeyName),
		Content:                 dcl.ValueOrEmptyString(resource.Content),
	}

	return p
}

// ApplyObject handles the gRPC request by passing it to the underlying Object Apply() method.
func (s *ObjectServer) applyObject(ctx context.Context, c *storage.Client, request *storagepb.ApplyStorageObjectRequest) (*storagepb.StorageObject, error) {
	p := ProtoToObject(request.GetResource())
	res, err := c.ApplyObject(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ObjectToProto(res)
	return r, nil
}

// ApplyObject handles the gRPC request by passing it to the underlying Object Apply() method.
func (s *ObjectServer) ApplyStorageObject(ctx context.Context, request *storagepb.ApplyStorageObjectRequest) (*storagepb.StorageObject, error) {
	cl, err := createConfigObject(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyObject(ctx, cl, request)
}

// DeleteObject handles the gRPC request by passing it to the underlying Object Delete() method.
func (s *ObjectServer) DeleteStorageObject(ctx context.Context, request *storagepb.DeleteStorageObjectRequest) (*emptypb.Empty, error) {

	cl, err := createConfigObject(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteObject(ctx, ProtoToObject(request.GetResource()))

}

// ListStorageObject handles the gRPC request by passing it to the underlying ObjectList() method.
func (s *ObjectServer) ListStorageObject(ctx context.Context, request *storagepb.ListStorageObjectRequest) (*storagepb.ListStorageObjectResponse, error) {
	cl, err := createConfigObject(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListObject(ctx, request.Bucket)
	if err != nil {
		return nil, err
	}
	var protos []*storagepb.StorageObject
	for _, r := range resources.Items {
		rp := ObjectToProto(r)
		protos = append(protos, rp)
	}
	return &storagepb.ListStorageObjectResponse{Items: protos}, nil
}

func createConfigObject(ctx context.Context, service_account_file string) (*storage.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return storage.NewClient(conf), nil
}
