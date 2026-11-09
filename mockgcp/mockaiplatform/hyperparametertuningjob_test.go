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
	"testing"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestHyperparameterTuningJobLifecycle(t *testing.T) {
	ctx := context.Background()
	_, service := setupMockJobService()

	parent := "projects/test-project/locations/us-central1"

	// 1. Create HyperparameterTuningJob (minimal)
	createReqMinimal := &pb.CreateHyperparameterTuningJobRequest{
		Parent: parent,
		HyperparameterTuningJob: &pb.HyperparameterTuningJob{
			DisplayName:        "minimal-hptjob-test",
			MaxTrialCount:      1,
			ParallelTrialCount: 1,
			StudySpec: &pb.StudySpec{
				Metrics: []*pb.StudySpec_MetricSpec{
					{
						MetricId: "val_loss",
						Goal:     pb.StudySpec_MetricSpec_MINIMIZE,
					},
				},
				Parameters: []*pb.StudySpec_ParameterSpec{
					{
						ParameterId: "learning_rate",
						ParameterValueSpec: &pb.StudySpec_ParameterSpec_DoubleValueSpec_{
							DoubleValueSpec: &pb.StudySpec_ParameterSpec_DoubleValueSpec{
								MinValue: 0.001,
								MaxValue: 0.1,
							},
						},
					},
				},
			},
			TrialJobSpec: &pb.CustomJobSpec{
				WorkerPoolSpecs: []*pb.WorkerPoolSpec{
					{
						MachineSpec: &pb.MachineSpec{
							MachineType: "n1-standard-4",
						},
						ReplicaCount: 1,
						Task: &pb.WorkerPoolSpec_ContainerSpec{
							ContainerSpec: &pb.ContainerSpec{
								ImageUri: "alpine:latest",
								Command:  []string{"sh", "-c", "echo hello world"},
							},
						},
					},
				},
			},
		},
	}

	createdMinimal, err := service.CreateHyperparameterTuningJob(ctx, createReqMinimal)
	if err != nil {
		t.Fatalf("CreateHyperparameterTuningJob (minimal) failed: %v", err)
	}
	if createdMinimal == nil || createdMinimal.Name == "" {
		t.Fatalf("expected non-empty created job with Name")
	}
	if createdMinimal.Name != "projects/1234567890/locations/us-central1/hyperparameterTuningJobs/4195856845425868800" {
		t.Errorf("expected minimal ID 4195856845425868800, got %s", createdMinimal.Name)
	}
	if createdMinimal.State != pb.JobState_JOB_STATE_PENDING {
		t.Errorf("expected state JOB_STATE_PENDING, got %v", createdMinimal.State)
	}
	// Verify defaults
	if len(createdMinimal.TrialJobSpec.WorkerPoolSpecs) != 1 {
		t.Fatalf("expected 1 workerPoolSpec, got %d", len(createdMinimal.TrialJobSpec.WorkerPoolSpecs))
	}
	wp := createdMinimal.TrialJobSpec.WorkerPoolSpecs[0]
	if wp.DiskSpec == nil || wp.DiskSpec.BootDiskSizeGb != 100 || wp.DiskSpec.BootDiskType != "pd-ssd" {
		t.Errorf("expected diskSpec defaults (100 GB, pd-ssd), got %+v", wp.DiskSpec)
	}

	// 2. Get HyperparameterTuningJob
	getReq := &pb.GetHyperparameterTuningJobRequest{Name: createdMinimal.Name}
	getResp, err := service.GetHyperparameterTuningJob(ctx, getReq)
	if err != nil {
		t.Fatalf("GetHyperparameterTuningJob failed: %v", err)
	}
	if getResp.DisplayName != "minimal-hptjob-test" {
		t.Errorf("expected displayName 'minimal-hptjob-test', got %q", getResp.DisplayName)
	}
	if getResp.StartTime == nil {
		t.Errorf("expected startTime to be populated on Get")
	}

	// 3. List HyperparameterTuningJobs
	listReq := &pb.ListHyperparameterTuningJobsRequest{Parent: parent}
	listResp, err := service.ListHyperparameterTuningJobs(ctx, listReq)
	if err != nil {
		t.Fatalf("ListHyperparameterTuningJobs failed: %v", err)
	}
	if len(listResp.HyperparameterTuningJobs) != 1 {
		t.Fatalf("expected 1 job in list, got %d", len(listResp.HyperparameterTuningJobs))
	}

	// 4. Cancel HyperparameterTuningJob
	cancelReq := &pb.CancelHyperparameterTuningJobRequest{Name: createdMinimal.Name}
	if _, err := service.CancelHyperparameterTuningJob(ctx, cancelReq); err != nil {
		t.Fatalf("CancelHyperparameterTuningJob failed: %v", err)
	}

	// Cancelling again should return FailedPrecondition
	if _, err := service.CancelHyperparameterTuningJob(ctx, cancelReq); status.Code(err) != codes.FailedPrecondition {
		t.Fatalf("expected FailedPrecondition when cancelling already cancelled job, got %v", err)
	}

	// Verify state is CANCELLED
	getAfterCancel, err := service.GetHyperparameterTuningJob(ctx, getReq)
	if err != nil {
		t.Fatalf("GetHyperparameterTuningJob after cancel failed: %v", err)
	}
	if getAfterCancel.State != pb.JobState_JOB_STATE_CANCELLED {
		t.Errorf("expected state JOB_STATE_CANCELLED, got %v", getAfterCancel.State)
	}
	if getAfterCancel.Error == nil || getAfterCancel.Error.Code != int32(codes.Canceled) {
		t.Errorf("expected error code Canceled, got %+v", getAfterCancel.Error)
	}

	// 5. Delete HyperparameterTuningJob
	deleteReq := &pb.DeleteHyperparameterTuningJobRequest{Name: createdMinimal.Name}
	deleteOp, err := service.DeleteHyperparameterTuningJob(ctx, deleteReq)
	if err != nil {
		t.Fatalf("DeleteHyperparameterTuningJob failed: %v", err)
	}
	if deleteOp == nil {
		t.Fatalf("expected non-nil LRO from DeleteHyperparameterTuningJob")
	}

	// 6. Verify Deleted (NotFound)
	_, err = service.GetHyperparameterTuningJob(ctx, getReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got %v", err)
	}

	// 7. Create HyperparameterTuningJob (maximal)
	createReqMaximal := &pb.CreateHyperparameterTuningJobRequest{
		Parent: parent,
		HyperparameterTuningJob: &pb.HyperparameterTuningJob{
			DisplayName: "maximal-hptjob-test",
		},
	}
	createdMaximal, err := service.CreateHyperparameterTuningJob(ctx, createReqMaximal)
	if err != nil {
		t.Fatalf("CreateHyperparameterTuningJob (maximal) failed: %v", err)
	}
	if createdMaximal.Name != "projects/1234567890/locations/us-central1/hyperparameterTuningJobs/3529700133551734784" {
		t.Errorf("expected maximal ID 3529700133551734784, got %s", createdMaximal.Name)
	}
}

func TestParseHyperparameterTuningJobName(t *testing.T) {
	service, _ := setupMockJobService()

	validName := "projects/test-project/locations/us-central1/hyperparameterTuningJobs/1234567890"
	parsed, err := service.parseHyperparameterTuningJobName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid name: %v", err)
	}
	if parsed.Location != "us-central1" {
		t.Errorf("expected location us-central1, got %s", parsed.Location)
	}
	if parsed.HyperparameterTuningJob != "1234567890" {
		t.Errorf("expected hyperparameterTuningJob 1234567890, got %s", parsed.HyperparameterTuningJob)
	}

	// Invalid name (missing parts)
	_, err = service.parseHyperparameterTuningJobName("projects/test-project/locations/us-central1")
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for incomplete name, got %v", err)
	}

	// Non-numeric ID
	_, err = service.parseHyperparameterTuningJobName("projects/test-project/locations/us-central1/hyperparameterTuningJobs/non-numeric")
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for non-numeric ID, got %v", err)
	}
}
