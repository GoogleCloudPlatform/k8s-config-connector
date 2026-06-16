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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ComputeSnapshotGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeSnapshot",
}

var _ refs.Ref = &ComputeSnapshotRef{}

// ComputeSnapshotRef is a reference to a Google Cloud ComputeSnapshot.
type ComputeSnapshotRef struct {
	// A reference to an externally managed ComputeSnapshot resource. Should be in the format "projects/{{projectID}}/global/snapshots/{{snapshotID}}" or "projects/{{projectID}}/regions/{{region}}/snapshots/{{snapshotID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeSnapshot resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeSnapshot resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeSnapshotRef{}, &ComputeSnapshot{})
}

func (r *ComputeSnapshotRef) GetGVK() schema.GroupVersionKind {
	return ComputeSnapshotGVK
}

func (r *ComputeSnapshotRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeSnapshotRef) GetExternal() string {
	return r.External
}

func (r *ComputeSnapshotRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeSnapshotRef) ValidateExternal(ref string) error {
	id := &ComputeSnapshotIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeSnapshotRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeSnapshotIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeSnapshotRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*ComputeSnapshot](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromComputeSnapshotSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
