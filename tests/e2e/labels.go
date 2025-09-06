// Copyright 2024 Google LLC
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
	"encoding/json"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var TestingInvalidLabels = map[string]string{
	// K8s recommended label names.
	"app.kubernetes.io/name":       "mock-app",
	"app.kubernetes.io/instance":   "mock-instance",
	"app.kubernetes.io/version":    "v1.0.0",
	"app.kubernetes.io/component":  "mock-component",
	"app.kubernetes.io/part-of":    "mock-part-of",
	"app.kubernetes.io/managed-by": "configmanagement.gke.io",

	// K8s common tool-applied labels
	"applyset.kubernetes.io/id":              "mock-applyset-id",
	"configmanagement.gke.io/sync-name":      "mock-sync-name",
	"configmanagement.gke.io/sync-namespace": "mock-sync-namespace",

	// Valid GCP labels.
	"managed-by-cnrm-test": "true",
}

// TestingLabelName is a valid GCP label name used for testing purposes. It should be reserved and shown in the http golden log.
const TestingLabelName = "managed-by-cnrm-test"

// AddInvalidGCPLabels adds additional K8s common labels to the test object. Those labels are not valid GCP labels, but are used to test the label validation logic in the controller.
// If the controller uses K8s labels as GCP labels, it should remove the invalid ones so these testing label should not fail the resource reconciliation.
func AddTestingGCPLabels(t *testing.T, u *unstructured.Unstructured) {
	labelInSpec := true

	totalLabels, found, err := unstructured.NestedStringMap(u.Object, "spec", "labels")
	if !found || err != nil {
		labelInSpec = false
		totalLabels = u.GetLabels()

	}

	// Inject the testing label
	if totalLabels[TestingLabelName] != "" {
		t.Fatalf("'managed-by-cnrm-test' is reserved label name, it should not be set. expect nil, got %q", totalLabels["managed-by-cnrm-test"])
	}
	totalLabels[TestingLabelName] = "true"

	// Add the testing invalid labels for metadata validation.
	if !labelInSpec {
		for k8sLabel, k8sLabelValue := range TestingInvalidLabels {
			if _, exists := totalLabels[k8sLabel]; !exists {
				totalLabels[k8sLabel] = k8sLabelValue
			} else {
				t.Logf("Label %q already exists with value %q, skipping addition", k8sLabel, totalLabels[k8sLabel])
			}
		}
		u.SetLabels(totalLabels)
	} else {
		err := unstructured.SetNestedStringMap(u.Object, totalLabels, "spec", "labels")
		if err != nil {
			t.Fatalf("failed to set labels in spec: %v", err)
		}
	}
}

func VerifyTestingLabels(t *testing.T, u *unstructured.Unstructured) {
	labelInSpec := true

	existing, found, err := unstructured.NestedStringMap(u.Object, "spec", "labels")
	if !found || err != nil {
		labelInSpec = false
		existing = u.GetLabels()
	}

	// clean up the testing label
	if _, ok := existing[TestingLabelName]; !ok {
		t.Fatalf("unexpected label wipe out. controller should not wipe out labels, but compute the labels based on the existing ones. Got nil labels")
	}
	delete(existing, TestingLabelName)

	// Remove the testing invalid labels
	if !labelInSpec {
		newLabels := make(map[string]string)
		for k, v := range existing {
			if _, ok := TestingInvalidLabels[k]; ok {
				//	t.Fatalf("label %q with value %q is not a valid GCP label, it should not be set", k, v)
				continue
			}
			newLabels[k] = v
		}
		u.SetLabels(newLabels)
	} else {
		err := unstructured.SetNestedStringMap(u.Object, existing, "spec", "labels")
		if err != nil {
			t.Fatalf("failed to set labels in spec: %v", err)
		}
	}
}

func VerifyLabelsInGoldenHTTPLog(t *testing.T, u *unstructured.Unstructured, events test.LogEntries) {
	// if support gcp labels, the label should exist in the event
	existing, found, err := unstructured.NestedStringMap(u.Object, "spec", "labels")
	if !found || err != nil {
		existing = u.GetLabels()
	}
	if existing == nil {
		return
	}
	// testing label unset
	if _, ok := existing[TestingLabelName]; !ok {
		return
	}

	// Verify testing labels are reconciled in GCP HTTP log, and remove them.
	testLabelFoundAndRemoved := false
	for i := range events {
		// Handle Response
		if events[i].Response.Body != "" {
			body := events[i].Response.ParseBody()
			if body != nil {
				if findAndRemoveTestLabel(body) {
					testLabelFoundAndRemoved = true
					updatedBodyBytes, err := json.Marshal(body)
					if err != nil {
						t.Fatalf("failed to marshal updated response body: %v", err)
					}
					events[i].Response.Body = string(updatedBodyBytes)
				}
			}
		}

		// Handle Request
		if events[i].Request.Body != "" {
			body := events[i].Request.ParseBody()
			if body != nil {
				if findAndRemoveTestLabel(body) {
					testLabelFoundAndRemoved = true
					updatedBodyBytes, err := json.Marshal(body)
					if err != nil {
						t.Fatalf("failed to marshal updated request body: %v", err)
					}
					events[i].Request.Body = string(updatedBodyBytes)
				}
			}
		}
	}

	if !testLabelFoundAndRemoved {
		t.Errorf("expected testing label %q to be present in GCP HTTP log, but it was not found", TestingLabelName)
	}
}

// findAndRemoveTestLabel recursively searches for the testing label in the response body and removes it.
// Returns true if the label was found and removed.
func findAndRemoveTestLabel(data interface{}) bool {
	found := false
	switch v := data.(type) {
	case map[string]interface{}:
		// Check for labels map at the current level
		if labels, ok := v["labels"].(map[string]interface{}); ok {
			if _, exists := labels[TestingLabelName]; exists {
				delete(labels, TestingLabelName)
				found = true
			}
		}
		// Recurse into nested maps and slices
		for _, value := range v {
			if findAndRemoveTestLabel(value) {
				found = true
			}
		}
	case []interface{}:
		for _, item := range v {
			if findAndRemoveTestLabel(item) {
				found = true
			}
		}
	}
	return found
}
