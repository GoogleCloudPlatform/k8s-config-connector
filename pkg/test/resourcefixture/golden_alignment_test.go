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

package resourcefixture

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

var mockGCPSkipGroupKinds = map[schema.GroupKind]bool{}

func TestGoldenLogAlignment(t *testing.T) {
	rootDir := "testdata/basic"
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		t.Fatalf("failed to get absolute path for %s: %v", rootDir, err)
	}

	err = filepath.WalkDir(absRootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			realLogPath := filepath.Join(path, "_http.log")
			mockLogPath := filepath.Join(path, "_http_mock.log")

			if fileExists(realLogPath) {
				createPath := filepath.Join(path, "create.yaml")
				if fileExists(createPath) {
					gvk, err := getGVKFromYAML(createPath)
					if err == nil {
						gk := gvk.GroupKind()
						if !mockGCPSkipGroupKinds[gk] && !fileExists(mockLogPath) {
							t.Errorf("fixture %q: resource must have _http_mock.log golden file", path)
						}
					}
				}

				if fileExists(mockLogPath) {
					relPath, _ := filepath.Rel(absRootDir, path)
					t.Run(relPath, func(t *testing.T) {
						compareLogs(t, realLogPath, mockLogPath)
					})
				}
			}
		}

		return nil
	})

	if err != nil {
		t.Fatalf("error walking directory: %v", err)
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

type pathMethodEvents map[string]map[string][]httpEvent

func groupByPathAndMethod(events []httpEvent) pathMethodEvents {
	grouped := make(pathMethodEvents)
	for _, ev := range events {
		if ev.Method == "GET" {
			continue // Skip GET entirely
		}
		basePath := strings.Split(cleanURL(ev.URL), "?")[0]
		if _, ok := grouped[basePath]; !ok {
			grouped[basePath] = make(map[string][]httpEvent)
		}
		grouped[basePath][ev.Method] = append(grouped[basePath][ev.Method], ev)
	}
	return grouped
}

func compareLogs(t *testing.T, realPath, mockPath string) {
	realEvents := readLog(t, realPath)
	mockEvents := readLog(t, mockPath)

	realGrouped := groupByPathAndMethod(realEvents)
	mockGrouped := groupByPathAndMethod(mockEvents)

	compareGroupedLogs(t, realGrouped, mockGrouped)
}

func readLog(t *testing.T, path string) []httpEvent {
	bytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}
	return parseLog(t, string(bytes))
}

func normalizeAPIVersion(path string) string {
	// Replaces path segments like "/v1/", "/v1beta1/", "/v1beta2/", "/v2/", "/v3/", "/v1alpha1/" etc.
	// with "/api_version/"
	re := regexp.MustCompile(`/(v[0-9]+[a-zA-Z0-9]*)/`)
	path = re.ReplaceAllString(path, "/api_version/")

	// Normalize project number and project ID placeholders
	path = strings.ReplaceAll(path, "${projectNumber}", "_project_")
	path = strings.ReplaceAll(path, "${projectId}", "_project_")
	return path
}

