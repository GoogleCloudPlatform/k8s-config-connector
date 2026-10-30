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

package validator

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/classifier"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/specdiff"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestValidator_Validate(t *testing.T) {
	baseFixture := func() *fixture.Data {
		return &fixture.Data{
			FixtureName: "widget-test",
			PrimaryCreate: &fixture.KRMResource{
				Kind: "ExampleWidget",
				Name: "widget-1",
				Obj:  &unstructured.Unstructured{Object: map[string]any{"kind": "ExampleWidget"}},
			},
			PrimaryUpdate: &fixture.KRMResource{
				Kind: "ExampleWidget",
				Name: "widget-1",
				Obj:  &unstructured.Unstructured{Object: map[string]any{"kind": "ExampleWidget"}},
			},
		}
	}

	createOK := &classifier.ClassifiedEntry{
		Index:         1,
		Phase:         logparser.PhaseCreate,
		Role:          classifier.RolePrimary,
		ShortEndpoint: "widgets?widgetId=widget-1",
		Entry: &logparser.Entry{
			Request: logparser.HTTPRequest{
				Method: "POST",
				URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets?widgetId=widget-1",
				Body:   "{\n  \"description\": \"initial\"\n}",
			},
			Response: logparser.HTTPResponse{
				Status:     "200 OK",
				StatusCode: 200,
				Body:       "{\n  \"description\": \"initial\"\n}",
			},
		},
	}

	updateOK := &classifier.ClassifiedEntry{
		Index:         2,
		Phase:         logparser.PhaseUpdate,
		Role:          classifier.RolePrimary,
		ShortEndpoint: "widgets/widget-1?updateMask=description",
		UpdateMask:    []string{"description"},
		Entry: &logparser.Entry{
			Request: logparser.HTTPRequest{
				Method: "PATCH",
				URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-1?updateMask=description",
				Body:   "{\n  \"description\": \"updated\"\n}",
			},
			Response: logparser.HTTPResponse{
				Status:     "200 OK",
				StatusCode: 200,
				Body:       "{\n  \"description\": \"updated\"\n}",
			},
		},
	}

	deleteOK := &classifier.ClassifiedEntry{
		Index:         3,
		Phase:         logparser.PhaseDelete,
		Role:          classifier.RolePrimary,
		ShortEndpoint: "widgets/widget-1",
		Entry: &logparser.Entry{
			Request: logparser.HTTPRequest{
				Method: "DELETE",
				URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-1",
			},
			Response: logparser.HTTPResponse{
				Status:     "200 OK",
				StatusCode: 200,
			},
		},
	}

	cleanDiff := []specdiff.Item{
		{
			Path:             "spec.description",
			CreateVal:        "initial",
			UpdateVal:        "updated",
			SentInUpdateRPC:  true,
			UpdateRPCDetail:  "Yes (#2 `PATCH`)",
			FinalStateMatch:  true,
			FinalStateDetail: "Yes",
		},
	}

	badHeaders := make(http.Header)
	badHeaders.Set("Authorization", "Bearer dummy-test-token")
	badHeaders.Set("X-Goog-Request-Params", "parent=projects/p/locations/us-east1")

	tests := []struct {
		name               string
		fix                *fixture.Data
		entries            []*classifier.ClassifiedEntry
		diffs              []specdiff.Item
		wantPassed         bool
		wantFailedCheckIDs []string
		wantReasonSubstrs  []string
	}{
		{
			name:       "clean lifecycle passes V1-V7",
			fix:        baseFixture(),
			entries:    []*classifier.ClassifiedEntry{createOK, updateOK, deleteOK},
			diffs:      cleanDiff,
			wantPassed: true,
		},
		{
			name: "V1 fails on 400 Bad Request during create",
			fix:  baseFixture(),
			entries: []*classifier.ClassifiedEntry{
				{
					Index:         1,
					Phase:         logparser.PhaseCreate,
					Role:          classifier.RolePrimary,
					ShortEndpoint: "widgets?widgetId=widget-1",
					Entry: &logparser.Entry{
						Request: logparser.HTTPRequest{
							Method: "POST",
							URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets?widgetId=widget-1",
							Body:   "{\n  \"description\": \"initial\"\n}",
						},
						Response: logparser.HTTPResponse{
							Status:     "400 Bad Request",
							StatusCode: 400,
						},
					},
				},
				updateOK,
				deleteOK,
			},
			diffs:              cleanDiff,
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V1"},
			wantReasonSubstrs:  []string{"400 Bad Request"},
		},
		{
			name:    "V2 and V3 fail on silent update no-op",
			fix:     baseFixture(),
			entries: []*classifier.ClassifiedEntry{createOK, deleteOK},
			diffs: []specdiff.Item{
				{
					Path:             "spec.description",
					CreateVal:        "initial",
					UpdateVal:        "updated",
					SentInUpdateRPC:  false,
					UpdateRPCDetail:  "No (0 Update RPCs)",
					FinalStateMatch:  false,
					FinalStateDetail: "No",
				},
			},
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V2", "V3"},
			wantReasonSubstrs:  []string{"SILENT UPDATE NO-OP DETECTED"},
		},
		{
			name:               "V4 fails on missing DELETE RPC",
			fix:                baseFixture(),
			entries:            []*classifier.ClassifiedEntry{createOK, updateOK},
			diffs:              cleanDiff,
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V4"},
			wantReasonSubstrs:  []string{"MISSING DELETE RPC"},
		},
		{
			name: "V5 fails on duplicate consecutive PATCH",
			fix:  baseFixture(),
			entries: []*classifier.ClassifiedEntry{
				createOK,
				updateOK,
				{
					Index:         3,
					Phase:         logparser.PhaseUpdate,
					Role:          classifier.RolePrimary,
					ShortEndpoint: "widgets/widget-1?updateMask=description",
					UpdateMask:    []string{"description"},
					Entry: &logparser.Entry{
						Request: logparser.HTTPRequest{
							Method: "PATCH",
							URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-1?updateMask=description",
							Body:   "{\n  \"description\": \"updated\"\n}",
						},
						Response: logparser.HTTPResponse{
							Status:     "200 OK",
							StatusCode: 200,
							Body:       "{\n  \"description\": \"updated\"\n}",
						},
					},
				},
				deleteOK,
			},
			diffs:              cleanDiff,
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V5"},
			wantReasonSubstrs:  []string{"Duplicate consecutive mutation RPCs"},
		},
		{
			name: "V6 and V7 fail on mockgcp leak, Authorization header, non-canonical JSON, and X-Goog-Request-Params mismatch",
			fix: &fixture.Data{
				FixtureName:   "widget-bad",
				RawLogContent: "// (mockgcp) leaked signature",
				PrimaryCreate: &fixture.KRMResource{
					Kind: "ExampleWidget",
					Name: "widget-1",
				},
			},
			entries: []*classifier.ClassifiedEntry{
				{
					Index:         1,
					Phase:         logparser.PhaseCreate,
					Role:          classifier.RolePrimary,
					ShortEndpoint: "locations/us-central1/widgets",
					Entry: &logparser.Entry{
						Request: logparser.HTTPRequest{
							Method: "POST",
							URL:    "https://example.googleapis.com/v1/projects/p/locations/us-central1/widgets",
							Header: badHeaders,
							Body:   `{"b": 1, "a": 2}`,
						},
						Response: logparser.HTTPResponse{
							Status:     "200 OK",
							StatusCode: 200,
						},
					},
				},
				deleteOK,
			},
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V6", "V7"},
			wantReasonSubstrs:  []string{"(mockgcp)", "Authorization", "Non-canonical JSON", "X-Goog-Request-Params"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			v := New(tc.fix, tc.entries, tc.entries, tc.diffs)

			// Act
			checks, passed := v.Validate()

			// Assert
			if passed != tc.wantPassed {
				t.Fatalf("passed = %v, want %v (checks: %+v)", passed, tc.wantPassed, checks)
			}
			var gotFailedIDs []string
			var combined strings.Builder
			for _, c := range checks {
				if !c.Passed {
					gotFailedIDs = append(gotFailedIDs, c.ID)
					combined.WriteString(c.Summary + " " + c.Hint + "\n")
				}
			}
			if !reflect.DeepEqual(gotFailedIDs, tc.wantFailedCheckIDs) {
				t.Errorf("failed check IDs = %v, want %v", gotFailedIDs, tc.wantFailedCheckIDs)
			}
			for _, sub := range tc.wantReasonSubstrs {
				if !strings.Contains(combined.String(), sub) {
					t.Errorf("combined failure reasons %q missing substring %q", combined.String(), sub)
				}
			}
		})
	}
}
