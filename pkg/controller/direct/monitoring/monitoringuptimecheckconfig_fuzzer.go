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
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".period")
	f.SpecFields.Insert(".timeout")
	f.SpecFields.Insert(".content_matchers")
	f.SpecFields.Insert(".selected_regions")

	// Resource (oneof)
	f.SpecFields.Insert(".monitored_resource")
	f.SpecFields.Insert(".resource_group")

	// CheckRequestType (oneof)
	f.SpecFields.Insert(".http_check")
	f.SpecFields.Insert(".tcp_check")

	// Unimplemented / ignored top-level fields
	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".checker_type")
	f.UnimplementedFields.Insert(".is_internal")
	f.UnimplementedFields.Insert(".internal_checkers")
	f.UnimplementedFields.Insert(".user_labels")
	f.UnimplementedFields.Insert(".synthetic_monitor")

	// Unimplemented nested fields in HTTPCheck
	f.UnimplementedFields.Insert(".http_check.accepted_response_status_codes")
	f.UnimplementedFields.Insert(".http_check.ping_config")
	f.UnimplementedFields.Insert(".http_check.service_agent_authentication")
	f.UnimplementedFields.Insert(".http_check.custom_content_type")
	f.UnimplementedFields.Insert(".http_check.auth_info.password")

	// Unimplemented nested fields in TCPCheck
	f.UnimplementedFields.Insert(".tcp_check.ping_config")

	// Unimplemented nested fields in ContentMatchers
	f.UnimplementedFields.Insert(".content_matchers[].json_path_matcher")

	return f
}
