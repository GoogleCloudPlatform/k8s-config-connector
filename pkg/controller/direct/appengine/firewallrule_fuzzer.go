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
// proto.message: google.appengine.v1.AppEngineFirewallRule
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
		AppEngineFirewallRuleSpec_FromProto, AppEngineFirewallRuleSpec_ToProto,
		AppEngineFirewallRuleObservedState_FromProto, AppEngineFirewallRuleObservedState_ToProto,
	)

	f.SpecField(".priority")
	f.SpecField(".action")
	f.SpecField(".source_range")
	f.SpecField(".description")

	return f
}
