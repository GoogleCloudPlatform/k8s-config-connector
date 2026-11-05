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

package logparser

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name            string
		rawLog          string
		wantPhases      []string
		wantStatusCodes []int
		wantReqBody0    string
	}{
		{
			name: "explicit Phase comments with sticky inheritance",
			rawLog: `# Phase: CREATE
GET https://example.googleapis.com/v1/projects/${projectId}/locations/us-central1/widgets/w1?%24alt=json
User-Agent: kcc/${kccVersion}

404 Not Found

{
  "error": {
    "code": 404
  }
}

---

POST https://example.googleapis.com/v1/projects/${projectId}/locations/us-central1/widgets?widgetId=w1
Content-Type: application/json

{
  "description": "initial"
}

200 OK

{
  "description": "initial"
}

---

# Phase: UPDATE
PATCH https://example.googleapis.com/v1/projects/${projectId}/locations/us-central1/widgets/w1?updateMask=description
Content-Type: application/json

{
  "description": "updated"
}

200 OK

{
  "description": "updated"
}
`,
			wantPhases:      []string{PhaseCreate, PhaseCreate, PhaseUpdate},
			wantStatusCodes: []int{404, 200, 200},
		},
		{
			name: "status-like string 200 OK inside request JSON body",
			rawLog: `# Phase: CREATE
POST https://example.googleapis.com/v1/projects/${projectId}/locations/us-central1/widgets?widgetId=w1
Content-Type: application/json

{
  "description": "200 OK\ninside a string should not split request body"
}

200 OK

{
  "name": "projects/${projectId}/locations/us-central1/widgets/w1"
}
`,
			wantPhases:      []string{PhaseCreate},
			wantStatusCodes: []int{200},
			wantReqBody0:    "{\n  \"description\": \"200 OK\\ninside a string should not split request body\"\n}",
		},
		{
			name: "gRPC status header OK and NotFound",
			rawLog: `# Phase: CREATE
GRPC /google.example.v1.WidgetService/GetWidget
Request-Type: google.example.v1.GetWidgetRequest

{
  "name": "projects/p/locations/l/widgets/w1"
}

error: code = NotFound desc = widget not found

---

GRPC /google.example.v1.WidgetService/CreateWidget
Request-Type: google.example.v1.CreateWidgetRequest

{
  "widget": {
    "name": "w1"
  }
}

OK

{
  "name": "projects/p/locations/l/widgets/w1"
}
`,
			wantPhases:      []string{PhaseCreate, PhaseCreate},
			wantStatusCodes: []int{404, 200},
		},
		{
			name: "unmarked log entries default to empty phase",
			rawLog: `GET https://example.googleapis.com/v1/projects/p/locations/l/widgets/w1

200 OK

{
  "name": "w1"
}
`,
			wantPhases:      []string{""},
			wantStatusCodes: []int{200},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange & Act
			entries, err := Parse(tc.rawLog)

			// Assert
			if err != nil {
				t.Fatalf("Parse() returned unexpected error: %v", err)
			}
			if len(entries) != len(tc.wantPhases) {
				t.Fatalf("len(entries) = %d, want %d", len(entries), len(tc.wantPhases))
			}
			var gotPhases []string
			var gotCodes []int
			for _, e := range entries {
				gotPhases = append(gotPhases, e.Phase)
				gotCodes = append(gotCodes, e.Response.StatusCode)
			}
			if !reflect.DeepEqual(gotPhases, tc.wantPhases) {
				t.Errorf("phases = %v, want %v", gotPhases, tc.wantPhases)
			}
			if !reflect.DeepEqual(gotCodes, tc.wantStatusCodes) {
				t.Errorf("statusCodes = %v, want %v", gotCodes, tc.wantStatusCodes)
			}
			if tc.wantReqBody0 != "" && entries[0].Request.Body != tc.wantReqBody0 {
				t.Errorf("entries[0].Request.Body = %q, want %q", entries[0].Request.Body, tc.wantReqBody0)
			}
		})
	}
}

func TestStripAndFormatPhaseComments(t *testing.T) {
	tests := []struct {
		name           string
		entries        []*Entry
		wantSubstrings []string
	}{
		{
			name: "formats phase transitions and strips phase comments cleanly",
			entries: []*Entry{
				{
					Phase: PhaseCreate,
					Request: HTTPRequest{
						Method: "POST",
						URL:    "https://example.googleapis.com/v1/projects/${projectId}/widgets",
						Body:   "{\n  \"name\": \"w1\"\n}",
					},
					Response: HTTPResponse{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       "{}",
					},
				},
				{
					Phase: PhaseDelete,
					Request: HTTPRequest{
						Method: "DELETE",
						URL:    "https://example.googleapis.com/v1/projects/${projectId}/widgets/w1",
					},
					Response: HTTPResponse{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       "{}",
					},
				},
			},
			wantSubstrings: []string{"# Phase: CREATE", "# Phase: DELETE"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			formatted := FormatWithPhaseComments(tc.entries)
			stripped := StripPhaseComments(formatted)

			// Assert
			for _, sub := range tc.wantSubstrings {
				if !strings.Contains(formatted, sub) {
					t.Errorf("FormatWithPhaseComments() missing %q in:\n%s", sub, formatted)
				}
			}
			if strings.Contains(stripped, "# Phase:") {
				t.Errorf("StripPhaseComments() left phase comments in output:\n%s", stripped)
			}
		})
	}
}
