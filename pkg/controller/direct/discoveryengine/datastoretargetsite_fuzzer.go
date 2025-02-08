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
// proto.message: google.cloud.discoveryengine.v1.TargetSite

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"k8s.io/apimachinery/pkg/util/sets"
)

func init() {
	fuzztesting.RegisterFuzzer(dataStoreTargetSiteSpecFuzzer().FuzzSpec)
	fuzztesting.RegisterFuzzer(dataStoreTargetSiteObservedStateFuzzer().FuzzObservedState)
}

var dataStoreTargetSiteKrmFields = fuzztesting.KRMFields{
	UnimplementedFields: sets.New(".name"),
	SpecFields: sets.New(".exact_match",
		".provided_uri_pattern",
		".type"),
	ObservedStateFields: sets.New(".generated_uri_pattern",
		".root_domain_uri",
		".site_verification_info",
		".indexing_status",
		".update_time",
		".failure_reason"),
}

func dataStoreTargetSiteSpecFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TargetSite{},
		DiscoveryEngineDataStoreTargetSiteSpec_FromProto, DiscoveryEngineDataStoreTargetSiteSpec_ToProto,
	)
	f.KRMFields = dataStoreTargetSiteKrmFields
	return f
}

func dataStoreTargetSiteObservedStateFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.TargetSite{},
		DiscoveryEngineDataStoreTargetSiteObservedState_FromProto, DiscoveryEngineDataStoreTargetSiteObservedState_ToProto,
	)
	f.KRMFields = dataStoreTargetSiteKrmFields
	return f
}
