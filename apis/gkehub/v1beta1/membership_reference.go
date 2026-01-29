// Copyright 2024 Google LLC
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

package v1beta1

import (
	"context"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &GKEHubMembershipRef{}

type GKEHubMembershipRef struct {
	/* The name of the membership. Allowed value: The Google Cloud resource name of a `GKEHubMembership` resource (format: `projects/{{project}}/locations/{{location}}/memberships/{{name}}`).*/
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

func (r *GKEHubMembershipRef) GetGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("GKEHubMembership")
}

func (r *GKEHubMembershipRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *GKEHubMembershipRef) GetExternal() string {
	return r.External
}

func (r *GKEHubMembershipRef) SetExternal(ref string) {
	r.External = ref
}

func (r *GKEHubMembershipRef) ValidateExternal(ref string) error {
	return (&GKEHubMembershipIdentity{}).FromExternal(ref)
}

func (r *GKEHubMembershipRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		// Fallback for DCL resources or if status.externalRef is missing.
		// Attempt to construct the external ID from spec fields.

		// 1. Get Resource ID
		resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
		if resourceID == "" {
			resourceID = u.GetName()
		}

		// 2. Get Location
		location, _, _ := unstructured.NestedString(u.Object, "spec", "location")
		if location == "" {
			location = "global" // Default to global
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

		// Handle project ID format (strip "projects/" prefix if present, as NewMembershipIdentity expects it?)
		// Actually, NewMembershipIdentity uses ParentIdentity which formats as "projects/%s".
		// If project is "projects/foo", ParentIdentity string becomes "projects/projects/foo".
		// We need to be careful. GetProjectID helper in refsv1beta1/helper.go?
		// No, I don't see GetProjectID in the list_dir output or helper.go view.
		// I must implement it inline or use a known helper.
		// "projectref.go" might have helpers but they are likely private or specific.
		// Let's implement basic stripping.

		// strip "projects/" prefix
		if len(project) > 9 && project[:9] == "projects/" {
			project = project[9:]
		}

		return NewGKEHubMembershipIdentity(project, location, resourceID).String()
	})
}
