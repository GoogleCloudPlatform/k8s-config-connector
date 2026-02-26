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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/fielddesc"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestAllCRDsGetSpecAndStatusDescription(t *testing.T) {
	crds, err := crdloader.LoadCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}
	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			fd := fielddesc.GetSpecDescription(&crd, version.Name)
			expectedType := "object"
			if fd.Type != expectedType {
				t.Fatalf("unexpected type: got '%v', want' %v'", fd.Type, expectedType)
			}
			fd = getStatusDescription(t, &crd, version.Name)
			if fd.Type != expectedType {
				t.Fatalf("unexpected type: got '%v', want' %v'", fd.Type, expectedType)
			}
		}
	}
}

func getStatusDescription(t *testing.T, crd *apiextensions.CustomResourceDefinition, version string) fielddesc.FieldDescription {
	fd, err := fielddesc.GetStatusDescription(crd, version)
	if err != nil {
		t.Fatalf("error getting status description")
	}
	return fd
}
