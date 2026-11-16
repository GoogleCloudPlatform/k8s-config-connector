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
// proto.message: google.cloud.dialogflow.cx.v3.Flow
// api.group: dialogflowcx.cnrm.cloud.google.com

package dialogflowcx

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflowcx/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(dialogflowCXFlowFuzzer())
}

func dialogflowCXFlowFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer[*pb.Flow, krm.DialogflowCXFlowSpec](&pb.Flow{},
		DialogflowCXFlowSpec_FromProto,
		DialogflowCXFlowSpec_ToProto,
	)

	// +proto=google.cloud.dialogflow.cx.v3.Flow
	// +krm=dialogflowcx.cnrm.cloud.google.com/DialogflowCXFlow
	//
	// Field Comparison:
	// | KRM Spec Field                       | Proto Field                                                  | Notes |
	// |--------------------------------------|--------------------------------------------------------------|-------|
	// | parent                               | (none)                                                       | KCC resource reference / hierarchy |
	// | description                          | description                                                  |       |
	// | displayName                          | display_name                                                 |       |
	// | eventHandlers                        | event_handlers                                               |       |
	// | languageCode                         | (none)                                                       | In CreateFlowRequest / UpdateFlowRequest |
	// | nluSettings                          | nlu_settings                                                 |       |
	// | transitionRouteGroups                | transition_route_groups                                      |       |
	// | transitionRoutes                     | transition_routes                                            |       |
	// | resourceID                           | (none)                                                       | KCC acquisition ID |

	// Identity field
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".display_name")
	f.SpecField(".event_handlers[].event")
	f.SpecField(".event_handlers[].name")
	f.SpecField(".event_handlers[].target_flow")
	f.SpecField(".event_handlers[].target_page")
	f.SpecField(".event_handlers[].trigger_fulfillment.conditional_cases[].cases")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].channel")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].conversation_success.metadata")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].live_agent_handoff.metadata")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].output_audio_text.allow_playback_interruption")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].output_audio_text.ssml")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].output_audio_text.text")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].payload")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].play_audio.allow_playback_interruption")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].play_audio.audio_uri")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].telephony_transfer_call.phone_number")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].text.allow_playback_interruption")
	f.SpecField(".event_handlers[].trigger_fulfillment.messages[].text.text")
	f.SpecField(".event_handlers[].trigger_fulfillment.return_partial_responses")
	f.SpecField(".event_handlers[].trigger_fulfillment.set_parameter_actions[].parameter")
	f.SpecField(".event_handlers[].trigger_fulfillment.set_parameter_actions[].value")
	f.SpecField(".event_handlers[].trigger_fulfillment.tag")
	f.SpecField(".event_handlers[].trigger_fulfillment.webhook")
	f.SpecField(".nlu_settings.classification_threshold")
	f.SpecField(".nlu_settings.model_training_mode")
	f.SpecField(".nlu_settings.model_type")
	f.SpecField(".transition_route_groups")
	f.SpecField(".transition_routes[].condition")
	f.SpecField(".transition_routes[].intent")
	f.SpecField(".transition_routes[].name")
	f.SpecField(".transition_routes[].target_flow")
	f.SpecField(".transition_routes[].target_page")
	f.SpecField(".transition_routes[].trigger_fulfillment.conditional_cases[].cases")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].channel")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].conversation_success.metadata")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].live_agent_handoff.metadata")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].output_audio_text.allow_playback_interruption")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].output_audio_text.ssml")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].output_audio_text.text")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].payload")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].play_audio.allow_playback_interruption")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].play_audio.audio_uri")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].telephony_transfer_call.phone_number")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].text.allow_playback_interruption")
	f.SpecField(".transition_routes[].trigger_fulfillment.messages[].text.text")
	f.SpecField(".transition_routes[].trigger_fulfillment.return_partial_responses")
	f.SpecField(".transition_routes[].trigger_fulfillment.set_parameter_actions[].parameter")
	f.SpecField(".transition_routes[].trigger_fulfillment.set_parameter_actions[].value")
	f.SpecField(".transition_routes[].trigger_fulfillment.tag")
	f.SpecField(".transition_routes[].trigger_fulfillment.webhook")

	// Unimplemented / Not yet triaged fields
	f.Unimplemented_NotYetTriaged(".advanced_settings")
	f.Unimplemented_NotYetTriaged(".event_handlers[].target_playbook")
	f.Unimplemented_NotYetTriaged(".event_handlers[].trigger_fulfillment.advanced_settings")
	f.Unimplemented_NotYetTriaged(".event_handlers[].trigger_fulfillment.enable_generative_fallback")
	f.Unimplemented_NotYetTriaged(".event_handlers[].trigger_fulfillment.generators")
	f.Unimplemented_NotYetTriaged(".event_handlers[].trigger_fulfillment.messages[].end_interaction")
	f.Unimplemented_NotYetTriaged(".event_handlers[].trigger_fulfillment.messages[].knowledge_info_card")
	f.Unimplemented_NotYetTriaged(".event_handlers[].trigger_fulfillment.messages[].mixed_audio")
	f.Unimplemented_NotYetTriaged(".event_handlers[].trigger_fulfillment.messages[].response_type")
	f.Unimplemented_NotYetTriaged(".event_handlers[].trigger_fulfillment.messages[].tool_call")
	f.Unimplemented_NotYetTriaged(".input_parameter_definitions")
	f.Unimplemented_NotYetTriaged(".knowledge_connector_settings")
	f.Unimplemented_NotYetTriaged(".locked")
	f.Unimplemented_NotYetTriaged(".multi_language_settings")
	f.Unimplemented_NotYetTriaged(".output_parameter_definitions")
	f.Unimplemented_NotYetTriaged(".transition_routes[].description")
	f.Unimplemented_NotYetTriaged(".transition_routes[].trigger_fulfillment.advanced_settings")
	f.Unimplemented_NotYetTriaged(".transition_routes[].trigger_fulfillment.enable_generative_fallback")
	f.Unimplemented_NotYetTriaged(".transition_routes[].trigger_fulfillment.generators")
	f.Unimplemented_NotYetTriaged(".transition_routes[].trigger_fulfillment.messages[].end_interaction")
	f.Unimplemented_NotYetTriaged(".transition_routes[].trigger_fulfillment.messages[].knowledge_info_card")
	f.Unimplemented_NotYetTriaged(".transition_routes[].trigger_fulfillment.messages[].mixed_audio")
	f.Unimplemented_NotYetTriaged(".transition_routes[].trigger_fulfillment.messages[].response_type")
	f.Unimplemented_NotYetTriaged(".transition_routes[].trigger_fulfillment.messages[].tool_call")

	f.FilterSpec = func(in *pb.Flow) {
		for _, eh := range in.EventHandlers {
			if eh != nil && eh.TriggerFulfillment != nil {
				normalizeFulfillment(eh.TriggerFulfillment)
			}
		}
		for _, tr := range in.TransitionRoutes {
			if tr != nil && tr.TriggerFulfillment != nil {
				normalizeFulfillment(tr.TriggerFulfillment)
			}
		}
	}

	return f
}
