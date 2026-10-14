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

package mocknetworkservices

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockProjectStore struct {
	projectsByID     map[string]*projects.ProjectData
	projectsByNumber map[string]*projects.ProjectData
}

func (m *mockProjectStore) GetProject(projectName *projects.ProjectName) (*projects.ProjectData, error) {
	return m.GetProjectByIDOrNumber(projectName.OriginalValue)
}

func (m *mockProjectStore) GetProjectByID(projectID string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByID[projectID]; ok {
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", projectID)
}

func (m *mockProjectStore) GetProjectByNumber(projectNumberAsString string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByNumber[projectNumberAsString]; ok {
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", projectNumberAsString)
}

func (m *mockProjectStore) GetProjectByIDOrNumber(projectIDOrNumber string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByID[projectIDOrNumber]; ok {
		return p, nil
	}
	if p, ok := m.projectsByNumber[projectIDOrNumber]; ok {
		return p, nil
	}
	if n, err := strconv.ParseInt(projectIDOrNumber, 10, 64); err == nil {
		p := &projects.ProjectData{
			Number: n,
			ID:     "project-" + projectIDOrNumber,
		}
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", projectIDOrNumber)
}

func setupMockNetworkServicesService() (*MockService, *NetworkServicesServer) {
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
	server := &NetworkServicesServer{MockService: s}
	s.v1 = server
	return s, server
}

func TestServiceLbPolicyLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockNetworkServicesService()

	parent := "projects/test-project/locations/global"
	policyID := "test-policy"
	name := parent + "/serviceLbPolicies/" + policyID

	// 1. Get before creation -> 404 NotFound
	_, err := server.GetServiceLbPolicy(ctx, &pb.GetServiceLbPolicyRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent policy, got: %v", err)
	}

	// 2. Create policy
	createReq := &pb.CreateServiceLbPolicyRequest{
		Parent:            parent,
		ServiceLbPolicyId: policyID,
		ServiceLbPolicy: &pb.ServiceLbPolicy{
			Description:            "Initial description",
			LoadBalancingAlgorithm: pb.ServiceLbPolicy_WATERFALL_BY_REGION,
			Labels: map[string]string{
				"env": "test",
			},
		},
	}
	op, err := server.CreateServiceLbPolicy(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateServiceLbPolicy failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from CreateServiceLbPolicy")
	}

	// 3. Get after creation
	policy, err := server.GetServiceLbPolicy(ctx, &pb.GetServiceLbPolicyRequest{Name: name})
	if err != nil {
		t.Fatalf("GetServiceLbPolicy failed: %v", err)
	}
	if policy.GetName() != name {
		t.Errorf("expected name %q, got %q", name, policy.GetName())
	}
	if policy.GetDescription() != "Initial description" {
		t.Errorf("expected description 'Initial description', got %q", policy.GetDescription())
	}
	if policy.GetLoadBalancingAlgorithm() != pb.ServiceLbPolicy_WATERFALL_BY_REGION {
		t.Errorf("expected algorithm WATERFALL_BY_REGION, got %v", policy.GetLoadBalancingAlgorithm())
	}
	if policy.GetLabels()["env"] != "test" {
		t.Errorf("expected label env=test, got %v", policy.GetLabels())
	}

	// 4. Update policy
	updateReq := &pb.UpdateServiceLbPolicyRequest{
		ServiceLbPolicy: &pb.ServiceLbPolicy{
			Name:        name,
			Description: "Updated description",
			Labels: map[string]string{
				"env": "prod",
			},
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"description", "labels"}},
	}
	op, err = server.UpdateServiceLbPolicy(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateServiceLbPolicy failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from UpdateServiceLbPolicy")
	}

	policy, err = server.GetServiceLbPolicy(ctx, &pb.GetServiceLbPolicyRequest{Name: name})
	if err != nil {
		t.Fatalf("GetServiceLbPolicy after update failed: %v", err)
	}
	if policy.GetDescription() != "Updated description" {
		t.Errorf("expected description 'Updated description', got %q", policy.GetDescription())
	}
	if policy.GetLabels()["env"] != "prod" {
		t.Errorf("expected label env=prod, got %v", policy.GetLabels())
	}

	// 5. Update with invalid update mask path -> InvalidArgument
	invalidUpdateReq := &pb.UpdateServiceLbPolicyRequest{
		ServiceLbPolicy: &pb.ServiceLbPolicy{
			Name: name,
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"invalidField"}},
	}
	if _, err := server.UpdateServiceLbPolicy(ctx, invalidUpdateReq); status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for invalid update mask path, got: %v", err)
	}

	// 6. List policies
	listResp, err := server.ListServiceLbPolicies(ctx, &pb.ListServiceLbPoliciesRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListServiceLbPolicies failed: %v", err)
	}
	if len(listResp.GetServiceLbPolicies()) != 1 {
		t.Fatalf("expected 1 policy, got %d", len(listResp.GetServiceLbPolicies()))
	}

	// 7. List with invalid parent -> InvalidArgument
	if _, err := server.ListServiceLbPolicies(ctx, &pb.ListServiceLbPoliciesRequest{Parent: "invalid-parent"}); status.Code(err) != codes.InvalidArgument {
		t.Errorf("expected InvalidArgument for invalid parent in ListServiceLbPolicies, got: %v", err)
	}

	// 8. Delete policy
	op, err = server.DeleteServiceLbPolicy(ctx, &pb.DeleteServiceLbPolicyRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteServiceLbPolicy failed: %v", err)
	}
	if op == nil {
		t.Fatalf("expected non-nil Operation from DeleteServiceLbPolicy")
	}

	// 9. Get after delete -> 404 NotFound
	_, err = server.GetServiceLbPolicy(ctx, &pb.GetServiceLbPolicyRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", err)
	}
}

func TestParseServiceLbPolicyName(t *testing.T) {
	_, server := setupMockNetworkServicesService()

	validName := "projects/test-project/locations/global/serviceLbPolicies/test-policy"
	parsed, err := server.parseServiceLbPolicyName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid policy name: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "global" || parsed.ServiceLbPolicyName != "test-policy" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/global",
		"projects/test-project/locations/global/serviceLbPolicies",
		"invalid/test-project/locations/global/serviceLbPolicies/test-policy",
	}
	for _, name := range invalidNames {
		if _, err := server.parseServiceLbPolicyName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}

func TestParseServiceLbPolicyParent(t *testing.T) {
	_, server := setupMockNetworkServicesService()

	validParent := "projects/test-project/locations/global"
	parsed, err := server.parseServiceLbPolicyParent(validParent)
	if err != nil {
		t.Fatalf("unexpected error parsing valid parent: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "global" {
		t.Errorf("unexpected parsed parent: %+v", parsed)
	}

	invalidParents := []string{
		"",
		"projects/test-project",
		"locations/global",
		"projects/test-project/locations/global/serviceLbPolicies",
		"invalid/test-project/locations/global",
	}
	for _, p := range invalidParents {
		if _, err := server.parseServiceLbPolicyParent(p); err == nil {
			t.Errorf("expected error for invalid parent %q, got nil", p)
		}
	}
}
