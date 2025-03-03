// Copyright 2024 Google LLC. All Rights Reserved.
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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/filestore/beta/filestore_beta_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/beta"
)

// BackupServer implements the gRPC interface for Backup.
type BackupServer struct{}

// ProtoToBackupStateEnum converts a BackupStateEnum enum from its proto representation.
func ProtoToFilestoreBetaBackupStateEnum(e betapb.FilestoreBetaBackupStateEnum) *beta.BackupStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FilestoreBetaBackupStateEnum_name[int32(e)]; ok {
		e := beta.BackupStateEnum(n[len("FilestoreBetaBackupStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackupSourceInstanceTierEnum converts a BackupSourceInstanceTierEnum enum from its proto representation.
func ProtoToFilestoreBetaBackupSourceInstanceTierEnum(e betapb.FilestoreBetaBackupSourceInstanceTierEnum) *beta.BackupSourceInstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.FilestoreBetaBackupSourceInstanceTierEnum_name[int32(e)]; ok {
		e := beta.BackupSourceInstanceTierEnum(n[len("FilestoreBetaBackupSourceInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackup converts a Backup resource from its proto representation.
func ProtoToBackup(p *betapb.FilestoreBetaBackup) *beta.Backup {
	obj := &beta.Backup{
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		State:              ProtoToFilestoreBetaBackupStateEnum(p.GetState()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		CapacityGb:         dcl.Int64OrNil(p.GetCapacityGb()),
		StorageBytes:       dcl.Int64OrNil(p.GetStorageBytes()),
		SourceInstance:     dcl.StringOrNil(p.GetSourceInstance()),
		SourceFileShare:    dcl.StringOrNil(p.GetSourceFileShare()),
		SourceInstanceTier: ProtoToFilestoreBetaBackupSourceInstanceTierEnum(p.GetSourceInstanceTier()),
		DownloadBytes:      dcl.Int64OrNil(p.GetDownloadBytes()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// BackupStateEnumToProto converts a BackupStateEnum enum to its proto representation.
func FilestoreBetaBackupStateEnumToProto(e *beta.BackupStateEnum) betapb.FilestoreBetaBackupStateEnum {
	if e == nil {
		return betapb.FilestoreBetaBackupStateEnum(0)
	}
	if v, ok := betapb.FilestoreBetaBackupStateEnum_value["BackupStateEnum"+string(*e)]; ok {
		return betapb.FilestoreBetaBackupStateEnum(v)
	}
	return betapb.FilestoreBetaBackupStateEnum(0)
}

// BackupSourceInstanceTierEnumToProto converts a BackupSourceInstanceTierEnum enum to its proto representation.
func FilestoreBetaBackupSourceInstanceTierEnumToProto(e *beta.BackupSourceInstanceTierEnum) betapb.FilestoreBetaBackupSourceInstanceTierEnum {
	if e == nil {
		return betapb.FilestoreBetaBackupSourceInstanceTierEnum(0)
	}
	if v, ok := betapb.FilestoreBetaBackupSourceInstanceTierEnum_value["BackupSourceInstanceTierEnum"+string(*e)]; ok {
		return betapb.FilestoreBetaBackupSourceInstanceTierEnum(v)
	}
	return betapb.FilestoreBetaBackupSourceInstanceTierEnum(0)
}

// BackupToProto converts a Backup resource to its proto representation.
func BackupToProto(resource *beta.Backup) *betapb.FilestoreBetaBackup {
	p := &betapb.FilestoreBetaBackup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(FilestoreBetaBackupStateEnumToProto(resource.State))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetCapacityGb(dcl.ValueOrEmptyInt64(resource.CapacityGb))
	p.SetStorageBytes(dcl.ValueOrEmptyInt64(resource.StorageBytes))
	p.SetSourceInstance(dcl.ValueOrEmptyString(resource.SourceInstance))
	p.SetSourceFileShare(dcl.ValueOrEmptyString(resource.SourceFileShare))
	p.SetSourceInstanceTier(FilestoreBetaBackupSourceInstanceTierEnumToProto(resource.SourceInstanceTier))
	p.SetDownloadBytes(dcl.ValueOrEmptyInt64(resource.DownloadBytes))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)

	return p
}

// applyBackup handles the gRPC request by passing it to the underlying Backup Apply() method.
func (s *BackupServer) applyBackup(ctx context.Context, c *beta.Client, request *betapb.ApplyFilestoreBetaBackupRequest) (*betapb.FilestoreBetaBackup, error) {
	p := ProtoToBackup(request.GetResource())
	res, err := c.ApplyBackup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BackupToProto(res)
	return r, nil
}

// applyFilestoreBetaBackup handles the gRPC request by passing it to the underlying Backup Apply() method.
func (s *BackupServer) ApplyFilestoreBetaBackup(ctx context.Context, request *betapb.ApplyFilestoreBetaBackupRequest) (*betapb.FilestoreBetaBackup, error) {
	cl, err := createConfigBackup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyBackup(ctx, cl, request)
}

// DeleteBackup handles the gRPC request by passing it to the underlying Backup Delete() method.
func (s *BackupServer) DeleteFilestoreBetaBackup(ctx context.Context, request *betapb.DeleteFilestoreBetaBackupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBackup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBackup(ctx, ProtoToBackup(request.GetResource()))

}

// ListFilestoreBetaBackup handles the gRPC request by passing it to the underlying BackupList() method.
func (s *BackupServer) ListFilestoreBetaBackup(ctx context.Context, request *betapb.ListFilestoreBetaBackupRequest) (*betapb.ListFilestoreBetaBackupResponse, error) {
	cl, err := createConfigBackup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBackup(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.FilestoreBetaBackup
	for _, r := range resources.Items {
		rp := BackupToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListFilestoreBetaBackupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigBackup(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
