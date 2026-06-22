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

var _ refs.Ref = &NetworkSecurityDnsThreatDetectorRef{}

// NetworkSecurityDnsThreatDetectorRef defines the resource reference to NetworkSecurityDnsThreatDetector
type NetworkSecurityDnsThreatDetectorRef struct {
	// A reference to an externally managed NetworkSecurityDnsThreatDetector resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/dnsThreatDetectors/{{dns_threat_detector}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkSecurityDnsThreatDetector resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkSecurityDnsThreatDetector resource.
	Namespace string `json:"namespace,omitempty"`
}

// GetGVK returns the GroupVersionKind for NetworkSecurityDnsThreatDetector
func (r *NetworkSecurityDnsThreatDetectorRef) GetGVK() schema.GroupVersionKind {
	return NetworkSecurityDnsThreatDetectorGVK
}

// GetNamespacedName returns the NamespacedName for NetworkSecurityDnsThreatDetector
func (r *NetworkSecurityDnsThreatDetectorRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

// GetExternal returns the External string for NetworkSecurityDnsThreatDetector
func (r *NetworkSecurityDnsThreatDetectorRef) GetExternal() string {
	return r.External
}

// SetExternal sets the External string for NetworkSecurityDnsThreatDetector
func (r *NetworkSecurityDnsThreatDetectorRef) SetExternal(external string) {
	r.External = external
}

// ValidateExternal checks if the External string is valid
func (r *NetworkSecurityDnsThreatDetectorRef) ValidateExternal(external string) error {
	identity := &NetworkSecurityDnsThreatDetectorIdentity{}
	return identity.FromExternal(external)
}

// ParseExternalToIdentity parses the External string to an identity
func (r *NetworkSecurityDnsThreatDetectorRef) ParseExternalToIdentity() (any, error) {
	identity := &NetworkSecurityDnsThreatDetectorIdentity{}
	err := identity.FromExternal(r.External)
	return identity, err
}

// Normalize normalizes the reference
func (r *NetworkSecurityDnsThreatDetectorRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromNetworkSecurityDnsThreatDetectorSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	})
}

func init() {
	refs.Register(&NetworkSecurityDnsThreatDetectorRef{})
}
