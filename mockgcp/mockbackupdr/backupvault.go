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
// proto.service: google.cloud.backupdr.v1.BackupDR
// proto.message: google.cloud.backupdr.v1.BackupVault

package mockbackupdr

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/backupdr/v1"
)

func (s *BackupDRV1) GetBackupVault(ctx context.Context, req *pb.GetBackupVaultRequest) (*pb.BackupVault, error) {
	name, err := s.parseBackupVaultName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupVault{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupDRV1) CreateBackupVault(ctx context.Context, req *pb.CreateBackupVaultRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/backupVaults/%s", req.GetParent(), req.GetBackupVaultId())
	name, err := s.parseBackupVaultName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.BackupVault).(*pb.BackupVault)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pb.BackupVault_CREATING
	obj.Etag = proto.String(fields.ComputeWeakEtag(obj))
	obj.Deletable = proto.Bool(true) // default to true

	s.populateDefaultsForBackupVault(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		obj.State = pb.BackupVault_ACTIVE
		lroMetadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (s *BackupDRV1) UpdateBackupVault(ctx context.Context, req *pb.UpdateBackupVaultRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupVaultName(req.GetBackupVault().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.BackupVault{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "description":
			obj.Description = proto.String(req.GetBackupVault().GetDescription())
		case "labels":
			obj.Labels = req.GetBackupVault().GetLabels()
		case "annotations":
			obj.Annotations = req.GetBackupVault().GetAnnotations()
		case "backupMinimumEnforcedRetentionDuration":
			obj.BackupMinimumEnforcedRetentionDuration = req.GetBackupVault().GetBackupMinimumEnforcedRetentionDuration()
		case "effectiveTime":
			obj.EffectiveTime = req.GetBackupVault().GetEffectiveTime()
		case "accessRestriction":
			obj.AccessRestriction = req.GetBackupVault().GetAccessRestriction()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	obj.UpdateTime = timestamppb.New(time.Now())

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "update",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})
}

func (s *BackupDRV1) DeleteBackupVault(ctx context.Context, req *pb.DeleteBackupVaultRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupVaultName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.BackupVault{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "delete",
	}
	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

func (s *BackupDRV1) populateDefaultsForBackupVault(obj *pb.BackupVault, name *backupVaultName) {
	if obj.BackupMinimumEnforcedRetentionDuration == nil {
		obj.BackupMinimumEnforcedRetentionDuration = durationpb.New(24 * 7 * time.Hour) // 7 days
	}
	if obj.Uid == "" {
		obj.Uid = "b8271390-a8aa-11ee-9847-26abc4d7b854"
	}
	if obj.AccessRestriction == pb.BackupVault_ACCESS_RESTRICTION_UNSPECIFIED {
		obj.AccessRestriction = pb.BackupVault_WITHIN_ORGANIZATION
	}
	if obj.ServiceAccount == "" {
		obj.ServiceAccount = fmt.Sprintf("vault-%d-12345@gcp-sa-backupdr-pr.iam.gserviceaccount.com", name.Project.Number)
	}
}

type backupVaultName struct {
	Project       *projects.ProjectData
	Location      string
	BackupVaultID string
}

func (n *backupVaultName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/backupVaults/" + n.BackupVaultID
}

// parseBackupVaultName parses a string into a backupVaultName.
// The expected form is `projects/*/locations/*/backupVaults/*`.
func (s *BackupDRV1) parseBackupVaultName(name string) (*backupVaultName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupVaults" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupVaultName{
			Project:       project,
			Location:      tokens[3],
			BackupVaultID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
