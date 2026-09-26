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

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
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
			ref:  "projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-datasource/backups/my-backup",
			want: &BackupIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				BackupVault: "my-vault",
				DataSource:  "my-datasource",
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
			ref:  "https://backupdr.googleapis.com/projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-datasource/backups/my-backup",
			want: &BackupIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				BackupVault: "my-vault",
				DataSource:  "my-datasource",
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
		DataSource:  "my-datasource",
		Backup:      "my-backup",
	}
	want := "projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-datasource/backups/my-backup"
	if got := id.String(); got != want {
		t.Errorf("String() = %v, want %v", got, want)
	}
}

func TestBackupIdentity_ParentString(t *testing.T) {
	id := &BackupIdentity{
		Project:     "my-project",
		Location:    "us-central1",
		BackupVault: "my-vault",
		DataSource:  "my-datasource",
		Backup:      "my-backup",
	}
	want := "projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-datasource"
	if got := id.ParentString(); got != want {
		t.Errorf("ParentString() = %v, want %v", got, want)
	}
}

func TestBackupDRBackup_GetIdentity(t *testing.T) {
	ctx := context.Background()
	reader := fake.NewClientBuilder().Build()

	tests := []struct {
		name    string
		obj     *BackupDRBackup
		want    *BackupIdentity
		wantErr bool
	}{
		{
			name: "valid resource from spec",
			obj: &BackupDRBackup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-backup",
					Namespace: "default",
				},
				Spec: BackupDRBackupSpec{
					ProjectRef:  &refsv1beta1.ProjectRef{External: "my-project"},
					Location:    ptr.To("us-central1"),
					BackupVault: ptr.To("my-vault"),
					DataSource:  ptr.To("my-datasource"),
				},
			},
			want: &BackupIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				BackupVault: "my-vault",
				DataSource:  "my-datasource",
				Backup:      "my-backup",
			},
		},
		{
			name: "valid resource with matching externalRef in status",
			obj: &BackupDRBackup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-backup",
					Namespace: "default",
				},
				Spec: BackupDRBackupSpec{
					ProjectRef:  &refsv1beta1.ProjectRef{External: "my-project"},
					Location:    ptr.To("us-central1"),
					BackupVault: ptr.To("my-vault"),
					DataSource:  ptr.To("my-datasource"),
				},
				Status: BackupDRBackupStatus{
					ExternalRef: ptr.To("projects/my-project/locations/us-central1/backupVaults/my-vault/dataSources/my-datasource/backups/my-backup"),
				},
			},
			want: &BackupIdentity{
				Project:     "my-project",
				Location:    "us-central1",
				BackupVault: "my-vault",
				DataSource:  "my-datasource",
				Backup:      "my-backup",
			},
		},
		{
			name: "mismatching externalRef in status",
			obj: &BackupDRBackup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-backup",
					Namespace: "default",
				},
				Spec: BackupDRBackupSpec{
					ProjectRef:  &refsv1beta1.ProjectRef{External: "my-project"},
					Location:    ptr.To("us-central1"),
					BackupVault: ptr.To("my-vault"),
					DataSource:  ptr.To("my-datasource"),
				},
				Status: BackupDRBackupStatus{
					ExternalRef: ptr.To("projects/other-project/locations/us-central1/backupVaults/my-vault/dataSources/my-datasource/backups/my-backup"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing backupVault in spec",
			obj: &BackupDRBackup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-backup",
					Namespace: "default",
				},
				Spec: BackupDRBackupSpec{
					ProjectRef: &refsv1beta1.ProjectRef{External: "my-project"},
					Location:   ptr.To("us-central1"),
					DataSource: ptr.To("my-datasource"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing dataSource in spec",
			obj: &BackupDRBackup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-backup",
					Namespace: "default",
				},
				Spec: BackupDRBackupSpec{
					ProjectRef:  &refsv1beta1.ProjectRef{External: "my-project"},
					Location:    ptr.To("us-central1"),
					BackupVault: ptr.To("my-vault"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := tt.obj.GetIdentity(ctx, reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got, ok := id.(*BackupIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned type %T, want *BackupIdentity", id)
				}
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
