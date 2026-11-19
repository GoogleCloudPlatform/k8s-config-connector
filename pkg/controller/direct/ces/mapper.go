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

	pb "cloud.google.com/go/ces/apiv1/cespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ces/v1alpha1"
	krmdiscoveryenginev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AudioRecordingConfig_FromProto(mapCtx *direct.MapContext, in *pb.AudioRecordingConfig) *krm.AudioRecordingConfig {
	if in == nil {
		return nil
	}
	out := &krm.AudioRecordingConfig{}
	// Map every parseable bucket to gcsBucketRef, so that CompareGreenfieldSpec
	// never merges a GCP-held gcsBucket into a spec that sets only gcsBucketRef
	// (gcsBucket takes precedence in ToProto and would hide the change).
	if gcsBucket := in.GetGcsBucket(); gcsBucket != "" {
		bucketID := &storagev1beta1.StorageBucketIdentity{}
		if bucketID.FromExternal(gcsBucket) == nil {
			out.GCSBucketRef = &storagev1beta1.StorageBucketRef{External: "gs://" + bucketID.BucketName()}
		} else {
			out.GCSBucket = direct.LazyPtr(gcsBucket)
		}
	}
	out.GCSPathPrefix = direct.LazyPtr(in.GetGcsPathPrefix())
	return out
}

// normalizeAppForDiff canonicalizes loggingSettings.audioRecordingConfig.gcsBucket
// (for example "gs://b/" to "gs://b") on both sides of the diff, because
// FromProto renders every parseable GCP value as "gs://<bucket>".
func normalizeAppForDiff(_ context.Context, app *pb.App) error {
	cfg := app.GetLoggingSettings().GetAudioRecordingConfig()
	if cfg.GetGcsBucket() == "" {
		return nil
	}
	bucketID := &storagev1beta1.StorageBucketIdentity{}
	if bucketID.FromExternal(cfg.GetGcsBucket()) == nil {
		cfg.GcsBucket = "gs://" + bucketID.BucketName()
	}
	return nil
}

func AudioRecordingConfig_ToProto(mapCtx *direct.MapContext, in *krm.AudioRecordingConfig) *pb.AudioRecordingConfig {
	if in == nil {
		return nil
	}
	out := &pb.AudioRecordingConfig{}
	if in.GCSBucket != nil {
		out.GcsBucket = *in.GCSBucket
	} else if in.GCSBucketRef != nil && in.GCSBucketRef.External != "" {
		bucketID := &storagev1beta1.StorageBucketIdentity{}
		if err := bucketID.FromExternal(in.GCSBucketRef.External); err != nil {
			mapCtx.Errorf("parsing gcsBucketRef.external %q: %v", in.GCSBucketRef.External, err)
		} else {
			out.GcsBucket = "gs://" + bucketID.BucketName()
		}
	}
	out.GcsPathPrefix = direct.ValueOf(in.GCSPathPrefix)
	return out
}

func DataStore_ToProto(mapCtx *direct.MapContext, in *krm.DataStore) *pb.DataStore {
	if in == nil {
		return nil
	}
	out := &pb.DataStore{}
	if in.DataStoreRef != nil {
		out.Name = in.DataStoreRef.External
	}
	return out
}

func DataStore_FromProto(mapCtx *direct.MapContext, in *pb.DataStore) *krm.DataStore {
	if in == nil {
		return nil
	}
	out := &krm.DataStore{}
	if in.GetName() != "" {
		out.DataStoreRef = &krmdiscoveryenginev1alpha1.DiscoveryEngineDataStoreRef{External: in.GetName()}
	}
	return out
}

func DataStoreSettings_ToProto(mapCtx *direct.MapContext, in *krm.DataStoreSettings) *pb.DataStoreSettings {
	if in == nil {
		return nil
	}
	return &pb.DataStoreSettings{}
}

func DataStoreSettings_FromProto(mapCtx *direct.MapContext, in *pb.DataStoreSettings) *krm.DataStoreSettings {
	if in == nil {
		return nil
	}
	return &krm.DataStoreSettings{}
}

func Schema_ToProto(mapCtx *direct.MapContext, in *krm.Schema) *pb.Schema {
	if in == nil {
		return nil
	}
	return &pb.Schema{}
}

func Schema_FromProto(mapCtx *direct.MapContext, in *pb.Schema) *krm.Schema {
	if in == nil {
		return nil
	}
	return &krm.Schema{}
}

func TriggerAction_Response_FromProto(mapCtx *direct.MapContext, in *pb.TriggerAction_Response) *krm.TriggerAction_Response {
	if in == nil {
		return nil
	}
	out := &krm.TriggerAction_Response{}
	out.Text = direct.LazyPtr(in.GetText())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	return out
}

func TriggerAction_Response_ToProto(mapCtx *direct.MapContext, in *krm.TriggerAction_Response) *pb.TriggerAction_Response {
	if in == nil {
		return nil
	}
	out := &pb.TriggerAction_Response{}
	out.Text = direct.ValueOf(in.Text)
	out.Disabled = direct.ValueOf(in.Disabled)
	return out
}
