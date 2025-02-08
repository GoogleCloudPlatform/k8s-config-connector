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

package monitoring

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func MonitoringNotificationChannelDescriptorObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannelDescriptor) *krm.MonitoringNotificationChannelDescriptorObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringNotificationChannelDescriptorObservedState{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: SupportedTiers
	// MISSING: LaunchStage
	return out
}
func MonitoringNotificationChannelDescriptorObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringNotificationChannelDescriptorObservedState) *pb.NotificationChannelDescriptor {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannelDescriptor{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: SupportedTiers
	// MISSING: LaunchStage
	return out
}
func MonitoringNotificationChannelDescriptorSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannelDescriptor) *krm.MonitoringNotificationChannelDescriptorSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringNotificationChannelDescriptorSpec{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: SupportedTiers
	// MISSING: LaunchStage
	return out
}
func MonitoringNotificationChannelDescriptorSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringNotificationChannelDescriptorSpec) *pb.NotificationChannelDescriptor {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannelDescriptor{}
	// MISSING: Name
	// MISSING: Type
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: SupportedTiers
	// MISSING: LaunchStage
	return out
}
func NotificationChannelDescriptor_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannelDescriptor) *krm.NotificationChannelDescriptor {
	if in == nil {
		return nil
	}
	out := &krm.NotificationChannelDescriptor{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.LazyPtr(in.GetType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = direct.Slice_FromProto(mapCtx, in.Labels, LabelDescriptor_FromProto)
	out.SupportedTiers = direct.EnumSlice_FromProto(mapCtx, in.SupportedTiers)
	out.LaunchStage = direct.Enum_FromProto(mapCtx, in.GetLaunchStage())
	return out
}
func NotificationChannelDescriptor_ToProto(mapCtx *direct.MapContext, in *krm.NotificationChannelDescriptor) *pb.NotificationChannelDescriptor {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannelDescriptor{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.ValueOf(in.Type)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Labels = direct.Slice_ToProto(mapCtx, in.Labels, LabelDescriptor_ToProto)
	out.SupportedTiers = direct.EnumSlice_ToProto[pb.ServiceTier](mapCtx, in.SupportedTiers)
	out.LaunchStage = direct.Enum_ToProto[pb.LaunchStage](mapCtx, in.LaunchStage)
	return out
}
