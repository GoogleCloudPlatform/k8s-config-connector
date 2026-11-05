// Copyright 2025 Google LLC
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
// proto.message: google.cloud.netapp.v1.BackupVault

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
)

func (s *backupVaultsService) GetBackupVault(ctx context.Context, req *pb.GetBackupVaultRequest) (*pb.BackupVault, error) {
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

func (s *backupVaultsService) CreateBackupVault(ctx context.Context, req *pb.CreateBackupVaultRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/backupVaults/" + req.BackupVaultId
	name, err := s.parseBackupVaultName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	obj := proto.CloneOf(req.GetBackupVault())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(now)
	obj.State = pb.BackupVault_READY

	s.populateDefaultsForBackupVault(obj)

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

func (s *backupVaultsService) UpdateBackupVault(ctx context.Context, req *pb.UpdateBackupVaultRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupVaultName(req.BackupVault.Name)
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackupVault{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Requested entity was not found.")
		}
		return nil, err
	}
	if err := fields.UpdateByFieldMask(obj, req.BackupVault, req.UpdateMask.Paths); err != nil {
		return nil, err
	}

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

func (s *backupVaultsService) DeleteBackupVault(ctx context.Context, req *pb.DeleteBackupVaultRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseBackupVaultName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	existing := &pb.BackupVault{}
	err = s.storage.Delete(ctx, fqn, existing)
	if err != nil {
		return &longrunningpb.Operation{}, err
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

func (s *backupVaultsService) populateDefaultsForBackupVault(obj *pb.BackupVault) {
	if obj.BackupVaultType == pb.BackupVault_BACKUP_VAULT_TYPE_UNSPECIFIED {
		obj.BackupVaultType = pb.BackupVault_IN_REGION
	}
	if obj.BackupRetentionPolicy == nil {
		obj.BackupRetentionPolicy = &pb.BackupVault_BackupRetentionPolicy{}
	}
}

type backupVaultName struct {
	Project       *projects.ProjectData
	Location      string
	BackupVaultId string
}

func (n *backupVaultName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", n.Project.ID, n.Location, n.BackupVaultId)
}

// parseBackupVaultName parses a string into a backupVaultName.
// The expected form is `projects/*/locations/*/backupVaults/*`.
func (s *backupVaultsService) parseBackupVaultName(name string) (*backupVaultName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "backupVaults" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupVaultName{
			Project:       project,
			Location:      tokens[3],
			BackupVaultId: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
