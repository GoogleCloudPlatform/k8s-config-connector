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

package resourceskeleton_test

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/serviceclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceskeleton"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestNewProject(t *testing.T) {
	projectID := "my-project-id"
	u, err := resourceskeleton.NewProject(projectID, testservicemappingloader.New(t))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.GetName() != projectID {
		t.Fatalf("unexpected value for name: got '%v', want '%v'", u.GetName(), projectID)
	}
	expectedKind := "Project"
	if u.GetKind() != expectedKind {
		t.Fatalf("unexpected value for kind: got '%v', want '%v'", u.GetKind(), expectedKind)
	}
	if _, ok := k8s.GetAnnotation(k8s.FolderIDAnnotation, u); !ok {
		t.Fatalf("expected annotations to contain the folder id annotation '%v'", k8s.ProjectIDAnnotation)
	}
}

type URISkeletonTestCase struct {
	ResourceConfigID string
	URI              string
	ExpectedSkeleton *unstructured.Unstructured
}

func TestNewFromURI(t *testing.T) {
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	smLoader := testservicemappingloader.New(t)
	idToRC := make(map[string]v1alpha1.ResourceConfig)
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			id := resourcefixture.GetUniqueResourceConfigID(rc)
			idToRC[id] = rc
		}
	}
	testCases := loadURISkeletonTestCases(t)
	ensureURITCExistsForEachResourceConfig(t, smLoader, testCases)
	for _, tc := range testCases {
		rc := idToRC[tc.ResourceConfigID]
		if !*rc.IDTemplateCanBeUsedToMatchResourceName {
			if tc.ExpectedSkeleton != nil {
				t.Errorf("error skeleton is not null but resource doesn't support NewFromUri '%v': %v", tc.ResourceConfigID, tc.URI)
			}
			continue
		}
		u, err := resourceskeleton.NewFromURI(tc.URI, smLoader, tfProvider)
		if err != nil {
			t.Fatalf("error getting skeleton unstruct from uri '%v': %v", tc.URI, err)
		}
		if !reflect.DeepEqual(tc.ExpectedSkeleton, u) {
			diff := cmp.Diff(tc.ExpectedSkeleton, u)
			t.Fatalf("mismatch between skeletons for resource id '%v', diff:\n%v", tc.ResourceConfigID, diff)
		}
	}
}

type AssetSkeletonTestCase struct {
	ResourceConfigID string                     `json:"resourceConfigId,omitempty"`
	Asset            *asset.Asset               `json:"asset,omitempty"`
	ExpectedSkeleton *unstructured.Unstructured `json:"expectedSkeleton,omitempty"`
}

func TestNewFromAsset(t *testing.T) {
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	smLoader := testservicemappingloader.New(t)
	mockServiceClient := serviceclient.NewMockServiceClient(t)
	testCases := loadAssetSkeletonTestCases(t)
	ensureAssetTCExistsForEachResourceConfig(t, smLoader, testCases)
	for _, tc := range testCases {
		a := tc.Asset
		if a == nil {
			continue
		}
		u, err := resourceskeleton.NewFromAsset(a, smLoader, tfProvider, &mockServiceClient)
		if err != nil {
			t.Fatalf("error getting skeleton unstruct from asset '%v' with type '%v': %v", a.Name, a.AssetType, err)
		}
		if !reflect.DeepEqual(tc.ExpectedSkeleton, u) {
			diff := cmp.Diff(tc.ExpectedSkeleton, u)
			t.Fatalf("mismatch between skeletons for resource id '%v', diff:\n%v", tc.ResourceConfigID, diff)
		}
	}
}

func ensureAssetTCExistsForEachResourceConfig(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, testCases []AssetSkeletonTestCase) {
	t.Helper()
	rcIDToTC := make(map[string]AssetSkeletonTestCase)
	for _, tc := range testCases {
		rcIDToTC[tc.ResourceConfigID] = tc
	}
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			// TODO: remove 'Direct' field from ResourceConfig and remove the if statement.
			// The 'Direct' indicator won't be needed after we finish all the migrations.
			// The 'Direct' indicator is necessary during the migration so
			// that Config Connector uses direct approach to generate CRDs
			// but still allow TF-based controller to reconcile the resource.
			if rc.Direct {
				continue
			}
			if rc.AutoGenerated {
				continue
			}
			id := resourcefixture.GetUniqueResourceConfigID(rc)
			if _, ok := rcIDToTC[id]; ok {
				continue
			}
			t.Fatalf("missing test case for resource config with id '%v' and TF name '%v'", id, rc.Name)
		}
	}
}

func ensureURITCExistsForEachResourceConfig(t *testing.T, smLoader *servicemappingloader.ServiceMappingLoader, testCases []URISkeletonTestCase) {
	t.Helper()
	rcIDToTC := make(map[string]URISkeletonTestCase)
	for _, tc := range testCases {
		rcIDToTC[tc.ResourceConfigID] = tc
	}
	for _, sm := range smLoader.GetServiceMappings() {
		for _, rc := range sm.Spec.Resources {
			// TODO: remove 'Direct' field from ResourceConfig and remove the if statement.
			// The 'Direct' indicator won't be needed after we finish all the migrations.
			// The 'Direct' indicator is necessary during the migration so
			// that Config Connector uses direct approach to generate CRDs
			// but still allow TF-based controller to reconcile the resource.
			if rc.Direct {
				continue
			}
			if rc.AutoGenerated {
				continue
			}
			id := resourcefixture.GetUniqueResourceConfigID(rc)
			if _, ok := rcIDToTC[id]; ok {
				continue
			}
			t.Fatalf("missing test case for resource config with id '%v' and TF name '%v'", id, rc.Name)
		}
	}
}

const (
	assetResourceSkeletonPath = "testdata/asset-skeleton.yaml"
	uriResourceSkeletonPath   = "testdata/uri-skeleton.yaml"
)

func loadURISkeletonTestCases(t *testing.T) []URISkeletonTestCase {
	var value []URISkeletonTestCase
	testyaml.UnmarshalFile(t, uriResourceSkeletonPath, &value)
	return value
}

func loadAssetSkeletonTestCases(t *testing.T) []AssetSkeletonTestCase {
	var value []AssetSkeletonTestCase
	testyaml.UnmarshalFile(t, assetResourceSkeletonPath, &value)
	return value
}
