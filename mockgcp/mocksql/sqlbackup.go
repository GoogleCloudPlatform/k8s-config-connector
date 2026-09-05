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

package mocksql

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type sqlBackupRunsServer struct {
	*MockService
	pb.UnimplementedSqlBackupRunsServiceServer
}

type BackupName struct {
	Project      *projects.ProjectData
	InstanceName string
	BackupID     int64
}

func (n *BackupName) String() string {
	return "projects/" + n.Project.ID + "/instances/" + n.InstanceName + "/backupRuns/" + strconv.FormatInt(n.BackupID, 10)
}

func (s *MockService) buildBackupName(projectID, instanceName string, backupID int64) (*BackupName, error) {
	project, err := s.projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}
	return &BackupName{
		Project:      project,
		InstanceName: instanceName,
		BackupID:     backupID,
	}, nil
}

func (s *sqlBackupRunsServer) Insert(ctx context.Context, req *pb.SqlBackupRunsInsertRequest) (*pb.Operation, error) {
	// Assign a timestamp-based backup ID
	backupID := time.Now().UnixNano() / int64(time.Millisecond)

	name, err := s.buildBackupName(req.GetProject(), req.GetInstance(), backupID)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.GetBody()).(*pb.BackupRun)
	obj.Id = backupID
	obj.Instance = name.InstanceName
	obj.Kind = "sql#backupRun"
	obj.SelfLink = fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s/backupRuns/%d",
		name.Project.ID, name.InstanceName, backupID)

	staticTime := timestamppb.New(time.Date(2024, 4, 1, 12, 34, 56, 123456000, time.UTC))
	obj.EnqueuedTime = staticTime
	obj.StartTime = staticTime
	obj.EndTime = staticTime
	obj.Status = pb.SqlBackupRunStatus_SUCCESSFUL
	obj.Type = pb.SqlBackupRunType_ON_DEMAND

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_BACKUP_VOLUME,
		Status:        pb.Operation_DONE,
		BackupContext: &pb.BackupContext{
			BackupId: backupID,
			Kind:     "sql#backupContext",
		},
		Name: fmt.Sprintf("operation-%d", backupID),
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sqlBackupRunsServer) Get(ctx context.Context, req *pb.SqlBackupRunsGetRequest) (*pb.BackupRun, error) {
	name, err := s.buildBackupName(req.GetProject(), req.GetInstance(), req.GetId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupRun{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not Found")
		}
		return nil, err
	}

	return obj, nil
}

func (s *sqlBackupRunsServer) Delete(ctx context.Context, req *pb.SqlBackupRunsDeleteRequest) (*pb.Operation, error) {
	name, err := s.buildBackupName(req.GetProject(), req.GetInstance(), req.GetId())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.BackupRun{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not Found")
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_DELETE,
		Status:        pb.Operation_DONE,
		Name:          fmt.Sprintf("operation-delete-%d", name.BackupID),
	}

	return s.operations.startLRO(ctx, op, deleted, func() (proto.Message, error) {
		return deleted, nil
	})
}
