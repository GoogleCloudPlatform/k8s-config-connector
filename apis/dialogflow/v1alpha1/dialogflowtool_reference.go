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

var _ refs.Ref = &DialogflowToolRef{}

// DialogflowToolRef is a reference to a GCP DialogflowTool.
type DialogflowToolRef struct {
	// A reference to an externally managed DialogflowTool resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/agents/{{agentID}}/tools/{{toolID}}".
	External string `json:"external,omitempty"`

	// The name of a DialogflowTool resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DialogflowTool resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DialogflowToolRef{})
}

func (r *DialogflowToolRef) GetGVK() schema.GroupVersionKind {
	return DialogflowToolGVK
}

func (r *DialogflowToolRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DialogflowToolRef) GetExternal() string {
	return r.External
}

func (r *DialogflowToolRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DialogflowToolRef) ValidateExternal(ref string) error {
	id := &DialogflowToolIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DialogflowToolRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DialogflowToolIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DialogflowToolRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
