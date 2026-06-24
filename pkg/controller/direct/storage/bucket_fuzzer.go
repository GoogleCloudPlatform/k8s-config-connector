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
// See the License for the() Apache License, Version 2.0.

// +tool:fuzz-gen
// proto.message: google.storage.v1.Bucket
// api.group: storage.cnrm.cloud.google.com

package storage

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	pb "google.golang.org/genproto/googleapis/storage/v1"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(StorageBucketFuzzer())
}

func StorageBucketFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.Bucket{},
		StorageBucketSpec_FromProto, StorageBucketSpec_ToProto,
	)

	// Comparison of KRM StorageBucketSpec fields with google.storage.v1.Bucket proto fields:
	//
	// Mapped KRM Spec fields:
	// - Autoclass (StorageBucketAutoclass)              -> f.SpecField(".autoclass")
	// - Cors ([]StorageBucketCors)                      -> f.SpecField(".cors")
	// - DefaultEventBasedHold (*bool)                    -> f.SpecField(".default_event_based_hold")
	// - Encryption (*StorageBucketEncryption)           -> f.SpecField(".encryption")
	// - LifecycleRule ([]StorageBucketLifecycleRule)    -> f.SpecField(".lifecycle")
	// - Location (*string)                              -> f.SpecField(".location")
	// - Logging (*StorageBucketLogging)                 -> f.SpecField(".logging")
	// - ResourceID (*string)                            -> f.Unimplemented_Identity(".name") / f.Unimplemented_Identity(".id")
	// - RetentionPolicy (*StorageBucketRetentionPolicy) -> f.SpecField(".retention_policy")
	// - StorageClass (*string)                          -> f.SpecField(".storage_class")
	// - Versioning (*StorageBucketVersioning)           -> f.SpecField(".versioning")
	// - Website (*StorageBucketWebsite)                 -> f.SpecField(".website")
	// - PublicAccessPrevention (*string)                 -> f.SpecField(".iam_configuration")
	// - UniformBucketLevelAccess (*bool)                -> f.SpecField(".iam_configuration")
	// - RequesterPays (*bool)                           -> f.SpecField(".billing")
	//
	// Unmapped KRM Spec fields:
	// - BucketPolicyOnly (*bool)                         -> Unmapped/Deprecated (corresponds to iam_configuration.bucket_policy_only in proto, which is untriaged)
	// - CustomPlacementConfig (*StorageBucketCustomPlacementConfig) -> Unmapped in KCC direct mapper (no corresponding field in google.storage.v1.Bucket proto)
	// - SoftDeletePolicy (*StorageBucketSoftDeletePolicy) -> Unmapped in direct mapper (no corresponding field in google.storage.v1.Bucket proto)

	f.SpecField(".cors")
	f.SpecField(".location")
	f.SpecField(".storage_class")
	f.SpecField(".default_event_based_hold")
	f.SpecField(".website")
	f.SpecField(".versioning")
	f.SpecField(".logging")
	f.SpecField(".encryption")
	f.SpecField(".retention_policy")
	f.SpecField(".autoclass")
	f.SpecField(".lifecycle")
	f.SpecField(".billing")
	f.SpecField(".iam_configuration")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".id")
	f.Unimplemented_Identity(".project_number")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_Etag()

	f.Unimplemented_NotYetTriaged(".acl")
	f.Unimplemented_NotYetTriaged(".default_object_acl")
	f.Unimplemented_NotYetTriaged(".time_created")
	f.Unimplemented_NotYetTriaged(".metageneration")
	f.Unimplemented_NotYetTriaged(".updated")
	f.Unimplemented_NotYetTriaged(".owner")
	f.Unimplemented_NotYetTriaged(".location_type")
	f.Unimplemented_NotYetTriaged(".zone_affinity")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".autoclass.toggle_time")
	f.Unimplemented_NotYetTriaged(".retention_policy.effective_time")
	f.Unimplemented_NotYetTriaged(".lifecycle.rule[].condition.matches_pattern")
	f.Unimplemented_NotYetTriaged(".iam_configuration.uniform_bucket_level_access.locked_time")

	f.FilterSpec = func(in *pb.Bucket) {
		if in.Lifecycle != nil {
			if len(in.Lifecycle.Rule) == 0 {
				in.Lifecycle = nil
			} else {
				var activeRules []*pb.Bucket_Lifecycle_Rule
				for _, r := range in.Lifecycle.Rule {
					if r == nil {
						continue
					}
					// If both Action and Condition are empty, skip the rule
					if r.Action == nil && r.Condition == nil {
						continue
					}
					// If Action is empty and Condition is empty, skip
					if (r.Action == nil || (r.Action.Type == "" && r.Action.StorageClass == "")) &&
						(r.Condition == nil || (r.Condition.Age == 0 && r.Condition.CreatedBefore == nil &&
							r.Condition.CustomTimeBefore == nil && r.Condition.DaysSinceCustomTime == 0 &&
							r.Condition.DaysSinceNoncurrentTime == 0 && len(r.Condition.MatchesPrefix) == 0 &&
							len(r.Condition.MatchesStorageClass) == 0 && len(r.Condition.MatchesSuffix) == 0 &&
							r.Condition.NoncurrentTimeBefore == nil && r.Condition.NumNewerVersions == 0 &&
							r.Condition.IsLive == nil)) {
						continue
					}
					activeRules = append(activeRules, r)
				}
				if len(activeRules) == 0 {
					in.Lifecycle = nil
				} else {
					in.Lifecycle.Rule = activeRules
				}
			}
		}

		if in.IamConfiguration != nil {
			if in.IamConfiguration.PublicAccessPrevention == pb.Bucket_IamConfiguration_PUBLIC_ACCESS_PREVENTION_UNSPECIFIED &&
				in.IamConfiguration.UniformBucketLevelAccess == nil {
				in.IamConfiguration = nil
			} else if in.IamConfiguration.UniformBucketLevelAccess != nil {
				in.IamConfiguration.UniformBucketLevelAccess.LockedTime = nil
			}
		}
	}

	return f
}
