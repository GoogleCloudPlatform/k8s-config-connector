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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &NotebooksRuntimeRef{}

// NotebooksRuntimeRef is a reference to a GCP NotebooksRuntime.
type NotebooksRuntimeRef struct {
	// A reference to an externally managed NotebooksRuntime resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/runtimes/{{runtimeID}}".
	External string `json:"external,omitempty"`

	// The name of a NotebooksRuntime resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NotebooksRuntime resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&NotebooksRuntimeRef{}, &NotebooksRuntime{})
}

func (r *NotebooksRuntimeRef) GetGVK() schema.GroupVersionKind {
	return NotebooksRuntimeGVK
}

func (r *NotebooksRuntimeRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NotebooksRuntimeRef) GetExternal() string {
	return r.External
}

func (r *NotebooksRuntimeRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *NotebooksRuntimeRef) ValidateExternal(ref string) error {
	id := &NotebooksRuntimeIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *NotebooksRuntimeRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NotebooksRuntimeIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NotebooksRuntimeRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
