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

package netapp

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
)
func NetappStoragePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.StoragePool) *krm.NetappStoragePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetappStoragePoolObservedState{}
	// MISSING: Name
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	// MISSING: VolumeCapacityGib
	// MISSING: VolumeCount
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Network
	// MISSING: ActiveDirectory
	// MISSING: KMSConfig
	// MISSING: LdapEnabled
	// MISSING: PsaRange
	// MISSING: EncryptionType
	// MISSING: GlobalAccessAllowed
	// MISSING: AllowAutoTiering
	// MISSING: ReplicaZone
	// MISSING: Zone
	return out
}
func NetappStoragePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetappStoragePoolObservedState) *pb.StoragePool {
	if in == nil {
		return nil
	}
	out := &pb.StoragePool{}
	// MISSING: Name
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	// MISSING: VolumeCapacityGib
	// MISSING: VolumeCount
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Network
	// MISSING: ActiveDirectory
	// MISSING: KMSConfig
	// MISSING: LdapEnabled
	// MISSING: PsaRange
	// MISSING: EncryptionType
	// MISSING: GlobalAccessAllowed
	// MISSING: AllowAutoTiering
	// MISSING: ReplicaZone
	// MISSING: Zone
	return out
}
func NetappStoragePoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.StoragePool) *krm.NetappStoragePoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetappStoragePoolSpec{}
	// MISSING: Name
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	// MISSING: VolumeCapacityGib
	// MISSING: VolumeCount
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Network
	// MISSING: ActiveDirectory
	// MISSING: KMSConfig
	// MISSING: LdapEnabled
	// MISSING: PsaRange
	// MISSING: EncryptionType
	// MISSING: GlobalAccessAllowed
	// MISSING: AllowAutoTiering
	// MISSING: ReplicaZone
	// MISSING: Zone
	return out
}
func NetappStoragePoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetappStoragePoolSpec) *pb.StoragePool {
	if in == nil {
		return nil
	}
	out := &pb.StoragePool{}
	// MISSING: Name
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	// MISSING: VolumeCapacityGib
	// MISSING: VolumeCount
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Network
	// MISSING: ActiveDirectory
	// MISSING: KMSConfig
	// MISSING: LdapEnabled
	// MISSING: PsaRange
	// MISSING: EncryptionType
	// MISSING: GlobalAccessAllowed
	// MISSING: AllowAutoTiering
	// MISSING: ReplicaZone
	// MISSING: Zone
	return out
}
func StoragePool_FromProto(mapCtx *direct.MapContext, in *pb.StoragePool) *krm.StoragePool {
	if in == nil {
		return nil
	}
	out := &krm.StoragePool{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ServiceLevel = direct.Enum_FromProto(mapCtx, in.GetServiceLevel())
	out.CapacityGib = direct.LazyPtr(in.GetCapacityGib())
	// MISSING: VolumeCapacityGib
	// MISSING: VolumeCount
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	out.Network = direct.LazyPtr(in.GetNetwork())
	out.ActiveDirectory = direct.LazyPtr(in.GetActiveDirectory())
	out.KMSConfig = direct.LazyPtr(in.GetKmsConfig())
	out.LdapEnabled = direct.LazyPtr(in.GetLdapEnabled())
	out.PsaRange = direct.LazyPtr(in.GetPsaRange())
	// MISSING: EncryptionType
	out.GlobalAccessAllowed = in.GlobalAccessAllowed
	out.AllowAutoTiering = direct.LazyPtr(in.GetAllowAutoTiering())
	out.ReplicaZone = direct.LazyPtr(in.GetReplicaZone())
	out.Zone = direct.LazyPtr(in.GetZone())
	return out
}
func StoragePool_ToProto(mapCtx *direct.MapContext, in *krm.StoragePool) *pb.StoragePool {
	if in == nil {
		return nil
	}
	out := &pb.StoragePool{}
	out.Name = direct.ValueOf(in.Name)
	out.ServiceLevel = direct.Enum_ToProto[pb.ServiceLevel](mapCtx, in.ServiceLevel)
	out.CapacityGib = direct.ValueOf(in.CapacityGib)
	// MISSING: VolumeCapacityGib
	// MISSING: VolumeCount
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	out.Network = direct.ValueOf(in.Network)
	out.ActiveDirectory = direct.ValueOf(in.ActiveDirectory)
	out.KmsConfig = direct.ValueOf(in.KMSConfig)
	out.LdapEnabled = direct.ValueOf(in.LdapEnabled)
	out.PsaRange = direct.ValueOf(in.PsaRange)
	// MISSING: EncryptionType
	out.GlobalAccessAllowed = in.GlobalAccessAllowed
	out.AllowAutoTiering = direct.ValueOf(in.AllowAutoTiering)
	out.ReplicaZone = direct.ValueOf(in.ReplicaZone)
	out.Zone = direct.ValueOf(in.Zone)
	return out
}
func StoragePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.StoragePool) *krm.StoragePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StoragePoolObservedState{}
	// MISSING: Name
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	out.VolumeCapacityGib = direct.LazyPtr(in.GetVolumeCapacityGib())
	out.VolumeCount = direct.LazyPtr(in.GetVolumeCount())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateDetails = direct.LazyPtr(in.GetStateDetails())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Network
	// MISSING: ActiveDirectory
	// MISSING: KMSConfig
	// MISSING: LdapEnabled
	// MISSING: PsaRange
	out.EncryptionType = direct.Enum_FromProto(mapCtx, in.GetEncryptionType())
	// MISSING: GlobalAccessAllowed
	// MISSING: AllowAutoTiering
	// MISSING: ReplicaZone
	// MISSING: Zone
	return out
}
func StoragePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StoragePoolObservedState) *pb.StoragePool {
	if in == nil {
		return nil
	}
	out := &pb.StoragePool{}
	// MISSING: Name
	// MISSING: ServiceLevel
	// MISSING: CapacityGib
	out.VolumeCapacityGib = direct.ValueOf(in.VolumeCapacityGib)
	out.VolumeCount = direct.ValueOf(in.VolumeCount)
	out.State = direct.Enum_ToProto[pb.StoragePool_State](mapCtx, in.State)
	out.StateDetails = direct.ValueOf(in.StateDetails)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Network
	// MISSING: ActiveDirectory
	// MISSING: KMSConfig
	// MISSING: LdapEnabled
	// MISSING: PsaRange
	out.EncryptionType = direct.Enum_ToProto[pb.EncryptionType](mapCtx, in.EncryptionType)
	// MISSING: GlobalAccessAllowed
	// MISSING: AllowAutoTiering
	// MISSING: ReplicaZone
	// MISSING: Zone
	return out
}
