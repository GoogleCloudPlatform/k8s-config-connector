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
	"math/rand"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type sqlBackupRunServer struct {
	*MockService
	pb.UnimplementedSqlBackupRunsServiceServer
}

type sqlBackupRunName struct {
	Project      *projects.ProjectData
	InstanceName string
	BackupRunID  int64
}

func (n *sqlBackupRunName) String() string {
	return fmt.Sprintf("projects/%s/instances/%s/backupRuns/%d", n.Project.ID, n.InstanceName, n.BackupRunID)
}

func (s *MockService) buildBackupRunName(projectID string, instance string, id int64) (*sqlBackupRunName, error) {
	project, err := s.projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}
	return &sqlBackupRunName{
		Project:      project,
		InstanceName: instance,
		BackupRunID:  id,
	}, nil
}

func (s *sqlBackupRunServer) Insert(ctx context.Context, req *pb.SqlBackupRunsInsertRequest) (*pb.Operation, error) {
	// Generate a unique backup ID
	backupID := time.Now().UnixNano()/int64(time.Millisecond) + rand.Int63n(100)
	name, err := s.buildBackupRunName(req.GetProject(), req.GetInstance(), backupID)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.BackupRun{}
	if req.GetBody() != nil {
		obj = proto.Clone(req.GetBody()).(*pb.BackupRun)
	}

	obj.Id = backupID
	obj.Instance = req.GetInstance()
	obj.Kind = "sql#backupRun"
	obj.Type = pb.SqlBackupRunType_ON_DEMAND
	obj.Status = pb.SqlBackupRunStatus_SUCCESSFUL

	now := time.Unix(backupID/1000, 0)
	obj.StartTime = timestamppb.New(now)
	obj.EndTime = timestamppb.New(now.Add(1 * time.Minute))
	obj.EnqueuedTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		TargetId:      fmt.Sprintf("%d", backupID),
		TargetLink:    fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s/backupRuns/%d", name.Project.ID, name.InstanceName, backupID),
		OperationType: pb.Operation_BACKUP,
		Status:        pb.Operation_DONE,
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sqlBackupRunServer) Get(ctx context.Context, req *pb.SqlBackupRunsGetRequest) (*pb.BackupRun, error) {
	name, err := s.buildBackupRunName(req.GetProject(), req.GetInstance(), req.GetId())
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

func (s *sqlBackupRunServer) Delete(ctx context.Context, req *pb.SqlBackupRunsDeleteRequest) (*pb.Operation, error) {
	name, err := s.buildBackupRunName(req.GetProject(), req.GetInstance(), req.GetId())
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

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_DELETE_BACKUP,
		Status:        pb.Operation_DONE,
	}

	return s.operations.startLRO(ctx, op, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *sqlBackupRunServer) List(ctx context.Context, req *pb.SqlBackupRunsListRequest) (*pb.BackupRunsListResponse, error) {
	project, err := s.projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/instances/%s/backupRuns/", project.ID, req.GetInstance())

	var items []*pb.BackupRun
	var list pb.BackupRunsListResponse

	findKind := (&pb.BackupRun{}).ProtoReflect().Descriptor()
	findFunc := func(obj proto.Message) error {
		run := obj.(*pb.BackupRun)
		items = append(items, run)
		return nil
	}

	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: prefix}, findFunc); err != nil {
		return nil, err
	}

	list.Items = items
	list.Kind = "sql#backupRunsList"

	return &list, nil
}
