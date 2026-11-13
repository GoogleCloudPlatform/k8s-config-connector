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

package linter

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	DeletionPolicyAnnotation = "cnrm.cloud.google.com/deletion-policy"
	ConflictPolicyAnnotation = "cnrm.cloud.google.com/management-conflict-prevention-policy"
	StateIntoSpecAnnotation  = "cnrm.cloud.google.com/state-into-spec"
)

type LintIssue struct {
	FilePath string
	Kind     string
	Name     string
	Severity string // "ERROR" or "WARNING"
	Message  string
}

type LintReport struct {
	TotalFiles     int
	TotalResources int
	Issues         []LintIssue
	Passed         bool
}

// LintDirectory walks through directory and validates all YAML files against migration requirements.
func LintDirectory(dirPath string, strictSafety bool) (*LintReport, error) {
	report := &LintReport{
		Passed: true,
	}

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".yaml") && !strings.HasSuffix(path, ".yml") {
			return nil
		}

		report.TotalFiles++
		data, err := os.ReadFile(path)
		if err != nil {
			report.Issues = append(report.Issues, LintIssue{
				FilePath: path,
				Severity: "ERROR",
				Message:  fmt.Sprintf("Failed to read file: %v", err),
			})
			report.Passed = false
			return nil
		}

		decoder := yaml.NewDecoder(bytes.NewReader(data))
		for {
			var doc map[string]interface{}
			if err := decoder.Decode(&doc); err != nil {
				if err == io.EOF {
					break
				}
				report.Issues = append(report.Issues, LintIssue{
					FilePath: path,
					Severity: "ERROR",
					Message:  fmt.Sprintf("YAML decode error: %v", err),
				})
				report.Passed = false
				break
			}
			if len(doc) == 0 {
				continue
			}

			report.TotalResources++
			issues := LintResource(path, doc, strictSafety)
			for _, issue := range issues {
				report.Issues = append(report.Issues, issue)
				if issue.Severity == "ERROR" {
					report.Passed = false
				}
			}
		}
		return nil
	})

	return report, err
}

// LintResource checks a single resource.
func LintResource(filePath string, doc map[string]interface{}, strictSafety bool) []LintIssue {
	var issues []LintIssue

	apiVersion, _ := doc["apiVersion"].(string)
	kind, _ := doc["kind"].(string)

	rawMetadata, _ := doc["metadata"].(map[string]interface{})
	name := "unnamed"
	if rawMetadata != nil {
		if n, ok := rawMetadata["name"].(string); ok {
			name = n
		}
	}

	// 1. Check for basic K8s fields
	if apiVersion == "" {
		issues = append(issues, LintIssue{
			FilePath: filePath,
			Kind:     kind,
			Name:     name,
			Severity: "ERROR",
			Message:  "Missing 'apiVersion'",
		})
	}
	if kind == "" {
		issues = append(issues, LintIssue{
			FilePath: filePath,
			Kind:     kind,
			Name:     name,
			Severity: "ERROR",
			Message:  "Missing 'kind'",
		})
	}

	// 2. Check if status block leaked through
	if _, hasStatus := doc["status"]; hasStatus {
		issues = append(issues, LintIssue{
			FilePath: filePath,
			Kind:     kind,
			Name:     name,
			Severity: "ERROR",
			Message:  "Resource contains a 'status' block; must be stripped before migration",
		})
	}

	// 3. If KCC managed resource, verify safety annotations (exclude core control plane configs)
	if strings.Contains(apiVersion, ".cnrm.cloud.google.com") && !strings.HasPrefix(apiVersion, "core.cnrm.cloud.google.com") {
		var annotations map[string]interface{}
		if rawMetadata != nil {
			if ann, ok := rawMetadata["annotations"].(map[string]interface{}); ok {
				annotations = ann
			}
		}

		deletionPolicy, _ := annotations[DeletionPolicyAnnotation].(string)
		if deletionPolicy != "abandon" {
			sev := "WARNING"
			if strictSafety {
				sev = "ERROR"
			}
			issues = append(issues, LintIssue{
				FilePath: filePath,
				Kind:     kind,
				Name:     name,
				Severity: sev,
				Message:  fmt.Sprintf("Missing '%s: abandon'. Required to prevent GCP resource deletion upon CR delete.", DeletionPolicyAnnotation),
			})
		}

		conflictPolicy, _ := annotations[ConflictPolicyAnnotation].(string)
		if conflictPolicy != "none" {
			sev := "WARNING"
			if strictSafety {
				sev = "ERROR"
			}
			issues = append(issues, LintIssue{
				FilePath: filePath,
				Kind:     kind,
				Name:     name,
				Severity: sev,
				Message:  fmt.Sprintf("Missing '%s: none'. Required to prevent adoption conflicts on pre-existing GCP resources.", ConflictPolicyAnnotation),
			})
		}

		stateIntoSpec, _ := annotations[StateIntoSpecAnnotation].(string)
		if stateIntoSpec != "absent" {
			issues = append(issues, LintIssue{
				FilePath: filePath,
				Kind:     kind,
				Name:     name,
				Severity: "WARNING",
				Message:  fmt.Sprintf("Missing '%s: absent'. Recommended to prevent immutability webhook rejection on subsequent 3-way merge apply.", StateIntoSpecAnnotation),
			})
		}

		// Check for server-managed runtime leaks in metadata
		if rawMetadata != nil {
			for _, forbidden := range []string{"uid", "resourceVersion", "generation", "managedFields"} {
				if _, exists := rawMetadata[forbidden]; exists {
					issues = append(issues, LintIssue{
						FilePath: filePath,
						Kind:     kind,
						Name:     name,
						Severity: "ERROR",
						Message:  fmt.Sprintf("Metadata contains runtime field '%s'; must be stripped", forbidden),
					})
				}
			}
		}
	}

	return issues
}
