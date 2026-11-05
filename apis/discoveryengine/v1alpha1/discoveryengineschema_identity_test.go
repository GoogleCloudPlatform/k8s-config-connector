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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDiscoveryEngineSchemaIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *DiscoveryEngineSchemaIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/global/dataStores/my-datastore/schemas/my-schema",
			want: &DiscoveryEngineSchemaIdentity{
				Project:   "my-project",
				Location:  "global",
				Datastore: "my-datastore",
				Schema:    "my-schema",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://discoveryengine.googleapis.com/projects/my-project/locations/global/dataStores/my-datastore/schemas/my-schema",
			want: &DiscoveryEngineSchemaIdentity{
				Project:   "my-project",
				Location:  "global",
				Datastore: "my-datastore",
				Schema:    "my-schema",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &DiscoveryEngineSchemaIdentity{}
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

func TestDiscoveryEngineSchema_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name    string
		obj     *DiscoveryEngineSchema
		wantErr bool
		want    *DiscoveryEngineSchemaIdentity
	}{
		{
			name: "GetIdentity with resourceID",
			obj: &DiscoveryEngineSchema{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-schema",
				},
				Spec: DiscoveryEngineSchemaSpec{
					DataStoreRef: &DiscoveryEngineDataStoreRef{
						External: "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore",
					},
					ResourceID: common.LazyPtr("my-schema"),
				},
			},
			want: &DiscoveryEngineSchemaIdentity{
				Project:   "my-project",
				Location:  "global",
				Datastore: "my-datastore",
				Schema:    "my-schema",
			},
		},
		{
			name: "GetIdentity falling back to metadata.name",
			obj: &DiscoveryEngineSchema{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-schema",
				},
				Spec: DiscoveryEngineSchemaSpec{
					DataStoreRef: &DiscoveryEngineDataStoreRef{
						External: "projects/my-project/locations/global/collections/default_collection/dataStores/my-datastore",
					},
				},
			},
			want: &DiscoveryEngineSchemaIdentity{
				Project:   "my-project",
				Location:  "global",
				Datastore: "my-datastore",
				Schema:    "my-schema",
			},
		},
		{
			name: "missing DataStoreRef",
			obj: &DiscoveryEngineSchema{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-schema",
				},
				Spec: DiscoveryEngineSchemaSpec{},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
