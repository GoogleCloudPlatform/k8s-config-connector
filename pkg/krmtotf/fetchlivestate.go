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
	"context"
	"fmt"
	"regexp"
	"slices"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// FetchLiveState is a combination of a resource import + read. It returns the state of the
// underlying resource as seen by the TF provider.
func FetchLiveState(ctx context.Context, resource *Resource, provider *tfschema.Provider, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*terraform.InstanceState, error) {
	// Get the ID to pass to the import
	id, err := resource.GetImportID(kubeClient, smLoader)
	if err != nil {
		if _, ok := k8s.AsServerGeneratedIDNotFoundError(err); ok {
			// If the import ID cannot be determined because it requires a server-
			// generated ID that has not been set, this means the resource has not
			// yet been created. Return as if the read returned a non-existent
			// resource.
			return &terraform.InstanceState{}, nil
		}
		return nil, fmt.Errorf("error getting ID for resource: %w", err)
	}
	return fetchLiveStateFromID(ctx, id, resource, provider, kubeClient, smLoader)

}

// ShouldResolveParentForDelete
// Special handling for KMSCryptoKey that still lives after its parent KMSKeyRing is deleted.
// For KMSCryptoKey resource, we can import the tf state directly from its selfLink instead of sourcing for its parent.
// More info in b/279485255#comment14
func ShouldResolveParentForDelete(resource *Resource) bool {
	allowlist := []string{"KMSCryptoKey"}
	return !slices.Contains(allowlist, resource.Kind) || hasEmptySelfLink(resource)
}

func hasEmptySelfLink(resource *Resource) bool {
	id, err := resource.SelfLinkAsID()
	if err != nil || id == "" {
		return true
	}
	return false
}

// ShouldCheckParentReadyForDelete
// Special handling for allowlist resources, when parent exists but has deletion failed error.
// Due to their API design, the allowlisted resources are deletable even if their parents are not ready.
// See b/306583728#comment8 for details.
func ShouldCheckParentReadyForDelete(resource *Resource, parent *k8s.Resource) bool {
	allowlist := []string{"AlloyDBInstance", "EdgeContainerNodePool", "TagsTagValue"}
	// If the resource kind in allowlist, we skip checking parent.
	if slices.Contains(allowlist, resource.Kind) {
		return false
	}
	// Skip checking parent with nested resource deletion error.
	return !isDeletionFailureDueToExistingDependent(parent)
}

func isDeletionFailureDueToExistingDependent(r *k8s.Resource) bool {
	if k8s.IsResourceReady(r) {
		return false
	}
	cond, _ := k8s.GetReadyCondition(r)
	// Full error message:
	// Resource '"projects/project/locations/location/clusters/cluster"' has nested resources.
	// If the API supports cascading delete, set 'force' to true to delete it and its nested resources.
	errorMessageRegex := ".*Resource .* has nested resources.*"
	match, _ := regexp.MatchString(errorMessageRegex, cond.Message)
	return match
}

func FetchLiveStateForDelete(ctx context.Context, resource *Resource, provider *tfschema.Provider, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*terraform.InstanceState, error) {
	if !ShouldResolveParentForDelete(resource) {
		id, err := resource.SelfLinkAsID()
		if err != nil {
			return nil, err
		}
		if id != "" {
			return fetchLiveStateFromID(ctx, id, resource, provider, kubeClient, smLoader)
		}
	}
	return FetchLiveState(ctx, resource, provider, kubeClient, smLoader)
}

func fetchLiveStateFromID(ctx context.Context, id string, resource *Resource, provider *tfschema.Provider, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*terraform.InstanceState, error) {
	// Get the imported resource
	var state *terraform.InstanceState
	var err error
	if resource.ResourceConfig.SkipImport {
		state = &terraform.InstanceState{ID: id}
	} else {
		state, err = ImportState(ctx, id, resource.TFInfo, provider)
		if err != nil {
			return nil, err
		}
	}

	// Given that some fields are input-only or may only be returned on creation,
	// e.g. private key, we need to stick with the previously captured values.
	state, err = presetFieldsForRead(resource, state, kubeClient, smLoader)
	if err != nil {
		return nil, err
	}
	state, diagnostics := resource.TFResource.RefreshWithoutUpgrade(ctx, state, provider.Meta())
	if err := NewErrorFromDiagnostics(diagnostics); err != nil {
		return nil, fmt.Errorf("error reading underlying resource: %w", err)
	}
	return state, nil
}

