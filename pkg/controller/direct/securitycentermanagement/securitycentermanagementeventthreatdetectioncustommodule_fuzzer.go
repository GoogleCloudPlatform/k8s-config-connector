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
// proto.message: google.cloud.securitycentermanagement.v1.EventThreatDetectionCustomModule
// api.group: securitycentermanagement.cnrm.cloud.google.com

package securitycentermanagement

import (
	pb "cloud.google.com/go/securitycentermanagement/apiv1/securitycentermanagementpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(securityCenterManagementEventThreatDetectionCustomModuleFuzzer())
}

func securityCenterManagementEventThreatDetectionCustomModuleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.EventThreatDetectionCustomModule{},
		SecurityCenterManagementEventThreatDetectionCustomModuleSpec_FromProto, SecurityCenterManagementEventThreatDetectionCustomModuleSpec_ToProto,
		SecurityCenterManagementEventThreatDetectionCustomModuleObservedState_FromProto, SecurityCenterManagementEventThreatDetectionCustomModuleObservedState_ToProto,
	)

	f.IdentityField(".name")

	// Spec fields:
	f.SpecField(".config")
	f.SpecField(".enablement_state")
	f.SpecField(".type")
	f.SpecField(".display_name")
	f.SpecField(".description")

	// Output only / observed state fields:
	f.StatusField(".ancestor_module")
	f.StatusField(".update_time")
	f.StatusField(".last_editor")

	return f
}
