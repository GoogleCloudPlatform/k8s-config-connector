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
// proto.message: google.monitoring.metricsscope.v1.MonitoredProject
// krm.group: monitoring.cnrm.cloud.google.com
// krm.kind: MonitoringMonitoredProject

package monitoring

import (
	metricsscopepb "cloud.google.com/go/monitoring/metricsscope/apiv1/metricsscopepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(monitoringMonitoredProjectFuzzer())
}

func monitoringMonitoredProjectFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*metricsscopepb.MonitoredProject, krm.MonitoringMonitoredProjectSpec, krm.MonitoringMonitoredProjectStatus](
		&metricsscopepb.MonitoredProject{},
		MonitoringMonitoredProjectSpec_FromProto,
		MonitoringMonitoredProjectSpec_ToProto,
		MonitoringMonitoredProjectStatus_FromProto,
		MonitoringMonitoredProjectStatus_ToProto,
	)

	// Field comparison between KRM Spec (MonitoringMonitoredProjectSpec) and GCP Proto (google.monitoring.metricsscope.v1.MonitoredProject):
	// - MetricsScope (KRM Spec) and ResourceID (KRM Spec) are combined into `.name` in Proto representing
	//   "locations/global/metricsScopes/{metrics_scope}/projects/{project}".
	// - CreateTime (KRM Status) maps to `.create_time` (GCP Proto).

	// Status fields
	f.StatusField(".create_time")

	// Identity and Unimplemented fields
	// - Name (GCP Proto) contains both the parent metrics scope and the project ID.
	f.Unimplemented_Identity(".name")

	return f
}
