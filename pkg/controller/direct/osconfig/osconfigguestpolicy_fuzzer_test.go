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

package osconfig

import (
	"math/rand"
	"testing"
	"time"
)

func TestFuzzerRoundtrip(t *testing.T) {
	fuzzer := osConfigGuestPolicyFuzzer()
	seed := time.Now().UnixNano()
	t.Logf("Fuzzer seed: %d", seed)
	randStream := rand.New(rand.NewSource(seed))

	for i := 0; i < 100; i++ {
		fuzzer.FuzzSpec(t, randStream.Int63())
		fuzzer.FuzzStatus(t, randStream.Int63())
	}
}

func TestParseBucketName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "projects/my-project/buckets/my-bucket",
			expected: "my-bucket",
		},
		{
			input:    "projects/my-project/buckets/my-bucket/",
			expected: "my-bucket",
		},
		{
			input:    "my-bucket-short",
			expected: "my-bucket-short",
		},
		{
			input:    "my-bucket-short/",
			expected: "my-bucket-short",
		},
		{
			input:    "some-other/unparseable/format/",
			expected: "format",
		},
		{
			input:    "",
			expected: "",
		},
	}

	for _, tc := range tests {
		actual := parseBucketName(tc.input)
		if actual != tc.expected {
			t.Errorf("parseBucketName(%q) = %q; expected %q", tc.input, actual, tc.expected)
		}
	}
}
