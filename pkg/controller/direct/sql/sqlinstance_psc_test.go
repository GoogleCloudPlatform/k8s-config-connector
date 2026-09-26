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

package sql

import (
	"context"
	"testing"

	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/sqladmin/v1beta4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestDiffPscConfig_EmptyAndNil(t *testing.T) {
	t.Parallel()

	// nil desired and empty actual
	diff1 := DiffPscConfig(nil, &api.PscConfig{})
	if diff1.HasDiff() {
		t.Errorf("DiffPscConfig(nil, empty) reported diff: %v", diff1.Fields)
	}

	// empty desired and nil actual
	diff2 := DiffPscConfig(&api.PscConfig{}, nil)
	if diff2.HasDiff() {
		t.Errorf("DiffPscConfig(empty, nil) reported diff: %v", diff2.Fields)
	}

	// both having empty slice
	diff3 := DiffPscConfig(
		&api.PscConfig{PscAutoConnections: []*api.PscAutoConnectionConfig{}},
		&api.PscConfig{PscAutoConnections: []*api.PscAutoConnectionConfig{}},
	)
	if diff3.HasDiff() {
		t.Errorf("DiffPscConfig(empty slice, empty slice) reported diff: %v", diff3.Fields)
	}
}

func TestDiffPscConfig_URIFormats(t *testing.T) {
	t.Parallel()

	desired := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{
				ConsumerNetwork: "projects/my-project/global/networks/my-network",
			},
		},
	}

	// GCP returns full REST URI
	actual := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{
				ConsumerNetwork: "https://www.googleapis.com/compute/v1/projects/my-project/global/networks/my-network",
			},
		},
	}

	diff := DiffPscConfig(desired, actual)
	if diff.HasDiff() {
		t.Errorf("DiffPscConfig failed to normalize URI format: %v", diff.Fields)
	}
}

func TestDiffPscConfig_ReorderedList(t *testing.T) {
	t.Parallel()

	desired := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{ConsumerNetwork: "projects/proj/global/networks/network-b"},
			{ConsumerNetwork: "projects/proj/global/networks/network-a"},
		},
	}

	actual := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{ConsumerNetwork: "projects/proj/global/networks/network-a"},
			{ConsumerNetwork: "projects/proj/global/networks/network-b"},
		},
	}

	diff := DiffPscConfig(desired, actual)
	if diff.HasDiff() {
		t.Errorf("DiffPscConfig reported false drift on reordered list: %v", diff.Fields)
	}
}

func TestDiffPscConfig_IgnoreOutputFields(t *testing.T) {
	t.Parallel()

	desired := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{
				ConsumerNetwork: "projects/my-project/global/networks/default",
				ConsumerProject: "my-service-project",
			},
		},
	}

	// Actual GCP response with runtime server-generated fields
	actual := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{
				ConsumerNetwork:                       "projects/my-project/global/networks/default",
				ConsumerProject:                       "my-service-project",
				IpAddress:                             "10.128.0.50",
				Status:                                "ACTIVE",
				ConsumerNetworkStatus:                 "POLICY_CREATED",
				ServiceConnectionPolicy:               "projects/my-project/regions/us-central1/serviceConnectionPolicies/default-12345",
				ServiceConnectionPolicyCreationResult: "SUCCESS",
			},
		},
	}

	diff := DiffPscConfig(desired, actual)
	if diff.HasDiff() {
		t.Errorf("DiffPscConfig reported unexpected drift from server-generated output fields: %v", diff.Fields)
	}
}

