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
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis"
	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	tfresource "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/resource"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/klog/v2"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

// ResolveSpecAndStatus returns the resolved spec and status in different formats
// gated by the 'state-into-spec' annotation.
//
// If the annotation takes the 'merge' value, the function returns spec as a mix of k8s user managed fields and defaulted state from APIs
// and returns status with the legacy format containing observed state for output-only fields only.
//
// If the annotation takes the 'absent' value, the function will delegate to resolveDesiredStateInSpecAndObservedStateInStatus() to resolve
// the spec and the status.
func ResolveSpecAndStatus(resource *Resource, state *terraform.InstanceState) (
	spec map[string]interface{}, status map[string]interface{}) {
	val, found := k8s.GetAnnotation(k8s.StateIntoSpecAnnotation, resource)
	if !found || val == apis.StateMergeIntoSpec {
		return GetSpecAndStatusFromState(resource, state)
	}
	return resolveDesiredStateInSpecAndObservedStateInStatus(resource, state)
}

// GetSpecAndStatusFromState converts state into separate, KRM-compatible spec and status
// objects.
//
// This function can handle partial state structs (ones that may fail if applied with terraform).
// The resource.Spec that is passed is assumed to be the desired state of the user, and as such
// fields that are specified by Kubernetes to be managed by Config Connector will use the values in
// resource.Spec rather than those in state in the returned spec and status. That said, this function
// returns spec as a mix of k8s user managed fields and defaulted state from APIs
// and returns status with the legacy format containing observed state for output-only fields only.
//
// See ConvertTFObjToKCCObj for a complete description of the merging behavior of
// state and resource.Spec (passed as prevSpec)
func GetSpecAndStatusFromState(resource *Resource, state *terraform.InstanceState) (
	spec map[string]interface{}, status map[string]interface{}) {
	unmodifiedState := InstanceStateToMap(resource.TFResource, state)
	krmState, krmStateWithIgnoredOutputOnlySpecFields := ConvertTFObjToKCCObj(unmodifiedState, resource.Spec, resource.TFResource.Schema,
		&resource.ResourceConfig, "", resource.ManagedFields)
	krmState = withCustomExpanders(krmState, resource, resource.Kind)
	krmStateWithIgnoredOutputOnlySpecFields = withCustomExpanders(krmStateWithIgnoredOutputOnlySpecFields, resource, resource.Kind)
	spec = make(map[string]interface{})
	status = make(map[string]interface{})
	for field, fieldSchema := range resource.TFResource.Schema {
		key := text.SnakeCaseToLowerCamelCase(field)
		if ok, refConfig := IsReferenceField(field, &resource.ResourceConfig); ok && refConfig.Key != "" {
			key = refConfig.Key
		}
		val := krmState[key]
		if val == nil {
			continue
		}
		target := &spec
		if !fieldSchema.Required && !fieldSchema.Optional {
			if k8s.OutputOnlyFieldsAreUnderObservedState(resource.GroupVersionKind()) {
				observedState, ok := status[k8s.ObservedStateFieldName]
				if !ok {
					// Always add the 'observedState' subfield if the resource
					// should have its computed field under the observed state.
					observedState = make(map[string]interface{})
					status[k8s.ObservedStateFieldName] = observedState
				}
				observedStateMap := observedState.(map[string]interface{})
				target = &(observedStateMap)
			} else {
				target = &status
				key = renameStatusFieldIfNeeded(resource.ResourceConfig.Name, key)
			}
		}
		(*target)[key] = val
	}
	if location, ok := getLocationValueFromResourceOrState(resource, unmodifiedState); ok {
		spec["location"] = location
	}
	if conditions, ok := resource.Status["conditions"]; ok {
		status["conditions"] = deepcopy.DeepCopy(conditions)
	}
	if observedGeneration, ok := resource.Status["observedGeneration"]; ok {
		status["observedGeneration"] = deepcopy.DeepCopy(observedGeneration)
	}
	if resource.ResourceConfig.ObservedFields != nil {
		observedFields := resolveObservedFields(resource, krmStateWithIgnoredOutputOnlySpecFields)
		if len(observedFields) > 0 {
			// Merge the observed fields into the observed state.
			observedState, ok := status[k8s.ObservedStateFieldName]
			if !ok {
				observedState = make(map[string]interface{})
			}
			for k, v := range observedFields {
				observedState.(map[string]interface{})[k] = v
			}
			status[k8s.ObservedStateFieldName] = observedState
		}
	}
	// Remove the 'observedState' subfield if it is empty.
	observedState, ok := status[k8s.ObservedStateFieldName]
	if ok && len(observedState.(map[string]interface{})) == 0 {
		delete(status, k8s.ObservedStateFieldName)
	}

	if len(spec) == 0 {
		spec = nil
	}
	if len(status) == 0 {
		status = nil
	}
	return spec, status
}

