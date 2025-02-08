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
// proto.message: google.cloud.discoveryengine.v1.DataStore

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"k8s.io/apimachinery/pkg/util/sets"
)

func init() {
	fuzztesting.RegisterFuzzer(dataStoreSpecFuzzer().FuzzSpec)
	fuzztesting.RegisterFuzzer(dataStoreObservedStateFuzzer().FuzzObservedState)
}

var dataStoreKrmFields = fuzztesting.KRMFields{
	UnimplementedFields: sets.New(
		".name",                       // special field
		".document_processing_config", // complex map[string]object, so not implementing yet
		".starting_schema",            // Tricky field, only on create
	),
	SpecFields: sets.New(
		".display_name",
		".industry_vertical",
		".solution_types",
		".content_config",
		".workspace_config"),
	ObservedStateFields: sets.New(
		".default_schema_id",
		".create_time",
		".billing_estimation"),
}

func dataStoreSpecFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DataStore{},
		DiscoveryEngineDataStoreSpec_FromProto, DiscoveryEngineDataStoreSpec_ToProto,
	)
	f.KRMFields = dataStoreKrmFields
	return f
}

func dataStoreObservedStateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DataStore{},
		DiscoveryEngineDataStoreObservedState_FromProto, DiscoveryEngineDataStoreObservedState_ToProto,
	)
	f.KRMFields = dataStoreKrmFields
	return f
}
