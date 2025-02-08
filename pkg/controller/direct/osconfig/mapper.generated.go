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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/osconfig/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/osconfig/apiv1/osconfigpb"
)
func Inventory_FromProto(mapCtx *direct.MapContext, in *pb.Inventory) *krm.Inventory {
	if in == nil {
		return nil
	}
	out := &krm.Inventory{}
	// MISSING: Name
	out.OsInfo = Inventory_OsInfo_FromProto(mapCtx, in.GetOsInfo())
	// MISSING: Items
	// MISSING: UpdateTime
	return out
}
func Inventory_ToProto(mapCtx *direct.MapContext, in *krm.Inventory) *pb.Inventory {
	if in == nil {
		return nil
	}
	out := &pb.Inventory{}
	// MISSING: Name
	out.OsInfo = Inventory_OsInfo_ToProto(mapCtx, in.OsInfo)
	// MISSING: Items
	// MISSING: UpdateTime
	return out
}
func InventoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Inventory) *krm.InventoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.InventoryObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: OsInfo
	// MISSING: Items
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func InventoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.InventoryObservedState) *pb.Inventory {
	if in == nil {
		return nil
	}
	out := &pb.Inventory{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: OsInfo
	// MISSING: Items
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func Inventory_Item_FromProto(mapCtx *direct.MapContext, in *pb.Inventory_Item) *krm.Inventory_Item {
	if in == nil {
		return nil
	}
	out := &krm.Inventory_Item{}
	out.ID = direct.LazyPtr(in.GetId())
	out.OriginType = direct.Enum_FromProto(mapCtx, in.GetOriginType())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.InstalledPackage = Inventory_SoftwarePackage_FromProto(mapCtx, in.GetInstalledPackage())
	out.AvailablePackage = Inventory_SoftwarePackage_FromProto(mapCtx, in.GetAvailablePackage())
	return out
}
func Inventory_Item_ToProto(mapCtx *direct.MapContext, in *krm.Inventory_Item) *pb.Inventory_Item {
	if in == nil {
		return nil
	}
	out := &pb.Inventory_Item{}
	out.Id = direct.ValueOf(in.ID)
	out.OriginType = direct.Enum_ToProto[pb.Inventory_Item_OriginType](mapCtx, in.OriginType)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Type = direct.Enum_ToProto[pb.Inventory_Item_Type](mapCtx, in.Type)
	if oneof := Inventory_SoftwarePackage_ToProto(mapCtx, in.InstalledPackage); oneof != nil {
		out.Details = &pb.Inventory_Item_InstalledPackage{InstalledPackage: oneof}
	}
	if oneof := Inventory_SoftwarePackage_ToProto(mapCtx, in.AvailablePackage); oneof != nil {
		out.Details = &pb.Inventory_Item_AvailablePackage{AvailablePackage: oneof}
	}
	return out
}
func Inventory_OsInfo_FromProto(mapCtx *direct.MapContext, in *pb.Inventory_OsInfo) *krm.Inventory_OsInfo {
	if in == nil {
		return nil
	}
	out := &krm.Inventory_OsInfo{}
	out.Hostname = direct.LazyPtr(in.GetHostname())
	out.LongName = direct.LazyPtr(in.GetLongName())
	out.ShortName = direct.LazyPtr(in.GetShortName())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Architecture = direct.LazyPtr(in.GetArchitecture())
	out.KernelVersion = direct.LazyPtr(in.GetKernelVersion())
	out.KernelRelease = direct.LazyPtr(in.GetKernelRelease())
	out.OsconfigAgentVersion = direct.LazyPtr(in.GetOsconfigAgentVersion())
	return out
}
func Inventory_OsInfo_ToProto(mapCtx *direct.MapContext, in *krm.Inventory_OsInfo) *pb.Inventory_OsInfo {
	if in == nil {
		return nil
	}
	out := &pb.Inventory_OsInfo{}
	out.Hostname = direct.ValueOf(in.Hostname)
	out.LongName = direct.ValueOf(in.LongName)
	out.ShortName = direct.ValueOf(in.ShortName)
	out.Version = direct.ValueOf(in.Version)
	out.Architecture = direct.ValueOf(in.Architecture)
	out.KernelVersion = direct.ValueOf(in.KernelVersion)
	out.KernelRelease = direct.ValueOf(in.KernelRelease)
	out.OsconfigAgentVersion = direct.ValueOf(in.OsconfigAgentVersion)
	return out
}
func Inventory_SoftwarePackage_FromProto(mapCtx *direct.MapContext, in *pb.Inventory_SoftwarePackage) *krm.Inventory_SoftwarePackage {
	if in == nil {
		return nil
	}
	out := &krm.Inventory_SoftwarePackage{}
	out.YumPackage = Inventory_VersionedPackage_FromProto(mapCtx, in.GetYumPackage())
	out.AptPackage = Inventory_VersionedPackage_FromProto(mapCtx, in.GetAptPackage())
	out.ZypperPackage = Inventory_VersionedPackage_FromProto(mapCtx, in.GetZypperPackage())
	out.GoogetPackage = Inventory_VersionedPackage_FromProto(mapCtx, in.GetGoogetPackage())
	out.ZypperPatch = Inventory_ZypperPatch_FromProto(mapCtx, in.GetZypperPatch())
	out.WuaPackage = Inventory_WindowsUpdatePackage_FromProto(mapCtx, in.GetWuaPackage())
	out.QfePackage = Inventory_WindowsQuickFixEngineeringPackage_FromProto(mapCtx, in.GetQfePackage())
	out.CosPackage = Inventory_VersionedPackage_FromProto(mapCtx, in.GetCosPackage())
	out.WindowsApplication = Inventory_WindowsApplication_FromProto(mapCtx, in.GetWindowsApplication())
	return out
}
func Inventory_SoftwarePackage_ToProto(mapCtx *direct.MapContext, in *krm.Inventory_SoftwarePackage) *pb.Inventory_SoftwarePackage {
	if in == nil {
		return nil
	}
	out := &pb.Inventory_SoftwarePackage{}
	if oneof := Inventory_VersionedPackage_ToProto(mapCtx, in.YumPackage); oneof != nil {
		out.Details = &pb.Inventory_SoftwarePackage_YumPackage{YumPackage: oneof}
	}
	if oneof := Inventory_VersionedPackage_ToProto(mapCtx, in.AptPackage); oneof != nil {
		out.Details = &pb.Inventory_SoftwarePackage_AptPackage{AptPackage: oneof}
	}
	if oneof := Inventory_VersionedPackage_ToProto(mapCtx, in.ZypperPackage); oneof != nil {
		out.Details = &pb.Inventory_SoftwarePackage_ZypperPackage{ZypperPackage: oneof}
	}
	if oneof := Inventory_VersionedPackage_ToProto(mapCtx, in.GoogetPackage); oneof != nil {
		out.Details = &pb.Inventory_SoftwarePackage_GoogetPackage{GoogetPackage: oneof}
	}
	if oneof := Inventory_ZypperPatch_ToProto(mapCtx, in.ZypperPatch); oneof != nil {
		out.Details = &pb.Inventory_SoftwarePackage_ZypperPatch{ZypperPatch: oneof}
	}
	if oneof := Inventory_WindowsUpdatePackage_ToProto(mapCtx, in.WuaPackage); oneof != nil {
		out.Details = &pb.Inventory_SoftwarePackage_WuaPackage{WuaPackage: oneof}
	}
	if oneof := Inventory_WindowsQuickFixEngineeringPackage_ToProto(mapCtx, in.QfePackage); oneof != nil {
		out.Details = &pb.Inventory_SoftwarePackage_QfePackage{QfePackage: oneof}
	}
	if oneof := Inventory_VersionedPackage_ToProto(mapCtx, in.CosPackage); oneof != nil {
		out.Details = &pb.Inventory_SoftwarePackage_CosPackage{CosPackage: oneof}
	}
	if oneof := Inventory_WindowsApplication_ToProto(mapCtx, in.WindowsApplication); oneof != nil {
		out.Details = &pb.Inventory_SoftwarePackage_WindowsApplication{WindowsApplication: oneof}
	}
	return out
}
func Inventory_VersionedPackage_FromProto(mapCtx *direct.MapContext, in *pb.Inventory_VersionedPackage) *krm.Inventory_VersionedPackage {
	if in == nil {
		return nil
	}
	out := &krm.Inventory_VersionedPackage{}
	out.PackageName = direct.LazyPtr(in.GetPackageName())
	out.Architecture = direct.LazyPtr(in.GetArchitecture())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func Inventory_VersionedPackage_ToProto(mapCtx *direct.MapContext, in *krm.Inventory_VersionedPackage) *pb.Inventory_VersionedPackage {
	if in == nil {
		return nil
	}
	out := &pb.Inventory_VersionedPackage{}
	out.PackageName = direct.ValueOf(in.PackageName)
	out.Architecture = direct.ValueOf(in.Architecture)
	out.Version = direct.ValueOf(in.Version)
	return out
}
func Inventory_WindowsApplication_FromProto(mapCtx *direct.MapContext, in *pb.Inventory_WindowsApplication) *krm.Inventory_WindowsApplication {
	if in == nil {
		return nil
	}
	out := &krm.Inventory_WindowsApplication{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DisplayVersion = direct.LazyPtr(in.GetDisplayVersion())
	out.Publisher = direct.LazyPtr(in.GetPublisher())
	out.InstallDate = Date_FromProto(mapCtx, in.GetInstallDate())
	out.HelpLink = direct.LazyPtr(in.GetHelpLink())
	return out
}
func Inventory_WindowsApplication_ToProto(mapCtx *direct.MapContext, in *krm.Inventory_WindowsApplication) *pb.Inventory_WindowsApplication {
	if in == nil {
		return nil
	}
	out := &pb.Inventory_WindowsApplication{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DisplayVersion = direct.ValueOf(in.DisplayVersion)
	out.Publisher = direct.ValueOf(in.Publisher)
	out.InstallDate = Date_ToProto(mapCtx, in.InstallDate)
	out.HelpLink = direct.ValueOf(in.HelpLink)
	return out
}
func Inventory_WindowsQuickFixEngineeringPackage_FromProto(mapCtx *direct.MapContext, in *pb.Inventory_WindowsQuickFixEngineeringPackage) *krm.Inventory_WindowsQuickFixEngineeringPackage {
	if in == nil {
		return nil
	}
	out := &krm.Inventory_WindowsQuickFixEngineeringPackage{}
	out.Caption = direct.LazyPtr(in.GetCaption())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.HotFixID = direct.LazyPtr(in.GetHotFixId())
	out.InstallTime = direct.StringTimestamp_FromProto(mapCtx, in.GetInstallTime())
	return out
}
func Inventory_WindowsQuickFixEngineeringPackage_ToProto(mapCtx *direct.MapContext, in *krm.Inventory_WindowsQuickFixEngineeringPackage) *pb.Inventory_WindowsQuickFixEngineeringPackage {
	if in == nil {
		return nil
	}
	out := &pb.Inventory_WindowsQuickFixEngineeringPackage{}
	out.Caption = direct.ValueOf(in.Caption)
	out.Description = direct.ValueOf(in.Description)
	out.HotFixId = direct.ValueOf(in.HotFixID)
	out.InstallTime = direct.StringTimestamp_ToProto(mapCtx, in.InstallTime)
	return out
}
func Inventory_WindowsUpdatePackage_FromProto(mapCtx *direct.MapContext, in *pb.Inventory_WindowsUpdatePackage) *krm.Inventory_WindowsUpdatePackage {
	if in == nil {
		return nil
	}
	out := &krm.Inventory_WindowsUpdatePackage{}
	out.Title = direct.LazyPtr(in.GetTitle())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Categories = direct.Slice_FromProto(mapCtx, in.Categories, Inventory_WindowsUpdatePackage_WindowsUpdateCategory_FromProto)
	out.KbArticleIds = in.KbArticleIds
	out.SupportURL = direct.LazyPtr(in.GetSupportUrl())
	out.MoreInfoUrls = in.MoreInfoUrls
	out.UpdateID = direct.LazyPtr(in.GetUpdateId())
	out.RevisionNumber = direct.LazyPtr(in.GetRevisionNumber())
	out.LastDeploymentChangeTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastDeploymentChangeTime())
	return out
}
func Inventory_WindowsUpdatePackage_ToProto(mapCtx *direct.MapContext, in *krm.Inventory_WindowsUpdatePackage) *pb.Inventory_WindowsUpdatePackage {
	if in == nil {
		return nil
	}
	out := &pb.Inventory_WindowsUpdatePackage{}
	out.Title = direct.ValueOf(in.Title)
	out.Description = direct.ValueOf(in.Description)
	out.Categories = direct.Slice_ToProto(mapCtx, in.Categories, Inventory_WindowsUpdatePackage_WindowsUpdateCategory_ToProto)
	out.KbArticleIds = in.KbArticleIds
	out.SupportUrl = direct.ValueOf(in.SupportURL)
	out.MoreInfoUrls = in.MoreInfoUrls
	out.UpdateId = direct.ValueOf(in.UpdateID)
	out.RevisionNumber = direct.ValueOf(in.RevisionNumber)
	out.LastDeploymentChangeTime = direct.StringTimestamp_ToProto(mapCtx, in.LastDeploymentChangeTime)
	return out
}
func Inventory_WindowsUpdatePackage_WindowsUpdateCategory_FromProto(mapCtx *direct.MapContext, in *pb.Inventory_WindowsUpdatePackage_WindowsUpdateCategory) *krm.Inventory_WindowsUpdatePackage_WindowsUpdateCategory {
	if in == nil {
		return nil
	}
	out := &krm.Inventory_WindowsUpdatePackage_WindowsUpdateCategory{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func Inventory_WindowsUpdatePackage_WindowsUpdateCategory_ToProto(mapCtx *direct.MapContext, in *krm.Inventory_WindowsUpdatePackage_WindowsUpdateCategory) *pb.Inventory_WindowsUpdatePackage_WindowsUpdateCategory {
	if in == nil {
		return nil
	}
	out := &pb.Inventory_WindowsUpdatePackage_WindowsUpdateCategory{}
	out.Id = direct.ValueOf(in.ID)
	out.Name = direct.ValueOf(in.Name)
	return out
}
func Inventory_ZypperPatch_FromProto(mapCtx *direct.MapContext, in *pb.Inventory_ZypperPatch) *krm.Inventory_ZypperPatch {
	if in == nil {
		return nil
	}
	out := &krm.Inventory_ZypperPatch{}
	out.PatchName = direct.LazyPtr(in.GetPatchName())
	out.Category = direct.LazyPtr(in.GetCategory())
	out.Severity = direct.LazyPtr(in.GetSeverity())
	out.Summary = direct.LazyPtr(in.GetSummary())
	return out
}
func Inventory_ZypperPatch_ToProto(mapCtx *direct.MapContext, in *krm.Inventory_ZypperPatch) *pb.Inventory_ZypperPatch {
	if in == nil {
		return nil
	}
	out := &pb.Inventory_ZypperPatch{}
	out.PatchName = direct.ValueOf(in.PatchName)
	out.Category = direct.ValueOf(in.Category)
	out.Severity = direct.ValueOf(in.Severity)
	out.Summary = direct.ValueOf(in.Summary)
	return out
}
func OsconfigInventoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Inventory) *krm.OsconfigInventoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigInventoryObservedState{}
	// MISSING: Name
	// MISSING: OsInfo
	// MISSING: Items
	// MISSING: UpdateTime
	return out
}
func OsconfigInventoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigInventoryObservedState) *pb.Inventory {
	if in == nil {
		return nil
	}
	out := &pb.Inventory{}
	// MISSING: Name
	// MISSING: OsInfo
	// MISSING: Items
	// MISSING: UpdateTime
	return out
}
func OsconfigInventorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Inventory) *krm.OsconfigInventorySpec {
	if in == nil {
		return nil
	}
	out := &krm.OsconfigInventorySpec{}
	// MISSING: Name
	// MISSING: OsInfo
	// MISSING: Items
	// MISSING: UpdateTime
	return out
}
func OsconfigInventorySpec_ToProto(mapCtx *direct.MapContext, in *krm.OsconfigInventorySpec) *pb.Inventory {
	if in == nil {
		return nil
	}
	out := &pb.Inventory{}
	// MISSING: Name
	// MISSING: OsInfo
	// MISSING: Items
	// MISSING: UpdateTime
	return out
}
