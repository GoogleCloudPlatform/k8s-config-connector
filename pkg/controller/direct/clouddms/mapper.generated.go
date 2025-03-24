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

// +generated:mapper
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudDMSConversionWorkspaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.CloudDMSConversionWorkspaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSConversionWorkspaceObservedState{}
	out.HasUncommittedChanges = direct.LazyPtr(in.GetHasUncommittedChanges())
	out.LatestCommitID = direct.LazyPtr(in.GetLatestCommitId())
	out.LatestCommitTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLatestCommitTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func CloudDMSConversionWorkspaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSConversionWorkspaceObservedState) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	out.HasUncommittedChanges = direct.ValueOf(in.HasUncommittedChanges)
	out.LatestCommitId = direct.ValueOf(in.LatestCommitID)
	out.LatestCommitTime = direct.StringTimestamp_ToProto(mapCtx, in.LatestCommitTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func CloudDMSConversionWorkspaceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.CloudDMSConversionWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSConversionWorkspaceSpec{}
	out.Source = DatabaseEngineInfo_FromProto(mapCtx, in.GetSource())
	out.Destination = DatabaseEngineInfo_FromProto(mapCtx, in.GetDestination())
	out.GlobalSettings = in.GlobalSettings
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func CloudDMSConversionWorkspaceSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSConversionWorkspaceSpec) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	out.Source = DatabaseEngineInfo_ToProto(mapCtx, in.Source)
	out.Destination = DatabaseEngineInfo_ToProto(mapCtx, in.Destination)
	out.GlobalSettings = in.GlobalSettings
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func DatabaseEngineInfo_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseEngineInfo) *krm.DatabaseEngineInfo {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseEngineInfo{}
	out.Engine = direct.Enum_FromProto(mapCtx, in.GetEngine())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func DatabaseEngineInfo_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseEngineInfo) *pb.DatabaseEngineInfo {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseEngineInfo{}
	out.Engine = direct.Enum_ToProto[pb.DatabaseEngine](mapCtx, in.Engine)
	out.Version = direct.ValueOf(in.Version)
	return out
}
