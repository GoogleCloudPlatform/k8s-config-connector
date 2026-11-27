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
// proto.message: google.cloud.netapp.v1.Volume

package mocknetapp

import (
	"context"
	"fmt"
	"math"
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

func (s *backupVaultsService) GetVolume(ctx context.Context, req *pb.GetVolumeRequest) (*pb.Volume, error) {
	name, err := s.parseVolumeName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Volume{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	cloned := proto.Clone(obj).(*pb.Volume)
	if cloned.SnapshotPolicy != nil {
		if cloned.SnapshotPolicy.HourlySchedule != nil && cloned.SnapshotPolicy.HourlySchedule.GetSnapshotsToKeep() == 0 {
			cloned.SnapshotPolicy.HourlySchedule = nil
		}
		if cloned.SnapshotPolicy.MonthlySchedule != nil && cloned.SnapshotPolicy.MonthlySchedule.GetSnapshotsToKeep() == 0 {
			cloned.SnapshotPolicy.MonthlySchedule = nil
		}
		if cloned.SnapshotPolicy.WeeklySchedule != nil && cloned.SnapshotPolicy.WeeklySchedule.GetSnapshotsToKeep() == 0 {
			cloned.SnapshotPolicy.WeeklySchedule = nil
		}
	}

	return cloned, nil
}

func (s *backupVaultsService) ListVolumes(ctx context.Context, req *pb.ListVolumesRequest) (*pb.ListVolumesResponse, error) {
	parent := req.GetParent()
	response := &pb.ListVolumesResponse{}

	findPrefix := fmt.Sprintf("%s/volumes/", parent)

	volumeKind := (&pb.Volume{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, volumeKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		cloned := proto.Clone(obj).(*pb.Volume)
		response.Volumes = append(response.Volumes, cloned)
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *backupVaultsService) CreateVolume(ctx context.Context, req *pb.CreateVolumeRequest) (*longrunningpb.Operation, error) {
	reqName := req.GetParent() + "/volumes/" + req.GetVolumeId()
	name, err := s.parseVolumeName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.Clone(req.GetVolume()).(*pb.Volume)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)

	s.populateDefaultsForVolume(ctx, obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroObj := proto.Clone(obj).(*pb.Volume)
	s.populateLRODefaultsForVolume(lroObj)

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
		return lroObj, nil
	})
}

func (s *backupVaultsService) UpdateVolume(ctx context.Context, req *pb.UpdateVolumeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVolumeName(req.GetVolume().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Volume{}
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

	if err := fields.UpdateByFieldMask(obj, req.GetVolume(), paths); err != nil {
		return nil, fmt.Errorf("update field_mask.paths: %w", err)
	}

	s.populateDefaultsForVolume(ctx, obj, name)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroObj := proto.Clone(obj).(*pb.Volume)
	s.populateLRODefaultsForVolume(lroObj)

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
		return lroObj, nil
	})
}

func (s *backupVaultsService) DeleteVolume(ctx context.Context, req *pb.DeleteVolumeRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseVolumeName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.Volume{}
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

func (s *backupVaultsService) populateDefaultsForVolume(ctx context.Context, obj *pb.Volume, name *volumeName) {
	if obj.State == pb.Volume_STATE_UNSPECIFIED {
		obj.State = pb.Volume_READY
	}
	if obj.StateDetails == "" {
		obj.StateDetails = "Available for use"
	}
	if obj.EncryptionType == pb.EncryptionType_ENCRYPTION_TYPE_UNSPECIFIED {
		obj.EncryptionType = pb.EncryptionType_SERVICE_MANAGED
	}
	if obj.SecurityStyle == pb.SecurityStyle_SECURITY_STYLE_UNSPECIFIED {
		obj.SecurityStyle = pb.SecurityStyle_UNIX
	}
	if obj.UnixPermissions == "" {
		obj.UnixPermissions = "0770"
	}

	// Format storagePool as short ID
	storagePoolID := obj.StoragePool
	if tokens := strings.Split(obj.StoragePool, "/"); len(tokens) > 0 {
		storagePoolID = tokens[len(tokens)-1]
		obj.StoragePool = storagePoolID
	}

	// Try to inherit network and serviceLevel from StoragePool
	if storagePoolID != "" {
		poolFQN := fmt.Sprintf("projects/%s/locations/%s/storagePools/%s", name.Project.ID, name.Location, storagePoolID)
		pool := &pb.StoragePool{}
		if err := s.storage.Get(ctx, poolFQN, pool); err == nil {
			if obj.ServiceLevel == pb.ServiceLevel_SERVICE_LEVEL_UNSPECIFIED {
				obj.ServiceLevel = pool.ServiceLevel
			}
			if obj.Network == "" {
				obj.Network = pool.Network
			}
		}
	}

	if obj.ServiceLevel == pb.ServiceLevel_SERVICE_LEVEL_UNSPECIFIED {
		obj.ServiceLevel = pb.ServiceLevel_STANDARD
	}
	if obj.Network == "" {
		obj.Network = fmt.Sprintf("projects/%d/global/networks/default", name.Project.Number)
	}

	// Calculate throughput if capacity is set
	if obj.CapacityGib > 0 {
		var raw float64
		switch obj.ServiceLevel {
		case pb.ServiceLevel_PREMIUM:
			raw = float64(obj.CapacityGib) / 16.0
		case pb.ServiceLevel_EXTREME:
			raw = float64(obj.CapacityGib) / 8.0
		default:
			// STANDARD: 16 MiB/s per TiB (1024 GiB)
			raw = float64(obj.CapacityGib) / 64.0
		}
		obj.ThroughputMibps = math.RoundToEven(raw*1000.0) / 1000.0
	}

	// Populate mountOptions
	if len(obj.MountOptions) == 0 && obj.ShareName != "" {
		ip := "192.168.22.36"
		instructions := fmt.Sprintf("Setting up your instance\nOpen an SSH client and connect to your instance.\nInstall the nfs client on your instance.\nOn Red Hat Enterprise Linux or SuSE Linux instance:\nsudo yum install -y nfs-utils\nOn an Ubuntu or Debian instance:\nsudo apt-get install nfs-common\n\nMounting your volume\nCreate a new directory on your instance, such as \"/%s\":\nsudo mkdir /%s\nMount your volume using the example command below:\nsudo mount -t nfs -o rw,hard,rsize=65536,wsize=65536,vers=3,tcp %s:/%s /%s\nNote. Please use mount options appropriate for your specific workloads when known.",
			obj.ShareName, obj.ShareName, ip, obj.ShareName, obj.ShareName)
		obj.MountOptions = []*pb.MountOption{
			{
				Export:       fmt.Sprintf("/%s", obj.ShareName),
				ExportFull:   fmt.Sprintf("%s:/%s", ip, obj.ShareName),
				Instructions: instructions,
				IpAddress:    ip,
				Protocol:     pb.Protocols_NFSV3,
			},
		}
	}

	// ExportPolicy defaults
	if obj.ExportPolicy == nil {
		obj.ExportPolicy = &pb.ExportPolicy{}
	} else {
		for _, rule := range obj.ExportPolicy.Rules {
			if rule.HasRootAccess == nil {
				rule.HasRootAccess = proto.String("true")
			}
			if rule.SquashMode == nil {
				mode := pb.SimpleExportPolicyRule_NO_ROOT_SQUASH
				rule.SquashMode = &mode
			}
			if rule.Kerberos_5ReadOnly == nil {
				rule.Kerberos_5ReadOnly = proto.Bool(false)
			}
			if rule.Kerberos_5ReadWrite == nil {
				rule.Kerberos_5ReadWrite = proto.Bool(false)
			}
			if rule.Kerberos_5IReadOnly == nil {
				rule.Kerberos_5IReadOnly = proto.Bool(false)
			}
			if rule.Kerberos_5IReadWrite == nil {
				rule.Kerberos_5IReadWrite = proto.Bool(false)
			}
			if rule.Kerberos_5PReadOnly == nil {
				rule.Kerberos_5PReadOnly = proto.Bool(false)
			}
			if rule.Kerberos_5PReadWrite == nil {
				rule.Kerberos_5PReadWrite = proto.Bool(false)
			}
		}
	}
}

func (s *backupVaultsService) populateLRODefaultsForVolume(obj *pb.Volume) {
	if obj.SnapshotPolicy != nil && obj.SnapshotPolicy.GetEnabled() {
		if obj.SnapshotPolicy.HourlySchedule == nil {
			obj.SnapshotPolicy.HourlySchedule = &pb.HourlySchedule{
				Minute:          proto.Float64(0),
				SnapshotsToKeep: proto.Float64(0),
			}
		}
		if obj.SnapshotPolicy.MonthlySchedule == nil {
			obj.SnapshotPolicy.MonthlySchedule = &pb.MonthlySchedule{
				DaysOfMonth:     proto.String(""),
				Hour:            proto.Float64(0),
				Minute:          proto.Float64(0),
				SnapshotsToKeep: proto.Float64(0),
			}
		}
		if obj.SnapshotPolicy.WeeklySchedule == nil {
			obj.SnapshotPolicy.WeeklySchedule = &pb.WeeklySchedule{
				Day:             proto.String(""),
				Hour:            proto.Float64(0),
				Minute:          proto.Float64(0),
				SnapshotsToKeep: proto.Float64(0),
			}
		}
	}
}

type volumeName struct {
	Project  *projects.ProjectData
	Location string
	VolumeId string
}

func (n *volumeName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/volumes/%s", n.Project.ID, n.Location, n.VolumeId)
}

func (s *backupVaultsService) parseVolumeName(name string) (*volumeName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "volumes" {
		project, err := s.Projects.GetProjectByIDOrNumber(tokens[1])
		if err != nil {
			return nil, err
		}

		return &volumeName{
			Project:  project,
			Location: tokens[3],
			VolumeId: tokens[5],
		}, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
