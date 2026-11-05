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

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReachable(t *testing.T) {
	// a leads to b, b leads to c and back to a, and d leads nowhere.
	graph := map[string][]string{"a": {"b"}, "b": {"c", "a"}}
	next := func(n string) []string { return graph[n] }
	stopAtC := func(n string) bool { return n == "c" }
	for _, tc := range []struct {
		name  string
		seeds []string
		stop  func(string) bool
		want  map[string]bool
	}{
		{name: "follows every edge, through a cycle", seeds: []string{"a"}, want: map[string]bool{"a": true, "b": true, "c": true}},
		{name: "doesn't enter a stop node", seeds: []string{"a"}, stop: stopAtC, want: map[string]bool{"a": true, "b": true}},
		{name: "doesn't enter a seed that is a stop node", seeds: []string{"c"}, stop: stopAtC, want: map[string]bool{}},
		{name: "starts from every seed", seeds: []string{"c", "d"}, want: map[string]bool{"c": true, "d": true}},
		{name: "no seeds", want: map[string]bool{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got := Reachable(tc.seeds, next, tc.stop)

			// Assert
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Reachable() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
