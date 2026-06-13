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

var _ refs.Ref = &ComputeRouterRef{}

// ComputeRouterRef is a reference to a ComputeRouter.
type ComputeRouterRef struct {
	// A reference to an externally managed ComputeRouter resource.
	// Should be in the format "projects/{{projectID}}/regions/{{region}}/routers/{{routerID}}".
	External string `json:"external,omitempty"`

	// The name of a ComputeRouter resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ComputeRouter resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ComputeRouterRef{})
}

func (r *ComputeRouterRef) GetGVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeRouter",
	}
}

func (r *ComputeRouterRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeRouterRef) GetExternal() string {
	return r.External
}

func (r *ComputeRouterRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ComputeRouterRef) ValidateExternal(ref string) error {
	id := &ComputeRouterIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeRouterRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ComputeRouterIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ComputeRouterRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
		if selfLink != "" {
			return refs.TrimComputeURIPrefix(selfLink)
		}
		return ""
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
