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

package kcclite

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/deepcopy"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/pathslice"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/typeutil"

	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/structured-merge-diff/v4/fieldpath"
)

// ToKCCLite will convert the KRM representation to a state that all fields that can be used without using the API server.
// More specifically speaking, it will
// 1) resolve the resource reference and store the value in 'external' field
// 2) resolve the secret reference and store the value in 'value' field
func ToKCCLite(resource *dcl.Resource, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader,
	kubeClient client.Client) (*unstructured.Unstructured, error) {
	kccLite, _, err := convertToKCCLite(resource, smLoader, schemaLoader, serviceMappingLoader, kubeClient, true)
	return kccLite, err
}

func ToKCCLiteBestEffort(resource *dcl.Resource, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader,
	kubeClient client.Client) (*unstructured.Unstructured, error) {
	kccLite, _, err := convertToKCCLite(resource, smLoader, schemaLoader, serviceMappingLoader, kubeClient, false)
	return kccLite, err
}

func ToKCCLiteAndSecretVersions(resource *dcl.Resource, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader,
	kubeClient client.Client) (kccLite *unstructured.Unstructured, secretVersions map[string]string, err error) {
	return convertToKCCLite(resource, smLoader, schemaLoader, serviceMappingLoader, kubeClient, true)
}

func convertToKCCLite(resource *dcl.Resource, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader,
	kubeClient client.Client, mustResolveAllFields bool) (kccLite *unstructured.Unstructured, secretVersions map[string]string, err error) {
	lite, err := resource.MarshalAsUnstructured()
	if err != nil {
		return nil, nil, err
	}
	config, found, err := unstructured.NestedFieldNoCopy(lite.Object, "spec")
	if err != nil {
		return nil, nil, err
	}
	if !found || config == nil {
		return lite, nil, nil
	}

	secretVersions = make(map[string]string)
	convertedSpec, err := convertConfig(config.(map[string]interface{}), []string{}, resource.Schema, smLoader, schemaLoader, serviceMappingLoader, resource.GetNamespace(), kubeClient, mustResolveAllFields, secretVersions)

	if err != nil {
		return nil, nil, err
	}
	if err := unstructured.SetNestedMap(lite.Object, convertedSpec, "spec"); err != nil {
		return nil, nil, err
	}
	return lite, secretVersions, nil
}

func convertConfig(config map[string]interface{}, path []string, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader, schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader, namespace string, kubeClient client.Client, mustResolveAllFields bool, secretVersions map[string]string) (map[string]interface{}, error) {
	if len(config) == 0 {
		return config, nil
	}
	if schema.Type != "object" {
		return nil, fmt.Errorf("expect the schame type to be 'object', but got %v", schema.Type)
	}
	for f, s := range schema.Properties {
		if dclextension.IsReferenceField(s) {
			if err := handleReferenceField(append(path, f), config, s, smLoader, schemaLoader, serviceMappingLoader, kubeClient, namespace, mustResolveAllFields); err != nil {
				return nil, fmt.Errorf("error resolving reference field %v: %w", f, err)
			}
			continue
		}
		if config[f] != nil {
			convertedVal, err := convertVal(config[f], append(path, f), s, smLoader, schemaLoader, serviceMappingLoader, namespace, kubeClient, mustResolveAllFields, secretVersions)
			if err != nil {
				return nil, err
			}
			delete(config, f)
			// It's possible that convertVal() returns nil value (e.g. when a Secret
			// is not found) when mustResolveAllFields is false. We should ignore
			// unresolved field.
			if convertedVal != nil {
				dcl.AddToMap(f, convertedVal, config)
			}
		}
	}
	return config, nil
}

