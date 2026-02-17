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

package storage

import (
	"fmt"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	gcp "google.golang.org/api/storage/v1"
)

func StorageBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageBucketSpec) *gcp.Bucket {
	if in == nil {
		return nil
	}
	out := &gcp.Bucket{}
	out.Location = direct.ValueOf(in.Location)
	out.StorageClass = direct.ValueOf(in.StorageClass)
	out.DefaultEventBasedHold = direct.ValueOf(in.DefaultEventBasedHold)
	if in.DefaultEventBasedHold != nil {
		out.ForceSendFields = append(out.ForceSendFields, "DefaultEventBasedHold")
	}
	if in.CustomPlacementConfig != nil {
		out.CustomPlacementConfig = &gcp.BucketCustomPlacementConfig{
			DataLocations: in.CustomPlacementConfig.DataLocations,
		}
	}
	if in.Encryption != nil {
		out.Encryption = &gcp.BucketEncryption{
			DefaultKmsKeyName: in.Encryption.KmsKeyRef.External,
		}
	}
	if in.IPFilter != nil {
		out.IpFilter = &gcp.BucketIpFilter{
			Mode:                       direct.ValueOf(in.IPFilter.Mode),
			AllowCrossOrgVPCs:          direct.ValueOf(in.IPFilter.AllowCrossOrgVPCs),
			AllowAllServiceAgentAccess: direct.ValueOf(in.IPFilter.AllowAllServiceAgentAccess),
		}
		if in.IPFilter.AllowCrossOrgVPCs != nil {
			out.IpFilter.ForceSendFields = append(out.IpFilter.ForceSendFields, "AllowCrossOrgVPCs")
		}
		if in.IPFilter.AllowAllServiceAgentAccess != nil {
			out.IpFilter.ForceSendFields = append(out.IpFilter.ForceSendFields, "AllowAllServiceAgentAccess")
		}
		if in.IPFilter.PublicNetworkSource != nil {
			out.IpFilter.PublicNetworkSource = &gcp.BucketIpFilterPublicNetworkSource{
				AllowedIpCidrRanges: in.IPFilter.PublicNetworkSource.AllowedIPRanges,
			}
		}
		for _, vpc := range in.IPFilter.VpcNetworkSources {
			gcpVpc := &gcp.BucketIpFilterVpcNetworkSources{
				AllowedIpCidrRanges: vpc.AllowedIPRanges,
				Network:             vpc.NetworkRef.External,
			}
			out.IpFilter.VpcNetworkSources = append(out.IpFilter.VpcNetworkSources, gcpVpc)
		}
	}
	if in.Autoclass != nil {
		out.Autoclass = &gcp.BucketAutoclass{
			Enabled: in.Autoclass.Enabled,
		}
		out.Autoclass.ForceSendFields = append(out.Autoclass.ForceSendFields, "Enabled")
	}
	if in.Versioning != nil {
		out.Versioning = &gcp.BucketVersioning{
			Enabled: in.Versioning.Enabled,
		}
		out.Versioning.ForceSendFields = append(out.Versioning.ForceSendFields, "Enabled")
	}
	if in.Website != nil {
		out.Website = &gcp.BucketWebsite{
			MainPageSuffix: direct.ValueOf(in.Website.MainPageSuffix),
			NotFoundPage:   direct.ValueOf(in.Website.NotFoundPage),
		}
	}
	if in.Logging != nil {
		out.Logging = &gcp.BucketLogging{
			LogBucket:       in.Logging.LogBucket,
			LogObjectPrefix: direct.ValueOf(in.Logging.LogObjectPrefix),
		}
	}
	if in.RetentionPolicy != nil {
		out.RetentionPolicy = &gcp.BucketRetentionPolicy{
			IsLocked:        direct.ValueOf(in.RetentionPolicy.IsLocked),
			RetentionPeriod: in.RetentionPolicy.RetentionPeriod,
		}
		if in.RetentionPolicy.IsLocked != nil {
			out.RetentionPolicy.ForceSendFields = append(out.RetentionPolicy.ForceSendFields, "IsLocked")
		}
	}
	if in.SoftDeletePolicy != nil {
		out.SoftDeletePolicy = &gcp.BucketSoftDeletePolicy{
			RetentionDurationSeconds: direct.ValueOf(in.SoftDeletePolicy.RetentionDurationSeconds),
		}
		out.SoftDeletePolicy.ForceSendFields = append(out.SoftDeletePolicy.ForceSendFields, "RetentionDurationSeconds")
	}

	if in.RequesterPays != nil {
		out.Billing = &gcp.BucketBilling{
			RequesterPays: *in.RequesterPays,
		}
		out.Billing.ForceSendFields = append(out.Billing.ForceSendFields, "RequesterPays")
	}

	// Map top-level IAM fields to nested IamConfiguration
	if in.BucketPolicyOnly != nil || in.UniformBucketLevelAccess != nil || in.PublicAccessPrevention != nil {
		if out.IamConfiguration == nil {
			out.IamConfiguration = &gcp.BucketIamConfiguration{}
		}
		if in.BucketPolicyOnly != nil {
			out.IamConfiguration.BucketPolicyOnly = &gcp.BucketIamConfigurationBucketPolicyOnly{
				Enabled: *in.BucketPolicyOnly,
			}
			out.IamConfiguration.BucketPolicyOnly.ForceSendFields = append(out.IamConfiguration.BucketPolicyOnly.ForceSendFields, "Enabled")
		}
		if in.UniformBucketLevelAccess != nil {
			out.IamConfiguration.UniformBucketLevelAccess = &gcp.BucketIamConfigurationUniformBucketLevelAccess{
				Enabled: *in.UniformBucketLevelAccess,
			}
			out.IamConfiguration.UniformBucketLevelAccess.ForceSendFields = append(out.IamConfiguration.UniformBucketLevelAccess.ForceSendFields, "Enabled")
		}
		if in.PublicAccessPrevention != nil {
			out.IamConfiguration.PublicAccessPrevention = *in.PublicAccessPrevention
		}
	}

	if len(in.LifecycleRule) > 0 {
		out.Lifecycle = &gcp.BucketLifecycle{}
		for _, rule := range in.LifecycleRule {
			gcpRule := &gcp.BucketLifecycleRule{
				Action: &gcp.BucketLifecycleRuleAction{
					Type:         rule.Action.Type,
					StorageClass: direct.ValueOf(rule.Action.StorageClass),
				},
				Condition: &gcp.BucketLifecycleRuleCondition{
					Age:                     rule.Condition.Age,
					CreatedBefore:           direct.ValueOf(rule.Condition.CreatedBefore),
					CustomTimeBefore:        direct.ValueOf(rule.Condition.CustomTimeBefore),
					DaysSinceCustomTime:     direct.ValueOf(rule.Condition.DaysSinceCustomTime),
					DaysSinceNoncurrentTime: direct.ValueOf(rule.Condition.DaysSinceNoncurrentTime),
					MatchesPrefix:           rule.Condition.MatchesPrefix,
					MatchesStorageClass:     rule.Condition.MatchesStorageClass,
					MatchesSuffix:           rule.Condition.MatchesSuffix,
					NoncurrentTimeBefore:    direct.ValueOf(rule.Condition.NoncurrentTimeBefore),
					NumNewerVersions:        direct.ValueOf(rule.Condition.NumNewerVersions),
				},
			}
			if rule.Condition.WithState != nil {
				gcpRule.Condition.IsLive = direct.LazyPtr(*rule.Condition.WithState == "LIVE")
			}
			out.Lifecycle.Rule = append(out.Lifecycle.Rule, gcpRule)
		}
	}

	if len(in.Cors) > 0 {
		for _, cors := range in.Cors {
			gcpCors := &gcp.BucketCors{
				MaxAgeSeconds:  direct.ValueOf(cors.MaxAgeSeconds),
				Method:         cors.Method,
				Origin:         cors.Origin,
				ResponseHeader: cors.ResponseHeader,
			}
			out.Cors = append(out.Cors, gcpCors)
		}
	}

	return out
}

