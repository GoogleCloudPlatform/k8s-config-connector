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

package dialogflow

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/apiv2beta1/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Document_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.Document {
	if in == nil {
		return nil
	}
	out := &krm.Document{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.MimeType = direct.LazyPtr(in.GetMimeType())
	out.KnowledgeTypes = direct.EnumSlice_FromProto(mapCtx, in.KnowledgeTypes)
	out.ContentURI = direct.LazyPtr(in.GetContentUri())
	out.Content = direct.LazyPtr(in.GetContent())
	out.RawContent = in.GetRawContent()
	out.EnableAutoReload = direct.LazyPtr(in.GetEnableAutoReload())
	// MISSING: LatestReloadStatus
	out.Metadata = in.Metadata
	// MISSING: State
	return out
}
func Document_ToProto(mapCtx *direct.MapContext, in *krm.Document) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.MimeType = direct.ValueOf(in.MimeType)
	out.KnowledgeTypes = direct.EnumSlice_ToProto[pb.Document_KnowledgeType](mapCtx, in.KnowledgeTypes)
	if oneof := Document_ContentUri_ToProto(mapCtx, in.ContentURI); oneof != nil {
		out.Source = oneof
	}
	if oneof := Document_Content_ToProto(mapCtx, in.Content); oneof != nil {
		out.Source = oneof
	}
	if oneof := Document_RawContent_ToProto(mapCtx, in.RawContent); oneof != nil {
		out.Source = oneof
	}
	out.EnableAutoReload = direct.ValueOf(in.EnableAutoReload)
	// MISSING: LatestReloadStatus
	out.Metadata = in.Metadata
	// MISSING: State
	return out
}
func DocumentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Document) *krm.DocumentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MimeType
	// MISSING: KnowledgeTypes
	// MISSING: ContentURI
	// MISSING: Content
	// MISSING: RawContent
	// MISSING: EnableAutoReload
	out.LatestReloadStatus = Document_ReloadStatus_FromProto(mapCtx, in.GetLatestReloadStatus())
	// MISSING: Metadata
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func DocumentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentObservedState) *pb.Document {
	if in == nil {
		return nil
	}
	out := &pb.Document{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: MimeType
	// MISSING: KnowledgeTypes
	// MISSING: ContentURI
	// MISSING: Content
	// MISSING: RawContent
	// MISSING: EnableAutoReload
	out.LatestReloadStatus = Document_ReloadStatus_ToProto(mapCtx, in.LatestReloadStatus)
	// MISSING: Metadata
	out.State = direct.Enum_ToProto[pb.Document_State](mapCtx, in.State)
	return out
}
func Document_ReloadStatus_FromProto(mapCtx *direct.MapContext, in *pb.Document_ReloadStatus) *krm.Document_ReloadStatus {
	if in == nil {
		return nil
	}
	out := &krm.Document_ReloadStatus{}
	out.Time = direct.StringTimestamp_FromProto(mapCtx, in.GetTime())
	out.Status = Status_FromProto(mapCtx, in.GetStatus())
	return out
}
func Document_ReloadStatus_ToProto(mapCtx *direct.MapContext, in *krm.Document_ReloadStatus) *pb.Document_ReloadStatus {
	if in == nil {
		return nil
	}
	out := &pb.Document_ReloadStatus{}
	out.Time = direct.StringTimestamp_ToProto(mapCtx, in.Time)
	out.Status = Status_ToProto(mapCtx, in.Status)
	return out
}