func convertVal(val interface{}, path []string, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader, schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader, namespace string, kubeClient client.Client, mustResolveAllFields bool, secretVersions map[string]string) (interface{}, error) {
	switch schema.Type {
	case "object":
		obj, ok := val.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the value to be map[string]interface{} but was actually %T", val)
		}
		if schema.AdditionalProperties != nil {
			if typeutil.IsPrimitiveType(schema.AdditionalProperties.Type) {
				return val, nil
			}
			if schema.AdditionalProperties.Type == "object" {
				res := make(map[string]interface{})
				for k, v := range obj {
					convertedVal, err := convertVal(v, append(path, k), schema.AdditionalProperties, smLoader, schemaLoader, serviceMappingLoader, namespace, kubeClient, mustResolveAllFields, secretVersions)
					if err != nil {
						return nil, fmt.Errorf("error converting the object value for key %v: %w", k, err)
					}
					res[k] = convertedVal
				}
				return res, nil
			}
			return nil, fmt.Errorf("not supported type for AdditionalProperties %v", schema.AdditionalProperties.Type)
		}
		return convertConfig(obj, path, schema, smLoader, schemaLoader, serviceMappingLoader, namespace, kubeClient, mustResolveAllFields, secretVersions)
	case "array":
		if typeutil.IsPrimitiveType(schema.Items.Type) {
			return val, nil
		}
		items, ok := val.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the value to be []interface{} but was actually %T", val)
		}
		res := make([]interface{}, 0)
		for _, item := range items {
			processedItem, err := convertVal(item, path, schema.Items, smLoader, schemaLoader, serviceMappingLoader, namespace, kubeClient, mustResolveAllFields, secretVersions)
			if err != nil {
				return nil, fmt.Errorf("error converting list item: %w", err)
			}
			res = append(res, processedItem)
		}
		return res, nil
	case "string":
		if ok, _ := dclextension.IsSensitiveField(schema); ok {
			field := corekccv1alpha1.SensitiveField{}
			if err := util.Marshal(val, &field); err != nil {
				return nil, fmt.Errorf("error parsing %v onto a SensitiveField struct: %w", val, err)
			}

			if field.Value != nil {
				return map[string]interface{}{"value": *field.Value}, nil
			}

			secretKeyRef := field.ValueFrom.SecretKeyRef
			secretVal, secretVer, err := k8s.GetSecretVal(secretKeyRef, namespace, kubeClient)
			if err != nil {
				if mustResolveAllFields {
					return nil, err
				}

				// If the secret can't be found but it is not required to be resolved, then
				// return nil secret value and nil error.
				return nil, nil
			}
			secretVersions[secretKeyRef.Name] = secretVer
			return map[string]interface{}{"value": secretVal}, nil
		}
		return val, nil
	case "boolean", "number", "integer":
		return val, nil
	default:
		return nil, fmt.Errorf("unknown schema type %v", schema.Type)
	}
}

func handleReferenceField(path []string, config map[string]interface{}, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader,
	kubeClient client.Client, namespace string, mustResolveAllFields bool) error {
	if dcl.IsMultiTypeParentReferenceField(path) {
		return handleMultiTypeParentReferenceField(config, schema, smLoader, kubeClient, namespace, mustResolveAllFields)
	}
	if schema.Type == "array" {
		return handleListOfReferencesField(path, config, schema, smLoader, schemaLoader, serviceMappingLoader, kubeClient, namespace, mustResolveAllFields)
	}
	return handleRegularReferenceField(path, config, schema, smLoader, schemaLoader, serviceMappingLoader, kubeClient, namespace, mustResolveAllFields)
}

func handleMultiTypeParentReferenceField(config map[string]interface{}, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader,
	kubeClient client.Client, namespace string, mustResolveAllFields bool) error {
	rawVal, tc, err := dcl.GetHierarchicalRefFromConfigForMultiParentResource(config, schema, smLoader)
	if err != nil {
		return fmt.Errorf("error getting hierarchical reference from config for multi-parent resource: %w", err)
	}
	if rawVal == nil {
		return nil
	}
	refField := tc.Key
	refObj, ok := rawVal.(map[string]interface{})
	if !ok {
		return fmt.Errorf("expected the value to be map[string]interface{} for reference field %v but was actually %T", refField, rawVal)
	}
	val, err := resolveHierarchicalReferenceForMultiParentResource(refObj, tc, namespace, kubeClient, smLoader)
	if err != nil {
		if mustResolveAllFields {
			return err
		}
		delete(config, refField)
		return nil
	}
	delete(config, refField)
	dcl.AddToMap(refField, val, config)
	return nil
}

