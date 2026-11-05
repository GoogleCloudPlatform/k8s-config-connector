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
// proto.message: google.appengine.v1.FirewallRule
// api.group: appengine.cnrm.cloud.google.com

package appengine

import (
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(appEngineFirewallRuleFuzzer())
}

func appEngineFirewallRuleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.FirewallRule{},
		AppEngineFirewallRuleSpec_v1alpha1_FromProto, AppEngineFirewallRuleSpec_v1alpha1_ToProto,
		AppEngineFirewallRuleObservedState_v1alpha1_FromProto, AppEngineFirewallRuleObservedState_v1alpha1_ToProto,
	)

	// +-------------------------+---------------------------+----------------------------------------------+
	// | KRM Spec Field          | Proto Field               | Fuzzer/Mapping Status                        |
	// +-------------------------+---------------------------+----------------------------------------------+
	// | action                  | action                    | SpecField (".action")                        |
	// | description             | description               | SpecField (".description")                   |
	// | project                 | N/A                       | KRM spec-only GCP project reference          |
	// | resourceID              | priority                  | SpecField (".priority")                      |
	// | sourceRange             | source_range              | SpecField (".source_range")                  |
	// +-------------------------+---------------------------+----------------------------------------------+

	// Spec fields
	f.SpecField(".action")
	f.SpecField(".description")
	f.SpecField(".priority")
	f.SpecField(".source_range")

	return f
}
