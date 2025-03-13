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
// proto.service: google.cloud.aiplatform.v1.ScheduleService
// proto.message: google.cloud.aiplatform.v1.Schedule

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/aiplatform/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
)

type scheduleService struct {
	*MockService
	pb.UnimplementedScheduleServiceServer
}

func (s *scheduleService) GetSchedule(ctx context.Context, req *pb.GetScheduleRequest) (*pb.Schedule, error) {
	name, err := s.parseScheduleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Schedule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Name = strings.Replace(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number), -1)
	return obj, nil
}

func (s *scheduleService) ListSchedules(ctx context.Context, req *pb.ListSchedulesRequest) (*pb.ListSchedulesResponse, error) {
	response := &pb.ListSchedulesResponse{}

	prefix := req.GetParent()
	scheduleKind := (&pb.Schedule{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, scheduleKind, storage.ListOptions{
		Prefix: prefix,
	}, func(obj proto.Message) error {
		schedule := obj.(*pb.Schedule)

		scheduleCopy := deepcopy.DeepCopy(*schedule).(pb.Schedule)
		scheduleCopy.UpdateTime = nil
		name, err := s.parseScheduleName(scheduleCopy.Name)
		if err != nil {
			return err
		}
		scheduleCopy.Name = strings.Replace(scheduleCopy.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number), -1)

		response.Schedules = append(response.Schedules, &scheduleCopy)
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

func (s *scheduleService) CreateSchedule(ctx context.Context, req *pb.CreateScheduleRequest) (*pb.Schedule, error) {
	reqName := req.Parent + "/schedules/" + req.Schedule.DisplayName
	name, err := s.parseScheduleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := proto.Clone(req.Schedule).(*pb.Schedule)
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())
	obj.NextRunTime = timestamppb.New(time.Now())
	obj.StartTime = timestamppb.New(time.Now())
	obj.State = pb.Schedule_ACTIVE

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.Name = strings.Replace(obj.Name, name.Project.ID, fmt.Sprintf("%v", name.Project.Number), -1)
	return obj, nil
}

func (s *scheduleService) UpdateSchedule(ctx context.Context, req *pb.UpdateScheduleRequest) (*pb.Schedule, error) {
	name, err := s.parseScheduleName(req.GetSchedule().GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()
	obj := &pb.Schedule{}

	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	paths := req.GetUpdateMask().GetPaths()
	if len(paths) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "update_mask must be provided")
	}

	for _, path := range paths {
		switch path {
		case "display_name", "displayName":
			obj.DisplayName = req.GetSchedule().GetDisplayName()
		case "start_time", "startTime":
			obj.StartTime = req.GetSchedule().GetStartTime()
		case "end_time", "endTime":
			obj.EndTime = req.GetSchedule().GetEndTime()
		case "cron":
			obj.TimeSpecification = &pb.Schedule_Cron{Cron: req.GetSchedule().GetCron()}
		case "max_concurrent_run_count", "maxConcurrentRunCount":
			obj.MaxConcurrentRunCount = req.GetSchedule().GetMaxConcurrentRunCount()
		case "allow_queueing", "allowQueueing":
			obj.AllowQueueing = req.GetSchedule().GetAllowQueueing()
		default:
			return nil, status.Errorf(codes.InvalidArgument, "field %q is not supported in update_mask", path)
		}
	}
	obj.UpdateTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *scheduleService) DeleteSchedule(ctx context.Context, req *pb.DeleteScheduleRequest) (*longrunning.Operation, error) {
	name, err := s.parseScheduleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.Schedule{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(time.Now()),
		UpdateTime: timestamppb.New(time.Now()),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

func (s *scheduleService) PauseSchedule(ctx context.Context, req *pb.PauseScheduleRequest) (*emptypb.Empty, error) {
	name, err := s.parseScheduleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Schedule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.State = pb.Schedule_PAUSED
	obj.LastPauseTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *scheduleService) ResumeSchedule(ctx context.Context, req *pb.ResumeScheduleRequest) (*emptypb.Empty, error) {
	name, err := s.parseScheduleName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := &pb.Schedule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	obj.State = pb.Schedule_ACTIVE
	obj.LastResumeTime = timestamppb.New(time.Now())
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

type scheduleName struct {
	Project      *projects.ProjectData
	Location     string
	ScheduleName string
}

func (n *scheduleName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/schedules/%s", n.Project.ID, n.Location, n.ScheduleName)
}

// parseScheduleName parses a string into a scheduleName.
// The expected form is `projects/*/locations/*/schedules/*`.
func (s *MockService) parseScheduleName(name string) (*scheduleName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "schedules" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &scheduleName{
			Project:      project,
			Location:     tokens[3],
			ScheduleName: tokens[5],
		}

		return name, nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
