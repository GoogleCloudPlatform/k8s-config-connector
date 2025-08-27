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

package reference

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ExternalNormalizer interface {
	// NormalizedExternal expects the implemented struct has a "External" field, and this function
	// assigns a value to the "External" field if it is empty.
	// In general, it retrieves the corresponding ConfigConnector object from the cluster, using
	// the `status.externalRef` or other field as the "External" value
	NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error)
}

type Ref interface {
	// GetGVK returns the schema.GroupVersionKind of the reference type
	GetGVK() schema.GroupVersionKind

	// GetNamespacedName returns the types.NamespacedName of a given reference
	GetNamespacedName() types.NamespacedName

	// GetExternal returns the external reference string (if set) of the reference.
	GetExternal() string

	// SetExternal sets the external reference string for a reference.
	SetExternal(ref string)

	// ValidateExternal returns nil if the given external reference string has a valid format for the reference.
	// Otherwise, it returns an error.
	ValidateExternal(ref string) error

	// Normalize ensures the "External" reference (in string format) is
	// set for a given Ref, and that it has the correct format.
	//
	// If "External" is already set, the format will be validated.
	//
	// If "External" is not set, the NamespacedName will be used to query the
	// referenced object from the K8s API and fetch it's external reference
	// value. If "Namespace" is not specified in the reference, the
	// `defaultNamespace“ will be used instead.
	Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error

	// NormalizeOnTemplate returns the resolved resource name formatted with the given template.
	NormalizeOnTemplate(ctx context.Context, kube client.Reader, defaultNamespace string, tpl string) error
}

// Normalize is a general-purpose reference resolver that can be used to
// implement the "Normalize" interface method for most Ref types.
func Normalize(ctx context.Context, reader client.Reader, ref Ref, defaultNamespace string) error {
	if ref.GetExternal() == "" {
		u, err := getReferencedObject(ctx, reader, ref, defaultNamespace)
		if err != nil {
			return err
		}
		externalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
		if err != nil {
			return fmt.Errorf("reading status.externalRef: %w", err)
		}
		if externalRef == "" {
			return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), ref.GetNamespacedName())
		}
		ref.SetExternal(externalRef)
		return err
	}
	return ref.ValidateExternal(ref.GetExternal())
}

// NormalizeOnTemplate is a reference resolver that can be used to
// implement the "NormalizeOnTemplate" interface method for Ref types that require a format different from the default.
func NormalizeOnTemplate(ctx context.Context, reader client.Reader, ref Ref, defaultNamespace string, tpl string) error {
	if tpl == "" {
		return ref.Normalize(ctx, reader, defaultNamespace)
	}

	// If external field is provided:
	// Verify if it matches the given template.
	external := ref.GetExternal()
	if external != "" {
		if matchTemplate(external, tpl) {
			return nil
		}
		return fmt.Errorf("format of %s external=%q was not known (use %s)", ref.GetGVK().Kind, external, tpl)
	}

	// Otherwise, name and name fields are provided:
	// Get the referenced object.
	u, err := getReferencedObject(ctx, reader, ref, defaultNamespace)
	if err != nil {
		return err
	}

	// Resolve placeholders from referenced object's status.externalRef field, if it exists
	processedTemplate := tpl
	externalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("reading status.externalRef: %w", err)
	}
	if externalRef != "" {
		components, err := ParseExternalRef(externalRef)
		if err != nil {
			return fmt.Errorf("parsing resource's externalRef %q: %w", externalRef, err)
		}
		// Replace all placeholders using the parsed externalRef components.
		for key, val := range components {
			placeholder := "{{" + key + "}}"
			processedTemplate = strings.ReplaceAll(processedTemplate, placeholder, val)
		}
	}

	// Resolve remaining placeholders from referenced object's status fields other than status.externalRef
	re := regexp.MustCompile(`\{\{(.*?)\}\}`)
	matches := re.FindAllStringSubmatch(processedTemplate, -1)
	if len(matches) == 0 {
		ref.SetExternal(processedTemplate)
		return nil
	}

	// Loop through remaining placeholders and replace them with values from referenced object's status fields
	for _, match := range matches {
		placeholder := match[0] // e.g., "{{status.selfLink}}"
		fieldName := match[1]   // e.g., "status.selfLink"

		fieldParts := strings.Split(fieldName, ".")
		value, found, err := unstructured.NestedString(u.Object, fieldParts...)
		if err != nil {
			return fmt.Errorf("error getting value for field %q from referenced object %s %s: %w", fieldName, ref.GetGVK(), u.GetNamespace()+"/"+u.GetName(), err)
		}
		if !found {
			return fmt.Errorf("template field %q not found in referenced object %s %s", fieldName, ref.GetGVK(), u.GetNamespace()+"/"+u.GetName())
		}
		processedTemplate = strings.Replace(processedTemplate, placeholder, value, 1)
	}

	// Final validation to ensure no placeholders are left
	remaining := re.FindAllString(processedTemplate, -1)
	if len(remaining) > 0 {
		return fmt.Errorf("template could not be fully resolved, unresolved placeholders: %v. Please check field existence in referenced object", remaining)
	}
	ref.SetExternal(processedTemplate)
	return nil
}

func getReferencedObject(ctx context.Context, reader client.Reader, ref Ref, defaultNamespace string) (*unstructured.Unstructured, error) {
	key := ref.GetNamespacedName()
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ref.GetGVK())
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("reading referenced %s %s: %w", ref.GetGVK(), key, err)
	}
	return u, nil
}

func matchTemplate(s, template string) bool {
	placeholderRegex := regexp.MustCompile(`\{\{.+?\}\}`)

	// Replace the placeholder with the regex pattern `([^/]+)`
	// This pattern matches one or more characters that are not a `/`
	// For example: http://{{host}}/api/{{version}} -> http://([^/]+)/api/([^/]+)
	regexTemplate := placeholderRegex.ReplaceAllString(template, "([^/]+)")

	// Add `^` at the beginning and `$` to the regexTemplate, ensure entire string matches
	re, err := regexp.Compile("^" + regexTemplate + "$")
	if err != nil {
		return false
	}
	return re.MatchString(s)
}
