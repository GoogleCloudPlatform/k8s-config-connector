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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ComputePublicDelegatedPrefixGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputePublicDelegatedPrefix",
}

var _ refs.Ref = &ComputePublicDelegatedPrefixRef{}

// ComputePublicDelegatedPrefixRef is a reference to a GCP ComputePublicDelegatedPrefix.
type ComputePublicDelegatedPrefixRef struct {
	// A reference to an externally managed ComputePublicDelegatedPrefix resource.
	// Should be in the format "projects/{{projectID}}/global/publicDelegatedPrefixes/{{publicDelegatedPrefixID}}"
	// or "projects/{{projectID}}/regions/{{region}}/publicDelegatedPrefixes/{{publicDelegatedPrefixID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputePublicDelegatedPrefix resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputePublicDelegatedPrefix resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputePublicDelegatedPrefixRef{})
}

func (r *ComputePublicDelegatedPrefixRef) GetGVK() schema.GroupVersionKind {
	return ComputePublicDelegatedPrefixGVK
}

func (r *ComputePublicDelegatedPrefixRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputePublicDelegatedPrefixRef) GetExternal() string {
	return r.External
}

func (r *ComputePublicDelegatedPrefixRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputePublicDelegatedPrefixRef) ValidateExternal(ref string) error {
	trimmedRef := apirefs.TrimComputeURIPrefix(ref)
	id := &ComputePublicDelegatedPrefixIdentity{}
	if err := id.FromExternal(trimmedRef); err != nil {
		return err
	}
	return nil
}

func (r *ComputePublicDelegatedPrefixRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputePublicDelegatedPrefixIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputePublicDelegatedPrefixRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" {
		r.External = apirefs.TrimComputeURIPrefix(r.External)
	}

	// Since ComputePublicDelegatedPrefix is not yet managed by KCC, we do not have a fallback
	// for obtaining the external reference from a KCC resource's status.
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, nil)
}

var (
	_ identity.IdentityV2 = &ComputePublicDelegatedPrefixIdentity{}
)

var ComputeGlobalPublicDelegatedPrefixIdentityFormat = gcpurls.Template[ComputePublicDelegatedPrefixIdentity]("compute.googleapis.com", "projects/{project}/global/publicDelegatedPrefixes/{publicDelegatedPrefix}")
var ComputeRegionalPublicDelegatedPrefixIdentityFormat = gcpurls.Template[ComputePublicDelegatedPrefixIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/publicDelegatedPrefixes/{publicDelegatedPrefix}")

var ComputeGlobalPublicDelegatedPrefixIdentityPartialFormat = gcpurls.Template[ComputePublicDelegatedPrefixIdentity]("compute.googleapis.com", "global/publicDelegatedPrefixes/{publicDelegatedPrefix}")
var ComputeRegionalPublicDelegatedPrefixIdentityPartialFormat = gcpurls.Template[ComputePublicDelegatedPrefixIdentity]("compute.googleapis.com", "regions/{region}/publicDelegatedPrefixes/{publicDelegatedPrefix}")

// ComputePublicDelegatedPrefixIdentity is the identity of a GCP ComputePublicDelegatedPrefix resource.
// +k8s:deepcopy-gen=false
type ComputePublicDelegatedPrefixIdentity struct {
	Project               string
	Region                string
	PublicDelegatedPrefix string
}

func (i *ComputePublicDelegatedPrefixIdentity) IsGlobal() bool {
	return i.Region == "" || i.Region == "global"
}

func (i *ComputePublicDelegatedPrefixIdentity) String() string {
	if !i.IsGlobal() {
		return ComputeRegionalPublicDelegatedPrefixIdentityFormat.ToString(*i)
	}
	return ComputeGlobalPublicDelegatedPrefixIdentityFormat.ToString(*i)
}

func (i *ComputePublicDelegatedPrefixIdentity) FromExternal(ref string) error {
	ref = apirefs.TrimComputeURIPrefix(ref)

	if parsed, match, _ := ComputeGlobalPublicDelegatedPrefixIdentityFormat.Parse(ref); match {
		*i = *parsed
		i.Region = "global"
		return nil
	}
	if parsed, match, _ := ComputeRegionalPublicDelegatedPrefixIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	if parsed, match, _ := ComputeGlobalPublicDelegatedPrefixIdentityPartialFormat.Parse(ref); match {
		*i = *parsed
		i.Region = "global"
		return nil
	}
	if parsed, match, _ := ComputeRegionalPublicDelegatedPrefixIdentityPartialFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputePublicDelegatedPrefix external=%q was not known (use %s or %s)", ref, ComputeGlobalPublicDelegatedPrefixIdentityFormat.CanonicalForm(), ComputeRegionalPublicDelegatedPrefixIdentityFormat.CanonicalForm())
}

func (i *ComputePublicDelegatedPrefixIdentity) Host() string {
	return ComputeGlobalPublicDelegatedPrefixIdentityFormat.Host()
}

func (i *ComputePublicDelegatedPrefixIdentity) ParentString() string {
	if !i.IsGlobal() {
		return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
	}
	return fmt.Sprintf("projects/%s/global", i.Project)
}
