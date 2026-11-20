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
		{
			name: "composite slice with omitted empty slice, int64 string coercion, and server default field",
			foundVal: []any{
				map[string]any{
					"weight":        "2",
					"serverDefault": "Y1",
					"logsPanel":     map[string]any{"filter": "foo"},
				},
			},
			expectedVal: []any{
				map[string]any{
					"weight": int64(2),
					"logsPanel": map[string]any{
						"filter":        "foo",
						"resourceNames": []any{},
					},
				},
			},
			createVal: []any{
				map[string]any{
					"weight": int64(1),
				},
			},
			diffPath: "spec.columnLayout.columns",
			want:     true,
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

func TestDiffer_Compute_GoldenObjectRetainsValue(t *testing.T) {
	createObj := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "example.cnrm.cloud.google.com/v1beta1",
			"kind":       "ExampleSecretVersion",
			"metadata": map[string]any{
				"name": "sv-1",
				"labels": map[string]any{
					"app.kubernetes.io/name": "retained-app",
				},
			},
			"spec": map[string]any{
				"enabled": true,
				"secretData": map[string]any{
					"valueFrom": map[string]any{
						"secretKeyRef": map[string]any{
							"key":  "secretData",
							"name": "dep-secret",
						},
					},
				},
				"droppedField": "should-still-diff",
			},
		},
	}
	updateObj := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "example.cnrm.cloud.google.com/v1beta1",
			"kind":       "ExampleSecretVersion",
			"metadata": map[string]any{
				"name": "sv-1",
			},
			"spec": map[string]any{
				"enabled":    false,
				"resourceID": "sv-1",
			},
		},
	}
	goldenObj := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "example.cnrm.cloud.google.com/v1beta1",
			"kind":       "ExampleSecretVersion",
			"metadata": map[string]any{
				"name": "sv-1",
				"labels": map[string]any{
					"app.kubernetes.io/name": "retained-app",
				},
			},
			"spec": map[string]any{
				"enabled": false,
				"secretData": map[string]any{
					"valueFrom": map[string]any{
						"secretKeyRef": map[string]any{
							"key":  "secretData",
							"name": "dep-secret",
						},
					},
				},
			},
		},
	}

	fix := &fixture.Data{
		PrimaryCreate: &fixture.KRMResource{Kind: "ExampleSecretVersion", Name: "sv-1", Obj: createObj},
		PrimaryUpdate: &fixture.KRMResource{Kind: "ExampleSecretVersion", Name: "sv-1", Obj: updateObj},
		GoldenObj:     goldenObj,
	}
	items := New(fix, nil).Compute()

	gotPaths := make(map[string]bool)
	for _, it := range items {
		gotPaths[it.Path] = true
	}
	if gotPaths["spec.secretData.valueFrom.secretKeyRef"] {
		t.Errorf("expected retained *Ref field in GoldenObj to be skipped, got items: %+v", items)
	}
	if gotPaths["metadata.labels.app.kubernetes.io/name"] {
		t.Errorf("expected retained dotted label in GoldenObj to be skipped, got items: %+v", items)
	}
	if !gotPaths["spec.droppedField"] {
		t.Errorf("expected non-retained omitted field spec.droppedField to remain in diffs, got: %+v", items)
	}
	if !gotPaths["spec.enabled"] {
		t.Errorf("expected changed field spec.enabled to remain in diffs, got: %+v", items)
	}
	if !IsClientSideKRMField("spec.resourceID") {
		t.Errorf("expected IsClientSideKRMField(spec.resourceID) = true")
	}
}

func TestRefKeysMatchAndLookup(t *testing.T) {
	tests := []struct {
		name      string
		body      map[string]any
		diffPath  string
		wantVal   any
		wantFound bool
	}{
		{
			name:      "kmsKeyRef matches kmsKeyName",
			body:      map[string]any{"kmsKeyName": "projects/p/locations/l/keyRings/r/cryptoKeys/k"},
			diffPath:  "spec.kmsKeyRef",
			wantVal:   "projects/p/locations/l/keyRings/r/cryptoKeys/k",
			wantFound: true,
		},
		{
			name:      "keyRef matches keyName",
			body:      map[string]any{"keyName": "projects/p/locations/l/keyRings/r/cryptoKeys/k"},
			diffPath:  "spec.keyRef",
			wantVal:   "projects/p/locations/l/keyRings/r/cryptoKeys/k",
			wantFound: true,
		},
		{
			name:      "exact match takes precedence over fuzzy ref match",
			body:      map[string]any{"network": "exact-net", "networkId": "fuzzy-id"},
			diffPath:  "spec.networkRef",
			wantVal:   "exact-net",
			wantFound: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := findValueForDiffPath(tc.body, tc.diffPath)
			if ok != tc.wantFound || got != tc.wantVal {
				t.Errorf("findValueForDiffPath(%v, %q) = (%v, %v), want (%v, %v)",
					tc.body, tc.diffPath, got, ok, tc.wantVal, tc.wantFound)
			}
		})
	}
}

