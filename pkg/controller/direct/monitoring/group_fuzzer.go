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
// proto.message: google.monitoring.v3.Group
// api.group: monitoring.cnrm.cloud.google.com

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/monitoring/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(monitoringGroupFuzzer())
}

func monitoringGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer[*pb.Group, krm.MonitoringGroupSpec, krm.MonitoringGroupStatus](&pb.Group{},
		MonitoringGroupSpec_FromProto, MonitoringGroupSpec_ToProto,
		nil, nil,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".filter")
	f.SpecField(".is_cluster")
	f.SpecField(".parent_name")

	// Identity fields
	f.Unimplemented_Identity(".name")

	return f
}
