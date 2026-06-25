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
// proto.message: google.cloud.cloudsecuritycompliance.v1.FrameworkDeployment
// api.group: cloudsecuritycompliance.cnrm.cloud.google.com

package cloudsecuritycompliance

import (
	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudSecurityComplianceFrameworkDeploymentFuzzer())
}

func cloudSecurityComplianceFrameworkDeploymentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FrameworkDeployment{},
		CloudSecurityComplianceFrameworkDeploymentSpec_FromProto, CloudSecurityComplianceFrameworkDeploymentSpec_ToProto,
		CloudSecurityComplianceFrameworkDeploymentObservedState_FromProto, CloudSecurityComplianceFrameworkDeploymentObservedState_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".framework")
	f.SpecField(".target_resource_config")
	f.SpecField(".cloud_control_metadata")
	f.SpecField(".etag")

	f.Unimplemented_NotYetTriaged(".cloud_control_metadata[].cloud_control_details.parameters[].parameter_value.string_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_metadata[].cloud_control_details.parameters[].parameter_value.bool_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_metadata[].cloud_control_details.parameters[].parameter_value.string_list_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_metadata[].cloud_control_details.parameters[].parameter_value.number_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_metadata[].cloud_control_details.parameters[].parameter_value.oneof_value")

	f.StatusField(".computed_target_resource")
	f.StatusField(".deployment_state")
	f.StatusField(".cc_deployments")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".cc_group_deployments")
	f.StatusField(".target_resource_display_name")
	f.StatusField(".cloud_control_deployment_references")

	f.Unimplemented_Identity(".name")

	return f
}
