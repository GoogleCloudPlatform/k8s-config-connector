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
	_ refs.Ref            = &SaaSServiceMgmtTenantRef{}
	_ identity.IdentityV2 = &SaaSServiceMgmtTenantIdentity{}
)

var SaaSServiceMgmtTenantGVK = GroupVersion.WithKind("SaaSServiceMgmtTenant")

var SaaSServiceMgmtTenantIdentityFormat = gcpurls.Template[SaaSServiceMgmtTenantIdentity]("saasservicemgmt.googleapis.com", "projects/{project}/locations/{location}/tenants/{tenant}")

// SaaSServiceMgmtTenantIdentity is the identity of a GCP SaaSServiceMgmtTenant resource.
// +k8s:deepcopy-gen=false
type SaaSServiceMgmtTenantIdentity struct {
	Project  string
	Location string
	Tenant   string
}

func (i *SaaSServiceMgmtTenantIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *SaaSServiceMgmtTenantIdentity) String() string {
	return SaaSServiceMgmtTenantIdentityFormat.ToString(*i)
}

func (i *SaaSServiceMgmtTenantIdentity) FromExternal(ref string) error {
	parsed, match, err := SaaSServiceMgmtTenantIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SaaSServiceMgmtTenant external=%q was not known (use %s): %w", ref, SaaSServiceMgmtTenantIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SaaSServiceMgmtTenant external=%q was not known (use %s)", ref, SaaSServiceMgmtTenantIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *SaaSServiceMgmtTenantIdentity) Host() string {
	return SaaSServiceMgmtTenantIdentityFormat.Host()
}

// SaaSServiceMgmtTenantRef is a reference to a GCP SaaSServiceMgmtTenant.
type SaaSServiceMgmtTenantRef struct {
	// A reference to an externally managed SaaSServiceMgmtTenant resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/tenants/{{tenantID}}".
	External string `json:"external,omitempty"`

	// The name of a SaaSServiceMgmtTenant resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SaaSServiceMgmtTenant resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&SaaSServiceMgmtTenantRef{})
}

func (r *SaaSServiceMgmtTenantRef) GetGVK() schema.GroupVersionKind {
	return SaaSServiceMgmtTenantGVK
}

func (r *SaaSServiceMgmtTenantRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *SaaSServiceMgmtTenantRef) GetExternal() string {
	return r.External
}

func (r *SaaSServiceMgmtTenantRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *SaaSServiceMgmtTenantRef) ValidateExternal(ref string) error {
	id := &SaaSServiceMgmtTenantIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *SaaSServiceMgmtTenantRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SaaSServiceMgmtTenantIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SaaSServiceMgmtTenantRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}
