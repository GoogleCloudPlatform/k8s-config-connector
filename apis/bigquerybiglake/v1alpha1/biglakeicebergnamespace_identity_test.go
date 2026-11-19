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

func TestBigLakeIcebergNamespace_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *BigLakeIcebergNamespace
		want    *BigLakeIcebergNamespaceIdentity
		wantErr bool
	}{
		{
			name: "parentRef external and explicit resourceID",
			obj: &BigLakeIcebergNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-namespace",
					Namespace: "default",
				},
				Spec: BigLakeIcebergNamespaceSpec{
					ParentRef: &BigLakeIcebergCatalogRef{
						External: "projects/test-project/catalogs/test-catalog",
					},
					ResourceID: ptr.To("ns-1"),
				},
			},
			want: &BigLakeIcebergNamespaceIdentity{
				Project:   "test-project",
				Catalog:   "test-catalog",
				Namespace: "ns-1",
			},
			wantErr: false,
		},
		{
			name: "missing parentRef",
			obj: &BigLakeIcebergNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-namespace",
					Namespace: "default",
				},
				Spec: BigLakeIcebergNamespaceSpec{
					ResourceID: ptr.To("ns-1"),
				},
			},
			wantErr: true,
		},
		{
			name: "server generated id resolved from status.externalRef",
			obj: &BigLakeIcebergNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-namespace",
					Namespace: "default",
				},
				Spec: BigLakeIcebergNamespaceSpec{
					ParentRef: &BigLakeIcebergCatalogRef{
						External: "projects/test-project/catalogs/test-catalog",
					},
				},
				Status: BigLakeIcebergNamespaceStatus{
					ExternalRef: ptr.To("projects/test-project/catalogs/test-catalog/namespaces/server-ns"),
				},
			},
			want: &BigLakeIcebergNamespaceIdentity{
				Project:   "test-project",
				Catalog:   "test-catalog",
				Namespace: "server-ns",
			},
			wantErr: false,
		},
		{
			name: "spec and status drift error",
			obj: &BigLakeIcebergNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-namespace",
					Namespace: "default",
				},
				Spec: BigLakeIcebergNamespaceSpec{
					ParentRef: &BigLakeIcebergCatalogRef{
						External: "projects/test-project/catalogs/test-catalog",
					},
					ResourceID: ptr.To("spec-ns"),
				},
				Status: BigLakeIcebergNamespaceStatus{
					ExternalRef: ptr.To("projects/test-project/catalogs/test-catalog/namespaces/status-ns"),
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
