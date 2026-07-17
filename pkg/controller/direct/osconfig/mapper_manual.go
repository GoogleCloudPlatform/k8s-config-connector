// Copyright 2026 Google LLC
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
	"time"

	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
	osconfigpb "cloud.google.com/go/osconfig/apiv1beta/osconfigpb"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FixedOrPercent_Fixed_ToProto(mapCtx *direct.MapContext, in *int64) *pb.FixedOrPercent_Fixed {
	if in == nil {
		return nil
	}
	return &pb.FixedOrPercent_Fixed{Fixed: int32(*in)}
}

func FixedOrPercent_Percent_ToProto(mapCtx *direct.MapContext, in *int64) *pb.FixedOrPercent_Percent {
	if in == nil {
		return nil
	}
	return &pb.FixedOrPercent_Percent{Percent: int32(*in)}
}

func FixedOrPercent_FromProto(mapCtx *direct.MapContext, in *pb.FixedOrPercent) *krm.FixedOrPercent {
	if in == nil {
		return nil
	}
	out := &krm.FixedOrPercent{}
	if in.Mode != nil {
		switch m := in.Mode.(type) {
		case *pb.FixedOrPercent_Fixed:
			out.Fixed = direct.LazyPtr(int64(m.Fixed))
		case *pb.FixedOrPercent_Percent:
			out.Percent = direct.LazyPtr(int64(m.Percent))
		}
	}
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

// The following functions were manually copied from the generated code
// to support primitive values (instead of pointers) for required fields.

func OSPolicy_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy) *krm.OSPolicy {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy{}
	out.ID = in.GetId()
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Mode = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetMode()))
	out.ResourceGroups = direct.Slice_FromProto(mapCtx, in.ResourceGroups, OSPolicy_ResourceGroup_FromProto)
	out.AllowNoResourceGroupMatch = direct.LazyPtr(in.GetAllowNoResourceGroupMatch())
	return out
}

func OSPolicy_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy) *pb.OSPolicy {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy{}
	out.Id = in.ID
	out.Description = direct.ValueOf(in.Description)
	out.Mode = direct.Enum_ToProto[pb.OSPolicy_Mode](mapCtx, &in.Mode)
	out.ResourceGroups = direct.Slice_ToProto(mapCtx, in.ResourceGroups, OSPolicy_ResourceGroup_ToProto)
	out.AllowNoResourceGroupMatch = direct.ValueOf(in.AllowNoResourceGroupMatch)
	return out
}

func OSPolicyAssignment_InstanceFilter_Inventory_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment_InstanceFilter_Inventory) *krm.OSPolicyAssignment_InstanceFilter_Inventory {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignment_InstanceFilter_Inventory{}
	out.OSShortName = in.GetOsShortName()
	out.OSVersion = direct.LazyPtr(in.GetOsVersion())
	return out
}

func OSPolicyAssignment_InstanceFilter_Inventory_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignment_InstanceFilter_Inventory) *pb.OSPolicyAssignment_InstanceFilter_Inventory {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment_InstanceFilter_Inventory{}
	out.OsShortName = in.OSShortName
	out.OsVersion = direct.ValueOf(in.OSVersion)
	return out
}

func OSPolicyAssignment_Rollout_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicyAssignment_Rollout) *krm.OSPolicyAssignment_Rollout {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicyAssignment_Rollout{}
	out.DisruptionBudget = FixedOrPercent_FromProto(mapCtx, in.GetDisruptionBudget())
	out.MinWaitDuration = direct.ValueOf(direct.StringDuration_FromProto(mapCtx, in.GetMinWaitDuration()))
	return out
}

func OSPolicyAssignment_Rollout_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicyAssignment_Rollout) *pb.OSPolicyAssignment_Rollout {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicyAssignment_Rollout{}
	out.DisruptionBudget = FixedOrPercent_ToProto(mapCtx, in.DisruptionBudget)
	out.MinWaitDuration = direct.StringDuration_ToProto(mapCtx, &in.MinWaitDuration)
	return out
}

