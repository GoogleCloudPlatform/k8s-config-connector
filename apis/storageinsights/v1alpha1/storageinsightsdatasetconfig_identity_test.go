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

package v1alpha1_test

import (
	"context"
	"testing"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/storageinsights/v1alpha1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// mockReader implements client.Reader for testing
type mockReader struct {
	client.Reader
}

func (m *mockReader) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	return nil
}

func TestStorageInsightsDatasetConfigIdentity(t *testing.T) {
	ctx := context.TODO()
	reader := &mockReader{}

	tests := []struct {
		name     string
		obj      *v1alpha1.StorageInsightsDatasetConfig
		expected string
		wantErr  bool
	}{
		{
			name: "basic format",
			obj: &v1alpha1.StorageInsightsDatasetConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config",
				},
				Spec: v1alpha1.StorageInsightsDatasetConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "test-project",
					},
					Location: "us-central1",
				},
			},
			expected: "projects/test-project/locations/us-central1/datasetConfigs/test-config",
		},
		{
			name: "with resource id",
			obj: &v1alpha1.StorageInsightsDatasetConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config",
				},
				Spec: v1alpha1.StorageInsightsDatasetConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "test-project",
					},
					Location:   "us-central1",
					ResourceID: ptr("custom-config"),
				},
			},
			expected: "projects/test-project/locations/us-central1/datasetConfigs/custom-config",
		},
		{
			name: "cross-check with status external ref success",
			obj: &v1alpha1.StorageInsightsDatasetConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config",
				},
				Spec: v1alpha1.StorageInsightsDatasetConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "test-project",
					},
					Location: "us-central1",
				},
				Status: v1alpha1.StorageInsightsDatasetConfigStatus{
					ExternalRef: ptr("projects/test-project/locations/us-central1/datasetConfigs/test-config"),
				},
			},
			expected: "projects/test-project/locations/us-central1/datasetConfigs/test-config",
		},
		{
			name: "cross-check with status external ref failure",
			obj: &v1alpha1.StorageInsightsDatasetConfig{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-config",
				},
				Spec: v1alpha1.StorageInsightsDatasetConfigSpec{
					ProjectRef: &refsv1beta1.ProjectRef{
						External: "test-project",
					},
					Location: "us-central1",
				},
				Status: v1alpha1.StorageInsightsDatasetConfigStatus{
					ExternalRef: ptr("projects/test-project/locations/us-central1/datasetConfigs/different-config"),
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id, err := tc.obj.GetIdentity(ctx, reader)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if id.String() != tc.expected {
				t.Errorf("expected identity %q, got %q", tc.expected, id.String())
			}
		})
	}
}

func TestStorageInsightsDatasetConfigIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		ref      string
		expected v1alpha1.StorageInsightsDatasetConfigIdentity
		wantErr  bool
	}{
		{
			name: "valid ref",
			ref:  "projects/test-project/locations/us-central1/datasetConfigs/test-config",
			expected: v1alpha1.StorageInsightsDatasetConfigIdentity{
				Project:       "test-project",
				Location:      "us-central1",
				DatasetConfig: "test-config",
			},
		},
		{
			name:    "invalid format",
			ref:     "projects/test-project/locations/us-central1/invalid/test-config",
			wantErr: true,
		},
		{
			name:    "missing dataset config id",
			ref:     "projects/test-project/locations/us-central1/datasetConfigs/",
			wantErr: true,
		},
		{
			name:    "missing project id",
			ref:     "projects//locations/us-central1/datasetConfigs/test-config",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			id := &v1alpha1.StorageInsightsDatasetConfigIdentity{}
			err := id.FromExternal(tc.ref)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.expected, *id); diff != "" {
				t.Errorf("FromExternal() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func ptr(s string) *string {
	return &s
}
