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

package resolverefs

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

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

// TODO(yuhou): Location can be optional. Use provider default location when it's unset.
func GetLocation(obj *unstructured.Unstructured) (string, error) {
	// TODO(yuhou): field can be "location" or "region".
	location, _, err := unstructured.NestedString(obj.Object, "spec", "location")
	if err != nil {
		return "", fmt.Errorf("cannot get location for referenced %s %v: %w", obj.GetKind(), obj.GetNamespace(), err)
	}
	if location == "" {
		return "", fmt.Errorf("cannot get location for referenced %s %v (spec.location not set)", obj.GetKind(), obj.GetNamespace())
	}
	return location, nil
}

func GetProjectID(ctx context.Context, reader client.Reader, obj *unstructured.Unstructured) (string, error) {
	projectRefExternal, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "external")
	if projectRefExternal != "" {
		projectRef := v1beta1.ProjectRef{
			External: projectRefExternal,
		}

		project, err := ResolveProjectRef(ctx, reader, obj, &projectRef)
		if err != nil {
			return "", fmt.Errorf("cannot parse projectRef.external %q in %v %v/%v: %w", projectRefExternal, obj.GetKind(), obj.GetNamespace(), obj.GetName(), err)
		}
		return project.ProjectID, nil
	}

	projectRefName, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "name")
	if projectRefName != "" {
		projectRefNamespace, _, _ := unstructured.NestedString(obj.Object, "spec", "projectRef", "namespace")

		projectRef := v1beta1.ProjectRef{
			Name:      projectRefName,
			Namespace: projectRefNamespace,
		}
		if projectRef.Namespace == "" {
			projectRef.Namespace = obj.GetNamespace()
		}

		project, err := ResolveProjectRef(ctx, reader, obj, &projectRef)
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
