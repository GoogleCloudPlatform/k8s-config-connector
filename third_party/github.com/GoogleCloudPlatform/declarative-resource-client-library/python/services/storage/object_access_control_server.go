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

// Server implements the gRPC interface for ObjectAccessControl.
type ObjectAccessControlServer struct{}

// ProtoToObjectAccessControlProjectTeamTeamEnum converts a ObjectAccessControlProjectTeamTeamEnum enum from its proto representation.
func ProtoToStorageObjectAccessControlProjectTeamTeamEnum(e storagepb.StorageObjectAccessControlProjectTeamTeamEnum) *storage.ObjectAccessControlProjectTeamTeamEnum {
	if e == 0 {
		return nil
	}
	if n, ok := storagepb.StorageObjectAccessControlProjectTeamTeamEnum_name[int32(e)]; ok {
		e := storage.ObjectAccessControlProjectTeamTeamEnum(n[len("StorageObjectAccessControlProjectTeamTeamEnum"):])
		return &e
	}
	return nil
}

// ProtoToObjectAccessControlRoleEnum converts a ObjectAccessControlRoleEnum enum from its proto representation.
func ProtoToStorageObjectAccessControlRoleEnum(e storagepb.StorageObjectAccessControlRoleEnum) *storage.ObjectAccessControlRoleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := storagepb.StorageObjectAccessControlRoleEnum_name[int32(e)]; ok {
		e := storage.ObjectAccessControlRoleEnum(n[len("StorageObjectAccessControlRoleEnum"):])
		return &e
	}
	return nil
}

// ProtoToObjectAccessControlProjectTeam converts a ObjectAccessControlProjectTeam resource from its proto representation.
func ProtoToStorageObjectAccessControlProjectTeam(p *storagepb.StorageObjectAccessControlProjectTeam) *storage.ObjectAccessControlProjectTeam {
	if p == nil {
		return nil
	}
	obj := &storage.ObjectAccessControlProjectTeam{
		ProjectNumber: dcl.StringOrNil(p.ProjectNumber),
		Team:          ProtoToStorageObjectAccessControlProjectTeamTeamEnum(p.GetTeam()),
	}
	return obj
}

// ProtoToObjectAccessControl converts a ObjectAccessControl resource from its proto representation.
func ProtoToObjectAccessControl(p *storagepb.StorageObjectAccessControl) *storage.ObjectAccessControl {
	obj := &storage.ObjectAccessControl{
		Project:     dcl.StringOrNil(p.Project),
		Bucket:      dcl.StringOrNil(p.Bucket),
		Domain:      dcl.StringOrNil(p.Domain),
		Email:       dcl.StringOrNil(p.Email),
		Entity:      dcl.StringOrNil(p.Entity),
		EntityId:    dcl.StringOrNil(p.EntityId),
		ProjectTeam: ProtoToStorageObjectAccessControlProjectTeam(p.GetProjectTeam()),
		Role:        ProtoToStorageObjectAccessControlRoleEnum(p.GetRole()),
		Id:          dcl.StringOrNil(p.Id),
		Object:      dcl.StringOrNil(p.Object),
		Generation:  dcl.Int64OrNil(p.Generation),
	}
	return obj
}

// ObjectAccessControlProjectTeamTeamEnumToProto converts a ObjectAccessControlProjectTeamTeamEnum enum to its proto representation.
func StorageObjectAccessControlProjectTeamTeamEnumToProto(e *storage.ObjectAccessControlProjectTeamTeamEnum) storagepb.StorageObjectAccessControlProjectTeamTeamEnum {
	if e == nil {
		return storagepb.StorageObjectAccessControlProjectTeamTeamEnum(0)
	}
	if v, ok := storagepb.StorageObjectAccessControlProjectTeamTeamEnum_value["ObjectAccessControlProjectTeamTeamEnum"+string(*e)]; ok {
		return storagepb.StorageObjectAccessControlProjectTeamTeamEnum(v)
	}
	return storagepb.StorageObjectAccessControlProjectTeamTeamEnum(0)
}

