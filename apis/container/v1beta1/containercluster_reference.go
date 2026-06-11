// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ContainerClusterRef{}
var ContainerClusterGVK = GroupVersion.WithKind("ContainerCluster")

func init() {
	refsv1beta1.Register(&ContainerClusterRef{})
}

type ContainerClusterRef struct {
	// The GKE cluster. Valid formats:
	//  `projects/{projectID}/locations/{location}/clusters/{clusterID}`
	//  `projects/{projectID}/zones/{zone}/clusters/{clusterID}`
	External string `json:"external,omitempty"`

	// Name of the project resource. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	// Namespace of the project resource. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

func (r *ContainerClusterRef) GetGVK() schema.GroupVersionKind {
	return ContainerClusterGVK
}

func (r *ContainerClusterRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ContainerClusterRef) GetExternal() string {
	return r.External
}

func (r *ContainerClusterRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *ContainerClusterRef) ValidateExternal(ref string) error {
	id := &ContainerClusterIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ContainerClusterRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &ContainerClusterIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *ContainerClusterRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		id, err := getIdentityFromContainerClusterSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return id.String()
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// NormalizedExternal provision the "External" value for other resource that depends on ContainerCluster.
// If the "External" is given in the other resource's spec.ContainerClusterRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual ContainerCluster object from the cluster.
// NOTE: ContainerCluster is currently a direct resource in this PR, but other resources still use this method.
// Deprecated: Use Normalize instead.
func (r *ContainerClusterRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
