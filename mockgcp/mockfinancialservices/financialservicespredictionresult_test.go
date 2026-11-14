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

package mockfinancialservices

import (
	"context"
	"testing"

	pb "cloud.google.com/go/financialservices/apiv1/financialservicespb"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFinancialServicesPredictionResultLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockFinancialServicesService()

	parent := "projects/test-project/locations/us-central1/instances/test-instance"
	predID := "test-pred"
	name := parent + "/predictionResults/" + predID

	// 1. Get before creation -> 404 NotFound
	_, err := server.GetPredictionResult(ctx, &pb.GetPredictionResultRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent prediction result, got: %v", err)
	}

	// 2. Create prediction result with nil payload -> InvalidArgument
	nilCreateReq := &pb.CreatePredictionResultRequest{
		Parent:             parent,
		PredictionResultId: predID,
		PredictionResult:   nil,
	}
	_, err = server.CreatePredictionResult(ctx, nilCreateReq)
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument when creating prediction result with nil payload, got: %v", err)
	}

	// Create prediction result before parent instance exists -> 404 NotFound
	createReq := &pb.CreatePredictionResultRequest{
		Parent:             parent,
		PredictionResultId: predID,
		PredictionResult: &pb.PredictionResult{
			Dataset: parent + "/datasets/test-ds",
			Model:   parent + "/models/test-model",
		},
	}
	_, err = server.CreatePredictionResult(ctx, createReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound when creating prediction result with non-existent parent instance, got: %v", err)
	}

	// Create parent instance first
	_, err = server.CreateInstance(ctx, &pb.CreateInstanceRequest{
		Parent:     "projects/test-project/locations/us-central1",
		InstanceId: "test-instance",
		Instance:   &pb.Instance{},
	})
	if err != nil {
		t.Fatalf("CreateInstance failed: %v", err)
	}

	// 3. Create prediction result after parent instance exists -> Success
	op, err := server.CreatePredictionResult(ctx, createReq)
	if err != nil {
		t.Fatalf("CreatePredictionResult failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from CreatePredictionResult")
	}

	// 4. Get after creation -> Active
	obj, err := server.GetPredictionResult(ctx, &pb.GetPredictionResultRequest{Name: name})
	if err != nil {
		t.Fatalf("GetPredictionResult failed: %v", err)
	}
	if obj.GetName() != name {
		t.Errorf("expected name %q, got %q", name, obj.GetName())
	}
	if obj.GetState() != pb.PredictionResult_ACTIVE {
		t.Errorf("expected state ACTIVE, got %v", obj.GetState())
	}

	// 5. Update with nil payload -> InvalidArgument
	_, err = server.UpdatePredictionResult(ctx, &pb.UpdatePredictionResultRequest{PredictionResult: nil})
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument when updating prediction result with nil payload, got: %v", err)
	}

	// Update non-existent prediction result -> 404 NotFound
	nonExistentUpdateReq := &pb.UpdatePredictionResultRequest{
		PredictionResult: &pb.PredictionResult{
			Name: parent + "/predictionResults/non-existent-pred",
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"labels"}},
	}
	_, err = server.UpdatePredictionResult(ctx, nonExistentUpdateReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound when updating non-existent prediction result, got: %v", err)
	}

	// 6. Update prediction result (field mask update)
	updateReq := &pb.UpdatePredictionResultRequest{
		PredictionResult: &pb.PredictionResult{
			Name: name,
			Labels: map[string]string{
				"env": "test",
			},
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"labels"}},
	}
	op, err = server.UpdatePredictionResult(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdatePredictionResult failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from UpdatePredictionResult")
	}

	// Verify update persisted
	updatedObj, err := server.GetPredictionResult(ctx, &pb.GetPredictionResultRequest{Name: name})
	if err != nil {
		t.Fatalf("GetPredictionResult after update failed: %v", err)
	}
	if updatedObj.GetLabels()["env"] != "test" {
		t.Errorf("expected label env=test, got %v", updatedObj.GetLabels())
	}
	if updatedObj.GetState() != pb.PredictionResult_ACTIVE {
		t.Errorf("expected state ACTIVE preserved, got %v", updatedObj.GetState())
	}

	// Update prediction result without field mask (verifying system fields preserved)
	fullUpdateReq := &pb.UpdatePredictionResultRequest{
		PredictionResult: &pb.PredictionResult{
			Name: name,
			Labels: map[string]string{
				"env": "prod",
			},
		},
	}
	op, err = server.UpdatePredictionResult(ctx, fullUpdateReq)
	if err != nil {
		t.Fatalf("UpdatePredictionResult without field mask failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from UpdatePredictionResult")
	}
	fullUpdatedObj, err := server.GetPredictionResult(ctx, &pb.GetPredictionResultRequest{Name: name})
	if err != nil {
		t.Fatalf("GetPredictionResult after full update failed: %v", err)
	}
	if fullUpdatedObj.GetLabels()["env"] != "prod" {
		t.Errorf("expected label env=prod, got %v", fullUpdatedObj.GetLabels())
	}
	if fullUpdatedObj.GetState() != pb.PredictionResult_ACTIVE {
		t.Errorf("expected state ACTIVE preserved after full update, got %v", fullUpdatedObj.GetState())
	}
	if fullUpdatedObj.GetName() != name {
		t.Errorf("expected name %q preserved, got %q", name, fullUpdatedObj.GetName())
	}

	// 7. Delete prediction result
	op, err = server.DeletePredictionResult(ctx, &pb.DeletePredictionResultRequest{Name: name})
	if err != nil {
		t.Fatalf("DeletePredictionResult failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from DeletePredictionResult")
	}

	// 8. Get after delete -> 404 NotFound
	_, err = server.GetPredictionResult(ctx, &pb.GetPredictionResultRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", err)
	}
}

func TestParsePredictionResultName(t *testing.T) {
	s, _ := setupMockFinancialServicesService()

	validName := "projects/test-project/locations/us-central1/instances/test-instance/predictionResults/test-pred"
	parsed, err := s.parsePredictionResultName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid prediction result name: %v", err)
	}
	if parsed.Instance.Project.ID != "test-project" || parsed.Instance.Location != "us-central1" || parsed.Instance.Instance != "test-instance" || parsed.PredictionResult != "test-pred" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/us-central1/instances/test-instance",
		"projects/test-project/locations/us-central1/instances/test-instance/predictionResults",
		"projects/test-project/locations/us-central1/instances/test-instance/predictionResults/",
		"invalid/test-project/locations/us-central1/instances/test-instance/predictionResults/test-pred",
	}
	for _, name := range invalidNames {
		if _, err := s.parsePredictionResultName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}
