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

var _ refs.Ref = &NetworkSecurityMirroringDeploymentRef{}

// NetworkSecurityMirroringDeploymentRef is a reference to a NetworkSecurityMirroringDeployment.
type NetworkSecurityMirroringDeploymentRef struct {
	// A reference to an externally managed NetworkSecurityMirroringDeployment resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/mirroringDeployments/{{mirroring_deployment}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkSecurityMirroringDeployment resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkSecurityMirroringDeployment resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&NetworkSecurityMirroringDeploymentRef{}, &NetworkSecurityMirroringDeployment{})
}

func (r *NetworkSecurityMirroringDeploymentRef) GetGVK() schema.GroupVersionKind {
	return NetworkSecurityMirroringDeploymentGVK
}

func (r *NetworkSecurityMirroringDeploymentRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkSecurityMirroringDeploymentRef) GetExternal() string {
	return r.External
}

func (r *NetworkSecurityMirroringDeploymentRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

func (r *NetworkSecurityMirroringDeploymentRef) ValidateExternal(external string) error {
	id := &NetworkSecurityMirroringDeploymentIdentity{}
	return id.FromExternal(external)
}

func (r *NetworkSecurityMirroringDeploymentRef) ParseExternalToIdentity(external string) (interface{}, error) {
	id := &NetworkSecurityMirroringDeploymentIdentity{}
	err := id.FromExternal(external)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkSecurityMirroringDeploymentRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromNetworkSecurityMirroringDeploymentSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
