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
	"fmt"
	"sort"
	"strings"
	"testing"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/stateintospec"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

func TestOutputOnlyFieldsAreUnderObservedState(t *testing.T) {
	crds, err := crdloader.LoadCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatalf("error getting new service mapping loader: %v", err)
	}

	var errs []string

	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			gvk := k8sschema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: version.Name,
				Kind:    crd.Spec.Names.Kind,
			}

			openAPISchema := version.Schema.OpenAPIV3Schema
			prop := findOpenAPIProperty(openAPISchema, "status", "observedState")
			hasObservedState := prop != nil

			// 'status.observedState' should exist for CRDs that
			// (1) have all the output-only fields under 'status.observedState'
			//     instead of 'status', or
			// (2) have observedFields configured in the resource config to
			//     expose selected output 'spec' fields.
			// 'status.observedState' does not exist for CRDs that
			// (1) don't have any output-only fields, or
			// (2) have all the output-only fields under 'status' but don't
			//     have observedFields configured.
			mayHaveObservedState := stateintospec.OutputOnlyFieldsAreUnderObservedState(gvk)

			rcs, err := smLoader.GetResourceConfigs(gvk)
			// Ignore not found error because there are handwritten and
			// DCL-based resources.
			if err != nil && !strings.Contains(err.Error(), "no mapping with name") {
				t.Errorf("error getting resource config(s) for gvk %+v: %v", gvk, err)
			}
			shouldHaveObservedState := false
			for _, rc := range rcs {
				if rc.ObservedFields != nil && len(*rc.ObservedFields) > 0 {
					shouldHaveObservedState = true
				}
			}

			// If shouldHaveObservedState is true, we should expect there is
			// 'status.observedState' in the CRD.
			if shouldHaveObservedState && !hasObservedState {
				errs = append(errs, fmt.Sprintf("crd=%s version=%v: 'status.observedState' doesn't exist, but it should", crd.Name, version.Name))
			}

			// If mayHaveObservedState is true, it means 'status' will only
			// have subfields 'conditions' and 'observedGeneration', and
			// optionally subfield 'observedState'.
			if mayHaveObservedState {
				statusProp := findOpenAPIProperty(openAPISchema, "status")
				if statusProp == nil {
					errs = append(errs, fmt.Sprintf("crd=%s version=%v: 'status' doesn't exist, but it should", crd.Name, version.Name))
				} else {
					requiredFieldsMap := map[string]bool{"observedGeneration": true, "conditions": true}
					optionalFieldsMap := map[string]bool{"observedState": true}

					// for direct resources, we will use the "externalRef" prop under the "status"
					// to track the KCC full resource ID
					optionalFieldsMap["externalRef"] = true

					for k := range statusProp.Properties {
						foundInMaps := false
						if _, ok := requiredFieldsMap[k]; ok {
							foundInMaps = true
							delete(requiredFieldsMap, k)
						}
						if _, ok := optionalFieldsMap[k]; ok {
							foundInMaps = true
							delete(optionalFieldsMap, k)
						}
						if !foundInMaps {
							errs = append(errs, fmt.Sprintf("crd=%s version=%v: CRD has non-boilerplate field '%v' under 'status'", crd.Name, version.Name, k))
						}
					}

					if len(requiredFieldsMap) > 0 {
						// Convert map keys to slice for consistent error message
						missing := make([]string, 0, len(requiredFieldsMap))
						for k := range requiredFieldsMap {
							missing = append(missing, k)
						}
						sort.Strings(missing)
						errs = append(errs, fmt.Sprintf("crd=%s version=%v: CRD doesn't have enough subfields under 'status' field: it should at least have fields %v", crd.Name, version.Name, missing))
					}
				}
			}
		}
	}

	sort.Strings(errs)
	want := strings.Join(errs, "\n")
	test.CompareGoldenFile(t, "testdata/exceptions/observed_state.txt", want)
}

func findOpenAPIProperty(schema *apiextensionsv1.JSONSchemaProps, path ...string) *apiextensionsv1.JSONSchemaProps {
	pos := schema
	for _, k := range path {
		p, found := pos.Properties[k]
		if !found {
			return nil
		}
		pos = &p
	}
	return pos
}
