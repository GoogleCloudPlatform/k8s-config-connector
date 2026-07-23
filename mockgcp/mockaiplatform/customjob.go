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

package mockaiplatform

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

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type jobService struct {
	*MockService
	pb.UnimplementedJobServiceServer
}

func (s *jobService) CreateCustomJob(ctx context.Context, req *pb.CreateCustomJobRequest) (*pb.CustomJob, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	reqName := req.Parent + "/customJobs/" + id
	name, err := s.parseCustomJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.CustomJob).(*pb.CustomJob)
	obj.Name = fqn
	obj.State = pb.JobState_JOB_STATE_SUCCEEDED
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *jobService) GetCustomJob(ctx context.Context, req *pb.GetCustomJobRequest) (*pb.CustomJob, error) {
	name, err := s.parseCustomJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	resource := &pb.CustomJob{}
	if err := s.storage.Get(ctx, fqn, resource); err != nil {
		return nil, err
	}

	return resource, nil
}

func (s *jobService) DeleteCustomJob(ctx context.Context, req *pb.DeleteCustomJobRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseCustomJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	resource := &pb.CustomJob{}
	if err := s.storage.Delete(ctx, fqn, resource); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, &emptypb.Empty{})
}

type CustomJobName struct {
	Project     *projects.ProjectData
	Location    string
	CustomJobID string
}

func (n *CustomJobName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/customJobs/%s", n.Project.Number, n.Location, n.CustomJobID)
}

func (s *MockService) parseCustomJobName(name string) (*CustomJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "customJobs" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &CustomJobName{
			Project:     project,
			Location:    tokens[3],
			CustomJobID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
