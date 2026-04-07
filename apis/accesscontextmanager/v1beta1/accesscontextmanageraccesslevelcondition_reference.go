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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &AccessContextManagerAccessLevelConditionRef{}

// AccessContextManagerAccessLevelConditionRef is a reference to an AccessContextManagerAccessLevelCondition resource.
type AccessContextManagerAccessLevelConditionRef struct {
	// A reference to an externally managed AccessContextManagerAccessLevelCondition resource.
	// Should be in the format "accessPolicies/{accessPolicy}/accessLevels/{accessLevel}/condition".
	External string `json:"external,omitempty"`

	// The name of an AccessContextManagerAccessLevelCondition resource.
	Name string `json:"name,omitempty"`

	// The namespace of an AccessContextManagerAccessLevelCondition resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&AccessContextManagerAccessLevelConditionRef{})
}

func (r *AccessContextManagerAccessLevelConditionRef) GetGVK() schema.GroupVersionKind {
	return AccessContextManagerAccessLevelConditionGVK
}

func (r *AccessContextManagerAccessLevelConditionRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *AccessContextManagerAccessLevelConditionRef) GetExternal() string {
	return r.External
}

func (r *AccessContextManagerAccessLevelConditionRef) SetExternal(ref string) {
	r.External = ref
}

func (r *AccessContextManagerAccessLevelConditionRef) ValidateExternal(ref string) error {
	id := &AccessContextManagerAccessLevelConditionIdentity{}
	return id.FromExternal(ref)
}

func (r *AccessContextManagerAccessLevelConditionRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &AccessContextManagerAccessLevelConditionIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *AccessContextManagerAccessLevelConditionRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// Condition doesn't have a direct name field in GCP, but we can try to get it from externalRef
		if externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef"); externalRef != "" {
			return externalRef
		}
		return ""
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
