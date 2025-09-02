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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestResourceContents(t *testing.T) {
	fixtures := resourcefixture.Load(t)
	for _, f := range fixtures {
		if f.Update == nil {
			continue
		}
		createUnstruct := test.ToUnstruct(t, f.Create)
		updateUnstruct := test.ToUnstruct(t, f.Update)
		testCreateAndUpdateUnstructMatchingProperties(t, f, createUnstruct, updateUnstruct)
		dependencies := make([]*unstructured.Unstructured, 0)
		testAllNamesAreUnique(t, f, append(dependencies, createUnstruct)...)
	}
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
