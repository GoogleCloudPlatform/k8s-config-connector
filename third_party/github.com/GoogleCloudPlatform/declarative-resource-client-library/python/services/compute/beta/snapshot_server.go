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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for Snapshot.
type SnapshotServer struct{}

// ProtoToSnapshotSnapshotEncryptionKey converts a SnapshotSnapshotEncryptionKey resource from its proto representation.
func ProtoToComputeBetaSnapshotSnapshotEncryptionKey(p *betapb.ComputeBetaSnapshotSnapshotEncryptionKey) *beta.SnapshotSnapshotEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.SnapshotSnapshotEncryptionKey{
		RawKey: dcl.StringOrNil(p.RawKey),
		Sha256: dcl.StringOrNil(p.Sha256),
	}
	return obj
}

// ProtoToSnapshotSourceDiskEncryptionKey converts a SnapshotSourceDiskEncryptionKey resource from its proto representation.
func ProtoToComputeBetaSnapshotSourceDiskEncryptionKey(p *betapb.ComputeBetaSnapshotSourceDiskEncryptionKey) *beta.SnapshotSourceDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.SnapshotSourceDiskEncryptionKey{
		RawKey: dcl.StringOrNil(p.RawKey),
	}
	return obj
}

// ProtoToSnapshot converts a Snapshot resource from its proto representation.
func ProtoToSnapshot(p *betapb.ComputeBetaSnapshot) *beta.Snapshot {
	obj := &beta.Snapshot{
		Name:                    dcl.StringOrNil(p.Name),
		Description:             dcl.StringOrNil(p.Description),
		SourceDisk:              dcl.StringOrNil(p.SourceDisk),
		DiskSizeGb:              dcl.Int64OrNil(p.DiskSizeGb),
		StorageBytes:            dcl.Int64OrNil(p.StorageBytes),
		SnapshotEncryptionKey:   ProtoToComputeBetaSnapshotSnapshotEncryptionKey(p.GetSnapshotEncryptionKey()),
		SourceDiskEncryptionKey: ProtoToComputeBetaSnapshotSourceDiskEncryptionKey(p.GetSourceDiskEncryptionKey()),
		SelfLink:                dcl.StringOrNil(p.SelfLink),
		Project:                 dcl.StringOrNil(p.Project),
		Zone:                    dcl.StringOrNil(p.Zone),
		Id:                      dcl.Int64OrNil(p.Id),
	}
	for _, r := range p.GetLicense() {
		obj.License = append(obj.License, r)
	}
	return obj
}

// SnapshotSnapshotEncryptionKeyToProto converts a SnapshotSnapshotEncryptionKey resource to its proto representation.
func ComputeBetaSnapshotSnapshotEncryptionKeyToProto(o *beta.SnapshotSnapshotEncryptionKey) *betapb.ComputeBetaSnapshotSnapshotEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaSnapshotSnapshotEncryptionKey{
		RawKey: dcl.ValueOrEmptyString(o.RawKey),
		Sha256: dcl.ValueOrEmptyString(o.Sha256),
	}
	return p
}

// SnapshotSourceDiskEncryptionKeyToProto converts a SnapshotSourceDiskEncryptionKey resource to its proto representation.
func ComputeBetaSnapshotSourceDiskEncryptionKeyToProto(o *beta.SnapshotSourceDiskEncryptionKey) *betapb.ComputeBetaSnapshotSourceDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaSnapshotSourceDiskEncryptionKey{
		RawKey: dcl.ValueOrEmptyString(o.RawKey),
	}
	return p
}

// SnapshotToProto converts a Snapshot resource to its proto representation.
func SnapshotToProto(resource *beta.Snapshot) *betapb.ComputeBetaSnapshot {
	p := &betapb.ComputeBetaSnapshot{
		Name:                    dcl.ValueOrEmptyString(resource.Name),
		Description:             dcl.ValueOrEmptyString(resource.Description),
		SourceDisk:              dcl.ValueOrEmptyString(resource.SourceDisk),
		DiskSizeGb:              dcl.ValueOrEmptyInt64(resource.DiskSizeGb),
		StorageBytes:            dcl.ValueOrEmptyInt64(resource.StorageBytes),
		SnapshotEncryptionKey:   ComputeBetaSnapshotSnapshotEncryptionKeyToProto(resource.SnapshotEncryptionKey),
		SourceDiskEncryptionKey: ComputeBetaSnapshotSourceDiskEncryptionKeyToProto(resource.SourceDiskEncryptionKey),
		SelfLink:                dcl.ValueOrEmptyString(resource.SelfLink),
		Project:                 dcl.ValueOrEmptyString(resource.Project),
		Zone:                    dcl.ValueOrEmptyString(resource.Zone),
		Id:                      dcl.ValueOrEmptyInt64(resource.Id),
	}
	for _, r := range resource.License {
		p.License = append(p.License, r)
	}

	return p
}

// ApplySnapshot handles the gRPC request by passing it to the underlying Snapshot Apply() method.
func (s *SnapshotServer) applySnapshot(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaSnapshotRequest) (*betapb.ComputeBetaSnapshot, error) {
	p := ProtoToSnapshot(request.GetResource())
	res, err := c.ApplySnapshot(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SnapshotToProto(res)
	return r, nil
}

// ApplySnapshot handles the gRPC request by passing it to the underlying Snapshot Apply() method.
func (s *SnapshotServer) ApplyComputeBetaSnapshot(ctx context.Context, request *betapb.ApplyComputeBetaSnapshotRequest) (*betapb.ComputeBetaSnapshot, error) {
	cl, err := createConfigSnapshot(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySnapshot(ctx, cl, request)
}

// DeleteSnapshot handles the gRPC request by passing it to the underlying Snapshot Delete() method.
func (s *SnapshotServer) DeleteComputeBetaSnapshot(ctx context.Context, request *betapb.DeleteComputeBetaSnapshotRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSnapshot(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSnapshot(ctx, ProtoToSnapshot(request.GetResource()))

}

// ListComputeBetaSnapshot handles the gRPC request by passing it to the underlying SnapshotList() method.
func (s *SnapshotServer) ListComputeBetaSnapshot(ctx context.Context, request *betapb.ListComputeBetaSnapshotRequest) (*betapb.ListComputeBetaSnapshotResponse, error) {
	cl, err := createConfigSnapshot(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSnapshot(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaSnapshot
	for _, r := range resources.Items {
		rp := SnapshotToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaSnapshotResponse{Items: protos}, nil
}

func createConfigSnapshot(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