func handleListOfReferencesField(path []string, config map[string]interface{}, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader,
	kubeClient client.Client, namespace string, mustResolveAllFields bool) error {
	refField, err := dclextension.GetReferenceFieldName(path, schema)
	if err != nil {
		return fmt.Errorf("error getting the reference field name %w", err)
	}
	if config[refField] == nil {
		return nil
	}
	rawVal := config[refField]
	items, ok := rawVal.([]interface{})
	if !ok {
		return fmt.Errorf("expected the value to be []interface{} for reference field %v but was actually %T", refField, rawVal)
	}
	res := make([]interface{}, 0)
	for _, item := range items {
		refObj, ok := item.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected the value for item reference to be map[string]interface{}, but was actually %T", item)
		}
		refVal, err := resolveResourceReference(refObj, schema.Items, smLoader, schemaLoader, serviceMappingLoader, kubeClient, namespace)
		if err != nil {
			if mustResolveAllFields {
				return err
			}
			continue
		}
		res = append(res, refVal)
	}
	delete(config, refField)
	if len(res) != 0 {
		config[refField] = res
	}
	return nil
}

func handleRegularReferenceField(path []string, config map[string]interface{}, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader,
	kubeClient client.Client, namespace string, mustResolveAllFields bool) error {
	refField, err := dclextension.GetReferenceFieldName(path, schema)
	if err != nil {
		return fmt.Errorf("error getting the reference field name %w", err)
	}
	if config[refField] == nil {
		return nil
	}
	rawVal := config[refField]
	refObj, ok := rawVal.(map[string]interface{})
	if !ok {
		return fmt.Errorf("expected the value to be map[string]interface{} for reference field %v but was actually %T", refField, rawVal)
	}
	val, err := resolveResourceReference(refObj, schema, smLoader, schemaLoader, serviceMappingLoader, kubeClient, namespace)
	if err != nil {
		if mustResolveAllFields {
			return err
		}
		delete(config, refField)
		return nil
	}
	delete(config, refField)
	dcl.AddToMap(refField, val, config)
	return nil
}

func resolveHierarchicalReferenceForMultiParentResource(resourceRefValRaw map[string]interface{}, tc *corekccv1alpha1.TypeConfig, ns string,
	kubeClient client.Client, _ dclmetadata.ServiceMetadataLoader) (map[string]interface{}, error) {
	val, err := resolveReferenceObject(resourceRefValRaw, tc, ns, kubeClient)
	if err != nil {
		return nil, err
	}
	// Use the original resolved value of the reference object here. That is,
	// don't do any canonicalization even if the target field is the referenced
	// resource's name (i.e. what we do for other resource references in
	// resolveResourceReference()). This is because multi-type parent reference
	// fields in DCL expect a different format from the one provided by the
	// referenced resource's x-dcl-id (e.g. "projects/{project_id}",
	// "folders/{folder_id}", etc.). Since we handle the formatting of
	// multi-type parent reference fields at the KCCLite->DCL layer (e.g.
	// convert "project-id" to "projects/project-id"), let us just use the
	// original resolved value of the reference object here at the KCC->KCCLite
	// layer (e.g. "project-id").
	return map[string]interface{}{
		"external": val,
	}, nil
}

func resolveResourceReference(resourceRefValRaw map[string]interface{}, schema *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader,
	kubeClient client.Client, ns string) (map[string]interface{}, error) {
	tcs, err := dcl.GetReferenceTypeConfigs(schema, smLoader)
	if err != nil {
		return nil, err
	}

	// If the reference object was originally an external reference, don't do
	// anything extra.
	if _, ok := resourceRefValRaw["external"]; ok {
		return resourceRefValRaw, nil
	}

	tc, refResource, err := getMultiTypeReferencedResource(resourceRefValRaw, tcs, ns, kubeClient)
	if err != nil {
		return nil, err
	}
	val, err := resolveTargetFieldValue(refResource, tc)
	if err != nil {
		return nil, fmt.Errorf("error resolving target field value for referenced resource %v with GroupVersionKind %v: %w",
			k8s.GetNamespacedName(refResource), refResource.GroupVersionKind(), err)
	}
	// Canonicalize the resolved target field value if the target field value
	// is the referenced resource's name.
	if tc.TargetField == "name" {
		s, err := dclschemaloader.GetDCLSchemaForGVK(tc.GVK, smLoader, schemaLoader)
		if err != nil {
			return nil, fmt.Errorf("error getting DCL schema for referenced GroupVersionKind %v: %w", tc.GVK, err)
		}
		template, err := dclextension.GetNameValueTemplate(s)
		if err != nil {
			return nil, fmt.Errorf("error getting name value template for referenced GroupVersionKind %v: %w", tc.GVK, err)
		}
		canonicalizedVal, err := CanonicalizeReferencedResourceName(val, template, refResource, smLoader, schemaLoader, serviceMappingLoader, kubeClient)
		if err != nil {
			return nil, fmt.Errorf("error canonicalizing name of referenced resource %v with GroupVersionKind %v: %w",
				k8s.GetNamespacedName(refResource), refResource.GroupVersionKind(), err)
		}
		return map[string]interface{}{
			"external": canonicalizedVal,
		}, nil
	}
	return map[string]interface{}{
		"external": val,
	}, nil
}

