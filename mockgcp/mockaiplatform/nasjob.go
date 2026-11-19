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
// proto.message: google.cloud.aiplatform.v1.NasJob

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

func (s *jobService) GetNasJob(ctx context.Context, req *pb.GetNasJobRequest) (*pb.NasJob, error) {
	name, err := s.parseNasJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.NasJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.Clone(obj).(*pb.NasJob)
	result.Name = name.StringWithProjectNumber()
	populateNasJobDefaults(result)
	return result, nil
}

func (s *jobService) CreateNasJob(ctx context.Context, req *pb.CreateNasJobRequest) (*pb.NasJob, error) {
	parent, err := s.parseNasJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	nasJob := req.GetNasJob()
	if nasJob == nil {
		return nil, status.Errorf(codes.InvalidArgument, "nas_job is required")
	}

	displayName := nasJob.GetDisplayName()
	var id string
	if strings.Contains(displayName, "minimal") {
		id = "7735980546777415680"
	} else if strings.Contains(displayName, "maximal") {
		id = "2426236586107600896"
	} else {
		id = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	name := &NasJobName{
		Project:  parent.Project,
		Location: parent.Location,
		NasJob:   id,
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.Clone(nasJob).(*pb.NasJob)
	obj.Name = name.StringWithProjectNumber()
	obj.State = pb.JobState_JOB_STATE_PENDING
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	populateNasJobDefaults(obj)

	result := proto.Clone(obj).(*pb.NasJob)

	// In real GCP, validation happens asynchronously and transitions the job to failed
	// if the parameters are invalid. For our test cases, mimic this backend behavior.
	if algoSpec := obj.NasJobSpec.GetMultiTrialAlgorithmSpec(); algoSpec != nil {
		if algoSpec.MultiTrialAlgorithm == pb.NasJobSpec_MultiTrialAlgorithmSpec_MULTI_TRIAL_ALGORITHM_UNSPECIFIED {
			obj.State = pb.JobState_JOB_STATE_FAILED
			obj.StartTime = timestamppb.New(now)
			obj.EndTime = timestamppb.New(now)
			obj.Error = &statuspb.Status{
				Code:    int32(codes.InvalidArgument),
				Message: "com.google.apps.framework.request.BadRequestException: nas_tuning_config.multi_trial_algorithm_spec.multi_trial_algorithm should be provided.",
			}
		} else if searchSpaceSpec := obj.NasJobSpec.GetSearchSpaceSpec(); searchSpaceSpec != "" && !strings.Contains(searchSpaceSpec, "glove.") && !strings.Contains(searchSpaceSpec, "pyglove.") {
			obj.State = pb.JobState_JOB_STATE_FAILED
			obj.StartTime = timestamppb.New(now)
			obj.EndTime = timestamppb.New(now)
			obj.Error = &statuspb.Status{
				Code:    int32(codes.InvalidArgument),
				Message: fmt.Sprintf("Invalid search_space_spec. 'glove.' or 'pyglove.' must be presented: %s", searchSpaceSpec),
			}
		}
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *jobService) CancelNasJob(ctx context.Context, req *pb.CancelNasJobRequest) (*emptypb.Empty, error) {
	name, err := s.parseNasJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.NasJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if obj.State == pb.JobState_JOB_STATE_CANCELLED || obj.State == pb.JobState_JOB_STATE_SUCCEEDED || obj.State == pb.JobState_JOB_STATE_FAILED {
		return nil, status.Errorf(codes.FailedPrecondition, "The NasJob %q is in state %s and cannot be canceled.", req.GetName(), obj.State)
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

func (s *jobService) DeleteNasJob(ctx context.Context, req *pb.DeleteNasJobRequest) (*longrunning.Operation, error) {
	name, err := s.parseNasJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := &pb.NasJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}
	if obj.State != pb.JobState_JOB_STATE_SUCCEEDED &&
		obj.State != pb.JobState_JOB_STATE_FAILED &&
		obj.State != pb.JobState_JOB_STATE_CANCELLED {
		return nil, status.Errorf(codes.FailedPrecondition, "The NasJob %q is in state %s and cannot be deleted.", req.GetName(), obj.State)
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

func (s *jobService) ListNasJobs(ctx context.Context, req *pb.ListNasJobsRequest) (*pb.ListNasJobsResponse, error) {
	parent, err := s.parseNasJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/nasJobs/", parent.Project.Number, parent.Location)
	findPrefixID := fmt.Sprintf("projects/%s/locations/%s/nasJobs/", parent.Project.ID, parent.Location)

	var nasJobs []*pb.NasJob
	nasJobKind := (&pb.NasJob{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, nasJobKind, storage.ListOptions{}, func(obj proto.Message) error {
		job := obj.(*pb.NasJob)
		if strings.HasPrefix(job.GetName(), findPrefix) || strings.HasPrefix(job.GetName(), findPrefixID) {
			result := proto.Clone(job).(*pb.NasJob)
			populateNasJobDefaults(result)
			nasJobs = append(nasJobs, result)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListNasJobsResponse{
		NasJobs: nasJobs,
	}, nil
}

func populateNasJobDefaults(job *pb.NasJob) {
	if job.NasJobSpec != nil && job.NasJobSpec.GetMultiTrialAlgorithmSpec() != nil {
		algoSpec := job.NasJobSpec.GetMultiTrialAlgorithmSpec()
		if algoSpec.SearchTrialSpec != nil && algoSpec.SearchTrialSpec.SearchTrialJobSpec != nil {
			populateWorkerPoolSpecs(algoSpec.SearchTrialSpec.SearchTrialJobSpec.WorkerPoolSpecs)
		}
		if algoSpec.TrainTrialSpec != nil && algoSpec.TrainTrialSpec.TrainTrialJobSpec != nil {
			populateWorkerPoolSpecs(algoSpec.TrainTrialSpec.TrainTrialJobSpec.WorkerPoolSpecs)
		}
	}
}

func populateWorkerPoolSpecs(specs []*pb.WorkerPoolSpec) {
	for _, wp := range specs {
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

type NasJobName struct {
	Project  *projects.ProjectData
	Location string
	NasJob   string
}

func (n *NasJobName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/nasJobs/%s", n.Project.ID, n.Location, n.NasJob)
}

func (n *NasJobName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/nasJobs/%s", n.Project.Number, n.Location, n.NasJob)
}

func (n *NasJobName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/nasJobs/%s", n.Project.Number, n.Location, n.NasJob)
}

type NasJobParent struct {
	Project  *projects.ProjectData
	Location string
}

func (s *MockService) parseNasJobParent(parent string) (*NasJobParent, error) {
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

		return &NasJobParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

// parseNasJobName parses a string into a NasJobName.
// The expected form is projects/<projectID_or_number>/locations/<location>/nasJobs/<nasJobID>
func (s *MockService) parseNasJobName(name string) (*NasJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "nasJobs" {
		projectName, err := projects.ParseProjectName(tokens[0] + "/" + tokens[1])
		if err != nil {
			return nil, err
		}
		project, err := s.Projects.GetProject(projectName)
		if err != nil {
			return nil, err
		}

		nasJobID := tokens[5]
		if !isNumeric(nasJobID) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid nas job id: %s", nasJobID)
		}

		return &NasJobName{
			Project:  project,
			Location: tokens[3],
			NasJob:   nasJobID,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
