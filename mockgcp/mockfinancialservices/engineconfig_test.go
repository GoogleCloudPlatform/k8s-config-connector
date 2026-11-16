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
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestEngineConfigLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockFinancialServicesService()

	parent := "projects/test-project/locations/us-central1"
	instanceID := "test-instance"
	instanceName := parent + "/instances/" + instanceID
	engineConfigID := "test-engineconfig"
	name := instanceName + "/engineConfigs/" + engineConfigID

	// 1. Try to create EngineConfig before parent Instance exists -> 404 NotFound
	createReq := &pb.CreateEngineConfigRequest{
		Parent:         instanceName,
		EngineConfigId: engineConfigID,
		EngineConfig: &pb.EngineConfig{
			EngineVersion: "projects/test-project/locations/us-central1/instances/test-instance/engineVersions/aml-retail-v001",
		},
	}
	_, err := server.CreateEngineConfig(ctx, createReq)
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound when parent instance does not exist, got: %v", err)
	}

	// 2. Create parent Instance
	createInstReq := &pb.CreateInstanceRequest{
		Parent:     parent,
		InstanceId: instanceID,
		Instance:   &pb.Instance{},
	}
	if _, err := server.CreateInstance(ctx, createInstReq); err != nil {
		t.Fatalf("CreateInstance failed: %v", err)
	}

	// 3. Get before creation -> 404 NotFound
	_, err = server.GetEngineConfig(ctx, &pb.GetEngineConfigRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent engineConfig, got: %v", err)
	}

	// 4. Create EngineConfig
	op, err := server.CreateEngineConfig(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateEngineConfig failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from CreateEngineConfig")
	}

	// 5. Get after creation -> Active
	ec, err := server.GetEngineConfig(ctx, &pb.GetEngineConfigRequest{Name: name})
	if err != nil {
		t.Fatalf("GetEngineConfig failed: %v", err)
	}
	if ec.GetName() != name {
		t.Errorf("expected name %q, got %q", name, ec.GetName())
	}
	if ec.GetState() != pb.EngineConfig_ACTIVE {
		t.Errorf("expected state ACTIVE, got %v", ec.GetState())
	}

	// 6. Update EngineConfig
	updateReq := &pb.UpdateEngineConfigRequest{
		EngineConfig: &pb.EngineConfig{
			Name:          name,
			EngineVersion: "projects/test-project/locations/us-central1/instances/test-instance/engineVersions/aml-retail-v002",
		},
	}
	op, err = server.UpdateEngineConfig(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateEngineConfig failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from UpdateEngineConfig")
	}

	ec, err = server.GetEngineConfig(ctx, &pb.GetEngineConfigRequest{Name: name})
	if err != nil {
		t.Fatalf("GetEngineConfig after update failed: %v", err)
	}
	if ec.GetEngineVersion() != "projects/test-project/locations/us-central1/instances/test-instance/engineVersions/aml-retail-v002" {
		t.Errorf("unexpected engineVersion after update: %v", ec.GetEngineVersion())
	}

	// 7. Delete EngineConfig
	op, err = server.DeleteEngineConfig(ctx, &pb.DeleteEngineConfigRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteEngineConfig failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from DeleteEngineConfig")
	}
	opFinished, err := server.operations.GetOperation(ctx, &longrunning.GetOperationRequest{Name: op.GetName()})
	if err != nil {
		t.Fatalf("GetOperation failed: %v", err)
	}
	if !opFinished.GetDone() {
		t.Fatalf("expected Operation to be done")
	}
	if opFinished.GetResponse() == nil {
		t.Fatalf("expected non-nil Operation response")
	}
	resp := &emptypb.Empty{}
	if err := opFinished.GetResponse().UnmarshalTo(resp); err != nil {
		t.Fatalf("expected Operation response to unmarshal to Empty, got: %v", err)
	}

	// 8. Get after delete -> 404 NotFound
	_, err = server.GetEngineConfig(ctx, &pb.GetEngineConfigRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", err)
	}
}

func TestParseEngineConfigName(t *testing.T) {
	s, _ := setupMockFinancialServicesService()

	validName := "projects/test-project/locations/us-central1/instances/test-instance/engineConfigs/test-engineconfig"
	parsed, err := s.parseEngineConfigName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid engineConfig name: %v", err)
	}
	if parsed.Instance.Project.ID != "test-project" || parsed.Instance.Location != "us-central1" || parsed.Instance.Instance != "test-instance" || parsed.EngineConfig != "test-engineconfig" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/us-central1/instances/test-instance",
		"projects/test-project/locations/us-central1/instances/test-instance/engineConfigs",
		"invalid/test-project/locations/us-central1/instances/test-instance/engineConfigs/test-engineconfig",
	}
	for _, name := range invalidNames {
		if _, err := s.parseEngineConfigName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}
