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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &NetworkConnectivityTransportRef{}
var _ refs.ExternalRef = &NetworkConnectivityTransportRef{}

// NetworkConnectivityTransportRef is a reference to a GCP NetworkConnectivityTransport.
type NetworkConnectivityTransportRef struct {
	// A reference to an externally managed NetworkConnectivityTransport resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/transports/{{transport}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkConnectivityTransport resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkConnectivityTransport resource.
	Namespace string `json:"namespace,omitempty"`
}

// GetGVK returns the GroupVersionKind for NetworkConnectivityTransport
func (r *NetworkConnectivityTransportRef) GetGVK() schema.GroupVersionKind {
	return NetworkConnectivityTransportGVK
}

// GetNamespacedName returns the NamespacedName for NetworkConnectivityTransport
func (r *NetworkConnectivityTransportRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

// GetExternal returns the External string for NetworkConnectivityTransport
func (r *NetworkConnectivityTransportRef) GetExternal() string {
	return r.External
}

// SetExternal sets the External string for NetworkConnectivityTransport
func (r *NetworkConnectivityTransportRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

// ValidateExternal checks if the External string is valid
func (r *NetworkConnectivityTransportRef) ValidateExternal(external string) error {
	identity := &NetworkConnectivityTransportIdentity{}
	return identity.FromExternal(external)
}

// ParseExternalToIdentity parses the External string to an identity
func (r *NetworkConnectivityTransportRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkConnectivityTransportIdentity{}
	err := id.FromExternal(r.External)
	return id, err
}

// Normalize normalizes the reference
func (r *NetworkConnectivityTransportRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.Normalize(ctx, reader, r, otherNamespace)
}

var _ refs.Ref = &NetworkConnectivityRemoteTransportProfileRef{}
var _ refs.ExternalRef = &NetworkConnectivityRemoteTransportProfileRef{}
var _ identity.IdentityV2 = &NetworkConnectivityRemoteTransportProfileIdentity{}

var NetworkConnectivityRemoteTransportProfileGVK = schema.GroupVersionKind{
	Group:   "networkconnectivity.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "NetworkConnectivityRemoteTransportProfile",
}

// NetworkConnectivityRemoteTransportProfileRef is a reference to a GCP NetworkConnectivityRemoteTransportProfile.
type NetworkConnectivityRemoteTransportProfileRef struct {
	// A reference to an externally managed NetworkConnectivityRemoteTransportProfile resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/remoteTransportProfiles/{{remoteTransportProfile}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkConnectivityRemoteTransportProfile resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkConnectivityRemoteTransportProfile resource.
	Namespace string `json:"namespace,omitempty"`
}

// GetGVK returns the GroupVersionKind for NetworkConnectivityRemoteTransportProfile
func (r *NetworkConnectivityRemoteTransportProfileRef) GetGVK() schema.GroupVersionKind {
	return NetworkConnectivityRemoteTransportProfileGVK
}

// GetNamespacedName returns the NamespacedName for NetworkConnectivityRemoteTransportProfile
func (r *NetworkConnectivityRemoteTransportProfileRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

// GetExternal returns the External string for NetworkConnectivityRemoteTransportProfile
func (r *NetworkConnectivityRemoteTransportProfileRef) GetExternal() string {
	return r.External
}

// SetExternal sets the External string for NetworkConnectivityRemoteTransportProfile
func (r *NetworkConnectivityRemoteTransportProfileRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

var NetworkConnectivityRemoteTransportProfileIdentityFormat = gcpurls.Template[NetworkConnectivityRemoteTransportProfileIdentity]("networkconnectivity.googleapis.com", "projects/{project}/locations/{location}/remoteTransportProfiles/{remoteTransportProfile}")

// NetworkConnectivityRemoteTransportProfileIdentity is the identity of a GCP NetworkConnectivityRemoteTransportProfile resource.
// +k8s:deepcopy-gen=false
type NetworkConnectivityRemoteTransportProfileIdentity struct {
	Project                string
	Location               string
	RemoteTransportProfile string
}

func (i *NetworkConnectivityRemoteTransportProfileIdentity) String() string {
	return NetworkConnectivityRemoteTransportProfileIdentityFormat.ToString(*i)
}

func (i *NetworkConnectivityRemoteTransportProfileIdentity) FromExternal(ref string) error {
	parsed, match, err := NetworkConnectivityRemoteTransportProfileIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NetworkConnectivityRemoteTransportProfile external=%q was not known (use %s): %w", ref, NetworkConnectivityRemoteTransportProfileIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NetworkConnectivityRemoteTransportProfile external=%q was not known (use %s)", ref, NetworkConnectivityRemoteTransportProfileIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NetworkConnectivityRemoteTransportProfileIdentity) Host() string {
	return NetworkConnectivityRemoteTransportProfileIdentityFormat.Host()
}

func (i *NetworkConnectivityRemoteTransportProfileIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

// ValidateExternal checks if the External string is valid
func (r *NetworkConnectivityRemoteTransportProfileRef) ValidateExternal(external string) error {
	id := &NetworkConnectivityRemoteTransportProfileIdentity{}
	return id.FromExternal(external)
}

// ParseExternalToIdentity parses the External string to an identity
func (r *NetworkConnectivityRemoteTransportProfileRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkConnectivityRemoteTransportProfileIdentity{}
	err := id.FromExternal(r.External)
	return id, err
}

// Normalize normalizes the reference
func (r *NetworkConnectivityRemoteTransportProfileRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.Normalize(ctx, reader, r, otherNamespace)
}

func init() {
	refs.Register(&NetworkConnectivityTransportRef{}, &NetworkConnectivityTransport{})
	refs.Register(&NetworkConnectivityRemoteTransportProfileRef{}, nil)
}
