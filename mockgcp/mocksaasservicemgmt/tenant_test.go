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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	pb "cloud.google.com/go/saasplatform/saasservicemgmt/apiv1beta1/saasservicemgmtpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
)

func setupMockDeploymentsService() (*MockService, *SaasDeploymentsServer) {
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
	}
	s.saasDeploymentsServer = &SaasDeploymentsServer{MockService: s}
	return s, s.saasDeploymentsServer
}

func TestTenantLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockDeploymentsService()

	parent := "projects/test-project/locations/us-central1"
	tenantID := "test-tenant"
	name := parent + "/tenants/" + tenantID

	// 1. Get before creation -> 404 NotFound
	_, err := server.GetTenant(ctx, &pb.GetTenantRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound for non-existent tenant, got: %v", err)
	}

	// 2. Create tenant
	createReq := &pb.CreateTenantRequest{
		Parent:   parent,
		TenantId: tenantID,
		Tenant: &pb.Tenant{
			Saas:             "projects/test-project/locations/us-central1/saas/test-saas",
			ConsumerResource: "projects/test-project/locations/us-central1/shoes/test-shoe",
			Labels: map[string]string{
				"env": "test",
			},
		},
	}
	created, err := server.CreateTenant(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateTenant failed: %v", err)
	}
	if created.GetName() != name {
		t.Errorf("expected name %q, got %q", name, created.GetName())
	}

	// 3. Get after creation
	got, err := server.GetTenant(ctx, &pb.GetTenantRequest{Name: name})
	if err != nil {
		t.Fatalf("GetTenant failed: %v", err)
	}
	if got.GetName() != name {
		t.Errorf("expected name %q, got %q", name, got.GetName())
	}
	if got.GetSaas() != "projects/test-project/locations/us-central1/saas/test-saas" {
		t.Errorf("expected saas test-saas, got %v", got.GetSaas())
	}
	if got.GetConsumerResource() != "projects/test-project/locations/us-central1/shoes/test-shoe" {
		t.Errorf("expected consumerResource test-shoe, got %v", got.GetConsumerResource())
	}
	if got.GetLabels()["env"] != "test" {
		t.Errorf("expected label env=test, got %v", got.GetLabels())
	}

	// 4. Update tenant
	updateReq := &pb.UpdateTenantRequest{
		Tenant: &pb.Tenant{
			Name:             name,
			Saas:             "projects/test-project/locations/us-central1/saas/test-saas",
			ConsumerResource: "projects/test-project/locations/us-central1/shoes/updated-shoe",
			Labels: map[string]string{
				"env": "prod",
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"consumer_resource", "labels"},
		},
	}
	updated, err := server.UpdateTenant(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateTenant failed: %v", err)
	}
	if updated.GetName() != name {
		t.Errorf("expected name %q, got %q", name, updated.GetName())
	}

	// 5. Get after update
	got, err = server.GetTenant(ctx, &pb.GetTenantRequest{Name: name})
	if err != nil {
		t.Fatalf("GetTenant failed: %v", err)
	}
	if got.GetConsumerResource() != "projects/test-project/locations/us-central1/shoes/updated-shoe" {
		t.Errorf("expected consumerResource updated-shoe, got %v", got.GetConsumerResource())
	}
	if got.GetLabels()["env"] != "prod" {
		t.Errorf("expected label env=prod, got %v", got.GetLabels())
	}

	// 6. List tenants
	listResp, err := server.ListTenants(ctx, &pb.ListTenantsRequest{Parent: parent})
	if err != nil {
		t.Fatalf("ListTenants failed: %v", err)
	}
	if len(listResp.GetTenants()) != 1 {
		t.Fatalf("expected 1 tenant in list, got %d", len(listResp.GetTenants()))
	}
	if listResp.GetTenants()[0].GetName() != name {
		t.Errorf("expected list tenant name %q, got %q", name, listResp.GetTenants()[0].GetName())
	}

	// 7. Delete tenant
	_, err = server.DeleteTenant(ctx, &pb.DeleteTenantRequest{Name: name})
	if err != nil {
		t.Fatalf("DeleteTenant failed: %v", err)
	}

	// 8. Get after delete -> 404 NotFound
	_, err = server.GetTenant(ctx, &pb.GetTenantRequest{Name: name})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got: %v", err)
	}
}

func TestParseTenantName(t *testing.T) {
	s, _ := setupMockDeploymentsService()

	tests := []struct {
		name        string
		input       string
		wantProject string
		wantLoc     string
		wantTenant  string
		wantErr     bool
	}{
		{
			name:        "valid name",
			input:       "projects/test-project/locations/us-central1/tenants/my-tenant",
			wantProject: "test-project",
			wantLoc:     "us-central1",
			wantTenant:  "my-tenant",
			wantErr:     false,
		},
		{
			name:    "invalid segment count",
			input:   "projects/test-project/locations/us-central1/tenants",
			wantErr: true,
		},
		{
			name:    "invalid prefix",
			input:   "foos/test-project/locations/us-central1/tenants/my-tenant",
			wantErr: true,
		},
		{
			name:    "invalid collection",
			input:   "projects/test-project/locations/us-central1/foos/my-tenant",
			wantErr: true,
		},
		{
			name:    "unknown project",
			input:   "projects/non-existent-project/locations/us-central1/tenants/my-tenant",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := s.parseTenantName(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q, got nil", tc.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for %q: %v", tc.input, err)
			}
			if got.Project.ID != tc.wantProject {
				t.Errorf("expected project %q, got %q", tc.wantProject, got.Project.ID)
			}
			if got.Location != tc.wantLoc {
				t.Errorf("expected location %q, got %q", tc.wantLoc, got.Location)
			}
			if got.Tenant != tc.wantTenant {
				t.Errorf("expected tenant %q, got %q", tc.wantTenant, got.Tenant)
			}
			if got.String() != tc.input {
				t.Errorf("expected String() %q, got %q", tc.input, got.String())
			}
		})
	}
}
