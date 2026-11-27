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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &NetAppSnapshotRef{}
var NetAppSnapshotGVK = GroupVersion.WithKind("NetAppSnapshot")

// NetAppSnapshotRef is a reference to a NetAppSnapshot.
type NetAppSnapshotRef struct {
	/* A reference to an externally managed NetAppSnapshot resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/volumes/{{volumeID}}/snapshots/{{snapshotID}}". */
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a NetAppSnapshot resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetAppSnapshot resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

func (r *NetAppSnapshotRef) GetGVK() schema.GroupVersionKind {
	return NetAppSnapshotGVK
}

func (r *NetAppSnapshotRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *NetAppSnapshotRef) GetExternal() string {
	return r.External
}

func (r *NetAppSnapshotRef) SetExternal(ref string) {
	r.External = ref
}

func (r *NetAppSnapshotRef) ValidateExternal(ref string) error {
	id := &NetAppSnapshotIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *NetAppSnapshotRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetAppSnapshotIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetAppSnapshotRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", NetAppSnapshotGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
