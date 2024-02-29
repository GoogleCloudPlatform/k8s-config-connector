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

package crdgeneration

import (
	"errors"
	"fmt"
	"log"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration/crdboilerplate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	dclmetatda "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/pathslice"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	"github.com/nasa9084/go-openapi"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	Dcl2CRDLabel = "cnrm.cloud.google.com/dcl2crd"
)

var (
	ErrUnsupportedReferencedResource = fmt.Errorf("referenced resource is unsupported by KCC")
)

type DCL2CRDGenerator struct {
	metadataLoader   dclmetatda.ServiceMetadataLoader
	schemaLoader     dclschemaloader.DCLSchemaLoader
	allSupportedGVKs []schema.GroupVersionKind
}

func New(metadataLoader dclmetatda.ServiceMetadataLoader, schemaLoader dclschemaloader.DCLSchemaLoader, allSupportedGVKs []schema.GroupVersionKind) *DCL2CRDGenerator {
	return &DCL2CRDGenerator{
		metadataLoader:   metadataLoader,
		schemaLoader:     schemaLoader,
		allSupportedGVKs: allSupportedGVKs,
	}
}

// GenerateCRDFromOpenAPISchema returns a CustomResourceDefinition given the DCL OpenAPI schema
func (a *DCL2CRDGenerator) GenerateCRDFromOpenAPISchema(schema *openapi.Schema, gvk schema.GroupVersionKind) (*apiextensions.CustomResourceDefinition, error) {
	r, found := a.metadataLoader.GetResourceWithGVK(gvk)
	if !found {
		return nil, fmt.Errorf("ServiceMetadata for resource with GVK %v is not found", gvk)
	}
	openAPIV3Schema, err := a.generateOpenAPIV3Schema(schema, r)
	if err != nil {
		return nil, fmt.Errorf("error generating CRD schema for %v: %w", gvk.Kind, err)
	}
	crd := GetCustomResourceDefinition(gvk.Kind, gvk.Group, []string{gvk.Version}, "", openAPIV3Schema, Dcl2CRDLabel)
	if r.DCLVersion == "alpha" {
		crd.ObjectMeta.Labels[k8s.KCCStabilityLabel] = k8s.StabilityLevelAlpha
	} else {
		crd.ObjectMeta.Labels[k8s.KCCStabilityLabel] = k8s.StabilityLevelStable
	}
	return crd, nil
}

func (a *DCL2CRDGenerator) generateOpenAPIV3Schema(schema *openapi.Schema, resource metadata.Resource) (*apiextensions.JSONSchemaProps, error) {
	var err error
	crdSchema := crdboilerplate.GetOpenAPIV3SchemaSkeleton()
	specJSONSchema, err := a.generateSpecJSONSchema(schema, resource)
	if err != nil {
		return nil, fmt.Errorf("error generating spec schema %w", err)
	}
	statusJSONSchema, err := generateStatusJSONSchema(schema)
	if err != nil {
		return nil, fmt.Errorf("error generating status schema %w", err)
	}
	if len(specJSONSchema.Properties) > 0 {
		crdSchema.Properties["spec"] = *specJSONSchema
		if len(specJSONSchema.Required) > 0 {
			crdSchema.Required = slice.IncludeString(crdSchema.Required, "spec")
		}
	}
	if statusJSONSchema != nil {
		statusJSONSchema, err = k8s.RenameStatusFieldsWithReservedNames(statusJSONSchema)
		if err != nil {
			return nil, fmt.Errorf("error renaming status fields with reserved names: %w", err)
		}
		for k, v := range statusJSONSchema.Properties {
			crdSchema.Properties["status"].Properties[k] = v
		}
	}
	return crdSchema, nil
}

