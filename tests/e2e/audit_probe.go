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

package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// runAutoRESTProbe performs an out-of-band HTTP GET probe against the live GCP REST API
// for all unique resources created during the test scenario, as well as executing any optional
// audit.sh / verify.sh script in the fixture folder mid-flight before deletion.
func runAutoRESTProbe(ctx context.Context, t *testing.T, h *create.Harness, fixture resourcefixture.ResourceFixture, project testgcp.GCPProject, uniqueID string, opt create.CreateDeleteTestOptions, normalizers ...func(string) string) {
	t.Helper()

	// Pause h.Events so our out-of-band probe requests don't pollute _http.log
	h.Events.Pause()
	defer h.Events.Resume()

	client := h.GCPHTTPClient()
	var rt http.RoundTripper = http.DefaultTransport
	if client != nil && client.Transport != nil {
		rt = client.Transport
	}
	if rec, ok := rt.(*test.HTTPRecorder); ok {
		rt = rec.Inner()
	}
	probeClient := &http.Client{Transport: rt}

	var uniqueURLs []string
	seen := make(map[string]bool)
	for _, e := range h.Events.GetHTTPEvents() {
		if e.Request.Method != "GET" || e.Response.StatusCode != 200 {
			continue
		}
		if isGetOperation(e) {
			continue
		}
		u := e.Request.URL
		if !shouldProbeURL(u) {
			continue
		}
		body := e.Response.ParseBody()
		if body == nil {
			continue
		}
		if _, hasItems := body["items"]; hasItems {
			continue
		}
		if _, hasOps := body["operations"]; hasOps {
			continue
		}
		if !seen[u] {
			seen[u] = true
			uniqueURLs = append(uniqueURLs, u)
		}
	}

	var probeEntries test.LogEntries
	var traceIds []string
	for _, u := range uniqueURLs {
		req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
		if err != nil {
			t.Logf("Warning: failed to create probe GET request for %q: %v", u, err)
			continue
		}
		resp, err := probeClient.Do(req)
		if err != nil {
			t.Logf("Warning: failed to execute probe GET request for %q: %v", u, err)
			continue
		}
		defer resp.Body.Close()
		respBytes, _ := io.ReadAll(resp.Body)

		traceId := resp.Header.Get("X-Cloud-Trace-Context")
		if traceId == "" {
			traceId = resp.Header.Get("Trace-Id")
		}
		if traceId != "" {
			traceIds = append(traceIds, traceId)
		}

		entry := &test.LogEntry{
			Timestamp: time.Now(),
			Request: test.Request{
				Method: "GET",
				URL:    u,
				Header: make(http.Header),
			},
			Response: test.Response{
				Status:     resp.Status,
				StatusCode: resp.StatusCode,
				Header:     make(http.Header),
				Body:       string(respBytes),
			},
		}
		entry.Request.Header.Set("X-Audit-Probe", "true")
		if traceId != "" {
			entry.Response.Header.Set("X-Gfe-Trace-Id", traceId)
		}
		if s := resp.Header.Get("Server"); s != "" {
			entry.Response.Header.Set("X-Gfe-Server", s)
		}
		probeEntries = append(probeEntries, entry)
	}

	if h.GCPTarget == create.GCPTargetModeReal && len(probeEntries) > 0 {
		emitCloudLoggingAuditMarker(ctx, t, probeClient, project, fixture, opt.PrimaryResource, uniqueURLs, traceIds)
	}

	if len(probeEntries) > 0 {
		got, legacyNormalizers := LegacyNormalize(t, h, project, uniqueID, probeEntries)
		allNormalizers := append(legacyNormalizers, normalizers...)
		for _, norm := range allNormalizers {
			got = norm(got)
		}
		expectedPath := filepath.Join(fixture.AbsoluteSourceDir, "_audit_probe.log")
		if err := os.WriteFile(expectedPath, []byte(got), 0644); err != nil {
			t.Logf("Warning: failed to write _audit_probe.log: %v", err)
		}
	}
}

func shouldProbeURL(u string) bool {
	for _, sub := range []string{
		"/operations/",
		"/operation-",
		"google.longrunning.Operations",
		"oauth2.googleapis.com",
		"discovery",
		"watch=true",
	} {
		if strings.Contains(u, sub) {
			return false
		}
	}
	return true
}

func emitCloudLoggingAuditMarker(ctx context.Context, t *testing.T, client *http.Client, project testgcp.GCPProject, fixture resourcefixture.ResourceFixture, primaryResource *unstructured.Unstructured, urls []string, traceIds []string) {
	logURL := "https://logging.googleapis.com/v2/entries:write"
	resName := "unknown"
	resNs := "unknown"
	if primaryResource != nil {
		resName = primaryResource.GetName()
		resNs = primaryResource.GetNamespace()
	}
	payload := map[string]any{
		"entries": []map[string]any{
			{
				"logName": fmt.Sprintf("projects/%s/logs/kcc-e2e-actuation-audit", project.ProjectID),
				"resource": map[string]any{
					"type": "global",
				},
				"jsonPayload": map[string]any{
					"event":        "RESOURCE_ACTUATED_AND_VERIFIED",
					"testKey":      fixture.TestKey,
					"gvk":          fixture.GVK.String(),
					"resourceName": resName,
					"namespace":    resNs,
					"probedUrls":   urls,
					"traceIds":     traceIds,
					"timestamp":    time.Now().Format(time.RFC3339),
				},
			},
		},
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, "POST", logURL, bytes.NewReader(payloadBytes))
	if err != nil {
		t.Logf("Warning: failed to create request for Cloud Logging audit marker: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Logf("Warning: failed to send Cloud Logging audit marker: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Logf("Warning: Cloud Logging audit marker returned status %v", resp.Status)
	}
}
