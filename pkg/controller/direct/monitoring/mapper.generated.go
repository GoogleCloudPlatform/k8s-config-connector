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
// krm.group: monitoring.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.monitoring.v3

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MonitoringNotificationChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannel) *krm.MonitoringNotificationChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringNotificationChannelObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: UserLabels
	out.VerificationStatus = direct.Enum_FromProto(mapCtx, in.GetVerificationStatus())
	out.CreationRecord = MutationRecord_FromProto(mapCtx, in.GetCreationRecord())
	out.MutationRecords = direct.Slice_FromProto(mapCtx, in.MutationRecords, MutationRecord_FromProto)
	return out
}
func MonitoringNotificationChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringNotificationChannelObservedState) *pb.NotificationChannel {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannel{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: UserLabels
	out.VerificationStatus = direct.Enum_ToProto[pb.NotificationChannel_VerificationStatus](mapCtx, in.VerificationStatus)
	out.CreationRecord = MutationRecord_ToProto(mapCtx, in.CreationRecord)
	out.MutationRecords = direct.Slice_ToProto(mapCtx, in.MutationRecords, MutationRecord_ToProto)
	return out
}
func MonitoringNotificationChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannel) *krm.MonitoringNotificationChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringNotificationChannelSpec{}
	out.Type = direct.LazyPtr(in.GetType())
	// MISSING: Name
	// MISSING: DisplayName
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	// MISSING: UserLabels
	out.Enabled = direct.BoolValue_FromProto(mapCtx, in.GetEnabled())
	return out
}
func MonitoringNotificationChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringNotificationChannelSpec) *pb.NotificationChannel {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannel{}
	out.Type = direct.ValueOf(in.Type)
	// MISSING: Name
	// MISSING: DisplayName
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	// MISSING: UserLabels
	out.Enabled = direct.BoolValue_ToProto(mapCtx, in.Enabled)
	return out
}
func MonitoringNotificationChannelStatus_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannel) *krm.MonitoringNotificationChannelStatus {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringNotificationChannelStatus{}
	// MISSING: Type
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: UserLabels
	out.VerificationStatus = direct.Enum_FromProto(mapCtx, in.GetVerificationStatus())
	// MISSING: Enabled
	// MISSING: CreationRecord
	// MISSING: MutationRecords
	return out
}
func MonitoringNotificationChannelStatus_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringNotificationChannelStatus) *pb.NotificationChannel {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannel{}
	// MISSING: Type
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: UserLabels
	out.VerificationStatus = direct.Enum_ToProto[pb.NotificationChannel_VerificationStatus](mapCtx, in.VerificationStatus)
	// MISSING: Enabled
	// MISSING: CreationRecord
	// MISSING: MutationRecords
	return out
}
func MutationRecord_FromProto(mapCtx *direct.MapContext, in *pb.MutationRecord) *krm.MutationRecord {
	if in == nil {
		return nil
	}
	out := &krm.MutationRecord{}
	out.MutateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMutateTime())
	out.MutatedBy = direct.LazyPtr(in.GetMutatedBy())
	return out
}
func MutationRecord_ToProto(mapCtx *direct.MapContext, in *krm.MutationRecord) *pb.MutationRecord {
	if in == nil {
		return nil
	}
	out := &pb.MutationRecord{}
	out.MutateTime = direct.StringTimestamp_ToProto(mapCtx, in.MutateTime)
	out.MutatedBy = direct.ValueOf(in.MutatedBy)
	return out
}
