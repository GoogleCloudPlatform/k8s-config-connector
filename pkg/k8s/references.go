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
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var errContainerAnnotationNotFound = fmt.Errorf("no container annotation found")

func GetReferencedResourceIfReady(resourceRef *corekccv1alpha1.ResourceReference, gvk schema.GroupVersionKind, resourceNamespace string, kubeClient client.Client) (*Resource, error) {
	r, err := GetReferencedResource(resourceRef, gvk, resourceNamespace, kubeClient)
	if err != nil {
		return nil, err
	}
	if !IsResourceReady(r) {
		return nil, NewReferenceNotReadyErrorForResource(r)
	}
	return r, nil
}

func GetReferencedResource(resourceRef *corekccv1alpha1.ResourceReference, gvk schema.GroupVersionKind, resourceNamespace string, kubeClient client.Client) (*Resource, error) {
	u, err := GetReferencedResourceAsUnstruct(resourceRef, gvk, resourceNamespace, kubeClient)
	if err != nil {
		return nil, err
	}
	r := &Resource{}
	if err := util.Marshal(u, r); err != nil {
		return nil, fmt.Errorf("error marshalling unstruct for referenced resource %v with GroupVersionKind %v to k8s resource: %w",
			GetNamespacedName(u), u.GroupVersionKind(), err)
	}
	return r, nil
}

func GetReferencedResourceAsUnstruct(resourceRef *corekccv1alpha1.ResourceReference, gvk schema.GroupVersionKind, resourceNamespace string, kubeClient client.Client) (*unstructured.Unstructured, error) {
	name := resourceRef.Name
	if name == "" {
		return nil, fmt.Errorf("resource reference is missing required 'name' field")
	}
	namespace := resourceRef.Namespace
	if namespace == "" {
		namespace = resourceNamespace
	}
	nn := types.NamespacedName{Name: name, Namespace: namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	if err := kubeClient.Get(context.TODO(), nn, u); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, NewReferenceNotFoundError(gvk, nn)
		}
		return nil, fmt.Errorf("error getting referenced resource %v with GroupVersionKind %v from API server: %w", nn, gvk, err)
	}
	return u, nil
}

// EnsureHierarchicalReference ensures that the given resource has a
// hierarchical reference and will set one if none is found.
func EnsureHierarchicalReference(ctx context.Context, resource *Resource, hierarchicalRefs []corekccv1alpha1.HierarchicalReference, containers []corekccv1alpha1.Container, c client.Client) error {
	if len(hierarchicalRefs) == 0 {
		return nil
	}

	ns := corev1.Namespace{}
	if err := c.Get(ctx, types.NamespacedName{Name: resource.GetNamespace()}, &ns); err != nil {
		return fmt.Errorf("error getting namespace %v: %w", resource.GetNamespace(), err)
	}

	return SetDefaultHierarchicalReference(resource, &ns, hierarchicalRefs, containers)
}

