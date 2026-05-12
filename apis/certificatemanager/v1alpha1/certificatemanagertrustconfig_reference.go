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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &CertificateManagerTrustConfigRef{}

// CertificateManagerTrustConfigRef defines the resource reference to CertificateManagerTrustConfig, which "External" field
// holds the GCP identifier for the KRM object.
type CertificateManagerTrustConfigRef struct {
	// A reference to an externally managed CertificateManagerTrustConfig resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/trustConfigs/{{trustConfigID}}".
	External string `json:"external,omitempty"`

	// The name of a CertificateManagerTrustConfig resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CertificateManagerTrustConfig resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CertificateManagerTrustConfigRef{})
}

func (r *CertificateManagerTrustConfigRef) GetGVK() schema.GroupVersionKind {
	return CertificateManagerTrustConfigGVK
}

func (r *CertificateManagerTrustConfigRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CertificateManagerTrustConfigRef) GetExternal() string {
	return r.External
}

func (r *CertificateManagerTrustConfigRef) SetExternal(ref string) {
	r.External = ref
}

func (r *CertificateManagerTrustConfigRef) ValidateExternal(ref string) error {
	id := &CertificateManagerTrustConfigIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CertificateManagerTrustConfigRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CertificateManagerTrustConfigIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CertificateManagerTrustConfigRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromCertificateManagerTrustConfigSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
