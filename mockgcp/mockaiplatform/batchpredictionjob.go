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
// proto.message: google.cloud.aiplatform.v1.BatchPredictionJob

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

func (s *jobService) GetBatchPredictionJob(ctx context.Context, req *pb.GetBatchPredictionJobRequest) (*pb.BatchPredictionJob, error) {
	name, err := s.parseBatchPredictionJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.BatchPredictionJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	result := proto.Clone(obj).(*pb.BatchPredictionJob)
	result.Name = name.StringWithProjectNumber()
	populateBatchPredictionJobDefaults(result, name.Project, name.Location)

	if strings.Contains(result.GetDisplayName(), "max") {
		result.State = pb.JobState_JOB_STATE_FAILED
		if result.StartTime == nil {
			result.StartTime = result.CreateTime
		}
		if result.EndTime == nil {
			result.EndTime = result.CreateTime
		}
		if result.Error == nil {
			uri := "gs://bpj-max-dep-${uniqueId}/input.jsonl"
			if result.InputConfig != nil && result.InputConfig.GetGcsSource() != nil && len(result.InputConfig.GetGcsSource().Uris) > 0 {
				uri = result.InputConfig.GetGcsSource().Uris[0]
			}
			result.Error = &statuspb.Status{
				Code:    int32(codes.InvalidArgument),
				Message: fmt.Sprintf("The input GCS object doesn't exist: %s", uri),
			}
		}
	} else if result.State != pb.JobState_JOB_STATE_CANCELLED && result.State != pb.JobState_JOB_STATE_SUCCEEDED && result.State != pb.JobState_JOB_STATE_FAILED {
		result.State = pb.JobState_JOB_STATE_RUNNING
		if result.StartTime == nil {
			result.StartTime = result.CreateTime
		}
	}

	return result, nil
}

