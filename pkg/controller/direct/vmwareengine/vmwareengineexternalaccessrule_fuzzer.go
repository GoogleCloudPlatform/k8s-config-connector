// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.vmwareengine.v1.ExternalAccessRule
// api.group: vmwareengine.cnrm.cloud.google.com

package vmwareengine

import (
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vmwareEngineExternalAccessRuleFuzzer())
}

func vmwareEngineExternalAccessRuleFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ExternalAccessRule{},
		VMwareEngineExternalAccessRuleSpec_FromProto, VMwareEngineExternalAccessRuleSpec_ToProto,
		VMwareEngineExternalAccessRuleObservedState_FromProto, VMwareEngineExternalAccessRuleObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".priority")
	f.SpecFields.Insert(".action")
	f.SpecFields.Insert(".ip_protocol")
	f.SpecFields.Insert(".source_ip_ranges")
	f.SpecFields.Insert(".source_ports")
	f.SpecFields.Insert(".destination_ip_ranges")
	f.SpecFields.Insert(".destination_ports")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".state")
	f.StatusFields.Insert(".uid")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
