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

package ces

import (
	"context"
	"testing"

	pb "cloud.google.com/go/ces/apiv1/cespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ces/v1alpha1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestAudioRecordingConfig_ToProto(t *testing.T) {
	tests := []struct {
		name string
		in   *krm.AudioRecordingConfig
		want *pb.AudioRecordingConfig
	}{
		{
			name: "maps_canonical_gcs_bucket_ref_to_gs_uri",
			in: &krm.AudioRecordingConfig{
				GCSBucketRef:  &storagev1beta1.StorageBucketRef{External: "projects/my-proj/buckets/audio-bkt"},
				GCSPathPrefix: direct.LazyPtr("recordings/$session"),
			},
			want: &pb.AudioRecordingConfig{
				GcsBucket:     "gs://audio-bkt",
				GcsPathPrefix: "recordings/$session",
			},
		},
		{
			name: "maps_bare_gcs_bucket_ref_to_gs_uri",
			in: &krm.AudioRecordingConfig{
				GCSBucketRef: &storagev1beta1.StorageBucketRef{External: "audio-bkt"},
			},
			want: &pb.AudioRecordingConfig{
				GcsBucket: "gs://audio-bkt",
			},
		},
		{
			name: "maps_gs_uri_bucket_ref_with_trailing_slash",
			in: &krm.AudioRecordingConfig{
				GCSBucketRef: &storagev1beta1.StorageBucketRef{External: "gs://audio-bkt/"},
			},
			want: &pb.AudioRecordingConfig{
				GcsBucket: "gs://audio-bkt",
			},
		},
		{
			name: "preserves_explicit_gcs_bucket_string",
			in: &krm.AudioRecordingConfig{
				GCSBucket: direct.LazyPtr("gs://direct-bkt"),
			},
			want: &pb.AudioRecordingConfig{
				GcsBucket: "gs://direct-bkt",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mapCtx := &direct.MapContext{}
			got := AudioRecordingConfig_ToProto(mapCtx, tc.in)

			if mapCtx.Err() != nil {
				t.Fatalf("AudioRecordingConfig_ToProto(%+v) unexpected error: %v", tc.in, mapCtx.Err())
			}
			if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("AudioRecordingConfig_ToProto(%+v) mismatch (-want +got):\n%s", tc.in, diff)
			}
		})
	}
}

func TestAudioRecordingConfig_FromProtoAndCompare(t *testing.T) {
	tests := []struct {
		name       string
		actual     *pb.AudioRecordingConfig
		desired    *krm.AudioRecordingConfig
		wantDiff   bool
		wantBucket string
		wantErr    bool
	}{
		{
			name: "exports_gs_uri_to_gcs_bucket_ref_and_converges_with_zero_diff",
			actual: &pb.AudioRecordingConfig{
				GcsBucket:     "gs://audio-bkt",
				GcsPathPrefix: "recordings/$session",
			},
			desired: &krm.AudioRecordingConfig{
				GCSBucketRef:  &storagev1beta1.StorageBucketRef{External: "projects/my-proj/buckets/audio-bkt"},
				GCSPathPrefix: direct.LazyPtr("recordings/$session"),
			},
			wantBucket: "gs://audio-bkt",
		},
		{
			name:   "detects_gcs_bucket_ref_update",
			actual: &pb.AudioRecordingConfig{GcsBucket: "gs://old-bkt"},
			desired: &krm.AudioRecordingConfig{
				GCSBucketRef: &storagev1beta1.StorageBucketRef{External: "projects/my-proj/buckets/new-bkt"},
			},
			wantDiff:   true,
			wantBucket: "gs://new-bkt",
		},
		{
			name:   "detects_gcs_bucket_ref_update_when_gcp_holds_trailing_slash",
			actual: &pb.AudioRecordingConfig{GcsBucket: "gs://old-bkt/"},
			desired: &krm.AudioRecordingConfig{
				GCSBucketRef: &storagev1beta1.StorageBucketRef{External: "projects/my-proj/buckets/new-bkt"},
			},
			wantDiff:   true,
			wantBucket: "gs://new-bkt",
		},
		{
			name:   "preserves_legacy_gcs_bucket_updates_after_merge",
			actual: &pb.AudioRecordingConfig{GcsBucket: "gs://old-bkt"},
			desired: &krm.AudioRecordingConfig{
				GCSBucket: direct.LazyPtr("gs://new-direct-bkt"),
			},
			wantDiff:   true,
			wantBucket: "gs://new-direct-bkt",
		},
		{
			name:   "no_diff_for_matching_gcs_bucket_with_trailing_slash",
			actual: &pb.AudioRecordingConfig{GcsBucket: "gs://audio-bkt/"},
			desired: &krm.AudioRecordingConfig{
				GCSBucket: direct.LazyPtr("gs://audio-bkt/"),
			},
			wantBucket: "gs://audio-bkt/",
		},
		{
			name:   "no_diff_for_matching_gs_uri_gcs_bucket_ref",
			actual: &pb.AudioRecordingConfig{GcsBucket: "gs://audio-bkt"},
			desired: &krm.AudioRecordingConfig{
				GCSBucketRef: &storagev1beta1.StorageBucketRef{External: "gs://audio-bkt"},
			},
			wantBucket: "gs://audio-bkt",
		},
		{
			name:   "rejects_unparseable_gcs_bucket_ref",
			actual: &pb.AudioRecordingConfig{GcsBucket: "gs://audio-bkt"},
			desired: &krm.AudioRecordingConfig{
				GCSBucketRef: &storagev1beta1.StorageBucketRef{External: "gs://"},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := &pb.App{
				DisplayName:     "My App",
				LoggingSettings: &pb.LoggingSettings{AudioRecordingConfig: tc.actual},
			}
			desired := &krm.CESAppSpec{
				DisplayName:     direct.LazyPtr("My App"),
				LoggingSettings: &krm.LoggingSettings{AudioRecordingConfig: tc.desired},
			}

			merged, diff, updateMask, err := common.CompareGreenfieldSpec(
				context.Background(),
				desired,
				actual,
				CESAppSpec_FromProto,
				CESAppSpec_ToProto,
				normalizeAppForDiff,
			)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("CompareGreenfieldSpec() = nil error, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("CompareGreenfieldSpec() unexpected error: %v", err)
			}
			if diff.HasDiff() != tc.wantDiff {
				t.Errorf("CompareGreenfieldSpec() HasDiff() = %v (paths %v), want %v", diff.HasDiff(), updateMask.GetPaths(), tc.wantDiff)
			}
			if got := merged.GetLoggingSettings().GetAudioRecordingConfig().GetGcsBucket(); got != tc.wantBucket {
				t.Errorf("CompareGreenfieldSpec() merged GcsBucket = %q, want %q", got, tc.wantBucket)
			}
		})
	}
}
