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

package dataplex

import (
	"testing"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestAsset_ResourceSpec_ToProto(t *testing.T) {
	storageType := "STORAGE_BUCKET"

	tests := []struct {
		name     string
		external string
		wantName string
	}{
		{
			name:     "canonical_project_bucket_path",
			external: "projects/my-proj/buckets/my-bucket",
			wantName: "projects/my-proj/buckets/my-bucket",
		},
		{
			name:     "host_prefixed_bucket_path",
			external: "//storage.googleapis.com/projects/my-proj/buckets/my-bucket",
			wantName: "projects/my-proj/buckets/my-bucket",
		},
		{
			name:     "gs_uri_bucket",
			external: "gs://my-bucket",
			wantName: "projects/_/buckets/my-bucket",
		},
		{
			name:     "bare_bucket_name",
			external: "my-bucket",
			wantName: "projects/_/buckets/my-bucket",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			in := &krm.Asset_ResourceSpec{
				Type:             &storageType,
				StorageBucketRef: &storagev1beta1.StorageBucketRef{External: tc.external},
			}

			got := Asset_ResourceSpec_ToProto(mapCtx, in)

			if mapCtx.Err() != nil {
				t.Fatalf("Asset_ResourceSpec_ToProto(%q) unexpected error: %v", tc.external, mapCtx.Err())
			}
			if got.GetName() != tc.wantName {
				t.Errorf("Asset_ResourceSpec_ToProto(%q).Name = %q, want %q", tc.external, got.GetName(), tc.wantName)
			}
		})
	}
}

func TestAsset_ResourceSpec_FromProto(t *testing.T) {
	mapCtx := &direct.MapContext{}
	in := &pb.Asset_ResourceSpec{
		Type:           pb.Asset_ResourceSpec_STORAGE_BUCKET,
		Name:           "my-bucket",
		ReadAccessMode: pb.Asset_ResourceSpec_DIRECT,
	}

	got := Asset_ResourceSpec_FromProto(mapCtx, in)

	if mapCtx.Err() != nil {
		t.Fatalf("Asset_ResourceSpec_FromProto(%v) unexpected error: %v", in, mapCtx.Err())
	}
	if got.StorageBucketRef == nil || got.StorageBucketRef.External != "my-bucket" {
		t.Fatalf("Asset_ResourceSpec_FromProto(%v).StorageBucketRef = %+v, want External=%q", in, got.StorageBucketRef, "my-bucket")
	}
	if err := got.StorageBucketRef.ValidateExternal(got.StorageBucketRef.External); err != nil {
		t.Errorf("ValidateExternal(%q) unexpected error on exported StorageBucketRef: %v", got.StorageBucketRef.External, err)
	}
}
