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

package logginglogbucket

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func LoggingLogBucketSpec_FromProto(mapCtx *direct.MapContext, in *pb.LogBucket) *krm.LoggingLogBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogBucketSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.RetentionDays != 0 {
		days := int64(in.RetentionDays)
		out.RetentionDays = &days
	}
	out.Locked = direct.LazyPtr(in.GetLocked())
	out.EnableAnalytics = direct.LazyPtr(in.GetAnalyticsEnabled())

	if len(in.GetRestrictedFields()) > 0 {
		out.RestrictedFields = in.GetRestrictedFields()
	}

	if len(in.GetIndexConfigs()) > 0 {
		out.IndexConfigs = make([]krm.LogBucketIndexConfig, len(in.GetIndexConfigs()))
		for i, v := range in.GetIndexConfigs() {
			out.IndexConfigs[i] = krm.LogBucketIndexConfig{
				FieldPath: direct.LazyPtr(v.GetFieldPath()),
				Type:      direct.Enum_FromProto(mapCtx, v.GetType()),
			}
		}
	}

	if in.GetCmekSettings() != nil && in.GetCmekSettings().GetKmsKeyName() != "" {
		out.CmekSettings = &krm.LogBucketCmekSettings{
			KmsKeyRef: &refsv1beta1.KMSCryptoKeyRef{External: in.GetCmekSettings().GetKmsKeyName()},
		}
	}

	return out
}

func LoggingLogBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogBucketSpec) *pb.LogBucket {
	if in == nil {
		return nil
	}
	out := &pb.LogBucket{}
	out.Description = direct.ValueOf(in.Description)
	if in.RetentionDays != nil {
		out.RetentionDays = int32(*in.RetentionDays)
	}
	out.Locked = direct.ValueOf(in.Locked)
	out.AnalyticsEnabled = direct.ValueOf(in.EnableAnalytics)

	if len(in.RestrictedFields) > 0 {
		out.RestrictedFields = in.RestrictedFields
	}

	if len(in.IndexConfigs) > 0 {
		out.IndexConfigs = make([]*pb.IndexConfig, len(in.IndexConfigs))
		for i, v := range in.IndexConfigs {
			out.IndexConfigs[i] = &pb.IndexConfig{
				FieldPath: direct.ValueOf(v.FieldPath),
				Type:      direct.Enum_ToProto[pb.IndexType](mapCtx, v.Type),
			}
		}
	}

	if in.CmekSettings != nil && in.CmekSettings.KmsKeyRef != nil {
		out.CmekSettings = &pb.CmekSettings{
			KmsKeyName: in.CmekSettings.KmsKeyRef.External,
		}
	}

	return out
}

func LoggingLogBucketStatus_FromProto(mapCtx *direct.MapContext, in *pb.LogBucket) *krm.LoggingLogBucketStatus {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogBucketStatus{}
	out.ExternalRef = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LifecycleState = direct.Enum_FromProto(mapCtx, in.GetLifecycleState())
	return out
}

func LoggingLogBucketStatus_ToProto(mapCtx *direct.MapContext, in *krm.LoggingLogBucketStatus) *pb.LogBucket {
	if in == nil {
		return nil
	}
	out := &pb.LogBucket{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LifecycleState = direct.Enum_ToProto[pb.LifecycleState](mapCtx, in.LifecycleState)
	return out
}
