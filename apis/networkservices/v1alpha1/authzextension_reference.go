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

var _ refsv1beta1.Ref = &NetworkServicesAuthzExtensionRef{}

type NetworkServicesAuthzExtensionRef struct {
	/* The name of the authzextension. Allowed value: The Google Cloud resource name of a `NetworkServicesAuthzExtension` resource (format: `projects/{{project}}/locations/{{location}}/authzExtensions/{{name}}`).*/
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

func (r *NetworkServicesAuthzExtensionRef) GetGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("NetworkServicesAuthzExtension")
}

func (r *NetworkServicesAuthzExtensionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkServicesAuthzExtensionRef) GetExternal() string {
	return r.External
}

func (r *NetworkServicesAuthzExtensionRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *NetworkServicesAuthzExtensionRef) ValidateExternal(ref string) error {
	return (&AuthzExtensionIdentity{}).FromExternal(ref)
}

func (r *NetworkServicesAuthzExtensionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name != "" {
		return fmt.Errorf("cannot specify both name and external on NetworkServicesAuthzExtensionRef")
	}
	if r.External != "" {
		if err := r.ValidateExternal(r.External); err != nil {
			return err
		}
		return nil
	}

	key := types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}

	obj := &NetworkServicesAuthzExtension{}
	if err := reader.Get(ctx, key, obj); err != nil {
		return err
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return err
	}
	r.External = id.String()
	return nil
}

func (r *NetworkServicesAuthzExtensionRef) NormalizeWithFallback(ctx context.Context, reader client.Reader, defaultNamespace string) error {
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
		// We need to strip the "projects/" prefix because AuthzExtensionIdentity expects the raw project ID.
		if len(project) > 9 && project[:9] == "projects/" {
			project = project[9:]
		}

		return fmt.Sprintf("projects/%s/locations/%s/authzExtensions/%s", project, location, resourceID)
	})
}
