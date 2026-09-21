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
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/yaml"
)

var legacyScenarios = func() map[string]bool {
	m := make(map[string]bool)
	legacyScenariosRaw, err := os.ReadFile("../testdata/legacy_scenarios.txt")
	if err != nil {
		panic(fmt.Sprintf("failed to read legacy_scenarios.txt: %v", err))
	}
	lines := strings.Split(string(legacyScenariosRaw), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		m[line] = true
	}
	return m
}()

func TestGoldenAlignment(t *testing.T) {
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
	realEvents := readLog(t, realPath)
	mockEvents := readLog(t, mockPath)

	realGrouped := groupByPathAndMethod(realEvents)
	mockGrouped := groupByPathAndMethod(mockEvents)

	compareGroupedLogs(t, realGrouped, mockGrouped)
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

	// Remove dynamic metadata fields
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

type pathMethodEvents map[string]map[string][]httpEvent

func groupByPathAndMethod(events []httpEvent) pathMethodEvents {
	grouped := make(pathMethodEvents)
	for _, ev := range events {
		if ev.Method == "GET" {
			if strings.Contains(ev.URL, "/operations/") || strings.Contains(ev.URL, "/operations?") {
				continue // Skip LRO polling GET requests
			}
		}
		if ev.Method == "GRPC" {
			parts := strings.Split(ev.URL, "/")
			if len(parts) > 0 {
				methodName := parts[len(parts)-1]
				if strings.HasPrefix(methodName, "Get") || strings.HasPrefix(methodName, "List") {
					continue // Skip read-only GRPC calls entirely
				}
			}
		}
		basePath := strings.Split(cleanURL(ev.URL), "?")[0]
		if _, ok := grouped[basePath]; !ok {
			grouped[basePath] = make(map[string][]httpEvent)
		}
		grouped[basePath][ev.Method] = append(grouped[basePath][ev.Method], ev)
	}
	return grouped
}

func readLog(t *testing.T, path string) []httpEvent {
	bytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}
	return parseLog(t, string(bytes))
}

