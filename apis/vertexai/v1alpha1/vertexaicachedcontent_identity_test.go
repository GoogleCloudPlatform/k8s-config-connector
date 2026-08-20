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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestVertexAICachedContentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *VertexAICachedContentIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/cachedContents/my-cachedcontent",
			want: &VertexAICachedContentIdentity{
				Project:       "my-project",
				Location:      "us-central1",
				CachedContent: "my-cachedcontent",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://aiplatform.googleapis.com/projects/my-project/locations/us-central1/cachedContents/my-cachedcontent",
			want: &VertexAICachedContentIdentity{
				Project:       "my-project",
				Location:      "us-central1",
				CachedContent: "my-cachedcontent",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &VertexAICachedContentIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, i); diff != "" {
					t.Errorf("FromExternal mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestGetIdentityFromVertexAICachedContentSpec(t *testing.T) {
	tests := []struct {
		name    string
		obj     *VertexAICachedContent
		want    *VertexAICachedContentIdentity
		wantErr bool
	}{
		{
			name: "resourceID is simple name",
			obj: &VertexAICachedContent{
				Spec: VertexAICachedContentSpec{
					Location:   "us-central1",
					ResourceID: common.LazyPtr("my-cachedcontent"),
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
			},
			want: &VertexAICachedContentIdentity{
				Project:       "my-project",
				Location:      "us-central1",
				CachedContent: "my-cachedcontent",
			},
		},
		{
			name: "resourceID is empty",
			obj: &VertexAICachedContent{
				Spec: VertexAICachedContentSpec{
					Location: "us-central1",
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
			},
			want: &VertexAICachedContentIdentity{
				Project:       "my-project",
				Location:      "us-central1",
				CachedContent: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIdentityFromVertexAICachedContentSpec(context.Background(), nil, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIdentityFromVertexAICachedContentSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("getIdentityFromVertexAICachedContentSpec() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestVertexAICachedContent_GetIdentity(t *testing.T) {
	tests := []struct {
		name    string
		obj     *VertexAICachedContent
		want    *VertexAICachedContentIdentity
		wantErr bool
	}{
		{
			name: "resourceID is nil, externalRef exists",
			obj: &VertexAICachedContent{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-cachedcontent-krm",
				},
				Spec: VertexAICachedContentSpec{
					Location: "us-central1",
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "my-project",
					},
				},
				Status: VertexAICachedContentStatus{
					ExternalRef: common.LazyPtr("projects/my-project/locations/us-central1/cachedContents/123456789"),
				},
			},
			want: &VertexAICachedContentIdentity{
				Project:       "my-project",
				Location:      "us-central1",
				CachedContent: "123456789",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(context.Background(), nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