func (a *DCL2CRDGenerator) generateSpecJSONSchema(schema *openapi.Schema, resource metadata.Resource) (*apiextensions.JSONSchemaProps, error) {
	var err error
	if schema.Type != "object" {
		return nil, fmt.Errorf("expect the entry level DCL OpenAPI schema to be object type, but got %v", schema.Type)
	}
	jsonSchema := &apiextensions.JSONSchemaProps{
		Type:       "object",
		Properties: make(map[string]apiextensions.JSONSchemaProps),
	}
	required := make([]string, 0)
	dclLabelsField, _, dclLabelsFieldFound, err := extension.GetLabelsFieldSchema(schema)
	if err != nil {
		return nil, fmt.Errorf("error extracting DCL labels field schema: %w", err)
	}
	for k, v := range schema.Properties {
		if !v.ReadOnly {
			if k == "name" {
				s, err := handleNameField(v)
				if err != nil {
					return nil, fmt.Errorf("error handling 'name' field: %w", err)
				}
				jsonSchema.Properties["resourceID"] = *s
				continue
			}
			// TODO(b/164208968): handle edge cases where dclLabelsField field is not at the top level
			if dclLabelsFieldFound && k == dclLabelsField {
				continue
			}
			// TODO(b/186159460): Delete this if-block once all resources
			// support hierarchical references.
			if dcl.IsContainerField([]string{k}) && !resource.SupportsHierarchicalReferences {
				continue
			}
			// Multi-type parent reference fields (i.e. "parent" fields) need
			// to be split up into separate resource reference fields
			if dcl.IsMultiTypeParentReferenceField([]string{k}) {
				// TODO(b/186159460): Delete this if-block once all resources
				// support hierarchical references.
				if !resource.SupportsHierarchicalReferences {
					return nil, fmt.Errorf("resource supports 'parent' field but doesn't support hierarchical references")
				}
				refs, err := a.multiTypeParentFieldToHierarchicalRefs(v)
				if err != nil {
					return nil, err
				}
				for fieldName, s := range refs {
					s, err := prependImmutableToDescriptionIfImmutable(s, v)
					if err != nil {
						return nil, fmt.Errorf("error prepending Immutable to description of hierarchical reference field %v if field is immutable: %w", fieldName, err)
					}
					jsonSchema.Properties[fieldName] = *s
				}
				continue
			}
			fieldName, s, err := a.dclSchemaToSpecJSONSchema([]string{k}, v, false)
			if err != nil {
				return nil, err
			}
			if isRequiredField(schema, k) {
				required = slice.IncludeString(required, fieldName)
			}
			jsonSchema.Properties[fieldName] = *s
		}
	}
	if len(required) != 0 {
		jsonSchema.Required = required
	}
	jsonSchema, err = a.addSchemaRulesForHierarchicalReferences(jsonSchema, schema, resource)
	if err != nil {
		return nil, err
	}
	return jsonSchema, nil
}

func generateStatusJSONSchema(schema *openapi.Schema) (*apiextensions.JSONSchemaProps, error) {
	if schema.Type != "object" {
		return nil, fmt.Errorf("expect the entry level DCL OpenAPI schema to be object type, but got %v", schema.Type)
	}
	return getStatusSchema(schema), nil
}

func getStatusSchema(schema *openapi.Schema) *apiextensions.JSONSchemaProps {
	// we treat read-only fields as status fields
	if schema.ReadOnly {
		return dclSchemaToStatusJSONSchema(schema)
	}
	if schema.Type == "object" {
		jsonSchema := &apiextensions.JSONSchemaProps{
			Type:       "object",
			Properties: make(map[string]apiextensions.JSONSchemaProps),
		}
		for k, v := range schema.Properties {
			s := getStatusSchema(v)
			if s != nil {
				jsonSchema.Properties[k] = *s
			}
		}
		if len(jsonSchema.Properties) == 0 {
			return nil
		}
		return jsonSchema
	}
	// for now, don't split nested read-only fields in an array of objects into status
	// those fields only make sense along with the whole object.
	return nil
}

