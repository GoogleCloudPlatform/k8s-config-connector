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
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"

	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

func disableAdditionalProperties(s *spec.Schema) {
	if s == nil {
		return
	}
	if s.Type.Contains("object") || len(s.Properties) > 0 {
		// Only set additionalProperties=false if there is a defined set of properties.
		if len(s.Properties) > 0 && s.AdditionalProperties == nil {
			s.AdditionalProperties = &spec.SchemaOrBool{Allows: false}
		}
		for k, v := range s.Properties {
			disableAdditionalProperties(&v)
			s.Properties[k] = v
		}
	}
	if s.Items != nil && s.Items.Schema != nil {
		disableAdditionalProperties(s.Items.Schema)
	}
}

// Run this test with WRITE_GOLDEN_OUTPUT=1 set to update the exceptions list.
func TestYAMLValidatesAgainstCRDSchema(t *testing.T) {
	t.Parallel()

	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading CRDs: %v", err)
	}

	validators := make(map[schema.GroupVersionKind]validation.SchemaValidator)

	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			gvk := schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: version.Name,
				Kind:    crd.Spec.Names.Kind,
			}

			if version.Schema == nil || version.Schema.OpenAPIV3Schema == nil {
				continue
			}

			b, err := json.Marshal(version.Schema.OpenAPIV3Schema)
			if err != nil {
				t.Fatalf("error marshalling schema for %v: %v", gvk, err)
			}
			var s spec.Schema
			if err := json.Unmarshal(b, &s); err != nil {
				t.Fatalf("error unmarshalling schema for %v: %v", gvk, err)
			}

			// Enforce strict schema validation by rejecting unspecified properties
			if specSchema, ok := s.Properties["spec"]; ok {
				disableAdditionalProperties(&specSchema)
				s.Properties["spec"] = specSchema
			}

			validator := validation.NewSchemaValidatorFromOpenAPI(&s)
			validators[gvk] = validator
		}
	}

	fixtures := resourcefixture.Load(t)
	var errors []string

	// Replace all test variables like ${PROJECT_ID} with a mock string
	// to avoid schema validation errors (e.g., getting "null" or empty string when
	// environment variables are not set during apichecks).
	testVarRegex := regexp.MustCompile("\\$\\{([a-zA-Z0-9_]+)\\}")

	for _, fixture := range fixtures {
		var yamlBytes [][]byte
		if len(fixture.Create) > 0 {
			yamlBytes = append(yamlBytes, testyaml.SplitYAML(t, fixture.Create)...)
		}
		if len(fixture.Update) > 0 {
			yamlBytes = append(yamlBytes, testyaml.SplitYAML(t, fixture.Update)...)
		}
		if len(fixture.Dependencies) > 0 {
			yamlBytes = append(yamlBytes, testyaml.SplitYAML(t, fixture.Dependencies)...)
		}

		for _, b := range yamlBytes {
			if len(b) == 0 {
				continue
			}

			b = testVarRegex.ReplaceAllFunc(b, func(match []byte) []byte {
				matches := testVarRegex.FindSubmatch(match)
				if len(matches) > 1 {
					varName := strings.ToLower(string(matches[1]))
					varName = strings.ReplaceAll(varName, "_", "")
					return []byte("mock-" + varName)
				}
				return match
			})

			unstruct := ToUnstruct(t, b)
			if unstruct == nil || unstruct.Object == nil {
				continue
			}
			gvk := unstruct.GroupVersionKind()
			validator, ok := validators[gvk]
			if !ok {
				// We don't have a validator for this GVK (e.g. core k8s resources like Secret, ConfigMap)
				continue
			}

			errs := validation.ValidateCustomResource(field.NewPath(""), unstruct.Object, validator)
			for _, err := range errs {
				errors = append(errors, fmt.Sprintf("[schema_violation] file=%s gvk=%v: %v", fixture.TestKey, gvk, err))
			}
		}
	}

	// deduplicate and sort errors
	errMap := make(map[string]bool)
	for _, err := range errors {
		errMap[err] = true
	}
	var uniqueErrors []string
	for err := range errMap {
		uniqueErrors = append(uniqueErrors, err)
	}
	sort.Strings(uniqueErrors)

	want := strings.Join(uniqueErrors, "\n")
	test.CompareGoldenFile(t, "testdata/exceptions/schema_violations.txt", want)
}
