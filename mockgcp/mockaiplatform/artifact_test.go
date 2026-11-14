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

package mockaiplatform

import (
	"context"
	"testing"

	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func setupMockArtifactService() (*MockService, *metadataStoreService) {
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
	return s, &metadataStoreService{MockService: s}
}

func TestParseArtifactName(t *testing.T) {
	s, _ := setupMockArtifactService()

	validNames := []struct {
		input           string
		expectedProject string
		expectedLoc     string
		expectedStore   string
		expectedArt     string
	}{
		{
			input:           "projects/test-project/locations/us-central1/metadataStores/my-store/artifacts/my-artifact",
			expectedProject: "test-project",
			expectedLoc:     "us-central1",
			expectedStore:   "my-store",
			expectedArt:     "my-artifact",
		},
		{
			input:           "projects/1234567890/locations/us-central1/metadataStores/my-store/artifacts/my-artifact",
			expectedProject: "test-project",
			expectedLoc:     "us-central1",
			expectedStore:   "my-store",
			expectedArt:     "my-artifact",
		},
	}

	for _, tc := range validNames {
		parsed, err := s.parseArtifactName(tc.input)
		if err != nil {
			t.Fatalf("unexpected error parsing %q: %v", tc.input, err)
		}
		if parsed.Project.ID != tc.expectedProject {
			t.Errorf("expected project %q, got %q", tc.expectedProject, parsed.Project.ID)
		}
		if parsed.Location != tc.expectedLoc {
			t.Errorf("expected location %q, got %q", tc.expectedLoc, parsed.Location)
		}
		if parsed.MetadataStoreId != tc.expectedStore {
			t.Errorf("expected metadataStore %q, got %q", tc.expectedStore, parsed.MetadataStoreId)
		}
		if parsed.ArtifactId != tc.expectedArt {
			t.Errorf("expected artifactId %q, got %q", tc.expectedArt, parsed.ArtifactId)
		}
	}

	invalidNames := []string{
		"",
		"projects/test-project",
		"projects/test-project/locations/us-central1/metadataStores/my-store",
		"projects/test-project/locations/us-central1/artifacts/my-artifact",
	}

	for _, inv := range invalidNames {
		_, err := s.parseArtifactName(inv)
		if err == nil {
			t.Errorf("expected error parsing invalid name %q, got nil", inv)
		}
	}
}

func TestArtifactLifecycle(t *testing.T) {
	ctx := context.Background()
	_, service := setupMockArtifactService()

	parent := "projects/test-project/locations/us-central1/metadataStores/default"
	artifactId := "test-artifact"

	// 1. Get non-existent -> NotFound
	_, err := service.GetArtifact(ctx, &pb.GetArtifactRequest{
		Name: parent + "/artifacts/" + artifactId,
	})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound, got %v", err)
	}

	// 2. Create Artifact
	created, err := service.CreateArtifact(ctx, &pb.CreateArtifactRequest{
		Parent:     parent,
		ArtifactId: artifactId,
		Artifact: &pb.Artifact{
			DisplayName: "Test Artifact",
			SchemaTitle: "system.Artifact",
			Uri:         "gs://test-bucket/model",
		},
	})
	if err != nil {
		t.Fatalf("CreateArtifact failed: %v", err)
	}

	expectedName := "projects/1234567890/locations/us-central1/metadataStores/default/artifacts/test-artifact"
	if created.GetName() != expectedName {
		t.Errorf("expected name %q, got %q", expectedName, created.GetName())
	}
	if created.GetSchemaVersion() != "0.0.1" {
		t.Errorf("expected defaulted schemaVersion 0.0.1, got %q", created.GetSchemaVersion())
	}
	if created.GetMetadata() == nil {
		t.Errorf("expected non-nil metadata struct")
	}

	// 3. Get Artifact
	got, err := service.GetArtifact(ctx, &pb.GetArtifactRequest{
		Name: parent + "/artifacts/" + artifactId,
	})
	if err != nil {
		t.Fatalf("GetArtifact failed: %v", err)
	}
	if got.GetDisplayName() != "Test Artifact" {
		t.Errorf("expected displayName 'Test Artifact', got %q", got.GetDisplayName())
	}

	// 4. Update Artifact
	updated, err := service.UpdateArtifact(ctx, &pb.UpdateArtifactRequest{
		Artifact: &pb.Artifact{
			Name:        parent + "/artifacts/" + artifactId,
			DisplayName: "Updated Artifact",
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{"displayName"},
		},
	})
	if err != nil {
		t.Fatalf("UpdateArtifact failed: %v", err)
	}
	if updated.GetDisplayName() != "Updated Artifact" {
		t.Errorf("expected updated displayName, got %q", updated.GetDisplayName())
	}
	// Verify Uri wasn't cleared
	if updated.GetUri() != "gs://test-bucket/model" {
		t.Errorf("expected uri preserved, got %q", updated.GetUri())
	}

	// 5. Delete Artifact
	op, err := service.DeleteArtifact(ctx, &pb.DeleteArtifactRequest{
		Name: parent + "/artifacts/" + artifactId,
	})
	if err != nil {
		t.Fatalf("DeleteArtifact failed: %v", err)
	}
	if !op.GetDone() {
		t.Errorf("expected operation to be done")
	}

	// 6. Verify deleted -> NotFound
	_, err = service.GetArtifact(ctx, &pb.GetArtifactRequest{
		Name: parent + "/artifacts/" + artifactId,
	})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound after delete, got %v", err)
	}
}
