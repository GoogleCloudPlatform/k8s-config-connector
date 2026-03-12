// Copyright 2025 Google LLC
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

package httptogrpc

import (
	"reflect"
	"strings"
	"testing"
)

func TestPathMatcher(t *testing.T) {
	tests := []struct {
		httpPath      string
		requestPath   string
		expectMatch   bool
		expectMapping map[string]string
	}{
		// Matches with single wildcards
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/*}",
			requestPath: "v1/projects/proj1/locations/us-central1/instances/inst1",
			expectMatch: true,
			expectMapping: map[string]string{
				"name": "projects/proj1/locations/us-central1/instances/inst1",
			},
		},
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/*}",
			requestPath: "v1/projects/proj1/locations/us-central1/instances/inst1/extra",
			expectMatch: false,
		},
		// Matches with trailing literals
		{
			httpPath:    "v1/{parent=projects/*/locations/*}/instances",
			requestPath: "v1/projects/proj1/locations/us-central1/instances",
			expectMatch: true,
			expectMapping: map[string]string{
				"parent": "projects/proj1/locations/us-central1",
			},
		},
		{
			httpPath:    "v1/{parent=projects/*/locations/*}/instances/bar",
			requestPath: "v1/projects/proj1/locations/us-central1/instances",
			expectMatch: false,
		},
		// Multi-segment wildcards (**)
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/*}",
			requestPath: "v1/projects/proj1/locations/us-central1",
			expectMatch: false,
		},
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/**}",
			requestPath: "v1/projects/proj1/locations/us-central1/instances/foo/bar/baz",
			expectMatch: true,
			expectMapping: map[string]string{
				"name": "projects/proj1/locations/us-central1/instances/foo/bar/baz",
			},
		},
		// Multi-segment wildcards with trailing literals
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/**}/tail-literal",
			requestPath: "v1/projects/proj1/locations/us-central1/instances/foo/bar/baz/tail-literal",
			expectMatch: true,
			expectMapping: map[string]string{
				"name": "projects/proj1/locations/us-central1/instances/foo/bar/baz",
			},
		},
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/**}/tail-literal",
			requestPath: "v1/projects/proj1/locations/us-central1/instances/foo/bar/baz",
			expectMatch: false,
		},
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/**}/tail-literal",
			requestPath: "v1/projects/proj1/locations/us-central1/instances/foo/bar/wrong-tail",
			expectMatch: false,
		},
		// Multi-segment wildcards with additional variables
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/**}/{foo}",
			requestPath: "v1/projects/proj1/locations/us-central1/instances/foo/bar/baz/ending",
			expectMatch: true,
			expectMapping: map[string]string{
				"name": "projects/proj1/locations/us-central1/instances/foo/bar/baz",
				"foo":  "ending",
			},
		},
		// Wildcards matching empty segments
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/**}",
			requestPath: "v1/projects/proj1/locations/us-central1/instances/",
			expectMatch: true,
			expectMapping: map[string]string{
				"name": "projects/proj1/locations/us-central1/instances/",
			},
		},
		{
			httpPath:    "v1/{name=projects/*/locations/*/instances/**}/foo",
			requestPath: "v1/projects/proj1/locations/us-central1/instances/foo",
			expectMatch: true,
			expectMapping: map[string]string{
				"name": "projects/proj1/locations/us-central1/instances",
			},
		},
	}

	for _, test := range tests {
		matcher, err := newPathMatcher(test.httpPath)
		if err != nil {
			t.Fatalf("newPathMatcher(%q) failed: %v", test.httpPath, err)
		}
		tokens := strings.Split(test.requestPath, "/")
		mapping, matched := matcher.Match(tokens)
		if matched != test.expectMatch {
			t.Errorf("matcher.match(%q) vs %q = %v, want %v", test.requestPath, test.httpPath, matched, test.expectMatch)
		}
		if matched && !reflect.DeepEqual(mapping, test.expectMapping) {
			t.Errorf("matcher.match(%q) vs %q = %v, want %v", test.requestPath, test.httpPath, mapping, test.expectMapping)
		}
	}
}
