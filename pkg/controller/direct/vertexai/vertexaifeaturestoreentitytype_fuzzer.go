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
// proto.message: google.cloud.aiplatform.v1beta1.EntityType
// api.group: vertexai.cnrm.cloud.google.com

package vertexai

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIFeaturestoreEntityTypeFuzzer())
}

func vertexAIFeaturestoreEntityTypeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.EntityType{},
		VertexAIFeaturestoreEntityTypeSpec_v1alpha1_FromProto,
		VertexAIFeaturestoreEntityTypeSpec_v1alpha1_ToProto,
		VertexAIFeaturestoreEntityTypeStatus_v1alpha1_FromProto,
		VertexAIFeaturestoreEntityTypeStatus_v1alpha1_ToProto,
	)

	// Fields comparison:
	// KRM Spec:
	// - description -> pb.EntityType.description
	// - featurestore -> part of identity / parent resource URL (not in pb.EntityType)
	// - monitoringConfig -> pb.EntityType.monitoring_config
	// - offlineStorageTtlDays -> pb.EntityType.offline_storage_ttl_days
	// - resourceID -> part of identity / GCP resource name (f.Unimplemented_Identity(".name"))
	//
	// KRM Status:
	// - conditions -> KRM standard status field
	// - createTime -> pb.EntityType.create_time
	// - etag -> pb.EntityType.etag
	// - observedGeneration -> KRM standard status field
	// - region -> KRM status field
	// - updateTime -> pb.EntityType.update_time

	// Identity field
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".description")
	f.SpecField(".monitoring_config")
	f.SpecField(".offline_storage_ttl_days")

	// Status fields
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".etag")

	// Unimplemented / Not yet triaged fields in pb.EntityType
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".satisfies_pzs")
	f.Unimplemented_NotYetTriaged(".satisfies_pzi")

	// FilterSpec handles the required non-pointer `value` float64 in KRM
	// FeaturestoreMonitoringConfig_ThresholdConfig (which defaults to 0 on round-trip
	// if the proto oneof `threshold` was originally unset/nil).
	f.FilterSpec = func(in *pb.EntityType) {
		if in.MonitoringConfig != nil {
			if tc := in.MonitoringConfig.CategoricalThresholdConfig; tc != nil && tc.Threshold == nil {
				tc.Threshold = &pb.FeaturestoreMonitoringConfig_ThresholdConfig_Value{Value: 0}
			}
			if tc := in.MonitoringConfig.NumericalThresholdConfig; tc != nil && tc.Threshold == nil {
				tc.Threshold = &pb.FeaturestoreMonitoringConfig_ThresholdConfig_Value{Value: 0}
			}
		}
	}

	return f
}
