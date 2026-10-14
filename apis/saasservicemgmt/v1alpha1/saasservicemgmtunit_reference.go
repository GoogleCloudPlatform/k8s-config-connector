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

var _ refs.Ref = &SaaSServiceMgmtUnitRef{}

// SaaSServiceMgmtUnitRef is a reference to a GCP SaaSServiceMgmtUnit.
type SaaSServiceMgmtUnitRef struct {
	// A reference to an externally managed SaaSServiceMgmtUnit resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/units/{{unitID}}".
	External string `json:"external,omitempty"`

	// The name of a SaaSServiceMgmtUnit resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SaaSServiceMgmtUnit resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&SaaSServiceMgmtUnitRef{}, &SaaSServiceMgmtUnit{})
}

func (r *SaaSServiceMgmtUnitRef) GetGVK() schema.GroupVersionKind {
	return SaaSServiceMgmtUnitGVK
}

func (r *SaaSServiceMgmtUnitRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SaaSServiceMgmtUnitRef) GetExternal() string {
	return r.External
}

func (r *SaaSServiceMgmtUnitRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *SaaSServiceMgmtUnitRef) ValidateExternal(ref string) error {
	id := &SaaSServiceMgmtUnitIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *SaaSServiceMgmtUnitRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SaaSServiceMgmtUnitIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SaaSServiceMgmtUnitRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

var (
	_ refs.Ref            = &SaaSServiceMgmtUnitKindRef{}
	_ identity.IdentityV2 = &SaaSServiceMgmtUnitKindIdentity{}
)

var SaaSServiceMgmtUnitKindGVK = GroupVersion.WithKind("SaaSServiceMgmtUnitKind")

var SaaSServiceMgmtUnitKindIdentityFormat = gcpurls.Template[SaaSServiceMgmtUnitKindIdentity]("saasservicemgmt.googleapis.com", "projects/{project}/locations/{location}/unitKinds/{unitKind}")

// SaaSServiceMgmtUnitKindIdentity is the identity of a GCP SaaSServiceMgmtUnitKind resource.
// +k8s:deepcopy-gen=false
type SaaSServiceMgmtUnitKindIdentity struct {
	Project  string
	Location string
	UnitKind string
}

func (i *SaaSServiceMgmtUnitKindIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *SaaSServiceMgmtUnitKindIdentity) String() string {
	return SaaSServiceMgmtUnitKindIdentityFormat.ToString(*i)
}

func (i *SaaSServiceMgmtUnitKindIdentity) FromExternal(ref string) error {
	parsed, match, err := SaaSServiceMgmtUnitKindIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SaaSServiceMgmtUnitKind external=%q was not known (use %s): %w", ref, SaaSServiceMgmtUnitKindIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SaaSServiceMgmtUnitKind external=%q was not known (use %s)", ref, SaaSServiceMgmtUnitKindIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SaaSServiceMgmtUnitKindIdentity) Host() string {
	return SaaSServiceMgmtUnitKindIdentityFormat.Host()
}

// SaaSServiceMgmtUnitKindRef is a reference to a GCP SaaSServiceMgmtUnitKind.
type SaaSServiceMgmtUnitKindRef struct {
	// A reference to an externally managed SaaSServiceMgmtUnitKind resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/unitKinds/{{unitKindID}}".
	External string `json:"external,omitempty"`

	// The name of a SaaSServiceMgmtUnitKind resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SaaSServiceMgmtUnitKind resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&SaaSServiceMgmtUnitKindRef{})
}

func (r *SaaSServiceMgmtUnitKindRef) GetGVK() schema.GroupVersionKind {
	return SaaSServiceMgmtUnitKindGVK
}

func (r *SaaSServiceMgmtUnitKindRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SaaSServiceMgmtUnitKindRef) GetExternal() string {
	return r.External
}

func (r *SaaSServiceMgmtUnitKindRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *SaaSServiceMgmtUnitKindRef) ValidateExternal(ref string) error {
	id := &SaaSServiceMgmtUnitKindIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *SaaSServiceMgmtUnitKindRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SaaSServiceMgmtUnitKindIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SaaSServiceMgmtUnitKindRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
