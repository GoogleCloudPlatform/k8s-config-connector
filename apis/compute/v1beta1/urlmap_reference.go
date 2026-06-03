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

	// The name of a ComputeURLMap resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeURLMap resource.
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

func (r *ComputeURLMapRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.GetExternal() != "" {
		return r.ValidateExternal(r.GetExternal())
	}
	key := r.GetNamespacedName()
	if key.Namespace == "" {
		key.Namespace = defaultNamespace
	}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.GetGVK())
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return fmt.Errorf("reading referenced %s %s: %w", r.GetGVK(), key, err)
	}

	// Get external from status.externalRef. This is the most trustworthy place.
	externalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return fmt.Errorf("reading status.externalRef: %w", err)
	}
	if externalRef == "" {
		if externalRef, err = urlMapLegacyExternalRef(ctx, reader, u); err != nil {
			return err
		}
	}
	if externalRef == "" {
		return k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.SetExternal(externalRef)
	return nil
}

func urlMapLegacyExternalRef(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (string, error) {
	selfLink, found, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil {
		return "", fmt.Errorf("reading status.selfLink: %w", err)
	}
	if !found || selfLink == "" {
		return "", nil
	}
	return selfLink, nil
}
