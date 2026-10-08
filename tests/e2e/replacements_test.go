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

import "testing"

func TestReplaceTagResourceIDs(t *testing.T) {
	responseBody := `{"tagBindings":[{"name":"tagBindings/parent/tagValues/987654321","tagValue":"tagValues/987654321"}],"parent":"tagKeys/123456789"}`

	got := ReplaceTagResourceIDs(responseBody)
	want := `{"tagBindings":[{"name":"tagBindings/parent/tagValues/${tagValueID}","tagValue":"tagValues/${tagValueID}"}],"parent":"tagKeys/${tagKeyID}"}`
	if got != want {
		t.Fatalf("ReplaceTagResourceIDs() = %q, want %q", got, want)
	}
}

func TestExtractTagResourceIDsFromLinks(t *testing.T) {
	replacements := NewReplacements()
	replacements.ExtractIDsFromLinks("tagKeys/123456789")
	replacements.ExtractIDsFromLinks("tagValues/987654321")

	got := replacements.ApplyReplacements("tagKeys/123456789 tagValues/987654321")
	want := "tagKeys/${tagKeyID} tagValues/${tagValueID}"
	if got != want {
		t.Fatalf("ApplyReplacements() = %q, want %q", got, want)
	}
}
