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
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

var mockGCPSkipFixtures = map[string]bool{
	"devicestreaming/v1alpha1/devicestreamingsession/devicestreamingsession-maximal": true,
	"devicestreaming/v1alpha1/devicestreamingsession/devicestreamingsession-minimal": true,
	// TODO(https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/12388): Align outdated ComposerEnvironment mock logs with real GCP
	"composer/v1beta1/composerenvironment/composerenvironmentwithkms":    true,
	"composer/v1beta1/composerenvironment/composerenvironmentwithrefs":   true,
	"composer/v1beta1/composerenvironment/composerenvironmentnodeconfig": true,
}

var realGCPSkipFixtures = map[string]bool{
	// Resource Manager Tags are org level thus requiring an owned test org.
	"container/v1beta1/containercluster/containercluster-resourcemanagertags-autopilot": true,
	"container/v1beta1/containercluster/containercluster-resourcemanagertags-standard":  true,
	"container/v1beta1/containernodepool/containernodepool-resourcemanagertags":         true,
	// SecurityCenter MuteConfig requires organization-level permissions and quota project.
	"securitycenter/v1alpha1/securitycentermuteconfig/securitycentermuteconfig-dynamic": true,
	"securitycenter/v1alpha1/securitycentermuteconfig/securitycentermuteconfig-maximal": true,
	"securitycenter/v1alpha1/securitycentermuteconfig/securitycentermuteconfig-minimal": true,
	// GKEBackup BackupChannel requires distinct source and destination projects.
	"gkebackup/v1alpha1/gkebackupbackupchannel/gkebackupbackupchannel-maximal": true,
	"gkebackup/v1alpha1/gkebackupbackupchannel/gkebackupbackupchannel-minimal": true,
	// Tags acquire and project tag key tests have environment-dependent real GCP payloads
	"tags/v1beta1/tagstagkey/tagkeyacquire":        true,
	"tags/v1beta1/tagstagvalue/tagvalueacquire":    true,
	"tags/v1beta1/tagstagkey/tagkeyprojectautogen": true,

	// Skipped for patch release: golden log alignment issues with IAM dependency traffic are fixed
	// on master (in commits between the cherry-picked PR and release base). Skipped here to keep
	// the patch isolated without backporting test infra diffs.
	"composer/v1beta1/composerenvironment/composerenvironmentbasic":           true,
	"composer/v1beta1/composerenvironment/composerenvironmentmultipleupdates": true,
}

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
			createPath := filepath.Join(path, "create.yaml")
			if fileExists(createPath) {
				relPath, _ := filepath.Rel(absRootDir, path)
				if realGCPSkipFixtures[relPath] || mockGCPSkipFixtures[relPath] {
					return nil
				}

				realLogPath := filepath.Join(path, "_http.log")
				mockLogPath := filepath.Join(path, "_http_mock.log")

				if fileExists(realLogPath) && fileExists(mockLogPath) {
					t.Run(relPath, func(t *testing.T) {
						primaryKind, err := getPrimaryKind(createPath)
						if err != nil {
							t.Fatalf("failed to get primary kind for %s: %v", path, err)
						}
						dependenciesPath := filepath.Join(path, "dependencies.yaml")
						depKinds, err := getDependencyKinds(dependenciesPath)
						if err != nil {
							t.Fatalf("failed to get dependency kinds for %s: %v", path, err)
						}
						compareLogs(t, realLogPath, mockLogPath, depKinds, primaryKind)
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

func TestRealHTTPLogsDoNotContainMockGCP(t *testing.T) {
	rootDir := "testdata/basic"
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		t.Fatalf("failed to get absolute path for %s: %v", rootDir, err)
	}

	err = filepath.WalkDir(absRootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && d.Name() == "_http.log" {
			dirPath := filepath.Dir(path)
			relPath, _ := filepath.Rel(absRootDir, dirPath)
			if realGCPSkipFixtures[relPath] || mockGCPSkipFixtures[relPath] {
				return nil
			}

			data, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading %s: %w", path, err)
			}
			if strings.Contains(string(data), "(mockgcp)") {
				t.Errorf("real GCP log %s contains '(mockgcp)'! Never copy _http_mock.log to _http.log", path)
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

func compareLogs(t *testing.T, realPath, mockPath string, depKinds map[string]string, primaryKind string) {
	realEvents := readLog(t, realPath)
	mockEvents := readLog(t, mockPath)

	if len(depKinds) > 0 {
		realEvents = filterDependencyEvents(realEvents, depKinds, primaryKind)
		mockEvents = filterDependencyEvents(mockEvents, depKinds, primaryKind)
	}

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

			// Sort events by their RequestBody to ensure deterministic order for concurrent sibling operations
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
					allowed = true // Allow extra retries/reconciliations across GET and mock calls
				}
				// Allow generateServiceIdentity to have fewer calls in mock because the direct controller
				// optimizes and avoids duplicate POST calls.
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
				continue // Subresource Router NAT updates modify the parent Cloud Router array via iterative PATCH loops with differing intermediate call ordering between real and mock
			}

			for i := 0; i < compareCount; i++ {
				if is404OrEmptyOnDeletedParent(path, realEvs[i], mockGrouped) || is404OrEmptyOnDeletedParent(path, mockEvs[i], mockGrouped) {
					continue
				}
				if method == "GET" && strings.Contains(realEvs[i].Status, "404") && strings.Contains(mockEvs[i].Status, "404") {
					continue // Both real and mock confirm resource does not exist right before create / after delete
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

	// Normalize any UUIDs to dummy UUID to align real and mock logs
	uuidRegex := regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)
	realJSON = uuidRegex.ReplaceAllString(realJSON, "00000000-0000-0000-0000-000000000001")
	mockJSON = uuidRegex.ReplaceAllString(mockJSON, "00000000-0000-0000-0000-000000000001")

	// Normalize aiplatform v1beta1 to v1 to align real and mock logs
	realJSON = strings.ReplaceAll(realJSON, "aiplatform.v1beta1", "aiplatform.v1")
	mockJSON = strings.ReplaceAll(mockJSON, "aiplatform.v1beta1", "aiplatform.v1")

	// Remove doneTime to align LROs (mock LROs are done immediately, real are active)
	doneTimeRegex := regexp.MustCompile(`\s*"doneTime":\s*"[^"]*",?\s*`)
	realJSON = doneTimeRegex.ReplaceAllString(realJSON, "")
	mockJSON = doneTimeRegex.ReplaceAllString(mockJSON, "")

	// Mock server does not return "done: false" in metadata
	doneRegex := regexp.MustCompile(`\s*"done":\s*false,?\s*`)
	realJSON = doneRegex.ReplaceAllString(realJSON, "")
	mockJSON = doneRegex.ReplaceAllString(mockJSON, "")

	secretVersionRegex := regexp.MustCompile(`/secrets/kcc-test-([a-z-]+)/versions/[0-9]+`)
	realJSON = secretVersionRegex.ReplaceAllString(realJSON, `/secrets/kcc-test-$1/versions/_version_`)
	mockJSON = secretVersionRegex.ReplaceAllString(mockJSON, `/secrets/kcc-test-$1/versions/_version_`)

	// Normalize certificate manager prefix
	realJSON = strings.ReplaceAll(realJSON, "//certificatemanager.googleapis.com/", "")
	mockJSON = strings.ReplaceAll(mockJSON, "//certificatemanager.googleapis.com/", "")

	// Normalize unhyphenated Composer Airflow URI tokens
	composerUriRegex := regexp.MustCompile(`https://[0-9a-f]{32}-dot-`)
	realJSON = composerUriRegex.ReplaceAllString(realJSON, "https://00000000000000000000000000000001-dot-")
	mockJSON = composerUriRegex.ReplaceAllString(mockJSON, "https://00000000000000000000000000000001-dot-")

	// Normalize Composer auto-generated bucket hashes
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
		delete(v, "done")
		delete(v, "requestedCancellation")
		delete(v, "endTime")
		delete(v, "statusMessage")
		delete(v, "createTime")
		delete(v, "updateTime")
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
			if opType, ok := v["operationType"].(string); ok && opType == "UPGRADE_NODES" {
				delete(v, "operationType")
			}
		} else if name, ok := v["name"].(string); ok && (strings.Contains(name, "operation") || strings.Contains(name, "/operations/")) {
			v["name"] = "operations/${operationID}"
			delete(v, "metadata")
			if status, ok := v["status"].(string); ok && status == "PENDING" {
				v["status"] = "RUNNING"
			}
		}
		if kind, ok := v["kind"].(string); ok && kind == "compute#backendService" {
			delete(v, "port")
			delete(v, "portName")
			delete(v, "protocol")
		}
		if kind, ok := v["kind"].(string); ok && kind == "compute#instanceGroupManager" {
			return map[string]interface{}{"kind": kind}
		}
		if kind, ok := v["kind"].(string); ok && kind == "compute#network" {
			delete(v, "subnetworks")
			delete(v, "peerings")
			delete(v, "routingConfig")
		}
		if kind, ok := v["kind"].(string); ok && kind == "compute#subnetwork" {
			delete(v, "allowSubnetCidrRoutesOverlap")
			delete(v, "enableFlowLogs")
			delete(v, "stackType")
		}
		if kind, ok := v["kind"].(string); ok && kind == "storage#objects" {
			delete(v, "prefixes")
		}
		if _, isCluster := v["monitoringService"]; isCluster {
			delete(v, "currentMasterVersion")
			delete(v, "currentNodeVersion")
			delete(v, "initialClusterVersion")
			delete(v, "releaseChannel")
			delete(v, "currentNodeCount")
			delete(v, "nodeCreationConfig")
			delete(v, "controlPlaneEgress")
			delete(v, "master")
			delete(v, "privateCluster")
			delete(v, "anonymousAuthenticationConfig")
			delete(v, "ipAllocationPolicy")
			delete(v, "masterAuth")
			delete(v, "controlPlaneEndpointsConfig")
			delete(v, "addonsConfig")
			delete(v, "zone")
		}
		if cluster, ok := v["cluster"].(map[string]interface{}); ok {
			delete(cluster, "initialClusterVersion")
		}
		if config, ok := v["config"].(map[string]interface{}); ok {
			if containerdConfig, ok := config["containerdConfig"].(map[string]interface{}); ok {
				delete(containerdConfig, "registryHosts")
				delete(containerdConfig, "writableCgroups")
				delete(containerdConfig, "privateRegistryAccessConfig")
			}
		}
		if containerdConfig, ok := v["containerdConfig"].(map[string]interface{}); ok {
			delete(containerdConfig, "registryHosts")
			delete(containerdConfig, "writableCgroups")
			delete(containerdConfig, "privateRegistryAccessConfig")
		}
		if kubelet, ok := v["kubeletConfig"].(map[string]interface{}); ok {
			delete(kubelet, "maxParallelImagePulls")
		}
		if _, isDnsAuth := v["dnsResourceRecord"]; isDnsAuth {
			if t, ok := v["type"].(float64); ok && t == 1 {
				v["type"] = "FIXED_RECORD"
			}
			if rec, ok := v["dnsResourceRecord"].(map[string]interface{}); ok {
				if data, ok := rec["data"].(string); ok && (data == "authorize.certificatemanager.goog." || data == "dns-resource-record-data-placeholder") {
					rec["data"] = "_NORMALIZED_DNS_DATA_"
				}
				if name, ok := rec["name"].(string); ok {
					rec["name"] = regexp.MustCompile(`_acme-challenge\.[^.]+\.`).ReplaceAllString(name, "_acme-challenge._NORMALIZED_DOMAIN_.")
				}
			}
		}
		if _, hasNodePools := v["nodePools"]; hasNodePools {
			delete(v, "nodePools")
			delete(v, "nodeConfig")
			delete(v, "networkConfig")
		}
		if _, isNodePool := v["initialNodeCount"]; isNodePool {
			delete(v, "instanceGroupUrls")
			delete(v, "version")
			delete(v, "networkConfig")
			delete(v, "etag")
			delete(v, "locations")
			delete(v, "kubeletCertInfo")
			if sl, ok := v["selfLink"].(string); ok {
				v["selfLink"] = strings.ReplaceAll(sl, "/zones/", "/locations/")
			}
			if cfg, ok := v["config"].(map[string]interface{}); ok {
				delete(cfg, "nodeImageConfig")
			}
		}
		if auto, ok := v["autoCreateSubnetworks"].(bool); ok && auto {
			if _, hasSubnets := v["subnetworks"]; hasSubnets {
				v["subnetworks"] = []interface{}{"https://www.googleapis.com/compute/v1/projects/_project_/regions/_all_/subnetworks/_auto_"}
			}
		}
		if ula, ok := v["enableUlaInternalIpv6"].(bool); ok && !ula {
			delete(v, "enableUlaInternalIpv6")
		}
		if enc, ok := v["encryptedInterconnectRouter"].(bool); ok && !enc {
			delete(v, "encryptedInterconnectRouter")
		}
		if timeout, ok := v["effectiveTcpTimeWaitTimeoutSec"].(float64); ok && timeout == 120 {
			delete(v, "effectiveTcpTimeWaitTimeoutSec")
		}
		if desc, ok := v["description"].(string); ok && desc == "" {
			delete(v, "description")
		}
		if preview, ok := v["preview"].(bool); ok && !preview {
			delete(v, "preview")
		}
		if dyn, ok := v["enableDynamicPortAllocation"].(bool); ok && !dyn {
			delete(v, "enableDynamicPortAllocation")
		}
		if state, ok := v["state"].(string); ok && state == "READY" && v["kind"] == "compute#subnetwork" {
			delete(v, "state")
		}
		if slice, ok := v["drainNatIps"].([]interface{}); ok && len(slice) == 0 {
			delete(v, "drainNatIps")
		}
		if slice, ok := v["natIps"].([]interface{}); ok && len(slice) == 0 {
			delete(v, "natIps")
		}
		if slice, ok := v["nats"].([]interface{}); ok && len(slice) == 0 {
			delete(v, "nats")
		}
		if v["logConfig"] == nil {
			delete(v, "logConfig")
		}
		if val, ok := v["icmpIdleTimeoutSec"].(float64); ok && val == 30 {
			delete(v, "icmpIdleTimeoutSec")
		}
		if val, ok := v["udpIdleTimeoutSec"].(float64); ok && val == 30 {
			delete(v, "udpIdleTimeoutSec")
		}
		if val, ok := v["tcpTransitoryIdleTimeoutSec"].(float64); ok && val == 30 {
			delete(v, "tcpTransitoryIdleTimeoutSec")
		}
		if val, ok := v["tcpEstablishedIdleTimeoutSec"].(float64); ok && val == 1200 {
			delete(v, "tcpEstablishedIdleTimeoutSec")
		}
		if val, ok := v["tcpTimeWaitTimeoutSec"].(float64); ok && (val == 120 || val == 0) {
			delete(v, "tcpTimeWaitTimeoutSec")
		}
		if tier, ok := v["autoNetworkTier"].(string); ok && tier == "PREMIUM" {
			delete(v, "autoNetworkTier")
		}
		if rangeStr, ok := v["internalIpv6Range"].(string); ok && strings.HasPrefix(rangeStr, "fd") {
			v["internalIpv6Range"] = "fd00:0000:0000:0:0:0:0:0/48"
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

func getPrimaryKind(path string) (string, error) {
	gvk, err := getGVKFromYAML(path)
	if err != nil {
		return "", err
	}
	return gvk.Kind, nil
}

func getDependencyKinds(path string) (map[string]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	depKinds := make(map[string]string)
	// Split by "---" to support multi-document YAML files
	docs := strings.Split(string(bytes), "\n---")
	for _, doc := range docs {
		doc = strings.TrimSpace(doc)
		if doc == "" {
			continue
		}
		var u unstructured.Unstructured
		if err := yaml.Unmarshal([]byte(doc), &u); err != nil {
			// Some files might have comments or other non-YAML content, or split issues
			continue
		}
		kind := u.GetKind()
		if kind == "Project" || kind == "Folder" || kind == "Organization" {
			continue
		}
		name := u.GetName()
		if len(name) >= 3 {
			depKinds[name] = kind
		}
		if resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID"); len(resourceID) >= 3 {
			depKinds[resourceID] = kind
		}
		for _, ph := range getPlaceholdersForKind(kind) {
			depKinds[ph] = kind
		}
	}
	return depKinds, nil
}

func getPlaceholdersForKind(kind string) []string {
	switch kind {
	case "ComputeNetwork":
		return []string{"${networkID}"}
	case "ComputeSubnetwork":
		return []string{"${subnetworkID}"}
	case "ComputeAddress":
		return []string{"${addressID}"}
	case "ComputeForwardingRule":
		return []string{"${forwardingRuleID}"}
	case "ComputeFirewall":
		return []string{"${firewallID}"}
	case "ComputeRouter":
		return []string{"${routerID}"}
	case "ComputeRoute":
		return []string{"${routeID}"}
	case "ComputeDisk":
		return []string{"${diskID}"}
	case "ComputeInstance":
		return []string{"${instanceID}"}
	case "KMSKeyRing":
		return []string{"${kmsKeyRingID}", "${keyRingID}"}
	case "KMSCryptoKey":
		return []string{"${kmsCryptoKeyID}", "${cryptoKeyID}"}
	case "IAMServiceAccount":
		return []string{"${serviceAccountID}", "${serviceAccountEmail}"}
	case "CertificateManagerCertificateMap":
		return []string{"${certificateMapID}"}
	case "CertificateManagerCertificate":
		return []string{"${certificateID}"}
	case "CertificateManagerDNSAuthorization":
		return []string{"${dnsAuthorizationID}"}
	}
	return nil
}

func filterDependencyEvents(events []httpEvent, depKinds map[string]string, primaryKind string) []httpEvent {
	var filtered []httpEvent
	for _, ev := range events {
		if !isDependencyEvent(ev, depKinds, primaryKind) {
			filtered = append(filtered, ev)
		}
	}
	return filtered
}

func isDependencyEvent(ev httpEvent, depKinds map[string]string, primaryKind string) bool {
	// PATCH requests are never part of dependency creation/setup, so they are always kept
	if ev.Method == "PATCH" {
		return false
	}

	//  tested resource is a virtual resource (i.e. primaryKind == "RedisClusterEndpoint"), any HTTP traffic
	//  containing /clusters/ (the parent RedisCluster) is not treated as dependency traffic and is kept in the comparison.
	if primaryKind == "RedisClusterEndpoint" {
		if strings.Contains(ev.URL, "/clusters/") {
			return false
		}
	}

	isIAM := primaryKind == "IAMPolicy" || primaryKind == "IAMPolicyMember" || primaryKind == "IAMPartialPolicy" || primaryKind == "IAMAuditConfig"

	for depName, kind := range depKinds {
		// If a dependency resource has the same kind as the primary resource under test,
		// we keep all of its events.
		if kind == primaryKind {
			continue
		}

		// Clean the URL to get the path
		urlPath := strings.Split(cleanURL(ev.URL), "?")[0]

		// If the path doesn't contain the dependency name, check if it's a POST to create it
		if !strings.Contains(urlPath, depName) {
			if ev.Method == "POST" && (strings.Contains(ev.RequestBody, depName) || strings.Contains(ev.URL, depName)) {
				return true
			}
			continue
		}

		// If the path contains depName, check if it is a valid segment (boundary check)
		idx := strings.Index(urlPath, depName)
		if idx == -1 {
			continue
		}

		if idx > 0 {
			before := urlPath[idx-1]
			if before != '/' && before != ':' && before != '=' {
				continue
			}
		}

		suffix := urlPath[idx+len(depName):]
		// If suffix is empty or /, it is the dependency resource itself
		if suffix == "" || suffix == "/" {
			return true
		}

		// If suffix starts with /operations, /locations, or /regions, it's an LRO or location metadata on the dependency itself
		if strings.HasPrefix(suffix, "/operations") || strings.HasPrefix(suffix, "/locations") || strings.HasPrefix(suffix, "/regions") {
			return true
		}

		// If suffix has more than one path segment (i.e. contains a '/' after the first segment),
		// it is a child resource of the dependency, so it belongs to the primary resource under test.
		trimmedSuffix := strings.TrimPrefix(suffix, "/")
		if strings.Contains(trimmedSuffix, "/") {
			continue
		}

		// If the primary resource is NOT an IAM resource, and suffix is an IAM path segment, it is an IAM policy request on the dependency itself
		if !isIAM {
			if strings.HasPrefix(suffix, "/iam") ||
				strings.HasPrefix(suffix, ":setIamPolicy") ||
				strings.HasPrefix(suffix, ":getIamPolicy") ||
				strings.HasPrefix(suffix, ":testIamPermissions") {
				return true
			}
		}

		// Any other single-segment suffix (e.g. /setLabels) is a custom subresource/action on the dependency itself, so filter it out
		return true
	}
	return false
}
