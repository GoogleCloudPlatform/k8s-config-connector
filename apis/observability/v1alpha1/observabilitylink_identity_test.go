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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestObservabilityLink_GetIdentity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		obj     *ObservabilityLink
		want    *ObservabilityLinkIdentity
		wantErr bool
	}{
		{
			name: "spec with parentRef and resourceID",
			obj: &ObservabilityLink{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-link",
					Namespace: "test-ns",
				},
				Spec: ObservabilityLinkSpec{
					ParentRef: &ObservabilityDatasetRef{
						External: "projects/my-project/locations/us-central1/buckets/my-bucket/datasets/my-dataset",
					},
					ResourceID: direct.LazyPtr("custom-link-id"),
				},
			},
			want: &ObservabilityLinkIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Bucket:   "my-bucket",
				Dataset:  "my-dataset",
				Link:     "custom-link-id",
			},
		},
		{
			name: "spec with parentRef and fallback to metadata.name",
			obj: &ObservabilityLink{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-link-name",
					Namespace: "test-ns",
				},
				Spec: ObservabilityLinkSpec{
					ParentRef: &ObservabilityDatasetRef{
						External: "projects/my-project/locations/us-central1/buckets/my-bucket/datasets/my-dataset",
					},
				},
			},
			want: &ObservabilityLinkIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Bucket:   "my-bucket",
				Dataset:  "my-dataset",
				Link:     "k8s-link-name",
			},
		},
		{
			name: "status.externalRef matches spec",
			obj: &ObservabilityLink{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-link",
					Namespace: "test-ns",
				},
				Spec: ObservabilityLinkSpec{
					ParentRef: &ObservabilityDatasetRef{
						External: "projects/my-project/locations/us-central1/buckets/my-bucket/datasets/my-dataset",
					},
				},
				Status: ObservabilityLinkStatus{
					ExternalRef: direct.LazyPtr("projects/my-project/locations/us-central1/buckets/my-bucket/datasets/my-dataset/links/my-link"),
				},
			},
			want: &ObservabilityLinkIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Bucket:   "my-bucket",
				Dataset:  "my-dataset",
				Link:     "my-link",
			},
		},
		{
			name: "status.externalRef drift error",
			obj: &ObservabilityLink{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-link",
					Namespace: "test-ns",
				},
				Spec: ObservabilityLinkSpec{
					ParentRef: &ObservabilityDatasetRef{
						External: "projects/my-project/locations/us-central1/buckets/my-bucket/datasets/my-dataset",
					},
				},
				Status: ObservabilityLinkStatus{
					ExternalRef: direct.LazyPtr("projects/my-project/locations/us-central1/buckets/other-bucket/datasets/my-dataset/links/my-link"),
				},
			},
			wantErr: true,
		},
		{
			name: "missing parentRef",
			obj: &ObservabilityLink{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-link",
					Namespace: "test-ns",
				},
				Spec: ObservabilityLinkSpec{},
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
				gotIdentity, ok := got.(*ObservabilityLinkIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned unexpected type %T", got)
				}
				if diff := cmp.Diff(tt.want, gotIdentity); diff != "" {
					t.Errorf("(-want +got):\n%s", diff)
				}
			}
		})
	}
}