func resolveObservedFields(resource *Resource, krmState map[string]interface{}) map[string]interface{} {
	observedFields := make(map[string]interface{})
	for _, f := range *resource.ResourceConfig.ObservedFields {
		// TODO(b/314840974): Remove the check once the reference fields are supported.
		if ok, _ := IsReferenceField(f, &resource.ResourceConfig); ok {
			panic(fmt.Errorf("reference fields are not supported as observed fields"))
		}
		// TODO(b/314841141): Remove the check once the labels fields are supported.
		if isMetadataMappingLabelsField(f, &resource.ResourceConfig) {
			panic(fmt.Errorf("fields mapping to metadata.labels are not supported as observed fields"))
		}
		// TODO(b/314842047): Remove the check once the resource name fields are supported.
		if isMetadataMappingNameField(f, &resource.ResourceConfig) ||
			isServerGeneratedIDField(f, &resource.ResourceConfig) {
			panic(fmt.Errorf("fields of resource names are not supported as observed fields"))
		}
		addFieldIfExists(strings.Split(f, "."), resource.TFResource.Schema, krmState, observedFields)
	}
	return observedFields
}

func addFieldIfExists(path []string, tfSchemas map[string]*tfschema.Schema, source, parent map[string]interface{}) {
	if len(path) == 0 {
		return
	}
	field := text.SnakeCaseToLowerCamelCase(path[0])
	fieldState, ok := source[field]
	if !ok {
		return
	}

	fieldSchema, ok := tfSchemas[path[0]]
	if !ok {
		panic(fmt.Errorf("field %v not existent in the TF schema", path[0]))
	}
	// TODO(b/314841744): Remove after sensitive fields are supported.
	if tfresource.IsSensitiveField(fieldSchema) {
		panic(fmt.Errorf("sensitive fields are not supported as observed fields"))
	}

	if len(path) == 1 {
		parent[field] = fieldState
		return
	}

	switch fieldSchema.Type {
	case tfschema.TypeList, tfschema.TypeSet:
		if fieldSchema.MaxItems != 1 {
			panic(fmt.Errorf("invalid max items size %v of schema type tfschema.TypeList / tfschema.TypeSet for a nested field %v", fieldSchema.MaxItems, path[0]))
		}
		// Support the nested object field.
		subResource, ok := fieldSchema.Elem.(*tfschema.Resource)
		if !ok {
			panic(fmt.Errorf("type for schema elem under field %v should be *tfschema.Resource but got %T", path[0], fieldSchema.Elem))
		}
		subSchema := subResource.Schema
		fieldStateMap, ok := fieldState.(map[string]interface{})
		if !ok {
			panic(fmt.Errorf("retrieved fieldState of nested field %v is not of type map[string]interface{}: %v", field, fieldState))
		}
		value, ok := parent[field].(map[string]interface{})
		if !ok {
			value = make(map[string]interface{})
		}
		addFieldIfExists(path[1:], subSchema, fieldStateMap, value)
		if len(value) > 0 {
			parent[field] = value
		}
		return
	default:
		// TODO(b/312581557): Handle array types.
		panic(fmt.Errorf("invalid schema type %v for a nested field %v", fieldSchema.Type, path[0]))
	}
}

// ResolveSpecAndStatusWithResourceID returns the resolved spec and status with the `resourceID`
// field is populated in the KRM spec.
func ResolveSpecAndStatusWithResourceID(resource *Resource, state *terraform.InstanceState) (
	spec map[string]interface{}, status map[string]interface{}) {
	spec, status = ResolveSpecAndStatus(resource, state)

	resourceID, ok := getResourceIDIfSupported(resource, status)
	if !ok {
		return spec, status
	}

	if spec == nil {
		spec = make(map[string]interface{})
	}
	spec[k8s.ResourceIDFieldName] = resourceID
	return spec, status
}

// resolveDesiredStateInSpecAndObservedStateInStatus resolves spec as desired state and persists observed state in status.
func resolveDesiredStateInSpecAndObservedStateInStatus(resource *Resource, state *terraform.InstanceState) (
	spec map[string]interface{}, status map[string]interface{}) {
	spec = deepcopy.MapStringInterface(resource.Spec)
	returnedSpec, status := GetSpecAndStatusFromState(resource, state)
	// 'spec.additionalExperiments' in DataflowJob needs to be replaced by the
	// returned value because the API also updated the field and we need to
	// match the desired state with the live state.
	if resource.Kind == "DataflowJob" {
		if _, ok := spec["additionalExperiments"]; ok {
			if _, ok := returnedSpec["additionalExperiments"]; !ok {
				klog.Error("Kind", resource.Kind, "NamespacedName", resource.GetNamespacedName(), "The returned state doesn't contain additionalExperiments but it should")
			} else {
				spec["additionalExperiments"] = returnedSpec["additionalExperiments"]
			}
		}
	}
	return spec, status
}

