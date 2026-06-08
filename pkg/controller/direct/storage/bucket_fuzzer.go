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
	// - Location (*string)                              -> f.SpecField(".location")
	// - Logging (*StorageBucketLogging)                 -> f.SpecField(".logging")
	// - ResourceID (*string)                            -> f.Unimplemented_Identity(".name") / f.Unimplemented_Identity(".id")
	// - RetentionPolicy (*StorageBucketRetentionPolicy) -> f.SpecField(".retention_policy")
	// - StorageClass (*string)                          -> f.SpecField(".storage_class")
	// - Versioning (*StorageBucketVersioning)           -> f.SpecField(".versioning")
	// - Website (*StorageBucketWebsite)                 -> f.SpecField(".website")
	//
	// Unmapped KRM Spec fields:
	// - BucketPolicyOnly (*bool)                         -> Unmapped/Deprecated (corresponds to iam_configuration.bucket_policy_only in proto, which is untriaged)
	// - CustomPlacementConfig (*StorageBucketCustomPlacementConfig) -> Unmapped in KCC direct mapper (no corresponding field in google.storage.v1.Bucket proto)
	// - LifecycleRule ([]StorageBucketLifecycleRule)    -> f.Unimplemented_NotYetTriaged(".lifecycle")
	// - PublicAccessPrevention (*string)                 -> f.Unimplemented_NotYetTriaged(".iam_configuration")
	// - RequesterPays (*bool)                           -> f.Unimplemented_NotYetTriaged(".billing")
	// - SoftDeletePolicy (*StorageBucketSoftDeletePolicy) -> Unmapped in direct mapper (no corresponding field in google.storage.v1.Bucket proto)
	// - UniformBucketLevelAccess (*bool)                -> f.Unimplemented_NotYetTriaged(".iam_configuration")

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

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".id")
	f.Unimplemented_Identity(".project_number")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_Etag()

	f.Unimplemented_NotYetTriaged(".acl")
	f.Unimplemented_NotYetTriaged(".default_object_acl")
	f.Unimplemented_NotYetTriaged(".lifecycle")
	f.Unimplemented_NotYetTriaged(".time_created")
	f.Unimplemented_NotYetTriaged(".metageneration")
	f.Unimplemented_NotYetTriaged(".updated")
	f.Unimplemented_NotYetTriaged(".owner")
	f.Unimplemented_NotYetTriaged(".billing")
	f.Unimplemented_NotYetTriaged(".location_type")
	f.Unimplemented_NotYetTriaged(".iam_configuration")
	f.Unimplemented_NotYetTriaged(".zone_affinity")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".autoclass.toggle_time")
	f.Unimplemented_NotYetTriaged(".retention_policy.effective_time")

	return f
}
