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
// proto.message: google.cloud.aiplatform.v1beta1.Feature
// api.group: vertexai.cnrm.cloud.google.com

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIFeaturestoreEntityTypeFeatureFuzzer())
}

func vertexAIFeaturestoreEntityTypeFeatureFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Feature{},
		VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_FromProto,
		VertexAIFeaturestoreEntityTypeFeatureSpec_v1alpha1_ToProto,
		VertexAIFeaturestoreEntityTypeFeatureStatus_v1alpha1_FromProto,
		VertexAIFeaturestoreEntityTypeFeatureStatus_v1alpha1_ToProto,
	)

	// Fields comparison:
	// KRM Spec:
	// - description -> pb.Feature.description
	// - entitytype -> part of identity / parent resource URL (not in pb.Feature)
	// - resourceID -> part of identity / GCP resource name (f.Unimplemented_Identity(".name"))
	// - valueType -> pb.Feature.value_type
	//
	// KRM Status:
	// - conditions -> KRM standard status field
	// - createTime -> pb.Feature.create_time
	// - etag -> pb.Feature.etag
	// - observedGeneration -> KRM standard status field
	// - region -> KRM status field
	// - updateTime -> pb.Feature.update_time

	// Identity field
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".value_type")

	// Status fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".etag")

	// Unimplemented / Not yet triaged fields in pb.Feature
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".monitoring_config")
	f.Unimplemented_NotYetTriaged(".disable_monitoring")
	f.Unimplemented_NotYetTriaged(".monitoring_stats")
	f.Unimplemented_NotYetTriaged(".monitoring_stats_anomalies")
	f.Unimplemented_NotYetTriaged(".feature_stats_and_anomaly")
	f.Unimplemented_NotYetTriaged(".version_column_name")
	f.Unimplemented_NotYetTriaged(".point_of_contact")

	return f
}
