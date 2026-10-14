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
	"strings"
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
	s.v1 = &NetworkServicesServer{MockService: s}
	return s, s.v1
}

func TestParseLbTrafficExtensionName(t *testing.T) {
	_, server := setupMockNetworkServicesService()

	validName := "projects/test-project/locations/us-central1/lbTrafficExtensions/ext-1"
	parsed, err := server.parseLbTrafficExtensionName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing valid name: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "us-central1" || parsed.LbTrafficExtensionName != "ext-1" {
		t.Errorf("unexpected parsed name: %+v", parsed)
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/us-central1",
		"projects/test-project/locations/us-central1/lbTrafficExtensions",
		"projects/test-project/locations/us-central1/lbTrafficExtensions/",
		"projects//locations/us-central1/lbTrafficExtensions/ext-1",
		"projects/test-project/locations//lbTrafficExtensions/ext-1",
		"invalid/test-project/locations/us-central1/lbTrafficExtensions/ext-1",
		"projects/nonexistent/locations/us-central1/lbTrafficExtensions/ext-1",
	}
	for _, name := range invalidNames {
		if _, err := server.parseLbTrafficExtensionName(name); err == nil {
			t.Errorf("expected error for invalid name %q, got nil", name)
		}
	}
}

func TestParseLbTrafficExtensionParent(t *testing.T) {
	_, server := setupMockNetworkServicesService()

	validParent := "projects/test-project/locations/us-central1"
	parsed, err := server.parseLbTrafficExtensionParent(validParent)
	if err != nil {
		t.Fatalf("unexpected error parsing valid parent: %v", err)
	}
	if parsed.Project.ID != "test-project" || parsed.Location != "us-central1" {
		t.Errorf("unexpected parsed parent: %+v", parsed)
	}

	invalidParents := []string{
		"",
		"projects",
		"projects/test-project",
		"projects/test-project/locations",
		"projects/test-project/locations/",
		"projects//locations/us-central1",
		"invalid/test-project/locations/us-central1",
		"projects/nonexistent/locations/us-central1",
	}
	for _, parent := range invalidParents {
		if _, err := server.parseLbTrafficExtensionParent(parent); err == nil {
			t.Errorf("expected error for invalid parent %q, got nil", parent)
		}
	}
}

func TestLbTrafficExtensionLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockNetworkServicesService()

	parent := "projects/test-project/locations/us-central1"
	extensionID := "ext-test"
	name := parent + "/lbTrafficExtensions/" + extensionID

	// 1. Get non-existent extension -> 404 NotFound
	if _, err := server.GetLbTrafficExtension(ctx, &pb.GetLbTrafficExtensionRequest{Name: name}); status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent extension, got: %v", err)
	}

	// 2. Create validation errors
	if _, err := server.CreateLbTrafficExtension(ctx, &pb.CreateLbTrafficExtensionRequest{
		Parent: parent,
	}); status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument for empty id, got: %v", err)
	}
	if _, err := server.CreateLbTrafficExtension(ctx, &pb.CreateLbTrafficExtensionRequest{
		Parent:               parent,
		LbTrafficExtensionId: extensionID,
		LbTrafficExtension:   nil,
	}); status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument for nil extension, got: %v", err)
	}

	// 3. Create valid extension
	createReq := &pb.CreateLbTrafficExtensionRequest{
		Parent:               parent,
		LbTrafficExtensionId: extensionID,
		LbTrafficExtension: &pb.LbTrafficExtension{
			Description: "Initial description",
			Labels: map[string]string{
				"env": "test",
			},
			ForwardingRules: []string{
				"https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/forwardingRules/fwd-rule",
			},
			ExtensionChains: []*pb.ExtensionChain{
				{
					Name: "chain-1",
					MatchCondition: &pb.ExtensionChain_MatchCondition{
						CelExpression: "request.headers['x-test'] == 'true'",
					},
					Extensions: []*pb.ExtensionChain_Extension{
						{
							Name:    "ext-backend",
							Service: "https://www.googleapis.com/compute/v1/projects/test-project/regions/us-central1/backendServices/be-svc",
						},
					},
				},
			},
			LoadBalancingScheme: pb.LoadBalancingScheme_INTERNAL_MANAGED,
		},
	}
	createOp, err := server.CreateLbTrafficExtension(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateLbTrafficExtension failed: %v", err)
	}
	if createOp == nil {
		t.Fatalf("expected non-nil LRO from CreateLbTrafficExtension")
	}

	// 4. List extensions
	listResp, err := server.ListLbTrafficExtensions(ctx, &pb.ListLbTrafficExtensionsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListLbTrafficExtensions failed: %v", err)
	}
	if len(listResp.LbTrafficExtensions) != 1 {
		t.Fatalf("expected 1 extension in list, got %d", len(listResp.LbTrafficExtensions))
	}

	// List with project number parent
	listRespNum, err := server.ListLbTrafficExtensions(ctx, &pb.ListLbTrafficExtensionsRequest{
		Parent: "projects/1234567890/locations/us-central1",
	})
	if err != nil {
		t.Fatalf("ListLbTrafficExtensions with project number parent failed: %v", err)
	}
	if len(listRespNum.LbTrafficExtensions) != 1 {
		t.Fatalf("expected 1 extension with project number parent, got %d", len(listRespNum.LbTrafficExtensions))
	}

	// List with invalid parent
	if _, err := server.ListLbTrafficExtensions(ctx, &pb.ListLbTrafficExtensionsRequest{
		Parent: "invalid-parent",
	}); status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument for invalid list parent, got: %v", err)
	}

	// 5. Get created extension and verify canonicalization (project ID -> project Number)
	getResp, err := server.GetLbTrafficExtension(ctx, &pb.GetLbTrafficExtensionRequest{Name: name})
	if err != nil {
		t.Fatalf("GetLbTrafficExtension failed: %v", err)
	}
	if getResp.Name != name {
		t.Errorf("expected name %q, got %q", name, getResp.Name)
	}
	if getResp.Description != "Initial description" {
		t.Errorf("expected description 'Initial description', got %q", getResp.Description)
	}
	if len(getResp.ForwardingRules) != 1 || !strings.Contains(getResp.ForwardingRules[0], "1234567890") {
		t.Errorf("expected forwardingRule project number canonicalization, got %+v", getResp.ForwardingRules)
	}
	if len(getResp.ExtensionChains) != 1 || !strings.Contains(getResp.ExtensionChains[0].Extensions[0].Service, "1234567890") {
		t.Errorf("expected service project number canonicalization, got %+v", getResp.ExtensionChains)
	}

	// 6. Update validations
	if _, err := server.UpdateLbTrafficExtension(ctx, &pb.UpdateLbTrafficExtensionRequest{}); status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument for nil extension update, got: %v", err)
	}

	// Immutable loadBalancingScheme check
	if _, err := server.UpdateLbTrafficExtension(ctx, &pb.UpdateLbTrafficExtensionRequest{
		LbTrafficExtension: &pb.LbTrafficExtension{
			Name:                name,
			LoadBalancingScheme: pb.LoadBalancingScheme_EXTERNAL_MANAGED,
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"loadBalancingScheme"}},
	}); status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument for immutable loadBalancingScheme update, got: %v", err)
	}

	// Invalid updateMask path
	if _, err := server.UpdateLbTrafficExtension(ctx, &pb.UpdateLbTrafficExtensionRequest{
		LbTrafficExtension: &pb.LbTrafficExtension{
			Name: name,
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"invalid_path"}},
	}); status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument for invalid updateMask path, got: %v", err)
	}

	// Update non-existent resource
	if _, err := server.UpdateLbTrafficExtension(ctx, &pb.UpdateLbTrafficExtensionRequest{
		LbTrafficExtension: &pb.LbTrafficExtension{
			Name: "projects/test-project/locations/us-central1/lbTrafficExtensions/non-existent",
		},
		UpdateMask: &field_mask.FieldMask{Paths: []string{"description"}},
	}); status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent update, got: %v", err)
	}

	// 7. Valid Update
	updateReq := &pb.UpdateLbTrafficExtensionRequest{
		LbTrafficExtension: &pb.LbTrafficExtension{
			Name:        name,
			Description: "Updated description",
			Labels: map[string]string{
				"env": "prod",
			},
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"description", "labels"},
		},
	}
	updateOp, err := server.UpdateLbTrafficExtension(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateLbTrafficExtension failed: %v", err)
	}
	if updateOp == nil {
		t.Fatalf("expected non-nil LRO from UpdateLbTrafficExtension")
	}

	getResp, err = server.GetLbTrafficExtension(ctx, &pb.GetLbTrafficExtensionRequest{Name: name})
	if err != nil {
		t.Fatalf("GetLbTrafficExtension after update failed: %v", err)
	}
	if getResp.Description != "Updated description" {
		t.Errorf("expected updated description 'Updated description', got %q", getResp.Description)
	}
	if getResp.Labels["env"] != "prod" {
		t.Errorf("expected updated label env=prod, got %v", getResp.Labels)
	}

	// 8. Delete
	// Delete non-existent
	if _, err := server.DeleteLbTrafficExtension(ctx, &pb.DeleteLbTrafficExtensionRequest{
		Name: "projects/test-project/locations/us-central1/lbTrafficExtensions/non-existent",
	}); status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound deleting non-existent extension, got: %v", err)
	}

	// Delete created
	deleteOp, err := server.DeleteLbTrafficExtension(ctx, &pb.DeleteLbTrafficExtensionRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteLbTrafficExtension failed: %v", err)
	}
	if deleteOp == nil {
		t.Fatalf("expected non-nil LRO from DeleteLbTrafficExtension")
	}

	// Verify deleted
	if _, err := server.GetLbTrafficExtension(ctx, &pb.GetLbTrafficExtensionRequest{Name: name}); status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got: %v", err)
	}

	listAfterDel, err := server.ListLbTrafficExtensions(ctx, &pb.ListLbTrafficExtensionsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListLbTrafficExtensions after delete failed: %v", err)
	}
	if len(listAfterDel.LbTrafficExtensions) != 0 {
		t.Fatalf("expected 0 extensions after delete, got %d", len(listAfterDel.LbTrafficExtensions))
	}
}
