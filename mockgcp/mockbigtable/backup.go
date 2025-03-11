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
// proto.service: google.bigtable.admin.v2.BigtableTableAdmin
// proto.message: google.bigtable.admin.v2.Backup

package mockbigtable

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

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"

	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *tableAdminServer) CreateBackup(ctx context.Context, req *pb.CreateBackupRequest) (*longrunningpb.Operation, error) {
	reqName := req.Parent + "/backups/" + req.BackupId
	name, err := s.parseBackupName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Backup).(*pb.Backup)
	obj.Name = fqn

	s.populateDefaultsForBackup(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("operations/%s/locations/%s", name.String(), "us-east1-c")
	lroMetadata := &pb.CreateBackupMetadata{
		Name:        req.BackupId,
		SourceTable: req.Backup.SourceTable,
		StartTime:   timestamppb.New(time.Now()),
	}
	return s.operations.StartLRO(ctx, prefix, lroMetadata, func() (proto.Message, error) {
		return obj, nil
	})

}

func (s *tableAdminServer) populateDefaultsForBackup(obj *pb.Backup) {
	if obj.GetBackupType() == pb.Backup_BACKUP_TYPE_UNSPECIFIED {
		obj.BackupType = pb.Backup_STANDARD
	}
	if obj.GetState() == pb.Backup_STATE_UNSPECIFIED {
		obj.State = pb.Backup_READY
	}

	if obj.StartTime == nil {
		obj.StartTime = timestamppb.Now()
	}

	if obj.SourceTable == "" {
		obj.SourceTable = "sample-table"
	}

	if obj.EncryptionInfo == nil {
		obj.EncryptionInfo = &pb.EncryptionInfo{EncryptionType: pb.EncryptionInfo_GOOGLE_DEFAULT_ENCRYPTION}
	}

	if obj.EndTime == nil {
		obj.EndTime = timestamppb.New(time.Now().Add(5 * time.Minute))
	}
}

func (s *tableAdminServer) GetBackup(ctx context.Context, req *pb.GetBackupRequest) (*pb.Backup, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Backup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "backup %q not found", name)
		}
		return nil, err
	}

	return obj, nil
}

func (s *tableAdminServer) UpdateBackup(ctx context.Context, req *pb.UpdateBackupRequest) (*pb.Backup, error) {
	name, err := s.parseBackupName(req.Backup.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Backup{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "expire_time":
			obj.ExpireTime = req.Backup.ExpireTime
		default:
			return nil, status.Errorf(codes.InvalidArgument, "update_mask path %q not valid", path)
		}
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *tableAdminServer) DeleteBackup(ctx context.Context, req *pb.DeleteBackupRequest) (*emptypb.Empty, error) {
	name, err := s.parseBackupName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Backup{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

type backupName struct {
	Project  *projects.ProjectData
	Instance string
	Cluster  string
	BackupID string
}

func (n *backupName) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/clusters/%s/backups/%s", n.Project.ID, n.Instance, n.Cluster, n.BackupID)
}

// parseBackupName parses a string into a backupName.
// The expected form is `projects/*/instances/*/clusters/*/backups/*`.
func (s *tableAdminServer) parseBackupName(name string) (*backupName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "instances" && tokens[4] == "clusters" && tokens[6] == "backups" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &backupName{
			Project:  project,
			Instance: tokens[3],
			Cluster:  tokens[5],
			BackupID: tokens[7],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