// ObjectAccessControlRoleEnumToProto converts a ObjectAccessControlRoleEnum enum to its proto representation.
func StorageObjectAccessControlRoleEnumToProto(e *storage.ObjectAccessControlRoleEnum) storagepb.StorageObjectAccessControlRoleEnum {
	if e == nil {
		return storagepb.StorageObjectAccessControlRoleEnum(0)
	}
	if v, ok := storagepb.StorageObjectAccessControlRoleEnum_value["ObjectAccessControlRoleEnum"+string(*e)]; ok {
		return storagepb.StorageObjectAccessControlRoleEnum(v)
	}
	return storagepb.StorageObjectAccessControlRoleEnum(0)
}

// ObjectAccessControlProjectTeamToProto converts a ObjectAccessControlProjectTeam resource to its proto representation.
func StorageObjectAccessControlProjectTeamToProto(o *storage.ObjectAccessControlProjectTeam) *storagepb.StorageObjectAccessControlProjectTeam {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageObjectAccessControlProjectTeam{
		ProjectNumber: dcl.ValueOrEmptyString(o.ProjectNumber),
		Team:          StorageObjectAccessControlProjectTeamTeamEnumToProto(o.Team),
	}
	return p
}

// ObjectAccessControlToProto converts a ObjectAccessControl resource to its proto representation.
func ObjectAccessControlToProto(resource *storage.ObjectAccessControl) *storagepb.StorageObjectAccessControl {
	p := &storagepb.StorageObjectAccessControl{
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Bucket:      dcl.ValueOrEmptyString(resource.Bucket),
		Domain:      dcl.ValueOrEmptyString(resource.Domain),
		Email:       dcl.ValueOrEmptyString(resource.Email),
		Entity:      dcl.ValueOrEmptyString(resource.Entity),
		EntityId:    dcl.ValueOrEmptyString(resource.EntityId),
		ProjectTeam: StorageObjectAccessControlProjectTeamToProto(resource.ProjectTeam),
		Role:        StorageObjectAccessControlRoleEnumToProto(resource.Role),
		Id:          dcl.ValueOrEmptyString(resource.Id),
		Object:      dcl.ValueOrEmptyString(resource.Object),
		Generation:  dcl.ValueOrEmptyInt64(resource.Generation),
	}

	return p
}

// ApplyObjectAccessControl handles the gRPC request by passing it to the underlying ObjectAccessControl Apply() method.
func (s *ObjectAccessControlServer) applyObjectAccessControl(ctx context.Context, c *storage.Client, request *storagepb.ApplyStorageObjectAccessControlRequest) (*storagepb.StorageObjectAccessControl, error) {
	p := ProtoToObjectAccessControl(request.GetResource())
	res, err := c.ApplyObjectAccessControl(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ObjectAccessControlToProto(res)
	return r, nil
}

// ApplyObjectAccessControl handles the gRPC request by passing it to the underlying ObjectAccessControl Apply() method.
func (s *ObjectAccessControlServer) ApplyStorageObjectAccessControl(ctx context.Context, request *storagepb.ApplyStorageObjectAccessControlRequest) (*storagepb.StorageObjectAccessControl, error) {
	cl, err := createConfigObjectAccessControl(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyObjectAccessControl(ctx, cl, request)
}

// DeleteObjectAccessControl handles the gRPC request by passing it to the underlying ObjectAccessControl Delete() method.
func (s *ObjectAccessControlServer) DeleteStorageObjectAccessControl(ctx context.Context, request *storagepb.DeleteStorageObjectAccessControlRequest) (*emptypb.Empty, error) {

	cl, err := createConfigObjectAccessControl(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteObjectAccessControl(ctx, ProtoToObjectAccessControl(request.GetResource()))

}

// ListStorageObjectAccessControl handles the gRPC request by passing it to the underlying ObjectAccessControlList() method.
func (s *ObjectAccessControlServer) ListStorageObjectAccessControl(ctx context.Context, request *storagepb.ListStorageObjectAccessControlRequest) (*storagepb.ListStorageObjectAccessControlResponse, error) {
	cl, err := createConfigObjectAccessControl(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListObjectAccessControl(ctx, request.Project, request.Bucket, request.Object)
	if err != nil {
		return nil, err
	}
	var protos []*storagepb.StorageObjectAccessControl
	for _, r := range resources.Items {
		rp := ObjectAccessControlToProto(r)
		protos = append(protos, rp)
	}
	return &storagepb.ListStorageObjectAccessControlResponse{Items: protos}, nil
}

func createConfigObjectAccessControl(ctx context.Context, service_account_file string) (*storage.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return storage.NewClient(conf), nil
}
