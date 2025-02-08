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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/osconfig/apiv1beta/osconfigpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1alpha1"
)
func AptRepository_FromProto(mapCtx *direct.MapContext, in *pb.AptRepository) *krm.AptRepository {
	if in == nil {
		return nil
	}
	out := &krm.AptRepository{}
	out.ArchiveType = direct.Enum_FromProto(mapCtx, in.GetArchiveType())
	out.URI = direct.LazyPtr(in.GetUri())
	out.Distribution = direct.LazyPtr(in.GetDistribution())
	out.Components = in.Components
	out.GpgKey = direct.LazyPtr(in.GetGpgKey())
	return out
}
func AptRepository_ToProto(mapCtx *direct.MapContext, in *krm.AptRepository) *pb.AptRepository {
	if in == nil {
		return nil
	}
	out := &pb.AptRepository{}
	out.ArchiveType = direct.Enum_ToProto[pb.AptRepository_ArchiveType](mapCtx, in.ArchiveType)
	out.Uri = direct.ValueOf(in.URI)
	out.Distribution = direct.ValueOf(in.Distribution)
	out.Components = in.Components
	out.GpgKey = direct.ValueOf(in.GpgKey)
	return out
}
func Assignment_FromProto(mapCtx *direct.MapContext, in *pb.Assignment) *krm.Assignment {
	if in == nil {
		return nil
	}
	out := &krm.Assignment{}
	out.GroupLabels = direct.Slice_FromProto(mapCtx, in.GroupLabels, Assignment_GroupLabel_FromProto)
	out.Zones = in.Zones
	out.Instances = in.Instances
	out.InstanceNamePrefixes = in.InstanceNamePrefixes
	out.OsTypes = direct.Slice_FromProto(mapCtx, in.OsTypes, Assignment_OsType_FromProto)
	return out
}
func Assignment_ToProto(mapCtx *direct.MapContext, in *krm.Assignment) *pb.Assignment {
	if in == nil {
		return nil
	}
	out := &pb.Assignment{}
	out.GroupLabels = direct.Slice_ToProto(mapCtx, in.GroupLabels, Assignment_GroupLabel_ToProto)
	out.Zones = in.Zones
	out.Instances = in.Instances
	out.InstanceNamePrefixes = in.InstanceNamePrefixes
	out.OsTypes = direct.Slice_ToProto(mapCtx, in.OsTypes, Assignment_OsType_ToProto)
	return out
}
func Assignment_GroupLabel_FromProto(mapCtx *direct.MapContext, in *pb.Assignment_GroupLabel) *krm.Assignment_GroupLabel {
	if in == nil {
		return nil
	}
	out := &krm.Assignment_GroupLabel{}
	out.Labels = in.Labels
	return out
}
func Assignment_GroupLabel_ToProto(mapCtx *direct.MapContext, in *krm.Assignment_GroupLabel) *pb.Assignment_GroupLabel {
	if in == nil {
		return nil
	}
	out := &pb.Assignment_GroupLabel{}
	out.Labels = in.Labels
	return out
}
func Assignment_OsType_FromProto(mapCtx *direct.MapContext, in *pb.Assignment_OsType) *krm.Assignment_OsType {
	if in == nil {
		return nil
	}
	out := &krm.Assignment_OsType{}
	out.OsShortName = direct.LazyPtr(in.GetOsShortName())
	out.OsVersion = direct.LazyPtr(in.GetOsVersion())
	out.OsArchitecture = direct.LazyPtr(in.GetOsArchitecture())
	return out
}
func Assignment_OsType_ToProto(mapCtx *direct.MapContext, in *krm.Assignment_OsType) *pb.Assignment_OsType {
	if in == nil {
		return nil
	}
	out := &pb.Assignment_OsType{}
	out.OsShortName = direct.ValueOf(in.OsShortName)
	out.OsVersion = direct.ValueOf(in.OsVersion)
	out.OsArchitecture = direct.ValueOf(in.OsArchitecture)
	return out
}
func GooRepository_FromProto(mapCtx *direct.MapContext, in *pb.GooRepository) *krm.GooRepository {
	if in == nil {
		return nil
	}
	out := &krm.GooRepository{}
	out.Name = direct.LazyPtr(in.GetName())
	out.URL = direct.LazyPtr(in.GetUrl())
	return out
}
func GooRepository_ToProto(mapCtx *direct.MapContext, in *krm.GooRepository) *pb.GooRepository {
	if in == nil {
		return nil
	}
	out := &pb.GooRepository{}
	out.Name = direct.ValueOf(in.Name)
	out.Url = direct.ValueOf(in.URL)
	return out
}
func GuestPolicy_FromProto(mapCtx *direct.MapContext, in *pb.GuestPolicy) *krm.GuestPolicy {
	if in == nil {
		return nil
	}
	out := &krm.GuestPolicy{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Assignment = Assignment_FromProto(mapCtx, in.GetAssignment())
	out.Packages = direct.Slice_FromProto(mapCtx, in.Packages, Package_FromProto)
	out.PackageRepositories = direct.Slice_FromProto(mapCtx, in.PackageRepositories, PackageRepository_FromProto)
	out.Recipes = direct.Slice_FromProto(mapCtx, in.Recipes, SoftwareRecipe_FromProto)
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func GuestPolicy_ToProto(mapCtx *direct.MapContext, in *krm.GuestPolicy) *pb.GuestPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GuestPolicy{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Assignment = Assignment_ToProto(mapCtx, in.Assignment)
	out.Packages = direct.Slice_ToProto(mapCtx, in.Packages, Package_ToProto)
	out.PackageRepositories = direct.Slice_ToProto(mapCtx, in.PackageRepositories, PackageRepository_ToProto)
	out.Recipes = direct.Slice_ToProto(mapCtx, in.Recipes, SoftwareRecipe_ToProto)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
func GuestPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GuestPolicy) *krm.GuestPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GuestPolicyObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Assignment
	// MISSING: Packages
	// MISSING: PackageRepositories
	// MISSING: Recipes
	// MISSING: Etag
	return out
}
func GuestPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GuestPolicyObservedState) *pb.GuestPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GuestPolicy{}
	// MISSING: Name
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Assignment
	// MISSING: Packages
	// MISSING: PackageRepositories
	// MISSING: Recipes
	// MISSING: Etag
	return out
}
func OsconfigGuestPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GuestPolicy) *krm.OsconfigGuestPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigGuestPolicyObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Assignment
	// MISSING: Packages
	// MISSING: PackageRepositories
	// MISSING: Recipes
	// MISSING: Etag
	return out
}
func OsconfigGuestPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigGuestPolicyObservedState) *pb.GuestPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GuestPolicy{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Assignment
	// MISSING: Packages
	// MISSING: PackageRepositories
	// MISSING: Recipes
	// MISSING: Etag
	return out
}
func OsconfigGuestPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.GuestPolicy) *krm.OsconfigGuestPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigGuestPolicySpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Assignment
	// MISSING: Packages
	// MISSING: PackageRepositories
	// MISSING: Recipes
	// MISSING: Etag
	return out
}
func OsconfigGuestPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigGuestPolicySpec) *pb.GuestPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GuestPolicy{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Assignment
	// MISSING: Packages
	// MISSING: PackageRepositories
	// MISSING: Recipes
	// MISSING: Etag
	return out
}
func Package_FromProto(mapCtx *direct.MapContext, in *pb.Package) *krm.Package {
	if in == nil {
		return nil
	}
	out := &krm.Package{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DesiredState = direct.Enum_FromProto(mapCtx, in.GetDesiredState())
	out.Manager = direct.Enum_FromProto(mapCtx, in.GetManager())
	return out
}
func Package_ToProto(mapCtx *direct.MapContext, in *krm.Package) *pb.Package {
	if in == nil {
		return nil
	}
	out := &pb.Package{}
	out.Name = direct.ValueOf(in.Name)
	out.DesiredState = direct.Enum_ToProto[pb.DesiredState](mapCtx, in.DesiredState)
	out.Manager = direct.Enum_ToProto[pb.Package_Manager](mapCtx, in.Manager)
	return out
}
func PackageRepository_FromProto(mapCtx *direct.MapContext, in *pb.PackageRepository) *krm.PackageRepository {
	if in == nil {
		return nil
	}
	out := &krm.PackageRepository{}
	out.Apt = AptRepository_FromProto(mapCtx, in.GetApt())
	out.Yum = YumRepository_FromProto(mapCtx, in.GetYum())
	out.Zypper = ZypperRepository_FromProto(mapCtx, in.GetZypper())
	out.Goo = GooRepository_FromProto(mapCtx, in.GetGoo())
	return out
}
func PackageRepository_ToProto(mapCtx *direct.MapContext, in *krm.PackageRepository) *pb.PackageRepository {
	if in == nil {
		return nil
	}
	out := &pb.PackageRepository{}
	if oneof := AptRepository_ToProto(mapCtx, in.Apt); oneof != nil {
		out.Repository = &pb.PackageRepository_Apt{Apt: oneof}
	}
	if oneof := YumRepository_ToProto(mapCtx, in.Yum); oneof != nil {
		out.Repository = &pb.PackageRepository_Yum{Yum: oneof}
	}
	if oneof := ZypperRepository_ToProto(mapCtx, in.Zypper); oneof != nil {
		out.Repository = &pb.PackageRepository_Zypper{Zypper: oneof}
	}
	if oneof := GooRepository_ToProto(mapCtx, in.Goo); oneof != nil {
		out.Repository = &pb.PackageRepository_Goo{Goo: oneof}
	}
	return out
}
func SoftwareRecipe_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe) *krm.SoftwareRecipe {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Artifacts = direct.Slice_FromProto(mapCtx, in.Artifacts, SoftwareRecipe_Artifact_FromProto)
	out.InstallSteps = direct.Slice_FromProto(mapCtx, in.InstallSteps, SoftwareRecipe_Step_FromProto)
	out.UpdateSteps = direct.Slice_FromProto(mapCtx, in.UpdateSteps, SoftwareRecipe_Step_FromProto)
	out.DesiredState = direct.Enum_FromProto(mapCtx, in.GetDesiredState())
	return out
}
func SoftwareRecipe_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe) *pb.SoftwareRecipe {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe{}
	out.Name = direct.ValueOf(in.Name)
	out.Version = direct.ValueOf(in.Version)
	out.Artifacts = direct.Slice_ToProto(mapCtx, in.Artifacts, SoftwareRecipe_Artifact_ToProto)
	out.InstallSteps = direct.Slice_ToProto(mapCtx, in.InstallSteps, SoftwareRecipe_Step_ToProto)
	out.UpdateSteps = direct.Slice_ToProto(mapCtx, in.UpdateSteps, SoftwareRecipe_Step_ToProto)
	out.DesiredState = direct.Enum_ToProto[pb.DesiredState](mapCtx, in.DesiredState)
	return out
}
func SoftwareRecipe_Artifact_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Artifact) *krm.SoftwareRecipe_Artifact {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Artifact{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Remote = SoftwareRecipe_Artifact_Remote_FromProto(mapCtx, in.GetRemote())
	out.Gcs = SoftwareRecipe_Artifact_Gcs_FromProto(mapCtx, in.GetGcs())
	out.AllowInsecure = direct.LazyPtr(in.GetAllowInsecure())
	return out
}
func SoftwareRecipe_Artifact_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Artifact) *pb.SoftwareRecipe_Artifact {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Artifact{}
	out.Id = direct.ValueOf(in.ID)
	if oneof := SoftwareRecipe_Artifact_Remote_ToProto(mapCtx, in.Remote); oneof != nil {
		out.Artifact = &pb.SoftwareRecipe_Artifact_Remote_{Remote: oneof}
	}
	if oneof := SoftwareRecipe_Artifact_Gcs_ToProto(mapCtx, in.Gcs); oneof != nil {
		out.Artifact = &pb.SoftwareRecipe_Artifact_Gcs_{Gcs: oneof}
	}
	out.AllowInsecure = direct.ValueOf(in.AllowInsecure)
	return out
}
func SoftwareRecipe_Artifact_Gcs_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Artifact_Gcs) *krm.SoftwareRecipe_Artifact_Gcs {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Artifact_Gcs{}
	out.Bucket = direct.LazyPtr(in.GetBucket())
	out.Object = direct.LazyPtr(in.GetObject())
	out.Generation = direct.LazyPtr(in.GetGeneration())
	return out
}
func SoftwareRecipe_Artifact_Gcs_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Artifact_Gcs) *pb.SoftwareRecipe_Artifact_Gcs {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Artifact_Gcs{}
	out.Bucket = direct.ValueOf(in.Bucket)
	out.Object = direct.ValueOf(in.Object)
	out.Generation = direct.ValueOf(in.Generation)
	return out
}
func SoftwareRecipe_Artifact_Remote_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Artifact_Remote) *krm.SoftwareRecipe_Artifact_Remote {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Artifact_Remote{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.Checksum = direct.LazyPtr(in.GetChecksum())
	return out
}
func SoftwareRecipe_Artifact_Remote_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Artifact_Remote) *pb.SoftwareRecipe_Artifact_Remote {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Artifact_Remote{}
	out.Uri = direct.ValueOf(in.URI)
	out.Checksum = direct.ValueOf(in.Checksum)
	return out
}
func SoftwareRecipe_Step_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Step) *krm.SoftwareRecipe_Step {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step{}
	out.FileCopy = SoftwareRecipe_Step_CopyFile_FromProto(mapCtx, in.GetFileCopy())
	out.ArchiveExtraction = SoftwareRecipe_Step_ExtractArchive_FromProto(mapCtx, in.GetArchiveExtraction())
	out.MsiInstallation = SoftwareRecipe_Step_InstallMsi_FromProto(mapCtx, in.GetMsiInstallation())
	out.DpkgInstallation = SoftwareRecipe_Step_InstallDpkg_FromProto(mapCtx, in.GetDpkgInstallation())
	out.RpmInstallation = SoftwareRecipe_Step_InstallRpm_FromProto(mapCtx, in.GetRpmInstallation())
	out.FileExec = SoftwareRecipe_Step_ExecFile_FromProto(mapCtx, in.GetFileExec())
	out.ScriptRun = SoftwareRecipe_Step_RunScript_FromProto(mapCtx, in.GetScriptRun())
	return out
}
func SoftwareRecipe_Step_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step) *pb.SoftwareRecipe_Step {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Step{}
	if oneof := SoftwareRecipe_Step_CopyFile_ToProto(mapCtx, in.FileCopy); oneof != nil {
		out.Step = &pb.SoftwareRecipe_Step_FileCopy{FileCopy: oneof}
	}
	if oneof := SoftwareRecipe_Step_ExtractArchive_ToProto(mapCtx, in.ArchiveExtraction); oneof != nil {
		out.Step = &pb.SoftwareRecipe_Step_ArchiveExtraction{ArchiveExtraction: oneof}
	}
	if oneof := SoftwareRecipe_Step_InstallMsi_ToProto(mapCtx, in.MsiInstallation); oneof != nil {
		out.Step = &pb.SoftwareRecipe_Step_MsiInstallation{MsiInstallation: oneof}
	}
	if oneof := SoftwareRecipe_Step_InstallDpkg_ToProto(mapCtx, in.DpkgInstallation); oneof != nil {
		out.Step = &pb.SoftwareRecipe_Step_DpkgInstallation{DpkgInstallation: oneof}
	}
	if oneof := SoftwareRecipe_Step_InstallRpm_ToProto(mapCtx, in.RpmInstallation); oneof != nil {
		out.Step = &pb.SoftwareRecipe_Step_RpmInstallation{RpmInstallation: oneof}
	}
	if oneof := SoftwareRecipe_Step_ExecFile_ToProto(mapCtx, in.FileExec); oneof != nil {
		out.Step = &pb.SoftwareRecipe_Step_FileExec{FileExec: oneof}
	}
	if oneof := SoftwareRecipe_Step_RunScript_ToProto(mapCtx, in.ScriptRun); oneof != nil {
		out.Step = &pb.SoftwareRecipe_Step_ScriptRun{ScriptRun: oneof}
	}
	return out
}
func SoftwareRecipe_Step_CopyFile_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Step_CopyFile) *krm.SoftwareRecipe_Step_CopyFile {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_CopyFile{}
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	out.Destination = direct.LazyPtr(in.GetDestination())
	out.Overwrite = direct.LazyPtr(in.GetOverwrite())
	out.Permissions = direct.LazyPtr(in.GetPermissions())
	return out
}
func SoftwareRecipe_Step_CopyFile_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_CopyFile) *pb.SoftwareRecipe_Step_CopyFile {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Step_CopyFile{}
	out.ArtifactId = direct.ValueOf(in.ArtifactID)
	out.Destination = direct.ValueOf(in.Destination)
	out.Overwrite = direct.ValueOf(in.Overwrite)
	out.Permissions = direct.ValueOf(in.Permissions)
	return out
}
func SoftwareRecipe_Step_ExecFile_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Step_ExecFile) *krm.SoftwareRecipe_Step_ExecFile {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_ExecFile{}
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	out.LocalPath = direct.LazyPtr(in.GetLocalPath())
	out.Args = in.Args
	out.AllowedExitCodes = in.AllowedExitCodes
	return out
}
func SoftwareRecipe_Step_ExecFile_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_ExecFile) *pb.SoftwareRecipe_Step_ExecFile {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Step_ExecFile{}
	if oneof := SoftwareRecipe_Step_ExecFile_ArtifactId_ToProto(mapCtx, in.ArtifactID); oneof != nil {
		out.LocationType = oneof
	}
	if oneof := SoftwareRecipe_Step_ExecFile_LocalPath_ToProto(mapCtx, in.LocalPath); oneof != nil {
		out.LocationType = oneof
	}
	out.Args = in.Args
	out.AllowedExitCodes = in.AllowedExitCodes
	return out
}
func SoftwareRecipe_Step_ExtractArchive_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Step_ExtractArchive) *krm.SoftwareRecipe_Step_ExtractArchive {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_ExtractArchive{}
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	out.Destination = direct.LazyPtr(in.GetDestination())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	return out
}
func SoftwareRecipe_Step_ExtractArchive_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_ExtractArchive) *pb.SoftwareRecipe_Step_ExtractArchive {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Step_ExtractArchive{}
	out.ArtifactId = direct.ValueOf(in.ArtifactID)
	out.Destination = direct.ValueOf(in.Destination)
	out.Type = direct.Enum_ToProto[pb.SoftwareRecipe_Step_ExtractArchive_ArchiveType](mapCtx, in.Type)
	return out
}
func SoftwareRecipe_Step_InstallDpkg_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Step_InstallDpkg) *krm.SoftwareRecipe_Step_InstallDpkg {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_InstallDpkg{}
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	return out
}
func SoftwareRecipe_Step_InstallDpkg_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_InstallDpkg) *pb.SoftwareRecipe_Step_InstallDpkg {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Step_InstallDpkg{}
	out.ArtifactId = direct.ValueOf(in.ArtifactID)
	return out
}
func SoftwareRecipe_Step_InstallMsi_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Step_InstallMsi) *krm.SoftwareRecipe_Step_InstallMsi {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_InstallMsi{}
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	out.Flags = in.Flags
	out.AllowedExitCodes = in.AllowedExitCodes
	return out
}
func SoftwareRecipe_Step_InstallMsi_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_InstallMsi) *pb.SoftwareRecipe_Step_InstallMsi {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Step_InstallMsi{}
	out.ArtifactId = direct.ValueOf(in.ArtifactID)
	out.Flags = in.Flags
	out.AllowedExitCodes = in.AllowedExitCodes
	return out
}
func SoftwareRecipe_Step_InstallRpm_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Step_InstallRpm) *krm.SoftwareRecipe_Step_InstallRpm {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_InstallRpm{}
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	return out
}
func SoftwareRecipe_Step_InstallRpm_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_InstallRpm) *pb.SoftwareRecipe_Step_InstallRpm {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Step_InstallRpm{}
	out.ArtifactId = direct.ValueOf(in.ArtifactID)
	return out
}
func SoftwareRecipe_Step_RunScript_FromProto(mapCtx *direct.MapContext, in *pb.SoftwareRecipe_Step_RunScript) *krm.SoftwareRecipe_Step_RunScript {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_RunScript{}
	out.Script = direct.LazyPtr(in.GetScript())
	out.AllowedExitCodes = in.AllowedExitCodes
	out.Interpreter = direct.Enum_FromProto(mapCtx, in.GetInterpreter())
	return out
}
func SoftwareRecipe_Step_RunScript_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_RunScript) *pb.SoftwareRecipe_Step_RunScript {
	if in == nil {
		return nil
	}
	out := &pb.SoftwareRecipe_Step_RunScript{}
	out.Script = direct.ValueOf(in.Script)
	out.AllowedExitCodes = in.AllowedExitCodes
	out.Interpreter = direct.Enum_ToProto[pb.SoftwareRecipe_Step_RunScript_Interpreter](mapCtx, in.Interpreter)
	return out
}
func YumRepository_FromProto(mapCtx *direct.MapContext, in *pb.YumRepository) *krm.YumRepository {
	if in == nil {
		return nil
	}
	out := &krm.YumRepository{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BaseURL = direct.LazyPtr(in.GetBaseUrl())
	out.GpgKeys = in.GpgKeys
	return out
}
func YumRepository_ToProto(mapCtx *direct.MapContext, in *krm.YumRepository) *pb.YumRepository {
	if in == nil {
		return nil
	}
	out := &pb.YumRepository{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BaseUrl = direct.ValueOf(in.BaseURL)
	out.GpgKeys = in.GpgKeys
	return out
}
func ZypperRepository_FromProto(mapCtx *direct.MapContext, in *pb.ZypperRepository) *krm.ZypperRepository {
	if in == nil {
		return nil
	}
	out := &krm.ZypperRepository{}
	out.ID = direct.LazyPtr(in.GetId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BaseURL = direct.LazyPtr(in.GetBaseUrl())
	out.GpgKeys = in.GpgKeys
	return out
}
func ZypperRepository_ToProto(mapCtx *direct.MapContext, in *krm.ZypperRepository) *pb.ZypperRepository {
	if in == nil {
		return nil
	}
	out := &pb.ZypperRepository{}
	out.Id = direct.ValueOf(in.ID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BaseUrl = direct.ValueOf(in.BaseURL)
	out.GpgKeys = in.GpgKeys
	return out
}
