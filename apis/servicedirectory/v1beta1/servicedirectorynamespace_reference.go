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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ServiceDirectoryNamespaceRef{}

// ServiceDirectoryNamespaceRef is a reference to a GCP ServiceDirectoryNamespace.
type ServiceDirectoryNamespaceRef struct {
	// A reference to an externally managed ServiceDirectoryNamespace resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/namespaces/{{namespaceID}}".
	External string `json:"external,omitempty"`

	// The name of a ServiceDirectoryNamespace resource.
	Name string `json:"name,omitempty"`

	// The namespace of a ServiceDirectoryNamespace resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&ServiceDirectoryNamespaceRef{}, &ServiceDirectoryNamespace{})
}

func (r *ServiceDirectoryNamespaceRef) GetGVK() schema.GroupVersionKind {
	return ServiceDirectoryNamespaceGVK
}

func (r *ServiceDirectoryNamespaceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ServiceDirectoryNamespaceRef) GetExternal() string {
	return r.External
}

func (r *ServiceDirectoryNamespaceRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ServiceDirectoryNamespaceRef) ValidateExternal(ref string) error {
	id := &ServiceDirectoryNamespaceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ServiceDirectoryNamespaceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ServiceDirectoryNamespaceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ServiceDirectoryNamespaceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*ServiceDirectoryNamespace](u)
		if err != nil {
			return ""
		}
		identity, err := getIdentityFromServiceDirectoryNamespaceSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
