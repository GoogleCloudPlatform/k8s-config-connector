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

package advisorynotifications

import (
	pb "cloud.google.com/go/advisorynotifications/apiv1/advisorynotificationspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/advisorynotifications/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AdvisorynotificationsNotificationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Notification) *krm.AdvisorynotificationsNotificationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AdvisorynotificationsNotificationObservedState{}
	// MISSING: Name
	// MISSING: Subject
	// MISSING: Messages
	// MISSING: CreateTime
	// MISSING: NotificationType
	return out
}
func AdvisorynotificationsNotificationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AdvisorynotificationsNotificationObservedState) *pb.Notification {
	if in == nil {
		return nil
	}
	out := &pb.Notification{}
	// MISSING: Name
	// MISSING: Subject
	// MISSING: Messages
	// MISSING: CreateTime
	// MISSING: NotificationType
	return out
}
func AdvisorynotificationsNotificationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Notification) *krm.AdvisorynotificationsNotificationSpec {
	if in == nil {
		return nil
	}
	out := &krm.AdvisorynotificationsNotificationSpec{}
	// MISSING: Name
	// MISSING: Subject
	// MISSING: Messages
	// MISSING: CreateTime
	// MISSING: NotificationType
	return out
}
func AdvisorynotificationsNotificationSpec_ToProto(mapCtx *direct.MapContext, in *krm.AdvisorynotificationsNotificationSpec) *pb.Notification {
	if in == nil {
		return nil
	}
	out := &pb.Notification{}
	// MISSING: Name
	// MISSING: Subject
	// MISSING: Messages
	// MISSING: CreateTime
	// MISSING: NotificationType
	return out
}
func Attachment_FromProto(mapCtx *direct.MapContext, in *pb.Attachment) *krm.Attachment {
	if in == nil {
		return nil
	}
	out := &krm.Attachment{}
	out.Csv = Csv_FromProto(mapCtx, in.GetCsv())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func Attachment_ToProto(mapCtx *direct.MapContext, in *krm.Attachment) *pb.Attachment {
	if in == nil {
		return nil
	}
	out := &pb.Attachment{}
	if oneof := Csv_ToProto(mapCtx, in.Csv); oneof != nil {
		out.Data = &pb.Attachment_Csv{Csv: oneof}
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func Csv_FromProto(mapCtx *direct.MapContext, in *pb.Csv) *krm.Csv {
	if in == nil {
		return nil
	}
	out := &krm.Csv{}
	out.Headers = in.Headers
	out.DataRows = direct.Slice_FromProto(mapCtx, in.DataRows, Csv_CsvRow_FromProto)
	return out
}
func Csv_ToProto(mapCtx *direct.MapContext, in *krm.Csv) *pb.Csv {
	if in == nil {
		return nil
	}
	out := &pb.Csv{}
	out.Headers = in.Headers
	out.DataRows = direct.Slice_ToProto(mapCtx, in.DataRows, Csv_CsvRow_ToProto)
	return out
}
func Csv_CsvRow_FromProto(mapCtx *direct.MapContext, in *pb.Csv_CsvRow) *krm.Csv_CsvRow {
	if in == nil {
		return nil
	}
	out := &krm.Csv_CsvRow{}
	out.Entries = in.Entries
	return out
}
func Csv_CsvRow_ToProto(mapCtx *direct.MapContext, in *krm.Csv_CsvRow) *pb.Csv_CsvRow {
	if in == nil {
		return nil
	}
	out := &pb.Csv_CsvRow{}
	out.Entries = in.Entries
	return out
}
func Message_FromProto(mapCtx *direct.MapContext, in *pb.Message) *krm.Message {
	if in == nil {
		return nil
	}
	out := &krm.Message{}
	out.Body = Message_Body_FromProto(mapCtx, in.GetBody())
	out.Attachments = direct.Slice_FromProto(mapCtx, in.Attachments, Attachment_FromProto)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.LocalizationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLocalizationTime())
	return out
}
func Message_ToProto(mapCtx *direct.MapContext, in *krm.Message) *pb.Message {
	if in == nil {
		return nil
	}
	out := &pb.Message{}
	out.Body = Message_Body_ToProto(mapCtx, in.Body)
	out.Attachments = direct.Slice_ToProto(mapCtx, in.Attachments, Attachment_ToProto)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.LocalizationTime = direct.StringTimestamp_ToProto(mapCtx, in.LocalizationTime)
	return out
}
func Message_Body_FromProto(mapCtx *direct.MapContext, in *pb.Message_Body) *krm.Message_Body {
	if in == nil {
		return nil
	}
	out := &krm.Message_Body{}
	out.Text = Text_FromProto(mapCtx, in.GetText())
	return out
}
func Message_Body_ToProto(mapCtx *direct.MapContext, in *krm.Message_Body) *pb.Message_Body {
	if in == nil {
		return nil
	}
	out := &pb.Message_Body{}
	out.Text = Text_ToProto(mapCtx, in.Text)
	return out
}
func Notification_FromProto(mapCtx *direct.MapContext, in *pb.Notification) *krm.Notification {
	if in == nil {
		return nil
	}
	out := &krm.Notification{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Subject = Subject_FromProto(mapCtx, in.GetSubject())
	out.Messages = direct.Slice_FromProto(mapCtx, in.Messages, Message_FromProto)
	// MISSING: CreateTime
	out.NotificationType = direct.Enum_FromProto(mapCtx, in.GetNotificationType())
	return out
}
func Notification_ToProto(mapCtx *direct.MapContext, in *krm.Notification) *pb.Notification {
	if in == nil {
		return nil
	}
	out := &pb.Notification{}
	out.Name = direct.ValueOf(in.Name)
	out.Subject = Subject_ToProto(mapCtx, in.Subject)
	out.Messages = direct.Slice_ToProto(mapCtx, in.Messages, Message_ToProto)
	// MISSING: CreateTime
	out.NotificationType = direct.Enum_ToProto[pb.NotificationType](mapCtx, in.NotificationType)
	return out
}
func NotificationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Notification) *krm.NotificationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NotificationObservedState{}
	// MISSING: Name
	// MISSING: Subject
	// MISSING: Messages
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: NotificationType
	return out
}
func NotificationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NotificationObservedState) *pb.Notification {
	if in == nil {
		return nil
	}
	out := &pb.Notification{}
	// MISSING: Name
	// MISSING: Subject
	// MISSING: Messages
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: NotificationType
	return out
}
func Subject_FromProto(mapCtx *direct.MapContext, in *pb.Subject) *krm.Subject {
	if in == nil {
		return nil
	}
	out := &krm.Subject{}
	out.Text = Text_FromProto(mapCtx, in.GetText())
	return out
}
func Subject_ToProto(mapCtx *direct.MapContext, in *krm.Subject) *pb.Subject {
	if in == nil {
		return nil
	}
	out := &pb.Subject{}
	out.Text = Text_ToProto(mapCtx, in.Text)
	return out
}
func Text_FromProto(mapCtx *direct.MapContext, in *pb.Text) *krm.Text {
	if in == nil {
		return nil
	}
	out := &krm.Text{}
	out.EnText = direct.LazyPtr(in.GetEnText())
	out.LocalizedText = direct.LazyPtr(in.GetLocalizedText())
	out.LocalizationState = direct.Enum_FromProto(mapCtx, in.GetLocalizationState())
	return out
}
func Text_ToProto(mapCtx *direct.MapContext, in *krm.Text) *pb.Text {
	if in == nil {
		return nil
	}
	out := &pb.Text{}
	out.EnText = direct.ValueOf(in.EnText)
	out.LocalizedText = direct.ValueOf(in.LocalizedText)
	out.LocalizationState = direct.Enum_ToProto[pb.LocalizationState](mapCtx, in.LocalizationState)
	return out
}
