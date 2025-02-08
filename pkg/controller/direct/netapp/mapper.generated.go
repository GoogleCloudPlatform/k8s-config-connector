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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BackupVault_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupVault {
	if in == nil {
		return nil
	}
	out := &krm.BackupVault{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: State
	// MISSING: CreateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	return out
}
func BackupVault_ToProto(mapCtx *direct.MapContext, in *krm.BackupVault) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: State
	// MISSING: CreateTime
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	return out
}
func BackupVaultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.BackupVaultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupVaultObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Description
	// MISSING: Labels
	return out
}
func BackupVaultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupVaultObservedState) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.BackupVault_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Description
	// MISSING: Labels
	return out
}
func NetappBackupVaultObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.NetappBackupVaultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetappBackupVaultObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	return out
}
func NetappBackupVaultObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetappBackupVaultObservedState) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	return out
}
func NetappBackupVaultSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.NetappBackupVaultSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetappBackupVaultSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	return out
}
func NetappBackupVaultSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetappBackupVaultSpec) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	// MISSING: Name
	// MISSING: State
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	return out
}
