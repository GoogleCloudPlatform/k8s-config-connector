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

package conversion

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/pathslice"

	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// A Converter knows how to convert between KRM and DCL format.
type Converter struct {
	SchemaLoader   dclschemaloader.DCLSchemaLoader
	MetadataLoader dclmetadata.ServiceMetadataLoader
}

// New returns a Converter.
func New(schemaLoader dclschemaloader.DCLSchemaLoader, metadataLoader dclmetadata.ServiceMetadataLoader) *Converter {
	c := &Converter{}
	c.SchemaLoader = schemaLoader
	c.MetadataLoader = metadataLoader
	return c
}

// KRMObjectToDCLObject converts the given KCC Lite KRM resource to a DCL resource.
func (c *Converter) KRMObjectToDCLObject(obj *unstructured.Unstructured) (*dclunstruct.Resource, error) {
	gvk := obj.GroupVersionKind()
	stv, err := dclmetadata.ToServiceTypeVersion(gvk, c.MetadataLoader)
	if err != nil {
		return nil, fmt.Errorf("error resolving the DCL ServiceTypeVersion from GroupVersionKind %v: %w", gvk, err)
	}
	resourceMetadata, found := c.MetadataLoader.GetResourceWithGVK(gvk)
	if !found {
		return nil, fmt.Errorf("ServiceMetadata for resource with GroupVersionKind %v not found", gvk)
	}
	// load DCL schema per DCL ServiceTypeVersion
	dclSchema, err := c.SchemaLoader.GetDCLSchema(stv)
	if err != nil {
		return nil, fmt.Errorf("error getting the DCL schema for ServiceTypeVersion %v: %w", stv, err)
	}
	if dclSchema.Type != "object" {
		return nil, fmt.Errorf("expect the entry level DCL OpenAPI schema to be 'object' type, but got %v", dclSchema.Type)
	}

	r := &dclunstruct.Resource{
		STV: stv,
	}
	// Only convert spec fields for now since output-only fields are not enforceable for DCL.
	spec := obj.UnstructuredContent()["spec"]
	if spec == nil {
		spec = make(map[string]interface{})
	}
	dclObj, err := convertToDCL(spec, []string{}, dclSchema, c.MetadataLoader, false)
	if err != nil {
		return nil, fmt.Errorf("error converting the spec of resource %v/%v to DCL object: %w", obj.GetNamespace(), obj.GetName(), err)
	}
	dclObjMap, ok := dclObj.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the converted DCL object to be map[string]interface{} but was actually %T", dclObj)
	}
	r.Object = dclObjMap
	// special handling on name, container and labels fields
	if err := convertToDCLNameField(obj, r, dclSchema); err != nil {
		return nil, fmt.Errorf("error resolving the value of 'name' field for resource %v/%v: %w", obj.GetNamespace(), obj.GetName(), err)
	}
	// TODO(b/186159460): Delete this if-block once all resources support
	// hierarchical references.
	if !resourceMetadata.SupportsHierarchicalReferences {
		if err := convertToDCLContainerField(obj, r, dclSchema); err != nil {
			return nil, fmt.Errorf("error resolving the value of the container field for resource %v/%v: %w", obj.GetNamespace(), obj.GetName(), err)
		}
	}
	if err := convertToDCLLabelsField(obj, r, dclSchema); err != nil {
		return nil, fmt.Errorf("error converting Kubernetes labels into DCL labels for resource %v/%v: %w", obj.GetNamespace(), obj.GetName(), err)
	}
	return r, nil
}

