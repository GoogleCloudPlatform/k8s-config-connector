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

package backupdr

import (
	pb "cloud.google.com/go/backupdr/apiv1/backupdrpb"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/backupdr/v1alpha1"
	compute "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BackupDRBackupPlanAssociationSpec_Resource_FromProto(mapCtx *direct.MapContext, resource string, resourceType string) *krm.Resource {
	if resource == "" {
		return nil
	}
	out := &krm.Resource{}
	if resourceType == krm.ResourceType_ComputeInstance {
		out.ComputeInstanceRef = &compute.InstanceRef{External: resource}
	}
	return out
}
func BackupDRBackupPlanAssociationSpec_Resource_ToProto(mapCtx *direct.MapContext, in *krm.Resource) string {
	if in == nil {
		return ""
	}
	if in.ComputeInstanceRef != nil {
		return in.ComputeInstanceRef.External
	}
	return ""
}
func BackupDRBackupPlanAssociationSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupPlanAssociation) *krm.BackupDRBackupPlanAssociationSpec {
	if in == nil {
		return nil
	}
	out := &krm.BackupDRBackupPlanAssociationSpec{}
	// MISSING: Name
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.Resource = BackupDRBackupPlanAssociationSpec_Resource_FromProto(mapCtx, in.GetResource(), in.GetResourceType())
	if in.GetBackupPlan() != "" {
		out.BackupPlanRef = &krm.BackupPlanRef{External: in.GetBackupPlan()}
	}
	return out
}
func BackupDRBackupPlanAssociationSpec_ToProto(mapCtx *direct.MapContext, in *krm.BackupDRBackupPlanAssociationSpec) *pb.BackupPlanAssociation {
	if in == nil {
		return nil
	}
	out := &pb.BackupPlanAssociation{}
	// MISSING: Name
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.Resource = BackupDRBackupPlanAssociationSpec_Resource_ToProto(mapCtx, in.Resource)
	if in.BackupPlanRef != nil {
		out.BackupPlan = in.BackupPlanRef.External
	}
	return out
}
func Status_ToProto(mapCtx *direct.MapContext, in *krm.Status) *statuspb.Status {
	if in == nil {
		return nil
	}
	out := &statuspb.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	out.Details = direct.Slice_ToProto(mapCtx, in.Details, Detail_ToProto)
	return out
}
func Status_FromProto(mapCtx *direct.MapContext, in *statuspb.Status) *krm.Status {
	if in == nil {
		return nil
	}
	out := &krm.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.Details = direct.Slice_FromProto(mapCtx, in.GetDetails(), Detail_FromProto)
	return out
}
func Detail_ToProto(mapCtx *direct.MapContext, in *krm.Any) *anypb.Any {
	if in == nil {
		return nil
	}
	out := &anypb.Any{}
	out.TypeUrl = direct.ValueOf(in.TypeURL)
	out.Value = in.Value
	return out
}
func Detail_FromProto(mapCtx *direct.MapContext, in *anypb.Any) *krm.Any {
	if in == nil {
		return nil
	}
	out := &krm.Any{}
	out.TypeURL = direct.LazyPtr(in.GetTypeUrl())
	out.Value = in.GetValue()
	return out
}
