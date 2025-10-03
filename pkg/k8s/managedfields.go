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

package k8s

import (
	"bytes"
	"fmt"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
	"sigs.k8s.io/structured-merge-diff/v4/schema"
	"sigs.k8s.io/structured-merge-diff/v4/typed"
	"sigs.k8s.io/structured-merge-diff/v4/value"
)

/*

This file contains all the functionality around field management in KCC.
KCC supports the concept of "externally-managed" fields
(https://cloud.google.com/config-connector/docs/concepts/managing-fields-externally),
which respects the live value for fields on the underlying API that no Kubernetes user manages.
Thus, we need to parse and perform operations using the core Kubernetes field management metadata.

For more information on core Kubernetes field management, see the Server-Side
Apply Kubernetes documentation:
https://kubernetes.io/docs/reference/using-api/server-side-apply

This package makes heavy use of the functionality exposed by the
`sigs.k8s.io/structured-merge-diff` library. This is the library used by the
Kubernetes API server to perform all its operations surrounding its managed
field metadata.

The following packages from `sigs.k8s.io/structured-merge-diff` are leveraged:

* sigs.k8s.io/structured-merge-diff/v4/fieldpath: We make use of the
  `fieldpath.Set` type. This exposes standard set operations like Union,
  Difference, Intersection for sets of fields. A JSON-encoding of this type is
  what is directly stored in the core Kubernetes
  `metadata.managedFields[].fieldsV1` field.

* sigs.k8s.io/structured-merge-diff/v4/value: We make use of the `value.Value`
  type. This is just a wrapper type for a JSON type. All we need to do with
  this structure is create one with a map[string]interface{}, and extract the
  updated map[string]interface{} at the end.

* sigs.k8s.io/structured-merge-diff/v4/typed: We make use of the
  `typed.TypedValue` type. The struct itself is just a wrapper for a pair of a
  `value.Value` and a `schema.Schema` (explained below). However, it exposes a
  bunch of handy functions around comparing and extracting values from objects.
  For example:
  * The `Compare` function takes two objects (and old state and an updated
    state) and returns a Comparison object with field sets of everything that
    was added, modified, and removed.
  * The `FieldSet` function returns a set of all the fields that show up in the
    object. This is useful for comparing to the core Kubernetes managed field
    metadata.
  * The `Merge` function takes a partial state and merges it with an object.
  * The `ExtractItems` function can take a field set and extract a partial
    state from an object with just those fields represented.
  * The `RemoveItems` function can take a field set and remove all those fields
    from an object.

* sigs.k8s.io/structured-merge-diff/v4/schema: We use the `schema.Schema` type.
  This is a schema type exclusive to structured-merge-diff (abbreviated as SMD)
  operations. It is a requirement to create a `typed.TypedValue`, but no
  operations are used on it directly. This just decouples the schema
  information relevant to SMD operations from other schema types like
  `apiextensions.JSONSchemaProps`.

*/

func IsK8sManaged(key string, specObj map[string]interface{}, managedFields *fieldpath.Set) bool {
	pe := fieldpath.PathElement{FieldName: &key}
	if managedFields == nil {
		// If no managed field information present, treat values specified in the
		// spec as k8s-managed.
		_, ok := specObj[key]
		return ok
	}
	if managedFields.Members.Has(pe) {
		return true
	}
	// We must also check the nested objects within the managed fields for management
	// in order to handle the following cases:
	// - Sensitive fields are converted from strings to complex reference objects
	// - Maps, though nested objects and technically able to be merged, must be treated
	//   as atomic in order to allow for users to clear entries
	_, found := managedFields.Children.Get(pe)
	return found
}

// ConstructManagedFieldsV1Set takes the given managed field entries and constructs a
// set of all the k8s-managed fields from the spec.
func ConstructManagedFieldsV1Set(managedFields []v1.ManagedFieldsEntry) (*fieldpath.Set, error) {
	res := fieldpath.NewSet()
	for _, managedFieldEntry := range managedFields {
		if managedFieldEntry.Manager == ControllerManagedFieldManager {
			continue
		}
		if managedFieldEntry.FieldsType != ManagedFieldsTypeFieldsV1 {
			return nil, fmt.Errorf(
				"expected managed field entry for manager '%v' and operation '%v' of type '%v', got type '%v'",
				managedFieldEntry.Manager, managedFieldEntry.Operation, ManagedFieldsTypeFieldsV1,
				managedFieldEntry.FieldsType)
		}
		fieldsV1 := managedFieldEntry.FieldsV1
		if managedFieldEntry.FieldsV1 == nil {
			return nil, fmt.Errorf("managed field entry for manager '%v' and operation '%v' has empty fieldsV1",
				managedFieldEntry.Manager, managedFieldEntry.Operation)
		}
		entrySet := fieldpath.NewSet()
		if err := entrySet.FromJSON(bytes.NewReader(fieldsV1.Raw)); err != nil {
			return nil, fmt.Errorf("error marshaling managed fields for manager '%v' and operation '%v' from JSON: %w",
				managedFieldEntry.Manager, managedFieldEntry.Operation, err)
		}
		specFieldName := "spec"
		specSet, found := entrySet.Children.Get(fieldpath.PathElement{FieldName: &specFieldName})
		if !found {
			continue
		}
		res = res.Union(specSet)
	}
	return res, nil
}