func OSPolicy_InventoryFilter_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_InventoryFilter) *krm.OSPolicy_InventoryFilter {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_InventoryFilter{}
	out.OSShortName = in.GetOsShortName()
	out.OSVersion = direct.LazyPtr(in.GetOsVersion())
	return out
}

func OSPolicy_InventoryFilter_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_InventoryFilter) *pb.OSPolicy_InventoryFilter {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_InventoryFilter{}
	out.OsShortName = in.OSShortName
	out.OsVersion = direct.ValueOf(in.OSVersion)
	return out
}

func OSPolicy_Resource_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource) *krm.OSPolicy_Resource {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource{}
	out.ID = in.GetId()
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
	out.Id = in.ID
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
		out.ResourceType = &pb.OSPolicy_Resource_File_{File: oneof}
	}
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
	out.Interpreter = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetInterpreter()))
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
	out.Interpreter = direct.Enum_ToProto[pb.OSPolicy_Resource_ExecResource_Exec_Interpreter](mapCtx, &in.Interpreter)
	out.OutputFilePath = direct.ValueOf(in.OutputFilePath)
	return out
}

func OSPolicy_Resource_FileResource_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_FileResource) *krm.OSPolicy_Resource_FileResource {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_FileResource{}
	out.File = OSPolicy_Resource_File_FromProto(mapCtx, in.GetFile())
	out.Content = direct.LazyPtr(in.GetContent())
	out.Path = in.GetPath()
	out.State = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetState()))
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
	out.Path = in.Path
	out.State = direct.Enum_ToProto[pb.OSPolicy_Resource_FileResource_DesiredState](mapCtx, &in.State)
	out.Permissions = direct.ValueOf(in.Permissions)
	return out
}

func OSPolicy_Resource_File_GCS_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_File_Gcs) *krm.OSPolicy_Resource_File_GCS {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_File_GCS{}
	out.Bucket = in.GetBucket()
	out.Object = in.GetObject()
	out.Generation = direct.LazyPtr(in.GetGeneration())
	return out
}

func OSPolicy_Resource_File_GCS_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_File_GCS) *pb.OSPolicy_Resource_File_Gcs {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_File_Gcs{}
	out.Bucket = in.Bucket
	out.Object = in.Object
	out.Generation = direct.ValueOf(in.Generation)
	return out
}

func OSPolicy_Resource_File_Remote_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_File_Remote) *krm.OSPolicy_Resource_File_Remote {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_File_Remote{}
	out.URI = in.GetUri()
	out.Sha256Checksum = direct.LazyPtr(in.GetSha256Checksum())
	return out
}

func OSPolicy_Resource_File_Remote_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_File_Remote) *pb.OSPolicy_Resource_File_Remote {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_File_Remote{}
	out.Uri = in.URI
	out.Sha256Checksum = direct.ValueOf(in.Sha256Checksum)
	return out
}

func OSPolicy_Resource_PackageResource_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource) *krm.OSPolicy_Resource_PackageResource {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource{}
	out.DesiredState = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetDesiredState()))
	out.Apt = OSPolicy_Resource_PackageResource_Apt_FromProto(mapCtx, in.GetApt())
	out.Deb = OSPolicy_Resource_PackageResource_Deb_FromProto(mapCtx, in.GetDeb())
	out.Yum = OSPolicy_Resource_PackageResource_Yum_FromProto(mapCtx, in.GetYum())
	out.Zypper = OSPolicy_Resource_PackageResource_Zypper_FromProto(mapCtx, in.GetZypper())
	out.Rpm = OSPolicy_Resource_PackageResource_Rpm_FromProto(mapCtx, in.GetRpm())
	out.Googet = OSPolicy_Resource_PackageResource_GooGet_FromProto(mapCtx, in.GetGooget())
	out.Msi = OSPolicy_Resource_PackageResource_Msi_FromProto(mapCtx, in.GetMsi())
	return out
}

