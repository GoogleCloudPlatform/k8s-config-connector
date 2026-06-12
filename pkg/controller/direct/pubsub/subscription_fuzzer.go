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
// proto.message: google.pubsub.v1.Subscription
// api.group: pubsub.cnrm.cloud.google.com

package pubsub

import (
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(pubSubSubscriptionFuzzer())
}

func pubSubSubscriptionFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.Subscription{},
		PubSubSubscriptionSpec_FromProto, PubSubSubscriptionSpec_ToProto,
	)

	f.SpecField(".topic")
	f.SpecField(".ack_deadline_seconds")
	f.SpecField(".push_config")
	f.SpecField(".bigquery_config")
	f.SpecField(".cloud_storage_config")
	f.SpecField(".dead_letter_policy")
	f.SpecField(".enable_exactly_once_delivery")
	f.SpecField(".enable_message_ordering")
	f.SpecField(".expiration_policy")
	f.SpecField(".filter")
	f.SpecField(".message_retention_duration")
	f.SpecField(".retain_acked_messages")
	f.SpecField(".retry_policy")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_LabelsAnnotations(".tags")

	f.Unimplemented_NotYetTriaged(".detached")
	f.Unimplemented_NotYetTriaged(".message_transforms")
	f.Unimplemented_NotYetTriaged(".topic_message_retention_duration")
	f.Unimplemented_NotYetTriaged(".state")
	f.Unimplemented_NotYetTriaged(".analytics_hub_subscription_info")

	f.Unimplemented_NotYetTriaged(".bigquery_config.service_account_email")
	f.Unimplemented_NotYetTriaged(".bigquery_config.state")
	f.Unimplemented_NotYetTriaged(".bigquery_config.use_table_schema")

	f.Unimplemented_NotYetTriaged(".cloud_storage_config.filename_datetime_format")
	f.Unimplemented_NotYetTriaged(".cloud_storage_config.text_config")
	f.Unimplemented_NotYetTriaged(".cloud_storage_config.service_account_email")
	f.Unimplemented_NotYetTriaged(".cloud_storage_config.max_messages")
	f.Unimplemented_NotYetTriaged(".cloud_storage_config.avro_config.use_topic_schema")

	f.Unimplemented_NotYetTriaged(".push_config.pubsub_wrapper")

	return f
}
