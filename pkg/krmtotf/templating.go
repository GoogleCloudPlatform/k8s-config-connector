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

package krmtotf

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var fieldRegex = regexp.MustCompile("{{([0-9A-Za-z_.?]+)}}")

func ResolveValueTemplate(template string, val string, r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (string, error) {
	if template == "" {
		// An empty template is defined as no expansion necessary.
		return val, nil
	}
	ret := strings.ReplaceAll(template, "{{value}}", val)
	// Check to see if there are any Terraform fields to expand. If so, send it through field
	// expansion.
	if fieldRegex.MatchString(ret) {
		return expandTemplate(ret, r, c, smLoader)
	}
	return ret, nil
}

func resolveValueTemplateFromInterface(template string, val interface{}, r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (interface{}, error) {
	if template == "" {
		// An empty template is defined as no expansion necessary.
		return val, nil
	}
	switch valAsType := val.(type) {
	case string:
		return ResolveValueTemplate(template, valAsType, r, c, smLoader)
	default:
		return nil, fmt.Errorf("cannot resolve value template for non-string type")
	}
}

// returns true if the value could have been obtained from the template, i.e. if 'template' is "folders/{{value}}"
//   - if 'value' is "folders/1234567" the result will be true
//   - if 'value' is "organizations/1234567" the result will be false
func valueMatchesTemplate(template string, value string) bool {
	if template == "" {
		return true
	}
	template = strings.ReplaceAll(template, "{{value}}", ".*")
	r := regexp.MustCompile(template)
	return r.MatchString(value)
}

func expandTemplate(template string, r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (string, error) {
	leftBracketIdx := strings.Index(template, "[")
	if leftBracketIdx == -1 {
		if isORTemplate(template) {
			return expandOrTemplate(template, r, c, smLoader)
		}
		return expandFieldTemplate(template, r, c, smLoader)
	}
	rightBracketIdx := strings.LastIndex(template, "]")
	if rightBracketIdx == -1 {
		return "", fmt.Errorf("template '%v' has a left bracket but is missing corresponding right bracket", template)
	}
	subTemplate := template[leftBracketIdx+1 : rightBracketIdx]
	result, err := expandTemplate(subTemplate, r, c, smLoader)
	if err != nil {
		return "", fmt.Errorf("error resolving sub-template: %w", err)
	}
	resolvedTemplate := fmt.Sprintf("%v%v%v", template[:leftBracketIdx], result, template[rightBracketIdx+1:])
	return expandTemplate(resolvedTemplate, r, c, smLoader)
}

func isORTemplate(template string) bool {
	return strings.Contains(template, "|")
}

// expandFieldTemplate expands the given template by replacing any occurrence of {{field_name}} with
// the value of that field. Note that this may trigger further recursive field expansions, if the
// field's value depends on further template expansion.
func expandFieldTemplate(template string, r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (string, error) {
	var resolutionError error
	resolveFunc := func(s string) string {
		field := fieldRegex.FindStringSubmatch(s)[1]
		isRequired := true
		if strings.HasSuffix(field, "?") {
			isRequired = false
			field = strings.TrimSuffix(field, "?")
		}
		if isRef, refConfig := IsReferenceField(field, &r.ResourceConfig); isRef {
			val, found, err := getValueFromReference(refConfig, r, c, smLoader)
			if err != nil {
				resolutionError = fmt.Errorf("error getting value from reference: %w", err)
				return ""
			}
			if !found {
				if isRequired {
					resolutionError = fmt.Errorf("required reference '%v' could not be found in spec", field)
				}
				return ""
			}
			return val
		}
		if field == r.ResourceConfig.ResourceID.TargetField {
			val, err := resolveResourceID(r, c, smLoader)
			if err != nil {
				resolutionError = fmt.Errorf("error resolving resource ID: %w", err)
				return ""
			}
			return val
		}
		if field == r.ResourceConfig.MetadataMapping.Name {
			val, err := resolveNameMetadataMapping(r, c, smLoader)
			if err != nil {
				resolutionError = fmt.Errorf("error resolving metadata name mapping value: %w", err)
				return ""
			}
			return val
		}

		if field == r.ResourceConfig.MetadataMapping.Labels {
			resolutionError = fmt.Errorf("cannot map labels (map[string]string) to string field '%v'", field)
			return ""
		}
		if field == "region" && r.ResourceConfig.Locationality == gcp.Regional ||
			field == "zone" && r.ResourceConfig.Locationality == gcp.Zonal {
			if val, exists, _ := unstructured.NestedString(r.Spec, "location"); exists {
				return val
			}
		}
		if !SupportsHierarchicalReferences(&r.ResourceConfig) {
			// TODO(b/193177782): Delete this if-block once all resources
			// support hierarchical references.
			for _, c := range r.ResourceConfig.Containers {
				if field == c.TFField {
					annotation := k8s.GetAnnotationForContainerType(c.Type)
					val, ok := k8s.GetAnnotation(annotation, r)
					if (!ok || val == "") && isRequired {
						resolutionError = fmt.Errorf("no value found for annotation %v", annotation)
						return ""
					}
					return val
				}
			}
		}
		for _, d := range r.ResourceConfig.Directives {
			if field == d {
				annotation := k8s.FormatAnnotation(text.SnakeCaseToKebabCase(d))
				val, ok := k8s.GetAnnotation(annotation, r)
				if (!ok || val == "") && isRequired {
					resolutionError = fmt.Errorf("no value found for annotation %v", annotation)
					return ""
				}
				return val
			}
		}
		path := text.SnakeCaseToLowerCamelCase(field)
		if val, exists, _ := unstructured.NestedString(r.Spec, strings.Split(path, ".")...); exists {
			return val
		}
		if val, exists, _ := unstructured.NestedString(r.GetStatusOrObservedState(), strings.Split(path, ".")...); exists {
			return val
		}

		// todo(yuhou): Special handling to resolve project from DCL-based resource
		// Only used for fixtures/globalcomputebackendservicesecuritysettings test for now
		if path == "project" {
			dclPath := "projectRef.external"
			if val, exists, _ := unstructured.NestedString(r.Spec, strings.Split(dclPath, ".")...); exists {
				return val
			}
		}

		if isRequired {
			resolutionError = fmt.Errorf("unable to resolve missing value: %v", field)
		}
		return ""
	}
	return fieldRegex.ReplaceAllStringFunc(template, resolveFunc), resolutionError
}

func expandOrTemplate(template string, r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (string, error) {
	templates := strings.Split(template, "|")
	expectedCount := 2
	if len(templates) != expectedCount {
		return "", fmt.Errorf("unexpected template format: after splitting on '|' there are '%v' templates when %v were expected",
			len(templates), expectedCount)
	}
	t1 := templates[0]
	t2 := templates[1]
	result, err := expandTemplate(t1, r, c, smLoader)
	if err == nil {
		return result, nil
	}
	result, err = expandTemplate(t2, r, c, smLoader)
	if err != nil {
		return "", fmt.Errorf("error resolving both sides of an '|' template: %w", err)
	}
	return result, nil
}

func getValueFromReference(refConfig *corekccv1alpha1.ReferenceConfig, r *Resource, c client.Client,
	smLoader *servicemappingloader.ServiceMappingLoader) (val string, found bool, err error) {
	pathToRef := getPathToReferenceKey(refConfig)
	refObj, ok, err := unstructured.NestedMap(r.Spec, pathToRef...)
	if err != nil {
		return "", false, fmt.Errorf("error getting reference object '%v': %w", strings.Join(pathToRef, "."), err)
	}
	if !ok {
		return "", false, nil
	}
	retRaw, err := ResolveReferenceObject(refObj, *refConfig, r, c, smLoader)
	if err != nil {
		return "", false, fmt.Errorf("error resolving reference field: %w", err)
	}
	ret, ok := retRaw.(string)
	if !ok {
		return "", false, fmt.Errorf("could not parse reference resolution value '%+v' as string", retRaw)
	}
	return ret, true, nil
}

func extractValueSegmentFromIDInStatus(idInStatus, template string) (string, error) {
	if template == "" {
		return idInStatus, nil
	}

	// Convert the template to be a compilable regular expression.
	template = strings.ReplaceAll(template, "{{value}}", "([^/]+)")

	// Template starting with the parent field value may contain additional
	// slashes that are not in the template.
	// E.g. DataCatalogPolicyTag has resource ID template:
	// "{{taxonomy}}/policyTags/{{value}}", however the value of the parent
	// field, 'spec.taxonomy', is "projects/test-project/locations/us/taxonomies/tid",
	// which contains additional slashes that are not captured in the template.
	if strings.HasPrefix(template, "{{") {
		re := regexp.MustCompile(`^({{[a-z]([a-z_]*[a-z])*}})/.*$`)
		matched := re.FindStringSubmatch(template)
		if len(matched) == 0 {
			return "", fmt.Errorf("error extracting the parent field name from resource ID template %v", template)
		}
		parentField := matched[1]
		template = strings.ReplaceAll(template, parentField, "[^/](.*[^/])*")
	}

	template = fieldRegex.ReplaceAllString(template, "[^/]+")
	template = fmt.Sprintf("%s%s%s", "^", template, "$")

	// Extract out the resourceID from the value.
	templateRegex := regexp.MustCompile(template)
	subMatches := templateRegex.FindStringSubmatch(idInStatus)
	if len(subMatches) < 2 {
		return "", fmt.Errorf("error extracting out the value segment "+
			"from the idInStatus template '%s' using template '%s'", idInStatus,
			template)
	}

	return subMatches[len(subMatches)-1], nil
}
