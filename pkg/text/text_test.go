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

func TestSnakeCaseToLowerCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "snake_case_input",
			expected: "snakecaseinput",
		},
		{
			input:    "input",
			expected: "input",
		},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			output := SnakeCaseToLowerCase(tc.input)
			if tc.expected != output {
				t.Errorf("error parsing snake case to lower case: got %v, want %v", output, tc.expected)
			}
		})
	}
}

func TestIsPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "PrivateCA",
			expected: true,
		},
		{
			input:    "CAPool",
			expected: true,
		},
		{
			input:    "IAM",
			expected: true,
		},
		{
			input:    "Bigtable",
			expected: true,
		},
		{
			input:    "bigquery",
			expected: false,
		},
		{
			input:    "MD5",
			expected: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			output := IsPascalCase(tc.input)
			if tc.expected != output {
				t.Errorf("error checking if the input is pascal case: got %v, want %v", output, tc.expected)
			}
		})
	}
}

func TestRemoveSpecialCharacters(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Config Connector 123",
			expected: "Config Connector 123",
		},
		{
			input:    "project_id",
			expected: "projectid",
		},
		{
			input:    "${uniqueId}",
			expected: "uniqueId",
		},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			output := RemoveSpecialCharacters(tc.input)
			if tc.expected != output {
				t.Errorf("error checking if the input is pascal case: got %v, want %v", output, tc.expected)
			}
		})
	}
}
