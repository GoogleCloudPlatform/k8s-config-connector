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

package specdiff

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/classifier"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestValuesCompatible(t *testing.T) {
	tests := []struct {
		name        string
		foundVal    any
		expectedVal any
		createVal   any
		diffPath    string
		want        bool
	}{
		{
			name:        "mismatched integers 150 vs 200",
			foundVal:    float64(150),
			expectedVal: int64(200),
			createVal:   int64(100),
			diffPath:    "spec.priority",
			want:        false,
		},
		{
			name:        "matching integers 200 vs 200",
			foundVal:    float64(200),
			expectedVal: int64(200),
			createVal:   int64(100),
			diffPath:    "spec.priority",
			want:        true,
		},
		{
			name:        "mismatched booleans true vs false",
			foundVal:    true,
			expectedVal: false,
			createVal:   nil,
			diffPath:    "spec.enabled",
			want:        false,
		},
		{
			name:        "mismatched string slices",
			foundVal:    []any{"a"},
			expectedVal: []any{"b"},
			createVal:   []any{"c"},
			diffPath:    "spec.scopes",
			want:        false,
		},
		{
			name:        "enum integer vs string name",
			foundVal:    float64(2),
			expectedVal: "ACTIVE",
			createVal:   "INACTIVE",
			diffPath:    "spec.state",
			want:        true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := valuesCompatible(tc.foundVal, tc.expectedVal, tc.createVal, tc.diffPath)

			// Assert
			if got != tc.want {
				t.Errorf("valuesCompatible(%v, %v, %v, %q) = %v, want %v",
					tc.foundVal, tc.expectedVal, tc.createVal, tc.diffPath, got, tc.want)
			}
		})
	}
}

func TestUpdateMaskCoversPath(t *testing.T) {
	tests := []struct {
		name     string
		mask     []string
		diffPath string
		want     bool
	}{
		{
			name:     "rejects child mask for parent clearing",
			mask:     []string{"network_config.subnetwork"},
			diffPath: "spec.networkConfig",
			want:     false,
		},
		{
			name:     "accepts exact child mask",
			mask:     []string{"network_config.subnetwork"},
			diffPath: "spec.networkConfig.subnetwork",
			want:     true,
		},
		{
			name:     "accepts parent mask covering child path",
			mask:     []string{"network_config"},
			diffPath: "spec.networkConfig.subnetwork",
			want:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := updateMaskCoversPath(tc.mask, tc.diffPath)

			// Assert
			if got != tc.want {
				t.Errorf("updateMaskCoversPath(%v, %q) = %v, want %v", tc.mask, tc.diffPath, got, tc.want)
			}
		})
	}
}

