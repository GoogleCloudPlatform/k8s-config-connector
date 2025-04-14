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
// proto.message: google.cloud.discoveryengine.v1.Engine

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(engineFuzzer())
}

func engineFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Engine{},
		DiscoveryEngineEngineSpec_FromProto, DiscoveryEngineEngineSpec_ToProto,
		EngineObservedState_FromProto, EngineObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".chat_engine_metadata") // Could be status
	f.UnimplementedFields.Insert(".create_time")          // Could be status
	f.UnimplementedFields.Insert(".update_time")          // Could be status

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".common_config")
	f.SpecFields.Insert(".chat_engine_config")
	f.SpecFields.Insert(".search_engine_config")
	f.SpecFields.Insert(".solution_type")
	f.SpecFields.Insert(".data_store_ids")
	f.SpecFields.Insert(".industry_vertical")
	f.SpecFields.Insert(".disable_analytics")

	// f.StatusFields.Insert(".create_time")
	// f.StatusFields.Insert(".update_time")

	return f
}
