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
// proto.message: google.cloud.backupdr.v1.BackupPlanAssociation

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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/backupdr/v1"
	"github.com/google/uuid"
)

func (s *BackupDRV1) GetBackupPlanAssociation(ctx context.Context, req *pb.GetBackupPlanAssociationRequest) (*pb.BackupPlanAssociation, error) {
	name, err := s.parseBackupPlanAssociationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupPlanAssociation{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupDRV1) CreateBackupPlanAssociation(ctx context.Context, req *pb.CreateBackupPlanAssociationRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/backupPlanAssociations/%s", req.GetParent(), req.GetBackupPlanAssociationId())
	name, err := s.parseBackupPlanAssociationName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.BackupPlanAssociation).(*pb.BackupPlanAssociation)

	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.State = pb.BackupPlanAssociation_CREATING
	s.populateDefaultsForBackupPlanAssociation(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.Now(),
		Target:     name.String(),
		Verb:       "create",
		AdditionalInfo: map[string]string{
			"backupPlan":   obj.GetBackupPlan(),
			"resource":     obj.GetResource(),
			"resourceType": obj.GetResourceType(),
		},
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.Now()

		obj.State = pb.BackupPlanAssociation_ACTIVE
		// change project ID to project Number
		obj.BackupPlan = fmt.Sprintf("projects/%d/locations/%s/backupPlans/%s", name.Project.Number, name.Location, strings.TrimPrefix(req.BackupPlanAssociation.GetBackupPlan(), fmt.Sprintf("projects/%s/locations/%s/backupPlans/", name.Project.ID, name.Location)))
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}

		return obj, nil
	})
}

func (s *BackupDRV1) DeleteBackupPlanAssociation(ctx context.Context, req *pb.DeleteBackupPlanAssociationRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanAssociationName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.BackupPlanAssociation{}
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
		lroMetadata.EndTime = timestamppb.Now()
		return &emptypb.Empty{}, nil
	})
}

type backupPlanAssociationName struct {
	Project                 *projects.ProjectData
	Location                string
	BackupPlanAssociationID string
}

func (n *backupPlanAssociationName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/backupPlanAssociations/" + n.BackupPlanAssociationID
}

// parseBackupPlanAssociationName parses a string into a backupPlanAssociationName.gi
// The expected form is `projects/*/locations/*/backupPlanAssociations/*`.
func (s *BackupDRV1) parseBackupPlanAssociationName(name string) (*backupPlanAssociationName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupPlanAssociations" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupPlanAssociationName{
			Project:                 project,
			Location:                tokens[3],
			BackupPlanAssociationID: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func (s *MockService) populateDefaultsForBackupPlanAssociation(obj *pb.BackupPlanAssociation, name *backupPlanAssociationName) {
	if obj.RulesConfigInfo == nil {
		obj.RulesConfigInfo = []*pb.RuleConfigInfo{
			{
				LastBackupState: pb.RuleConfigInfo_FIRST_BACKUP_PENDING,
				RuleId:          "rule-1",
			},
		}
	}
	uuid := uuid.New().String()
	backupVaultName := name.BackupPlanAssociationID // backup vault name is refenreced by the backup plan, which we cannot get from the request, just use the name of the backup plan association
	obj.DataSource = fmt.Sprintf("projects/%d/locations/%s/backupVaults/%s/dataSources/%s", name.Project.Number, name.Location, backupVaultName, uuid)
}
