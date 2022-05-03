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

package fielddesc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/fielddesc"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/testutils"

	"gopkg.in/yaml.v2"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestAllCRDsGetSpecAndStatusDescription(t *testing.T) {
	crds, err := crdloader.LoadCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}
	for _, crd := range crds {
		if crd.Spec.Names.Kind == "ServiceMapping" {
			continue
		}
		fd := fielddesc.GetSpecDescription(&crd)
		expectedType := "object"
		if fd.Type != expectedType {
			t.Fatalf("unexpected type: got '%v', want' %v'", fd.Type, expectedType)
		}
		fd = getStatusDescription(t, &crd)
		if fd.Type != expectedType {
			t.Fatalf("unexpected type: got '%v', want' %v'", fd.Type, expectedType)
		}
	}
}

// note: when updating the schema of the CRDs below these tests will likely fail. You can update the
// expected test data by running the test with the -update flag.
func TestOutputMatches(t *testing.T) {
	testOutputMatches(t, "AccessContextManagerAccessLevel")
	testOutputMatches(t, "BinaryAuthorizationPolicy")
	testOutputMatches(t, "PubSubSubscription")
}

func testOutputMatches(t *testing.T, resourceKind string) {
	crd, err := crdloader.GetCRDForKind(resourceKind)
	if err != nil {
		t.Fatalf("error getting crd '%v': %v", resourceKind, err)
	}
	fd := fielddesc.GetSpecDescription(crd)
	bytes := fieldDescToYAML(t, fd)
	testutils.VerifyContentsMatch(t, bytes, fmt.Sprintf("testdata/%v-spec.golden.yaml", strings.ToLower(resourceKind)))
	fd = getStatusDescription(t, crd)
	bytes = fieldDescToYAML(t, fd)
	testutils.VerifyContentsMatch(t, bytes, fmt.Sprintf("testdata/%v-status.golden.yaml", strings.ToLower(resourceKind)))
}

func getStatusDescription(t *testing.T, crd *apiextensions.CustomResourceDefinition) fielddesc.FieldDescription {
	fd, err := fielddesc.GetStatusDescription(crd)
	if err != nil {
		t.Fatalf("error getting status description")
	}
	return fd
}

func fieldDescToYAML(t *testing.T, fieldDesc fielddesc.FieldDescription) []byte {
	bytes, err := yaml.Marshal(fieldDesc)
	if err != nil {
		t.Fatalf("error marshalling to yaml: %v", err)
	}
	return bytes
}
