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
// krm.group: vertexai.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.aiplatform.v1beta1

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krmcolabv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/colab/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krmvertexaiv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ColabImage_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ColabImage) *krmcolabv1alpha1.ColabImage {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.ColabImage{}
	out.ReleaseName = direct.LazyPtr(in.GetReleaseName())
	// MISSING: Description
	return out
}
func ColabImage_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.ColabImage) *pb.ColabImage {
	if in == nil {
		return nil
	}
	out := &pb.ColabImage{}
	out.ReleaseName = direct.ValueOf(in.ReleaseName)
	// MISSING: Description
	return out
}
func ColabImageObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ColabImage) *krmcolabv1alpha1.ColabImageObservedState {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.ColabImageObservedState{}
	// MISSING: ReleaseName
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func ColabImageObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.ColabImageObservedState) *pb.ColabImage {
	if in == nil {
		return nil
	}
	out := &pb.ColabImage{}
	// MISSING: ReleaseName
	out.Description = direct.ValueOf(in.Description)
	return out
}
func ColabRuntimeObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntime) *krmcolabv1alpha1.ColabRuntimeObservedState {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.ColabRuntimeObservedState{}
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
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_v1alpha1_FromProto(mapCtx, in.GetIdleShutdownConfig())
	// MISSING: EUCConfig
	// MISSING: ShieldedVMConfig
	// MISSING: SoftwareConfig
	out.EncryptionSpec = EncryptionSpecObservedState_v1alpha1_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ColabRuntimeObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.ColabRuntimeObservedState) *pb.NotebookRuntime {
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
	// MISSING: MachineSpec
	// MISSING: DataPersistentDiskSpec
	// MISSING: NetworkSpec
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_v1alpha1_ToProto(mapCtx, in.IdleShutdownConfig)
	// MISSING: EUCConfig
	// MISSING: ShieldedVMConfig
	// MISSING: SoftwareConfig
	out.EncryptionSpec = EncryptionSpecObservedState_v1alpha1_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func ColabRuntimeSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntime) *krmcolabv1alpha1.ColabRuntimeSpec {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.ColabRuntimeSpec{}
	// MISSING: Name
	out.RuntimeUser = direct.LazyPtr(in.GetRuntimeUser())
	// MISSING: NotebookRuntimeTemplateRef
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
func ColabRuntimeSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.ColabRuntimeSpec) *pb.NotebookRuntime {
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
func ColabRuntimeTemplateObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeTemplate) *krmcolabv1alpha1.ColabRuntimeTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.ColabRuntimeTemplateObservedState{}
	// MISSING: Name
	// MISSING: IsDefault
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.EUCConfig = NotebookEUCConfigObservedState_v1alpha1_FromProto(mapCtx, in.GetEucConfig())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: SoftwareConfig
	return out
}
func ColabRuntimeTemplateObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.ColabRuntimeTemplateObservedState) *pb.NotebookRuntimeTemplate {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntimeTemplate{}
	// MISSING: Name
	// MISSING: IsDefault
	out.Etag = direct.ValueOf(in.Etag)
	out.EucConfig = NotebookEUCConfigObservedState_v1alpha1_ToProto(mapCtx, in.EUCConfig)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: SoftwareConfig
	return out
}
func ColabRuntimeTemplateSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookRuntimeTemplate) *krmcolabv1alpha1.ColabRuntimeTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.ColabRuntimeTemplateSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: IsDefault
	out.MachineSpec = MachineSpec_v1alpha1_FromProto(mapCtx, in.GetMachineSpec())
	out.DataPersistentDiskSpec = PersistentDiskSpec_v1alpha1_FromProto(mapCtx, in.GetDataPersistentDiskSpec())
	out.NetworkSpec = NetworkSpec_v1alpha1_FromProto(mapCtx, in.GetNetworkSpec())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	out.Labels = in.Labels
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_v1alpha1_FromProto(mapCtx, in.GetIdleShutdownConfig())
	out.EUCConfig = NotebookEUCConfig_v1alpha1_FromProto(mapCtx, in.GetEucConfig())
	out.NotebookRuntimeType = direct.Enum_FromProto(mapCtx, in.GetNotebookRuntimeType())
	out.ShieldedVMConfig = ShieldedVMConfig_v1alpha1_FromProto(mapCtx, in.GetShieldedVmConfig())
	out.NetworkTags = in.NetworkTags
	out.EncryptionSpec = EncryptionSpec_v1alpha1_FromProto(mapCtx, in.GetEncryptionSpec())
	// MISSING: SoftwareConfig
	return out
}
func ColabRuntimeTemplateSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.ColabRuntimeTemplateSpec) *pb.NotebookRuntimeTemplate {
	if in == nil {
		return nil
	}
	out := &pb.NotebookRuntimeTemplate{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: IsDefault
	out.MachineSpec = MachineSpec_v1alpha1_ToProto(mapCtx, in.MachineSpec)
	out.DataPersistentDiskSpec = PersistentDiskSpec_v1alpha1_ToProto(mapCtx, in.DataPersistentDiskSpec)
	out.NetworkSpec = NetworkSpec_v1alpha1_ToProto(mapCtx, in.NetworkSpec)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Labels = in.Labels
	out.IdleShutdownConfig = NotebookIdleShutdownConfig_v1alpha1_ToProto(mapCtx, in.IdleShutdownConfig)
	out.EucConfig = NotebookEUCConfig_v1alpha1_ToProto(mapCtx, in.EUCConfig)
	out.NotebookRuntimeType = direct.Enum_ToProto[pb.NotebookRuntimeType](mapCtx, in.NotebookRuntimeType)
	out.ShieldedVmConfig = ShieldedVMConfig_v1alpha1_ToProto(mapCtx, in.ShieldedVMConfig)
	out.NetworkTags = in.NetworkTags
	out.EncryptionSpec = EncryptionSpec_v1alpha1_ToProto(mapCtx, in.EncryptionSpec)
	// MISSING: SoftwareConfig
	return out
}
func EncryptionSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krmcolabv1alpha1.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.EncryptionSpec{}
	// MISSING: KMSKeyName
	return out
}
func EncryptionSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	// MISSING: KMSKeyName
	return out
}
func EncryptionSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krm.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionSpec{}
	// MISSING: KMSKeyName
	return out
}
func EncryptionSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionSpec) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	// MISSING: KMSKeyName
	return out
}
func EncryptionSpecObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionSpec) *krmcolabv1alpha1.EncryptionSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.EncryptionSpecObservedState{}
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	return out
}
func EncryptionSpecObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.EncryptionSpecObservedState) *pb.EncryptionSpec {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionSpec{}
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	return out
}
func EnvVar_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.EnvVar) *krmcolabv1alpha1.EnvVar {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.EnvVar{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func EnvVar_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.EnvVar) *pb.EnvVar {
	if in == nil {
		return nil
	}
	out := &pb.EnvVar{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func Featurestore_OnlineServingConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore_OnlineServingConfig) *krmvertexaiv1alpha1.Featurestore_OnlineServingConfig {
	if in == nil {
		return nil
	}
	out := &krmvertexaiv1alpha1.Featurestore_OnlineServingConfig{}
	out.FixedNodeCount = direct.LazyPtr(in.GetFixedNodeCount())
	out.Scaling = Featurestore_OnlineServingConfig_Scaling_v1alpha1_FromProto(mapCtx, in.GetScaling())
	return out
}
func Featurestore_OnlineServingConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmvertexaiv1alpha1.Featurestore_OnlineServingConfig) *pb.Featurestore_OnlineServingConfig {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore_OnlineServingConfig{}
	out.FixedNodeCount = direct.ValueOf(in.FixedNodeCount)
	out.Scaling = Featurestore_OnlineServingConfig_Scaling_v1alpha1_ToProto(mapCtx, in.Scaling)
	return out
}
func Featurestore_OnlineServingConfig_Scaling_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore_OnlineServingConfig_Scaling) *krmvertexaiv1alpha1.Featurestore_OnlineServingConfig_Scaling {
	if in == nil {
		return nil
	}
	out := &krmvertexaiv1alpha1.Featurestore_OnlineServingConfig_Scaling{}
	out.MinNodeCount = direct.LazyPtr(in.GetMinNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	out.CPUUtilizationTarget = direct.LazyPtr(in.GetCpuUtilizationTarget())
	return out
}
func Featurestore_OnlineServingConfig_Scaling_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmvertexaiv1alpha1.Featurestore_OnlineServingConfig_Scaling) *pb.Featurestore_OnlineServingConfig_Scaling {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore_OnlineServingConfig_Scaling{}
	out.MinNodeCount = direct.ValueOf(in.MinNodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	out.CpuUtilizationTarget = direct.ValueOf(in.CPUUtilizationTarget)
	return out
}
func MachineSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.MachineSpec) *krmcolabv1alpha1.MachineSpec {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.MachineSpec{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.AcceleratorType = direct.Enum_FromProto(mapCtx, in.GetAcceleratorType())
	out.AcceleratorCount = direct.LazyPtr(in.GetAcceleratorCount())
	out.GpuPartitionSize = direct.LazyPtr(in.GetGpuPartitionSize())
	out.TpuTopology = direct.LazyPtr(in.GetTpuTopology())
	out.MultihostGpuNodeCount = direct.LazyPtr(in.GetMultihostGpuNodeCount())
	out.ReservationAffinity = ReservationAffinity_v1alpha1_FromProto(mapCtx, in.GetReservationAffinity())
	return out
}
func MachineSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.MachineSpec) *pb.MachineSpec {
	if in == nil {
		return nil
	}
	out := &pb.MachineSpec{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.AcceleratorType = direct.Enum_ToProto[pb.AcceleratorType](mapCtx, in.AcceleratorType)
	out.AcceleratorCount = direct.ValueOf(in.AcceleratorCount)
	out.GpuPartitionSize = direct.ValueOf(in.GpuPartitionSize)
	out.TpuTopology = direct.ValueOf(in.TpuTopology)
	out.MultihostGpuNodeCount = direct.ValueOf(in.MultihostGpuNodeCount)
	out.ReservationAffinity = ReservationAffinity_v1alpha1_ToProto(mapCtx, in.ReservationAffinity)
	return out
}
func MetadataStore_DataplexConfig_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore_DataplexConfig) *krm.MetadataStore_DataplexConfig {
	if in == nil {
		return nil
	}
	out := &krm.MetadataStore_DataplexConfig{}
	out.EnabledPipelinesLineage = direct.LazyPtr(in.GetEnabledPipelinesLineage())
	return out
}
func MetadataStore_DataplexConfig_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.MetadataStore_DataplexConfig) *pb.MetadataStore_DataplexConfig {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore_DataplexConfig{}
	out.EnabledPipelinesLineage = direct.ValueOf(in.EnabledPipelinesLineage)
	return out
}
func MetadataStore_MetadataStoreState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore_MetadataStoreState) *krm.MetadataStore_MetadataStoreState {
	if in == nil {
		return nil
	}
	out := &krm.MetadataStore_MetadataStoreState{}
	out.DiskUtilizationBytes = direct.LazyPtr(in.GetDiskUtilizationBytes())
	return out
}
func MetadataStore_MetadataStoreState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.MetadataStore_MetadataStoreState) *pb.MetadataStore_MetadataStoreState {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore_MetadataStoreState{}
	out.DiskUtilizationBytes = direct.ValueOf(in.DiskUtilizationBytes)
	return out
}
func NetworkSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NetworkSpec) *krmcolabv1alpha1.NetworkSpec {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.NetworkSpec{}
	out.EnableInternetAccess = direct.LazyPtr(in.GetEnableInternetAccess())
	if in.GetNetwork() != "" {
		out.NetworkRef = &refsv1beta1.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &refsv1beta1.ComputeSubnetworkRef{External: in.GetSubnetwork()}
	}
	return out
}
func NetworkSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.NetworkSpec) *pb.NetworkSpec {
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
func NotebookEUCConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookEucConfig) *krmcolabv1alpha1.NotebookEUCConfig {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.NotebookEUCConfig{}
	out.EUCDisabled = direct.LazyPtr(in.GetEucDisabled())
	// MISSING: BypassActasCheck
	return out
}
func NotebookEUCConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.NotebookEUCConfig) *pb.NotebookEucConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotebookEucConfig{}
	out.EucDisabled = direct.ValueOf(in.EUCDisabled)
	// MISSING: BypassActasCheck
	return out
}
func NotebookEUCConfigObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookEucConfig) *krmcolabv1alpha1.NotebookEUCConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.NotebookEUCConfigObservedState{}
	// MISSING: EUCDisabled
	out.BypassActasCheck = direct.LazyPtr(in.GetBypassActasCheck())
	return out
}
func NotebookEUCConfigObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.NotebookEUCConfigObservedState) *pb.NotebookEucConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotebookEucConfig{}
	// MISSING: EUCDisabled
	out.BypassActasCheck = direct.ValueOf(in.BypassActasCheck)
	return out
}
func NotebookIdleShutdownConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookIdleShutdownConfig) *krmcolabv1alpha1.NotebookIdleShutdownConfig {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.NotebookIdleShutdownConfig{}
	out.IdleTimeout = direct.StringDuration_FromProto(mapCtx, in.GetIdleTimeout())
	out.IdleShutdownDisabled = direct.LazyPtr(in.GetIdleShutdownDisabled())
	return out
}
func NotebookIdleShutdownConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.NotebookIdleShutdownConfig) *pb.NotebookIdleShutdownConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotebookIdleShutdownConfig{}
	out.IdleTimeout = direct.StringDuration_ToProto(mapCtx, in.IdleTimeout)
	out.IdleShutdownDisabled = direct.ValueOf(in.IdleShutdownDisabled)
	return out
}
func NotebookSoftwareConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookSoftwareConfig) *krmcolabv1alpha1.NotebookSoftwareConfig {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.NotebookSoftwareConfig{}
	out.ColabImage = ColabImage_v1alpha1_FromProto(mapCtx, in.GetColabImage())
	out.Env = direct.Slice_FromProto(mapCtx, in.Env, EnvVar_v1alpha1_FromProto)
	out.PostStartupScriptConfig = PostStartupScriptConfig_v1alpha1_FromProto(mapCtx, in.GetPostStartupScriptConfig())
	return out
}
func NotebookSoftwareConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.NotebookSoftwareConfig) *pb.NotebookSoftwareConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotebookSoftwareConfig{}
	if oneof := ColabImage_v1alpha1_ToProto(mapCtx, in.ColabImage); oneof != nil {
		out.RuntimeImage = &pb.NotebookSoftwareConfig_ColabImage{ColabImage: oneof}
	}
	out.Env = direct.Slice_ToProto(mapCtx, in.Env, EnvVar_v1alpha1_ToProto)
	out.PostStartupScriptConfig = PostStartupScriptConfig_v1alpha1_ToProto(mapCtx, in.PostStartupScriptConfig)
	return out
}
func NotebookSoftwareConfigObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.NotebookSoftwareConfig) *krmcolabv1alpha1.NotebookSoftwareConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.NotebookSoftwareConfigObservedState{}
	out.ColabImage = ColabImageObservedState_v1alpha1_FromProto(mapCtx, in.GetColabImage())
	// MISSING: Env
	// MISSING: PostStartupScriptConfig
	return out
}
func NotebookSoftwareConfigObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.NotebookSoftwareConfigObservedState) *pb.NotebookSoftwareConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotebookSoftwareConfig{}
	if oneof := ColabImageObservedState_v1alpha1_ToProto(mapCtx, in.ColabImage); oneof != nil {
		out.RuntimeImage = &pb.NotebookSoftwareConfig_ColabImage{ColabImage: oneof}
	}
	// MISSING: Env
	// MISSING: PostStartupScriptConfig
	return out
}
func PersistentDiskSpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PersistentDiskSpec) *krmcolabv1alpha1.PersistentDiskSpec {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.PersistentDiskSpec{}
	out.DiskType = direct.LazyPtr(in.GetDiskType())
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	return out
}
func PersistentDiskSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.PersistentDiskSpec) *pb.PersistentDiskSpec {
	if in == nil {
		return nil
	}
	out := &pb.PersistentDiskSpec{}
	out.DiskType = direct.ValueOf(in.DiskType)
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	return out
}
func PostStartupScriptConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.PostStartupScriptConfig) *krmcolabv1alpha1.PostStartupScriptConfig {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.PostStartupScriptConfig{}
	out.PostStartupScript = direct.LazyPtr(in.GetPostStartupScript())
	out.PostStartupScriptURL = direct.LazyPtr(in.GetPostStartupScriptUrl())
	out.PostStartupScriptBehavior = direct.Enum_FromProto(mapCtx, in.GetPostStartupScriptBehavior())
	return out
}
func PostStartupScriptConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.PostStartupScriptConfig) *pb.PostStartupScriptConfig {
	if in == nil {
		return nil
	}
	out := &pb.PostStartupScriptConfig{}
	out.PostStartupScript = direct.ValueOf(in.PostStartupScript)
	out.PostStartupScriptUrl = direct.ValueOf(in.PostStartupScriptURL)
	out.PostStartupScriptBehavior = direct.Enum_ToProto[pb.PostStartupScriptConfig_PostStartupScriptBehavior](mapCtx, in.PostStartupScriptBehavior)
	return out
}
func ReservationAffinity_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krmcolabv1alpha1.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.ReservationAffinity{}
	out.ReservationAffinityType = direct.Enum_FromProto(mapCtx, in.GetReservationAffinityType())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Values = in.Values
	return out
}
func ReservationAffinity_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.ReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ReservationAffinityType = direct.Enum_ToProto[pb.ReservationAffinity_Type](mapCtx, in.ReservationAffinityType)
	out.Key = direct.ValueOf(in.Key)
	out.Values = in.Values
	return out
}
func ShieldedVMConfig_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedVmConfig) *krmcolabv1alpha1.ShieldedVMConfig {
	if in == nil {
		return nil
	}
	out := &krmcolabv1alpha1.ShieldedVMConfig{}
	out.EnableSecureBoot = direct.LazyPtr(in.GetEnableSecureBoot())
	return out
}
func ShieldedVMConfig_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcolabv1alpha1.ShieldedVMConfig) *pb.ShieldedVmConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedVmConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
	return out
}
func VertexAIFeaturestoreObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.Featurestore) *krmvertexaiv1alpha1.VertexAIFeaturestoreObservedState {
	if in == nil {
		return nil
	}
	out := &krmvertexaiv1alpha1.VertexAIFeaturestoreObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Etag
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIFeaturestoreObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmvertexaiv1alpha1.VertexAIFeaturestoreObservedState) *pb.Featurestore {
	if in == nil {
		return nil
	}
	out := &pb.Featurestore{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Etag
	out.State = direct.Enum_ToProto[pb.Featurestore_State](mapCtx, in.State)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func VertexAIMetadataStoreObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore) *krm.VertexAIMetadataStoreObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIMetadataStoreObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = MetadataStore_MetadataStoreState_v1beta1_FromProto(mapCtx, in.GetState())
	return out
}
func VertexAIMetadataStoreObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIMetadataStoreObservedState) *pb.MetadataStore {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = MetadataStore_MetadataStoreState_v1beta1_ToProto(mapCtx, in.State)
	return out
}
func VertexAIMetadataStoreSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.MetadataStore) *krm.VertexAIMetadataStoreSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIMetadataStoreSpec{}
	out.EncryptionSpec = EncryptionSpec_v1beta1_FromProto(mapCtx, in.GetEncryptionSpec())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DataplexConfig = MetadataStore_DataplexConfig_v1beta1_FromProto(mapCtx, in.GetDataplexConfig())
	return out
}
func VertexAIMetadataStoreSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIMetadataStoreSpec) *pb.MetadataStore {
	if in == nil {
		return nil
	}
	out := &pb.MetadataStore{}
	out.EncryptionSpec = EncryptionSpec_v1beta1_ToProto(mapCtx, in.EncryptionSpec)
	out.Description = direct.ValueOf(in.Description)
	out.DataplexConfig = MetadataStore_DataplexConfig_v1beta1_ToProto(mapCtx, in.DataplexConfig)
	return out
}
