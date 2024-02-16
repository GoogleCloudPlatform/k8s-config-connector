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

package dcl

import (
	"fmt"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/pathslice"

	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// The following variables represent the fields that are used to reference
	// parent resources. DCL resources can support either one multi-type parent
	// reference field or one of the single-type parent reference fields.
	multiTypeParentReferenceField              = "parent"
	singleTypeParentReferenceFieldProject      = "project"
	singleTypeParentReferenceFieldFolder       = "folder"
	singleTypeParentReferenceFieldOrganization = "organization"
	singleTypeParentReferenceFields            = map[string]bool{
		singleTypeParentReferenceFieldProject:      true,
		singleTypeParentReferenceFieldFolder:       true,
		singleTypeParentReferenceFieldOrganization: true,
	}

	// The following variables represent the different hierarchical reference
	// configurations that can be had by DCL-based KCC resources.
	hierarchicalReferenceProject = corekccv1alpha1.HierarchicalReference{
		Key:  "projectRef",
		Type: corekccv1alpha1.HierarchicalReferenceTypeProject,
	}
	hierarchicalReferenceFolder = corekccv1alpha1.HierarchicalReference{
		Key:  "folderRef",
		Type: corekccv1alpha1.HierarchicalReferenceTypeFolder,
	}
	hierarchicalReferenceOrganization = corekccv1alpha1.HierarchicalReference{
		Key:  "organizationRef",
		Type: corekccv1alpha1.HierarchicalReferenceTypeOrganization,
	}
	hierarchicalReferenceBillingAccount = corekccv1alpha1.HierarchicalReference{
		Key:  "billingAccountRef",
		Type: corekccv1alpha1.HierarchicalReferenceTypeBillingAccount,
	}
)

/*
*
dclReferenceExtensionElem represents an element in a 'x-dcl-references' list.
```

	x-dcl-references:
	- field: name
	  parent: true
	  resource: SomeService/SomeResourceType

```
*
*/
type dclReferenceExtensionElem struct {
	// Resource indicates the referenced resource type in the format: Service/ResourceKind, e.g. Compute/BackendBucket.
	Resource string `json:"resource"`
	// Field is the referenced resource's field from which to extract the desired value, e.g. "name", "selfLink"
	Field string `json:"field"`
	// Parent specifies whether the referenced resource is a parent. If the parent
	// is successfully deleted, it's assumed that this resource may be deleted without any call to the
	// underlying API.
	Parent bool `json:"parent,omitempty"`
}

func GetReferenceTypeConfigs(schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader) ([]corekccv1alpha1.TypeConfig, error) {
	refExtensionElems, err := getDCLReferenceExtensionElems(schema)
	if err != nil {
		return nil, err
	}
	res := make([]corekccv1alpha1.TypeConfig, 0)
	for _, e := range refExtensionElems {
		tc, err := toTypeConfig(e, smLoader)
		if err != nil {
			return nil, err
		}
		res = append(res, *tc)
	}
	return res, nil
}

// ToTypeConfig converts a 'x-dcl-references' element to a TypeConfig.
func ToTypeConfig(rawElem map[interface{}]interface{}, smLoader dclmetadata.ServiceMetadataLoader) (*corekccv1alpha1.TypeConfig, error) {
	e, err := toDCLReferenceExtensionElem(rawElem)
	if err != nil {
		return nil, err
	}
	return toTypeConfig(e, smLoader)
}

func toTypeConfig(e dclReferenceExtensionElem, smLoader dclmetadata.ServiceMetadataLoader) (*corekccv1alpha1.TypeConfig, error) {
	if err := validateDCLReferenceExtensionElem(e); err != nil {
		return nil, err
	}
	tc := &corekccv1alpha1.TypeConfig{}
	tc.TargetField = e.Field
	tc.Parent = e.Parent
	refGVK, err := getReferenceGVK(e.Resource, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error resolving the GVK for referenced resource: %w", err)
	}
	tc.GVK = refGVK
	tc.Key = text.LowercaseInitial(k8s.KindWithoutServicePrefix(tc.GVK)) + "Ref"
	return tc, nil
}

func validateDCLReferenceExtensionElem(e dclReferenceExtensionElem) error {
	if e.Resource == "" {
		return fmt.Errorf("required 'resource' attribute is not specified in 'x-dcl-references' extension")
	}
	if e.Field == "" {
		return fmt.Errorf("required 'field' attribute is not specified in 'x-dcl-references' extension")
	}
	return nil
}

func convertToStringInterfaceMap(in map[interface{}]interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	for k, v := range in {
		s, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("wrong type for the key: %T, expect to have string", k)
		}
		out[s] = v
	}
	return out, nil
}