// There are three scenarios for which we get the location value
//  1. It is in the spec.location and has been supplied by the customer and is likely an "easy" value like us-central1
//  2. It is in the state after a terraform import, i.e. the resource name was converted to TF state, and again it has an easy value like #1
//  3. It is in the state after a terraform Read, in which case it likely came from in a GET response from the given service
//     and is likely the fully qualified region / zone URL
//
// It is desired that we retain the 'easy' names from #1 and #2 so those are given precedence
func getLocationValueFromResourceOrState(resource *Resource, state map[string]interface{}) (interface{}, bool) {
	// the value for 'zone' returned by GCP APIs is the full ResourceName, i.e.
	//   https://www.googleapis.com/compute/v1/projects/project-id/zones/us-west2-a
	// to prevent overwriting the 'easier' zone value that most people use in their configs, return the previous
	// value if it is in the spec
	if location, ok := resource.Spec["location"]; ok {
		return location, true
	}
	if resource.ResourceConfig.Locationality == "" {
		return "", false
	}
	switch resource.ResourceConfig.Locationality {
	case gcp.Global:
		return "global", true
	case gcp.Regional, gcp.Zonal:
		locationFieldName := getTFFieldNameForLocation(resource.ResourceConfig.Locationality)
		value, ok := state[locationFieldName]
		if !ok {
			return nil, false
		}
		return value, true
	}
	panic(fmt.Errorf("unknown location type: %v", resource.ResourceConfig.Locationality))
}

func getTFFieldNameForLocation(locType string) string {
	switch locType {
	case gcp.Regional:
		return "region"
	case gcp.Zonal:
		return "zone"
	case gcp.Global:
		return "global"
	}
	panic(fmt.Errorf("unknown location type: %v", locType))
}

func getResourceIDIfSupported(resource *Resource, status map[string]interface{}) (interface{}, bool) {
	if !SupportsResourceIDField(&resource.ResourceConfig) {
		return nil, false
	}

	if resourceID, ok := resource.Spec[k8s.ResourceIDFieldName]; ok {
		return resourceID, true
	}

	if IsResourceIDFieldServerGenerated(&resource.ResourceConfig) {
		serverGeneratedIDFromStatus, exists, err :=
			getServerGeneratedIDFromStatus(&resource.ResourceConfig, resource.GroupVersionKind(), status)
		if !exists || err != nil {
			panic(fmt.Errorf("server-generated resource ID not "+
				"returned for resource Kind '%s', Name '%s', Namespace '%s'",
				resource.Kind, resource.Name, resource.Namespace))
		}

		resourceID, err := extractValueSegmentFromIDInStatus(
			serverGeneratedIDFromStatus,
			resource.ResourceConfig.ResourceID.ValueTemplate)
		if err != nil {
			panic(fmt.Errorf("incorrect format of server-generated "+
				"resource ID for resource Kind '%s', Name '%s', Namespace "+
				"'%s': %w", resource.Kind, resource.Name, resource.Namespace, err))
		}

		return resourceID, true
	}

	resourceID := resource.GetName()
	if resourceID == "" {
		panic(fmt.Errorf("user-specified resource ID not found for resource "+
			"Kind '%s', Name '%s', Namespace '%s'", resource.Kind, resource.Name,
			resource.Namespace))
	}

	return resourceID, true
}

func GetLabelsFromState(resource *Resource, rawState *terraform.InstanceState) map[string]string {
	state := InstanceStateToMap(resource.TFResource, rawState)
	labelsValue, ok := getNestedMapFromState(state, strings.Split(resource.ResourceConfig.MetadataMapping.Labels, ".")...)
	if !ok {
		return make(map[string]string)
	}
	result := make(map[string]string, len(labelsValue))
	for k, v := range labelsValue {
		result[k] = v.(string)
	}
	return result
}

func GetEtagFromState(resource *Resource, rawState *terraform.InstanceState) string {
	state := InstanceStateToMap(resource.TFResource, rawState)
	etagValue, ok := getNestedFieldFromState(state, "etag")
	if ok && etagValue != nil {
		return etagValue.(string)
	}
	return ""
}

