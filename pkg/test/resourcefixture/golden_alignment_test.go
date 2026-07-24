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

var mockGCPSkipGroupKinds = map[schema.GroupKind]bool{
	schema.GroupKind{
		Group: "devicestreaming.cnrm.cloud.google.com",
		Kind:  "DeviceStreamingSession",
	}: true,
}

var realGCPSkipGroupKinds = map[schema.GroupKind]bool{
	schema.GroupKind{
		Group: "securitycenter.cnrm.cloud.google.com",
		Kind:  "SecurityCenterMuteConfig",
	}: true,
	schema.GroupKind{
		Group: "gkebackup.cnrm.cloud.google.com",
		Kind:  "GKEBackupBackupChannel",
	}: true,
	schema.GroupKind{
		Group: "edgecontainer.cnrm.cloud.google.com",
		Kind:  "EdgeContainerCluster",
	}: true,
}

var realGCPSkipFixtures = map[string]bool{
	"container/v1beta1/containercluster/containercluster-resourcemanagertags-autopilot": true,
	"container/v1beta1/containercluster/containercluster-resourcemanagertags-standard":  true,
	"container/v1beta1/containernodepool/containernodepool-resourcemanagertags":         true,
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
			relPath, _ := filepath.Rel(absRootDir, path)
			if realGCPSkipFixtures[relPath] {
				return nil
			}

			realLogPath := filepath.Join(path, "_http.log")
			mockLogPath := filepath.Join(path, "_http_mock.log")

			if fileExists(realLogPath) {
				createPath := filepath.Join(path, "create.yaml")
				if fileExists(createPath) {
					gvk, err := getGVKFromYAML(createPath)
					if err == nil {
						gk := gvk.GroupKind()
						if mockGCPSkipGroupKinds[gk] || realGCPSkipGroupKinds[gk] {
							return nil
						}
						if !mockGCPSkipGroupKinds[gk] && !fileExists(mockLogPath) {
							t.Errorf("fixture %q: resource must have _http_mock.log golden file", path)
						}
					}
				}

				if fileExists(mockLogPath) {
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
			if realGCPSkipFixtures[relPath] {
				return nil
			}

			createPath := filepath.Join(dirPath, "create.yaml")
			if fileExists(createPath) {
				gvk, err := getGVKFromYAML(createPath)
				if err == nil && realGCPSkipGroupKinds[gvk.GroupKind()] {
					return nil
				}
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

	// Remove doneTime to align LROs (mock LROs are done immediately, real are active)
	doneTimeRegex := regexp.MustCompile(`\s*"doneTime":\s*"[^"]*",?\s*`)
	realJSON = doneTimeRegex.ReplaceAllString(realJSON, "")
	mockJSON = doneTimeRegex.ReplaceAllString(mockJSON, "")

	// Mock server does not return "done: false" in metadata
	doneRegex := regexp.MustCompile(`\s*"done":\s*false,?\s*`)
	realJSON = doneRegex.ReplaceAllString(realJSON, "")
	mockJSON = doneRegex.ReplaceAllString(mockJSON, "")

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
		if name, ok := v["name"].(string); ok && strings.Contains(name, "/operations/") {
			v["name"] = "operations/${operationID}"
			delete(v, "metadata")
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
		if kind, ok := v["kind"].(string); ok && kind == "storage#objects" {
			delete(v, "prefixes")
		}
		if _, isCluster := v["monitoringService"]; isCluster {
			delete(v, "currentMasterVersion")
			delete(v, "currentNodeVersion")
			delete(v, "initialClusterVersion")
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
