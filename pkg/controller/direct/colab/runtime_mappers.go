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
	if in == nil {
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

func ColabRuntimeSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntime) *krm.ColabRuntimeSpec {
	if in == nil {
		return nil
	}
	out := &krm.ColabRuntimeSpec{}
	// MISSING: Name
	out.RuntimeUser = direct.LazyPtr(in.GetRuntimeUser())
	out.ColabRuntimeTemplateRef = NotebookRuntimeTemplateRef_FromProto(mapCtx, in.GetNotebookRuntimeTemplateRef())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	// MISSING: EUCConfig
	// MISSING: ShieldedVMConfig
	out.NetworkTags = in.NetworkTags
	// MISSING: SoftwareConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}

func ColabRuntimeSpec_ToProto(mapCtx *direct.MapContext, in *krm.ColabRuntimeSpec) *pb.NotebookRuntime {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntime{}
	// MISSING: Name
	out.RuntimeUser = direct.ValueOf(in.RuntimeUser)
	out.NotebookRuntimeTemplateRef = NotebookRuntimeTemplateRef_ToProto(mapCtx, in.ColabRuntimeTemplateRef)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	// MISSING: EUCConfig
	// MISSING: ShieldedVMConfig
	out.NetworkTags = in.NetworkTags
	// MISSING: SoftwareConfig
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
