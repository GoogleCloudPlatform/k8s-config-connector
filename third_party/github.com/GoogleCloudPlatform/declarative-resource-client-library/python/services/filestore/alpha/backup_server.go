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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/filestore/alpha/filestore_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/filestore/alpha"
)

// BackupServer implements the gRPC interface for Backup.
type BackupServer struct{}

// ProtoToBackupStateEnum converts a BackupStateEnum enum from its proto representation.
func ProtoToFilestoreAlphaBackupStateEnum(e alphapb.FilestoreAlphaBackupStateEnum) *alpha.BackupStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaBackupStateEnum_name[int32(e)]; ok {
		e := alpha.BackupStateEnum(n[len("FilestoreAlphaBackupStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackupSourceInstanceTierEnum converts a BackupSourceInstanceTierEnum enum from its proto representation.
func ProtoToFilestoreAlphaBackupSourceInstanceTierEnum(e alphapb.FilestoreAlphaBackupSourceInstanceTierEnum) *alpha.BackupSourceInstanceTierEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.FilestoreAlphaBackupSourceInstanceTierEnum_name[int32(e)]; ok {
		e := alpha.BackupSourceInstanceTierEnum(n[len("FilestoreAlphaBackupSourceInstanceTierEnum"):])
		return &e
	}
	return nil
}

// ProtoToBackup converts a Backup resource from its proto representation.
func ProtoToBackup(p *alphapb.FilestoreAlphaBackup) *alpha.Backup {
	obj := &alpha.Backup{
		Name:               dcl.StringOrNil(p.GetName()),
		Description:        dcl.StringOrNil(p.GetDescription()),
		State:              ProtoToFilestoreAlphaBackupStateEnum(p.GetState()),
		CreateTime:         dcl.StringOrNil(p.GetCreateTime()),
		CapacityGb:         dcl.Int64OrNil(p.GetCapacityGb()),
		StorageBytes:       dcl.Int64OrNil(p.GetStorageBytes()),
		SourceInstance:     dcl.StringOrNil(p.GetSourceInstance()),
		SourceFileShare:    dcl.StringOrNil(p.GetSourceFileShare()),
		SourceInstanceTier: ProtoToFilestoreAlphaBackupSourceInstanceTierEnum(p.GetSourceInstanceTier()),
		DownloadBytes:      dcl.Int64OrNil(p.GetDownloadBytes()),
		Project:            dcl.StringOrNil(p.GetProject()),
		Location:           dcl.StringOrNil(p.GetLocation()),
	}
	return obj
}

// BackupStateEnumToProto converts a BackupStateEnum enum to its proto representation.
func FilestoreAlphaBackupStateEnumToProto(e *alpha.BackupStateEnum) alphapb.FilestoreAlphaBackupStateEnum {
	if e == nil {
		return alphapb.FilestoreAlphaBackupStateEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaBackupStateEnum_value["BackupStateEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaBackupStateEnum(v)
	}
	return alphapb.FilestoreAlphaBackupStateEnum(0)
}

// BackupSourceInstanceTierEnumToProto converts a BackupSourceInstanceTierEnum enum to its proto representation.
func FilestoreAlphaBackupSourceInstanceTierEnumToProto(e *alpha.BackupSourceInstanceTierEnum) alphapb.FilestoreAlphaBackupSourceInstanceTierEnum {
	if e == nil {
		return alphapb.FilestoreAlphaBackupSourceInstanceTierEnum(0)
	}
	if v, ok := alphapb.FilestoreAlphaBackupSourceInstanceTierEnum_value["BackupSourceInstanceTierEnum"+string(*e)]; ok {
		return alphapb.FilestoreAlphaBackupSourceInstanceTierEnum(v)
	}
	return alphapb.FilestoreAlphaBackupSourceInstanceTierEnum(0)
}

// BackupToProto converts a Backup resource to its proto representation.
func BackupToProto(resource *alpha.Backup) *alphapb.FilestoreAlphaBackup {
	p := &alphapb.FilestoreAlphaBackup{}
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetState(FilestoreAlphaBackupStateEnumToProto(resource.State))
	p.SetCreateTime(dcl.ValueOrEmptyString(resource.CreateTime))
	p.SetCapacityGb(dcl.ValueOrEmptyInt64(resource.CapacityGb))
	p.SetStorageBytes(dcl.ValueOrEmptyInt64(resource.StorageBytes))
	p.SetSourceInstance(dcl.ValueOrEmptyString(resource.SourceInstance))
	p.SetSourceFileShare(dcl.ValueOrEmptyString(resource.SourceFileShare))
	p.SetSourceInstanceTier(FilestoreAlphaBackupSourceInstanceTierEnumToProto(resource.SourceInstanceTier))
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
func (s *BackupServer) applyBackup(ctx context.Context, c *alpha.Client, request *alphapb.ApplyFilestoreAlphaBackupRequest) (*alphapb.FilestoreAlphaBackup, error) {
	p := ProtoToBackup(request.GetResource())
	res, err := c.ApplyBackup(ctx, p)
	if err != nil {
		return nil, err
	}
	r := BackupToProto(res)
	return r, nil
}

// applyFilestoreAlphaBackup handles the gRPC request by passing it to the underlying Backup Apply() method.
func (s *BackupServer) ApplyFilestoreAlphaBackup(ctx context.Context, request *alphapb.ApplyFilestoreAlphaBackupRequest) (*alphapb.FilestoreAlphaBackup, error) {
	cl, err := createConfigBackup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyBackup(ctx, cl, request)
}

// DeleteBackup handles the gRPC request by passing it to the underlying Backup Delete() method.
func (s *BackupServer) DeleteFilestoreAlphaBackup(ctx context.Context, request *alphapb.DeleteFilestoreAlphaBackupRequest) (*emptypb.Empty, error) {

	cl, err := createConfigBackup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteBackup(ctx, ProtoToBackup(request.GetResource()))

}

// ListFilestoreAlphaBackup handles the gRPC request by passing it to the underlying BackupList() method.
func (s *BackupServer) ListFilestoreAlphaBackup(ctx context.Context, request *alphapb.ListFilestoreAlphaBackupRequest) (*alphapb.ListFilestoreAlphaBackupResponse, error) {
	cl, err := createConfigBackup(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListBackup(ctx, request.GetProject(), request.GetLocation())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.FilestoreAlphaBackup
	for _, r := range resources.Items {
		rp := BackupToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListFilestoreAlphaBackupResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigBackup(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
