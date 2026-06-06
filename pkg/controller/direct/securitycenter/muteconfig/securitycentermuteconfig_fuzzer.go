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
// proto.message: google.cloud.securitycenter.v1.MuteConfig
// api.group: securitycenter.cnrm.cloud.google.com

package muteconfig

import (
	securitycenterpb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/securitycenter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(securityCenterMuteConfigFuzzer())
}

func securityCenterMuteConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&securitycenterpb.MuteConfig{},
		securitycenter.SecurityCenterMuteConfigSpec_FromProto, securitycenter.SecurityCenterMuteConfigSpec_ToProto,
		securitycenter.SecurityCenterMuteConfigObservedState_FromProto, securitycenter.SecurityCenterMuteConfigObservedState_ToProto,
	)

	f.IdentityField(".name")

	// Spec fields:
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".filter")
	f.SpecFields.Insert(".type")
	f.SpecFields.Insert(".expiry_time")

	// Output only / observed state fields:
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".most_recent_editor")

	// Unimplemented / Deprecated fields:
	f.UnimplementedFields.Insert(".display_name")

	return f
}
