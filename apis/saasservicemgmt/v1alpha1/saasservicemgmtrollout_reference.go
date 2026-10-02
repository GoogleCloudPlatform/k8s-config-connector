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

var (
	_ refs.Ref            = &SaaSServiceMgmtRolloutRef{}
	_ refs.Ref            = &SaaSServiceMgmtRolloutKindRef{}
	_ identity.IdentityV2 = &SaaSServiceMgmtRolloutKindIdentity{}
)

// SaaSServiceMgmtRolloutRef is a reference to a GCP SaaSServiceMgmtRollout.
type SaaSServiceMgmtRolloutRef struct {
	// A reference to an externally managed SaaSServiceMgmtRollout resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/rollouts/{{rolloutID}}".
	External string `json:"external,omitempty"`

	// The name of a SaaSServiceMgmtRollout resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SaaSServiceMgmtRollout resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&SaaSServiceMgmtRolloutRef{}, &SaaSServiceMgmtRollout{})
	refs.Register(&SaaSServiceMgmtRolloutKindRef{})
}

func (r *SaaSServiceMgmtRolloutRef) GetGVK() schema.GroupVersionKind {
	return SaaSServiceMgmtRolloutGVK
}

func (r *SaaSServiceMgmtRolloutRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SaaSServiceMgmtRolloutRef) GetExternal() string {
	return r.External
}

func (r *SaaSServiceMgmtRolloutRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

func (r *SaaSServiceMgmtRolloutRef) ValidateExternal(external string) error {
	id := &SaaSServiceMgmtRolloutIdentity{}
	return id.FromExternal(external)
}

func (r *SaaSServiceMgmtRolloutRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SaaSServiceMgmtRolloutIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SaaSServiceMgmtRolloutRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

var SaaSServiceMgmtRolloutKindGVK = GroupVersion.WithKind("SaaSServiceMgmtRolloutKind")

var SaaSServiceMgmtRolloutKindIdentityFormat = gcpurls.Template[SaaSServiceMgmtRolloutKindIdentity]("saasservicemgmt.googleapis.com", "projects/{project}/locations/{location}/rolloutKinds/{rolloutKind}")

// SaaSServiceMgmtRolloutKindIdentity is the identity of a GCP SaaSServiceMgmtRolloutKind resource.
// +k8s:deepcopy-gen=false
type SaaSServiceMgmtRolloutKindIdentity struct {
	Project     string
	Location    string
	RolloutKind string
}

func (i *SaaSServiceMgmtRolloutKindIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *SaaSServiceMgmtRolloutKindIdentity) String() string {
	return SaaSServiceMgmtRolloutKindIdentityFormat.ToString(*i)
}

func (i *SaaSServiceMgmtRolloutKindIdentity) FromExternal(ref string) error {
	parsed, match, err := SaaSServiceMgmtRolloutKindIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SaaSServiceMgmtRolloutKind external=%q was not known (use %s): %w", ref, SaaSServiceMgmtRolloutKindIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SaaSServiceMgmtRolloutKind external=%q was not known (use %s)", ref, SaaSServiceMgmtRolloutKindIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SaaSServiceMgmtRolloutKindIdentity) Host() string {
	return SaaSServiceMgmtRolloutKindIdentityFormat.Host()
}

// SaaSServiceMgmtRolloutKindRef is a reference to a GCP SaaSServiceMgmtRolloutKind.
type SaaSServiceMgmtRolloutKindRef struct {
	// A reference to an externally managed SaaSServiceMgmtRolloutKind resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/rolloutKinds/{{rolloutKindID}}".
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a SaaSServiceMgmtRolloutKind resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SaaSServiceMgmtRolloutKind resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

func (r *SaaSServiceMgmtRolloutKindRef) GetGVK() schema.GroupVersionKind {
	return SaaSServiceMgmtRolloutKindGVK
}

func (r *SaaSServiceMgmtRolloutKindRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *SaaSServiceMgmtRolloutKindRef) GetExternal() string {
	return r.External
}

func (r *SaaSServiceMgmtRolloutKindRef) SetExternal(external string) {
	r.External = external
}

func (r *SaaSServiceMgmtRolloutKindRef) ValidateExternal(external string) error {
	id := &SaaSServiceMgmtRolloutKindIdentity{}
	return id.FromExternal(external)
}

func (r *SaaSServiceMgmtRolloutKindRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SaaSServiceMgmtRolloutKindIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SaaSServiceMgmtRolloutKindRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
