// Copyright 2025 Google LLC
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

package mockalloydb

import (
	"testing"

	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

func TestSetInstanceFieldsPSCDefaults(t *testing.T) {
	name := &instanceName{
		Project:      &projects.ProjectData{ID: "my-project", Number: 123456789},
		Location:     "europe-west1",
		ClusterName:  "my-cluster",
		InstanceName: "my-instance",
	}

	// 1. Default (UNSPECIFIED) states should enable effective PSC auto DNS and connection policy.
	inst := &pb.Instance{
		InstanceType: pb.Instance_PRIMARY,
		PscInstanceConfig: &pb.Instance_PscInstanceConfig{
			AllowedConsumerProjects: []string{"123456789"},
		},
	}
	setInstanceFields(name, inst)

	if got := inst.GetIpAddress(); got != "" {
		t.Errorf("got IpAddress = %q for PSC instance, want empty", got)
	}
	if !inst.GetPscInstanceInfo().GetEffectivePscAutoDnsEnabled() {
		t.Errorf("got EffectivePscAutoDnsEnabled = false for UNSPECIFIED, want true")
	}
	if !inst.GetPscInstanceInfo().GetEffectivePscAutoConnectionPolicy() {
		t.Errorf("got EffectivePscAutoConnectionPolicy = false for UNSPECIFIED, want true")
	}

	// 2. DISABLED states should disable effective PSC auto DNS and connection policy.
	inst.PscInstanceConfig.PscAutoDnsState = pb.Instance_PSC_AUTO_DNS_STATE_DISABLED
	inst.PscInstanceConfig.PscAutoConnectionPolicyState = pb.Instance_PscInstanceConfig_DISABLED
	setInstanceFields(name, inst)

	if inst.GetPscInstanceInfo().GetEffectivePscAutoDnsEnabled() {
		t.Errorf("got EffectivePscAutoDnsEnabled = true for DISABLED, want false")
	}
	if inst.GetPscInstanceInfo().GetEffectivePscAutoConnectionPolicy() {
		t.Errorf("got EffectivePscAutoConnectionPolicy = true for DISABLED, want false")
	}

	// 3. Non-PSC instances should retain the default PSA IpAddress.
	psaInst := &pb.Instance{
		InstanceType: pb.Instance_PRIMARY,
	}
	setInstanceFields(name, psaInst)
	if got := psaInst.GetIpAddress(); got != "10.1.2.3" {
		t.Errorf("got IpAddress = %q for non-PSC instance, want 10.1.2.3", got)
	}
}