// DCLObjectToKRMObject converts the given DCL resource to a KCC Lite KRM resource.
func (c *Converter) DCLObjectToKRMObject(resource *dclunstruct.Resource) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	gvk, err := dclmetadata.ToGroupVersionKind(resource.STV, c.MetadataLoader)
	if err != nil {
		return nil, fmt.Errorf("error resolving GroupVersionKind from the DCL ServiceTypeVersion %v: %w", resource.STV, err)
	}
	obj.SetGroupVersionKind(gvk)
	resourceMetadata, found := c.MetadataLoader.GetResourceWithGVK(gvk)
	if !found {
		return nil, fmt.Errorf("ServiceMetadata for resource with GroupVersionKind %v not found", gvk)
	}
	dclSchema, err := c.SchemaLoader.GetDCLSchema(resource.STV)
	if err != nil {
		return nil, fmt.Errorf("error getting the DCL schema for ServiceTypeVersion %v", resource.STV)
	}
	if dclSchema.Type != "object" {
		return nil, fmt.Errorf("expect the entry level DCL OpenAPI schema to be object type, but got %v", dclSchema.Type)
	}

	// ensure DCL-returned state contains no nil values
	dcl.TrimNilFields(resource.Object)
	// convert dcl state to spec and status
	spec, err := convertToKRMSpec(resource.Object, []string{}, dclSchema, c.MetadataLoader, resourceMetadata, false)
	if err != nil {
		return nil, fmt.Errorf("error extracting the spec from the DCL resource %v/%v: %w", obj.GetNamespace(), obj.GetName(), err)
	}
	if spec != nil {
		specMap, ok := spec.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the converted spec to be map[string]interface{} but was actually %T", spec)
		}
		if len(specMap) > 0 {
			obj.Object["spec"] = specMap
		}
	}
	status, err := convertToKRMStatus(resource.Object, dclSchema)
	if err != nil {
		return nil, fmt.Errorf("error extracting the status from the DCL resource %v/%v: %w", obj.GetNamespace(), obj.GetName(), err)
	}
	if len(status) > 0 {
		obj.Object["status"] = status
	}
	// special handling on name, container and labels fields
	if err := liftDCLLabelsField(obj, dclSchema); err != nil {
		return nil, fmt.Errorf("error lifting 'labels' field to metadata.labels: %w", err)
	}
	// TODO(b/186159460): Delete this if-block once all resources support
	// hierarchical references.
	if !resourceMetadata.SupportsHierarchicalReferences {
		if err := liftDCLContainerField(obj, dclSchema); err != nil {
			return nil, fmt.Errorf("error lifting container field to annotation: %w", err)
		}
	}
	if err := convertToKRMResourceIDField(obj, dclSchema); err != nil {
		return nil, fmt.Errorf("error converting 'name' field to 'resourceID': %w", err)
	}
	return obj, nil
}

func convertToKRMSpec(val interface{}, path []string, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader,
	resourceMetadata dclmetadata.Resource, isCollectionItemSchema bool) (interface{}, error) {
	if val == nil {
		return nil, nil
	}
	switch schema.Type {
	case "object":
		obj, ok := val.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the value to be map[string]interface{}, but was actually %T", val)
		}
		res := make(map[string]interface{})
		// The field additionalProperties is mutually exclusive with properties in CustomResourceDefinition.
		if schema.AdditionalProperties != nil {
			for k, v := range obj {
				convertedVal, err := convertToKRMSpec(v, append(path, k), schema.AdditionalProperties, smLoader, resourceMetadata, true)
				if err != nil {
					return nil, err
				}
				dcl.AddToMap(k, convertedVal, res)
			}
			return res, nil
		}
		for field, fieldSchema := range schema.Properties {
			if fieldSchema.ReadOnly && !isCollectionItemSchema {
				continue
			}
			val, ok := obj[field]
			if !ok || val == nil {
				continue
			}
			isSensitive, err := extension.IsSensitiveField(fieldSchema)
			if err != nil {
				return nil, err
			}
			if !fieldSchema.ReadOnly && isSensitive {
				convertedVal, err := convertSensitiveFieldToKRM(val)
				if err != nil {
					return nil, fmt.Errorf("error resolving the value for sensitive field %v: %w", field, err)
				}
				dcl.AddToMap(field, convertedVal, res)
				continue
			}
			// For resources that don't support hierarchical references,
			// convert the container field to a primitive field in the spec.
			// This field will later be converted to an annotation by
			// liftDCLContainerField().
			// TODO(b/186159460): Delete this if-block once all resources support
			// hierarchical references.
			if !fieldSchema.ReadOnly && dcl.IsContainerField(append(path, field)) && !resourceMetadata.SupportsHierarchicalReferences {
				convertedVal, err := convertToKRMSpec(obj[field], append(path, field), fieldSchema, smLoader, resourceMetadata, isCollectionItemSchema)
				if err != nil {
					return nil, fmt.Errorf("error resolving the value for DCL field %v: %w", field, err)
				}
				dcl.AddToMap(field, convertedVal, res)
				continue
			}
			if !fieldSchema.ReadOnly && extension.IsReferenceField(fieldSchema) {
				refField, convertedVal, err := convertReferenceFieldToKRM(append(path, field), obj[field], fieldSchema, smLoader)
				if err != nil {
					return nil, fmt.Errorf("error converting the value for reference field %v: %w", field, err)
				}
				dcl.AddToMap(refField, convertedVal, res)
				continue
			}
			convertedVal, err := convertToKRMSpec(obj[field], append(path, field), fieldSchema, smLoader, resourceMetadata, isCollectionItemSchema)
			if err != nil {
				return nil, fmt.Errorf("error resolving the value for DCL field %v: %w", field, err)
			}
			dcl.AddToMap(field, convertedVal, res)
		}
		return res, nil
	case "array":
		items, ok := val.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the value to be array but was actually %T", val)
		}
		if len(items) == 0 {
			return nil, nil
		}
		res := make([]interface{}, 0)
		for _, item := range items {
			processedItem, err := convertToKRMSpec(item, path, schema.Items, smLoader, resourceMetadata, true)
			if err != nil {
				return nil, fmt.Errorf("error converting list item: %w", err)
			}
			res = append(res, processedItem)
		}
		return res, nil
	case "string", "boolean", "number", "integer":
		return val, nil
	default:
		return nil, fmt.Errorf("unknown schema type %v", schema.Type)
	}
}

