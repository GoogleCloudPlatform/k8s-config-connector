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

func TestBigLakeHiveTableIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BigLakeHiveTableIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/catalogs/my-catalog/databases/my-database/tables/my-table",
			want: &BigLakeHiveTableIdentity{
				Project:  "my-project",
				Catalog:  "my-catalog",
				Database: "my-database",
				Table:    "my-table",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "empty table ID",
			ref:     "projects/my-project/catalogs/my-catalog/databases/my-database/tables/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigLakeHiveTableIdentity{}
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

func TestBigLakeHiveTable_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *BigLakeHiveTable
		want    *BigLakeHiveTableIdentity
		wantErr bool
	}{
		{
			name: "parentRef external and explicit resourceID",
			obj: &BigLakeHiveTable{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-table",
					Namespace: "default",
				},
				Spec: BigLakeHiveTableSpec{
					ParentRef: &BigLakeHiveDatabaseRef{
						External: "projects/test-project/catalogs/test-catalog/databases/test-database",
					},
					ResourceID: ptr.To("table-1"),
				},
			},
			want: &BigLakeHiveTableIdentity{
				Project:  "test-project",
				Catalog:  "test-catalog",
				Database: "test-database",
				Table:    "table-1",
			},
			wantErr: false,
		},
		{
			name: "parentRef external and metadata.name as resourceID fallback",
			obj: &BigLakeHiveTable{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-table",
					Namespace: "default",
				},
				Spec: BigLakeHiveTableSpec{
					ParentRef: &BigLakeHiveDatabaseRef{
						External: "projects/test-project/catalogs/test-catalog/databases/test-database",
					},
				},
			},
			want: &BigLakeHiveTableIdentity{
				Project:  "test-project",
				Catalog:  "test-catalog",
				Database: "test-database",
				Table:    "test-table",
			},
			wantErr: false,
		},
		{
			name: "missing parentRef",
			obj: &BigLakeHiveTable{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-table",
					Namespace: "default",
				},
				Spec: BigLakeHiveTableSpec{
					ResourceID: ptr.To("table-1"),
				},
			},
			wantErr: true,
		},
		{
			name: "spec and status drift error",
			obj: &BigLakeHiveTable{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-table",
					Namespace: "default",
				},
				Spec: BigLakeHiveTableSpec{
					ParentRef: &BigLakeHiveDatabaseRef{
						External: "projects/test-project/catalogs/test-catalog/databases/test-database",
					},
					ResourceID: ptr.To("spec-table"),
				},
				Status: BigLakeHiveTableStatus{
					ExternalRef: ptr.To("projects/test-project/catalogs/test-catalog/databases/test-database/tables/status-table"),
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
