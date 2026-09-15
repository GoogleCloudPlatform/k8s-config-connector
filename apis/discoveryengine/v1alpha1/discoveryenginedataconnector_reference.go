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

var _ refsv1beta1.Ref = &DiscoveryEngineDataConnectorRef{}
var _ refsv1beta1.ExternalRef = &DiscoveryEngineDataConnectorRef{}
var _ refsv1beta1.ExternalNormalizer = &DiscoveryEngineDataConnectorRef{}

// DiscoveryEngineDataConnectorRef is a reference to a GCP DiscoveryEngineDataConnector.
type DiscoveryEngineDataConnectorRef struct {
	// A reference to an externally managed DiscoveryEngineDataConnector resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/collections/{{collection}}/dataConnector".
	External string `json:"external,omitempty"`

	// The name of a DiscoveryEngineDataConnector resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DiscoveryEngineDataConnector resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&DiscoveryEngineDataConnectorRef{}, &DiscoveryEngineDataConnector{})
}

func (r *DiscoveryEngineDataConnectorRef) GetGVK() schema.GroupVersionKind {
	return DiscoveryEngineDataConnectorGVK
}

func (r *DiscoveryEngineDataConnectorRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *DiscoveryEngineDataConnectorRef) GetExternal() string {
	return r.External
}

func (r *DiscoveryEngineDataConnectorRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *DiscoveryEngineDataConnectorRef) ValidateExternal(ref string) error {
	identity := &DiscoveryEngineDataConnectorIdentity{}
	return identity.FromExternal(ref)
}

func (r *DiscoveryEngineDataConnectorRef) ParseExternalToIdentity() (identity.Identity, error) {
	identity := &DiscoveryEngineDataConnectorIdentity{}
	if err := identity.FromExternal(r.External); err != nil {
		return nil, err
	}
	return identity, nil
}

// Normalize ensures the "External" reference (in string format) is set.
func (r *DiscoveryEngineDataConnectorRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

// NormalizedExternal provisions the "External" value for resources depending on DiscoveryEngineDataConnector (for compatibility).
func (r *DiscoveryEngineDataConnectorRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", DiscoveryEngineDataConnectorGVK.Kind)
	}

	// From given External
	if r.External != "" {
		identity := &DiscoveryEngineDataConnectorIdentity{}
		if err := identity.FromExternal(r.External); err != nil {
			return "", err
		}
		r.External = identity.String()
		return r.External, nil
	}

	// Delegate to the standard Normalize helper
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
