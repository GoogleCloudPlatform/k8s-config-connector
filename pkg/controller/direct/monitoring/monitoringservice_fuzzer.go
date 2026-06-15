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
// proto.message: google.monitoring.v3.Service
// api.group: monitoring.cnrm.cloud.google.com

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(monitoringServiceFuzzer())
}

func monitoringServiceFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer[*pb.Service, krm.MonitoringServiceSpec](&pb.Service{},
		MonitoringServiceSpec_FromProto, MonitoringServiceSpec_ToProto,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".telemetry")

	// Identity fields
	f.Unimplemented_Identity(".name")

	// Oneof and unsupported fields
	f.Unimplemented_NotYetTriaged(".custom")
	f.Unimplemented_NotYetTriaged(".app_engine")
	f.Unimplemented_NotYetTriaged(".cloud_endpoints")
	f.Unimplemented_NotYetTriaged(".cluster_istio")
	f.Unimplemented_NotYetTriaged(".mesh_istio")
	f.Unimplemented_NotYetTriaged(".istio_canonical_service")
	f.Unimplemented_NotYetTriaged(".cloud_run")
	f.Unimplemented_NotYetTriaged(".gke_namespace")
	f.Unimplemented_NotYetTriaged(".gke_workload")
	f.Unimplemented_NotYetTriaged(".gke_service")
	f.Unimplemented_NotYetTriaged(".basic_service")
	f.Unimplemented_NotYetTriaged(".user_labels")

	return f
}
