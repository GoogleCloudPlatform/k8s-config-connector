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
// proto.message: google.api.MetricDescriptor
// krm.group: monitoring.cnrm.cloud.google.com
// krm.kind: MonitoringMetricDescriptor

package monitoring

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(monitoringMetricDescriptorFuzzer())
}

func monitoringMetricDescriptorFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*metricpb.MetricDescriptor, krm.MonitoringMetricDescriptorSpec, krm.MonitoringMetricDescriptorStatus](
		&metricpb.MetricDescriptor{},
		MonitoringMetricDescriptorSpec_FromProto,
		MonitoringMetricDescriptorSpec_ToProto,
		MonitoringMetricDescriptorStatus_FromProto,
		MonitoringMetricDescriptorStatus_ToProto,
	)

	// Field comparison between KRM Spec (MonitoringMetricDescriptorSpec) and GCP Proto (google.api.MetricDescriptor):
	// - ProjectRef (KRM Spec) maps to GCP project URL segment, not part of standard MetricDescriptor fields.
	// - Description (KRM Spec) maps to `.description` (GCP Proto).
	// - DisplayName (KRM Spec) maps to `.display_name` (GCP Proto).
	// - Labels (KRM Spec) maps to `.labels` (GCP Proto).
	// - LaunchStage (KRM Spec) maps to `.launch_stage` (GCP Proto).
	// - Metadata (KRM Spec) maps to `.metadata` (GCP Proto).
	// - MetricKind (KRM Spec) maps to `.metric_kind` (GCP Proto).
	// - Type (KRM Spec) maps to `.type` (GCP Proto).
	// - Unit (KRM Spec) maps to `.unit` (GCP Proto).
	// - ValueType (KRM Spec) maps to `.value_type` (GCP Proto).

	f.SpecField(".description")
	f.SpecField(".display_name")
	f.SpecField(".labels")
	f.SpecField(".launch_stage")
	f.SpecField(".metadata")
	f.SpecField(".metric_kind")
	f.SpecField(".type")
	f.SpecField(".unit")
	f.SpecField(".value_type")

	// Status fields
	// - MonitoredResourceTypes (KRM Status) maps to `.monitored_resource_types` (GCP Proto).
	f.StatusField(".monitored_resource_types")

	// Identity and Unimplemented fields
	// - Name (GCP Proto) maps to Status.SelfLink. It represents the resource's GCP-assigned or constructed identifier.
	f.Unimplemented_Identity(".name")

	// Unimplemented fields on proto that are not exposed/mapped in KRM
	f.Unimplemented_NotYetTriaged(".metadata.time_series_resource_hierarchy_level")

	return f
}
