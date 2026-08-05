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
// proto.message: google.cloud.discoveryengine.v1beta.LicenseConfig

package discoveryengine

import (
	discoveryenginepb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(licenseConfigFuzzer())
}

func licenseConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&discoveryenginepb.LicenseConfig{},
		DiscoveryEngineLicenseConfigSpec_v1alpha1_FromProto, DiscoveryEngineLicenseConfigSpec_v1alpha1_ToProto,
		DiscoveryEngineLicenseConfigObservedState_v1alpha1_FromProto, DiscoveryEngineLicenseConfigObservedState_v1alpha1_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".license_count")
	f.SpecField(".subscription_tier")
	f.SpecField(".auto_renew")
	f.SpecField(".start_date")
	f.SpecField(".end_date")
	f.SpecField(".subscription_term")
	f.SpecField(".free_trial")

	f.StatusField(".state")
	f.StatusField(".gemini_bundle")
	f.StatusField(".early_terminated")
	f.StatusField(".early_termination_date")

	return f
}
