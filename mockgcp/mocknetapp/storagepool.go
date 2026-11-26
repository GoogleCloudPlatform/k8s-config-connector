// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:mockgcp-support
// proto.service: google.cloud.netapp.v1.NetApp
// proto.message: google.cloud.netapp.v1.StoragePool

package mocknetapp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func (s *backupVaultsService) GetStoragePool(ctx context.Context, req *pb.GetStoragePoolRequest) (*pb.StoragePool, error) {
	name, err := s.parseStoragePoolName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.StoragePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *backupVaultsService) ListStoragePools(ctx context.Context, req *pb.ListStoragePoolsRequest) (*pb.ListStoragePoolsResponse, error) {
	parent := req.GetParent()
	response := &pb.ListStoragePoolsResponse{}

	findPrefix := fmt.Sprintf("%s/storagePools/", parent)

	storagePoolKind := (&pb.StoragePool{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, storagePoolKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		cloned := proto.Clone(obj).(*pb.StoragePool)
		response.StoragePools = append(response.StoragePools, cloned)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *backupVaultsService) CreateStoragePool(ctx context.Context, req *pb.CreateStoragePoolRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetParent() + "/storagePools/" + req.GetStoragePoolId()
	name, err := s.parseStoragePoolName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetStoragePool()).(*pb.StoragePool)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	s.populateDefaultsForStoragePool(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                fqn,
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupVaultsService) UpdateStoragePool(ctx context.Context, req *pb.UpdateStoragePoolRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseStoragePoolName(req.GetStoragePool().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.StoragePool{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	if err := fields.UpdateByFieldMask(obj, req.GetStoragePool(), paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	s.populateDefaultsForStoragePool(obj, name)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                fqn,
		Verb:                  "update",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *backupVaultsService) DeleteStoragePool(ctx context.Context, req *pb.DeleteStoragePoolRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseStoragePoolName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.StoragePool{}
	err = s.storage.Delete(ctx, fqn, existing)
	if err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.Now(),
		Target:                fqn,
		Verb:                  "delete",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

func (s *backupVaultsService) populateDefaultsForStoragePool(obj *pb.StoragePool, name *storagePoolName) {
	if obj.State == pb.StoragePool_STATE_UNSPECIFIED {
		obj.State = pb.StoragePool_READY
	}
	if obj.EncryptionType == pb.EncryptionType_ENCRYPTION_TYPE_UNSPECIFIED {
		obj.EncryptionType = pb.EncryptionType_SERVICE_MANAGED
	}
	if obj.QosType == pb.QosType_QOS_TYPE_UNSPECIFIED {
		obj.QosType = pb.QosType_AUTO
	}
	if obj.Type == nil {
		typeVal := pb.StoragePoolType_FILE
		obj.Type = &typeVal
	}
	if obj.Mode == nil {
		modeVal := pb.Mode_MODE_UNSPECIFIED
		obj.Mode = &modeVal
	}
	// Zone and ReplicaZone are write-only / create-only parameters, not returned in GET/LRO responses
	obj.Zone = ""
	obj.ReplicaZone = ""

	// Convert network project ID to project number
	if obj.Network != "" {
		parts := strings.Split(obj.Network, "/")
		if len(parts) == 5 && parts[0] == "projects" && parts[2] == "global" && parts[3] == "networks" {
			proj, err := s.Projects.GetProjectByIDOrNumber(parts[1])
			if err == nil {
				obj.Network = fmt.Sprintf("projects/%d/global/networks/%s", proj.Number, parts[4])
			}
		}
	}

	// Calculate throughput if not custom performance
	if !obj.CustomPerformanceEnabled && obj.CapacityGib > 0 {
		var throughput int64
		switch obj.ServiceLevel {
		case pb.ServiceLevel_PREMIUM:
			throughput = obj.CapacityGib / 16
		case pb.ServiceLevel_EXTREME:
			throughput = obj.CapacityGib / 8
		default:
			// STANDARD: 16 MiB/s per TiB (1024 GiB)
			throughput = obj.CapacityGib / 64
		}
		obj.TotalThroughputMibps = throughput
		obj.AvailableThroughputMibps = float64(throughput)
	}
}

type storagePoolName struct {
	Project       *projects.ProjectData
	Location      string
	StoragePoolId string
}

func (n *storagePoolName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", n.Project.ID, n.Location, n.StoragePoolId)
}

func (s *backupVaultsService) parseStoragePoolName(name string) (*storagePoolName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "storagePools" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &storagePoolName{
			Project:       project,
			Location:      tokens[3],
			StoragePoolId: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
