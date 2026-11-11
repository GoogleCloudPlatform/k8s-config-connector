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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &CesToolRef{}
var CesToolGVK = GroupVersion.WithKind("CesTool")

// CesToolRef is a reference to a GCP CesTool.
type CesToolRef struct {
	/* A reference to an externally managed CesTool resource.
	Should be in the format "projects/{{projectID}}/locations/{{location}}/apps/{{appID}}/tools/{{toolID}}". */
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a CesTool resource.
	Name string `json:"name,omitempty"`

	// The namespace of a CesTool resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

func init() {
	refsv1beta1.Register(&CesToolRef{})
}

func (r *CesToolRef) GetGVK() schema.GroupVersionKind {
	return CesToolGVK
}

func (r *CesToolRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *CesToolRef) GetExternal() string {
	return r.External
}

func (r *CesToolRef) SetExternal(ref string) {
	r.External = ref
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
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", CesToolGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
