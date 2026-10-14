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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &BackupRef{}

// BackupRef is a reference to a GCP BackupDRBackup.
type BackupRef struct {
	// A reference to an externally managed BackupDRBackup resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/backupVaults/{{backupvaultID}}/dataSources/{{datasourceID}}/backups/{{backupID}}".
	External string `json:"external,omitempty"`

	// The name of a BackupDRBackup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BackupDRBackup resource.
	Namespace string `json:"namespace,omitempty"`
}

var _ refs.Ref = &DataSourceRef{}

var DataSourceGVK = GroupVersion.WithKind("BackupDRDataSource")

// DataSourceRef is a reference to a GCP BackupDRDataSource.
type DataSourceRef struct {
	// A reference to an externally managed BackupDRDataSource resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/backupVaults/{{backupvaultID}}/dataSources/{{datasourceID}}".
	External string `json:"external,omitempty"`

	/* NOTYET
	// The name of a BackupDRDataSource resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BackupDRDataSource resource.
	Namespace string `json:"namespace,omitempty"`
	*/
}

var (
	_ identity.IdentityV2 = &DataSourceIdentity{}
)

var DataSourceIdentityFormat = gcpurls.Template[DataSourceIdentity]("backupdr.googleapis.com", "projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}")

// DataSourceIdentity is the identity of a GCP BackupDRDataSource resource.
// +k8s:deepcopy-gen=false
type DataSourceIdentity struct {
	Project     string
	Location    string
	BackupVault string
	DataSource  string
}

func (i *DataSourceIdentity) String() string {
	return DataSourceIdentityFormat.ToString(*i)
}

func (i *DataSourceIdentity) Host() string {
	return DataSourceIdentityFormat.Host()
}

func (i *DataSourceIdentity) FromExternal(ref string) error {
	parsed, match, err := DataSourceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of BackupDRDataSource external=%q was not known (use %s): %w", ref, DataSourceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of BackupDRDataSource external=%q was not known (use %s)", ref, DataSourceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *DataSourceIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s/backupVaults/%s", i.Project, i.Location, i.BackupVault)
}

func init() {
	refs.Register(&BackupRef{}, &BackupDRBackup{})
	refs.Register(&DataSourceRef{})
}

func (r *BackupRef) GetGVK() schema.GroupVersionKind {
	return BackupDRBackupGVK
}

func (r *BackupRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *BackupRef) GetExternal() string {
	return r.External
}

func (r *BackupRef) SetExternal(ref string) {
	r.External = ref
	r.Name = ""
	r.Namespace = ""
}

func (r *BackupRef) ValidateExternal(ref string) error {
	id := &BackupIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *BackupRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &BackupIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *BackupRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refs.Normalize(ctx, reader, r, defaultNamespace)
}

// NormalizedExternal provisions the "External" value.
// Kept for backward compatibility with older callers.
func (r *BackupRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if err := r.Normalize(ctx, reader, otherNamespace); err != nil {
		return "", err
	}
	return r.External, nil
}

func (r *DataSourceRef) GetGVK() schema.GroupVersionKind {
	return DataSourceGVK
}

func (r *DataSourceRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{}
}

func (r *DataSourceRef) GetExternal() string {
	return r.External
}

func (r *DataSourceRef) SetExternal(ref string) {
	r.External = ref
}

func (r *DataSourceRef) ValidateExternal(ref string) error {
	id := &DataSourceIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *DataSourceRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &DataSourceIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}

func (r *DataSourceRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	if r.External == "" {
		return fmt.Errorf("external reference must be specified for %s", DataSourceGVK.Kind)
	}
	return r.ValidateExternal(r.External)
}
