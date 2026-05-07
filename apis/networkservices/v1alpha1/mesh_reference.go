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
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &NetworkServicesMeshRef{}

// NetworkServicesMeshRef defines the resource reference to NetworkServicesMesh, which "External" field
// holds the GCP identifier for the KRM object.
type NetworkServicesMeshRef struct {
	// A reference to an externally managed NetworkServicesMesh resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/meshes/{{meshID}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkServicesMesh resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkServicesMesh resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *NetworkServicesMeshRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "networkservices.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "NetworkServicesMesh",
	}
}

func (r *NetworkServicesMeshRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkServicesMeshRef) GetExternal() string {
	return r.External
}

func (r *NetworkServicesMeshRef) SetExternal(ref string) {
	r.External = ref
}

func (r *NetworkServicesMeshRef) ValidateExternal(ref string) error {
	return (&NetworkServicesMeshIdentity{}).FromExternal(ref)
}

func (r *NetworkServicesMeshRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name != "" {
		return fmt.Errorf("cannot specify both name and external on %s reference", r.GetGVK().Kind)
	}

	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		// Fallback if status.externalRef is missing.
		// Attempt to construct the external ID from spec fields.

		// 1. Get Resource ID
		resourceID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return ""
		}

		// 2. Get Location
		location, _, _ := unstructured.NestedString(u.Object, "spec", "location")
		if location == "" {
			location = "global"
		}

		// 3. Get Project
		project, _, _ := unstructured.NestedString(u.Object, "spec", "projectRef", "external")
		if project == "" {
			// Check annotation
			annotations := u.GetAnnotations()
			if val, ok := annotations["cnrm.cloud.google.com/project-id"]; ok {
				project = val
			}
		}

		// If project is still empty, we can't construct the ID.
		if project == "" {
			return ""
		}

		// The project ID might be in the format "projects/my-project".
		// We need to strip the "projects/" prefix because NetworkServicesMeshIdentity expects the raw project ID.
		if len(project) > 9 && project[:9] == "projects/" {
			project = project[9:]
		}

		return NewNetworkServicesMeshIdentity(project, location, resourceID).String()
	})
}
