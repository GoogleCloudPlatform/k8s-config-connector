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

package dataform

import (
	pb "cloud.google.com/go/dataform/apiv1/dataformpb"
	krmaiplatformv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	krmdataformv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1alpha1"
	krmstoragev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NotebookRuntimeOptions_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeOptions) *krmdataformv1alpha1.NotebookRuntimeOptions {
	if in == nil {
		return nil
	}
	out := &krmdataformv1alpha1.NotebookRuntimeOptions{}
	if in.GetGcsOutputBucket() != "" {
		out.GCSOutputBucketRef = &krmstoragev1beta1.StorageBucketRef{External: in.GetGcsOutputBucket()}
	}
	if in.GetAiPlatformNotebookRuntimeTemplate() != "" {
		out.AiPlatformNotebookRuntimeTemplateRef = &krmaiplatformv1alpha1.AIPlatformNotebookRuntimeTemplateRef{External: in.GetAiPlatformNotebookRuntimeTemplate()}
	}
	return out
}

func NotebookRuntimeOptions_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmdataformv1alpha1.NotebookRuntimeOptions) *pb.NotebookRuntimeOptions {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntimeOptions{}
	if in.GCSOutputBucketRef != nil {
		out.ExecutionSink = &pb.NotebookRuntimeOptions_GcsOutputBucket{
			GcsOutputBucket: in.GCSOutputBucketRef.External,
		}
	}
	if in.AiPlatformNotebookRuntimeTemplateRef != nil {
		out.AiPlatformNotebookRuntimeTemplate = in.AiPlatformNotebookRuntimeTemplateRef.External
	}
	return out
}