// FetchLiveStateForCreateAndUpdate is the same as FetchLiveState except for added special
// handling for certain types of resources during resource creation and update.
func FetchLiveStateForCreateAndUpdate(ctx context.Context, resource *Resource, provider *tfschema.Provider, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*terraform.InstanceState, error) {
	// Special handling for resource which cannot be imported or read, has user-specified resource ID,
	// and only contains top level fields that are immutable and/or computed.
	// For such resource, its fetched live state will always be identical to the user config,
	// regardless of the existence of the underlying GCP resource.
	// We need to set the live state to empty so that the controller can retrieve the
	// computed values via an explicit TF Apply() call.
	//
	// For example, ServiceIdentity is an unreadable resource with the user-specified
	// resource ID, and all of its `spec` fields are immutable. An empty InstanceState
	// ensures there can be a diff during the first reconciliation, so that TF
	// controller can retrieve the computed value of the `status.email` field (the
	// service identity) via the response of a TF Apply().
	if resource.Unreadable() &&
		resource.ResourceConfig.SkipImport &&
		!resource.HasServerGeneratedIDField() &&
		resource.AllTopLevelFieldsAreImmutableOrComputed() {
		return &terraform.InstanceState{}, nil
	}

	return FetchLiveState(ctx, resource, provider, kubeClient, smLoader)
}

// ImportState parses the given id into a TF state. Note that this function
// does not make any network calls; it simply does a best effort to determine
// TF state by parsing the id.
//
// As a result of this being best-effort, the returned state may not have
// every field required in a fully valid InstanceState.
func ImportState(ctx context.Context, id string, tfInfo *terraform.InstanceInfo, provider *tfschema.Provider) (*terraform.InstanceState, error) {
	importedResources, err := provider.ImportState(ctx, tfInfo, id)
	if err != nil {
		return nil, fmt.Errorf("error importing resource: %w", err)
	}
	if len(importedResources) != 1 {
		return nil, fmt.Errorf("import corresponds to more than one resource")
	}
	return importedResources[0], nil
}

func presetFieldsForRead(r *Resource, imported *terraform.InstanceState, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*terraform.InstanceState, error) {
	importedMap := InstanceStateToMap(r.TFResource, imported)
	ret, err := WithFieldsPresetForRead(importedMap, r, kubeClient, smLoader)
	if err != nil {
		return nil, err
	}
	return MapToInstanceState(r.TFResource, ret), nil
}

