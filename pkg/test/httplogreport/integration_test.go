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

package httplogreport

import (
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func findCheck(report *Report, id string) *ValidationCheck {
	for i := range report.Checks {
		if report.Checks[i].ID == id {
			return &report.Checks[i]
		}
	}
	return nil
}

func TestFixturesDataModel(t *testing.T) {
	type specDiffExpectation struct {
		sentInUpdateRPC bool
		matchesEndState bool
	}

	cases := []struct {
		dirName            string
		wantPassed         bool
		wantFailedCheckIDs []string
		wantReasonSubstr   string
		wantEntryPhases    []string
		wantSpecDiffCount  int
		wantSpecDiffByPath map[string]specDiffExpectation
	}{
		{
			dirName:           "good-composite-create",
			wantPassed:        true,
			wantEntryPhases:   []string{PhaseCreate, PhaseCreate, PhaseCreate, PhaseUpdate, PhaseDelete},
			wantSpecDiffCount: 1,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.description": {sentInUpdateRPC: true, matchesEndState: true},
			},
		},
		{
			dirName:           "good-field-by-field-update",
			wantPassed:        true,
			wantEntryPhases:   []string{PhaseCreate, PhaseUpdate, PhaseUpdate, PhaseDelete},
			wantSpecDiffCount: 2,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.description":    {sentInUpdateRPC: true, matchesEndState: true},
				"metadata.labels.env": {sentInUpdateRPC: true, matchesEndState: true},
			},
		},
		{
			dirName:           "good-maskless-put-update",
			wantPassed:        true,
			wantEntryPhases:   []string{PhaseCreate, PhaseUpdate, PhaseDelete},
			wantSpecDiffCount: 2,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.displayName": {sentInUpdateRPC: true, matchesEndState: true},
				"spec.priority":    {sentInUpdateRPC: true, matchesEndState: true},
			},
		},
		{
			dirName:            "bad-v1-create-400",
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V1"},
			wantReasonSubstr:   "400 Bad Request",
			wantEntryPhases:    []string{PhaseCreate, PhaseCreate, PhaseDelete},
			wantSpecDiffCount:  0,
		},
		{
			dirName:            "bad-v2-update-noop",
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V2", "V3"},
			wantReasonSubstr:   "SILENT UPDATE NO-OP DETECTED",
			wantEntryPhases:    []string{PhaseCreate, PhaseUpdate, PhaseDelete},
			wantSpecDiffCount:  1,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.description": {sentInUpdateRPC: false, matchesEndState: false},
			},
		},
		{
			dirName:            "bad-v2-dropped-field",
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V2", "V3"},
			wantReasonSubstr:   "DROPPED UPDATE FIELD(S)",
			wantEntryPhases:    []string{PhaseCreate, PhaseUpdate, PhaseDelete},
			wantSpecDiffCount:  2,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.description": {sentInUpdateRPC: true, matchesEndState: true},
				"spec.displayName": {sentInUpdateRPC: false, matchesEndState: false},
			},
		},
		{
			dirName:            "bad-v3-endstate-mismatch",
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V3"},
			wantReasonSubstr:   "END-STATE MISMATCH",
			wantEntryPhases:    []string{PhaseCreate, PhaseUpdate, PhaseUpdate, PhaseDelete},
			wantSpecDiffCount:  1,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.description": {sentInUpdateRPC: true, matchesEndState: false},
			},
		},
		{
			dirName:            "bad-v4-missing-delete",
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V4"},
			wantReasonSubstr:   "MISSING DELETE RPC",
			wantEntryPhases:    []string{PhaseCreate, PhaseUpdate, PhaseUpdate},
			wantSpecDiffCount:  1,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.description": {sentInUpdateRPC: true, matchesEndState: true},
			},
		},
		{
			dirName:            "bad-v5-duplicate-patch-loop",
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V5"},
			wantReasonSubstr:   "Duplicate consecutive mutation RPCs",
			wantEntryPhases:    []string{PhaseCreate, PhaseUpdate, PhaseUpdate, PhaseDelete},
			wantSpecDiffCount:  1,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.description": {sentInUpdateRPC: true, matchesEndState: true},
			},
		},
		{
			dirName:            "bad-v6-mockgcp-and-auth-leak",
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V6"},
			wantReasonSubstr:   "(mockgcp)",
			wantEntryPhases:    []string{PhaseCreate, PhaseUpdate, PhaseDelete},
			wantSpecDiffCount:  1,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.description": {sentInUpdateRPC: true, matchesEndState: true},
			},
		},
		{
			dirName:            "bad-v7-hand-edited-log",
			wantPassed:         false,
			wantFailedCheckIDs: []string{"V7"},
			wantReasonSubstr:   "Cross-stanza continuity mismatch",
			wantEntryPhases:    []string{PhaseCreate, PhaseCreate, PhaseCreate, PhaseUpdate, PhaseDelete},
			wantSpecDiffCount:  1,
			wantSpecDiffByPath: map[string]specDiffExpectation{
				"spec.description": {sentInUpdateRPC: true, matchesEndState: true},
			},
		},
	}

	allCheckIDs := []string{"V1", "V2", "V3", "V4", "V5", "V6", "V7"}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.dirName, func(t *testing.T) {
			// Arrange
			dir := filepath.Join("testdata", tc.dirName)

			// Act
			report, err := AnalyzeFixture(dir)
			if err != nil {
				t.Fatalf("AnalyzeFixture(%s) failed: %v", dir, err)
			}

			// Assert
			if report.Passed != tc.wantPassed {
				t.Fatalf("report.Passed = %v, want %v", report.Passed, tc.wantPassed)
			}

			for _, id := range allCheckIDs {
				if findCheck(report, id) == nil {
					t.Fatalf("expected check %s to be present in report.Checks", id)
				}
			}

			var gotFailedIDs []string
			var combinedReasons strings.Builder
			for _, chk := range report.Checks {
				if !chk.Passed {
					gotFailedIDs = append(gotFailedIDs, chk.ID)
					combinedReasons.WriteString(chk.Summary + " " + chk.Hint + "\n")
					if strings.TrimSpace(chk.Hint) == "" {
						t.Errorf("failed check %s in %s has empty Hint", chk.ID, tc.dirName)
					}
				}
			}
			if !reflect.DeepEqual(gotFailedIDs, tc.wantFailedCheckIDs) {
				t.Errorf("failed check IDs = %v, want %v", gotFailedIDs, tc.wantFailedCheckIDs)
			}
			if tc.wantReasonSubstr != "" && !strings.Contains(combinedReasons.String(), tc.wantReasonSubstr) {
				t.Errorf("failed check reasons %q did not contain substring %q", combinedReasons.String(), tc.wantReasonSubstr)
			}

			var gotPhases []string
			for _, e := range report.Entries {
				gotPhases = append(gotPhases, e.Phase)
			}
			if !reflect.DeepEqual(gotPhases, tc.wantEntryPhases) {
				t.Errorf("entry phases = %v, want %v", gotPhases, tc.wantEntryPhases)
			}

			if len(report.SpecDiff) != tc.wantSpecDiffCount {
				t.Fatalf("len(report.SpecDiff) = %d, want %d", len(report.SpecDiff), tc.wantSpecDiffCount)
			}
			for _, diff := range report.SpecDiff {
				if wantDiff, ok := tc.wantSpecDiffByPath[diff.Path]; ok {
					if diff.SentInUpdateRPC != wantDiff.sentInUpdateRPC {
						t.Errorf("SpecDiff[%s].SentInUpdateRPC = %v, want %v", diff.Path, diff.SentInUpdateRPC, wantDiff.sentInUpdateRPC)
					}
					if diff.FinalStateMatch != wantDiff.matchesEndState {
						t.Errorf("SpecDiff[%s].FinalStateMatch = %v, want %v", diff.Path, diff.FinalStateMatch, wantDiff.matchesEndState)
					}
				}
			}
		})
	}
}

