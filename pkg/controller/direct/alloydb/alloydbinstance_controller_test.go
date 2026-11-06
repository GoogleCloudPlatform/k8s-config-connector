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

package alloydb

import (
	"context"
	"reflect"
	"testing"

	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/alloydb/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func TestAlloyDBInstancePSCMappings(t *testing.T) {
	mapCtx := &direct.MapContext{}
	spec := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjects: []string{"123456789"},
			PSCAutoConnections: []krm.Instance_PSCAutoConnectionConfig{
				{
					ConsumerProject: direct.PtrTo("my-project"),
					ConsumerNetwork: direct.PtrTo("projects/my-project/global/networks/default"),
				},
			},
			PSCInterfaceConfigs: []krm.Instance_PSCInterfaceConfig{
				{
					NetworkAttachmentResource: direct.PtrTo("projects/my-project/regions/us-central1/networkAttachments/my-na"),
				},
			},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_ENABLED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("ENABLED"),
		},
	}

	protoInst := AlloyDBInstanceSpec_ToProto(mapCtx, spec)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in AlloyDBInstanceSpec_ToProto: %v", err)
	}
	if protoInst.GetPscInstanceConfig().GetPscAutoDnsState() != pb.Instance_PSC_AUTO_DNS_STATE_ENABLED {
		t.Errorf("got PscAutoDnsState = %v, want %v", protoInst.GetPscInstanceConfig().GetPscAutoDnsState(), pb.Instance_PSC_AUTO_DNS_STATE_ENABLED)
	}
	if protoInst.GetPscInstanceConfig().GetPscAutoConnectionPolicyState() != pb.Instance_PscInstanceConfig_ENABLED {
		t.Errorf("got PscAutoConnectionPolicyState = %v, want %v", protoInst.GetPscInstanceConfig().GetPscAutoConnectionPolicyState(), pb.Instance_PscInstanceConfig_ENABLED)
	}

	roundTripSpec := AlloyDBInstanceSpec_FromProto(mapCtx, protoInst)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in AlloyDBInstanceSpec_FromProto: %v", err)
	}
	if !reflect.DeepEqual(roundTripSpec.PSCInstanceConfig, spec.PSCInstanceConfig) {
		t.Errorf("roundTripSpec.PSCInstanceConfig = %+v, want %+v", roundTripSpec.PSCInstanceConfig, spec.PSCInstanceConfig)
	}

	// Verify ObservedState mapping for PSCInstanceConfig and PSCInstanceInfo.
	protoInst.PscInstanceConfig.ServiceAttachmentLink = "projects/123/regions/us-central1/serviceAttachments/sa-1"
	protoInst.PscInstanceConfig.PscDnsName = "uid.uid.us-central1.alloydb-psc.goog."
	protoInst.PscInstanceConfig.PscAutoConnections[0].IpAddress = "10.0.0.10"
	protoInst.PscInstanceConfig.PscAutoConnections[0].Status = "ACTIVE"
	protoInst.PscInstanceConfig.PscAutoConnections[0].ConsumerNetworkStatus = "VALID"
	protoInst.PscInstanceInfo = &pb.Instance_PscInstanceInfo{
		EffectivePscAutoDnsEnabled:       true,
		PscAutoDnsNames:                  []string{"inst.cluster.us-central1.my-project.alloydb.goog."},
		EffectivePscAutoConnectionPolicy: true,
		ServiceConnectionPolicy:          "projects/my-project/locations/us-central1/serviceConnectionPolicies/default-scp",
	}

	status := AlloyDBInstanceStatus_FromProto(mapCtx, protoInst)
	if err := mapCtx.Err(); err != nil {
		t.Fatalf("unexpected error in AlloyDBInstanceStatus_FromProto: %v", err)
	}
	if status.ObservedState == nil || status.ObservedState.PSCInstanceConfig == nil || status.ObservedState.PSCInstanceInfo == nil {
		t.Fatalf("expected ObservedState.PSCInstanceConfig and ObservedState.PSCInstanceInfo to be populated, got %+v", status.ObservedState)
	}
	if !direct.ValueOf(status.ObservedState.PSCInstanceInfo.EffectivePSCAutoDNSEnabled) {
		t.Errorf("got EffectivePSCAutoDNSEnabled = false, want true")
	}
	if !direct.ValueOf(status.ObservedState.PSCInstanceInfo.EffectivePSCAutoConnectionPolicy) {
		t.Errorf("got EffectivePSCAutoConnectionPolicy = false, want true")
	}
	if got := direct.ValueOf(status.ObservedState.PSCInstanceInfo.ServiceConnectionPolicy); got != "projects/my-project/locations/us-central1/serviceConnectionPolicies/default-scp" {
		t.Errorf("got ServiceConnectionPolicy = %q, want projects/my-project/locations/us-central1/serviceConnectionPolicies/default-scp", got)
	}
}

func TestCompareInstancePSC(t *testing.T) {
	ctx := context.Background()
	actual := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjects:      []string{"111"},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_ENABLED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("ENABLED"),
		},
	}
	desired := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjects:      []string{"111", "222"},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_DISABLED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("DISABLED"),
		},
	}

	paths, err := compareInstance(ctx, actual, desired)
	if err != nil {
		t.Fatalf("compareInstance returned unexpected error: %v", err)
	}
	wantPaths := []string{
		"psc_instance_config.allowed_consumer_projects",
		"psc_instance_config.psc_auto_dns_state",
		"psc_instance_config.psc_auto_connection_policy_state",
	}
	if !reflect.DeepEqual(paths, wantPaths) {
		t.Errorf("compareInstance paths = %v, want %v", paths, wantPaths)
	}

	// Verify that explicit *_UNSPECIFIED does not trigger a diff when actual enum is nil (0 in proto).
	actualUnspecified := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjects: []string{"111"},
		},
	}
	desiredUnspecified := &krm.AlloyDBInstanceSpec{
		PSCInstanceConfig: &krm.Instance_PSCInstanceConfig{
			AllowedConsumerProjects:      []string{"111"},
			PSCAutoDNSState:              direct.PtrTo("PSC_AUTO_DNS_STATE_UNSPECIFIED"),
			PSCAutoConnectionPolicyState: direct.PtrTo("PSC_AUTO_CONNECTION_POLICY_STATE_UNSPECIFIED"),
		},
	}
	unspecifiedPaths, err := compareInstance(ctx, actualUnspecified, desiredUnspecified)
	if err != nil {
		t.Fatalf("compareInstance returned unexpected error for UNSPECIFIED: %v", err)
	}
	if len(unspecifiedPaths) != 0 {
		t.Errorf("compareInstance paths for UNSPECIFIED = %v, want empty", unspecifiedPaths)
	}
}
