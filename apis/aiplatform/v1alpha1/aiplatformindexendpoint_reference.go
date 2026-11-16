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

var _ refs.Ref = &AIPlatformIndexEndpointRef{}

var AIPlatformIndexEndpointGVK = GroupVersion.WithKind("AIPlatformIndexEndpoint")

// AIPlatformIndexEndpointRef is a reference to an AIPlatformIndexEndpoint.
type AIPlatformIndexEndpointRef struct {
	// A reference to an externally managed AIPlatformIndexEndpoint resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/indexEndpoints/{{indexEndpointID}}".
	External string `json:"external,omitempty"`

	// The name of an AIPlatformIndexEndpoint resource.
	Name string `json:"name,omitempty"`

	// The namespace of an AIPlatformIndexEndpoint resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&AIPlatformIndexEndpointRef{})
}

func (r *AIPlatformIndexEndpointRef) GetGVK() schema.GroupVersionKind {
	return AIPlatformIndexEndpointGVK
}

func (r *AIPlatformIndexEndpointRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *AIPlatformIndexEndpointRef) GetExternal() string {
	return r.External
}

func (r *AIPlatformIndexEndpointRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *AIPlatformIndexEndpointRef) ValidateExternal(ref string) error {
	id := &AIPlatformIndexEndpointIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *AIPlatformIndexEndpointRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &AIPlatformIndexEndpointIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *AIPlatformIndexEndpointRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
