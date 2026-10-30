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
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
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

func TestTier1ExplicitVsTier2InferredParityOnFixtures(t *testing.T) {
	fixtureDirs := []string{
		filepath.Join("testdata", "good-composite-create"),
		filepath.Join("testdata", "good-field-by-field-update"),
		filepath.Join("testdata", "good-maskless-put-update"),
		filepath.Join("..", "resourcefixture", "testdata", "basic", "agentregistry", "v1alpha1", "agentregistrybinding", "agentregistrybinding-maximal"),
		filepath.Join("..", "resourcefixture", "testdata", "basic", "agentregistry", "v1alpha1", "agentregistrybinding", "agentregistrybinding-minimal"),
	}

	for _, dir := range fixtureDirs {
		dir := dir
		t.Run(filepath.Base(dir), func(t *testing.T) {
			// Arrange
			rawBytes, err := os.ReadFile(filepath.Join(dir, "_http.log"))
			if err != nil {
				t.Fatalf("reading _http.log in %s: %v", dir, err)
			}
			unmarkedLog := StripPhaseComments(string(rawBytes))

			// Act
			reportInferred, err := AnalyzeFixtureWithLogContent(dir, unmarkedLog)
			if err != nil {
				t.Fatalf("AnalyzeFixtureWithLogContent(inferred, %s) failed: %v", dir, err)
			}
			if reportInferred.PhaseSource != PhaseSourceInferred {
				t.Fatalf("expected PhaseSourceInferred, got %s", reportInferred.PhaseSource)
			}
			if !reportInferred.Passed {
				t.Fatalf("expected inferred report for %s to pass all checks", dir)
			}

			entries, err := ParseHTTPLog(unmarkedLog)
			if err != nil {
				t.Fatalf("ParseHTTPLog(%s) failed: %v", dir, err)
			}
			for i, ce := range reportInferred.Entries {
				entries[i].Phase = ce.Phase
			}
			explicitLog := FormatWithPhaseComments(entries)
			reportExplicit, err := AnalyzeFixtureWithLogContent(dir, explicitLog)
			if err != nil {
				t.Fatalf("AnalyzeFixtureWithLogContent(explicit, %s) failed: %v", dir, err)
			}

			// Assert
			if reportExplicit.PhaseSource != PhaseSourceExplicit {
				t.Fatalf("expected PhaseSourceExplicit, got %s", reportExplicit.PhaseSource)
			}
			if !reportExplicit.Passed {
				t.Fatalf("expected explicit report for %s to pass all checks", dir)
			}

			if len(reportInferred.Entries) != len(reportExplicit.Entries) {
				t.Fatalf("entry count mismatch: inferred=%d, explicit=%d", len(reportInferred.Entries), len(reportExplicit.Entries))
			}
			for i := range reportInferred.Entries {
				if reportInferred.Entries[i].Phase != reportExplicit.Entries[i].Phase {
					t.Errorf("entry #%d Phase mismatch: inferred=%s, explicit=%s", i+1, reportInferred.Entries[i].Phase, reportExplicit.Entries[i].Phase)
				}
				if reportInferred.Entries[i].Role != reportExplicit.Entries[i].Role {
					t.Errorf("entry #%d Role mismatch: inferred=%s, explicit=%s", i+1, reportInferred.Entries[i].Role, reportExplicit.Entries[i].Role)
				}
			}
			if !reflect.DeepEqual(reportInferred.SpecDiff, reportExplicit.SpecDiff) {
				t.Errorf("SpecDiff mismatch between inferred and explicit reports:\ninferred: %+v\nexplicit: %+v", reportInferred.SpecDiff, reportExplicit.SpecDiff)
			}
			if !reflect.DeepEqual(reportInferred.Checks, reportExplicit.Checks) {
				t.Errorf("Checks mismatch between inferred and explicit reports:\ninferred: %+v\nexplicit: %+v", reportInferred.Checks, reportExplicit.Checks)
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

func TestCommittedHTTPLogReports(t *testing.T) {
	basicRoot := filepath.Join("..", "resourcefixture", "testdata", "basic")
	var reportDirs []string
	var phaseStampedLogDirs []string

	err := filepath.WalkDir(basicRoot, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			return nil
		}
		switch d.Name() {
		case "_http_report.md":
			reportDirs = append(reportDirs, filepath.Dir(p))
		case "_http.log":
			if b, readErr := os.ReadFile(p); readErr == nil && strings.Contains(string(b), "# Phase:") {
				phaseStampedLogDirs = append(phaseStampedLogDirs, filepath.Dir(p))
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking %s: %v", basicRoot, err)
	}
	if len(reportDirs) == 0 {
		t.Fatalf("expected at least 1 committed _http_report.md under %s", basicRoot)
	}

	for _, dir := range phaseStampedLogDirs {
		reportPath := filepath.Join(dir, "_http_report.md")
		if !fixture.FileExists(reportPath) {
			t.Errorf("fixture %s has # Phase: markers in _http.log, but _http_report.md is missing. Run: go run ./cmd/httplog-report --fixture %s --write", dir, dir)
		}
	}

	for _, dir := range reportDirs {
		dir := dir
		t.Run(filepath.Base(dir), func(t *testing.T) {
			report, err := AnalyzeFixture(dir)
			if err != nil {
				t.Fatalf("AnalyzeFixture(%s) failed: %v", dir, err)
			}
			if !report.Passed {
				t.Errorf("fixture %s failed HTTP log validation:\n%s", dir, RenderTerminal(report, "normal"))
			}
			committedBytes, err := os.ReadFile(filepath.Join(dir, "_http_report.md"))
			if err != nil {
				t.Fatalf("reading committed _http_report.md: %v", err)
			}
			expectedMD := RenderMarkdown(report)
			if string(committedBytes) != expectedMD {
				t.Errorf("committed _http_report.md in %s is out of sync with fixture files (or Integrity Hash mismatch).\nRun: go run ./cmd/httplog-report --fixture %s --write", dir, dir)
			}
		})
	}
}