func getStatusFieldsWithValuePopulated(path string, val interface{}, schema *openapi.Schema, paths []string) ([]string, error) {
	if val == nil {
		return paths, nil
	}
	if schema.ReadOnly {
		paths = append(paths, path)
		return paths, nil
	}
	if schema.Type == "object" {
		obj, ok := val.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the value to be map[string]interface{} but was actually %T", val)
		}
		if len(obj) == 0 {
			return paths, nil
		}
		if schema.AdditionalProperties != nil {
			return getStatusFieldsWithValuePopulated(path, obj, schema.AdditionalProperties, paths)
		}
		var err error
		for k, v := range obj {
			fieldSchema, ok := schema.Properties[k]
			if !ok {
				continue
			}
			qualifiedName := k
			if path != "" {
				qualifiedName = path + "." + k
			}
			paths, err = getStatusFieldsWithValuePopulated(qualifiedName, v, fieldSchema, paths)
			if err != nil {
				return nil, fmt.Errorf("error getting status fields with value from %v: %w", qualifiedName, err)
			}
		}
		return paths, nil
	}
	return paths, nil
}

func convertToKRMStatus(obj map[string]interface{}, schema *openapi.Schema) (map[string]interface{}, error) {
	paths := make([]string, 0)
	paths, err := getStatusFieldsWithValuePopulated("", obj, schema, paths)
	if err != nil {
		return nil, fmt.Errorf("error getting status fields from DCL object: %w", err)
	}
	status := make(map[string]interface{})
	for _, path := range paths {
		val, found, err := unstructured.NestedFieldCopy(obj, strings.Split(path, ".")...)
		if err != nil {
			return nil, fmt.Errorf("error copying the value for status field %v: %w", path, err)
		}
		if !found {
			return nil, fmt.Errorf("couldn't find the value for status field %v", path)
		}
		splitPath := strings.Split(path, ".")
		splitPath = renameStatusFieldIfCollidesWithReservedName(splitPath)
		if err := unstructured.SetNestedField(status, val, splitPath...); err != nil {
			return nil, fmt.Errorf("error setting the value for status field %v: %w", path, err)
		}
	}
	return status, nil
}

