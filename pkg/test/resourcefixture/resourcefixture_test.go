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
	"strings"
	"testing"

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testservicemapping "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemapping"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ensure that every resource type is accounted for in the returned cover set
func TestGetBasicTypeSetCover(t *testing.T) {
	smLoader := testservicemappingloader.New(t)
	serviceMetadataLoader := dclmetadata.New()
	basicFixtures := resourcefixture.GetBasicTypeSetCover(t)
	resourceConfigIds := getResourceConfigIDSet(t, smLoader)
	notFoundResourceConfigIds := make(map[string]bool, len(resourceConfigIds))
	for k, v := range resourceConfigIds {
		notFoundResourceConfigIds[k] = v
	}
	for _, f := range basicFixtures {
		verifyResourceConfigIDAndDeleteFromNotFound(t, smLoader, serviceMetadataLoader, f.Create, resourceConfigIds, notFoundResourceConfigIds)
		for _, d := range testyaml.SplitYAML(t, f.Dependencies) {
			verifyResourceConfigIDAndDeleteFromNotFound(t, smLoader, serviceMetadataLoader, d, resourceConfigIds, notFoundResourceConfigIds)
		}
	}
	if len(notFoundResourceConfigIds) > 0 {
		resourceConfigIds := make([]string, 0, len(notFoundResourceConfigIds))
		for k := range notFoundResourceConfigIds {
			resourceConfigIds = append(resourceConfigIds, k)
		}
		t.Fatalf("resources covering the config(s) of '%v' are missing from the basic resource set cover", strings.Join(resourceConfigIds, ", "))
	}
}

func verifyResourceConfigIDAndDeleteFromNotFound(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, yamlBytes []byte, resourceConfigIds, notFound map[string]bool) {
	t.Helper()
	u := test.ToUnstruct(t, yamlBytes)
	if !resourcefixture.ShouldHaveResourceConfig(u, serviceMetadataLoader) {
		return
	}
	rc := testservicemapping.GetResourceConfig(t, smLoader, u)
	rcID := resourcefixture.GetUniqueResourceConfigID(*rc)
	// if this occurs then something is wrong with the way this test sets itself up
	if _, ok := resourceConfigIds[rcID]; !ok {
		t.Fatalf("unexpected missing resource name: %v", rcID)
	}
	delete(notFound, rcID)
}

func getResourceConfigIDSet(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader) map[string]bool {
	sms := smLoader.GetServiceMappings()
	resourceConfigIds := make(map[string]bool, 0)
	for _, sm := range sms {
		for _, rc := range sm.Spec.Resources {
			// TODO: remove 'Direct' field from ResourceConfig and remove the if statement.
			// The 'Direct' indicator won't be needed after we finish all the migrations.
			// The 'Direct' indicator is necessary during the migration so
			// that Config Connector uses direct approach to generate CRDs
			// but still allow TF-based controller to reconcile the resource.
			if rc.Direct {
				continue
			}
			// No need to cover auto-generated v1alpha1 CRDs when calculating
			// set cover.
			if sm.GetVersionFor(&rc) == k8s.KCCAPIVersionV1Alpha1 {
				continue
			}
			id := resourcefixture.GetUniqueResourceConfigID(rc)
			if _, ok := resourceConfigIds[id]; ok {
				t.Fatalf("unexpected scenario: resource config '%v' and another resource config resolve to the same id '%v'",
					rc.Name, id)
			}
			resourceConfigIds[id] = true
		}
	}
	// AccessContextManagerAccessPolicy is special in that there can only be a single one for the entire org and
	// for that reason it is OK if it is not included in the set of all resource types
	delete(resourceConfigIds, "AccessContextManagerAccessPolicy")
	return resourceConfigIds
}

func TestNamesAreUniqueWithinTheSameTestType(t *testing.T) {
	// all resource fixtures should have unique names
	fixtures := resourcefixture.Load(t)
	namesPerTestType := make(map[resourcefixture.TestType]map[string]bool)
	for _, f := range fixtures {
		if _, ok := namesPerTestType[f.Type]; !ok {
			namesPerTestType[f.Type] = make(map[string]bool)
		}
		names := namesPerTestType[f.Type]
		if _, ok := names[f.Name]; ok {
			t.Fatalf("cannot have two fixtures with the same type '%v' and name '%v'", f.Type, f.Name)
		}
		names[f.Name] = true
	}
}

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
