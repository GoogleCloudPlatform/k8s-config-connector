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
// proto.message: google.cloud.dataplex.v1.DataScan
// api.group: dataplex.cnrm.cloud.google.com

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataplexDataScanFuzzer())
}

func dataplexDataScanFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DataScan{},
		DataplexDataScanSpec_FromProto, DataplexDataScanSpec_ToProto,
		DataplexDataScanObservedState_FromProto, DataplexDataScanObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".data")
	f.SpecField(".execution_spec")
	f.SpecField(".data_quality_spec")
	f.SpecField(".data_profile_spec")
	f.SpecField(".data_discovery_spec")

	f.StatusField(".uid")
	f.StatusField(".state")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".execution_status")
	f.StatusField(".type")
	f.StatusField(".data_quality_result")
	f.StatusField(".data_profile_result")
	f.StatusField(".data_discovery_result")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")
	f.Unimplemented_NotYetTriaged(".data_documentation_spec")
	f.Unimplemented_NotYetTriaged(".data_documentation_result")
	f.Unimplemented_NotYetTriaged(".data_quality_spec.rules[].debug_queries")
	f.Unimplemented_NotYetTriaged(".data_quality_spec.catalog_publishing_enabled")
	f.Unimplemented_NotYetTriaged(".data_profile_result.catalog_publishing_status")
	f.Unimplemented_NotYetTriaged(".data_profile_result.post_scan_actions_result.bigquery_export_result")
	f.Unimplemented_NotYetTriaged(".data_discovery_result.bigquery_publishing.location")
	f.Unimplemented_NotYetTriaged(".data_profile_spec.catalog_publishing_enabled")
	f.Unimplemented_NotYetTriaged(".execution_spec.trigger.one_time")
	f.Unimplemented_NotYetTriaged(".data_quality_result.dimensions[].score")
	f.Unimplemented_NotYetTriaged(".data_discovery_spec.bigquery_publishing_config.location")
	f.Unimplemented_NotYetTriaged(".data_discovery_spec.bigquery_publishing_config.project")
	f.Unimplemented_NotYetTriaged(".data_quality_result.anomaly_detection_generated_assets")
	f.Unimplemented_NotYetTriaged(".data_quality_result.columns[].column")
	f.Unimplemented_NotYetTriaged(".data_quality_result.dimensions[].dimension")
	f.Unimplemented_NotYetTriaged(".data_quality_result.catalog_publishing_status")
	f.Unimplemented_NotYetTriaged(".data_quality_result.rules[].rule")
	f.Unimplemented_NotYetTriaged(".data_quality_result.rules[].passed")
	f.Unimplemented_NotYetTriaged(".data_quality_result.rules[].null_count")
	f.Unimplemented_NotYetTriaged(".data_quality_result.rules[].pass_ratio")
	f.Unimplemented_NotYetTriaged(".data_quality_result.rules[].failing_rows_query")
	f.Unimplemented_NotYetTriaged(".data_quality_result.rules[].assertion_row_count")
	f.Unimplemented_NotYetTriaged(".data_quality_result.rules[].evaluated_count")
	f.Unimplemented_NotYetTriaged(".data_quality_result.rules[].passed_count")
	f.Unimplemented_NotYetTriaged(".data_quality_result.rules[].debug_queries_result_sets")
	f.Unimplemented_NotYetTriaged(".data_quality_result.columns[].dimensions")
	f.Unimplemented_NotYetTriaged(".data_quality_result.columns[].score")
	f.Unimplemented_NotYetTriaged(".data_discovery_result.bigquery_publishing.dataset")
	f.Unimplemented_NotYetTriaged(".data_quality_result.dimensions[].passed")
	f.Unimplemented_NotYetTriaged(".data_quality_result.post_scan_actions_result.bigquery_export_result")
	f.Unimplemented_NotYetTriaged(".data_quality_result.columns[].passed")

	return f
}