// SetDefaultHierarchicalReference sets a hierarchical reference on the given
// resource if it doesn't have one. The resulting hierarchical reference is
// based on whichever of the following is found first:
// (1) Resource-level container annotations (if supported)
// (2) Namespace-level container annotations
// (3) Namespace name (if resource supports project references)
func SetDefaultHierarchicalReference(resource *Resource, ns *corev1.Namespace, hierarchicalRefs []corekccv1alpha1.HierarchicalReference, containers []corekccv1alpha1.Container) error {
	if len(hierarchicalRefs) == 0 {
		// No defaulting required.
		return nil
	}

	// If the resource already has a hierarchical reference, no modification is
	// required.
	ref, _, err := GetHierarchicalReference(resource, hierarchicalRefs)
	if err != nil {
		return fmt.Errorf("error getting hierarchical reference from object: %w", err)
	}
	if ref != nil {
		return nil
	}

	// If the resource supports the legacy behavior of resource-level
	// container annotations, use that to default a hierarchical reference.
	if len(containers) > 0 {
		annotations := resource.GetAnnotations()
		containerTypes := ContainerTypes(containers)
		err := setHierarchicalReferenceUsingContainerAnnotation(resource, annotations, hierarchicalRefs, containerTypes)
		if err != nil && !errors.Is(err, errContainerAnnotationNotFound) {
			return fmt.Errorf("error setting hierarchical reference using resource-level container annotation: %w", err)
		} else if err == nil {
			return nil
		}
	}

	// If the namespace has a container annotation, use that to default a
	// hierarchical reference.
	// Note: Use `hierarchicalRefs` to determine list of container types
	// supported by namespace instead of `containers` since new resources won't
	// have a `containers` field.
	nsAnnotations := ns.GetAnnotations()
	nsContainerTypes := ContainerTypesFor(hierarchicalRefs)
	err = setHierarchicalReferenceUsingContainerAnnotation(resource, nsAnnotations, hierarchicalRefs, nsContainerTypes)
	if err != nil && !errors.Is(err, errContainerAnnotationNotFound) {
		return fmt.Errorf("error setting hierarchical reference using namespace-level container annotation: %w", err)
	} else if err == nil {
		return nil
	}

	// For project-scoped resources, we can use the namespace name as the project ID.
	h := HierarchicalReferenceWithType(hierarchicalRefs, corekccv1alpha1.HierarchicalReferenceTypeProject)
	if h != nil {
		if err := SetHierarchicalReference(resource, h, ns.GetName()); err != nil {
			return fmt.Errorf("error setting hierarchical reference from using namespace name: %w", err)
		}
		return nil
	}

	possibleFields := HierarchicalReferencesToFields(hierarchicalRefs)
	possibleAnnotations := containerTypesToAnnotations(nsContainerTypes)
	return fmt.Errorf("resource must have one field among [%v], or namespace must have one annotation among [%v]",
		strings.Join(possibleFields, ", "), strings.Join(possibleAnnotations, ", "))
}

// GetHierarchicalReference gets the resource reference within the resource
// that corresponds to any of the given hierarchical reference configurations,
// as well as the hierarchical reference configuration associated with the
// resource reference. Returns a nil resource reference if none is found.
// Returns an error if multiple resource references are found (an invalid
// resource state as resources can have at most one hierarchical reference).
func GetHierarchicalReference(resource *Resource, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) (
	*corekccv1alpha1.ResourceReference, corekccv1alpha1.HierarchicalReference, error) {
	return GetHierarchicalReferenceFromSpec(resource.Spec, hierarchicalRefs)
}

func GetHierarchicalReferenceFromSpec(spec map[string]interface{}, hierarchicalRefs []corekccv1alpha1.HierarchicalReference) (
	*corekccv1alpha1.ResourceReference, corekccv1alpha1.HierarchicalReference, error) {
	var resourceRef *corekccv1alpha1.ResourceReference
	var hierarchicalRef corekccv1alpha1.HierarchicalReference
	for _, h := range hierarchicalRefs {
		val, ok, err := unstructured.NestedMap(spec, h.Key)
		if err != nil {
			return nil, corekccv1alpha1.HierarchicalReference{},
				fmt.Errorf("error reading spec.%v: %w", h.Key, err)
		}
		if !ok {
			continue
		}

		if resourceRef != nil {
			return nil, corekccv1alpha1.HierarchicalReference{},
				fmt.Errorf("only one of spec.%v and spec.%v can be specified", h.Key, hierarchicalRef.Key)
		}

		ref := &corekccv1alpha1.ResourceReference{}
		if err = util.Marshal(val, ref); err != nil {
			return nil, corekccv1alpha1.HierarchicalReference{},
				fmt.Errorf("error marshalling spec.%v into a resource reference: %w", h.Key, err)
		}
		resourceRef = ref
		hierarchicalRef = h
	}
	return resourceRef, hierarchicalRef, nil
}

