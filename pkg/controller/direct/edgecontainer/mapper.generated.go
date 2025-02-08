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

package edgecontainer

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
)
func EdgecontainerNodePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodePool) *krm.EdgecontainerNodePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EdgecontainerNodePoolObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NodeLocation
	// MISSING: NodeCount
	// MISSING: MachineFilter
	// MISSING: LocalDiskEncryption
	// MISSING: NodeVersion
	// MISSING: NodeConfig
	return out
}
func EdgecontainerNodePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EdgecontainerNodePoolObservedState) *pb.NodePool {
	if in == nil {
		return nil
	}
	out := &pb.NodePool{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NodeLocation
	// MISSING: NodeCount
	// MISSING: MachineFilter
	// MISSING: LocalDiskEncryption
	// MISSING: NodeVersion
	// MISSING: NodeConfig
	return out
}
func EdgecontainerNodePoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.NodePool) *krm.EdgecontainerNodePoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.EdgecontainerNodePoolSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NodeLocation
	// MISSING: NodeCount
	// MISSING: MachineFilter
	// MISSING: LocalDiskEncryption
	// MISSING: NodeVersion
	// MISSING: NodeConfig
	return out
}
func EdgecontainerNodePoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.EdgecontainerNodePoolSpec) *pb.NodePool {
	if in == nil {
		return nil
	}
	out := &pb.NodePool{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: NodeLocation
	// MISSING: NodeCount
	// MISSING: MachineFilter
	// MISSING: LocalDiskEncryption
	// MISSING: NodeVersion
	// MISSING: NodeConfig
	return out
}
func NodePool_FromProto(mapCtx *direct.MapContext, in *pb.NodePool) *krm.NodePool {
	if in == nil {
		return nil
	}
	out := &krm.NodePool{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.NodeLocation = direct.LazyPtr(in.GetNodeLocation())
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.MachineFilter = direct.LazyPtr(in.GetMachineFilter())
	out.LocalDiskEncryption = NodePool_LocalDiskEncryption_FromProto(mapCtx, in.GetLocalDiskEncryption())
	// MISSING: NodeVersion
	out.NodeConfig = NodePool_NodeConfig_FromProto(mapCtx, in.GetNodeConfig())
	return out
}
func NodePool_ToProto(mapCtx *direct.MapContext, in *krm.NodePool) *pb.NodePool {
	if in == nil {
		return nil
	}
	out := &pb.NodePool{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.NodeLocation = direct.ValueOf(in.NodeLocation)
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.MachineFilter = direct.ValueOf(in.MachineFilter)
	out.LocalDiskEncryption = NodePool_LocalDiskEncryption_ToProto(mapCtx, in.LocalDiskEncryption)
	// MISSING: NodeVersion
	out.NodeConfig = NodePool_NodeConfig_ToProto(mapCtx, in.NodeConfig)
	return out
}
func NodePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodePool) *krm.NodePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NodePoolObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: NodeLocation
	// MISSING: NodeCount
	// MISSING: MachineFilter
	out.LocalDiskEncryption = NodePool_LocalDiskEncryptionObservedState_FromProto(mapCtx, in.GetLocalDiskEncryption())
	out.NodeVersion = direct.LazyPtr(in.GetNodeVersion())
	// MISSING: NodeConfig
	return out
}
func NodePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NodePoolObservedState) *pb.NodePool {
	if in == nil {
		return nil
	}
	out := &pb.NodePool{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: NodeLocation
	// MISSING: NodeCount
	// MISSING: MachineFilter
	out.LocalDiskEncryption = NodePool_LocalDiskEncryptionObservedState_ToProto(mapCtx, in.LocalDiskEncryption)
	out.NodeVersion = direct.ValueOf(in.NodeVersion)
	// MISSING: NodeConfig
	return out
}
func NodePool_LocalDiskEncryption_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_LocalDiskEncryption) *krm.NodePool_LocalDiskEncryption {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_LocalDiskEncryption{}
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	// MISSING: KMSKeyActiveVersion
	// MISSING: KMSKeyState
	// MISSING: KMSStatus
	// MISSING: ResourceState
	return out
}
func NodePool_LocalDiskEncryption_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_LocalDiskEncryption) *pb.NodePool_LocalDiskEncryption {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_LocalDiskEncryption{}
	out.KmsKey = direct.ValueOf(in.KMSKey)
	// MISSING: KMSKeyActiveVersion
	// MISSING: KMSKeyState
	// MISSING: KMSStatus
	// MISSING: ResourceState
	return out
}
func NodePool_LocalDiskEncryptionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_LocalDiskEncryption) *krm.NodePool_LocalDiskEncryptionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_LocalDiskEncryptionObservedState{}
	// MISSING: KMSKey
	out.KMSKeyActiveVersion = direct.LazyPtr(in.GetKmsKeyActiveVersion())
	out.KMSKeyState = direct.Enum_FromProto(mapCtx, in.GetKmsKeyState())
	out.KMSStatus = Status_FromProto(mapCtx, in.GetKmsStatus())
	out.ResourceState = direct.Enum_FromProto(mapCtx, in.GetResourceState())
	return out
}
func NodePool_LocalDiskEncryptionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_LocalDiskEncryptionObservedState) *pb.NodePool_LocalDiskEncryption {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_LocalDiskEncryption{}
	// MISSING: KMSKey
	out.KmsKeyActiveVersion = direct.ValueOf(in.KMSKeyActiveVersion)
	out.KmsKeyState = direct.Enum_ToProto[pb.KmsKeyState](mapCtx, in.KMSKeyState)
	out.KmsStatus = Status_ToProto(mapCtx, in.KMSStatus)
	out.ResourceState = direct.Enum_ToProto[pb.ResourceState](mapCtx, in.ResourceState)
	return out
}
func NodePool_NodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_NodeConfig) *krm.NodePool_NodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_NodeConfig{}
	out.Labels = in.Labels
	out.NodeStorageSchema = direct.LazyPtr(in.GetNodeStorageSchema())
	return out
}
func NodePool_NodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_NodeConfig) *pb.NodePool_NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_NodeConfig{}
	out.Labels = in.Labels
	out.NodeStorageSchema = direct.ValueOf(in.NodeStorageSchema)
	return out
}
