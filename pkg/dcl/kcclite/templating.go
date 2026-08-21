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
	"path"
	"regexp"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var fieldRegex = regexp.MustCompile("{{([0-9A-Za-z]+)}}")

func CanonicalizeReferencedResourceName(name string, nameValueTemplate string, refResource *k8s.Resource,
	smLoader dclmetadata.ServiceMetadataLoader, schemaLoader dclschemaloader.DCLSchemaLoader,
	serviceMappingLoader *servicemappingloader.ServiceMappingLoader, kubeClient client.Client) (string, error) {
	ret := strings.ReplaceAll(nameValueTemplate, "{{name}}", name)
	var resolutionError error
	resolveFunc := func(s string) string {
		field := fieldRegex.FindStringSubmatch(s)[1]
		if dcl.IsParentReferenceField([]string{field}) {
			val, err := resolveParentReferenceFieldValue(field, refResource, smLoader, schemaLoader, serviceMappingLoader, kubeClient)
			if err != nil {
				resolutionError = decorateErrorForResolvingReferenceField(err, field, refResource)
				return ""
			}
			return val
		}
		val, exists, err := unstructured.NestedString(refResource.Spec, field)
		if err != nil {
			resolutionError = fmt.Errorf("error getting value for DCL field %v in spec of referenced resource %v with GroupVersionKind %v: %w",
				field, k8s.GetNamespacedName(refResource), refResource.GroupVersionKind(), err)
			return ""
		}
		if exists {
			return val
		}
		// Value not found in spec, so check status.
		val, exists, err = unstructured.NestedString(refResource.Status, field)
		if err != nil {
			resolutionError = fmt.Errorf("error getting value for DCL field %v in status of referenced resource %v with GroupVersionKind %v: %w",
				field, k8s.GetNamespacedName(refResource), refResource.GroupVersionKind(), err)
			return ""
		}
		if exists {
			return val
		}
		resolutionError = fmt.Errorf("no value found for DCL field %v in referenced resource %v with GroupVersionKind %v",
			field, k8s.GetNamespacedName(refResource), refResource.GroupVersionKind())
		return ""
	}
	return fieldRegex.ReplaceAllStringFunc(ret, resolveFunc), resolutionError
}

func resolveParentReferenceFieldValue(field string, resource *k8s.Resource, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, serviceMappingLoader *servicemappingloader.ServiceMappingLoader, kubeClient client.Client) (string, error) {
	if dclmetadata.IsDCLBasedResourceKind(resource.GroupVersionKind(), smLoader) {
		return resolveParentReferenceFieldValueForDCLResource(field, resource, smLoader, schemaLoader, kubeClient)
	}
	// If the resource is not a DCL-based resource, then assume it is TF-based.
	return resolveParentReferenceFieldValueForTFResource(field, resource, serviceMappingLoader, kubeClient)
}

func resolveParentReferenceFieldValueForDCLResource(field string, resource *k8s.Resource, smLoader dclmetadata.ServiceMetadataLoader,
	schemaLoader dclschemaloader.DCLSchemaLoader, kubeClient client.Client) (string, error) {
	r, found := smLoader.GetResourceWithGVK(resource.GroupVersionKind())
	if !found {
		return "", fmt.Errorf("ServiceMetadata for resource with GroupVersionKind %v not found", resource.GroupVersionKind())
	}
	if !r.Releasable {
		return "", fmt.Errorf("expected resource with GroupVersionKind %v to be supported via DCL, but it is not", resource.GroupVersionKind())
	}

	// TODO(b/186159460): Delete this if-block once all DCL-based resources
	// support hierarchical references.
	if !r.SupportsHierarchicalReferences {
		return resolveParentFieldFromContainerAnnotation(field, resource)
	}

	schema, err := dclschemaloader.GetDCLSchemaForGVK(resource.GroupVersionKind(), smLoader, schemaLoader)
	if err != nil {
		return "", fmt.Errorf("error getting DCL schema for GroupVersionKind %v: %w", resource.GroupVersionKind(), err)
	}
	fieldSchema, ok := schema.Properties[field]
	if !ok {
		return "", fmt.Errorf("could not find schema for DCL field '%v' for GroupVersionKind %v", field, resource.GroupVersionKind())
	}
	if dcl.IsMultiTypeParentReferenceField([]string{field}) {
		return resolveMultiTypeParentReferenceFieldValueForDCLResource(fieldSchema, resource, smLoader, kubeClient)
	}
	return resolveSingleTypeParentReferenceFieldValueForDCLResource(field, fieldSchema, resource, smLoader, kubeClient)
}

func resolveMultiTypeParentReferenceFieldValueForDCLResource(schema *openapi.Schema, resource *k8s.Resource,
	smLoader dclmetadata.ServiceMetadataLoader, kubeClient client.Client) (string, error) {
	rawVal, tc, err := dcl.GetHierarchicalRefFromConfigForMultiParentResource(resource.Spec, schema, smLoader)
	if err != nil {
		return "", fmt.Errorf("error getting hierarchical reference from config for multi-parent resource: %w", err)
	}
	if rawVal == nil {
		return "", fmt.Errorf("no hierarchical reference found for multi-parent resource")
	}
	refField := tc.Key
	refObj, ok := rawVal.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("expected the value to be map[string]interface{} for reference field %v but was actually %T", refField, rawVal)
	}
	val, err := resolveReferenceObject(refObj, tc, resource.GetNamespace(), kubeClient)
	if err != nil {
		return "", err
	}
	name := path.Base(val) // In case the resolved value is a path, use the last element of the path only.
	return fmt.Sprintf("%v%v", dcl.ParentPrefixForKind(tc.GVK.Kind), name), nil
}

