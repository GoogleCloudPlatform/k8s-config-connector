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

package codegen

import "testing"

func TestPlural(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want string
	}{
		{"jobscompany", "jobscompanies"},
		{"discoveryenginesamplequery", "discoveryenginesamplequeries"},
		{"gateway", "gateways"},
		{"proxy", "proxies"},
		{"vertexairagcorpus", "vertexairagcorpora"},
		{"corpus", "corpora"},
		{"setting", "settings"},
		{"settings", "settings"},
		{"iapsettings", "iapsettings"},
		{"metric", "metrics"},
		{"metrics", "metrics"},
		{"status", "statuses"},
		{"index", "indexes"},
		{"batch", "batches"},
		{"dish", "dishes"},
		{"shelf", "shelves"},
		{"knife", "knives"},
	} {
		if got := Plural(tc.in); got != tc.want {
			t.Errorf("Plural(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
