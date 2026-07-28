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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ v1beta1.ExternalRef = &VertexAIReasoningEngineRef{}

var VertexAIReasoningEngineRefGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "VertexAIReasoningEngine",
}

func init() {
	v1beta1.Register(&VertexAIReasoningEngineRef{})
}

// VertexAIReasoningEngineRef is a reference to a GCP VertexAIReasoningEngine.
type VertexAIReasoningEngineRef struct {
	// A reference to an externally managed VertexAIReasoningEngine resource. Should be in the format "projects/{{project}}/locations/{{location}}/reasoningEngines/{{reasoningEngine}}"
	External string `json:"external,omitempty"`

	// The name of a VertexAIReasoningEngine resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIReasoningEngine resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *VertexAIReasoningEngineRef) GetGVK() schema.GroupVersionKind {
	return VertexAIReasoningEngineRefGVK
}

func (r *VertexAIReasoningEngineRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{Namespace: r.Namespace, Name: r.Name}
}

func (r *VertexAIReasoningEngineRef) GetExternal() string {
	return r.External
}

func (r *VertexAIReasoningEngineRef) SetExternal(external string) {
	r.External = external
}

func (r *VertexAIReasoningEngineRef) ValidateExternal(external string) error {
	return (&VertexAIReasoningEngineIdentity{}).FromExternal(external)
}

func (r *VertexAIReasoningEngineRef) ParseExternalToIdentity() (identity.Identity, error) {
	identity := &VertexAIReasoningEngineIdentity{}
	if err := identity.FromExternal(r.External); err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *VertexAIReasoningEngineRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return v1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromVertexAIReasoningEngineSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	})
}
