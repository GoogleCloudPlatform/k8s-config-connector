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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestBigQueryConnectionConnectionIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name                 string
		ref                  string
		wantErr              bool
		want                 *BigQueryConnectionConnectionIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/connections/my-connection",
			want: &BigQueryConnectionConnectionIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Connection: "my-connection",
			},
			wantIdentitySpecified: true,
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://bigqueryconnection.googleapis.com/projects/my-project/locations/us-central1/connections/my-connection",
			want: &BigQueryConnectionConnectionIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Connection: "my-connection",
			},
			wantIdentitySpecified: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigQueryConnectionConnectionIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if got := i.HasIdentitySpecified(); got != tt.wantIdentitySpecified {
					t.Errorf("HasIdentitySpecified() = %v, want %v", got, tt.wantIdentitySpecified)
				}
			}
		})
	}
}

func TestBigQueryConnectionConnection_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name                 string
		obj                  *BigQueryConnectionConnection
		wantErr              bool
		want                 *BigQueryConnectionConnectionIdentity
		wantIdentitySpecified bool
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &BigQueryConnectionConnection{
				Spec: BigQueryConnectionConnectionSpec{
					Parent: Parent{
						ProjectRef: &refs.ProjectRef{
							External: "my-project",
						},
						Location: "us-central1",
					},
					ResourceID: common.LazyPtr("my-connection"),
				},
			},
			want: &BigQueryConnectionConnectionIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Connection: "my-connection",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with server-generated identity (empty spec resourceID, defaulted from status)",
			obj: &BigQueryConnectionConnection{
				Spec: BigQueryConnectionConnectionSpec{
					Parent: Parent{
						ProjectRef: &refs.ProjectRef{
							External: "my-project",
						},
						Location: "us-central1",
					},
				},
				Status: BigQueryConnectionConnectionStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/connections/server-generated-id"),
				},
			},
			want: &BigQueryConnectionConnectionIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Connection: "server-generated-id",
			},
			wantIdentitySpecified: true,
		},
		{
			name: "GetIdentity with no spec resourceID and empty status (not yet created)",
			obj: &BigQueryConnectionConnection{
				Spec: BigQueryConnectionConnectionSpec{
					Parent: Parent{
						ProjectRef: &refs.ProjectRef{
							External: "my-project",
						},
						Location: "us-central1",
					},
				},
			},
			want: &BigQueryConnectionConnectionIdentity{
				Project:    "my-project",
				Location:   "us-central1",
				Connection: "",
			},
			wantIdentitySpecified: false,
		},
		{
			name: "GetIdentity with server-generated identity conflict with spec resourceID",
			obj: &BigQueryConnectionConnection{
				Spec: BigQueryConnectionConnectionSpec{
					Parent: Parent{
						ProjectRef: &refs.ProjectRef{
							External: "my-project",
						},
						Location: "us-central1",
					},
					ResourceID: common.LazyPtr("conflict-id"),
				},
				Status: BigQueryConnectionConnectionStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/connections/server-generated-id"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIdentity, err := tt.obj.GetIdentity(ctx, fakeClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				got, ok := gotIdentity.(*BigQueryConnectionConnectionIdentity)
				if !ok {
					t.Fatalf("returned identity is not *BigQueryConnectionConnectionIdentity, got %T", gotIdentity)
				}
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
				if gotSpecified := got.HasIdentitySpecified(); gotSpecified != tt.wantIdentitySpecified {
					t.Errorf("got.HasIdentitySpecified() = %v, want %v", gotSpecified, tt.wantIdentitySpecified)
				}
			}
		})
	}
}