func GetNameFromState(resource *Resource, rawState *terraform.InstanceState) string {
	if resource.ResourceConfig.MetadataMapping.Name == "" {
		return ""
	}
	state := InstanceStateToMap(resource.TFResource, rawState)
	nameValue, ok := getNestedFieldFromState(state, strings.Split(resource.ResourceConfig.MetadataMapping.Name, ".")...)
	if ok && nameValue != nil {
		return nameValue.(string)
	}
	return ""
}

// returns a nested map from within the state by traversing down the state with the path defined by the 'fields' list parameter
// if no such field path exists then the second parameter is false
// if there is a type mismatch then panic
func getNestedMapFromState(state map[string]interface{}, fields ...string) (map[string]interface{}, bool) {
	value, ok := getNestedFieldFromState(state, fields...)
	if !ok || value == nil {
		return nil, false
	}
	result, ok := value.(map[string]interface{})
	if !ok {
		panic(fmt.Sprintf("expected type '%v' instead got '%v'", reflect.TypeOf(make(map[string]interface{})).Name(),
			reflect.TypeOf(value).Name()))
	}
	return result, true
}

// returns a nested field from within the state by traversing down the state with the path defined by the 'fields' list parameter,
// stripping out the lists of length one that terraform inserts
// if no such field path exists then the second parameter is false
// if a field within the path is not a map then panic
func getNestedFieldFromState(state map[string]interface{}, fields ...string) (interface{}, bool) {
	var result interface{}
	result = state
	for i := 0; i < len(fields); i++ {
		subMap, ok := result.(map[string]interface{})
		if !ok {
			panic(formatUnexpectedValueTypeInStateMessage(result, i-1, fields...))
		}
		result, ok = getFieldFromStateMap(subMap, fields[i])
		// an 'ok' value of false indicates no value, but there are cases where the stored value is 'nil' so we need to check for that as well
		if !ok || result == nil {
			return nil, false
		}
	}
	return result, true
}

func formatUnexpectedValueTypeInStateMessage(value interface{}, fieldNum int, fields ...string) string {
	expectedType := reflect.TypeOf(make(map[string]interface{})).Name()
	actualType := reflect.TypeOf(value).Name()
	if fieldNum < 0 {
		return fmt.Sprintf("expected type '%v' instead got '%v'", expectedType, actualType)
	}
	return fmt.Sprintf("expected '%v' to be of type '%v' instead got '%v'", fields[fieldNum], expectedType, actualType)
}

func getFieldFromStateMap(state map[string]interface{}, field string) (interface{}, bool) {
	value, ok := state[field]
	if !ok {
		return nil, false
	}
	// the response returned by terraform will insert a list of size 1 for nested fields
	if listVal, ok := value.([]interface{}); ok {
		return listVal[0], true
	}
	return value, true
}

// Get the directives and container annotation(s) from the state
func GetAnnotationsFromState(resource *Resource, rawState *terraform.InstanceState) map[string]string {
	annotations := make(map[string]string, len(resource.ResourceConfig.Directives)+1)
	state := InstanceStateToMap(resource.TFResource, rawState)
	for _, directive := range resource.ResourceConfig.Directives {
		if isIgnoredField(directive, &resource.ResourceConfig) {
			continue
		}
		value, ok := getValueFromState(state, directive)
		if !ok {
			continue
		}
		key := k8s.FormatAnnotation(text.SnakeCaseToKebabCase(directive))
		annotations[key] = value
	}
	if !SupportsHierarchicalReferences(&resource.ResourceConfig) {
		// TODO(b/193177782): Delete this if-block once all resources support
		// hierarchical references.
		for _, c := range resource.ResourceConfig.Containers {
			value, ok := getValueFromState(state, c.TFField)
			if !ok {
				continue
			}
			if valueMatchesTemplate(c.ValueTemplate, value) {
				key := k8s.GetAnnotationForContainerType(c.Type)
				annotations[key] = value
			}
		}
	}
	return annotations
}

func getValueFromState(state map[string]interface{}, key string) (string, bool) {
	value, ok := state[key]
	// the state map contains all possible keys with 'nil' for missing values
	if !ok || value == nil {
		return "", false
	}
	stringValue := fmt.Sprintf("%v", value)
	if stringValue == "" {
		return "", false
	}
	return stringValue, true
}

