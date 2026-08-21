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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &NetworkServicesEdgeCacheServiceRef{}

// NetworkServicesEdgeCacheServiceRef is a reference to a GCP NetworkServicesEdgeCacheService.
type NetworkServicesEdgeCacheServiceRef struct {
	// A reference to an externally managed NetworkServicesEdgeCacheService resource.
	// Should be in the format "projects/{{projectID}}/locations/global/edgeCacheServices/{{edgeCacheServiceID}}".
	External string `json:"external,omitempty"`

	// The name of a NetworkServicesEdgeCacheService resource.
	Name string `json:"name,omitempty"`

	// The namespace of a NetworkServicesEdgeCacheService resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refsv1beta1.Register(&NetworkServicesEdgeCacheServiceRef{}, &NetworkServicesEdgeCacheService{})
}

func (r *NetworkServicesEdgeCacheServiceRef) GetGVK() schema.GroupVersionKind {
	return NetworkServicesEdgeCacheServiceGVK
}

func (r *NetworkServicesEdgeCacheServiceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *NetworkServicesEdgeCacheServiceRef) GetExternal() string {
	return r.External
}

func (r *NetworkServicesEdgeCacheServiceRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *NetworkServicesEdgeCacheServiceRef) ValidateExternal(ref string) error {
	id := &NetworkServicesEdgeCacheServiceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *NetworkServicesEdgeCacheServiceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &NetworkServicesEdgeCacheServiceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *NetworkServicesEdgeCacheServiceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External != "" && r.Name != "" {
		return fmt.Errorf("cannot specify both name and external on %s reference", r.GetGVK().Kind)
	}
	fallback := func(u *unstructured.Unstructured) string {
		obj, err := common.ToStructuredType[*NetworkServicesEdgeCacheService](u)
		if err != nil {
			return ""
		}
		identity, err := NewNetworkServicesEdgeCacheServiceIdentity(ctx, reader, obj)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
