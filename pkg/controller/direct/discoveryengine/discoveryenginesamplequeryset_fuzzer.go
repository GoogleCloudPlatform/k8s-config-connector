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
// proto.message: google.cloud.discoveryengine.v1beta.SampleQuerySet

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(sampleQuerySetFuzzer())
}

func sampleQuerySetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.SampleQuerySet{},
		DiscoveryEngineSampleQuerySetSpec_FromProto, DiscoveryEngineSampleQuerySetSpec_ToProto,
		DiscoveryEngineSampleQuerySetObservedState_FromProto, DiscoveryEngineSampleQuerySetObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".display_name")
	f.SpecField(".description")

	f.StatusField(".create_time")

	return f
}
