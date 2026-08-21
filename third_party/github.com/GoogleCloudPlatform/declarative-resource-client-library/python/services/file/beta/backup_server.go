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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/file/beta/file_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/file/beta"
)

// Server implements the gRPC interface for Backup.
type BackupServer struct{}

// ProtoToBackupStateEnum converts a BackupStateEnum enum from its proto representation.
func ProtoToFileBetaBackupStateEnum(e betapb.FileBetaBackupStateEnum) *beta.BackupStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FileBetaBackupStateEnum_name[int32(e)]; ok {
		e := beta.BackupStateEnum(n[len("FileBetaBackupStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackupSourceInstanceTierEnum converts a BackupSourceInstanceTierEnum enum from its proto representation.
func ProtoToFileBetaBackupSourceInstanceTierEnum(e betapb.FileBetaBackupSourceInstanceTierEnum) *beta.BackupSourceInstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FileBetaBackupSourceInstanceTierEnum_name[int32(e)]; ok {
		e := beta.BackupSourceInstanceTierEnum(n[len("FileBetaBackupSourceInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackup converts a Backup resource from its proto representation.
func ProtoToBackup(p *betapb.FileBetaBackup) *beta.Backup {
	obj := &beta.Backup{
		Name:               dcl.StringOrNil(p.Name),
		Description:        dcl.StringOrNil(p.Description),
		State:              ProtoToFileBetaBackupStateEnum(p.GetState()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		CapacityGb:         dcl.Int64OrNil(p.CapacityGb),
		StorageBytes:       dcl.Int64OrNil(p.StorageBytes),
		SourceInstance:     dcl.StringOrNil(p.SourceInstance),
		SourceFileShare:    dcl.StringOrNil(p.SourceFileShare),
		SourceInstanceTier: ProtoToFileBetaBackupSourceInstanceTierEnum(p.GetSourceInstanceTier()),
		DownloadBytes:      dcl.Int64OrNil(p.DownloadBytes),
		Project:            dcl.StringOrNil(p.Project),
		Location:           dcl.StringOrNil(p.Location),
	}
	return obj
}

// BackupStateEnumToProto converts a BackupStateEnum enum to its proto representation.
func FileBetaBackupStateEnumToProto(e *beta.BackupStateEnum) betapb.FileBetaBackupStateEnum {
	if e == nil {
		return betapb.FileBetaBackupStateEnum(0)
	}
	if v, ok := betapb.FileBetaBackupStateEnum_value["BackupStateEnum"+string(*e)]; ok {
		return betapb.FileBetaBackupStateEnum(v)
	}
	return betapb.FileBetaBackupStateEnum(0)
}

// BackupSourceInstanceTierEnumToProto converts a BackupSourceInstanceTierEnum enum to its proto representation.
func FileBetaBackupSourceInstanceTierEnumToProto(e *beta.BackupSourceInstanceTierEnum) betapb.FileBetaBackupSourceInstanceTierEnum {
	if e == nil {
		return betapb.FileBetaBackupSourceInstanceTierEnum(0)
	}
	if v, ok := betapb.FileBetaBackupSourceInstanceTierEnum_value["BackupSourceInstanceTierEnum"+string(*e)]; ok {
		return betapb.FileBetaBackupSourceInstanceTierEnum(v)
	}
	return betapb.FileBetaBackupSourceInstanceTierEnum(0)
}

// BackupToProto converts a Backup resource to its proto representation.
func BackupToProto(resource *beta.Backup) *betapb.FileBetaBackup {
	p := &betapb.FileBetaBackup{
		Name:               dcl.ValueOrEmptyString(resource.Name),
		Description:        dcl.ValueOrEmptyString(resource.Description),
		State:              FileBetaBackupStateEnumToProto(resource.State),
		CreateTime:         dcl.ValueOrEmptyString(resource.CreateTime),
		CapacityGb:         dcl.ValueOrEmptyInt64(resource.CapacityGb),
		StorageBytes:       dcl.ValueOrEmptyInt64(resource.StorageBytes),
		SourceInstance:     dcl.ValueOrEmptyString(resource.SourceInstance),
		SourceFileShare:    dcl.ValueOrEmptyString(resource.SourceFileShare),
		SourceInstanceTier: FileBetaBackupSourceInstanceTierEnumToProto(resource.SourceInstanceTier),
		DownloadBytes:      dcl.ValueOrEmptyInt64(resource.DownloadBytes),
		Project:            dcl.ValueOrEmptyString(resource.Project),
		Location:           dcl.ValueOrEmptyString(resource.Location),
	}

	return p
}

// ApplyBackup handles the gRPC request by passing it to the underlying Backup Apply() method.
func (s *BackupServer) applyBackup(ctx context.Context, c *beta.Client, request *betapb.ApplyFileBetaBackupRequest) (*betapb.FileBetaBackup, error) {
	p := ProtoToBackup(request.GetResource())
	res, err := c.ApplyBackup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BackupToProto(res)
	return r, nil
}

// ApplyBackup handles the gRPC request by passing it to the underlying Backup Apply() method.
func (s *BackupServer) ApplyFileBetaBackup(ctx context.Context, request *betapb.ApplyFileBetaBackupRequest) (*betapb.FileBetaBackup, error) {
	cl, err := createConfigBackup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyBackup(ctx, cl, request)
}

// DeleteBackup handles the gRPC request by passing it to the underlying Backup Delete() method.
func (s *BackupServer) DeleteFileBetaBackup(ctx context.Context, request *betapb.DeleteFileBetaBackupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBackup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBackup(ctx, ProtoToBackup(request.GetResource()))

}

// ListFileBetaBackup handles the gRPC request by passing it to the underlying BackupList() method.
func (s *BackupServer) ListFileBetaBackup(ctx context.Context, request *betapb.ListFileBetaBackupRequest) (*betapb.ListFileBetaBackupResponse, error) {
	cl, err := createConfigBackup(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBackup(ctx, ProtoToBackup(request.GetResource()))
	if err != nil {
		return nil, err
	}
	var protos []*betapb.FileBetaBackup
	for _, r := range resources.Items {
		rp := BackupToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListFileBetaBackupResponse{Items: protos}, nil
}

func createConfigBackup(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
