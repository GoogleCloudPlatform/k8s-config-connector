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
	"fmt"
	"strconv"
	"testing"

	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	resourcemanagerpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/resourcemanager/v3"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockresourcemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestListDatasetConfigsPagination(t *testing.T) {
	ctx := context.Background()
	st := storage.NewInMemoryStorage()
	env := &common.MockEnvironment{
		Projects: mockresourcemanager.NewProjectStore(st),
	}

	// 1. Create a mock project in storage so parseDatasetConfigName succeeds
	projectID := "test-project"
	projectNumber := int64(123456789)
	projObj := &resourcemanagerpb.Project{
		ProjectId: projectID,
		Name:      "projects/" + strconv.FormatInt(projectNumber, 10),
	}
	if err := st.Create(ctx, "projects/"+projectID, projObj); err != nil {
		t.Fatalf("failed to create mock project: %v", err)
	}

	// 2. Initialize the service
	svc := New(env, st).(*MockService)

	// 3. Populate 5 dataset configs in the storage
	location := "us-central1"
	parent := fmt.Sprintf("projects/%s/locations/%s", projectID, location)
	numConfigs := 5
	for i := 1; i <= numConfigs; i++ {
		name := fmt.Sprintf("dataset-config-%d", i)
		fqn := fmt.Sprintf("%s/datasetConfigs/%s", parent, name)
		config := &pb.DatasetConfig{
			Name: fqn,
		}
		if err := st.Create(ctx, fqn, config); err != nil {
			t.Fatalf("failed to create mock dataset config: %v", err)
		}
	}

	server := svc.storageInsightsServer

	tests := []struct {
		name         string
		pageSize     int32
		pageToken    string
		expectedLen  int
		expectedNext string
		expectedCode codes.Code
	}{
		{
			name:         "default page size and empty token",
			pageSize:     0,
			pageToken:    "",
			expectedLen:  5,
			expectedNext: "",
			expectedCode: codes.OK,
		},
		{
			name:         "page size 2 and empty token",
			pageSize:     2,
			pageToken:    "",
			expectedLen:  2,
			expectedNext: "2",
			expectedCode: codes.OK,
		},
		{
			name:         "page size 2 and token 2",
			pageSize:     2,
			pageToken:    "2",
			expectedLen:  2,
			expectedNext: "4",
			expectedCode: codes.OK,
		},
		{
			name:         "negative page token (should not panic, clamp to 0)",
			pageSize:     2,
			pageToken:    "-3",
			expectedLen:  2,
			expectedNext: "2",
			expectedCode: codes.OK,
		},
		{
			name:         "extremely large page token (should not panic, clamp to end)",
			pageSize:     2,
			pageToken:    "100",
			expectedLen:  0,
			expectedNext: "",
			expectedCode: codes.OK,
		},
		{
			name:         "invalid non-integer page token (should return error)",
			pageSize:     2,
			pageToken:    "abc",
			expectedLen:  0,
			expectedNext: "",
			expectedCode: codes.InvalidArgument,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.ListDatasetConfigsRequest{
				Parent:    parent,
				PageSize:  tc.pageSize,
				PageToken: tc.pageToken,
			}

			resp, err := server.ListDatasetConfigs(ctx, req)

			if tc.expectedCode != codes.OK {
				if err == nil {
					t.Fatalf("expected error code %v, got success", tc.expectedCode)
				}
				statusErr, ok := status.FromError(err)
				if !ok {
					t.Fatalf("expected status error, got: %v", err)
				}
				if statusErr.Code() != tc.expectedCode {
					t.Fatalf("expected error code %v, got %v", tc.expectedCode, statusErr.Code())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(resp.DatasetConfigs) != tc.expectedLen {
				t.Errorf("expected dataset configs length %d, got %d", tc.expectedLen, len(resp.DatasetConfigs))
			}

			if resp.NextPageToken != tc.expectedNext {
				t.Errorf("expected NextPageToken %q, got %q", tc.expectedNext, resp.NextPageToken)
			}
		})
	}
}
