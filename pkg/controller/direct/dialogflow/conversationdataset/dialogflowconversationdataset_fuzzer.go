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
// proto.message: google.cloud.dialogflow.v2.ConversationDataset
// api.group: dialogflow.cnrm.cloud.google.com

package conversationdataset

import (
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dialogflowConversationDatasetFuzzer())
}

func dialogflowConversationDatasetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConversationDataset{},
		DialogflowConversationDatasetSpec_FromProto,
		DialogflowConversationDatasetSpec_ToProto,
		DialogflowConversationDatasetObservedState_FromProto,
		DialogflowConversationDatasetObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")

	f.StatusField(".create_time")
	f.StatusField(".input_config")
	f.StatusField(".conversation_info")
	f.StatusField(".conversation_count")
	f.StatusField(".satisfies_pzi")
	f.StatusField(".satisfies_pzs")

	f.Unimplemented_Identity(".name")

	return f
}
