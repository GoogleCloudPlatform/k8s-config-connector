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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var CesAppVersionGVK = schema.GroupVersionKind{
	Group:   "ces.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "CesAppVersion",
}

var _ refs.Ref = &CesAppVersionRef{}

// CesAppVersionRef is a reference to a GCP CesAppVersion.
type CesAppVersionRef struct {
	// A reference to an externally managed CesAppVersion resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}/versions/{{versionID}}".
	External string `json:"external,omitempty"`

	// The name of a CesAppVersion resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CesAppVersion resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&CesAppVersionRef{})
}

func (r *CesAppVersionRef) GetGVK() schema.GroupVersionKind {
	return CesAppVersionGVK
}

func (r *CesAppVersionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *CesAppVersionRef) GetExternal() string {
	return r.External
}

func (r *CesAppVersionRef) SetExternal(ref string) {
	r.External = ref
}

func (r *CesAppVersionRef) ValidateExternal(ref string) error {
	id := &CesAppVersionIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *CesAppVersionRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &CesAppVersionIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *CesAppVersionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", CesAppVersionGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
