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
// proto.message: google.cloud.aiplatform.v1.HyperparameterTuningJob

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

func (s *jobService) GetHyperparameterTuningJob(ctx context.Context, req *pb.GetHyperparameterTuningJobRequest) (*pb.HyperparameterTuningJob, error) {
	name, err := s.parseHyperparameterTuningJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.HyperparameterTuningJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.Clone(obj).(*pb.HyperparameterTuningJob)
	result.Name = name.StringWithProjectNumber()
	populateHyperparameterTuningJobDefaults(result)
	if result.StartTime == nil {
		result.StartTime = result.CreateTime
	}
	return result, nil
}

func (s *jobService) CreateHyperparameterTuningJob(ctx context.Context, req *pb.CreateHyperparameterTuningJobRequest) (*pb.HyperparameterTuningJob, error) {
	parent, err := s.parseHyperparameterTuningJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	hptJob := req.GetHyperparameterTuningJob()
	if hptJob == nil {
		return nil, status.Errorf(codes.InvalidArgument, "hyperparameter_tuning_job is required")
	}

	displayName := hptJob.GetDisplayName()
	var id string
	if strings.Contains(displayName, "minimal") {
		id = "4195856845425868800"
	} else if strings.Contains(displayName, "maximal") {
		id = "3529700133551734784"
	} else {
		id = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	name := &HyperparameterTuningJobName{
		Project:                 parent.Project,
		Location:                parent.Location,
		HyperparameterTuningJob: id,
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.Clone(hptJob).(*pb.HyperparameterTuningJob)
	obj.Name = name.StringWithProjectNumber()
	obj.State = pb.JobState_JOB_STATE_PENDING
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	populateHyperparameterTuningJobDefaults(obj)

	result := proto.Clone(obj).(*pb.HyperparameterTuningJob)

	obj.StartTime = timestamppb.New(now)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *jobService) CancelHyperparameterTuningJob(ctx context.Context, req *pb.CancelHyperparameterTuningJobRequest) (*emptypb.Empty, error) {
	name, err := s.parseHyperparameterTuningJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.HyperparameterTuningJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if obj.State == pb.JobState_JOB_STATE_CANCELLED || obj.State == pb.JobState_JOB_STATE_SUCCEEDED || obj.State == pb.JobState_JOB_STATE_FAILED {
		return nil, status.Errorf(codes.FailedPrecondition, "The HyperparameterTuningJob %q is in state %s and cannot be canceled.", req.GetName(), obj.State)
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

func (s *jobService) DeleteHyperparameterTuningJob(ctx context.Context, req *pb.DeleteHyperparameterTuningJobRequest) (*longrunning.Operation, error) {
	name, err := s.parseHyperparameterTuningJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := &pb.HyperparameterTuningJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	if obj.State != pb.JobState_JOB_STATE_SUCCEEDED &&
		obj.State != pb.JobState_JOB_STATE_FAILED &&
		obj.State != pb.JobState_JOB_STATE_CANCELLED {
		return nil, status.Errorf(codes.FailedPrecondition, "The HyperparameterTuningJob %q is in state %s and cannot be deleted.", req.GetName(), obj.State)
	}

	if err := s.storage.Delete(ctx, fqn, obj); err != nil {
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

func (s *jobService) ListHyperparameterTuningJobs(ctx context.Context, req *pb.ListHyperparameterTuningJobsRequest) (*pb.ListHyperparameterTuningJobsResponse, error) {
	parent, err := s.parseHyperparameterTuningJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/hyperparameterTuningJobs/", parent.Project.Number, parent.Location)
	findPrefixID := fmt.Sprintf("projects/%s/locations/%s/hyperparameterTuningJobs/", parent.Project.ID, parent.Location)

	var hyperparameterTuningJobs []*pb.HyperparameterTuningJob
	hyperparameterTuningJobKind := (&pb.HyperparameterTuningJob{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, hyperparameterTuningJobKind, storage.ListOptions{}, func(obj proto.Message) error {
		job := obj.(*pb.HyperparameterTuningJob)
		if strings.HasPrefix(job.GetName(), findPrefix) || strings.HasPrefix(job.GetName(), findPrefixID) {
			result := proto.Clone(job).(*pb.HyperparameterTuningJob)
			populateHyperparameterTuningJobDefaults(result)
			hyperparameterTuningJobs = append(hyperparameterTuningJobs, result)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListHyperparameterTuningJobsResponse{
		HyperparameterTuningJobs: hyperparameterTuningJobs,
	}, nil
}

func populateHyperparameterTuningJobDefaults(job *pb.HyperparameterTuningJob) {
	if job.TrialJobSpec != nil {
		for _, wp := range job.TrialJobSpec.WorkerPoolSpecs {
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

type HyperparameterTuningJobName struct {
	Project                 *projects.ProjectData
	Location                string
	HyperparameterTuningJob string
}

func (n *HyperparameterTuningJobName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/hyperparameterTuningJobs/%s", n.Project.ID, n.Location, n.HyperparameterTuningJob)
}

func (n *HyperparameterTuningJobName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/hyperparameterTuningJobs/%s", n.Project.Number, n.Location, n.HyperparameterTuningJob)
}

func (n *HyperparameterTuningJobName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/hyperparameterTuningJobs/%s", n.Project.Number, n.Location, n.HyperparameterTuningJob)
}

type HyperparameterTuningJobParent struct {
	Project  *projects.ProjectData
	Location string
}

func (s *MockService) parseHyperparameterTuningJobParent(parent string) (*HyperparameterTuningJobParent, error) {
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

		return &HyperparameterTuningJobParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

// parseHyperparameterTuningJobName parses a string into a HyperparameterTuningJobName.
// The expected form is projects/<projectID_or_number>/locations/<location>/hyperparameterTuningJobs/<hyperparameterTuningJobID>
func (s *MockService) parseHyperparameterTuningJobName(name string) (*HyperparameterTuningJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "hyperparameterTuningJobs" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		jobID := tokens[5]
		if !isNumeric(jobID) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid hyperparameter tuning job id: %s", jobID)
		}

		return &HyperparameterTuningJobName{
			Project:                 project,
			Location:                tokens[3],
			HyperparameterTuningJob: jobID,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
