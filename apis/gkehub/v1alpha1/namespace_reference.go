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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &GKEHubNamespaceRef{}

// GKEHubNamespaceRef is a reference to a GKEHubNamespace.
type GKEHubNamespaceRef struct {
	// A reference to an externally managed GKEHubNamespace resource.
	// Should be in the format "projects/{{project}}/locations/{{location}}/scopes/{{scope}}/namespaces/{{namespace}}".
	External string `json:"external,omitempty"`

	// The name of a GKEHubNamespace resource.
	Name string `json:"name,omitempty"`

	// The namespace of a GKEHubNamespace resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&GKEHubNamespaceRef{})
}

func (r *GKEHubNamespaceRef) GetGVK() schema.GroupVersionKind {
	return GKEHubNamespaceGVK
}

func (r *GKEHubNamespaceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *GKEHubNamespaceRef) GetExternal() string {
	return r.External
}

func (r *GKEHubNamespaceRef) SetExternal(ref string) {
	r.External = ref
}

func (r *GKEHubNamespaceRef) ValidateExternal(ref string) error {
	id := &GKEHubNamespaceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *GKEHubNamespaceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &GKEHubNamespaceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *GKEHubNamespaceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj := &GKEHubNamespace{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
			return ""
		}
		identity, err := getIdentityFromGKEHubNamespaceSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

func ResolveGKEHubNamespaceRef(ctx context.Context, reader client.Reader, obj client.Object, ref *GKEHubNamespaceRef) (*GKEHubNamespaceIdentity, error) {
	if ref == nil {
		return nil, nil
	}
	if err := ref.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	id := &GKEHubNamespaceIdentity{}
	if err := id.FromExternal(ref.External); err != nil {
		return nil, err
	}
	return id, nil
}
