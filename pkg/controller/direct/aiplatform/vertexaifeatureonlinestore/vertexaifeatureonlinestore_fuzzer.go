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
// proto.message: google.cloud.aiplatform.v1.FeatureOnlineStore
// api.group: aiplatform.cnrm.cloud.google.com

package vertexaifeatureonlinestore

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/aiplatform"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(featureonlinestoreFuzzer())
}

func featureonlinestoreFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FeatureOnlineStore{},
		aiplatform.VertexAIFeatureOnlineStoreSpec_FromProto, aiplatform.VertexAIFeatureOnlineStoreSpec_ToProto,
		aiplatform.VertexAIFeatureOnlineStoreObservedState_FromProto, aiplatform.VertexAIFeatureOnlineStoreObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".satisfies_pzi")
	f.UnimplementedFields.Insert(".satisfies_pzs")

	f.Unimplemented_NotYetTriaged(".dedicated_serving_endpoint.public_endpoint_domain_name")
	f.Unimplemented_NotYetTriaged(".dedicated_serving_endpoint.service_attachment")
	f.Unimplemented_NotYetTriaged(".bigtable.enable_direct_bigtable_access")
	f.Unimplemented_NotYetTriaged(".bigtable.bigtable_metadata")
	f.Unimplemented_NotYetTriaged(".bigtable.zone")
	f.Unimplemented_NotYetTriaged(".dedicated_serving_endpoint.private_service_connect_config.psc_automation_configs[].ip_address")
	f.Unimplemented_NotYetTriaged(".dedicated_serving_endpoint.private_service_connect_config.psc_automation_configs[].forwarding_rule")
	f.Unimplemented_NotYetTriaged(".dedicated_serving_endpoint.private_service_connect_config.psc_automation_configs[].state")
	f.Unimplemented_NotYetTriaged(".dedicated_serving_endpoint.private_service_connect_config.psc_automation_configs[].error_message")
	f.Unimplemented_NotYetTriaged(".dedicated_serving_endpoint.private_service_connect_config.service_attachment")

	f.SpecFields.Insert(".bigtable")
	f.SpecFields.Insert(".optimized")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".dedicated_serving_endpoint")
	f.SpecFields.Insert(".encryption_spec")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")

	f.Unimplemented_Etag()

	return f
}
