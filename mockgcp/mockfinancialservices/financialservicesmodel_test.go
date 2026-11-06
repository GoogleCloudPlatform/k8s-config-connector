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

	// 2. Create model
	createReq := &pb.CreateModelRequest{
		Parent:  parent,
		ModelId: modelID,
		Model: &pb.Model{
			EngineConfig:   "projects/test-project/locations/us-central1/instances/test-instance/engineConfigs/test-ec",
			PrimaryDataset: "projects/test-project/locations/us-central1/instances/test-instance/datasets/test-ds",
		},
	}
	op, err := server.CreateModel(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateModel failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from CreateModel")
	}

	// 3. Get after creation -> Active
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

	// 4. Update model
	updateReq := &pb.UpdateModelRequest{
		Model: &pb.Model{
			Name:    name,
			EndTime: nil,
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

	// 5. List models
	listResp, err := server.ListModels(ctx, &pb.ListModelsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListModels failed: %v", err)
	}
	if len(listResp.GetModels()) != 1 {
		t.Fatalf("expected 1 model, got %d", len(listResp.GetModels()))
	}

	// 6. Delete model
	op, err = server.DeleteModel(ctx, &pb.DeleteModelRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteModel failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from DeleteModel")
	}

	// 7. Get after delete -> 404 NotFound
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
		"invalid/test-project/locations/us-central1/instances/test-instance/models/test-model",
	}
	for _, name := range invalidNames {
		if _, err := s.parseModelName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}
