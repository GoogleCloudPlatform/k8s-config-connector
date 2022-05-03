// Copyright 2022 Google LLC
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

package text

import "testing"

func TestAsSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "lowerCamelCaseInput",
			expected: "lower_camel_case_input",
		},
		{
			input:    "UpperCamelCaseInput",
			expected: "upper_camel_case_input",
		},
		{
			input:    "snake_case_input",
			expected: "snake_case_input",
		},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			output := AsSnakeCase(tc.input)
			if tc.expected != output {
				t.Errorf("expected: %v, actual: %v", tc.expected, output)
			}
		})
	}
}
