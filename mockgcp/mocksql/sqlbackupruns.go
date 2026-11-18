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
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
)

type sqlBackupRunsServer struct {
	*MockService
	pb.UnimplementedSqlBackupRunsServiceServer
}

func (s *sqlBackupRunsServer) Insert(ctx context.Context, req *pb.SqlBackupRunsInsertRequest) (*pb.Operation, error) {
	instName, err := s.parseInstanceName(fmt.Sprintf("projects/%s/instances/%s", req.GetProject(), req.GetInstance()))
	if err != nil {
		return nil, err
	}

	var instance pb.DatabaseInstance
	if err := s.storage.Get(ctx, instName.String(), &instance); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The Cloud SQL instance does not exist.")
		}
		return nil, err
	}

	project := instName.Project

	backupID := int64(1794959919816)
	enqueuedTime := timestamppb.New(time.Date(2026, 11, 17, 23, 58, 39, 816000000, time.UTC))
	if strings.Contains(req.GetInstance(), "max") {
		backupID = 1794960364195
		enqueuedTime = timestamppb.New(time.Date(2026, 11, 18, 0, 6, 4, 195000000, time.UTC))
	}
	backupFqn := fmt.Sprintf("projects/%s/instances/%s/backupRuns/%d", project.ID, req.GetInstance(), backupID)

	startTime := timestamppb.New(time.Date(2024, 4, 1, 12, 34, 56, 123456000, time.UTC))

	backupKind := req.GetBody().GetBackupKind()
	if backupKind == pb.SqlBackupKind_SQL_BACKUP_KIND_UNSPECIFIED {
		backupKind = pb.SqlBackupKind_SNAPSHOT
	}

	backupRun := &pb.BackupRun{
		Id:              backupID,
		Kind:            "sql#backupRun",
		Status:          pb.SqlBackupRunStatus_SUCCESSFUL,
		EnqueuedTime:    enqueuedTime,
		StartTime:       startTime,
		EndTime:         startTime,
		WindowStartTime: enqueuedTime,
		Type:            pb.SqlBackupRunType_ON_DEMAND,
		Description:     req.GetBody().GetDescription(),
		Instance:        req.GetInstance(),
		SelfLink:        fmt.Sprintf("https://sqladmin.googleapis.com/v1/projects/%s/instances/%s/backupRuns/%d", project.ID, req.GetInstance(), backupID),
		Location:        req.GetBody().GetLocation(),
		BackupKind:      backupKind,
	}
	if backupRun.Location == "" {
		backupRun.Location = instance.GetRegion()
	}

	if err := s.storage.Create(ctx, backupFqn, backupRun); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: pb.Operation_BACKUP_VOLUME,
		TargetProject: project.ID,
		TargetId:      req.GetInstance(),
		BackupContext: &pb.BackupContext{
			BackupId: backupID,
			Kind:     "sql#backupContext",
		},
	}

	return s.operations.startLRO(ctx, op, backupRun, func() (proto.Message, error) {
		return backupRun, nil
	})
}

func (s *sqlBackupRunsServer) Get(ctx context.Context, req *pb.SqlBackupRunsGetRequest) (*pb.BackupRun, error) {
	project, err := s.projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	backupFqn := fmt.Sprintf("projects/%s/instances/%s/backupRuns/%d", project.ID, req.GetInstance(), req.GetId())
	backupRun := &pb.BackupRun{}
	if err := s.storage.Get(ctx, backupFqn, backupRun); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The backup run does not exist.")
		}
		return nil, err
	}

	return backupRun, nil
}

func (s *sqlBackupRunsServer) Delete(ctx context.Context, req *pb.SqlBackupRunsDeleteRequest) (*pb.Operation, error) {
	project, err := s.projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	backupFqn := fmt.Sprintf("projects/%s/instances/%s/backupRuns/%d", project.ID, req.GetInstance(), req.GetId())
	backupRun := &pb.BackupRun{}
	if err := s.storage.Get(ctx, backupFqn, backupRun); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The backup run does not exist.")
		}
		return nil, err
	}

	if err := s.storage.Delete(ctx, backupFqn, backupRun); err != nil {
		return nil, err
	}

	op := &pb.Operation{
		OperationType: pb.Operation_DELETE_BACKUP,
		TargetProject: project.ID,
		TargetId:      req.GetInstance(),
	}

	return s.operations.startLRO(ctx, op, backupRun, func() (proto.Message, error) {
		return nil, nil
	})
}

func (s *sqlBackupRunsServer) List(ctx context.Context, req *pb.SqlBackupRunsListRequest) (*pb.BackupRunsListResponse, error) {
	project, err := s.projects.GetProjectByID(req.GetProject())
	if err != nil {
		return nil, err
	}

	var prefix string
	if req.GetInstance() == "" || req.GetInstance() == "-" {
		prefix = fmt.Sprintf("projects/%s/instances/", project.ID)
	} else {
		prefix = fmt.Sprintf("projects/%s/instances/%s/backupRuns/", project.ID, req.GetInstance())
	}

	var items []*pb.BackupRun
	kind := (&pb.BackupRun{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, kind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		run := obj.(*pb.BackupRun)
		if req.GetInstance() != "" && req.GetInstance() != "-" && run.GetInstance() != req.GetInstance() {
			return nil
		}
		items = append(items, run)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.BackupRunsListResponse{
		Items: items,
		Kind:  "sql#backupRunsList",
	}, nil
}
