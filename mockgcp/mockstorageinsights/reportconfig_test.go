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

package mockstorageinsights

import (
	"context"
	"strconv"
	"testing"

	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/genproto/googleapis/type/date"
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

func setupMockStorageInsightsService() (*MockService, *storageInsightsServer) {
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
	server := &storageInsightsServer{MockService: s}
	s.v1 = server
	return s, server
}

func TestReportConfigLifecycle(t *testing.T) {
	ctx := context.Background()
	_, server := setupMockStorageInsightsService()

	parent := "projects/test-project/locations/us-central1"

	// 1. Create ReportConfig
	createReq := &pb.CreateReportConfigRequest{
		Parent: parent,
		ReportConfig: &pb.ReportConfig{
			DisplayName: "Test Report Config",
			FrequencyOptions: &pb.FrequencyOptions{
				Frequency: pb.FrequencyOptions_DAILY,
				StartDate: &date.Date{
					Year:  2050,
					Month: 1,
					Day:   1,
				},
				EndDate: &date.Date{
					Year:  2050,
					Month: 1,
					Day:   2,
				},
			},
			ReportFormat: &pb.ReportConfig_CsvOptions{
				CsvOptions: &pb.CSVOptions{},
			},
			ReportKind: &pb.ReportConfig_ObjectMetadataReportOptions{
				ObjectMetadataReportOptions: &pb.ObjectMetadataReportOptions{
					MetadataFields: []string{"bucket", "project", "name"},
					Filter: &pb.ObjectMetadataReportOptions_StorageFilters{
						StorageFilters: &pb.CloudStorageFilters{
							Bucket: "test-bucket",
						},
					},
					DestinationOptions: &pb.ObjectMetadataReportOptions_StorageDestinationOptions{
						StorageDestinationOptions: &pb.CloudStorageDestinationOptions{
							Bucket:          "test-bucket",
							DestinationPath: "reports",
						},
					},
				},
			},
		},
	}

	created, err := server.CreateReportConfig(ctx, createReq)
	if err != nil {
		t.Fatalf("CreateReportConfig failed: %v", err)
	}

	if created.GetName() == "" {
		t.Fatalf("expected non-empty Name")
	}
	if created.GetCsvOptions().GetDelimiter() != "," {
		t.Errorf("expected default delimiter ',' but got %q", created.GetCsvOptions().GetDelimiter())
	}
	if created.GetCsvOptions().GetRecordSeparator() != "\n" {
		t.Errorf("expected default record separator '\n' but got %q", created.GetCsvOptions().GetRecordSeparator())
	}

	// 2. Get ReportConfig
	got, err := server.GetReportConfig(ctx, &pb.GetReportConfigRequest{
		Name: created.GetName(),
	})
	if err != nil {
		t.Fatalf("GetReportConfig failed: %v", err)
	}
	if got.GetDisplayName() != "Test Report Config" {
		t.Errorf("expected DisplayName 'Test Report Config', got %q", got.GetDisplayName())
	}

	// 3. Update ReportConfig
	updateReq := &pb.UpdateReportConfigRequest{
		ReportConfig: &pb.ReportConfig{
			Name:        created.GetName(),
			DisplayName: "Updated Report Config",
			FrequencyOptions: &pb.FrequencyOptions{
				Frequency: pb.FrequencyOptions_WEEKLY,
				StartDate: &date.Date{
					Year:  2050,
					Month: 1,
					Day:   1,
				},
				EndDate: &date.Date{
					Year:  2050,
					Month: 1,
					Day:   2,
				},
			},
			ReportKind: &pb.ReportConfig_ObjectMetadataReportOptions{
				ObjectMetadataReportOptions: &pb.ObjectMetadataReportOptions{
					MetadataFields: []string{"bucket", "project", "name", "size"},
					DestinationOptions: &pb.ObjectMetadataReportOptions_StorageDestinationOptions{
						StorageDestinationOptions: &pb.CloudStorageDestinationOptions{
							DestinationPath: "updated-reports",
						},
					},
				},
			},
		},
		UpdateMask: &field_mask.FieldMask{
			Paths: []string{
				"displayName",
				"frequencyOptions",
				"objectMetadataReportOptions.metadataFields",
				"objectMetadataReportOptions.storageDestinationOptions.destinationPath",
			},
		},
	}

	updated, err := server.UpdateReportConfig(ctx, updateReq)
	if err != nil {
		t.Fatalf("UpdateReportConfig failed: %v", err)
	}
	if updated.GetDisplayName() != "Updated Report Config" {
		t.Errorf("expected DisplayName 'Updated Report Config', got %q", updated.GetDisplayName())
	}
	if updated.GetFrequencyOptions().GetFrequency() != pb.FrequencyOptions_WEEKLY {
		t.Errorf("expected Frequency WEEKLY, got %v", updated.GetFrequencyOptions().GetFrequency())
	}
	if updated.GetObjectMetadataReportOptions().GetStorageDestinationOptions().GetDestinationPath() != "updated-reports" {
		t.Errorf("expected destination path 'updated-reports', got %q", updated.GetObjectMetadataReportOptions().GetStorageDestinationOptions().GetDestinationPath())
	}

	// 4. List ReportConfigs
	listResp, err := server.ListReportConfigs(ctx, &pb.ListReportConfigsRequest{
		Parent: parent,
	})
	if err != nil {
		t.Fatalf("ListReportConfigs failed: %v", err)
	}
	if len(listResp.GetReportConfigs()) != 1 {
		t.Fatalf("expected 1 report config, got %d", len(listResp.GetReportConfigs()))
	}

	// 5. Delete ReportConfig
	if _, err := server.DeleteReportConfig(ctx, &pb.DeleteReportConfigRequest{
		Name: created.GetName(),
	}); err != nil {
		t.Fatalf("DeleteReportConfig failed: %v", err)
	}

	// 6. Verify NotFound after delete
	_, err = server.GetReportConfig(ctx, &pb.GetReportConfigRequest{
		Name: created.GetName(),
	})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("expected NotFound error, got %v", err)
	}
}
