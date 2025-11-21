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

package mockfirestore

import (
	"context"
	"strings"
	"time"

	pb "cloud.google.com/go/firestore/apiv1/admin/adminpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"k8s.io/apimachinery/pkg/util/uuid"
)

func (s *firestoreAdminServer) CreateBackupSchedule(ctx context.Context, req *pb.CreateBackupScheduleRequest) (*pb.BackupSchedule, error) {
	name, err := s.parseBackupScheduleName(req.GetParent() + "/backupSchedules/" + string(uuid.NewUUID()))
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	now := time.Now()

	obj := req.BackupSchedule
	if err := s.populateBackupScheduleDefaults(obj); err != nil {
		return nil, err
	}

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	obj.Name = fqn

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *firestoreAdminServer) GetBackupSchedule(ctx context.Context, req *pb.GetBackupScheduleRequest) (*pb.BackupSchedule, error) {
	name, err := s.parseBackupScheduleName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackupSchedule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *firestoreAdminServer) ListBackupSchedules(ctx context.Context, req *pb.ListBackupSchedulesRequest) (*pb.ListBackupSchedulesResponse, error) {
	parent, err := s.parseDatabaseName(req.GetParent())
	if err != nil {
		return nil, err
	}

	prefix := parent.String() + "/backupSchedules/"

	response := &pb.ListBackupSchedulesResponse{}

	backupScheduleKind := (&pb.BackupSchedule{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, backupScheduleKind, storage.ListOptions{Prefix: prefix}, func(obj proto.Message) error {
		response.BackupSchedules = append(response.BackupSchedules, obj.(*pb.BackupSchedule))
		return nil
	}); err != nil {
		return nil, err
	}

	return response, nil
}

func (s *firestoreAdminServer) UpdateBackupSchedule(ctx context.Context, req *pb.UpdateBackupScheduleRequest) (*pb.BackupSchedule, error) {
	name, err := s.parseBackupScheduleName(req.GetBackupSchedule().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.BackupSchedule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	updatePaths := req.GetUpdateMask().GetPaths()
	if len(updatePaths) == 0 {
		updatePaths = fields.ComputeImpliedFieldMask(ctx, req.GetBackupSchedule(), "name")
	}

	for _, path := range updatePaths {
		switch path {
		case "retention":
			obj.Retention = req.BackupSchedule.Retention
		case "recurrence":
			obj.Recurrence = req.BackupSchedule.Recurrence
		default:
			return nil, status.Errorf(codes.InvalidArgument, "Update mask path %q not supported (by mockgcp)", path)
		}
	}

	// The service seems to drop DailyRecurrence if it's empty, but that seems more like a bug.

	s.populateBackupScheduleDefaults(obj)

	obj.UpdateTime = timestamppb.Now()

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *firestoreAdminServer) DeleteBackupSchedule(ctx context.Context, req *pb.DeleteBackupScheduleRequest) (*emptypb.Empty, error) {
	name, err := s.parseBackupScheduleName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	if err := s.storage.Delete(ctx, fqn, &pb.BackupSchedule{}); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *firestoreAdminServer) populateBackupScheduleDefaults(obj *pb.BackupSchedule) error {
	return nil
}

type backupScheduleName struct {
	Project    *projects.ProjectData
	DatabaseID string
	ScheduleID string
}

func (n *backupScheduleName) String() string {
	return "projects/" + n.Project.ID + "/databases/" + n.DatabaseID + "/backupSchedules/" + n.ScheduleID
}

func (s *firestoreAdminServer) parseBackupScheduleName(name string) (*backupScheduleName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "databases" && tokens[4] == "backupSchedules" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}
		name := &backupScheduleName{
			Project:    project,
			DatabaseID: tokens[3],
			ScheduleID: tokens[5],
		}
		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
