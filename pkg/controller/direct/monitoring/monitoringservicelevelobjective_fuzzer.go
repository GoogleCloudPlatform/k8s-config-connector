// Copyright 2026 Google LLC
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
// proto.message: google.monitoring.v3.ServiceLevelObjective
// api.group: monitoring.cnrm.cloud.google.com

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	calendarperiodpb "google.golang.org/genproto/googleapis/type/calendarperiod"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(monitoringServiceLevelObjectiveFuzzer())
}

func monitoringServiceLevelObjectiveFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.ServiceLevelObjective{},
		MonitoringServiceLevelObjectiveSpec_FromProto, MonitoringServiceLevelObjectiveSpec_ToProto,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".service_level_indicator")
	f.SpecField(".goal")
	f.SpecField(".rolling_period")
	f.SpecField(".calendar_period")

	// Identity/URL fields
	f.Unimplemented_Identity(".name")

	// Unimplemented fields
	f.Unimplemented_NotYetTriaged(".user_labels")

	f.FilterSpec = func(in *pb.ServiceLevelObjective) {
		if calendarPeriod, ok := in.Period.(*pb.ServiceLevelObjective_CalendarPeriod); ok {
			if calendarPeriod.CalendarPeriod == calendarperiodpb.CalendarPeriod_CALENDAR_PERIOD_UNSPECIFIED {
				in.Period = nil
			}
		}
	}

	return f
}
