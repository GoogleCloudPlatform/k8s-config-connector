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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

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
				if strings.HasSuffix(fieldPath, "Refs[]") || strings.HasSuffix(fieldPath, "Refs") {
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

// Looks for fields that looks like refs, but are in the status.
// These fields should not be refs, they should be "external style" links.
func TestNoRefsInStatus(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	var errs []string
	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			visitCRDVersion(version, func(field *CRDField) {
				fieldPath := field.FieldPath

				// Only consider status
				if !strings.HasPrefix(fieldPath, ".status.") {
					return
				}

				// Well-known exception
				if fieldPath == ".status.externalRef" {
					return
				}

				// Check if this is named like a ref
				isRef := false
				if strings.HasSuffix(fieldPath, "Ref") {
					isRef = true
				}
				if strings.HasSuffix(fieldPath, "Refs[]") || strings.HasSuffix(fieldPath, "Refs") {
					isRef = true
				}

				if isRef {
					errs = append(errs, fmt.Sprintf("[no_refs_in_status] crd=%s version=%v: reference field %q should not be in status", crd.Name, version.Name, fieldPath))
				}
			})
		}
	}

	sort.Strings(errs)

	want := strings.Join(errs, "\n")

	test.CompareGoldenFile(t, "testdata/exceptions/no_refs_in_status.txt", want)
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
					case "iam":
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

// Avoid passing sensitive data as plain text in the CRD
func TestNoSensitiveField(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	var errs []string

	sensitiveKeywords := []string{
		"password",
	}
	for _, crd := range crds {

		for _, version := range crd.Spec.Versions {
			totalPaths := sets.Set[string]{}
			skepticalFieldPaths := sets.Set[string]{}
			visitCRDVersion(version, func(field *CRDField) {
				fieldPath := field.FieldPath
				isSensitiveSkeptical := false
				field.FieldPath = strings.ToLower(field.FieldPath)
				for _, sensitiveWord := range sensitiveKeywords {
					if strings.HasSuffix(field.FieldPath, sensitiveWord) {
						isSensitiveSkeptical = true
						break
					}
				}
				if isSensitiveSkeptical {
					skepticalFieldPaths.Insert(fieldPath)
				}
				totalPaths.Insert(fieldPath)
			})
			for skeptical := range skepticalFieldPaths {
				if totalPaths.Has(skeptical + ".valueFrom.secretKeyRef.key") {
					continue
				}
				if totalPaths.Has(skeptical + ".secretRef.name") {
					continue
				}
				errs = append(errs, fmt.Sprintf("crd=%s version=%v: field %q is sensitive data, should use secretRef", crd.Name, version.Name, skeptical))
			}

		}
	}

	sort.Strings(errs)
	want := strings.Join(errs, "\n")
	test.CompareGoldenFile(t, "testdata/exceptions/sensitive.txt", want)
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
