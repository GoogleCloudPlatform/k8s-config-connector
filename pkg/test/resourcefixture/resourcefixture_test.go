// Copyright 2022 Google LLC
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

package resourcefixture_test

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var registerModelGVKRegex = regexp.MustCompile(`RegisterModel\(krm\.(\w+)GVK`)

func TestResourceContents(t *testing.T) {
	fixtures := resourcefixture.Load(t)
	immutableKinds, specOnlyGreenfieldKinds := loadDirectControllerUpdateCapabilities(t)
	for _, f := range fixtures {
		if f.Update == nil {
			continue
		}
		createUnstruct := test.ToUnstruct(t, f.Create)
		updateUnstruct := test.ToUnstruct(t, f.Update)
		testCreateAndUpdateUnstructMatchingProperties(t, f, createUnstruct, updateUnstruct)
		testCreateAndUpdateUnstructShouldDiffer(t, f, createUnstruct, updateUnstruct, immutableKinds, specOnlyGreenfieldKinds)
		dependencies := make([]*unstructured.Unstructured, 0)
		testAllNamesAreUnique(t, f, append(dependencies, createUnstruct)...)
	}
}

func loadDirectControllerUpdateCapabilities(t *testing.T) (immutableKinds map[string]bool, specOnlyGreenfieldKinds map[string]bool) {
	t.Helper()
	immutableKinds = make(map[string]bool)
	specOnlyGreenfieldKinds = make(map[string]bool)

	directDir := filepath.Join(repo.GetRootOrTestFatal(t), "pkg", "controller", "direct")
	err := filepath.WalkDir(directDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), "_controller.go") {
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(b)
		matches := registerModelGVKRegex.FindStringSubmatch(content)
		if len(matches) < 2 {
			return nil
		}
		kind := matches[1]
		if strings.Contains(content, "is not supported, fields") ||
			strings.Contains(content, fmt.Sprintf("updating %s is not supported", kind)) ||
			strings.Contains(content, fmt.Sprintf("update %s is not supported", kind)) {
			immutableKinds[kind] = true
		} else if strings.Contains(content, "CompareGreenfieldSpec(") &&
			!strings.Contains(content, "CompareGreenfieldSpecAndLabels(") {
			specOnlyGreenfieldKinds[kind] = true
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error scanning direct controllers in %s: %v", directDir, err)
	}
	return immutableKinds, specOnlyGreenfieldKinds
}

func testCreateAndUpdateUnstructMatchingProperties(t *testing.T, fixture resourcefixture.ResourceFixture,
	createUnstruct, updateUnstruct *unstructured.Unstructured) {
	testCreateAndUpdateUnstructShouldMatchNameAndNamespace(t, fixture, createUnstruct, updateUnstruct)
	testCreateAndUpdateUnstructShouldHaveTheSameGVK(t, fixture, createUnstruct, updateUnstruct)
}

func testCreateAndUpdateUnstructShouldMatchNameAndNamespace(t *testing.T, fixture resourcefixture.ResourceFixture,
	createUnstruct, updateUnstruct *unstructured.Unstructured) {
	if createUnstruct.GetName() != updateUnstruct.GetName() {
		t.Errorf("name mismatch between create and update for fixture '%v': create(%v), update(%v)",
			fixture.Name, createUnstruct.GetName(), updateUnstruct.GetName())
	}
	if createUnstruct.GetNamespace() != updateUnstruct.GetNamespace() {
		t.Errorf("namespace mismatch between create and update for fixture '%v': create(%v), update(%v)",
			fixture.Name, createUnstruct.GetNamespace(), updateUnstruct.GetNamespace())
	}
}

func testCreateAndUpdateUnstructShouldHaveTheSameGVK(t *testing.T, fixture resourcefixture.ResourceFixture,
	createUnstruct, updateUnstruct *unstructured.Unstructured) {
	if createUnstruct.GroupVersionKind() != updateUnstruct.GroupVersionKind() {
		t.Errorf("gvk mismatch between create and update for fixture '%v': create(%v), update(%v)",
			fixture.Name, createUnstruct.GroupVersionKind(), updateUnstruct.GroupVersionKind())
	}
}

func testAllNamesAreUnique(t *testing.T, fixture resourcefixture.ResourceFixture, resources ...*unstructured.Unstructured) {
	names := make(map[string]*unstructured.Unstructured)
	for _, r := range resources {
		if _, ok := names[r.GetName()]; ok {
			t.Errorf("fixture '%v' has more than one resource with name '%v': all resource names within a fixture must be unique",
				fixture.Name, r.GetName())
		}
		names[r.GetName()] = r
	}
}

func testCreateAndUpdateUnstructShouldDiffer(t *testing.T, fixture resourcefixture.ResourceFixture,
	createUnstruct, updateUnstruct *unstructured.Unstructured,
	immutableKinds, specOnlyGreenfieldKinds map[string]bool) {
	if err := checkUpdateFixtureDiff(fixture.Name, createUnstruct, updateUnstruct, immutableKinds, specOnlyGreenfieldKinds); err != nil {
		t.Error(err)
	}
}