func containsUnsupportedFieldTypes(managedFields []v1.ManagedFieldsEntry) bool {
	for _, entry := range managedFields {
		// Only FieldsV1 is currently supported.
		if entry.FieldsType != ManagedFieldsTypeFieldsV1 {
			return true
		}
	}
	return false
}

func GetK8sManagedFields(u *unstructured.Unstructured) (*fieldpath.Set, error) {
	managedFields := u.GetManagedFields()
	if managedFields != nil && !containsUnsupportedFieldTypes(managedFields) {
		res, err := ConstructManagedFieldsV1Set(managedFields)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, nil
}

// OverlayManagedFieldsOntoState overlays the fields managed by Kubernetes managers onto the
// KRM-ified live state.
//
// The return value is the union of stateAsKRM with managed fields from spec.
func OverlayManagedFieldsOntoState(spec, stateAsKRM map[string]interface{}, managedFields *fieldpath.Set,
	jsonSchema *apiextensions.JSONSchemaProps, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) (map[string]interface{}, error) {
	config, err := overlayManagedFieldsOntoState(spec, stateAsKRM, managedFields, jsonSchema)
	if err != nil {
		return nil, err
	}

	// Add the resource's hierarchical reference back to the config if the
	// resource supports hierarchical references. We do this because
	// overlayManagedFieldsOntoState() strips out the hierarchical reference if
	// it was defaulted by the KCC webhook/controller. This is because the
	// resource's managedFields metadata doesn't include the hierarchical
	// reference if it was defaulted by KCC rather than explicitly set by the
	// user. Since we know that a resource's hierarchical reference is part of
	// the desired state even if it's not in the managedFields metadata, let's
	// add it back in.
	// TODO(b/184319410): Come up with a more generic solution to ensure that
	// fields added by webhooks/controllers are preserved as part of the
	// desired state if they truly belong there.
	if len(hierarchicalRefs) > 0 {
		config, err = addHierarchicalReferenceToConfig(config, spec, hierarchicalRefs)
		if err != nil {
			return nil, fmt.Errorf("error adding hierarchical reference to config: %w", err)
		}
	}

	return config, nil
}

func overlayManagedFieldsOntoState(spec, stateAsKRM map[string]interface{}, managedFields *fieldpath.Set,
	jsonSchema *apiextensions.JSONSchemaProps) (map[string]interface{}, error) {
	if jsonSchema == nil {
		return nil, fmt.Errorf("JSON schema is required")
	}
	specSchema, ok := jsonSchema.Properties["spec"]
	if !ok {
		if spec != nil && len(spec) > 0 {
			return nil, fmt.Errorf("cannot parse spec with no spec schema available")
		}
		return make(map[string]interface{}), nil
	}

	// Wrap the spec and live state objects with `typed.TypedValue`
	// objects, as this type has methods defined on it for operations
	// like calculating field sets, extracting partial states, and
	// merging partial states.
	smdSchema := jsonSchemaToSMDSchema(&specSchema)
	specAsTyped, err := toTypedValue(spec, smdSchema)
	if err != nil {
		return nil, fmt.Errorf("error converting spec to typed value: %w", err)
	}
	stateAsTyped, err := toTypedValue(stateAsKRM, smdSchema)
	if err != nil {
		return nil, fmt.Errorf("error converting state to typed value: %w", err)
	}

	if managedFields == nil {
		managedFields = &fieldpath.Set{}
	}

	// Construct a set of only the leaves in order to avoid the behavior
	// during ValueType.Merge where a parent element being a member of the
	// set (i.e. listed as "." in the managed fields) causes the child
	// elements to be left unmerged and the parent object taken wholesale.
	managedFields = managedFields.Leaves()

	// Treat _atomic_ list fields as k8s-managed. We do this since the controller
	// assumes management over atomic lists to be able to default values inside
	// objects in lists. Therefore, to avoid suppressing intentional diffs,
	// treat any atomic list as k8s-managed.
	atomicListFields, err := getAtomicListFields(smdSchema)
	if err != nil {
		return nil, fmt.Errorf("error getting atomic list fields: %w", err)
	}
	specFieldSet, err := specAsTyped.ToFieldSet()
	if err != nil {
		return nil, fmt.Errorf("error constructing field set for spec: %w", err)
	}
	atomicListFields = atomicListFields.Intersection(specFieldSet)
	managedFields = managedFields.Union(atomicListFields)

	// Extract k8s-managed fields from the spec, and then merge them with the live
	// state. In cases where the k8s-managed state and the live state both have the
	// same field, the value from the k8s-managed state is taken.
	k8sManagedPartialState := specAsTyped.ExtractItems(managedFields)
	overlaidState, err := stateAsTyped.Merge(k8sManagedPartialState)
	if err != nil {
		return nil, fmt.Errorf("error merging partial managed state with live state: %w", err)
	}

	// Unwrap the `typed.TypedValue` back to a map[string]interface{} that
	// the rest of reconciliation expects.
	overlaidStateRaw := overlaidState.AsValue().Unstructured()
	if overlaidStateRaw == nil {
		return make(map[string]interface{}), nil
	}
	res, ok := overlaidStateRaw.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("overlaid state unstructured not of type map[string]interface{}")
	}
	return res, nil
}

