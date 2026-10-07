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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var DiscoveryEngineCollectionGVK = schema.GroupVersionKind{
	Group:   "discoveryengine.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "DiscoveryEngineCollection",
}

var _ refsv1beta1.Ref = &DiscoveryEngineCollectionRef{}
var _ refsv1beta1.ExternalRef = &DiscoveryEngineCollectionRef{}
var _ refsv1beta1.ExternalNormalizer = &DiscoveryEngineCollectionRef{}

// DiscoveryEngineCollectionRef is a reference to a GCP DiscoveryEngineCollection.
type DiscoveryEngineCollectionRef struct {
	// A reference to an externally managed DiscoveryEngineCollection resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/collections/{{collectionID}}".
	External string `json:"external,omitempty"`

	// The name of a DiscoveryEngineCollection resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DiscoveryEngineCollection resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&DiscoveryEngineCollectionRef{})
}

func (r *DiscoveryEngineCollectionRef) GetGVK() schema.GroupVersionKind {
	return DiscoveryEngineCollectionGVK
}

func (r *DiscoveryEngineCollectionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *DiscoveryEngineCollectionRef) GetExternal() string {
	return r.External
}

func (r *DiscoveryEngineCollectionRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *DiscoveryEngineCollectionRef) ValidateExternal(ref string) error {
	identity := &DiscoveryEngineCollectionIdentity{}
	return identity.FromExternal(ref)
}

func (r *DiscoveryEngineCollectionRef) ParseExternalToIdentity() (identity.Identity, error) {
	identity := &DiscoveryEngineCollectionIdentity{}
	if err := identity.FromExternal(r.External); err != nil {
		return nil, err
	}
	return identity, nil
}

// Normalize ensures the "External" reference (in string format) is set.
func (r *DiscoveryEngineCollectionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, nil)
}

// NormalizedExternal provisions the "External" value for resources depending on DiscoveryEngineCollection (for compatibility).
func (r *DiscoveryEngineCollectionRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on %s reference", DiscoveryEngineCollectionGVK.Kind)
	}

	// From given External
	if r.External != "" {
		identity := &DiscoveryEngineCollectionIdentity{}
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

var (
	_ identity.IdentityV2 = &DiscoveryEngineCollectionIdentity{}
)

var DiscoveryEngineCollectionIdentityFormat = gcpurls.Template[DiscoveryEngineCollectionIdentity]("discoveryengine.googleapis.com", "projects/{project}/locations/{location}/collections/{collection}")

// DiscoveryEngineCollectionIdentity is the identity of a GCP DiscoveryEngineCollection resource.
// +k8s:deepcopy-gen=false
type DiscoveryEngineCollectionIdentity struct {
	Project    string
	Location   string
	Collection string
}

func (i *DiscoveryEngineCollectionIdentity) String() string {
	return DiscoveryEngineCollectionIdentityFormat.ToString(*i)
}

func (i *DiscoveryEngineCollectionIdentity) FromExternal(ref string) error {
	parsed, match, err := DiscoveryEngineCollectionIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of DiscoveryEngineCollection external=%q was not known (use %s): %w", ref, DiscoveryEngineCollectionIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of DiscoveryEngineCollection external=%q was not known (use %s)", ref, DiscoveryEngineCollectionIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DiscoveryEngineCollectionIdentity) Host() string {
	return DiscoveryEngineCollectionIdentityFormat.Host()
}

func (i *DiscoveryEngineCollectionIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}
