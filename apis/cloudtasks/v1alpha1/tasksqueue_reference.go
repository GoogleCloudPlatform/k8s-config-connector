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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &TasksQueueRef{}

// TasksQueueRef is a reference to a TasksQueue.
type TasksQueueRef struct {
	// A reference to an externally managed TasksQueue resource.
	// Should be in the format "projects/{{project}}/locations/{{location}}/queues/{{queue}}".
	External string `json:"external,omitempty"`

	// The name of a TasksQueue resource.
	Name string `json:"name,omitempty"`

	// The namespace of a TasksQueue resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&TasksQueueRef{}, &TasksQueue{})
}

func (r *TasksQueueRef) GetGVK() schema.GroupVersionKind {
	return TasksQueueGVK.GroupVersion().WithKind(TasksQueueGVK.Kind)
}

func (r *TasksQueueRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}

func (r *TasksQueueRef) GetExternal() string {
	return r.External
}

func (r *TasksQueueRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

func (r *TasksQueueRef) ValidateExternal(ref string) error {
	id := &TasksQueueIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *TasksQueueRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &TasksQueueIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *TasksQueueRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromTasksQueueSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	})
}
