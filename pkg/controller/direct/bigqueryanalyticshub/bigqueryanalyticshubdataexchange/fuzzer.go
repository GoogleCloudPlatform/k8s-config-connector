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
// proto.message: google.cloud.bigquery.analyticshub.v1.DataExchange
// api.group: bigqueryanalyticshub.cnrm.cloud.google.com

package bigqueryanalyticshubdataexchange

import (
	pb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(BigQueryAnalyticsHubDataExchangeFuzzer())
}

func BigQueryAnalyticsHubDataExchangeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DataExchange{},
		BigQueryAnalyticsHubDataExchangeSpec_FromProto, BigQueryAnalyticsHubDataExchangeSpec_ToProto,
		BigQueryAnalyticsHubDataExchangeObservedState_FromProto, BigQueryAnalyticsHubDataExchangeObservedState_ToProto,
	)

	f.FilterSpec = func(in *pb.DataExchange) {
		if in.DiscoveryType != nil && *in.DiscoveryType == pb.DiscoveryType_DISCOVERY_TYPE_UNSPECIFIED {
			in.DiscoveryType = nil
		}
	}

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".primary_contact")
	f.SpecField(".documentation")
	f.SpecField(".discovery_type")

	f.StatusField(".listing_count")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".icon")
	f.Unimplemented_NotYetTriaged(".sharing_environment_config")
	f.Unimplemented_NotYetTriaged(".log_linked_dataset_query_user_email")

	return f
}
