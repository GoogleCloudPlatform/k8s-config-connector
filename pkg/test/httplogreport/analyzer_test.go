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

package httplogreport

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAnalyzer_AnalyzeFixture(t *testing.T) {
	createYAML := `apiVersion: example.cnrm.cloud.google.com/v1alpha1
kind: ExampleWidget
metadata:
  name: widget-test
spec:
  description: "initial"
`
	updateYAML := `apiVersion: example.cnrm.cloud.google.com/v1alpha1
kind: ExampleWidget
metadata:
  name: widget-test
spec:
  description: "updated"
`

	tests := []struct {
		name             string
		httpLog          string
		wantPassed       bool
		wantReportSubstr string
	}{
		{
			name: "passing fixture end-to-end report generation and WriteReport",
			httpLog: `# Phase: CREATE
POST https://example.googleapis.com/v1/projects/p/locations/l/widgets?widgetId=widget-test HTTP/1.1
Content-Type: application/json

{
  "description": "initial"
}

HTTP/1.1 200 OK
Content-Type: application/json

{
  "description": "initial",
  "name": "projects/p/locations/l/widgets/widget-test"
}

---

# Phase: UPDATE
PATCH https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-test?updateMask=description HTTP/1.1
Content-Type: application/json

{
  "description": "updated"
}

HTTP/1.1 200 OK
Content-Type: application/json

{
  "description": "updated",
  "name": "projects/p/locations/l/widgets/widget-test"
}

---

# Phase: DELETE
DELETE https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-test HTTP/1.1

HTTP/1.1 200 OK
`,
			wantPassed:       true,
			wantReportSubstr: "PASS (7/7 checks passed)",
		},
		{
			name: "failing fixture with missing delete",
			httpLog: `# Phase: CREATE
POST https://example.googleapis.com/v1/projects/p/locations/l/widgets?widgetId=widget-test HTTP/1.1
Content-Type: application/json

{
  "description": "initial"
}

HTTP/1.1 200 OK
Content-Type: application/json

{
  "description": "initial",
  "name": "projects/p/locations/l/widgets/widget-test"
}

---

# Phase: UPDATE
PATCH https://example.googleapis.com/v1/projects/p/locations/l/widgets/widget-test?updateMask=description HTTP/1.1
Content-Type: application/json

{
  "description": "updated"
}

HTTP/1.1 200 OK
Content-Type: application/json

{
  "description": "updated",
  "name": "projects/p/locations/l/widgets/widget-test"
}
`,
			wantPassed:       false,
			wantReportSubstr: "FAIL (6/7 checks passed)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			if err := os.WriteFile(filepath.Join(tmpDir, "create.yaml"), []byte(createYAML), 0644); err != nil {
				t.Fatalf("writing create.yaml: %v", err)
			}
			if err := os.WriteFile(filepath.Join(tmpDir, "update.yaml"), []byte(updateYAML), 0644); err != nil {
				t.Fatalf("writing update.yaml: %v", err)
			}
			if err := os.WriteFile(filepath.Join(tmpDir, "_http.log"), []byte(tc.httpLog), 0644); err != nil {
				t.Fatalf("writing _http.log: %v", err)
			}
			analyzer := NewAnalyzer()

			// Act
			dirs, err := DiscoverFixtureDirs(tmpDir)
			if err != nil {
				t.Fatalf("DiscoverFixtureDirs failed: %v", err)
			}
			report, err := analyzer.AnalyzeFixture(tmpDir)
			if err != nil {
				t.Fatalf("AnalyzeFixture failed: %v", err)
			}
			reportPath, err := analyzer.WriteReport(tmpDir, report)
			if err != nil {
				t.Fatalf("WriteReport failed: %v", err)
			}

			// Assert
			if len(dirs) != 1 || dirs[0] != tmpDir {
				t.Errorf("expected DiscoverFixtureDirs to return [%s], got %v", tmpDir, dirs)
			}
			if report.Passed != tc.wantPassed {
				t.Errorf("report.Passed = %v, want %v: %+v", report.Passed, tc.wantPassed, report.Checks)
			}
			if report.PhaseSource != PhaseSourceExplicit {
				t.Errorf("expected PhaseSourceExplicit, got %s", report.PhaseSource)
			}
			writtenBytes, err := os.ReadFile(reportPath)
			if err != nil {
				t.Fatalf("reading written report: %v", err)
			}
			if !strings.Contains(string(writtenBytes), tc.wantReportSubstr) {
				t.Errorf("expected written report to contain %q, got:\n%s", tc.wantReportSubstr, string(writtenBytes))
			}
		})
	}
}
