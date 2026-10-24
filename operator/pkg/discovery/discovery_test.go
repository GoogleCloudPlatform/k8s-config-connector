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

package discovery_test

import (
	"errors"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/discovery"
)

func TestFilterCRDNames(t *testing.T) {
	mockCatalog := []string{
		"configconnectors.core.cnrm.cloud.google.com",
		"computeinstances.compute.cnrm.cloud.google.com",
		"computenetworks.compute.cnrm.cloud.google.com",
		"storagebuckets.storage.cnrm.cloud.google.com",
		"pubsubtopics.pubsub.cnrm.cloud.google.com",
		"iamserviceaccounts.iam.cnrm.cloud.google.com",
		"runservices.run.cnrm.cloud.google.com",
		// Unsupported services in Sovereign Germany
		"spannerinstances.spanner.cnrm.cloud.google.com",
		"alloydbclusters.alloydb.cnrm.cloud.google.com",
		"aiplatformmodels.aiplatform.cnrm.cloud.google.com",
		"apigeeorganizations.apigee.cnrm.cloud.google.com",
	}

	// Enabled services including run.googleapis.com
	enabledInGermany := []string{
		"compute.googleapis.com",
		"storage.googleapis.com",
		"pubsub.googleapis.com",
		"iam.googleapis.com",
		"run.googleapis.com",
	}

	retained, report := discovery.FilterCRDNames(mockCatalog, enabledInGermany)

	if report.TotalCRDsInCatalog != 11 {
		t.Fatalf("expected 11 total CRDs, got %d", report.TotalCRDsInCatalog)
	}

	// Expected retained: core (1) + compute (2) + storage (1) + pubsub (1) + iam (1) + run (1) = 7
	if len(retained) != 7 {
		t.Fatalf("expected 7 retained CRDs, got %d: %v", len(retained), retained)
	}

	if report.PrunedCRDs != 4 {
		t.Fatalf("expected 4 pruned CRDs, got %d", report.PrunedCRDs)
	}

	// Verify unsupported groups were pruned
	for _, crd := range retained {
		if crd == "spannerinstances.spanner.cnrm.cloud.google.com" ||
			crd == "alloydbclusters.alloydb.cnrm.cloud.google.com" ||
			crd == "aiplatformmodels.aiplatform.cnrm.cloud.google.com" ||
			crd == "apigeeorganizations.apigee.cnrm.cloud.google.com" {
			t.Errorf("found unsupported CRD in retained list: %s", crd)
		}
	}
}

func TestFilterCRDNames_WithUnavailableServices(t *testing.T) {
	mockCatalog := []string{
		"configconnectors.core.cnrm.cloud.google.com",
		"computeinstances.compute.cnrm.cloud.google.com",
		"computenetworks.compute.cnrm.cloud.google.com",
		"storagebuckets.storage.cnrm.cloud.google.com",
		"pubsubtopics.pubsub.cnrm.cloud.google.com",
		"iamserviceaccounts.iam.cnrm.cloud.google.com",
		"runservices.run.cnrm.cloud.google.com",
	}

	// In sovereign cloud, run.googleapis.com appears in available services list,
	// but is not visible to the partition (SERVICE_NOT_VISIBLE).
	enabledServices := []string{
		"compute.googleapis.com",
		"storage.googleapis.com",
		"pubsub.googleapis.com",
		"iam.googleapis.com",
		"run.googleapis.com",
	}

	unavailableServices := []string{
		"run.googleapis.com",
	}

	retained, report := discovery.FilterCRDNames(mockCatalog, enabledServices, unavailableServices)

	// run.googleapis.com should be pruned because it is in unavailableServices
	for _, crd := range retained {
		if crd == "runservices.run.cnrm.cloud.google.com" {
			t.Errorf("expected RunService to be pruned when service is marked unavailable")
		}
	}

	// Expected retained: core (1) + compute (2) + storage (1) + pubsub (1) + iam (1) = 6
	if len(retained) != 6 {
		t.Fatalf("expected 6 retained CRDs, got %d: %v", len(retained), retained)
	}

	if report.PrunedCRDs != 1 {
		t.Fatalf("expected 1 pruned CRD, got %d", report.PrunedCRDs)
	}
}

func TestIsServiceNotVisibleError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "SERVICE_NOT_VISIBLE reason",
			err:      errors.New("googleapi: Error 403: The Cloud Run Admin API API is not visible to project partition:sample-project (or it may not exist), reason: SERVICE_NOT_VISIBLE"),
			expected: true,
		},
		{
			name:     "Generic 404 not found",
			err:      errors.New("googleapi: Error 404: Not found"),
			expected: false,
		},
		{
			name:     "Service not visible message text",
			err:      errors.New("The Cloud Run Admin API API is not visible to project partition:sample-project"),
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if actual := discovery.IsServiceNotVisibleError(tc.err); actual != tc.expected {
				t.Errorf("IsServiceNotVisibleError() = %v, want %v", actual, tc.expected)
			}
		})
	}
}