func convertToDCL(val interface{}, path []string, schema *openapi.Schema,
	smLoader dclmetadata.ServiceMetadataLoader, isCollectionItemSchema bool) (interface{}, error) {
	if val == nil {
		return nil, nil
	}
	switch schema.Type {
	case "object":
		obj, ok := val.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the value to be map[string]interface{} but was actually %T", val)
		}
		// The field additionalProperties is mutually exclusive with properties in CustomResourceDefinition.
		if schema.AdditionalProperties != nil {
			res := make(map[string]interface{})
			for k, v := range obj {
				dclVal, err := convertToDCL(v, append(path, k), schema.AdditionalProperties, smLoader, true)
				if err != nil {
					return nil, fmt.Errorf("error converting item with AdditionalProperties schema: %w", err)
				}
				dcl.AddToMap(k, dclVal, res)
			}
			return res, nil
		}
		res := make(map[string]interface{})
		for field, fieldSchema := range schema.Properties {
			// Avoid dropping read-only fields from objects in collections
			// (maps, arrays) to avoid the possibility of ending up with empty
			// objects, as this would result in diffs given that the live value
			// of the object is non-empty. Note: we only need to do this for
			// objects in collections as that is the only case where read-only
			// fields can end up in the spec instead of status.
			if fieldSchema.ReadOnly && !isCollectionItemSchema {
				continue
			}

			// It is expected that convertToDCL() will be first called with an
			// entry-level DCL schema, which is always an object, and then be
			// called in the recursion with non-entry-level schemas. So we only
			// check sensitive fields here (instead of checking it in the string
			// type) before calling convertToDCL() recursively to process field
			// generically.
			isSensitive, err := extension.IsSensitiveField(fieldSchema)
			if err != nil {
				return nil, err
			}
			if !fieldSchema.ReadOnly && isSensitive {
				convertedVal, err := convertSensitiveFieldToDCL(append(path, field), obj)
				if err != nil {
					return nil, fmt.Errorf("error resolving the value for sensitive field %v: %w", field, err)
				}
				dcl.AddToMap(field, convertedVal, res)
				continue
			}
			if !fieldSchema.ReadOnly && extension.IsReferenceField(fieldSchema) {
				convertedVal, err := convertReferenceFieldToDCL(append(path, field), obj, fieldSchema, smLoader)
				if err != nil {
					return nil, fmt.Errorf("error resolving the value for reference field %v: %w", field, err)
				}
				dcl.AddToMap(field, convertedVal, res)
				continue
			}
			convertedVal, err := convertToDCL(obj[field], append(path, field), fieldSchema, smLoader, isCollectionItemSchema)
			if err != nil {
				return nil, fmt.Errorf("error resolving the value for DCL field %v: %w", field, err)
			}
			dcl.AddToMap(field, convertedVal, res)
		}
		return res, nil
	case "array":
		items, ok := val.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the value to be []interface{} but was actually %T", val)
		}
		res := make([]interface{}, 0)
		for _, item := range items {
			processedItem, err := convertToDCL(item, path, schema.Items, smLoader, true)
			if err != nil {
				return nil, fmt.Errorf("error converting the item: %w", err)
			}
			res = append(res, processedItem)
		}
		return res, nil
	case "integer":
		return dcl.CanonicalizeIntegerValue(val)
	case "number":
		return dcl.CanonicalizeNumberValue(val)
	case "string", "boolean":
		return val, nil
	default:
		return nil, fmt.Errorf("unknown schema type %v", schema.Type)
	}
}

func convertSensitiveFieldToDCL(path []string, obj map[string]interface{}) (interface{}, error) {
	field := pathslice.Base(path)
	if obj[field] == nil {
		return nil, nil
	}
	raw := obj[field]
	secretRef, ok := raw.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the value to be map[string]interface{} for field %v but was actually %T", field, raw)
	}
	return secretRef["value"], nil
}

func convertSensitiveFieldToKRM(val interface{}) (interface{}, error) {
	if val == nil {
		return nil, nil
	}
	sensitiveFieldStruct := make(map[string]interface{})
	sensitiveFieldStruct["value"] = val
	return sensitiveFieldStruct, nil
}

func convertReferenceFieldToDCL(path []string, obj map[string]interface{}, schema *openapi.Schema,
	smLoader dclmetadata.ServiceMetadataLoader) (interface{}, error) {
	if dcl.IsMultiTypeParentReferenceField(path) {
		return convertMultiTypeParentReferenceFieldToDCL(obj, schema, smLoader)
	}
	if schema.Type == "array" {
		return convertListOfReferencesFieldToDCL(path, obj, schema)
	}
	return convertRegularReferenceFieldToDCL(path, obj, schema)
}