func TestRenderMarkdownPrinter(t *testing.T) {
	tests := []struct {
		name            string
		dirName         string
		verbosity       string
		wantMDSubstrs   []string
		wantTermSubstrs []string
	}{
		{
			name:      "passing fixture markdown and summary terminal output",
			dirName:   "good-composite-create",
			verbosity: "summary",
			wantMDSubstrs: []string{
				"# HTTP Lifecycle Report: `good-composite-create`",
				"- **Validation Status:** PASS (7/7 checks passed)",
				"- **Integrity Hash:** `sha256:",
				"## 1. Spec Diff (`create.yaml` -> `update.yaml`)",
				"## 2. Primary Resource RPC Summary",
				"## 3. Automated Validation Checks",
				"- [x] **V1 (Create):**",
			},
			wantTermSubstrs: []string{
				"- **Validation Status:** PASS (7/7 checks passed)",
			},
		},
		{
			name:      "failing fixture markdown and mutations terminal output with agent hint",
			dirName:   "bad-v2-dropped-field",
			verbosity: "mutations",
			wantMDSubstrs: []string{
				"- **Validation Status:** FAIL (5/7 checks passed)",
				"- [ ] **V2 (Update):**",
				"`<hint_for_agent>`: DROPPED UPDATE FIELD(S):",
			},
			wantTermSubstrs: []string{
				"## 4. Detailed RPC Payloads",
				"<hint_for_agent>[V2 Update]",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			report, err := AnalyzeFixture(filepath.Join("testdata", tc.dirName))
			if err != nil {
				t.Fatalf("AnalyzeFixture(%s) failed: %v", tc.dirName, err)
			}

			// Act
			md := RenderMarkdown(report)
			termOut := RenderTerminal(report, tc.verbosity)

			// Assert
			for _, sub := range tc.wantMDSubstrs {
				if !strings.Contains(md, sub) {
					t.Errorf("RenderMarkdown(%s) missing expected substring %q", tc.dirName, sub)
				}
			}
			for _, sub := range tc.wantTermSubstrs {
				if !strings.Contains(termOut, sub) {
					t.Errorf("RenderTerminal(%s, %s) missing expected substring %q", tc.dirName, tc.verbosity, sub)
				}
			}
		})
	}
}
