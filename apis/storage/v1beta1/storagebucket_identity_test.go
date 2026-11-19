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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestStorageBucketIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name           string
		ref            string
		want           *StorageBucketIdentity
		wantStringForm string
		wantErr        bool
	}{
		{
			name: "valid_full_url",
			ref:  "//storage.googleapis.com/projects/my-project/buckets/my-bucket",
			want: &StorageBucketIdentity{
				Project: "my-project",
				Bucket:  "my-bucket",
			},
			wantStringForm: "projects/my-project/buckets/my-bucket",
		},
		{
			name: "valid_relative_path",
			ref:  "projects/my-project/buckets/my-bucket",
			want: &StorageBucketIdentity{
				Project: "my-project",
				Bucket:  "my-bucket",
			},
			wantStringForm: "projects/my-project/buckets/my-bucket",
		},
		{
			name: "valid_wildcard_project_path",
			ref:  "projects/_/buckets/my-bucket",
			want: &StorageBucketIdentity{
				Project: "_",
				Bucket:  "my-bucket",
			},
			wantStringForm: "projects/_/buckets/my-bucket",
		},
		{
			name: "valid_bare_bucket_name",
			ref:  "my-bucket",
			want: &StorageBucketIdentity{
				Project: "_",
				Bucket:  "my-bucket",
			},
			wantStringForm: "projects/_/buckets/my-bucket",
		},
		{
			name: "valid_bare_dotted_bucket_name",
			ref:  "my.bucket.example.com",
			want: &StorageBucketIdentity{
				Project: "_",
				Bucket:  "my.bucket.example.com",
			},
			wantStringForm: "projects/_/buckets/my.bucket.example.com",
		},
		{
			name: "valid_www_googleapis_self_link",
			ref:  "https://www.googleapis.com/storage/v1/b/my-bucket",
			want: &StorageBucketIdentity{
				Project: "_",
				Bucket:  "my-bucket",
			},
			wantStringForm: "projects/_/buckets/my-bucket",
		},
		{
			name: "valid_storage_googleapis_self_link",
			ref:  "https://storage.googleapis.com/storage/v1/b/my-bucket",
			want: &StorageBucketIdentity{
				Project: "_",
				Bucket:  "my-bucket",
			},
			wantStringForm: "projects/_/buckets/my-bucket",
		},
		{
			name: "valid_gs_scheme",
			ref:  "gs://my-bucket",
			want: &StorageBucketIdentity{
				Project: "_",
				Bucket:  "my-bucket",
			},
			wantStringForm: "projects/_/buckets/my-bucket",
		},
		{
			name: "valid_gs_scheme_with_single_trailing_slash",
			ref:  "gs://my-bucket/",
			want: &StorageBucketIdentity{
				Project: "_",
				Bucket:  "my-bucket",
			},
			wantStringForm: "projects/_/buckets/my-bucket",
		},
		{
			name:    "invalid_empty",
			ref:     "",
			wantErr: true,
		},
		{
			name:    "invalid_gs_empty",
			ref:     "gs://",
			wantErr: true,
		},
		{
			name:    "invalid_gs_whitespace",
			ref:     "gs://my bucket",
			wantErr: true,
		},
		{
			name:    "invalid_gs_double_trailing_slash",
			ref:     "gs://my-bucket//",
			wantErr: true,
		},
		{
			name:    "invalid_gs_with_subpath",
			ref:     "gs://my-bucket/subpath",
			wantErr: true,
		},
		{
			name:    "invalid_bare_uppercase",
			ref:     "My-Bucket",
			wantErr: true,
		},
		{
			name:    "invalid_bare_whitespace",
			ref:     "my bucket",
			wantErr: true,
		},
		{
			name:    "invalid_self_link_with_object_path",
			ref:     "https://www.googleapis.com/storage/v1/b/my-bucket/o/my-object",
			wantErr: true,
		},
		{
			name:    "invalid_missing_project",
			ref:     "buckets/my-bucket",
			wantErr: true,
		},
		{
			name:    "invalid_extra_segments",
			ref:     "projects/my-project/buckets/my-bucket/extra",
			wantErr: true,
		},
		{
			name:    "invalid_wrong_collection",
			ref:     "projects/my-project/foos/my-bucket",
			wantErr: true,
		},
		{
			name:    "invalid_empty_bucket_segment",
			ref:     "projects/my-project/buckets/",
			wantErr: true,
		},
		{
			name:    "invalid_path_uppercase_bucket",
			ref:     "projects/my-project/buckets/My-Bucket",
			wantErr: true,
		},
		{
			name:    "invalid_leading_hyphen",
			ref:     "-bkt",
			wantErr: true,
		},
		{
			name:    "invalid_trailing_dot",
			ref:     "bkt.",
			wantErr: true,
		},
		{
			name:    "invalid_double_dot",
			ref:     "my..bkt",
			wantErr: true,
		},
		{
			name:    "invalid_too_long",
			ref:     strings.Repeat("a", 223),
			wantErr: true,
		},
		{
			name: "valid_max_length_222",
			ref:  strings.Repeat("a", 222),
			want: &StorageBucketIdentity{
				Project: "_",
				Bucket:  strings.Repeat("a", 222),
			},
			wantStringForm: "projects/_/buckets/" + strings.Repeat("a", 222),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			i := &StorageBucketIdentity{}
			err := i.FromExternal(tc.ref)

			if (err != nil) != tc.wantErr {
				t.Fatalf("FromExternal(%q) error = %v, wantErr %v", tc.ref, err, tc.wantErr)
			}
			if !tc.wantErr {
				if diff := cmp.Diff(tc.want, i); diff != "" {
					t.Errorf("FromExternal(%q) mismatch (-want +got):\n%s", tc.ref, diff)
				}
				if got := i.String(); got != tc.wantStringForm {
					t.Errorf("FromExternal(%q).String() = %q, want %q", tc.ref, got, tc.wantStringForm)
				}
			}
		})
	}
}