func StorageBucketSpec_FromProto(mapCtx *direct.MapContext, in *gcp.Bucket) *krm.StorageBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketSpec{}
	out.Location = direct.LazyPtr(in.Location)
	out.StorageClass = direct.LazyPtr(in.StorageClass)
	out.DefaultEventBasedHold = direct.LazyPtr(in.DefaultEventBasedHold)
	if in.CustomPlacementConfig != nil {
		out.CustomPlacementConfig = &krm.BucketCustomPlacementConfig{
			DataLocations: in.CustomPlacementConfig.DataLocations,
		}
	}
	if in.Encryption != nil {
		out.Encryption = &krm.BucketEncryption{
			KmsKeyRef: refsv1beta1.KMSCryptoKeyRef{External: in.Encryption.DefaultKmsKeyName},
		}
	}
	if in.IpFilter != nil {
		out.IPFilter = &krm.BucketIPFilter{
			Mode:                       direct.LazyPtr(in.IpFilter.Mode),
			AllowCrossOrgVPCs:          direct.LazyPtr(in.IpFilter.AllowCrossOrgVPCs),
			AllowAllServiceAgentAccess: direct.LazyPtr(in.IpFilter.AllowAllServiceAgentAccess),
		}
		if in.IpFilter.PublicNetworkSource != nil {
			out.IPFilter.PublicNetworkSource = &krm.BucketIPFilterPublicNetworkSource{
				AllowedIPRanges: in.IpFilter.PublicNetworkSource.AllowedIpCidrRanges,
			}
		}
		for _, vpc := range in.IpFilter.VpcNetworkSources {
			krmVpc := krm.BucketIPFilterVpcNetworkSource{
				AllowedIPRanges: vpc.AllowedIpCidrRanges,
				NetworkRef:      computev1beta1.ComputeNetworkRef{External: vpc.Network},
			}
			out.IPFilter.VpcNetworkSources = append(out.IPFilter.VpcNetworkSources, krmVpc)
		}
	}
	if in.Autoclass != nil {
		out.Autoclass = &krm.BucketAutoclass{
			Enabled: in.Autoclass.Enabled,
		}
	}
	if in.Versioning != nil {
		out.Versioning = &krm.BucketVersioning{
			Enabled: in.Versioning.Enabled,
		}
	}
	if in.Website != nil {
		out.Website = &krm.BucketWebsite{
			MainPageSuffix: direct.LazyPtr(in.Website.MainPageSuffix),
			NotFoundPage:   direct.LazyPtr(in.Website.NotFoundPage),
		}
	}
	if in.Logging != nil {
		out.Logging = &krm.BucketLogging{
			LogBucket:       in.Logging.LogBucket,
			LogObjectPrefix: direct.LazyPtr(in.Logging.LogObjectPrefix),
		}
	}
	if in.RetentionPolicy != nil {
		out.RetentionPolicy = &krm.BucketRetentionPolicy{
			IsLocked:        direct.LazyPtr(in.RetentionPolicy.IsLocked),
			RetentionPeriod: in.RetentionPolicy.RetentionPeriod,
		}
	}
	if in.SoftDeletePolicy != nil {
		out.SoftDeletePolicy = &krm.BucketSoftDeletePolicy{
			RetentionDurationSeconds: direct.LazyPtr(in.SoftDeletePolicy.RetentionDurationSeconds),
		}
	}

	if in.Billing != nil {
		out.RequesterPays = direct.LazyPtr(in.Billing.RequesterPays)
	}

	if in.IamConfiguration != nil {
		if in.IamConfiguration.BucketPolicyOnly != nil {
			out.BucketPolicyOnly = direct.LazyPtr(in.IamConfiguration.BucketPolicyOnly.Enabled)
		}
		if in.IamConfiguration.UniformBucketLevelAccess != nil {
			out.UniformBucketLevelAccess = direct.LazyPtr(in.IamConfiguration.UniformBucketLevelAccess.Enabled)
		}
		if in.IamConfiguration.PublicAccessPrevention != "" {
			out.PublicAccessPrevention = direct.LazyPtr(in.IamConfiguration.PublicAccessPrevention)
		}
	}

	if in.Lifecycle != nil {
		for _, rule := range in.Lifecycle.Rule {
			krmRule := krm.BucketLifecycleRule{
				Action: krm.BucketAction{
					Type:         rule.Action.Type,
					StorageClass: direct.LazyPtr(rule.Action.StorageClass),
				},
				Condition: krm.BucketCondition{
					Age:                     rule.Condition.Age,
					CreatedBefore:           direct.LazyPtr(rule.Condition.CreatedBefore),
					CustomTimeBefore:        direct.LazyPtr(rule.Condition.CustomTimeBefore),
					DaysSinceCustomTime:     direct.LazyPtr(rule.Condition.DaysSinceCustomTime),
					DaysSinceNoncurrentTime: direct.LazyPtr(rule.Condition.DaysSinceNoncurrentTime),
					MatchesPrefix:           rule.Condition.MatchesPrefix,
					MatchesStorageClass:     rule.Condition.MatchesStorageClass,
					MatchesSuffix:           rule.Condition.MatchesSuffix,
					NoncurrentTimeBefore:    direct.LazyPtr(rule.Condition.NoncurrentTimeBefore),
					NumNewerVersions:        direct.LazyPtr(rule.Condition.NumNewerVersions),
				},
			}
			if rule.Condition.IsLive != nil {
				if *rule.Condition.IsLive {
					krmRule.Condition.WithState = direct.LazyPtr("LIVE")
				} else {
					krmRule.Condition.WithState = direct.LazyPtr("ARCHIVED")
				}
			}
			out.LifecycleRule = append(out.LifecycleRule, krmRule)
		}
	}

	if len(in.Cors) > 0 {
		for _, cors := range in.Cors {
			krmCors := krm.BucketCors{
				MaxAgeSeconds:  direct.LazyPtr(cors.MaxAgeSeconds),
				Method:         cors.Method,
				Origin:         cors.Origin,
				ResponseHeader: cors.ResponseHeader,
			}
			out.Cors = append(out.Cors, krmCors)
		}
	}

	return out
}

func StorageBucketObservedState_FromProto(mapCtx *direct.MapContext, in *gcp.Bucket) *krm.BucketObservedStateStatus {
	if in == nil {
		return nil
	}
	out := &krm.BucketObservedStateStatus{}
	if in.SoftDeletePolicy != nil {
		out.SoftDeletePolicy = &krm.BucketSoftDeletePolicyStatus{
			EffectiveTime:            direct.LazyPtr(in.SoftDeletePolicy.EffectiveTime),
			RetentionDurationSeconds: direct.LazyPtr(in.SoftDeletePolicy.RetentionDurationSeconds),
		}
	}
	return out
}

func StorageBucketStatus_FromProto(mapCtx *direct.MapContext, in *gcp.Bucket) *krm.StorageBucketStatus {
	if in == nil {
		return nil
	}
	out := &krm.StorageBucketStatus{}
	out.ObservedState = StorageBucketObservedState_FromProto(mapCtx, in)
	out.SelfLink = direct.LazyPtr(in.SelfLink)
	if in.Name != "" {
		out.Url = direct.LazyPtr(fmt.Sprintf("gs://%s", in.Name))
	}
	return out
}
