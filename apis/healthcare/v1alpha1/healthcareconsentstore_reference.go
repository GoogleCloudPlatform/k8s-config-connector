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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &HealthcareConsentStoreRef{}

// HealthcareConsentStoreRef is a reference to a GCP HealthcareConsentStore.
type HealthcareConsentStoreRef struct {
	// A reference to an externally managed HealthcareConsentStore resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/datasets/{{datasetID}}/consentStores/{{consentStoreID}}".
	External string `json:"external,omitempty"`

	// The name of a HealthcareConsentStore resource.
	Name string `json:"name,omitempty"`

	// The namespace of a HealthcareConsentStore resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&HealthcareConsentStoreRef{}, &HealthcareConsentStore{})
}

func (r *HealthcareConsentStoreRef) GetGVK() schema.GroupVersionKind {
	return HealthcareConsentStoreGVK
}

func (r *HealthcareConsentStoreRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *HealthcareConsentStoreRef) GetExternal() string {
	return r.External
}

func (r *HealthcareConsentStoreRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *HealthcareConsentStoreRef) ValidateExternal(ref string) error {
	id := &HealthcareConsentStoreIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *HealthcareConsentStoreRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &HealthcareConsentStoreIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *HealthcareConsentStoreRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*HealthcareConsentStore](u)
		if err != nil {
			return ""
		}
		ready := false
		for _, cond := range obj.Status.Conditions {
			if cond.Type == "Ready" && cond.Status == "True" {
				ready = true
				break
			}
		}
		if !ready {
			return ""
		}
		id, err := getIdentityFromHealthcareConsentStoreSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
