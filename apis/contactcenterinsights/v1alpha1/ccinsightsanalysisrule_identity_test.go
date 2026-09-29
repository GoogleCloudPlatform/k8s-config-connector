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

package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCCInsightsAnalysisRuleIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *CCInsightsAnalysisRuleIdentity
		hasError bool
	}{
		{
			name:  "Full resource name",
			input: "projects/my-project/locations/us-central1/analysisRules/my-analysis-rule",
			expected: &CCInsightsAnalysisRuleIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				AnalysisRule: "my-analysis-rule",
			},
			hasError: false,
		},
		{
			name:  "Full resource name with host",
			input: "contactcenterinsights.googleapis.com/projects/my-project/locations/us-central1/analysisRules/my-analysis-rule",
			expected: &CCInsightsAnalysisRuleIdentity{
				Project:      "my-project",
				Location:     "us-central1",
				AnalysisRule: "my-analysis-rule",
			},
			hasError: false,
		},
		{
			name:     "Invalid format",
			input:    "projects/my-project/locations/us-central1/invalid/my-analysis-rule",
			expected: nil,
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &CCInsightsAnalysisRuleIdentity{}
			err := id.FromExternal(tc.input)
			if tc.hasError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.expected, id); diff != "" {
				t.Errorf("CCInsightsAnalysisRuleIdentity mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCCInsightsAnalysisRuleRef_ValidateExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
	}{
		{
			name:    "valid reference",
			ref:     "projects/my-project/locations/us-central1/analysisRules/my-analysis-rule",
			wantErr: false,
		},
		{
			name:    "invalid prefix",
			ref:     "invalid/my-project/locations/us-central1/analysisRules/my-analysis-rule",
			wantErr: true,
		},
		{
			name:    "missing location",
			ref:     "projects/my-project/analysisRules/my-analysis-rule",
			wantErr: true,
		},
		{
			name:    "missing analysis rule",
			ref:     "projects/my-project/locations/us-central1",
			wantErr: true,
		},
		{
			name:    "empty string",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CCInsightsAnalysisRuleRef{}
			if err := r.ValidateExternal(tt.ref); (err != nil) != tt.wantErr {
				t.Errorf("CCInsightsAnalysisRuleRef.ValidateExternal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
