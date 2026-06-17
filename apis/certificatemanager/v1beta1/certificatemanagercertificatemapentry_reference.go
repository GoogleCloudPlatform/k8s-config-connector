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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &CertificateManagerCertificateMapEntryRef{}

// CertificateManagerCertificateMapEntryRef is a reference to a GCP CertificateManagerCertificateMapEntry.
type CertificateManagerCertificateMapEntryRef struct {
	// A reference to an externally managed CertificateManagerCertificateMapEntry resource. Should be in the format "projects/{{projectID}}/locations/global/certificateMaps/{{certificateMapID}}/certificateMapEntries/{{certificateMapEntryID}}".
	External string `json:"external,omitempty"`

	// The name of a CertificateManagerCertificateMapEntry resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CertificateManagerCertificateMapEntry resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CertificateManagerCertificateMapEntryRef{})
}

func (r *CertificateManagerCertificateMapEntryRef) GetGVK() schema.GroupVersionKind {
	return CertificateManagerCertificateMapEntryGVK
}

func (r *CertificateManagerCertificateMapEntryRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CertificateManagerCertificateMapEntryRef) GetExternal() string {
	return r.External
}

func (r *CertificateManagerCertificateMapEntryRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CertificateManagerCertificateMapEntryRef) ValidateExternal(ref string) error {
	id := &CertificateManagerCertificateMapEntryIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CertificateManagerCertificateMapEntryRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CertificateManagerCertificateMapEntryIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CertificateManagerCertificateMapEntryRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*CertificateManagerCertificateMapEntry](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromCertificateManagerCertificateMapEntrySpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
