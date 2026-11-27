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
// proto.message: google.cloud.dialogflow.cx.v3.Intent
// api.group: dialogflowcx.cnrm.cloud.google.com

package dialogflowcx

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dialogflowCXIntentFuzzer())
}

func dialogflowCXIntentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Intent{},
		DialogflowCXIntentSpec_FromProto, DialogflowCXIntentSpec_ToProto,
		DialogflowCXIntentObservedState_FromProto, DialogflowCXIntentObservedState_ToProto,
	)

	// +------------------------------------------+------------------------------------------+-------------------------------------------------------+
	// | KRM Spec / Status Field                  | Proto Field                              | Fuzzer/Mapping Status                                 |
	// +------------------------------------------+------------------------------------------+-------------------------------------------------------+
	// | spec.description                         | description                              | SpecField (".description")                            |
	// | spec.displayName                         | display_name                             | SpecField (".display_name")                           |
	// | spec.isFallback                          | is_fallback                              | SpecField (".is_fallback")                            |
	// | spec.languageCode                        | N/A                                      | API parameter, not on proto                           |
	// | spec.parameters                          | parameters                               | SpecField (".parameters")                             |
	// | spec.parameters[].entityType             | parameters[].entity_type                 | SpecField (".parameters[].entity_type")               |
	// | spec.parameters[].id                     | parameters[].id                          | SpecField (".parameters[].id")                        |
	// | spec.parameters[].isList                 | parameters[].is_list                     | SpecField (".parameters[].is_list")                   |
	// | spec.parameters[].redact                 | parameters[].redact                      | SpecField (".parameters[].redact")                    |
	// | spec.parent                              | N/A                                      | Parent URL, not on proto                              |
	// | spec.priority                            | priority                                 | SpecField (".priority")                               |
	// | spec.resourceID                          | name                                     | Unimplemented_Identity (".name")                      |
	// | spec.trainingPhrases                     | training_phrases                         | SpecField (".training_phrases")                       |
	// | spec.trainingPhrases[].id                 | training_phrases[].id                    | SpecField (".training_phrases[].id")                  |
	// | spec.trainingPhrases[].parts              | training_phrases[].parts                 | SpecField (".training_phrases[].parts")               |
	// | spec.trainingPhrases[].parts[].parameterId| training_phrases[].parts[].parameter_id  | SpecField (".training_phrases[].parts[].parameter_id")|
	// | spec.trainingPhrases[].parts[].text       | training_phrases[].parts[].text         | SpecField (".training_phrases[].parts[].text")        |
	// | spec.trainingPhrases[].repeatCount        | training_phrases[].repeat_count          | SpecField (".training_phrases[].repeat_count")         |
	// | status.conditions                        | N/A                                      | K8s condition, not on proto                           |
	// | status.externalRef                       | N/A                                      | KCC external ref, not on proto                        |
	// | status.name                              | name                                     | Unimplemented_Identity (".name")                      |
	// | status.observedGeneration                | N/A                                      | K8s observed gen, not on proto                        |
	// | N/A                                      | dtmf_pattern                             | Unimplemented_NotYetTriaged (".dtmf_pattern")         |
	// | N/A                                      | labels                                   | Unimplemented_NotYetTriaged (".labels")               |
	// +------------------------------------------+------------------------------------------+-------------------------------------------------------+

	f.Unimplemented_Identity(".name")

	f.SpecField(".description")
	f.SpecField(".display_name")
	f.SpecField(".is_fallback")
	f.SpecField(".parameters")
	f.SpecField(".parameters[].entity_type")
	f.SpecField(".parameters[].id")
	f.SpecField(".parameters[].is_list")
	f.SpecField(".parameters[].redact")
	f.SpecField(".priority")
	f.SpecField(".training_phrases")
	f.SpecField(".training_phrases[].id")
	f.SpecField(".training_phrases[].parts")
	f.SpecField(".training_phrases[].parts[].parameter_id")
	f.SpecField(".training_phrases[].parts[].text")
	f.SpecField(".training_phrases[].repeat_count")

	f.Unimplemented_NotYetTriaged(".dtmf_pattern")
	f.Unimplemented_NotYetTriaged(".labels")

	return f
}
