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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &GDCHardwareManagementHardwareRef{}

// GDCHardwareManagementHardwareRef is a reference to a GCP GDCHardwareManagementHardware.
type GDCHardwareManagementHardwareRef struct {
	// A reference to an externally managed GDCHardwareManagementHardware resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/hardware/{{hardware}}"
	External string `json:"external,omitempty"`

	// The name of a GDCHardwareManagementHardware resource.
	Name string `json:"name,omitempty"`

	// The namespace of a GDCHardwareManagementHardware resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&GDCHardwareManagementHardwareRef{}, &GDCHardwareManagementHardware{})
	refs.Register(&GDCHardwareManagementOrderRef{})
	refs.Register(&GDCHardwareManagementSiteRef{})
}

func (r *GDCHardwareManagementHardwareRef) GetGVK() schema.GroupVersionKind {
	return GDCHardwareManagementHardwareGVK
}

func (r *GDCHardwareManagementHardwareRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *GDCHardwareManagementHardwareRef) GetExternal() string {
	return r.External
}

func (r *GDCHardwareManagementHardwareRef) SetExternal(external string) {
	r.External = external
}

func (r *GDCHardwareManagementHardwareRef) ValidateExternal(ref string) error {
	id := &GDCHardwareManagementHardwareIdentity{}
	return id.FromExternal(ref)
}

func (r *GDCHardwareManagementHardwareRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &GDCHardwareManagementHardwareIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *GDCHardwareManagementHardwareRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		structured, err := common.ToStructuredType[*GDCHardwareManagementHardware](u)
		if err != nil {
			return ""
		}
		id, err := getIdentityFromGDCHardwareManagementHardwareSpec(ctx, reader, structured)
		if err != nil {
			return ""
		}
		return id.String()
	})
}

var (
	GDCHardwareManagementOrderGVK = GroupVersion.WithKind("GDCHardwareManagementOrder")
	GDCHardwareManagementSiteGVK  = GroupVersion.WithKind("GDCHardwareManagementSite")
)

var _ refs.Ref = &GDCHardwareManagementOrderRef{}

// GDCHardwareManagementOrderRef is a reference to a GCP GDCHardwareManagementOrder.
type GDCHardwareManagementOrderRef struct {
	// A reference to an externally managed GDCHardwareManagementOrder resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/orders/{{orderID}}"
	External string `json:"external,omitempty"`

	// The name of a GDCHardwareManagementOrder resource.
	Name string `json:"name,omitempty"`

	// The namespace of a GDCHardwareManagementOrder resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *GDCHardwareManagementOrderRef) GetGVK() schema.GroupVersionKind {
	return GDCHardwareManagementOrderGVK
}

func (r *GDCHardwareManagementOrderRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *GDCHardwareManagementOrderRef) GetExternal() string {
	return r.External
}

func (r *GDCHardwareManagementOrderRef) SetExternal(external string) {
	r.External = external
}

func (r *GDCHardwareManagementOrderRef) ValidateExternal(ref string) error {
	return nil
}

func (r *GDCHardwareManagementOrderRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("must specify external on GDCHardwareManagementOrder reference (namespaced references not supported as the resource is not managed by Config Connector)")
	}
	return nil
}

var _ refs.Ref = &GDCHardwareManagementSiteRef{}

// GDCHardwareManagementSiteRef is a reference to a GCP GDCHardwareManagementSite.
type GDCHardwareManagementSiteRef struct {
	// A reference to an externally managed GDCHardwareManagementSite resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/sites/{{siteID}}"
	External string `json:"external,omitempty"`

	// The name of a GDCHardwareManagementSite resource.
	Name string `json:"name,omitempty"`

	// The namespace of a GDCHardwareManagementSite resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *GDCHardwareManagementSiteRef) GetGVK() schema.GroupVersionKind {
	return GDCHardwareManagementSiteGVK
}

func (r *GDCHardwareManagementSiteRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *GDCHardwareManagementSiteRef) GetExternal() string {
	return r.External
}

func (r *GDCHardwareManagementSiteRef) SetExternal(external string) {
	r.External = external
}

func (r *GDCHardwareManagementSiteRef) ValidateExternal(ref string) error {
	return nil
}

func (r *GDCHardwareManagementSiteRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("must specify external on GDCHardwareManagementSite reference (namespaced references not supported as the resource is not managed by Config Connector)")
	}
	return nil
}
