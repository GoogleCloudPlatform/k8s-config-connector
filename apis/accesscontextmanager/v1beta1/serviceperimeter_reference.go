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

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &AccessContextManagerServicePerimeterRef{}

// AccessContextManagerServicePerimeterRef defines the resource reference to AccessContextManagerServicePerimeter, which "External" field
// holds the GCP identifier for the KRM object.
type AccessContextManagerServicePerimeterRef struct {
	// A reference to an externally managed AccessContextManagerServicePerimeter resource.
	// Should be in the format "accessPolicies/{{accessPolicyID}}/servicePerimeters/{{servicePerimeter}}".
	External *string `json:"external,omitempty"`

	// The name of a AccessContextManagerServicePerimeter resource.
	Name *string `json:"name,omitempty"`

	// The namespace of a AccessContextManagerServicePerimeter resource.
	Namespace *string `json:"namespace,omitempty"`
}

func (r *AccessContextManagerServicePerimeterRef) GetGVK() schema.GroupVersionKind {
	return AccessContextManagerServicePerimeterGVK
}

func (r *AccessContextManagerServicePerimeterRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      direct.ValueOf(r.Name),
		Namespace: direct.ValueOf(r.Namespace),
	}
}

func (r *AccessContextManagerServicePerimeterRef) GetExternal() string {
	return direct.ValueOf(r.External)
}

func (r *AccessContextManagerServicePerimeterRef) SetExternal(ref string) {
	r.External = direct.LazyPtr(ref)
}

func (r *AccessContextManagerServicePerimeterRef) ValidateExternal(ref string) error {
	id := &ServicePerimeterIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *AccessContextManagerServicePerimeterRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		actualExternalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
		return actualExternalRef
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

func init() {
	refs.Register(&AccessContextManagerServicePerimeterRef{})
}
