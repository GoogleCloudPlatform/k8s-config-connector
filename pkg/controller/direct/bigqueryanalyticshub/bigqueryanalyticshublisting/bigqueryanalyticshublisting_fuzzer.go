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
	f.SpecField(".categories")
	f.SpecField(".publisher")
	f.SpecField(".request_access")
	f.SpecField(".discovery_type")
	f.SpecField(".bigquery_dataset")

	f.StatusField(".state")

	f.Unimplemented_NotYetTriaged(".icon")
	f.Unimplemented_NotYetTriaged(".restricted_export_config")
	f.Unimplemented_Identity(".name")

	return f
}
