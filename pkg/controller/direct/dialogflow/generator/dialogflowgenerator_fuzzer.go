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
// proto.message: google.cloud.dialogflow.v2.Generator
// api.group: dialogflow.cnrm.cloud.google.com

package generator

import (
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dialogflowGeneratorFuzzer())
}

func dialogflowGeneratorFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Generator{},
		DialogflowGeneratorSpec_FromProto,
		DialogflowGeneratorSpec_ToProto,
		DialogflowGeneratorObservedState_FromProto,
		DialogflowGeneratorObservedState_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".free_form_context")
	f.SpecField(".summarization_context")
	f.SpecField(".inference_parameter")
	f.SpecField(".trigger_event")
	f.SpecField(".published_model")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	f.Unimplemented_Identity(".name")

	f.Unimplemented_NotYetTriaged(".agent_coaching_context")
	f.Unimplemented_NotYetTriaged(".tools")
	f.Unimplemented_NotYetTriaged(".suggestion_deduping_config")
	f.Unimplemented_NotYetTriaged(".toolset_tools")
	f.Unimplemented_NotYetTriaged(".ces_tool_specs")
	f.Unimplemented_NotYetTriaged(".ces_app_specs")
	f.Unimplemented_NotYetTriaged(".summarization_context.few_shot_examples[].output.agent_coaching_suggestion")
	f.Unimplemented_NotYetTriaged(".summarization_context.few_shot_examples[].output.tool_call_info")

	return f
}
