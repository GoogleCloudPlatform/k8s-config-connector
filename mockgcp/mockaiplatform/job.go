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

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/google/uuid"
)

type jobService struct {
	*MockService
	pb.UnimplementedJobServiceServer
}

type DataLabelingJobName struct {
	Project           *projects.ProjectData
	Location          string
	DataLabelingJobID string
}

func (n *DataLabelingJobName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/dataLabelingJobs/%s", n.Project.Number, n.Location, n.DataLabelingJobID)
}

func (s *MockService) parseDataLabelingJobName(name string) (*DataLabelingJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "dataLabelingJobs" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &DataLabelingJobName{
			Project:           project,
			Location:          tokens[3],
			DataLabelingJobID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *jobService) GetDataLabelingJob(ctx context.Context, req *pb.GetDataLabelingJobRequest) (*pb.DataLabelingJob, error) {
	name, err := s.parseDataLabelingJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.DataLabelingJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *jobService) CreateDataLabelingJob(ctx context.Context, req *pb.CreateDataLabelingJobRequest) (*pb.DataLabelingJob, error) {
	id := uuid.NewString()
	reqName := req.Parent + "/dataLabelingJobs/" + id
	name, err := s.parseDataLabelingJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.DataLabelingJob).(*pb.DataLabelingJob)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.State = pb.JobState_JOB_STATE_SUCCEEDED

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *jobService) DeleteDataLabelingJob(ctx context.Context, req *pb.DeleteDataLabelingJobRequest) (*longrunningpb.Operation, error) {
	name, err := s.parseDataLabelingJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.DataLabelingJob{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	op := &pb.DeleteOperationMetadata{}
	op.GenericMetadata = &pb.GenericOperationMetadata{
		CreateTime: timestamppb.New(now),
		UpdateTime: timestamppb.New(now),
	}
	opPrefix := fmt.Sprintf("projects/%d/locations/%s", name.Project.Number, name.Location)
	return s.operations.DoneLRO(ctx, opPrefix, op, nil)
}
