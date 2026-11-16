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

func TestFinancialServicesModelLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockFinancialServicesService()

	parent := "projects/test-project/locations/us-central1/instances/test-instance"
	modelID := "test-model"
	name := parent + "/models/" + modelID

	// 1. Get before creation -> 404 NotFound
	_, err := server.GetModel(ctx, &pb.GetModelRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent model, got: %v", err)
	}

	// 2. Create model with nil Model -> InvalidArgument
	nilModelCreateReq := &pb.CreateModelRequest{
		Parent:  parent,
		ModelId: modelID,
		Model:   nil,
	}
	_, err = server.CreateModel(ctx, nilModelCreateReq)
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument when creating model with nil Model, got: %v", err)
	}

	// Create model before parent instance exists -> 404 NotFound
	createReq := &pb.CreateModelRequest{
		Parent:  parent,
		ModelId: modelID,
		Model: &pb.Model{
			EngineConfig:   "projects/test-project/locations/us-central1/instances/test-instance/engineConfigs/test-ec",
			PrimaryDataset: "projects/test-project/locations/us-central1/instances/test-instance/datasets/test-ds",
		},
	}
	_, err = server.CreateModel(ctx, createReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound when creating model with non-existent parent instance, got: %v", err)
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

	// 3. Create model after parent instance exists -> Success
	op, err := server.CreateModel(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateModel failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from CreateModel")
	}

	// 4. Get after creation -> Active
	model, err := server.GetModel(ctx, &pb.GetModelRequest{Name: name})
	if err != nil {
		t.Fatalf("GetModel failed: %v", err)
	}
	if model.GetName() != name {
		t.Errorf("expected name %q, got %q", name, model.GetName())
	}
	if model.GetState() != pb.Model_ACTIVE {
		t.Errorf("expected state ACTIVE, got %v", model.GetState())
	}

	// 5. Update with nil Model -> InvalidArgument
	_, err = server.UpdateModel(ctx, &pb.UpdateModelRequest{Model: nil})
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument when updating model with nil Model, got: %v", err)
	}

	// Update non-existent model -> 404 NotFound
	nonExistentUpdateReq := &pb.UpdateModelRequest{
		Model: &pb.Model{
			Name: parent + "/models/non-existent-model",
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"labels"}},
	}
	_, err = server.UpdateModel(ctx, nonExistentUpdateReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound when updating non-existent model, got: %v", err)
	}

	// 6. Update model (field mask update)
	updateReq := &pb.UpdateModelRequest{
		Model: &pb.Model{
			Name: name,
			Labels: map[string]string{
				"env": "test",
			},
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"labels"}},
	}
	op, err = server.UpdateModel(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateModel failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from UpdateModel")
	}

	// Verify update persisted
	updatedModel, err := server.GetModel(ctx, &pb.GetModelRequest{Name: name})
	if err != nil {
		t.Fatalf("GetModel after update failed: %v", err)
	}
	if updatedModel.GetLabels()["env"] != "test" {
		t.Errorf("expected label env=test, got %v", updatedModel.GetLabels())
	}
	if updatedModel.GetState() != pb.Model_ACTIVE {
		t.Errorf("expected state ACTIVE preserved, got %v", updatedModel.GetState())
	}

	// Update model without field mask (verifying system fields preserved)
	fullUpdateReq := &pb.UpdateModelRequest{
		Model: &pb.Model{
			Name: name,
			Labels: map[string]string{
				"env": "prod",
			},
		},
	}
	op, err = server.UpdateModel(ctx, fullUpdateReq)
	if err != nil {
		t.Fatalf("UpdateModel without field mask failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from UpdateModel")
	}
	fullUpdatedModel, err := server.GetModel(ctx, &pb.GetModelRequest{Name: name})
	if err != nil {
		t.Fatalf("GetModel after full update failed: %v", err)
	}
	if fullUpdatedModel.GetLabels()["env"] != "prod" {
		t.Errorf("expected label env=prod, got %v", fullUpdatedModel.GetLabels())
	}
	if fullUpdatedModel.GetState() != pb.Model_ACTIVE {
		t.Errorf("expected state ACTIVE preserved after full update, got %v", fullUpdatedModel.GetState())
	}
	if fullUpdatedModel.GetName() != name {
		t.Errorf("expected name %q preserved, got %q", name, fullUpdatedModel.GetName())
	}

	// 7. List models with invalid parent -> InvalidArgument
	_, err = server.ListModels(ctx, &pb.ListModelsRequest{Parent: "invalid-parent"})
	if status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument for invalid parent, got: %v", err)
	}

	// 8. List models
	listResp, err := server.ListModels(ctx, &pb.ListModelsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListModels failed: %v", err)
	}
	if len(listResp.GetModels()) != 1 {
		t.Fatalf("expected 1 model, got %d", len(listResp.GetModels()))
	}

	// 9. Delete model
	op, err = server.DeleteModel(ctx, &pb.DeleteModelRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteModel failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from DeleteModel")
	}

	// 10. Get after delete -> 404 NotFound
	_, err = server.GetModel(ctx, &pb.GetModelRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", err)
	}
}

func TestParseModelName(t *testing.T) {
	s, _ := setupMockFinancialServicesService()

	validName := "projects/test-project/locations/us-central1/instances/test-instance/models/test-model"
	parsed, err := s.parseModelName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid model name: %v", err)
	}
	if parsed.Instance.Project.ID != "test-project" || parsed.Instance.Location != "us-central1" || parsed.Instance.Instance != "test-instance" || parsed.Model != "test-model" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/us-central1/instances/test-instance",
		"projects/test-project/locations/us-central1/instances/test-instance/models",
		"projects/test-project/locations/us-central1/instances/test-instance/models/",
		"invalid/test-project/locations/us-central1/instances/test-instance/models/test-model",
	}
	for _, name := range invalidNames {
		if _, err := s.parseModelName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}
