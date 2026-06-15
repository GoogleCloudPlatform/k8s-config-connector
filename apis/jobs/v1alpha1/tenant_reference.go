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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var CloudTalentSolutionTenantGVK = GroupVersion.WithKind("CloudTalentSolutionTenant")

var _ refs.Ref = &CloudTalentSolutionTenantRef{}

// CloudTalentSolutionTenantRef is a reference to a CloudTalentSolutionTenant.
type CloudTalentSolutionTenantRef struct {
	// A reference to an externally managed CloudTalentSolutionTenant resource.
	// Should be in the format "projects/{{projectID}}/tenants/{{tenantID}}".
	External string `json:"external,omitempty"`

	// The name of a CloudTalentSolutionTenant resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CloudTalentSolutionTenant resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CloudTalentSolutionTenantRef{}, nil)
}

func (r *CloudTalentSolutionTenantRef) GetGVK() schema.GroupVersionKind {
	return CloudTalentSolutionTenantGVK
}

func (r *CloudTalentSolutionTenantRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CloudTalentSolutionTenantRef) GetExternal() string {
	return r.External
}

func (r *CloudTalentSolutionTenantRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CloudTalentSolutionTenantRef) ValidateExternal(ref string) error {
	id := &CloudTalentSolutionTenantIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CloudTalentSolutionTenantRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CloudTalentSolutionTenantIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CloudTalentSolutionTenantRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		return ""
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

var CloudTalentSolutionTenantIdentityFormat = gcpurls.Template[CloudTalentSolutionTenantIdentity]("jobs.googleapis.com", "projects/{project}/tenants/{tenant}")

// CloudTalentSolutionTenantIdentity represents the GCP Identity of a CloudTalentSolutionTenant resource.
// +k8s:deepcopy-gen=false
type CloudTalentSolutionTenantIdentity struct {
	Project string
	Tenant  string
}

func (i *CloudTalentSolutionTenantIdentity) String() string {
	return CloudTalentSolutionTenantIdentityFormat.ToString(*i)
}

func (i *CloudTalentSolutionTenantIdentity) FromExternal(ref string) error {
	parsed, match, err := CloudTalentSolutionTenantIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CloudTalentSolutionTenant external=%q was not known (use %s): %w", ref, CloudTalentSolutionTenantIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CloudTalentSolutionTenant external=%q was not known (use %s)", ref, CloudTalentSolutionTenantIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *CloudTalentSolutionTenantIdentity) Host() string {
	return CloudTalentSolutionTenantIdentityFormat.Host()
}
