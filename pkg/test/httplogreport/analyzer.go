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
	"fmt"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/classifier"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/fixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/logparser"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/printer"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/specdiff"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/httplogreport/validator"
)

// Re-exported constants and type aliases for callers of pkg/test/httplogreport.
const (
	PhaseCreate      = logparser.PhaseCreate
	PhaseUpdate      = logparser.PhaseUpdate
	PhaseRereconcile = logparser.PhaseRereconcile
	PhaseDelete      = logparser.PhaseDelete

	PhaseSourceExplicit = classifier.PhaseSourceExplicit
	PhaseSourceInferred = classifier.PhaseSourceInferred

	RolePrimary    = classifier.RolePrimary
	RoleDependency = classifier.RoleDependency
)

type (
	HTTPEntry       = logparser.Entry
	HTTPRequest     = logparser.HTTPRequest
	HTTPResponse    = logparser.HTTPResponse
	FixtureData     = fixture.Data
	ResourceMeta    = fixture.KRMResource
	ClassifiedEntry = classifier.ClassifiedEntry
	LROSummary      = classifier.LROSummary
	SpecDiffItem    = specdiff.Item
	ValidationCheck = validator.Check
)

// Report holds the complete analyzed and validated lifecycle data for a fixture.
type Report struct {
	Fixture         *fixture.Data
	Entries         []*classifier.ClassifiedEntry
	PrimaryEntries  []*classifier.ClassifiedEntry
	DependencyCount int
	TotalEntries    int
	PhaseSource     string
	SpecDiff        []specdiff.Item
	Checks          []validator.Check
	Passed          bool
}

// Analyzer orchestrates fixture loading, HTTP log parsing, classification,
// spec diff correlation, V1-V7 validation, and Markdown/terminal rendering.
type Analyzer struct {
	printer *printer.Printer
}

// NewAnalyzer constructs a new Analyzer with default subpackage collaborators.
func NewAnalyzer() *Analyzer {
	return &Analyzer{
		printer: printer.New(),
	}
}

// AnalyzeFixture loads, classifies, diffs, and validates a test fixture directory.
func AnalyzeFixture(fixtureDir string) (*Report, error) {
	return NewAnalyzer().AnalyzeFixture(fixtureDir)
}

// AnalyzeFixtureWithLogContent analyzes a fixture directory using in-memory rawLogContent
// instead of reading _http.log from disk.
func AnalyzeFixtureWithLogContent(fixtureDir string, rawLogContent string) (*Report, error) {
	return NewAnalyzer().AnalyzeFixtureWithLogContent(fixtureDir, rawLogContent)
}

// AnalyzeFixture loads, classifies, diffs, and validates a test fixture directory.
func (a *Analyzer) AnalyzeFixture(fixtureDir string) (*Report, error) {
	fd, err := fixture.Load(fixtureDir)
	if err != nil {
		return nil, err
	}
	return a.analyzeLoadedFixture(fd)
}

// AnalyzeFixtureWithLogContent analyzes a fixture directory using in-memory rawLogContent.
func (a *Analyzer) AnalyzeFixtureWithLogContent(fixtureDir string, rawLogContent string) (*Report, error) {
	fd, err := fixture.LoadWithLogContent(fixtureDir, rawLogContent)
	if err != nil {
		return nil, err
	}
	return a.analyzeLoadedFixture(fd)
}

