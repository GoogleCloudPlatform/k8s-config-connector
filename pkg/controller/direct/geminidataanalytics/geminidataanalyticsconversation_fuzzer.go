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
// proto.message: google.cloud.geminidataanalytics.v1beta.Conversation
// api.group: geminidataanalytics.cnrm.cloud.google.com

package geminidataanalytics

import (
	pb "cloud.google.com/go/geminidataanalytics/apiv1beta/geminidataanalyticspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(geminiDataAnalyticsConversationFuzzer())
}

func geminiDataAnalyticsConversationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Conversation{},
		GeminiDataAnalyticsConversationSpec_FromProto, GeminiDataAnalyticsConversationSpec_ToProto,
		GeminiDataAnalyticsConversationObservedState_FromProto, GeminiDataAnalyticsConversationObservedState_ToProto,
	)

	f.IdentityField(".name")

	f.SpecField(".agents")
	f.SpecField(".labels")

	f.StatusField(".create_time")
	f.StatusField(".last_used_time")

	f.Unimplemented_NotYetTriaged(".kms_key")
	f.Unimplemented_NotYetTriaged(".memory_paused")

	return f
}
