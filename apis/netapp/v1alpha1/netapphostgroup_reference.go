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

var _ refs.Ref = &NetAppHostGroupRef{}

// NetAppHostGroupRef is a reference to a GCP NetAppHostGroup.
type NetAppHostGroupRef struct {
	// A reference to an externally managed NetAppHostGroup resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/hostGroups/{{hostGroupID}}".
	External string `json:"external,omitempty"`

	// The name of a NetAppHostGroup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetAppHostGroup resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&NetAppHostGroupRef{}, &NetAppHostGroup{})
}

func (r *NetAppHostGroupRef) GetGVK() schema.GroupVersionKind {
	return NetAppHostGroupGVK
}

func (r *NetAppHostGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetAppHostGroupRef) GetExternal() string {
	return r.External
}

func (r *NetAppHostGroupRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *NetAppHostGroupRef) ValidateExternal(ref string) error {
	id := &NetAppHostGroupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *NetAppHostGroupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetAppHostGroupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetAppHostGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
