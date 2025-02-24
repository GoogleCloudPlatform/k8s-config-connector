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

// Server implements the gRPC interface for DefaultObjectAccessControl.
type DefaultObjectAccessControlServer struct{}

// ProtoToDefaultObjectAccessControlProjectTeamTeamEnum converts a DefaultObjectAccessControlProjectTeamTeamEnum enum from its proto representation.
func ProtoToStorageDefaultObjectAccessControlProjectTeamTeamEnum(e storagepb.StorageDefaultObjectAccessControlProjectTeamTeamEnum) *storage.DefaultObjectAccessControlProjectTeamTeamEnum {
	if e == 0 {
		return nil
	}
	if n, ok := storagepb.StorageDefaultObjectAccessControlProjectTeamTeamEnum_name[int32(e)]; ok {
		e := storage.DefaultObjectAccessControlProjectTeamTeamEnum(n[len("StorageDefaultObjectAccessControlProjectTeamTeamEnum"):])
		return &e
	}
	return nil
}

// ProtoToDefaultObjectAccessControlRoleEnum converts a DefaultObjectAccessControlRoleEnum enum from its proto representation.
func ProtoToStorageDefaultObjectAccessControlRoleEnum(e storagepb.StorageDefaultObjectAccessControlRoleEnum) *storage.DefaultObjectAccessControlRoleEnum {
	if e == 0 {
		return nil
	}
	if n, ok := storagepb.StorageDefaultObjectAccessControlRoleEnum_name[int32(e)]; ok {
		e := storage.DefaultObjectAccessControlRoleEnum(n[len("StorageDefaultObjectAccessControlRoleEnum"):])
		return &e
	}
	return nil
}

// ProtoToDefaultObjectAccessControlProjectTeam converts a DefaultObjectAccessControlProjectTeam resource from its proto representation.
func ProtoToStorageDefaultObjectAccessControlProjectTeam(p *storagepb.StorageDefaultObjectAccessControlProjectTeam) *storage.DefaultObjectAccessControlProjectTeam {
	if p == nil {
		return nil
	}
	obj := &storage.DefaultObjectAccessControlProjectTeam{
		ProjectNumber: dcl.StringOrNil(p.ProjectNumber),
		Team:          ProtoToStorageDefaultObjectAccessControlProjectTeamTeamEnum(p.GetTeam()),
	}
	return obj
}

// ProtoToDefaultObjectAccessControl converts a DefaultObjectAccessControl resource from its proto representation.
func ProtoToDefaultObjectAccessControl(p *storagepb.StorageDefaultObjectAccessControl) *storage.DefaultObjectAccessControl {
	obj := &storage.DefaultObjectAccessControl{
		Project:     dcl.StringOrNil(p.Project),
		Bucket:      dcl.StringOrNil(p.Bucket),
		Domain:      dcl.StringOrNil(p.Domain),
		Email:       dcl.StringOrNil(p.Email),
		Entity:      dcl.StringOrNil(p.Entity),
		EntityId:    dcl.StringOrNil(p.EntityId),
		ProjectTeam: ProtoToStorageDefaultObjectAccessControlProjectTeam(p.GetProjectTeam()),
		Role:        ProtoToStorageDefaultObjectAccessControlRoleEnum(p.GetRole()),
	}
	return obj
}

// DefaultObjectAccessControlProjectTeamTeamEnumToProto converts a DefaultObjectAccessControlProjectTeamTeamEnum enum to its proto representation.
func StorageDefaultObjectAccessControlProjectTeamTeamEnumToProto(e *storage.DefaultObjectAccessControlProjectTeamTeamEnum) storagepb.StorageDefaultObjectAccessControlProjectTeamTeamEnum {
	if e == nil {
		return storagepb.StorageDefaultObjectAccessControlProjectTeamTeamEnum(0)
	}
	if v, ok := storagepb.StorageDefaultObjectAccessControlProjectTeamTeamEnum_value["DefaultObjectAccessControlProjectTeamTeamEnum"+string(*e)]; ok {
		return storagepb.StorageDefaultObjectAccessControlProjectTeamTeamEnum(v)
	}
	return storagepb.StorageDefaultObjectAccessControlProjectTeamTeamEnum(0)
}

// DefaultObjectAccessControlRoleEnumToProto converts a DefaultObjectAccessControlRoleEnum enum to its proto representation.
func StorageDefaultObjectAccessControlRoleEnumToProto(e *storage.DefaultObjectAccessControlRoleEnum) storagepb.StorageDefaultObjectAccessControlRoleEnum {
	if e == nil {
		return storagepb.StorageDefaultObjectAccessControlRoleEnum(0)
	}
	if v, ok := storagepb.StorageDefaultObjectAccessControlRoleEnum_value["DefaultObjectAccessControlRoleEnum"+string(*e)]; ok {
		return storagepb.StorageDefaultObjectAccessControlRoleEnum(v)
	}
	return storagepb.StorageDefaultObjectAccessControlRoleEnum(0)
}

