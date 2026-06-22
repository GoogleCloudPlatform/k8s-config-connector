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
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &DataprocClusterRef{}
var DataprocClusterGVK = GroupVersion.WithKind("DataprocCluster")

// DataprocClusterRef defines the resource reference to DataprocCluster, which "External" field
// holds the GCP identifier for the KRM object.
type DataprocClusterRef struct {
	// A reference to an externally managed DataprocCluster resource.
	// Should be in the format "v1/projects/{{projectID}}/regions/{{region}}/clusters/{{clusterName}}" or "projects/{{projectID}}/regions/{{region}}/clusters/{{clusterName}}".
	External string `json:"external,omitempty"`

	// The name of a DataprocCluster resource.
	Name string `json:"name,omitempty"`

	// The namespace of a DataprocCluster resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&DataprocClusterRef{})
}

func (r *DataprocClusterRef) GetGVK() schema.GroupVersionKind {
	return DataprocClusterGVK
}

func (r *DataprocClusterRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DataprocClusterRef) GetExternal() string {
	return r.External
}

func (r *DataprocClusterRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DataprocClusterRef) ValidateExternal(ref string) error {
	id := &DataprocClusterIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DataprocClusterRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DataprocClusterIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DataprocClusterRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		identity, err := getIdentityFromDataprocClusterSpec(ctx, reader, u)
		if err != nil {
			return ""
		}
		return identity.String()
	}
	return refs.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}