// ConvertTFObjToKCCObj takes the state (which should be a Terraform resource),
// and returns two maps: the first one is formatted to KCC's custom resource
// schema for the appropriate Kind, the second one contains additional
// output-only fields that are used in observed state only.
//
// prevSpec is used for multiple purposes:
//   - ensures the returned result has a similar order for objects in lists, reducing
//     the perceived diff when applied.
//   - if server-side apply is used, the prevSpec value for a field will be used over
//     the value in state if it is managed by KCC.
//   - for sets (which are represented as lists), the result is a merger of both the
//     state and the prevSpec.
func ConvertTFObjToKCCObj(state map[string]interface{}, prevSpec map[string]interface{},
	schemas map[string]*tfschema.Schema, rc *corekccv1alpha1.ResourceConfig, prefix string,
	managedFields *fieldpath.Set) (krmState, krmStateWithIgnoredOutputOnlySpecFields map[string]interface{}) {
	rawKRMState := convertTFMapToKCCMap(state, prevSpec, schemas, rc, prefix, managedFields, true)
	rawKRMStateWithIgnoredOutputOnlySpecFields := deepcopy.DeepCopy(rawKRMState)
	if rc.IgnoredOutputOnlySpecFields != nil {
		rawKRMStateWithIgnoredOutputOnlySpecFields =
			convertTFMapToKCCMap(state, prevSpec, schemas, rc, prefix, managedFields, false)
	}
	// Round-trip via JSON in order to ensure consistency with unstructured.Unstructured's Object type.
	var retKRMState map[string]interface{}
	if err := util.Marshal(rawKRMState, &retKRMState); err != nil {
		panic(fmt.Errorf("error normalizing KRM-ified object: %w", err))
	}
	var retKRMStateWithIgnoredOutputOnlySpecFields map[string]interface{}
	if err := util.Marshal(rawKRMStateWithIgnoredOutputOnlySpecFields, &retKRMStateWithIgnoredOutputOnlySpecFields); err != nil {
		panic(fmt.Errorf("error normalizing KRM-ified object: %w", err))
	}
	return retKRMState, retKRMStateWithIgnoredOutputOnlySpecFields
}

