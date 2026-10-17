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

package resourceoverrides_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestLoggingLogSinkPreTerraformExport(t *testing.T) {
	tests := []struct {
		name         string
		attributes   map[string]string
		expectedType string
		expectErr    bool
	}{
		{
			name: "project scoped sink",
			attributes: map[string]string{
				"project": "sample-project",
			},
			expectedType: "google_logging_project_sink",
			expectErr:    false,
		},
		{
			name: "folder scoped sink",
			attributes: map[string]string{
				"folder": "folders/123456789012",
			},
			expectedType: "google_logging_folder_sink",
			expectErr:    false,
		},
		{
			name: "organization scoped sink",
			attributes: map[string]string{
				"org_id": "organizations/123456789012",
			},
			expectedType: "google_logging_organization_sink",
			expectErr:    false,
		},
		{
			name: "both project and folder set",
			attributes: map[string]string{
				"project": "sample-project",
				"folder":  "folders/123456789012",
			},
			expectErr: true,
		},
		{
			name: "no scope set",
			attributes: map[string]string{
				"name": "sample-sink",
			},
			expectErr: true,
		},
	}

	ro := resourceoverrides.GetLoggingLogSinkResourceOverrides()
	if len(ro.Overrides) == 0 || ro.Overrides[0].PreTerraformExport == nil {
		t.Fatalf("expected PreTerraformExport override for LoggingLogSink")
	}
	preTerraformExport := ro.Overrides[0].PreTerraformExport

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			op := &operations.TerraformExport{
				TerraformState: &terraform.InstanceState{
					Attributes: tc.attributes,
				},
				TerraformInfo: &terraform.InstanceInfo{},
			}
			err := preTerraformExport(context.Background(), op)
			if tc.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tc.expectErr && op.TerraformInfo.Type != tc.expectedType {
				t.Fatalf("expected type %q, got %q", tc.expectedType, op.TerraformInfo.Type)
			}
		})
	}
}