func getProjectID(path string) string {
	re := regexp.MustCompile(`/projects/([^/]+)`)
	matches := re.FindStringSubmatch(path)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func hasDeletedParent(path string, mockGrouped pathMethodEvents) bool {
	normalizedPath := normalizeAPIVersion(path)
	segments := strings.Split(normalizedPath, "/")

	// Create a map of normalized mockGrouped paths to check
	normalizedMockPaths := make(map[string]map[string][]httpEvent)
	for mockPath, methods := range mockGrouped {
		normalizedMockPaths[normalizeAPIVersion(mockPath)] = methods
	}

	// 1. Standard prefix-based parent check
	for i := len(segments) - 1; i > 0; i-- {
		parentPath := strings.Join(segments[:i], "/")
		if parentPath == "" {
			continue
		}
		if parentMethods, ok := normalizedMockPaths[parentPath]; ok {
			if deleteEvs, found := parentMethods["DELETE"]; found && len(deleteEvs) > 0 {
				return true
			}
		}
	}

	// 2. Sibling dependency check (e.g. Subnetwork/Route/Firewall depending on Network)
	projectID := getProjectID(normalizedPath)
	if projectID != "" {
		isNetworkDependent := strings.Contains(path, "/subnetworks") ||
			strings.Contains(path, "/routes") ||
			strings.Contains(path, "/firewalls") ||
			strings.Contains(path, "/servicenetworking")

		if isNetworkDependent {
			for mockPath, methods := range normalizedMockPaths {
				if strings.Contains(mockPath, "/networks/") && getProjectID(mockPath) == projectID {
					if deleteEvs, found := methods["DELETE"]; found && len(deleteEvs) > 0 {
						return true
					}
				}
			}
		}
	}

	return false
}

func isMock404OrEmptyOnDeletedParent(path string, mockEv httpEvent, mockGrouped pathMethodEvents) bool {
	if !hasDeletedParent(path, mockGrouped) {
		return false
	}
	if strings.Contains(mockEv.Status, "404") {
		return true
	}
	if strings.Contains(mockEv.ResponseBody, `"code": 404`) || strings.Contains(mockEv.ResponseBody, `"code":404`) {
		return true
	}
	return false
}

func compareGroupedLogs(t *testing.T, realGrouped, mockGrouped pathMethodEvents) {
	// Check all paths in realGrouped
	for path, realMethods := range realGrouped {
		mockMethods, pathExistsInMock := mockGrouped[path]

		for method, realEvs := range realMethods {
			mockEvs := mockMethods[method]

			if !pathExistsInMock {
				// If DELETE is missing entirely, we check if it is allowed via deleted parent
				if method == "DELETE" && hasDeletedParent(path, mockGrouped) {
					continue
				}
				t.Errorf("path %q present in real log but missing in mock log", path)
				continue
			}

			if len(mockEvs) == 0 {
				if method == "DELETE" && hasDeletedParent(path, mockGrouped) {
					continue
				}
				t.Errorf("path %q: method %s present in real log but missing in mock log", path, method)
				continue
			}

			// Sort events by their RequestBody to ensure deterministic order for concurrent sibling operations
			sort.SliceStable(realEvs, func(i, j int) bool {
				return realEvs[i].RequestBody < realEvs[j].RequestBody
			})
			sort.SliceStable(mockEvs, func(i, j int) bool {
				return mockEvs[i].RequestBody < mockEvs[j].RequestBody
			})

			if len(realEvs) != len(mockEvs) {
				allowed := false
				if method == "DELETE" && len(mockEvs) < len(realEvs) {
					if hasDeletedParent(path, mockGrouped) {
						allowed = true
					}
				}
				if len(mockEvs) > len(realEvs) {
					allowed = true // Allow extra retries/reconciliations in mock
				}
				if !allowed {
					t.Errorf("path %q, method %s: mismatched number of calls: real has %d, mock has %d", path, method, len(realEvs), len(mockEvs))
					continue
				}
			}

			// Compare only up to the number of calls present in both (or up to realEvs size if mock is larger)
			compareCount := len(mockEvs)
			if len(realEvs) < compareCount {
				compareCount = len(realEvs)
			}

			for i := 0; i < compareCount; i++ {
				if isMock404OrEmptyOnDeletedParent(path, mockEvs[i], mockGrouped) {
					continue
				}
				compareJSON(t, fmt.Sprintf("path %s, method %s, call %d request body", path, method, i), realEvs[i].RequestBody, mockEvs[i].RequestBody)
				compareJSON(t, fmt.Sprintf("path %s, method %s, call %d response body", path, method, i), realEvs[i].ResponseBody, mockEvs[i].ResponseBody)
			}
		}
	}

	// Also check if mockGrouped has any paths/methods that realGrouped doesn't have!
	for path, mockMethods := range mockGrouped {
		realMethods, pathExistsInReal := realGrouped[path]
		if !pathExistsInReal {
			t.Errorf("path %q present in mock log but missing in real log", path)
			continue
		}
		for method, mockEvs := range mockMethods {
			realEvs := realMethods[method]
			if len(realEvs) == 0 && len(mockEvs) > 0 {
				t.Errorf("path %q: method %s present in mock log but missing in real log", path, method)
			}
		}
	}
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
		// Skip request headers
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

		// Skip response headers
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

func cleanURL(u string) string {
	if protoIdx := strings.Index(u, "://"); protoIdx != -1 {
		u = u[protoIdx+3:]
	}
	if slashIdx := strings.Index(u, "/"); slashIdx != -1 {
		u = u[slashIdx:]
	}
	return u
}

func compareJSON(t *testing.T, context, realJSON, mockJSON string) {
	if realJSON == "" && mockJSON == "" {
		return
	}

	var realObj, mockObj interface{}

	if realJSON != "" {
		if err := json.Unmarshal([]byte(realJSON), &realObj); err != nil {
			if realJSON != mockJSON {
				t.Errorf("%s: string mismatch:\n  real: %q\n  mock: %q", context, realJSON, mockJSON)
			}
			return
		}
	}

	if mockJSON != "" {
		if err := json.Unmarshal([]byte(mockJSON), &mockObj); err != nil {
			if realJSON != mockJSON {
				t.Errorf("%s: string mismatch:\n  real: %q\n  mock: %q", context, realJSON, mockJSON)
			}
			return
		}
	}

	if !reflect.DeepEqual(realObj, mockObj) {
		t.Errorf("%s: JSON mismatch:\n  real: %s\n  mock: %s", context, realJSON, mockJSON)
	}
}

func getGVKFromYAML(path string) (schema.GroupVersionKind, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return schema.GroupVersionKind{}, err
	}
	var u unstructured.Unstructured
	if err := yaml.Unmarshal(bytes, &u); err != nil {
		return schema.GroupVersionKind{}, err
	}
	return u.GroupVersionKind(), nil
}
