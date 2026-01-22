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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var BigQueryDatasetGVK = schema.GroupVersionKind{
	Group:   "bigquery.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "BigQueryDataset",
}

var BigQueryTableGVK = schema.GroupVersionKind{
	Group:   "bigquery.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "BigQueryTable",
}

func init() {
	Register(&BigQueryDatasetRef{})
	Register(&BigQueryTableRef{})
}

// BigQueryDatasetRef
type BigQueryDatasetRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

func (r *BigQueryDatasetRef) GetGVK() schema.GroupVersionKind {
	return BigQueryDatasetGVK
}

func (r *BigQueryDatasetRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BigQueryDatasetRef) GetExternal() string {
	return r.External
}

func (r *BigQueryDatasetRef) SetExternal(ref string) {
	r.External = ref
}

func (r *BigQueryDatasetRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}

func (r *BigQueryDatasetRef) ValidateExternal(ref string) error {
	id := &BigQueryDatasetIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BigQueryDatasetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BigQueryDatasetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

type BigQueryDatasetIdentity struct {
	ProjectID string
	DatasetID string
}

var _ identity.Identity = &BigQueryDatasetIdentity{}

var BigQueryDatasetFormat = gcpurls.Template[BigQueryDatasetIdentity]("bigquery.googleapis.com", "projects/{projectID}/datasets/{datasetID}")

func (id *BigQueryDatasetIdentity) Host() string {
	return BigQueryDatasetFormat.Host()
}

func (id *BigQueryDatasetIdentity) String() string {
	return BigQueryDatasetFormat.ToString(*id)
}

func (id *BigQueryDatasetIdentity) FromExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("BigQueryDataset external reference cannot be empty")
	}

	parsed, match, err := BigQueryDatasetFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryDataset external=%q was not known (use %s): %w", ref, BigQueryDatasetFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryDataset external=%q was not known (use %s)", ref, BigQueryDatasetFormat.CanonicalForm())
	}
	*id = *parsed
	return nil
}

// BigQueryTableRef
type BigQueryTableRef struct {
	External  string `json:"external,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

func (r *BigQueryTableRef) GetGVK() schema.GroupVersionKind {
	return BigQueryTableGVK
}

func (r *BigQueryTableRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BigQueryTableRef) GetExternal() string {
	return r.External
}

func (r *BigQueryTableRef) SetExternal(ref string) {
	r.External = ref
}

func (r *BigQueryTableRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return Normalize(ctx, reader, r, defaultNamespace)
}

func (r *BigQueryTableRef) ValidateExternal(ref string) error {
	id := &BigQueryTableIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BigQueryTableRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BigQueryTableIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

type BigQueryTableIdentity struct {
	ProjectID string
	DatasetID string
	TableID   string
}

var _ identity.Identity = &BigQueryTableIdentity{}

var BigQueryTableFormat = gcpurls.Template[BigQueryTableIdentity]("bigquery.googleapis.com", "projects/{projectID}/datasets/{datasetID}/tables/{tableID}")

func (id *BigQueryTableIdentity) Host() string {
	return BigQueryTableFormat.Host()
}

func (id *BigQueryTableIdentity) String() string {
	return BigQueryTableFormat.ToString(*id)
}

func (id *BigQueryTableIdentity) FromExternal(ref string) error {
	if ref == "" {
		return fmt.Errorf("BigQueryTable external reference cannot be empty")
	}

	// Workaround for some cases where we might get "datasets/..." if only partial external (though not expected for canonical)
	// But let's stick to full format.

	parsed, match, err := BigQueryTableFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BigQueryTable external=%q was not known (use %s): %w", ref, BigQueryTableFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BigQueryTable external=%q was not known (use %s)", ref, BigQueryTableFormat.CanonicalForm())
	}
	*id = *parsed
	return nil
}