func dclSchemaToStatusJSONSchema(schema *openapi.Schema) *apiextensions.JSONSchemaProps {
	jsonSchema := apiextensions.JSONSchemaProps{}
	jsonSchema.Type = schema.Type
	jsonSchema.Description = schema.Description
	jsonSchema.Format = schema.Format

	switch schema.Type {
	case "object":
		// The field additionalProperties is mutually exclusive with properties in CustomResourceDefinition
		if schema.AdditionalProperties != nil {
			jsonSchema.AdditionalProperties = &apiextensions.JSONSchemaPropsOrBool{
				Schema: dclSchemaToStatusJSONSchema(schema.AdditionalProperties),
			}
			break
		}
		jsonSchema.Properties = make(map[string]apiextensions.JSONSchemaProps)
		for k, v := range schema.Properties {
			s := dclSchemaToStatusJSONSchema(v)
			jsonSchema.Properties[k] = *s
		}
	case "array":
		itemSchema := dclSchemaToStatusJSONSchema(schema.Items)
		jsonSchema.Items = &apiextensions.JSONSchemaPropsOrArray{
			Schema: itemSchema,
		}
	case "boolean", "number", "string", "integer":
		jsonSchema.Type = schema.Type
	default:
		log.Fatalf("unknown schema type %v", schema.Type)
	}
	return &jsonSchema
}

func (a *DCL2CRDGenerator) dclSchemaToSpecJSONSchema(path []string, schema *openapi.Schema, isCollectionItemSchema bool) (string, *apiextensions.JSONSchemaProps, error) {
	field := pathslice.Base(path)
	if !schema.ReadOnly && extension.IsReferenceField(schema) {
		refFieldName, err := extension.GetReferenceFieldName(path, schema)
		if err != nil {
			return "", nil, fmt.Errorf("error resolving the name for reference field %s: %w", field, err)
		}
		refSchema, err := a.handleReferenceField(path, schema)
		if err != nil {
			return "", nil, fmt.Errorf("error resolving the reference schema for field %v: %w", field, err)
		}
		refSchema, err = prependImmutableToDescriptionIfImmutable(refSchema, schema)
		if err != nil {
			return "", nil, fmt.Errorf("error prepending Immutable to description of field %v if field is immutable: %w", field, err)
		}
		return refFieldName, refSchema, nil
	}
	isSensitive, err := extension.IsSensitiveField(schema)
	if err != nil {
		return "", nil, fmt.Errorf("error checking sensitivity for field %v: %w", field, err)
	}
	if !schema.ReadOnly && isSensitive {
		s := crdboilerplate.GetSensitiveFieldSchemaBoilerplate()
		s.Description = schema.Description
		jsonSchema, err := prependImmutableToDescriptionIfImmutable(&s, schema)
		if err != nil {
			return "", nil, fmt.Errorf("error prepending Immutable to description of field %v if field is immutable: %w", field, err)
		}
		return field, jsonSchema, nil
	}
	jsonSchema := &apiextensions.JSONSchemaProps{}
	jsonSchema.Type = schema.Type
	jsonSchema.Description = schema.Description
	jsonSchema.Format = schema.Format

	jsonSchema, err = prependImmutableToDescriptionIfImmutable(jsonSchema, schema)
	if err != nil {
		return "", nil, fmt.Errorf("error prepending Immutable to description of field %v if field is immutable: %w", field, err)
	}

	fieldName := field
	switch schema.Type {
	case "object":
		// The field additionalProperties is mutually exclusive with properties in CustomResourceDefinition
		if schema.AdditionalProperties != nil {
			_, s, err := a.dclSchemaToSpecJSONSchema(path, schema.AdditionalProperties, true)
			if err != nil {
				return "", nil, err
			}
			jsonSchema.AdditionalProperties = &apiextensions.JSONSchemaPropsOrBool{
				Schema: s,
			}
			break
		}
		jsonSchema.Properties = make(map[string]apiextensions.JSONSchemaProps)
		required := make([]string, 0)
		for k, v := range schema.Properties {
			if !v.ReadOnly || isCollectionItemSchema {
				fieldName, s, err := a.dclSchemaToSpecJSONSchema(append(path, k), v, isCollectionItemSchema)
				if err != nil {
					return "", nil, err
				}
				jsonSchema.Properties[fieldName] = *s
				if isRequiredField(schema, k) {
					required = slice.IncludeString(required, fieldName)
				}
			}
		}
		if len(required) != 0 {
			jsonSchema.Required = required
		}
	case "array":
		f, itemSchema, err := a.dclSchemaToSpecJSONSchema(path, schema.Items, true)
		if err != nil {
			return "", nil, err
		}
		fieldName = f
		jsonSchema.Items = &apiextensions.JSONSchemaPropsOrArray{
			Schema: itemSchema,
		}
	case "boolean", "number", "string", "integer":
		jsonSchema.Type = schema.Type
	default:
		log.Fatalf("unknown schema type %v for field %v", schema.Type, field)
	}
	return fieldName, jsonSchema, nil
}