func TestMapContainsKeyedSliceAndDurationMatch(t *testing.T) {
	if !ValuesMatch("3600s", "1h") {
		t.Errorf("expected ValuesMatch(\"3600s\", \"1h\") = true")
	}
	if ValuesMatch("60s", "2m") {
		t.Errorf("expected ValuesMatch(\"60s\", \"2m\") = false")
	}

	tests := []struct {
		name   string
		actual map[string]any
		expect []any
		want   bool
	}{
		{
			name: "map value repeats key field",
			actual: map[string]any{
				"item-1": map[string]any{"name": "item-1", "weight": float64(10)},
			},
			expect: []any{
				map[string]any{"name": "item-1", "weight": int64(10)},
			},
			want: true,
		},
		{
			name: "map value omits key field present as map key",
			actual: map[string]any{
				"item-1": map[string]any{"weight": float64(10)},
			},
			expect: []any{
				map[string]any{"name": "item-1", "weight": int64(10)},
			},
			want: true,
		},
		{
			name: "scalar map value with key/value entries",
			actual: map[string]any{
				"env": "prod",
			},
			expect: []any{
				map[string]any{"key": "env", "value": "prod"},
			},
			want: true,
		},
		{
			name: "extra stale key in actual map fails length check",
			actual: map[string]any{
				"item-1": map[string]any{"weight": float64(10)},
				"stale":  map[string]any{"weight": float64(5)},
			},
			expect: []any{
				map[string]any{"name": "item-1", "weight": int64(10)},
			},
			want: false,
		},
		{
			name: "duplicate key in expected slice cannot mask unchecked actual entry",
			actual: map[string]any{
				"item-1": map[string]any{"weight": float64(10)},
				"extra":  map[string]any{"weight": float64(5)},
			},
			expect: []any{
				map[string]any{"name": "item-1", "weight": int64(10)},
				map[string]any{"name": "item-1", "weight": int64(10)},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := mapContainsKeyedSlice(tc.actual, tc.expect); got != tc.want {
				t.Errorf("mapContainsKeyedSlice() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestResponseBodiesMatchDiffItem_WriteOnlyScoping(t *testing.T) {
	tests := []struct {
		name       string
		bodies     []map[string]any
		item       Item
		wantMatch  bool
		wantDetail string
	}{
		{
			name:   "omitted true write-only password passes",
			bodies: []map[string]any{{"username": "admin"}},
			item: Item{
				Path:            "spec.password.valueFrom.secretKeyRef",
				CreateVal:       "secret-1",
				UpdateVal:       "secret-2",
				SentInUpdateRPC: true,
			},
			wantMatch:  true,
			wantDetail: "Yes (write-only)",
		},
		{
			name:   "redacted write-only password passes",
			bodies: []map[string]any{{"password": "REDACTED"}},
			item: Item{
				Path:            "spec.password.value",
				CreateVal:       "old-pass",
				UpdateVal:       "new-pass",
				SentInUpdateRPC: true,
			},
			wantMatch:  true,
			wantDetail: "Yes (write-only)",
		},
		{
			name:   "unredacted stale password in response still fails",
			bodies: []map[string]any{{"password": "old-pass"}},
			item: Item{
				Path:            "spec.password.value",
				CreateVal:       "old-pass",
				UpdateVal:       "new-pass",
				SentInUpdateRPC: true,
			},
			wantMatch:  false,
			wantDetail: "No (mismatch)",
		},
		{
			name:   "readable passwordValidationPolicy.minLength mismatch is not suppressed",
			bodies: []map[string]any{{"passwordValidationPolicy": map[string]any{"minLength": float64(8)}}},
			item: Item{
				Path:            "spec.passwordValidationPolicy.minLength",
				CreateVal:       int64(8),
				UpdateVal:       int64(12),
				SentInUpdateRPC: true,
			},
			wantMatch:  false,
			wantDetail: "No (mismatch)",
		},
		{
			name:   "readable tokenURL mismatch is not suppressed",
			bodies: []map[string]any{{"tokenURL": "https://old.example.com/token"}},
			item: Item{
				Path:            "spec.tokenURL",
				CreateVal:       "https://old.example.com/token",
				UpdateVal:       "https://new.example.com/token",
				SentInUpdateRPC: true,
			},
			wantMatch:  false,
			wantDetail: "No (mismatch)",
		},
		{
			name:   "boolean token toggle mismatch is not suppressed",
			bodies: []map[string]any{{"idToken": true}},
			item: Item{
				Path:            "spec.idToken",
				CreateVal:       true,
				UpdateVal:       false,
				SentInUpdateRPC: true,
			},
			wantMatch:  false,
			wantDetail: "No (mismatch)",
		},
		{
			name:   "readable secretRef mismatch is not suppressed",
			bodies: []map[string]any{{"secret": "projects/p/secrets/old"}},
			item: Item{
				Path:            "spec.secretRef",
				CreateVal:       "projects/p/secrets/old",
				UpdateVal:       "projects/p/secrets/new",
				SentInUpdateRPC: true,
			},
			wantMatch:  false,
			wantDetail: "No (mismatch)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotMatch, gotDetail := responseBodiesMatchDiffItem(tc.bodies, tc.item)
			if gotMatch != tc.wantMatch || gotDetail != tc.wantDetail {
				t.Errorf("responseBodiesMatchDiffItem() = (%v, %q), want (%v, %q)",
					gotMatch, gotDetail, tc.wantMatch, tc.wantDetail)
			}
		})
	}
}
