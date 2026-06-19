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

package v1beta1

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	VertexAITensorboardGVK = schema.GroupVersionKind{
		Group:   "vertexai.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "VertexAITensorboard",
	}

	VertexAIPersistentResourceGVK = schema.GroupVersionKind{
		Group:   "vertexai.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "VertexAIPersistentResource",
	}
)

func init() {
	Register(&VertexAITensorboardRef{})
	Register(&VertexAIPersistentResourceRef{})
}

type VertexAITensorboardRef struct {
	/* A reference to an externally managed Vertex AI Tensorboard resource.
	Should be of the format `projects/{{projectID}}/locations/{{location}}/tensorboards/{{tensorboardID}}`. */
	External string `json:"external,omitempty"`

	/* The `name` of a `VertexAITensorboard` resource. */
	Name string `json:"name,omitempty"`

	/* The `namespace` of a `VertexAITensorboard` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *VertexAITensorboardRef) GetGVK() schema.GroupVersionKind {
	return VertexAITensorboardGVK
}

func (r *VertexAITensorboardRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *VertexAITensorboardRef) GetExternal() string {
	return r.External
}

func (r *VertexAITensorboardRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAITensorboardRef) ValidateExternal(ref string) error {
	return nil
}

func (r *VertexAITensorboardRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	_, err := r.NormalizedExternal(ctx, reader, defaultNamespace)
	return err
}

func (r *VertexAITensorboardRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" {
		return r.External, nil
	}
	if r.Name == "" {
		return "", fmt.Errorf("must specify either name or external on VertexAITensorboardRef")
	}

	key := types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
	if key.Namespace == "" {
		key.Namespace = otherNamespace
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(VertexAITensorboardGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", fmt.Errorf("referenced VertexAITensorboard %v not found", key)
		}
		return "", fmt.Errorf("error reading referenced VertexAITensorboard %v: %w", key, err)
	}

	external, err := GetResourceID(u)
	if err != nil {
		return "", err
	}
	r.External = external
	return r.External, nil
}

type VertexAIPersistentResourceRef struct {
	/* A reference to an externally managed Vertex AI PersistentResource resource.
	Should be of the format `projects/{{projectID}}/locations/{{location}}/persistentResources/{{persistentResourceID}}`. */
	External string `json:"external,omitempty"`

	/* The `name` of a `VertexAIPersistentResource` resource. */
	Name string `json:"name,omitempty"`

	/* The `namespace` of a `VertexAIPersistentResource` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *VertexAIPersistentResourceRef) GetGVK() schema.GroupVersionKind {
	return VertexAIPersistentResourceGVK
}

func (r *VertexAIPersistentResourceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *VertexAIPersistentResourceRef) GetExternal() string {
	return r.External
}

func (r *VertexAIPersistentResourceRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAIPersistentResourceRef) ValidateExternal(ref string) error {
	return nil
}

func (r *VertexAIPersistentResourceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	_, err := r.NormalizedExternal(ctx, reader, defaultNamespace)
	return err
}

func (r *VertexAIPersistentResourceRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" {
		return r.External, nil
	}
	return "", fmt.Errorf("VertexAIPersistentResource is not supported as a local KCC resource reference, please use external")
}