func convertMultiTypeParentReferenceFieldToDCL(obj map[string]interface{}, schema *openapi.Schema,
	smLoader dclmetadata.ServiceMetadataLoader) (interface{}, error) {
	rawVal, tc, err := dcl.GetHierarchicalRefFromConfigForMultiParentResource(obj, schema, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting hierarchical reference from config for multi-parent resource: %w", err)
	}
	if rawVal == nil {
		return nil, fmt.Errorf("no hierarchical reference found for multi-parent resource")
	}
	refField := tc.Key
	refObj, ok := rawVal.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the value to be map[string]interface{} for reference field %v but was actually %T", refField, rawVal)
	}

	// Prefix the 'external' value with the parent prefix (e.g. "projects/")
	// since that is how DCL distinguishes between different types of parents
	// for multi-type parent reference fields.
	rawExternalVal, ok := refObj["external"]
	if !ok {
		return nil, fmt.Errorf("'external' was unexpectedly not set for reference field %v", refField)
	}
	externalVal, ok := rawExternalVal.(string)
	if !ok {
		return nil, fmt.Errorf("expected the value of 'external' to be string for reference field %v but was actually %T", refField, rawExternalVal)
	}
	if externalVal == "" {
		return externalVal, nil
	}
	parentPrefix := dcl.ParentPrefixForKind(tc.GVK.Kind)
	if strings.HasPrefix(externalVal, parentPrefix) {
		return externalVal, nil
	}
	return fmt.Sprintf("%v%v", parentPrefix, externalVal), nil
}

func convertListOfReferencesFieldToDCL(path []string, obj map[string]interface{}, schema *openapi.Schema) (interface{}, error) {
	refField, err := extension.GetReferenceFieldName(path, schema)
	if err != nil {
		return nil, fmt.Errorf("error getting the reference field name %w", err)
	}
	if obj[refField] == nil {
		return nil, nil
	}
	rawVal := obj[refField]
	items, ok := rawVal.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the value to be []interface{} for reference field %v but was actually %T", refField, rawVal)
	}
	res := make([]interface{}, 0)
	for _, item := range items {
		refObj, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the value for the reference to be map[string]interface{}, but was actually %T", item)
		}
		refVal, err := resolveReferenceValue(refObj, schema.Items)
		if err != nil {
			return nil, err
		}
		res = append(res, refVal)
	}
	return res, nil
}

func convertRegularReferenceFieldToDCL(path []string, obj map[string]interface{}, schema *openapi.Schema) (interface{}, error) {
	refField, err := extension.GetReferenceFieldName(path, schema)
	if err != nil {
		return nil, fmt.Errorf("error getting the reference field name %w", err)
	}
	if obj[refField] == nil {
		return nil, nil
	}
	rawVal := obj[refField]
	refObj, ok := rawVal.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the value to be map[string]interface{} for reference field %v but was actually %T", refField, rawVal)
	}
	return resolveReferenceValue(refObj, schema)
}

func convertReferenceFieldToKRM(path []string, val interface{}, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader) (string, interface{}, error) {
	if dcl.IsMultiTypeParentReferenceField(path) {
		return convertMultiTypeParentReferenceFieldToKRM(path, val, schema, smLoader)
	}
	if schema.Type == "array" {
		return convertListOfReferencesFieldToKRM(path, val, schema)
	}
	return convertRegularReferenceToKRM(path, val, schema)
}

func convertMultiTypeParentReferenceFieldToKRM(path []string, val interface{}, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader) (string, interface{}, error) {
	field := pathslice.Base(path)
	tcs, err := dcl.GetReferenceTypeConfigs(schema, smLoader)
	if err != nil {
		return "", nil, fmt.Errorf("error getting reference type configs for DCL field '%v': %w", field, err)
	}
	v, ok := val.(string)
	if !ok {
		return "", nil, fmt.Errorf("expected the value to be string for DCL field '%v' but was actually %T", field, val)
	}
	if v == "" {
		return "", nil, fmt.Errorf("value of DCL field '%v' is unexpectedly an empty string", field)
	}
	for _, tc := range tcs {
		// Multi-type parent reference fields in DCL use parent prefixes (e.g.
		// "projects/") to denote the type of resource being referenced.
		if strings.HasPrefix(v, dcl.ParentPrefixForKind(tc.GVK.Kind)) {
			convertedVal := map[string]interface{}{
				"external": v,
			}
			return tc.Key, convertedVal, nil
		}
	}
	return "", nil, fmt.Errorf("value for DCL field %v could not be recognized as a valid parent: %v", field, v)
}

