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
// proto.message: google.cloud.backupdr.v1.BackupPlan

package mockbackupdr

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
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/backupdr/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *BackupDRV1) GetBackupPlan(ctx context.Context, req *pb.GetBackupPlanRequest) (*pb.BackupPlan, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupPlan{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *BackupDRV1) CreateBackupPlan(ctx context.Context, req *pb.CreateBackupPlanRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/backupPlans/%s", req.GetParent(), req.GetBackupPlanId())
	name, err := s.parseBackupPlanName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.Clone(req.GetBackupPlan()).(*pb.BackupPlan)
	obj.Name = fqn
	now := time.Now()
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.BackupPlan_CREATING
	setDefaultServiceAccount(obj, name)

	obj.Etag = fields.ComputeWeakEtag(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "create",
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		obj.State = pb.BackupPlan_ACTIVE

		// mimic the behavior of the real GCP API
		obj.BackupVault = strings.ReplaceAll(obj.BackupVault, name.Project.ID, fmt.Sprintf("%v", name.Project.Number))

		lroMetadata.EndTime = timestamppb.New(time.Now())
		if err := s.storage.Update(ctx, fqn, obj); err != nil {
			return nil, err
		}
		return obj, nil
	})

}

func (s *BackupDRV1) DeleteBackupPlan(ctx context.Context, req *pb.DeleteBackupPlanRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupPlanName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	deleted := &pb.BackupPlan{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}
	deleted.State = pb.BackupPlan_DELETING

	lroMetadata := &pb.OperationMetadata{
		ApiVersion: "v1",
		CreateTime: timestamppb.New(now),
		Target:     name.String(),
		Verb:       "delete",
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type backupPlanName struct {
	Project    *projects.ProjectData
	Location   string
	BackupPlan string
}

func (n *backupPlanName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", n.Project.ID, n.Location, n.BackupPlan)
}

// parseBackupPlanName parses a string into an backupPlanName.
// The expected form is `projects/*/locations/*/backupPlans/*`.
func (s *MockService) parseBackupPlanName(name string) (*backupPlanName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupPlans" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupPlanName{
			Project:    project,
			Location:   tokens[3],
			BackupPlan: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}

func setDefaultServiceAccount(obj *pb.BackupPlan, name *backupPlanName) {
	if obj.BackupVaultServiceAccount == "" {
		obj.BackupVaultServiceAccount = fmt.Sprintf("vault-%d-12345@gcp-sa-backupdr-pr.iam.gserviceaccount.com", name.Project.Number)
	}
}
