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
// proto.message: google.cloud.cloudsecuritycompliance.v1.CloudControl
// api.group: cloudsecuritycompliance.cnrm.cloud.google.com

package cloudsecuritycompliancecloudcontrol

import (
	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/cloudsecuritycompliance"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudSecurityComplianceCloudControlFuzzer())
}

func cloudSecurityComplianceCloudControlFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.CloudControl{},
		cloudsecuritycompliance.CloudSecurityComplianceCloudControlSpec_FromProto, cloudsecuritycompliance.CloudSecurityComplianceCloudControlSpec_ToProto,
		cloudsecuritycompliance.CloudSecurityComplianceCloudControlObservedState_FromProto, cloudsecuritycompliance.CloudSecurityComplianceCloudControlObservedState_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".display_name")
	f.SpecField(".parameter_spec")
	f.SpecField(".rules")
	f.SpecField(".severity")
	f.SpecField(".finding_category")
	f.SpecField(".supported_cloud_providers")
	f.SpecField(".remediation_steps")
	f.SpecField(".categories")
	f.SpecField(".supported_target_resource_types")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".parameter_spec[].sub_parameters[].sub_parameters")

	f.StatusField(".major_revision_id")
	f.StatusField(".supported_enforcement_modes")
	f.StatusField(".related_frameworks")
	f.StatusField(".create_time")

	return f
}