func resolveReferenceObject(resourceRefValRaw map[string]interface{}, tc *corekccv1alpha1.TypeConfig,
	ns string, kubeClient client.Client) (string, error) {
	if rawVal, ok := resourceRefValRaw["external"]; ok {
		val, ok := rawVal.(string)
		if !ok {
			return "", fmt.Errorf("expected the value of 'external' in the resource reference object to be string, but was actually %T", rawVal)
		}
		return val, nil
	}
	refResource, err := getReferencedResource(resourceRefValRaw, tc, ns, kubeClient)
	if err != nil {
		return "", err
	}
	val, err := resolveTargetFieldValue(refResource, tc)
	if err != nil {
		return "", fmt.Errorf("error resolving target field value for referenced resource %v with GroupVersionKind %v: %w",
			k8s.GetNamespacedName(refResource), refResource.GroupVersionKind(), err)
	}
	return val, nil
}

func getMultiTypeReferencedResource(resourceRefValRaw map[string]interface{}, tcs []corekccv1alpha1.TypeConfig,
	ns string, kubeClient client.Client) (*corekccv1alpha1.TypeConfig, *k8s.Resource, error) {

	if len(tcs) == 0 {
		return nil, nil, fmt.Errorf("error resolving resource reference, no resource type information found")
	}

	if rawVal, ok := resourceRefValRaw["kind"]; ok {
		// "kind" is specified in rawVal
		kind, ok := rawVal.(string)
		if !ok {
			return nil, nil, fmt.Errorf("expected the value of 'kind' in the resource reference object to be string, but was actually %T", rawVal)
		}
		if len(tcs) == 1 {
			// "single-kind" resource ref should not have "kind" in rawVal
			return nil, nil, fmt.Errorf("'kind' is found in the single-type resource reference")
		}

		// "multi-kind" resource ref looks for matching "kind" in tcs
		for i := range tcs {
			tc := &tcs[i]
			if kind == tc.GVK.Kind {
				refResource, err := getReferencedResource(resourceRefValRaw, tc, ns, kubeClient)
				return tc, refResource, err
			}
		}
		return nil, nil, fmt.Errorf("the value of 'kind': '%v' is not supported in the resource reference", kind)
	}

	// "kind" is not specified in rawVal
	if len(tcs) == 1 {
		// "single-kind" resource ref uses default kind in tcs[0]
		tc := &tcs[0]
		refResource, err := getReferencedResource(resourceRefValRaw, tc, ns, kubeClient)
		return tc, refResource, err
	}

	// "multi-kind" resource ref requires "kind"
	return nil, nil, fmt.Errorf("'kind' is missing in the multi-type resource reference")
}

func getReferencedResource(resourceRefValRaw map[string]interface{}, tc *corekccv1alpha1.TypeConfig,
	ns string, kubeClient client.Client) (*k8s.Resource, error) {
	resourceRef := &v1alpha1.ResourceReference{}
	if err := util.Marshal(resourceRefValRaw, resourceRef); err != nil {
		return nil, fmt.Errorf("error marshalling raw resource reference object to resource reference struct: %w", err)
	}
	refResource, err := k8s.GetReferencedResourceIfReady(resourceRef, tc.GVK, ns, kubeClient)
	if err != nil {
		return nil, err
	}
	return refResource, nil
}

