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
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// resolveUnmanagedFields sets fields which no other server-side apply manager aside from
// KCC has expressed an opinion on to their live state value.
func resolveUnmanagedFields(spec map[string]interface{}, r *Resource, liveState *terraform.InstanceState,
	jsonSchema *apiextensions.JSONSchemaProps) (map[string]interface{}, error) {
	if r.ManagedFields == nil {
		// If this resource does not have SSA enabled, there is no way to distinguish
		// fields that are k8s-managed from fields that are externally-managed.
		// In this case, treat the whole spec in etcd as the desired state.
		return spec, nil
	}

	var stateAsKRM map[string]interface{}
	var err error

	if liveState.Empty() {
		stateAsKRM = make(map[string]interface{})
	} else {
		stateAsKRM, _ = GetSpecAndStatusFromState(r, liveState)
	}

	switch r.Kind {
	// TODO(b/223303389): Roll out ability to switch between conflicting fields
	// to remaining resources.
	case "BigtableAppProfile", "CloudBuildTrigger", "ResourceManagerPolicy":
		if err = RemoveFieldsFromStateThatConflictWithSpec(stateAsKRM, spec, r.ResourceConfig, []string{}, r.TFResource.Schema); err != nil {
			return nil, fmt.Errorf("error stripping fields from state that conflict with fields already in spec: %w", err)
		}
	}

	return k8s.OverlayManagedFieldsOntoState(spec, stateAsKRM, r.ManagedFields, jsonSchema, r.ResourceConfig.HierarchicalReferences)
}

// RemoveFieldsFromStateThatConflictWithSpec removes fields from 'state' that
// conflict with any of the fields found in 'spec'. This is useful for when we
// want to overlay 'state' onto 'spec' without ending up with an invalid
// resource configuration.
func RemoveFieldsFromStateThatConflictWithSpec(state map[string]interface{}, spec map[string]interface{},
	rc corekccv1alpha1.ResourceConfig, tfPath []string, schemaMap map[string]*tfschema.Schema) error {
	if len(state) == 0 || len(spec) == 0 {
		return nil
	}
	for k, s := range schemaMap {
		tfPath := append(tfPath, k)
		tfField := strings.Join(tfPath, ".")
		krmPath := text.SnakeCaseStrsToLowerCamelCaseStrs(tfPath)
		if ok, refConfig := IsReferenceField(tfField, &rc); ok {
			krmPath = getPathToReferenceKey(refConfig)
		}

		_, found, err := unstructured.NestedFieldNoCopy(state, krmPath...)
		if err != nil {
			return fmt.Errorf("error checking for existence of field with path %v in state: %w", krmPath, err)
		}
		if !found {
			continue
		}

		// Determine the keys which conflict with the current key. Note:
		// ConflictsWith and ExactlyOneOf specify keys in format
		// "parent_field.0.child_field.0.grandchild_field".
		conflictingTFKeys := slice.RemoveStringFromStringSlice(
			slice.ConcatStringSlices(s.ConflictsWith, s.ExactlyOneOf),
			strings.Join(tfPath, ".0."), // Remove current key which is included in ExactlyOneOf.
		)
		for _, conflictingTFKey := range conflictingTFKeys {
			conflictingTFPath := strings.Split(conflictingTFKey, ".0.")
			conflictingTFField := strings.Join(conflictingTFPath, ".")
			conflictingKRMPath := text.SnakeCaseStrsToLowerCamelCaseStrs(conflictingTFPath)
			if ok, refConfig := IsReferenceField(conflictingTFField, &rc); ok {
				conflictingKRMPath = getPathToReferenceKey(refConfig)
			}

			_, found, err := unstructured.NestedFieldNoCopy(spec, conflictingKRMPath...)
			if err != nil {
				return fmt.Errorf("error checking for existence of conflicting field with path %v in spec: %w", krmPath, err)
			}
			if !found {
				continue
			}

			if err := removeFieldAndAnyEmptyAncestorsFromObject(krmPath, state); err != nil {
				return fmt.Errorf("error removing field at path %v from state: %w", krmPath, err)
			}
		}

		// If field is an object, check if any of its subfields need to be
		// removed. Note: ConflictsWith and ExactlyOneOf can only point to
		// top-level fields or fields nested in objects. They cannot point to
		// fields in objects in lists/sets.
		if !tfresource.IsObjectField(s) {
			continue
		}
		if err := RemoveFieldsFromStateThatConflictWithSpec(state, spec, rc, tfPath, s.Elem.(*tfschema.Resource).Schema); err != nil {
			return err
		}
	}
	return nil
}

// removeFieldAndAnyEmptyAncestorsFromObject removes the field at 'path' from
// 'obj'. And, if removing the field results in an empty object at the parent
// field, the function removes the parent field too, and then the grandparent
// field, and so on.
func removeFieldAndAnyEmptyAncestorsFromObject(path []string, obj map[string]interface{}) error {
	unstructured.RemoveNestedField(obj, path...)

	// If the field is a top-level field, we can stop now since there are no
	// ancestor fields to potentially remove.
	if len(path) <= 1 {
		return nil
	}

	// Remove ancestor fields that are now empty, starting from the parent
	// field and then going up the tree.
	for i := len(path) - 2; i >= 0; i-- {
		ancestorPath := path[0 : i+1]
		val, found, err := unstructured.NestedFieldNoCopy(obj, ancestorPath...)
		if err != nil {
			return fmt.Errorf("error checking for existence of field ancestor at path: %v: %w", ancestorPath, err)
		}
		if !found {
			return fmt.Errorf("unexpectedly failed to find field ancestor at path %v", ancestorPath)
		}
		valAsObj, ok := val.(map[string]interface{})
		if !ok {
			return fmt.Errorf("field ancestor at path %v is unexpectedly not a map[string]interface{}", ancestorPath)
		}
		if len(valAsObj) > 0 {
			return nil
		}
		unstructured.RemoveNestedField(obj, ancestorPath...)
	}
	return nil
}