func WithFieldsPresetForRead(imported map[string]interface{}, r *Resource, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (map[string]interface{}, error) {
	var config *terraform.ResourceConfig
	var secretVersions map[string]string
	var err error
	// As we are directly calling the `krmResourceToTFResourceConfig`
	// helper method instead of using the exported wrapping functions,
	// define variables for all the arguments to improve readability.
	mustResolveSensitiveFields := !k8s.IsDeleted(&r.ObjectMeta)
	importedAsInstanceState := MapToInstanceState(r.TFResource, imported)
	var jsonSchema *apiextensions.JSONSchemaProps
	config, secretVersions, err = KRMResourceToTFResourceConfigFull(
		r, kubeClient, smLoader, importedAsInstanceState, jsonSchema, mustResolveSensitiveFields,
	)
	if err != nil {
		return nil, fmt.Errorf("error converting resource config: %w", err)
	}

	ret := withImmutableFields(imported, ResourceConfigToMap(config), r.TFResource.Schema)
	ret, err = withMutableButUnreadableFields(ret, r, secretVersions, kubeClient)
	if err != nil {
		return nil, fmt.Errorf("error presetting mutable but unreadable fields for read: %w", err)
	}
	ret = withDirectives(ret, r)
	ret, err = withStatusFields(ret, r, kubeClient, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error presetting status fields for read: %w", err)
	}
	return ret, nil
}

func withImmutableFields(imported, config map[string]interface{}, schemas map[string]*tfschema.Schema) map[string]interface{} {
	ret := deepcopy.MapStringInterface(imported)
	if ret == nil {
		ret = make(map[string]interface{})
	}
	for field, schema := range schemas {
		configVal := config[field]
		if schema.ForceNew {
			if configVal == nil {
				// If no value is specified by the user, prefill with the zero value.
				// This happens due to pruning of default zero values returned from
				// the read.
				ret[field] = getZeroValueForType(schema.Type)
			} else {
				ret[field] = configVal
			}
			continue
		}
		if configVal == nil {
			continue
		}
		// The current field is mutable, but may be a list of objects containing a field that
		// is immutable.
		switch schema.Type {
		case tfschema.TypeList, tfschema.TypeSet:
			switch elem := schema.Elem.(type) {
			case *tfschema.Resource:
				// Note: we assume that indexes in the list are preserved between the imported structure
				// and the expanded config structure.
				configList := configVal.([]interface{})
				importedList, _ := imported[field].([]interface{})
				retList := make([]interface{}, 0)
				for idx, expandedItem := range configList {
					expandedItem := expandedItem.(map[string]interface{})
					var importedItem map[string]interface{}
					if len(importedList) > idx {
						importedItem = importedList[idx].(map[string]interface{})
					}
					retList = append(retList, withImmutableFields(importedItem, expandedItem, elem.Schema))
				}
				ret[field] = retList
			}
		}
	}
	return ret
}

func withMutableButUnreadableFields(imported map[string]interface{}, r *Resource, currSecretVersions map[string]string, kubeClient client.Client) (map[string]interface{}, error) {
	if len(r.ResourceConfig.MutableButUnreadableFields) == 0 {
		return imported, nil
	}

	mutableButUnreadableFields, err := getMutableButUnreadableFieldsFromAnnotations(r)
	if err != nil {
		return nil, err
	}
	if len(mutableButUnreadableFields) == 0 {
		return imported, nil
	}

	secretVersions, err := k8s.GetSecretVersionsFromAnnotations(&r.Resource)
	if err != nil {
		return nil, err
	}
	// When secretVersions are not found in the annotations, there is a possibility
	// that this is either (1) a resource acquisition or (2) a resource created
	// before it supported the annotation. To avoid unnecessarily updating the
	// resource in both cases, use the current Secret versions.
	if secretVersions == nil {
		secretVersions = currSecretVersions
	}

	return setMutableButUnreadableFields(imported, mutableButUnreadableFields["spec"].(map[string]interface{}), r.TFResource.Schema, secretVersions, r.GetNamespace(), kubeClient)
}

func setMutableButUnreadableFields(imported, mutableButUnreadableSpec map[string]interface{}, schemas map[string]*tfschema.Schema, secretVersions map[string]string, namespace string, kubeClient client.Client) (map[string]interface{}, error) {
	ret := deepcopy.MapStringInterface(imported)
	for k, v := range mutableButUnreadableSpec {
		tfKey := text.AsSnakeCase(k)
		schema, ok := schemas[tfKey]
		if !ok {
			return nil, fmt.Errorf("could not find a schema for field %v", tfKey)
		}
		switch schema.Type {
		case tfschema.TypeString:
			if !schema.Sensitive {
				ret[tfKey] = v
				continue
			}

			sensitiveField := corekccv1alpha1.SensitiveField{}
			if err := util.Marshal(v, &sensitiveField); err != nil {
				return nil, fmt.Errorf("error parsing %v onto a SensitiveField struct: %w", v, err)
			}

			if sensitiveField.Value != nil {
				ret[tfKey] = *sensitiveField.Value
				continue
			}

			secretKeyRef := sensitiveField.ValueFrom.SecretKeyRef
			secretVal, secretVer, err := k8s.GetSecretVal(secretKeyRef, namespace, kubeClient)
			if err != nil {
				// If a previously referenced Secret cannot be resolved, it is
				// possible it simply no longer exists. Don't error out in this
				// case; just skip the presetting of the field in the live
				// state so that a diff is generated if the field had been
				// updated in the spec. If the field still points to the same
				// Secret in the spec, then the KRM2TF conversion will
				// appropriately error out due to the non-existent Secret.
				continue
			}
			// Preset sensitive field only if we can be sure that the
			// referenced Secret had not been changed.
			prevSecretVer, ok := secretVersions[secretKeyRef.Name]
			if ok && secretVer == prevSecretVer {
				ret[tfKey] = secretVal
			}
		case tfschema.TypeBool, tfschema.TypeFloat, tfschema.TypeInt:
			ret[tfKey] = v
		case tfschema.TypeMap:
			ret[tfKey] = deepcopy.DeepCopy(v)
		case tfschema.TypeList, tfschema.TypeSet:
			switch elem := schema.Elem.(type) {
			// List/set of primitives
			case *tfschema.Schema:
				ret[tfKey] = deepcopy.DeepCopy(v)
			// List/set of objects OR nested object
			case *tfschema.Resource:
				// List/set of objects
				if schema.MaxItems != 1 {
					panic(fmt.Errorf("error presetting field %v: presetting mutable-but-unreadable fields in objects contained in lists/sets is not yet supported", tfKey))
				}
				// Nested object
				prevObj, ok := v.(map[string]interface{}) // Nested objects are represented as maps in KRM
				if !ok {
					return nil, fmt.Errorf("expected field %v in %v to be a map, but it is not", k, k8s.MutableButUnreadableFieldsAnnotation)
				}
				importedObj, err := getObjectAtFieldInState(imported, tfKey)
				if err != nil {
					return nil, fmt.Errorf("error getting object at field %v from state map: %w", tfKey, err)
				}
				obj, err := setMutableButUnreadableFields(importedObj, prevObj, elem.Schema, secretVersions, namespace, kubeClient)
				if err != nil {
					return nil, err
				}
				if len(obj) == 0 {
					continue
				}
				ret[tfKey] = []interface{}{obj} // Nested objects are represented as lists with one item in TF
			}
		}
	}
	return ret, nil
}

// getObjectAtFieldInState gets the object at field 'tfKey' in the TF state map 'state'
func getObjectAtFieldInState(state map[string]interface{}, tfKey string) (map[string]interface{}, error) {
	v, ok := getNestedFieldFromState(state, tfKey)
	if !ok {
		return make(map[string]interface{}), nil
	}
	obj, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected field %v to be a nested object, but it is not", tfKey)
	}
	return obj, nil
}

