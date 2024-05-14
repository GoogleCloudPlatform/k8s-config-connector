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

package crdtemplate_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	crdtemplate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/template"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestAllCRDsShouldConvertToYAML(t *testing.T) {
	crds, err := crdloader.LoadCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}
	for _, crd := range crds {
		specToYAML(t, &crd)
		statusToYAML(t, &crd)
	}
}

func TestSpecAndStatusToYAML(t *testing.T) {
	// when adding a new type or updating the test data file run this test with the 'WRITE_GOLDEN_OUTPUT' env var to update the 'golden files'
	testToYAML(t, "ComputeInstance")
	testToYAML(t, "ComputeBackendService")
	testToYAML(t, "PubSubTopic")
	testToYAML(t, "PubSubSubscription")
	testToYAML(t, "SpannerDatabase")
}

func TestAllLoadedCRDHaveManagedByKCCLabel(t *testing.T) {
	crds, err := crdloader.LoadCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	for _, crd := range crds {
		if _, ok := crd.Labels[crdgeneration.ManagedByKCCLabel]; !ok {
			t.Errorf("%v CRD missing %v label", crd.Kind, crdgeneration.ManagedByKCCLabel)
		}
	}
}

func testToYAML(t *testing.T, resourceKind string) {
	crd := getCRDForKind(t, resourceKind)
	bytes := specToYAML(t, crd)
	test.CompareGoldenFile(t, fmt.Sprintf("testdata/%v-spec.yaml.golden", strings.ToLower(resourceKind)), string(bytes), test.IgnoreLeadingComments)
	bytes = statusToYAML(t, crd)
	test.CompareGoldenFile(t, fmt.Sprintf("testdata/%v-status.yaml.golden", strings.ToLower(resourceKind)), string(bytes), test.IgnoreLeadingComments)
}

func getCRDForKind(t *testing.T, kind string) *apiextensions.CustomResourceDefinition {
	crd, err := crdloader.GetCRDForKind(kind)
	if err != nil {
		t.Fatalf("error getting crd: %v", err)
	}
	return crd
}

func specToYAML(t *testing.T, crd *apiextensions.CustomResourceDefinition) []byte {
	bytes, err := crdtemplate.SpecToYAML(crd)
	if err != nil {
		t.Fatalf("error converting crd spec to YAML: %v", err)
	}
	return bytes
}

func statusToYAML(t *testing.T, crd *apiextensions.CustomResourceDefinition) []byte {
	bytes, err := crdtemplate.StatusToYAML(crd)
	if err != nil {
		t.Fatalf("error converting crd spec to YAML: %v", err)
	}
	return bytes
}
