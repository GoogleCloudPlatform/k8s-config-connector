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
// proto.message: google.cloud.aiplatform.v1.Study
// api.group: aiplatform.cnrm.cloud.google.com

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIStudyFuzzer())
}

func vertexAIStudyFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Study{},
		VertexAIStudySpec_FromProto, VertexAIStudySpec_ToProto,
		VertexAIStudyObservedState_FromProto, VertexAIStudyObservedState_ToProto,
	)

	// Identity and special fields
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".study_spec")

	// Status fields (ObservedState)
	f.StatusField(".state")
	f.StatusField(".create_time")

	// Unmapped or unhandled fields to ignore in the fuzzer roundtrip
	f.Unimplemented_NotYetTriaged(".inactive_reason")
	f.Unimplemented_NotYetTriaged(".study_spec.parameters[].conditional_parameter_specs[].parameter_spec")

	return f
}
