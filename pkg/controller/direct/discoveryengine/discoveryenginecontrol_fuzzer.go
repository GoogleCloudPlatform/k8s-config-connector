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
// proto.message: google.cloud.discoveryengine.v1.Control

package discoveryengine

import (
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMSpecFuzzer(controlFuzzer())
}

func controlFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedSpecFuzzer(&pb.Control{},
		DiscoveryEngineControlSpec_v1alpha1_FromProto, DiscoveryEngineControlSpec_v1alpha1_ToProto,
	)

	f.SpecField(".display_name")
	f.SpecField(".solution_type")
	f.SpecField(".use_cases")
	f.SpecField(".conditions")

	f.SpecField(".boost_action")
	f.SpecField(".filter_action")
	f.SpecField(".redirect_action")
	f.SpecField(".synonyms_action")
	f.SpecField(".promote_action")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".associated_serving_config_ids")

	return f
}
