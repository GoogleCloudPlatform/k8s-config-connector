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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/reference"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ reference.Reference = &SnapshotRef{}

// SnapshotRef defines the resource reference to PubSubSnapshot, which "External" field
// holds the GCP identifier for the KRM object.
type SnapshotRef struct {
	// A reference to an externally managed PubSubSnapshot resource.
	// Should be in the format "projects/{{projectID}}/snapshots/{{snapshotID}}".
	External string `json:"external,omitempty"`

	// The name of a PubSubSnapshot resource.
	Name string `json:"name,omitempty"`

	// The namespace of a PubSubSnapshot resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *SnapshotRef) GetGVK() schema.GroupVersionKind {
	return PubSubSnapshotGVK
}

func (r *SnapshotRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SnapshotRef) GetExternal() string {
	return r.External
}

func (r *SnapshotRef) SetExternal(ref string) {
	r.External = ref
}

func (r *SnapshotRef) ValidateExternal() error {
	if _, _, err := ParseSnapshotExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *SnapshotRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return reference.Normalize(ctx, reader, r, defaultNamespace)
}
