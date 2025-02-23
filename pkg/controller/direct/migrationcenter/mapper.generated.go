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

package migrationcenter

import (
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ImportDataFile_FromProto(mapCtx *direct.MapContext, in *pb.ImportDataFile) *krm.ImportDataFile {
	if in == nil {
		return nil
	}
	out := &krm.ImportDataFile{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	// MISSING: CreateTime
	// MISSING: State
	out.UploadFileInfo = UploadFileInfo_FromProto(mapCtx, in.GetUploadFileInfo())
	return out
}
func ImportDataFile_ToProto(mapCtx *direct.MapContext, in *krm.ImportDataFile) *pb.ImportDataFile {
	if in == nil {
		return nil
	}
	out := &pb.ImportDataFile{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Format = direct.Enum_ToProto[pb.ImportJobFormat](mapCtx, in.Format)
	// MISSING: CreateTime
	// MISSING: State
	if oneof := UploadFileInfo_ToProto(mapCtx, in.UploadFileInfo); oneof != nil {
		out.FileInfo = &pb.ImportDataFile_UploadFileInfo{UploadFileInfo: oneof}
	}
	return out
}
func ImportDataFileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImportDataFile) *krm.ImportDataFileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ImportDataFileObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Format
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UploadFileInfo = UploadFileInfoObservedState_FromProto(mapCtx, in.GetUploadFileInfo())
	return out
}
func ImportDataFileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ImportDataFileObservedState) *pb.ImportDataFile {
	if in == nil {
		return nil
	}
	out := &pb.ImportDataFile{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Format
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[pb.ImportDataFile_State](mapCtx, in.State)
	if oneof := UploadFileInfoObservedState_ToProto(mapCtx, in.UploadFileInfo); oneof != nil {
		out.FileInfo = &pb.ImportDataFile_UploadFileInfo{UploadFileInfo: oneof}
	}
	return out
}
func MigrationcenterImportDataFileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ImportDataFile) *krm.MigrationcenterImportDataFileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterImportDataFileObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Format
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: UploadFileInfo
	return out
}
func MigrationcenterImportDataFileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterImportDataFileObservedState) *pb.ImportDataFile {
	if in == nil {
		return nil
	}
	out := &pb.ImportDataFile{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Format
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: UploadFileInfo
	return out
}
func MigrationcenterImportDataFileSpec_FromProto(mapCtx *direct.MapContext, in *pb.ImportDataFile) *krm.MigrationcenterImportDataFileSpec {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterImportDataFileSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Format
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: UploadFileInfo
	return out
}
func MigrationcenterImportDataFileSpec_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterImportDataFileSpec) *pb.ImportDataFile {
	if in == nil {
		return nil
	}
	out := &pb.ImportDataFile{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Format
	// MISSING: CreateTime
	// MISSING: State
	// MISSING: UploadFileInfo
	return out
}
func UploadFileInfo_FromProto(mapCtx *direct.MapContext, in *pb.UploadFileInfo) *krm.UploadFileInfo {
	if in == nil {
		return nil
	}
	out := &krm.UploadFileInfo{}
	// MISSING: SignedURI
	// MISSING: Headers
	// MISSING: URIExpirationTime
	return out
}
func UploadFileInfo_ToProto(mapCtx *direct.MapContext, in *krm.UploadFileInfo) *pb.UploadFileInfo {
	if in == nil {
		return nil
	}
	out := &pb.UploadFileInfo{}
	// MISSING: SignedURI
	// MISSING: Headers
	// MISSING: URIExpirationTime
	return out
}
func UploadFileInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.UploadFileInfo) *krm.UploadFileInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.UploadFileInfoObservedState{}
	out.SignedURI = direct.LazyPtr(in.GetSignedUri())
	out.Headers = in.Headers
	out.URIExpirationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUriExpirationTime())
	return out
}
func UploadFileInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.UploadFileInfoObservedState) *pb.UploadFileInfo {
	if in == nil {
		return nil
	}
	out := &pb.UploadFileInfo{}
	out.SignedUri = direct.ValueOf(in.SignedURI)
	out.Headers = in.Headers
	out.UriExpirationTime = direct.StringTimestamp_ToProto(mapCtx, in.URIExpirationTime)
	return out
}
