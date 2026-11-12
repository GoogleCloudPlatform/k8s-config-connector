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
// proto.message: google.cloud.dialogflow.cx.v3.Agent
// api.group: dialogflowcx.cnrm.cloud.google.com

package dialogflowcx

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dialogflowCXAgentFuzzer())
}

func dialogflowCXAgentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Agent{},
		DialogflowCXAgentSpec_FromProto, DialogflowCXAgentSpec_ToProto,
		DialogflowCXAgentObservedState_FromProto, DialogflowCXAgentObservedState_ToProto,
	)

	// Spec fields
	f.SpecField(".avatar_uri")
	f.SpecField(".default_language_code")
	f.SpecField(".description")
	f.SpecField(".display_name")
	f.SpecField(".enable_spell_correction")
	f.SpecField(".enable_stackdriver_logging")
	f.SpecField(".security_settings")
	f.SpecField(".speech_to_text_settings")
	f.SpecField(".speech_to_text_settings.enable_speech_adaptation")
	f.SpecField(".supported_language_codes")
	f.SpecField(".time_zone")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Not yet triaged fields
	f.Unimplemented_NotYetTriaged(".advanced_settings")
	f.Unimplemented_NotYetTriaged(".answer_feedback_settings")
	f.Unimplemented_NotYetTriaged(".client_certificate_settings")
	f.Unimplemented_NotYetTriaged(".enable_multi_language_training")
	f.Unimplemented_NotYetTriaged(".gen_app_builder_settings")
	f.Unimplemented_NotYetTriaged(".git_integration_settings")
	f.Unimplemented_NotYetTriaged(".locked")
	f.Unimplemented_NotYetTriaged(".personalization_settings")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".start_flow")
	f.Unimplemented_NotYetTriaged(".start_playbook")
	f.Unimplemented_NotYetTriaged(".text_to_speech_settings")

	return f
}
