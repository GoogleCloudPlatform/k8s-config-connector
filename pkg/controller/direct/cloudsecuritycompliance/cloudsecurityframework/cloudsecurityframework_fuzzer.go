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

package cloudsecurityframework

import (
	pb "cloud.google.com/go/cloudsecuritycompliance/apiv1/cloudsecuritycompliancepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/cloudsecuritycompliance"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudSecurityFrameworkFuzzer())
}

func cloudSecurityFrameworkFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Framework{},
		cloudsecuritycompliance.CloudSecurityFrameworkSpec_FromProto,
		cloudsecuritycompliance.CloudSecurityFrameworkSpec_ToProto,
		cloudsecuritycompliance.CloudSecurityFrameworkObservedState_FromProto,
		cloudsecuritycompliance.CloudSecurityFrameworkObservedState_ToProto,
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
	f.StatusField(".supported_enforcement_modes")

	// Identity/Ignored fields
	f.Unimplemented_Identity(".name")

	// Untriaged or deeply nested fields
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].name")
	f.Unimplemented_NotYetTriaged(".cloud_control_details[].parameters")
	f.Unimplemented_NotYetTriaged(".supported_enforcement_modes")

	return f
}
