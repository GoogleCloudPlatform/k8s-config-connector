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
	ignoredTemplates := map[string]bool{
		// Add known exceptions here.
		// Example: "//some.googleapis.com/foo/{}/bar": true,
		"//storage.googleapis.com/projects/{}/buckets/{}":                   true,
		"//connectors.googleapis.com/projects/{}/locations/{}/providers/{}": true,
		"//bigtableadmin.googleapis.com/projects/{}/instances/{}/appProfiles/{}": true,
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
