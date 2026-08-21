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
// proto.message: google.cloud.cloudsecuritycompliance.v1.Framework
// api.group: cloudsecuritycompliance.cnrm.cloud.google.com

package cloudsecuritycomplianceframework

import (
	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/cloudsecuritycompliance"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudSecurityComplianceFrameworkFuzzer())
}

func cloudSecurityComplianceFrameworkFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Framework{},
		cloudsecuritycompliance.CloudSecurityComplianceFrameworkSpec_FromProto,
		cloudsecuritycompliance.CloudSecurityComplianceFrameworkSpec_ToProto,
		cloudsecuritycompliance.CloudSecurityComplianceFrameworkObservedState_FromProto,
		cloudsecuritycompliance.CloudSecurityComplianceFrameworkObservedState_ToProto,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".cloud_control_details")
	f.SpecField(".category")

	// Status / Observed state fields
	f.StatusField(".major_revision_id")
	f.StatusField(".type")
	f.StatusField(".supported_cloud_providers")
	f.StatusField(".supported_target_resource_types")

	// Identity/Ignored fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / recursive / untriaged fields
	f.Unimplemented_NotYetTriaged(".supported_enforcement_modes")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].name")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.bool_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.string_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.string_list_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.number_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.oneof_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.oneof_value.parameter_value.bool_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.oneof_value.parameter_value.string_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.oneof_value.parameter_value.string_list_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.oneof_value.parameter_value.number_value")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters[].parameter_value.oneof_value.parameter_value.oneof_value")

	return f
}
