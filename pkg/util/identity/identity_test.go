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

package identity

import (
	"regexp"
	"testing"
)

func TestProjectIDRegexp(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		parses bool
	}{
		{
			name:   "Normal parse",
			input:  "project-id-with-1-number",
			parses: true,
		},
		{
			name:   "Normal parse short",
			input:  "projec",
			parses: true,
		},
		{
			name:   "Normal parse long",
			input:  "project-with-30-characters-abc",
			parses: true,
		},
		{
			name:   "Normal parse, end number",
			input:  "project-id-last-1",
			parses: true,
		},
		{
			name:   "Too short",
			input:  "proje",
			parses: false,
		},
		{
			name:   "Too long",
			input:  "project-with-31-characters-abcd",
			parses: false,
		},
		{
			name:   "Capital letter",
			input:  "project-ID-with-1-number",
			parses: false,
		},
		{
			name:   "Capital letter first",
			input:  "Project-id-with-1-number",
			parses: false,
		},
		{
			name:   "Capital letter last",
			input:  "project-id-with-1-numbeR",
			parses: false,
		},
		{
			name:   "Number first",
			input:  "1-number",
			parses: false,
		},
		{
			name:   "Dash first",
			input:  "-number",
			parses: false,
		},
		{
			name:   "Dash last",
			input:  "number-",
			parses: false,
		},
	}

	for _, tc := range tests {
		parser, err := regexp.Compile("^" + ProjectIDRegexp + "$")
		if err != nil {
			t.Fatalf("Test %s expected error but did not get one %v", tc.name, err)
		}
		result := parser.MatchString(tc.input)
		if result != tc.parses {
			t.Fatalf("Test %q parse of %q gives bad result %v != %v", tc.name, tc.input, result, tc.parses)
		}
	}
}
