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
	pb "cloud.google.com/go/discoveryengine/apiv1alpha/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzDataStore())
}

func fuzzDataStore() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DataStore{},
		DiscoveryEngineDataStoreSpec_FromProto, DiscoveryEngineDataStoreSpec_ToProto,
		DataStoreObservedState_FromProto, DataStoreObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")                       // special field
	f.UnimplementedFields.Insert(".document_processing_config") // complex map[string]object, so not implementing yet
	f.UnimplementedFields.Insert(".starting_schema")            // Tricky field, only on create

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".industry_vertical")
	f.SpecFields.Insert(".solution_types")
	f.SpecFields.Insert(".content_config")
	f.SpecFields.Insert(".workspace_config")

	f.StatusFields.Insert(".default_schema_id")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".billing_estimation")

	return f
}
