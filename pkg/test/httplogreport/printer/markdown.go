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

package printer

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/classifier"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/specdiff"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/validator"
)

// ReportData holds the complete analyzed data model for rendering a lifecycle report.
type ReportData struct {
	Fixture         *fixture.Data
	Entries         []*classifier.ClassifiedEntry
	PrimaryEntries  []*classifier.ClassifiedEntry
	DependencyCount int
	TotalEntries    int
	SpecDiff        []specdiff.Item
	Checks          []validator.Check
	Passed          bool
}

// Printer renders Markdown and terminal reports and computes deterministic integrity hashes.
type Printer struct{}

// New constructs a new Printer.
func New() *Printer {
	return &Printer{}
}

// RenderMarkdown renders the deterministic _http_report.md content for report.
func (p *Printer) RenderMarkdown(report *ReportData) string {
	var b strings.Builder

	fixtureName := "unknown"
	kind := "Unknown"
	apiVersion := "unknown"
	resName := "unknown"

	if report != nil && report.Fixture != nil {
		fixtureName = report.Fixture.FixtureName
		if report.Fixture.PrimaryCreate != nil {
			kind = report.Fixture.PrimaryCreate.Kind
			apiVersion = report.Fixture.PrimaryCreate.APIVersion
			resName = report.Fixture.PrimaryCreate.Name
			if report.Fixture.PrimaryCreate.ResourceID != "" && report.Fixture.PrimaryCreate.ResourceID != resName {
				resName = fmt.Sprintf("%s (resourceID: %s)", resName, report.Fixture.PrimaryCreate.ResourceID)
			}
		}
	}

	passedCount := 0
	for _, c := range report.Checks {
		if c.Passed {
			passedCount++
		}
	}
	statusStr := fmt.Sprintf("PASS (%d/%d checks passed)", passedCount, len(report.Checks))
	if !report.Passed {
		statusStr = fmt.Sprintf("FAIL (%d/%d checks passed)", passedCount, len(report.Checks))
	}

	b.WriteString(fmt.Sprintf("# HTTP Lifecycle Report: `%s`\n\n", fixtureName))
	b.WriteString(fmt.Sprintf("- **Primary Resource:** `%s` (`%s`)\n", kind, apiVersion))
	b.WriteString(fmt.Sprintf("- **Resource Name:** `%s`\n", resName))
	b.WriteString(fmt.Sprintf("- **Source Log:** `_http.log` (%d total HTTP entries, %d dependency entries)\n", report.TotalEntries, report.DependencyCount))
	b.WriteString(fmt.Sprintf("- **Validation Status:** %s\n", statusStr))
	b.WriteString(fmt.Sprintf("- **Integrity Hash:** `sha256:%s`\n\n", p.ComputeIntegrityHash(report.Fixture)))

	// Section 1: Spec Diff
	b.WriteString("## 1. Spec Diff (`create.yaml` -> `update.yaml`)\n\n")
	if report.Fixture == nil || report.Fixture.PrimaryUpdate == nil {
		b.WriteString("_No `update.yaml` in this fixture._\n\n")
	} else if len(report.SpecDiff) == 0 {
		b.WriteString("_No `.spec` or `.metadata.labels` differences between `create.yaml` and `update.yaml`._\n\n")
	} else {
		b.WriteString("| Field | Change (`create` -> `update`) | Sent in Update? | Final GCP State |\n")
		b.WriteString("| :--- | :--- | :--- | :--- |\n")
		for _, item := range report.SpecDiff {
			fieldCell := formatWrappedFieldPath(item.Path)
			cVal := specdiff.FormatCompactValue(item.CreateVal)
			uVal := specdiff.FormatCompactValue(item.UpdateVal)
			changeCell := fmt.Sprintf("**from:** %s<br>**to:** %s", cVal, uVal)
			sent := item.UpdateRPCDetail
			if sent == "" {
				if item.SentInUpdateRPC {
					sent = "Yes"
				} else {
					sent = "No"
				}
			}
			finalMatch := formatConciseFinalState(item)
			b.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
				escapeMarkdownTableCell(fieldCell),
				escapeMarkdownTableCell(changeCell),
				escapeMarkdownTableCell(sent),
				escapeMarkdownTableCell(finalMatch)))
		}
		b.WriteString("\n")
	}

	// Optional Dependency Summary if dependencies are present
	if report.DependencyCount > 0 {
		b.WriteString("## Dependency Traffic Summary\n\n")
		b.WriteString("| Dependency Kind | Total RPCs | Phases |\n")
		b.WriteString("| :--- | :--- | :--- |\n")
		type depStat struct {
			count  int
			phases map[string]bool
		}
		stats := make(map[string]*depStat)
		var depKinds []string
		for _, ce := range report.Entries {
			if ce.Role != classifier.RoleDependency {
				continue
			}
			dk := ce.ResourceKind
			if dk == "" {
				dk = "Dependency"
			}
			st, ok := stats[dk]
			if !ok {
				st = &depStat{phases: make(map[string]bool)}
				stats[dk] = st
				depKinds = append(depKinds, dk)
			}
			st.count++
			st.phases[ce.Phase] = true
		}
		sort.Strings(depKinds)
		for _, dk := range depKinds {
			st := stats[dk]
			var phs []string
			for _, ph := range []string{logparser.PhaseCreate, logparser.PhaseUpdate, logparser.PhaseRereconcile, logparser.PhaseDelete} {
				if st.phases[ph] {
					phs = append(phs, fmt.Sprintf("`%s`", ph))
				}
			}
			b.WriteString(fmt.Sprintf("| `%s` | %d | %s |\n", dk, st.count, strings.Join(phs, ", ")))
		}
		b.WriteString("\n")
	}

	// Section 2: Primary Resource RPC Summary
	b.WriteString("## 2. Primary Resource RPC Summary\n\n")
	if len(report.PrimaryEntries) == 0 {
		b.WriteString("_No primary resource RPC entries recorded._\n\n")
	} else {
		b.WriteString("| # | Phase | Method / Status | Endpoint | Details / LRO / Mask |\n")
		b.WriteString("| :--- | :--- | :--- | :--- | :--- |\n")
		for _, ce := range report.PrimaryEntries {
			methodStatus := formatMethodStatusCell(ce)
			endpointCell := formatWrappedEndpointCell(ce.ShortEndpoint)
			detailsLROCell := formatDetailsLROCell(ce)
			b.WriteString(fmt.Sprintf("| %d | `%s` | %s | %s | %s |\n",
				ce.Index,
				ce.Phase,
				escapeMarkdownTableCell(methodStatus),
				escapeMarkdownTableCell(endpointCell),
				escapeMarkdownTableCell(detailsLROCell)))
		}
		b.WriteString("\n")
	}

	// Section 3: Automated Validation Checks
	b.WriteString("## 3. Automated Validation Checks\n\n")
	for _, c := range report.Checks {
		box := "x"
		if !c.Passed {
			box = " "
		}
		b.WriteString(fmt.Sprintf("- [%s] **%s (%s):** %s\n", box, c.ID, c.Title, c.Summary))
		if !c.Passed && c.Hint != "" {
			b.WriteString(fmt.Sprintf("  - `<hint_for_agent>`: %s\n", c.Hint))
		}
	}

	return b.String()
}

