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

package releases

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestReferenceDoc(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	var missing []string
	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			gvk := schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: version.Name,
				Kind:    crd.Spec.Names.Kind,
			}
			if strings.HasSuffix(gvk.Group, "core.cnrm.cloud.google.com") {
				continue
			}
			if strings.HasPrefix(gvk.Version, "v1alpha1") {
				continue
			}
			// IAMServiceAccount is a special case.
			if gvk.Kind == "IAMServiceAccount" {
				continue
			}
			if hasReferenceDoc(t, gvk) {
				continue
			}
			missing = append(missing, gvk.String())
		}
	}

	sort.Strings(missing)
	want := strings.Join(missing, "\n")
	test.CompareGoldenFile(t, "testdata/missing_reference.txt", want)
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
