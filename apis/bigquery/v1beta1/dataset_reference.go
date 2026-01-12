// Copyright 2024 Google LLC
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
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &DatasetRef{}

// DatasetRef defines the resource reference to BigQueryDataset, which "External" field
// holds the GCP identifier for the KRM object.
type DatasetRef struct {
	// A reference to an externally-managed BigQueryDataset resource.
	// Should be in the format "projects/{{projectID}}/datasets/{{datasetID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryDataset resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryDataset resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *DatasetRef) GetGVK() schema.GroupVersionKind {
	return BigQueryDatasetGVK
}

func (r *DatasetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *DatasetRef) GetExternal() string {
	return r.External
}

func (r *DatasetRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DatasetRef) ValidateExternal(ref string) error {
	id := &DatasetIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *DatasetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return ""
		}
		datasetID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("projects/%s/datasets/%s", projectID, datasetID)
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}

// NormalizedExternal provision the "External" value for other resource that depends on BigQueryDataset.
// If the "External" is given in the other resource's spec.BigQueryDatasetRef, the given value will be used.
// Otherwise, the "Name" and "Namespace" will be used to query the actual BigQueryDataset object from the cluster.
func (r *DatasetRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}
