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

package lint

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/config/tests/samples/create"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
)

// Look at fields of v1beta1+ CRDs and make sure that all settable fields are set
// by at least one CR across any samples or fixtures for a given CRD.
func TestFieldsSet(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	betaGVKsToFieldSet := make(map[schema.GroupVersionKind]map[string]bool)

	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			// ignore v1alpha1s for now
			if version.Name == "v1alpha1" {
				continue
			}
			gvk := schema.GroupVersionKind{
				Group: crd.Spec.Group, Version: version.Name, Kind: crd.Spec.Names.Kind,
			}

			// field set will describe every "spec.foo.bar" field
			fieldSet := make(map[string]bool)

			visitCRDVersion(version, func(field *CRDField) {
				fieldPath := field.FieldPath

				// Only consider spec
				if !strings.HasPrefix(fieldPath, ".spec.") {
					return
				}

				switch field.props.Type {
				case "string", "boolean", "integer", "number":
					// only mark "leaf" paths
					fieldSet[fieldPath] = false // initialize as not found
				}
			})

			betaGVKsToFieldSet[gvk] = fieldSet
		}
	}

	project := testgcp.GCPProject{
		ProjectID:     "tests-testnames",
		ProjectNumber: 1234567890,
	}
	for _, sample := range create.LoadAllSamples(t, project) {
		for _, resource := range sample.Resources {
			gvk := resource.GroupVersionKind()

			if gvk.Version == "v1alpha1" {
				continue
			}

			setFields := make(map[string]struct{})

			//t.Logf("Processing sample resource %s of gvk %s", resource.GetName(), gvk)

			spec, found, err := unstructured.NestedMap(resource.Object, "spec")
			if err != nil {
				t.Error("error processing spec", err)
			}
			if !found {
				t.Errorf("did not find spec for resource %s of gvk %s", resource.GetName(), gvk)
			}

			visitUnstructFields(spec, ".spec", func(fieldPath string) {
				// found a field!
				setFields[fieldPath] = struct{}{}
			})

			fieldsForGVK := betaGVKsToFieldSet[gvk]
			for field := range setFields {
				if _, ok := fieldsForGVK[field]; !ok {
					//t.Logf("field %s not found for gvk %s's fields %+v", field, gvk, fieldsForGVK)
					t.Logf("field %s not found in schema for gvk %s", field, gvk)
				} else {
					fieldsForGVK[field] = true
				}
			}
			betaGVKsToFieldSet[gvk] = fieldsForGVK
		}
	}

	for _, fixture := range resourcefixture.Load(t) {
		// just look at primary resources
		primaryResource := bytesToUnstructured(t, fixture.Create, testvariable.NewUniqueID(), project)

		gvk := primaryResource.GroupVersionKind()
		setFields := make(map[string]struct{})

		if gvk.Version == "v1alpha1" {
			continue
		}

		//t.Logf("Processing fixture resource %s of gvk %s", resource.GetName(), gvk)

		spec, found, err := unstructured.NestedMap(primaryResource.Object, "spec")
		if err != nil {
			t.Error("error processing spec", err)
		}
		if !found {
			t.Errorf("did not find spec for resource %s of gvk %s", primaryResource.GetName(), gvk)
		}

		visitUnstructFields(spec, ".spec", func(fieldPath string) {
			// found a field!
			setFields[fieldPath] = struct{}{}
		})

		fieldsForGVK := betaGVKsToFieldSet[gvk]
		for field := range setFields {
			if _, ok := fieldsForGVK[field]; !ok {
				//t.Logf("field %s not found for gvk %s's fields %+v", field, gvk, fieldsForGVK)
				t.Logf("field %s not found in schema for gvk %s", field, gvk)
			} else {
				fieldsForGVK[field] = true
			}
		}
		betaGVKsToFieldSet[gvk] = fieldsForGVK
	}

	// todo acpana temporary while ironing out kinks w test
	kindsToTest := map[string]struct{}{
		"AccessContextManagerAccessLevel": {},
	}
	var errs []string
	for gvk, fields := range betaGVKsToFieldSet {
		if _, ok := kindsToTest[gvk.Kind]; !ok {
			continue
		}

		for fieldKey := range fields {
			if !fields[fieldKey] {
				//t.Logf("field %s not seen in samples or fixtures for gvk %s", fieldKey, gvk)
				errs = append(errs, fmt.Sprintf("[field not set] gvk=%s: field %s not set in a sample or fixture resource", gvk, fieldKey))
			}
		}
	}

	sort.Strings(errs)
	want := strings.Join(errs, "\n")
	test.CompareGoldenFile(t, "testdata/exceptions/fields_not_set.txt", want)
}

