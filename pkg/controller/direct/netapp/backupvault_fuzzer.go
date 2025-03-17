// Copyright 2024 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.cloud.netapp.v1.BackupVault
// api.group: netapp.cnrm.cloud.google.com

package netapp

import (
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(netAppBackupVaultFuzzer())
}

func netAppBackupVaultFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.BackupVault{},
		netAppBackupVaultSpecFromProto, netAppBackupVaultSpecToProto,
		netAppBackupVaultObservedStateFromProto, netAppBackupVaultObservedStateToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".create_time")

	f.UnimplementedFields.Insert(".name")

	return f
}

func netAppBackupVaultSpecFromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.NetAppBackupVaultSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetAppBackupVaultSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	return out
}

func netAppBackupVaultSpecToProto(mapCtx *direct.MapContext, in *krm.NetAppBackupVaultSpec) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	return out
}

func netAppBackupVaultObservedStateFromProto(mapCtx *direct.MapContext, in *pb.BackupVault) *krm.NetAppBackupVaultObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetAppBackupVaultObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}

func netAppBackupVaultObservedStateToProto(mapCtx *direct.MapContext, in *krm.NetAppBackupVaultObservedState) *pb.BackupVault {
	if in == nil {
		return nil
	}
	out := &pb.BackupVault{}
	out.State = direct.Enum_ToProto[pb.BackupVault_State](mapCtx, in.State)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
