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

// krm.group: storage.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.storage.v1

package storage

import (
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "google.golang.org/genproto/googleapis/storage/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func StorageBucketCors_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Cors) *krm.StorageBucketCors {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketCors{}
	out.MaxAgeSeconds = direct.PtrTo(int(in.MaxAgeSeconds))
	out.Method = in.Method
	out.Origin = in.Origin
	out.ResponseHeader = in.ResponseHeader
	return out
}
func StorageBucketCors_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketCors) *pb.Bucket_Cors {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Cors{}
	if in.MaxAgeSeconds != nil {
		out.MaxAgeSeconds = int32(*in.MaxAgeSeconds)
	}
	out.Method = in.Method
	out.Origin = in.Origin
	out.ResponseHeader = in.ResponseHeader
	return out
}

func StorageBucketWebsite_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Website) *krm.StorageBucketWebsite {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketWebsite{}
	if in.MainPageSuffix != "" {
		out.MainPageSuffix = direct.PtrTo(in.MainPageSuffix)
	}
	if in.NotFoundPage != "" {
		out.NotFoundPage = direct.PtrTo(in.NotFoundPage)
	}
	return out
}
func StorageBucketWebsite_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketWebsite) *pb.Bucket_Website {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Website{}
	if in.MainPageSuffix != nil {
		out.MainPageSuffix = *in.MainPageSuffix
	}
	if in.NotFoundPage != nil {
		out.NotFoundPage = *in.NotFoundPage
	}
	return out
}

func StorageBucketVersioning_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Versioning) *krm.StorageBucketVersioning {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketVersioning{}
	out.Enabled = in.Enabled
	return out
}
func StorageBucketVersioning_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketVersioning) *pb.Bucket_Versioning {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Versioning{}
	out.Enabled = in.Enabled
	return out
}

func StorageBucketLogging_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Logging) *krm.StorageBucketLogging {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketLogging{}
	out.LogBucket = in.LogBucket
	if in.LogObjectPrefix != "" {
		out.LogObjectPrefix = direct.PtrTo(in.LogObjectPrefix)
	}
	return out
}
func StorageBucketLogging_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketLogging) *pb.Bucket_Logging {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Logging{}
	out.LogBucket = in.LogBucket
	if in.LogObjectPrefix != nil {
		out.LogObjectPrefix = *in.LogObjectPrefix
	}
	return out
}

func StorageBucketEncryption_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Encryption) *krm.StorageBucketEncryption {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketEncryption{}
	if in.DefaultKmsKeyName != "" {
		out.KmsKeyRef = &refsv1beta1.KMSCryptoKeyRef{External: in.DefaultKmsKeyName}
	}
	return out
}
func StorageBucketEncryption_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketEncryption) *pb.Bucket_Encryption {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Encryption{}
	if in.KmsKeyRef != nil {
		out.DefaultKmsKeyName = in.KmsKeyRef.External
	}
	return out
}

func StorageBucketRetentionPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_RetentionPolicy) *krm.StorageBucketRetentionPolicy {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketRetentionPolicy{}
	out.IsLocked = direct.PtrTo(in.IsLocked)
	out.RetentionPeriod = int(in.RetentionPeriod)
	return out
}
func StorageBucketRetentionPolicy_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketRetentionPolicy) *pb.Bucket_RetentionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_RetentionPolicy{}
	if in.IsLocked != nil {
		out.IsLocked = *in.IsLocked
	}
	out.RetentionPeriod = int64(in.RetentionPeriod)
	return out
}

func StorageBucketAutoclass_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Autoclass) *krm.StorageBucketAutoclass {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketAutoclass{}
	out.Enabled = direct.PtrTo(in.Enabled)
	return out
}
func StorageBucketAutoclass_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketAutoclass) *pb.Bucket_Autoclass {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Autoclass{}
	if in.Enabled != nil {
		out.Enabled = *in.Enabled
	}
	return out
}

func StorageBucketLifecycleRuleAction_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Lifecycle_Rule_Action) *krm.StorageBucketLifecycleRuleAction {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketLifecycleRuleAction{}
	if in.StorageClass != "" {
		out.StorageClass = &in.StorageClass
	}
	out.Type = in.Type
	return out
}

func StorageBucketLifecycleRuleAction_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketLifecycleRuleAction) *pb.Bucket_Lifecycle_Rule_Action {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Lifecycle_Rule_Action{}
	if in.StorageClass != nil {
		out.StorageClass = *in.StorageClass
	}
	out.Type = in.Type
	return out
}