func convertTFMapToKCCMap(state map[string]interface{}, prevSpec map[string]interface{},
	schemas map[string]*tfschema.Schema, rc *corekccv1alpha1.ResourceConfig, prefix string,
	managedFields *fieldpath.Set, ignoreOutputOnlySpecFields bool) map[string]interface{} {
	ret := make(map[string]interface{})
	for field, schema := range schemas {
		qualifiedName := field
		if prefix != "" {
			qualifiedName = prefix + "." + field
		}
		if isOverriddenField(qualifiedName, rc, ignoreOutputOnlySpecFields) {
			continue
		}
		if ok, refConfig := IsReferenceField(qualifiedName, rc); ok {
			key := GetKeyForReferenceField(refConfig)
			if val := convertTFReferenceToKCCReference(field, key, state, prevSpec, refConfig); val != nil {
				ret[key] = val
			}
			continue
		}
		key := text.SnakeCaseToLowerCamelCase(field)
		stateVal := state[field]
		prevSpecVal := prevSpec[key]
		if stateVal == nil {
			// Since partial terraform state are supported, if the next state is nil, we can
			// omit including them in the returned value.
			//
			// The one exception is if the field is managed by KCC. In this case,
			// it is assumed that "prevSpec" is the desired specification by the user,
			// and we replicate the managedField check that occurs when stateVal is non-nil.
			if prevSpecVal != nil && k8s.IsK8sManaged(key, prevSpec, managedFields) {
				ret[key] = prevSpecVal
			}
			continue
		}
		if isGCPManagedField(rc.Kind, qualifiedName) {
			ret[key] = stateVal
			continue
		}
		switch schema.Type {
		// Note:
		// - The provider will add defaulted "zero" values to any unset fields:
		// 	 https://github.com/hashicorp/terraform/blob/f9f73204383953e1b7fb91af6c56573cc0be2c02/helper/schema/field_reader.go#L287
		//   This adds a lot of noise to the KRM resource, so we prune these if they were not explicitly set
		//   by the user and does not have an explicit default (in which case the zero value would imply it was set explicitly).
		// - Certain APIs allow fields to be input in multiple different formats, and expand to
		//   a canonical form on the server. We want to keep the format that the user specified.
		case tfschema.TypeBool:
			if k8s.IsK8sManaged(key, prevSpec, managedFields) {
				ret[key] = prevSpecVal
			} else if schema.Required || stateVal.(bool) || (schema.Default != nil && schema.Default != stateVal) {
				ret[key] = stateVal
			}
		case tfschema.TypeFloat, tfschema.TypeInt:
			// The conversion from cty.Value to map[string]interface{} via JSON marshaling
			// will cause all numeric values to be of type float64.
			if k8s.IsK8sManaged(key, prevSpec, managedFields) {
				ret[key] = prevSpecVal
			} else if schema.Required || stateVal.(float64) != 0 || (schema.Default != nil && schema.Default != stateVal) {
				ret[key] = stateVal
			}
		case tfschema.TypeString:
			if k8s.IsK8sManaged(key, prevSpec, managedFields) {
				ret[key] = prevSpecVal
			} else {
				if stateVal.(string) == "" {
					continue
				}
				if tfresource.IsSensitiveConfigurableField(schema) {
					switch rc.Name {
					case "google_sql_database_instance",
						"google_compute_backend_service",
						"google_compute_region_backend_service":
						continue
					}
					val := stateVal.(string)
					ret[key] = corekccv1alpha1.SensitiveField{
						Value: &val,
					}
				} else {
					ret[key] = stateVal
				}
			}
		case tfschema.TypeList, tfschema.TypeSet:
			list, ok := stateVal.([]interface{})
			if !ok {
				panic(fmt.Sprintf("interface conversion for field %v in resource %v: interface {} is %T, not []interface {}. prevSpecVal: %v, stateVal: %v", qualifiedName, rc.Name, stateVal, prevSpecVal, stateVal))
			}
			if len(list) == 0 {
				continue
			}
			if schema.MaxItems == 1 {
				// A list with MaxItems == 1 is actually a nested object due to limitations with TF schemas.
				tfObjMap := list[0].(map[string]interface{})
				tfObjSchema := schema.Elem.(*tfschema.Resource).Schema
				prevObjMap, _ := prevSpecVal.(map[string]interface{})
				var nestedManagedFields *fieldpath.Set
				if managedFields != nil {
					pe := fieldpath.PathElement{FieldName: &key}
					var found bool
					nestedManagedFields, found = managedFields.Children.Get(pe)
					if !found {
						nestedManagedFields = fieldpath.NewSet()
					}
				}
				if val := convertTFMapToKCCMap(tfObjMap, prevObjMap, tfObjSchema, rc, qualifiedName, nestedManagedFields, ignoreOutputOnlySpecFields); val != nil {
					ret[key] = val
				}
				continue
			}
			if schema.Type == tfschema.TypeSet {
				// Sets in the spec require extra care in mapping elements from the previous spec
				// to the new one, as the ordering may have changed in the returned state. Sets in
				// the status can be treated the same as lists, as the new state is the definitive
				// source of truth and there is no reference resolution.
				if schema.Required || schema.Optional {
					retObj := convertTFSetToKCCSet(stateVal, prevSpecVal, schema, rc, qualifiedName, ignoreOutputOnlySpecFields)
					if retObj != nil {
						ret[key] = retObj
					}
					continue
				}
			}
			// A list may be either a list of primitives or a list of resources.
			switch schema.Elem.(type) {
			case *tfschema.Schema:
				// If it's a list of primitives, there is no conversion required
				ret[key] = deepcopy.DeepCopy(list)
			case *tfschema.Resource:
				// It's a list of the same type of resource. Convert each one using the same schema.
				prevList, _ := prevSpecVal.([]interface{})
				tfObjSchema := schema.Elem.(*tfschema.Resource).Schema
				retObjList := make([]interface{}, 0)
				for idx, elem := range list {
					tfObjMap := elem.(map[string]interface{})
					var prevObjMap map[string]interface{}
					if idx < len(prevList) {
						prevObjMap, _ = prevList[idx].(map[string]interface{})
					}
					if val := convertTFMapToKCCMap(tfObjMap, prevObjMap, tfObjSchema, rc, qualifiedName, nil, ignoreOutputOnlySpecFields); val != nil {
						retObjList = append(retObjList, val)
					}
				}
				if len(retObjList) == 0 {
					continue
				}
				ret[key] = retObjList
			}
		case tfschema.TypeMap:
			if k8s.IsK8sManaged(key, prevSpec, managedFields) {
				ret[key] = prevSpecVal
				continue
			}
			m := stateVal.(map[string]interface{})
			// Prune empty maps defaulted by the provider
			if len(m) == 0 {
				continue
			}
			// In this case, we do not convert from snake_case to camelCase, as the
			// keys here are user-provided
			ret[key] = deepcopy.DeepCopy(m)
		case tfschema.TypeInvalid:
			panic("invalid schema type")
		default:
			panic(fmt.Errorf("unrecognized schema type %v", schema.Type))
		}
	}
	if len(ret) == 0 {
		return nil
	}
	return ret
}