func (a *Analyzer) analyzeLoadedFixture(fd *fixture.Data) (*Report, error) {
	rawEntries, err := logparser.Parse(fd.RawLogContent)
	if err != nil {
		return nil, fmt.Errorf("parsing _http.log in %s: %w", fd.FixtureDir, err)
	}

	clf := classifier.New(fd)
	classified, phaseSource := clf.Classify(rawEntries)

	var primaryEntries []*classifier.ClassifiedEntry
	depCount := 0
	for _, ce := range classified {
		if ce.Role == classifier.RolePrimary && !ce.IsLROPoll {
			primaryEntries = append(primaryEntries, ce)
		}
		if ce.Role == classifier.RoleDependency {
			depCount++
		}
	}

	var diffs []specdiff.Item
	if fd.PrimaryCreate != nil && fd.PrimaryUpdate != nil {
		differ := specdiff.New(fd, primaryEntries)
		diffs = differ.Compute()
	}

	val := validator.New(fd, classified, primaryEntries, diffs)
	checks, passed := val.Validate()

	return &Report{
		Fixture:         fd,
		Entries:         classified,
		PrimaryEntries:  primaryEntries,
		DependencyCount: depCount,
		TotalEntries:    len(classified),
		PhaseSource:     phaseSource,
		SpecDiff:        diffs,
		Checks:          checks,
		Passed:          passed,
	}, nil
}

// RenderMarkdown renders the deterministic _http_report.md content for report.
func RenderMarkdown(report *Report) string {
	return NewAnalyzer().RenderMarkdown(report)
}

// RenderMarkdown renders the deterministic _http_report.md content for report.
func (a *Analyzer) RenderMarkdown(report *Report) string {
	return a.printer.RenderMarkdown(toPrinterData(report))
}

// RenderTerminal renders the report for CLI terminal output with optional verbosity.
func RenderTerminal(report *Report, verbosity string) string {
	return NewAnalyzer().RenderTerminal(report, verbosity)
}

// RenderTerminal renders the report for CLI terminal output with optional verbosity.
func (a *Analyzer) RenderTerminal(report *Report, verbosity string) string {
	return a.printer.RenderTerminal(toPrinterData(report), verbosity)
}

// ComputeIntegrityHash computes the 16-hex-char SHA-256 digest over the normalized fixture inputs.
func ComputeIntegrityHash(report *Report) string {
	return NewAnalyzer().ComputeIntegrityHash(report)
}

// ComputeIntegrityHash computes the 16-hex-char SHA-256 digest over the normalized fixture inputs.
func (a *Analyzer) ComputeIntegrityHash(report *Report) string {
	if report == nil {
		return a.printer.ComputeIntegrityHash(nil)
	}
	return a.printer.ComputeIntegrityHash(report.Fixture)
}

// WriteReport writes _http_report.md into fixtureDir and returns the written file path.
func WriteReport(fixtureDir string, report *Report) (string, error) {
	return NewAnalyzer().WriteReport(fixtureDir, report)
}

// WriteReport writes _http_report.md into fixtureDir and returns the written file path.
func (a *Analyzer) WriteReport(fixtureDir string, report *Report) (string, error) {
	outPath := filepath.Join(fixtureDir, "_http_report.md")
	content := a.RenderMarkdown(report)
	if err := os.WriteFile(outPath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("writing %s: %w", outPath, err)
	}
	return outPath, nil
}

func toPrinterData(report *Report) *printer.ReportData {
	if report == nil {
		return &printer.ReportData{}
	}
	return &printer.ReportData{
		Fixture:         report.Fixture,
		Entries:         report.Entries,
		PrimaryEntries:  report.PrimaryEntries,
		DependencyCount: report.DependencyCount,
		TotalEntries:    report.TotalEntries,
		SpecDiff:        report.SpecDiff,
		Checks:          report.Checks,
		Passed:          report.Passed,
	}
}

// DiscoverFixtureDirs finds all fixture directories under rootPath containing both create.yaml and _http.log.
func DiscoverFixtureDirs(rootPath string) ([]string, error) {
	return fixture.DiscoverDirs(rootPath)
}

// ParseHTTPLog parses a raw _http.log string into a slice of HTTPEntry values.
func ParseHTTPLog(raw string) ([]*HTTPEntry, error) {
	return logparser.Parse(raw)
}

// StripPhaseComments removes all `# Phase:` marker lines from raw _http.log content.
func StripPhaseComments(raw string) string {
	return logparser.StripPhaseComments(raw)
}

// FormatWithPhaseComments renders a slice of HTTPEntry back into canonical _http.log text.
func FormatWithPhaseComments(entries []*HTTPEntry) string {
	return logparser.FormatWithPhaseComments(entries)
}
