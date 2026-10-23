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

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBigQueryDatasetAccessIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name    string
		ref     string
		want    *BigQueryDatasetAccessIdentity
		wantErr bool
	}{
		{
			name: "valid external reference",
			ref:  "projects/test-project/datasets/test-dataset/access/test-access",
			want: &BigQueryDatasetAccessIdentity{
				Project: "test-project",
				Dataset: "test-dataset",
				Access:  "test-access",
			},
		},
		{
			name: "valid full URL",
			ref:  "//bigquery.googleapis.com/projects/test-project/datasets/test-dataset/access/test-access",
			want: &BigQueryDatasetAccessIdentity{
				Project: "test-project",
				Dataset: "test-dataset",
				Access:  "test-access",
			},
		},
		{
			name:    "invalid format missing segments",
			ref:     "projects/test-project/datasets/test-dataset",
			wantErr: true,
		},
		{
			name:    "invalid format wrong collection",
			ref:     "projects/test-project/datasets/test-dataset/tables/test-table",
			wantErr: true,
		},
		{
			name:    "empty reference",
			ref:     "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := &BigQueryDatasetAccessIdentity{}
			err := got.FromExternal(tc.ref)
			if (err != nil) != tc.wantErr {
				t.Fatalf("FromExternal() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !tc.wantErr {
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
				}
				if gotStr := got.String(); gotStr != "projects/test-project/datasets/test-dataset/access/test-access" {
					t.Errorf("String() = %q, want %q", gotStr, "projects/test-project/datasets/test-dataset/access/test-access")
				}
				if gotParent := got.ParentString(); gotParent != "projects/test-project/datasets/test-dataset" {
					t.Errorf("ParentString() = %q, want %q", gotParent, "projects/test-project/datasets/test-dataset")
				}
				if gotHost := got.Host(); gotHost != "bigquery.googleapis.com" {
					t.Errorf("Host() = %q, want %q", gotHost, "bigquery.googleapis.com")
				}
			}
		})
	}
}

func TestBigQueryDatasetAccess_GetIdentity(t *testing.T) {
	ctx := context.Background()

	customResourceID := "custom-access"
	tests := []struct {
		name    string
		obj     *BigQueryDatasetAccess
		want    *BigQueryDatasetAccessIdentity
		wantErr bool
	}{
		{
			name: "with explicit resourceID",
			obj: &BigQueryDatasetAccess{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-name",
					Namespace: "test-namespace",
				},
				Spec: BigQueryDatasetAccessSpec{
					ProjectRef: parent.ProjectRef{
						External: "test-project",
					},
					DatasetId:  "test-dataset",
					ResourceID: &customResourceID,
				},
			},
			want: &BigQueryDatasetAccessIdentity{
				Project: "test-project",
				Dataset: "test-dataset",
				Access:  "custom-access",
			},
		},
		{
			name: "fallback to metadata.name when resourceID is not set",
			obj: &BigQueryDatasetAccess{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "k8s-name",
					Namespace: "test-namespace",
				},
				Spec: BigQueryDatasetAccessSpec{
					ProjectRef: parent.ProjectRef{
						External: "test-project",
					},
					DatasetId: "test-dataset",
				},
			},
			want: &BigQueryDatasetAccessIdentity{
				Project: "test-project",
				Dataset: "test-dataset",
				Access:  "k8s-name",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, nil)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetIdentity() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !tc.wantErr {
				got, ok := id.(*BigQueryDatasetAccessIdentity)
				if !ok {
					t.Fatalf("GetIdentity() returned type %T, want *BigQueryDatasetAccessIdentity", id)
				}
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
