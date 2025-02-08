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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
)
func DialogflowVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.DialogflowVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowVersionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: NluSettings
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func DialogflowVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowVersionObservedState) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: NluSettings
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func DialogflowVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.DialogflowVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.DialogflowVersionSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: NluSettings
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func DialogflowVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.DialogflowVersionSpec) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: NluSettings
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func NluSettings_FromProto(mapCtx *direct.MapContext, in *pb.NluSettings) *krm.NluSettings {
	if in == nil {
		return nil
	}
	out := &krm.NluSettings{}
	out.ModelType = direct.Enum_FromProto(mapCtx, in.GetModelType())
	out.ClassificationThreshold = direct.LazyPtr(in.GetClassificationThreshold())
	out.ModelTrainingMode = direct.Enum_FromProto(mapCtx, in.GetModelTrainingMode())
	return out
}
func NluSettings_ToProto(mapCtx *direct.MapContext, in *krm.NluSettings) *pb.NluSettings {
	if in == nil {
		return nil
	}
	out := &pb.NluSettings{}
	out.ModelType = direct.Enum_ToProto[pb.NluSettings_ModelType](mapCtx, in.ModelType)
	out.ClassificationThreshold = direct.ValueOf(in.ClassificationThreshold)
	out.ModelTrainingMode = direct.Enum_ToProto[pb.NluSettings_ModelTrainingMode](mapCtx, in.ModelTrainingMode)
	return out
}
func Version_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.Version {
	if in == nil {
		return nil
	}
	out := &krm.Version{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: NluSettings
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func Version_ToProto(mapCtx *direct.MapContext, in *krm.Version) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: NluSettings
	// MISSING: CreateTime
	// MISSING: State
	return out
}
func VersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Version) *krm.VersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VersionObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.NluSettings = NluSettings_FromProto(mapCtx, in.GetNluSettings())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func VersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VersionObservedState) *pb.Version {
	if in == nil {
		return nil
	}
	out := &pb.Version{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	out.NluSettings = NluSettings_ToProto(mapCtx, in.NluSettings)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.State = direct.Enum_ToProto[pb.Version_State](mapCtx, in.State)
	return out
}
