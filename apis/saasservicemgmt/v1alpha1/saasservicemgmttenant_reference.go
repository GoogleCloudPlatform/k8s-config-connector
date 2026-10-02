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
	_ refs.Ref         = &SaaSServiceMgmtTenantRef{}
	_ refs.ExternalRef = &SaaSServiceMgmtTenantRef{}
)

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
	refs.Register(&SaaSServiceMgmtTenantRef{}, &SaaSServiceMgmtTenant{})
	refs.Register(&SaaSServiceMgmtSaaSRef{})
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

var (
	_ refs.Ref            = &SaaSServiceMgmtSaaSRef{}
	_ refs.ExternalRef    = &SaaSServiceMgmtSaaSRef{}
	_ identity.IdentityV2 = &SaaSServiceMgmtSaaSIdentity{}
)

var SaaSServiceMgmtSaaSGVK = GroupVersion.WithKind("SaaSServiceMgmtSaaS")

var SaaSServiceMgmtSaaSIdentityFormat = gcpurls.Template[SaaSServiceMgmtSaaSIdentity]("saasservicemgmt.googleapis.com", "projects/{project}/locations/{location}/saas/{saas}")

// SaaSServiceMgmtSaaSIdentity is the identity of a GCP SaaSServiceMgmtSaaS resource.
// +k8s:deepcopy-gen=false
type SaaSServiceMgmtSaaSIdentity struct {
	Project  string
	Location string
	Saas     string
}

func (i *SaaSServiceMgmtSaaSIdentity) String() string {
	return SaaSServiceMgmtSaaSIdentityFormat.ToString(*i)
}

func (i *SaaSServiceMgmtSaaSIdentity) Host() string {
	return SaaSServiceMgmtSaaSIdentityFormat.Host()
}

func (i *SaaSServiceMgmtSaaSIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func (i *SaaSServiceMgmtSaaSIdentity) FromExternal(ref string) error {
	parsed, match, err := SaaSServiceMgmtSaaSIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of SaaSServiceMgmtSaaS external=%q was not known (use %s): %w", ref, SaaSServiceMgmtSaaSIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of SaaSServiceMgmtSaaS external=%q was not known (use %s)", ref, SaaSServiceMgmtSaaSIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

// SaaSServiceMgmtSaaSRef is a reference to a GCP SaaSServiceMgmtSaaS.
type SaaSServiceMgmtSaaSRef struct {
	// A reference to an externally managed SaaSServiceMgmtSaaS resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/saas/{{saas}}".
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a SaaSServiceMgmtSaaS resource.
	Name string `json:"name,omitempty"`

	// The namespace of a SaaSServiceMgmtSaaS resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

func (r *SaaSServiceMgmtSaaSRef) GetGVK() schema.GroupVersionKind {
	return SaaSServiceMgmtSaaSGVK
}

func (r *SaaSServiceMgmtSaaSRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *SaaSServiceMgmtSaaSRef) GetExternal() string {
	return r.External
}

func (r *SaaSServiceMgmtSaaSRef) SetExternal(ref string) {
	r.External = ref
}

func (r *SaaSServiceMgmtSaaSRef) ValidateExternal(ref string) error {
	id := &SaaSServiceMgmtSaaSIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *SaaSServiceMgmtSaaSRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &SaaSServiceMgmtSaaSIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *SaaSServiceMgmtSaaSRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", SaaSServiceMgmtSaaSGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
