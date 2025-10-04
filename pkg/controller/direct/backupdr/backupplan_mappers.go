// Copyright 2025 Google LLC
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
// proto.message: google.cloud.backupdr.v1.ManagementServer
// api.group: backupdr.cnrm.cloud.google.com

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	krmbackupdrv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupDRBackupPlanSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlan) *krm.BackupDRBackupPlanSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupDRBackupPlanSpec{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.BackupRules = direct.Slice_FromProto(mapCtx, in.BackupRules, BackupRule_v1beta1_FromProto)
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	if in.GetBackupVault() != "" {
		out.BackupVaultRef = &krmbackupdrv1alpha1.BackupVaultRef{
			External: in.GetBackupVault(),
		}
	}
	return out
}
func BackupDRBackupPlanSpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRBackupPlanSpec) *pb.BackupPlan {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlan{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.BackupRules = direct.Slice_ToProto(mapCtx, in.BackupRules, BackupRule_v1beta1_ToProto)
	out.ResourceType = direct.ValueOf(in.ResourceType)
	if in.BackupVaultRef != nil {
		out.BackupVault = in.BackupVaultRef.External
	}
	return out
}