// DefaultObjectAccessControlProjectTeamToProto converts a DefaultObjectAccessControlProjectTeam resource to its proto representation.
func StorageDefaultObjectAccessControlProjectTeamToProto(o *storage.DefaultObjectAccessControlProjectTeam) *storagepb.StorageDefaultObjectAccessControlProjectTeam {
	if o == nil {
		return nil
	}
	p := &storagepb.StorageDefaultObjectAccessControlProjectTeam{
		ProjectNumber: dcl.ValueOrEmptyString(o.ProjectNumber),
		Team:          StorageDefaultObjectAccessControlProjectTeamTeamEnumToProto(o.Team),
	}
	return p
}

// DefaultObjectAccessControlToProto converts a DefaultObjectAccessControl resource to its proto representation.
func DefaultObjectAccessControlToProto(resource *storage.DefaultObjectAccessControl) *storagepb.StorageDefaultObjectAccessControl {
	p := &storagepb.StorageDefaultObjectAccessControl{
		Project:     dcl.ValueOrEmptyString(resource.Project),
		Bucket:      dcl.ValueOrEmptyString(resource.Bucket),
		Domain:      dcl.ValueOrEmptyString(resource.Domain),
		Email:       dcl.ValueOrEmptyString(resource.Email),
		Entity:      dcl.ValueOrEmptyString(resource.Entity),
		EntityId:    dcl.ValueOrEmptyString(resource.EntityId),
		ProjectTeam: StorageDefaultObjectAccessControlProjectTeamToProto(resource.ProjectTeam),
		Role:        StorageDefaultObjectAccessControlRoleEnumToProto(resource.Role),
	}

	return p
}

// ApplyDefaultObjectAccessControl handles the gRPC request by passing it to the underlying DefaultObjectAccessControl Apply() method.
func (s *DefaultObjectAccessControlServer) applyDefaultObjectAccessControl(ctx context.Context, c *storage.Client, request *storagepb.ApplyStorageDefaultObjectAccessControlRequest) (*storagepb.StorageDefaultObjectAccessControl, error) {
	p := ProtoToDefaultObjectAccessControl(request.GetResource())
	res, err := c.ApplyDefaultObjectAccessControl(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DefaultObjectAccessControlToProto(res)
	return r, nil
}

// ApplyDefaultObjectAccessControl handles the gRPC request by passing it to the underlying DefaultObjectAccessControl Apply() method.
func (s *DefaultObjectAccessControlServer) ApplyStorageDefaultObjectAccessControl(ctx context.Context, request *storagepb.ApplyStorageDefaultObjectAccessControlRequest) (*storagepb.StorageDefaultObjectAccessControl, error) {
	cl, err := createConfigDefaultObjectAccessControl(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyDefaultObjectAccessControl(ctx, cl, request)
}

// DeleteDefaultObjectAccessControl handles the gRPC request by passing it to the underlying DefaultObjectAccessControl Delete() method.
func (s *DefaultObjectAccessControlServer) DeleteStorageDefaultObjectAccessControl(ctx context.Context, request *storagepb.DeleteStorageDefaultObjectAccessControlRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDefaultObjectAccessControl(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDefaultObjectAccessControl(ctx, ProtoToDefaultObjectAccessControl(request.GetResource()))

}

// ListStorageDefaultObjectAccessControl handles the gRPC request by passing it to the underlying DefaultObjectAccessControlList() method.
func (s *DefaultObjectAccessControlServer) ListStorageDefaultObjectAccessControl(ctx context.Context, request *storagepb.ListStorageDefaultObjectAccessControlRequest) (*storagepb.ListStorageDefaultObjectAccessControlResponse, error) {
	cl, err := createConfigDefaultObjectAccessControl(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDefaultObjectAccessControl(ctx, request.Project, request.Bucket)
	if err != nil {
		return nil, err
	}
	var protos []*storagepb.StorageDefaultObjectAccessControl
	for _, r := range resources.Items {
		rp := DefaultObjectAccessControlToProto(r)
		protos = append(protos, rp)
	}
	return &storagepb.ListStorageDefaultObjectAccessControlResponse{Items: protos}, nil
}

func createConfigDefaultObjectAccessControl(ctx context.Context, service_account_file string) (*storage.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return storage.NewClient(conf), nil
}
