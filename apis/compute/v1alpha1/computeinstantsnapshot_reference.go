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

var _ refs.Ref = &ComputeInstantSnapshotRef{}

// ComputeInstantSnapshotRef is a reference to a GCP ComputeInstantSnapshot resource.
type ComputeInstantSnapshotRef struct {
	// A reference to an externally managed ComputeInstantSnapshot resource. Should be in the format "projects/{{projectID}}/zones/{{zone}}/instantSnapshots/{{instantSnapshot}}" or "projects/{{projectID}}/regions/{{region}}/instantSnapshots/{{instantSnapshot}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeInstantSnapshot resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeInstantSnapshot resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeInstantSnapshotRef{}, &ComputeInstantSnapshot{})
}

func (r *ComputeInstantSnapshotRef) GetGVK() schema.GroupVersionKind {
	return ComputeInstantSnapshotGVK
}

func (r *ComputeInstantSnapshotRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeInstantSnapshotRef) GetExternal() string {
	return r.External
}

func (r *ComputeInstantSnapshotRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeInstantSnapshotRef) ValidateExternal(ref string) error {
	id := &ComputeInstantSnapshotIdentity{}
	return id.FromExternal(ref)
}

func (r *ComputeInstantSnapshotRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeInstantSnapshotIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeInstantSnapshotRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
