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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for Snapshot.
type SnapshotServer struct{}

// ProtoToSnapshotSnapshotEncryptionKey converts a SnapshotSnapshotEncryptionKey resource from its proto representation.
func ProtoToComputeSnapshotSnapshotEncryptionKey(p *computepb.ComputeSnapshotSnapshotEncryptionKey) *compute.SnapshotSnapshotEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.SnapshotSnapshotEncryptionKey{
		RawKey: dcl.StringOrNil(p.RawKey),
		Sha256: dcl.StringOrNil(p.Sha256),
	}
	return obj
}

// ProtoToSnapshotSourceDiskEncryptionKey converts a SnapshotSourceDiskEncryptionKey resource from its proto representation.
func ProtoToComputeSnapshotSourceDiskEncryptionKey(p *computepb.ComputeSnapshotSourceDiskEncryptionKey) *compute.SnapshotSourceDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.SnapshotSourceDiskEncryptionKey{
		RawKey: dcl.StringOrNil(p.RawKey),
	}
	return obj
}

// ProtoToSnapshot converts a Snapshot resource from its proto representation.
func ProtoToSnapshot(p *computepb.ComputeSnapshot) *compute.Snapshot {
	obj := &compute.Snapshot{
		Name:                    dcl.StringOrNil(p.Name),
		Description:             dcl.StringOrNil(p.Description),
		SourceDisk:              dcl.StringOrNil(p.SourceDisk),
		DiskSizeGb:              dcl.Int64OrNil(p.DiskSizeGb),
		StorageBytes:            dcl.Int64OrNil(p.StorageBytes),
		SnapshotEncryptionKey:   ProtoToComputeSnapshotSnapshotEncryptionKey(p.GetSnapshotEncryptionKey()),
		SourceDiskEncryptionKey: ProtoToComputeSnapshotSourceDiskEncryptionKey(p.GetSourceDiskEncryptionKey()),
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
func ComputeSnapshotSnapshotEncryptionKeyToProto(o *compute.SnapshotSnapshotEncryptionKey) *computepb.ComputeSnapshotSnapshotEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeSnapshotSnapshotEncryptionKey{
		RawKey: dcl.ValueOrEmptyString(o.RawKey),
		Sha256: dcl.ValueOrEmptyString(o.Sha256),
	}
	return p
}

// SnapshotSourceDiskEncryptionKeyToProto converts a SnapshotSourceDiskEncryptionKey resource to its proto representation.
func ComputeSnapshotSourceDiskEncryptionKeyToProto(o *compute.SnapshotSourceDiskEncryptionKey) *computepb.ComputeSnapshotSourceDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeSnapshotSourceDiskEncryptionKey{
		RawKey: dcl.ValueOrEmptyString(o.RawKey),
	}
	return p
}

// SnapshotToProto converts a Snapshot resource to its proto representation.
func SnapshotToProto(resource *compute.Snapshot) *computepb.ComputeSnapshot {
	p := &computepb.ComputeSnapshot{
		Name:                    dcl.ValueOrEmptyString(resource.Name),
		Description:             dcl.ValueOrEmptyString(resource.Description),
		SourceDisk:              dcl.ValueOrEmptyString(resource.SourceDisk),
		DiskSizeGb:              dcl.ValueOrEmptyInt64(resource.DiskSizeGb),
		StorageBytes:            dcl.ValueOrEmptyInt64(resource.StorageBytes),
		SnapshotEncryptionKey:   ComputeSnapshotSnapshotEncryptionKeyToProto(resource.SnapshotEncryptionKey),
		SourceDiskEncryptionKey: ComputeSnapshotSourceDiskEncryptionKeyToProto(resource.SourceDiskEncryptionKey),
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
func (s *SnapshotServer) applySnapshot(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeSnapshotRequest) (*computepb.ComputeSnapshot, error) {
	p := ProtoToSnapshot(request.GetResource())
	res, err := c.ApplySnapshot(ctx, p)
	if err != nil {
		return nil, err
	}
	r := SnapshotToProto(res)
	return r, nil
}

// ApplySnapshot handles the gRPC request by passing it to the underlying Snapshot Apply() method.
func (s *SnapshotServer) ApplyComputeSnapshot(ctx context.Context, request *computepb.ApplyComputeSnapshotRequest) (*computepb.ComputeSnapshot, error) {
	cl, err := createConfigSnapshot(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applySnapshot(ctx, cl, request)
}

// DeleteSnapshot handles the gRPC request by passing it to the underlying Snapshot Delete() method.
func (s *SnapshotServer) DeleteComputeSnapshot(ctx context.Context, request *computepb.DeleteComputeSnapshotRequest) (*emptypb.Empty, error) {

	cl, err := createConfigSnapshot(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteSnapshot(ctx, ProtoToSnapshot(request.GetResource()))

}

// ListComputeSnapshot handles the gRPC request by passing it to the underlying SnapshotList() method.
func (s *SnapshotServer) ListComputeSnapshot(ctx context.Context, request *computepb.ListComputeSnapshotRequest) (*computepb.ListComputeSnapshotResponse, error) {
	cl, err := createConfigSnapshot(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListSnapshot(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeSnapshot
	for _, r := range resources.Items {
		rp := SnapshotToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeSnapshotResponse{Items: protos}, nil
}

func createConfigSnapshot(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
