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
	"time"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestModelDeploymentMonitoringJobCRUD(t *testing.T) {
	ctx := context.Background()
	_, service := setupMockJobService()

	parent := "projects/test-project/locations/us-central1"
	endpoint := "projects/test-project/locations/us-central1/endpoints/12345678"

	// 1. Create ModelDeploymentMonitoringJob
	createReq := &pb.CreateModelDeploymentMonitoringJobRequest{
		Parent: parent,
		ModelDeploymentMonitoringJob: &pb.ModelDeploymentMonitoringJob{
			DisplayName: "test-monitoring-job",
			Endpoint:    endpoint,
			ModelDeploymentMonitoringObjectiveConfigs: []*pb.ModelDeploymentMonitoringObjectiveConfig{
				{
					DeployedModelId: "deployed-model-1",
					ObjectiveConfig: &pb.ModelMonitoringObjectiveConfig{
						PredictionDriftDetectionConfig: &pb.ModelMonitoringObjectiveConfig_PredictionDriftDetectionConfig{
							DefaultDriftThreshold: &pb.ThresholdConfig{
								Threshold: &pb.ThresholdConfig_Value{Value: 0.1},
							},
						},
					},
				},
			},
			ModelDeploymentMonitoringScheduleConfig: &pb.ModelDeploymentMonitoringScheduleConfig{
				MonitorInterval: durationpb.New(3600 * time.Second),
			},
			LoggingSamplingStrategy: &pb.SamplingStrategy{
				RandomSampleConfig: &pb.SamplingStrategy_RandomSampleConfig{
					SampleRate: 0.5,
				},
			},
			Labels: map[string]string{
				"env": "test",
			},
		},
	}

	created, err := service.CreateModelDeploymentMonitoringJob(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateModelDeploymentMonitoringJob failed: %v", err)
	}
	if created.GetName() == "" {
		t.Fatalf("expected non-empty name, got %q", created.GetName())
	}
	if created.GetState() != pb.JobState_JOB_STATE_RUNNING {
		t.Errorf("expected state %v, got %v", pb.JobState_JOB_STATE_RUNNING, created.GetState())
	}
	if created.GetScheduleState() != pb.ModelDeploymentMonitoringJob_OFFLINE {
		t.Errorf("expected schedule state %v, got %v", pb.ModelDeploymentMonitoringJob_OFFLINE, created.GetScheduleState())
	}
	if len(created.GetBigqueryTables()) == 0 {
		t.Errorf("expected populated BigqueryTables, got empty")
	}

	// 2. Get ModelDeploymentMonitoringJob
	getReq := &pb.GetModelDeploymentMonitoringJobRequest{
		Name: created.GetName(),
	}
	got, err := service.GetModelDeploymentMonitoringJob(ctx, getReq)
	if err != nil {
		t.Fatalf("GetModelDeploymentMonitoringJob failed: %v", err)
	}
	if got.GetDisplayName() != "test-monitoring-job" {
		t.Errorf("expected displayName %q, got %q", "test-monitoring-job", got.GetDisplayName())
	}

	// 3. Update ModelDeploymentMonitoringJob
	updateReq := &pb.UpdateModelDeploymentMonitoringJobRequest{
		ModelDeploymentMonitoringJob: &pb.ModelDeploymentMonitoringJob{
			Name:        created.GetName(),
			DisplayName: "updated-monitoring-job",
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"display_name"},
		},
	}
	op, err := service.UpdateModelDeploymentMonitoringJob(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateModelDeploymentMonitoringJob failed: %v", err)
	}
	if !op.GetDone() {
		t.Errorf("expected operation to be done")
	}

	gotAfterUpdate, err := service.GetModelDeploymentMonitoringJob(ctx, getReq)
	if err != nil {
		t.Fatalf("GetModelDeploymentMonitoringJob after update failed: %v", err)
	}
	if gotAfterUpdate.GetDisplayName() != "updated-monitoring-job" {
		t.Errorf("expected updated displayName %q, got %q", "updated-monitoring-job", gotAfterUpdate.GetDisplayName())
	}

	// 4. Pause & Resume
	_, err = service.PauseModelDeploymentMonitoringJob(ctx, &pb.PauseModelDeploymentMonitoringJobRequest{Name: created.GetName()})
	if err != nil {
		t.Fatalf("PauseModelDeploymentMonitoringJob failed: %v", err)
	}
	paused, err := service.GetModelDeploymentMonitoringJob(ctx, getReq)
	if err != nil {
		t.Fatalf("GetModelDeploymentMonitoringJob after pause failed: %v", err)
	}
	if paused.GetState() != pb.JobState_JOB_STATE_PAUSED {
		t.Errorf("expected state PAUSED, got %v", paused.GetState())
	}

	_, err = service.ResumeModelDeploymentMonitoringJob(ctx, &pb.ResumeModelDeploymentMonitoringJobRequest{Name: created.GetName()})
	if err != nil {
		t.Fatalf("ResumeModelDeploymentMonitoringJob failed: %v", err)
	}
	resumed, err := service.GetModelDeploymentMonitoringJob(ctx, getReq)
	if err != nil {
		t.Fatalf("GetModelDeploymentMonitoringJob after resume failed: %v", err)
	}
	if resumed.GetState() != pb.JobState_JOB_STATE_RUNNING {
		t.Errorf("expected state RUNNING, got %v", resumed.GetState())
	}

	// 5. List ModelDeploymentMonitoringJobs
	listResp, err := service.ListModelDeploymentMonitoringJobs(ctx, &pb.ListModelDeploymentMonitoringJobsRequest{
		Parent: parent,
	})
	if err != nil {
		t.Fatalf("ListModelDeploymentMonitoringJobs failed: %v", err)
	}
	if len(listResp.GetModelDeploymentMonitoringJobs()) == 0 {
		t.Errorf("expected non-empty jobs list")
	}

	// 6. Delete ModelDeploymentMonitoringJob
	deleteOp, err := service.DeleteModelDeploymentMonitoringJob(ctx, &pb.DeleteModelDeploymentMonitoringJobRequest{
		Name: created.GetName(),
	})
	if err != nil {
		t.Fatalf("DeleteModelDeploymentMonitoringJob failed: %v", err)
	}
	if !deleteOp.GetDone() {
		t.Errorf("expected delete operation to be done")
	}

	// Verify not found after delete
	_, err = service.GetModelDeploymentMonitoringJob(ctx, getReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound error after delete, got %v", err)
	}
}
