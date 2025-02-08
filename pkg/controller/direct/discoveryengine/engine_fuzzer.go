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
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"k8s.io/apimachinery/pkg/util/sets"
)

func init() {
	fuzztesting.RegisterFuzzer(engineSpecFuzzer().FuzzSpec)
}

var engineKrmFields = fuzztesting.KRMFields{
	UnimplementedFields: sets.New(
		".chat_engine_metadata", // Could be status
		".name",                 // special field
		".create_time",          // Could be status
		".update_time"),         // Could be status
	SpecFields: sets.New(".display_name",
		".common_config",
		".chat_engine_config",
		".search_engine_config",
		".solution_type",
		".data_store_ids",
		".industry_vertical",
		".disable_analytics"),
	//ObservedStateFields: sets.New(
	//".create_time",
	//".update_time"),
}

func engineSpecFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Engine{},
		DiscoveryEngineEngineSpec_FromProto, DiscoveryEngineEngineSpec_ToProto,
	)
	f.KRMFields = engineKrmFields
	return f
}
