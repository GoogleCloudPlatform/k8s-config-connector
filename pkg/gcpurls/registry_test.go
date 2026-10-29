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
	"path/filepath"
	"regexp"
	"testing"

	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1alpha1"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/commerceproducer/v1alpha1"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/filestore/v1beta1"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/apis/healthcare/v1alpha1"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/yaml"
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
		// AI Platform
		"//aiplatform.googleapis.com/projects/{}/locations/{}/ragCorpora/{}": true,

		// AlloyDB

		// Apigee
		"//apigee.googleapis.com/organizations/{}/addonsConfig":      true,
		"//apigee.googleapis.com/organizations/{}/syncAuthorization": true,

		// Apigee Registry
		"//apigeeregistry.googleapis.com/projects/{}/locations/{}/apis/{}":      true,
		"//apigeeregistry.googleapis.com/projects/{}/locations/{}/artifacts/{}": true,
		"//apigeeregistry.googleapis.com/projects/{}/locations/{}/instances/{}": true,

		// Artifact Registry

		// AutoML
		"//automl.googleapis.com/projects/{}/locations/{}/datasets/{}": true,

		// BigLake

		// BigQuery
		"//bigquery.googleapis.com/projects/{}/datasets/{}/access/{}": true,

		// Bigtable
		"//bigtable.googleapis.com/projects/{}/instances/{}/tables/{}/columnFamilies/{}": true,

		// Commerce Producer
		"//commerceproducer.googleapis.com/projects/{}/locations/{}/privateOffers/{}":              true,
		"//commerceproducer.googleapis.com/projects/{}/locations/{}/services/{}/standardOffers/{}": true,
		"//commerceproducer.googleapis.com/projects/{}/locations/{}/services/{}/skus/{}":           true,

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
		"//compute.googleapis.com/locations/global/securityPolicies/{}/associations/{}":   true,
		"//compute.googleapis.com/projects/{}/global/backendServices/{}/signedUrlKeys/{}": true,
		"//compute.googleapis.com/projects/{}/regions/{}/instanceGroupManagers/{}/{}":     true,
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
		"//oracledatabase.googleapis.com/projects/{}/locations/{}/autonomousDatabaseBackups/{}": true,

		// OSLogin

		// Rapid Migration Assessment

		// Service Usage
		"//serviceusage.googleapis.com/projects/{}/services/{}/identity": true,

		// Storage
		"//storage.googleapis.com/projects/{}/buckets/{}":            true,
		"//storage.googleapis.com/projects/{}/buckets/{}/objects/{}": true,

		// Translate
		"//translate.googleapis.com/projects/{}/locations/{}/adaptiveMtDatasets/{}": true,

		// Vision

		// Vision AI
		"//visionai.googleapis.com/projects/{}/locations/{}/applications/{}": true,

		// Workflow Executions

	}
	loadJudgementQueueCAIIgnores(t, "../../apis/*/judgement_queue.yaml", ignoredTemplates)

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

func TestLoadJudgementQueueCAIIgnores(t *testing.T) {
	tmpDir := t.TempDir()
	qFile := filepath.Join(tmpDir, "judgement_queue.yaml")
	content := `entries:
  - kind: ExampleResource
    group: example.cnrm.cloud.google.com
    reason: identity-not-in-cai
    detail: template "//example.googleapis.com/projects/{project}/locations/{location}/widgets/{widget}" is not listed in docs/ai/metadata/cloudassetinventory_names.jsonl; verify the URL format or add an exception in pkg/gcpurls/registry_test.go
    status: open
  - kind: RejectedResource
    group: example.cnrm.cloud.google.com
    reason: identity-not-in-cai
    detail: template "//example.googleapis.com/projects/{project}/locations/{location}/badWidgets/{widget}" is not listed in docs/ai/metadata/cloudassetinventory_names.jsonl
    status: rejected
  - kind: OtherResource
    group: example.cnrm.cloud.google.com
    reason: identity-multi-pattern
    detail: template "//example.googleapis.com/projects/{project}/bars/{bar}"
    status: open
`
	if err := os.WriteFile(qFile, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write temp queue file: %v", err)
	}
	ignored := make(map[string]bool)
	loadJudgementQueueCAIIgnores(t, filepath.Join(tmpDir, "*.yaml"), ignored)
	want := "//example.googleapis.com/projects/{}/locations/{}/widgets/{}"
	if !ignored[want] {
		t.Errorf("expected %q to be in ignored map, got %v", want, ignored)
	}
	if len(ignored) != 1 {
		t.Errorf("expected exactly 1 ignored entry, got %v", ignored)
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

type judgementQueueFile struct {
	Entries []struct {
		Reason string `json:"reason"`
		Detail string `json:"detail"`
		Status string `json:"status"`
	} `json:"entries"`
}

var judgementTemplateRegex = regexp.MustCompile(`template "([^"]+)"`)

func loadJudgementQueueCAIIgnores(t *testing.T, globPattern string, ignored map[string]bool) {
	t.Helper()
	files, err := filepath.Glob(globPattern)
	if err != nil {
		t.Fatalf("failed to glob judgement queues %s: %v", globPattern, err)
	}
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("failed to read judgement queue %s: %v", f, err)
		}
		var q judgementQueueFile
		if err := yaml.Unmarshal(data, &q); err != nil {
			t.Fatalf("failed to unmarshal judgement queue %s: %v", f, err)
		}
		for _, entry := range q.Entries {
			if entry.Reason != "identity-not-in-cai" || entry.Status == "rejected" {
				continue
			}
			if m := judgementTemplateRegex.FindStringSubmatch(entry.Detail); len(m) == 2 {
				ignored[normalizeTemplateFormat(m[1])] = true
			}
		}
	}
}
