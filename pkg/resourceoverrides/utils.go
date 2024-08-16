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

package resourceoverrides

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration/crdboilerplate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/crdutil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

type MultiKindRef struct {
	Kind        string
	TargetField string
}

// KeepTopLevelFieldOptionalWithDefault decorates the input CRD to modify the given top field as optional with the default.
func KeepTopLevelFieldOptionalWithDefault(crd *apiextensions.CustomResourceDefinition, defaultValue interface{}, field string) error {
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	spec := schema.Properties["spec"]
	fieldSchema := spec.Properties[field]
	bytes, err := json.Marshal(defaultValue)
	if err != nil {
		return err
	}
	fieldSchema.Default = &apiextensions.JSON{
		Raw: bytes,
	}
	spec.Properties[field] = fieldSchema
	// mark the given field optional
	required := make([]string, 0)
	for _, v := range spec.Required {
		if v != field {
			required = append(required, v)
		}
	}
	spec.Required = required
	schema.Properties["spec"] = spec
	if len(spec.Required) == 0 {
		schema.Required = []string{}
	}
	return nil
}

func getSchemaForPath(schema *apiextensions.JSONSchemaProps, path []string) (*apiextensions.JSONSchemaProps, error) {
	if len(path) == 0 {
		return schema, nil
	}

	var subSchema apiextensions.JSONSchemaProps
	var ok bool
	if schema.Properties != nil {
		if subSchema, ok = schema.Properties[path[0]]; !ok {
			return nil, fmt.Errorf("can't find field %s under properties", path[0])
		}
	} else if schema.AdditionalProperties != nil {
		// TODO: Handle this edge case once there is a concrete example.
		return nil, fmt.Errorf("field %s is under a map field: this path can't be handled", path[0])
	} else {
		return nil, fmt.Errorf("the schema doesn't support any subfield")
	}

	if len(path) == 1 {
		return &subSchema, nil
	}

	switch subSchema.Type {
	case "array":
		return getSchemaForPath(subSchema.Items.Schema, path[1:])
	default:
		return getSchemaForPath(&subSchema, path[1:])
	}
}

