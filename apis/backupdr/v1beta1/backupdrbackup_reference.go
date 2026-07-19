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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refs.Ref = &BackupRef{}

// BackupRef is a reference to a GCP BackupDRBackup.
type BackupRef struct {
	// A reference to an externally managed BackupDRBackup resource. Should be in the format "projects/{{projectID}}/locations/{{location}}/backupVaults/{{backupVault}}/dataSources/{{dataSource}}/backups/{{backup}}".
	External string `json:"external,omitempty"`

	// The name of a BackupDRBackup resource.
	Name string `json:"name,omitempty"`

	// The namespace of a BackupDRBackup resource.
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	refs.Register(&BackupRef{}, &BackupDRBackup{})
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
