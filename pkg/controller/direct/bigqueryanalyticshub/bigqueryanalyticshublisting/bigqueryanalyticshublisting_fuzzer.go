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
// proto.message: google.cloud.bigquery.analyticshub.v1.Listing
// api.group: bigqueryanalyticshub.cnrm.cloud.google.com

package bigqueryanalyticshublisting

import (
	pb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(bigQueryAnalyticsHubListingFuzzer())
}

func bigQueryAnalyticsHubListingFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Listing{},
		BigQueryAnalyticsHubListingSpec_FromProto, BigQueryAnalyticsHubListingSpec_ToProto,
		BigQueryAnalyticsHubListingObservedState_FromProto, BigQueryAnalyticsHubListingObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".primary_contact")
	f.SpecField(".documentation")
	f.SpecField(".data_provider")
	f.SpecField(".publisher")
	f.SpecField(".request_access")
	f.SpecField(".discovery_type")
	f.SpecField(".bigquery_dataset")

	f.StatusField(".state")

	f.Unimplemented_NotYetTriaged(".icon")
	f.Unimplemented_NotYetTriaged(".restricted_export_config")
	f.Unimplemented_Identity(".name")

	f.Unimplemented_NotYetTriaged(".allow_only_metadata_sharing")
	f.Unimplemented_NotYetTriaged(".categories")
	f.Unimplemented_NotYetTriaged(".bigquery_dataset.effective_replicas")
	f.Unimplemented_NotYetTriaged(".bigquery_dataset.effective_replicas[].location")
	f.Unimplemented_NotYetTriaged(".bigquery_dataset.effective_replicas[].primary_state")
	f.Unimplemented_NotYetTriaged(".bigquery_dataset.replica_locations")
	f.Unimplemented_NotYetTriaged(".bigquery_dataset.restricted_export_policy.enabled")
	f.Unimplemented_NotYetTriaged(".bigquery_dataset.restricted_export_policy.restrict_direct_table_access")
	f.Unimplemented_NotYetTriaged(".bigquery_dataset.restricted_export_policy.restrict_query_result")
	f.Unimplemented_NotYetTriaged(".bigquery_dataset.selected_resources")
	f.Unimplemented_NotYetTriaged(".bigquery_dataset.selected_resources[].routine")
	f.Unimplemented_NotYetTriaged(".commercial_info")
	f.Unimplemented_NotYetTriaged(".commercial_info.cloud_marketplace")
	f.Unimplemented_NotYetTriaged(".commercial_info.cloud_marketplace.commercial_state")
	f.Unimplemented_NotYetTriaged(".commercial_info.cloud_marketplace.service")
	f.Unimplemented_NotYetTriaged(".log_linked_dataset_query_user_email")
	f.Unimplemented_NotYetTriaged(".pubsub_topic")
	f.Unimplemented_NotYetTriaged(".pubsub_topic.data_affinity_regions")
	f.Unimplemented_NotYetTriaged(".pubsub_topic.topic")
	f.Unimplemented_NotYetTriaged(".resource_type")
	f.Unimplemented_NotYetTriaged(".stored_procedure_config")
	f.Unimplemented_NotYetTriaged(".stored_procedure_config.allowed_stored_procedure_types")
	f.Unimplemented_NotYetTriaged(".stored_procedure_config.enabled")

	f.FilterSpec = func(in *pb.Listing) {
		if in.DiscoveryType != nil && *in.DiscoveryType == pb.DiscoveryType_DISCOVERY_TYPE_UNSPECIFIED {
			in.DiscoveryType = nil
		}
	}

	return f
}
