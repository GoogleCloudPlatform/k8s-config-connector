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
)

func TestBigLakeHiveCatalogIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BigLakeHiveCatalogIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/catalogs/my-catalog",
			want: &BigLakeHiveCatalogIdentity{
				Project: "my-project",
				Catalog: "my-catalog",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "empty catalog ID",
			ref:     "projects/my-project/catalogs/",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigLakeHiveCatalogIdentity{}
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

func TestBigLakeHiveCatalog_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *BigLakeHiveCatalog
		want    *BigLakeHiveCatalogIdentity
		wantErr bool
	}{
		{
			name: "projectRef external and explicit resourceID",
			obj: &BigLakeHiveCatalog{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-catalog",
					Namespace: "default",
				},
				Spec: BigLakeHiveCatalogSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "test-project",
					},
					ResourceID: ptr.To("cat-1"),
				},
			},
			want: &BigLakeHiveCatalogIdentity{
				Project: "test-project",
				Catalog: "cat-1",
			},
			wantErr: false,
		},
		{
			name: "projectRef external and metadata.name as resourceID fallback",
			obj: &BigLakeHiveCatalog{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-catalog",
					Namespace: "default",
				},
				Spec: BigLakeHiveCatalogSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "test-project",
					},
				},
			},
			want: &BigLakeHiveCatalogIdentity{
				Project: "test-project",
				Catalog: "test-catalog",
			},
			wantErr: false,
		},
		{
			name: "spec and status drift error",
			obj: &BigLakeHiveCatalog{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-catalog",
					Namespace: "default",
				},
				Spec: BigLakeHiveCatalogSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "test-project",
					},
					ResourceID: ptr.To("cat-1"),
				},
				Status: BigLakeHiveCatalogStatus{
					ExternalRef: ptr.To("projects/different-project/catalogs/cat-1"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				id, ok := got.(*BigLakeHiveCatalogIdentity)
				if !ok {
					t.Fatalf("unexpected identity type: %T", got)
				}
				if diff := cmp.Diff(tt.want, id); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