func withDirectives(imported map[string]interface{}, r *Resource) map[string]interface{} {
	ret := deepcopy.MapStringInterface(imported)
	for _, d := range r.ResourceConfig.Directives {
		key := k8s.FormatAnnotation(text.SnakeCaseToKebabCase(d))
		if v, ok := k8s.GetAnnotation(key, r); ok {
			ret[d] = v
		} else {
			if r.TFResource.Schema[d].Default != nil {
				ret[d] = r.TFResource.Schema[d].Default
			}
		}
	}
	return ret
}

func withStatusFields(imported map[string]interface{}, r *Resource, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (map[string]interface{}, error) {
	ret := deepcopy.MapStringInterface(imported)
	tfStatus, err := KRMObjectToTFObject(r.GetStatusOrObservedState(), r.TFResource)
	if err != nil {
		return nil, fmt.Errorf("error converting status object: %w", err)
	}
	for k, v := range tfStatus {
		ret[k] = v
	}

	if SupportsResourceIDField(&r.ResourceConfig) && IsResourceIDFieldServerGenerated(&r.ResourceConfig) {
		idInStatus, err := r.ConstructServerGeneratedIDInStatusFromResourceID(kubeClient, smLoader)
		if err != nil {
			return nil, fmt.Errorf("error syncing the server-generated ID: %w", err)
		}
		if idInStatus != "" {
			ret[r.ResourceConfig.ServerGeneratedIDField] = idInStatus
		}
	}

	return ret, nil
}

func getZeroValueForType(valueType tfschema.ValueType) interface{} {
	switch valueType {
	case tfschema.TypeBool:
		return false
	case tfschema.TypeFloat, tfschema.TypeInt:
		return float64(0)
	case tfschema.TypeString:
		return ""
	case tfschema.TypeList, tfschema.TypeMap, tfschema.TypeSet:
		return nil
	default:
		panic(fmt.Sprintf("unknown value type %v", valueType))
	}
}
