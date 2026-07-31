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

	longrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type pipelineService struct {
	*MockService
	pb.UnimplementedPipelineServiceServer
}

func (s *pipelineService) GetPipelineJob(ctx context.Context, req *pb.GetPipelineJobRequest) (*pb.PipelineJob, error) {
	name, err := s.parsePipelineJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PipelineJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *pipelineService) CreatePipelineJob(ctx context.Context, req *pb.CreatePipelineJobRequest) (*pb.PipelineJob, error) {
	reqName := req.Parent + "/pipelineJobs/" + req.PipelineJobId
	name, err := s.parsePipelineJobName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.PipelineJob).(*pb.PipelineJob)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.StartTime = timestamppb.New(now)
	obj.EndTime = timestamppb.New(now.Add(5 * time.Second))

	obj.State = pb.PipelineState_PIPELINE_STATE_SUCCEEDED

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *pipelineService) DeletePipelineJob(ctx context.Context, req *pb.DeletePipelineJobRequest) (*longrunning.Operation, error) {
	name, err := s.parsePipelineJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.PipelineJob{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
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

type PipelineJobName struct {
	Project       *projects.ProjectData
	Location      string
	PipelineJobID string
}

func (n *PipelineJobName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/pipelineJobs/%s", n.Project.Number, n.Location, n.PipelineJobID)
}

// parsePipelineJobName parses a string into a PipelineJobName.
// The expected form of input string is projects/<projectID>/locations/<location>/pipelineJobs/<pipelineJobID>
func (s *MockService) parsePipelineJobName(name string) (*PipelineJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "pipelineJobs" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &PipelineJobName{
			Project:       project,
			Location:      tokens[3],
			PipelineJobID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}

func (s *pipelineService) GetTrainingPipeline(ctx context.Context, req *pb.GetTrainingPipelineRequest) (*pb.TrainingPipeline, error) {
	name, err := s.parseTrainingPipelineName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.TrainingPipeline{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *pipelineService) CreateTrainingPipeline(ctx context.Context, req *pb.CreateTrainingPipelineRequest) (*pb.TrainingPipeline, error) {
	id := fmt.Sprintf("tp-%d", time.Now().UnixNano())
	if req.GetTrainingPipeline().GetName() != "" {
		if name, err := s.parseTrainingPipelineName(req.GetTrainingPipeline().GetName()); err == nil {
			id = name.TrainingPipelineID
		}
	}

	reqName := req.Parent + "/trainingPipelines/" + id
	name, err := s.parseTrainingPipelineName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.TrainingPipeline).(*pb.TrainingPipeline)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.StartTime = timestamppb.New(now)
	obj.EndTime = timestamppb.New(now.Add(5 * time.Second))

	obj.State = pb.PipelineState_PIPELINE_STATE_SUCCEEDED

	// If there's a ModelToUpload, let's make sure its metadata and names are populated
	if obj.ModelToUpload != nil {
		obj.ModelToUpload.Name = fqn + "/model"
		staticTime := timestamppb.New(time.Date(2024, 4, 1, 12, 34, 56, 123456, time.UTC))
		obj.ModelToUpload.CreateTime = staticTime
		obj.ModelToUpload.UpdateTime = staticTime
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *pipelineService) DeleteTrainingPipeline(ctx context.Context, req *pb.DeleteTrainingPipelineRequest) (*longrunning.Operation, error) {
	name, err := s.parseTrainingPipelineName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()
	now := time.Now()

	deleted := &pb.TrainingPipeline{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
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

type TrainingPipelineName struct {
	Project            *projects.ProjectData
	Location           string
	TrainingPipelineID string
}

func (n *TrainingPipelineName) String() string {
	return fmt.Sprintf("projects/%d/locations/%s/trainingPipelines/%s", n.Project.Number, n.Location, n.TrainingPipelineID)
}

func (s *MockService) parseTrainingPipelineName(name string) (*TrainingPipelineName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "trainingPipelines" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		name := &TrainingPipelineName{
			Project:            project,
			Location:           tokens[3],
			TrainingPipelineID: tokens[5],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
