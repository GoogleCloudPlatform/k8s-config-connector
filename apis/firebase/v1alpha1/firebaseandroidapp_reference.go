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

var _ refs.Ref = &FirebaseAndroidAppRef{}

// FirebaseAndroidAppRef is a reference to a GCP FirebaseAndroidApp.
type FirebaseAndroidAppRef struct {
	// A reference to an externally managed FirebaseAndroidApp resource.
	// Should be in the format "projects/{{projectID}}/androidApps/{{appID}}".
	External string `json:"external,omitempty"`

	// The name of a FirebaseAndroidApp resource.
	Name string `json:"name,omitempty"`

	// The namespace of a FirebaseAndroidApp resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&FirebaseAndroidAppRef{}, &FirebaseAndroidApp{})
}

func (r *FirebaseAndroidAppRef) GetGVK() schema.GroupVersionKind {
	return FirebaseAndroidAppGVK
}

func (r *FirebaseAndroidAppRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *FirebaseAndroidAppRef) GetExternal() string {
	return r.External
}

func (r *FirebaseAndroidAppRef) SetExternal(ref string) {
	r.External = ref
}

func (r *FirebaseAndroidAppRef) ValidateExternal(ref string) error {
	id := &FirebaseAndroidAppIdentity{}
	return id.FromExternal(ref)
}

func (r *FirebaseAndroidAppRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &FirebaseAndroidAppIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *FirebaseAndroidAppRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*FirebaseAndroidApp](u)
		if err != nil {
			return ""
		}
		id, err := obj.GetIdentity(ctx, reader)
		if err != nil || id == nil {
			return ""
		}
		return id.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
