// Copyright 2024 Google LLC
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

package testutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// MustReadFile reads the specified file, failing the test on error.
func MustReadFile(t *testing.T, p string) []byte {
	t.Helper()
	b, err := os.ReadFile(p)
	if err != nil {
		absPath, _ := filepath.Abs(p)
		t.Fatalf("error reading file '%v' (absolute path %v): %v", p, absPath, err)
	}
	return b
}

// CompareGoldenFile ensures that the specified golden file matches the output we got.
func CompareGoldenFile(t *testing.T, p string, got string, normalizers ...func(s string) string) {
	if os.Getenv("WRITE_GOLDEN_OUTPUT") != "" {
		// Short-circuit when the output is correct
		b, err := os.ReadFile(p)
		if err == nil {
			want := string(b)
			for _, normalizer := range normalizers {
				got = normalizer(got)
				want = normalizer(want)
			}
			if want == got {
				return
			}
		}

		if err := os.WriteFile(p, []byte(got), 0644); err != nil {
			t.Fatalf("failed to write golden output %s: %v", p, err)
		}
		t.Errorf("wrote output to %s", p)
	} else {
		want := string(MustReadFile(t, p))

		for _, normalizer := range normalizers {
			got = normalizer(got)
			want = normalizer(want)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected diff in %s: %s", p, diff)
		}
	}
}
