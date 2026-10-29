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

package mocksaasservicemgmt

import (
	"context"
	"testing"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRolloutLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockRolloutsService()

	parent := "projects/test-project/locations/us-central1"
	rolloutID := "test-rollout"
	expectedName := parent + "/rollouts/" + rolloutID

	// 1. Get before create -> NotFound
	_, err := server.GetRollout(ctx, &pb.GetRolloutRequest{Name: expectedName})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got %v", err)
	}

	// 2. Create Rollout
	created, err := server.CreateRollout(ctx, &pb.CreateRolloutRequest{
		Parent:    parent,
		RolloutId: rolloutID,
		Rollout: &pb.Rollout{
			Release:                      parent + "/releases/test-release",
			RolloutKind:                  parent + "/rolloutKinds/test-rollout-kind",
			RolloutOrchestrationStrategy: "Google.Cloud.Simple.AllAtOnce",
			UnitFilter:                   "true",
			Labels: map[string]string{
				"env": "test",
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateRollout failed: %v", err)
	}

	if created.Name != expectedName {
		t.Errorf("expected name %q, got %q", expectedName, created.Name)
	}
	if created.EffectiveUnitFilter != "true" {
		t.Errorf("expected EffectiveUnitFilter 'true', got %q", created.EffectiveUnitFilter)
	}
	if created.Control == nil || created.Control.Action != pb.RolloutAction_ROLLOUT_ACTION_RUN {
		t.Errorf("expected default control action ROLLOUT_ACTION_RUN, got %+v", created.Control)
	}
	if created.CreateTime != nil || created.UpdateTime != nil || created.Etag != "" {
		t.Errorf("CreateRollout response should not have CreateTime/UpdateTime/Etag")
	}

	// 3. Get Rollout after create
	got, err := server.GetRollout(ctx, &pb.GetRolloutRequest{Name: expectedName})
	if err != nil {
		t.Fatalf("GetRollout failed: %v", err)
	}
	if got.CreateTime == nil || got.UpdateTime == nil || got.Etag == "" {
		t.Errorf("GetRollout response should have CreateTime/UpdateTime/Etag")
	}

	// 4. Update Rollout
	updated, err := server.UpdateRollout(ctx, &pb.UpdateRolloutRequest{
		Rollout: &pb.Rollout{
			Name:                         expectedName,
			RolloutOrchestrationStrategy: "Google.Cloud.Simple.OneLocationAtATime",
			Control: &pb.RolloutControl{
				Action: pb.RolloutAction_ROLLOUT_ACTION_RUN,
				ActionParams: &pb.RolloutControl_RunParams{
					RunParams: &pb.RolloutControl_RunRolloutActionParams{
						RetryFailedOperations: true,
					},
				},
			},
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"rollout_orchestration_strategy", "control"},
		},
	})
	if err != nil {
		t.Fatalf("UpdateRollout failed: %v", err)
	}
	if updated.RolloutOrchestrationStrategy != "Google.Cloud.Simple.OneLocationAtATime" {
		t.Errorf("expected updated strategy 'Google.Cloud.Simple.OneLocationAtATime', got %q", updated.RolloutOrchestrationStrategy)
	}
	if updated.RootRollout != "" || updated.State != pb.Rollout_ROLLOUT_STATE_UNSPECIFIED {
		t.Errorf("UpdateRollout response should strip rootRollout and state")
	}

	// 5. Get after update has state and rootRollout
	gotAfterUpdate, err := server.GetRollout(ctx, &pb.GetRolloutRequest{Name: expectedName})
	if err != nil {
		t.Fatalf("GetRollout after update failed: %v", err)
	}
	if gotAfterUpdate.RootRollout != expectedName {
		t.Errorf("expected rootRollout %q, got %q", expectedName, gotAfterUpdate.RootRollout)
	}
	if gotAfterUpdate.State != pb.Rollout_ROLLOUT_STATE_RUNNING {
		t.Errorf("expected state RUNNING, got %v", gotAfterUpdate.State)
	}

	// 6. List Rollouts
	listResp, err := server.ListRollouts(ctx, &pb.ListRolloutsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListRollouts failed: %v", err)
	}
	if len(listResp.Rollouts) != 1 {
		t.Fatalf("expected 1 rollout, got %d", len(listResp.Rollouts))
	}

	// 7. Delete Rollout
	_, err = server.DeleteRollout(ctx, &pb.DeleteRolloutRequest{Name: expectedName})
	if err != nil {
		t.Fatalf("DeleteRollout failed: %v", err)
	}

	// 8. Get after delete -> NotFound
	_, err = server.GetRollout(ctx, &pb.GetRolloutRequest{Name: expectedName})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got %v", err)
	}
}