func resolveTargetFieldValue(refResource *k8s.Resource, typeConfig *corekccv1alpha1.TypeConfig) (string, error) {
	if typeConfig.TargetField == "name" {
		// When resolving target field from direct resources, get the value(resourceID) from externalRef
		if supportedgvks.IsDirectByGVK(refResource.GroupVersionKind()) || k8s.IsDirectByAnnotation(refResource) {
			val, _, err := unstructured.NestedString(refResource.Status, "externalRef")
			if err != nil {
				return "", err
			}
			tokens := strings.Split(val, "/")
			return tokens[len(tokens)-1], nil
		}
		val, ok, err := unstructured.NestedString(refResource.Spec, k8s.ResourceIDFieldName)
		if err != nil {
			return "", err
		}
		if !ok {
			return "", fmt.Errorf("couldn't resolve the resource Id")
		}
		return val, nil
	}
	if val, exist, _ := unstructured.NestedString(refResource.Status, strings.Split(typeConfig.TargetField, ".")...); exist {
		return val, nil
	}
	if val, exist, _ := unstructured.NestedString(refResource.Spec, strings.Split(typeConfig.TargetField, ".")...); exist {
		return val, nil
	}
	return "", fmt.Errorf("couldn't resolve the value for target field %v from the referenced resource %v", typeConfig.TargetField, refResource.GetNamespacedName())
}

// ResolveSpecAndStatus returns the resolved spec and status in different formats
// gated by the 'state-into-spec' annotation.
//
// If the annotation takes the 'merge' value, the function returns the spec as a mix of k8s user managed fields and defaulted state from APIs
// and returns the status with the legacy format containing observed state for output-only fields only.
//
// If the annotation takes the 'absent' value, the function will delegate to resolveDesiredStateInSpecAndObservedStateInStatus() to resolve
// the spec and the status.
func ResolveSpecAndStatus(state *unstructured.Unstructured, resource *dcl.Resource,
	smLoader dclmetadata.ServiceMetadataLoader) (spec map[string]interface{}, status map[string]interface{}, err error) {
	val, found := k8s.GetAnnotation(k8s.StateIntoSpecAnnotation, resource)
	if !found || val == k8s.StateMergeIntoSpec {
		spec, status, err = resolveMixedSpecAndLegacyStatus(state, resource, smLoader)
	} else {
		spec, status, err = resolveDesiredStateInSpecAndObservedStateInStatus(state, resource, smLoader)
	}
	if err != nil {
		return nil, nil, err
	}

	// marshal via JSON in order to ensure consistency with dcl.Resource
	normalizedSpec := make(map[string]interface{})
	normalizedStatus := make(map[string]interface{})
	if err := util.Marshal(&spec, &normalizedSpec); err != nil {
		return nil, nil, fmt.Errorf("error normalizing the spec: %w", err)
	}
	if err := util.Marshal(&status, &normalizedStatus); err != nil {
		return nil, nil, fmt.Errorf("error normalizing the status: %w", err)
	}
	return normalizedSpec, normalizedStatus, nil
}

// resolveMixedSpecAndLegacyStatus returns spec as a mix of k8s user managed fields and defaulted state from APIs
// and returns status with the legacy format containing observed state for output-only fields only.
func resolveMixedSpecAndLegacyStatus(state *unstructured.Unstructured, resource *dcl.Resource,
	smLoader dclmetadata.ServiceMetadataLoader) (spec map[string]interface{}, status map[string]interface{}, err error) {
	status, found, err := unstructured.NestedMap(state.Object, "status")
	if err != nil {
		return nil, nil, fmt.Errorf("error getting status from the state: %w", err)
	}
	if !found {
		status = make(map[string]interface{})
	}
	conditions, found, err := unstructured.NestedFieldCopy(resource.Status, "conditions")
	if err != nil {
		return nil, nil, fmt.Errorf("error resolving conditions from resource status: %w", err)
	}
	if found {
		status["conditions"] = conditions
	}
	// preserve the observedGeneration value
	g, found, err := unstructured.NestedFieldCopy(resource.Status, "observedGeneration")
	if err != nil {
		return nil, nil, fmt.Errorf("error resolving observedGeneration from resource status: %w", err)
	}
	if found {
		status["observedGeneration"] = g
	}

	stateSpec, found, err := unstructured.NestedMap(state.Object, "spec")
	if err != nil {
		return nil, nil, fmt.Errorf("error getting spec from the state: %w", err)
	}
	if !found {
		stateSpec = make(map[string]interface{})
	}
	mergedSpec, err := mergeSpecWithLiteState(stateSpec, resource.Spec, []string{}, resource.Schema, resource.ManagedFields, smLoader)
	if err != nil {
		return nil, nil, fmt.Errorf("error merging spec from the live state and the raw spec: %w", err)
	}

	if err := populateResourceIDFieldInSpec(state, resource, mergedSpec); err != nil {
		return nil, nil, fmt.Errorf("error populating 'resourceID' field in spec: %w", err)
	}

	// return nil rather than empty maps to simplify the resource representation in etcd
	if len(mergedSpec) == 0 {
		mergedSpec = nil
	}
	if len(status) == 0 {
		status = nil
	}
	return mergedSpec, status, nil
}

