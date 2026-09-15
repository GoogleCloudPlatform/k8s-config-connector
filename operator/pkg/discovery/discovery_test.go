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

package discovery_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/discovery"
)

func TestFilterCRDNames(t *testing.T) {
	// Sample of 10 CRDs across supported and unsupported services
	mockCatalog := []string{
		"configconnectors.core.cnrm.cloud.google.com",
		"computeinstances.compute.cnrm.cloud.google.com",
		"computenetworks.compute.cnrm.cloud.google.com",
		"storagebuckets.storage.cnrm.cloud.google.com",
		"pubsubtopics.pubsub.cnrm.cloud.google.com",
		"iamserviceaccounts.iam.cnrm.cloud.google.com",
		// Unsupported services in Sovereign Germany
		"spannerinstances.spanner.cnrm.cloud.google.com",
		"alloydbclusters.alloydb.cnrm.cloud.google.com",
		"aiplatformmodels.aiplatform.cnrm.cloud.google.com",
		"apigeeorganizations.apigee.cnrm.cloud.google.com",
	}

	// Enabled services in Sovereign Germany
	enabledInGermany := []string{
		"compute.googleapis.com",
		"storage.googleapis.com",
		"pubsub.googleapis.com",
		"iam.googleapis.com",
	}

	retained, report := discovery.FilterCRDNames(mockCatalog, enabledInGermany)

	if report.TotalCRDsInCatalog != 10 {
		t.Fatalf("expected 10 total CRDs, got %d", report.TotalCRDsInCatalog)
	}

	// Expected retained: core (1) + compute (2) + storage (1) + pubsub (1) + iam (1) = 6
	if len(retained) != 6 {
		t.Fatalf("expected 6 retained CRDs, got %d: %v", len(retained), retained)
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
