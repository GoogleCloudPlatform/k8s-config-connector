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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/advisorynotifications/apiv1/advisorynotificationspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/advisorynotifications/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
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
func AdvisorynotificationsSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.AdvisorynotificationsSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AdvisorynotificationsSettingsObservedState{}
	// MISSING: Name
	// MISSING: NotificationSettings
	// MISSING: Etag
	return out
}
func AdvisorynotificationsSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AdvisorynotificationsSettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: NotificationSettings
	// MISSING: Etag
	return out
}
func AdvisorynotificationsSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.AdvisorynotificationsSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.AdvisorynotificationsSettingsSpec{}
	// MISSING: Name
	// MISSING: NotificationSettings
	// MISSING: Etag
	return out
}
func AdvisorynotificationsSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.AdvisorynotificationsSettingsSpec) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: NotificationSettings
	// MISSING: Etag
	return out
}
func NotificationSettings_FromProto(mapCtx *direct.MapContext, in *pb.NotificationSettings) *krm.NotificationSettings {
	if in == nil {
		return nil
	}
	out := &krm.NotificationSettings{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func NotificationSettings_ToProto(mapCtx *direct.MapContext, in *krm.NotificationSettings) *pb.NotificationSettings {
	if in == nil {
		return nil
	}
	out := &pb.NotificationSettings{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func Settings_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.Settings {
	if in == nil {
		return nil
	}
	out := &krm.Settings{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: NotificationSettings
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func Settings_ToProto(mapCtx *direct.MapContext, in *krm.Settings) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: NotificationSettings
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
