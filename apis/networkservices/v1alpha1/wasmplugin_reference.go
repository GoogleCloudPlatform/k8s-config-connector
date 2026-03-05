// Copyright 2026 Google LLC
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

package v1alpha1

import (
	"context"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &NetworkServicesWasmPluginRef{}

type NetworkServicesWasmPluginRef struct {
	/* The name of the wasmplugin. Allowed value: The Google Cloud resource name of a `NetworkServicesWasmPlugin` resource (format: `projects/{{project}}/locations/{{location}}/wasmPlugins/{{name}}`).*/
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

func (r *NetworkServicesWasmPluginRef) GetGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("NetworkServicesWasmPlugin")
}

func (r *NetworkServicesWasmPluginRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkServicesWasmPluginRef) GetExternal() string {
	return r.External
}

func (r *NetworkServicesWasmPluginRef) SetExternal(ref string) {
	r.External = ref
}

func (r *NetworkServicesWasmPluginRef) ValidateExternal(ref string) error {
	return (&NetworkServicesWasmPluginIdentity{}).FromExternal(ref)
}

func (r *NetworkServicesWasmPluginRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		// Fallback if status.externalRef is missing.
		// Attempt to construct the external ID from spec fields.

		// 1. Get Resource ID
		resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
		if resourceID == "" {
			resourceID = u.GetName()
		}

		// 2. Get Location
		location, _, _ := unstructured.NestedString(u.Object, "spec", "location")

		// 3. Get Project
		project, _, _ := unstructured.NestedString(u.Object, "spec", "projectRef", "external")
		if project == "" {
			// Check annotation
			annotations := u.GetAnnotations()
			if val, ok := annotations["cnrm.cloud.google.com/project-id"]; ok {
				project = val
			}
		}

		// If project or location is still empty, we can't construct the ID.
		if project == "" || location == "" {
			return ""
		}

		// The project ID might be in the format "projects/my-project".
		// We need to strip the "projects/" prefix because NetworkServicesWasmPluginIdentity expects the raw project ID.
		if len(project) > 9 && project[:9] == "projects/" {
			project = project[9:]
		}

		return NewNetworkServicesWasmPluginIdentity(project, location, resourceID).String()
	})
}
