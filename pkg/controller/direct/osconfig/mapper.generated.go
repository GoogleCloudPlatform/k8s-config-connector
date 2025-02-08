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

package osconfig

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func FixedOrPercent_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &krm.FixedOrPercent{}
	out.Fixed = direct.LazyPtr(in.GetFixed())
	out.Percent = direct.LazyPtr(in.GetPercent())
	return out
}
func FixedOrPercent_ToProto(mapCtx *direct.MapContext, in *krm.FixedOrPercent) *pb.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &pb.FixedOrPercent{}
	if oneof := FixedOrPercent_Fixed_ToProto(mapCtx, in.Fixed); oneof != nil {
		out.Mode = oneof
	}
	if oneof := FixedOrPercent_Percent_ToProto(mapCtx, in.Percent); oneof != nil {
		out.Mode = oneof
	}
	return out
}
func OSPolicy_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy) *krm.OSPolicy {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.ResourceGroups = direct.Slice_FromProto(mapCtx, in.ResourceGroups, OSPolicy_ResourceGroup_FromProto)
	out.AllowNoResourceGroupMatch = direct.LazyPtr(in.GetAllowNoResourceGroupMatch())
	return out
}
func OSPolicy_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy) *pb.OSPolicy {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy{}
	out.Id = direct.ValueOf(in.ID)
	out.Description = direct.ValueOf(in.Description)
	out.Mode = direct.Enum_ToProto[pb.OSPolicy_Mode](mapCtx, in.Mode)
	out.ResourceGroups = direct.Slice_ToProto(mapCtx, in.ResourceGroups, OSPolicy_ResourceGroup_ToProto)
	out.AllowNoResourceGroupMatch = direct.ValueOf(in.AllowNoResourceGroupMatch)
	return out
}
func OSPolicyAssignment_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment) *krm.OSPolicyAssignment {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignment{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.OsPolicies = direct.Slice_FromProto(mapCtx, in.OsPolicies, OSPolicy_FromProto)
	out.InstanceFilter = OSPolicyAssignment_InstanceFilter_FromProto(mapCtx, in.GetInstanceFilter())
	out.Rollout = OSPolicyAssignment_Rollout_FromProto(mapCtx, in.GetRollout())
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	out.Etag = direct.LazyPtr(in.GetEtag())
	// MISSING: RolloutState
	// MISSING: Baseline
	// MISSING: Deleted
	// MISSING: Reconciling
	// MISSING: Uid
	return out
}
func OSPolicyAssignment_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignment) *pb.OSPolicyAssignment {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.OsPolicies = direct.Slice_ToProto(mapCtx, in.OsPolicies, OSPolicy_ToProto)
	out.InstanceFilter = OSPolicyAssignment_InstanceFilter_ToProto(mapCtx, in.InstanceFilter)
	out.Rollout = OSPolicyAssignment_Rollout_ToProto(mapCtx, in.Rollout)
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	out.Etag = direct.ValueOf(in.Etag)
	// MISSING: RolloutState
	// MISSING: Baseline
	// MISSING: Deleted
	// MISSING: Reconciling
	// MISSING: Uid
	return out
}
func OSPolicyAssignmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment) *krm.OSPolicyAssignmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignmentObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OsPolicies
	// MISSING: InstanceFilter
	// MISSING: Rollout
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	// MISSING: Etag
	out.RolloutState = direct.Enum_FromProto(mapCtx, in.GetRolloutState())
	out.Baseline = direct.LazyPtr(in.GetBaseline())
	out.Deleted = direct.LazyPtr(in.GetDeleted())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func OSPolicyAssignmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignmentObservedState) *pb.OSPolicyAssignment {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OsPolicies
	// MISSING: InstanceFilter
	// MISSING: Rollout
	out.RevisionId = direct.ValueOf(in.RevisionID)
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	// MISSING: Etag
	out.RolloutState = direct.Enum_ToProto[pb.OSPolicyAssignment_RolloutState](mapCtx, in.RolloutState)
	out.Baseline = direct.ValueOf(in.Baseline)
	out.Deleted = direct.ValueOf(in.Deleted)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func OSPolicyAssignment_InstanceFilter_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment_InstanceFilter) *krm.OSPolicyAssignment_InstanceFilter {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignment_InstanceFilter{}
	out.All = direct.LazyPtr(in.GetAll())
	out.InclusionLabels = direct.Slice_FromProto(mapCtx, in.InclusionLabels, OSPolicyAssignment_LabelSet_FromProto)
	out.ExclusionLabels = direct.Slice_FromProto(mapCtx, in.ExclusionLabels, OSPolicyAssignment_LabelSet_FromProto)
	out.Inventories = direct.Slice_FromProto(mapCtx, in.Inventories, OSPolicyAssignment_InstanceFilter_Inventory_FromProto)
	return out
}
func OSPolicyAssignment_InstanceFilter_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignment_InstanceFilter) *pb.OSPolicyAssignment_InstanceFilter {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment_InstanceFilter{}
	out.All = direct.ValueOf(in.All)
	out.InclusionLabels = direct.Slice_ToProto(mapCtx, in.InclusionLabels, OSPolicyAssignment_LabelSet_ToProto)
	out.ExclusionLabels = direct.Slice_ToProto(mapCtx, in.ExclusionLabels, OSPolicyAssignment_LabelSet_ToProto)
	out.Inventories = direct.Slice_ToProto(mapCtx, in.Inventories, OSPolicyAssignment_InstanceFilter_Inventory_ToProto)
	return out
}
func OSPolicyAssignment_InstanceFilter_Inventory_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment_InstanceFilter_Inventory) *krm.OSPolicyAssignment_InstanceFilter_Inventory {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignment_InstanceFilter_Inventory{}
	out.OsShortName = direct.LazyPtr(in.GetOsShortName())
	out.OsVersion = direct.LazyPtr(in.GetOsVersion())
	return out
}
func OSPolicyAssignment_InstanceFilter_Inventory_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignment_InstanceFilter_Inventory) *pb.OSPolicyAssignment_InstanceFilter_Inventory {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment_InstanceFilter_Inventory{}
	out.OsShortName = direct.ValueOf(in.OsShortName)
	out.OsVersion = direct.ValueOf(in.OsVersion)
	return out
}
func OSPolicyAssignment_LabelSet_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment_LabelSet) *krm.OSPolicyAssignment_LabelSet {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignment_LabelSet{}
	out.Labels = in.Labels
	return out
}
func OSPolicyAssignment_LabelSet_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignment_LabelSet) *pb.OSPolicyAssignment_LabelSet {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment_LabelSet{}
	out.Labels = in.Labels
	return out
}
func OSPolicyAssignment_Rollout_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment_Rollout) *krm.OSPolicyAssignment_Rollout {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignment_Rollout{}
	out.DisruptionBudget = FixedOrPercent_FromProto(mapCtx, in.GetDisruptionBudget())
	out.MinWaitDuration = direct.StringDuration_FromProto(mapCtx, in.GetMinWaitDuration())
	return out
}
func OSPolicyAssignment_Rollout_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignment_Rollout) *pb.OSPolicyAssignment_Rollout {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment_Rollout{}
	out.DisruptionBudget = FixedOrPercent_ToProto(mapCtx, in.DisruptionBudget)
	out.MinWaitDuration = direct.StringDuration_ToProto(mapCtx, in.MinWaitDuration)
	return out
}
func OSPolicy_InventoryFilter_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_InventoryFilter) *krm.OSPolicy_InventoryFilter {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_InventoryFilter{}
	out.OsShortName = direct.LazyPtr(in.GetOsShortName())
	out.OsVersion = direct.LazyPtr(in.GetOsVersion())
	return out
}
func OSPolicy_InventoryFilter_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_InventoryFilter) *pb.OSPolicy_InventoryFilter {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_InventoryFilter{}
	out.OsShortName = direct.ValueOf(in.OsShortName)
	out.OsVersion = direct.ValueOf(in.OsVersion)
	return out
}
func OSPolicy_Resource_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource) *krm.OSPolicy_Resource {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Pkg = OSPolicy_Resource_PackageResource_FromProto(mapCtx, in.GetPkg())
	out.Repository = OSPolicy_Resource_RepositoryResource_FromProto(mapCtx, in.GetRepository())
	out.Exec = OSPolicy_Resource_ExecResource_FromProto(mapCtx, in.GetExec())
	out.File = OSPolicy_Resource_FileResource_FromProto(mapCtx, in.GetFile())
	return out
}
func OSPolicy_Resource_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource) *pb.OSPolicy_Resource {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource{}
	out.Id = direct.ValueOf(in.ID)
	if oneof := OSPolicy_Resource_PackageResource_ToProto(mapCtx, in.Pkg); oneof != nil {
		out.ResourceType = &pb.OSPolicy_Resource_Pkg{Pkg: oneof}
	}
	if oneof := OSPolicy_Resource_RepositoryResource_ToProto(mapCtx, in.Repository); oneof != nil {
		out.ResourceType = &pb.OSPolicy_Resource_Repository{Repository: oneof}
	}
	if oneof := OSPolicy_Resource_ExecResource_ToProto(mapCtx, in.Exec); oneof != nil {
		out.ResourceType = &pb.OSPolicy_Resource_Exec{Exec: oneof}
	}
	if oneof := OSPolicy_Resource_FileResource_ToProto(mapCtx, in.File); oneof != nil {
		out.ResourceType = &pb.OSPolicy_Resource_File{File: oneof}
	}
	return out
}
func OSPolicy_ResourceGroup_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_ResourceGroup) *krm.OSPolicy_ResourceGroup {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_ResourceGroup{}
	out.InventoryFilters = direct.Slice_FromProto(mapCtx, in.InventoryFilters, OSPolicy_InventoryFilter_FromProto)
	out.Resources = direct.Slice_FromProto(mapCtx, in.Resources, OSPolicy_Resource_FromProto)
	return out
}
func OSPolicy_ResourceGroup_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_ResourceGroup) *pb.OSPolicy_ResourceGroup {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_ResourceGroup{}
	out.InventoryFilters = direct.Slice_ToProto(mapCtx, in.InventoryFilters, OSPolicy_InventoryFilter_ToProto)
	out.Resources = direct.Slice_ToProto(mapCtx, in.Resources, OSPolicy_Resource_ToProto)
	return out
}
func OSPolicy_Resource_ExecResource_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_ExecResource) *krm.OSPolicy_Resource_ExecResource {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_ExecResource{}
	out.Validate = OSPolicy_Resource_ExecResource_Exec_FromProto(mapCtx, in.GetValidate())
	out.Enforce = OSPolicy_Resource_ExecResource_Exec_FromProto(mapCtx, in.GetEnforce())
	return out
}
func OSPolicy_Resource_ExecResource_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_ExecResource) *pb.OSPolicy_Resource_ExecResource {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_ExecResource{}
	out.Validate = OSPolicy_Resource_ExecResource_Exec_ToProto(mapCtx, in.Validate)
	out.Enforce = OSPolicy_Resource_ExecResource_Exec_ToProto(mapCtx, in.Enforce)
	return out
}
func OSPolicy_Resource_ExecResource_Exec_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_ExecResource_Exec) *krm.OSPolicy_Resource_ExecResource_Exec {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_ExecResource_Exec{}
	out.File = OSPolicy_Resource_File_FromProto(mapCtx, in.GetFile())
	out.Script = direct.LazyPtr(in.GetScript())
	out.Args = in.Args
	out.Interpreter = direct.Enum_FromProto(mapCtx, in.GetInterpreter())
	out.OutputFilePath = direct.LazyPtr(in.GetOutputFilePath())
	return out
}
func OSPolicy_Resource_ExecResource_Exec_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_ExecResource_Exec) *pb.OSPolicy_Resource_ExecResource_Exec {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_ExecResource_Exec{}
	if oneof := OSPolicy_Resource_File_ToProto(mapCtx, in.File); oneof != nil {
		out.Source = &pb.OSPolicy_Resource_ExecResource_Exec_File{File: oneof}
	}
	if oneof := OSPolicy_Resource_ExecResource_Exec_Script_ToProto(mapCtx, in.Script); oneof != nil {
		out.Source = oneof
	}
	out.Args = in.Args
	out.Interpreter = direct.Enum_ToProto[pb.OSPolicy_Resource_ExecResource_Exec_Interpreter](mapCtx, in.Interpreter)
	out.OutputFilePath = direct.ValueOf(in.OutputFilePath)
	return out
}
func OSPolicy_Resource_File_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_File) *krm.OSPolicy_Resource_File {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_File{}
	out.Remote = OSPolicy_Resource_File_Remote_FromProto(mapCtx, in.GetRemote())
	out.Gcs = OSPolicy_Resource_File_Gcs_FromProto(mapCtx, in.GetGcs())
	out.LocalPath = direct.LazyPtr(in.GetLocalPath())
	out.AllowInsecure = direct.LazyPtr(in.GetAllowInsecure())
	return out
}
func OSPolicy_Resource_File_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_File) *pb.OSPolicy_Resource_File {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_File{}
	if oneof := OSPolicy_Resource_File_Remote_ToProto(mapCtx, in.Remote); oneof != nil {
		out.Type = &pb.OSPolicy_Resource_File_Remote_{Remote: oneof}
	}
	if oneof := OSPolicy_Resource_File_Gcs_ToProto(mapCtx, in.Gcs); oneof != nil {
		out.Type = &pb.OSPolicy_Resource_File_Gcs_{Gcs: oneof}
	}
	if oneof := OSPolicy_Resource_File_LocalPath_ToProto(mapCtx, in.LocalPath); oneof != nil {
		out.Type = oneof
	}
	out.AllowInsecure = direct.ValueOf(in.AllowInsecure)
	return out
}
func OSPolicy_Resource_FileResource_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_FileResource) *krm.OSPolicy_Resource_FileResource {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_FileResource{}
	out.File = OSPolicy_Resource_File_FromProto(mapCtx, in.GetFile())
	out.Content = direct.LazyPtr(in.GetContent())
	out.Path = direct.LazyPtr(in.GetPath())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Permissions = direct.LazyPtr(in.GetPermissions())
	return out
}
func OSPolicy_Resource_FileResource_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_FileResource) *pb.OSPolicy_Resource_FileResource {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_FileResource{}
	if oneof := OSPolicy_Resource_File_ToProto(mapCtx, in.File); oneof != nil {
		out.Source = &pb.OSPolicy_Resource_FileResource_File{File: oneof}
	}
	if oneof := OSPolicy_Resource_FileResource_Content_ToProto(mapCtx, in.Content); oneof != nil {
		out.Source = oneof
	}
	out.Path = direct.ValueOf(in.Path)
	out.State = direct.Enum_ToProto[pb.OSPolicy_Resource_FileResource_DesiredState](mapCtx, in.State)
	out.Permissions = direct.ValueOf(in.Permissions)
	return out
}
func OSPolicy_Resource_File_Gcs_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_File_Gcs) *krm.OSPolicy_Resource_File_Gcs {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_File_Gcs{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.Object = direct.LazyPtr(in.GetObject())
	out.Generation = direct.LazyPtr(in.GetGeneration())
	return out
}
func OSPolicy_Resource_File_Gcs_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_File_Gcs) *pb.OSPolicy_Resource_File_Gcs {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_File_Gcs{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.Object = direct.ValueOf(in.Object)
	out.Generation = direct.ValueOf(in.Generation)
	return out
}
func OSPolicy_Resource_File_Remote_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_File_Remote) *krm.OSPolicy_Resource_File_Remote {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_File_Remote{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Sha256Checksum = direct.LazyPtr(in.GetSha256Checksum())
	return out
}
func OSPolicy_Resource_File_Remote_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_File_Remote) *pb.OSPolicy_Resource_File_Remote {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_File_Remote{}
	out.Uri = direct.ValueOf(in.URI)
	out.Sha256Checksum = direct.ValueOf(in.Sha256Checksum)
	return out
}
func OSPolicy_Resource_PackageResource_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource) *krm.OSPolicy_Resource_PackageResource {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource{}
	out.DesiredState = direct.Enum_FromProto(mapCtx, in.GetDesiredState())
	out.Apt = OSPolicy_Resource_PackageResource_APT_FromProto(mapCtx, in.GetApt())
	out.Deb = OSPolicy_Resource_PackageResource_Deb_FromProto(mapCtx, in.GetDeb())
	out.Yum = OSPolicy_Resource_PackageResource_YUM_FromProto(mapCtx, in.GetYum())
	out.Zypper = OSPolicy_Resource_PackageResource_Zypper_FromProto(mapCtx, in.GetZypper())
	out.Rpm = OSPolicy_Resource_PackageResource_RPM_FromProto(mapCtx, in.GetRpm())
	out.Googet = OSPolicy_Resource_PackageResource_GooGet_FromProto(mapCtx, in.GetGooget())
	out.Msi = OSPolicy_Resource_PackageResource_MSI_FromProto(mapCtx, in.GetMsi())
	return out
}
func OSPolicy_Resource_PackageResource_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource) *pb.OSPolicy_Resource_PackageResource {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource{}
	out.DesiredState = direct.Enum_ToProto[pb.OSPolicy_Resource_PackageResource_DesiredState](mapCtx, in.DesiredState)
	if oneof := OSPolicy_Resource_PackageResource_APT_ToProto(mapCtx, in.Apt); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Apt{Apt: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_Deb_ToProto(mapCtx, in.Deb); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Deb_{Deb: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_YUM_ToProto(mapCtx, in.Yum); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Yum{Yum: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_Zypper_ToProto(mapCtx, in.Zypper); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Zypper_{Zypper: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_RPM_ToProto(mapCtx, in.Rpm); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Rpm{Rpm: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_GooGet_ToProto(mapCtx, in.Googet); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Googet{Googet: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_MSI_ToProto(mapCtx, in.Msi); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Msi{Msi: oneof}
	}
	return out
}
func OSPolicy_Resource_PackageResource_APT_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_APT) *krm.OSPolicy_Resource_PackageResource_APT {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_APT{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func OSPolicy_Resource_PackageResource_APT_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_APT) *pb.OSPolicy_Resource_PackageResource_APT {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_APT{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func OSPolicy_Resource_PackageResource_Deb_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_Deb) *krm.OSPolicy_Resource_PackageResource_Deb {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_Deb{}
	out.Source = OSPolicy_Resource_File_FromProto(mapCtx, in.GetSource())
	out.PullDeps = direct.LazyPtr(in.GetPullDeps())
	return out
}
func OSPolicy_Resource_PackageResource_Deb_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_Deb) *pb.OSPolicy_Resource_PackageResource_Deb {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_Deb{}
	out.Source = OSPolicy_Resource_File_ToProto(mapCtx, in.Source)
	out.PullDeps = direct.ValueOf(in.PullDeps)
	return out
}
func OSPolicy_Resource_PackageResource_GooGet_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_GooGet) *krm.OSPolicy_Resource_PackageResource_GooGet {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_GooGet{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func OSPolicy_Resource_PackageResource_GooGet_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_GooGet) *pb.OSPolicy_Resource_PackageResource_GooGet {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_GooGet{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func OSPolicy_Resource_PackageResource_MSI_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_MSI) *krm.OSPolicy_Resource_PackageResource_MSI {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_MSI{}
	out.Source = OSPolicy_Resource_File_FromProto(mapCtx, in.GetSource())
	out.Properties = in.Properties
	return out
}
func OSPolicy_Resource_PackageResource_MSI_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_MSI) *pb.OSPolicy_Resource_PackageResource_MSI {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_MSI{}
	out.Source = OSPolicy_Resource_File_ToProto(mapCtx, in.Source)
	out.Properties = in.Properties
	return out
}
func OSPolicy_Resource_PackageResource_RPM_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_RPM) *krm.OSPolicy_Resource_PackageResource_RPM {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_RPM{}
	out.Source = OSPolicy_Resource_File_FromProto(mapCtx, in.GetSource())
	out.PullDeps = direct.LazyPtr(in.GetPullDeps())
	return out
}
func OSPolicy_Resource_PackageResource_RPM_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_RPM) *pb.OSPolicy_Resource_PackageResource_RPM {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_RPM{}
	out.Source = OSPolicy_Resource_File_ToProto(mapCtx, in.Source)
	out.PullDeps = direct.ValueOf(in.PullDeps)
	return out
}
func OSPolicy_Resource_PackageResource_YUM_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_YUM) *krm.OSPolicy_Resource_PackageResource_YUM {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_YUM{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func OSPolicy_Resource_PackageResource_YUM_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_YUM) *pb.OSPolicy_Resource_PackageResource_YUM {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_YUM{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func OSPolicy_Resource_PackageResource_Zypper_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_Zypper) *krm.OSPolicy_Resource_PackageResource_Zypper {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_Zypper{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func OSPolicy_Resource_PackageResource_Zypper_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_Zypper) *pb.OSPolicy_Resource_PackageResource_Zypper {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_Zypper{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func OSPolicy_Resource_RepositoryResource_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_RepositoryResource) *krm.OSPolicy_Resource_RepositoryResource {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_RepositoryResource{}
	out.Apt = OSPolicy_Resource_RepositoryResource_AptRepository_FromProto(mapCtx, in.GetApt())
	out.Yum = OSPolicy_Resource_RepositoryResource_YumRepository_FromProto(mapCtx, in.GetYum())
	out.Zypper = OSPolicy_Resource_RepositoryResource_ZypperRepository_FromProto(mapCtx, in.GetZypper())
	out.Goo = OSPolicy_Resource_RepositoryResource_GooRepository_FromProto(mapCtx, in.GetGoo())
	return out
}
func OSPolicy_Resource_RepositoryResource_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_RepositoryResource) *pb.OSPolicy_Resource_RepositoryResource {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_RepositoryResource{}
	if oneof := OSPolicy_Resource_RepositoryResource_AptRepository_ToProto(mapCtx, in.Apt); oneof != nil {
		out.Repository = &pb.OSPolicy_Resource_RepositoryResource_Apt{Apt: oneof}
	}
	if oneof := OSPolicy_Resource_RepositoryResource_YumRepository_ToProto(mapCtx, in.Yum); oneof != nil {
		out.Repository = &pb.OSPolicy_Resource_RepositoryResource_Yum{Yum: oneof}
	}
	if oneof := OSPolicy_Resource_RepositoryResource_ZypperRepository_ToProto(mapCtx, in.Zypper); oneof != nil {
		out.Repository = &pb.OSPolicy_Resource_RepositoryResource_Zypper{Zypper: oneof}
	}
	if oneof := OSPolicy_Resource_RepositoryResource_GooRepository_ToProto(mapCtx, in.Goo); oneof != nil {
		out.Repository = &pb.OSPolicy_Resource_RepositoryResource_Goo{Goo: oneof}
	}
	return out
}
func OSPolicy_Resource_RepositoryResource_AptRepository_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_RepositoryResource_AptRepository) *krm.OSPolicy_Resource_RepositoryResource_AptRepository {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_RepositoryResource_AptRepository{}
	out.ArchiveType = direct.Enum_FromProto(mapCtx, in.GetArchiveType())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Distribution = direct.LazyPtr(in.GetDistribution())
	out.Components = in.Components
	out.GpgKey = direct.LazyPtr(in.GetGpgKey())
	return out
}
func OSPolicy_Resource_RepositoryResource_AptRepository_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_RepositoryResource_AptRepository) *pb.OSPolicy_Resource_RepositoryResource_AptRepository {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_RepositoryResource_AptRepository{}
	out.ArchiveType = direct.Enum_ToProto[pb.OSPolicy_Resource_RepositoryResource_AptRepository_ArchiveType](mapCtx, in.ArchiveType)
	out.Uri = direct.ValueOf(in.URI)
	out.Distribution = direct.ValueOf(in.Distribution)
	out.Components = in.Components
	out.GpgKey = direct.ValueOf(in.GpgKey)
	return out
}
func OSPolicy_Resource_RepositoryResource_GooRepository_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_RepositoryResource_GooRepository) *krm.OSPolicy_Resource_RepositoryResource_GooRepository {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_RepositoryResource_GooRepository{}
	out.Name = direct.LazyPtr(in.GetName())
	out.URL = direct.LazyPtr(in.GetUrl())
	return out
}
func OSPolicy_Resource_RepositoryResource_GooRepository_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_RepositoryResource_GooRepository) *pb.OSPolicy_Resource_RepositoryResource_GooRepository {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_RepositoryResource_GooRepository{}
	out.Name = direct.ValueOf(in.Name)
	out.Url = direct.ValueOf(in.URL)
	return out
}
func OSPolicy_Resource_RepositoryResource_YumRepository_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_RepositoryResource_YumRepository) *krm.OSPolicy_Resource_RepositoryResource_YumRepository {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_RepositoryResource_YumRepository{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BaseURL = direct.LazyPtr(in.GetBaseUrl())
	out.GpgKeys = in.GpgKeys
	return out
}
func OSPolicy_Resource_RepositoryResource_YumRepository_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_RepositoryResource_YumRepository) *pb.OSPolicy_Resource_RepositoryResource_YumRepository {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_RepositoryResource_YumRepository{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BaseUrl = direct.ValueOf(in.BaseURL)
	out.GpgKeys = in.GpgKeys
	return out
}
func OSPolicy_Resource_RepositoryResource_ZypperRepository_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_RepositoryResource_ZypperRepository) *krm.OSPolicy_Resource_RepositoryResource_ZypperRepository {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_RepositoryResource_ZypperRepository{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BaseURL = direct.LazyPtr(in.GetBaseUrl())
	out.GpgKeys = in.GpgKeys
	return out
}
func OSPolicy_Resource_RepositoryResource_ZypperRepository_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_RepositoryResource_ZypperRepository) *pb.OSPolicy_Resource_RepositoryResource_ZypperRepository {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_RepositoryResource_ZypperRepository{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BaseUrl = direct.ValueOf(in.BaseURL)
	out.GpgKeys = in.GpgKeys
	return out
}
func OsconfigOSPolicyAssignmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment) *krm.OsconfigOSPolicyAssignmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigOSPolicyAssignmentObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OsPolicies
	// MISSING: InstanceFilter
	// MISSING: Rollout
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	// MISSING: Etag
	// MISSING: RolloutState
	// MISSING: Baseline
	// MISSING: Deleted
	// MISSING: Reconciling
	// MISSING: Uid
	return out
}
func OsconfigOSPolicyAssignmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigOSPolicyAssignmentObservedState) *pb.OSPolicyAssignment {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OsPolicies
	// MISSING: InstanceFilter
	// MISSING: Rollout
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	// MISSING: Etag
	// MISSING: RolloutState
	// MISSING: Baseline
	// MISSING: Deleted
	// MISSING: Reconciling
	// MISSING: Uid
	return out
}
func OsconfigOSPolicyAssignmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment) *krm.OsconfigOSPolicyAssignmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigOSPolicyAssignmentSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OsPolicies
	// MISSING: InstanceFilter
	// MISSING: Rollout
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	// MISSING: Etag
	// MISSING: RolloutState
	// MISSING: Baseline
	// MISSING: Deleted
	// MISSING: Reconciling
	// MISSING: Uid
	return out
}
func OsconfigOSPolicyAssignmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigOSPolicyAssignmentSpec) *pb.OSPolicyAssignment {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: OsPolicies
	// MISSING: InstanceFilter
	// MISSING: Rollout
	// MISSING: RevisionID
	// MISSING: RevisionCreateTime
	// MISSING: Etag
	// MISSING: RolloutState
	// MISSING: Baseline
	// MISSING: Deleted
	// MISSING: Reconciling
	// MISSING: Uid
	return out
}
