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

package logging

import (
	"testing"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestLoggingLogSinkSpec_StorageBucketDestination(t *testing.T) {
	tests := []struct {
		name     string
		external string
		wantDest string
	}{
		{
			name:     "canonical_project_bucket_path",
			external: "projects/my-proj/buckets/sink-bkt",
			wantDest: "storage.googleapis.com/sink-bkt",
		},
		{
			name:     "gs_uri_bucket",
			external: "gs://sink-bkt",
			wantDest: "storage.googleapis.com/sink-bkt",
		},
		{
			name:     "gs_uri_with_trailing_slash",
			external: "gs://sink-bkt/",
			wantDest: "storage.googleapis.com/sink-bkt",
		},
		{
			name:     "self_link_bucket",
			external: "https://www.googleapis.com/storage/v1/b/sink-bkt",
			wantDest: "storage.googleapis.com/sink-bkt",
		},
		{
			name:     "host_prefixed_bucket",
			external: "//storage.googleapis.com/projects/my-proj/buckets/sink-bkt",
			wantDest: "storage.googleapis.com/sink-bkt",
		},
		{
			name:     "bare_bucket_name",
			external: "sink-bkt",
			wantDest: "storage.googleapis.com/sink-bkt",
		},
		{
			name:     "logging_destination_form",
			external: "storage.googleapis.com/sink-bkt",
			wantDest: "storage.googleapis.com/sink-bkt",
		},
		{
			name:     "unparseable_external_passes_through",
			external: "Exported_Bkt",
			wantDest: "storage.googleapis.com/Exported_Bkt",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			in := &krm.LoggingLogSinkSpec{
				Destination: krm.LogSinkDestination{
					StorageBucketRef: &storagev1beta1.StorageBucketRef{External: tc.external},
				},
			}

			got := LoggingLogSinkSpec_ToProto(mapCtx, in)

			if mapCtx.Err() != nil {
				t.Fatalf("LoggingLogSinkSpec_ToProto(%q) unexpected error: %v", tc.external, mapCtx.Err())
			}
			if got.GetDestination() != tc.wantDest {
				t.Errorf("LoggingLogSinkSpec_ToProto(%q).Destination = %q, want %q", tc.external, got.GetDestination(), tc.wantDest)
			}
		})
	}
}

func TestLoggingLogSinkSpec_FromProto_ExportedStorageBucketValidates(t *testing.T) {
	mapCtx := &direct.MapContext{}
	in := &pb.LogSink{
		Destination: "storage.googleapis.com/exported-bkt",
	}

	got := LoggingLogSinkSpec_FromProto(mapCtx, in)

	if mapCtx.Err() != nil {
		t.Fatalf("LoggingLogSinkSpec_FromProto(%v) unexpected error: %v", in, mapCtx.Err())
	}
	if got.Destination.StorageBucketRef == nil || got.Destination.StorageBucketRef.External != "exported-bkt" {
		t.Fatalf("LoggingLogSinkSpec_FromProto(%v) StorageBucketRef = %+v, want External=%q", in, got.Destination.StorageBucketRef, "exported-bkt")
	}
	if err := got.Destination.StorageBucketRef.ValidateExternal(got.Destination.StorageBucketRef.External); err != nil {
		t.Errorf("ValidateExternal(%q) unexpected error: %v", got.Destination.StorageBucketRef.External, err)
	}
}
