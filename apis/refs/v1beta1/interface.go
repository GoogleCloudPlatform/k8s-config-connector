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

package v1beta1

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

// NormalizeOnTemplate returns the resolved resource name formatted with the given template.
// Used for Ref types that require a format different from the default.
func NormalizeOnTemplate(ctx context.Context, reader client.Reader, ref Ref, defaultNamespace string, tpl string) error {
	// Default case: Use status.externalRef(direct) or normalize to the same format(legacy).
	if tpl == "" {
		return ref.Normalize(ctx, reader, defaultNamespace)
	}

	// todo(yuhou): Double-check whether we need this or not
	// If external is provided: verify if it matches the given template.

	//external := ref.GetExternal()
	//if external != "" {
	//	if matchTemplate(external, tpl) {
	//		return nil
	//	}
	//	return fmt.Errorf("format of %s external=%q was not known (use %s)", ref.GetGVK().Kind, external, tpl)
	//}

	// Otherwise, name and name fields are provided:
	// Get the referenced object.
	u, err := getReferencedObject(ctx, reader, ref, defaultNamespace)
	if err != nil {
		return err
	}

	processedTemplate := tpl

	// Case 1: Resolve placeholders from referenced object's resourceId
	// Get resourceId from the last part of the normalized external
	err = ref.Normalize(ctx, reader, defaultNamespace)
	if err != nil {
		return err
	}
	normalizedExternal := ref.GetExternal()
	tokens := strings.Split(normalizedExternal, "/")
	resourceId := tokens[len(tokens)-1]
	processedTemplate = strings.Replace(processedTemplate, "{{resourceId}}", resourceId, 1)

	// Case 2: Resolve placeholders from referenced object's `status` or `status.observedState` fields
	// Get placeholders
	re := regexp.MustCompile(`\{\{(.*?)\}\}`)
	matches := re.FindAllStringSubmatch(processedTemplate, -1)
	if len(matches) == 0 {
		ref.SetExternal(processedTemplate)
		return nil
	}

	for _, match := range matches {
		placeholder := match[0] // "{{selfLink}}"
		fieldName := match[1]   // "selfLink"

		fieldParts := strings.Split(fieldName, ".")
		observedStateFieldParts := append([]string{"status", "observedState"}, fieldParts...)
		value, found, err := unstructured.NestedString(u.Object, observedStateFieldParts...)
		if err != nil {
			return fmt.Errorf("error getting value for field status.observedState.%s from referenced object %s %s: %w", fieldName, ref.GetGVK(), u.GetNamespace()+"/"+u.GetName(), err)
		}
		if !found {
			statusFieldParts := append([]string{"status"}, fieldParts...)
			value, found, err = unstructured.NestedString(u.Object, statusFieldParts...)
			if err != nil {
				return fmt.Errorf("error getting value for field status.%s from referenced object %s %s: %w", fieldName, ref.GetGVK(), u.GetNamespace()+"/"+u.GetName(), err)
			}
			if !found {
				return fmt.Errorf("template field %q not found in referenced object %s %s (looked in 'status.observedState' and 'status')", fieldName, ref.GetGVK(), u.GetNamespace()+"/"+u.GetName())
			}
		}
		processedTemplate = strings.Replace(processedTemplate, placeholder, value, 1)
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

// todo(yuhou): Double-check whether we need this or not
//func matchTemplate(s, template string) bool {
//	placeholderRegex := regexp.MustCompile(`\{\{.+?\}\}`)
//
//	// Replace the placeholder with the regex pattern `([^/]+)`
//	// This pattern matches one or more characters that are not a `/`
//	// For example: http://{{host}}/api/{{version}} -> http://([^/]+)/api/([^/]+)
//	regexTemplate := placeholderRegex.ReplaceAllString(template, "([^/]+)")
//
//	// Add `^` at the beginning and `$` to the regexTemplate, ensure entire string matches
//	re, err := regexp.Compile("^" + regexTemplate + "$")
//	if err != nil {
//		return false
//	}
//	return re.MatchString(s)
//}
