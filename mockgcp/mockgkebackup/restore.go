// Copyright 2024 Google LLC
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
// proto.service: google.cloud.gkebackup.v1.BackupForGKE
// proto.message: google.cloud.gkebackup.v1.Restore

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

func (s *BackupForGKEV1) GetRestore(ctx context.Context, req *pb.GetRestoreRequest) (*pb.Restore, error) {
	name, err := s.parseRestoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Restore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Restore %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupForGKEV1) CreateRestore(ctx context.Context, req *pb.CreateRestoreRequest) (*longrunningpb.Operation, error) {
	// Parse the parent RestorePlan name
	parentName, err := s.parseRestorePlanName(req.GetParent())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parent name: %v", err)
	}

	// Construct the full Restore name
	reqName := fmt.Sprintf("%s/restores/%s", req.GetParent(), req.GetRestoreId())
	name, err := s.parseRestoreName(reqName)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	// Get the parent RestorePlan to inherit properties
	parentPlan := &pb.RestorePlan{}
	if err := s.storage.Get(ctx, parentName.String(), parentPlan); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "RestorePlan %q not found", parentName.String())
		}
		return nil, status.Errorf(codes.Internal, "error getting parent RestorePlan: %v", err)
	}

	// Basic validation: Check if the referenced backup exists
	// Note: Full validation would also check if the backup belongs to the plan's backup_plan
	backupName, err := s.parseBackupName(req.GetRestore().GetBackup())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid backup name: %v", err)
	}
	backup := &pb.Backup{}
	if err := s.storage.Get(ctx, backupName.String(), backup); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.FailedPrecondition, "Backup %q not found", backupName.String())
		}
		return nil, status.Errorf(codes.Internal, "error getting backup: %v", err)
	}
	// TODO: Add check: if !strings.HasPrefix(backup.Name, parentPlan.BackupPlan)

	now := time.Now()

	obj := proto.Clone(req.GetRestore()).(*pb.Restore)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Uid = name.RestoreID // Consider generating a real UUID if needed
	obj.State = pb.Restore_CREATING
	obj.StateReason = "RestoreJob is being created."
	obj.Cluster = parentPlan.GetCluster()
	obj.RestoreConfig = parentPlan.GetRestoreConfig() // Inherit config

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
		// Simulate restore process
		obj.State = pb.Restore_IN_PROGRESS
		obj.StateReason = "Restoring resources."
		obj.UpdateTime = timestamppb.New(time.Now())
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			// Should we mark as failed?
			return nil, err
		}

		// Simulate completion
		lroMetadata.EndTime = timestamppb.New(time.Now())
		obj.State = pb.Restore_SUCCEEDED
		obj.StateReason = "Restore completed."
		obj.CompleteTime = lroMetadata.EndTime
		obj.UpdateTime = lroMetadata.EndTime
		// Set dummy counts
		obj.ResourcesRestoredCount = 10
		obj.ResourcesExcludedCount = 1
		obj.ResourcesFailedCount = 0
		obj.VolumesRestoredCount = 2

		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *BackupForGKEV1) UpdateRestore(ctx context.Context, req *pb.UpdateRestoreRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRestoreName(req.GetRestore().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Restore{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	now := time.Now()

	// Required. The update mask applies to the resource.
	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	// Check for updates to immutable fields
	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = req.GetRestore().GetDescription()
		case "labels":
			obj.Labels = req.GetRestore().GetLabels()
		default:
			// Check if the path refers to an immutable field or subfield.
			// Since Restore only allows description and labels updates post-creation,
			// any other path implies an attempt to modify an immutable field.
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q attempts to modify an immutable field", path)
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
		updatedObj := &pb.Restore{}
		if err := s.storage.Get(ctx, fqn, updatedObj); err != nil {
			return nil, err
		}
		return updatedObj, nil
	})
}

func (s *BackupForGKEV1) DeleteRestore(ctx context.Context, req *pb.DeleteRestoreRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseRestoreName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deletedObj := &pb.Restore{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Restore %q not found", fqn)
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

type restoreName struct {
	Project       *projects.ProjectData
	Location      string
	RestorePlanID string
	RestoreID     string
}

func (n *restoreName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/restorePlans/%s/restores/%s", n.Project.ID, n.Location, n.RestorePlanID, n.RestoreID)
}

// parseRestoreName parses a string into a restoreName.
// The expected form is `projects/*/locations/*/restorePlans/*/restores/*`.
func (s *MockService) parseRestoreName(name string) (*restoreName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "restorePlans" && tokens[6] == "restores" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &restoreName{
			Project:       project,
			Location:      tokens[3],
			RestorePlanID: tokens[5],
			RestoreID:     tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid format for a restore", name)
}