func TestDiffPscConfig_DriftDetection(t *testing.T) {
	t.Parallel()

	desired := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{
				ConsumerNetwork: "projects/my-project/global/networks/network-1",
				ConsumerProject: "proj-1",
			},
		},
	}

	// Case 1: Different network
	actualDifferentNet := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{
				ConsumerNetwork: "projects/my-project/global/networks/network-2",
				ConsumerProject: "proj-1",
			},
		},
	}
	diffNet := DiffPscConfig(desired, actualDifferentNet)
	if !diffNet.HasDiff() {
		t.Errorf("Expected diff on different network, got none")
	}

	// Case 2: Different project
	actualDifferentProj := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{
				ConsumerNetwork: "projects/my-project/global/networks/network-1",
				ConsumerProject: "proj-2",
			},
		},
	}
	diffProj := DiffPscConfig(desired, actualDifferentProj)
	if !diffProj.HasDiff() {
		t.Errorf("Expected diff on different project, got none")
	}

	// Case 3: Added connection in desired
	desiredTwo := &api.PscConfig{
		PscEnabled: true,
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{ConsumerNetwork: "projects/my-project/global/networks/network-1"},
			{ConsumerNetwork: "projects/my-project/global/networks/network-2"},
		},
	}
	diffCount := DiffPscConfig(desiredTwo, actualDifferentNet)
	if !diffCount.HasDiff() {
		t.Errorf("Expected diff when desired has extra connection, got none")
	}
}

func TestPscAutoConnections_KRMToGCP_Mappings(t *testing.T) {
	t.Parallel()

	krmPsc := []krm.InstancePscConfig{
		{
			PscEnabled:                     direct.PtrTo(true),
			PscAutoConnectionPolicyEnabled: direct.PtrTo(true),
			AllowedConsumerProjects:        []string{"proj-a", "proj-b"},
			PscAutoConnections: []krm.InstancePscAutoConnectionConfig{
				{
					ConsumerNetwork: direct.PtrTo("projects/host-p/global/networks/net-1"),
					// Test alias ConsumerServiceProjectId
					ConsumerServiceProjectId: direct.PtrTo("svc-p1"),
				},
				{
					// Test fallback from ConsumerNetworkRef
					ConsumerNetworkRef: &computerefs.ComputeNetworkRef{
						External: "projects/host-p/global/networks/net-2",
					},
					ConsumerProject: direct.PtrTo("svc-p2"),
				},
			},
		},
	}

	gcpObj := InstancePscConfigKRMToGCP(krmPsc)
	if gcpObj == nil {
		t.Fatalf("InstancePscConfigKRMToGCP returned nil")
	}

	if !gcpObj.PscEnabled {
		t.Errorf("expected PscEnabled=true")
	}
	if !gcpObj.PscAutoConnectionPolicyEnabled {
		t.Errorf("expected PscAutoConnectionPolicyEnabled=true")
	}
	if len(gcpObj.PscAutoConnections) != 2 {
		t.Fatalf("expected 2 PscAutoConnections, got %d", len(gcpObj.PscAutoConnections))
	}

	conn1 := gcpObj.PscAutoConnections[0]
	if conn1.ConsumerNetwork != "projects/host-p/global/networks/net-1" {
		t.Errorf("conn1 network mismatch: got %s", conn1.ConsumerNetwork)
	}
	if conn1.ConsumerProject != "svc-p1" {
		t.Errorf("conn1 project mismatch: got %s", conn1.ConsumerProject)
	}

	conn2 := gcpObj.PscAutoConnections[1]
	if conn2.ConsumerNetwork != "projects/host-p/global/networks/net-2" {
		t.Errorf("conn2 network mismatch: got %s", conn2.ConsumerNetwork)
	}
	if conn2.ConsumerProject != "svc-p2" {
		t.Errorf("conn2 project mismatch: got %s", conn2.ConsumerProject)
	}
}