func populateResourceIDFieldInSpec(state *unstructured.Unstructured, resource *dcl.Resource, spec map[string]interface{}) error {
	// preserve the resourceID if specified in spec
	if val, ok := resource.Spec[k8s.ResourceIDFieldName]; ok {
		spec[k8s.ResourceIDFieldName] = val
		return nil
	}
	// If it's the case that the resource with server-generated id is initially created,
	// read the resource ID from state and store it
	val, found, err := unstructured.NestedString(state.Object, "spec", k8s.ResourceIDFieldName)
	if err != nil {
		return err
	}
	if found {
		spec[k8s.ResourceIDFieldName] = val
	}
	return nil
}

// resolveSpecAndObservedStateInStatus resolves spec as desired state and persists observed state in status.
// TODO(b/193928224): persist the full observed state including both configurable fields and output-only fields in status.
func resolveDesiredStateInSpecAndObservedStateInStatus(state *unstructured.Unstructured, resource *dcl.Resource,
	smLoader dclmetadata.ServiceMetadataLoader) (
	spec map[string]interface{}, status map[string]interface{}, err error) {
	if resource.Spec != nil {
		spec = deepcopy.MapStringInterface(resource.Spec)
	}
	if err := populateResourceIDFieldInSpec(state, resource, spec); err != nil {
		return nil, nil, fmt.Errorf("error populating 'resourceID' field in spec: %w", err)
	}
	_, status, err = resolveMixedSpecAndLegacyStatus(state, resource, smLoader)
	if err != nil {
		return nil, nil, err
	}
	return spec, status, nil
}

