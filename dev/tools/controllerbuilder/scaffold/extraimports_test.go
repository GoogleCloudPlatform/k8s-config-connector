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

package scaffold

import (
	"strings"
	"testing"
)

// TestExtraImportsFor verifies that ExtraImportsFor correctly detects external
// type qualifiers across provided struct bodies and outputs corresponding aliased import statements.
func TestExtraImportsFor(t *testing.T) {
	tests := []struct {
		name       string
		spec       string
		status     string
		wantSubstr string
		wantNone   bool
	}{
		{name: "apiextensionsv1 in spec", spec: "\tConfig apiextensionsv1.JSON `json:\"config\"`",
			wantSubstr: `apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"`},
		{name: "common in status", status: "\tError *common.Status `json:\"error\"`",
			wantSubstr: `common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"`},
		{name: "nothing special", spec: "\tName *string `json:\"name\"`", wantNone: true},
	}
	for _, tt := range tests {
		got := ExtraImportsFor(tt.spec, tt.status)
		if tt.wantNone {
			if len(got) != 0 {
				t.Errorf("%s: want no imports, got %v", tt.name, got)
			}
			continue
		}
		if len(got) != 1 || !strings.Contains(got[0], tt.wantSubstr) {
			t.Errorf("%s: got %v, want one line containing %q", tt.name, got, tt.wantSubstr)
		}
		// An import line without an alias leaves the qualifier unresolved.
		if !strings.Contains(got[0], " \"") {
			t.Errorf("%s: import line %q has no alias", tt.name, got[0])
		}
	}
}
