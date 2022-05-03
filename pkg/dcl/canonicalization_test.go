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

package dcl_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
)

func TestCanonicalizeIntegerValue(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected int64
		hasError bool
	}{
		{
			name:     "int64 to int64",
			value:    int64(10),
			expected: int64(10),
		},
		{
			name:     "int to int64",
			value:    int(20),
			expected: int64(20),
		},
		{
			name:     "float64 to int64",
			value:    float64(30.5),
			expected: int64(30),
		},
		{
			name:     "non integer value",
			value:    "wrong value",
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			canonicalizedVal, err := dcl.CanonicalizeIntegerValue(tc.value)
			if err != nil {
				if !tc.hasError {
					t.Fatalf("error canonicalizing integer value: got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := canonicalizedVal, tc.expected; got != want {
				t.Fatalf("error canonicalizing integer value: got %v, want %v", got, want)
			}
		})
	}
}

func TestCanonicalizeNumberValue(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected float64
		hasError bool
	}{
		{
			name:     "int64 to float64",
			value:    int64(10),
			expected: float64(10),
		},
		{
			name:     "int to float64",
			value:    int(20),
			expected: float64(20),
		},
		{
			name:     "float64 to float64",
			value:    float64(30.5),
			expected: float64(30.5),
		},
		{
			name:     "float32 to float64",
			value:    float32(40.8),
			expected: float64(float32(40.8)),
		},
		{
			name:     "non number value",
			value:    "wrong value",
			hasError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			canonicalizedVal, err := dcl.CanonicalizeNumberValue(tc.value)
			if err != nil {
				if !tc.hasError {
					t.Fatalf("error canonicalizing integer value: got an error, but want no error: %v", err)
				}
				return
			}
			if got, want := canonicalizedVal, tc.expected; got != want {
				t.Fatalf("error canonicalizing integer value: got %v, want %v", got, want)
			}
		})
	}
}
