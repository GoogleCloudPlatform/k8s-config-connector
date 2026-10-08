// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestComputeRatchet(t *testing.T) {
	header := "# a documented baseline\n# second header line\n"

	tests := []struct {
		name        string
		baseline    string
		got         string
		wantAdded   []string
		wantRemoved []string
	}{
		{
			name:     "unchanged",
			baseline: header + "violation a\nviolation b\n",
			got:      "violation a\nviolation b",
		},
		{
			name:      "new violation is added",
			baseline:  header + "violation a\n",
			got:       "violation a\nviolation b",
			wantAdded: []string{"violation b"},
		},
		{
			name:        "fixed violation is removed",
			baseline:    header + "violation a\nviolation b\n",
			got:         "violation a",
			wantRemoved: []string{"violation b"},
		},
		{
			name:        "added and removed together",
			baseline:    header + "violation a\n",
			got:         "violation b",
			wantAdded:   []string{"violation b"},
			wantRemoved: []string{"violation a"},
		},
		{
			// Regression: comments were once counted as entries, so a documented
			// baseline saw its own header as fixed violations and pruned it away.
			name:     "comments are not entries",
			baseline: header + "violation a\n",
			got:      "violation a",
		},
		{
			name:     "blank lines and whitespace are ignored",
			baseline: header + "\n  violation a  \n\n",
			got:      "violation a\n\n",
		},
		{
			name:      "empty baseline reports everything as added",
			baseline:  header,
			got:       "violation a",
			wantAdded: []string{"violation a"},
		},
		{
			name:        "empty got reports everything as removed",
			baseline:    header + "violation a\n",
			got:         "",
			wantRemoved: []string{"violation a"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ComputeRatchet(tc.baseline, tc.got)
			if diff := cmp.Diff(tc.wantAdded, got.Added, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("Added diff (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tc.wantRemoved, got.Removed, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("Removed diff (-want +got):\n%s", diff)
			}
		})
	}
}

// TestComputeRatchetPreservesHeader is the regression test for the bug where
// pruning rewrote the file from the entry set alone, deleting the baseline's own
// documentation.
func TestComputeRatchetPreservesHeader(t *testing.T) {
	baseline := "# what this file is for\n# how to add an entry\nviolation a\nviolation b\n"

	got := ComputeRatchet(baseline, "violation a")

	if len(got.Removed) != 1 {
		t.Fatalf("expected 1 removed entry, got %v", got.Removed)
	}
	for _, want := range []string{"# what this file is for", "# how to add an entry"} {
		if !strings.Contains(got.Pruned, want) {
			t.Errorf("pruned content lost header line %q:\n%s", want, got.Pruned)
		}
	}
	if !strings.Contains(got.Pruned, "violation a") {
		t.Errorf("pruned content lost the surviving entry:\n%s", got.Pruned)
	}
	if strings.Contains(got.Pruned, "violation b") {
		t.Errorf("pruned content still has the fixed entry:\n%s", got.Pruned)
	}
}

// TestComputeRatchetPrunedEndsWithNewline is a regression test. Pruning once
// wrote the file without a trailing newline, so a later `cat >> file` appended
// the new entry onto the last existing one and silently corrupted both. These
// baselines are appended to by hand and by agents, so the newline is load-bearing.
func TestComputeRatchetPrunedEndsWithNewline(t *testing.T) {
	got := ComputeRatchet("# header\nviolation a\nviolation b\n", "violation a")
	if !strings.HasSuffix(got.Pruned, "\n") {
		t.Errorf("pruned content must end with a newline, got %q", got.Pruned)
	}
}

// TestShouldPrune covers the property the whole mechanism exists for: when there
// are new violations the baseline is never rewritten, even with
// WRITE_GOLDEN_OUTPUT set. Absorbing new entries during the regenerate-goldens
// workflow is exactly what turns the check into a no-op.
func TestShouldPrune(t *testing.T) {
	tests := []struct {
		name         string
		decision     RatchetDecision
		writeEnabled bool
		want         bool
	}{
		{
			name:         "new violations are never written, even with write enabled",
			decision:     RatchetDecision{Added: []string{"new"}, Removed: []string{"fixed"}},
			writeEnabled: true,
			want:         false,
		},
		{
			name:         "new violations only, write enabled",
			decision:     RatchetDecision{Added: []string{"new"}},
			writeEnabled: true,
			want:         false,
		},
		{
			name:         "fixed violations are pruned when write enabled",
			decision:     RatchetDecision{Removed: []string{"fixed"}},
			writeEnabled: true,
			want:         true,
		},
		{
			name:         "fixed violations are not pruned without write enabled",
			decision:     RatchetDecision{Removed: []string{"fixed"}},
			writeEnabled: false,
			want:         false,
		},
		{
			name:         "nothing to do",
			decision:     RatchetDecision{},
			writeEnabled: true,
			want:         false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ShouldPrune(tc.decision, tc.writeEnabled); got != tc.want {
				t.Errorf("ShouldPrune() = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestCompareRatchetFilePrunes covers the other direction: a fixed violation is
// removed when WRITE_GOLDEN_OUTPUT is set, and the header survives.
func TestCompareRatchetFilePrunes(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "baseline.txt")
	if err := os.WriteFile(path, []byte("# header\nviolation a\nviolation b\n"), 0644); err != nil {
		t.Fatalf("writing baseline: %v", err)
	}

	t.Setenv("WRITE_GOLDEN_OUTPUT", "1")
	CompareRatchetFile(t, path, "violation a")

	after, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading baseline after: %v", err)
	}
	got := string(after)
	if !strings.Contains(got, "# header") {
		t.Errorf("prune dropped the header:\n%s", got)
	}
	if strings.Contains(got, "violation b") {
		t.Errorf("prune kept the fixed entry:\n%s", got)
	}
	if !strings.Contains(got, "violation a") {
		t.Errorf("prune dropped the surviving entry:\n%s", got)
	}
}

// TestCompareRatchetFileNoBaselineFile covers the bootstrap case: a missing file
// is fine only when there is nothing to record.
func TestCompareRatchetFileNoBaselineFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "does-not-exist.txt")
	CompareRatchetFile(t, path, "")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("expected no baseline file to be created, got err=%v", err)
	}
}
