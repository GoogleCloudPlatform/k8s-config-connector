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
// proto.message: google.pubsub.v1.Topic
// api.group: pubsub.cnrm.cloud.google.com

package pubsub

import (
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(pubSubTopicFuzzer())
}

func pubSubTopicFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.Topic{},
		PubSubTopicSpec_FromProto, PubSubTopicSpec_ToProto,
	)

	f.SpecField(".kms_key_name")
	f.SpecField(".message_retention_duration")
	f.SpecField(".message_storage_policy")
	f.SpecField(".message_storage_policy.allowed_persistence_regions")
	f.SpecField(".schema_settings")
	f.SpecField(".schema_settings.schema")
	f.SpecField(".schema_settings.encoding")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".state")
	f.Unimplemented_NotYetTriaged(".ingestion_data_source_settings")
	f.Unimplemented_NotYetTriaged(".message_transforms")
	f.Unimplemented_NotYetTriaged(".tags")
	f.Unimplemented_NotYetTriaged(".schema_settings.first_revision_id")
	f.Unimplemented_NotYetTriaged(".schema_settings.last_revision_id")
	f.Unimplemented_NotYetTriaged(".message_storage_policy.enforce_in_transit")

	return f
}
