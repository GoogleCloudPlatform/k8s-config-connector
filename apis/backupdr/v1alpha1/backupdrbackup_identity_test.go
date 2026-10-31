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
	"testing"

	backupdrv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func TestBackupIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BackupIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-source/backups/my-backup",
			want: &BackupIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				BackupVault: "my-vault",
				DataSource:  "my-source",
				Backup:      "my-backup",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://backupdr.googleapis.com/projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-source/backups/my-backup",
			want: &BackupIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				BackupVault: "my-vault",
				DataSource:  "my-source",
				Backup:      "my-backup",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BackupIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestBackupIdentity_String(t *testing.T) {
	id := &BackupIdentity{
		Project:     "my-project",
		Location:    "us-central1",
		BackupVault: "my-vault",
		DataSource:  "my-source",
		Backup:      "my-backup",
	}
	want := "projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-source/backups/my-backup"
	if got := id.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestBackupIdentity_ParentString(t *testing.T) {
	id := &BackupIdentity{
		Project:     "my-project",
		Location:    "us-central1",
		BackupVault: "my-vault",
		DataSource:  "my-source",
		Backup:      "my-backup",
	}
	want := "projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-source"
	if got := id.ParentString(); got != want {
		t.Errorf("ParentString() = %v, want %v", got, want)
	}
}

func TestDataSourceIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DataSourceIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-source",
			want: &DataSourceIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				BackupVault: "my-vault",
				DataSource:  "my-source",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://backupdr.googleapis.com/projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-source",
			want: &DataSourceIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				BackupVault: "my-vault",
				DataSource:  "my-source",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DataSourceIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestDataSourceIdentity_String(t *testing.T) {
	id := &DataSourceIdentity{
		Project:     "my-project",
		Location:    "us-central1",
		BackupVault: "my-vault",
		DataSource:  "my-source",
	}
	want := "projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-source"
	if got := id.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestGetIdentityFromBackupSpec(t *testing.T) {
	ctx := context.Background()

	obj := &BackupDRBackup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-backup",
			Namespace: "test-ns",
		},
		Spec: BackupDRBackupSpec{
			ProjectRef: &refsv1beta1.ProjectRef{
				External: "my-project",
			},
			Location: ptr.To("us-central1"),
			BackupVaultRef: &backupdrv1beta1.BackupVaultRef{
				External: "projects/my-project/locations/us-central1/backupVaults/my-vault",
			},
			DataSourceRef: &DataSourceRef{
				External: "projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-source",
			},
		},
	}

	got, err := getIdentityFromBackupSpec(ctx, nil, obj)
	if err != nil {
		t.Fatalf("getIdentityFromBackupSpec() unexpected error = %v", err)
	}

	want := &BackupIdentity{
		Project:     "my-project",
		Location:    "us-central1",
		BackupVault: "my-vault",
		DataSource:  "my-source",
		Backup:      "test-backup",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("getIdentityFromBackupSpec() mismatch (-want +got):\n%s", diff)
	}
}
