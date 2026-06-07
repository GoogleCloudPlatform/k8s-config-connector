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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &CertificateManagerCertificateIssuanceConfigRef{}

// CertificateManagerCertificateIssuanceConfigRef is a reference to a CertificateManagerCertificateIssuanceConfig.
type CertificateManagerCertificateIssuanceConfigRef struct {
	// A reference to an externally managed CertificateManagerCertificateIssuanceConfig resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/certificateIssuanceConfigs/{{certificateIssuanceConfigID}}".
	External string `json:"external,omitempty"`

	// The name of a CertificateManagerCertificateIssuanceConfig resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CertificateManagerCertificateIssuanceConfig resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CertificateManagerCertificateIssuanceConfigRef{})
}

func (r *CertificateManagerCertificateIssuanceConfigRef) GetGVK() schema.GroupVersionKind {
	return CertificateManagerCertificateIssuanceConfigGVK
}

func (r *CertificateManagerCertificateIssuanceConfigRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *CertificateManagerCertificateIssuanceConfigRef) GetExternal() string {
	return r.External
}

func (r *CertificateManagerCertificateIssuanceConfigRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *CertificateManagerCertificateIssuanceConfigRef) ValidateExternal(ref string) error {
	id := &CertificateManagerCertificateIssuanceConfigIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CertificateManagerCertificateIssuanceConfigRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CertificateManagerCertificateIssuanceConfigIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CertificateManagerCertificateIssuanceConfigRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*CertificateManagerCertificateIssuanceConfig](u)
		if err != nil {
			return ""
		}
		// Manually populate obj.Spec.ParentRef if it was lost during conversion because of ',inline' tag on a pointer.
		if obj.Spec.ParentRef == nil || obj.Spec.ParentRef.ProjectRef == nil {
			projectRefVal, found, _ := unstructured.NestedMap(u.Object, "spec", "projectRef")
			if !found {
				projectRefVal, found, _ = unstructured.NestedMap(u.Object, "spec", "parentRef", "projectRef")
			}
			if !found {
				projectRefVal, found, _ = unstructured.NestedMap(u.Object, "spec", "ParentRef", "projectRef")
			}
			if found {
				var projRef refs.ProjectRef
				if err := runtime.DefaultUnstructuredConverter.FromUnstructured(projectRefVal, &projRef); err == nil {
					if obj.Spec.ParentRef == nil {
						obj.Spec.ParentRef = &parent.ProjectAndLocationRef{}
					}
					obj.Spec.ParentRef.ProjectRef = &projRef
				}
			}
			locationVal, found, _ := unstructured.NestedString(u.Object, "spec", "location")
			if !found {
				locationVal, found, _ = unstructured.NestedString(u.Object, "spec", "parentRef", "location")
			}
			if !found {
				locationVal, found, _ = unstructured.NestedString(u.Object, "spec", "ParentRef", "location")
			}
			if found {
				if obj.Spec.ParentRef == nil {
					obj.Spec.ParentRef = &parent.ProjectAndLocationRef{}
				}
				obj.Spec.ParentRef.Location = locationVal
			}
		}
		identity, err := getIdentityFromCertificateManagerCertificateIssuanceConfigSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
