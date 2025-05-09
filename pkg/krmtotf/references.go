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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetReferencedResource(r *Resource, typeConfig corekccv1alpha1.TypeConfig,
	resourceRef *v1alpha1.ResourceReference, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (rsrc *Resource, err error) {
	if resourceRef.External != "" {
		return nil, fmt.Errorf("reference is external: %v", resourceRef.External)
	}
	u, err := k8s.GetReferencedResourceAsUnstruct(resourceRef, typeConfig.GVK, r.GetNamespace(), kubeClient)
	if err != nil {
		return nil, err
	}
	rsrc = &Resource{}
	if err := util.Marshal(u, rsrc); err != nil {
		return nil, fmt.Errorf("error parsing %v", u.GetName())
	}
	if typeConfig.DCLBasedResource {
		return rsrc, nil
	}

	rc, err := smLoader.GetResourceConfig(u)
	if err != nil {
		return nil, fmt.Errorf("error getting ResourceConfig for referenced resource %v: %w", r.GetName(), err)
	}
	rsrc.ResourceConfig = *rc
	return rsrc, nil
}

func handleResourceReference(config map[string]interface{}, refConfig v1alpha1.ReferenceConfig, r *Resource, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) error {
	path := strings.Split(refConfig.TFField, ".")
	return ResolveResourceReference(path, config, refConfig, r, c, smLoader)
}

func ResolveResourceReference(path []string, obj interface{}, refConfig v1alpha1.ReferenceConfig,
	r *Resource, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) error {
	if obj == nil {
		return nil
	}

	// If the object is a list, resolve the reference for each list item
	var config map[string]interface{}
	switch objAsType := obj.(type) {
	case []interface{}:
		for _, item := range objAsType {
			if err := ResolveResourceReference(path, item, refConfig, r, kubeClient, smLoader); err != nil {
				return err
			}
		}
		return nil
	case map[string]interface{}:
		config = objAsType
	default:
		return fmt.Errorf("error resolving reference: cannot iterate through type that is not object or list of objects")
	}

	field := text.SnakeCaseToLowerCamelCase(path[0])

	// Iterate down the chain of fields
	if len(path) > 1 {
		return ResolveResourceReference(path[1:], config[field], refConfig, r, kubeClient, smLoader)
	}

	// Base case. We have found the object that holds the field that is the reference. Resolve its value.
	key := field
	if refConfig.Key != "" {
		key = refConfig.Key
	}
	ref := config[key]
	if ref == nil {
		return nil
	}

	var resolvedVal interface{}
	switch refAsType := ref.(type) {
	case map[string]interface{}:
		var err error
		resolvedVal, err = ResolveReferenceObject(refAsType, refConfig, r, kubeClient, smLoader)
		if err != nil {
			return err
		}
	case []interface{}:
		resolvedList := make([]interface{}, 0)
		for _, item := range refAsType {
			itemAsMap, ok := item.(map[string]interface{})
			if !ok {
				return fmt.Errorf("expected reference %v to be object but was not", key)
			}
			resolvedVal, err := ResolveReferenceObject(itemAsMap, refConfig, r, kubeClient, smLoader)
			if err != nil {
				return err
			}
			resolvedList = append(resolvedList, resolvedVal)
		}
		resolvedVal = resolvedList
	default:
		return fmt.Errorf("unexpected type for reference field %v", path[0])
	}
	config[field] = resolvedVal
	if field != key {
		delete(config, key)
	}
	return nil
}

func ResolveReferenceObject(resourceRefValRaw map[string]interface{},
	refConfig corekccv1alpha1.ReferenceConfig, r *Resource, kubeClient client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (interface{}, error) {
	typeConfig := refConfig.TypeConfig
	if len(refConfig.Types) > 0 {
		var (
			nestedRefValRaw interface{}
			err             error
			ok              bool
			found           bool
		)
		for _, typeConfig = range refConfig.Types {
			nestedRefValRaw, found, err = unstructured.NestedFieldNoCopy(resourceRefValRaw, typeConfig.Key)
			if err != nil {
				return nil, err
			}
			if found {
				if typeConfig.JSONSchemaType != "" {
					// This is not actually a reference, but an explicit value that should be used.
					return resolveValueTemplateFromInterface(typeConfig.ValueTemplate, nestedRefValRaw, r, kubeClient, smLoader)
				}
				resourceRefValRaw, ok = nestedRefValRaw.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expected reference to be object")
				}
				break
			}
		}
		if !found {
			return nil, nil
		}
	}
	resourceRef := &v1alpha1.ResourceReference{}
	if err := util.Marshal(resourceRefValRaw, resourceRef); err != nil {
		return nil, fmt.Errorf("field %v is a wrong format", typeConfig.Key)
	}

	// Resource references usually point to K8s resources except when the
	// resource reference is an external resource reference. In the case of an
	// external resource reference, the 'external' field is used to specify a
	// string identifier for the referenced resource.
	if resourceRef.External != "" {
		return resourceRef.External, nil
	}

	// Deletions do some limited config expansion in order to preset immutable fields before the read
	// from the underlying API. Do a best-effort setting of these fields, as any unresolvable references
	// (due to, say, the reference having been deleted before) will be learned as a result of the read.
	deleting := k8s.IsDeleted(&r.ObjectMeta)

	refResource, err := GetReferencedResource(r, typeConfig, resourceRef, kubeClient, smLoader)
	if err != nil {
		if k8s.IsReferenceNotFoundError(err) {
			if deleting {
				return nil, nil
			}
			return nil, err
		}
		return nil, fmt.Errorf("error getting referenced resource from API server: %w", err)
	}

	if !deleting && !k8s.IsResourceReady(&refResource.Resource) {
		return nil, k8s.NewReferenceNotReadyErrorForResource(&refResource.Resource)
	}

	resolvedVal, err := resolveTargetFieldValue(refResource, typeConfig)
	if err != nil {
		return nil, fmt.Errorf("error resolving value of target field of "+
			"referenced resource %v %v: %v", refResource.GroupVersionKind(),
			refResource.GetNamespacedName(), err)
	}

	if deleting && typeConfig.TargetField == "" && typeConfig.ValueTemplate == "" {
		return resolvedVal, nil
	}

	return resolveValueTemplateFromInterface(typeConfig.ValueTemplate, resolvedVal, refResource, kubeClient, smLoader)
}

func resolveTargetFieldValue(r *Resource, tc corekccv1alpha1.TypeConfig) (interface{}, error) {
	key := text.SnakeCaseToLowerCamelCase(tc.TargetField)
	switch key {
	case "":
		return resolveDefaultTargetFieldValue(r, tc)
	default:
		if val, exists, _ := unstructured.NestedString(r.Spec, strings.Split(key, ".")...); exists {
			return val, nil
		}
		if val, exists, _ := unstructured.NestedString(r.GetStatusOrObservedState(), strings.Split(key, ".")...); exists {
			return val, nil
		}
		// Check 'status.observedState' for the observed, optional 'spec' fields.
		if val, exists, _ := unstructured.NestedString(getObservedStateFromStatus(r.Status), strings.Split(key, ".")...); exists {
			return val, nil
		}
		// For now, we do not support recursive target field resolution (i.e. targeting a field in
		// the referenced resource that itself is a reference to a third resource, which would require
		// its own target field resolution).
		return nil, fmt.Errorf("referenced resource's target field %v is unsupported", tc.TargetField)
	}
}

func resolveDefaultTargetFieldValue(r *Resource, tc corekccv1alpha1.TypeConfig) (interface{}, error) {
	// When resolving default target field from direct resources, get the value(resourceID) from externalRef
	if supportedgvks.IsDirectByGVK(r.Resource.GroupVersionKind()) || k8s.IsDirectByAnnotation(&r.Resource) {
		val, _, err := unstructured.NestedString(r.Status, "externalRef")
		if err != nil {
			return "", err
		}
		tokens := strings.Split(val, "/")
		return tokens[len(tokens)-1], nil
	}
	if !tc.DCLBasedResource && !SupportsResourceIDField(&r.ResourceConfig) {
		return r.GetName(), nil
	}

	id, err := r.GetResourceID()
	if err != nil {
		return "", err
	}

	return id, nil
}

func IsReferenceField(qualifiedName string, rc *corekccv1alpha1.ResourceConfig) (bool, *corekccv1alpha1.ReferenceConfig) {
	for _, refConfig := range rc.ResourceReferences {
		if qualifiedName == refConfig.TFField {
			return true, &refConfig
		}
	}
	return false, nil
}

func containsReferenceField(qualifiedName string, rc *corekccv1alpha1.ResourceConfig) bool {
	for _, refConfig := range rc.ResourceReferences {
		if strings.HasPrefix(refConfig.TFField, qualifiedName) {
			return true
		}
	}
	return false
}

func GetKeyForReferenceField(refConfig *corekccv1alpha1.ReferenceConfig) string {
	if refConfig.Key != "" {
		return refConfig.Key
	}
	parts := strings.Split(refConfig.TFField, ".")
	containerField := text.SnakeCaseToLowerCamelCase(parts[len(parts)-1])
	return containerField
}

func getPathToReferenceKey(refConfig *corekccv1alpha1.ReferenceConfig) []string {
	fieldCamelCase := text.SnakeCaseToLowerCamelCase(refConfig.TFField)
	path := strings.Split(fieldCamelCase, ".")
	if refConfig.Key != "" {
		path[len(path)-1] = refConfig.Key
	}
	return path
}

func IsHierarchicalReference(ref corekccv1alpha1.ReferenceConfig, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) bool {
	// Hierarchical references can only be at the root level, but this
	// reference is not a root-level field.
	if strings.Contains(ref.TFField, ".") {
		return false
	}
	key := GetKeyForReferenceField(&ref)
	for _, h := range hierarchicalRefs {
		if h.Key == key {
			return true
		}
	}
	return false
}

func IsRequiredParentReference(ref corekccv1alpha1.ReferenceConfig, resource *Resource) bool {
	if ref.Parent {
		return true
	}
	if !IsHierarchicalReference(ref, resource.ResourceConfig.HierarchicalReferences) {
		return false
	}
	// For projects and folders, we shouldn't treat their hierarchical references as parent references
	// because their URLs only contain their own project_id or folder_id, i.e. folders/{folder_id} and projects/{project_id}
	if resource.Kind == "Project" || resource.Kind == "Folder" {
		return false
	}
	return true
}

func GetReferenceConfigForHierarchicalReference(hierarchicalRef corekccv1alpha1.HierarchicalReference, rc *corekccv1alpha1.ResourceConfig) (*corekccv1alpha1.ReferenceConfig, error) {
	for _, ref := range rc.ResourceReferences {
		// Hierarchical references can only be at the root level, but this
		// reference is not a root-level field.
		if strings.Contains(ref.TFField, ".") {
			continue
		}
		key := GetKeyForReferenceField(&ref)
		if key == hierarchicalRef.Key {
			return &ref, nil
		}
	}
	return nil, fmt.Errorf("no reference config found for hierarchical reference field %v", hierarchicalRef.Key)
}
