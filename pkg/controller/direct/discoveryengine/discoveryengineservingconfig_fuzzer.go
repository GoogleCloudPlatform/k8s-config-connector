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
// proto.message: google.cloud.discoveryengine.v1beta.ServingConfig

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(servingConfigFuzzer())
}

func servingConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ServingConfig{},
		DiscoveryEngineServingConfigSpec_v1alpha1_FromProto, DiscoveryEngineServingConfigSpec_v1alpha1_ToProto,
		DiscoveryEngineServingConfigObservedState_v1alpha1_FromProto, DiscoveryEngineServingConfigObservedState_v1alpha1_ToProto,
	)

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".answer_generation_spec")
	f.Unimplemented_NotYetTriaged(".promote_control_ids")
	f.Unimplemented_NotYetTriaged(".generic_config.content_search_spec.summary_spec.multimodal_spec")

	f.SpecField(".display_name")
	f.SpecField(".solution_type")
	f.SpecField(".model_id")
	f.SpecField(".diversity_level")
	f.SpecField(".embedding_config")
	f.SpecField(".ranking_expression")
	f.SpecField(".filter_control_ids")
	f.SpecField(".boost_control_ids")
	f.SpecField(".redirect_control_ids")
	f.SpecField(".synonyms_control_ids")
	f.SpecField(".oneway_synonyms_control_ids")
	f.SpecField(".dissociate_control_ids")
	f.SpecField(".replacement_control_ids")
	f.SpecField(".ignore_control_ids")
	f.SpecField(".personalization_spec")
	f.SpecField(".media_config")
	f.SpecField(".generic_config")

	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
