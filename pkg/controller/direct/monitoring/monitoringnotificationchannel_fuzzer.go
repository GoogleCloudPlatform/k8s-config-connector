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
// proto.message: google.monitoring.v3.NotificationChannel
// api.group: monitoring.cnrm.cloud.google.com

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(monitoringNotificationChannelFuzzer())
}

func monitoringNotificationChannelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NotificationChannel{},
		MonitoringNotificationChannelSpec_FromProto, MonitoringNotificationChannelSpec_ToProto,
		MonitoringNotificationChannelStatus_FromProto, MonitoringNotificationChannelStatus_ToProto,
	)

	// Spec fields
	// .type is mapped to Spec.Type
	f.SpecField(".type")
	// .description is mapped to Spec.Description
	f.SpecField(".description")
	// .labels is mapped to Spec.Labels
	f.SpecField(".labels")
	// .enabled is mapped to Spec.Enabled
	f.SpecField(".enabled")

	// Status fields
	// .name is mapped to Status.Name
	f.StatusField(".name")
	// .verification_status is mapped to Status.VerificationStatus
	f.StatusField(".verification_status")

	// Identity fields
	// .name is used as the resource name/identifier
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not Yet Triaged fields
	f.Unimplemented_NotYetTriaged(".display_name")
	f.Unimplemented_NotYetTriaged(".creation_record")
	f.Unimplemented_NotYetTriaged(".mutation_records")
	f.Unimplemented_NotYetTriaged(".user_labels")

	return f
}
