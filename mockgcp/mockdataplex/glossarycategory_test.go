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

package mockdataplex

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockProjectStore struct {
	projectsByID map[string]*projects.ProjectData
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
	return m.GetProjectByIDOrNumber(projectNumberAsString)
}

func (m *mockProjectStore) GetProjectByIDOrNumber(projectIDOrNumber string) (*projects.ProjectData, error) {
	if p, ok := m.projectsByID[projectIDOrNumber]; ok {
		return p, nil
	}
	if n, err := strconv.ParseInt(projectIDOrNumber, 10, 64); err == nil {
		p := &projects.ProjectData{
			Number: n,
			ID:     projectIDOrNumber,
		}
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", projectIDOrNumber)
}

func TestListGlossaryCategories_PrefixIsolation(t *testing.T) {
	ctx := context.Background()
	store := &mockProjectStore{
		projectsByID: map[string]*projects.ProjectData{
			"test-project": {
				Number: 1234567890,
				ID:     "test-project",
			},
		},
	}
	env := &common.MockEnvironment{
		Projects: store,
	}
	st := storage.NewInMemoryStorage()
	mockSvc := New(env, st)
	glossarySvc := mockSvc.businessGlossaryService

	// Create category in "myglossary"
	cat1, err := glossarySvc.CreateGlossaryCategory(ctx, &pb.CreateGlossaryCategoryRequest{
		Parent:     "projects/test-project/locations/us-central1/glossaries/myglossary",
		CategoryId: "cat-1",
		Category: &pb.GlossaryCategory{
			DisplayName: "Category 1 in myglossary",
		},
	})
	if err != nil {
		t.Fatalf("failed to create category in myglossary: %v", err)
	}

	// Create category in sibling glossary whose name starts with "myglossary"
	cat2, err := glossarySvc.CreateGlossaryCategory(ctx, &pb.CreateGlossaryCategoryRequest{
		Parent:     "projects/test-project/locations/us-central1/glossaries/myglossary-sibling",
		CategoryId: "cat-2",
		Category: &pb.GlossaryCategory{
			DisplayName: "Category 2 in myglossary-sibling",
		},
	})
	if err != nil {
		t.Fatalf("failed to create category in myglossary-sibling: %v", err)
	}

	// List categories in "myglossary" - should only return cat1
	res1, err := glossarySvc.ListGlossaryCategories(ctx, &pb.ListGlossaryCategoriesRequest{
		Parent: "projects/test-project/locations/us-central1/glossaries/myglossary",
	})
	if err != nil {
		t.Fatalf("ListGlossaryCategories for myglossary failed: %v", err)
	}
	if len(res1.Categories) != 1 || res1.Categories[0].Name != cat1.Name {
		t.Errorf("ListGlossaryCategories(myglossary) = %v; want [%s]", res1.Categories, cat1.Name)
	}

	// List categories in "myglossary-sibling" - should only return cat2
	res2, err := glossarySvc.ListGlossaryCategories(ctx, &pb.ListGlossaryCategoriesRequest{
		Parent: "projects/test-project/locations/us-central1/glossaries/myglossary-sibling",
	})
	if err != nil {
		t.Fatalf("ListGlossaryCategories for myglossary-sibling failed: %v", err)
	}
	if len(res2.Categories) != 1 || res2.Categories[0].Name != cat2.Name {
		t.Errorf("ListGlossaryCategories(myglossary-sibling) = %v; want [%s]", res2.Categories, cat2.Name)
	}
}
