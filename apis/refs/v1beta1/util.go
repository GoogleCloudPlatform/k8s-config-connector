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
	"fmt"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ValidateResourceRef(ref *v1alpha1.ResourceRef) error {
	if ref == nil {
		return nil
	}
	if ref.Name == "" && ref.External == "" {
		return fmt.Errorf("must specify either name or external on reference")
	}
	if ref.Name != "" && ref.External != "" {
		return fmt.Errorf("cannot specify both name and external on reference")
	}
	return nil
}

func ParseResourceRef(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef, group, version, kind string) (*unstructured.Unstructured, error) {
	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	resource := &unstructured.Unstructured{}
	resource.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   group,
		Version: version,
		Kind:    kind,
	})
	if err := reader.Get(ctx, key, resource); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced %s %v not found", kind, key)
		}
		return nil, fmt.Errorf("error reading referenced %s %v: %w", kind, key, err)
	}
	return resource, nil
}

func GetResourceID(u *unstructured.Unstructured) (string, error) {
	resourceID, _, err := unstructured.NestedString(u.Object, "spec", "resourceID")
	if err != nil {
		return "", fmt.Errorf("reading spec.resourceID from %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
	}
	if resourceID == "" {
		resourceID = u.GetName()
	}
	return resourceID, nil
}

func GetProjectID(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	projectRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "external")
	if projectRefExternal != "" {
		projectRef := ProjectRef{
			External: projectRefExternal,
		}

		project, err := ResolveProject(ctx, reader, obj, &projectRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse projectRef.external %q in %v %v/%v: %w", projectRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return project.ProjectID, nil
	}

	projectRefName, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "name")
	if projectRefName != "" {
		projectRefNamespace, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "namespace")

		projectRef := ProjectRef{
			Name:      projectRefName,
			Namespace: projectRefNamespace,
		}
		if projectRef.Namespace == "" {
			projectRef.Namespace = obj.GetNamespace()
		}

		project, err := ResolveProject(ctx, reader, obj, &projectRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse projectRef in %v %v/%v: %w", obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return project.ProjectID, nil
	}

	if projectID := obj.GetAnnotations()["cnrm.cloud.google.com/project-id"]; projectID != "" {
		return projectID, nil
	}

	return "", fmt.Errorf("cannot find project id for %v %v/%v", obj.GetKind(), obj.GetNamespace(), obj.GetName())
}

// todo(yuhou): Location can be optional. Use default location when it's unset.
func GetLocation(obj *unstructured.Unstructured) (string, error) {
	location, _, err := unstructured.NestedString(obj.Object, "spec", "location")
	if err != nil {
		return "", fmt.Errorf("cannot get location for referenced %s %v: %w", obj.GetKind(), obj.GetNamespace(), err)
	}
	if location == "" {
		return "", fmt.Errorf("cannot get location for referenced %s %v (spec.location not set)", obj.GetKind(), obj.GetNamespace())
	}
	return location, nil
}
