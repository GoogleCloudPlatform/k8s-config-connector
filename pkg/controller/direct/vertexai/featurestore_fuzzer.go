// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.aiplatform.v1beta1.Featurestore
// api.group: aiplatform.cnrm.cloud.google.com

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(featurestoreFuzzer())
}

func featurestoreFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Featurestore{},
		VertexAIFeaturestoreSpec_FromProto, VertexAIFeaturestoreSpec_ToProto,
		VertexAIFeaturestoreObservedState_FromProto, VertexAIFeaturestoreObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")          // special field
	f.UnimplementedFields.Insert(".satisfies_pzi") // NOTYET
	f.UnimplementedFields.Insert(".satisfies_pzs") // NOTYET

	f.SpecFields.Insert(".etag")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".online_serving_config")
	f.SpecFields.Insert(".online_storage_ttl_days")
	f.SpecFields.Insert(".encryption_spec")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")

	return f
}