func (a *DCL2CRDGenerator) multiTypeParentFieldToHierarchicalRefs(schema *openapi.Schema) (map[string]*apiextensions.JSONSchemaProps, error) {
	tcs, err := dcl.GetReferenceTypeConfigs(schema, a.metadataLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting reference type configs for DCL field 'parent': %w", err)
	}

	keys := make([]string, 0)
	for _, tc := range tcs {
		keys = append(keys, tc.Key)
	}

	hierarchicalRefs := make(map[string]*apiextensions.JSONSchemaProps)
	for _, tc := range tcs {
		s, err := a.resolveResourceReferenceJSONSchemaPerType(&tc, "")
		if err != nil {
			return nil, err
		}
		s.Description = fmt.Sprintf("The %v that this resource belongs to. Only one of [%v] may be specified.", tc.GVK.Kind, strings.Join(keys, ", "))
		hierarchicalRefs[tc.Key] = s
	}
	return hierarchicalRefs, nil
}

func (a *DCL2CRDGenerator) handleReferenceField(path []string, schema *openapi.Schema) (*apiextensions.JSONSchemaProps, error) {
	if dcl.IsSingleTypeParentReferenceField(path) {
		return a.handleSingleTypeParentReferenceField(path, schema)
	}
	if schema.Type == "array" {
		return a.handleListOfReferencesField(schema)
	}
	return a.resolveResourceReferenceJSONSchema(schema)
}

func (a *DCL2CRDGenerator) handleSingleTypeParentReferenceField(path []string, schema *openapi.Schema) (*apiextensions.JSONSchemaProps, error) {
	field := pathslice.Base(path)
	refSchema, err := a.resolveResourceReferenceJSONSchema(schema)
	if err != nil {
		return nil, err
	}
	refSchema.Description = fmt.Sprintf("The %v that this resource belongs to.", strings.Title(field))
	return refSchema, nil
}

func (a *DCL2CRDGenerator) handleListOfReferencesField(schema *openapi.Schema) (*apiextensions.JSONSchemaProps, error) {
	refSchema, err := a.resolveResourceReferenceJSONSchema(schema.Items)
	if err != nil {
		return nil, err
	}
	res := &apiextensions.JSONSchemaProps{
		Type: "array",
		Items: &apiextensions.JSONSchemaPropsOrArray{
			Schema: refSchema,
		},
	}
	return res, nil
}

func handleNameField(schema *openapi.Schema) (*apiextensions.JSONSchemaProps, error) {
	isServerGenerated, err := extension.IsResourceIDFieldServerGenerated(schema)
	if err != nil {
		return nil, err
	}

	var description string
	if isServerGenerated {
		description = GenerateResourceIDFieldDescription("name", true)
	} else {
		description = GenerateResourceIDFieldDescription("name", false)
	}

	return &apiextensions.JSONSchemaProps{
		Type:        schema.Type,
		Description: description,
	}, nil
}

func (a *DCL2CRDGenerator) resolveResourceReferenceJSONSchema(schema *openapi.Schema) (*apiextensions.JSONSchemaProps, error) {
	tcs, err := dcl.GetReferenceTypeConfigs(schema, a.metadataLoader)
	if err != nil {
		return nil, err
	}

	// We name reference field as {fieldName}Ref.
	if len(tcs) == 1 {
		refSchema, err := a.resolveResourceReferenceJSONSchemaPerType(&tcs[0], schema.Description)
		return refSchema, err
	}

	refSchema, err := a.resolveResourceReferenceJSONSchemaMultiTypes(tcs, schema.Description)
	return refSchema, err
}

