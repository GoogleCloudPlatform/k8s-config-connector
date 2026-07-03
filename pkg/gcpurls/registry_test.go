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

package gcpurls_test

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"testing"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/filestore/v1beta1"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

type CAIEntry struct {
	ResourceType string   `json:"resourceType"`
	NameFormats  []string `json:"nameFormats"`
}

func TestRegisteredTemplatesMatchCAI(t *testing.T) {
	// Load CAI definitions
	caiFormats := make(map[string]bool)

	// Path relative to pkg/gcpurls
	caiPath := "../../docs/ai/metadata/cloudassetinventory_names.jsonl"
	file, err := os.Open(caiPath)
	if err != nil {
		t.Fatalf("failed to open CAI metadata at %s: %v", caiPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry CAIEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			t.Fatalf("failed to unmarshal CAI entry: %v", err)
		}
		for _, format := range entry.NameFormats {
			normalized := normalizeCAIFormat(format)
			caiFormats[normalized] = true
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("scanner error: %v", err)
	}

	templates := gcpurls.AllTemplates()
	if len(templates) == 0 {
		t.Fatal("no templates registered")
	}
	t.Logf("Checking %d registered templates", len(templates))

	// Exceptions for templates that are known not to match CAI or are not in CAI.
	// We use the normalized format for the key.
	//
	// NOTE ON "WRONG" PATTERNS / MISMATCHES:
	// If Cloud Asset Inventory added support for an asset, and we had given it a different "url template":
	ignoredTemplates := map[string]bool{
		// AI Platform
		"//aiplatform.googleapis.com/projects/{}/locations/{}/exampleStores/{}": true,

		// AlloyDB
		"//alloydb.googleapis.com/projects/{}/locations/{}/clusters/{}/users/{}": true,

		// Apigee Registry
		"//apigeeregistry.googleapis.com/projects/{}/locations/{}/apis/{}":      true,
		"//apigeeregistry.googleapis.com/projects/{}/locations/{}/artifacts/{}": true,
		"//apigeeregistry.googleapis.com/projects/{}/locations/{}/instances/{}": true,

		// Artifact Registry
		"//artifactregistry.googleapis.com/projects/{}/locations/{}/vpcscConfig": true,

		// AI Streams
		"//aistreams.googleapis.com/projects/{}/locations/{}/clusters/{}": true,

		// AutoML
		"//automl.googleapis.com/projects/{}/locations/{}/datasets/{}": true,

		// Batch
		"//batch.googleapis.com/projects/{}/locations/{}/resourceAllowances/{}": true,

		// BigLake
		"//biglake.googleapis.com/projects/{}/locations/{}/catalogs/{}": true,

		// BigQuery Connection
		"//bigqueryconnection.googleapis.com/projects/{}/locations/{}/connections/{}": true,

		// Bigtable
		"//bigtable.googleapis.com/projects/{}/instances/{}/tables/{}/columnFamilies/{}": true,

		// Billing Budgets
		"//billingbudgets.googleapis.com/billingAccounts/{}/budgets/{}": true,

		// Cloud KMS
		"//cloudkms.googleapis.com/projects/{}/locations/{}/keyRings/{}/cryptoKeys/{}/ciphertext/{}": true,

		// Cloud Security Compliance
		"//cloudsecuritycompliance.googleapis.com/organizations/{}/locations/{}/cloudControls/{}":      true,
		"//cloudsecuritycompliance.googleapis.com/organizations/{}/locations/{}/cloudControlGroups/{}": true,
		"//cloudsecuritycompliance.googleapis.com/organizations/{}/locations/{}/controls/{}":           true,
		"//cloudsecuritycompliance.googleapis.com/projects/{}/locations/{}/frameworks/{}":              true,
		"//cloudsecuritycompliance.googleapis.com/organizations/{}/locations/{}/frameworks/{}":         true,

		// Compute
		"//compute.googleapis.com/global/publicDelegatedPrefixes/{}":                      true,
		"//compute.googleapis.com/projects/{}/global/backendServices/{}/signedUrlKeys/{}": true,
		"//compute.googleapis.com/projects/{}/global/images/family/{}":                    true,
		"//compute.googleapis.com/projects/{}/zones/{}/disks/{}/{}":                       true,
		"//compute.googleapis.com/projects/{}/zones/{}/futureReservations/{}":             true,
		"//compute.googleapis.com/regions/{}/publicDelegatedPrefixes/{}":                  true,
		"//compute.googleapis.com/projects/{}/regions/{}/routers/{}/interfaces/{}":        true,
		"//compute.googleapis.com/projects/{}/regions/{}/routers/{}/{}":                   true,
		"//compute.googleapis.com/projects/{}/zones/{}/networkEndpointGroups/{}/{}/{}/{}": true,
		"//compute.googleapis.com/projects/{}/zones/{}/networkEndpointGroups/{}//{}/{}":   true,

		// Connectors
		"//connectors.googleapis.com/projects/{}/locations/{}/providers/{}": true,

		// Content Warehouse
		"//contentwarehouse.googleapis.com/projects/{}/locations/{}/documentSchemas/{}": true,
		"//contentwarehouse.googleapis.com/projects/{}/locations/{}/ruleSets/{}":        true,

		// Data Labeling
		"//datalabeling.googleapis.com/projects/{}/evaluationJobs/{}": true,

		// Dataplex
		"//dataplex.googleapis.com/projects/{}/locations/{}/aspectTypes/{}":           true,
		"//dataplex.googleapis.com/projects/{}/locations/{}/dataAttributeBindings/{}": true,
		"//dataplex.googleapis.com/projects/{}/locations/{}/dataTaxonomies/{}":        true,
		"//dataplex.googleapis.com/projects/{}/locations/{}/entryGroups/{}":           true,
		"//dataplex.googleapis.com/projects/{}/locations/{}/entryTypes/{}":            true,

		// Dataproc
		"//dataproc.googleapis.com/projects/{}/locations/{}/sessionTemplates/{}": true,
		"//dataproc.googleapis.com/v1/projects/{}/regions/{}/clusters/{}":        true,

		// Dialogflow
		"//dialogflow.googleapis.com/projects/{}/locations/{}/generators/{}":       true,
		"//dialogflow.googleapis.com/projects/{}/locations/{}/securitySettings/{}": true,
		"//dialogflow.googleapis.com/projects/{}/locations/{}/sipTrunks/{}":        true,

		// Discovery Engine
		"//discoveryengine.googleapis.com/projects/{}/locations/{}/dataStores/{}/controls/{}":      true,
		"//discoveryengine.googleapis.com/projects/{}/locations/{}/dataStores/{}/conversations/{}": true,
		"//discoveryengine.googleapis.com/projects/{}/locations/{}/identityMappingStores/{}":       true,

		// DLP
		"//dlp.googleapis.com/projects/{}/locations/{}/connections/{}": true,

		// DNS
		"//dns.googleapis.com/projects/{}/managedZones/{}/rrsets/{}":    true,
		"//dns.googleapis.com/projects/{}/responsePolicies/{}":          true,
		"//dns.googleapis.com/projects/{}/responsePolicies/{}/rules/{}": true,

		// Firestore
		"//firestore.googleapis.com/projects/{}/databases/{}/backupSchedules/{}":             true,
		"//firestore.googleapis.com/projects/{}/databases/{}/collectionGroups/{}":            true,
		"//firestore.googleapis.com/projects/{}/databases/{}/collectionGroups/{}/indexes/{}": true,

		// IAM
		"//iam.googleapis.com/policies/{}/denypolicies/{}": true,

		// IAP
		"//iap.googleapis.com/projects/{}/brands/{}": true,

		// Logging
		"//logging.googleapis.com/billingAccounts/{}/exclusions/{}": true,
		"//logging.googleapis.com/folders/{}/exclusions/{}":         true,
		"//logging.googleapis.com/organizations/{}/exclusions/{}":   true,
		"//logging.googleapis.com/projects/{}/exclusions/{}":        true,

		// Monitoring
		"//monitoring.googleapis.com/locations/global/metricsScopes/{}/projects/{}": true,
		"//monitoring.googleapis.com/projects/{}/groups/{}":                         true,
		"//monitoring.googleapis.com/projects/{}/metricDescriptors/{}":              true,
		"//monitoring.googleapis.com/projects/{}/services/{}":                       true,

		// Migration Center
		"//migrationcenter.googleapis.com/projects/{}/locations/{}/groups/{}": true,

		// Network Connectivity
		"//networkconnectivity.googleapis.com/projects/{}/locations/{}/regionalEndpoints/{}": true,

		// Network Security
		"//networksecurity.googleapis.com/projects/{}/locations/{}/backendAuthenticationConfigs/{}": true,
		"//networksecurity.googleapis.com/projects/{}/locations/{}/sacRealms/{}":                    true,
		"//networksecurity.googleapis.com/projects/{}/locations/{}/securityProfiles/{}":             true,

		// Network Services
		"//networkservices.googleapis.com/projects/{}/locations/global/edgeCacheServices/{}": true,

		// Notebooks
		"//notebooks.googleapis.com/projects/{}/locations/{}/environments/{}": true,
		"//notebooks.googleapis.com/projects/{}/locations/{}/executions/{}":   true,

		// OSConfig
		"//osconfig.googleapis.com/projects/{}/guestPolicies/{}": true,

		// Privileged Access Manager
		"//privilegedaccessmanager.googleapis.com/folders/{}/locations/{}/entitlements/{}":       true,
		"//privilegedaccessmanager.googleapis.com/organizations/{}/locations/{}/entitlements/{}": true,
		"//privilegedaccessmanager.googleapis.com/projects/{}/locations/{}/entitlements/{}":      true,

		// Service Usage
		"//serviceusage.googleapis.com/projects/{}/services/{}/identity": true,

		// Storage
		"//storage.googleapis.com/projects/{}/buckets/{}": true,

		// Workflow Executions
		"//workflowexecutions.googleapis.com/projects/{}/locations/{}/workflows/{}/executions/{}": true,
	}
	for _, tmpl := range templates {
		fullURL := "//" + tmpl.Host() + "/" + tmpl.CanonicalForm()
		normalized := normalizeTemplateFormat(fullURL)
		if tmpl.Host() == "" || tmpl.Host() == "example.com" {
			continue
		}

		if ignoredTemplates[normalized] {
			continue
		}

		if !caiFormats[normalized] {
			t.Errorf("Registered template %q (normalized: %q) not found in CAI definitions", fullURL, normalized)
		}
	}
}

var caiVarRegex = regexp.MustCompile(`\{\{[^}]+\}\}`)
var tmplVarRegex = regexp.MustCompile(`\{[^}]+\}`)

func normalizeCAIFormat(s string) string {
	return caiVarRegex.ReplaceAllString(s, "{}")
}

func normalizeTemplateFormat(s string) string {
	return tmplVarRegex.ReplaceAllString(s, "{}")
}
