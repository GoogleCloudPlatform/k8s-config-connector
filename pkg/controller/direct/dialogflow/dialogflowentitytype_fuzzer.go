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
// proto.message: google.cloud.dialogflow.v2.EntityType
// api.group: dialogflow.cnrm.cloud.google.com

package dialogflow

import (
	pb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(dialogflowEntityTypeFuzzer())
}

func dialogflowEntityTypeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.EntityType{},
		DialogflowEntityTypeSpec_FromProto, DialogflowEntityTypeSpec_ToProto,
	)

	// Documented Field Comparison (KRM Spec -> Proto):
	// - displayName (string): Mapped to .display_name.
	// - enableFuzzyExtraction (*bool): Mapped to .enable_fuzzy_extraction.
	// - entities ([]EntityType_Entity): Mapped to .entities.
	//   - entities[].value (string): Mapped to .entities[].value.
	//   - entities[].synonyms ([]string): Mapped to .entities[].synonyms.
	// - kind (string): Mapped to .kind.
	// - projectRef (*refs.ProjectRef): KCC reference field, not part of GCP proto.
	// - resourceID (*string): KCC acquisition field, maps to the last segment of .name.

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".enable_fuzzy_extraction")
	f.SpecField(".entities")
	f.SpecField(".entities[].value")
	f.SpecField(".entities[].synonyms")
	f.SpecField(".kind")

	// Identity / Special fields
	f.Unimplemented_Identity(".name")

	// Unimplemented / Untriaged proto fields
	f.Unimplemented_NotYetTriaged(".auto_expansion_mode")

	return f
}