// Construct the trimmed spec that only contains k8s managed fields.
//
// The DCL SDK's Apply() function can take a partial state that only contains fields that users have
// an opinion on. Here we will look into the managed-fields set and trim the full spec to only preserve
// fields that are k8s-managed (i.e. users want KCC to enforce those fields to their desired state).
// DCL will take the generated partial state, enforce specified fields and ignore unspecified fields
// by preserving live values from the underlying API.
func ConstructTrimmedSpecWithManagedFields(resource *Resource, jsonSchema *apiextensions.JSONSchemaProps,
	hierarchicalRefs []corekccv1alpha1.HierarchicalReference) (map[string]interface{}, error) {
	if resource.ManagedFields == nil {
		// If no managed field information present, treat values specified in the
		// spec as k8s-managed.
		return resource.Spec, nil
	}
	if resource.Spec == nil {
		return nil, nil
	}
	// construct an empty state map to overlay onto
	emptyState := make(map[string]interface{})
	trimmed, err := OverlayManagedFieldsOntoState(resource.Spec, emptyState,
		resource.ManagedFields, jsonSchema, hierarchicalRefs)
	if err != nil {
		return nil, fmt.Errorf("error constructing trimmed state with managed fields: %w", err)
	}
	return trimmed, nil
}

// getAtomicListFields returns a field set of all atomic list fields in the
// given schema.
func getAtomicListFields(s *schema.Schema) (*fieldpath.Set, error) {
	if len(s.Types) != 1 {
		return nil, fmt.Errorf("expected schema to have 1 type, got %v", len(s.Types))
	}
	typeDef := s.Types[0]
	if typeDef.Map == nil {
		return nil, fmt.Errorf("type definition is not map")
	}
	return getAtomicListFieldsFromMap(typeDef.Map), nil
}

func getAtomicListFieldsFromMap(m *schema.Map) *fieldpath.Set {
	res := &fieldpath.Set{}
	for _, structField := range m.Fields {
		fieldName := structField.Name
		pe := fieldpath.PathElement{FieldName: &fieldName}
		fieldAtom := structField.Type.Inlined
		if fieldAtom.List != nil {
			switch fieldAtom.List.ElementRelationship {
			case schema.Atomic:
				res.Members.Insert(pe)
			case schema.Associative:
				elemAtom := fieldAtom.List.ElementType.Inlined
				if elemAtom.Map != nil {
					nestedSet := getAtomicListFieldsFromMap(elemAtom.Map)
					if !nestedSet.Empty() {
						insertChild(res, nestedSet, pe)
					}
				}
			}
		} else if fieldAtom.Map != nil {
			nestedSet := getAtomicListFieldsFromMap(fieldAtom.Map)
			if !nestedSet.Empty() {
				insertChild(res, nestedSet, pe)
			}
		}
	}
	return res
}

func insertChild(set, childSet *fieldpath.Set, pe fieldpath.PathElement) {
	childSetPtr := set.Children.Descend(pe)
	*childSetPtr = *childSet
}

