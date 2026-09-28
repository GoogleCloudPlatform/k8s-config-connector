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

func TestSingular(t *testing.T) {
	// One case per branch of the switch, so deleting a branch names itself
	// rather than surfacing as a missed match somewhere downstream.
	for _, tc := range []struct {
		branch string
		in     string
		want   string
	}{
		{"ies, longer than four", "policies", "policy"},
		{"ies, too short for that branch, falls to trailing s", "ies", "ie"},
		{"sses", "addresses", "address"},
		{"ches", "batches", "batch"},
		{"shes", "dishes", "dish"},
		{"xes", "indexes", "index"},
		{"ss is not a plural", "access", "access"},
		{"us is not a plural", "status", "status"},
		{"is is not a plural", "analysis", "analysis"},
		{"trailing s", "subnetworks", "subnetwork"},
		{"trailing s, too short", "as", "as"},
		{"no plural to strip", "policy", "policy"},

		// Two known-wrong answers, pinned so a change to them is deliberate.
		// Neither costs anything today, because no Kind in the tree ends in
		// Http, Cache or Size, and a rule that only decides whether to ask a
		// person a question need not be right about English.
		//
		// The ss/us/is guards do not catch an acronym, so "https" loses its s.
		// The ches/shes/xes/zes branch assumes the stem ends in a consonant,
		// which holds for batches and boxes but not for a stem already ending
		// in e: caches and sizes lose that e as well.
		{"acronym, mangled", "https", "http"},
		{"stem ends in e, mangled", "caches", "cach"},
		{"stem ends in e, mangled", "sizes", "siz"},
	} {
		if got := Singular(tc.in); got != tc.want {
			t.Errorf("%s: Singular(%q) = %q, want %q", tc.branch, tc.in, got, tc.want)
		}
	}
}
