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

package vertexaiindex

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIIndexFuzzer())
}

func vertexAIIndexFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Index{},
		VertexAIIndexSpec_FromProto, VertexAIIndexSpec_ToProto,
		VertexAIIndexObservedState_FromProto, VertexAIIndexObservedState_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".metadata_schema_uri")
	f.SpecField(".metadata")
	f.SpecField(".etag")
	f.SpecField(".index_update_method")
	f.SpecField(".encryption_spec")

	f.StatusField(".name")
	f.StatusField(".deployed_indexes")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".index_stats")
	f.StatusField(".satisfies_pzs")
	f.StatusField(".satisfies_pzi")

	f.Unimplemented_LabelsAnnotations(".labels")

	return f
}
