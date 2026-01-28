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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	refsv1beta1.Register(&DatasetRef{})
}

var _ refsv1beta1.Ref = &DatasetRef{}
var _ refsv1beta1.ExternalRef = &DatasetRef{}

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

// GetGVK returns the GroupVersionKind for BigQueryDataset.
func (r *DatasetRef) GetGVK() schema.GroupVersionKind {
	return BigQueryDatasetGVK
}

// GetNamespacedName returns the NamespacedName for the reference.
func (r *DatasetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

// GetExternal returns the external reference.
func (r *DatasetRef) GetExternal() string {
	return r.External
}

func (r *DatasetRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DatasetRef) ValidateExternal(ref string) error {
	id := &DatasetIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DatasetRef) Normalize(ctx context.Context, reader client.Reader, otherNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// BigQueryDataset is still TF based so we resolve the reference from the object
		// BUT only construct the reference if the resource is ready
		resource, err := k8s.NewResource(u)
		if err != nil {
			return ""
		}
		if !k8s.IsResourceReady(resource) {
			return ""
		}

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
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, otherNamespace, fallback)
}

// NormalizedExternal is a helper function to resolve the external reference.
// Deprecated: Use Normalize instead.
func (r *DatasetRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

func (r *DatasetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DatasetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}
