// Copyright 2023 Google LLC
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

package slice_test

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	"github.com/google/go-cmp/cmp"
)

func TestReverse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "nil array",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty array",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "array with content",
			input:    []string{"one", "two", "three", "four", "five"},
			expected: []string{"five", "four", "three", "two", "one"},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			sliceToReverse := tc.input
			slice.Reverse(sliceToReverse)
			if got, want := sliceToReverse, tc.expected; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff (-want +got): \n%v", cmp.Diff(want, got))
			}
		})
	}
}
