/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockkubeapiserver

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
)

func (s *MemoryStorage) objectChanged(u *unstructured.Unstructured) {
	gvk := u.GroupVersionKind()

	switch gvk.GroupKind() {
	case schema.GroupKind{Kind: "Namespace"}:
		s.namespaceChanged(u)
	case schema.GroupKind{Group: "apiextensions.k8s.io", Kind: "CustomResourceDefinition"}:
		if err := s.crdChanged(u); err != nil {
			klog.Warningf("crd change was invalid: %v", err)
		}
	}
}

func (s *MemoryStorage) namespaceChanged(u *unstructured.Unstructured) {
	// These changes seem to be done synchronously (similar to a mutating webhook)
	labels := u.GetLabels()
	name := u.GetName()
	if labels["kubernetes.io/metadata.name"] != name {
		if labels == nil {
			labels = make(map[string]string)
		}
		labels["kubernetes.io/metadata.name"] = name
		u.SetLabels(labels)
	}
	phase, _, _ := unstructured.NestedFieldNoCopy(u.Object, "status", "phase")
	if phase != "Active" {
		unstructured.SetNestedField(u.Object, "Active", "status", "phase")
	}
	found := false
	finalizers, _, _ := unstructured.NestedSlice(u.Object, "spec", "finalizers")
	for _, finalizer := range finalizers {
		if finalizer == "kubernetes" {
			found = true
		}
	}
	if !found {
		finalizers = append(finalizers, "kubernetes")
		unstructured.SetNestedSlice(u.Object, finalizers, "spec", "finalizers")
	}
}

func (s *MemoryStorage) crdChanged(u *unstructured.Unstructured) error {
	// TODO: Deleted / changed CRDs

	group, _, _ := unstructured.NestedString(u.Object, "spec", "group")
	if group == "" {
		return fmt.Errorf("spec.group not set")
	}

	kind, _, _ := unstructured.NestedString(u.Object, "spec", "names", "kind")
	if kind == "" {
		return fmt.Errorf("spec.names.kind not set")
	}

	resource, _, _ := unstructured.NestedString(u.Object, "spec", "names", "plural")
	if resource == "" {
		return fmt.Errorf("spec.names.plural not set")
	}

	scope, _, _ := unstructured.NestedString(u.Object, "spec", "scope")
	if scope == "" {
		return fmt.Errorf("spec.scope not set")
	}

	versionsObj, found, _ := unstructured.NestedFieldNoCopy(u.Object, "spec", "versions")
	if !found {
		return fmt.Errorf("spec.versions not set")
	}

	versions, ok := versionsObj.([]interface{})
	if !ok {
		return fmt.Errorf("spec.versions not a slice")
	}

	for _, versionObj := range versions {
		version, ok := versionObj.(map[string]interface{})
		if !ok {
			return fmt.Errorf("spec.versions element not an object")
		}

		versionName, _, _ := unstructured.NestedString(version, "name")
		if versionName == "" {
			return fmt.Errorf("version name not set")
		}
		gvk := schema.GroupVersionKind{Group: group, Version: versionName, Kind: kind}
		gvr := gvk.GroupVersion().WithResource(resource)
		gr := gvr.GroupResource()

		storage := &resourceStorage{
			GroupResource: gr,
			objects:       make(map[types.NamespacedName]*unstructured.Unstructured),
		}

		// TODO: share storage across different versions
		s.resourceStorages[gr] = storage

		r := &ResourceInfo{
			API: metav1.APIResource{
				Name:    resource,
				Group:   gvk.Group,
				Version: gvk.Version,
				Kind:    gvk.Kind,
			},
			GVK:     gvk,
			GVR:     gvr,
			storage: storage,
		}
		r.ListGVK = gvk.GroupVersion().WithKind(gvk.Kind + "List")

		// TODO: Set r.TypeInfo from schema

		switch scope {
		case "Namespaced":
			r.API.Namespaced = true
		case "Cluster":
			r.API.Namespaced = false
		default:
			return fmt.Errorf("scope %q is not recognized", scope)
		}

		s.schema.resources = append(s.schema.resources, r)
	}
	return nil
}