func bytesToUnstructured(t *testing.T, bytes []byte, testID string, project testgcp.GCPProject) *unstructured.Unstructured {
	t.Helper()
	updatedBytes := testcontroller.ReplaceTestVars(t, bytes, testID, project)
	return test.ToUnstructWithNamespace(t, updatedBytes, testID)
}

// Looks for fields that looks like refs, but are not
func TestMissingRefs(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	var errs []string
	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			visitCRDVersion(version, func(field *CRDField) {
				fieldPath := field.FieldPath

				// Only consider spec
				if strings.HasPrefix(fieldPath, ".status.") {
					return
				}

				// Check if this is already a ref
				if strings.HasSuffix(fieldPath, "Ref") {
					return
				}
				if strings.HasSuffix(fieldPath, "Refs[]") {
					return
				}
				if strings.HasSuffix(fieldPath, "Ref.external") {
					return
				}
				if strings.HasSuffix(fieldPath, "Ref.name") {
					return
				}

				isRef := false
				desc := field.props.Description
				// Heuristic: look for descriptions like "should be of the form projects/{projectID}/locations/{location}/bars/{name}"
				if strings.Contains(desc, " projects/") {
					isRef = true
				}

				if isRef {
					errs = append(errs, fmt.Sprintf("[refs] crd=%s version=%v: field %q should be a reference", crd.Name, version.Name, fieldPath))
				}
			})
		}
	}

	sort.Strings(errs)

	want := strings.Join(errs, "\n")

	test.CompareGoldenFile(t, "testdata/exceptions/missingrefs.txt", want)
}

func TestCRDsDoNotHaveFooUrlRef(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			visitCRDVersion(version, func(field *CRDField) {
				fieldPath := field.FieldPath
				lower := strings.ToLower(fieldPath)
				if strings.HasSuffix(lower, "urlref") && !strings.HasSuffix(lower, ".urlref") {
					// Prefer network_ref to network_url_ref
					// While we allow url_ref, network_url_ref is odd;
					// _url indicates the data type / representation of the field,
					// and we don't want two "types" in our field name.
					t.Errorf("invalid field name %q in %q; prefer fooRef to fooUrlRef", fieldPath, crd.Name)
				}
			})
		}
	}
}

// Enforces acronym capitalization on CRDs
// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#naming-conventions
// All letters in the acronym should have the same case, using the appropriate case for the situation.
// For example, at the beginning of a field name, the acronym should be all lowercase, such as "httpGet".
// Where used as a constant, all letters should be uppercase, such as "TCP" or "UDP".
func TestCRDsAcronyms(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	var errs []string
	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			visitCRDVersion(version, func(field *CRDField) {
				fieldPath := field.FieldPath
				tokens := splitCamelCase(fieldPath)

				// Special cases for common acronyms
				for i, token := range tokens {
					isAcronym := false

					switch strings.ToLower(token) {
					case "http", "https", "ssh", "tls", "udp", "tcp":
						isAcronym = true
					case "api":
						isAcronym = true

					case "ipv4", "ipv6", "ip", "cidr", "bgp":
						isAcronym = true

					case "id":
						isAcronym = true

					case "url":
						isAcronym = true
					case "cdn":
						isAcronym = true
					case "nat":
						isAcronym = true
					case "x509":
						isAcronym = true
					case "sso":
						isAcronym = true
					case "oauth2", "oidc":
						isAcronym = true
					case "iap":
						isAcronym = true
					case "os":
						isAcronym = true
					}

					// TODO: Ips, Cidrs

					// TODO: Src / Dest

					if isAcronym {
						if i == 0 {
							tokens[i] = strings.ToLower(token)
						} else {
							tokens[i] = strings.ToUpper(token)
						}
					}
				}

				corrected := strings.Join(tokens, "")

				if corrected != fieldPath {
					errs = append(errs, fmt.Sprintf("[acronyms] crd=%s version=%v: field %q should be %q", crd.Name, version.Name, fieldPath, corrected))
				}
			})
		}
	}

	sort.Strings(errs)

	want := strings.Join(errs, "\n")

	test.CompareGoldenFile(t, "testdata/exceptions/acronyms.txt", want)
}

