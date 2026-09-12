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

var _ refs.Ref = &AgentRegistryServiceRef{}
var _ refs.ExternalRef = &AgentRegistryServiceRef{}

// AgentRegistryServiceRef is a reference to a GCP AgentRegistryService.
type AgentRegistryServiceRef struct {
	// A reference to an externally managed AgentRegistryService resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/services/{{service}}"
	External string `json:"external,omitempty"`

	// The name of an AgentRegistryService resource.
	Name string `json:"name,omitempty"`

	// The namespace of an AgentRegistryService resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *AgentRegistryServiceRef) GetGVK() schema.GroupVersionKind {
	return AgentRegistryServiceGVK
}

func (r *AgentRegistryServiceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *AgentRegistryServiceRef) GetExternal() string {
	return r.External
}

func (r *AgentRegistryServiceRef) SetExternal(external string) {
	r.External = external
}

func (r *AgentRegistryServiceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &AgentRegistryServiceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *AgentRegistryServiceRef) ValidateExternal(external string) error {
	return (&AgentRegistryServiceIdentity{}).FromExternal(external)
}

func (r *AgentRegistryServiceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

func init() {
	refs.Register(&AgentRegistryServiceRef{}, &AgentRegistryService{})
}
