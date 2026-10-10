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

	// Paths relative to pkg/gcpurls
	metadataPaths := []string{
		"../../docs/ai/metadata/cloudassetinventory_names.jsonl",
		"../../docs/ai/metadata/discovery_names.jsonl",
	}
	for _, metadataPath := range metadataPaths {
		loadMetadataFormats(t, metadataPath, caiFormats)
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

		// AlloyDB

		// Apigee Registry
		"//apigeeregistry.googleapis.com/projects/{}/locations/{}/apis/{}":      true,
		"//apigeeregistry.googleapis.com/projects/{}/locations/{}/artifacts/{}": true,
		"//apigeeregistry.googleapis.com/projects/{}/locations/{}/instances/{}": true,

		// Artifact Registry

		// AutoML
		"//automl.googleapis.com/projects/{}/locations/{}/datasets/{}": true,

		// BigLake

		// Bigtable
		"//bigtable.googleapis.com/projects/{}/instances/{}/tables/{}/columnFamilies/{}": true,

		// Cloud KMS
		"//cloudkms.googleapis.com/projects/{}/locations/{}/keyRings/{}/cryptoKeys/{}/ciphertext/{}": true,

		// Cloud Number Registry

		// Cloud Security Compliance
		"//cloudsecuritycompliance.googleapis.com/organizations/{}/locations/{}/cloudControls/{}":      true,
		"//cloudsecuritycompliance.googleapis.com/organizations/{}/locations/{}/cloudControlGroups/{}": true,
		"//cloudsecuritycompliance.googleapis.com/organizations/{}/locations/{}/controls/{}":           true,
		"//cloudsecuritycompliance.googleapis.com/projects/{}/locations/{}/frameworks/{}":              true,
		"//cloudsecuritycompliance.googleapis.com/organizations/{}/locations/{}/frameworks/{}":         true,

		// Compute
		"//compute.googleapis.com/global/publicDelegatedPrefixes/{}":                      true,
		"//compute.googleapis.com/projects/{}/global/backendServices/{}/signedUrlKeys/{}": true,
		"//compute.googleapis.com/projects/{}/zones/{}/disks/{}/{}":                       true,
		"//compute.googleapis.com/regions/{}/publicDelegatedPrefixes/{}":                  true,
		"//compute.googleapis.com/projects/{}/regions/{}/routers/{}/interfaces/{}":        true,
		"//compute.googleapis.com/projects/{}/regions/{}/routers/{}/{}":                   true,
		"//compute.googleapis.com/projects/{}/zones/{}/networkEndpointGroups/{}/{}/{}/{}": true,
		"//compute.googleapis.com/projects/{}/zones/{}/networkEndpointGroups/{}//{}/{}":   true,

		// Config Delivery
		"//configdelivery.googleapis.com/projects/{}/locations/{}/fleetPackages/{}":   true,
		"//configdelivery.googleapis.com/projects/{}/locations/{}/resourceBundles/{}": true,

		// Connectors

		// Content Warehouse

		// Data Labeling
		"//datalabeling.googleapis.com/projects/{}/annotationSpecSets/{}": true,
		"//datalabeling.googleapis.com/projects/{}/datasets/{}":           true,
		"//datalabeling.googleapis.com/projects/{}/evaluationJobs/{}":     true,
		"//datalabeling.googleapis.com/projects/{}/instructions/{}":       true,

		// Dataplex
		"//dataplex.googleapis.com/projects/{}/locations/{}/dataAttributeBindings/{}": true,
		"//dataplex.googleapis.com/projects/{}/locations/{}/dataTaxonomies/{}":        true,
		// Dataproc
		"//dataproc.googleapis.com/v1/projects/{}/regions/{}/clusters/{}": true,

		// Device Streaming
		"//devicestreaming.googleapis.com/projects/{}/deviceSessions/{}": true,

		// Discovery Engine
		"//discoveryengine.googleapis.com/projects/{}/locations/{}/licenseConfigs/{}/": true,

		// DNS
		"//dns.googleapis.com/projects/{}/managedZones/{}/rrsets/{}":    true,
		"//dns.googleapis.com/projects/{}/responsePolicies/{}":          true,
		"//dns.googleapis.com/projects/{}/responsePolicies/{}/rules/{}": true,

		// Firebase Hosting

		// Firestore
		"//firestore.googleapis.com/projects/{}/databases/{}/collectionGroups/{}": true,

		// IAM
		"//iam.googleapis.com/policies/{}/denypolicies/{}": true,

		// Cloud Talent Solution

		// License Manager
		"//licensemanager.googleapis.com/projects/{}/locations/{}/configurations/{}": true,

		// Monitoring
		"//monitoring.googleapis.com/projects/{}/services/{}": true,

		// Map Management
		"//mapmanagement.googleapis.com/projects/{}/mapConfigs/{}":   true,
		"//mapmanagement.googleapis.com/projects/{}/styleConfigs/{}": true,

		// Model Armor
		"//modelarmor.googleapis.com/projects/{}/locations/{}/templates/{}":      true,
		"//modelarmor.googleapis.com/projects/{}/locations/{}/floorSetting":      true,
		"//modelarmor.googleapis.com/folders/{}/locations/{}/floorSetting":       true,
		"//modelarmor.googleapis.com/organizations/{}/locations/{}/floorSetting": true,

		// Network Security

		// Network Services
		"//networkservices.googleapis.com/projects/{}/locations/global/edgeCacheServices/{}": true,

		// Oracle Database

		// OSLogin

		// Rapid Migration Assessment

		// Service Usage
		"//serviceusage.googleapis.com/projects/{}/services/{}/identity": true,

		// Storage
		"//storage.googleapis.com/projects/{}/buckets/{}":            true,
		"//storage.googleapis.com/projects/{}/buckets/{}/objects/{}": true,

		// Vision

		// Workflow Executions

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

func loadMetadataFormats(t *testing.T, metadataPath string, out map[string]bool) {
	t.Helper()
	file, err := os.Open(metadataPath)
	if err != nil {
		t.Fatalf("failed to open metadata at %s: %v", metadataPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry CAIEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			t.Fatalf("failed to unmarshal metadata entry in %s: %v", metadataPath, err)
		}
		for _, format := range entry.NameFormats {
			normalized := normalizeCAIFormat(format)
			out[normalized] = true
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("scanner error reading %s: %v", metadataPath, err)
	}
}
