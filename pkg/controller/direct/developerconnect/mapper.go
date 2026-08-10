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

package developerconnect

import (
	pb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	krmapphubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apphub/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/developerconnect/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DevConnectInsightsConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.InsightsConfig) *krm.DevConnectInsightsConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DevConnectInsightsConfigObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.RuntimeConfigs = direct.Slice_FromProto(mapCtx, in.RuntimeConfigs, RuntimeConfigObservedState_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Annotations
	// MISSING: Labels
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	out.Errors = direct.Slice_FromProto(mapCtx, in.Errors, direct.Status_FromProto)
	return out
}

func DevConnectInsightsConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DevConnectInsightsConfigObservedState) *pb.InsightsConfig {
	if in == nil {
		return nil
	}
	out := &pb.InsightsConfig{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.RuntimeConfigs = direct.Slice_ToProto(mapCtx, in.RuntimeConfigs, RuntimeConfigObservedState_ToProto)
	out.State = direct.Enum_ToProto[pb.InsightsConfig_State](mapCtx, in.State)
	// MISSING: Annotations
	// MISSING: Labels
	out.Reconciling = direct.ValueOf(in.Reconciling)
	out.Errors = direct.Slice_ToProto(mapCtx, in.Errors, direct.Status_ToProto)
	return out
}

func DevConnectInsightsConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.InsightsConfig) *krm.DevConnectInsightsConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.DevConnectInsightsConfigSpec{}
	if in.GetAppHubApplication() != "" {
		out.AppHubApplicationRef = &krmapphubv1beta1.ApplicationRef{External: in.GetAppHubApplication()}
	}
	// MISSING: Name
	out.ArtifactConfigs = direct.Slice_FromProto(mapCtx, in.ArtifactConfigs, ArtifactConfig_FromProto)
	// MISSING: Annotations
	// MISSING: Labels
	return out
}

func DevConnectInsightsConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.DevConnectInsightsConfigSpec) *pb.InsightsConfig {
	if in == nil {
		return nil
	}
	out := &pb.InsightsConfig{}
	if in.AppHubApplicationRef != nil && in.AppHubApplicationRef.External != "" {
		out.InsightsConfigContext = &pb.InsightsConfig_AppHubApplication{
			AppHubApplication: in.AppHubApplicationRef.External,
		}
	}
	// MISSING: Name
	out.ArtifactConfigs = direct.Slice_ToProto(mapCtx, in.ArtifactConfigs, ArtifactConfig_ToProto)
	// MISSING: Annotations
	// MISSING: Labels
	return out
}
