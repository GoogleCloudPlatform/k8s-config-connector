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
// proto.message: google.monitoring.v3.UptimeCheckConfig
// api.group: monitoring.cnrm.cloud.google.com

package monitoring

import (
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(monitoringUptimeCheckConfigFuzzer())
}

func monitoringUptimeCheckConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.UptimeCheckConfig{},
		MonitoringUptimeCheckConfigSpec_FromProto, MonitoringUptimeCheckConfigSpec_ToProto,
		MonitoringUptimeCheckConfigStatus_FromProto, MonitoringUptimeCheckConfigStatus_ToProto,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".period")
	f.SpecField(".timeout")
	f.SpecField(".content_matchers")
	f.SpecField(".selected_regions")

	// Resource (oneof)
	f.SpecField(".monitored_resource")
	f.SpecField(".resource_group")

	// CheckRequestType (oneof)
	f.SpecField(".http_check")
	f.SpecField(".tcp_check")

	// Unimplemented / ignored top-level fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".checker_type")
	f.Unimplemented_NotYetTriaged(".is_internal")
	f.Unimplemented_NotYetTriaged(".internal_checkers")
	f.Unimplemented_NotYetTriaged(".user_labels")
	f.Unimplemented_NotYetTriaged(".synthetic_monitor")

	// Unimplemented nested fields in HTTPCheck
	f.Unimplemented_NotYetTriaged(".http_check.accepted_response_status_codes")
	f.Unimplemented_NotYetTriaged(".http_check.ping_config")
	f.Unimplemented_NotYetTriaged(".http_check.service_agent_authentication")
	f.Unimplemented_NotYetTriaged(".http_check.custom_content_type")
	f.Unimplemented_NotYetTriaged(".http_check.auth_info.password")

	// Unimplemented nested fields in TCPCheck
	f.Unimplemented_NotYetTriaged(".tcp_check.ping_config")

	// Unimplemented nested fields in ContentMatchers
	f.Unimplemented_NotYetTriaged(".content_matchers[].json_path_matcher")

	return f
}
