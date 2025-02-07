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

package bigquery

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerydatatransfer/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
)
func BigQueryDataTransferConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransferConfig) *krm.BigQueryDataTransferConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDataTransferConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DestinationDatasetID
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.NextRunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextRunTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UserID = direct.LazyPtr(in.GetUserId())
	out.DatasetRegion = direct.LazyPtr(in.GetDatasetRegion())
	// MISSING: NotificationPubsubTopic
	out.OwnerInfo = UserInfo_FromProto(mapCtx, in.GetOwnerInfo())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func BigQueryDataTransferConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDataTransferConfigObservedState) *pb.TransferConfig {
	if in == nil {
		return nil
	}
	out := &pb.TransferConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DestinationDatasetID
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.NextRunTime = direct.StringTimestamp_ToProto(mapCtx, in.NextRunTime)
	out.State = direct.Enum_ToProto[pb.TransferState](mapCtx, in.State)
	out.UserId = direct.ValueOf(in.UserID)
	out.DatasetRegion = direct.ValueOf(in.DatasetRegion)
	// MISSING: NotificationPubsubTopic
	if oneof := UserInfo_ToProto(mapCtx, in.OwnerInfo); oneof != nil {
		out.OwnerInfo = &pb.TransferConfig_OwnerInfo{OwnerInfo: oneof}
	}
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func BigQueryDataTransferConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.TransferConfig) *krm.BigQueryDataTransferConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDataTransferConfigSpec{}
	// MISSING: DestinationDatasetID
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DataSourceID = direct.LazyPtr(in.GetDataSourceId())
	out.Params = Params_FromProto(mapCtx, in.GetParams())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.ScheduleOptions = ScheduleOptions_FromProto(mapCtx, in.GetScheduleOptions())
	out.ScheduleOptionsV2 = ScheduleOptionsV2_FromProto(mapCtx, in.GetScheduleOptionsV2())
	out.DataRefreshWindowDays = direct.LazyPtr(in.GetDataRefreshWindowDays())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	// MISSING: NotificationPubsubTopic
	out.EmailPreferences = EmailPreferences_FromProto(mapCtx, in.GetEmailPreferences())
	out.EncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.GetEncryptionConfiguration())
	return out
}
func BigQueryDataTransferConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDataTransferConfigSpec) *pb.TransferConfig {
	if in == nil {
		return nil
	}
	out := &pb.TransferConfig{}
	// MISSING: DestinationDatasetID
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DataSourceId = direct.ValueOf(in.DataSourceID)
	out.Params = Params_ToProto(mapCtx, in.Params)
	out.Schedule = direct.ValueOf(in.Schedule)
	out.ScheduleOptions = ScheduleOptions_ToProto(mapCtx, in.ScheduleOptions)
	out.ScheduleOptionsV2 = ScheduleOptionsV2_ToProto(mapCtx, in.ScheduleOptionsV2)
	out.DataRefreshWindowDays = direct.ValueOf(in.DataRefreshWindowDays)
	out.Disabled = direct.ValueOf(in.Disabled)
	// MISSING: NotificationPubsubTopic
	out.EmailPreferences = EmailPreferences_ToProto(mapCtx, in.EmailPreferences)
	out.EncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.EncryptionConfiguration)
	return out
}
func BigqueryDataSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.BigqueryDataSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryDataSourceObservedState{}
	// MISSING: Name
	// MISSING: DataSourceID
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ClientID
	// MISSING: Scopes
	// MISSING: TransferType
	// MISSING: SupportsMultipleTransfers
	// MISSING: UpdateDeadlineSeconds
	// MISSING: DefaultSchedule
	// MISSING: SupportsCustomSchedule
	// MISSING: Parameters
	// MISSING: HelpURL
	// MISSING: AuthorizationType
	// MISSING: DataRefreshType
	// MISSING: DefaultDataRefreshWindowDays
	// MISSING: ManualRunsDisabled
	// MISSING: MinimumScheduleInterval
	return out
}
func BigqueryDataSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryDataSourceObservedState) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	// MISSING: Name
	// MISSING: DataSourceID
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ClientID
	// MISSING: Scopes
	// MISSING: TransferType
	// MISSING: SupportsMultipleTransfers
	// MISSING: UpdateDeadlineSeconds
	// MISSING: DefaultSchedule
	// MISSING: SupportsCustomSchedule
	// MISSING: Parameters
	// MISSING: HelpURL
	// MISSING: AuthorizationType
	// MISSING: DataRefreshType
	// MISSING: DefaultDataRefreshWindowDays
	// MISSING: ManualRunsDisabled
	// MISSING: MinimumScheduleInterval
	return out
}
func BigqueryDataSourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.BigqueryDataSourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryDataSourceSpec{}
	// MISSING: Name
	// MISSING: DataSourceID
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ClientID
	// MISSING: Scopes
	// MISSING: TransferType
	// MISSING: SupportsMultipleTransfers
	// MISSING: UpdateDeadlineSeconds
	// MISSING: DefaultSchedule
	// MISSING: SupportsCustomSchedule
	// MISSING: Parameters
	// MISSING: HelpURL
	// MISSING: AuthorizationType
	// MISSING: DataRefreshType
	// MISSING: DefaultDataRefreshWindowDays
	// MISSING: ManualRunsDisabled
	// MISSING: MinimumScheduleInterval
	return out
}
func BigqueryDataSourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryDataSourceSpec) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	// MISSING: Name
	// MISSING: DataSourceID
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: ClientID
	// MISSING: Scopes
	// MISSING: TransferType
	// MISSING: SupportsMultipleTransfers
	// MISSING: UpdateDeadlineSeconds
	// MISSING: DefaultSchedule
	// MISSING: SupportsCustomSchedule
	// MISSING: Parameters
	// MISSING: HelpURL
	// MISSING: AuthorizationType
	// MISSING: DataRefreshType
	// MISSING: DefaultDataRefreshWindowDays
	// MISSING: ManualRunsDisabled
	// MISSING: MinimumScheduleInterval
	return out
}
func BigqueryTransferConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransferConfig) *krm.BigqueryTransferConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryTransferConfigObservedState{}
	// MISSING: Name
	// MISSING: DestinationDatasetID
	// MISSING: DisplayName
	// MISSING: DataSourceID
	// MISSING: Params
	// MISSING: Schedule
	// MISSING: ScheduleOptions
	// MISSING: ScheduleOptionsV2
	// MISSING: DataRefreshWindowDays
	// MISSING: Disabled
	// MISSING: UpdateTime
	// MISSING: NextRunTime
	// MISSING: State
	// MISSING: UserID
	// MISSING: DatasetRegion
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	// MISSING: OwnerInfo
	// MISSING: EncryptionConfiguration
	// MISSING: Error
	return out
}
func BigqueryTransferConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryTransferConfigObservedState) *pb.TransferConfig {
	if in == nil {
		return nil
	}
	out := &pb.TransferConfig{}
	// MISSING: Name
	// MISSING: DestinationDatasetID
	// MISSING: DisplayName
	// MISSING: DataSourceID
	// MISSING: Params
	// MISSING: Schedule
	// MISSING: ScheduleOptions
	// MISSING: ScheduleOptionsV2
	// MISSING: DataRefreshWindowDays
	// MISSING: Disabled
	// MISSING: UpdateTime
	// MISSING: NextRunTime
	// MISSING: State
	// MISSING: UserID
	// MISSING: DatasetRegion
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	// MISSING: OwnerInfo
	// MISSING: EncryptionConfiguration
	// MISSING: Error
	return out
}
func BigqueryTransferConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.TransferConfig) *krm.BigqueryTransferConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryTransferConfigSpec{}
	// MISSING: Name
	// MISSING: DestinationDatasetID
	// MISSING: DisplayName
	// MISSING: DataSourceID
	// MISSING: Params
	// MISSING: Schedule
	// MISSING: ScheduleOptions
	// MISSING: ScheduleOptionsV2
	// MISSING: DataRefreshWindowDays
	// MISSING: Disabled
	// MISSING: UpdateTime
	// MISSING: NextRunTime
	// MISSING: State
	// MISSING: UserID
	// MISSING: DatasetRegion
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	// MISSING: OwnerInfo
	// MISSING: EncryptionConfiguration
	// MISSING: Error
	return out
}
func BigqueryTransferConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryTransferConfigSpec) *pb.TransferConfig {
	if in == nil {
		return nil
	}
	out := &pb.TransferConfig{}
	// MISSING: Name
	// MISSING: DestinationDatasetID
	// MISSING: DisplayName
	// MISSING: DataSourceID
	// MISSING: Params
	// MISSING: Schedule
	// MISSING: ScheduleOptions
	// MISSING: ScheduleOptionsV2
	// MISSING: DataRefreshWindowDays
	// MISSING: Disabled
	// MISSING: UpdateTime
	// MISSING: NextRunTime
	// MISSING: State
	// MISSING: UserID
	// MISSING: DatasetRegion
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	// MISSING: OwnerInfo
	// MISSING: EncryptionConfiguration
	// MISSING: Error
	return out
}
func BigqueryTransferRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransferRun) *krm.BigqueryTransferRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryTransferRunObservedState{}
	// MISSING: Name
	// MISSING: ScheduleTime
	// MISSING: RunTime
	// MISSING: ErrorStatus
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Params
	// MISSING: DestinationDatasetID
	// MISSING: DataSourceID
	// MISSING: State
	// MISSING: UserID
	// MISSING: Schedule
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	return out
}
func BigqueryTransferRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryTransferRunObservedState) *pb.TransferRun {
	if in == nil {
		return nil
	}
	out := &pb.TransferRun{}
	// MISSING: Name
	// MISSING: ScheduleTime
	// MISSING: RunTime
	// MISSING: ErrorStatus
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Params
	// MISSING: DestinationDatasetID
	// MISSING: DataSourceID
	// MISSING: State
	// MISSING: UserID
	// MISSING: Schedule
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	return out
}
func BigqueryTransferRunSpec_FromProto(mapCtx *direct.MapContext, in *pb.TransferRun) *krm.BigqueryTransferRunSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryTransferRunSpec{}
	// MISSING: Name
	// MISSING: ScheduleTime
	// MISSING: RunTime
	// MISSING: ErrorStatus
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Params
	// MISSING: DestinationDatasetID
	// MISSING: DataSourceID
	// MISSING: State
	// MISSING: UserID
	// MISSING: Schedule
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	return out
}
func BigqueryTransferRunSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryTransferRunSpec) *pb.TransferRun {
	if in == nil {
		return nil
	}
	out := &pb.TransferRun{}
	// MISSING: Name
	// MISSING: ScheduleTime
	// MISSING: RunTime
	// MISSING: ErrorStatus
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Params
	// MISSING: DestinationDatasetID
	// MISSING: DataSourceID
	// MISSING: State
	// MISSING: UserID
	// MISSING: Schedule
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	return out
}
func EmailPreferences_FromProto(mapCtx *direct.MapContext, in *pb.EmailPreferences) *krm.EmailPreferences {
	if in == nil {
		return nil
	}
	out := &krm.EmailPreferences{}
	out.EnableFailureEmail = direct.LazyPtr(in.GetEnableFailureEmail())
	return out
}
func EmailPreferences_ToProto(mapCtx *direct.MapContext, in *krm.EmailPreferences) *pb.EmailPreferences {
	if in == nil {
		return nil
	}
	out := &pb.EmailPreferences{}
	out.EnableFailureEmail = direct.ValueOf(in.EnableFailureEmail)
	return out
}
func EncryptionConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfiguration) *krm.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfiguration{}
	// MISSING: KMSKeyName
	return out
}
func EncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfiguration) *pb.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfiguration{}
	// MISSING: KMSKeyName
	return out
}
func EventDrivenSchedule_FromProto(mapCtx *direct.MapContext, in *pb.EventDrivenSchedule) *krm.EventDrivenSchedule {
	if in == nil {
		return nil
	}
	out := &krm.EventDrivenSchedule{}
	// MISSING: PubsubSubscription
	return out
}
func EventDrivenSchedule_ToProto(mapCtx *direct.MapContext, in *krm.EventDrivenSchedule) *pb.EventDrivenSchedule {
	if in == nil {
		return nil
	}
	out := &pb.EventDrivenSchedule{}
	// MISSING: PubsubSubscription
	return out
}
func ManualSchedule_FromProto(mapCtx *direct.MapContext, in *pb.ManualSchedule) *krm.ManualSchedule {
	if in == nil {
		return nil
	}
	out := &krm.ManualSchedule{}
	return out
}
func ManualSchedule_ToProto(mapCtx *direct.MapContext, in *krm.ManualSchedule) *pb.ManualSchedule {
	if in == nil {
		return nil
	}
	out := &pb.ManualSchedule{}
	return out
}
func ScheduleOptions_FromProto(mapCtx *direct.MapContext, in *pb.ScheduleOptions) *krm.ScheduleOptions {
	if in == nil {
		return nil
	}
	out := &krm.ScheduleOptions{}
	out.DisableAutoScheduling = direct.LazyPtr(in.GetDisableAutoScheduling())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func ScheduleOptions_ToProto(mapCtx *direct.MapContext, in *krm.ScheduleOptions) *pb.ScheduleOptions {
	if in == nil {
		return nil
	}
	out := &pb.ScheduleOptions{}
	out.DisableAutoScheduling = direct.ValueOf(in.DisableAutoScheduling)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func ScheduleOptionsV2_FromProto(mapCtx *direct.MapContext, in *pb.ScheduleOptionsV2) *krm.ScheduleOptionsV2 {
	if in == nil {
		return nil
	}
	out := &krm.ScheduleOptionsV2{}
	out.TimeBasedSchedule = TimeBasedSchedule_FromProto(mapCtx, in.GetTimeBasedSchedule())
	out.ManualSchedule = ManualSchedule_FromProto(mapCtx, in.GetManualSchedule())
	out.EventDrivenSchedule = EventDrivenSchedule_FromProto(mapCtx, in.GetEventDrivenSchedule())
	return out
}
func ScheduleOptionsV2_ToProto(mapCtx *direct.MapContext, in *krm.ScheduleOptionsV2) *pb.ScheduleOptionsV2 {
	if in == nil {
		return nil
	}
	out := &pb.ScheduleOptionsV2{}
	if oneof := TimeBasedSchedule_ToProto(mapCtx, in.TimeBasedSchedule); oneof != nil {
		out.Schedule = &pb.ScheduleOptionsV2_TimeBasedSchedule{TimeBasedSchedule: oneof}
	}
	if oneof := ManualSchedule_ToProto(mapCtx, in.ManualSchedule); oneof != nil {
		out.Schedule = &pb.ScheduleOptionsV2_ManualSchedule{ManualSchedule: oneof}
	}
	if oneof := EventDrivenSchedule_ToProto(mapCtx, in.EventDrivenSchedule); oneof != nil {
		out.Schedule = &pb.ScheduleOptionsV2_EventDrivenSchedule{EventDrivenSchedule: oneof}
	}
	return out
}
func TimeBasedSchedule_FromProto(mapCtx *direct.MapContext, in *pb.TimeBasedSchedule) *krm.TimeBasedSchedule {
	if in == nil {
		return nil
	}
	out := &krm.TimeBasedSchedule{}
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func TimeBasedSchedule_ToProto(mapCtx *direct.MapContext, in *krm.TimeBasedSchedule) *pb.TimeBasedSchedule {
	if in == nil {
		return nil
	}
	out := &pb.TimeBasedSchedule{}
	out.Schedule = direct.ValueOf(in.Schedule)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func TransferRun_FromProto(mapCtx *direct.MapContext, in *pb.TransferRun) *krm.TransferRun {
	if in == nil {
		return nil
	}
	out := &krm.TransferRun{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ScheduleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduleTime())
	out.RunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRunTime())
	out.ErrorStatus = Status_FromProto(mapCtx, in.GetErrorStatus())
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Params
	// MISSING: DestinationDatasetID
	// MISSING: DataSourceID
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.UserID = direct.LazyPtr(in.GetUserId())
	// MISSING: Schedule
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	return out
}
func TransferRun_ToProto(mapCtx *direct.MapContext, in *krm.TransferRun) *pb.TransferRun {
	if in == nil {
		return nil
	}
	out := &pb.TransferRun{}
	out.Name = direct.ValueOf(in.Name)
	out.ScheduleTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduleTime)
	out.RunTime = direct.StringTimestamp_ToProto(mapCtx, in.RunTime)
	out.ErrorStatus = Status_ToProto(mapCtx, in.ErrorStatus)
	// MISSING: StartTime
	// MISSING: EndTime
	// MISSING: UpdateTime
	// MISSING: Params
	// MISSING: DestinationDatasetID
	// MISSING: DataSourceID
	out.State = direct.Enum_ToProto[pb.TransferState](mapCtx, in.State)
	out.UserId = direct.ValueOf(in.UserID)
	// MISSING: Schedule
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	return out
}
func TransferRunObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransferRun) *krm.TransferRunObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TransferRunObservedState{}
	// MISSING: Name
	// MISSING: ScheduleTime
	// MISSING: RunTime
	// MISSING: ErrorStatus
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Params = Params_FromProto(mapCtx, in.GetParams())
	out.DestinationDatasetID = direct.LazyPtr(in.GetDestinationDatasetId())
	out.DataSourceID = direct.LazyPtr(in.GetDataSourceId())
	// MISSING: State
	// MISSING: UserID
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.NotificationPubsubTopic = direct.LazyPtr(in.GetNotificationPubsubTopic())
	out.EmailPreferences = EmailPreferences_FromProto(mapCtx, in.GetEmailPreferences())
	return out
}
func TransferRunObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TransferRunObservedState) *pb.TransferRun {
	if in == nil {
		return nil
	}
	out := &pb.TransferRun{}
	// MISSING: Name
	// MISSING: ScheduleTime
	// MISSING: RunTime
	// MISSING: ErrorStatus
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Params = Params_ToProto(mapCtx, in.Params)
	if oneof := TransferRunObservedState_DestinationDatasetId_ToProto(mapCtx, in.DestinationDatasetID); oneof != nil {
		out.Destination = oneof
	}
	out.DataSourceId = direct.ValueOf(in.DataSourceID)
	// MISSING: State
	// MISSING: UserID
	out.Schedule = direct.ValueOf(in.Schedule)
	out.NotificationPubsubTopic = direct.ValueOf(in.NotificationPubsubTopic)
	out.EmailPreferences = EmailPreferences_ToProto(mapCtx, in.EmailPreferences)
	return out
}
func UserInfo_FromProto(mapCtx *direct.MapContext, in *pb.UserInfo) *krm.UserInfo {
	if in == nil {
		return nil
	}
	out := &krm.UserInfo{}
	out.Email = in.Email
	return out
}
func UserInfo_ToProto(mapCtx *direct.MapContext, in *krm.UserInfo) *pb.UserInfo {
	if in == nil {
		return nil
	}
	out := &pb.UserInfo{}
	out.Email = in.Email
	return out
}
