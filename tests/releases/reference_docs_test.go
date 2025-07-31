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

package references

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestReferenceDoc(t *testing.T) {
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatalf("error creating service mapping loader: %v", err)
	}
	serviceMetadataLoader := metadata.New()
	allGVKs, err := supportedgvks.All(smLoader, serviceMetadataLoader)
	if err != nil {
		t.Fatalf("error getting all GVKs: %v", err)
	}

	var missing []string
	for _, gvk := range allGVKs {
		if strings.HasPrefix(gvk.Version, "v1alpha1") {
			continue
		}
		// IAMServiceAccount is a special case, it is not a real resource.
		if gvk.Kind == "IAMServiceAccount" {
			continue
		}
		if hasReferenceDoc(t, gvk) {
			continue
		}
		missing = append(missing, gvk.String())
	}

	if len(missing) == 0 {
		// Create an empty file if it does not exist
		goldenFile := "testdata/missing_reference.txt"
		if _, err := os.Stat(goldenFile); os.IsNotExist(err) {
			if err := ioutil.WriteFile(goldenFile, []byte{}, 0644); err != nil {
				t.Fatalf("error creating empty golden file: %v", err)
			}
		}
		return
	}

	sort.Strings(missing)
	missingRefs := strings.Join(missing, "\n") + "\n"

	goldenFile := "testdata/missing_reference.txt"
	if os.Getenv("UPDATE_GOLDEN") != "" {
		if err := ioutil.WriteFile(goldenFile, []byte(missingRefs), 0644); err != nil {
			t.Fatalf("error writing golden file: %v", err)
		}
		return
	}

	bytes, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatalf("error reading golden file: %v", err)
	}
	if diff := cmp.Diff(string(bytes), missingRefs); diff != "" {
		t.Errorf("unexpected diff in missing references (-want +got):\n%s", diff)
	}
}

func hasReferenceDoc(t *testing.T, gvk schema.GroupVersionKind) bool {
	// This logic is based on scripts/generate-google3-docs/resource-reference/main.go
	templateFileName := templateFileNameForGVK(gvk)
	templatePath := filepath.Join(repo.GetG3ResourceReferenceTemplatesPath(), templateFileName)
	if _, err := os.Stat(templatePath); err == nil {
		return true
	} else if !os.IsNotExist(err) {
		t.Fatalf("error stating template file %s: %v", templatePath, err)
	}
	return false
}

func templateFileNameForGVK(gvk schema.GroupVersionKind) string {
	// This logic is copied from scripts/generate-google3-docs/resource-reference/main.go
	return strings.ToLower(strings.Join([]string{groupName(gvk), gvk.Kind}, "_")) + ".tmpl"
}

func groupName(gvk schema.GroupVersionKind) string {
	// This logic is copied from scripts/generate-google3-docs/resource-reference/main.go
	return strings.SplitN(gvk.Group, ".", 2)[0]
}