func StorageBucketLifecycleRuleCondition_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Lifecycle_Rule_Condition) *krm.StorageBucketLifecycleRuleCondition {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketLifecycleRuleCondition{}
	if in.Age != 0 {
		out.Age = direct.PtrTo(int(in.Age))
	}
	out.CreatedBefore = direct.StringTimestamp_FromProto(mapCtx, in.CreatedBefore)
	out.CustomTimeBefore = direct.StringTimestamp_FromProto(mapCtx, in.CustomTimeBefore)
	if in.DaysSinceCustomTime != 0 {
		out.DaysSinceCustomTime = direct.PtrTo(int(in.DaysSinceCustomTime))
	}
	if in.DaysSinceNoncurrentTime != 0 {
		out.DaysSinceNoncurrentTime = direct.PtrTo(int(in.DaysSinceNoncurrentTime))
	}
	out.MatchesPrefix = in.MatchesPrefix
	out.MatchesStorageClass = in.MatchesStorageClass
	out.MatchesSuffix = in.MatchesSuffix
	out.NoncurrentTimeBefore = direct.StringTimestamp_FromProto(mapCtx, in.NoncurrentTimeBefore)
	if in.NumNewerVersions != 0 {
		out.NumNewerVersions = direct.PtrTo(int(in.NumNewerVersions))
	}
	if in.IsLive != nil {
		if in.IsLive.Value {
			out.WithState = direct.PtrTo("LIVE")
		} else {
			out.WithState = direct.PtrTo("ARCHIVED")
		}
	}
	return out
}

func StorageBucketLifecycleRuleCondition_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketLifecycleRuleCondition) *pb.Bucket_Lifecycle_Rule_Condition {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Lifecycle_Rule_Condition{}
	if in.Age != nil {
		out.Age = int32(*in.Age)
	}
	out.CreatedBefore = direct.StringTimestamp_ToProto(mapCtx, in.CreatedBefore)
	out.CustomTimeBefore = direct.StringTimestamp_ToProto(mapCtx, in.CustomTimeBefore)
	if in.DaysSinceCustomTime != nil {
		out.DaysSinceCustomTime = int32(*in.DaysSinceCustomTime)
	}
	if in.DaysSinceNoncurrentTime != nil {
		out.DaysSinceNoncurrentTime = int32(*in.DaysSinceNoncurrentTime)
	}
	out.MatchesPrefix = in.MatchesPrefix
	out.MatchesStorageClass = in.MatchesStorageClass
	out.MatchesSuffix = in.MatchesSuffix
	out.NoncurrentTimeBefore = direct.StringTimestamp_ToProto(mapCtx, in.NoncurrentTimeBefore)
	if in.NumNewerVersions != nil {
		out.NumNewerVersions = int32(*in.NumNewerVersions)
	}
	if in.WithState != nil {
		switch *in.WithState {
		case "LIVE":
			out.IsLive = &wrapperspb.BoolValue{Value: true}
		case "ARCHIVED":
			out.IsLive = &wrapperspb.BoolValue{Value: false}
		}
	}
	return out
}

func StorageBucketLifecycleRule_FromProto(mapCtx *direct.MapContext, in *pb.Bucket_Lifecycle_Rule) *krm.StorageBucketLifecycleRule {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketLifecycleRule{}
	out.Action = StorageBucketLifecycleRuleAction_FromProto(mapCtx, in.Action)
	out.Condition = StorageBucketLifecycleRuleCondition_FromProto(mapCtx, in.Condition)
	return out
}

func StorageBucketLifecycleRule_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketLifecycleRule) *pb.Bucket_Lifecycle_Rule {
	if in == nil {
		return nil
	}
	out := &pb.Bucket_Lifecycle_Rule{}
	out.Action = StorageBucketLifecycleRuleAction_ToProto(mapCtx, in.Action)
	out.Condition = StorageBucketLifecycleRuleCondition_ToProto(mapCtx, in.Condition)
	return out
}

