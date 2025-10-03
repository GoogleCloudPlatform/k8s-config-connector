// Copyright 2025 Google LLC
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
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
)

// CheckTransportMetrics verifies that the transport layer is emitting the expected metrics
// by scraping the Prometheus metrics endpoint and checking for the specific metrics we expect.
// This function can be called from scenario tests or used independently.
func CheckTransportMetrics(t *testing.T, h *create.Harness) {
	// Wait a bit for any existing metrics to be collected
	time.Sleep(5 * time.Second)

	// Scrape the metrics endpoint to verify the transport metrics are being emitted
	metricsEndpoint := "http://localhost:8888/metrics"
	t.Logf("scraping metrics endpoint: %s", metricsEndpoint)

	resp, err := http.Get(metricsEndpoint)
	if err != nil {
		t.Fatalf("failed to scrape metrics endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("metrics endpoint returned status %d, expected 200", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read metrics response: %v", err)
	}

	metricsBody := string(body)
	t.Logf("metrics endpoint returned %d bytes", len(metricsBody))

	// Check for the specific transport metrics we expect
	expectedMetrics := []string{
		"configconnector_gcp_api_requests_total",
		//"configconnector_gcp_api_request_duration_seconds",
		//"configconnector_gcp_api_errors_total",
	}

	missingMetrics := []string{}
	for _, metric := range expectedMetrics {
		if !strings.Contains(metricsBody, metric) {
			missingMetrics = append(missingMetrics, metric)
		}
	}

	if len(missingMetrics) > 0 {
		t.Errorf("missing expected transport metrics: %v", missingMetrics)
	} else {
		t.Logf("all expected transport metrics found")
	}

	// Check that we have metrics with proper labels
	// Look for metrics with method, status_code, and controller_name labels
	labelPatterns := []string{
		"configconnector_gcp_api_requests_total{",
		"configconnector_gcp_api_request_duration_seconds{",
	}

	missingLabels := []string{}
	for _, pattern := range labelPatterns {
		if !strings.Contains(metricsBody, pattern) {
			missingLabels = append(missingLabels, pattern)
		}
	}

	if len(missingLabels) > 0 {
		t.Errorf("missing metrics with expected label structure: %v", missingLabels)
	} else {
		t.Logf("metrics with proper label structure found")
	}

	// Check for specific label values we expect
	expectedLabelValues := []string{
		"method=\"",
		"status_code=\"",
		"controller_name=\"",
	}

	missingLabelValues := []string{}
	for _, labelValue := range expectedLabelValues {
		if !strings.Contains(metricsBody, labelValue) {
			missingLabelValues = append(missingLabelValues, labelValue)
		}
	}

	if len(missingLabelValues) > 0 {
		t.Errorf("missing expected label values: %v", missingLabelValues)
	} else {
		t.Logf("expected label values found")
	}

	// Check for specific controller names we expect from our test resources
	expectedControllerNames := []string{
		"project-controller",               // For Project resources
		"logginglogmetric-controller",      // For LoggingLogMetric resources
		"computenetwork-controller",        // For ComputeNetwork resources
		"monitoringalertpolicy-controller", // For MonitoringAlertPolicy resources
	}

	missingControllerNames := []string{}
	for _, controllerName := range expectedControllerNames {
		// Look for metrics with this specific controller name
		controllerPattern := fmt.Sprintf("controller_name=\"%s\"", controllerName)
		if !strings.Contains(metricsBody, controllerPattern) {
			missingControllerNames = append(missingControllerNames, controllerName)
		}
	}

	if len(missingControllerNames) > 0 {
		t.Errorf("missing metrics with expected controller names: %v", missingControllerNames)
	} else {
		t.Logf("all expected controller names found in metrics")
	}

	// Log the controller names we found for debugging
	t.Logf("checking for controller names in metrics...")
	for _, controllerName := range expectedControllerNames {
		controllerPattern := fmt.Sprintf("controller_name=\"%s\"", controllerName)
		if strings.Contains(metricsBody, controllerPattern) {
			t.Logf("  ✓ Found controller: %s", controllerName)
		} else {
			t.Logf("  ✗ Missing controller: %s", controllerName)
		}
	}

	t.Logf("transport metrics verification completed:")
	t.Logf("  - metrics endpoint scraped successfully")
	t.Logf("  - transport metrics verified on Prometheus endpoint")
	t.Logf("  - controller names verified for all resource types")
}
