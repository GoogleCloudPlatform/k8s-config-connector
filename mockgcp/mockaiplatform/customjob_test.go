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
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func setupMockJobService() (*MockService, *jobService) {
	store := &mockProjectStore{
		projectsByID: map[string]*projects.ProjectData{
			"test-project": {
				Number: 1234567890,
				ID:     "test-project",
			},
		},
		projectsByNumber: map[string]*projects.ProjectData{
			"1234567890": {
				Number: 1234567890,
				ID:     "test-project",
			},
		},
	}
	env := &common.MockEnvironment{
		Projects: store,
	}
	st := storage.NewInMemoryStorage()
	s := &MockService{
		MockEnvironment: env,
		storage:         st,
		operations:      operations.NewOperationsService(st),
	}
	return s, &jobService{MockService: s}
}

func TestCustomJobLifecycle(t *testing.T) {
	ctx := context.Background()
	_, service := setupMockJobService()

	parent := "projects/test-project/locations/us-central1"

	// 1. Create CustomJob
	createReq := &pb.CreateCustomJobRequest{
		Parent: parent,
		CustomJob: &pb.CustomJob{
			DisplayName: "test-job",
			JobSpec: &pb.CustomJobSpec{
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

	created, err := service.CreateCustomJob(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateCustomJob failed: %v", err)
	}
	if created == nil || created.Name == "" {
		t.Fatalf("expected non-empty created custom job with Name")
	}
	if created.State != pb.JobState_JOB_STATE_PENDING {
		t.Errorf("expected state JOB_STATE_PENDING, got %v", created.State)
	}
	// Verify defaults
	if len(created.JobSpec.WorkerPoolSpecs) != 1 {
		t.Fatalf("expected 1 workerPoolSpec, got %d", len(created.JobSpec.WorkerPoolSpecs))
	}
	wp := created.JobSpec.WorkerPoolSpecs[0]
	if wp.DiskSpec == nil || wp.DiskSpec.BootDiskSizeGb != 100 || wp.DiskSpec.BootDiskType != "pd-ssd" {
		t.Errorf("expected diskSpec defaults (100 GB, pd-ssd), got %+v", wp.DiskSpec)
	}

	// 2. Get CustomJob
	getReq := &pb.GetCustomJobRequest{Name: created.Name}
	getResp, err := service.GetCustomJob(ctx, getReq)
	if err != nil {
		t.Fatalf("GetCustomJob failed: %v", err)
	}
	if getResp.DisplayName != "test-job" {
		t.Errorf("expected displayName 'test-job', got %q", getResp.DisplayName)
	}

	// 3. List CustomJobs
	listReq := &pb.ListCustomJobsRequest{Parent: parent}
	listResp, err := service.ListCustomJobs(ctx, listReq)
	if err != nil {
		t.Fatalf("ListCustomJobs failed: %v", err)
	}
	if len(listResp.CustomJobs) != 1 {
		t.Fatalf("expected 1 job in list, got %d", len(listResp.CustomJobs))
	}

	// 4. Cancel CustomJob
	cancelReq := &pb.CancelCustomJobRequest{Name: created.Name}
	if _, err := service.CancelCustomJob(ctx, cancelReq); err != nil {
		t.Fatalf("CancelCustomJob failed: %v", err)
	}

	// Verify state is CANCELLED
	getAfterCancel, err := service.GetCustomJob(ctx, getReq)
	if err != nil {
		t.Fatalf("GetCustomJob after cancel failed: %v", err)
	}
	if getAfterCancel.State != pb.JobState_JOB_STATE_CANCELLED {
		t.Errorf("expected state JOB_STATE_CANCELLED, got %v", getAfterCancel.State)
	}
	if getAfterCancel.Error == nil || getAfterCancel.Error.Code != int32(codes.Canceled) {
		t.Errorf("expected error code Canceled, got %+v", getAfterCancel.Error)
	}

	// 5. Delete CustomJob
	deleteReq := &pb.DeleteCustomJobRequest{Name: created.Name}
	deleteOp, err := service.DeleteCustomJob(ctx, deleteReq)
	if err != nil {
		t.Fatalf("DeleteCustomJob failed: %v", err)
	}
	if deleteOp == nil {
		t.Fatalf("expected non-nil LRO from DeleteCustomJob")
	}

	// 6. Verify Deleted (NotFound)
	_, err = service.GetCustomJob(ctx, getReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got %v", err)
	}
}

func TestParseCustomJobName(t *testing.T) {
	service, _ := setupMockJobService()

	validName := "projects/test-project/locations/us-central1/customJobs/1234567890"
	parsed, err := service.parseCustomJobName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid name: %v", err)
	}
	if parsed.Location != "us-central1" {
		t.Errorf("expected location us-central1, got %s", parsed.Location)
	}
	if parsed.CustomJob != "1234567890" {
		t.Errorf("expected customJob 1234567890, got %s", parsed.CustomJob)
	}

	// Invalid name (missing parts)
	_, err = service.parseCustomJobName("projects/test-project/locations/us-central1")
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for incomplete name, got %v", err)
	}

	// Non-numeric ID
	_, err = service.parseCustomJobName("projects/test-project/locations/us-central1/customJobs/non-numeric")
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for non-numeric ID, got %v", err)
	}
}