func setHierarchicalReferenceUsingContainerAnnotation(resource *Resource, annotations map[string]string,
	hierarchicalRefs []corekccv1alpha1.HierarchicalReference, containerTypes []corekccv1alpha1.ContainerType) error {

	containerVal, containerType, err := GetContainerAnnotation(annotations, containerTypes)
	if err != nil {
		return fmt.Errorf("error getting container annotation from annotations: %w", err)
	}
	if containerVal == "" {
		return errContainerAnnotationNotFound
	}

	h := HierarchicalReferenceWithType(hierarchicalRefs, HierarchicalReferenceTypeFor(containerType))
	if h == nil {
		return fmt.Errorf("no hierarchical reference found for container type %v", containerType)
	}

	return SetHierarchicalReference(resource, h, containerVal)
}

func SetHierarchicalReference(resource *Resource, hierarchicalRef *corekccv1alpha1.HierarchicalReference, externalVal string) error {
	val := corekccv1alpha1.ResourceReference{
		External: externalVal,
	}
	var valAsMap map[string]interface{}
	if err := util.Marshal(val, &valAsMap); err != nil {
		return fmt.Errorf("error marshalling resource reference to map: %w", err)
	}
	if resource.Spec == nil {
		resource.Spec = make(map[string]interface{})
	}
	return unstructured.SetNestedMap(resource.Spec, valAsMap, hierarchicalRef.Key)
}

func HierarchicalReferenceWithType(hierarchicalRefs []corekccv1alpha1.HierarchicalReference, hType corekccv1alpha1.HierarchicalReferenceType) *corekccv1alpha1.HierarchicalReference {
	for _, h := range hierarchicalRefs {
		if h.Type == hType {
			return &h
		}
	}
	return nil
}

func HierarchicalReferencesToFields(hierarchicalRefs []corekccv1alpha1.HierarchicalReference) []string {
	fields := make([]string, 0)
	for _, h := range hierarchicalRefs {
		fields = append(fields, "spec."+h.Key)
	}
	return fields
}

func HierarchicalReferenceTypeFor(containerType corekccv1alpha1.ContainerType) corekccv1alpha1.HierarchicalReferenceType {
	switch containerType {
	case v1alpha1.ContainerTypeProject:
		return v1alpha1.HierarchicalReferenceTypeProject
	case v1alpha1.ContainerTypeFolder:
		return v1alpha1.HierarchicalReferenceTypeFolder
	case v1alpha1.ContainerTypeOrganization:
		return v1alpha1.HierarchicalReferenceTypeOrganization
	default:
		panic(fmt.Errorf("unrecognized container type: %v", containerType))
	}
}

func ContainerTypesFor(hierarchicalRefs []corekccv1alpha1.HierarchicalReference) []corekccv1alpha1.ContainerType {
	types := make([]corekccv1alpha1.ContainerType, 0)
	for _, h := range hierarchicalRefs {
		switch h.Type {
		case corekccv1alpha1.HierarchicalReferenceTypeBillingAccount:
			// Skip conversion of billing account hierarchical references to a
			// container type since there is no container type that corresponds
			// to billing accounts.
			continue
		default:
			types = append(types, containerTypeFor(h))
		}
	}
	return types
}

func containerTypeFor(hierarchicalRef corekccv1alpha1.HierarchicalReference) corekccv1alpha1.ContainerType {
	switch hierarchicalRef.Type {
	case corekccv1alpha1.HierarchicalReferenceTypeProject:
		return corekccv1alpha1.ContainerTypeProject
	case corekccv1alpha1.HierarchicalReferenceTypeFolder:
		return corekccv1alpha1.ContainerTypeFolder
	case corekccv1alpha1.HierarchicalReferenceTypeOrganization:
		return corekccv1alpha1.ContainerTypeOrganization
	case corekccv1alpha1.HierarchicalReferenceTypeBillingAccount:
		panic(fmt.Errorf("there is no container type equivalent to the hierarchical reference type %v", hierarchicalRef.Type))
	default:
		panic(fmt.Errorf("unrecognized hierarchical reference type %v", hierarchicalRef.Type))
	}
}
