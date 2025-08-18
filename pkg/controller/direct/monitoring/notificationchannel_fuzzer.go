// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
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

	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".enabled")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".verification_status")

	// Identity fields
	f.UnimplementedFields.Insert(".name")

	// Fields that are not in the terraform so we don't want to implement
	{
		f.UnimplementedFields.Insert(".display_name")
		f.UnimplementedFields.Insert(".creation_record")
		f.UnimplementedFields.Insert(".mutation_records")
		f.UnimplementedFields.Insert(".description")
		f.UnimplementedFields.Insert(".user_labels")
		f.UnimplementedFields.Insert(".display_name")
	}

	return f
}