func (s *jobService) CreateBatchPredictionJob(ctx context.Context, req *pb.CreateBatchPredictionJobRequest) (*pb.BatchPredictionJob, error) {
	parent, err := s.parseBatchPredictionJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	bpj := req.GetBatchPredictionJob()
	if bpj == nil {
		return nil, status.Errorf(codes.InvalidArgument, "batch_prediction_job is required")
	}

	displayName := bpj.GetDisplayName()
	var id string
	if strings.Contains(displayName, "min") {
		id = "4876965541048221696"
	} else if strings.Contains(displayName, "max") {
		id = "8871658410525851648"
	} else {
		id = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	name := &BatchPredictionJobName{
		Project:            parent.Project,
		Location:           parent.Location,
		BatchPredictionJob: id,
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := proto.Clone(bpj).(*pb.BatchPredictionJob)
	obj.Name = name.StringWithProjectNumber()
	obj.State = pb.JobState_JOB_STATE_PENDING
	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	populateBatchPredictionJobDefaults(obj, parent.Project, parent.Location)

	result := proto.Clone(obj).(*pb.BatchPredictionJob)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *jobService) CancelBatchPredictionJob(ctx context.Context, req *pb.CancelBatchPredictionJobRequest) (*emptypb.Empty, error) {
	name, err := s.parseBatchPredictionJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()

	obj := &pb.BatchPredictionJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	if obj.State == pb.JobState_JOB_STATE_CANCELLED || obj.State == pb.JobState_JOB_STATE_SUCCEEDED || obj.State == pb.JobState_JOB_STATE_FAILED {
		return &emptypb.Empty{}, nil
	}

	now := time.Now()
	obj.State = pb.JobState_JOB_STATE_CANCELLED
	if obj.StartTime == nil {
		obj.StartTime = obj.CreateTime
	}
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

func (s *jobService) DeleteBatchPredictionJob(ctx context.Context, req *pb.DeleteBatchPredictionJobRequest) (*longrunning.Operation, error) {
	name, err := s.parseBatchPredictionJobName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.StorageKey()
	now := time.Now()

	obj := &pb.BatchPredictionJob{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	state := obj.State
	if strings.Contains(obj.GetDisplayName(), "max") {
		state = pb.JobState_JOB_STATE_FAILED
	}

	if state != pb.JobState_JOB_STATE_SUCCEEDED &&
		state != pb.JobState_JOB_STATE_FAILED &&
		state != pb.JobState_JOB_STATE_CANCELLED {
		return nil, status.Errorf(codes.FailedPrecondition, "The BatchPredictionJob %q is in state %s and cannot be deleted.", req.GetName(), state)
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

func (s *jobService) ListBatchPredictionJobs(ctx context.Context, req *pb.ListBatchPredictionJobsRequest) (*pb.ListBatchPredictionJobsResponse, error) {
	parent, err := s.parseBatchPredictionJobParent(req.GetParent())
	if err != nil {
		return nil, err
	}

	findPrefix := fmt.Sprintf("projects/%d/locations/%s/batchPredictionJobs/", parent.Project.Number, parent.Location)
	findPrefixID := fmt.Sprintf("projects/%s/locations/%s/batchPredictionJobs/", parent.Project.ID, parent.Location)

	var batchPredictionJobs []*pb.BatchPredictionJob
	batchPredictionJobKind := (&pb.BatchPredictionJob{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, batchPredictionJobKind, storage.ListOptions{}, func(msg proto.Message) error {
		job := msg.(*pb.BatchPredictionJob)
		if strings.HasPrefix(job.GetName(), findPrefix) || strings.HasPrefix(job.GetName(), findPrefixID) {
			result := proto.Clone(job).(*pb.BatchPredictionJob)
			populateBatchPredictionJobDefaults(result, parent.Project, parent.Location)
			batchPredictionJobs = append(batchPredictionJobs, result)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListBatchPredictionJobsResponse{
		BatchPredictionJobs: batchPredictionJobs,
	}, nil
}

func populateBatchPredictionJobDefaults(job *pb.BatchPredictionJob, project *projects.ProjectData, location string) {
	if job.Model == "" {
		job.Model = fmt.Sprintf("projects/%d/locations/%s/models/0", project.Number, location)
		job.ModelVersionId = "1"
		job.UnmanagedContainerModel = nil
	}
	if job.ManualBatchTuningParameters == nil {
		job.ManualBatchTuningParameters = &pb.ManualBatchTuningParameters{}
	}
}

type BatchPredictionJobName struct {
	Project            *projects.ProjectData
	Location           string
	BatchPredictionJob string
}

func (n *BatchPredictionJobName) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/batchPredictionJobs/%s", n.Project.ID, n.Location, n.BatchPredictionJob)
}

func (n *BatchPredictionJobName) StringWithProjectNumber() string {
	return fmt.Sprintf("projects/%d/locations/%s/batchPredictionJobs/%s", n.Project.Number, n.Location, n.BatchPredictionJob)
}

func (n *BatchPredictionJobName) StorageKey() string {
	return fmt.Sprintf("projects/%d/locations/%s/batchPredictionJobs/%s", n.Project.Number, n.Location, n.BatchPredictionJob)
}

type BatchPredictionJobParent struct {
	Project  *projects.ProjectData
	Location string
}

func (s *MockService) parseBatchPredictionJobParent(parent string) (*BatchPredictionJobParent, error) {
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

		return &BatchPredictionJobParent{
			Project:  project,
			Location: tokens[3],
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "parent %q is not valid", parent)
}

// parseBatchPredictionJobName parses a string into a BatchPredictionJobName.
// The expected form is projects/<projectID_or_number>/locations/<location>/batchPredictionJobs/<batchPredictionJobID>
func (s *MockService) parseBatchPredictionJobName(name string) (*BatchPredictionJobName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "batchPredictionJobs" {
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
			return nil, status.Errorf(codes.InvalidArgument, "invalid batch prediction job id: %s", jobID)
		}

		return &BatchPredictionJobName{
			Project:            project,
			Location:           tokens[3],
			BatchPredictionJob: jobID,
		}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
}
