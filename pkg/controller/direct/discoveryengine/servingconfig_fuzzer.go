// Copyright 2024 Google LLC
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
// api.group: discoveryengine.cnrm.cloud.google.com

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(discoveryEngineServingConfigFuzzer())
}

func discoveryEngineServingConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ServingConfig{},
		DiscoveryEngineServingConfigSpec_FromProto, DiscoveryEngineServingConfigSpec_ToProto,
		DiscoveryEngineServingConfigObservedState_FromProto, DiscoveryEngineServingConfigObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".media_config")
	f.SpecFields.Insert(".generic_config")
	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".solution_type")
	f.SpecFields.Insert(".model_id")
	f.SpecFields.Insert(".diversity_level")
	f.SpecFields.Insert(".embedding_config")
	f.SpecFields.Insert(".ranking_expression")
	f.SpecFields.Insert(".filter_control_ids")
	f.SpecFields.Insert(".boost_control_ids")
	f.SpecFields.Insert(".redirect_control_ids")
	f.SpecFields.Insert(".synonyms_control_ids")
	f.SpecFields.Insert(".oneway_synonyms_control_ids")
	f.SpecFields.Insert(".dissociate_control_ids")
	f.SpecFields.Insert(".replacement_control_ids")
	f.SpecFields.Insert(".ignore_control_ids")
	f.SpecFields.Insert(".personalization_spec")

	// Fields we don't want to implement (yet) because they are very volatile
	f.UnimplementedFields.Insert(".create_time")
	f.UnimplementedFields.Insert(".update_time")

	return f
}