func TestDiffer_Compute(t *testing.T) {
	makeFixture := func() *fixture.Data {
		createObj := &unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "example.cnrm.cloud.google.com/v1alpha1",
				"kind":       "ExampleWidget",
				"metadata":   map[string]any{"name": "widget-1"},
				"spec": map[string]any{
					"description": "initial",
					"priority":    int64(100),
				},
			},
		}
		updateObj := &unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": "example.cnrm.cloud.google.com/v1alpha1",
				"kind":       "ExampleWidget",
				"metadata":   map[string]any{"name": "widget-1"},
				"spec": map[string]any{
					"description": "updated",
					"priority":    int64(200),
				},
			},
		}
		return &fixture.Data{
			PrimaryCreate: &fixture.KRMResource{Kind: "ExampleWidget", Name: "widget-1", Obj: createObj},
			PrimaryUpdate: &fixture.KRMResource{Kind: "ExampleWidget", Name: "widget-1", Obj: updateObj},
		}
	}

	tests := []struct {
		name                 string
		primary              []*classifier.ClassifiedEntry
		wantSentByPath       map[string]bool
		wantFinalMatchByPath map[string]bool
	}{
		{
			name: "update sent in PATCH and matched in final GET",
			primary: []*classifier.ClassifiedEntry{
				{
					Index:      1,
					Phase:      logparser.PhaseUpdate,
					Role:       classifier.RolePrimary,
					UpdateMask: []string{"description", "priority"},
					Entry: &logparser.Entry{
						Request: logparser.HTTPRequest{
							Method: "PATCH",
							URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-1?updateMask=description,priority",
							Body:   `{"description": "updated", "priority": 200}`,
						},
						Response: logparser.HTTPResponse{
							Status:     "200 OK",
							StatusCode: 200,
							Body:       `{"description": "updated", "priority": 200}`,
						},
					},
				},
			},
			wantSentByPath: map[string]bool{
				"spec.description": true,
				"spec.priority":    true,
			},
			wantFinalMatchByPath: map[string]bool{
				"spec.description": true,
				"spec.priority":    true,
			},
		},
		{
			name: "dropped field missing from updateMask",
			primary: []*classifier.ClassifiedEntry{
				{
					Index:      1,
					Phase:      logparser.PhaseUpdate,
					Role:       classifier.RolePrimary,
					UpdateMask: []string{"description"},
					Entry: &logparser.Entry{
						Request: logparser.HTTPRequest{
							Method: "PATCH",
							URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-1?updateMask=description",
							Body:   `{"description": "updated"}`,
						},
						Response: logparser.HTTPResponse{
							Status:     "200 OK",
							StatusCode: 200,
							Body:       `{"description": "updated", "priority": 100}`,
						},
					},
				},
			},
			wantSentByPath: map[string]bool{
				"spec.description": true,
				"spec.priority":    false,
			},
			wantFinalMatchByPath: map[string]bool{
				"spec.description": true,
				"spec.priority":    false,
			},
		},
		{
			name: "end-state value mismatch in final GET",
			primary: []*classifier.ClassifiedEntry{
				{
					Index:      1,
					Phase:      logparser.PhaseUpdate,
					Role:       classifier.RolePrimary,
					UpdateMask: []string{"description", "priority"},
					Entry: &logparser.Entry{
						Request: logparser.HTTPRequest{
							Method: "PATCH",
							URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-1?updateMask=description,priority",
							Body:   `{"description": "updated", "priority": 200}`,
						},
						Response: logparser.HTTPResponse{
							Status:     "200 OK",
							StatusCode: 200,
							Body:       `{"name": "projects/p/locations/l/operations/op-1", "done": true}`,
						},
					},
				},
				{
					Index: 2,
					Phase: logparser.PhaseUpdate,
					Role:  classifier.RolePrimary,
					Entry: &logparser.Entry{
						Request: logparser.HTTPRequest{
							Method: "GET",
							URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-1",
						},
						Response: logparser.HTTPResponse{
							Status:     "200 OK",
							StatusCode: 200,
							Body:       `{"description": "updated", "priority": 100}`,
						},
					},
				},
			},
			wantSentByPath: map[string]bool{
				"spec.description": true,
				"spec.priority":    true,
			},
			wantFinalMatchByPath: map[string]bool{
				"spec.description": true,
				"spec.priority":    false,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			differ := New(makeFixture(), tc.primary)

			// Act
			items := differ.Compute()

			// Assert
			if len(items) != 2 {
				t.Fatalf("expected 2 diff items, got %d", len(items))
			}
			for _, item := range items {
				if wantSent, ok := tc.wantSentByPath[item.Path]; ok && item.SentInUpdateRPC != wantSent {
					t.Errorf("item[%s].SentInUpdateRPC = %v, want %v (%s)", item.Path, item.SentInUpdateRPC, wantSent, item.UpdateRPCDetail)
				}
				if wantFinal, ok := tc.wantFinalMatchByPath[item.Path]; ok && item.FinalStateMatch != wantFinal {
					t.Errorf("item[%s].FinalStateMatch = %v, want %v (%s)", item.Path, item.FinalStateMatch, wantFinal, item.FinalStateDetail)
				}
			}
		})
	}
}