func (a *DCL2CRDGenerator) resolveResourceReferenceJSONSchemaPerType(tc *corekccv1alpha1.TypeConfig, description string) (*apiextensions.JSONSchemaProps, error) {
	supported, err := a.validateReferencedResourceKind(tc)
	if err != nil {
		return nil, err
	}
	externalRefDescription, err := a.getDescriptionForExternalRef(tc, description)
	if err != nil {
		return nil, err
	}
	refSchema := crdboilerplate.GetResourceReferenceSchemaBoilerplate(externalRefDescription)
	if !supported {
		MarkReferencedKindsNotSupported(refSchema, []string{tc.GVK.Kind})
	}
	return refSchema, nil
}

func (a *DCL2CRDGenerator) resolveResourceReferenceJSONSchemaMultiTypes(tcs []corekccv1alpha1.TypeConfig, description string) (*apiextensions.JSONSchemaProps, error) {
	supportedKinds := make([]string, 0)
	unsupportedKinds := make([]string, 0)
	for _, tc := range tcs {
		supported, err := a.validateReferencedResourceKind(&tc)
		if err != nil {
			return nil, err
		}
		if !supported {
			unsupportedKinds = append(unsupportedKinds, tc.GVK.Kind)
		} else {
			supportedKinds = append(supportedKinds, tc.GVK.Kind)
		}
	}
	externalRefDescription, err := a.getDescriptionForMultiKindExternalRef(tcs, description)
	if err != nil {
		return nil, err
	}
	refSchema := crdboilerplate.GetMultiKindResourceReferenceSchemaBoilerplate(externalRefDescription, supportedKinds)
	if len(unsupportedKinds) > 0 {
		MarkReferencedKindsNotSupported(refSchema, unsupportedKinds)
	}
	return refSchema, nil
}

func (a *DCL2CRDGenerator) validateReferencedResourceKind(tc *corekccv1alpha1.TypeConfig) (supported bool, err error) {
	if !k8s.GVKListContains(a.allSupportedGVKs, tc.GVK) {
		return false, nil
	}
	// On runtime, we need to load the DCL schema for the referenced resource type to resolve the standardized resource name
	if tc.TargetField == "name" {
		_, err := dclschemaloader.GetDCLSchemaForGVK(tc.GVK, a.metadataLoader, a.schemaLoader)
		if err != nil {
			return false, fmt.Errorf("error getting the DCL schema for %v: %w; if it's a supported tf-based resource type, "+
				"ensure that it has been declared in pkg/dcl/metadata/metadata.go with releasable flag as false, "+
				"and that its service is imported in pkg/dcl/schema/dclschemaloader/dclschemaloader.go, "+
				"since we need to load its OpenAPI schema for 'x-dcl-id' template", tc.GVK, err)
		}
	}
	return true, nil
}

func isRequiredField(schema *openapi.Schema, field string) bool {
	for _, item := range schema.Required {
		if field == item {
			return true
		}
	}
	return false
}

func (a *DCL2CRDGenerator) addSchemaRulesForHierarchicalReferences(jsonSchema *apiextensions.JSONSchemaProps, schema *openapi.Schema, resource metadata.Resource) (*apiextensions.JSONSchemaProps, error) {
	// TODO(b/186159460): Delete this if-block once all resources support
	// hierarchical references.
	if !resource.SupportsHierarchicalReferences {
		return jsonSchema, nil
	}
	hierarchicalRefs, err := dcl.GetHierarchicalReferenceConfigFromDCLSchema(schema, a.metadataLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting hierarchical reference config for resource: %w", err)
	}
	if resource.SupportsContainerAnnotations {
		// If resource supports resource-level container annotations, mark
		// hierarchical references optional since users can use the annotations
		// to configure the references.
		return MarkHierarchicalReferencesOptionalButMutuallyExclusive(jsonSchema, hierarchicalRefs), nil
	}
	return MarkHierarchicalReferencesRequiredButMutuallyExclusive(jsonSchema, hierarchicalRefs), nil
}

