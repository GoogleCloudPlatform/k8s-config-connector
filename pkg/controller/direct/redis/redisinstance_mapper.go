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

package redis

import (
	redispb "cloud.google.com/go/redis/apiv1/redispb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	dayofweekpb "google.golang.org/genproto/googleapis/type/dayofweek"
)

func InstanceWeeklyMaintenanceWindow_FromProto(mapCtx *direct.MapContext, in *redispb.WeeklyMaintenanceWindow) *krm.InstanceWeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.InstanceWeeklyMaintenanceWindow{}
	out.Day = direct.ValueOf(direct.Enum_FromProto(mapCtx, in.GetDay()))
	if in.GetStartTime() != nil {
		out.StartTime = *TimeOfDay_FromProto(mapCtx, in.GetStartTime())
	}
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	return out
}

func InstanceWeeklyMaintenanceWindow_ToProto(mapCtx *direct.MapContext, in *krm.InstanceWeeklyMaintenanceWindow) *redispb.WeeklyMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &redispb.WeeklyMaintenanceWindow{}
	out.Day = direct.Enum_ToProto[dayofweekpb.DayOfWeek](mapCtx, &in.Day)
	if in.StartTime.Hours != nil || in.StartTime.Minutes != nil || in.StartTime.Seconds != nil || in.StartTime.Nanos != nil {
		out.StartTime = TimeOfDay_ToProto(mapCtx, &in.StartTime)
	}
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	return out
}

func RedisInstanceSpec_FromProto(mapCtx *direct.MapContext, in *redispb.Instance) *krm.RedisInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.RedisInstanceSpec{}
	out.AlternativeLocationId = direct.LazyPtr(in.GetAlternativeLocationId())
	out.AuthEnabled = direct.LazyPtr(in.GetAuthEnabled())
	if in.GetAuthorizedNetwork() != "" {
		out.AuthorizedNetworkRef = &krm.InstanceAuthorizedNetworkRef{External: direct.LazyPtr(in.GetAuthorizedNetwork())}
	}
	out.ConnectMode = direct.Enum_FromProto(mapCtx, in.GetConnectMode())
	if in.GetCustomerManagedKey() != "" {
		out.CustomerManagedKeyRef = &krm.InstanceCustomerManagedKeyRef{External: direct.LazyPtr(in.GetCustomerManagedKey())}
	}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.LocationId = direct.LazyPtr(in.GetLocationId())
	out.MaintenancePolicy = InstanceMaintenancePolicy_FromProto(mapCtx, in.GetMaintenancePolicy())
	if v := in.GetMaintenanceSchedule(); v != nil {
		out.MaintenanceSchedule = []krm.InstanceMaintenanceSchedule{*InstanceMaintenanceSchedule_FromProto(mapCtx, v)}
	}
	out.MemorySizeGb = int64(in.GetMemorySizeGb())
	out.PersistenceConfig = InstancePersistenceConfig_FromProto(mapCtx, in.GetPersistenceConfig())
	out.ReadReplicasMode = direct.Enum_FromProto(mapCtx, in.GetReadReplicasMode())
	out.RedisConfigs = in.GetRedisConfigs()
	out.RedisVersion = direct.LazyPtr(in.GetRedisVersion())
	out.ReplicaCount = direct.LazyPtr(int64(in.GetReplicaCount()))
	out.ReservedIpRange = direct.LazyPtr(in.GetReservedIpRange())
	out.SecondaryIpRange = direct.LazyPtr(in.GetSecondaryIpRange())
	out.Tier = direct.Enum_FromProto(mapCtx, in.GetTier())
	out.TransitEncryptionMode = direct.Enum_FromProto(mapCtx, in.GetTransitEncryptionMode())
	return out
}

func RedisInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.RedisInstanceSpec) *redispb.Instance {
	if in == nil {
		return nil
	}
	out := &redispb.Instance{}
	out.AlternativeLocationId = direct.ValueOf(in.AlternativeLocationId)
	out.AuthEnabled = direct.ValueOf(in.AuthEnabled)
	if in.AuthorizedNetworkRef != nil {
		out.AuthorizedNetwork = direct.ValueOf(in.AuthorizedNetworkRef.External)
	}
	out.ConnectMode = direct.Enum_ToProto[redispb.Instance_ConnectMode](mapCtx, in.ConnectMode)
	if in.CustomerManagedKeyRef != nil {
		out.CustomerManagedKey = direct.ValueOf(in.CustomerManagedKeyRef.External)
	}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.LocationId = direct.ValueOf(in.LocationId)
	out.MaintenancePolicy = InstanceMaintenancePolicy_ToProto(mapCtx, in.MaintenancePolicy)
	if len(in.MaintenanceSchedule) > 0 {
		out.MaintenanceSchedule = InstanceMaintenanceSchedule_ToProto(mapCtx, &in.MaintenanceSchedule[0])
	}
	out.MemorySizeGb = int32(in.MemorySizeGb)
	out.PersistenceConfig = InstancePersistenceConfig_ToProto(mapCtx, in.PersistenceConfig)
	out.ReadReplicasMode = direct.Enum_ToProto[redispb.Instance_ReadReplicasMode](mapCtx, in.ReadReplicasMode)
	out.RedisConfigs = in.RedisConfigs
	out.RedisVersion = direct.ValueOf(in.RedisVersion)
	out.ReplicaCount = int32(direct.ValueOf(in.ReplicaCount))
	out.ReservedIpRange = direct.ValueOf(in.ReservedIpRange)
	out.SecondaryIpRange = direct.ValueOf(in.SecondaryIpRange)
	out.Tier = direct.Enum_ToProto[redispb.Instance_Tier](mapCtx, in.Tier)
	out.TransitEncryptionMode = direct.Enum_ToProto[redispb.Instance_TransitEncryptionMode](mapCtx, in.TransitEncryptionMode)
	return out
}

func InstanceObservedStateStatus_FromProto(mapCtx *direct.MapContext, in *redispb.Instance) *krm.InstanceObservedStateStatus {
	if in == nil {
		return nil
	}
	out := &krm.InstanceObservedStateStatus{}
	return out
}

func InstanceObservedStateStatus_ToProto(mapCtx *direct.MapContext, in *krm.InstanceObservedStateStatus) *redispb.Instance {
	if in == nil {
		return nil
	}
	out := &redispb.Instance{}
	return out
}