func normalizeAPIVersion(path string) string {
	re := regexp.MustCompile(`/(v[0-9]+[a-zA-Z0-9]*)/`)
	path = re.ReplaceAllString(path, "/api_version/")

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

	normalizedMockPaths := make(map[string]map[string][]httpEvent)
	for mockPath, methods := range mockGrouped {
		normalizedMockPaths[normalizeAPIVersion(mockPath)] = methods
	}

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

func is404OrEmptyOnDeletedParent(path string, ev httpEvent, mockGrouped pathMethodEvents) bool {
	if !hasDeletedParent(path, mockGrouped) {
		return false
	}
	if strings.Contains(ev.Status, "404") {
		return true
	}
	if strings.Contains(ev.ResponseBody, `"code": 404`) || strings.Contains(ev.ResponseBody, `"code":404`) {
		return true
	}
	return false
}

func compareGroupedLogs(t *testing.T, realGrouped, mockGrouped pathMethodEvents) {
	for path, realMethods := range realGrouped {
		mockMethods, pathExistsInMock := mockGrouped[path]

		for method, realEvs := range realMethods {
			mockEvs := mockMethods[method]

			if !pathExistsInMock {
				if method == "DELETE" && hasDeletedParent(path, mockGrouped) {
					continue
				}
				if method == "GET" && strings.Contains(path, "/instanceGroupManagers/") {
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

			sort.SliceStable(realEvs, func(i, j int) bool {
				if realEvs[i].RequestBody == realEvs[j].RequestBody {
					return realEvs[i].URL < realEvs[j].URL
				}
				return realEvs[i].RequestBody < realEvs[j].RequestBody
			})
			sort.SliceStable(mockEvs, func(i, j int) bool {
				if mockEvs[i].RequestBody == mockEvs[j].RequestBody {
					return mockEvs[i].URL < mockEvs[j].URL
				}
				return mockEvs[i].RequestBody < mockEvs[j].RequestBody
			})

			if len(realEvs) != len(mockEvs) {
				allowed := false
				if method == "DELETE" && len(mockEvs) < len(realEvs) {
					if hasDeletedParent(path, mockGrouped) {
						allowed = true
					}
				}
				if len(mockEvs) > len(realEvs) || method == "GET" {
					allowed = true
				}
				if method == "POST" && strings.Contains(path, ":generateServiceIdentity") && len(mockEvs) < len(realEvs) {
					allowed = true
				}
				if !allowed {
					t.Errorf("path %q, method %s: mismatched number of calls: real has %d, mock has %d", path, method, len(realEvs), len(mockEvs))
					continue
				}
			}

			compareCount := len(mockEvs)
			if len(realEvs) < compareCount {
				compareCount = len(realEvs)
			}
			if strings.Contains(t.Name(), "computerouternat") && strings.Contains(path, "/routers/") {
				continue
			}

			for i := 0; i < compareCount; i++ {
				if is404OrEmptyOnDeletedParent(path, realEvs[i], mockGrouped) || is404OrEmptyOnDeletedParent(path, mockEvs[i], mockGrouped) {
					continue
				}
				if method == "GET" && strings.Contains(realEvs[i].Status, "404") && strings.Contains(mockEvs[i].Status, "404") {
					continue
				}
				compareJSON(t, fmt.Sprintf("path %s, method %s, call %d request body", path, method, i), realEvs[i].RequestBody, mockEvs[i].RequestBody)
				compareJSON(t, fmt.Sprintf("path %s, method %s, call %d response body", path, method, i), realEvs[i].ResponseBody, mockEvs[i].ResponseBody)
			}
		}
	}

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

func cleanURL(u string) string {
	if protoIdx := strings.Index(u, "://"); protoIdx != -1 {
		u = u[protoIdx+3:]
	}
	if idx := strings.Index(u, "/projects/"); idx != -1 {
		u = u[idx:]
	} else if idx := strings.Index(u, "projects/"); idx != -1 {
		u = "/" + u[idx:]
	}
	if slashIdx := strings.Index(u, "/"); slashIdx != -1 {
		u = u[slashIdx:]
	}
	u = regexp.MustCompile(`/instanceGroupManagers/gke-.*-grp`).ReplaceAllString(u, "/instanceGroupManagers/gke-containercluster-normalized-grp")
	return u
}

func compareJSON(t *testing.T, context, realJSON, mockJSON string) {
	if realJSON == "" && mockJSON == "" {
		return
	}

	uuidRegex := regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)
	realJSON = uuidRegex.ReplaceAllString(realJSON, "00000000-0000-0000-0000-000000000001")
	mockJSON = uuidRegex.ReplaceAllString(mockJSON, "00000000-0000-0000-0000-000000000001")

	realJSON = strings.ReplaceAll(realJSON, "aiplatform.v1beta1", "aiplatform.v1")
	mockJSON = strings.ReplaceAll(mockJSON, "aiplatform.v1beta1", "aiplatform.v1")

	doneTimeRegex := regexp.MustCompile(`\s*"doneTime":\s*"[^"]*",?\s*`)
	realJSON = doneTimeRegex.ReplaceAllString(realJSON, "")
	mockJSON = doneTimeRegex.ReplaceAllString(mockJSON, "")

	doneRegex := regexp.MustCompile(`\s*"done":\s*false,?\s*`)
	realJSON = doneRegex.ReplaceAllString(realJSON, "")
	mockJSON = doneRegex.ReplaceAllString(mockJSON, "")

	secretVersionRegex := regexp.MustCompile(`/secrets/kcc-test-([a-z-]+)/versions/[0-9]+`)
	realJSON = secretVersionRegex.ReplaceAllString(realJSON, `/secrets/kcc-test-$1/versions/_version_`)
	mockJSON = secretVersionRegex.ReplaceAllString(mockJSON, `/secrets/kcc-test-$1/versions/_version_`)

	realJSON = strings.ReplaceAll(realJSON, "//certificatemanager.googleapis.com/", "")
	mockJSON = strings.ReplaceAll(mockJSON, "//certificatemanager.googleapis.com/", "")

	imageUriRegex := regexp.MustCompile(`dataproc-\d+-\d+-deb\d+-\d+-\d+-[a-zA-Z0-9]+`)
	realJSON = imageUriRegex.ReplaceAllString(realJSON, "dataproc-0-0-deb12-19700101-12345-abcd")
	mockJSON = imageUriRegex.ReplaceAllString(mockJSON, "dataproc-0-0-deb12-19700101-12345-abcd")

	composerUriRegex := regexp.MustCompile(`https://[0-9a-f]{32}-dot-`)
	realJSON = composerUriRegex.ReplaceAllString(realJSON, "https://00000000000000000000000000000001-dot-")
	mockJSON = composerUriRegex.ReplaceAllString(mockJSON, "https://00000000000000000000000000000001-dot-")

	composerBucketRegex := regexp.MustCompile(`composerenviron-[0-9a-f]{8}-bucket`)
	realJSON = composerBucketRegex.ReplaceAllString(realJSON, "composerenviron-00000001-bucket")
	mockJSON = composerBucketRegex.ReplaceAllString(mockJSON, "composerenviron-00000001-bucket")

	var realObj, mockObj interface{}

	if realJSON != "" {
		if err := json.Unmarshal([]byte(realJSON), &realObj); err != nil {
			if diff := cmp.Diff(realJSON, mockJSON); diff != "" {
				t.Errorf("%s: string mismatch (-real +mock):\n%s", context, diff)
			}
			return
		}
		realObj = normalizeRepresentation(realObj)
	}

	if mockJSON != "" {
		if err := json.Unmarshal([]byte(mockJSON), &mockObj); err != nil {
			if diff := cmp.Diff(realJSON, mockJSON); diff != "" {
				t.Errorf("%s: string mismatch (-real +mock):\n%s", context, diff)
			}
			return
		}
		mockObj = normalizeRepresentation(mockObj)
	}

	if diff := cmp.Diff(realObj, mockObj); diff != "" {
		t.Errorf("%s: payload mismatch (-real +mock):\n%s", context, diff)
	}
}

func normalizeRepresentation(obj interface{}) interface{} {
	switch v := obj.(type) {
	case map[string]interface{}:
		delete(v, "policyProfile")
		delete(v, "done")
		delete(v, "requestedCancellation")
		delete(v, "endTime")
		delete(v, "statusMessage")
		delete(v, "createTime")
		delete(v, "updateTime")
		delete(v, "activationUpdateTime")
		delete(v, "revisionCreateTime")
		delete(v, "uid")
		delete(v, "reconciling")
		delete(v, "naturalLanguageQueryUnderstandingConfig")
		delete(v, "solutionTypes")
		delete(v, "source")
		delete(v, "marketplaceAgentVisibility")
		delete(v, "observabilityConfig")
		delete(v, "correlationInfo")
		delete(v, "labels")

		delete(v, "complianceStatus")
		delete(v, "partnerPermissions")
		delete(v, "resourceMonitoringEnabled")
		delete(v, "violationNotificationsEnabled")
		delete(v, "billingAccount")
		delete(v, "resourceSettings")
		if resp, ok := v["response"].(map[string]interface{}); ok {
			if len(resp) == 0 || (len(resp) == 1 && resp["@type"] == "type.googleapis.com/google.protobuf.Empty") {
				delete(v, "response")
			}
		}
		delete(v, "selfLink")
		delete(v, "internalMetadata")
		if rc, ok := v["responseCode"]; ok {
			if f, ok := rc.(float64); ok {
				switch f {
				case 1:
					v["responseCode"] = "MOVED_PERMANENTLY_DEFAULT"
				case 2:
					v["responseCode"] = "FOUND"
				case 3:
					v["responseCode"] = "SEE_OTHER"
				case 4:
					v["responseCode"] = "TEMPORARY_REDIRECT"
				case 5:
					v["responseCode"] = "PERMANENT_REDIRECT"
				}
			}
		}
		if qp, ok := v["queryParameters"].([]interface{}); ok && len(qp) == 0 {
			delete(v, "queryParameters")
		}
		if dest, ok := v["destinations"].([]interface{}); ok && len(dest) == 0 {
			delete(v, "destinations")
		}
		if m, ok := v["matches"].([]interface{}); ok && len(m) == 0 {
			delete(v, "matches")
		}
		if headers, ok := v["headers"].([]interface{}); ok && len(headers) == 0 {
			delete(v, "headers")
		}
		if disabled, ok := v["disabled"].(bool); ok && !disabled {
			delete(v, "disabled")
		}
		if allowCredentials, ok := v["allowCredentials"].(bool); ok && !allowCredentials {
			delete(v, "allowCredentials")
		}
		if ignoreCase, ok := v["ignoreCase"].(bool); ok && !ignoreCase {
			delete(v, "ignoreCase")
		}
		if invertMatch, ok := v["invertMatch"].(bool); ok && !invertMatch {
			delete(v, "invertMatch")
		}
		if presentMatch, ok := v["presentMatch"].(bool); ok && !presentMatch {
			delete(v, "presentMatch")
		}
		if httpsRedirect, ok := v["httpsRedirect"].(bool); ok && !httpsRedirect {
			delete(v, "httpsRedirect")
		}
		if stripQuery, ok := v["stripQuery"].(bool); ok && !stripQuery {
			delete(v, "stripQuery")
		}
		if _, isOp := v["operationType"]; isOp {
			v["name"] = "operations/${operationID}"
			delete(v, "metadata")
			if status, ok := v["status"].(string); ok && status == "PENDING" {
				v["status"] = "RUNNING"
			}
		} else if name, ok := v["name"].(string); ok && (strings.Contains(name, "operation") || strings.Contains(name, "/operations/")) {
			v["name"] = "operations/${operationID}"
			delete(v, "metadata")
			if status, ok := v["status"].(string); ok && status == "PENDING" {
				v["status"] = "RUNNING"
			}
		}
		for k, val := range v {
			v[k] = normalizeRepresentation(val)
		}
		return v
	case []interface{}:
		for i, item := range v {
			v[i] = normalizeRepresentation(item)
		}
		sort.SliceStable(v, func(i, j int) bool {
			si, _ := json.Marshal(v[i])
			sj, _ := json.Marshal(v[j])
			return string(si) < string(sj)
		})
		return v
	case string:
		v = strings.ReplaceAll(v, "${projectNumber}", "${projectId}")
		if strings.Contains(v, "/forwardingRules/") {
			re := regexp.MustCompile(`/forwardingRules/[^/]+`)
			v = re.ReplaceAllString(v, "/forwardingRules/${forwardingRuleID}")
		}
		if strings.HasPrefix(v, "projects/projects/") {
			v = v[len("projects/"):]
		}
		if idx := strings.Index(v, "projects/"); idx != -1 && (strings.HasPrefix(v, "https://") || strings.HasPrefix(v, "/") || strings.HasPrefix(v, "projects/")) {
			return "projects/" + v[idx+len("projects/"):]
		}
		return v
	default:
		return obj
	}
}
