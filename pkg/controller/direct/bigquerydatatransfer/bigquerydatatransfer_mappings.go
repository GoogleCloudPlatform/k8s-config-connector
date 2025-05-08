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
	bigquery "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerydatatransfer/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func BigQueryDataTransferConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TransferConfig) *krm.BigQueryDataTransferConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDataTransferConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.NextRunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetNextRunTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.DatasetRegion = direct.LazyPtr(in.GetDatasetRegion())
	out.OwnerInfo = UserInfo_FromProto(mapCtx, in.GetOwnerInfo())
	out.UserID = direct.LazyPtr(in.GetUserId())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	return out
}
func BigQueryDataTransferConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDataTransferConfigObservedState) *pb.TransferConfig {
	if in == nil {
		return nil
	}
	out := &pb.TransferConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.NextRunTime = direct.StringTimestamp_ToProto(mapCtx, in.NextRunTime)
	out.State = direct.Enum_ToProto[pb.TransferState](mapCtx, in.State)
	out.DatasetRegion = direct.ValueOf(in.DatasetRegion)
	out.OwnerInfo = UserInfo_ToProto(mapCtx, in.OwnerInfo)
	out.UserId = direct.ValueOf(in.UserID)
	out.Error = Status_ToProto(mapCtx, in.Error)
	return out
}
func BigQueryDataTransferConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.TransferConfig) *krm.BigQueryDataTransferConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryDataTransferConfigSpec{}
	if in.GetDestinationDatasetId() != "" {
		out.DatasetRef = &bigquery.DatasetRef{External: in.GetDestinationDatasetId()}
	}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.DataSourceID = direct.LazyPtr(in.GetDataSourceId())
	out.Params = Params_FromProto(mapCtx, in.GetParams())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.ScheduleOptions = ScheduleOptions_FromProto(mapCtx, in.GetScheduleOptions())
	out.DataRefreshWindowDays = direct.LazyPtr(in.GetDataRefreshWindowDays())
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	if in.GetNotificationPubsubTopic() != "" {
		out.PubSubTopicRef = &refv1beta1.PubSubTopicRef{External: in.GetNotificationPubsubTopic()}
	}
	out.EmailPreferences = EmailPreferences_FromProto(mapCtx, in.GetEmailPreferences())
	out.EncryptionConfiguration = EncryptionConfiguration_FromProto(mapCtx, in.GetEncryptionConfiguration())
	out.ScheduleOptionsV2 = ScheduleOptionsV2_FromProto(mapCtx, in.GetScheduleOptionsV2())
	return out
}
func BigQueryDataTransferConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryDataTransferConfigSpec) *pb.TransferConfig {
	if in == nil {
		return nil
	}
	out := &pb.TransferConfig{}
	if in.DatasetRef != nil {
		out.Destination = &pb.TransferConfig_DestinationDatasetId{DestinationDatasetId: in.DatasetRef.External}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.DataSourceId = direct.ValueOf(in.DataSourceID)
	out.Params = Params_ToProto(mapCtx, in.Params)
	out.Schedule = direct.ValueOf(in.Schedule)
	out.ScheduleOptions = ScheduleOptions_ToProto(mapCtx, in.ScheduleOptions)
	out.DataRefreshWindowDays = direct.ValueOf(in.DataRefreshWindowDays)
	out.Disabled = direct.ValueOf(in.Disabled)
	if in.PubSubTopicRef != nil {
		out.NotificationPubsubTopic = in.PubSubTopicRef.External
	}
	out.EmailPreferences = EmailPreferences_ToProto(mapCtx, in.EmailPreferences)
	out.EncryptionConfiguration = EncryptionConfiguration_ToProto(mapCtx, in.EncryptionConfiguration)
	out.ScheduleOptionsV2 = ScheduleOptionsV2_ToProto(mapCtx, in.ScheduleOptionsV2)
	return out
}
func EncryptionConfiguration_FromProto(mapCtx *direct.MapContext, in *pb.EncryptionConfiguration) *krm.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &krm.EncryptionConfiguration{}
	if in.GetKmsKeyName() != nil {
		out.KMSKeyRef = &kmsv1beta1.KMSKeyRef_OneOf{External: in.GetKmsKeyName().GetValue()}
	}
	return out
}
func EncryptionConfiguration_ToProto(mapCtx *direct.MapContext, in *krm.EncryptionConfiguration) *pb.EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := &pb.EncryptionConfiguration{}
	if in.KMSKeyRef != nil {
		out.KmsKeyName = wrapperspb.String(in.KMSKeyRef.External)
	}
	return out
}
func Params_FromProto(mapCtx *direct.MapContext, in *structpb.Struct) map[string]string {
	if in == nil {
		return nil
	}
	out := make(map[string]string)
	for k, v := range in.GetFields() {
		// TODO: if we need to support more types, we need to change KRM type of Params to map[string]interface{}
		if stringValue, ok := v.Kind.(*structpb.Value_StringValue); ok {
			out[k] = stringValue.StringValue
		} else {
			mapCtx.Errorf("unexpected non-string value for key %s", k)
		}
	}
	return out
}
func Params_ToProto(mapCtx *direct.MapContext, in map[string]string) *structpb.Struct {
	if in == nil {
		return nil
	}
	out := &structpb.Struct{
		Fields: map[string]*structpb.Value{},
	}
	for k, v := range in {
		// TODO: if we need to support more types, we need to change KRM type of Params to map[string]interface{}
		out.Fields[k] = structpb.NewStringValue(v)
	}
	return out
}
func Status_ToProto(mapCtx *direct.MapContext, in *krm.Status) *statuspb.Status {
	if in == nil {
		return nil
	}
	out := &statuspb.Status{}
	out.Code = direct.ValueOf(in.Code)
	out.Message = direct.ValueOf(in.Message)
	// NOTYET
	// out.Details
	return out
}
func Status_FromProto(mapCtx *direct.MapContext, in *statuspb.Status) *krm.Status {
	if in == nil {
		return nil
	}
	out := &krm.Status{}
	out.Code = direct.LazyPtr(in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	// NOTYET
	// out.Details
	return out
}
func EventDrivenSchedule_FromProto(mapCtx *direct.MapContext, in *pb.EventDrivenSchedule) *krm.EventDrivenSchedule {
	if in == nil {
		return nil
	}
	out := &krm.EventDrivenSchedule{}
	if in.GetPubsubSubscription() != "" {
		out.PubSubSubscriptionRef = &refv1beta1.PubSubSubscriptionRef{External: in.GetPubsubSubscription()}
	}
	return out
}
func EventDrivenSchedule_ToProto(mapCtx *direct.MapContext, in *krm.EventDrivenSchedule) *pb.EventDrivenSchedule {
	if in == nil {
		return nil
	}
	out := &pb.EventDrivenSchedule{}
	if in.PubSubSubscriptionRef != nil {
		out.PubsubSubscription = in.PubSubSubscriptionRef.External
	}
	return out
}