func getReferenceGVK(resource string, smLoader dclmetadata.ServiceMetadataLoader) (schema.GroupVersionKind, error) {
	// The 'resource' attribute of a 'x-dcl-references' element has the format
	// "Service/DCLType" (e.g. Compute/BackendBucket)
	components := strings.Split(resource, "/")
	if len(components) != 2 {
		return schema.GroupVersionKind{}, fmt.Errorf("invalid format for 'resource' attribute in 'x-dcl-references' extension: %v", resource)
	}
	service := components[0]
	dclType := components[1]

	sm, found := smLoader.GetServiceMetadata(service)
	if !found {
		return schema.GroupVersionKind{}, fmt.Errorf("ServiceMetadata for service %v is not found", service)
	}
	r, found := sm.GetResourceWithType(dclType)
	if !found {
		return schema.GroupVersionKind{}, fmt.Errorf("resource with DCL type %v not supported in service %v", dclType, service)
	}
	return dclmetadata.GVKForResource(sm, r), nil
}

func getDCLReferenceExtensionElems(schema *openapi.Schema) ([]dclReferenceExtensionElem, error) {
	extension, ok := schema.Extension["x-dcl-references"]
	if !ok {
		return nil, fmt.Errorf("no 'x-dcl-references' extension found")
	}
	extensionAsList, ok := extension.([]interface{})
	if !ok {
		return nil, fmt.Errorf("wrong type for 'x-dcl-references' extension: %T, expect to have []interface{}", extension)
	}
	res := make([]dclReferenceExtensionElem, 0)
	for _, elem := range extensionAsList {
		elemAsMap, ok := elem.(map[interface{}]interface{})
		if !ok {
			return nil, fmt.Errorf("wrong type for element in 'x-dcl-references' extension %T, expect to have map[interface{}]interface{}", elem)
		}
		e, err := toDCLReferenceExtensionElem(elemAsMap)
		if err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, nil
}

func toDCLReferenceExtensionElem(rawElem map[interface{}]interface{}) (dclReferenceExtensionElem, error) {
	m, err := convertToStringInterfaceMap(rawElem)
	if err != nil {
		return dclReferenceExtensionElem{}, fmt.Errorf("error converting 'x-dcl-references' element to map[string]interface{}: %w", err)
	}
	e := dclReferenceExtensionElem{}
	if err := util.Marshal(m, &e); err != nil {
		return dclReferenceExtensionElem{}, fmt.Errorf("error marshalling 'x-dcl-references' element to struct: %w", err)
	}
	return e, nil
}

func GetHierarchicalReferencesForGVK(gvk schema.GroupVersionKind, smLoader dclmetadata.ServiceMetadataLoader, schemaLoader dclschemaloader.DCLSchemaLoader) ([]corekccv1alpha1.HierarchicalReference, error) {
	r, found := smLoader.GetResourceWithGVK(gvk)
	if !found {
		return nil, fmt.Errorf("ServiceMetadata for resource with GroupVersionKind %v not found", gvk)
	}
	// TODO(b/186159460): Delete this if-block once all resources support
	// hierarchical references.
	if !r.SupportsHierarchicalReferences {
		return nil, nil
	}
	stv, err := dclmetadata.ToServiceTypeVersion(gvk, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting DCL ServiceTypeVersion for GroupVersionKind %v: %w", gvk, err)
	}
	dclSchema, err := schemaLoader.GetDCLSchema(stv)
	if err != nil {
		return nil, fmt.Errorf("error getting the DCL Schema for GroupVersionKind %v: %w", gvk, err)
	}
	hierarchicalRefs, err := GetHierarchicalReferenceConfigFromDCLSchema(dclSchema, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error resolving the hierarchical reference config from DCL schema for GroupVersionKind %v: %w", gvk, err)
	}
	return hierarchicalRefs, nil
}

func GetHierarchicalReferenceConfigFromDCLSchema(schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader) ([]corekccv1alpha1.HierarchicalReference, error) {
	// Resource supports multiple parent types
	if SupportsMultipleParentTypes(schema) {
		res, err := GetHierarchicalReferenceConfigForMultiParentResource(schema, smLoader)
		if err != nil {
			return nil, fmt.Errorf("error getting hierarchical reference config for resource that supports multiple parent types: %w", err)
		}
		return res, nil
	}

	// Resource supports one parent type or none at all
	field, err := getSingleTypeParentReferenceField(schema)
	if err != nil {
		return nil, fmt.Errorf("error getting single-type parent reference field for resource: %w", err)
	}
	if field == "" {
		// Resource doesn't support any parent reference fields
		return nil, nil
	}
	switch field {
	case singleTypeParentReferenceFieldProject:
		return []corekccv1alpha1.HierarchicalReference{hierarchicalReferenceProject}, nil
	case singleTypeParentReferenceFieldFolder:
		return []corekccv1alpha1.HierarchicalReference{hierarchicalReferenceFolder}, nil
	case singleTypeParentReferenceFieldOrganization:
		return []corekccv1alpha1.HierarchicalReference{hierarchicalReferenceOrganization}, nil
	default:
		panic(fmt.Errorf("unrecognized single-type parent reference field: %v", field))
	}
}

func GetHierarchicalReferenceConfigForMultiParentResource(schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader) ([]corekccv1alpha1.HierarchicalReference, error) {
	if !SupportsMultipleParentTypes(schema) {
		return nil, fmt.Errorf("resource does not support multiple parent types")
	}

	parentFieldSchema := schema.Properties["parent"]
	tcs, err := GetReferenceTypeConfigs(parentFieldSchema, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error getting reference type configs for DCL field 'parent': %w", err)
	}

	res := make([]corekccv1alpha1.HierarchicalReference, 0)
	for _, tc := range tcs {
		switch tc.GVK.Kind {
		case "Project":
			res = append(res, hierarchicalReferenceProject)
		case "Folder":
			res = append(res, hierarchicalReferenceFolder)
		case "Organization":
			res = append(res, hierarchicalReferenceOrganization)
		case "BillingAccount":
			res = append(res, hierarchicalReferenceBillingAccount)
		default:
			panic(fmt.Errorf("'parent' field references an unsupported resource kind: %v", tc.GVK.Kind))
		}
	}
	return res, nil
}

func getSingleTypeParentReferenceField(schema *openapi.Schema) (string, error) {
	if SupportsMultipleParentTypes(schema) {
		return "", fmt.Errorf("resource supports multiple parent types, not one")
	}
	var res string
	for f := range singleTypeParentReferenceFields {
		_, ok := schema.Properties[f]
		if !ok {
			continue
		}
		if res != "" {
			return "", fmt.Errorf("resource unexpectedly has more than one single-type parent reference field")
		}
		res = f
	}
	return res, nil
}

func ParentReferenceFields() []string {
	return []string{
		singleTypeParentReferenceFieldProject,
		singleTypeParentReferenceFieldFolder,
		singleTypeParentReferenceFieldOrganization,
		multiTypeParentReferenceField,
	}
}

func IsParentReferenceField(path []string) bool {
	return IsSingleTypeParentReferenceField(path) ||
		IsMultiTypeParentReferenceField(path)
}

func IsSingleTypeParentReferenceField(path []string) bool {
	if len(path) > 1 {
		return false
	}
	field := pathslice.Base(path)
	_, ok := singleTypeParentReferenceFields[field]
	return ok
}

func IsMultiTypeParentReferenceField(path []string) bool {
	if len(path) > 1 {
		return false
	}
	field := pathslice.Base(path)
	return field == multiTypeParentReferenceField
}

func SupportsMultipleParentTypes(schema *openapi.Schema) bool {
	_, ok := schema.Properties[multiTypeParentReferenceField]
	return ok
}

// GetHierarchicalRefFromConfigForMultiParentResource gets the value and
// TypeConfig of the hierarchical reference in the KRM config, assuming that
// the config is for a multi-parent resource (i.e. supports more than one
// hierarchical reference). Returns nil if no hierarchical reference is found.
// Returns an error if multiple are found.
func GetHierarchicalRefFromConfigForMultiParentResource(config map[string]interface{}, schema *openapi.Schema,
	smLoader dclmetadata.ServiceMetadataLoader) (interface{}, *corekccv1alpha1.TypeConfig, error) {
	tcs, err := GetReferenceTypeConfigs(schema, smLoader)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting reference type configs: %w", err)
	}
	var rawVal interface{}
	var typeConfig corekccv1alpha1.TypeConfig
	for _, tc := range tcs {
		if v, ok := config[tc.Key]; ok && v != nil {
			if rawVal != nil {
				return nil, nil, fmt.Errorf("multiple hierarchical references found in config: %v and %v", typeConfig.Key, tc.Key)
			}
			rawVal = v
			typeConfig = tc
		}
	}
	if rawVal == nil {
		return nil, nil, nil
	}
	return rawVal, &typeConfig, nil
}

// ParentPrefixForKind gets the parent prefix for the given kind (e.g.
// "Project" => "projects/"). This is used for setting/parsing values of
// multi-type parent reference fields in DCL which require a parent
// prefix to denote the parent type.
func ParentPrefixForKind(kind string) string {
	switch kind {
	case "Project", "Folder", "Organization", "BillingAccount":
		return fmt.Sprintf("%vs/", text.LowercaseInitial(kind))
	default:
		panic(fmt.Errorf("tried to get parent prefix for kind %v which is not recognized as a hierarchical resource", kind))
	}
}
