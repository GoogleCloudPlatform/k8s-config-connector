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
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestBigQueryReservationCapacityCommitmentIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		wantErr bool
		want    *BigQueryReservationCapacityCommitmentIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/locations/us-central1/capacityCommitments/my-commitment",
			want: &BigQueryReservationCapacityCommitmentIdentity{
				Project:            "my-project",
				Location:           "us-central1",
				CapacityCommitment: "my-commitment",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name:    "empty commitment segment",
			ref:     "projects/my-project/locations/us-central1/capacityCommitments/",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://bigqueryreservation.googleapis.com/projects/my-project/locations/us-central1/capacityCommitments/my-commitment",
			want: &BigQueryReservationCapacityCommitmentIdentity{
				Project:            "my-project",
				Location:           "us-central1",
				CapacityCommitment: "my-commitment",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BigQueryReservationCapacityCommitmentIdentity{}
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

func TestBigQueryReservationCapacityCommitment_GetIdentity(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	tests := []struct {
		name    string
		obj     *BigQueryReservationCapacityCommitment
		wantErr bool
		want    *BigQueryReservationCapacityCommitmentIdentity
	}{
		{
			name: "GetIdentity with specified resourceID",
			obj: &BigQueryReservationCapacityCommitment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-name",
				},
				Spec: BigQueryReservationCapacityCommitmentSpec{
					ProjectRef: apirefs.ProjectRef{
						External: "my-project",
					},
					Location:   "us-central1",
					ResourceID: common.LazyPtr("my-commitment"),
				},
			},
			want: &BigQueryReservationCapacityCommitmentIdentity{
				Project:            "my-project",
				Location:           "us-central1",
				CapacityCommitment: "my-commitment",
			},
		},
		{
			name: "GetIdentity fallback to metadata.name",
			obj: &BigQueryReservationCapacityCommitment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-name",
				},
				Spec: BigQueryReservationCapacityCommitmentSpec{
					ProjectRef: apirefs.ProjectRef{
						External: "my-project",
					},
					Location: "us-central1",
				},
			},
			want: &BigQueryReservationCapacityCommitmentIdentity{
				Project:            "my-project",
				Location:           "us-central1",
				CapacityCommitment: "k8s-name",
			},
		},
		{
			name: "GetIdentity mismatch with status name",
			obj: &BigQueryReservationCapacityCommitment{
				ObjectMeta: metav1.ObjectMeta{
					Name: "k8s-name",
				},
				Spec: BigQueryReservationCapacityCommitmentSpec{
					ProjectRef: apirefs.ProjectRef{
						External: "my-project",
					},
					Location:   "us-central1",
					ResourceID: common.LazyPtr("my-commitment"),
				},
				Status: BigQueryReservationCapacityCommitmentStatus{
					Name: common.LazyPtr("projects/my-project/locations/us-central1/capacityCommitments/other-commitment"),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.obj.GetIdentity(ctx, fakeClient)
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
