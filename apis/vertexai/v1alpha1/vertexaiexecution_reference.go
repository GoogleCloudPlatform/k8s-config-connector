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

var _ v1beta1.ExternalRef = &VertexAIExecutionRef{}

var VertexAIExecutionRefGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "VertexAIExecution",
}

func init() {
	v1beta1.Register(&VertexAIExecutionRef{})
}

// VertexAIExecutionRef is a reference to a GCP VertexAIExecution.
type VertexAIExecutionRef struct {
	// A reference to an externally managed VertexAIExecution resource. Should be in the format "projects/{{project}}/locations/{{location}}/metadataStores/{{metadataStore}}/executions/{{execution}}"
	External string `json:"external,omitempty"`

	// The name of a VertexAIExecution resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAIExecution resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *VertexAIExecutionRef) GetGVK() schema.GroupVersionKind {
	return VertexAIExecutionRefGVK
}

func (r *VertexAIExecutionRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{Namespace: r.Namespace, Name: r.Name}
}

func (r *VertexAIExecutionRef) GetExternal() string {
	return r.External
}

func (r *VertexAIExecutionRef) SetExternal(external string) {
	r.External = external
}

func (r *VertexAIExecutionRef) ValidateExternal(external string) error {
	return (&VertexAIExecutionIdentity{}).FromExternal(external)
}

func (r *VertexAIExecutionRef) ParseExternalToIdentity() (identity.Identity, error) {
	identity := &VertexAIExecutionIdentity{}
	if err := identity.FromExternal(r.External); err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *VertexAIExecutionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return v1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromVertexAIExecutionSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	})
}
