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
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestNewAnywhereCacheIdentity(t *testing.T) {
	strPtr := func(s string) *string { return &s }

	tests := []struct {
		name       string
		bucketExt  string
		resourceID *string
		wantString string
		wantErr    bool
	}{
		{
			name:       "canonical_bucket_external",
			bucketExt:  "projects/my-project/buckets/my-bucket",
			resourceID: strPtr("us-central1-a"),
			wantString: "projects/_/buckets/my-bucket/anywhereCaches/us-central1-a",
		},
		{
			name:       "gs_uri_bucket_external",
			bucketExt:  "gs://my-bucket",
			resourceID: strPtr("us-central1-a"),
			wantString: "projects/_/buckets/my-bucket/anywhereCaches/us-central1-a",
		},
		{
			name:       "host_prefixed_bucket_external",
			bucketExt:  "//storage.googleapis.com/projects/my-project/buckets/my-bucket",
			resourceID: strPtr("us-central1-a"),
			wantString: "projects/_/buckets/my-bucket/anywhereCaches/us-central1-a",
		},
		{
			name:       "self_link_bucket_external",
			bucketExt:  "https://www.googleapis.com/storage/v1/b/my-bucket",
			resourceID: strPtr("us-central1-a"),
			wantString: "projects/_/buckets/my-bucket/anywhereCaches/us-central1-a",
		},
		{
			name:       "bare_bucket_external",
			bucketExt:  "my-bucket",
			resourceID: strPtr("us-central1-a"),
			wantString: "projects/_/buckets/my-bucket/anywhereCaches/us-central1-a",
		},
		{
			name:       "invalid_bucket_external",
			bucketExt:  "gs://",
			resourceID: strPtr("us-central1-a"),
			wantErr:    true,
		},
		{
			name:       "nil_bucket_ref_returns_error",
			resourceID: strPtr("us-central1-a"),
			wantErr:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).Build()
			var bucketRef *StorageBucketRef
			if tc.bucketExt != "" {
				bucketRef = &StorageBucketRef{External: tc.bucketExt}
			}
			obj := &StorageAnywhereCache{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-cache",
					Namespace: "default",
				},
				Spec: StorageAnywhereCacheSpec{
					BucketRef:  bucketRef,
					ResourceID: tc.resourceID,
				},
			}

			got, err := NewAnywhereCacheIdentity(t.Context(), reader, obj)

			if (err != nil) != tc.wantErr {
				t.Fatalf("NewAnywhereCacheIdentity(%q) error = %v, wantErr %v", tc.bucketExt, err, tc.wantErr)
			}
			if !tc.wantErr && got.String() != tc.wantString {
				t.Errorf("NewAnywhereCacheIdentity(%q).String() = %q, want %q", tc.bucketExt, got.String(), tc.wantString)
			}
		})
	}
}
