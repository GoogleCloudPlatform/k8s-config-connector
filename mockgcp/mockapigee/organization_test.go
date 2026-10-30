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

package mockapigee

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/apigee/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockProjectStore struct {
	projects map[string]*projects.ProjectData
}

func (m *mockProjectStore) GetProject(name *projects.ProjectName) (*projects.ProjectData, error) {
	return m.GetProjectByID(name.OriginalValue)
}

func (m *mockProjectStore) GetProjectByID(id string) (*projects.ProjectData, error) {
	if p, ok := m.projects[id]; ok {
		return p, nil
	}
	return &projects.ProjectData{ID: id, Number: 123456789012}, nil
}

func (m *mockProjectStore) GetProjectByNumber(num string) (*projects.ProjectData, error) {
	return &projects.ProjectData{ID: "test-project", Number: 123456789012}, nil
}

func (m *mockProjectStore) GetProjectByIDOrNumber(idOrNum string) (*projects.ProjectData, error) {
	return &projects.ProjectData{ID: idOrNum, Number: 123456789012}, nil
}

func setupMockOrganizationsServer() (*MockService, *organizationsServer) {
	st := storage.NewInMemoryStorage()
	env := &common.MockEnvironment{
		Projects: &mockProjectStore{
			projects: make(map[string]*projects.ProjectData),
		},
	}
	s := &MockService{
		MockEnvironment: env,
		storage:         st,
		operations:      operations.NewOperationsService(st),
	}
	return s, &organizationsServer{MockService: s}
}

func TestOrganization_DefaultAutoPopulation(t *testing.T) {
	// Arrange
	ctx := context.Background()
	_, server := setupMockOrganizationsServer()
	orgName := "organizations/unseeded-org-1"

	// Act
	getResp, err := server.GetOrganization(ctx, &pb.GetOrganizationRequest{
		Name: orgName,
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error getting unseeded organization: %v", err)
	}
	if getResp.GetName() != "unseeded-org-1" {
		t.Errorf("expected organization name %q, got %q", "unseeded-org-1", getResp.GetName())
	}
	if getResp.GetState() != "ACTIVE" {
		t.Errorf("expected state ACTIVE, got %q", getResp.GetState())
	}
	if getResp.GetBillingType() != "PAYG" {
		t.Errorf("expected billingType PAYG, got %q", getResp.GetBillingType())
	}
	if getResp.GetRuntimeType() != "CLOUD" {
		t.Errorf("expected runtimeType CLOUD, got %q", getResp.GetRuntimeType())
	}
	if getResp.GetApigeeProjectId() != defaultApigeeProjectID {
		t.Errorf("expected apigeeProjectId %q, got %q", defaultApigeeProjectID, getResp.GetApigeeProjectId())
	}
	if len(getResp.GetCaCertificate()) == 0 {
		t.Errorf("expected non-empty CaCertificate")
	}
}

func TestOrganization_ConcurrencyHandling(t *testing.T) {
	// Arrange
	ctx := context.Background()
	_, server := setupMockOrganizationsServer()
	orgName := "organizations/concurrent-test-org"
	concurrency := 20
	errChan := make(chan error, concurrency)
	var wg sync.WaitGroup

	// Act
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := server.GetOrganization(ctx, &pb.GetOrganizationRequest{
				Name: orgName,
			})
			if err != nil {
				errChan <- err
			}
		}()
	}
	wg.Wait()
	close(errChan)

	// Assert
	for err := range errChan {
		t.Errorf("concurrent GetOrganization failed with error: %v", err)
	}

	finalOrg, err := server.GetOrganization(ctx, &pb.GetOrganizationRequest{Name: orgName})
	if err != nil {
		t.Fatalf("failed to retrieve organization after concurrent creation: %v", err)
	}
	if finalOrg.GetName() != "concurrent-test-org" {
		t.Errorf("expected final organization name %q, got %q", "concurrent-test-org", finalOrg.GetName())
	}
}

func TestOrganization_DeletionSemantics(t *testing.T) {
	// Arrange
	ctx := context.Background()
	_, server := setupMockOrganizationsServer()
	orgName := "organizations/deletion-test-org"

	// Seed / auto-populate the organization
	initialOrg, err := server.GetOrganization(ctx, &pb.GetOrganizationRequest{Name: orgName})
	if err != nil {
		t.Fatalf("failed to auto-populate organization: %v", err)
	}
	if initialOrg == nil {
		t.Fatalf("expected non-nil initial organization")
	}

	// Act
	deleteOp, err := server.DeleteOrganization(ctx, &pb.DeleteOrganizationRequest{Name: orgName})
	if err != nil {
		t.Fatalf("DeleteOrganization failed: %v", err)
	}
	if deleteOp == nil {
		t.Fatalf("expected non-nil operation from DeleteOrganization")
	}

	// Assert: subsequent GetOrganization must return NotFound and must NOT re-populate
	_, getErr := server.GetOrganization(ctx, &pb.GetOrganizationRequest{Name: orgName})
	if status.Code(getErr) != codes.NotFound {
		t.Fatalf("expected NotFound after deletion, got: %v", getErr)
	}

	// Assert: subsequent SetAddonsOrganization must also return NotFound
	_, setAddonsErr := server.SetAddonsOrganization(ctx, &pb.SetAddonsOrganizationRequest{
		Org: orgName,
		Organization: &pb.GoogleCloudApigeeV1SetAddonsRequest{
			AddonsConfig: &pb.GoogleCloudApigeeV1AddonsConfig{
				ApiSecurityConfig: &pb.GoogleCloudApigeeV1ApiSecurityConfig{
					Enabled: true,
				},
			},
		},
	})
	if status.Code(setAddonsErr) != codes.NotFound {
		t.Fatalf("expected NotFound for SetAddonsOrganization on deleted organization, got: %v", setAddonsErr)
	}

	// Assert: explicit CreateOrganization clears deletion tombstone and recreates the organization
	createOp, createErr := server.CreateOrganization(ctx, &pb.CreateOrganizationRequest{
		Parent: fmt.Sprintf("projects/%s", "deletion-test-org"),
		Organization: &pb.GoogleCloudApigeeV1Organization{
			DisplayName: "recreated-org",
		},
	})
	if createErr != nil {
		t.Fatalf("CreateOrganization failed after deletion: %v", createErr)
	}
	if createOp == nil {
		t.Fatalf("expected non-nil operation from CreateOrganization")
	}

	recreatedOrg, err := server.GetOrganization(ctx, &pb.GetOrganizationRequest{Name: orgName})
	if err != nil {
		t.Fatalf("GetOrganization failed on recreated organization: %v", err)
	}
	if recreatedOrg.GetDisplayName() != "recreated-org" {
		t.Errorf("expected displayName %q, got %q", "recreated-org", recreatedOrg.GetDisplayName())
	}
}