func mergeSpecWithLiteState(state map[string]interface{}, spec map[string]interface{}, path []string,
	schema *openapi.Schema, managedFields *fieldpath.Set, smLoader dclmetadata.ServiceMetadataLoader) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	for f, s := range schema.Properties {
		if dclextension.IsReferenceField(s) {
			refField, val, err := mergeReferenceField(state, spec, append(path, f), s, smLoader)
			if err != nil {
				return nil, err
			}
			if val != nil {
				res[refField] = val
			}
			continue
		}

		stateVal := state[f]
		specVal := spec[f]
		if stateVal == nil && specVal == nil {
			continue
		}
		// for non-returnable values, use the last captured or user specified value
		if stateVal == nil {
			res[f] = deepcopy.DeepCopy(specVal)
			continue
		}
		if specVal == nil {
			res[f] = deepcopy.DeepCopy(stateVal)
			continue
		}

		switch s.Type {
		case "object":
			if s.AdditionalProperties != nil {
				if typeutil.IsPrimitiveType(s.AdditionalProperties.Type) {
					val, err := mergePrimitiveMap(state, spec, append(path, f), managedFields)
					if err != nil {
						return nil, err
					}
					dcl.AddToMap(f, val, res)
					continue
				}
				if s.AdditionalProperties.Type == "object" {
					val, err := mergeObjectMap(state, spec, append(path, f), s.AdditionalProperties, managedFields, smLoader)
					if err != nil {
						return nil, err
					}
					dcl.AddToMap(f, val, res)
					continue
				}
				return nil, fmt.Errorf("unsupported AdditionalProperties.Type for field '%v': %v", f, s.AdditionalProperties.Type)
			}

			val, err := mergeNestedObject(state, spec, append(path, f), s, managedFields, smLoader)
			if err != nil {
				return nil, err
			}
			if val != nil {
				res[f] = val
			}
		case "array":
			if typeutil.IsPrimitiveType(s.Items.Type) {
				listVal, ok := stateVal.([]interface{})
				if !ok {
					return nil, fmt.Errorf("expected the value for field '%v' to be []interface{} but was actually %T", f, stateVal)
				}
				if len(listVal) != 0 {
					res[f] = deepcopy.DeepCopy(listVal)
				}
				continue
			}
			if s.Items.Type == "object" {
				retObjList, err := mergeObjectArray(stateVal, specVal, append(path, f), s, smLoader)
				if err != nil {
					return nil, err
				}
				if len(retObjList) == 0 {
					continue
				}
				res[f] = retObjList
				continue
			}
			return nil, fmt.Errorf("unsupported Items.Type for the array field '%v': %v", f, s.Items.Type)
		case "string":
			if k8s.IsK8sManaged(f, spec, managedFields) {
				res[f] = specVal
			} else {
				isSensitiveField, err := dclextension.IsSensitiveField(s)
				if err != nil {
					return nil, err
				}
				if isSensitiveField {
					// assert the stateVal is in the right format by marshalling into SensitiveField struct
					sensitiveVal := corekccv1alpha1.SensitiveField{}
					if err := util.Marshal(stateVal, &sensitiveVal); err != nil {
						return nil, err
					}
					res[f] = stateVal
				} else {
					res[f] = stateVal
				}
			}
		case "boolean", "number", "integer":
			if k8s.IsK8sManaged(f, spec, managedFields) {
				res[f] = specVal
			} else {
				res[f] = stateVal
			}
		default:
			return nil, fmt.Errorf("unknown schema type %v", schema.Type)
		}
	}
	return res, nil
}

func mergeReferenceField(state map[string]interface{}, spec map[string]interface{}, path []string,
	s *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader) (string, interface{}, error) {
	if dcl.IsMultiTypeParentReferenceField(path) {
		return mergeMultiTypeParentReferenceField(state, spec, s, smLoader)
	}
	return mergeResourceReference(state, spec, path, s)
}

func mergeMultiTypeParentReferenceField(state map[string]interface{}, spec map[string]interface{},
	s *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader) (string, interface{}, error) {
	// See if the user already specified a value for one of the hierarchical
	// references supported by the resource.
	specVal, tc, err := dcl.GetHierarchicalRefFromConfigForMultiParentResource(spec, s, smLoader)
	if err != nil {
		return "", nil, fmt.Errorf("error getting hierarchical reference from spec for multi-parent resource: %w", err)
	}
	if specVal != nil {
		return tc.Key, specVal, nil
	}

	// See if one of the hierarchical references was set in the state.
	stateVal, tc, err := dcl.GetHierarchicalRefFromConfigForMultiParentResource(state, s, smLoader)
	if err != nil {
		return "", nil, fmt.Errorf("error getting hierarchical reference from state for multi-parent resource: %w", err)
	}
	if stateVal != nil {
		return tc.Key, stateVal, nil
	}

	return "", nil, nil
}

func mergeResourceReference(state map[string]interface{}, spec map[string]interface{}, path []string, s *openapi.Schema) (string, interface{}, error) {
	refField, err := dclextension.GetReferenceFieldName(path, s)
	if err != nil {
		return "", nil, err
	}
	if specVal, ok := spec[refField]; ok {
		// The user already specified a value for the KCC reference field in
		// the previous spec. Preserve it.
		return refField, specVal, nil
	} else if stateVal, ok := state[refField]; ok && stateVal != nil {
		return refField, stateVal, nil
	}
	return "", nil, nil
}

