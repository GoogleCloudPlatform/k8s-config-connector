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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &APIHubAPIRef{}

// APIHubAPIRef is a reference to a GCP APIHubAPI.
type APIHubAPIRef struct {
	// A reference to an externally managed APIHubAPI resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/apis/{{api}}"
	External string `json:"external,omitempty"`

	// The name of an APIHubAPI resource.
	Name string `json:"name,omitempty"`

	// The namespace of an APIHubAPI resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&APIHubAPIRef{})
}

func (r *APIHubAPIRef) GetGVK() schema.GroupVersionKind {
	return APIHubAPIGVK
}

func (r *APIHubAPIRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *APIHubAPIRef) GetExternal() string {
	return r.External
}

func (r *APIHubAPIRef) SetExternal(external string) {
	r.External = external
}

func (r *APIHubAPIRef) ValidateExternal(ref string) error {
	id := &APIHubAPIIdentity{}
	return id.FromExternal(ref)
}

func (r *APIHubAPIRef) ParseExternalToIdentity() (any, error) {
	id := &APIHubAPIIdentity{}
	err := id.FromExternal(r.External)
	return id, err
}

func (r *APIHubAPIRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*APIHubAPI](u)
		if err != nil {
			return ""
		}
		id, err := getIdentityFromAPIHubAPISpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