func convertListOfReferencesFieldToKRM(path []string, val interface{}, schema *openapi.Schema) (string, interface{}, error) {
	field := pathslice.Base(path)
	refField, err := extension.GetReferenceFieldName(path, schema)
	if err != nil {
		return "", nil, fmt.Errorf("error getting the reference field name %w", err)
	}
	items, ok := val.([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("expected the value to be []interface{} for reference field %v but was actually %T", refField, val)
	}
	res := make([]interface{}, 0)
	for _, item := range items {
		convertedVal, err := convertReferenceValueToKRM(item, schema.Items)
		if err != nil {
			return "", nil, err
		}
		res = append(res, convertedVal)
	}
	return field, res, nil
}

func convertRegularReferenceToKRM(path []string, val interface{}, schema *openapi.Schema) (string, interface{}, error) {
	field := pathslice.Base(path)
	refField, err := extension.GetReferenceFieldName(path, schema)
	if err != nil {
		return "", nil, fmt.Errorf("error getting the reference field name %w", err)
	}
	convertedVal, err := convertReferenceValueToKRM(val, schema)
	if err != nil {
		return "", nil, fmt.Errorf("error converting reference value %v to KRM format for field %v: %w", val, field, err)
	}
	return refField, convertedVal, nil
}

func convertReferenceValueToKRM(val interface{}, schema *openapi.Schema) (interface{}, error) {
	refConfigs, err := getDCLReferenceExtension(schema)
	if err != nil {
		return nil, fmt.Errorf("error getting DCL reference extension for reference value %v:  %w", val, err)
	}
	if len(refConfigs) >= 1 {
		res := make(map[string]interface{})
		res["external"] = val
		return res, nil
	}
	return nil, fmt.Errorf("getting empty resource types list for reference value")
}

func getDCLReferenceExtension(schema *openapi.Schema) ([]interface{}, error) {
	raw, ok := schema.Extension["x-dcl-references"]
	if !ok {
		return nil, fmt.Errorf("'x-dcl-references' extension is not defined")
	}
	refConfigs, ok := raw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong type for 'x-dcl-references' extension: %T, expect to have []interface{}", raw)
	}
	return refConfigs, nil
}

func resolveReferenceValue(obj map[string]interface{}, schema *openapi.Schema) (interface{}, error) {
	if obj == nil {
		return nil, nil
	}
	refConfigs, err := getDCLReferenceExtension(schema)
	if err != nil {
		return nil, fmt.Errorf("error getting DCL reference extension: %w", err)
	}
	if len(refConfigs) >= 1 {
		return obj["external"], nil
	}
	return nil, fmt.Errorf("couldn't resolve the reference value from %v", obj)
}

func convertToDCLContainerField(obj *unstructured.Unstructured, r *dclunstruct.Resource, schema *openapi.Schema) error {
	container, found, err := getContainerFieldName(schema)
	if err != nil {
		return fmt.Errorf("error getting the container field name %w", err)
	}
	if !found {
		return nil
	}
	annotations := obj.GetAnnotations()
	key := fmt.Sprintf("%s/%s-id", k8s.CNRMGroup, container)
	containerID := annotations[key]
	if containerID == "" {
		return fmt.Errorf("couldn't resolve the value for container field %s", container)
	}
	r.Object[container] = containerID
	return nil
}

func liftDCLContainerField(obj *unstructured.Unstructured, schema *openapi.Schema) error {
	container, found, err := getContainerFieldName(schema)
	if err != nil {
		return fmt.Errorf("error getting the container field name %w", err)
	}
	if !found {
		return nil
	}
	val, ok, err := unstructured.NestedString(obj.Object, "spec", container)
	if err != nil || !ok {
		return fmt.Errorf("couldn't get the value for container field %s: %w", container, err)
	}
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	key := fmt.Sprintf("%s/%s-id", k8s.CNRMGroup, container)
	annotations[key] = val
	obj.SetAnnotations(annotations)
	unstructured.RemoveNestedField(obj.Object, "spec", container)
	return nil
}

func getContainerFieldName(schema *openapi.Schema) (string, bool, error) {
	raw, ok := schema.Extension["x-dcl-parent-container"]
	if !ok {
		return "", false, nil
	}
	// DCL currently doesn't support resources that could have multiple container kinds.
	container, ok := raw.(string)
	if !ok {
		return "", false, fmt.Errorf("wrong type for 'x-dcl-parent-container' extension: %T, expect to have string type", raw)
	}
	return container, true, nil
}

