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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &NetworkSecurityMirroringEndpointGroupAssociationRef{}

// NetworkSecurityMirroringEndpointGroupAssociationRef defines the resource reference to NetworkSecurityMirroringEndpointGroupAssociation, which "External" field
// holds the GCP identifier for the KRM object.
type NetworkSecurityMirroringEndpointGroupAssociationRef struct {
	// A reference to an externally managed NetworkSecurityMirroringEndpointGroupAssociation resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/mirroringEndpointGroupAssociations/{{mirroring_endpoint_group_association}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkSecurityMirroringEndpointGroupAssociation resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkSecurityMirroringEndpointGroupAssociation resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&NetworkSecurityMirroringEndpointGroupAssociationRef{})
}

func (r *NetworkSecurityMirroringEndpointGroupAssociationRef) GetGVK() schema.GroupVersionKind {
	return NetworkSecurityMirroringEndpointGroupAssociationGVK
}

func (r *NetworkSecurityMirroringEndpointGroupAssociationRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkSecurityMirroringEndpointGroupAssociationRef) GetExternal() string {
	return r.External
}

func (r *NetworkSecurityMirroringEndpointGroupAssociationRef) SetExternal(external string) {
	r.External = external
}

func (r *NetworkSecurityMirroringEndpointGroupAssociationRef) ValidateExternal(external string) error {
	id := &NetworkSecurityMirroringEndpointGroupAssociationIdentity{}
	return id.FromExternal(external)
}

func (r *NetworkSecurityMirroringEndpointGroupAssociationRef) ParseExternalToIdentity(external string) (interface{}, error) {
	id := &NetworkSecurityMirroringEndpointGroupAssociationIdentity{}
	err := id.FromExternal(external)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkSecurityMirroringEndpointGroupAssociationRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromNetworkSecurityMirroringEndpointGroupAssociationSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
