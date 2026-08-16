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

package aiplatform

import (
	"strings"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func PipelineTemplateMetadata_FromProto(mapCtx *direct.MapContext, in *pb.PipelineTemplateMetadata) *krm.PipelineTemplateMetadata {
	if in == nil {
		return nil
	}
	out := &krm.PipelineTemplateMetadata{}
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}

func PipelineTemplateMetadata_ToProto(mapCtx *direct.MapContext, in *krm.PipelineTemplateMetadata) *pb.PipelineTemplateMetadata {
	if in == nil {
		return nil
	}
	out := &pb.PipelineTemplateMetadata{}
	out.Version = direct.ValueOf(in.Version)
	return out
}

func PipelineJobRuntimeConfig_FromProto(mapCtx *direct.MapContext, in *pb.PipelineJob_RuntimeConfig) *krm.PipelineJobRuntimeConfig {
	if in == nil {
		return nil
	}
	out := &krm.PipelineJobRuntimeConfig{}
	if in.GetGcsOutputDirectory() != "" {
		out.GCSOutputDirectoryRef = &storagev1beta1.StorageBucketRef{
			External: in.GetGcsOutputDirectory(),
		}
	}
	out.FailurePolicy = direct.Enum_FromProto(mapCtx, in.GetFailurePolicy())
	return out
}

func PipelineJobRuntimeConfig_ToProto(mapCtx *direct.MapContext, in *krm.PipelineJobRuntimeConfig) *pb.PipelineJob_RuntimeConfig {
	if in == nil {
		return nil
	}
	out := &pb.PipelineJob_RuntimeConfig{}
	if in.GCSOutputDirectoryRef != nil {
		out.GcsOutputDirectory = gcsOutputDirectoryToProto(mapCtx, in.GCSOutputDirectoryRef)
	}
	out.FailurePolicy = direct.Enum_ToProto[pb.PipelineFailurePolicy](mapCtx, in.FailurePolicy)
	return out
}

func gcsOutputDirectoryToProto(mapCtx *direct.MapContext, ref *storagev1beta1.StorageBucketRef) string {
	if ref == nil {
		return ""
	}
	if strings.HasPrefix(ref.External, "gs://") {
		return ref.External
	}
	id := &storagev1beta1.StorageBucketIdentity{}
	if err := id.FromExternal(ref.External); err == nil {
		return "gs://" + id.Bucket
	}
	// Fallback/identity for fuzzer or unparsed raw values: keep exactly as-is to ensure perfect roundtrip
	return ref.External
}
