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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var BackupDRDataSourceGVK = schema.GroupVersionKind{
	Group:   "backupdr.cnrm.cloud.google.com",
	Version: "v1alpha1",
	Kind:    "BackupDRDataSource",
}

var _ refs.Ref = &BackupDRDataSourceRef{}

// BackupDRDataSourceRef is a reference to a BackupDRDataSource.
type BackupDRDataSourceRef struct {
	// A reference to an externally managed BackupDRDataSource resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/backupVaults/{{backupVaultID}}/dataSources/{{dataSourceID}}".
	External string `json:"external,omitempty"`

	// The name of a BackupDRDataSource resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BackupDRDataSource resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&BackupDRDataSourceRef{}, nil)
}

func (r *BackupDRDataSourceRef) GetGVK() schema.GroupVersionKind {
	return BackupDRDataSourceGVK
}

func (r *BackupDRDataSourceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *BackupDRDataSourceRef) GetExternal() string {
	return r.External
}

func (r *BackupDRDataSourceRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *BackupDRDataSourceRef) ValidateExternal(ref string) error {
	id := &BackupDRDataSourceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BackupDRDataSourceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BackupDRDataSourceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BackupDRDataSourceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", BackupDRDataSourceGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
