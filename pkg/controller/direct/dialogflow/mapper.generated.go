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
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ConversationDataset_FromProto(mapCtx *direct.MapContext, in *pb.ConversationDataset) *krm.ConversationDataset {
	if in == nil {
		return nil
	}
	out := &krm.ConversationDataset{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: CreateTime
	// MISSING: InputConfig
	// MISSING: ConversationInfo
	// MISSING: ConversationCount
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	return out
}
func ConversationDataset_ToProto(mapCtx *direct.MapContext, in *krm.ConversationDataset) *pb.ConversationDataset {
	if in == nil {
		return nil
	}
	out := &pb.ConversationDataset{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: CreateTime
	// MISSING: InputConfig
	// MISSING: ConversationInfo
	// MISSING: ConversationCount
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	return out
}
func ConversationDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationDataset) *krm.ConversationDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversationDatasetObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.InputConfig = InputConfig_FromProto(mapCtx, in.GetInputConfig())
	out.ConversationInfo = ConversationInfo_FromProto(mapCtx, in.GetConversationInfo())
	out.ConversationCount = direct.LazyPtr(in.GetConversationCount())
	out.SatisfiesPzi = in.SatisfiesPzi
	out.SatisfiesPzs = in.SatisfiesPzs
	return out
}
func ConversationDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversationDatasetObservedState) *pb.ConversationDataset {
	if in == nil {
		return nil
	}
	out := &pb.ConversationDataset{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.InputConfig = InputConfig_ToProto(mapCtx, in.InputConfig)
	out.ConversationInfo = ConversationInfo_ToProto(mapCtx, in.ConversationInfo)
	out.ConversationCount = direct.ValueOf(in.ConversationCount)
	out.SatisfiesPzi = in.SatisfiesPzi
	out.SatisfiesPzs = in.SatisfiesPzs
	return out
}
func ConversationInfo_FromProto(mapCtx *direct.MapContext, in *pb.ConversationInfo) *krm.ConversationInfo {
	if in == nil {
		return nil
	}
	out := &krm.ConversationInfo{}
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	return out
}
func ConversationInfo_ToProto(mapCtx *direct.MapContext, in *krm.ConversationInfo) *pb.ConversationInfo {
	if in == nil {
		return nil
	}
	out := &pb.ConversationInfo{}
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	return out
}
func DialogflowConversationDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversationDataset) *krm.DialogflowConversationDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationDatasetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: InputConfig
	// MISSING: ConversationInfo
	// MISSING: ConversationCount
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	return out
}
func DialogflowConversationDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationDatasetObservedState) *pb.ConversationDataset {
	if in == nil {
		return nil
	}
	out := &pb.ConversationDataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: InputConfig
	// MISSING: ConversationInfo
	// MISSING: ConversationCount
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	return out
}
func DialogflowConversationDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversationDataset) *krm.DialogflowConversationDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowConversationDatasetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: InputConfig
	// MISSING: ConversationInfo
	// MISSING: ConversationCount
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	return out
}
func DialogflowConversationDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowConversationDatasetSpec) *pb.ConversationDataset {
	if in == nil {
		return nil
	}
	out := &pb.ConversationDataset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: InputConfig
	// MISSING: ConversationInfo
	// MISSING: ConversationCount
	// MISSING: SatisfiesPzi
	// MISSING: SatisfiesPzs
	return out
}
func GcsSources_FromProto(mapCtx *direct.MapContext, in *pb.GcsSources) *krm.GcsSources {
	if in == nil {
		return nil
	}
	out := &krm.GcsSources{}
	out.Uris = in.Uris
	return out
}
func GcsSources_ToProto(mapCtx *direct.MapContext, in *krm.GcsSources) *pb.GcsSources {
	if in == nil {
		return nil
	}
	out := &pb.GcsSources{}
	out.Uris = in.Uris
	return out
}
func InputConfig_FromProto(mapCtx *direct.MapContext, in *pb.InputConfig) *krm.InputConfig {
	if in == nil {
		return nil
	}
	out := &krm.InputConfig{}
	out.GcsSource = GcsSources_FromProto(mapCtx, in.GetGcsSource())
	return out
}
func InputConfig_ToProto(mapCtx *direct.MapContext, in *krm.InputConfig) *pb.InputConfig {
	if in == nil {
		return nil
	}
	out := &pb.InputConfig{}
	if oneof := GcsSources_ToProto(mapCtx, in.GcsSource); oneof != nil {
		out.Source = &pb.InputConfig_GcsSource{GcsSource: oneof}
	}
	return out
}
