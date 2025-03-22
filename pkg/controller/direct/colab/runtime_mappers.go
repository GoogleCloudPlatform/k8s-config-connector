// Copyright 2025 Google LLC
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

package colab

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/colab/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NotebookRuntimeTemplateRef_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeTemplateRef) *krm.NotebookRuntimeTemplateRef {
	if in == nil || in.NotebookRuntimeTemplate == "" {
		return nil
	}
	return &krm.NotebookRuntimeTemplateRef{
		External: in.NotebookRuntimeTemplate,
	}
}
func NotebookRuntimeTemplateRef_ToProto(mapCtx *direct.MapContext, in *krm.NotebookRuntimeTemplateRef) *pb.NotebookRuntimeTemplateRef {
	if in == nil {
		return nil
	}
	return &pb.NotebookRuntimeTemplateRef{
		NotebookRuntimeTemplate: in.External,
	}
}