func OSPolicy_Resource_PackageResource_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource) *pb.OSPolicy_Resource_PackageResource {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource{}
	out.DesiredState = direct.Enum_ToProto[pb.OSPolicy_Resource_PackageResource_DesiredState](mapCtx, &in.DesiredState)
	if oneof := OSPolicy_Resource_PackageResource_Apt_ToProto(mapCtx, in.Apt); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Apt{Apt: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_Deb_ToProto(mapCtx, in.Deb); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Deb_{Deb: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_Yum_ToProto(mapCtx, in.Yum); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Yum{Yum: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_Zypper_ToProto(mapCtx, in.Zypper); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Zypper_{Zypper: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_Rpm_ToProto(mapCtx, in.Rpm); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Rpm{Rpm: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_GooGet_ToProto(mapCtx, in.Googet); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Googet{Googet: oneof}
	}
	if oneof := OSPolicy_Resource_PackageResource_Msi_ToProto(mapCtx, in.Msi); oneof != nil {
		out.SystemPackage = &pb.OSPolicy_Resource_PackageResource_Msi{Msi: oneof}
	}
	return out
}

func OSPolicy_Resource_PackageResource_Apt_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_APT) *krm.OSPolicy_Resource_PackageResource_Apt {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_Apt{}
	out.Name = in.GetName()
	return out
}

func OSPolicy_Resource_PackageResource_Apt_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_Apt) *pb.OSPolicy_Resource_PackageResource_APT {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_APT{}
	out.Name = in.Name
	return out
}

func OSPolicy_Resource_PackageResource_GooGet_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_GooGet) *krm.OSPolicy_Resource_PackageResource_GooGet {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_GooGet{}
	out.Name = in.GetName()
	return out
}

func OSPolicy_Resource_PackageResource_GooGet_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_GooGet) *pb.OSPolicy_Resource_PackageResource_GooGet {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_GooGet{}
	out.Name = in.Name
	return out
}

func OSPolicy_Resource_PackageResource_Yum_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_YUM) *krm.OSPolicy_Resource_PackageResource_Yum {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_Yum{}
	out.Name = in.GetName()
	return out
}

func OSPolicy_Resource_PackageResource_Yum_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_Yum) *pb.OSPolicy_Resource_PackageResource_YUM {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_YUM{}
	out.Name = in.Name
	return out
}

func OSPolicy_Resource_PackageResource_Zypper_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_PackageResource_Zypper) *krm.OSPolicy_Resource_PackageResource_Zypper {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_PackageResource_Zypper{}
	out.Name = in.GetName()
	return out
}

func OSPolicy_Resource_PackageResource_Zypper_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_PackageResource_Zypper) *pb.OSPolicy_Resource_PackageResource_Zypper {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_PackageResource_Zypper{}
	out.Name = in.Name
	return out
}

func OSPolicy_Resource_RepositoryResource_AptRepository_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_RepositoryResource_AptRepository) *krm.OSPolicy_Resource_RepositoryResource_AptRepository {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_RepositoryResource_AptRepository{}
	out.ArchiveType = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetArchiveType()))
	out.URI = in.GetUri()
	out.Distribution = in.GetDistribution()
	out.Components = in.Components
	out.GpgKey = direct.LazyPtr(in.GetGpgKey())
	return out
}

func OSPolicy_Resource_RepositoryResource_AptRepository_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_RepositoryResource_AptRepository) *pb.OSPolicy_Resource_RepositoryResource_AptRepository {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_RepositoryResource_AptRepository{}
	out.ArchiveType = direct.Enum_ToProto[pb.OSPolicy_Resource_RepositoryResource_AptRepository_ArchiveType](mapCtx, &in.ArchiveType)
	out.Uri = in.URI
	out.Distribution = in.Distribution
	out.Components = in.Components
	out.GpgKey = direct.ValueOf(in.GpgKey)
	return out
}

func OSPolicy_Resource_RepositoryResource_GooRepository_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_RepositoryResource_GooRepository) *krm.OSPolicy_Resource_RepositoryResource_GooRepository {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_RepositoryResource_GooRepository{}
	out.Name = in.GetName()
	out.URL = in.GetUrl()
	return out
}

