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

package e2e

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

func TestIsListOrOperationResponse(t *testing.T) {
	tests := []struct {
		name string
		body map[string]any
		want bool
	}{
		{
			name: "nil body",
			body: nil,
			want: true,
		},
		{
			name: "operations list",
			body: map[string]any{
				"operations": []any{},
			},
			want: true,
		},
		{
			name: "compute list with kind suffix List",
			body: map[string]any{
				"kind":  "compute#networkList",
				"items": []any{},
			},
			want: true,
		},
		{
			name: "list response with items and no top-level name",
			body: map[string]any{
				"items": []any{
					map[string]any{"name": "item-1"},
				},
			},
			want: true,
		},
		{
			name: "single resource with items field and top-level name (e.g. NetworkSecurityAddressGroup)",
			body: map[string]any{
				"name":     "projects/test-proj/locations/global/addressGroups/ag-1",
				"capacity": 100,
				"items":    []any{"10.0.0.0/24"},
			},
			want: false,
		},
		{
			name: "standard single resource without items",
			body: map[string]any{
				"name": "projects/test-proj/locations/us-central1/clusters/c-1",
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isListOrOperationResponse(tc.body); got != tc.want {
				t.Errorf("isListOrOperationResponse(%v) = %v, want %v", tc.body, got, tc.want)
			}
		})
	}
}

func TestShouldProbeURL(t *testing.T) {
	tests := []struct {
		url  string
		want bool
	}{
		{
			url:  "https://discoveryengine.googleapis.com/v1/projects/my-project/locations/global/dataStores/my-ds/schemas/default_schema",
			want: true,
		},
		{
			url:  "https://compute.googleapis.com/compute/v1/projects/my-project/global/networks/my-net",
			want: true,
		},
		{
			url:  "https://www.googleapis.com/discovery/v1/apis",
			want: false,
		},
		{
			url:  "https://discovery.googleapis.com/v1/apis",
			want: false,
		},
		{
			url:  "https://discoveryengine.googleapis.com/v1/projects/123/locations/global/collections/default_collection/dataStores/ds/schemas/s/operations/op-123",
			want: false,
		},
		{
			url:  "https://oauth2.googleapis.com/token",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.url, func(t *testing.T) {
			if got := shouldProbeURL(tc.url); got != tc.want {
				t.Errorf("shouldProbeURL(%q) = %v, want %v", tc.url, got, tc.want)
			}
		})
	}
}

func TestGRPCEventToRESTProbeURL(t *testing.T) {
	tests := []struct {
		name    string
		entry   *test.LogEntry
		wantURL string
		wantOK  bool
	}{
		{
			name: "BigtableInstanceAdmin GetMemoryLayer",
			entry: &test.LogEntry{
				Request: test.Request{
					Method: "GRPC",
					URL:    "/google.bigtable.admin.v2.BigtableInstanceAdmin/GetMemoryLayer",
					Body:   `{"name":"projects/my-proj/instances/my-inst/clusters/my-cluster/memoryLayer"}`,
				},
				Response: test.Response{
					Status: "OK",
					Body:   `{"name":"projects/my-proj/instances/my-inst/clusters/my-cluster/memoryLayer","state":"READY"}`,
				},
			},
			wantURL: "https://bigtableadmin.googleapis.com/v2/projects/my-proj/instances/my-inst/clusters/my-cluster/memoryLayer",
			wantOK:  true,
		},
		{
			name: "BigtableTableAdmin GetAuthorizedView",
			entry: &test.LogEntry{
				Request: test.Request{
					Method: "GRPC",
					URL:    "/google.bigtable.admin.v2.BigtableTableAdmin/GetAuthorizedView",
					Body:   `{"name":"projects/my-proj/instances/my-inst/tables/my-tbl/authorizedViews/my-av"}`,
				},
				Response: test.Response{
					Status: "OK",
					Body:   `{"name":"projects/my-proj/instances/my-inst/tables/my-tbl/authorizedViews/my-av"}`,
				},
			},
			wantURL: "https://bigtableadmin.googleapis.com/v2/projects/my-proj/instances/my-inst/tables/my-tbl/authorizedViews/my-av",
			wantOK:  true,
		},
		{
			name: "DataplexService GetLake",
			entry: &test.LogEntry{
				Request: test.Request{
					Method: "GRPC",
					URL:    "/google.cloud.dataplex.v1.DataplexService/GetLake",
					Body:   `{"name":"projects/my-proj/locations/us-central1/lakes/my-lake"}`,
				},
				Response: test.Response{
					Status: "OK",
					Body:   `{"name":"projects/my-proj/locations/us-central1/lakes/my-lake"}`,
				},
			},
			wantURL: "https://dataplex.googleapis.com/v1/projects/my-proj/locations/us-central1/lakes/my-lake",
			wantOK:  true,
		},
		{
			name: "Non-Get gRPC method (UpdateMemoryLayer) is skipped",
			entry: &test.LogEntry{
				Request: test.Request{
					Method: "GRPC",
					URL:    "/google.bigtable.admin.v2.BigtableInstanceAdmin/UpdateMemoryLayer",
					Body:   `{"memoryLayer":{"name":"projects/my-proj/instances/my-inst/clusters/my-cluster/memoryLayer"}}`,
				},
				Response: test.Response{
					Status: "OK",
					Body:   `{"name":"operations/123"}`,
				},
			},
			wantURL: "",
			wantOK:  false,
		},
		{
			name: "Operations GetOperation is skipped",
			entry: &test.LogEntry{
				Request: test.Request{
					Method: "GRPC",
					URL:    "/google.longrunning.Operations/GetOperation",
					Body:   `{"name":"operations/projects/my-proj/operations/123"}`,
				},
				Response: test.Response{
					Status: "OK",
					Body:   `{"name":"operations/projects/my-proj/operations/123","done":true}`,
				},
			},
			wantURL: "",
			wantOK:  false,
		},
		{
			name: "Non-OK gRPC status is skipped",
			entry: &test.LogEntry{
				Request: test.Request{
					Method: "GRPC",
					URL:    "/google.bigtable.admin.v2.BigtableInstanceAdmin/GetMemoryLayer",
					Body:   `{"name":"projects/my-proj/instances/my-inst/clusters/my-cluster/memoryLayer"}`,
				},
				Response: test.Response{
					Status: "NOT_FOUND",
					Body:   `{}`,
				},
			},
			wantURL: "",
			wantOK:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotURL, gotOK := grpcEventToRESTProbeURL(tc.entry)
			if gotOK != tc.wantOK || gotURL != tc.wantURL {
				t.Errorf("grpcEventToRESTProbeURL() = (%q, %v), want (%q, %v)", gotURL, gotOK, tc.wantURL, tc.wantOK)
			}
		})
	}
}
