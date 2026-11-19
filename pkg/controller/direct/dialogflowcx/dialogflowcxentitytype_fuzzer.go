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
// proto.message: google.cloud.dialogflow.cx.v3.EntityType
// api.group: dialogflowcx.cnrm.cloud.google.com

package dialogflowcx

import (
	pb "cloud.google.com/go/dialogflow/cx/apiv3/cxpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dialogflowCXEntityTypeFuzzer())
}

func dialogflowCXEntityTypeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.EntityType{},
		DialogflowCXEntityTypeSpec_FromProto, DialogflowCXEntityTypeSpec_ToProto,
		DialogflowCXEntityTypeObservedState_FromProto, DialogflowCXEntityTypeObservedState_ToProto,
	)

	// +----------------------------------+------------------------------------+-------------------------------------+
	// | KRM Spec / Status Field          | Proto Field                        | Fuzzer/Mapping Status               |
	// +----------------------------------+------------------------------------+-------------------------------------+
	// | spec.autoExpansionMode           | auto_expansion_mode                | SpecField (".auto_expansion_mode")  |
	// | spec.displayName                 | display_name                       | SpecField (".display_name")         |
	// | spec.enableFuzzyExtraction       | enable_fuzzy_extraction            | SpecField (".enable_fuzzy_extraction") |
	// | spec.entities                    | entities                           | SpecField (".entities")             |
	// | spec.excludedPhrases             | excluded_phrases                   | SpecField (".excluded_phrases")     |
	// | spec.kind                        | kind                               | SpecField (".kind")                 |
	// | spec.languageCode                | N/A                                | API parameter, not on proto         |
	// | spec.parent                      | N/A                                | Parent URL, not on proto            |
	// | spec.redact                      | redact                             | SpecField (".redact")               |
	// | spec.resourceID                  | name                               | Unimplemented_Identity (".name")    |
	// | status.name                      | name                               | Unimplemented_Identity (".name")    |
	// +----------------------------------+------------------------------------+-------------------------------------+

	f.Unimplemented_Identity(".name")

	f.SpecField(".auto_expansion_mode")
	f.SpecField(".display_name")
	f.SpecField(".enable_fuzzy_extraction")
	f.SpecField(".entities")
	f.SpecField(".excluded_phrases")
	f.SpecField(".kind")
	f.SpecField(".redact")

	return f
}
