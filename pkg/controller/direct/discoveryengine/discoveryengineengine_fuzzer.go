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
// proto.message: google.cloud.discoveryengine.v1.Engine

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(engineFuzzer())
}

func engineFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.Engine{},
		DiscoveryEngineEngineSpec_FromProto, DiscoveryEngineEngineSpec_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.Unimplemented_NotYetTriaged(".chat_engine_metadata")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")

	f.Unimplemented_NotYetTriaged(".media_recommendation_engine_config")
	f.Unimplemented_NotYetTriaged(".chat_engine_config.allow_cross_region")

	f.Unimplemented_NotYetTriaged(".knowledge_graph_config")
	f.Unimplemented_NotYetTriaged(".app_type")
	f.Unimplemented_NotYetTriaged(".session_config")
	f.Unimplemented_NotYetTriaged(".features")
	f.Unimplemented_NotYetTriaged(".cmek_config")
	f.Unimplemented_NotYetTriaged(".configurable_billing_approach")
	f.Unimplemented_NotYetTriaged(".model_configs")
	f.Unimplemented_NotYetTriaged(".observability_config")
	f.Unimplemented_NotYetTriaged(".connector_tenant_info")
	f.Unimplemented_NotYetTriaged(".agent_gateway_setting")
	f.Unimplemented_NotYetTriaged(".marketplace_agent_visibility")
	f.Unimplemented_NotYetTriaged(".procurement_contact_emails")
	f.Unimplemented_NotYetTriaged(".associated_agent_registry")

	f.SpecField(".display_name")
	f.SpecField(".common_config")
	f.SpecField(".chat_engine_config")
	f.SpecField(".search_engine_config")
	f.SpecField(".solution_type")
	f.SpecField(".data_store_ids")
	f.SpecField(".industry_vertical")
	f.SpecField(".disable_analytics")

	return f
}
