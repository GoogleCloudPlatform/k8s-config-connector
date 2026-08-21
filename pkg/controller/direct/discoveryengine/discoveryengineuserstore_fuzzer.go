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
// proto.message: google.cloud.discoveryengine.v1beta.UserStore

package discoveryengine

import (
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzUserStore())
}

func fuzzUserStore() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&discoveryenginepb.UserStore{},
		DiscoveryEngineUserStoreSpec_v1alpha1_FromProto, DiscoveryEngineUserStoreSpec_v1alpha1_ToProto,
		DiscoveryEngineUserStoreObservedState_v1alpha1_FromProto, DiscoveryEngineUserStoreObservedState_v1alpha1_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".display_name")
	f.SpecField(".default_license_config")
	f.SpecField(".enable_license_auto_register")
	f.SpecField(".enable_expired_license_auto_update")

	return f
}
