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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &VertexAITrainingPipelineRef{}

// VertexAITrainingPipelineRef is a reference to a GCP VertexAITrainingPipeline.
type VertexAITrainingPipelineRef struct {
	// A reference to an externally managed VertexAITrainingPipeline resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/trainingPipelines/{{trainingPipelineID}}"
	External string `json:"external,omitempty"`

	// The name of a VertexAITrainingPipeline resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAITrainingPipeline resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&VertexAITrainingPipelineRef{})
}

// GetGVK returns the GroupVersionKind of the resource.
func (r *VertexAITrainingPipelineRef) GetGVK() schema.GroupVersionKind {
	return VertexAITrainingPipelineGVK
}

// GetNamespacedName returns the namespaced name of the referenced resource.
func (r *VertexAITrainingPipelineRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

// GetExternal returns the external reference string.
func (r *VertexAITrainingPipelineRef) GetExternal() string {
	return r.External
}

// SetExternal sets the external reference string.
func (r *VertexAITrainingPipelineRef) SetExternal(ref string) {
	r.External = ref
}

// ValidateExternal validates the format of the external reference.
func (r *VertexAITrainingPipelineRef) ValidateExternal(ref string) error {
	id := &VertexAITrainingPipelineIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

// ParseExternalToIdentity parses the external reference into an Identity.
func (r *VertexAITrainingPipelineRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &VertexAITrainingPipelineIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

// Normalize normalizes the reference using fallback lookup.
func (r *VertexAITrainingPipelineRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		structured, err := common.ToStructuredType[*VertexAITrainingPipeline](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromVertexAITrainingPipelineSpec(ctx, reader, structured)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
