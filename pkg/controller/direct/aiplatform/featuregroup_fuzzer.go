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

package aiplatform

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIFeatureGroupFuzzer())
}

func vertexAIFeatureGroupFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FeatureGroup{},
		VertexAIFeatureGroupSpec_FromProto, VertexAIFeatureGroupSpec_ToProto,
		VertexAIFeatureGroupObservedState_FromProto, VertexAIFeatureGroupObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".big_query")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.Unimplemented_Etag()

	return f
}
