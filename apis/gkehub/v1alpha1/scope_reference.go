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

var _ refsv1beta1.Ref = &GKEHubScopeRef{}

type GKEHubScopeRef struct {
	/* The name of the scope. Allowed value: The Google Cloud resource name of a `GKEHubScope` resource (format: `projects/{{project}}/locations/{{location}}/scopes/{{name}}`).*/
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

func (r *GKEHubScopeRef) GetGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("GKEHubScope")
}

func (r *GKEHubScopeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *GKEHubScopeRef) GetExternal() string {
	return r.External
}

func (r *GKEHubScopeRef) SetExternal(ref string) {
	r.External = ref
}

func (r *GKEHubScopeRef) ValidateExternal(ref string) error {
	return (&GKEHubScopeIdentity{}).FromExternal(ref)
}

func (r *GKEHubScopeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
		if resourceID == "" {
			resourceID = u.GetName()
		}
		location, _, _ := unstructured.NestedString(u.Object, "spec", "location")
		if location == "" {
			location = "global"
		}
		project, _, _ := unstructured.NestedString(u.Object, "spec", "projectRef", "external")
		if project == "" {
			annotations := u.GetAnnotations()
			if val, ok := annotations["cnrm.cloud.google.com/project-id"]; ok {
				project = val
			}
		}
		if project == "" {
			return ""
		}
		if len(project) > 9 && project[:9] == "projects/" {
			project = project[9:]
		}
		return NewGKEHubScopeIdentity(project, location, resourceID).String()
	})
}
