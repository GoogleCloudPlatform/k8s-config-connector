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

func init() {
	refsv1beta1.Register(&CesAppRef{}, &CESApp{})
}

var _ refsv1beta1.Ref = &CesAppRef{}

// CesAppRef is a reference to a GCP CESApp resource.
type CesAppRef struct {
	// A reference to an externally managed CESApp resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}".
	External string `json:"external,omitempty"`

	// The name of a CESApp resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CESApp resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CesAppRef) GetGVK() schema.GroupVersionKind {
	return CESAppGVK
}

func (r *CesAppRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CesAppRef) GetExternal() string {
	return r.External
}

func (r *CesAppRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CesAppRef) ValidateExternal(ref string) error {
	id := &CESAppIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CesAppRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CESAppIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CesAppRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

var (
	CesGuardrailIdentityFormat                     = gcpurls.Template[CesGuardrailIdentity]("ces.googleapis.com", "projects/{project}/locations/{location}/apps/{app}/guardrails/{guardrail}")
	_                          identity.IdentityV2 = &CesGuardrailIdentity{}
	CesGuardrailGVK                                = GroupVersion.WithKind("CesGuardrail")
	_                          refsv1beta1.Ref     = &CesGuardrailRef{}
)

// +k8s:deepcopy-gen=false
type CesGuardrailIdentity struct {
	Project   string
	Location  string
	App       string
	Guardrail string
}

func (i *CesGuardrailIdentity) String() string {
	return CesGuardrailIdentityFormat.ToString(*i)
}

func (i *CesGuardrailIdentity) FromExternal(ref string) error {
	parsed, match, err := CesGuardrailIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CesGuardrail external=%q was not known (use %s): %w", ref, CesGuardrailIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CesGuardrail external=%q was not known (use %s)", ref, CesGuardrailIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *CesGuardrailIdentity) Host() string {
	return CesGuardrailIdentityFormat.Host()
}

func (i *CesGuardrailIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location + "/apps/" + i.App
}

// CesGuardrailRef is a reference to a GCP CesGuardrail resource.
type CesGuardrailRef struct {
	// A reference to an externally managed CesGuardrail resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}/guardrails/{{guardrailID}}".
	External string `json:"external,omitempty"`

	// The name of a CesGuardrail resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CesGuardrail resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CesGuardrailRef) GetGVK() schema.GroupVersionKind {
	return CesGuardrailGVK
}

func (r *CesGuardrailRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CesGuardrailRef) GetExternal() string {
	return r.External
}

func (r *CesGuardrailRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CesGuardrailRef) ValidateExternal(ref string) error {
	id := &CesGuardrailIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CesGuardrailRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CesGuardrailIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CesGuardrailRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

var (
	CesToolIdentityFormat                     = gcpurls.Template[CesToolIdentity]("ces.googleapis.com", "projects/{project}/locations/{location}/apps/{app}/tools/{tool}")
	_                     identity.IdentityV2 = &CesToolIdentity{}
	CesToolGVK                                = GroupVersion.WithKind("CesTool")
	_                     refsv1beta1.Ref     = &CesToolRef{}
)

// +k8s:deepcopy-gen=false
type CesToolIdentity struct {
	Project  string
	Location string
	App      string
	Tool     string
}

func (i *CesToolIdentity) String() string {
	return CesToolIdentityFormat.ToString(*i)
}

func (i *CesToolIdentity) FromExternal(ref string) error {
	parsed, match, err := CesToolIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CesTool external=%q was not known (use %s): %w", ref, CesToolIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CesTool external=%q was not known (use %s)", ref, CesToolIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *CesToolIdentity) Host() string {
	return CesToolIdentityFormat.Host()
}

func (i *CesToolIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location + "/apps/" + i.App
}

// CesToolRef is a reference to a GCP CesTool resource.
type CesToolRef struct {
	// A reference to an externally managed CesTool resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}/tools/{{toolID}}".
	External string `json:"external,omitempty"`

	// The name of a CesTool resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CesTool resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CesToolRef) GetGVK() schema.GroupVersionKind {
	return CesToolGVK
}

func (r *CesToolRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CesToolRef) GetExternal() string {
	return r.External
}

func (r *CesToolRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CesToolRef) ValidateExternal(ref string) error {
	id := &CesToolIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CesToolRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CesToolIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CesToolRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

var (
	CesToolsetIdentityFormat                     = gcpurls.Template[CesToolsetIdentity]("ces.googleapis.com", "projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}")
	_                        identity.IdentityV2 = &CesToolsetIdentity{}
	CesToolsetGVK                                = GroupVersion.WithKind("CesToolset")
	_                        refsv1beta1.Ref     = &CesToolsetRef{}
)

// +k8s:deepcopy-gen=false
type CesToolsetIdentity struct {
	Project  string
	Location string
	App      string
	Toolset  string
}

func (i *CesToolsetIdentity) String() string {
	return CesToolsetIdentityFormat.ToString(*i)
}

func (i *CesToolsetIdentity) FromExternal(ref string) error {
	parsed, match, err := CesToolsetIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of CesToolset external=%q was not known (use %s): %w", ref, CesToolsetIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of CesToolset external=%q was not known (use %s)", ref, CesToolsetIdentityFormat.CanonicalForm())
	}
	*i = *parsed
	return nil
}

func (i *CesToolsetIdentity) Host() string {
	return CesToolsetIdentityFormat.Host()
}

func (i *CesToolsetIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location + "/apps/" + i.App
}

// CesToolsetRef is a reference to a GCP CesToolset resource.
type CesToolsetRef struct {
	// A reference to an externally managed CesToolset resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}/toolsets/{{toolsetID}}".
	External string `json:"external,omitempty"`

	// The name of a CesToolset resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CesToolset resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *CesToolsetRef) GetGVK() schema.GroupVersionKind {
	return CesToolsetGVK
}

func (r *CesToolsetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CesToolsetRef) GetExternal() string {
	return r.External
}

func (r *CesToolsetRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CesToolsetRef) ValidateExternal(ref string) error {
	id := &CesToolsetIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CesToolsetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CesToolsetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CesToolsetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

// +kcc:proto=google.cloud.ces.v1.Agent.AgentToolset
