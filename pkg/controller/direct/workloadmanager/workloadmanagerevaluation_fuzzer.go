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
// proto.message: google.cloud.workloadmanager.v1.Evaluation
// api.group: workloadmanager.cnrm.cloud.google.com

package workloadmanager

import (
	pb "cloud.google.com/go/workloadmanager/apiv1/workloadmanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(workloadManagerEvaluationFuzzer())
}

func workloadManagerEvaluationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Evaluation{},
		WorkloadManagerEvaluationSpec_FromProto, WorkloadManagerEvaluationSpec_ToProto,
		WorkloadManagerEvaluationObservedState_FromProto, WorkloadManagerEvaluationObservedState_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".rule_names")
	f.SpecField(".schedule")
	f.SpecField(".custom_rules_bucket")
	f.SpecField(".evaluation_type")

	f.SpecField(".big_query_destination")
	f.SpecField(".big_query_destination.destination_dataset")
	f.SpecField(".big_query_destination.create_new_results_table")

	f.SpecField(".resource_filter")
	f.SpecField(".resource_filter.resource_id_patterns")
	f.SpecField(".resource_filter.inclusion_labels")
	f.SpecField(".resource_filter.gce_instance_filter")
	f.SpecField(".resource_filter.gce_instance_filter.service_accounts")

	f.StatusField(".resource_status")
	f.StatusField(".resource_status.state")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_LabelsAnnotations(".labels")

	f.Unimplemented_NotYetTriaged(".kms_key")
	f.Unimplemented_NotYetTriaged(".resource_filter.scopes")

	return f
}