func TestStorageBucketIdentity_Methods(t *testing.T) {
	i := &StorageBucketIdentity{
		Project: "my-project",
		Bucket:  "my-bucket",
	}

	if got := i.String(); got != "projects/my-project/buckets/my-bucket" {
		t.Errorf("String() = %q, want %q", got, "projects/my-project/buckets/my-bucket")
	}
	if got := i.BucketName(); got != "my-bucket" {
		t.Errorf("BucketName() = %q, want %q", got, "my-bucket")
	}
}

func TestStorageBucketGetIdentity(t *testing.T) {
	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() unexpected error: %v", err)
	}
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	t.Run("valid_bucket_identity", func(t *testing.T) {
		obj := &StorageBucket{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-bucket",
				Namespace: "test-namespace",
				Annotations: map[string]string{
					"cnrm.cloud.google.com/project-id": "test-project",
				},
			},
		}

		gotIdentity, err := obj.GetIdentity(t.Context(), fakeClient)

		if err != nil {
			t.Fatalf("GetIdentity() unexpected error: %v", err)
		}
		wantIdentity := &StorageBucketIdentity{
			Project: "test-project",
			Bucket:  "test-bucket",
		}
		if diff := cmp.Diff(wantIdentity, gotIdentity); diff != "" {
			t.Errorf("GetIdentity() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("missing_project_wraps_underlying_error", func(t *testing.T) {
		obj := &StorageBucket{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-bucket",
				Namespace: "test-namespace",
			},
		}

		got, err := obj.GetIdentity(t.Context(), fakeClient)

		if err == nil {
			t.Fatalf("GetIdentity() = %+v, want error", got)
		}
		if !strings.Contains(err.Error(), "cannot find project id") {
			t.Errorf("GetIdentity() error = %q, want wrapped project error", err.Error())
		}
	})

	t.Run("invalid_resource_id_path_traversal_rejected", func(t *testing.T) {
		badID := "bucket/extra"
		obj := &StorageBucket{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-bucket",
				Namespace: "test-namespace",
				Annotations: map[string]string{
					"cnrm.cloud.google.com/project-id": "test-project",
				},
			},
			Spec: StorageBucketSpec{
				ResourceID: &badID,
			},
		}

		got, err := obj.GetIdentity(t.Context(), fakeClient)

		if err == nil {
			t.Fatalf("GetIdentity() = %+v, want error", got)
		}
	})
}
