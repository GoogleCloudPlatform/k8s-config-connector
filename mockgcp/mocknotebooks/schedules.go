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

// +tool:mockgcp-support
// proto.service: google.cloud.notebooks.v1.NotebookService
// proto.message: google.cloud.notebooks.v1.Schedule

package mocknotebooks

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

	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *NotebookServiceV1) GetSchedule(ctx context.Context, req *pb.GetScheduleRequest) (*pb.Schedule, error) {
	name, err := s.parseScheduleName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	obj := &pb.Schedule{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "schedule %q not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *NotebookServiceV1) CreateSchedule(ctx context.Context, req *pb.CreateScheduleRequest) (*longrunningpb.Operation, error) {
	reqName := fmt.Sprintf("%s/schedules/%s", req.GetParent(), req.GetScheduleId())
	name, err := s.parseScheduleName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	obj := proto.CloneOf(req.GetSchedule())
	obj.Name = fqn
	obj.CreateTime = timestamppb.New(time.Now())
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "create",
		Endpoint:              "CreateSchedule",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return obj, nil
	})
}

func (s *NotebookServiceV1) DeleteSchedule(ctx context.Context, req *pb.DeleteScheduleRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseScheduleName(req.GetName())
	if err != nil {
		return nil, err
	}
	fqn := name.String()

	deletedObj := &pb.Schedule{}
	if err := s.storage.Delete(ctx, fqn, deletedObj); err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	metadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(time.Now()),
		RequestedCancellation: false,
		Target:                name.String(),
		Verb:                  "delete",
		Endpoint:              "DeleteSchedule",
	}
	return s.operations.StartLRO(ctx, prefix, metadata, func() (proto.Message, error) {
		metadata.EndTime = timestamppb.New(time.Now())
		return &emptypb.Empty{}, nil
	})
}

type scheduleName struct {
	Project  *projects.ProjectData
	Location string
	Schedule string
}

func (n *scheduleName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/schedules/" + n.Schedule
}

// parseScheduleName parses a string into a scheduleName.
// The expected form is projects/<projectID>/locations/<location>/schedules/<schedule>.
func (s *MockService) parseScheduleName(name string) (*scheduleName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "schedules" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &scheduleName{
			Project:  project,
			Location: tokens[3],
			Schedule: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
