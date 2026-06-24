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
// proto.message: google.cloud.aiplatform.v1beta1.Dataset
// api.group: vertexai.cnrm.cloud.google.com

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIDatasetFuzzer())
}

func vertexAIDatasetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Dataset{},
		VertexAIDatasetSpec_FromProto, VertexAIDatasetSpec_ToProto,
		VertexAIDatasetObservedState_FromProto, VertexAIDatasetObservedState_ToProto,
	)

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".metadata_schema_uri")
	f.SpecField(".encryption_spec")

	// Status fields
	f.StatusField(".name")
	f.StatusField(".create_time")

	// Unimplemented / Not yet triaged fields
	f.Unimplemented_Etag()
	f.Unimplemented_NotYetTriaged(".description")
	f.Unimplemented_NotYetTriaged(".metadata")
	f.Unimplemented_NotYetTriaged(".data_item_count")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".saved_queries")
	f.Unimplemented_NotYetTriaged(".metadata_artifact")
	f.Unimplemented_NotYetTriaged(".model_reference")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")

	return f
}
