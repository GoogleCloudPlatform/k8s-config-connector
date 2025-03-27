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

// +generated:mapper
// krm.group: colab.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.aiplatform.v1beta1

package colab

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/colab/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ColabRuntimeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntime) *krm.ColabRuntimeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ColabRuntimeObservedState{}
	// MISSING: Name
	// MISSING: NotebookRuntimeTemplateRef
	out.ProxyURI = direct.LazyPtr(in.GetProxyUri())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.HealthState = direct.Enum_FromProto(mapCtx, in.GetHealthState())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.RuntimeState = direct.Enum_FromProto(mapCtx, in.GetRuntimeState())
	out.IsUpgradable = direct.LazyPtr(in.GetIsUpgradable())
	out.ExpirationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpirationTime())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.NotebookRuntimeType = direct.Enum_FromProto(mapCtx, in.GetNotebookRuntimeType())
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_FromProto(mapCtx, in.GetIdleShutdownConfig())
	out.EncryptionSpec = EncryptionSpecObservedState_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ColabRuntimeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ColabRuntimeObservedState) *pb.NotebookRuntime {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntime{}
	// MISSING: Name
	// MISSING: NotebookRuntimeTemplateRef
	out.ProxyUri = direct.ValueOf(in.ProxyURI)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.HealthState = direct.Enum_ToProto[pb.NotebookRuntime_HealthState](mapCtx, in.HealthState)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	out.RuntimeState = direct.Enum_ToProto[pb.NotebookRuntime_RuntimeState](mapCtx, in.RuntimeState)
	out.IsUpgradable = direct.ValueOf(in.IsUpgradable)
	out.ExpirationTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpirationTime)
	out.Version = direct.ValueOf(in.Version)
	out.NotebookRuntimeType = direct.Enum_ToProto[pb.NotebookRuntimeType](mapCtx, in.NotebookRuntimeType)
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_ToProto(mapCtx, in.IdleShutdownConfig)
	out.EncryptionSpec = EncryptionSpecObservedState_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ColabRuntimeSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntime) *krm.ColabRuntimeSpec {
	if in == nil {
		return nil
	}
	out := &krm.ColabRuntimeSpec{}
	// MISSING: Name
	out.RuntimeUser = direct.LazyPtr(in.GetRuntimeUser())
	// MISSING: NotebookRuntimeTemplateRef
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.NetworkTags = in.NetworkTags
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
	// MISSING: NotebookRuntimeTemplateRef
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.NetworkTags = in.NetworkTags
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ColabRuntimeTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeTemplate) *krm.ColabRuntimeTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ColabRuntimeTemplateObservedState{}
	// MISSING: Name
	// MISSING: IsDefault
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.EUCConfig = NotebookEUCConfigObservedState_FromProto(mapCtx, in.GetEucConfig())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ColabRuntimeTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ColabRuntimeTemplateObservedState) *pb.NotebookRuntimeTemplate {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntimeTemplate{}
	// MISSING: Name
	// MISSING: IsDefault
	out.Etag = direct.ValueOf(in.Etag)
	out.EucConfig = NotebookEUCConfigObservedState_ToProto(mapCtx, in.EUCConfig)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func ColabRuntimeTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeTemplate) *krm.ColabRuntimeTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.ColabRuntimeTemplateSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: IsDefault
	out.MachineSpec = MachineSpec_FromProto(mapCtx, in.GetMachineSpec())
	out.DataPersistentDiskSpec = PersistentDiskSpec_FromProto(mapCtx, in.GetDataPersistentDiskSpec())
	out.NetworkSpec = NetworkSpec_FromProto(mapCtx, in.GetNetworkSpec())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refs.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.Labels = in.Labels
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_FromProto(mapCtx, in.GetIdleShutdownConfig())
	out.EUCConfig = NotebookEUCConfig_FromProto(mapCtx, in.GetEucConfig())
	out.NotebookRuntimeType = direct.Enum_FromProto(mapCtx, in.GetNotebookRuntimeType())
	out.ShieldedVMConfig = ShieldedVMConfig_FromProto(mapCtx, in.GetShieldedVmConfig())
	out.NetworkTags = in.NetworkTags
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	return out
}
func ColabRuntimeTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.ColabRuntimeTemplateSpec) *pb.NotebookRuntimeTemplate {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntimeTemplate{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: IsDefault
	out.MachineSpec = MachineSpec_ToProto(mapCtx, in.MachineSpec)
	out.DataPersistentDiskSpec = PersistentDiskSpec_ToProto(mapCtx, in.DataPersistentDiskSpec)
	out.NetworkSpec = NetworkSpec_ToProto(mapCtx, in.NetworkSpec)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Labels = in.Labels
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_ToProto(mapCtx, in.IdleShutdownConfig)
	out.EucConfig = NotebookEUCConfig_ToProto(mapCtx, in.EUCConfig)
	out.NotebookRuntimeType = direct.Enum_ToProto[pb.NotebookRuntimeType](mapCtx, in.NotebookRuntimeType)
	out.ShieldedVmConfig = ShieldedVMConfig_ToProto(mapCtx, in.ShieldedVMConfig)
	out.NetworkTags = in.NetworkTags
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	return out
}
func EncryptionSpec_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	// MISSING: KMSKeyName
	return out
}
func EncryptionSpec_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	// MISSING: KMSKeyName
	return out
}
func EncryptionSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpecObservedState{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func EncryptionSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpecObservedState) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func Featurestore_OnlineServingConfig_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore_OnlineServingConfig) *krm.Featurestore_OnlineServingConfig {
	if in == nil {
		return nil
	}
	out := &krm.Featurestore_OnlineServingConfig{}
	out.FixedNodeCount = direct.LazyPtr(in.GetFixedNodeCount())
	out.Scaling = Featurestore_OnlineServingConfig_Scaling_FromProto(mapCtx, in.GetScaling())
	return out
}
func Featurestore_OnlineServingConfig_ToProto(mapCtx *direct.MapContext, in *krm.Featurestore_OnlineServingConfig) *pb.Featurestore_OnlineServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore_OnlineServingConfig{}
	out.FixedNodeCount = direct.ValueOf(in.FixedNodeCount)
	out.Scaling = Featurestore_OnlineServingConfig_Scaling_ToProto(mapCtx, in.Scaling)
	return out
}
func Featurestore_OnlineServingConfig_Scaling_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore_OnlineServingConfig_Scaling) *krm.Featurestore_OnlineServingConfig_Scaling {
	if in == nil {
		return nil
	}
	out := &krm.Featurestore_OnlineServingConfig_Scaling{}
	out.MinNodeCount = direct.LazyPtr(in.GetMinNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	out.CPUUtilizationTarget = direct.LazyPtr(in.GetCpuUtilizationTarget())
	return out
}
func Featurestore_OnlineServingConfig_Scaling_ToProto(mapCtx *direct.MapContext, in *krm.Featurestore_OnlineServingConfig_Scaling) *pb.Featurestore_OnlineServingConfig_Scaling {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore_OnlineServingConfig_Scaling{}
	out.MinNodeCount = direct.ValueOf(in.MinNodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	out.CpuUtilizationTarget = direct.ValueOf(in.CPUUtilizationTarget)
	return out
}
func MachineSpec_FromProto(mapCtx *direct.MapContext, in *pb.MachineSpec) *krm.MachineSpec {
	if in == nil {
		return nil
	}
	out := &krm.MachineSpec{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorType = direct.Enum_FromProto(mapCtx, in.GetAcceleratorType())
	out.AcceleratorCount = direct.LazyPtr(in.GetAcceleratorCount())
	out.TpuTopology = direct.LazyPtr(in.GetTpuTopology())
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	return out
}
func MachineSpec_ToProto(mapCtx *direct.MapContext, in *krm.MachineSpec) *pb.MachineSpec {
	if in == nil {
		return nil
	}
	out := &pb.MachineSpec{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorType = direct.Enum_ToProto[pb.AcceleratorType](mapCtx, in.AcceleratorType)
	out.AcceleratorCount = direct.ValueOf(in.AcceleratorCount)
	out.TpuTopology = direct.ValueOf(in.TpuTopology)
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	return out
}
func NetworkSpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkSpec) *krm.NetworkSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkSpec{}
	out.EnableInternetAccess = direct.LazyPtr(in.GetEnableInternetAccess())
	if in.GetNetwork() != "" {
		out.NetworkRef = &refs.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &refs.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	return out
}
func NetworkSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkSpec) *pb.NetworkSpec {
	if in == nil {
		return nil
	}
	out := &pb.NetworkSpec{}
	out.EnableInternetAccess = direct.ValueOf(in.EnableInternetAccess)
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	return out
}
func NotebookEUCConfig_FromProto(mapCtx *direct.MapContext, in *pb.NotebookEucConfig) *krm.NotebookEUCConfig {
	if in == nil {
		return nil
	}
	out := &krm.NotebookEUCConfig{}
	out.EUCDisabled = direct.LazyPtr(in.GetEucDisabled())
	// MISSING: BypassActasCheck
	return out
}
func NotebookEUCConfig_ToProto(mapCtx *direct.MapContext, in *krm.NotebookEUCConfig) *pb.NotebookEucConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotebookEucConfig{}
	out.EucDisabled = direct.ValueOf(in.EUCDisabled)
	// MISSING: BypassActasCheck
	return out
}
func NotebookEUCConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotebookEucConfig) *krm.NotebookEUCConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NotebookEUCConfigObservedState{}
	// MISSING: EUCDisabled
	out.BypassActasCheck = direct.LazyPtr(in.GetBypassActasCheck())
	return out
}
func NotebookEUCConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NotebookEUCConfigObservedState) *pb.NotebookEucConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotebookEucConfig{}
	// MISSING: EUCDisabled
	out.BypassActasCheck = direct.ValueOf(in.BypassActasCheck)
	return out
}
func NotebookIdleShutdownConfig_FromProto(mapCtx *direct.MapContext, in *pb.NotebookIdleShutdownConfig) *krm.NotebookIdleShutdownConfig {
	if in == nil {
		return nil
	}
	out := &krm.NotebookIdleShutdownConfig{}
	out.IdleTimeout = direct.StringDuration_FromProto(mapCtx, in.GetIdleTimeout())
	out.IdleShutdownDisabled = direct.LazyPtr(in.GetIdleShutdownDisabled())
	return out
}
func NotebookIdleShutdownConfig_ToProto(mapCtx *direct.MapContext, in *krm.NotebookIdleShutdownConfig) *pb.NotebookIdleShutdownConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotebookIdleShutdownConfig{}
	out.IdleTimeout = direct.StringDuration_ToProto(mapCtx, in.IdleTimeout)
	out.IdleShutdownDisabled = direct.ValueOf(in.IdleShutdownDisabled)
	return out
}
func PersistentDiskSpec_FromProto(mapCtx *direct.MapContext, in *pb.PersistentDiskSpec) *krm.PersistentDiskSpec {
	if in == nil {
		return nil
	}
	out := &krm.PersistentDiskSpec{}
	out.DiskType = direct.LazyPtr(in.GetDiskType())
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	return out
}
func PersistentDiskSpec_ToProto(mapCtx *direct.MapContext, in *krm.PersistentDiskSpec) *pb.PersistentDiskSpec {
	if in == nil {
		return nil
	}
	out := &pb.PersistentDiskSpec{}
	out.DiskType = direct.ValueOf(in.DiskType)
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	return out
}
func ReservationAffinity_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krm.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krm.ReservationAffinity{}
	out.ReservationAffinityType = direct.Enum_FromProto(mapCtx, in.GetReservationAffinityType())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Values = in.Values
	return out
}
func ReservationAffinity_ToProto(mapCtx *direct.MapContext, in *krm.ReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ReservationAffinityType = direct.Enum_ToProto[pb.ReservationAffinity_Type](mapCtx, in.ReservationAffinityType)
	out.Key = direct.ValueOf(in.Key)
	out.Values = in.Values
	return out
}
func ShieldedVMConfig_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedVmConfig) *krm.ShieldedVMConfig {
	if in == nil {
		return nil
	}
	out := &krm.ShieldedVMConfig{}
	out.EnableSecureBoot = direct.LazyPtr(in.GetEnableSecureBoot())
	return out
}
func ShieldedVMConfig_ToProto(mapCtx *direct.MapContext, in *krm.ShieldedVMConfig) *pb.ShieldedVmConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedVmConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
	return out
}
func VertexAIFeaturestoreObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.VertexAIFeaturestoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIFeaturestoreObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIFeaturestoreObservedState) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Featurestore_State](mapCtx, in.State)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreSpec_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krm.VertexAIFeaturestoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIFeaturestoreSpec{}
	// MISSING: Name
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.Labels = in.Labels
	out.OnlineServingConfig = Featurestore_OnlineServingConfig_FromProto(mapCtx, in.GetOnlineServingConfig())
	out.OnlineStorageTTLDays = direct.LazyPtr(in.GetOnlineStorageTtlDays())
	out.EncryptionSpec = EncryptionSpec_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIFeaturestoreSpec) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	out.Etag = direct.ValueOf(in.Etag)
	out.Labels = in.Labels
	out.OnlineServingConfig = Featurestore_OnlineServingConfig_ToProto(mapCtx, in.OnlineServingConfig)
	out.OnlineStorageTtlDays = direct.ValueOf(in.OnlineStorageTTLDays)
	out.EncryptionSpec = EncryptionSpec_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
