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
	"context"
	"reflect"
	"strconv"
	"testing"

	pb "cloud.google.com/go/alloydb/apiv1beta/alloydbpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/operations"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	ncpb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type fakeProjectStore struct {
	byID     map[string]*projects.ProjectData
	byNumber map[int64]*projects.ProjectData
}

func newFakeProjectStore(items ...*projects.ProjectData) *fakeProjectStore {
	s := &fakeProjectStore{
		byID:     make(map[string]*projects.ProjectData, len(items)),
		byNumber: make(map[int64]*projects.ProjectData, len(items)),
	}
	for _, p := range items {
		s.byID[p.ID] = p
		s.byNumber[p.Number] = p
	}
	return s
}

func (f *fakeProjectStore) GetProject(project *projects.ProjectName) (*projects.ProjectData, error) {
	if project.ProjectNumber != 0 {
		return f.GetProjectByNumber(strconv.FormatInt(project.ProjectNumber, 10))
	}
	return f.GetProjectByID(project.ProjectID)
}

func (f *fakeProjectStore) GetProjectByID(projectID string) (*projects.ProjectData, error) {
	if p, ok := f.byID[projectID]; ok {
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project %q not found", projectID)
}

func (f *fakeProjectStore) GetProjectByNumber(projectNumberAsString string) (*projects.ProjectData, error) {
	num, err := strconv.ParseInt(projectNumberAsString, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid project number %q", projectNumberAsString)
	}
	if p, ok := f.byNumber[num]; ok {
		return p, nil
	}
	return nil, status.Errorf(codes.NotFound, "project number %d not found", num)
}

func (f *fakeProjectStore) GetProjectByIDOrNumber(projectIDOrNumber string) (*projects.ProjectData, error) {
	if num, err := strconv.ParseInt(projectIDOrNumber, 10, 64); err == nil {
		return f.GetProjectByNumber(strconv.FormatInt(num, 10))
	}
	return f.GetProjectByID(projectIDOrNumber)
}

func newTestAlloyDBServer() (*AlloyDBAdminV1, storage.Storage) {
	st := storage.NewInMemoryStorage()
	projStore := newFakeProjectStore(
		&projects.ProjectData{ID: "my-project", Number: 123456789},
		&projects.ProjectData{ID: "other-project", Number: 987654321},
	)
	env := &common.MockEnvironment{
		Projects: projStore,
	}
	svc := &MockService{
		MockEnvironment: env,
		storage:         st,
		operations:      operations.NewOperationsService(st),
	}
	svc.v1 = &AlloyDBAdminV1{MockService: svc}
	return svc.v1, st
}

func TestSetInstanceFieldsPSCDefaults(t *testing.T) {
	ctx := context.Background()
	srv, _ := newTestAlloyDBServer()
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
	srv.setInstanceFields(ctx, name, inst)

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
	srv.setInstanceFields(ctx, name, inst)

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
	srv.setInstanceFields(ctx, name, psaInst)
	if got := psaInst.GetIpAddress(); got != "10.1.2.3" {
		t.Errorf("got IpAddress = %q for non-PSC instance, want 10.1.2.3", got)
	}
}

func TestCreateInstanceInPSCClusterWithoutPSCInstanceConfig(t *testing.T) {
	ctx := context.Background()
	srv, st := newTestAlloyDBServer()

	clusterFQN := "projects/my-project/locations/europe-west1/clusters/psc-cluster"
	if err := st.Create(ctx, clusterFQN, &pb.Cluster{
		Name: clusterFQN,
		PscConfig: &pb.Cluster_PscConfig{
			PscEnabled: true,
		},
	}); err != nil {
		t.Fatalf("failed to seed PSC cluster: %v", err)
	}

	if _, err := srv.CreateInstance(ctx, &pb.CreateInstanceRequest{
		Parent:     clusterFQN,
		InstanceId: "psc-inst",
		Instance: &pb.Instance{
			InstanceType: pb.Instance_PRIMARY,
			// Intentionally leave PscInstanceConfig nil on a PSC-enabled cluster.
		},
	}); err != nil {
		t.Fatalf("CreateInstance failed: %v", err)
	}

	got, err := srv.GetInstance(ctx, &pb.GetInstanceRequest{
		Name: clusterFQN + "/instances/psc-inst",
	})
	if err != nil {
		t.Fatalf("GetInstance failed: %v", err)
	}
	if got.GetIpAddress() != "" {
		t.Errorf("got IpAddress = %q, want empty on PSC cluster", got.GetIpAddress())
	}
	if got.GetPscInstanceConfig().GetServiceAttachmentLink() == "" {
		t.Errorf("expected ServiceAttachmentLink to be populated on PSC cluster")
	}
	if got.GetPscInstanceConfig().GetPscDnsName() == "" {
		t.Errorf("expected PscDnsName to be populated on PSC cluster")
	}
	if !got.GetPscInstanceInfo().GetEffectivePscAutoDnsEnabled() {
		t.Errorf("got EffectivePscAutoDnsEnabled = false, want true")
	}
	if !got.GetPscInstanceInfo().GetEffectivePscAutoConnectionPolicy() {
		t.Errorf("got EffectivePscAutoConnectionPolicy = false, want true")
	}
}

func TestCreateAndGetInstancePSCServiceConnectionPolicyFlow(t *testing.T) {
	ctx := context.Background()
	srv, st := newTestAlloyDBServer()

	clusterFQN := "projects/my-project/locations/europe-west1/clusters/psc-cluster"
	if err := st.Create(ctx, clusterFQN, &pb.Cluster{
		Name: clusterFQN,
		PscConfig: &pb.Cluster_PscConfig{
			PscEnabled: true,
		},
	}); err != nil {
		t.Fatalf("failed to seed PSC cluster: %v", err)
	}

	instFQN := clusterFQN + "/instances/psc-inst"
	if _, err := srv.CreateInstance(ctx, &pb.CreateInstanceRequest{
		Parent:     clusterFQN,
		InstanceId: "psc-inst",
		Instance: &pb.Instance{
			InstanceType: pb.Instance_PRIMARY,
			PscInstanceConfig: &pb.Instance_PscInstanceConfig{
				AllowedConsumerProjects: []string{"my-project", "123456789"},
				PscAutoConnections: []*pb.Instance_PscAutoConnectionConfig{
					{
						ConsumerProject: "123456789",
						ConsumerNetwork: "projects/123456789/global/networks/my-net",
					},
					{
						ConsumerProject: "my-project",
						ConsumerNetwork: "https://www.googleapis.com/compute/v1/projects/my-project/global/networks/my-net",
					},
				},
			},
		},
	}); err != nil {
		t.Fatalf("CreateInstance failed: %v", err)
	}

	// 1. Without a ServiceConnectionPolicy, ConsumerNetworkStatus should be CONNECTION_POLICY_MISSING.
	gotMissing, err := srv.GetInstance(ctx, &pb.GetInstanceRequest{Name: instFQN})
	if err != nil {
		t.Fatalf("GetInstance failed: %v", err)
	}
	wantProjects := []string{"123456789"}
	if !reflect.DeepEqual(gotMissing.GetPscInstanceConfig().GetAllowedConsumerProjects(), wantProjects) {
		t.Errorf("got AllowedConsumerProjects = %v, want %v", gotMissing.GetPscInstanceConfig().GetAllowedConsumerProjects(), wantProjects)
	}
	conns := gotMissing.GetPscInstanceConfig().GetPscAutoConnections()
	if len(conns) != 1 {
		t.Fatalf("got %d PscAutoConnections, want 1 deduplicated entry", len(conns))
	}
	if conns[0].GetConsumerProject() != "my-project" {
		t.Errorf("got ConsumerProject = %q, want my-project", conns[0].GetConsumerProject())
	}
	if conns[0].GetConsumerNetwork() != "projects/my-project/global/networks/my-net" {
		t.Errorf("got ConsumerNetwork = %q, want projects/my-project/global/networks/my-net", conns[0].GetConsumerNetwork())
	}
	if conns[0].GetConsumerNetworkStatus() != "CONNECTION_POLICY_MISSING" {
		t.Errorf("got ConsumerNetworkStatus = %q, want CONNECTION_POLICY_MISSING", conns[0].GetConsumerNetworkStatus())
	}
	if len(gotMissing.GetPscInstanceInfo().GetPscAutoDnsNames()) != 0 {
		t.Errorf("got PscAutoDnsNames = %v, want empty without policy", gotMissing.GetPscInstanceInfo().GetPscAutoDnsNames())
	}

	// 2. Seed a matching ServiceConnectionPolicy and verify VALID/ACTIVE + PscAutoDnsNames + ServiceConnectionPolicy.
	scpName := "projects/my-project/locations/europe-west1/serviceConnectionPolicies/default-scp"
	if err := st.Create(ctx, scpName, &ncpb.ServiceConnectionPolicy{
		Name:         scpName,
		ServiceClass: "google-cloud-alloydb",
		Network:      "projects/123456789/global/networks/my-net",
	}); err != nil {
		t.Fatalf("failed to seed ServiceConnectionPolicy: %v", err)
	}

	gotActive, err := srv.GetInstance(ctx, &pb.GetInstanceRequest{Name: instFQN})
	if err != nil {
		t.Fatalf("GetInstance with policy failed: %v", err)
	}
	activeConn := gotActive.GetPscInstanceConfig().GetPscAutoConnections()[0]
	if activeConn.GetConsumerNetworkStatus() != "VALID" {
		t.Errorf("got ConsumerNetworkStatus = %q, want VALID", activeConn.GetConsumerNetworkStatus())
	}
	if activeConn.GetStatus() != "ACTIVE" {
		t.Errorf("got Status = %q, want ACTIVE", activeConn.GetStatus())
	}
	if activeConn.GetIpAddress() != "10.0.0.2" {
		t.Errorf("got IpAddress = %q, want 10.0.0.2", activeConn.GetIpAddress())
	}
	if gotActive.GetPscInstanceInfo().GetServiceConnectionPolicy() != scpName {
		t.Errorf("got ServiceConnectionPolicy = %q, want %q", gotActive.GetPscInstanceInfo().GetServiceConnectionPolicy(), scpName)
	}
	wantDNS := []string{"00000000-0000-0000-0000-000000000001.00000000-0000-0000-0000-000000000001.europe-west1.alloydb.goog."}
	if !reflect.DeepEqual(gotActive.GetPscInstanceInfo().GetPscAutoDnsNames(), wantDNS) {
		t.Errorf("got PscAutoDnsNames = %v, want %v", gotActive.GetPscInstanceInfo().GetPscAutoDnsNames(), wantDNS)
	}

	// 3. Disable PSCAutoDNSState and PSCAutoConnectionPolicyState via UpdateInstance.
	if _, err := srv.UpdateInstance(ctx, &pb.UpdateInstanceRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{
				"pscInstanceConfig.pscAutoDnsState",
				"pscInstanceConfig.pscAutoConnectionPolicyState",
			},
		},
		Instance: &pb.Instance{
			Name: instFQN,
			PscInstanceConfig: &pb.Instance_PscInstanceConfig{
				PscAutoDnsState:              pb.Instance_PSC_AUTO_DNS_STATE_DISABLED,
				PscAutoConnectionPolicyState: pb.Instance_PscInstanceConfig_DISABLED,
			},
		},
	}); err != nil {
		t.Fatalf("UpdateInstance failed: %v", err)
	}

	gotDisabled, err := srv.GetInstance(ctx, &pb.GetInstanceRequest{Name: instFQN})
	if err != nil {
		t.Fatalf("GetInstance after update failed: %v", err)
	}
	if gotDisabled.GetPscInstanceInfo().GetEffectivePscAutoDnsEnabled() {
		t.Errorf("got EffectivePscAutoDnsEnabled = true, want false")
	}
	if gotDisabled.GetPscInstanceInfo().GetEffectivePscAutoConnectionPolicy() {
		t.Errorf("got EffectivePscAutoConnectionPolicy = true, want false")
	}
	if len(gotDisabled.GetPscInstanceInfo().GetPscAutoDnsNames()) != 0 {
		t.Errorf("got PscAutoDnsNames = %v, want empty when DNS disabled", gotDisabled.GetPscInstanceInfo().GetPscAutoDnsNames())
	}
	if gotDisabled.GetPscInstanceInfo().GetServiceConnectionPolicy() != "" {
		t.Errorf("got ServiceConnectionPolicy = %q, want empty when policy automation disabled", gotDisabled.GetPscInstanceInfo().GetServiceConnectionPolicy())
	}
}
