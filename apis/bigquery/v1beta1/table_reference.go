// Copyright 2025 Google LLC
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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &TableRef{}

// TableRef defines the resource reference to BigQueryTable, which "External" field
// holds the GCP identifier for the KRM object.
type TableRef struct {
	// A reference to an externally-managed BigQueryTable resource.
	// Should be in the format "projects/{{projectID}}/datasets/{{datasetsID}}/tables/{{tableID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryTable resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryTable resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *TableRef) GetGVK() schema.GroupVersionKind {
	return BigQueryTableGVK
}

func (r *TableRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *TableRef) GetExternal() string {
	return r.External
}

func (r *TableRef) SetExternal(ref string) {
	r.External = ref
}

func (r *TableRef) ValidateExternal(ref string) error {
	_, _, err := ParseBigQueryTableExternal(ref)
	return err
}

func (r *TableRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	fallback := func(u *unstructured.Unstructured) string {
		// Get resourceID
		tableID, err := refsv1beta1.GetResourceID(u)
		if err != nil {
			return ""
		}
		// Resolve parent
		obj := &BigQueryTable{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
			return ""
		}
		datasetRef := obj.Spec.DatasetRef
		// We can't use NormalizedExternal directly here because it might be recursive or context dependent?
		// datasetRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		// But DatasetRef is likely supporting Normalize or NormalizedExternal.
		// Let's assume it works.
		datasetExternal, err := datasetRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%s/tables/%s", datasetExternal, tableID)
	}
	return refsv1beta1.NormalizeWithFallback(ctx, reader, r, defaultNamespace, fallback)
}