// splitCamelCase splits the string on capital letters, so camelCase => []string{"camel", "Case"}
func splitCamelCase(s string) []string {
	var tokens []string

	var token string
	for _, r := range s {
		if unicode.IsUpper(r) {
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
		}
		token += string(r)
	}
	if token != "" {
		tokens = append(tokens, token)
	}
	return tokens
}

type CRDField struct {
	FieldPath string
	props     *apiextensions.JSONSchemaProps
}

func visitCRDVersion(version apiextensions.CustomResourceDefinitionVersion, callback func(crdField *CRDField)) {
	visitProps(version.Schema.OpenAPIV3Schema, "", callback)
}

func visitProps(props *apiextensions.JSONSchemaProps, fieldPath string, callback func(crdField *CRDField)) {
	callback(&CRDField{
		FieldPath: fieldPath,
		props:     props,
	})

	switch props.Type {
	case "object":
		for k := range props.Properties {
			child := props.Properties[k]
			visitProps(&child, fieldPath+"."+k, callback)
		}

	case "array":
		if props.Items != nil {
			for _, child := range props.Items.JSONSchemas {
				visitProps(&child, fieldPath+"[]", callback)
			}
			if props.Items.Schema != nil {
				visitProps(props.Items.Schema, fieldPath+"[]", callback)
			}
		}

	case "string", "boolean", "integer", "number":
		// No child properties
	default:
		klog.Fatalf("unhandled props.Type %q in %+v", props.Type, props)
	}
}

func visitUnstructFields(fields interface{}, fieldPath string, callback func(field string)) {
	switch actual := fields.(type) {
	case nil:
		// nothing to do
	case []interface{}:
		for _, k := range actual {
			switch k.(type) {
			case string, bool, int, int32, int64, float64:
				// this is a terminal stage
				callback(fieldPath)
				continue
			}
			visitUnstructFields(k, fieldPath+"[]", callback)
		}
	case map[string]interface{}:
		for k, v := range actual {
			switch v.(type) {
			case string, bool, int, int32, int64, float64:
				// this is a terminal stage
				callback(fieldPath + "." + k)

				continue
			}
			visitUnstructFields(v, fieldPath+"."+k, callback)
		}
	case string, bool, int, int32, int64, float64:
		callback(fieldPath)
	default:
		klog.Fatalf("unhandled value.Type %T in %+v", fields, fields)
	}
}

func TestCRDCamelCase(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}
	var errs []string
	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			visitCRDVersion(version, func(field *CRDField) {
				fieldPath := field.FieldPath
				first := func() int32 {
					tokens := strings.Split(fieldPath, ".")
					// Only check the last token to avoid duplication.
					for _, first := range tokens[len(tokens)-1] {
						return first
					}
					return 0
				}()
				if unicode.IsUpper(first) {
					errs = append(errs, fmt.Sprintf("[jsonNaming] crd=%s version=%v: field %q should use camel case", crd.Name, version.Name, field.FieldPath))
				}
			})
		}
	}
	sort.Strings(errs)
	if len(errs) != 0 {
		t.Fatal(errs)
	}
}