// convertTFReferenceToKCCReference converts the value of a TF reference field
// to a KCC reference value. The value of a TF reference field can either be a
// string or a list of strings. This function handles both cases.
func convertTFReferenceToKCCReference(tfField, specKey string, state map[string]interface{}, prevSpec map[string]interface{}, refConfig *corekccv1alpha1.ReferenceConfig) interface{} {
	if prevSpecVal, ok := prevSpec[specKey]; ok {
		// The user already specified a value for the KCC reference field in
		// the previous spec. Preserve it.
		return prevSpecVal
	}

	if state[tfField] == nil {
		return nil
	}

	// The user did not specify a value for the KCC reference field in the
	// previous spec, but the TF state has a value for the TF reference field.
	// Convert the value of the TF reference field to a KCC reference field
	// value.
	switch stateVal := state[tfField].(type) {
	case string:
		if stateVal == "" {
			return nil
		}
		if len(refConfig.Types) > 0 {
			// Get the first item in the list of types -- for now this is the defaulted ref
			defaultType := refConfig.Types[0]
			if defaultType.JSONSchemaType != "" {
				return map[string]interface{}{
					defaultType.Key: stateVal,
				}
			}

			return map[string]interface{}{
				defaultType.Key: corekccv1alpha1.ResourceReference{
					External: stateVal,
				},
			}
		}
		return corekccv1alpha1.ResourceReference{
			External: stateVal,
		}
	case []interface{}:
		if len(stateVal) == 0 {
			return nil
		}
		refs := make([]interface{}, 0)
		for _, elem := range stateVal {
			var newRef interface{}
			newRef = corekccv1alpha1.ResourceReference{
				External: elem.(string),
			}
			// this is a repeat of the same short-term fix made above for the string case when Types is an array
			if len(refConfig.Types) > 0 {
				newRef = map[string]interface{}{
					refConfig.Types[0].Key: newRef,
				}
			}
			refs = append(refs, newRef)
		}
		return refs
	default:
		panic(fmt.Errorf("value of TF reference field '%v' was neither a string nor a list", tfField))
	}
}

// convertTFSetToKCCSet converts a set object in Terraform to a KCC set object
func convertTFSetToKCCSet(stateVal, prevSpecVal interface{}, schema *tfschema.Schema, rc *corekccv1alpha1.ResourceConfig, prefix string, ignoreOutputOnlySpecFields bool) interface{} {
	if containsReferenceField(prefix, rc) {
		// TODO(kcc-eng): Support the case where the hashing function depends on resolved values from
		//  resource references. For the time being, fall back to the declared state.
		return prevSpecVal
	}
	list := stateVal.([]interface{})
	if len(list) == 0 {
		return nil
	}
	// Get the hash for each of the values in our new state.
	hashFunc := getHashFuncForSchema(schema)
	stateHashMap := make(map[int]interface{})
	for _, val := range list {
		stateHashMap[hashFunc(asHashable(val, schema.Elem))] = val
	}
	// convert each element from the state, but adhering to the ordering from the user-defined spec in order to
	// keep consistency when the user applies their config
	prevList, _ := prevSpecVal.([]interface{})
	retObjList := make([]interface{}, 0)
	for _, prevElem := range prevList {
		if prevElem == nil {
			retObjList = append(retObjList, nil)
			continue
		}
		var prevHashable interface{}
		switch schemaElem := schema.Elem.(type) {
		case *tfschema.Schema:
			prevHashable = asHashable(prevElem, schemaElem)
		case *tfschema.Resource:
			// convert the KRM previous spec object to a TF object so that we can calculate the correct hash
			prevElemAsTFObject, err := KRMObjectToTFObject(prevElem.(map[string]interface{}), schemaElem)
			if err != nil {
				panic(fmt.Errorf("error converting set object: %w", err))
			}
			prevHashable = asHashable(prevElemAsTFObject, schemaElem)
		default:
			panic(fmt.Errorf("unknown schema element type %v", schemaElem))
		}
		hash := hashFunc(prevHashable)
		stateElem, ok := stateHashMap[hash]
		// if the value is not in the stateHashMap, then it has been removed from the
		// new spec.
		if ok {
			delete(stateHashMap, hash)
		} else {
			stateElem = map[string]interface{}{}
		}
		retObjList = append(retObjList,
			convertTFElemToKCCElem(schema.Elem, stateElem, prevElem, rc, prefix, ignoreOutputOnlySpecFields))
	}
	// append any new elements in the list to the end
	for _, newElem := range stateHashMap {
		retObjList = append(retObjList,
			convertTFElemToKCCElem(schema.Elem, newElem, nil, rc, prefix, ignoreOutputOnlySpecFields))
	}
	if len(retObjList) == 0 {
		return nil
	}
	return retObjList
}

func getHashFuncForSchema(schema *tfschema.Schema) tfschema.SchemaSetFunc {
	// Determine the hashing function. If none is provided by the provider, then the defaults are used.
	hashFunc := schema.Set
	if hashFunc == nil {
		switch schemaElem := schema.Elem.(type) {
		case *tfschema.Schema:
			hashFunc = tfschema.HashSchema(schemaElem)
		case *tfschema.Resource:
			hashFunc = tfschema.HashResource(schemaElem)
		}
	}
	return hashFunc
}

