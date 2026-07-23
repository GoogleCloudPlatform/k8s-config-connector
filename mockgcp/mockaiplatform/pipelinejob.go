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
	"google.golang.org/protobuf/types/known/structpb"
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

	// Progress state from PENDING to SUCCEEDED after some time
	if obj.State == pb.PipelineState_PIPELINE_STATE_PENDING {
		createTime := obj.CreateTime.AsTime()
		if time.Since(createTime) >= 4*time.Second {
			obj.State = pb.PipelineState_PIPELINE_STATE_SUCCEEDED
			obj.StartTime = timestamppb.New(createTime.Add(1 * time.Second))
			obj.EndTime = timestamppb.New(createTime.Add(3 * time.Second))
			if obj.JobDetail == nil {
				obj.JobDetail = &pb.PipelineJobDetail{}
			}
			// Save the updated state
			if err := s.storage.Update(ctx, fqn, obj); err != nil {
				return nil, err
			}
		}
	}

	return obj, nil
}

func alignPipelineSpec(spec *structpb.Struct) {
	if spec == nil || spec.Fields == nil {
		return
	}
	configVal, ok := spec.Fields["deploymentConfig"]
	if !ok {
		return
	}
	configStruct := configVal.GetStructValue()
	if configStruct == nil || configStruct.Fields == nil {
		return
	}
	executors, ok := configStruct.Fields["executors"]
	if !ok {
		return
	}

	// Move executors to deploymentSpec
	spec.Fields["deploymentSpec"] = structpb.NewStructValue(&structpb.Struct{
		Fields: map[string]*structpb.Value{
			"executors": executors,
		},
	})

	// Delete from deploymentConfig
	delete(configStruct.Fields, "executors")
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

	alignPipelineSpec(obj.PipelineSpec)

	obj.State = pb.PipelineState_PIPELINE_STATE_PENDING
	obj.ServiceAccount = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", name.Project.Number)

	if obj.Labels == nil {
		obj.Labels = make(map[string]string)
	}
	obj.Labels["vertex-ai-pipelines-run-billing-id"] = "619702208161644544"

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

func (s *pipelineService) CancelPipelineJob(ctx context.Context, req *pb.CancelPipelineJobRequest) (*emptypb.Empty, error) {
	name, err := s.parsePipelineJobName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.PipelineJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.PipelineState_PIPELINE_STATE_CANCELLED
	obj.UpdateTime = timestamppb.New(time.Now())

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
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