func StorageBucketSpec_FromProto(mapCtx *direct.MapContext, in *pb.Bucket) *krm.StorageBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketSpec{}
	// MISSING: Acl
	// MISSING: DefaultObjectAcl
	if in.Lifecycle != nil {
		out.LifecycleRule = direct.Slice_FromProto(mapCtx, in.Lifecycle.Rule, StorageBucketLifecycleRule_FromProto)
	}
	// MISSING: TimeCreated
	// MISSING: ID
	// MISSING: Name
	// MISSING: ProjectNumber
	// MISSING: Metageneration
	out.Cors = direct.Slice_FromProto(mapCtx, in.Cors, StorageBucketCors_FromProto)
	out.Location = direct.LazyPtr(in.GetLocation())
	out.StorageClass = direct.LazyPtr(in.GetStorageClass())
	// MISSING: Etag
	// MISSING: Updated
	out.DefaultEventBasedHold = direct.LazyPtr(in.GetDefaultEventBasedHold())
	// MISSING: Labels
	out.Website = StorageBucketWebsite_FromProto(mapCtx, in.GetWebsite())
	out.Versioning = StorageBucketVersioning_FromProto(mapCtx, in.GetVersioning())
	out.Logging = StorageBucketLogging_FromProto(mapCtx, in.GetLogging())
	// MISSING: Owner
	out.Encryption = StorageBucketEncryption_FromProto(mapCtx, in.GetEncryption())
	// MISSING: Billing
	out.RetentionPolicy = StorageBucketRetentionPolicy_FromProto(mapCtx, in.GetRetentionPolicy())
	// MISSING: LocationType
	if in.IamConfiguration != nil {
		if in.IamConfiguration.UniformBucketLevelAccess != nil {
			out.UniformBucketLevelAccess = direct.PtrTo(in.IamConfiguration.UniformBucketLevelAccess.Enabled)
			out.BucketPolicyOnly = direct.PtrTo(in.IamConfiguration.UniformBucketLevelAccess.Enabled)
		}
		if in.IamConfiguration.PublicAccessPrevention != pb.Bucket_IamConfiguration_PUBLIC_ACCESS_PREVENTION_UNSPECIFIED {
			switch in.IamConfiguration.PublicAccessPrevention {
			case pb.Bucket_IamConfiguration_ENFORCED:
				out.PublicAccessPrevention = direct.LazyPtr("enforced")
			case pb.Bucket_IamConfiguration_INHERITED:
				out.PublicAccessPrevention = direct.LazyPtr("inherited")
			}
		}
	}
	// MISSING: ZoneAffinity
	// MISSING: SatisfiesPzs
	out.Autoclass = StorageBucketAutoclass_FromProto(mapCtx, in.GetAutoclass())
	return out
}

func StorageBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketSpec) *pb.Bucket {
	if in == nil {
		return nil
	}
	out := &pb.Bucket{}
	// MISSING: Acl
	// MISSING: DefaultObjectAcl
	if len(in.LifecycleRule) > 0 {
		out.Lifecycle = &pb.Bucket_Lifecycle{
			Rule: direct.Slice_ToProto(mapCtx, in.LifecycleRule, StorageBucketLifecycleRule_ToProto),
		}
	}
	// MISSING: TimeCreated
	// MISSING: ID
	// MISSING: Name
	// MISSING: ProjectNumber
	// MISSING: Metageneration
	out.Cors = direct.Slice_ToProto(mapCtx, in.Cors, StorageBucketCors_ToProto)
	out.Location = direct.ValueOf(in.Location)
	out.StorageClass = direct.ValueOf(in.StorageClass)
	// MISSING: Etag
	// MISSING: Updated
	out.DefaultEventBasedHold = direct.ValueOf(in.DefaultEventBasedHold)
	// MISSING: Labels
	out.Website = StorageBucketWebsite_ToProto(mapCtx, in.Website)
	out.Versioning = StorageBucketVersioning_ToProto(mapCtx, in.Versioning)
	out.Logging = StorageBucketLogging_ToProto(mapCtx, in.Logging)
	// MISSING: Owner
	out.Encryption = StorageBucketEncryption_ToProto(mapCtx, in.Encryption)
	// MISSING: Billing
	out.RetentionPolicy = StorageBucketRetentionPolicy_ToProto(mapCtx, in.RetentionPolicy)
	// MISSING: LocationType
	if in.UniformBucketLevelAccess != nil || in.PublicAccessPrevention != nil || in.BucketPolicyOnly != nil {
		out.IamConfiguration = &pb.Bucket_IamConfiguration{}
		ubla := false
		if in.UniformBucketLevelAccess != nil {
			ubla = *in.UniformBucketLevelAccess
		} else if in.BucketPolicyOnly != nil {
			ubla = *in.BucketPolicyOnly
		}
		out.IamConfiguration.UniformBucketLevelAccess = &pb.Bucket_IamConfiguration_UniformBucketLevelAccess{
			Enabled: ubla,
		}

		if in.PublicAccessPrevention != nil {
			switch strings.ToLower(*in.PublicAccessPrevention) {
			case "enforced":
				out.IamConfiguration.PublicAccessPrevention = pb.Bucket_IamConfiguration_ENFORCED
			case "inherited":
				out.IamConfiguration.PublicAccessPrevention = pb.Bucket_IamConfiguration_INHERITED
			default:
				out.IamConfiguration.PublicAccessPrevention = pb.Bucket_IamConfiguration_PUBLIC_ACCESS_PREVENTION_UNSPECIFIED
			}
		}
	}
	// MISSING: ZoneAffinity
	// MISSING: SatisfiesPzs
	out.Autoclass = StorageBucketAutoclass_ToProto(mapCtx, in.Autoclass)
	return out
}