func toTypedValue(obj map[string]interface{}, smdSchema *schema.Schema) (*typed.TypedValue, error) {
	// The SMD schema constructed by jsonSchemaToSMDSchema is hardcoded to have
	// only one type definition, with a blank name. Thus, our type reference
	// here should refer to the empty string.
	name := ""
	return typed.AsTyped(value.NewValueInterface(obj),
		smdSchema,
		schema.TypeRef{
			NamedType: &name,
		},
	)
}

// jsonSchemaToSMDSchema constructs a structured-merge-diff (SMD) schema.Schema
// object from the given JSON schema. The SMD schema is a requirement to
// create the typed.TypedValue objects.
//
// Note that `k8s.io/apiserver` defines a `TypeConverter` interface type that
// is intended to directly generate a `typed.TypedValue` object from a
// `runtime.Object`. However, the only concrete implementation takes a
// schema as input in the form of `k8s.io/openapi/pkg/util/proto.Models`
// rather than the `apiextensions.JSONSchemaProps` our controller has access
// to. As the SMD schema generation is quite straightforward and allows for
// only depending on `sigs.k8s.io/structured-merge-diff`, we choose to own
// our own conversion logic.
func jsonSchemaToSMDSchema(jsonSchema *apiextensions.JSONSchemaProps) *schema.Schema {
	return &schema.Schema{
		Types: []schema.TypeDef{
			{
				// Type definitions are named. However, since this schema only
				// contains one definition, we can just leave the name blank.
				Name: "",
				Atom: jsonSchemaToAtom(jsonSchema),
			},
		}}
}

func jsonSchemaToAtom(jsonSchema *apiextensions.JSONSchemaProps) schema.Atom {
	res := schema.Atom{}
	var (
		// scalarPtr is a helper function that allows us to easily reference
		// the built-in schema.Scalar constants in a context where a pointer
		// is required (since Go does not allow pointers to constants).
		scalarPtr = func(s schema.Scalar) *schema.Scalar { return &s }

		// scalarUnknown is a custom scalar type for use when the JSON schema
		// has no schema available for map elements. We must include some
		// sort of schema value, as map validation fails otherwise. Merges
		// on custom scalar types are supported by the SMD library.
		scalarUnknown = schema.Untyped
	)
	switch jsonSchema.Type {
	case "object":
		res.Map = &schema.Map{}
		if jsonSchema.Properties != nil {
			for field, fieldSchema := range jsonSchema.Properties {
				res.Map.Fields = append(res.Map.Fields, schema.StructField{
					Name: field,
					Type: schema.TypeRef{
						Inlined: jsonSchemaToAtom(&fieldSchema),
					},
				})
			}
			break
		}
		if jsonSchema.AdditionalProperties != nil {
			res.Map.ElementType = schema.TypeRef{
				Inlined: jsonSchemaToAtom(jsonSchema.AdditionalProperties.Schema),
			}
			break
		}
		res.Map.ElementType = schema.TypeRef{
			Inlined: schema.Atom{Scalar: &scalarUnknown},
		}
	case "array":
		res.List = &schema.List{
			ElementType: schema.TypeRef{
				Inlined: jsonSchemaToAtom(jsonSchema.Items.Schema),
			},
			ElementRelationship: schema.Atomic,
		}
	case "integer", "number":
		res.Scalar = scalarPtr(schema.Numeric)
	case "string":
		res.Scalar = scalarPtr(schema.String)
	case "boolean":
		res.Scalar = scalarPtr(schema.Boolean)
	default:
		panic(fmt.Sprintf("unknown JSON schema type %v", jsonSchema.Type))
	}
	return res
}

func addHierarchicalReferenceToConfig(config, spec map[string]interface{}, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) (map[string]interface{}, error) {
	modifiedConfig := deepcopy.MapStringInterface(config)
	resourceRef, hierarchicalRef, err := GetHierarchicalReferenceFromSpec(spec, hierarchicalRefs)
	if err != nil {
		return nil, fmt.Errorf("error getting hierarchical reference: %w", err)
	}
	if resourceRef == nil {
		return modifiedConfig, nil
	}
	var resourceRefRaw map[string]interface{}
	if err := util.Marshal(resourceRef, &resourceRefRaw); err != nil {
		return nil, fmt.Errorf("error marshalling hierarchical reference to map[string]interface{}: %w", err)
	}
	if err := unstructured.SetNestedField(modifiedConfig, resourceRefRaw, hierarchicalRef.Key); err != nil {
		return nil, fmt.Errorf("error setting hierarchical reference in config: %w", err)
	}
	return modifiedConfig, nil
}
