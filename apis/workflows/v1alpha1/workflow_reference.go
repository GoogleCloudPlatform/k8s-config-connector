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

package v1alpha1

import (
	"context"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &WorkflowsWorkflowRef{}

// WorkflowsWorkflowRef defines the resource reference to WorkflowsWorkflow, which "External" field
// holds the GCP identifier for the KRM object.
type WorkflowsWorkflowRef struct {
	// A reference to an externally managed WorkflowsWorkflow resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/workflows/{{workflowID}}".
	External string `json:"external,omitempty"`

	// The name of a WorkflowsWorkflow resource.
	Name string `json:"name,omitempty"`

	// The namespace of a WorkflowsWorkflow resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *WorkflowsWorkflowRef) GetGVK() schema.GroupVersionKind {
	return WorkflowsWorkflowGVK
}

func (r *WorkflowsWorkflowRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *WorkflowsWorkflowRef) GetExternal() string {
	return r.External
}

func (r *WorkflowsWorkflowRef) SetExternal(ref string) {
	r.External = ref
}

func (r *WorkflowsWorkflowRef) ValidateExternal(ref string) error {
	id := &WorkflowsWorkflowIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *WorkflowsWorkflowRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}
