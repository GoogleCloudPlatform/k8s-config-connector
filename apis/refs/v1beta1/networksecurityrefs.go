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

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type NetworkSecurityInterceptDeploymentGroupRef struct {
	/* A reference to an externally managed NetworkSecurityInterceptDeploymentGroup resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/interceptDeploymentGroups/{{interceptDeploymentGroupID}}". */
	External string `json:"external,omitempty"`
	/* The name field of a NetworkSecurityInterceptDeploymentGroup resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a NetworkSecurityInterceptDeploymentGroup resource. */
	Namespace string `json:"namespace,omitempty"`
}

type NetworkSecurityMirroringDeploymentGroupRef struct {
	/* A reference to an externally managed NetworkSecurityMirroringDeploymentGroup resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/mirroringDeploymentGroups/{{mirroringDeploymentGroupID}}". */
	External string `json:"external,omitempty"`
	/* The name field of a NetworkSecurityMirroringDeploymentGroup resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a NetworkSecurityMirroringDeploymentGroup resource. */
	Namespace string `json:"namespace,omitempty"`
}

type NetworkSecurityInterceptEndpointGroupRef struct {
	/* A reference to an externally managed NetworkSecurityInterceptEndpointGroup resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/interceptEndpointGroups/{{interceptEndpointGroupID}}". */
	External string `json:"external,omitempty"`
	/* The name field of a NetworkSecurityInterceptEndpointGroup resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a NetworkSecurityInterceptEndpointGroup resource. */
	Namespace string `json:"namespace,omitempty"`
}

type NetworkSecurityMirroringEndpointGroupRef struct {
	/* A reference to an externally managed NetworkSecurityMirroringEndpointGroup resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/mirroringEndpointGroups/{{mirroringEndpointGroupID}}". */
	External string `json:"external,omitempty"`
	/* The name field of a NetworkSecurityMirroringEndpointGroup resource. */
	Name string `json:"name,omitempty"`
	/* The namespace field of a NetworkSecurityMirroringEndpointGroup resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *NetworkSecurityInterceptDeploymentGroupRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkSecurityInterceptDeploymentGroup"}
}
func (r *NetworkSecurityInterceptDeploymentGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}
func (r *NetworkSecurityInterceptDeploymentGroupRef) GetExternal() string               { return r.External }
func (r *NetworkSecurityInterceptDeploymentGroupRef) SetExternal(ref string)            { r.External = ref }
func (r *NetworkSecurityInterceptDeploymentGroupRef) ValidateExternal(ref string) error { return nil }
func (r *NetworkSecurityInterceptDeploymentGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}

func (r *NetworkSecurityInterceptEndpointGroupRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkSecurityInterceptEndpointGroup"}
}
func (r *NetworkSecurityInterceptEndpointGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}
func (r *NetworkSecurityInterceptEndpointGroupRef) GetExternal() string               { return r.External }
func (r *NetworkSecurityInterceptEndpointGroupRef) SetExternal(ref string)            { r.External = ref }
func (r *NetworkSecurityInterceptEndpointGroupRef) ValidateExternal(ref string) error { return nil }
func (r *NetworkSecurityInterceptEndpointGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}

func (r *NetworkSecurityMirroringEndpointGroupRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkSecurityMirroringEndpointGroup"}
}
func (r *NetworkSecurityMirroringEndpointGroupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
}
func (r *NetworkSecurityMirroringEndpointGroupRef) GetExternal() string               { return r.External }
func (r *NetworkSecurityMirroringEndpointGroupRef) SetExternal(ref string)            { r.External = ref }
func (r *NetworkSecurityMirroringEndpointGroupRef) ValidateExternal(ref string) error { return nil }
func (r *NetworkSecurityMirroringEndpointGroupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}