func TestPscAutoConnections_GCPToKRM_Mappings(t *testing.T) {
	t.Parallel()

	gcpObj := &api.PscConfig{
		PscEnabled:                     true,
		PscAutoConnectionPolicyEnabled: true,
		AllowedConsumerProjects:        []string{"proj-a"},
		PscAutoConnections: []*api.PscAutoConnectionConfig{
			{
				ConsumerNetwork:       "projects/host-p/global/networks/net-1",
				ConsumerProject:       "svc-p1",
				ConsumerNetworkStatus: "POLICY_CREATED",
				IpAddress:             "10.0.1.20",
				Status:                "ACTIVE",
			},
		},
	}

	krmList := InstancePscConfigGCPToKRM(gcpObj)
	if len(krmList) != 1 {
		t.Fatalf("expected 1 InstancePscConfig, got %d", len(krmList))
	}

	krmPsc := krmList[0]
	if len(krmPsc.PscAutoConnections) != 1 {
		t.Fatalf("expected 1 PscAutoConnection, got %d", len(krmPsc.PscAutoConnections))
	}

	pac := krmPsc.PscAutoConnections[0]
	if pac.ConsumerNetwork == nil || *pac.ConsumerNetwork != "projects/host-p/global/networks/net-1" {
		t.Errorf("ConsumerNetwork mismatch: got %v", pac.ConsumerNetwork)
	}
	if pac.ConsumerProject == nil || *pac.ConsumerProject != "svc-p1" {
		t.Errorf("ConsumerProject mismatch: got %v", pac.ConsumerProject)
	}
	if pac.ConsumerServiceProjectId == nil || *pac.ConsumerServiceProjectId != "svc-p1" {
		t.Errorf("ConsumerServiceProjectId mismatch: got %v", pac.ConsumerServiceProjectId)
	}
	if pac.IpAddress == nil || *pac.IpAddress != "10.0.1.20" {
		t.Errorf("IpAddress mismatch: got %v", pac.IpAddress)
	}
	if pac.Status == nil || *pac.Status != "ACTIVE" {
		t.Errorf("Status mismatch: got %v", pac.Status)
	}
	if pac.ConsumerNetworkStatus == nil || *pac.ConsumerNetworkStatus != "POLICY_CREATED" {
		t.Errorf("ConsumerNetworkStatus mismatch: got %v", pac.ConsumerNetworkStatus)
	}
	if pac.ConsumerNetworkRef == nil || pac.ConsumerNetworkRef.External != "projects/host-p/global/networks/net-1" {
		t.Errorf("ConsumerNetworkRef mismatch: got %v", pac.ConsumerNetworkRef)
	}
}

func TestResolvePscAutoConnectionsRefs(t *testing.T) {
	t.Parallel()

	scheme := runtime.NewScheme()
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	sqlInstance := &krm.SQLInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-sql",
			Namespace: "test-ns",
		},
		Spec: krm.SQLInstanceSpec{
			Settings: krm.InstanceSettings{
				IpConfiguration: &krm.InstanceIpConfiguration{
					PscConfig: []krm.InstancePscConfig{
						{
							PscAutoConnections: []krm.InstancePscAutoConnectionConfig{
								{
									ConsumerNetworkRef: &computerefs.ComputeNetworkRef{
										External: "projects/proj-123/global/networks/my-vpc",
									},
									ConsumerServiceProjectId: direct.PtrTo("svc-project-456"),
								},
							},
						},
					},
				},
			},
		},
	}

	ctx := context.Background()
	if err := resolvePscAutoConnectionsRefs(ctx, fakeClient, sqlInstance); err != nil {
		t.Fatalf("resolvePscAutoConnectionsRefs failed: %v", err)
	}

	conn := sqlInstance.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections[0]
	if conn.ConsumerNetwork == nil || *conn.ConsumerNetwork != "projects/proj-123/global/networks/my-vpc" {
		t.Errorf("ConsumerNetwork was not populated from ConsumerNetworkRef: got %v", conn.ConsumerNetwork)
	}
	if conn.ConsumerProject == nil || *conn.ConsumerProject != "svc-project-456" {
		t.Errorf("ConsumerProject was not populated from ConsumerServiceProjectId: got %v", conn.ConsumerProject)
	}
}
