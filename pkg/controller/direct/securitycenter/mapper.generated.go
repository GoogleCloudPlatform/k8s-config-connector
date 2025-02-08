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

package securitycenter

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func NotificationConfig_FromProto(mapCtx *direct.MapContext, in *pb.NotificationConfig) *krm.NotificationConfig {
	if in == nil {
		return nil
	}
	out := &krm.NotificationConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PubsubTopic = direct.LazyPtr(in.GetPubsubTopic())
	// MISSING: ServiceAccount
	out.StreamingConfig = NotificationConfig_StreamingConfig_FromProto(mapCtx, in.GetStreamingConfig())
	return out
}
func NotificationConfig_ToProto(mapCtx *direct.MapContext, in *krm.NotificationConfig) *pb.NotificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotificationConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.PubsubTopic = direct.ValueOf(in.PubsubTopic)
	// MISSING: ServiceAccount
	if oneof := NotificationConfig_StreamingConfig_ToProto(mapCtx, in.StreamingConfig); oneof != nil {
		out.NotifyConfig = &pb.NotificationConfig_StreamingConfig_{StreamingConfig: oneof}
	}
	return out
}
func NotificationConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotificationConfig) *krm.NotificationConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NotificationConfigObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTopic
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	// MISSING: StreamingConfig
	return out
}
func NotificationConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NotificationConfigObservedState) *pb.NotificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotificationConfig{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTopic
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	// MISSING: StreamingConfig
	return out
}
func NotificationConfig_StreamingConfig_FromProto(mapCtx *direct.MapContext, in *pb.NotificationConfig_StreamingConfig) *krm.NotificationConfig_StreamingConfig {
	if in == nil {
		return nil
	}
	out := &krm.NotificationConfig_StreamingConfig{}
	out.Filter = direct.LazyPtr(in.GetFilter())
	return out
}
func NotificationConfig_StreamingConfig_ToProto(mapCtx *direct.MapContext, in *krm.NotificationConfig_StreamingConfig) *pb.NotificationConfig_StreamingConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotificationConfig_StreamingConfig{}
	out.Filter = direct.ValueOf(in.Filter)
	return out
}
func SecuritycenterNotificationConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NotificationConfig) *krm.SecuritycenterNotificationConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterNotificationConfigObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTopic
	// MISSING: ServiceAccount
	// MISSING: StreamingConfig
	return out
}
func SecuritycenterNotificationConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterNotificationConfigObservedState) *pb.NotificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotificationConfig{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTopic
	// MISSING: ServiceAccount
	// MISSING: StreamingConfig
	return out
}
func SecuritycenterNotificationConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.NotificationConfig) *krm.SecuritycenterNotificationConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterNotificationConfigSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTopic
	// MISSING: ServiceAccount
	// MISSING: StreamingConfig
	return out
}
func SecuritycenterNotificationConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterNotificationConfigSpec) *pb.NotificationConfig {
	if in == nil {
		return nil
	}
	out := &pb.NotificationConfig{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTopic
	// MISSING: ServiceAccount
	// MISSING: StreamingConfig
	return out
}
