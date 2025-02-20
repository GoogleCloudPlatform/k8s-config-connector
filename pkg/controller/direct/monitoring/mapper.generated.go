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
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func MonitoringNotificationChannelObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannel) *krm.MonitoringNotificationChannelObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringNotificationChannelObservedState{}
	// MISSING: Type
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: UserLabels
	// MISSING: VerificationStatus
	// MISSING: Enabled
	// MISSING: CreationRecord
	// MISSING: MutationRecords
	return out
}
func MonitoringNotificationChannelObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringNotificationChannelObservedState) *pb.NotificationChannel {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannel{}
	// MISSING: Type
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: UserLabels
	// MISSING: VerificationStatus
	// MISSING: Enabled
	// MISSING: CreationRecord
	// MISSING: MutationRecords
	return out
}
func MonitoringNotificationChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotificationChannel) *krm.MonitoringNotificationChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.MonitoringNotificationChannelSpec{}
	// MISSING: Type
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: UserLabels
	// MISSING: VerificationStatus
	// MISSING: Enabled
	// MISSING: CreationRecord
	// MISSING: MutationRecords
	return out
}
func MonitoringNotificationChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.MonitoringNotificationChannelSpec) *pb.NotificationChannel {
	if in == nil {
		return nil
	}
	out := &pb.NotificationChannel{}
	// MISSING: Type
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Labels
	// MISSING: UserLabels
	// MISSING: VerificationStatus
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
