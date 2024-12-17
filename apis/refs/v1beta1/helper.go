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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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

func GetLocation(u *unstructured.Unstructured) (string, error) {
	location, _, err := unstructured.NestedString(u.Object, "spec", "location")
	if err != nil {
		return "", fmt.Errorf("reading spec.location from %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
	}
	if location == "" {
		return "", fmt.Errorf("spec.location not set in %v %v/%v: %w", u.GroupVersionKind().Kind, u.GetNamespace(), u.GetName(), err)
	}
	return location, nil
}

func ResolveResourceName(ctx context.Context, reader client.Reader, key client.ObjectKey, gvk schema.GroupVersionKind) (*unstructured.Unstructured, error) {
	resource := &unstructured.Unstructured{}
	resource.SetGroupVersionKind(gvk)
	if err := reader.Get(ctx, key, resource); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(resource.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced %v %v: %w", gvk.Kind, key, err)
	}

	return resource, nil
}
