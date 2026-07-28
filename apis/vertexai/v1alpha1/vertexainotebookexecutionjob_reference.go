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

var _ v1beta1.ExternalRef = &VertexAINotebookExecutionJobRef{}

var VertexAINotebookExecutionJobRefGVK = schema.GroupVersionKind{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
	Kind:    "VertexAINotebookExecutionJob",
}

func init() {
	v1beta1.Register(&VertexAINotebookExecutionJobRef{})
}

// VertexAINotebookExecutionJobRef is a reference to a GCP VertexAINotebookExecutionJob.
type VertexAINotebookExecutionJobRef struct {
	// A reference to an externally managed VertexAINotebookExecutionJob resource. Should be in the format "projects/{{project}}/locations/{{location}}/notebookExecutionJobs/{{notebookExecutionJob}}"
	External string `json:"external,omitempty"`

	// The name of a VertexAINotebookExecutionJob resource.
	Name string `json:"name,omitempty"`

	// The namespace of a VertexAINotebookExecutionJob resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *VertexAINotebookExecutionJobRef) GetGVK() schema.GroupVersionKind {
	return VertexAINotebookExecutionJobRefGVK
}

func (r *VertexAINotebookExecutionJobRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{Namespace: r.Namespace, Name: r.Name}
}

func (r *VertexAINotebookExecutionJobRef) GetExternal() string {
	return r.External
}

func (r *VertexAINotebookExecutionJobRef) SetExternal(external string) {
	r.External = external
}

func (r *VertexAINotebookExecutionJobRef) ValidateExternal(external string) error {
	return (&VertexAINotebookExecutionJobIdentity{}).FromExternal(external)
}

func (r *VertexAINotebookExecutionJobRef) ParseExternalToIdentity() (identity.Identity, error) {
	identity := &VertexAINotebookExecutionJobIdentity{}
	if err := identity.FromExternal(r.External); err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *VertexAINotebookExecutionJobRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return v1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromVertexAINotebookExecutionJobSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	})
}
