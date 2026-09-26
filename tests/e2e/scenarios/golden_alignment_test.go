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

package scenarios

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/yaml"
)

func loadLegacyScenarios(t *testing.T, path string) map[string]bool {
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read legacy scenarios: %v", err)
	}
	m := make(map[string]bool)
	for _, line := range strings.Split(string(raw), "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			m[line] = true
		}
	}
	return m
}

func TestGoldenAlignment(t *testing.T) {
	legacyScenariosPath := "../testdata/legacy_scenarios.txt"
	legacyScenarios := loadLegacyScenarios(t, legacyScenariosPath)

	scenarioDir := "../testdata/scenarios"
	absScenarioDir, err := filepath.Abs(scenarioDir)
	if err != nil {
		t.Fatalf("failed to get absolute path for %s: %v", scenarioDir, err)
	}

	err = filepath.WalkDir(absScenarioDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			scriptPath := filepath.Join(path, "script.yaml")
			if fileExists(scriptPath) {
				relPath, _ := filepath.Rel(absScenarioDir, path)
				parts := strings.Split(relPath, string(filepath.Separator))
				if len(parts) == 0 {
					return nil
				}
				suite := parts[0]
				if legacyScenarios[suite] {
					return nil
				}

				t.Run(relPath, func(t *testing.T) {
					// Scan steps 0 to 99
					for i := 0; i < 100; i++ {
						realHTTPPath := filepath.Join(path, fmt.Sprintf("_http%02d.log", i))
						mockHTTPPath := filepath.Join(path, fmt.Sprintf("_http%02d_mock.log", i))

						if fileExists(realHTTPPath) && fileExists(mockHTTPPath) {
							t.Run(fmt.Sprintf("http-step%02d", i), func(t *testing.T) {
								compareHTTPLogs(t, realHTTPPath, mockHTTPPath)
							})
						}

						realObjPath := filepath.Join(path, fmt.Sprintf("_object%02d.yaml", i))
						mockObjPath := filepath.Join(path, fmt.Sprintf("_object%02d_mock.yaml", i))

						if fileExists(realObjPath) && fileExists(mockObjPath) {
							t.Run(fmt.Sprintf("object-step%02d", i), func(t *testing.T) {
								compareKRMObjects(t, realObjPath, mockObjPath)
							})
						}

						realExportPath := filepath.Join(path, fmt.Sprintf("_export%d.yaml", i))
						mockExportPath := filepath.Join(path, fmt.Sprintf("_export%d_mock.yaml", i))

						if fileExists(realExportPath) && fileExists(mockExportPath) {
							t.Run(fmt.Sprintf("export-step%d", i), func(t *testing.T) {
								compareKRMObjects(t, realExportPath, mockExportPath)
							})
						}
					}
				})
			}
		}

		return nil
	})

	if err != nil {
		t.Fatalf("error walking directory: %v", err)
	}
}

func compareHTTPLogs(t *testing.T, realPath, mockPath string) {
	realBytes, err := os.ReadFile(realPath)
	if err != nil {
		t.Fatalf("failed to read %s: %v", realPath, err)
	}
	mockBytes, err := os.ReadFile(mockPath)
	if err != nil {
		t.Fatalf("failed to read %s: %v", mockPath, err)
	}

	realEvents := parseLog(t, string(realBytes))
	mockEvents := parseLog(t, string(mockBytes))

	if len(realEvents) != len(mockEvents) {
		t.Fatalf("mismatched number of HTTP events: real has %d, mock has %d", len(realEvents), len(mockEvents))
	}

	for i := 0; i < len(realEvents); i++ {
		r := realEvents[i]
		m := mockEvents[i]

		if r.Method != m.Method {
			t.Errorf("Event %d: Method mismatch: real %q, mock %q", i, r.Method, m.Method)
		}

		if r.URL != m.URL {
			t.Errorf("Event %d: URL mismatch: real %q, mock %q", i, r.URL, m.URL)
		}

		compareBodies(t, fmt.Sprintf("Event %d Request Body", i), r.RequestBody, m.RequestBody)
		compareBodies(t, fmt.Sprintf("Event %d Response Body", i), r.ResponseBody, m.ResponseBody)
	}
}

