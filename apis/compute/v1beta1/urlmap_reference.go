// Copyright 2025 Google LLC
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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputeURLMapRef{}

var ComputeURLMapGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeURLMap",
}

// A reference to a ComputeURLMap resource.
type ComputeURLMapRef struct {
	// Allowed value: The `selfLink` field of a `ComputeURLMap` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeURLMapRef) GetGVK() schema.GroupVersionKind {
	return ComputeURLMapGVK
}

func (r *ComputeURLMapRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeURLMapRef) GetExternal() string {
	return r.External
}

func (r *ComputeURLMapRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeURLMapRef) ValidateExternal(ref string) error {
	if !strings.HasPrefix(ref, "projects/") && !strings.HasPrefix(ref, "https://www.googleapis.com/") {
		return fmt.Errorf("external reference format %q is not known; expected projects/<project>/global/urlMaps/<name> or https://www.googleapis.com/compute/v1/projects/<project>/global/urlMaps/<name>", ref)
	}
	return nil
}

func (r *ComputeURLMapRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := refsv1beta1.ValidateNameAndExternal(r.Name, r.External); err != nil {
		return "", fmt.Errorf("in ComputeURLMap reference: %w", err)
	}
	if r.External != "" {
		return r.External, nil
	}

	namespace := r.Namespace
	if namespace == "" {
		namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeURLMapGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced ComputeURLMap %s: %w", key, err)
	}

	selfLink, found, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil {
		return "", fmt.Errorf("reading status.selfLink for referenced ComputeURLMap %v: %w", key, err)
	}
	if !found || selfLink == "" {
		return "", fmt.Errorf("cannot get selfLink for referenced ComputeURLMap %v (status.selfLink is empty)", key)
	}
	return selfLink, nil
}

func (r *ComputeURLMapRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		return selfLink
	})
}