func liftDCLLabelsField(obj *unstructured.Unstructured, dclSchema *openapi.Schema) error {
	labelsField, _, found, err := extension.GetLabelsFieldSchema(dclSchema)
	if err != nil {
		return fmt.Errorf("error getting DCL labels field : '%w'", err)
	}
	if !found {
		return nil
	}
	//TODO(b/164208968): handle edge cases where 'labels' field is not at the top level
	valMap, found, err := unstructured.NestedMap(obj.Object, "spec", labelsField)
	if err != nil {
		return fmt.Errorf("error getting labels %w", err)
	}
	if !found {
		return nil
	}
	labels := make(map[string]string)
	for k, v := range valMap {
		labels[k] = v.(string)
	}
	obj.SetLabels(labels)
	unstructured.RemoveNestedField(obj.Object, "spec", labelsField)
	return nil
}

// The DCL schema use 'name' field to define the resource ID segment of the resource,
// this function will convert it to the unified 'resourceID' field across KCC resources
func convertToKRMResourceIDField(obj *unstructured.Unstructured, schema *openapi.Schema) error {
	// If there is no 'name' field defined in the DCL schema, skip it.
	// One example is DNSRecordSet resource.
	s, found := extension.GetNameFieldSchema(schema)
	if !found {
		return nil
	}
	// If the 'name' field is read-only, skip it.
	// This could happen to resources that don't have the 'name' field as a part of their URLs;
	// however, the REST API returns a output-only `name` field in the response.
	if s.ReadOnly {
		return nil
	}
	val, found, err := unstructured.NestedString(obj.Object, "spec", "name")
	if err != nil {
		return fmt.Errorf("error getting the value of 'name' field: %w", err)
	}
	if !found {
		return fmt.Errorf("'name' field is not found")
	}
	if err := unstructured.SetNestedField(obj.Object, val, "spec", "resourceID"); err != nil {
		return fmt.Errorf("error setting resourceID field: %w", err)
	}
	unstructured.RemoveNestedField(obj.Object, "spec", "name")
	return nil
}

// convertToDCLNameField converts 'resourceID' field in KCC to 'name' field in DCL
func convertToDCLNameField(obj *unstructured.Unstructured, r *dclunstruct.Resource, schema *openapi.Schema) error {
	// If there is no 'name' field defined in the DCL schema, skip it.
	// One example is DNSRecordSet resource.
	s, found := extension.GetNameFieldSchema(schema)
	if !found {
		return nil
	}
	// If the 'name' field is read-only, skip it.
	// This could happen to resources that don't have the 'name' field as a part of their URLs;
	// however, the REST API returns a output-only `name` field in the response.
	if s.ReadOnly {
		return nil
	}

	isServerGeneratedID, err := extension.IsResourceIDFieldServerGenerated(s)
	if err != nil {
		return fmt.Errorf("error parsing 'name' field schema: %w", err)
	}
	// convert 'resourceID' field to DCL's 'name' field
	val, found, err := unstructured.NestedString(obj.UnstructuredContent(), "spec", "resourceID")
	if err != nil {
		return fmt.Errorf("error getting the value of %s: %w", k8s.ResourceIDFieldPath, err)
	}
	if !found {
		// If the resource has a server-generated id, unspecified 'resourceID' field means creating a brand new resource.
		// Leave 'name' field in DCL unspecified also.
		if isServerGeneratedID {
			return nil
		}

		// If the resource allows user-specified name, use metadata.name as default if 'resourceID' field is not specified in spec
		val = obj.GetName()
	}
	if val == "" {
		return fmt.Errorf("the resolved value for 'name' field is invalid: '' (empty string)")
	}
	r.Object["name"] = val
	return nil
}

func convertToDCLLabelsField(obj *unstructured.Unstructured, r *dclunstruct.Resource, schema *openapi.Schema) error {
	labelsField, _, _, err := extension.GetLabelsFieldSchema(schema)
	if err != nil {
		return fmt.Errorf("error getting DCL labels field: '%w'", err)
	}
	//TODO(b/164208968): handle edge cases where 'labels' field is not at the top level
	if _, ok := schema.Properties[labelsField]; !ok {
		return nil
	}
	labels := label.ToJSONCompatibleFormat(label.NewGCPLabelsFromK8sLabels(obj.GetLabels()))
	if len(labels) != 0 {
		r.Object[labelsField] = labels
	}
	return nil
}

func renameStatusFieldIfCollidesWithReservedName(path []string) []string {
	reservedNames := k8s.ReservedStatusFieldNames()
	if _, ok := reservedNames[path[0]]; ok {
		path[0] = k8s.RenameStatusFieldWithReservedName(path[0])
	}
	return path
}
