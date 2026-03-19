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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TODO: Move this to service_types.go once it is generated/implemented
var MetastoreServiceGVK = schema.GroupVersionKind{
	Group:   "metastore.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "MetastoreService",
}

var _ refs.Ref = &MetastoreServiceRef{}
var _ refs.ExternalRef = &MetastoreServiceRef{}

// MetastoreServiceRef is a reference to a MetastoreService resource.
type MetastoreServiceRef struct {
	// A reference to an externally managed MetastoreService resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/services/{{serviceID}}".
	External string `json:"external,omitempty"`

	// The name of a MetastoreService resource.
	Name string `json:"name,omitempty"`

	// The namespace of a MetastoreService resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&MetastoreServiceRef{})
}

func (r *MetastoreServiceRef) GetGVK() schema.GroupVersionKind {
	return MetastoreServiceGVK
}

func (r *MetastoreServiceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *MetastoreServiceRef) GetExternal() string {
	return r.External
}

func (r *MetastoreServiceRef) SetExternal(ref string) {
	r.External = ref
}

func (r *MetastoreServiceRef) ValidateExternal(ref string) error {
	id := &MetastoreServiceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *MetastoreServiceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &MetastoreServiceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *MetastoreServiceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		projectID, err := refs.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		location, err := refs.GetLocation(u)
		if err != nil {
			return ""
		}
		resourceID, err := refs.GetResourceID(u)
		if err != nil {
			return ""
		}
		identity := &MetastoreServiceIdentity{
			Project:  projectID,
			Location: location,
			Service:  resourceID,
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// MetastoreServiceIdentity defines the identity for a MetastoreService resource.
type MetastoreServiceIdentity struct {
	Project  string
	Location string
	Service  string
}

var _ identity.Identity = &MetastoreServiceIdentity{}
var _ identity.IdentityV2 = &MetastoreServiceIdentity{}

var MetastoreServiceIdentityFormat = gcpurls.Template[MetastoreServiceIdentity]("metastore.googleapis.com", "projects/{project}/locations/{location}/services/{service}")

func (i *MetastoreServiceIdentity) String() string {
	return MetastoreServiceIdentityFormat.ToString(*i)
}

func (i *MetastoreServiceIdentity) FromExternal(ref string) error {
	parsed, match, err := MetastoreServiceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of MetastoreService external=%q was not known (use %s): %w", ref, MetastoreServiceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of MetastoreService external=%q was not known (use %s)", ref, MetastoreServiceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *MetastoreServiceIdentity) Host() string {
	return MetastoreServiceIdentityFormat.Host()
}
