// Copyright 2024 Google LLC
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

package bigquerydatatransfer

import (
	pb "cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerydatatransfer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryDataTransferConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransferConfig) *krm.BigQueryDataTransferConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDataTransferConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DestinationDatasetID
	// MISSING: DisplayName
	// MISSING: DataSourceID
	// MISSING: Params
	// MISSING: Schedule
	// MISSING: ScheduleOptions
	// MISSING: DataRefreshWindowDays
	// MISSING: Disabled
	out.UpdateTime = TransferConfig_UpdateTime_FromProto(mapCtx, in.GetUpdateTime())
	out.NextRunTime = TransferConfig_NextRunTime_FromProto(mapCtx, in.GetNextRunTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: UserID
	out.DatasetRegion = direct.LazyPtr(in.GetDatasetRegion())
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	out.OwnerInfo = UserInfo_FromProto(mapCtx, in.GetOwnerInfo())
	// MISSING: EncryptionConfiguration
	return out
}
func BigQueryDataTransferConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDataTransferConfigObservedState) *pb.TransferConfig {
	if in == nil {
		return nil
	}
	out := &pb.TransferConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DestinationDatasetID
	// MISSING: DisplayName
	// MISSING: DataSourceID
	// MISSING: Params
	// MISSING: Schedule
	// MISSING: ScheduleOptions
	// MISSING: DataRefreshWindowDays
	// MISSING: Disabled
	out.UpdateTime = TransferConfig_UpdateTime_ToProto(mapCtx, in.UpdateTime)
	out.NextRunTime = TransferConfig_NextRunTime_ToProto(mapCtx, in.NextRunTime)
	out.State = direct.Enum_ToProto[pb.TransferState](mapCtx, in.State)
	// MISSING: UserID
	out.DatasetRegion = direct.ValueOf(in.DatasetRegion)
	// MISSING: NotificationPubsubTopic
	// MISSING: EmailPreferences
	if oneof := UserInfo_ToProto(mapCtx, in.OwnerInfo); oneof != nil {
		out.OwnerInfo = &pb.TransferConfig_OwnerInfo{OwnerInfo: oneof}
	}
	// MISSING: EncryptionConfiguration
	return out
}
func BigQueryDataTransferConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.TransferConfig) *krm.BigQueryDataTransferConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDataTransferConfigSpec{}
	// MISSING: Name
	// MISSING: DestinationDatasetID
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DataSourceID = direct.LazyPtr(in.GetDataSourceId())
	out.Params = map[string]string_FromProto(mapCtx, in.GetParams())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.ScheduleOptions = ScheduleOptions_FromProto(mapCtx, in.GetScheduleOptions())
	out.DataRefreshWindowDays = direct.LazyPtr(in.GetDataRefreshWindowDays())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	// MISSING: UpdateTime
	// MISSING: NextRunTime
	// MISSING: State
	out.UserID = direct.LazyPtr(in.GetUserId())
	// MISSING: DatasetRegion
	// MISSING: NotificationPubsubTopic
	out.EmailPreferences = EmailPreferences_FromProto(mapCtx, in.GetEmailPreferences())
	// MISSING: OwnerInfo
	out.EncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.GetEncryptionConfiguration())
	return out
}
func BigQueryDataTransferConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDataTransferConfigSpec) *pb.TransferConfig {
	if in == nil {
		return nil
	}
	out := &pb.TransferConfig{}
	// MISSING: Name
	// MISSING: DestinationDatasetID
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DataSourceId = direct.ValueOf(in.DataSourceID)
	out.Params = map[string]string_ToProto(mapCtx, in.Params)
	out.Schedule = direct.ValueOf(in.Schedule)
	out.ScheduleOptions = ScheduleOptions_ToProto(mapCtx, in.ScheduleOptions)
	out.DataRefreshWindowDays = direct.ValueOf(in.DataRefreshWindowDays)
	out.Disabled = direct.ValueOf(in.Disabled)
	// MISSING: UpdateTime
	// MISSING: NextRunTime
	// MISSING: State
	out.UserId = direct.ValueOf(in.UserID)
	// MISSING: DatasetRegion
	// MISSING: NotificationPubsubTopic
	out.EmailPreferences = EmailPreferences_ToProto(mapCtx, in.EmailPreferences)
	// MISSING: OwnerInfo
	out.EncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.EncryptionConfiguration)
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
	// MISSING: KmsKeyName
	return out
}
func EncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfiguration) *pb.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfiguration{}
	// MISSING: KmsKeyName
	return out
}
func ScheduleOptions_FromProto(mapCtx *direct.MapContext, in *pb.ScheduleOptions) *krm.ScheduleOptions {
	if in == nil {
		return nil
	}
	out := &krm.ScheduleOptions{}
	out.DisableAutoScheduling = direct.LazyPtr(in.GetDisableAutoScheduling())
	out.StartTime = ScheduleOptions_StartTime_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = ScheduleOptions_EndTime_FromProto(mapCtx, in.GetEndTime())
	return out
}
func ScheduleOptions_ToProto(mapCtx *direct.MapContext, in *krm.ScheduleOptions) *pb.ScheduleOptions {
	if in == nil {
		return nil
	}
	out := &pb.ScheduleOptions{}
	out.DisableAutoScheduling = direct.ValueOf(in.DisableAutoScheduling)
	out.StartTime = ScheduleOptions_StartTime_ToProto(mapCtx, in.StartTime)
	out.EndTime = ScheduleOptions_EndTime_ToProto(mapCtx, in.EndTime)
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
