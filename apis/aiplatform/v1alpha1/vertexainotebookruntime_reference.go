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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ColabRuntimeTemplateGVK = schema.GroupVersionKind{
	Group:   "colab.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "ColabRuntimeTemplate",
}

var _ refs.Ref = &VertexAINotebookRuntimeRef{}
var _ refs.ExternalNormalizer = &NotebookRuntimeTemplateRef{}

// VertexAINotebookRuntimeRef is a reference to a VertexAINotebookRuntime.
type VertexAINotebookRuntimeRef struct {
	// A reference to an externally managed VertexAINotebookRuntime resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/notebookRuntimes/{{notebookRuntimeID}}".
	External string `json:"external,omitempty"`

	// The name of a VertexAINotebookRuntime resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAINotebookRuntime resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAINotebookRuntimeRef{}, &VertexAINotebookRuntime{})
}

func (r *VertexAINotebookRuntimeRef) GetGVK() schema.GroupVersionKind {
	return VertexAINotebookRuntimeGVK
}

func (r *VertexAINotebookRuntimeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *VertexAINotebookRuntimeRef) GetExternal() string {
	return r.External
}

func (r *VertexAINotebookRuntimeRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *VertexAINotebookRuntimeRef) ValidateExternal(ref string) error {
	id := &VertexAINotebookRuntimeIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *VertexAINotebookRuntimeRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAINotebookRuntimeIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *VertexAINotebookRuntimeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		structuredObj, err := common.ToStructuredType[*VertexAINotebookRuntime](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromVertexAINotebookRuntimeSpec(ctx, reader, structuredObj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// NotebookRuntimeTemplateRef is a reference to a ColabRuntimeTemplate.
type NotebookRuntimeTemplateRef struct {
	// A reference to an externally managed ColabRuntimeTemplate resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/notebookRuntimeTemplates/{{notebookruntimetemplateID}}".
	External string `json:"external,omitempty"`

	// The name of a ColabRuntimeTemplate resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ColabRuntimeTemplate resource.
	Namespace string `json:"namespace,omitempty"`
}

// NormalizedExternal provision the "External" value for other resource that depends on ColabRuntimeTemplate.
// If the "External" is given in the other resource's spec.ColabRuntimeTemplateRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ColabRuntimeTemplate object from the cluster.
func (r *NotebookRuntimeTemplateRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", ColabRuntimeTemplateGVK.Kind)
	}
	// From given External
	if r.External != "" {
		tokens := strings.Split(r.External, "/")
		if len(tokens) != 6 || tokens[0] != "projects" || tokens[2] != "locations" || tokens[4] != "notebookRuntimeTemplates" {
			return "", fmt.Errorf("format of NotebookRuntimeTemplate external=%q was not known (use projects/{{projectID}}/locations/{{location}}/notebookRuntimeTemplates/{{notebookruntimetemplateID}})", r.External)
		}
		return r.External, nil
	}

	// From the Config Connector object
	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ColabRuntimeTemplateGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced %s %s: %w", ColabRuntimeTemplateGVK, key, err)
	}
	// Get external from status.externalRef. This is the most trustworthy place.
	actualExternalRef, _, err := unstructured.NestedString(u.Object, "status", "externalRef")
	if err != nil {
		return "", fmt.Errorf("reading status.externalRef: %w", err)
	}
	if actualExternalRef == "" {
		return "", k8s.NewReferenceNotReadyError(u.GroupVersionKind(), key)
	}
	r.External = actualExternalRef
	return r.External, nil
}
