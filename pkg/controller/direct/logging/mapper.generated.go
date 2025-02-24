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

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CmekSettings_FromProto(mapCtx *direct.MapContext, in *pb.CmekSettings) *krm.CmekSettings {
	if in == nil {
		return nil
	}
	out := &krm.CmekSettings{}
	// MISSING: Name
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	out.KMSKeyVersionName = direct.LazyPtr(in.GetKmsKeyVersionName())
	// MISSING: ServiceAccountID
	return out
}
func CmekSettings_ToProto(mapCtx *direct.MapContext, in *krm.CmekSettings) *pb.CmekSettings {
	if in == nil {
		return nil
	}
	out := &pb.CmekSettings{}
	// MISSING: Name
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	out.KmsKeyVersionName = direct.ValueOf(in.KMSKeyVersionName)
	// MISSING: ServiceAccountID
	return out
}
func CmekSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CmekSettings) *krm.CmekSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CmekSettingsObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	out.ServiceAccountID = direct.LazyPtr(in.GetServiceAccountId())
	return out
}
func CmekSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CmekSettingsObservedState) *pb.CmekSettings {
	if in == nil {
		return nil
	}
	out := &pb.CmekSettings{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: KMSKeyName
	// MISSING: KMSKeyVersionName
	out.ServiceAccountId = direct.ValueOf(in.ServiceAccountID)
	return out
}
func IndexConfig_FromProto(mapCtx *direct.MapContext, in *pb.IndexConfig) *krm.IndexConfig {
	if in == nil {
		return nil
	}
	out := &krm.IndexConfig{}
	out.FieldPath = direct.LazyPtr(in.GetFieldPath())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: CreateTime
	return out
}
func IndexConfig_ToProto(mapCtx *direct.MapContext, in *krm.IndexConfig) *pb.IndexConfig {
	if in == nil {
		return nil
	}
	out := &pb.IndexConfig{}
	out.FieldPath = direct.ValueOf(in.FieldPath)
	out.Type = direct.Enum_ToProto[pb.IndexType](mapCtx, in.Type)
	// MISSING: CreateTime
	return out
}
func IndexConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.IndexConfig) *krm.IndexConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IndexConfigObservedState{}
	// MISSING: FieldPath
	// MISSING: Type
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func IndexConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IndexConfigObservedState) *pb.IndexConfig {
	if in == nil {
		return nil
	}
	out := &pb.IndexConfig{}
	// MISSING: FieldPath
	// MISSING: Type
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func LoggingLogBucketObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LogBucket) *krm.LoggingLogBucketObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogBucketObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RetentionDays
	// MISSING: Locked
	// MISSING: LifecycleState
	// MISSING: AnalyticsEnabled
	// MISSING: RestrictedFields
	// MISSING: IndexConfigs
	// MISSING: CmekSettings
	return out
}
func LoggingLogBucketObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogBucketObservedState) *pb.LogBucket {
	if in == nil {
		return nil
	}
	out := &pb.LogBucket{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RetentionDays
	// MISSING: Locked
	// MISSING: LifecycleState
	// MISSING: AnalyticsEnabled
	// MISSING: RestrictedFields
	// MISSING: IndexConfigs
	// MISSING: CmekSettings
	return out
}
func LoggingLogBucketSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogBucket) *krm.LoggingLogBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogBucketSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RetentionDays
	// MISSING: Locked
	// MISSING: LifecycleState
	// MISSING: AnalyticsEnabled
	// MISSING: RestrictedFields
	// MISSING: IndexConfigs
	// MISSING: CmekSettings
	return out
}
func LoggingLogBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogBucketSpec) *pb.LogBucket {
	if in == nil {
		return nil
	}
	out := &pb.LogBucket{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RetentionDays
	// MISSING: Locked
	// MISSING: LifecycleState
	// MISSING: AnalyticsEnabled
	// MISSING: RestrictedFields
	// MISSING: IndexConfigs
	// MISSING: CmekSettings
	return out
}