// PreserveMutuallyExclusiveNonReferenceField adds back the non-ref field to keep the
// CRD backwards compatible.
func PreserveMutuallyExclusiveNonReferenceField(crd *apiextensions.CustomResourceDefinition, parentPath []string, referenceFieldName, nonReferenceFieldName string) error {
	if referenceFieldName == "" || nonReferenceFieldName == "" {
		return fmt.Errorf("both 'referenceFieldName' and 'nonReferenceFieldName' must be specified")
	}

	// 1. Get the parent schema of the fields.
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	var err error
	var parent *apiextensions.JSONSchemaProps
	// Prepend the top-level 'spec' field into path.
	if len(parentPath) == 0 {
		parentPath = []string{"spec"}
	} else {
		parentPath = append([]string{"spec"}, parentPath...)
	}
	parentPathStr := strings.Join(parentPath, ".")
	parent, err = getSchemaForPath(schema, parentPath)
	if err != nil {
		return fmt.Errorf("can't get schema for path '%v' in CRD %s: %w", parentPathStr, crd.Name, err)
	}

	// 2. Check if the reference field is required.
	requiredFields, err := crdutil.GetRequiredRuleForObjectOrArray(parent)
	if err != nil {
		return fmt.Errorf("error getting the required rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
	}
	var required bool
	for _, field := range requiredFields {
		if field == referenceFieldName {
			required = true
			break
		}
	}

	// 3. Set the `oneOf` or `not` rule based on whether the reference field is
	// required.
	if required {
		oneOfRule, err := crdutil.GetOneOfRuleForObjectOrArray(parent)
		if err != nil {
			return fmt.Errorf("error getting the oneOf rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
		}
		// TODO(b/223688758): Handle multiple oneOf rules.
		if oneOfRule != nil {
			return fmt.Errorf("can't handle multiple pairs of required mutually exclustive fields under %s for field %s and %s in CRD %s", parentPath, referenceFieldName, referenceFieldName, crd.Name)
		}

		oneOfRule = []*apiextensions.JSONSchemaProps{
			{
				Required: []string{nonReferenceFieldName},
			},
			{
				Required: []string{referenceFieldName},
			},
		}
		if err := crdutil.SetOneOfRuleForObjectOrArray(parent, oneOfRule); err != nil {
			return fmt.Errorf("error setting the oneOf rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
		}

		var updatedRequiredFields []string
		for _, field := range requiredFields {
			if field != referenceFieldName {
				updatedRequiredFields = append(updatedRequiredFields, field)
			}
		}
		if err := crdutil.SetRequiredRuleForObjectOrArray(parent, updatedRequiredFields); err != nil {
			return fmt.Errorf("error setting the required rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
		}
	} else {
		notRule, err := crdutil.GetNotRuleForObjectOrArray(parent)
		if err != nil {
			return fmt.Errorf("error getting the not rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
		}
		// TODO(b/223688758): Handle multiple not rules.
		if notRule != nil {
			return fmt.Errorf("can't handling multiple pairs of optional mutually exclustive fields for %s in %s", referenceFieldName, crd.Name)
		}

		notRule = &apiextensions.JSONSchemaProps{
			Required: []string{nonReferenceFieldName, referenceFieldName},
		}
		if err := crdutil.SetNotRuleForObjectOrArray(parent, notRule); err != nil {
			return fmt.Errorf("error setting the not rule under path %s for CRD %s: %w", parentPath, crd.Name, err)
		}
	}

	// 4. Add the non-reference field into the schema.
	referenceFieldSchema, ok, err := crdutil.GetSchemaForFieldUnderObjectOrArray(referenceFieldName, parent)
	if err != nil {
		return fmt.Errorf("error getting schema for reference field %s under path %s for CRD %s: %w", referenceFieldName, parentPath, crd.Name, err)
	}
	if !ok {
		return fmt.Errorf("can't find reference field %s under path %s for CRD %s", referenceFieldName, parentPath, crd.Name)
	}
	fieldType := referenceFieldSchema.Type
	if fieldType != "object" && fieldType != "array" {
		return fmt.Errorf("wrong type for reference field %s under path %s for CRD %s: %s", referenceFieldName, parentPath, crd.Name, fieldType)
	}

	var nonReferenceFieldSchema *apiextensions.JSONSchemaProps
	description := fmt.Sprintf("DEPRECATED. Although this field is still available, there is limited support. "+
		"We recommend that you use `%s.%s` instead.", parentPathStr, referenceFieldName)
	if fieldType == "object" {
		nonReferenceFieldSchema = &apiextensions.JSONSchemaProps{
			Description: description,
			// When the type of a reference field is object, it means that the
			// original field type is string.
			Type: "string",
		}
	} else if fieldType == "array" {
		nonReferenceFieldSchema = &apiextensions.JSONSchemaProps{
			Description: description,
			// When the type of a reference field is an array, it means that the
			// original field type is an array of strings.
			Type: "array",
			Items: &apiextensions.JSONSchemaPropsOrArray{
				Schema: &apiextensions.JSONSchemaProps{
					Type: "string",
				},
			},
		}
	}

	if err := crdutil.SetSchemaForFieldUnderObjectOrArray(nonReferenceFieldName, parent, nonReferenceFieldSchema); err != nil {
		return fmt.Errorf("error setting schema for non-reference field %s under path %s for CRD %s: %w", nonReferenceFieldName, parentPath, crd.Name, err)
	}

	// 5. Set the updated schema.
	updatedSchema, err := getSchemaForPath(schema, parentPath[:len(parentPath)-1])
	if err != nil {
		return fmt.Errorf("can't get schema for path '%v' in CRD %s: %w", parentPathStr, crd.Name, err)
	}
	if err := crdutil.SetSchemaForFieldUnderObjectOrArray(parentPath[len(parentPath)-1], updatedSchema, parent); err != nil {
		return fmt.Errorf("error setting updated schema for parent path '%v' for CRD %s: %w", parentPathStr, crd.Name, err)
	}

	return nil
}

// EnsureReferenceFieldIsMultiKind adds the required `kind` field under the
// reference field if the `kind` field doesn't exist.
func EnsureReferenceFieldIsMultiKind(crd *apiextensions.CustomResourceDefinition, parentPath []string, referenceFieldName string, supportedKinds []MultiKindRef) error {
	if referenceFieldName == "" {
		return fmt.Errorf("param 'referenceFieldName' must be specified")
	}

	// 1. Get the parent schema of the fields.
	schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
	var err error
	var parent *apiextensions.JSONSchemaProps
	// Prepend the top-level 'spec' field into path.
	if len(parentPath) == 0 {
		parentPath = []string{"spec"}
	} else {
		parentPath = append([]string{"spec"}, parentPath...)
	}
	parentPathStr := strings.Join(parentPath, ".")
	parent, err = getSchemaForPath(schema, parentPath)
	if err != nil {
		return fmt.Errorf("can't get schema for path '%v' in CRD %s: %w", parentPathStr, crd.Name, err)
	}

	// 2. Ensure the required `kind` subfield is supported in the reference
	// field.
	referenceFieldSchema, ok, err := crdutil.GetSchemaForFieldUnderObjectOrArray(referenceFieldName, parent)
	if err != nil {
		return fmt.Errorf("error getting schema for reference field %s under path %s for CRD %s: %w", referenceFieldName, parentPath, crd.Name, err)
	}
	if !ok {
		return fmt.Errorf("can't find reference field %s under path %s for CRD %s", referenceFieldName, parentPath, crd.Name)
	}
	fieldType := referenceFieldSchema.Type
	if fieldType != "object" && fieldType != "array" {
		return fmt.Errorf("wrong type for reference field %s under path %s for CRD %s: %s", referenceFieldName, parentPath, crd.Name, fieldType)
	}

	// If the reference field is an array, it needs to be handled differently as
	// the schema for subfields is defined under .Items.Schema.
	referenceFieldSchemaForSubfields := referenceFieldSchema
	if fieldType == "array" {
		referenceFieldSchemaForSubfields = referenceFieldSchema.Items.Schema
	}
	if _, ok := referenceFieldSchemaForSubfields.Properties["kind"]; !ok {
		externalRefSchema, ok := referenceFieldSchemaForSubfields.Properties["external"]
		if !ok {
			return fmt.Errorf("can't find external field under reference %s in CRD %s", referenceFieldName, crd.Name)
		}

		if len(supportedKinds) == 0 {
			return fmt.Errorf("there must be at least one kind specified in 'supportedKinds' list")
		}
		if len(supportedKinds) > 1 {
			externalRefSchema.Description = NewMultiKindExternalRefDescription(supportedKinds)
		}
		kindsSlice := make([]string, len(supportedKinds))
		for i, kind := range supportedKinds {
			kindsSlice[i] = kind.Kind
		}
		referenceFieldSchemaWithKind := crdboilerplate.GetMultiKindResourceReferenceSchemaBoilerplate(externalRefSchema.Description, kindsSlice)

		if fieldType == "array" {
			referenceFieldSchema.Items.Schema = referenceFieldSchemaWithKind
		} else if fieldType == "object" {
			referenceFieldSchema = referenceFieldSchemaWithKind
		}
		if err := crdutil.SetSchemaForFieldUnderObjectOrArray(referenceFieldName, parent, referenceFieldSchema); err != nil {
			return fmt.Errorf("error setting schema for reference field %s under path %s for CRD %s: %w", referenceFieldName, parentPath, crd.Name, err)
		}
	}

	// 3. Set the updated schema.
	updatedSchema, err := getSchemaForPath(schema, parentPath[:len(parentPath)-1])
	if err != nil {
		return fmt.Errorf("can't get schema for path '%v' in CRD %s: %w", parentPathStr, crd.Name, err)
	}
	if err := crdutil.SetSchemaForFieldUnderObjectOrArray(parentPath[len(parentPath)-1], updatedSchema, parent); err != nil {
		return fmt.Errorf("error setting updated schema for parent path '%v' for CRD %s: %w", parentPathStr, crd.Name, err)
	}

	return nil
}

// PruneNoOpsField removes the no-ops field from spec if specified given the field path.
// It doesn't work for sub-fields in the parent field of array type.
func PruneNoOpsField(r *k8s.Resource, path ...string) error {
	unstructured.RemoveNestedField(r.Spec, path...)
	return nil
}

// PreserveUserSpecifiedLegacyField adds the user specified legacy field back to the reconciled resource.
// The reason to preserve the legacy field is because that users may be confused when the objects they try to create are different from what they get back.
func PreserveUserSpecifiedLegacyField(original, reconciled *k8s.Resource, path ...string) error {
	vo, found, err := unstructured.NestedFieldCopy(original.Spec, path...)
	if err != nil {
		return err
	}
	if !found {
		return nil
	}

	return unstructured.SetNestedField(reconciled.Spec, vo, path...)
}

// PreserveUserSpecifiedLegacyFieldUnderSlice iterates through the specified slice/array field in the reconciled and original resource, and adds
// the user-specified non-reference field(s) back into the reconciled resource. The reason to preserve the non-reference field is that users may be
// confused when the objects they try to create are different from what they get back.
// Note: This function assumed that the order of items in the slice are the
// same in the original and reconciled resources.
func PreserveUserSpecifiedLegacyFieldUnderSlice(original, reconciled *k8s.Resource, upToSlicePath []string, path []string) error {
	originalSlice, foundOriginal, err := unstructured.NestedSlice(original.Spec, upToSlicePath...)
	if err != nil {
		return fmt.Errorf("error getting the nested slice under path %s for the original resource: %w", strings.Join(upToSlicePath, "."), err)
	}
	reconciledSlice, foundReconciled, err := unstructured.NestedSlice(reconciled.Spec, upToSlicePath...)
	if err != nil {
		return fmt.Errorf("error getting the nested slice under path %s for the reconciled resource: %w", strings.Join(upToSlicePath, "."), err)
	}
	if !foundOriginal || !foundReconciled {
		return nil
	}
	for i, v := range originalSlice {
		pathFieldValue, foundPathField, err := unstructured.NestedFieldCopy(v.(map[string]interface{}), path...)
		if err != nil {
			return fmt.Errorf("error getting the user-specified value from the path %s within the slice field: %w", strings.Join(path, "."), err)
		}
		if !foundPathField {
			continue
		}
		if err := unstructured.SetNestedField(reconciledSlice[i].(map[string]interface{}), pathFieldValue, path...); err != nil {
			return fmt.Errorf("error setting original value to reconciled slice element: %w", err)
		}
	}
	if err := unstructured.SetNestedSlice(reconciled.Spec, reconciledSlice, upToSlicePath...); err != nil {
		return fmt.Errorf("error setting preserved values back into reconciled object: %w", err)
	}
	return nil
}

// PreserveUserSpecifiedLegacyArrayField adds the user specified legacy array
// field back to the reconciled resource.
// The reason to preserve the legacy array field is because that users may be
// confused when the objects they try to create are different from what they get
// back.
func PreserveUserSpecifiedLegacyArrayField(original, reconciled *k8s.Resource, path ...string) error {
	vo, found, err := unstructured.NestedSlice(original.Spec, path...)
	if err != nil {
		return err
	}
	if !found {
		return nil
	}

	return unstructured.SetNestedSlice(reconciled.Spec, vo, path...)
}

// PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecified prune the defaulted authoritative field from the reconciled resource (post-actuation) if only
// the legacy field is specified in the original spec.
// Populating the new authoritative field into spec along with the legacy field will cause confusion if users only modify the legacy field in their configuration
// without being aware of the defaulted field in k8s object.
func PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecified(original, reconciled *k8s.Resource, legacyFieldPath, fieldPath []string) error {
	// If the authoritative field is specified in the original spec, do nothing.
	_, found, err := unstructured.NestedFieldCopy(original.Spec, fieldPath...)
	if err != nil {
		return err
	}
	if found {
		return nil
	}
	// If only the legacy field is specified in the original spec, prune the defaulted authoritative field.
	_, found, err = unstructured.NestedFieldCopy(original.Spec, legacyFieldPath...)
	if err != nil {
		return err
	}
	if found {
		unstructured.RemoveNestedField(reconciled.Spec, fieldPath...)
		return nil
	}
	return nil
}

// PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecifiedUnderSlice iterates through the specified slice/array field in the reconciled and original resource,
// and prune the defaulted reference field from the reconciled resource (post-actuation) if only the non-reference field is specified in the original spec.
// Populating the new reference field into spec along with the non-reference field will cause confusion if users only modify the non-reference field in their configuration
// without being aware of the defaulted field in k8s object.
// Note: This function assumed that the order of items in the slice are the
// same in the original and reconciled resources.
func PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecifiedUnderSlice(original, reconciled *k8s.Resource, pathUpToSlice, nonReferenceFieldPath, referenceFieldPath []string) error {
	originalSlice, foundOriginal, err := unstructured.NestedSlice(original.Spec, pathUpToSlice...)
	if err != nil {
		return fmt.Errorf("error getting the nested slice under path %s for the original resource: %w", strings.Join(pathUpToSlice, "."), err)
	}
	reconciledSlice, foundReconciled, err := unstructured.NestedSlice(reconciled.Spec, pathUpToSlice...)
	if err != nil {
		return fmt.Errorf("error getting the nested slice under path %s for the reconciled resource: %w", strings.Join(pathUpToSlice, "."), err)
	}
	if !foundOriginal || !foundReconciled {
		return nil
	}
	for i, v := range originalSlice {
		_, foundReferenceField, err := unstructured.NestedFieldCopy(v.(map[string]interface{}), referenceFieldPath...)
		if err != nil {
			return fmt.Errorf("error checking the existence of the nested reference field %s within the slice field of the original resource: %w", strings.Join(referenceFieldPath, "."), err)
		}
		_, foundNonReferenceField, err := unstructured.NestedFieldCopy(v.(map[string]interface{}), nonReferenceFieldPath...)
		if err != nil {
			return fmt.Errorf("error checking the existence of the nested non-reference field %s within the slice field of the original resource: %w", strings.Join(nonReferenceFieldPath, "."), err)
		}
		if !foundReferenceField && foundNonReferenceField {
			unstructured.RemoveNestedField(reconciledSlice[i].(map[string]interface{}), referenceFieldPath...)
		}
	}
	if err := unstructured.SetNestedSlice(reconciled.Spec, reconciledSlice, pathUpToSlice...); err != nil {
		return fmt.Errorf("error setting the altered slice field %s back into the reconciled resource: %w", strings.Join(pathUpToSlice, "."), err)
	}
	return nil
}

// PruneDefaultedAuthoritativeArrayFieldIfOnlyLegacyArrayFieldSpecified prunes
// the defaulted authoritative array field from the reconciled resource
// (post-actuation) if only the legacy array field is specified in the original
// spec.
// Populating the new authoritative array field into spec along with the legacy
// array field will cause confusion if users only modify the legacy array field
// in their configuration without being aware of the defaulted field in k8s
// object.
func PruneDefaultedAuthoritativeArrayFieldIfOnlyLegacyArrayFieldSpecified(original, reconciled *k8s.Resource, legacyFieldPath, fieldPath []string) error {
	// If the authoritative field is specified in the original spec, do nothing.
	_, found, err := unstructured.NestedSlice(original.Spec, fieldPath...)
	if err != nil {
		return err
	}
	if found {
		return nil
	}

	// If only the legacy field is specified in the original spec, prune the
	// defaulted authoritative field.
	_, found, err = unstructured.NestedSlice(original.Spec, legacyFieldPath...)
	if err != nil {
		return err
	}
	if found {
		unstructured.RemoveNestedField(reconciled.Spec, fieldPath...)
	}
	return nil
}

// FavorAuthoritativeFieldOverLegacyField favor the value of the authoritative field if it's set;
// otherwise, it takes the value from the legacy field and populate it into the authoritative field and then prune the legacy field.
// If the legacy field is specified, this function will also mark the authoritative field as managed fields.
func FavorAuthoritativeFieldOverLegacyField(r *k8s.Resource, legacyFieldPath, fieldPath []string) error {
	if err := validateIfAuthoritativeFieldAndLegacyFieldTakeDifferentValues(r, legacyFieldPath, fieldPath); err != nil {
		return err
	}

	// If the authoritative field is set, keep the value.
	_, found, err := unstructured.NestedFieldCopy(r.Spec, fieldPath...)
	if err != nil {
		return err
	}
	if found {
		unstructured.RemoveNestedField(r.Spec, legacyFieldPath...)
		return nil
	}

	// If only the legacy field is set, populate the value to the authoritative field before actuation.
	v, found, err := unstructured.NestedFieldCopy(r.Spec, legacyFieldPath...)
	if err != nil {
		return err
	}
	if !found {
		return nil
	}
	if isReferenceFieldPath(fieldPath) {
		if err := unstructured.SetNestedField(r.Spec, v, append(fieldPath, "external")...); err != nil {
			return err
		}
		// Mark "external" under the authoritative field as user managed fields because the user has set the legacy field for the same feature.
		if err := markFieldAsManaged(r, append(fieldPath, "external")...); err != nil {
			return err
		}
		// Mark the authoritative field as user managed fields because the user has set the legacy field for the same feature.
		if err := markFieldAsManaged(r, fieldPath...); err != nil {
			return err
		}
		unstructured.RemoveNestedField(r.Spec, legacyFieldPath...)
		return nil
	}
	if err := unstructured.SetNestedField(r.Spec, v, fieldPath...); err != nil {
		return err
	}
	// Mark the authoritative field as user managed fields because the user has set the legacy field for the same feature.
	if err := markFieldAsManaged(r, fieldPath...); err != nil {
		return err
	}
	unstructured.RemoveNestedField(r.Spec, legacyFieldPath...)
	return nil
}

// isReferenceFieldPath will only identify non-array reference fields (reference array fields end with "Refs")
func isReferenceFieldPath(fieldPath []string) bool {
	field := fieldPath[len(fieldPath)-1]
	return strings.HasSuffix(field, "Ref")
}

// FavorReferenceFieldOverNonReferenceFieldUnderSlice returns an error if both fields are set; otherwise, take the value
// from the non-reference field and populate it into the "external" field in the reference field, and then prune the non-reference field.
func FavorReferenceFieldOverNonReferenceFieldUnderSlice(r *k8s.Resource, pathUpToSlice, nonReferenceFieldPath, referenceFieldPath []string) error {
	if err := validateAtMostOneFieldIsSetUnderSlice(r, pathUpToSlice, nonReferenceFieldPath, referenceFieldPath); err != nil {
		return err
	}
	sliceVal, found, err := unstructured.NestedSlice(r.Spec, pathUpToSlice...)
	if err != nil {
		return fmt.Errorf("error getting nested slice field %s from resource: %w", strings.Join(pathUpToSlice, "."), err)
	}
	if !found {
		return nil
	}
	for i, v := range sliceVal {
		nonReferenceVal, found, err := unstructured.NestedFieldCopy(v.(map[string]interface{}), nonReferenceFieldPath...)
		if err != nil {
			return fmt.Errorf("error getting non-reference field %s from slice field: %w", strings.Join(nonReferenceFieldPath, "."), err)
		}
		if !found {
			continue
		}
		if err := unstructured.SetNestedField(sliceVal[i].(map[string]interface{}), nonReferenceVal, append(referenceFieldPath, "external")...); err != nil {
			return fmt.Errorf("error setting non-reference value to reference field path %s: %w", strings.Join(append(referenceFieldPath, "external"), "."), err)
		}
		unstructured.RemoveNestedField(sliceVal[i].(map[string]interface{}), nonReferenceFieldPath...)
	}
	if err := unstructured.SetNestedSlice(r.Spec, sliceVal, pathUpToSlice...); err != nil {
		return fmt.Errorf("error setting altered slice %s back into resource: %w", strings.Join(pathUpToSlice, "."), err)
	}
	return nil
}

// validateIfAuthoritativeFieldAndLegacyFieldTakeDifferentValues returns error if the legacy field and the authoritative field are both present in spec
// but with different values.
func validateIfAuthoritativeFieldAndLegacyFieldTakeDifferentValues(r *k8s.Resource, legacyFieldPath, fieldPath []string) error {
	v1, f1, err := unstructured.NestedFieldCopy(r.Spec, fieldPath...)
	if err != nil {
		return err
	}

	v2, f2, err := unstructured.NestedFieldCopy(r.Spec, legacyFieldPath...)
	if err != nil {
		return err
	}

	if f1 && f2 && !reflect.DeepEqual(v1, v2) {
		authoritative := strings.Join(fieldPath, ".")
		legacy := strings.Join(legacyFieldPath, ".")
		return fmt.Errorf("'%v' field and '%v' field are both present in spec, but they take different values. It's recommended to only use '%v' in your configuration because '%v' has been deprecated by the API",
			authoritative, legacy, authoritative, legacy)
	}
	return nil
}

// validateAtMostOneFieldIsSetUnderSlice returns error if both field paths within the slice field are set
func validateAtMostOneFieldIsSetUnderSlice(r *k8s.Resource, fieldPathUpToSlice, fieldPath1, fieldPath2 []string) error {
	sliceField, found, err := unstructured.NestedSlice(r.Spec, fieldPathUpToSlice...)
	if err != nil {
		return fmt.Errorf("error getting slice field %s from resource: %w", strings.Join(fieldPathUpToSlice, "."), err)
	}
	if !found {
		return nil
	}
	for _, v := range sliceField {
		_, found1, err := unstructured.NestedFieldCopy(v.(map[string]interface{}), fieldPath1...)
		if err != nil {
			return fmt.Errorf("error checking existence of nested field %s within slice field: %w", strings.Join(fieldPath1, "."), err)
		}
		_, found2, err := unstructured.NestedFieldCopy(v.(map[string]interface{}), fieldPath2...)
		if err != nil {
			return fmt.Errorf("error checking existence of nested field %s within slice field: %w", strings.Join(fieldPath2, "."), err)
		}
		if !found1 || !found2 {
			continue
		}
		fullFieldPath1 := strings.Join(append(fieldPathUpToSlice, fieldPath1...), ".")
		fullFieldPath2 := strings.Join(append(fieldPathUpToSlice, fieldPath2...), ".")
		return fmt.Errorf("'%v' field and '%v' field are both set. Please remove one of the two fields", fullFieldPath1, fullFieldPath2)
	}
	return nil
}

func markFieldAsManaged(r *k8s.Resource, fieldPath ...string) error {
	if r.ManagedFields == nil {
		return nil
	}
	parts := make([]interface{}, 0, len(fieldPath))
	for _, v := range fieldPath {
		parts = append(parts, v)
	}
	p, err := fieldpath.MakePath(parts...)
	if err != nil {
		return err
	}
	r.ManagedFields.Insert(p)
	return nil
}

// FavorReferenceArrayFieldOverNonReferenceArrayField favor the value of the
// reference array field if it's set; otherwise, it takes the value from the
// non-reference array field, populates it into the 'external' subfield of the
// items in the reference field, and then prune the non-reference field.
func FavorReferenceArrayFieldOverNonReferenceArrayField(r *k8s.Resource, nonReferenceFieldPath, referenceFieldPath []string) error {
	nonRefArray, foundNonRef, err := unstructured.NestedStringSlice(r.Spec, nonReferenceFieldPath...)
	if err != nil {
		return fmt.Errorf("error getting the non-reference field '%v': %w", strings.Join(nonReferenceFieldPath, "."), err)
	}
	_, foundRef, err := unstructured.NestedSlice(r.Spec, referenceFieldPath...)
	if err != nil {
		return fmt.Errorf("error getting the reference field '%v': %w", strings.Join(referenceFieldPath, "."), err)
	}

	if foundNonRef && foundRef {
		return fmt.Errorf("mutually-exclusive fields '%v' and '%v' are both set", strings.Join(nonReferenceFieldPath, "."), strings.Join(referenceFieldPath, "."))
	}

	if !foundNonRef || len(nonRefArray) == 0 {
		return nil
	}

	// If only the non-reference array is set, populate its values to the
	// 'external' subfields of items under the reference array, and remove the
	// non-reference array field.
	refArray := make([]interface{}, len(nonRefArray))
	for i, val := range nonRefArray {
		refItem := make(map[string]interface{})
		if err := unstructured.SetNestedField(refItem, val, "external"); err != nil {
			return fmt.Errorf("error setting the 'external' field under the reference array '%v': %w", strings.Join(referenceFieldPath, "."), err)
		}
		refArray[i] = refItem
	}
	if err := unstructured.SetNestedSlice(r.Spec, refArray, referenceFieldPath...); err != nil {
		return err
	}
	unstructured.RemoveNestedField(r.Spec, nonReferenceFieldPath...)
	return nil
}

// removePrefixFromStringFieldInSpec removes the prefix from the field specified by
// the input path in the resource spec.
// The function returns error if the specified field is not a string.
// The function is no-op if the field is not found in resource spec.
// The function is no-op if the input prefix does not match the prefix of the specified field.
func removePrefixFromStringFieldInSpec(r *k8s.Resource, prefix string, path ...string) error {
	vo, found, err := unstructured.NestedString(r.Spec, path...)
	if err != nil {
		return err
	}
	if !found {
		return nil
	}
	if !strings.HasPrefix(vo, prefix) {
		return nil
	}
	v := strings.TrimPrefix(vo, prefix)

	return unstructured.SetNestedField(r.Spec, v, path...)
}

// fieldPath is the path to a Terraform field.
type fieldPath []string

// This function ensures backward compatibility on 'projectRef.external' field for resources that are
// migrated from DCL-based to TF-based implementation.
// DCL-based implementation uses format "projects/${PROJECT_ID?}".
// TF-based implementation may use format "${PROJECT_ID?}" or/and "projects/${PROJECT_ID?}".
// To ensure users can still use "projects/${PROJECT_ID?}" when the TF-based implementation only supports format "${PROJECT_ID?}",
// we need to trim the "projects/" prefix before TF actuation,
// and add back the prefix after TF actuation to preserve the user specified value.
func handleProjectsPrefixInProjectRefExternalFields(fps []fieldPath) ResourceOverride {
	o := ResourceOverride{}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		for _, fp := range fps {
			if err := removePrefixFromStringFieldInSpec(r, "projects/", fp...); err != nil {
				return fmt.Errorf("error removing 'projects/' prefix from field %v in pre-actuation transformation: %w", strings.Join(fp, "."), err)
			}
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		for _, fp := range fps {
			if err := PreserveUserSpecifiedLegacyField(original, reconciled, fp...); err != nil {
				return fmt.Errorf("error preserving field %v in post-actuation transformation: %w", strings.Join(fp, "."), err)
			}
		}
		return nil
	}
	// After status update, the http response is decoded to the original resource object, reverting the prefix removal.
	// So we need to remove the prefix again to ensure the prefix is removed.
	o.PostUpdateStatusTransform = func(r *k8s.Resource) error {
		for _, fp := range fps {
			if err := removePrefixFromStringFieldInSpec(r, "projects/", fp...); err != nil {
				return fmt.Errorf("error removing 'projects/' prefix from field %v in post-status update transformation: %w", strings.Join(fp, "."), err)
			}
		}
		return nil
	}
	return o
}

func NewMultiKindExternalRefDescription(supportedKinds []MultiKindRef) string {
	var description strings.Builder
	description.WriteString("Allowed values:")
	for _, ref := range supportedKinds {
		fmt.Fprintf(&description, " The `%s` field of a `%s` resource.", ref.TargetField, ref.Kind)
	}
	return description.String()
}
