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
// proto.message: google.cloud.dialogflow.v2.Agent
// api.group: dialogflow.cnrm.cloud.google.com

package dialogflow

import (
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dialogflow/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(dialogflowAgentFuzzer())
}

func dialogflowAgentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer[*pb.Agent, krm.DialogflowAgentSpec](&pb.Agent{},
		DialogflowAgentSpec_FromProto,
		DialogflowAgentSpec_ToProto,
	)

	// +proto=google.cloud.dialogflow.v2.Agent
	// +krm=dialogflow.cnrm.cloud.google.com/DialogflowAgent
	//
	// Field Comparison:
	// | KRM Spec Field          | Proto Field               | Notes                       |
	// |-------------------------|---------------------------|-----------------------------|
	// | apiVersion              | api_version               |                             |
	// | avatarUri               | avatar_uri                |                             |
	// | classificationThreshold | classification_threshold  |                             |
	// | defaultLanguageCode     | default_language_code     |                             |
	// | description             | description               |                             |
	// | displayName             | display_name              |                             |
	// | enableLogging           | enable_logging            |                             |
	// | matchMode               | match_mode                |                             |
	// | resourceID              |                           | Not in proto (KRM identity) |
	// | supportedLanguageCodes  | supported_language_codes  |                             |
	// | tier                    | tier                      |                             |
	// | timeZone                | time_zone                 |                             |

	// Identity field
	f.Unimplemented_Identity(".parent")

	// Spec fields
	f.SpecField(".api_version")
	f.SpecField(".avatar_uri")
	f.SpecField(".classification_threshold")
	f.SpecField(".default_language_code")
	f.SpecField(".description")
	f.SpecField(".display_name")
	f.SpecField(".enable_logging")
	f.SpecField(".match_mode")
	f.SpecField(".supported_language_codes")
	f.SpecField(".tier")
	f.SpecField(".time_zone")

	return f
}
