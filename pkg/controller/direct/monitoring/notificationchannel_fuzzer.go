// Copyright 2024 Google LLC
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
// crd.kind: MonitoringNotificationChannel

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(notificationChannelFuzzer())
}

func notificationChannelFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.NotificationChannel{},
		MonitoringNotificationChannelSpec_FromProto, MonitoringNotificationChannelSpec_ToProto,
		MonitoringNotificationChannelStatus_FromProto, MonitoringNotificationChannelStatus_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".enabled")
	f.SpecFields.Insert(".type")

	f.StatusFields.Insert(".verification_status")

	// System / naming fields
	f.UnimplementedFields.Insert(".name")

	// Labels
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".user_labels")

	// Volatile status fields we don't want to implement
	f.UnimplementedFields.Insert(".mutation_records")
	f.UnimplementedFields.Insert(".creation_record")

	return f
}
