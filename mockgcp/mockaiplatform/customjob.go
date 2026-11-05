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
// proto.service: google.cloud.aiplatform.v1.JobService
// proto.message: google.cloud.aiplatform.v1.CustomJob

package mockaiplatform

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type jobService struct {
	*MockService
	pb.UnimplementedJobServiceServer
}

func (s *jobService) GetCustomJob(ctx context.Context, req *pb.GetCustomJobRequest) (*pb.CustomJob, error) {
	name, err := s.parseCustomJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.CustomJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.Clone(obj).(*pb.CustomJob)
	result.Name = name.StringWithProjectNumber()
	populateCustomJobDefaults(result)
	return result, nil
}

func (s *jobService) CreateCustomJob(ctx context.Context, req *pb.CreateCustomJobRequest) (*pb.CustomJob, error) {
	parent, err := s.parseCustomJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	name := &CustomJobName{
		Project:   parent.Project,
		Location:  parent.Location,
		CustomJob: id,
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.Clone(req.GetCustomJob()).(*pb.CustomJob)
	obj.Name = name.StringWithProjectNumber()
	obj.State = pb.JobState_JOB_STATE_PENDING
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	populateCustomJobDefaults(obj)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.Clone(obj).(*pb.CustomJob)
	return result, nil
}

func (s *jobService) CancelCustomJob(ctx context.Context, req *pb.CancelCustomJobRequest) (*emptypb.Empty, error) {
	name, err := s.parseCustomJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.CustomJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.JobState_JOB_STATE_CANCELLED
	now := time.Now()
	obj.EndTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)
	obj.Error = &statuspb.Status{
		Code:    int32(codes.Canceled),
		Message: "CANCELED",
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *jobService) DeleteCustomJob(ctx context.Context, req *pb.DeleteCustomJobRequest) (*longrunning.Operation, error) {
	name, err := s.parseCustomJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	deleted := &pb.CustomJob{}
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

func (s *jobService) ListCustomJobs(ctx context.Context, req *pb.ListCustomJobsRequest) (*pb.ListCustomJobsResponse, error) {
	parent, err := s.parseCustomJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/customJobs/", parent.Project.Number, parent.Location)
	findPrefixID := fmt.Sprintf("projects/%s/locations/%s/customJobs/", parent.Project.ID, parent.Location)

	var customJobs []*pb.CustomJob
	customJobKind := (&pb.CustomJob{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, customJobKind, storage.ListOptions{}, func(obj proto.Message) error {
		job := obj.(*pb.CustomJob)
		if strings.HasPrefix(job.GetName(), findPrefix) || strings.HasPrefix(job.GetName(), findPrefixID) {
			result := proto.Clone(job).(*pb.CustomJob)
			populateCustomJobDefaults(result)
			customJobs = append(customJobs, result)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListCustomJobsResponse{
		CustomJobs: customJobs,
	}, nil
}

func populateCustomJobDefaults(job *pb.CustomJob) {
	if job.JobSpec != nil {
		for _, wp := range job.JobSpec.WorkerPoolSpecs {
			if wp.DiskSpec == nil {
				wp.DiskSpec = &pb.DiskSpec{
					BootDiskSizeGb: 100,
					BootDiskType:   "pd-ssd",
				}
			} else {
				if wp.DiskSpec.BootDiskSizeGb == 0 {
					wp.DiskSpec.BootDiskSizeGb = 100
				}
				if wp.DiskSpec.BootDiskType == "" {
					wp.DiskSpec.BootDiskType = "pd-ssd"
				}
			}
		}
	}
}

type CustomJobName struct {
	Project   *projects.ProjectData
	Location  string
	CustomJob string
}

func (n *CustomJobName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/customJobs/%s", n.Project.ID, n.Location, n.CustomJob)
}

func (n *CustomJobName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/customJobs/%s", n.Project.Number, n.Location, n.CustomJob)
}

func (n *CustomJobName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/customJobs/%s", n.Project.Number, n.Location, n.CustomJob)
}

type CustomJobParent struct {
	Project  *projects.ProjectData
	Location string
}

func (s *MockService) parseCustomJobParent(parent string) (*CustomJobParent, error) {
	tokens := strings.Split(parent, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "locations" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		return &CustomJobParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

// parseCustomJobName parses a string into a CustomJobName.
// The expected form is projects/<projectID_or_number>/locations/<location>/customJobs/<customJobID>
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

		customJobID := tokens[5]
		if !isNumeric(customJobID) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid custom job id: %s", customJobID)
		}

		return &CustomJobName{
			Project:   project,
			Location:  tokens[3],
			CustomJob: customJobID,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
