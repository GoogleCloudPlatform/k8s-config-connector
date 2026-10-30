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

package printer

import (
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/classifier"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/specdiff"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/validator"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestPrinter_Render(t *testing.T) {
	baseEntry := &classifier.ClassifiedEntry{
		Index:         1,
		Phase:         logparser.PhaseCreate,
		Role:          classifier.RolePrimary,
		ShortEndpoint: "widgets?widgetId=widget-sample",
		Details:       "`description`: `\"v1\"`",
		Entry: &logparser.Entry{
			Request: logparser.HTTPRequest{
				Method: "POST",
				URL:    "https://example.googleapis.com/v1/projects/p/locations/l/widgets?widgetId=widget-sample",
				Body:   "{\n  \"description\": \"v1\"\n}",
			},
			Response: logparser.HTTPResponse{
				Status:     "200 OK",
				StatusCode: 200,
				Body:       "{\n  \"description\": \"v1\"\n}",
			},
		},
	}

	tests := []struct {
		name              string
		report            *ReportData
		verbosity         string
		wantMDSubstrs     []string
		wantTermSubstrs   []string
		forbidTermSubstrs []string
	}{
		{
			name: "passing report markdown and summary terminal output",
			report: &ReportData{
				Fixture: &fixture.Data{
					FixtureName: "widget-pass",
					PrimaryCreate: &fixture.KRMResource{
						APIVersion: "example.cnrm.cloud.google.com/v1alpha1",
						Kind:       "ExampleWidget",
						Name:       "widget-pass",
						Obj:        &unstructured.Unstructured{},
					},
				},
				TotalEntries:   1,
				PrimaryEntries: []*classifier.ClassifiedEntry{baseEntry},
				Checks: []validator.Check{
					{ID: "V1", Title: "Create", Passed: true, Summary: "Primary resource created."},
				},
				Passed: true,
			},
			verbosity: "summary",
			wantMDSubstrs: []string{
				"# HTTP Lifecycle Report: `widget-pass`",
				"- **Validation Status:** PASS (1/1 checks passed)",
				"- **Integrity Hash:** `sha256:",
			},
			wantTermSubstrs: []string{
				"# HTTP Lifecycle Report: `widget-pass`",
				"- **Validation Status:** PASS (1/1 checks passed)",
			},
			forbidTermSubstrs: []string{
				"## 4. Detailed RPC Payloads",
				"<hint_for_agent>",
			},
		},
		{
			name: "failing report markdown and mutations terminal output with hint_for_agent",
			report: &ReportData{
				Fixture: &fixture.Data{
					FixtureName: "widget-sample",
					PrimaryCreate: &fixture.KRMResource{
						APIVersion: "example.cnrm.cloud.google.com/v1alpha1",
						Kind:       "ExampleWidget",
						Name:       "widget-sample",
						Obj:        &unstructured.Unstructured{},
					},
					PrimaryUpdate: &fixture.KRMResource{
						APIVersion: "example.cnrm.cloud.google.com/v1alpha1",
						Kind:       "ExampleWidget",
						Name:       "widget-sample",
						Obj:        &unstructured.Unstructured{},
					},
				},
				TotalEntries:   1,
				PrimaryEntries: []*classifier.ClassifiedEntry{baseEntry},
				SpecDiff: []specdiff.Item{
					{
						Path:             "spec.description",
						CreateVal:        "v1",
						UpdateVal:        "v2",
						SentInUpdateRPC:  false,
						UpdateRPCDetail:  "No (0 Update RPCs)",
						FinalStateMatch:  false,
						FinalStateDetail: "No (mismatch)",
					},
				},
				Checks: []validator.Check{
					{ID: "V1", Title: "Create", Passed: true, Summary: "Primary resource created."},
					{ID: "V2", Title: "Update", Passed: false, Summary: "0 mutation RPCs occurred in UPDATE phase.", Hint: "Check Update() diff detection."},
				},
				Passed: false,
			},
			verbosity: "mutations",
			wantMDSubstrs: []string{
				"# HTTP Lifecycle Report: `widget-sample`",
				"- **Validation Status:** FAIL (1/2 checks passed)",
			},
			wantTermSubstrs: []string{
				"## 4. Detailed RPC Payloads",
				"<hint_for_agent>[V2 Update] Check Update() diff detection.</hint_for_agent>",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			pr := New()

			// Act
			md := pr.RenderMarkdown(tc.report)
			term := pr.RenderTerminal(tc.report, tc.verbosity)

			// Assert
			for _, sub := range tc.wantMDSubstrs {
				if !strings.Contains(md, sub) {
					t.Errorf("RenderMarkdown() missing %q in:\n%s", sub, md)
				}
			}
			for _, sub := range tc.wantTermSubstrs {
				if !strings.Contains(term, sub) {
					t.Errorf("RenderTerminal(%q) missing %q in:\n%s", tc.verbosity, sub, term)
				}
			}
			for _, sub := range tc.forbidTermSubstrs {
				if strings.Contains(term, sub) {
					t.Errorf("RenderTerminal(%q) unexpectedly contained %q", tc.verbosity, sub)
				}
			}
		})
	}
}
