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

package v1alpha1

import (
	"context"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &BigQueryBigLakeDatabaseRef{}

// BigQueryBigLakeDatabaseRef is a reference to a BigQueryBigLakeDatabase resource.
type BigQueryBigLakeDatabaseRef struct {
	// A reference to an externally managed BigQueryBigLakeDatabase resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/catalogs/{{catalogID}}/databases/{{databaseID}}".
	External string `json:"external,omitempty"`

	// The name of a BigQueryBigLakeDatabase resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BigQueryBigLakeDatabase resource.
	Namespace string `json:"namespace,omitempty"`
}

func (r *BigQueryBigLakeDatabaseRef) GetGVK() schema.GroupVersionKind {
	return BigLakeDatabaseGVK
}

func (r *BigQueryBigLakeDatabaseRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BigQueryBigLakeDatabaseRef) GetExternal() string {
	return r.External
}

func (r *BigQueryBigLakeDatabaseRef) SetExternal(ref string) {
	r.External = ref
}

func (r *BigQueryBigLakeDatabaseRef) ValidateExternal(ref string) error {
	id := &DatabaseIdentity{}
	if err := id.FromExternal(r.GetExternal()); err != nil {
		return err
	}
	return nil
}

func (r *BigQueryBigLakeDatabaseRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}
