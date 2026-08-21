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
// proto.message: google.monitoring.v3.AlertPolicy
// api.group: monitoring.cnrm.cloud.google.com

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(monitoringAlertPolicyFuzzer())
}

func monitoringAlertPolicyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.AlertPolicy{},
		MonitoringAlertPolicySpec_FromProto, MonitoringAlertPolicySpec_ToProto,
		MonitoringAlertPolicyStatus_FromProto, MonitoringAlertPolicyStatus_ToProto,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".documentation")
	f.SpecField(".combiner")
	f.SpecField(".conditions")
	f.SpecField(".notification_channels")
	f.SpecField(".enabled")
	f.SpecField(".alert_strategy")
	f.SpecField(".severity")

	// Status fields
	f.StatusField(".creation_record")

	// Unimplemented / identity fields
	f.Unimplemented_Identity(".name")

	// Unimplemented fields on proto that are not exposed/mapped in KRM
	f.Unimplemented_Internal(".valid")
	f.Unimplemented_Internal(".mutation_record")
	f.Unimplemented_Internal(".user_labels")
	f.Unimplemented_Internal(".validity")
	f.Unimplemented_Internal(".documentation.subject")
	f.Unimplemented_Internal(".documentation.links")
	f.Unimplemented_Internal(".alert_strategy.notification_prompts")
	f.Unimplemented_Internal(".conditions[].condition_prometheus_query_language.disable_metric_validation")

	// Unimplemented SQL conditions nested fields
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.query")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.minutes")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.minutes.periodicity")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.hourly")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.hourly.periodicity")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.hourly.minute_offset")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.daily")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.daily.periodicity")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.daily.execution_time")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.daily.execution_time.hours")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.daily.execution_time.minutes")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.daily.execution_time.seconds")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.daily.execution_time.nanos")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.boolean_test")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.boolean_test.column")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.row_count_test")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.row_count_test.comparison")
	f.Unimplemented_NotYetTriaged(".conditions[].condition_sql.row_count_test.threshold")

	return f
}
