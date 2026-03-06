// Copyright 2024 Google LLC
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

package bigqueryanalyticshub

import (
	bigqueryanalyticshubpb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzDataExchange())
}

func fuzzDataExchange() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&bigqueryanalyticshubpb.DataExchange{},
		BigQueryAnalyticsHubDataExchangeSpec_FromProto, BigQueryAnalyticsHubDataExchangeSpec_ToProto,
		BigQueryAnalyticsHubDataExchangeObservedState_FromProto, BigQueryAnalyticsHubDataExchangeObservedState_ToProto,
	)

	f.UnimplementedFields.Insert([]string{
		".discovery_type", // does not round trip as it is a pointer to a enum value
		".sharing_environment_config",
		".icon",
		".name", // this is the same as our self reference
	}...)

	f.SpecFields.Insert([]string{
		".description",
		".documentation",
		".display_name",
		".primary_contact",
		".discovery_type",
	}...)

	// status fields are only listing_count for now
	f.StatusFields.Insert([]string{
		".listing_count",
	}...)

	return f
}
