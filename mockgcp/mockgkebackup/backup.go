// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:mockgcp-support
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.Backup

package mockgkebackup

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/gkebackup/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *BackupForGKEV1) GetBackup(ctx context.Context, req *pb.GetBackupRequest) (*pb.Backup, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Backup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Backup %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupForGKEV1) CreateBackup(ctx context.Context, req *pb.CreateBackupRequest) (*longrunningpb.Operation, error) {
	// Parse the parent BackupPlan name
	parentName, err := s.parseBackupPlanName(req.GetParent())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent name: %v", err)
	}

	// Construct the full Backup name
	reqName := fmt.Sprintf("%s/backups/%s", req.GetParent(), req.GetBackupId())
	name, err := s.parseBackupName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	// Get the parent BackupPlan to inherit properties
	parentPlan := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, parentName.String(), parentPlan); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "BackupPlan %q not found", parentName.String())
		}
		return nil, status.Errorf(codes.Internal, "error getting parent BackupPlan: %v", err)
	}

	now := time.Now()

	obj := proto.Clone(req.GetBackup()).(*pb.Backup)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = name.BackupID // Consider generating a real UUID if needed
	obj.Manual = true       // Created via API call
	obj.State = pb.Backup_CREATING
	obj.StateReason = "BackupJob is being pushed to target cluster."

	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime:            timestamppb.New(now),
		Target:                name.String(),
		Verb:                  "create",
		ApiVersion:            "v1",
		RequestedCancellation: false,
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(now)

		setBackupDefaultValuesAndInherit(obj, parentPlan)

		obj.State = pb.Backup_IN_PROGRESS
		obj.StateReason = "Starting to back up Kubernetes resources"
		obj.UpdateTime = timestamppb.New(now)
		obj.DeleteLockExpireTime = timestamppb.New(now.Add(time.Duration(obj.DeleteLockDays) * 24 * time.Hour))

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *BackupForGKEV1) UpdateBackup(ctx context.Context, req *pb.UpdateBackupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupName(req.GetBackup().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Backup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetBackup().GetDescription()
		case "labels":
			obj.Labels = req.GetBackup().GetLabels()
		case "retain_days":
			// Validation: retain_days can only be increased and must be >= delete_lock_days
			if req.GetBackup().GetRetainDays() < obj.RetainDays {
				return nil, status.Errorf(codes.InvalidArgument, "retain_days cannot be decreased")
			}
			if req.GetBackup().GetRetainDays() > 0 && req.GetBackup().GetRetainDays() < obj.DeleteLockDays {
				return nil, status.Errorf(codes.InvalidArgument, "retain_days must be >= delete_lock_days (%d)", obj.DeleteLockDays)
			}
			obj.RetainDays = req.GetBackup().GetRetainDays()
			if obj.RetainDays > 0 {
				expireTime := obj.CreateTime.AsTime().Add(time.Duration(obj.RetainDays) * 24 * time.Hour)
				obj.RetainExpireTime = timestamppb.New(expireTime)
			} else {
				obj.RetainExpireTime = nil // No automatic deletion
			}
		case "delete_lock_days":
			// Validation: delete_lock_days can only be increased
			if req.GetBackup().GetDeleteLockDays() < obj.DeleteLockDays {
				return nil, status.Errorf(codes.InvalidArgument, "delete_lock_days cannot be decreased")
			}
			// Update retain_days if it becomes invalid
			if obj.RetainDays > 0 && obj.RetainDays < req.GetBackup().GetDeleteLockDays() {
				return nil, status.Errorf(codes.InvalidArgument, "retain_days (%d) must be >= delete_lock_days (%d)", obj.RetainDays, req.GetBackup().GetDeleteLockDays())
			}
			obj.DeleteLockDays = req.GetBackup().GetDeleteLockDays()
			if obj.DeleteLockDays > 0 {
				expireTime := obj.CreateTime.AsTime().Add(time.Duration(obj.DeleteLockDays) * 24 * time.Hour)
				obj.DeleteLockExpireTime = timestamppb.New(expireTime)
			} else {
				obj.DeleteLockExpireTime = nil // No lock
			}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid for Backup update", path)
		}
	}
	obj.UpdateTime = timestamppb.New(now)
	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "update",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		// Refresh the object from storage in case the LRO implementation allows concurrent updates
		updatedObj := &pb.Backup{}
		if err := s.storage.Get(ctx, fqn, updatedObj); err != nil {
			return nil, err
		}
		return updatedObj, nil
	})
}

func (s *BackupForGKEV1) DeleteBackup(ctx context.Context, req *pb.DeleteBackupRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deletedObj := &pb.Backup{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Backup %q not found", fqn)
		}
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
		ApiVersion: "v1",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type backupName struct {
	Project      *projects.ProjectData
	Location     string
	BackupPlanID string
	BackupID     string
}

func (n *backupName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s/backups/%s", n.Project.ID, n.Location, n.BackupPlanID, n.BackupID)
}

// parseBackupName parses a string into a backupName.
// The expected form is `projects/*/locations/*/backupPlans/*/backups/*`.
func (s *MockService) parseBackupName(name string) (*backupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupPlans" && tokens[6] == "backups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupName{
			Project:      project,
			Location:     tokens[3],
			BackupPlanID: tokens[5],
			BackupID:     tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid format for a backup", name)
}

func setBackupDefaultValuesAndInherit(obj *pb.Backup, parentPlan *pb.BackupPlan) {
	obj.ClusterMetadata = &pb.Backup_ClusterMetadata{
		Cluster: parentPlan.Cluster,
		PlatformVersion: &pb.Backup_ClusterMetadata_GkeVersion{
			GkeVersion: "v1.31.6-gke.1020000",
		},
		K8SVersion: "1.31",
		BackupCrdVersions: map[string]string{
			"backupjobs.gkebackup.gke.io":                 "v1",
			"protectedapplicationgroups.gkebackup.gke.io": "v1",
			"protectedapplications.gkebackup.gke.io":      "v1",
			"restorejobs.gkebackup.gke.io":                "v1",
		},
	}

	if parentPlan.BackupConfig != nil {
		switch scope := parentPlan.BackupConfig.BackupScope.(type) {
		case *pb.BackupPlan_BackupConfig_AllNamespaces:
			obj.BackupScope = &pb.Backup_AllNamespaces{AllNamespaces: scope.AllNamespaces}
		case *pb.BackupPlan_BackupConfig_SelectedNamespaces:
			obj.BackupScope = &pb.Backup_SelectedNamespaces{SelectedNamespaces: scope.SelectedNamespaces}
		case *pb.BackupPlan_BackupConfig_SelectedApplications:
			obj.BackupScope = &pb.Backup_SelectedApplications{SelectedApplications: scope.SelectedApplications}
		}
		obj.ContainsVolumeData = parentPlan.BackupConfig.IncludeVolumeData
		obj.ContainsSecrets = parentPlan.BackupConfig.IncludeSecrets
	}
}