func OSPolicy_Resource_RepositoryResource_GooRepository_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_RepositoryResource_GooRepository) *pb.OSPolicy_Resource_RepositoryResource_GooRepository {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_RepositoryResource_GooRepository{}
	out.Name = in.Name
	out.Url = in.URL
	return out
}

func OSPolicy_Resource_RepositoryResource_YumRepository_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_RepositoryResource_YumRepository) *krm.OSPolicy_Resource_RepositoryResource_YumRepository {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_RepositoryResource_YumRepository{}
	out.ID = in.GetId()
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BaseURL = in.GetBaseUrl()
	out.GpgKeys = in.GpgKeys
	return out
}

func OSPolicy_Resource_RepositoryResource_YumRepository_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_RepositoryResource_YumRepository) *pb.OSPolicy_Resource_RepositoryResource_YumRepository {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_RepositoryResource_YumRepository{}
	out.Id = in.ID
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BaseUrl = in.BaseURL
	out.GpgKeys = in.GpgKeys
	return out
}

func OSPolicy_Resource_RepositoryResource_ZypperRepository_FromProto(mapCtx *direct.MapContext, in *pb.OSPolicy_Resource_RepositoryResource_ZypperRepository) *krm.OSPolicy_Resource_RepositoryResource_ZypperRepository {
	if in == nil {
		return nil
	}
	out := &krm.OSPolicy_Resource_RepositoryResource_ZypperRepository{}
	out.ID = in.GetId()
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BaseURL = in.GetBaseUrl()
	out.GpgKeys = in.GpgKeys
	return out
}

func OSPolicy_Resource_RepositoryResource_ZypperRepository_ToProto(mapCtx *direct.MapContext, in *krm.OSPolicy_Resource_RepositoryResource_ZypperRepository) *pb.OSPolicy_Resource_RepositoryResource_ZypperRepository {
	if in == nil {
		return nil
	}
	out := &pb.OSPolicy_Resource_RepositoryResource_ZypperRepository{}
	out.Id = in.ID
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BaseUrl = in.BaseURL
	out.GpgKeys = in.GpgKeys
	return out
}

func Assignment_Instances_FromProto(mapCtx *direct.MapContext, in []string) []computev1beta1.InstanceRef {
	if in == nil {
		return nil
	}
	out := make([]computev1beta1.InstanceRef, len(in))
	for i, v := range in {
		out[i] = computev1beta1.InstanceRef{External: v}
	}
	return out
}

func Assignment_Instances_ToProto(mapCtx *direct.MapContext, in []computev1beta1.InstanceRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = v.External
	}
	return out
}

func int32Slice_FromProto(in []int32) []int64 {
	if in == nil {
		return nil
	}
	out := make([]int64, len(in))
	for i, v := range in {
		out[i] = int64(v)
	}
	return out
}

func int32Slice_ToProto(in []int64) []int32 {
	if in == nil {
		return nil
	}
	out := make([]int32, len(in))
	for i, v := range in {
		out[i] = int32(v)
	}
	return out
}

func SoftwareRecipe_Step_ExecFile_FromProto(mapCtx *direct.MapContext, in *osconfigpb.SoftwareRecipe_Step_ExecFile) *krm.SoftwareRecipe_Step_ExecFile {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_ExecFile{}
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	out.LocalPath = direct.LazyPtr(in.GetLocalPath())
	out.Args = in.Args
	out.AllowedExitCodes = int32Slice_FromProto(in.AllowedExitCodes)
	return out
}

func SoftwareRecipe_Step_ExecFile_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_ExecFile) *osconfigpb.SoftwareRecipe_Step_ExecFile {
	if in == nil {
		return nil
	}
	out := &osconfigpb.SoftwareRecipe_Step_ExecFile{}
	if oneof := SoftwareRecipe_Step_ExecFile_ArtifactId_ToProto(mapCtx, in.ArtifactID); oneof != nil {
		out.LocationType = oneof
	}
	if oneof := SoftwareRecipe_Step_ExecFile_LocalPath_ToProto(mapCtx, in.LocalPath); oneof != nil {
		out.LocationType = oneof
	}
	out.Args = in.Args
	out.AllowedExitCodes = int32Slice_ToProto(in.AllowedExitCodes)
	return out
}