func compareBodies(t *testing.T, context string, realBody, mockBody string) {
	if realBody == "" && mockBody == "" {
		return
	}

	var realJSON, mockJSON interface{}
	realErr := json.Unmarshal([]byte(realBody), &realJSON)
	mockErr := json.Unmarshal([]byte(mockBody), &mockJSON)

	if realErr == nil && mockErr == nil {
		if diff := cmp.Diff(realJSON, mockJSON); diff != "" {
			t.Errorf("%s mismatch (-real +mock):\n%s", context, diff)
		}
	} else {
		if diff := cmp.Diff(realBody, mockBody); diff != "" {
			t.Errorf("%s mismatch (-real +mock):\n%s", context, diff)
		}
	}
}

func compareKRMObjects(t *testing.T, realPath, mockPath string) {
	realBytes, err := os.ReadFile(realPath)
	if err != nil {
		t.Fatalf("failed to read %s: %v", realPath, err)
	}
	mockBytes, err := os.ReadFile(mockPath)
	if err != nil {
		t.Fatalf("failed to read %s: %v", mockPath, err)
	}

	var realObj, mockObj map[string]interface{}
	if err := yaml.Unmarshal(realBytes, &realObj); err != nil {
		t.Fatalf("failed to unmarshal real KRM object %s: %v", realPath, err)
	}
	if err := yaml.Unmarshal(mockBytes, &mockObj); err != nil {
		t.Fatalf("failed to unmarshal mock KRM object %s: %v", mockPath, err)
	}

	cleanKRMFields(realObj)
	cleanKRMFields(mockObj)

	if diff := cmp.Diff(realObj, mockObj); diff != "" {
		t.Errorf("KRM object mismatch (-real +mock):\n%s", diff)
	}
}

func cleanKRMFields(obj map[string]interface{}) {
	if metadata, ok := obj["metadata"].(map[string]interface{}); ok {
		delete(metadata, "resourceVersion")
		delete(metadata, "uid")
		delete(metadata, "creationTimestamp")
		delete(metadata, "generation")
	}
}

type httpEvent struct {
	Method       string
	URL          string
	RequestBody  string
	Status       string
	ResponseBody string
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

var statusRegex = regexp.MustCompile(`^\d{3} `)

func parseLog(t *testing.T, content string) []httpEvent {
	var events []httpEvent
	rawEvents := strings.Split(content, "\n---\n")

	for _, raw := range rawEvents {
		raw = strings.TrimSpace(raw)
		if raw == "" {
			continue
		}

		lines := strings.Split(raw, "\n")
		var ev httpEvent

		reqParts := strings.SplitN(lines[0], " ", 2)
		if len(reqParts) < 2 {
			continue
		}
		ev.Method = reqParts[0]
		ev.URL = reqParts[1]

		idx := 1
		for idx < len(lines) && strings.TrimSpace(lines[idx]) != "" {
			idx++
		}
		if idx < len(lines) {
			idx++
		}

		var reqBodyLines []string
		for idx < len(lines) && !statusRegex.MatchString(lines[idx]) {
			reqBodyLines = append(reqBodyLines, lines[idx])
			idx++
		}
		ev.RequestBody = strings.TrimSpace(strings.Join(reqBodyLines, "\n"))

		if idx < len(lines) {
			ev.Status = lines[idx]
			idx++
		}

		for idx < len(lines) && strings.TrimSpace(lines[idx]) != "" {
			idx++
		}
		if idx < len(lines) {
			idx++
		}

		var respBodyLines []string
		for idx < len(lines) {
			respBodyLines = append(respBodyLines, lines[idx])
			idx++
		}
		ev.ResponseBody = strings.TrimSpace(strings.Join(respBodyLines, "\n"))

		events = append(events, ev)
	}

	return events
}