func asHashable(o, schemaElem interface{}) interface{} {
	if o == nil {
		return nil
	}
	switch schemaElem := schemaElem.(type) {
	case *tfschema.Schema:
		// There is no reader available except for a map reader, so for primitives we convert the field
		// to a map with just one element
		key := "k"
		reader := tfschema.MapFieldReader{
			Map:    tfschema.BasicMapReader(map[string]string{key: fmt.Sprintf("%v", o)}),
			Schema: map[string]*tfschema.Schema{key: schemaElem},
		}
		val, err := reader.ReadField([]string{key})
		if err != nil {
			panic(fmt.Errorf("unable to convert field to hashable: %w", err))
		}
		var ret interface{}
		if val.Exists {
			ret = val.Value
		}
		return ret
	case *tfschema.Resource:
		// In order to hash an object in a set, we must have the object represented in a form that
		// can be parsed by the Terraform hashing functions. This structure is exactly like our
		// map[string]interface{} representations of TF objects, but substitutes any sets represented
		// by []interface{} into a *tfschema.Set.
		m := o.(map[string]interface{})
		reader := tfschema.MapFieldReader{
			Map:    tfschema.BasicMapReader(MapToInstanceState(schemaElem, m).Attributes),
			Schema: schemaElem.Schema,
		}
		res := make(map[string]interface{})
		for k, s := range schemaElem.Schema {
			val, err := reader.ReadField([]string{k})
			if err != nil {
				panic(fmt.Errorf("unable to read field %v: %w", k, err))
			}
			if val.Exists {
				res[k] = val.Value
			} else {
				res[k] = getDefaultValueForTFType(s.Type)
			}
		}
		return res
	default:
		panic(fmt.Errorf("unknown schema element type %v", schemaElem))
	}
}

func getDefaultValueForTFType(tfType tfschema.ValueType) interface{} {
	switch tfType {
	case tfschema.TypeBool:
		return false
	case tfschema.TypeString:
		return ""
	case tfschema.TypeFloat:
		return 0.0
	case tfschema.TypeInt:
		return 0
	case tfschema.TypeList:
		return make([]interface{}, 0)
	case tfschema.TypeSet:
		return &tfschema.Set{}
	case tfschema.TypeMap:
		return make(map[string]interface{})
	case tfschema.TypeInvalid:
		panic("schema type is invalid")
	default:
		panic(fmt.Errorf("unrecognized schema type %v", tfType))
	}
}

func convertTFElemToKCCElem(elemSchema, tfObj, prevSpecObj interface{}, rc *corekccv1alpha1.ResourceConfig, prefix string, ignoreOutputOnlySpecFields bool) interface{} {
	switch elemSchema.(type) {
	case *tfschema.Schema:
		if prevSpecObj != nil {
			return prevSpecObj
		}
		return tfObj
	case *tfschema.Resource:
		tfObjSchema := elemSchema.(*tfschema.Resource).Schema
		tfObjMap, _ := tfObj.(map[string]interface{})
		prevObjMap, _ := prevSpecObj.(map[string]interface{})
		return convertTFMapToKCCMap(tfObjMap, prevObjMap, tfObjSchema, rc, prefix, nil, ignoreOutputOnlySpecFields)
	default:
		return prevSpecObj
	}
}

func isOverriddenField(field string, rc *corekccv1alpha1.ResourceConfig, ignoreOutputOnlySpecFields bool) bool {
	if field == rc.MetadataMapping.Name || field == rc.MetadataMapping.Labels {
		return true
	}
	if rc.Locationality != "" && (field == "zone" || field == "region") {
		return true
	}
	if isIgnoredField(field, rc) {
		return true
	}
	for _, f := range rc.Directives {
		if field == f {
			return true
		}
	}
	if !SupportsHierarchicalReferences(rc) {
		// TODO(b/193177782): Delete this if-block once all resources support
		// hierarchical references.
		for _, c := range rc.Containers {
			if field == c.TFField {
				return true
			}
		}
	}
	if ignoreOutputOnlySpecFields && rc.IgnoredOutputOnlySpecFields != nil {
		for _, f := range *rc.IgnoredOutputOnlySpecFields {
			if field == f {
				return true
			}
		}
	}
	return false
}

func isIgnoredField(field string, rc *corekccv1alpha1.ResourceConfig) bool {
	for _, f := range rc.IgnoredFields {
		if field == f {
			return true
		}
	}
	return false
}

func renameStatusFieldIfNeeded(tfResourceName, key string) string {
	reservedNames := k8s.ReservedStatusFieldNames()
	if _, found := reservedNames[key]; found {
		return k8s.RenameStatusFieldWithReservedNameIfResourceNotExcluded(tfResourceName, key)
	}
	return key
}
