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

package lint

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type CAIResourceEntry struct {
	ResourceType string   `json:"resourceType"`
	NameFormats  []string `json:"nameFormats"`
}

func TestCAIResourceCoverage(t *testing.T) {
	// Load all CRDs to find KCC resources
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading all CRDs: %v", err)
	}

	// Create a map of GroupKind to CRD for quick lookup
	gkToCRD := make(map[schema.GroupKind]bool)
	for _, crd := range crds {
		gk := schema.GroupKind{
			Group: crd.Spec.Group,
			Kind:  crd.Spec.Names.Kind,
		}
		gkToCRD[gk] = true
	}

	// Open the CAI metadata file
	caiPath := "../../docs/ai/metadata/cloudassetinventory_names.jsonl"
	file, err := os.Open(caiPath)
	if err != nil {
		t.Fatalf("failed to open CAI metadata at %s: %v", caiPath, err)
	}
	defer file.Close()

	var lines []string
	var correctCount, totalKCCCount int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry CAIResourceEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			t.Fatalf("failed to unmarshal CAI entry: %v", err)
		}

		// Identify matching KCC GVK via heuristics
		matchedGK := schema.GroupKind{}
		caiTokens := strings.Split(entry.ResourceType, "/")
		if len(caiTokens) != 2 {
			continue
		}
		serviceHost := caiTokens[0]
		caiResourceName := caiTokens[1]
		servicePrefix := strings.Split(serviceHost, ".")[0]

		for gk := range gkToCRD {
			if strings.HasPrefix(gk.Group, servicePrefix+".") {
				// Check if kind matches resource name, or servicePrefix + resourceName
				if strings.EqualFold(gk.Kind, caiResourceName) || 
					strings.EqualFold(gk.Kind, servicePrefix+caiResourceName) {
					matchedGK = gk
					break
				}
			}
		}

		gvkStr := "MISSING"
		isDirectController := false
		usesIdentityPattern := false
		usesRefPattern := false

		if matchedGK.Kind != "" {
			gvkStr = fmt.Sprintf("%s/%s", matchedGK.Group, matchedGK.Kind)
			
			// Check if direct
			if registry.IsDirectByGK(matchedGK) {
				isDirectController = true
			}

			// Check if identity is registered
			if _, err := refsv1beta1.NewIdentity(matchedGK); err == nil {
				usesIdentityPattern = true
			}

			// Check if reference is registered
			if _, err := refsv1beta1.NewRef(matchedGK); err == nil {
				usesRefPattern = true
			}

			if isDirectController && usesIdentityPattern && usesRefPattern {
				correctCount++
			}
			totalKCCCount++
		}

		reportLine := fmt.Sprintf("cai_type=%s kcc_gvk=%s isDirectController=%t usesIdentityPattern=%t usesRefPattern=%t",
			entry.ResourceType, gvkStr, isDirectController, usesIdentityPattern, usesRefPattern)
		lines = append(lines, reportLine)
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("scanner error: %v", err)
	}

	t.Logf("Count of fully modernized resources (isDirectController=true, usesIdentityPattern=true, usesRefPattern=true): %d out of %d", correctCount, totalKCCCount)

	sort.Strings(lines)
	want := strings.Join(lines, "\n") + "\n"

	test.CompareGoldenFile(t, "testdata/exceptions/cai_resource_coverage.txt", want)
}
