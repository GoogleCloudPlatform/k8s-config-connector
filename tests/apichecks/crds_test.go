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
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
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
				if strings.HasSuffix(fieldPath, "Refs[].external") {
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
				if strings.Contains(desc, "projects/{") {
					isRef = true
				}
				if strings.Contains(desc, "locations/{") {
					isRef = true
				}
				if strings.Contains(desc, "zones/{") {
					isRef = true
				}
				if strings.Contains(desc, "regions/{") {
					isRef = true
				}
				if strings.Contains(desc, "organizations/{") {
					isRef = true
				}
				if strings.Contains(desc, "folders/{") {
					isRef = true
				}

				if strings.HasSuffix(fieldPath, "erviceAccount") {
					isRef = true
				}
				// TODO: how to detect KMS Key

				if isRef {
					// We don't require refs for zones or regions, nor for instanceTypes
					switch {
					case strings.HasSuffix(fieldPath, ".zone"):
						// ok
					case strings.HasSuffix(fieldPath, ".location"):
						// ok
					case strings.HasSuffix(fieldPath, ".machineType"):
						// ok
					case strings.HasSuffix(fieldPath, ".acceleratorType"):
						// ok
					default:
						errs = append(errs, fmt.Sprintf("[refs] crd=%s version=%v: field %q should be a reference", crd.Name, version.Name, fieldPath))

					}
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

				for i, token := range tokens {
					var singular, pluralSuffix string

					if strings.HasSuffix(token, "ies") {
						singular = token[:len(token)-3] + "y"
						pluralSuffix = "ies"
					} else if strings.HasSuffix(token, "es") {
						singular = token[:len(token)-2]
						pluralSuffix = "es"
					} else if strings.HasSuffix(token, "s") {
						singular = token   // or token[:len(token)-1]
						pluralSuffix = "s" // maybe
					} else {
						singular = token
						pluralSuffix = ""
					}

					for _, acronym := range codegen.Acronyms {
						if pluralSuffix == "s" {
							if strings.EqualFold(acronym, singular) {
								pluralSuffix = ""
							} else if !strings.EqualFold(acronym, singular[:len(singular)-1]) {
								continue
							}
						} else {
							if !strings.EqualFold(acronym, singular) {
								continue
							}
						}

						switch pluralSuffix {
						case "ies": // y
							tokens[i] = acronym[:len(acronym)-1] + "ies"
						case "es":
							tokens[i] = acronym + "es"
						case "s":
							tokens[i] = acronym + "s"
						case "":
							tokens[i] = acronym
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

	// Add handling for google.protobuf.Value
	case "string", "boolean", "integer", "number", "":
		// No child properties
	default:
		// if preserveUnknownFields is true, we don't want to check the type
		// For recursive types, we don't want to recurse into schemaless fields
		if props.XPreserveUnknownFields != nil && *props.XPreserveUnknownFields {
			// We don't want to recurse into schemaless fields
			return
		}
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

func TestCRDShortNames(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading CRDs: %v", err)
	}

	var errs []string
	for _, crd := range crds {
		if len(crd.Spec.Names.ShortNames) == 0 {
			errs = append(errs, fmt.Sprintf("[shortnames] crd=%s: missing shortnames", crd.Name))
		}
		for _, sn := range crd.Spec.Names.ShortNames {
			if !strings.HasPrefix(sn, "gcp") {
				errs = append(errs, fmt.Sprintf("[shortnames] crd=%s: shortname %q does not start with gcp", crd.Name, sn))
			}
		}
	}

	sort.Strings(errs)
	want := strings.Join(errs, "\n")
	test.CompareGoldenFile(t, "testdata/exceptions/shortnames.txt", want)
}

// Run this test with WRITE_GOLDEN_OUTPUT set to update the exceptions list.
func TestCRDFieldPresenceInUnstructured(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading CRDs: %v", err)
	}

	unstructs := loadUnstructs(t)
	outputOnlySpecFields, err := loadOutputOnlySpecFields()
	if err != nil {
		t.Fatalf("error loading output-only spec fields from file: %v", err)
	}

	var errs []string
	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {

			if version.Name == "v1alpha1" {
				continue
			}

			kind := crd.Spec.Names.Kind
			visitCRDVersion(version, func(field *CRDField) {
				fieldPath := field.FieldPath
				// Only consider fields under `spec`
				if !strings.HasPrefix(fieldPath, ".spec.") {
					return
				}

				// skip the resource id field
				if strings.HasSuffix(fieldPath, ".resourceID") {
					return
				}

				// Check for "Ref" fields
				if strings.HasSuffix(fieldPath, "Ref") {
					hasExternal := false
					hasName := false

					// Check for specific related fields
					for _, obj := range unstructs {
						if obj.GetKind() != kind {
							continue
						}
						if hasField(obj.Object, fieldPath+".external") {
							hasExternal = true
						}
						if hasField(obj.Object, fieldPath+".name") {
							hasName = true
						}
					}

					// Only report an error if neither external nor name is set
					if !hasExternal && !hasName {
						errs = append(errs, fmt.Sprintf("[missing_field] crd=%s version=%v: field %q is not set; neither 'external' nor 'name' are set", crd.Name, version.Name, fieldPath))
					}
					return
				}

				// Skip non-terminal fields (fields with children or slices)
				if field.props != nil {
					if len(field.props.Properties) > 0 || field.props.Type == "object" {
						return
					}
					if field.props.Type == "array" && field.props.Items != nil {
						return // Skip the array itself; focus on its elements
					}
				}

				// Any XYZRef field was already handled and handling the children will just double count
				if strings.Contains(fieldPath, "Ref") {
					return
				}

				// Check if field exists in any unstructured object
				missing := true
				for _, obj := range unstructs {
					if obj.GetKind() != kind {
						continue
					}
					if hasField(obj.Object, fieldPath) {
						missing = false
						break
					}
				}

				// Exclude output-only spec fields.
				oosfLine := fmt.Sprintf("[output_only_spec_field] crd=%s version=%v: field %q is not set in unstructured objects", crd.Name, version.Name, fieldPath)
				if _, ok := outputOnlySpecFields[oosfLine]; ok {
					missing = false
				}

				if missing {
					errs = append(errs, fmt.Sprintf("[missing_field] crd=%s version=%v: field %q is not set in unstructured objects", crd.Name, version.Name, fieldPath))
				}
			})
		}
	}

	sort.Strings(errs)
	want := strings.Join(errs, "\n")
	test.CompareGoldenFile(t, "testdata/exceptions/missingfields.txt", want)
}

func loadOutputOnlySpecFields() (map[string]bool, error) {
	file, err := os.Open("testdata/exceptions/outputonlyspecfields.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	outputOnlySpecFields := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		outputOnlySpecFields[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return outputOnlySpecFields, nil
}

func loadUnstructs(t *testing.T) []*unstructured.Unstructured {
	t.Helper()
	unstructs := []*unstructured.Unstructured{}
	fixtures := resourcefixture.Load(t)

	for _, fixture := range fixtures {
		fixture := fixture
		createResource := bytesToUnstructured(t, fixture.Create)
		updateResource := bytesToUnstructured(t, fixture.Update)

		unstructs = append(unstructs, createResource, updateResource)
	}

	return unstructs
}

var (
	testID      = testvariable.NewUniqueID()
	testProject = testgcp.GCPProject{ProjectID: "test-skip", ProjectNumber: 123456789}
)

func bytesToUnstructured(t *testing.T, bytes []byte) *unstructured.Unstructured {
	t.Helper()

	updatedBytes := testcontroller.ReplaceTestVars(t, bytes, testID, testProject)
	return ToUnstruct(t, updatedBytes)
}

// hasField checks if an unstructured object contains the given field path.
// For list fields (indicated by [] in the path), it checks if any item in the list
// contains the specified field. If the path ends with [], checks if the field exists
// and is a non-empty list.
func hasField(obj map[string]interface{}, fieldPath string) bool {
	parts := strings.Split(strings.TrimPrefix(fieldPath, "."), ".")
	current := obj

	for i, part := range parts {
		if strings.HasSuffix(part, "[]") {
			listName := strings.TrimSuffix(part, "[]")
			if next, ok := current[listName]; ok {
				if items, ok := next.([]interface{}); ok {
					// 1. If this is the last part, return true if the list exists
					// For example, ".spec.automatedBackupPolicy.weeklySchedule.daysOfWeek[]"
					if i == len(parts)-1 {
						return true
					}
					// 2. Otherwise check remaining path in each item
					// For example, ".spec.automatedBackupPolicy.weeklySchedule.startTimes[].hours"
					remainingPath := strings.Join(parts[i+1:], ".")
					for _, item := range items {
						if itemMap, ok := item.(map[string]interface{}); ok {
							if hasField(itemMap, remainingPath) {
								return true // found the field in one of the items, we can stop searching
							}
						}
					}
				}
			}
			return false
		}

		if next, ok := current[part]; ok {
			if nextMap, ok := next.(map[string]interface{}); ok {
				current = nextMap
			} else {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func ToUnstruct(t *testing.T, bytes []byte) *unstructured.Unstructured {
	t.Helper()

	var obj map[string]interface{}
	err := yaml.Unmarshal(bytes, &obj)
	if err != nil {
		t.Errorf("error unmarshalling bytes %s to unstruct: %v", string(bytes), err)
	}

	return &unstructured.Unstructured{Object: obj}
}

// TestCRDShortNamePluralization checks for obviously incorrect pluralization in shortNames
func TestCRDShortNamePluralization(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading CRDs: %v", err)
	}

	var errs []string
	for _, crd := range crds {
		// Only check CRDs with exactly 2 shortNames (likely singular and plural forms)
		if len(crd.Spec.Names.ShortNames) != 2 {
			continue
		}

		// Get all CRD versions
		var versions []string
		for _, v := range crd.Spec.Versions {
			versions = append(versions, v.Name)
		}
		versionStr := strings.Join(versions, ",")
		if versionStr == "" {
			versionStr = "unknown"
		}

		// Sort shortNames by length to identify singular (shorter) and plural (longer)
		shortNames := make([]string, len(crd.Spec.Names.ShortNames))
		copy(shortNames, crd.Spec.Names.ShortNames)
		sort.Slice(shortNames, func(i, j int) bool {
			return len(shortNames[i]) < len(shortNames[j])
		})

		singular := shortNames[0]
		plural := shortNames[1]

		// Check if the plural form is valid according to English pluralization rules
		if !isValidPlural(singular, plural) {
			errs = append(errs, fmt.Sprintf("[shortname_plural] crd=%s version=%s: plural shortName %q appears to have incorrect pluralization of %q", crd.Name, versionStr, plural, singular))
		}
	}

	sort.Strings(errs)
	want := strings.Join(errs, "\n")
	test.CompareGoldenFile(t, "testdata/exceptions/shortname_pluralization.txt", want)
}

// isValidPlural checks if a string is a valid pluralization of another string
func isValidPlural(singular, plural string) bool {
	// Special cases for words that are already plural or don't follow standard rules
	alreadyPluralWords := []string{"settings", "metrics", "series", "data"}
	for _, word := range alreadyPluralWords {
		if strings.HasSuffix(singular, word) {
			return plural == singular // Already plural words should stay the same
		}
	}

	// Rule 1: If singular ends with 's', 'x', 'z', 'ch', 'sh', add 'es'
	if strings.HasSuffix(singular, "s") ||
		strings.HasSuffix(singular, "x") ||
		strings.HasSuffix(singular, "z") ||
		strings.HasSuffix(singular, "ch") ||
		strings.HasSuffix(singular, "sh") {
		return plural == singular+"es"
	}

	// Rule 2: If singular ends with 'y' preceded by a consonant, change 'y' to 'ies'
	if strings.HasSuffix(singular, "y") && len(singular) > 1 {
		// Check if the character before 'y' is a consonant
		r := rune(singular[len(singular)-2])
		if !isVowel(r) {
			return plural == singular[:len(singular)-1]+"ies"
		}
	}

	// Rule 3: If singular ends with 'f' or 'fe', change to 'ves'
	if strings.HasSuffix(singular, "f") {
		return plural == singular[:len(singular)-1]+"ves"
	}
	if strings.HasSuffix(singular, "fe") {
		return plural == singular[:len(singular)-2]+"ves"
	}

	// Rule 4: Default case - just add 's'
	return plural == singular+"s"
}

// Helper function to check if a rune is a vowel
func isVowel(r rune) bool {
	r = unicode.ToLower(r)
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}