// ComputeIntegrityHash computes a deterministic 16-hex-char SHA-256 digest over the
// normalized fixture inputs (create.yaml, update.yaml, dependencies.yaml, _http.log,
// and _generated_object_*.golden.yaml) so any hand-edit invalidates the report hash.
func (p *Printer) ComputeIntegrityHash(fd *fixture.Data) string {
	h := sha256.New()
	if fd == nil {
		sum := h.Sum(nil)
		return fmt.Sprintf("%x", sum[:8])
	}

	readNormalizedFile := func(filename string) string {
		if filename == "" || fd.FixtureDir == "" {
			return ""
		}
		path := filepath.Join(fd.FixtureDir, filename)
		b, err := os.ReadFile(path)
		if err != nil {
			return ""
		}
		return strings.TrimSpace(strings.ReplaceAll(string(b), "\r\n", "\n"))
	}

	h.Write([]byte("create.yaml:\n"))
	h.Write([]byte(readNormalizedFile("create.yaml")))
	h.Write([]byte("\nupdate.yaml:\n"))
	h.Write([]byte(readNormalizedFile("update.yaml")))
	h.Write([]byte("\ndependencies.yaml:\n"))
	h.Write([]byte(readNormalizedFile("dependencies.yaml")))
	h.Write([]byte("\n_http.log:\n"))
	cleanLog := strings.TrimSpace(logparser.StripPhaseComments(fd.RawLogContent))
	h.Write([]byte(cleanLog))
	h.Write([]byte("\ngolden.yaml:\n"))
	h.Write([]byte(readNormalizedFile(fd.GoldenFileName)))

	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum[:8])
}

func formatWrappedFieldPath(path string) string {
	if len(path) <= 32 {
		return fmt.Sprintf("`%s`", path)
	}
	segments := strings.Split(path, ".")
	if len(segments) <= 1 {
		return fmt.Sprintf("`%s`", path)
	}
	var parts []string
	cur := segments[0]
	for i := 1; i < len(segments); i++ {
		seg := segments[i]
		if len(cur)+1+len(seg) > 32 && cur != "" {
			parts = append(parts, fmt.Sprintf("`%s.`", cur))
			cur = "." + seg
		} else {
			cur = cur + "." + seg
		}
	}
	if cur != "" {
		parts = append(parts, fmt.Sprintf("`%s`", cur))
	}
	return strings.Join(parts, "<br>")
}