func SoftwareRecipe_Step_InstallMsi_FromProto(mapCtx *direct.MapContext, in *osconfigpb.SoftwareRecipe_Step_InstallMsi) *krm.SoftwareRecipe_Step_InstallMsi {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_InstallMsi{}
	out.ArtifactID = direct.LazyPtr(in.GetArtifactId())
	out.Flags = in.Flags
	out.AllowedExitCodes = int32Slice_FromProto(in.AllowedExitCodes)
	return out
}

func SoftwareRecipe_Step_InstallMsi_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_InstallMsi) *osconfigpb.SoftwareRecipe_Step_InstallMsi {
	if in == nil {
		return nil
	}
	out := &osconfigpb.SoftwareRecipe_Step_InstallMsi{}
	out.ArtifactId = direct.ValueOf(in.ArtifactID)
	out.Flags = in.Flags
	out.AllowedExitCodes = int32Slice_ToProto(in.AllowedExitCodes)
	return out
}

func SoftwareRecipe_Step_RunScript_FromProto(mapCtx *direct.MapContext, in *osconfigpb.SoftwareRecipe_Step_RunScript) *krm.SoftwareRecipe_Step_RunScript {
	if in == nil {
		return nil
	}
	out := &krm.SoftwareRecipe_Step_RunScript{}
	out.Script = direct.LazyPtr(in.GetScript())
	out.AllowedExitCodes = int32Slice_FromProto(in.AllowedExitCodes)
	out.Interpreter = direct.Enum_FromProto(mapCtx, in.GetInterpreter())
	return out
}

func SoftwareRecipe_Step_RunScript_ToProto(mapCtx *direct.MapContext, in *krm.SoftwareRecipe_Step_RunScript) *osconfigpb.SoftwareRecipe_Step_RunScript {
	if in == nil {
		return nil
	}
	out := &osconfigpb.SoftwareRecipe_Step_RunScript{}
	out.Script = direct.ValueOf(in.Script)
	out.AllowedExitCodes = int32Slice_ToProto(in.AllowedExitCodes)
	out.Interpreter = direct.Enum_ToProto[osconfigpb.SoftwareRecipe_Step_RunScript_Interpreter](mapCtx, in.Interpreter)
	return out
}

func OSConfigGuestPolicyStatus_FromProto(mapCtx *direct.MapContext, in *osconfigpb.GuestPolicy) *krm.OSConfigGuestPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.OSConfigGuestPolicyStatus{}
	if in.CreateTime != nil {
		out.CreateTime = direct.LazyPtr(in.CreateTime.AsTime().Format(time.RFC3339Nano))
	}
	if in.UpdateTime != nil {
		out.UpdateTime = direct.LazyPtr(in.UpdateTime.AsTime().Format(time.RFC3339Nano))
	}
	if in.Etag != "" {
		out.Etag = direct.LazyPtr(in.Etag)
	}
	return out
}

func OSConfigGuestPolicyStatus_ToProto(mapCtx *direct.MapContext, in *krm.OSConfigGuestPolicyStatus) *osconfigpb.GuestPolicy {
	if in == nil {
		return nil
	}
	out := &osconfigpb.GuestPolicy{}
	if in.CreateTime != nil {
		t, err := time.Parse(time.RFC3339Nano, *in.CreateTime)
		if err == nil {
			out.CreateTime = timestamppb.New(t)
		} else {
			mapCtx.Errorf("parsing createTime: %w", err)
		}
	}
	if in.UpdateTime != nil {
		t, err := time.Parse(time.RFC3339Nano, *in.UpdateTime)
		if err == nil {
			out.UpdateTime = timestamppb.New(t)
		} else {
			mapCtx.Errorf("parsing updateTime: %w", err)
		}
	}
	if in.Etag != nil {
		out.Etag = *in.Etag
	}
	return out
}