func resolveSingleTypeParentReferenceFieldValueForDCLResource(field string, schema *openapi.Schema, resource *k8s.Resource,
	smLoader dclmetadata.ServiceMetadataLoader, kubeClient client.Client) (string, error) {
	refField, err := dclextension.GetReferenceFieldName([]string{field}, schema)
	if err != nil {
		return "", fmt.Errorf("error getting the reference field name for DCL field '%v': %w", field, err)
	}
	rawVal, ok := resource.Spec[refField]
	if !ok || rawVal == nil {
		return "", fmt.Errorf("no hierarchical reference found for single-parent resource")
	}
	refObj, ok := rawVal.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("expected the value to be map[string]interface{} for reference field %v but was actually %T", refField, rawVal)
	}
	tcs, err := dcl.GetReferenceTypeConfigs(schema, smLoader)
	if err != nil {
		return "", fmt.Errorf("error getting type configs for DCL field '%v': %w", field, err)
	}
	if len(tcs) > 1 {
		return "", fmt.Errorf("unexpectedly got more than one type config for DCL field '%v' which is supposed to be a single-type parent reference field", field)
	}
	val, err := resolveReferenceObject(refObj, &tcs[0], resource.GetNamespace(), kubeClient)
	if err != nil {
		return "", err
	}
	name := path.Base(val) // In case the resolved value is a path, use the last element of the path only.
	return name, nil
}

func resolveParentReferenceFieldValueForTFResource(field string, resource *k8s.Resource,
	serviceMappingLoader *servicemappingloader.ServiceMappingLoader, kubeClient client.Client) (string, error) {
	u, err := resource.MarshalAsUnstructured()
	if err != nil {
		return "", fmt.Errorf("error marshalling resource to unstructured: %w", err)
	}
	rc, err := serviceMappingLoader.GetResourceConfig(u)
	if err != nil {
		return "", fmt.Errorf("error getting resource config for resource: %w", err)
	}
	krmResource := &krmtotf.Resource{
		Resource:       *resource,
		ResourceConfig: *rc,
	}
	// TODO(b/193177782): Delete this if-block once all TF-based resources
	// support hierarchical references.
	if !krmtotf.SupportsHierarchicalReferences(rc) {
		return resolveParentFieldFromContainerAnnotation(field, resource)
	}
	ref, hierarchicalRef, err := k8s.GetHierarchicalReference(resource, rc.HierarchicalReferences)
	if err != nil {
		return "", fmt.Errorf("error getting hierarchical reference: %w", err)
	}
	refConfig, err := krmtotf.GetReferenceConfigForHierarchicalReference(hierarchicalRef, rc)
	if err != nil {
		return "", fmt.Errorf("error getting reference config for hierarchical reference: %w", err)
	}
	var refObj map[string]interface{}
	if err := util.Marshal(ref, &refObj); err != nil {
		return "", fmt.Errorf("error marshalling hierarchical reference to map[string]interface{}: %w", err)
	}
	rawVal, err := krmtotf.ResolveReferenceObject(refObj, *refConfig, krmResource, kubeClient, serviceMappingLoader)
	if err != nil {
		return "", fmt.Errorf("error resolving resource reference object representing resource's hierarchical reference: %w", err)
	}
	val, ok := rawVal.(string)
	if !ok {
		return "", fmt.Errorf("expected the resolved value of the resource reference object to be string, but was actually %T", rawVal)
	}
	name := path.Base(val) // In case the resolved value is a path, use the last element of the path only.
	if dcl.IsMultiTypeParentReferenceField([]string{field}) {
		return fmt.Sprintf("%v%v", dcl.ParentPrefixForKind(refConfig.GVK.Kind), name), nil
	}
	return name, nil
}

func resolveParentFieldFromContainerAnnotation(field string, resource *k8s.Resource) (string, error) {
	annotation := k8s.GetAnnotationForContainerType(corekccv1alpha1.ContainerType(field))
	val, ok := k8s.GetAnnotation(annotation, resource)
	if !ok || val == "" {
		return "", fmt.Errorf("no value found for annotation %v in resource %v",
			annotation, k8s.GetNamespacedName(resource))
	}
	return val, nil
}

func decorateErrorForResolvingReferenceField(err error, field string, refResource *k8s.Resource) error {
	if unwrappedErr, ok := k8s.AsReferenceNotFoundError(err); ok {
		return k8s.NewTransitiveDependencyNotFoundError(unwrappedErr.RefResourceGVK, unwrappedErr.RefResource)
	}
	if unwrappedErr, ok := k8s.AsReferenceNotReadyError(err); ok {
		return k8s.NewTransitiveDependencyNotReadyError(unwrappedErr.RefResourceGVK, unwrappedErr.RefResource)
	}
	return fmt.Errorf("error getting value for DCL field %v in referenced resource %v with GroupVersionKind %v: %w",
		field, k8s.GetNamespacedName(refResource), refResource.GroupVersionKind(), err)
}
