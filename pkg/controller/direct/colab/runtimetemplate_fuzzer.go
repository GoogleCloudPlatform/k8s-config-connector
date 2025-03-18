// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate
// api.group: vertexai.cnrm.cloud.google.com

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func VertexAINotebookRuntimeTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeTemplate) *krm.VertexAINotebookRuntimeTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAINotebookRuntimeTemplateSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.MachineSpec = MachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.DataPersistentDiskSpec = PersistentDiskSpec_FromProto(mapCtx, in.GetDataPersistentDiskSpec())
	out.NetworkSpec = NetworkSpec_FromProto(mapCtx, in.GetNetworkSpec())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &krm.ServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.Labels = in.Labels
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_FromProto(mapCtx, in.GetIdleShutdownConfig())
	out.EUCConfig = NotebookEUCConfig_FromProto(mapCtx, in.GetEucConfig())
	out.NotebookRuntimeType = direct.LazyPtr(in.GetNotebookRuntimeType())
	out.ShieldedVMConfig = ShieldedVMConfig_FromProto(mapCtx, in.GetShieldedVmConfig())
	out.NetworkTags = in.NetworkTags
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	return out
}
func VertexAINotebookRuntimeTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAINotebookRuntimeTemplateSpec) *pb.NotebookRuntimeTemplate {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntimeTemplate{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.MachineSpec = MachineSpec_ToProto(mapCtx, in.MachineSpec)
	out.DataPersistentDiskSpec = PersistentDiskSpec_ToProto(mapCtx, in.DataPersistentDiskSpec)
	out.NetworkSpec = NetworkSpec_ToProto(mapCtx, in.NetworkSpec)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Labels = in.Labels
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_ToProto(mapCtx, in.IdleShutdownConfig)
	out.EucConfig = NotebookEUCConfig_ToProto(mapCtx, in.EUCConfig)
	out.NotebookRuntimeType = direct.ValueOf(in.NotebookRuntimeType)
	out.ShieldedVmConfig = ShieldedVMConfig_ToProto(mapCtx, in.ShieldedVMConfig)
	out.NetworkTags = in.NetworkTags
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	return out
}
func VertexAINotebookRuntimeTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeTemplate) *krm.VertexAINotebookRuntimeTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAINotebookRuntimeTemplateObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: IsDefault
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	// MISSING: ServiceAccount
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: Labels
	// MISSING: IdleShutdownConfig
	out.EUCConfig = NotebookEUCConfigObservedState_FromProto(mapCtx, in.GetEucConfig())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: NotebookRuntimeType
	// MISSING: ShieldedVMConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	return out
}
func VertexAINotebookRuntimeTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAINotebookRuntimeTemplateObservedState) *pb.NotebookRuntimeTemplate {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntimeTemplate{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: IsDefault
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	// MISSING: ServiceAccount
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: Labels
	// MISSING: IdleShutdownConfig
	out.EucConfig = NotebookEUCConfigObservedState_ToProto(mapCtx, in.EUCConfig)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: NotebookRuntimeType
	// MISSING: ShieldedVMConfig
	// MISSING: NetworkTags
	// MISSING: EncryptionSpec
	return out
}


