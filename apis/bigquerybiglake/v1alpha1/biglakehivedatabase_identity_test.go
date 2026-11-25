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

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

func TestBigLakeHiveDatabaseIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BigLakeHiveDatabaseIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/catalogs/my-catalog/databases/my-database",
			want: &BigLakeHiveDatabaseIdentity{
				Project:  "my-project",
				Catalog:  "my-catalog",
				Database: "my-database",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "empty database ID",
			ref:     "projects/my-project/catalogs/my-catalog/databases/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigLakeHiveDatabaseIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.String(); got != tt.ref {
					t.Errorf("String() = %v, want %v", got, tt.ref)
				}
			}
		})
	}
}

func TestBigLakeHiveDatabase_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *BigLakeHiveDatabase
		want    *BigLakeHiveDatabaseIdentity
		wantErr bool
	}{
		{
			name: "parentRef external and explicit resourceID",
			obj: &BigLakeHiveDatabase{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-database",
					Namespace: "default",
				},
				Spec: BigLakeHiveDatabaseSpec{
					ParentRef: &BigLakeHiveCatalogRef{
						External: "projects/test-project/catalogs/test-catalog",
					},
					ResourceID: ptr.To("db-1"),
				},
			},
			want: &BigLakeHiveDatabaseIdentity{
				Project:  "test-project",
				Catalog:  "test-catalog",
				Database: "db-1",
			},
			wantErr: false,
		},
		{
			name: "parentRef external and metadata.name as resourceID fallback",
			obj: &BigLakeHiveDatabase{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-database",
					Namespace: "default",
				},
				Spec: BigLakeHiveDatabaseSpec{
					ParentRef: &BigLakeHiveCatalogRef{
						External: "projects/test-project/catalogs/test-catalog",
					},
				},
			},
			want: &BigLakeHiveDatabaseIdentity{
				Project:  "test-project",
				Catalog:  "test-catalog",
				Database: "test-database",
			},
			wantErr: false,
		},
		{
			name: "missing parentRef",
			obj: &BigLakeHiveDatabase{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-database",
					Namespace: "default",
				},
				Spec: BigLakeHiveDatabaseSpec{
					ResourceID: ptr.To("db-1"),
				},
			},
			wantErr: true,
		},
		{
			name: "spec and status drift error",
			obj: &BigLakeHiveDatabase{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-database",
					Namespace: "default",
				},
				Spec: BigLakeHiveDatabaseSpec{
					ParentRef: &BigLakeHiveCatalogRef{
						External: "projects/test-project/catalogs/test-catalog",
					},
					ResourceID: ptr.To("spec-db"),
				},
				Status: BigLakeHiveDatabaseStatus{
					ExternalRef: ptr.To("projects/test-project/catalogs/test-catalog/databases/status-db"),
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.obj.GetIdentity(ctx, nil)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