func (a *DCL2CRDGenerator) getDescriptionForExternalRef(tc *corekccv1alpha1.TypeConfig, baseDescription string) (string, error) {
	exampleAllowedValue, err := a.getExampleAllowedValueForExternalRef(tc)
	if err != nil {
		if errors.Is(err, ErrUnsupportedReferencedResource) {
			return baseDescription, nil
		}
		return "", err
	}
	return text.AppendStrAsNewParagraph(
		baseDescription,
		fmt.Sprintf("Allowed value: %v", exampleAllowedValue),
	), nil
}

func (a *DCL2CRDGenerator) getDescriptionForMultiKindExternalRef(tcs []corekccv1alpha1.TypeConfig, baseDescription string) (string, error) {
	exampleAllowedValues := make([]string, 0)
	for _, tc := range tcs {
		v, err := a.getExampleAllowedValueForExternalRef(&tc)
		if err != nil {
			if errors.Is(err, ErrUnsupportedReferencedResource) {
				continue
			}
			return "", err
		}
		exampleAllowedValues = append(exampleAllowedValues, v)
	}
	if len(exampleAllowedValues) == 0 {
		return baseDescription, nil
	}
	return text.AppendStrAsNewParagraph(
		baseDescription,
		fmt.Sprintf("Allowed values:\n* %v", strings.Join(exampleAllowedValues, "\n* ")),
	), nil
}

func (a *DCL2CRDGenerator) getExampleAllowedValueForExternalRef(tc *corekccv1alpha1.TypeConfig) (string, error) {
	// Cannot programmatically determine values allowed by `external` if the
	// referenced resource is not yet supported by KCC.
	if !k8s.GVKListContains(a.allSupportedGVKs, tc.GVK) {
		switch tc.GVK.Kind {
		default:
			return "", ErrUnsupportedReferencedResource
		// For some resources, `external` allows the same value regardless of
		// the referencing resource, allowing us to just hardcode an allowed
		// value.
		case "Organization":
			return "The Google Cloud resource name of a Google Cloud Organization (format: `organizations/{{name}}`).", nil
		case "BillingAccount":
			return "The Google Cloud resource name of a Google Cloud Billing Account (format: `billingAccounts/{{name}}`).", nil
		}
	}

	article := text.IndefiniteArticleFor(tc.GVK.Kind)
	switch tc.TargetField {
	case "":
		return "", fmt.Errorf("reference field unexpectedly does not have a target field specified")
	case "name":
		s, err := dclschemaloader.GetDCLSchemaForGVK(tc.GVK, a.metadataLoader, a.schemaLoader)
		if err != nil {
			return "", fmt.Errorf("error getting DCL schema for GVK %v: %w", tc.GVK, err)
		}
		template, err := dclextension.GetNameValueTemplate(s)
		if err != nil {
			return "", fmt.Errorf("error getting name value template for GVK %v: %w", tc.GVK, err)
		}
		return fmt.Sprintf("The Google Cloud resource name of %v `%v` resource (format: `%v`).", article, tc.GVK.Kind, template), nil
	default:
		return fmt.Sprintf("The `%v` field of %v `%v` resource.", tc.TargetField, article, tc.GVK.Kind), nil
	}
}

func prependImmutableToDescriptionIfImmutable(jsonSchema *apiextensions.JSONSchemaProps, schema *openapi.Schema) (*apiextensions.JSONSchemaProps, error) {
	jsonSchemaCopy := jsonSchema.DeepCopy()
	ok, err := dclextension.IsImmutableField(schema)
	if err != nil {
		return nil, fmt.Errorf("error determining if field is immutable: %w", err)
	}
	if ok {
		jsonSchemaCopy.Description = strings.TrimSpace("Immutable. " + jsonSchemaCopy.Description)
	}
	return jsonSchemaCopy, nil
}
