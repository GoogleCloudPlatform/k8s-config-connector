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
// proto.message: google.cloud.contactcenterinsights.v1.Conversation
// api.group: contactcenterinsights.cnrm.cloud.google.com

package contactcenterinsights

import (
	pb "cloud.google.com/go/contactcenterinsights/apiv1/contactcenterinsightspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(conversationFuzzer())
}

func conversationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Conversation{},
		CCInsightsConversationSpec_FromProto, CCInsightsConversationSpec_ToProto,
		CCInsightsConversationObservedState_FromProto, CCInsightsConversationObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".dialogflow_intents")
	f.Unimplemented_NotYetTriaged(".latest_analysis")
	f.Unimplemented_NotYetTriaged(".runtime_annotations")
	f.Unimplemented_NotYetTriaged(".data_source.dialogflow_source.dialogflow_conversation")
	f.Unimplemented_NotYetTriaged(".latest_summary.conversation_model")

	// Spec fields
	f.SpecField(".call_metadata")
	f.SpecField(".expire_time")
	f.SpecField(".ttl")
	f.SpecField(".data_source")
	f.SpecField(".start_time")
	f.SpecField(".language_code")
	f.SpecField(".agent_id")
	f.SpecField(".labels")
	f.SpecField(".quality_metadata")
	f.SpecField(".metadata_json")
	f.SpecField(".medium")
	f.SpecField(".obfuscated_user_id")

	// Status/ObservedState fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".transcript")
	f.StatusField(".duration")
	f.StatusField(".turn_count")
	f.StatusField(".latest_summary")

	return f
}
