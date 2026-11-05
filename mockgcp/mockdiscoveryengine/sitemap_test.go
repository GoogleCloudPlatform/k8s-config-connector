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

package mockdiscoveryengine

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
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

func setupMockSiteSearchEngineService() (*MockService, *siteSearchEngineService) {
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
	return s, &siteSearchEngineService{MockService: s}
}

func TestParseSiteSearchEngineName(t *testing.T) {
	s, _ := setupMockSiteSearchEngineService()

	validName := "projects/test-project/locations/global/collections/default_collection/dataStores/ds-1/siteSearchEngine"
	parsed, err := s.parseSiteSearchEngineName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing siteSearchEngine name: %v", err)
	}
	if parsed.Project.Number != 1234567890 || parsed.Location != "global" || parsed.Collection != "default_collection" || parsed.DataStore != "ds-1" {
		t.Errorf("unexpected parsed fields: %+v", parsed)
	}

	invalidName := "projects/test-project/locations/global"
	if _, err := s.parseSiteSearchEngineName(invalidName); err == nil {
		t.Errorf("expected error parsing invalid siteSearchEngine name, got nil")
	}
}

func TestParseSitemapName(t *testing.T) {
	s, _ := setupMockSiteSearchEngineService()

	validName := "projects/test-project/locations/global/collections/default_collection/dataStores/ds-1/siteSearchEngine/sitemaps/sm-1"
	parsed, err := s.parseSitemapName(validName)
	if err != nil {
		t.Fatalf("unexpected error parsing sitemap name: %v", err)
	}
	if parsed.Sitemap != "sm-1" || parsed.DataStore != "ds-1" {
		t.Errorf("unexpected parsed sitemap name: %+v", parsed)
	}

	invalidName := "projects/test-project/locations/global/collections/default_collection/dataStores/ds-1/siteSearchEngine"
	if _, err := s.parseSitemapName(invalidName); err == nil {
		t.Errorf("expected error parsing invalid sitemap name, got nil")
	}
}

func TestSiteSearchEngineSitemapLifecycle(t *testing.T) {
	ctx := context.Background()
	_, svc := setupMockSiteSearchEngineService()

	parent := "projects/test-project/locations/global/collections/default_collection/dataStores/ds-1/siteSearchEngine"

	// 1. Create without enabling advanced site search should fail with InvalidArgument
	createReq := &pb.CreateSitemapRequest{
		Parent: parent,
		Sitemap: &pb.Sitemap{
			Feed: &pb.Sitemap_Uri{
				Uri: "https://example.com/sitemap.xml",
			},
		},
	}
	if _, err := svc.CreateSitemap(ctx, createReq); status.Code(err) != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument when Advanced Site Search is not enabled, got: %v", err)
	}

	// 2. Enable Advanced Site Search
	enableReq := &pb.EnableAdvancedSiteSearchRequest{
		SiteSearchEngine: parent,
	}
	enableOp, err := svc.EnableAdvancedSiteSearch(ctx, enableReq)
	if err != nil {
		t.Fatalf("unexpected error enabling advanced site search: %v", err)
	}
	if enableOp.GetName() == "" {
		t.Errorf("expected EnableAdvancedSiteSearch operation to have a name")
	}
	opResult, err := svc.operations.GetOperation(ctx, &longrunningpb.GetOperationRequest{
		Name: enableOp.GetName(),
	})
	if err != nil {
		t.Fatalf("unexpected error getting operation: %v", err)
	}
	if !opResult.GetDone() {
		t.Errorf("expected operation to be done")
	}

	// 3. Create sitemap should now succeed
	createOp, err := svc.CreateSitemap(ctx, createReq)
	if err != nil {
		t.Fatalf("unexpected error creating sitemap: %v", err)
	}
	if !createOp.GetDone() {
		t.Fatalf("expected CreateSitemap operation to be done")
	}

	// 4. Fetch sitemaps should list the created sitemap
	fetchReq := &pb.FetchSitemapsRequest{
		Parent: parent,
	}
	fetchResp, err := svc.FetchSitemaps(ctx, fetchReq)
	if err != nil {
		t.Fatalf("unexpected error fetching sitemaps: %v", err)
	}
	if len(fetchResp.GetSitemapsMetadata()) != 1 {
		t.Fatalf("expected 1 sitemap, got %d", len(fetchResp.GetSitemapsMetadata()))
	}
	smMeta := fetchResp.GetSitemapsMetadata()[0]
	if smMeta.GetSitemap().GetUri() != "https://example.com/sitemap.xml" {
		t.Errorf("unexpected uri: %q", smMeta.GetSitemap().GetUri())
	}
	if smMeta.GetSitemap().GetCreateTime() != nil {
		t.Errorf("expected createTime to be cleared in FetchSitemaps response, got: %v", smMeta.GetSitemap().GetCreateTime())
	}
	createdName := smMeta.GetSitemap().GetName()

	// 5. Delete sitemap
	deleteReq := &pb.DeleteSitemapRequest{
		Name: createdName,
	}
	deleteOp, err := svc.DeleteSitemap(ctx, deleteReq)
	if err != nil {
		t.Fatalf("unexpected error deleting sitemap: %v", err)
	}
	if !deleteOp.GetDone() {
		t.Errorf("expected DeleteSitemap operation to be done")
	}
	if deleteOp.GetResponse() != nil {
		t.Errorf("expected DeleteSitemap operation response to be nil matching real GCP, got: %v", deleteOp.GetResponse())
	}

	// 6. Fetch sitemaps should now be empty
	fetchRespAfterDelete, err := svc.FetchSitemaps(ctx, fetchReq)
	if err != nil {
		t.Fatalf("unexpected error fetching sitemaps after delete: %v", err)
	}
	if len(fetchRespAfterDelete.GetSitemapsMetadata()) != 0 {
		t.Errorf("expected 0 sitemaps after delete, got %d", len(fetchRespAfterDelete.GetSitemapsMetadata()))
	}

	// 7. Delete again should fail with NotFound
	if _, err := svc.DeleteSitemap(ctx, deleteReq); status.Code(err) != codes.NotFound {
		t.Errorf("expected NotFound deleting non-existent sitemap, got: %v", err)
	}
}
