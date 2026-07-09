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

var _ refs.ExternalRef = &NetworkSecurityUrlListRef{}

// NetworkSecurityUrlListRef is a reference to a GCP NetworkSecurityUrlList.
type NetworkSecurityUrlListRef struct {
	// A reference to an externally managed NetworkSecurityUrlList resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/urlLists/{{url_list}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkSecurityUrlList resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkSecurityUrlList resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&NetworkSecurityUrlListRef{})
}

func (r *NetworkSecurityUrlListRef) GetGVK() schema.GroupVersionKind {
	return NetworkSecurityUrlListGVK
}

func (r *NetworkSecurityUrlListRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkSecurityUrlListRef) GetExternal() string {
	return r.External
}

func (r *NetworkSecurityUrlListRef) SetExternal(external string) {
	r.External = external
	r.Name = ""
	r.Namespace = ""
}

func (r *NetworkSecurityUrlListRef) ValidateExternal(external string) error {
	id := &NetworkSecurityUrlListIdentity{}
	return id.FromExternal(external)
}

func (r *NetworkSecurityUrlListRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkSecurityUrlListIdentity{}
	err := id.FromExternal(r.External)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkSecurityUrlListRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, otherNamespace, func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*NetworkSecurityUrlList](u)
		if err != nil {
			return ""
		}
		id, err := getIdentityFromNetworkSecurityUrlListSpec(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return id.String()
	})
}