func formatConciseFinalState(item specdiff.Item) string {
	if item.FinalStateMatch {
		return "Yes"
	}
	return "No (mismatch)"
}

func formatMethodStatusCell(ce *classifier.ClassifiedEntry) string {
	method := ce.Entry.Request.Method
	if ce.Entry.Response.StatusCode > 0 {
		return fmt.Sprintf("`%s` `%d`", method, ce.Entry.Response.StatusCode)
	}
	if ce.Entry.Response.Status != "" {
		return fmt.Sprintf("`%s` `%s`", method, ce.Entry.Response.Status)
	}
	return fmt.Sprintf("`%s`", method)
}

func formatWrappedEndpointCell(shortEp string) string {
	if len(shortEp) <= 38 {
		return fmt.Sprintf("`%s`", shortEp)
	}
	pathPart, queryPart, hasQuery := strings.Cut(shortEp, "?")
	if len(pathPart) > 38 {
		if lastSlash := strings.LastIndex(pathPart, "/"); lastSlash > 0 && lastSlash < len(pathPart)-1 {
			prefix := pathPart[:lastSlash+1]
			resID := pathPart[lastSlash+1:]
			if hasQuery && queryPart != "" {
				return fmt.Sprintf("`%s`<br>`%s`<br>`?%s`", prefix, resID, queryPart)
			}
			return fmt.Sprintf("`%s`<br>`%s`", prefix, resID)
		}
	}
	if hasQuery && queryPart != "" {
		return fmt.Sprintf("`%s`<br>`?%s`", pathPart, queryPart)
	}
	return fmt.Sprintf("`%s`", shortEp)
}

func formatDetailsLROCell(ce *classifier.ClassifiedEntry) string {
	lroStr := formatLROCell(ce.LRO)
	details := strings.TrimSpace(ce.Details)
	if lroStr != "" && details != "" {
		return fmt.Sprintf("%s<br>%s", lroStr, details)
	}
	if lroStr != "" {
		return lroStr
	}
	if details != "" {
		return details
	}
	return "-"
}

func formatLROCell(lro *classifier.LROSummary) string {
	if lro == nil {
		return ""
	}
	verb := lro.Verb
	if verb == "" {
		verb = "op"
	}
	status := "done"
	if !lro.Done {
		status = "pending"
	}
	if lro.HasError {
		status = "error"
	}
	if len(lro.PollIndices) == 0 {
		return fmt.Sprintf("LRO: %s -> %s (sync)", verb, status)
	}
	if len(lro.PollIndices) == 1 {
		return fmt.Sprintf("LRO: %s -> %s (#%d)", verb, status, lro.PollIndices[0])
	}
	return fmt.Sprintf("LRO: %s -> %s (#%d..#%d)", verb, status, lro.PollIndices[0], lro.PollIndices[len(lro.PollIndices)-1])
}

func escapeMarkdownTableCell(s string) string {
	return strings.ReplaceAll(s, "|", "\\|")
}

// RenderTerminal renders the report for CLI terminal output with optional verbosity ("summary", "mutations", "full").
func (p *Printer) RenderTerminal(report *ReportData, verbosity string) string {
	var b strings.Builder
	b.WriteString(p.RenderMarkdown(report))

	verbosity = strings.ToLower(strings.TrimSpace(verbosity))

	if verbosity == "mutations" || verbosity == "full" {
		b.WriteString("\n## 4. Detailed RPC Payloads\n\n")
		for _, ce := range report.PrimaryEntries {
			if verbosity == "mutations" && !ce.Entry.IsMutation() {
				continue
			}
			b.WriteString(fmt.Sprintf("### Stanza #%d [%s] `%s %s` -> `%s`\n\n",
				ce.Index, ce.Phase, ce.Entry.Request.Method, ce.Entry.Request.URL, ce.Entry.Response.Status))
			if ce.Entry.Request.Body != "" {
				b.WriteString("**Request Body:**\n```json\n")
				b.WriteString(ce.Entry.Request.Body)
				b.WriteString("\n```\n\n")
			}
			if ce.Entry.Response.Body != "" {
				b.WriteString("**Response Body:**\n```json\n")
				b.WriteString(ce.Entry.Response.Body)
				b.WriteString("\n```\n\n")
			}
			if ce.LRO != nil && ce.LRO.FinalResponse != nil {
				b.WriteString(fmt.Sprintf("**Completed LRO Polls:** `%v`\n\n", ce.LRO.PollIndices))
			}
		}
	}

	if !report.Passed {
		b.WriteString("\n")
		for _, c := range report.Checks {
			if !c.Passed && c.Hint != "" {
				b.WriteString(fmt.Sprintf("<hint_for_agent>[%s %s] %s</hint_for_agent>\n", c.ID, c.Title, c.Hint))
			}
		}
	}

	return b.String()
}
