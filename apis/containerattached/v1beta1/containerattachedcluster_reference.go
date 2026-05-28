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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &ContainerAttachedClusterRef{}

func init() {
	refs.Register(&ContainerAttachedClusterRef{})
}

type ContainerAttachedClusterRef struct {
	/* A reference to an externally managed ContainerAttachedCluster resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/attachedClusters/{{containerattachedcluster}}" */
	External string `json:"external,omitempty"`
	/* The name of a ContainerAttachedCluster resource. */
	Name string `json:"name,omitempty"`
	/* The namespace of a ContainerAttachedCluster resource. */
	Namespace string `json:"namespace,omitempty"`
}

func (r *ContainerAttachedClusterRef) GetGVK() schema.GroupVersionKind {
	return ContainerAttachedClusterGVK
}

func (r *ContainerAttachedClusterRef) GetNamespacedName() client.ObjectKey {
	return client.ObjectKey{
		Namespace: r.Namespace,
		Name:      r.Name,
	}
}

func (r *ContainerAttachedClusterRef) GetExternal() string {
	return r.External
}

func (r *ContainerAttachedClusterRef) SetExternal(external string) {
	r.External = external
}

func (r *ContainerAttachedClusterRef) ValidateExternal(external string) error {
	identity := &ContainerAttachedClusterIdentity{}
	return identity.FromExternal(external)
}

func (r *ContainerAttachedClusterRef) ParseExternalToIdentity() (identity.Identity, error) {
	identity := &ContainerAttachedClusterIdentity{}
	if err := identity.FromExternal(r.External); err != nil {
		return nil, err
	}
	return identity, nil
}

func (r *ContainerAttachedClusterRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, func(u *unstructured.Unstructured) string {
		if id, err := getIdentityFromContainerAttachedClusterSpec(ctx, reader, u); err == nil {
			return id.String()
		}
		return ""
	})
}