func mergeObjectArray(stateVal, specVal interface{}, path []string, s *openapi.Schema, smLoader dclmetadata.ServiceMetadataLoader) ([]interface{}, error) {
	field := pathslice.Base(path)
	// DCL will return items in the original order.
	retObjList := make([]interface{}, 0)
	specList, ok := specVal.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the spec value for field '%v' to be []interface{} but was actually %T", field, specVal)
	}
	stateList, ok := stateVal.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the state value for field '%v' to be []interface{} but was actually %T", field, stateList)
	}
	if len(specList) > len(stateList) {
		return nil, fmt.Errorf("there are fewer items for field '%v' returned in state than configured in spec; state: %v, spec: %v", field, stateList, specList)
	}
	for idx, elem := range stateList {
		stateObjMap, ok := elem.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the item value from state for field '%v' to be map[string]interface{} but was actually %T", field, elem)
		}
		var specObjMap map[string]interface{}
		if idx < len(specList) {
			specObjMap, ok = specList[idx].(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("expected the item value from spec for field '%v' to be map[string]interface{} but was actually %T", field, elem)
			}
		}
		val, err := mergeSpecWithLiteState(stateObjMap, specObjMap, path, s.Items, nil, smLoader)
		if err != nil {
			return nil, err
		}
		if val != nil {
			retObjList = append(retObjList, val)
		}
	}
	return retObjList, nil
}

func mergeObjectMap(state map[string]interface{}, spec map[string]interface{}, path []string, s *openapi.Schema, managedFields *fieldpath.Set, smLoader dclmetadata.ServiceMetadataLoader) (map[string]interface{}, error) {
	field := pathslice.Base(path)
	var nestedManagedFields *fieldpath.Set
	if managedFields != nil {
		pe := fieldpath.PathElement{FieldName: &field}
		var found bool
		nestedManagedFields, found = managedFields.Children.Get(pe)
		if !found {
			nestedManagedFields = fieldpath.NewSet()
		}
	}
	retObjectMap := make(map[string]interface{})
	specObjectMap, ok := spec[field].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the spec value for field '%v' to be map[string]interface{} but was actually %T", field, spec[field])
	}
	stateObjectMap, ok := state[field].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the state value for field '%v' to be map[string]interface{} but was actually %T", field, state[field])
	}
	if len(specObjectMap) != len(stateObjectMap) {
		return nil, fmt.Errorf("the number of items for field '%v' returned in state is not the same as configured in spec; state: %v, spec: %v", field, stateObjectMap, specObjectMap)
	}

	for k := range stateObjectMap {
		if _, ok := specObjectMap[k]; !ok {
			return nil, fmt.Errorf("key '%v' is not configured in spec for field '%v'", k, field)
		}
		mergedVal, err := mergeNestedObject(stateObjectMap, specObjectMap, append(path, k), s, nestedManagedFields, smLoader)
		if err != nil {
			return nil, err
		}
		retObjectMap[k] = mergedVal
	}
	return retObjectMap, nil
}

func mergePrimitiveMap(state map[string]interface{}, spec map[string]interface{}, path []string, managedFields *fieldpath.Set) (interface{}, error) {
	field := pathslice.Base(path)
	if k8s.IsK8sManaged(field, spec, managedFields) {
		return deepcopy.DeepCopy(spec[field]), nil
	}
	if state[field] != nil {
		valMap, ok := state[field].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected the value for field '%v' to be map[string]interface{} but was actually %T", field, state[field])
		}
		if len(valMap) != 0 {
			return deepcopy.DeepCopy(valMap), nil
		}
	}
	return nil, nil
}

func mergeNestedObject(state map[string]interface{}, spec map[string]interface{}, path []string, s *openapi.Schema,
	managedFields *fieldpath.Set, smLoader dclmetadata.ServiceMetadataLoader) (map[string]interface{}, error) {
	field := pathslice.Base(path)
	var nestedManagedFields *fieldpath.Set
	if managedFields != nil {
		pe := fieldpath.PathElement{FieldName: &field}
		var found bool
		nestedManagedFields, found = managedFields.Children.Get(pe)
		if !found {
			nestedManagedFields = fieldpath.NewSet()
		}
	}
	stateConfigMap, ok := state[field].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the state value for field '%v' to be map[string]interface{} but was actually %T", field, state[field])
	}
	specConfigMap, ok := spec[field].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected the spec value for field '%v' to be map[string]interface{} but was actually %T", field, spec[field])
	}
	val, err := mergeSpecWithLiteState(stateConfigMap, specConfigMap, path, s, nestedManagedFields, smLoader)
	if err != nil {
		return nil, err
	}
	return val, nil
}