func checkUpdateFixtureDiff(fixtureName string, createUnstruct, updateUnstruct *unstructured.Unstructured,
	immutableKinds, specOnlyGreenfieldKinds map[string]bool) error {
	exemptedUpdateNoChange := map[string]bool{
		"bigquerydatasetnochangeupdate":                 true,
		"bigquerydatasetnochangeupdate-direct":          true,
		"bigquerytable-view":                            true,
		"certificatemanagercertificateissuanceconfig":   true,
		"containercluster-resourcemanagertags-standard": true,
		"dataformrepository-basic":                      true,
		"datalabelinginstruction-maximal":               true,
		"networkconnectivityregionalendpoint-maximal":   true,
		"networkconnectivityregionalendpoint-minimal":   true,
	}
	if exemptedUpdateNoChange[fixtureName] {
		return nil
	}

	kind := createUnstruct.GetKind()
	if immutableKinds[kind] {
		return fmt.Errorf("fixture %q has update.yaml, but %s is an immutable resource (its direct controller does not support updates); omit update.yaml for immutable resources", fixtureName, kind)
	}

	createSpec := createUnstruct.Object["spec"]
	updateSpec := updateUnstruct.Object["spec"]

	if specOnlyGreenfieldKinds[kind] && reflect.DeepEqual(createSpec, updateSpec) {
		return fmt.Errorf("fixture %q has update.yaml that does not modify spec (only metadata), but %s controller only reconciles spec via CompareGreenfieldSpec; either modify a mutable spec field in update.yaml or omit update.yaml if the resource/fixture has no mutable spec fields", fixtureName, kind)
	}

	createLabels := createUnstruct.GetLabels()
	updateLabels := updateUnstruct.GetLabels()

	createAnnotations := createUnstruct.GetAnnotations()
	updateAnnotations := updateUnstruct.GetAnnotations()

	if reflect.DeepEqual(createSpec, updateSpec) &&
		reflect.DeepEqual(createLabels, updateLabels) &&
		reflect.DeepEqual(createAnnotations, updateAnnotations) {
		return fmt.Errorf("fixture %q has update.yaml but it does not make any changes to spec, labels, or annotations compared to create.yaml. An update.yaml must trigger a real update/patch.", fixtureName)
	}
	return nil
}

func TestCheckUpdateFixtureDiff(t *testing.T) {
	immutableKinds := map[string]bool{
		"TranslateDataset": true,
	}
	specOnlyGreenfieldKinds := map[string]bool{
		"ArtifactRegistryRule": true,
	}

	tests := []struct {
		name           string
		fixtureName    string
		createObj      map[string]any
		updateObj      map[string]any
		wantErrContain string
	}{
		{
			name:        "mutable spec change succeeds",
			fixtureName: "artifactregistryrule-maximal",
			createObj: map[string]any{
				"apiVersion": "artifactregistry.cnrm.cloud.google.com/v1alpha1",
				"kind":       "ArtifactRegistryRule",
				"metadata":   map[string]any{"name": "test"},
				"spec":       map[string]any{"action": "ALLOW"},
			},
			updateObj: map[string]any{
				"apiVersion": "artifactregistry.cnrm.cloud.google.com/v1alpha1",
				"kind":       "ArtifactRegistryRule",
				"metadata":   map[string]any{"name": "test"},
				"spec":       map[string]any{"action": "DENY"},
			},
		},
		{
			name:        "immutable direct controller rejects update.yaml",
			fixtureName: "translatedataset-minimal",
			createObj: map[string]any{
				"apiVersion": "translate.cnrm.cloud.google.com/v1alpha1",
				"kind":       "TranslateDataset",
				"metadata":   map[string]any{"name": "test"},
				"spec":       map[string]any{"displayName": "ds-1"},
			},
			updateObj: map[string]any{
				"apiVersion": "translate.cnrm.cloud.google.com/v1alpha1",
				"kind":       "TranslateDataset",
				"metadata": map[string]any{
					"name":   "test",
					"labels": map[string]any{"env": "updated"},
				},
				"spec": map[string]any{"displayName": "ds-1"},
			},
			wantErrContain: "omit update.yaml for immutable resources",
		},
		{
			name:        "spec-only greenfield controller rejects metadata-only update.yaml",
			fixtureName: "artifactregistryrule-minimal",
			createObj: map[string]any{
				"apiVersion": "artifactregistry.cnrm.cloud.google.com/v1alpha1",
				"kind":       "ArtifactRegistryRule",
				"metadata":   map[string]any{"name": "test"},
				"spec":       map[string]any{"action": "ALLOW"},
			},
			updateObj: map[string]any{
				"apiVersion": "artifactregistry.cnrm.cloud.google.com/v1alpha1",
				"kind":       "ArtifactRegistryRule",
				"metadata": map[string]any{
					"name":   "test",
					"labels": map[string]any{"env": "updated"},
				},
				"spec": map[string]any{"action": "ALLOW"},
			},
			wantErrContain: "controller only reconciles spec via CompareGreenfieldSpec",
		},
		{
			name:        "identical create and update rejected",
			fixtureName: "some-legacy-fixture",
			createObj: map[string]any{
				"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
				"kind":       "StorageBucket",
				"metadata":   map[string]any{"name": "test"},
				"spec":       map[string]any{"location": "US"},
			},
			updateObj: map[string]any{
				"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
				"kind":       "StorageBucket",
				"metadata":   map[string]any{"name": "test"},
				"spec":       map[string]any{"location": "US"},
			},
			wantErrContain: "does not make any changes to spec, labels, or annotations",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			createUnstruct := &unstructured.Unstructured{Object: tc.createObj}
			updateUnstruct := &unstructured.Unstructured{Object: tc.updateObj}
			err := checkUpdateFixtureDiff(tc.fixtureName, createUnstruct, updateUnstruct, immutableKinds, specOnlyGreenfieldKinds)
			if tc.wantErrContain == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tc.wantErrContain)
			}
			if !strings.Contains(err.Error(), tc.wantErrContain) {
				t.Fatalf("error %q does not contain %q", err.Error(), tc.wantErrContain)
			}
		})
	}
}
