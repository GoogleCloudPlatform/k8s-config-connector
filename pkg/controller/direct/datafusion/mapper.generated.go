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

package datafusion

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/datafusion/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/datafusion/apiv1/datafusionpb"
)
func Accelerator_FromProto(mapCtx *direct.MapContext, in *pb.Accelerator) *krm.Accelerator {
	if in == nil {
		return nil
	}
	out := &krm.Accelerator{}
	out.AcceleratorType = direct.Enum_FromProto(mapCtx, in.GetAcceleratorType())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func Accelerator_ToProto(mapCtx *direct.MapContext, in *krm.Accelerator) *pb.Accelerator {
	if in == nil {
		return nil
	}
	out := &pb.Accelerator{}
	out.AcceleratorType = direct.Enum_ToProto[pb.Accelerator_AcceleratorType](mapCtx, in.AcceleratorType)
	out.State = direct.Enum_ToProto[pb.Accelerator_State](mapCtx, in.State)
	return out
}
func CryptoKeyConfig_FromProto(mapCtx *direct.MapContext, in *pb.CryptoKeyConfig) *krm.CryptoKeyConfig {
	if in == nil {
		return nil
	}
	out := &krm.CryptoKeyConfig{}
	out.KeyReference = direct.LazyPtr(in.GetKeyReference())
	return out
}
func CryptoKeyConfig_ToProto(mapCtx *direct.MapContext, in *krm.CryptoKeyConfig) *pb.CryptoKeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.CryptoKeyConfig{}
	out.KeyReference = direct.ValueOf(in.KeyReference)
	return out
}
func DatafusionInstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.DatafusionInstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatafusionInstanceObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Type
	// MISSING: EnableStackdriverLogging
	// MISSING: EnableStackdriverMonitoring
	// MISSING: PrivateInstance
	// MISSING: NetworkConfig
	// MISSING: Labels
	// MISSING: Options
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: ServiceEndpoint
	// MISSING: Zone
	// MISSING: Version
	// MISSING: ServiceAccount
	// MISSING: DisplayName
	// MISSING: AvailableVersion
	// MISSING: ApiEndpoint
	// MISSING: GcsBucket
	// MISSING: Accelerators
	// MISSING: P4ServiceAccount
	// MISSING: TenantProjectID
	// MISSING: DataprocServiceAccount
	// MISSING: EnableRbac
	// MISSING: CryptoKeyConfig
	// MISSING: DisabledReason
	return out
}
func DatafusionInstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatafusionInstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Type
	// MISSING: EnableStackdriverLogging
	// MISSING: EnableStackdriverMonitoring
	// MISSING: PrivateInstance
	// MISSING: NetworkConfig
	// MISSING: Labels
	// MISSING: Options
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: ServiceEndpoint
	// MISSING: Zone
	// MISSING: Version
	// MISSING: ServiceAccount
	// MISSING: DisplayName
	// MISSING: AvailableVersion
	// MISSING: ApiEndpoint
	// MISSING: GcsBucket
	// MISSING: Accelerators
	// MISSING: P4ServiceAccount
	// MISSING: TenantProjectID
	// MISSING: DataprocServiceAccount
	// MISSING: EnableRbac
	// MISSING: CryptoKeyConfig
	// MISSING: DisabledReason
	return out
}
func DatafusionInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.DatafusionInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.DatafusionInstanceSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Type
	// MISSING: EnableStackdriverLogging
	// MISSING: EnableStackdriverMonitoring
	// MISSING: PrivateInstance
	// MISSING: NetworkConfig
	// MISSING: Labels
	// MISSING: Options
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: ServiceEndpoint
	// MISSING: Zone
	// MISSING: Version
	// MISSING: ServiceAccount
	// MISSING: DisplayName
	// MISSING: AvailableVersion
	// MISSING: ApiEndpoint
	// MISSING: GcsBucket
	// MISSING: Accelerators
	// MISSING: P4ServiceAccount
	// MISSING: TenantProjectID
	// MISSING: DataprocServiceAccount
	// MISSING: EnableRbac
	// MISSING: CryptoKeyConfig
	// MISSING: DisabledReason
	return out
}
func DatafusionInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.DatafusionInstanceSpec) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Type
	// MISSING: EnableStackdriverLogging
	// MISSING: EnableStackdriverMonitoring
	// MISSING: PrivateInstance
	// MISSING: NetworkConfig
	// MISSING: Labels
	// MISSING: Options
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: ServiceEndpoint
	// MISSING: Zone
	// MISSING: Version
	// MISSING: ServiceAccount
	// MISSING: DisplayName
	// MISSING: AvailableVersion
	// MISSING: ApiEndpoint
	// MISSING: GcsBucket
	// MISSING: Accelerators
	// MISSING: P4ServiceAccount
	// MISSING: TenantProjectID
	// MISSING: DataprocServiceAccount
	// MISSING: EnableRbac
	// MISSING: CryptoKeyConfig
	// MISSING: DisabledReason
	return out
}
func Instance_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.Instance {
	if in == nil {
		return nil
	}
	out := &krm.Instance{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.EnableStackdriverLogging = direct.LazyPtr(in.GetEnableStackdriverLogging())
	out.EnableStackdriverMonitoring = direct.LazyPtr(in.GetEnableStackdriverMonitoring())
	out.PrivateInstance = direct.LazyPtr(in.GetPrivateInstance())
	out.NetworkConfig = NetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	out.Labels = in.Labels
	out.Options = in.Options
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: ServiceEndpoint
	out.Zone = direct.LazyPtr(in.GetZone())
	out.Version = direct.LazyPtr(in.GetVersion())
	// MISSING: ServiceAccount
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.AvailableVersion = direct.Slice_FromProto(mapCtx, in.AvailableVersion, Version_FromProto)
	// MISSING: ApiEndpoint
	// MISSING: GcsBucket
	out.Accelerators = direct.Slice_FromProto(mapCtx, in.Accelerators, Accelerator_FromProto)
	// MISSING: P4ServiceAccount
	// MISSING: TenantProjectID
	out.DataprocServiceAccount = direct.LazyPtr(in.GetDataprocServiceAccount())
	out.EnableRbac = direct.LazyPtr(in.GetEnableRbac())
	out.CryptoKeyConfig = CryptoKeyConfig_FromProto(mapCtx, in.GetCryptoKeyConfig())
	// MISSING: DisabledReason
	return out
}
func Instance_ToProto(mapCtx *direct.MapContext, in *krm.Instance) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.Type = direct.Enum_ToProto[pb.Instance_Type](mapCtx, in.Type)
	out.EnableStackdriverLogging = direct.ValueOf(in.EnableStackdriverLogging)
	out.EnableStackdriverMonitoring = direct.ValueOf(in.EnableStackdriverMonitoring)
	out.PrivateInstance = direct.ValueOf(in.PrivateInstance)
	out.NetworkConfig = NetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	out.Labels = in.Labels
	out.Options = in.Options
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateMessage
	// MISSING: ServiceEndpoint
	out.Zone = direct.ValueOf(in.Zone)
	out.Version = direct.ValueOf(in.Version)
	// MISSING: ServiceAccount
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.AvailableVersion = direct.Slice_ToProto(mapCtx, in.AvailableVersion, Version_ToProto)
	// MISSING: ApiEndpoint
	// MISSING: GcsBucket
	out.Accelerators = direct.Slice_ToProto(mapCtx, in.Accelerators, Accelerator_ToProto)
	// MISSING: P4ServiceAccount
	// MISSING: TenantProjectID
	out.DataprocServiceAccount = direct.ValueOf(in.DataprocServiceAccount)
	out.EnableRbac = direct.ValueOf(in.EnableRbac)
	out.CryptoKeyConfig = CryptoKeyConfig_ToProto(mapCtx, in.CryptoKeyConfig)
	// MISSING: DisabledReason
	return out
}
func InstanceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.InstanceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: Type
	// MISSING: EnableStackdriverLogging
	// MISSING: EnableStackdriverMonitoring
	// MISSING: PrivateInstance
	// MISSING: NetworkConfig
	// MISSING: Labels
	// MISSING: Options
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateMessage = direct.LazyPtr(in.GetStateMessage())
	out.ServiceEndpoint = direct.LazyPtr(in.GetServiceEndpoint())
	// MISSING: Zone
	// MISSING: Version
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	// MISSING: DisplayName
	// MISSING: AvailableVersion
	out.ApiEndpoint = direct.LazyPtr(in.GetApiEndpoint())
	out.GcsBucket = direct.LazyPtr(in.GetGcsBucket())
	// MISSING: Accelerators
	out.P4ServiceAccount = direct.LazyPtr(in.GetP4ServiceAccount())
	out.TenantProjectID = direct.LazyPtr(in.GetTenantProjectId())
	// MISSING: DataprocServiceAccount
	// MISSING: EnableRbac
	// MISSING: CryptoKeyConfig
	out.DisabledReason = direct.EnumSlice_FromProto(mapCtx, in.DisabledReason)
	return out
}
func InstanceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedState) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: Type
	// MISSING: EnableStackdriverLogging
	// MISSING: EnableStackdriverMonitoring
	// MISSING: PrivateInstance
	// MISSING: NetworkConfig
	// MISSING: Labels
	// MISSING: Options
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Instance_State](mapCtx, in.State)
	out.StateMessage = direct.ValueOf(in.StateMessage)
	out.ServiceEndpoint = direct.ValueOf(in.ServiceEndpoint)
	// MISSING: Zone
	// MISSING: Version
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	// MISSING: DisplayName
	// MISSING: AvailableVersion
	out.ApiEndpoint = direct.ValueOf(in.ApiEndpoint)
	out.GcsBucket = direct.ValueOf(in.GcsBucket)
	// MISSING: Accelerators
	out.P4ServiceAccount = direct.ValueOf(in.P4ServiceAccount)
	out.TenantProjectId = direct.ValueOf(in.TenantProjectID)
	// MISSING: DataprocServiceAccount
	// MISSING: EnableRbac
	// MISSING: CryptoKeyConfig
	out.DisabledReason = direct.EnumSlice_ToProto[pb.Instance_DisabledReason](mapCtx, in.DisabledReason)
	return out
}
func NetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NetworkConfig) *krm.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NetworkConfig{}
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.IPAllocation = direct.LazyPtr(in.GetIpAllocation())
	return out
}
func NetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NetworkConfig) *pb.NetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NetworkConfig{}
	out.Network = direct.ValueOf(in.Network)
	out.IpAllocation = direct.ValueOf(in.IPAllocation)
	return out
}
func Version_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.Version {
	if in == nil {
		return nil
	}
	out := &krm.Version{}
	out.VersionNumber = direct.LazyPtr(in.GetVersionNumber())
	out.DefaultVersion = direct.LazyPtr(in.GetDefaultVersion())
	out.AvailableFeatures = in.AvailableFeatures
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func Version_ToProto(mapCtx *direct.MapContext, in *krm.Version) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	out.VersionNumber = direct.ValueOf(in.VersionNumber)
	out.DefaultVersion = direct.ValueOf(in.DefaultVersion)
	out.AvailableFeatures = in.AvailableFeatures
	out.Type = direct.Enum_ToProto[pb.Version_Type](mapCtx, in.Type)
	return out
}
