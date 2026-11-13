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

package mockdlp

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/dlp/apiv2/dlppb"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
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

func setupMockDLPService() (*MockService, *dlpServiceServer) {
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
	server := newDLPServiceServer(s)
	return s, server
}

func TestCreateContentPolicy_Lifecycle(t *testing.T) {
	// Arrange
	ctx := context.Background()
	_, server := setupMockDLPService()
	parent := "projects/test-project/locations/global"
	policyID := "policy-1"
	createReq := &pb.CreateContentPolicyRequest{
		Parent:          parent,
		ContentPolicyId: policyID,
		ContentPolicy: &pb.ContentPolicy{
			DisplayName: "Test Content Policy",
			InspectConfig: &pb.InspectConfig{
				InfoTypes: []*pb.InfoType{
					{Name: "EMAIL_ADDRESS"},
				},
			},
		},
	}

	// Act
	created, err := server.CreateContentPolicy(ctx, createReq)

	// Assert
	if err != nil {
		t.Fatalf("CreateContentPolicy failed: %v", err)
	}
	expectedName := "projects/test-project/locations/global/contentPolicies/policy-1"
	if created.GetName() != expectedName {
		t.Errorf("expected name %q, got %q", expectedName, created.GetName())
	}
	if created.GetInspectConfig().GetLimits() == nil {
		t.Errorf("expected InspectConfig.Limits to be auto-populated")
	}
}

func TestListContentPolicies_DataRaceAndIsolation(t *testing.T) {
	// Arrange
	ctx := context.Background()
	_, server := setupMockDLPService()
	parent := "projects/test-project/locations/global"
	createReq := &pb.CreateContentPolicyRequest{
		Parent:          parent,
		ContentPolicyId: "policy-isolated",
		ContentPolicy: &pb.ContentPolicy{
			DisplayName: "Original Name",
			InspectConfig: &pb.InspectConfig{
				InfoTypes: []*pb.InfoType{
					{Name: "EMAIL_ADDRESS"},
				},
			},
		},
	}
	if _, err := server.CreateContentPolicy(ctx, createReq); err != nil {
		t.Fatalf("failed to create initial content policy: %v", err)
	}

	// Act
	listResp, err := server.ListContentPolicies(ctx, &pb.ListContentPoliciesRequest{
		Parent: parent,
	})
	if err != nil {
		t.Fatalf("ListContentPolicies failed: %v", err)
	}
	if len(listResp.GetContentPolicies()) != 1 {
		t.Fatalf("expected 1 content policy in list, got %d", len(listResp.GetContentPolicies()))
	}
	// Mutate the item returned from ListContentPolicies
	listResp.GetContentPolicies()[0].DisplayName = "Mutated In-Memory Name"
	getResp, err := server.GetContentPolicy(ctx, &pb.GetContentPolicyRequest{
		Name: "projects/test-project/locations/global/contentPolicies/policy-isolated",
	})

	// Assert
	if err != nil {
		t.Fatalf("GetContentPolicy failed: %v", err)
	}
	if getResp.GetDisplayName() != "Original Name" {
		t.Errorf("stored object was mutated! expected %q, got %q", "Original Name", getResp.GetDisplayName())
	}
}

func TestContentPolicy_ValidationErrors(t *testing.T) {
	// Arrange
	ctx := context.Background()
	_, server := setupMockDLPService()

	// Act
	_, errMissingID := server.CreateContentPolicy(ctx, &pb.CreateContentPolicyRequest{
		Parent: "projects/test-project/locations/global",
	})
	_, errInvalidParent := server.CreateContentPolicy(ctx, &pb.CreateContentPolicyRequest{
		Parent:          "invalid-parent",
		ContentPolicyId: "policy-1",
	})
	_, errMissingPolicy := server.CreateContentPolicy(ctx, &pb.CreateContentPolicyRequest{
		Parent:          "projects/test-project/locations/global",
		ContentPolicyId: "policy-1",
	})
	_, errNotFound := server.GetContentPolicy(ctx, &pb.GetContentPolicyRequest{
		Name: "projects/test-project/locations/global/contentPolicies/non-existent",
	})

	// Assert
	if status.Code(errMissingID) != codes.InvalidArgument {
		t.Errorf("errMissingID code = %v, want %v", status.Code(errMissingID), codes.InvalidArgument)
	}
	if status.Code(errInvalidParent) != codes.InvalidArgument {
		t.Errorf("errInvalidParent code = %v, want %v", status.Code(errInvalidParent), codes.InvalidArgument)
	}
	if status.Code(errMissingPolicy) != codes.InvalidArgument {
		t.Errorf("errMissingPolicy code = %v, want %v", status.Code(errMissingPolicy), codes.InvalidArgument)
	}
	if status.Code(errNotFound) != codes.NotFound {
		t.Errorf("errNotFound code = %v, want %v", status.Code(errNotFound), codes.NotFound)
	}
}

func TestUpdateAndDeleteContentPolicy(t *testing.T) {
	// Arrange
	ctx := context.Background()
	_, server := setupMockDLPService()
	parent := "projects/test-project/locations/global"
	policyName := "projects/test-project/locations/global/contentPolicies/policy-crud"
	_, err := server.CreateContentPolicy(ctx, &pb.CreateContentPolicyRequest{
		Parent:          parent,
		ContentPolicyId: "policy-crud",
		ContentPolicy: &pb.ContentPolicy{
			DisplayName: "Before Update",
			InspectConfig: &pb.InspectConfig{
				InfoTypes: []*pb.InfoType{{Name: "EMAIL_ADDRESS"}},
			},
		},
	})
	if err != nil {
		t.Fatalf("CreateContentPolicy failed: %v", err)
	}

	// Act
	updated, errUpdate := server.UpdateContentPolicy(ctx, &pb.UpdateContentPolicyRequest{
		Name: policyName,
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"display_name"},
		},
		ContentPolicy: &pb.ContentPolicy{
			DisplayName: "After Update",
		},
	})
	_, errDelete := server.DeleteContentPolicy(ctx, &pb.DeleteContentPolicyRequest{
		Name: policyName,
	})
	_, errGetAfterDelete := server.GetContentPolicy(ctx, &pb.GetContentPolicyRequest{
		Name: policyName,
	})

	// Assert
	if errUpdate != nil {
		t.Fatalf("UpdateContentPolicy failed: %v", errUpdate)
	}
	if updated.GetDisplayName() != "After Update" {
		t.Errorf("expected updated displayName %q, got %q", "After Update", updated.GetDisplayName())
	}
	if errDelete != nil {
		t.Fatalf("DeleteContentPolicy failed: %v", errDelete)
	}
	if status.Code(errGetAfterDelete) != codes.NotFound {
		t.Errorf("errGetAfterDelete code = %v, want %v", status.Code(errGetAfterDelete), codes.NotFound)
	}
}
