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

package mockrun

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/fields"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/run/v2"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/google/uuid"
	api "google.golang.org/genproto/googleapis/api"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RunV2 struct {
	*MockService
	pb.UnimplementedJobsServer
}

func (s *RunV2) GetJob(ctx context.Context, req *pb.GetJobRequest) (*pb.Job, error) {
	name, err := s.parseJobName(req.Name)
	if err != nil {
		return nil, err
	}

	project, err := s.Projects.GetProjectByID(name.Project.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found", name.Project.ID)
	}
	if project == nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found", name.Project.ID)
	}

	fqn := name.String()

	obj := &pb.Job{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (s *RunV2) CreateJob(ctx context.Context, req *pb.CreateJobRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/jobs/" + req.JobId
	name, err := s.parseJobName(reqName)
	if err != nil {
		return nil, err
	}

	project, err := s.Projects.GetProjectByID(name.Project.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found", name.Project.ID)
	}
	if project == nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found", name.Project.ID)
	}

	fqn := name.String()

	obj := proto.Clone(req.Job).(*pb.Job)
	obj.Name = fqn
	obj.CreateTime = timestamppb.Now()
	obj.UpdateTime = timestamppb.Now()
	obj.Etag = fields.ComputeWeakEtag(obj)
	obj.Creator = "test@google.com"

	obj.LastModifier = "test@google.com"
	obj.Generation = 1
	if obj.LatestCreatedExecution == nil {
		obj.LatestCreatedExecution = &pb.ExecutionReference{}
	}
	if obj.Template == nil {
		obj.Template = &pb.ExecutionTemplate{}
	}
	if obj.Template.TaskCount == 0 {
		obj.Template.TaskCount = 1
	}
	if obj.Template.Template == nil {
		obj.Template.Template = &pb.TaskTemplate{}
	}
	if obj.Template.Template.Timeout == nil {
		obj.Template.Template.Timeout = &duration.Duration{Seconds: 600}
	}
	if obj.Template.Template.ServiceAccount == "" {
		obj.Template.Template.ServiceAccount = fmt.Sprintf("%d-compute@developer.gserviceaccount.com", project.Number)
	}
	if obj.Template.Template.ExecutionEnvironment == 0 {
		obj.Template.Template.ExecutionEnvironment = pb.ExecutionEnvironment_EXECUTION_ENVIRONMENT_GEN2
	}
	for _, container := range obj.Template.Template.Containers {
		if container.Resources == nil {
			container.Resources = &pb.ResourceRequirements{
				Limits: map[string]string{
					"cpu":    "1000m",
					"memory": "512Mi",
				},
			}
		}
	}
	if obj.Template.Template.Retries == nil {
		obj.Template.Template.Retries = &pb.TaskTemplate_MaxRetries{MaxRetries: 3}
	}

	if obj.TerminalCondition == nil {
		obj.TerminalCondition = &pb.Condition{
			LastTransitionTime: timestamppb.Now(),
			State:              pb.Condition_CONDITION_SUCCEEDED,
			Type:               "Ready",
		}
	}
	obj.Uid = uuid.NewString()

	// Server-side defaults
	if obj.LaunchStage == 0 {
		obj.LaunchStage = api.LaunchStage_GA
	}
	// Note: We do not set obj.Status here, as it is expected to be empty on creation.
	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}
	return s.operations.StartLRO(ctx, req.Parent, obj, func() (proto.Message, error) {
		return obj, nil
	})
}

func (s *RunV2) UpdateJob(ctx context.Context, req *pb.UpdateJobRequest) (*longrunning.Operation, error) {
	name, err := s.parseJobName(req.GetJob().GetName())
	if err != nil {
		return nil, err
	}

	project, err := s.Projects.GetProjectByID(name.Project.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found", name.Project.ID)
	}
	if project == nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found", name.Project.ID)
	}

	fqn := name.String()

	obj := &pb.Job{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.UpdateTime = timestamppb.Now()
	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, obj, func() (protoreflect.ProtoMessage, error) {
		return obj, nil
	})
}

func (s *RunV2) DeleteJob(ctx context.Context, req *pb.DeleteJobRequest) (*longrunning.Operation, error) {
	name, err := s.parseJobName(req.Name)
	if err != nil {
		return nil, err
	}

	project, err := s.Projects.GetProjectByID(name.Project.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found", name.Project.ID)
	}
	if project == nil {
		return nil, status.Errorf(codes.NotFound, "project %q not found", name.Project.ID)
	}

	fqn := name.String()

	obj := &pb.Job{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	if err := s.storage.Delete(ctx, fqn, &pb.Job{}); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, lroPrefix, obj, func() (protoreflect.ProtoMessage, error) {
		return obj, nil
	})
}

type jobName struct {
	Project  *projects.ProjectData
	Location string
	Job      string
}

func (n *jobName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/jobs/%s", n.Project.ID, n.Location, n.Job)
}

func (s *MockService) parseJobName(name string) (*jobName, error) {
	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "jobs" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		n := &jobName{
			Project:  project,
			Location: tokens[3],
			Job:      tokens[5],
		}
		return n, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
