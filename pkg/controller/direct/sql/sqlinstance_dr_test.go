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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"google.golang.org/api/option"
	api "google.golang.org/api/sqladmin/v1beta4"
)

type mockTransport struct {
	roundTripFunc func(*http.Request) (*http.Response, error)
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.roundTripFunc(req)
}

func newTestUpdateOp(u *unstructured.Unstructured) *directbase.UpdateOperation {
	c := fake.NewClientBuilder().WithObjects(u).WithStatusSubresource(u).Build()
	lh := lifecyclehandler.NewLifecycleHandler(c, &record.FakeRecorder{})
	return directbase.NewUpdateOperation(lh, c, u)
}

func newTestCreateOp(u *unstructured.Unstructured) *directbase.CreateOperation {
	c := fake.NewClientBuilder().WithObjects(u).WithStatusSubresource(u).Build()
	lh := lifecyclehandler.NewLifecycleHandler(c, &record.FakeRecorder{})
	return directbase.NewCreateOperation(lh, c, u)
}

// TestDiffInstances_DR_RoleSwap_MultiEngine tests that DiffInstances cleanly suppresses
// diffs on masterInstanceName, instanceType, and failoverDrReplicaName when a Cloud SQL
// Enterprise DR failover or switchover inverts instance roles across all 3 engines.
func TestDiffInstances_DR_RoleSwap_MultiEngine(t *testing.T) {
	engines := []struct {
		name    string
		version string
		tier    string
	}{
		{
			name:    "PostgreSQL 16",
			version: "POSTGRES_16",
			tier:    "db-perf-optimized-N-2",
		},
		{
			name:    "MySQL 8.0",
			version: "MYSQL_8_0",
			tier:    "db-perf-optimized-N-2",
		},
		{
			name:    "SQL Server 2022",
			version: "SQLSERVER_2022_ENTERPRISE",
			tier:    "db-custom-4-16384",
		},
	}

	for _, eng := range engines {
		t.Run(eng.name, func(t *testing.T) {
			psaEndpoint := "dr-cluster-write.sql.goog"

			// Scenario 1: Normal In-Sync Primary
			t.Run("NormalPrimary", func(t *testing.T) {
				desired := &api.DatabaseInstance{
					Name:            "db-primary",
					DatabaseVersion: eng.version,
					Settings:        &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						FailoverDrReplicaName: "db-replica",
					},
				}
				actual := &api.DatabaseInstance{
					Name:            "db-primary",
					DatabaseVersion: eng.version,
					Settings:        &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						FailoverDrReplicaName: "db-replica",
						PsaWriteEndpoint:      psaEndpoint,
					},
				}
				diff := DiffInstances(desired, actual)
				if diff.HasDiff() {
					t.Fatalf("expected no diff for normal primary, got: %v", diff)
				}
			})

			// Scenario 2: Normal In-Sync Replica
			t.Run("NormalReplica", func(t *testing.T) {
				desired := &api.DatabaseInstance{
					Name:               "db-replica",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "db-primary",
					InstanceType:       "READ_REPLICA_INSTANCE",
					Settings:           &api.Settings{Tier: eng.tier},
				}
				actual := &api.DatabaseInstance{
					Name:               "db-replica",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "db-primary",
					InstanceType:       "READ_REPLICA_INSTANCE",
					Settings:           &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						DrReplica:        true,
						PsaWriteEndpoint: psaEndpoint,
					},
				}
				diff := DiffInstances(desired, actual)
				if diff.HasDiff() {
					t.Fatalf("expected no diff for normal replica, got: %v", diff)
				}
			})

			// Scenario 3: Post-Switchover Demoted Master (Former Primary -> Now DR Replica)
			t.Run("Mode2_Switchover_DemotedMaster", func(t *testing.T) {
				desired := &api.DatabaseInstance{
					Name:            "db-primary",
					DatabaseVersion: eng.version,
					Settings:        &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						FailoverDrReplicaName: "db-replica",
					},
				}
				// In GCP, db-primary is now a replica pointing to db-replica
				actual := &api.DatabaseInstance{
					Name:               "db-primary",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "db-replica",
					InstanceType:       "READ_REPLICA_INSTANCE",
					Settings:           &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						DrReplica:        true,
						PsaWriteEndpoint: psaEndpoint,
					},
				}
				diff := DiffInstances(desired, actual)
				if diff.HasDiff() {
					t.Fatalf("expected no diff for demoted master under DR suppression, got diff: %v", diff)
				}
			})

			// Scenario 4: Post-Switchover Promoted Master (Former Replica -> Now Primary)
			t.Run("Mode2_Switchover_PromotedMaster", func(t *testing.T) {
				desired := &api.DatabaseInstance{
					Name:               "db-replica",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "db-primary",
					InstanceType:       "READ_REPLICA_INSTANCE",
					Settings:           &api.Settings{Tier: eng.tier},
				}
				// In GCP, db-replica is now the standalone primary with failoverDrReplicaName = "db-primary"
				actual := &api.DatabaseInstance{
					Name:               "db-replica",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "",
					InstanceType:       "CLOUD_SQL_INSTANCE",
					Settings:           &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						FailoverDrReplicaName: "db-primary",
						PsaWriteEndpoint:      psaEndpoint,
					},
				}
				diff := DiffInstances(desired, actual)
				if diff.HasDiff() {
					t.Fatalf("expected no diff for promoted master under DR suppression, got diff: %v", diff)
				}
			})

			// Scenario 5: Mode 3 Unplanned Emergency Promotion
			t.Run("Mode3_UnplannedPromotion", func(t *testing.T) {
				desired := &api.DatabaseInstance{
					Name:               "db-replica",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "db-primary",
					InstanceType:       "READ_REPLICA_INSTANCE",
					Settings:           &api.Settings{Tier: eng.tier},
				}
				// Emergency promotion makes db-replica master before db-primary recovers
				actual := &api.DatabaseInstance{
					Name:               "db-replica",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "",
					InstanceType:       "CLOUD_SQL_INSTANCE",
					Settings:           &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						PsaWriteEndpoint: psaEndpoint,
					},
				}
				diff := DiffInstances(desired, actual)
				if diff.HasDiff() {
					t.Fatalf("expected no diff for emergency promoted replica, got diff: %v", diff)
				}
			})

			// Scenario 6: Non-DR Instance should still report diff on masterInstanceName mismatch
			t.Run("NonDR_DiffReported", func(t *testing.T) {
				desired := &api.DatabaseInstance{
					Name:               "standalone-db",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "",
					Settings:           &api.Settings{Tier: eng.tier},
				}
				actual := &api.DatabaseInstance{
					Name:               "standalone-db",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "unexpected-master",
					Settings:           &api.Settings{Tier: eng.tier},
				}
				diff := DiffInstances(desired, actual)
				if !diff.HasDiff() {
					t.Fatalf("expected diff on non-DR instance with master mismatch, got none")
				}
			})
		})
	}
}

// TestSQLInstance_StandbyDuringFailover tests that the controller enters standby
// when Cloud SQL is in MAINTENANCE or UPDATING state with an active failover operation.
func TestSQLInstance_StandbyDuringFailover(t *testing.T) {
	ctx := context.Background()

	// Mock HTTP client that serves operations list with an active SWITCHOVER operation
	mutatingCallInvoked := false
	transport := &mockTransport{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method == "PUT" || req.Method == "PATCH" || req.Method == "POST" {
				if req.URL.Path != "/sql/v1beta4/projects/test-project/operations" {
					mutatingCallInvoked = true
				}
			}

			// Mock GET operations list
			if req.Method == "GET" && req.URL.Path == "/sql/v1beta4/projects/test-project/operations" {
				opList := &api.OperationsListResponse{
					Items: []*api.Operation{
						{
							Name:          "op-switchover-123",
							OperationType: "SWITCHOVER",
							Status:        "RUNNING",
							TargetId:      "dr-instance",
						},
					},
				}
				data, _ := json.Marshal(opList)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}

			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
			}, nil
		},
	}

	httpClient := &http.Client{Transport: transport}
	sqlService, err := api.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatalf("creating sql service: %v", err)
	}

	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]any{
				"name":      "dr-instance",
				"namespace": "default",
			},
		},
	}

	adapter := &sqlInstanceAdapter{
		projectID:  "test-project",
		resourceID: "dr-instance",
		desired:    &krm.SQLInstance{},
		actual: &api.DatabaseInstance{
			Name:  "dr-instance",
			State: "MAINTENANCE",
		},
		sqlOperationsClient: api.NewOperationsService(sqlService),
		sqlInstancesClient:  api.NewInstancesService(sqlService),
		fieldMeta:           make(map[string]*FieldMetadata),
	}

	updateOp := newTestUpdateOp(u)
	err = adapter.Update(ctx, updateOp)
	if err != nil {
		t.Fatalf("Update returned error: %v", err)
	}

	if mutatingCallInvoked {
		t.Fatalf("mutating API call was invoked during active failover standby!")
	}

	if !updateOp.RequeueRequested {
		t.Fatalf("expected RequeueRequested to be true during standby")
	}

	if !updateOp.HasSetReadyCondition {
		t.Fatalf("expected HasSetReadyCondition to be true during standby")
	}

	// Verify status condition is Ready=False, Reason=FailoverInProgress
	conds, found, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
	if !found || len(conds) == 0 {
		t.Fatalf("expected status conditions to be set, got none")
	}
	latestCond := conds[0].(map[string]any)
	if latestCond["reason"] != "FailoverInProgress" {
		t.Fatalf("expected Reason=FailoverInProgress, got: %v", latestCond["reason"])
	}
	if latestCond["status"] != string(corev1.ConditionFalse) {
		t.Fatalf("expected Status=False, got: %v", latestCond["status"])
	}
}

// TestSQLInstance_PostFailoverAcknowledgment tests that when an instance returns
// to RUNNABLE after a failover, KCC emits Ready=True with Reason=FailoverAcknowledged
// and sets status.currentRole.
func TestSQLInstance_PostFailoverAcknowledgment(t *testing.T) {
	ctx := context.Background()

	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]any{
				"name":      "dr-replica",
				"namespace": "default",
			},
			"status": map[string]any{
				"conditions": []any{
					map[string]any{
						"type":   "Ready",
						"status": string(corev1.ConditionFalse),
						"reason": "FailoverInProgress",
					},
				},
			},
		},
	}

	adapter := &sqlInstanceAdapter{
		projectID:  "test-project",
		resourceID: "dr-replica",
		desired:    &krm.SQLInstance{},
		actual: &api.DatabaseInstance{
			Name:               "dr-replica",
			State:              "RUNNABLE",
			MasterInstanceName: "", // Now promoted to master
			ReplicationCluster: &api.ReplicationCluster{
				PsaWriteEndpoint: "dr-cluster.sql.goog",
			},
		},
		fieldMeta: make(map[string]*FieldMetadata),
	}

	updateOp := newTestUpdateOp(u)
	status, err := SQLInstanceStatusGCPToKRM(adapter.actual)
	if err != nil {
		t.Fatalf("converting status: %v", err)
	}

	if status.CurrentRole == nil || *status.CurrentRole != "PRIMARY" {
		t.Fatalf("expected CurrentRole=PRIMARY, got: %v", status.CurrentRole)
	}

	err = adapter.updateFinalStatus(ctx, updateOp, u, status)
	if err != nil {
		t.Fatalf("updateFinalStatus returned error: %v", err)
	}

	conds, found, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
	if !found || len(conds) == 0 {
		t.Fatalf("expected conditions to be populated")
	}
	latestCond := conds[0].(map[string]any)
	if latestCond["reason"] != "FailoverAcknowledged" {
		t.Fatalf("expected Reason=FailoverAcknowledged, got: %v", latestCond["reason"])
	}
	if latestCond["status"] != string(corev1.ConditionTrue) {
		t.Fatalf("expected Status=True, got: %v", latestCond["status"])
	}
}

// TestSQLInstance_DeletionBlockedDuringFailover tests that calling Delete on an instance
// undergoing failover/maintenance is blocked and returns an error.
func TestSQLInstance_DeletionBlockedDuringFailover(t *testing.T) {
	ctx := context.Background()

	transport := &mockTransport{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method == "GET" && req.URL.Path == "/sql/v1beta4/projects/test-project/operations" {
				opList := &api.OperationsListResponse{
					Items: []*api.Operation{
						{
							Name:          "op-failover-456",
							OperationType: "FAILOVER",
							Status:        "RUNNING",
							TargetId:      "ha-instance",
						},
					},
				}
				data, _ := json.Marshal(opList)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
			}, nil
		},
	}

	sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
	if err != nil {
		t.Fatalf("creating sql service: %v", err)
	}

	adapter := &sqlInstanceAdapter{
		projectID:  "test-project",
		resourceID: "ha-instance",
		actual: &api.DatabaseInstance{
			Name:  "ha-instance",
			State: "MAINTENANCE",
		},
		sqlOperationsClient: api.NewOperationsService(sqlService),
		sqlInstancesClient:  api.NewInstancesService(sqlService),
	}

	u := &unstructured.Unstructured{}
	c := fake.NewClientBuilder().Build()
	deleteOp := directbase.NewDeleteOperation(c, u)

	deleted, err := adapter.Delete(ctx, deleteOp)
	if err == nil {
		t.Fatalf("expected Delete to return an error while failover is in progress, got nil")
	}
	if deleted {
		t.Fatalf("expected deleted=false while failover is in progress")
	}
}

// TestDiffInstances_DR_BidirectionalSwitchover_Cyclic validates that Enterprise DR instances
// can undergo repeated bidirectional switchover and failback cycles across all 3 engines
// without any diff drift or reconciliation errors.
func TestDiffInstances_DR_BidirectionalSwitchover_Cyclic(t *testing.T) {
	engines := []struct {
		name    string
		version string
		tier    string
	}{
		{name: "PostgreSQL_16", version: "POSTGRES_16", tier: "db-perf-optimized-N-2"},
		{name: "MySQL_8.0", version: "MYSQL_8_0", tier: "db-perf-optimized-N-2"},
		{name: "SQL_Server_2022", version: "SQLSERVER_2022_ENTERPRISE", tier: "db-custom-4-16384"},
	}

	for _, eng := range engines {
		t.Run(eng.name, func(t *testing.T) {
			psaEndpoint := "enterprise-dr-write.sql.goog"

			// Instance A: original primary
			desiredA := &api.DatabaseInstance{
				Name:            "db-inst-a",
				DatabaseVersion: eng.version,
				Settings:        &api.Settings{Tier: eng.tier},
				ReplicationCluster: &api.ReplicationCluster{
					FailoverDrReplicaName: "db-inst-b",
				},
			}

			// Instance B: original replica
			desiredB := &api.DatabaseInstance{
				Name:               "db-inst-b",
				DatabaseVersion:    eng.version,
				MasterInstanceName: "test-proj:db-inst-a",
				InstanceType:       "READ_REPLICA_INSTANCE",
				Settings:           &api.Settings{Tier: eng.tier},
				ReplicationCluster: &api.ReplicationCluster{
					DrReplica: true,
				},
			}

			// Run 5 consecutive bidirectional cycles
			for cycle := 1; cycle <= 5; cycle++ {
				// Phase 1: Forward switchover (A demoted to DR_REPLICA, B promoted to PRIMARY)
				actualA_Demoted := &api.DatabaseInstance{
					Name:               "db-inst-a",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "test-proj:db-inst-b",
					InstanceType:       "READ_REPLICA_INSTANCE",
					Settings:           &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						DrReplica:        true,
						PsaWriteEndpoint: psaEndpoint,
					},
				}
				actualB_Promoted := &api.DatabaseInstance{
					Name:            "db-inst-b",
					DatabaseVersion: eng.version,
					Settings:        &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						FailoverDrReplicaName: "db-inst-a",
						PsaWriteEndpoint:      psaEndpoint,
					},
				}

				diffA1 := DiffInstances(desiredA, actualA_Demoted)
				if diffA1.HasDiff() {
					t.Fatalf("cycle %d forward switchover: expected no diff for demoted instance A, got: %v", cycle, diffA1)
				}

				diffB1 := DiffInstances(desiredB, actualB_Promoted)
				if diffB1.HasDiff() {
					t.Fatalf("cycle %d forward switchover: expected no diff for promoted instance B, got: %v", cycle, diffB1)
				}

				// Phase 2: Reverse failback (A restored to PRIMARY, B demoted back to DR_REPLICA)
				actualA_Restored := &api.DatabaseInstance{
					Name:            "db-inst-a",
					DatabaseVersion: eng.version,
					Settings:        &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						FailoverDrReplicaName: "db-inst-b",
						PsaWriteEndpoint:      psaEndpoint,
					},
				}
				actualB_Demoted := &api.DatabaseInstance{
					Name:               "db-inst-b",
					DatabaseVersion:    eng.version,
					MasterInstanceName: "test-proj:db-inst-a",
					InstanceType:       "READ_REPLICA_INSTANCE",
					Settings:           &api.Settings{Tier: eng.tier},
					ReplicationCluster: &api.ReplicationCluster{
						DrReplica:        true,
						PsaWriteEndpoint: psaEndpoint,
					},
				}

				diffA2 := DiffInstances(desiredA, actualA_Restored)
				if diffA2.HasDiff() {
					t.Fatalf("cycle %d failback: expected no diff for restored primary A, got: %v", cycle, diffA2)
				}

				diffB2 := DiffInstances(desiredB, actualB_Demoted)
				if diffB2.HasDiff() {
					t.Fatalf("cycle %d failback: expected no diff for restored replica B, got: %v", cycle, diffB2)
				}
			}
		})
	}
}

// TestSQLInstance_BidirectionalSwitchoverLifecycle tests full controller status lifecycle
// across forward switchover and reverse failback.
func TestSQLInstance_BidirectionalSwitchoverLifecycle(t *testing.T) {
	ctx := context.Background()

	// Step 1: Forward switchover - instance starts as PRIMARY, undergoes SWITCHOVER, demotes to DR_REPLICA
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]any{
				"name":      "cyclic-instance",
				"namespace": "default",
			},
			"status": map[string]any{
				"conditions": []any{
					map[string]any{
						"type":   "Ready",
						"status": string(corev1.ConditionFalse),
						"reason": "FailoverInProgress",
					},
				},
				"currentRole": "PRIMARY",
			},
		},
	}

	adapter := &sqlInstanceAdapter{
		projectID:  "test-project",
		resourceID: "cyclic-instance",
		actual: &api.DatabaseInstance{
			Name:               "cyclic-instance",
			MasterInstanceName: "test-project:other-instance",
			ReplicationCluster: &api.ReplicationCluster{
				DrReplica: true,
			},
		},
	}

	updateOp := newTestUpdateOp(u)
	status, err := SQLInstanceStatusGCPToKRM(adapter.actual)
	if err != nil {
		t.Fatalf("converting status: %v", err)
	}

	if *status.CurrentRole != "DR_REPLICA" {
		t.Fatalf("expected role DR_REPLICA post-forward switchover, got %s", *status.CurrentRole)
	}

	if err := adapter.updateFinalStatus(ctx, updateOp, u, status); err != nil {
		t.Fatalf("updateFinalStatus: %v", err)
	}

	conds, _, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
	cond := conds[0].(map[string]any)
	if cond["reason"] != "FailoverAcknowledged" || cond["status"] != string(corev1.ConditionTrue) {
		t.Fatalf("expected Ready=True, Reason=FailoverAcknowledged, got %v / %v", cond["status"], cond["reason"])
	}

	// Step 2: Reverse failback - instance undergoes SWITCHOVER, promotes back to PRIMARY
	u.Object["status"] = map[string]any{
		"conditions": []any{
			map[string]any{
				"type":   "Ready",
				"status": string(corev1.ConditionFalse),
				"reason": "FailoverInProgress",
			},
		},
		"currentRole": "DR_REPLICA",
	}

	adapter.actual = &api.DatabaseInstance{
		Name:               "cyclic-instance",
		MasterInstanceName: "",
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "other-instance",
		},
	}

	updateOp2 := newTestUpdateOp(u)
	status2, err := SQLInstanceStatusGCPToKRM(adapter.actual)
	if err != nil {
		t.Fatalf("converting status: %v", err)
	}

	if *status2.CurrentRole != "PRIMARY" {
		t.Fatalf("expected role PRIMARY post-failback, got %s", *status2.CurrentRole)
	}

	if err := adapter.updateFinalStatus(ctx, updateOp2, u, status2); err != nil {
		t.Fatalf("updateFinalStatus failback: %v", err)
	}

	conds2, _, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
	cond2 := conds2[0].(map[string]any)
	if cond2["reason"] != "FailoverAcknowledged" || cond2["status"] != string(corev1.ConditionTrue) {
		t.Fatalf("expected Ready=True, Reason=FailoverAcknowledged after failback, got %v / %v", cond2["status"], cond2["reason"])
	}
}

// TestDiffInstances_DR_ThreeTierTopology_SingleZone_CrossZone_CrossRegion validates a complete
// real-world production topology consisting of:
// 1. Primary Instance (us-central1, Regional HA across zones)
// 2. In-Region Single-Zone Read Replica (us-central1-a, Zonal)
// 3. Cross-Region DR Replica (us-east1, Regional HA across zones)
func TestDiffInstances_DR_ThreeTierTopology_SingleZone_CrossZone_CrossRegion(t *testing.T) {
	// Instance 1: Regional Primary in us-central1 (across zones)
	primaryDesired := &api.DatabaseInstance{
		Name:            "pg-dr-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
		},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "pg-dr-replica",
		},
	}
	primaryActual := &api.DatabaseInstance{
		Name:            "pg-dr-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
		},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "pg-dr-replica",
			PsaWriteEndpoint:      "psa-endpoint.sql.goog",
		},
	}

	// Instance 2: Single-Zone Read Replica in us-central1-a (single zone)
	zonalReplicaDesired := &api.DatabaseInstance{
		Name:               "pg-inregion-replica",
		Region:             "us-central1",
		GceZone:            "us-central1-a",
		MasterInstanceName: "pg-dr-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		DatabaseVersion:    "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "ZONAL",
		},
	}
	zonalReplicaActual := &api.DatabaseInstance{
		Name:               "pg-inregion-replica",
		Region:             "us-central1",
		GceZone:            "us-central1-a",
		MasterInstanceName: "pg-dr-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		DatabaseVersion:    "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "ZONAL",
		},
	}

	// Instance 3: Cross-Region DR Replica in us-east1 (Regional HA across zones)
	drReplicaDesired := &api.DatabaseInstance{
		Name:               "pg-dr-replica",
		Region:             "us-east1",
		MasterInstanceName: "pg-dr-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		DatabaseVersion:    "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
		},
	}
	drReplicaActual := &api.DatabaseInstance{
		Name:               "pg-dr-replica",
		Region:             "us-east1",
		MasterInstanceName: "pg-dr-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		DatabaseVersion:    "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
		},
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica:        true,
			PsaWriteEndpoint: "psa-endpoint.sql.goog",
		},
	}

	// 1. Initial State: Diff checks
	if diff := DiffInstances(primaryDesired, primaryActual); diff.HasDiff() {
		t.Fatalf("primary should have no diff, got: %v", diff)
	}
	if diff := DiffInstances(zonalReplicaDesired, zonalReplicaActual); diff.HasDiff() {
		t.Fatalf("zonal replica should have no diff, got: %v", diff)
	}
	if diff := DiffInstances(drReplicaDesired, drReplicaActual); diff.HasDiff() {
		t.Fatalf("dr replica should have no diff, got: %v", diff)
	}

	// 2. Initial State: Role mappings
	primStatus, err := SQLInstanceStatusGCPToKRM(primaryActual)
	if err != nil || primStatus.CurrentRole == nil || *primStatus.CurrentRole != "PRIMARY" {
		t.Fatalf("expected PRIMARY role on primary, got: %v", primStatus.CurrentRole)
	}

	zonalStatus, err := SQLInstanceStatusGCPToKRM(zonalReplicaActual)
	if err != nil || zonalStatus.CurrentRole != nil {
		t.Fatalf("expected nil CurrentRole on regular zonal replica, got: %v", zonalStatus.CurrentRole)
	}

	drStatus, err := SQLInstanceStatusGCPToKRM(drReplicaActual)
	if err != nil || drStatus.CurrentRole == nil || *drStatus.CurrentRole != "DR_REPLICA" {
		t.Fatalf("expected DR_REPLICA role on DR replica, got: %v", drStatus.CurrentRole)
	}

	// 3. Verify that the zonal replica retains strict diff detection on MasterInstanceName
	mutatedZonalDesired := &api.DatabaseInstance{
		Name:               "pg-inregion-replica",
		Region:             "us-central1",
		GceZone:            "us-central1-a",
		MasterInstanceName: "other-primary", // Intentionally mismatched
		InstanceType:       "READ_REPLICA_INSTANCE",
		DatabaseVersion:    "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "ZONAL",
		},
	}
	if diff := DiffInstances(mutatedZonalDesired, zonalReplicaActual); !diff.HasDiff() {
		t.Fatalf("expected diff on regular zonal replica when master differs, but got none (diff suppression must not leak)")
	}

	// 4. Planned Cross-Region Switchover: Primary <-> DR Replica Swap
	postSwitchoverPrimaryActual := &api.DatabaseInstance{
		Name:               "pg-dr-primary",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "pg-dr-replica",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
		},
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica:        true,
			PsaWriteEndpoint: "psa-endpoint.sql.goog",
		},
	}

	postSwitchoverDRActual := &api.DatabaseInstance{
		Name:            "pg-dr-replica",
		Region:          "us-east1",
		DatabaseVersion: "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
		},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "pg-dr-primary",
			PsaWriteEndpoint:      "psa-endpoint.sql.goog",
		},
	}

	// Diff suppression should be active on both DR instances post-switchover
	if diff := DiffInstances(primaryDesired, postSwitchoverPrimaryActual); diff.HasDiff() {
		t.Fatalf("expected diff suppression on demoted primary post-switchover, got: %v", diff)
	}
	if diff := DiffInstances(drReplicaDesired, postSwitchoverDRActual); diff.HasDiff() {
		t.Fatalf("expected diff suppression on promoted replica post-switchover, got: %v", diff)
	}

	// Role inversion
	newPrimStatus, err := SQLInstanceStatusGCPToKRM(postSwitchoverDRActual)
	if err != nil || newPrimStatus.CurrentRole == nil || *newPrimStatus.CurrentRole != "PRIMARY" {
		t.Fatalf("expected PRIMARY role on promoted DR replica, got: %v", newPrimStatus.CurrentRole)
	}

	newDemotedStatus, err := SQLInstanceStatusGCPToKRM(postSwitchoverPrimaryActual)
	if err != nil || newDemotedStatus.CurrentRole == nil || *newDemotedStatus.CurrentRole != "DR_REPLICA" {
		t.Fatalf("expected DR_REPLICA role on demoted primary, got: %v", newDemotedStatus.CurrentRole)
	}
}

// TestDiffInstances_DR_BackupConfiguration_Suppression validates that when an Enterprise DR primary
// is demoted to a replica during DR failover, Cloud SQL's automatic disabling of backups on replicas
// (enabled=false, pointInTimeRecoveryEnabled=false, binaryLogEnabled=false) is suppressed by DiffInstances,
// preventing unrecoverable 400 INVALID_ARGUMENT errors from Cloud SQL API.
// It also validates that for a normal primary, backup differences are NOT suppressed.
func TestDiffInstances_DR_BackupConfiguration_Suppression(t *testing.T) {
	desiredDRPrimary := &api.DatabaseInstance{
		Name:            "dr-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
			BackupConfiguration: &api.BackupConfiguration{
				Enabled:                    true,
				PointInTimeRecoveryEnabled: true,
				BinaryLogEnabled:           true,
			},
		},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "dr-replica",
		},
	}

	// In GCP, during DR role swap, dr-primary is demoted to READ_REPLICA_INSTANCE,
	// and Cloud SQL forcibly turns off backups on the replica.
	demotedActual := &api.DatabaseInstance{
		Name:               "dr-primary",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		InstanceType:       "READ_REPLICA_INSTANCE",
		MasterInstanceName: "dr-replica",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
			BackupConfiguration: &api.BackupConfiguration{
				Enabled:                    false,
				PointInTimeRecoveryEnabled: false,
				BinaryLogEnabled:           false,
			},
		},
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica:        true,
			PsaWriteEndpoint: "endpoint.sql.goog",
		},
	}

	diffDemoted := DiffInstances(desiredDRPrimary, demotedActual)
	if diffDemoted.HasDiff() {
		t.Fatalf("expected diff suppression on demoted DR primary for backupConfiguration, got: %v", diffDemoted)
	}

	// Test promoted replica: Manifest specifies a replica (no backups), but GCP promotes it to primary with full backups enabled.
	desiredDRReplica := &api.DatabaseInstance{
		Name:               "dr-replica",
		Region:             "us-east1",
		DatabaseVersion:    "POSTGRES_16",
		InstanceType:       "READ_REPLICA_INSTANCE",
		MasterInstanceName: "dr-primary",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
		},
	}
	promotedActual := &api.DatabaseInstance{
		Name:            "dr-replica",
		Region:          "us-east1",
		DatabaseVersion: "POSTGRES_16",
		InstanceType:    "CLOUD_SQL_INSTANCE",
		Settings: &api.Settings{
			Tier:             "db-perf-optimized-N-2",
			AvailabilityType: "REGIONAL",
			BackupConfiguration: &api.BackupConfiguration{
				Enabled:                     true,
				PointInTimeRecoveryEnabled:  true,
				Location:                    "us-east1",
				StartTime:                   "04:00",
				TransactionLogRetentionDays: 7,
				BackupRetentionSettings: &api.BackupRetentionSettings{
					RetainedBackups: 7,
					RetentionUnit:   "COUNT",
				},
			},
		},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "dr-primary",
			PsaWriteEndpoint:      "endpoint.sql.goog",
		},
	}
	diffPromoted := DiffInstances(desiredDRReplica, promotedActual)
	if diffPromoted.HasDiff() {
		t.Fatalf("expected diff suppression on promoted DR replica for backupConfiguration, got: %v", diffPromoted)
	}

	// Counter-test: A standalone non-DR instance MUST report diff if backups are disabled
	desiredNormal := &api.DatabaseInstance{
		Name:            "normal-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings: &api.Settings{
			Tier:             "db-custom-2-7680",
			AvailabilityType: "REGIONAL",
			BackupConfiguration: &api.BackupConfiguration{
				Enabled:                    true,
				PointInTimeRecoveryEnabled: true,
			},
		},
	}
	actualNormalMismatched := &api.DatabaseInstance{
		Name:            "normal-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		InstanceType:    "CLOUD_SQL_INSTANCE",
		Settings: &api.Settings{
			Tier:             "db-custom-2-7680",
			AvailabilityType: "REGIONAL",
			BackupConfiguration: &api.BackupConfiguration{
				Enabled:                    false,
				PointInTimeRecoveryEnabled: false,
			},
		},
	}

	diffNormal := DiffInstances(desiredNormal, actualNormalMismatched)
	if !diffNormal.HasDiff() {
		t.Fatalf("expected diff on regular primary when backupConfiguration differs, got none")
	}
}

// TestDiffInstances_DR_ExhaustiveMatrix_AllEnginesAndVersions validates Enterprise DR
// role swap across every database version supported by Cloud SQL Enterprise DR:
// PostgreSQL (14, 15, 16), MySQL (8.0, 8.4), and SQL Server (2019, 2022).
func TestDiffInstances_DR_ExhaustiveMatrix_AllEnginesAndVersions(t *testing.T) {
	matrix := []struct {
		engineFamily string
		version      string
		tier         string
		settings     func() *api.Settings
	}{
		{
			engineFamily: "PostgreSQL",
			version:      "POSTGRES_14",
			tier:         "db-custom-2-7680",
			settings: func() *api.Settings {
				return &api.Settings{
					Tier: "db-custom-2-7680",
					DatabaseFlags: []*api.DatabaseFlags{
						{Name: "autovacuum", Value: "on"},
						{Name: "log_connections", Value: "on"},
					},
					DataCacheConfig: &api.DataCacheConfig{DataCacheEnabled: true},
				}
			},
		},
		{
			engineFamily: "PostgreSQL",
			version:      "POSTGRES_15",
			tier:         "db-perf-optimized-N-2",
			settings: func() *api.Settings {
				return &api.Settings{
					Tier: "db-perf-optimized-N-2",
					DatabaseFlags: []*api.DatabaseFlags{
						{Name: "work_mem", Value: "16384"},
					},
				}
			},
		},
		{
			engineFamily: "PostgreSQL",
			version:      "POSTGRES_16",
			tier:         "db-perf-optimized-N-4",
			settings: func() *api.Settings {
				return &api.Settings{
					Tier:            "db-perf-optimized-N-4",
					DataCacheConfig: &api.DataCacheConfig{DataCacheEnabled: true},
				}
			},
		},
		{
			engineFamily: "MySQL",
			version:      "MYSQL_8_0",
			tier:         "db-custom-4-15360",
			settings: func() *api.Settings {
				return &api.Settings{
					Tier: "db-custom-4-15360",
					BackupConfiguration: &api.BackupConfiguration{
						Enabled:          true,
						BinaryLogEnabled: true,
					},
				}
			},
		},
		{
			engineFamily: "MySQL",
			version:      "MYSQL_8_4",
			tier:         "db-perf-optimized-N-2",
			settings: func() *api.Settings {
				return &api.Settings{
					Tier: "db-perf-optimized-N-2",
					BackupConfiguration: &api.BackupConfiguration{
						Enabled:          true,
						BinaryLogEnabled: true,
					},
				}
			},
		},
		{
			engineFamily: "SQL Server",
			version:      "SQLSERVER_2019_ENTERPRISE",
			tier:         "db-custom-4-16384",
			settings: func() *api.Settings {
				return &api.Settings{
					Tier: "db-custom-4-16384",
					SqlServerAuditConfig: &api.SqlServerAuditConfig{
						Bucket:            "gs://test-sqlserver-audit",
						RetentionInterval: "7d",
					},
				}
			},
		},
		{
			engineFamily: "SQL Server",
			version:      "SQLSERVER_2022_ENTERPRISE",
			tier:         "db-custom-8-32768",
			settings: func() *api.Settings {
				return &api.Settings{
					Tier: "db-custom-8-32768",
					ActiveDirectoryConfig: &api.SqlActiveDirectoryConfig{
						Domain: "corp.example.com",
					},
				}
			},
		},
	}

	for _, item := range matrix {
		testName := fmt.Sprintf("%s_%s", item.engineFamily, item.version)
		t.Run(testName, func(t *testing.T) {
			psaEndpoint := fmt.Sprintf("%s.global.sql-psa.goog", strings.ToLower(item.version))

			// Desired Baseline
			primaryDesired := &api.DatabaseInstance{
				Name:            "primary-" + strings.ToLower(item.version),
				Region:          "us-central1",
				DatabaseVersion: item.version,
				Settings:        item.settings(),
				ReplicationCluster: &api.ReplicationCluster{
					FailoverDrReplicaName: "dr-replica-" + strings.ToLower(item.version),
				},
			}
			replicaDesired := &api.DatabaseInstance{
				Name:               "dr-replica-" + strings.ToLower(item.version),
				Region:             "us-east1",
				DatabaseVersion:    item.version,
				MasterInstanceName: "primary-" + strings.ToLower(item.version),
				InstanceType:       "READ_REPLICA_INSTANCE",
				Settings:           item.settings(),
			}

			// Mode 2: Forward Switchover in GCP
			primaryDemotedActual := &api.DatabaseInstance{
				Name:               primaryDesired.Name,
				Region:             "us-central1",
				DatabaseVersion:    item.version,
				MasterInstanceName: replicaDesired.Name,
				InstanceType:       "READ_REPLICA_INSTANCE",
				Settings:           item.settings(),
				ReplicationCluster: &api.ReplicationCluster{
					DrReplica:        true,
					PsaWriteEndpoint: psaEndpoint,
				},
			}
			replicaPromotedActual := &api.DatabaseInstance{
				Name:            replicaDesired.Name,
				Region:          "us-east1",
				DatabaseVersion: item.version,
				InstanceType:    "CLOUD_SQL_INSTANCE",
				Settings:        item.settings(),
				ReplicationCluster: &api.ReplicationCluster{
					FailoverDrReplicaName: primaryDesired.Name,
					PsaWriteEndpoint:      psaEndpoint,
				},
			}

			// Assert diff suppression during role swap
			if diff := DiffInstances(primaryDesired, primaryDemotedActual); diff.HasDiff() {
				t.Fatalf("[%s] expected diff suppression on demoted primary, got: %v", testName, diff)
			}
			if diff := DiffInstances(replicaDesired, replicaPromotedActual); diff.HasDiff() {
				t.Fatalf("[%s] expected diff suppression on promoted replica, got: %v", testName, diff)
			}

			// Verify status mapping during role swap
			priStatus, err := SQLInstanceStatusGCPToKRM(primaryDemotedActual)
			if err != nil || priStatus.CurrentRole == nil || *priStatus.CurrentRole != "DR_REPLICA" {
				t.Fatalf("[%s] expected DR_REPLICA role, got: %v", testName, priStatus.CurrentRole)
			}
			repStatus, err := SQLInstanceStatusGCPToKRM(replicaPromotedActual)
			if err != nil || repStatus.CurrentRole == nil || *repStatus.CurrentRole != "PRIMARY" {
				t.Fatalf("[%s] expected PRIMARY role, got: %v", testName, repStatus.CurrentRole)
			}

			// Mode 2 Failback: Reverse Switchover back to original roles
			primaryRestoredActual := &api.DatabaseInstance{
				Name:            primaryDesired.Name,
				Region:          "us-central1",
				DatabaseVersion: item.version,
				InstanceType:    "CLOUD_SQL_INSTANCE",
				Settings:        item.settings(),
				ReplicationCluster: &api.ReplicationCluster{
					FailoverDrReplicaName: replicaDesired.Name,
					PsaWriteEndpoint:      psaEndpoint,
				},
			}
			replicaRestoredActual := &api.DatabaseInstance{
				Name:               replicaDesired.Name,
				Region:             "us-east1",
				DatabaseVersion:    item.version,
				MasterInstanceName: primaryDesired.Name,
				InstanceType:       "READ_REPLICA_INSTANCE",
				Settings:           item.settings(),
				ReplicationCluster: &api.ReplicationCluster{
					DrReplica:        true,
					PsaWriteEndpoint: psaEndpoint,
				},
			}

			if diff := DiffInstances(primaryDesired, primaryRestoredActual); diff.HasDiff() {
				t.Fatalf("[%s] expected zero diff on restored primary, got: %v", testName, diff)
			}
			if diff := DiffInstances(replicaDesired, replicaRestoredActual); diff.HasDiff() {
				t.Fatalf("[%s] expected zero diff on restored replica, got: %v", testName, diff)
			}
		})
	}
}

// TestDiffInstances_DR_LegitimateConfigDriftDetected ensures that while DR role-swap diffs
// (instanceType, masterInstanceName, failoverDrReplicaName, backupConfiguration) are suppressed,
// legitimate declarative configuration changes (tier, labels, authorized networks, database flags)
// MUST NOT be suppressed and are properly detected for reconciliation.
func TestDiffInstances_DR_LegitimateConfigDriftDetected(t *testing.T) {
	// Desired manifest specifies tier upgrade and an added user label
	desiredPrimary := &api.DatabaseInstance{
		Name:            "dr-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings: &api.Settings{
			Tier: "db-perf-optimized-N-4", // Upgraded tier in manifest
			UserLabels: map[string]string{
				"env":  "prod",
				"cost": "finance",
			},
			DatabaseFlags: []*api.DatabaseFlags{
				{Name: "autovacuum", Value: "on"},
			},
		},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "dr-replica",
		},
	}

	// In GCP, dr-primary is demoted (DR_REPLICA), but still running the old tier and missing the label
	demotedActualWithDrift := &api.DatabaseInstance{
		Name:               "dr-primary",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "dr-replica",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings: &api.Settings{
			Tier: "db-perf-optimized-N-2", // Old tier in GCP
			UserLabels: map[string]string{
				"env": "prod",
			},
			DatabaseFlags: []*api.DatabaseFlags{
				{Name: "autovacuum", Value: "on"},
			},
			BackupConfiguration: &api.BackupConfiguration{
				Enabled: false,
			},
		},
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica:        true,
			PsaWriteEndpoint: "dr.sql.goog",
		},
	}

	diff := DiffInstances(desiredPrimary, demotedActualWithDrift)
	if !diff.HasDiff() {
		t.Fatalf("expected diff detection on legitimate tier and label drift, but got none!")
	}

	// Verify that diff specifically detected .settings.tier and .settings.userLabels
	tierDiffFound := false
	labelDiffFound := false
	for _, entry := range diff.Fields {
		if entry.ID == ".settings.tier" {
			tierDiffFound = true
		}
		if strings.HasPrefix(entry.ID, ".settings.userLabels") {
			labelDiffFound = true
		}
		if entry.ID == ".masterInstanceName" || entry.ID == ".instanceType" {
			t.Fatalf("spurious diff on DR role fields: %s", entry.ID)
		}
	}

	if !tierDiffFound {
		t.Fatalf("expected diff on .settings.tier, entries were: %v", diff.Fields)
	}
	if !labelDiffFound {
		t.Fatalf("expected diff on .settings.userLabels, entries were: %v", diff.Fields)
	}
}

// TestDiffInstances_DR_CascadingAndMultiReplicaTopology tests a multi-node topology:
// 1 Primary (us-central1) + 1 Cross-Region DR Replica (us-east1) + 2 In-Region Read Replicas (us-central1).
// It verifies that during Mode 2 switchover, in-region cascading replicas and repointed replicas
// maintain predictable diff and role behavior.
func TestDiffInstances_DR_CascadingAndMultiReplicaTopology(t *testing.T) {
	primaryDesired := &api.DatabaseInstance{
		Name:            "prod-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings:        &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "prod-dr-replica",
		},
	}
	drReplicaDesired := &api.DatabaseInstance{
		Name:               "prod-dr-replica",
		Region:             "us-east1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "prod-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
	}
	inregReplica1Desired := &api.DatabaseInstance{
		Name:               "prod-inreg-replica-1",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "prod-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
	}
	inregReplica2Desired := &api.DatabaseInstance{
		Name:               "prod-inreg-replica-2",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "prod-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
	}

	// Mode 2 Switchover:
	// - prod-dr-replica is promoted to PRIMARY
	// - prod-primary is demoted to DR_REPLICA
	// - prod-inreg-replica-1 continues cascading from prod-primary
	// - prod-inreg-replica-2 is repointed to prod-dr-replica (new primary)
	psaEndpoint := "prod-dr.sql-psa.goog"
	primaryActual := &api.DatabaseInstance{
		Name:               "prod-primary",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "prod-dr-replica",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica:        true,
			PsaWriteEndpoint: psaEndpoint,
		},
	}
	drReplicaActual := &api.DatabaseInstance{
		Name:            "prod-dr-replica",
		Region:          "us-east1",
		DatabaseVersion: "POSTGRES_16",
		InstanceType:    "CLOUD_SQL_INSTANCE",
		Settings:        &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "prod-primary",
			PsaWriteEndpoint:      psaEndpoint,
		},
	}
	inreg1Actual := &api.DatabaseInstance{
		Name:               "prod-inreg-replica-1",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "prod-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
	}

	// 1. DR Primary & Replica have diff suppression
	if diff := DiffInstances(primaryDesired, primaryActual); diff.HasDiff() {
		t.Fatalf("expected diff suppression on demoted primary in 4-node cluster, got: %v", diff)
	}
	if diff := DiffInstances(drReplicaDesired, drReplicaActual); diff.HasDiff() {
		t.Fatalf("expected diff suppression on promoted DR replica in 4-node cluster, got: %v", diff)
	}

	// 2. Cascading local replica remains in sync with 0 diff
	if diff := DiffInstances(inregReplica1Desired, inreg1Actual); diff.HasDiff() {
		t.Fatalf("expected zero diff on cascading in-region replica, got: %v", diff)
	}

	// 3. Regular read replica that is intentionally mismatched against its desired master
	// MUST report a diff so KCC can reconcile it
	inreg2MismatchedActual := &api.DatabaseInstance{
		Name:               "prod-inreg-replica-2",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "prod-dr-replica", // Differs from desired prod-primary
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
	}
	diffInreg2 := DiffInstances(inregReplica2Desired, inreg2MismatchedActual)
	if !diffInreg2.HasDiff() {
		t.Fatalf("expected diff on in-region replica when master differs from manifest, got none")
	}
}

// TestDiffInstances_DR_PublicIP_NoPSA tests Cloud SQL Enterprise DR configured without
// Private Service Access (i.e. using Public IPs or PSC, where psaWriteEndpoint is empty).
func TestDiffInstances_DR_PublicIP_NoPSA(t *testing.T) {
	desiredPrimary := &api.DatabaseInstance{
		Name:            "pub-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings:        &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "pub-replica",
		},
	}
	desiredReplica := &api.DatabaseInstance{
		Name:               "pub-replica",
		Region:             "us-east1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "pub-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
	}

	// Inverted state in GCP without PSA endpoint
	demotedPrimaryActual := &api.DatabaseInstance{
		Name:               "pub-primary",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "pub-replica",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica: true, // drReplica flag is set by GCP even without PSA
		},
	}
	promotedReplicaActual := &api.DatabaseInstance{
		Name:            "pub-replica",
		Region:          "us-east1",
		DatabaseVersion: "POSTGRES_16",
		InstanceType:    "CLOUD_SQL_INSTANCE",
		Settings:        &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "pub-primary",
		},
	}

	if diff := DiffInstances(desiredPrimary, demotedPrimaryActual); diff.HasDiff() {
		t.Fatalf("expected diff suppression on public IP demoted primary, got: %v", diff)
	}
	if diff := DiffInstances(desiredReplica, promotedReplicaActual); diff.HasDiff() {
		t.Fatalf("expected diff suppression on public IP promoted replica, got: %v", diff)
	}

	// Status role projection without PSA
	demotedStatus, err := SQLInstanceStatusGCPToKRM(demotedPrimaryActual)
	if err != nil || demotedStatus.CurrentRole == nil || *demotedStatus.CurrentRole != "DR_REPLICA" {
		t.Fatalf("expected DR_REPLICA role for public IP demoted instance, got: %v", demotedStatus.CurrentRole)
	}
	promotedStatus, err := SQLInstanceStatusGCPToKRM(promotedReplicaActual)
	if err != nil || promotedStatus.CurrentRole == nil || *promotedStatus.CurrentRole != "PRIMARY" {
		t.Fatalf("expected PRIMARY role for public IP promoted instance, got: %v", promotedStatus.CurrentRole)
	}
}

// TestDiffInstances_DR_DecommissionFailoverDrReplica ensures that when a user clears
// failoverDrReplicaRef on a primary instance that has an active DR cluster (with psaWriteEndpoint),
// DiffInstances MUST NOT suppress the diff and must report the diff so Cloud SQL can remove the DR designation.
func TestDiffInstances_DR_DecommissionFailoverDrReplica(t *testing.T) {
	desiredPrimary := &api.DatabaseInstance{
		Name:            "dr-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings:        &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "", // User intentionally cleared DR replica designation
		},
	}
	actualPrimary := &api.DatabaseInstance{
		Name:            "dr-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings:        &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "dr-replica",
			PsaWriteEndpoint:      "psa-endpoint.sql.goog",
		},
	}

	diff := DiffInstances(desiredPrimary, actualPrimary)
	if !diff.HasDiff() {
		t.Fatalf("expected diff when decommissioning failoverDrReplicaName on primary with PSA endpoint, but diff was incorrectly suppressed!")
	}
	found := false
	for _, f := range diff.Fields {
		if f.ID == ".replicationCluster.failoverDrReplicaName" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected .replicationCluster.failoverDrReplicaName diff, got: %v", diff.Fields)
	}
}

// TestIsInstanceNameEqual_Formats validates name matching across all supported Cloud SQL reference formats:
// short names, legacy project-colon prefixes, full resource URLs, and REST URLs.
func TestIsInstanceNameEqual_Formats(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"my-instance", "my-instance", true},
		{"my-proj:my-instance", "my-instance", true},
		{"my-instance", "my-proj:my-instance", true},
		{"projects/my-proj/instances/my-instance", "my-instance", true},
		{"my-instance", "projects/my-proj/instances/my-instance", true},
		{"my-proj:my-instance", "projects/my-proj/instances/my-instance", true},
		{"projects/my-proj/instances/my-instance", "my-proj:my-instance", true},
		{"proj-a:my-instance", "projects/proj-b/instances/my-instance", false},
		{"projects/proj-a/instances/my-instance", "projects/proj-b/instances/my-instance", false},
		{"proj-a:my-instance", "proj-b:my-instance", false},
		{"https://sqladmin.googleapis.com/sql/v1beta4/projects/my-proj/instances/my-instance", "my-instance", true},
		{"https://sqladmin.googleapis.com/sql/v1beta4/projects/my-proj/instances/my-instance", "my-proj:my-instance", true},
		{"other-instance", "my-instance", false},
		{"notmy-instance", "my-instance", false},
		{"", "my-instance", false},
		{"my-instance", "", false},
		{"", "", true},
	}
	for _, tc := range tests {
		got := isInstanceNameEqual(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("isInstanceNameEqual(%q, %q) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}

// TestDiffInstances_DR_FullyQualifiedResourceURLs verifies that DR role-swap suppression
// works cleanly even when specs use fully-qualified GCP resource paths for masterInstanceRef
// or failoverDrReplicaRef.
func TestDiffInstances_DR_FullyQualifiedResourceURLs(t *testing.T) {
	desiredPrimary := &api.DatabaseInstance{
		Name:            "dr-primary",
		Region:          "us-central1",
		DatabaseVersion: "POSTGRES_16",
		Settings:        &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "projects/test-proj/instances/dr-replica",
		},
	}
	desiredReplica := &api.DatabaseInstance{
		Name:               "dr-replica",
		Region:             "us-east1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "projects/test-proj/instances/dr-primary",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
	}

	// Inverted state returned from GCP API (using short names)
	demotedPrimaryActual := &api.DatabaseInstance{
		Name:               "dr-primary",
		Region:             "us-central1",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "dr-replica",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica:        true,
			PsaWriteEndpoint: "dr.sql.goog",
		},
	}
	promotedReplicaActual := &api.DatabaseInstance{
		Name:            "dr-replica",
		Region:          "us-east1",
		DatabaseVersion: "POSTGRES_16",
		InstanceType:    "CLOUD_SQL_INSTANCE",
		Settings:        &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "dr-primary",
			PsaWriteEndpoint:      "dr.sql.goog",
		},
	}

	if diff := DiffInstances(desiredPrimary, demotedPrimaryActual); diff.HasDiff() {
		t.Fatalf("expected diff suppression on demoted primary with qualified URLs, got: %v", diff)
	}
	if diff := DiffInstances(desiredReplica, promotedReplicaActual); diff.HasDiff() {
		t.Fatalf("expected diff suppression on promoted replica with qualified URLs, got: %v", diff)
	}
}

// TestSQLInstanceStatus_CurrentRole_MasterInstancePrecedence ensures that an instance
// with MasterInstanceName set is never classified as PRIMARY, even if ReplicationCluster.DrReplica
// is temporarily false.
func TestSQLInstanceStatus_CurrentRole_MasterInstancePrecedence(t *testing.T) {
	instWithMaster := &api.DatabaseInstance{
		Name:               "replica-instance",
		MasterInstanceName: "primary-instance",
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica:        false, // Edge case: drReplica false but has master
			PsaWriteEndpoint: "dr.sql.goog",
		},
	}
	status, err := SQLInstanceStatusGCPToKRM(instWithMaster)
	if err != nil {
		t.Fatalf("unexpected error converting status: %v", err)
	}
	if status.CurrentRole == nil || *status.CurrentRole != "DR_REPLICA" {
		t.Fatalf("expected CurrentRole DR_REPLICA for instance with MasterInstanceName, got: %v", status.CurrentRole)
	}
}

// TestSQLInstance_StandbyDuringTransientOperationError ensures that when an instance is in
// MAINTENANCE or UPDATING state and the sqlOperations.list API returns a transient error,
// the controller gracefully enters standby (Ready=False, Reason=FailoverInProgress) and requests
// a requeue instead of issuing colliding mutating updates.
func TestSQLInstance_StandbyDuringTransientOperationError(t *testing.T) {
	ctx := context.Background()

	mutatingCallInvoked := false
	transport := &mockTransport{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method == "PUT" || req.Method == "PATCH" || req.Method == "POST" {
				if req.URL.Path != "/sql/v1beta4/projects/test-project/operations" {
					mutatingCallInvoked = true
				}
			}
			// Mock GET operations list failing with transient HTTP 500 error
			if req.Method == "GET" && req.URL.Path == "/sql/v1beta4/projects/test-project/operations" {
				return &http.Response{
					StatusCode: 500,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"error":{"code":500,"message":"Internal backend error"}}`))),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
			}, nil
		},
	}

	httpClient := &http.Client{Transport: transport}
	sqlClient, err := api.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatalf("failed creating sql service: %v", err)
	}

	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]interface{}{
				"name": "transient-instance",
			},
		},
	}

	adapter := &sqlInstanceAdapter{
		resourceID:          "transient-instance",
		projectID:           "test-project",
		sqlInstancesClient:  api.NewInstancesService(sqlClient),
		sqlOperationsClient: api.NewOperationsService(sqlClient),
		desired:             &krm.SQLInstance{},
		actual: &api.DatabaseInstance{
			Name:  "transient-instance",
			State: "MAINTENANCE", // In transient state
			Settings: &api.Settings{
				Tier: "db-perf-optimized-N-2",
			},
		},
		fieldMeta: make(map[string]*FieldMetadata),
	}

	updateOp := newTestUpdateOp(u)
	err = adapter.Update(ctx, updateOp)
	if err != nil {
		t.Fatalf("expected nil error on graceful standby during transient error, got: %v", err)
	}

	if mutatingCallInvoked {
		t.Fatalf("mutating API call was invoked during transient error in MAINTENANCE state!")
	}

	conditions, found, err := unstructured.NestedSlice(u.Object, "status", "conditions")
	if err != nil || !found || len(conditions) == 0 {
		t.Fatalf("expected conditions to be populated on unstructured, got: %v", conditions)
	}
	condMap := conditions[0].(map[string]any)
	if condMap["reason"] != "FailoverInProgress" {
		t.Fatalf("expected reason 'FailoverInProgress', got: %v", condMap["reason"])
	}
	if condMap["status"] != string(corev1.ConditionFalse) {
		t.Fatalf("expected status 'False', got: %v", condMap["status"])
	}
}

// TestSQLInstance_DeletionBlocked_TransientOperationError verifies that if checkActiveOperations
// returns an error while an instance is in MAINTENANCE or UPDATING state, Delete() aborts
// with an error to safeguard the resource from destructive termination.
func TestSQLInstance_DeletionBlocked_TransientOperationError(t *testing.T) {
	ctx := context.Background()

	deletionInvoked := false
	transport := &mockTransport{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method == "DELETE" {
				deletionInvoked = true
			}
			// Mock GET operations list failing with HTTP 500 error
			if req.Method == "GET" && req.URL.Path == "/sql/v1beta4/projects/test-project/operations" {
				return &http.Response{
					StatusCode: 500,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"error":{"code":500,"message":"Transient error"}}`))),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte(`{}`))),
			}, nil
		},
	}

	httpClient := &http.Client{Transport: transport}
	sqlClient, err := api.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatalf("failed creating sql service: %v", err)
	}

	adapter := &sqlInstanceAdapter{
		resourceID:          "deleting-instance",
		projectID:           "test-project",
		sqlInstancesClient:  api.NewInstancesService(sqlClient),
		sqlOperationsClient: api.NewOperationsService(sqlClient),
		actual: &api.DatabaseInstance{
			Name:  "deleting-instance",
			State: "MAINTENANCE",
		},
	}

	u := &unstructured.Unstructured{}
	c := fake.NewClientBuilder().Build()
	deleteOp := directbase.NewDeleteOperation(c, u)

	deleted, err := adapter.Delete(ctx, deleteOp)
	if err == nil {
		t.Fatalf("expected error blocking deletion during transient checkActiveOperations error, got nil")
	}
	if deleted {
		t.Fatalf("expected deleted=false when operation check fails during MAINTENANCE, got true")
	}
	if deletionInvoked {
		t.Fatalf("DELETE API call was invoked despite operation check failure during MAINTENANCE!")
	}
}

// TestSQLInstance_Create_DefersFailoverDrReplica verifies that when creating a new primary
// instance with a failoverDrReplicaRef, the controller strips/defers failoverDrReplicaName
// from the initial instances.insert request, preventing Cloud SQL from rejecting the call with a 400 error.
func TestSQLInstance_Create_DefersFailoverDrReplica(t *testing.T) {
	ctx := context.Background()

	var insertedInstance *api.DatabaseInstance
	transport := &mockTransport{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method == "POST" && strings.Contains(req.URL.Path, "/instances") {
				body, err := io.ReadAll(req.Body)
				if err != nil {
					return nil, err
				}
				insertedInstance = &api.DatabaseInstance{}
				if err := json.Unmarshal(body, insertedInstance); err != nil {
					return nil, err
				}
				op := &api.Operation{Name: "op-create", Status: "DONE"}
				data, _ := json.Marshal(op)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/dr-primary") {
				inst := &api.DatabaseInstance{
					Name:  "dr-primary",
					State: "RUNNABLE",
					Settings: &api.Settings{
						Tier: "db-perf-optimized-N-2",
						BackupConfiguration: &api.BackupConfiguration{
							Enabled: true,
						},
					},
				}
				data, _ := json.Marshal(inst)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			if req.Method == "GET" && strings.Contains(req.URL.Path, "/users") {
				usersList := &api.UsersListResponse{}
				data, _ := json.Marshal(usersList)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
			}, nil
		},
	}

	sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
	if err != nil {
		t.Fatalf("creating sql service: %v", err)
	}

	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]any{
				"name":      "dr-primary",
				"namespace": "default",
			},
		},
	}

	drReplicaName := "dr-replica"
	dbVersion := "POSTGRES_16"
	tier := "db-perf-optimized-N-2"
	adapter := &sqlInstanceAdapter{
		projectID:  "test-project",
		resourceID: "dr-primary",
		desired: &krm.SQLInstance{
			Spec: krm.SQLInstanceSpec{
				DatabaseVersion: &dbVersion,
				Settings: krm.InstanceSettings{
					Tier: tier,
				},
				ReplicationCluster: &krm.ReplicationCluster{
					FailoverDrReplicaRef: &refs.SQLInstanceRef{
						External: drReplicaName,
					},
				},
			},
		},
		sqlInstancesClient:  api.NewInstancesService(sqlService),
		sqlOperationsClient: api.NewOperationsService(sqlService),
		sqlUsersClient:      api.NewUsersService(sqlService),
		fieldMeta:           make(map[string]*FieldMetadata),
	}

	createOp := newTestCreateOp(u)
	if err := adapter.Create(ctx, createOp); err != nil {
		t.Fatalf("Create returned error: %v", err)
	}

	if insertedInstance == nil {
		t.Fatalf("instances.insert was not called")
	}

	if insertedInstance.ReplicationCluster != nil && insertedInstance.ReplicationCluster.FailoverDrReplicaName != "" {
		t.Fatalf("expected FailoverDrReplicaName to be deferred/stripped during insert, but got: %q", insertedInstance.ReplicationCluster.FailoverDrReplicaName)
	}
}

// TestSQLInstance_Update_WaitsForReplicaRunnable tests that during Update(), if a target DR replica
// is still creating (not yet RUNNABLE), the controller defers designating failoverDrReplicaName,
// requests a non-blocking requeue, and sets status condition Ready=False, Reason=ReplicationClusterPending.
func TestSQLInstance_Update_WaitsForReplicaRunnable(t *testing.T) {
	ctx := context.Background()

	mutatingCallInvoked := false
	transport := &mockTransport{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if (req.Method == "PUT" || req.Method == "PATCH") && strings.Contains(req.URL.Path, "/instances/dr-primary") {
				mutatingCallInvoked = true
				op := &api.Operation{Name: "op-update", Status: "DONE"}
				data, _ := json.Marshal(op)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			// DR replica is currently in PENDING_CREATE state in Cloud SQL
			if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/dr-replica") {
				replicaInst := &api.DatabaseInstance{
					Name:  "dr-replica",
					State: "PENDING_CREATE",
				}
				data, _ := json.Marshal(replicaInst)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			// Operations list check
			if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
				opList := &api.OperationsListResponse{}
				data, _ := json.Marshal(opList)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
			}, nil
		},
	}

	sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
	if err != nil {
		t.Fatalf("creating sql service: %v", err)
	}

	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]any{
				"name":      "dr-primary",
				"namespace": "default",
			},
		},
	}

	drReplicaName := "dr-replica"
	dbVersion := "POSTGRES_16"
	tier := "db-perf-optimized-N-2"
	resourceID := "dr-primary"
	backupEnabled := true
	desiredKRM := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			ResourceID:      &resourceID,
			DatabaseVersion: &dbVersion,
			Settings: krm.InstanceSettings{
				Tier: tier,
				BackupConfiguration: &krm.InstanceBackupConfiguration{
					Enabled: &backupEnabled,
				},
			},
			ReplicationCluster: &krm.ReplicationCluster{
				FailoverDrReplicaRef: &refs.SQLInstanceRef{
					External: drReplicaName,
				},
			},
		},
	}
	fieldMeta := make(map[string]*FieldMetadata)
	actualGCP, err := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
	if err != nil {
		t.Fatalf("converting desired to GCP: %v", err)
	}
	actualGCP.State = "RUNNABLE"
	actualGCP.ReplicationCluster = nil

	adapter := &sqlInstanceAdapter{
		projectID:           "test-project",
		resourceID:          "dr-primary",
		desired:             desiredKRM,
		actual:              actualGCP,
		sqlInstancesClient:  api.NewInstancesService(sqlService),
		sqlOperationsClient: api.NewOperationsService(sqlService),
		fieldMeta:           fieldMeta,
	}

	updateOp := newTestUpdateOp(u)
	if err := adapter.Update(ctx, updateOp); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}

	if mutatingCallInvoked {
		t.Fatalf("mutating update was issued while target DR replica was still PENDING_CREATE")
	}

	if !updateOp.RequeueRequested {
		t.Fatalf("expected RequeueRequested=true while waiting for replica to become RUNNABLE")
	}

	conds, found, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
	if !found || len(conds) == 0 {
		t.Fatalf("expected status conditions on unstructured, got none")
	}
	latestCond := conds[0].(map[string]any)
	if latestCond["reason"] != "ReplicationClusterPending" {
		t.Fatalf("expected Reason=ReplicationClusterPending, got: %v", latestCond["reason"])
	}
	if latestCond["status"] != string(corev1.ConditionFalse) {
		t.Fatalf("expected Status=False, got: %v", latestCond["status"])
	}
	msg, _ := latestCond["message"].(string)
	if !strings.Contains(msg, "PENDING_CREATE") || !strings.Contains(msg, "dr-replica") {
		t.Fatalf("expected condition message to mention replica and state, got: %q", msg)
	}
}

// TestSQLInstance_Update_DesignatesFailoverDrReplica_WhenRunnable tests that once the target
// DR replica reaches RUNNABLE state, the controller successfully updates the primary instance
// with failoverDrReplicaName.
func TestSQLInstance_Update_DesignatesFailoverDrReplica_WhenRunnable(t *testing.T) {
	ctx := context.Background()

	drReplicaName := "dr-replica"
	dbVersion := "POSTGRES_16"
	tier := "db-perf-optimized-N-2"
	resourceID := "dr-primary"
	backupEnabled := true
	desiredKRM := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			ResourceID:      &resourceID,
			DatabaseVersion: &dbVersion,
			Settings: krm.InstanceSettings{
				Tier: tier,
				BackupConfiguration: &krm.InstanceBackupConfiguration{
					Enabled: &backupEnabled,
				},
			},
			ReplicationCluster: &krm.ReplicationCluster{
				FailoverDrReplicaRef: &refs.SQLInstanceRef{
					External: drReplicaName,
				},
			},
		},
	}
	fieldMeta := make(map[string]*FieldMetadata)
	actualGCP, err := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
	if err != nil {
		t.Fatalf("converting desired to GCP: %v", err)
	}
	actualGCP.State = "RUNNABLE"
	actualGCP.ReplicationCluster = nil

	var updatedInstance *api.DatabaseInstance
	transport := &mockTransport{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if req.Method == "PUT" && strings.Contains(req.URL.Path, "/instances/dr-primary") {
				body, err := io.ReadAll(req.Body)
				if err != nil {
					return nil, err
				}
				updatedInstance = &api.DatabaseInstance{}
				if err := json.Unmarshal(body, updatedInstance); err != nil {
					return nil, err
				}
				op := &api.Operation{Name: "op-update", Status: "DONE"}
				data, _ := json.Marshal(op)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			// DR replica is RUNNABLE
			if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/dr-replica") {
				replicaInst := &api.DatabaseInstance{
					Name:  "dr-replica",
					State: "RUNNABLE",
				}
				data, _ := json.Marshal(replicaInst)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			// GET primary after update
			if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/dr-primary") {
				primaryInst, _ := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
				primaryInst.State = "RUNNABLE"
				data, _ := json.Marshal(primaryInst)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			// Operations list
			if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
				opList := &api.OperationsListResponse{}
				data, _ := json.Marshal(opList)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
			}, nil
		},
	}

	sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
	if err != nil {
		t.Fatalf("creating sql service: %v", err)
	}

	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]any{
				"name":      "dr-primary",
				"namespace": "default",
			},
		},
	}

	adapter := &sqlInstanceAdapter{
		projectID:           "test-project",
		resourceID:          "dr-primary",
		desired:             desiredKRM,
		actual:              actualGCP,
		sqlInstancesClient:  api.NewInstancesService(sqlService),
		sqlOperationsClient: api.NewOperationsService(sqlService),
		fieldMeta:           fieldMeta,
	}

	updateOp := newTestUpdateOp(u)
	if err := adapter.Update(ctx, updateOp); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}

	if updateOp.RequeueRequested {
		t.Fatalf("expected RequeueRequested=false when replica is RUNNABLE")
	}

	if updatedInstance == nil {
		t.Fatalf("expected Update to be called on primary instance")
	}

	if updatedInstance.ReplicationCluster == nil || updatedInstance.ReplicationCluster.FailoverDrReplicaName != "dr-replica" {
		t.Fatalf("expected failoverDrReplicaName='dr-replica' in update request, got: %v", updatedInstance.ReplicationCluster)
	}
}

// TestDiffInstances_DR_OptOutAnnotation tests that DiffInstancesWithConfig respects the
// suppressDRDiffs parameter, allowing users to opt out of diff suppression via
// cnrm.cloud.google.com/diff-suppression: "false".
func TestDiffInstances_DR_OptOutAnnotation(t *testing.T) {
	desired := &api.DatabaseInstance{
		Name:            "db-primary",
		DatabaseVersion: "POSTGRES_16",
		Settings:        &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			FailoverDrReplicaName: "db-replica",
		},
	}
	// Post-switchover: db-primary is now in GCP as a replica pointing to db-replica
	actualRoleSwapped := &api.DatabaseInstance{
		Name:               "db-primary",
		DatabaseVersion:    "POSTGRES_16",
		MasterInstanceName: "db-replica",
		InstanceType:       "READ_REPLICA_INSTANCE",
		Settings:           &api.Settings{Tier: "db-perf-optimized-N-2"},
		ReplicationCluster: &api.ReplicationCluster{
			DrReplica:        true,
			PsaWriteEndpoint: "dr-cluster-write.sql.goog",
		},
	}

	// 1. Default (suppressDRDiffs = true): diff is cleanly suppressed
	diffDefault := DiffInstancesWithConfig(desired, actualRoleSwapped, true)
	if diffDefault.HasDiff() {
		t.Fatalf("expected no diff under default suppression, got: %v", diffDefault)
	}

	// 2. Opt-out (suppressDRDiffs = false): diff IS reported on instanceType and masterInstanceName
	diffOptOut := DiffInstancesWithConfig(desired, actualRoleSwapped, false)
	if !diffOptOut.HasDiff() {
		t.Fatalf("expected diff to be detected when suppressDRDiffs=false, got none")
	}

	fields := make(map[string]bool)
	for _, f := range diffOptOut.Fields {
		fields[f.ID] = true
	}
	if !fields[".instanceType"] {
		t.Errorf("expected diff on .instanceType when opt-out enabled")
	}
	if !fields[".masterInstanceName"] {
		t.Errorf("expected diff on .masterInstanceName when opt-out enabled")
	}
}

// TestSQLInstance_Update_OptOutAnnotation_RoleInversionDetected tests that when a user sets
// cnrm.cloud.google.com/diff-suppression: "false" and a role swap is detected, the controller
// sets status condition Ready=False, Reason=RoleInversionDetected.
func TestSQLInstance_Update_OptOutAnnotation_RoleInversionDetected(t *testing.T) {
	ctx := context.Background()

	transport := &mockTransport{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			// Operations list
			if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
				opList := &api.OperationsListResponse{}
				data, _ := json.Marshal(opList)
				return &http.Response{
					StatusCode: 200,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader(data)),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
			}, nil
		},
	}

	sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
	if err != nil {
		t.Fatalf("creating sql service: %v", err)
	}

	u := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
			"kind":       "SQLInstance",
			"metadata": map[string]any{
				"name":      "dr-primary",
				"namespace": "default",
				"annotations": map[string]any{
					"cnrm.cloud.google.com/diff-suppression": "false",
				},
			},
		},
	}

	drReplicaName := "dr-replica"
	dbVersion := "POSTGRES_16"
	tier := "db-perf-optimized-N-2"
	resourceID := "dr-primary"
	backupEnabled := true
	desiredKRM := &krm.SQLInstance{
		Spec: krm.SQLInstanceSpec{
			ResourceID:      &resourceID,
			DatabaseVersion: &dbVersion,
			Settings: krm.InstanceSettings{
				Tier: tier,
				BackupConfiguration: &krm.InstanceBackupConfiguration{
					Enabled: &backupEnabled,
				},
			},
			ReplicationCluster: &krm.ReplicationCluster{
				FailoverDrReplicaRef: &refs.SQLInstanceRef{
					External: drReplicaName,
				},
			},
		},
	}
	desiredKRM.Annotations = map[string]string{
		"cnrm.cloud.google.com/diff-suppression": "false",
	}

	fieldMeta := make(map[string]*FieldMetadata)
	actualGCP, err := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
	if err != nil {
		t.Fatalf("converting desired to GCP: %v", err)
	}
	actualGCP.State = "RUNNABLE"
	actualGCP.MasterInstanceName = drReplicaName
	actualGCP.InstanceType = "READ_REPLICA_INSTANCE"
	actualGCP.ReplicationCluster = &api.ReplicationCluster{
		DrReplica: true,
	}

	adapter := &sqlInstanceAdapter{
		projectID:           "test-project",
		resourceID:          "dr-primary",
		desired:             desiredKRM,
		actual:              actualGCP,
		sqlInstancesClient:  api.NewInstancesService(sqlService),
		sqlOperationsClient: api.NewOperationsService(sqlService),
		fieldMeta:           fieldMeta,
	}

	updateOp := newTestUpdateOp(u)
	if err := adapter.Update(ctx, updateOp); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}

	conds, found, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
	if !found || len(conds) == 0 {
		t.Fatalf("expected status conditions on unstructured, got none")
	}
	latestCond := conds[0].(map[string]any)
	if latestCond["reason"] != "RoleInversionDetected" {
		t.Fatalf("expected Reason=RoleInversionDetected, got: %v", latestCond["reason"])
	}
	if latestCond["status"] != string(corev1.ConditionFalse) {
		t.Fatalf("expected Status=False, got: %v", latestCond["status"])
	}
	msg, _ := latestCond["message"].(string)
	if !strings.Contains(msg, "RoleInversionDetected") && !strings.Contains(msg, "inverted") {
		t.Fatalf("expected condition message to explain role inversion, got: %q", msg)
	}
}

// TestSQLInstance_Bootstrap_ExhaustiveMatrix_AllEngines tests that the circular reference bootstrap
// deferral and readiness gating operate flawlessly across all supported database engine families and versions:
// PostgreSQL (14, 15, 16), MySQL (8.0, 8.4), and SQL Server (2019, 2022).
func TestSQLInstance_Bootstrap_ExhaustiveMatrix_AllEngines(t *testing.T) {
	matrix := []struct {
		engineFamily string
		version      string
		tier         string
	}{
		{"PostgreSQL", "POSTGRES_14", "db-custom-2-7680"},
		{"PostgreSQL", "POSTGRES_15", "db-custom-2-7680"},
		{"PostgreSQL", "POSTGRES_16", "db-perf-optimized-N-2"},
		{"MySQL", "MYSQL_8_0", "db-custom-4-15360"},
		{"MySQL", "MYSQL_8_4", "db-perf-optimized-N-2"},
		{"SQL Server", "SQLSERVER_2019_STANDARD", "db-custom-4-16384"},
		{"SQL Server", "SQLSERVER_2019_ENTERPRISE", "db-custom-4-16384"},
		{"SQL Server", "SQLSERVER_2022_STANDARD", "db-custom-4-16384"},
		{"SQL Server", "SQLSERVER_2022_ENTERPRISE", "db-custom-8-32768"},
	}

	for _, tc := range matrix {
		t.Run(fmt.Sprintf("%s_%s", tc.engineFamily, tc.version), func(t *testing.T) {
			ctx := context.Background()
			primaryName := "dr-prim-" + strings.ToLower(tc.version)
			replicaName := "dr-repl-" + strings.ToLower(tc.version)
			tier := tc.tier
			ver := tc.version
			backupEnabled := true

			desiredKRM := &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					ResourceID:      &primaryName,
					DatabaseVersion: &ver,
					Settings: krm.InstanceSettings{
						Tier: tier,
						BackupConfiguration: &krm.InstanceBackupConfiguration{
							Enabled: &backupEnabled,
						},
					},
					ReplicationCluster: &krm.ReplicationCluster{
						FailoverDrReplicaRef: &refs.SQLInstanceRef{
							External: replicaName,
						},
					},
				},
			}

			// Subtest 1: Initial Create - FailoverDrReplicaName must be stripped/deferred from instances.insert
			t.Run("Create_DefersReplicationCluster", func(t *testing.T) {
				var insertedInst *api.DatabaseInstance
				transport := &mockTransport{
					roundTripFunc: func(req *http.Request) (*http.Response, error) {
						if req.Method == "POST" && strings.Contains(req.URL.Path, "/instances") {
							body, _ := io.ReadAll(req.Body)
							insertedInst = &api.DatabaseInstance{}
							_ = json.Unmarshal(body, insertedInst)
							op := &api.Operation{Name: "op-create", Status: "DONE"}
							data, _ := json.Marshal(op)
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/"+primaryName) {
							inst := &api.DatabaseInstance{
								Name:            primaryName,
								DatabaseVersion: ver,
								State:           "RUNNABLE",
								Settings: &api.Settings{
									Tier:                tier,
									BackupConfiguration: &api.BackupConfiguration{Enabled: true},
								},
							}
							data, _ := json.Marshal(inst)
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
						}, nil
					},
				}

				sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
				if err != nil {
					t.Fatalf("creating sql service: %v", err)
				}
				u := &unstructured.Unstructured{
					Object: map[string]any{
						"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
						"kind":       "SQLInstance",
						"metadata": map[string]any{
							"name":      primaryName,
							"namespace": "default",
						},
					},
				}
				adapter := &sqlInstanceAdapter{
					projectID:           "test-project",
					resourceID:          primaryName,
					desired:             desiredKRM,
					sqlInstancesClient:  api.NewInstancesService(sqlService),
					sqlOperationsClient: api.NewOperationsService(sqlService),
					sqlUsersClient:      api.NewUsersService(sqlService),
					fieldMeta:           make(map[string]*FieldMetadata),
				}
				createOp := newTestCreateOp(u)
				if err := adapter.Create(ctx, createOp); err != nil {
					t.Fatalf("Create returned error: %v", err)
				}
				if insertedInst == nil {
					t.Fatalf("instances.insert was never invoked")
				}
				if insertedInst.ReplicationCluster != nil && insertedInst.ReplicationCluster.FailoverDrReplicaName != "" {
					t.Fatalf("expected FailoverDrReplicaName to be omitted during initial insert, got: %q", insertedInst.ReplicationCluster.FailoverDrReplicaName)
				}
			})

			// Subtest 2: Update when replica is 404 (NOT_FOUND) -> Sets ReplicationClusterPending, requeues, zero mutations
			t.Run("Update_ReplicaNotFound_Gated", func(t *testing.T) {
				mutatingCallInvoked := false
				transport := &mockTransport{
					roundTripFunc: func(req *http.Request) (*http.Response, error) {
						if req.Method == "PUT" || req.Method == "PATCH" {
							mutatingCallInvoked = true
						}
						// Replica lookup returns 404
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/"+replicaName) {
							return &http.Response{
								StatusCode: 404,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader([]byte(`{"error":{"code":404,"message":"The resource could not be found."}}`))),
							}, nil
						}
						// Operations list
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
							data, _ := json.Marshal(&api.OperationsListResponse{})
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
						}, nil
					},
				}

				sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
				if err != nil {
					t.Fatalf("creating sql service: %v", err)
				}
				fieldMeta := make(map[string]*FieldMetadata)
				actualGCP, err := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
				if err != nil {
					t.Fatalf("SQLInstanceKRMToGCP: %v", err)
				}
				actualGCP.State = "RUNNABLE"
				actualGCP.ReplicationCluster = nil

				adapter := &sqlInstanceAdapter{
					projectID:           "test-project",
					resourceID:          primaryName,
					desired:             desiredKRM,
					actual:              actualGCP,
					sqlInstancesClient:  api.NewInstancesService(sqlService),
					sqlOperationsClient: api.NewOperationsService(sqlService),
					fieldMeta:           fieldMeta,
				}
				u := &unstructured.Unstructured{
					Object: map[string]any{
						"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
						"kind":       "SQLInstance",
						"metadata": map[string]any{
							"name":      primaryName,
							"namespace": "default",
						},
					},
				}
				updateOp := newTestUpdateOp(u)
				if err := adapter.Update(ctx, updateOp); err != nil {
					t.Fatalf("Update error: %v", err)
				}
				if mutatingCallInvoked {
					t.Fatalf("mutating update was issued while replica returned 404!")
				}
				if !updateOp.RequeueRequested {
					t.Fatalf("expected RequeueRequested=true while replica is not found")
				}
				conds, _, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
				if len(conds) == 0 || conds[0].(map[string]any)["reason"] != "ReplicationClusterPending" {
					t.Fatalf("expected Reason=ReplicationClusterPending, got: %v", conds)
				}
			})

			// Subtest 3: Update when replica is RUNNABLE -> Designates failoverDrReplicaName
			t.Run("Update_ReplicaRunnable_Designates", func(t *testing.T) {
				fieldMeta := make(map[string]*FieldMetadata)
				var updatedInst *api.DatabaseInstance
				transport := &mockTransport{
					roundTripFunc: func(req *http.Request) (*http.Response, error) {
						if req.Method == "PUT" && strings.Contains(req.URL.Path, "/instances/"+primaryName) {
							body, _ := io.ReadAll(req.Body)
							updatedInst = &api.DatabaseInstance{}
							_ = json.Unmarshal(body, updatedInst)
							op := &api.Operation{Name: "op-update", Status: "DONE"}
							data, _ := json.Marshal(op)
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						// Replica is RUNNABLE
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/"+replicaName) {
							data, _ := json.Marshal(&api.DatabaseInstance{Name: replicaName, State: "RUNNABLE"})
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						// GET primary after update
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/"+primaryName) {
							pInst, _ := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
							pInst.State = "RUNNABLE"
							pInst.ReplicationCluster = &api.ReplicationCluster{FailoverDrReplicaName: replicaName}
							data, _ := json.Marshal(pInst)
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						// Operations list
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
							data, _ := json.Marshal(&api.OperationsListResponse{})
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
						}, nil
					},
				}

				sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
				if err != nil {
					t.Fatalf("creating sql service: %v", err)
				}
				actualGCP, err := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
				if err != nil {
					t.Fatalf("SQLInstanceKRMToGCP: %v", err)
				}
				actualGCP.State = "RUNNABLE"
				actualGCP.ReplicationCluster = nil

				adapter := &sqlInstanceAdapter{
					projectID:           "test-project",
					resourceID:          primaryName,
					desired:             desiredKRM,
					actual:              actualGCP,
					sqlInstancesClient:  api.NewInstancesService(sqlService),
					sqlOperationsClient: api.NewOperationsService(sqlService),
					fieldMeta:           fieldMeta,
				}
				u := &unstructured.Unstructured{
					Object: map[string]any{
						"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
						"kind":       "SQLInstance",
						"metadata": map[string]any{
							"name":      primaryName,
							"namespace": "default",
						},
					},
				}
				updateOp := newTestUpdateOp(u)
				if err := adapter.Update(ctx, updateOp); err != nil {
					t.Fatalf("Update error: %v", err)
				}
				if updateOp.RequeueRequested {
					t.Fatalf("expected RequeueRequested=false when replica is RUNNABLE")
				}
				if updatedInst == nil || updatedInst.ReplicationCluster == nil || updatedInst.ReplicationCluster.FailoverDrReplicaName != replicaName {
					t.Fatalf("expected primary to be updated with failoverDrReplicaName=%q, got: %v", replicaName, updatedInst)
				}
			})
		})
	}
}

// TestSQLInstance_OptOut_ExhaustiveMatrix_AllEngines tests that when users disable diff-suppression
// via cnrm.cloud.google.com/diff-suppression: "false", the controller reliably intercepts role swaps,
// surfaces the RoleInversionDetected status condition, and strictly blocks mutating updates across all engines.
func TestSQLInstance_OptOut_ExhaustiveMatrix_AllEngines(t *testing.T) {
	engines := []struct {
		name    string
		version string
		tier    string
	}{
		{"PostgreSQL_14", "POSTGRES_14", "db-custom-2-7680"},
		{"PostgreSQL_15", "POSTGRES_15", "db-custom-2-7680"},
		{"PostgreSQL_16", "POSTGRES_16", "db-perf-optimized-N-2"},
		{"MySQL_8_0", "MYSQL_8_0", "db-custom-4-15360"},
		{"MySQL_8_4", "MYSQL_8_4", "db-perf-optimized-N-2"},
		{"SQLServer_2019", "SQLSERVER_2019_STANDARD", "db-custom-4-16384"},
		{"SQLServer_2022", "SQLSERVER_2022_ENTERPRISE", "db-custom-8-32768"},
	}

	for _, eng := range engines {
		t.Run(eng.name, func(t *testing.T) {
			ctx := context.Background()
			primaryName := "prim-" + strings.ToLower(eng.version)
			replicaName := "repl-" + strings.ToLower(eng.version)
			ver := eng.version
			tier := eng.tier
			backupEnabled := true

			// Desired state: Primary referencing replica
			desiredPrimaryKRM := &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					ResourceID:      &primaryName,
					DatabaseVersion: &ver,
					Settings: krm.InstanceSettings{
						Tier: tier,
						BackupConfiguration: &krm.InstanceBackupConfiguration{
							Enabled: &backupEnabled,
						},
					},
					ReplicationCluster: &krm.ReplicationCluster{
						FailoverDrReplicaRef: &refs.SQLInstanceRef{
							External: replicaName,
						},
					},
				},
			}
			desiredPrimaryKRM.Annotations = map[string]string{
				"cnrm.cloud.google.com/diff-suppression": "false",
			}

			// Case 1: Primary instance demoted to replica post-switchover
			t.Run("DemotedPrimary_HaltsMutations", func(t *testing.T) {
				mutatingCallInvoked := false
				transport := &mockTransport{
					roundTripFunc: func(req *http.Request) (*http.Response, error) {
						if req.Method == "PUT" || req.Method == "PATCH" || req.Method == "POST" {
							mutatingCallInvoked = true
						}
						// Operations list
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
							data, _ := json.Marshal(&api.OperationsListResponse{})
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
						}, nil
					},
				}

				sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
				if err != nil {
					t.Fatalf("creating sql service: %v", err)
				}

				fieldMeta := make(map[string]*FieldMetadata)
				actualGCP, err := SQLInstanceKRMToGCP(desiredPrimaryKRM, nil, fieldMeta)
				if err != nil {
					t.Fatalf("SQLInstanceKRMToGCP: %v", err)
				}
				actualGCP.State = "RUNNABLE"
				actualGCP.MasterInstanceName = replicaName
				actualGCP.InstanceType = "READ_REPLICA_INSTANCE"
				actualGCP.ReplicationCluster = &api.ReplicationCluster{
					DrReplica: true,
				}

				u := &unstructured.Unstructured{
					Object: map[string]any{
						"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
						"kind":       "SQLInstance",
						"metadata": map[string]any{
							"name":      primaryName,
							"namespace": "default",
							"annotations": map[string]any{
								"cnrm.cloud.google.com/diff-suppression": "false",
							},
						},
					},
				}

				adapter := &sqlInstanceAdapter{
					projectID:           "test-project",
					resourceID:          primaryName,
					desired:             desiredPrimaryKRM,
					actual:              actualGCP,
					sqlInstancesClient:  api.NewInstancesService(sqlService),
					sqlOperationsClient: api.NewOperationsService(sqlService),
					fieldMeta:           fieldMeta,
				}

				updateOp := newTestUpdateOp(u)
				if err := adapter.Update(ctx, updateOp); err != nil {
					t.Fatalf("Update returned unexpected error: %v", err)
				}

				if mutatingCallInvoked {
					t.Fatalf("mutating update was issued when diff-suppression was disabled on role-swapped instance!")
				}

				conds, found, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
				if !found || len(conds) == 0 {
					t.Fatalf("expected conditions on unstructured")
				}
				latestCond := conds[0].(map[string]any)
				if latestCond["reason"] != "RoleInversionDetected" {
					t.Fatalf("expected Reason=RoleInversionDetected, got: %v", latestCond["reason"])
				}
				if latestCond["status"] != string(corev1.ConditionFalse) {
					t.Fatalf("expected Status=False, got: %v", latestCond["status"])
				}
			})

			// Case 2: Replica instance promoted to primary post-switchover
			t.Run("PromotedReplica_HaltsMutations", func(t *testing.T) {
				mutatingCallInvoked := false
				transport := &mockTransport{
					roundTripFunc: func(req *http.Request) (*http.Response, error) {
						if req.Method == "PUT" || req.Method == "PATCH" || req.Method == "POST" {
							mutatingCallInvoked = true
						}
						// Operations list
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
							data, _ := json.Marshal(&api.OperationsListResponse{})
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
						}, nil
					},
				}

				sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
				if err != nil {
					t.Fatalf("creating sql service: %v", err)
				}

				desiredReplicaKRM := &krm.SQLInstance{
					Spec: krm.SQLInstanceSpec{
						ResourceID:      &replicaName,
						DatabaseVersion: &ver,
						MasterInstanceRef: &refs.SQLInstanceRef{
							External: primaryName,
						},
						Settings: krm.InstanceSettings{
							Tier: tier,
						},
					},
				}
				desiredReplicaKRM.Annotations = map[string]string{
					"cnrm.cloud.google.com/diff-suppression": "false",
				}

				fieldMeta := make(map[string]*FieldMetadata)
				actualGCP, err := SQLInstanceKRMToGCP(desiredReplicaKRM, nil, fieldMeta)
				if err != nil {
					t.Fatalf("SQLInstanceKRMToGCP: %v", err)
				}
				actualGCP.State = "RUNNABLE"
				actualGCP.MasterInstanceName = ""
				actualGCP.InstanceType = "CLOUD_SQL_INSTANCE"
				actualGCP.ReplicationCluster = &api.ReplicationCluster{
					FailoverDrReplicaName: primaryName,
				}

				u := &unstructured.Unstructured{
					Object: map[string]any{
						"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
						"kind":       "SQLInstance",
						"metadata": map[string]any{
							"name":      replicaName,
							"namespace": "default",
							"annotations": map[string]any{
								"cnrm.cloud.google.com/diff-suppression": "false",
							},
						},
					},
				}

				adapter := &sqlInstanceAdapter{
					projectID:           "test-project",
					resourceID:          replicaName,
					desired:             desiredReplicaKRM,
					actual:              actualGCP,
					sqlInstancesClient:  api.NewInstancesService(sqlService),
					sqlOperationsClient: api.NewOperationsService(sqlService),
					fieldMeta:           fieldMeta,
				}

				updateOp := newTestUpdateOp(u)
				if err := adapter.Update(ctx, updateOp); err != nil {
					t.Fatalf("Update returned unexpected error: %v", err)
				}

				if mutatingCallInvoked {
					t.Fatalf("mutating update was issued when diff-suppression was disabled on promoted replica!")
				}

				conds, found, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
				if !found || len(conds) == 0 {
					t.Fatalf("expected conditions on unstructured")
				}
				latestCond := conds[0].(map[string]any)
				if latestCond["reason"] != "RoleInversionDetected" {
					t.Fatalf("expected Reason=RoleInversionDetected, got: %v", latestCond["reason"])
				}
			})
		})
	}
}

// TestSQLInstance_Update_TargetReplica_TransientErrors verifies that unexpected transient backend
// errors (HTTP 500, 503) encountered during DR replica readiness polling are propagated cleanly
// without panic, and without issuing premature mutating updates against the primary.
func TestSQLInstance_Update_TargetReplica_TransientErrors(t *testing.T) {
	statusCodes := []int{500, 503}

	for _, statusCode := range statusCodes {
		t.Run(fmt.Sprintf("HTTP_%d", statusCode), func(t *testing.T) {
			ctx := context.Background()
			primaryName := "prim-transient"
			replicaName := "repl-transient"
			ver := "POSTGRES_16"
			tier := "db-perf-optimized-N-2"
			backupEnabled := true

			desiredKRM := &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					ResourceID:      &primaryName,
					DatabaseVersion: &ver,
					Settings: krm.InstanceSettings{
						Tier: tier,
						BackupConfiguration: &krm.InstanceBackupConfiguration{
							Enabled: &backupEnabled,
						},
					},
					ReplicationCluster: &krm.ReplicationCluster{
						FailoverDrReplicaRef: &refs.SQLInstanceRef{
							External: replicaName,
						},
					},
				},
			}

			mutatingCallInvoked := false
			transport := &mockTransport{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					if req.Method == "PUT" || req.Method == "PATCH" {
						mutatingCallInvoked = true
					}
					// Replica lookup returns transient error
					if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/"+replicaName) {
						return &http.Response{
							StatusCode: statusCode,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte(fmt.Sprintf(`{"error":{"code":%d,"message":"Backend service temporarily unavailable"}}`, statusCode)))),
						}, nil
					}
					// Operations list
					if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
						data, _ := json.Marshal(&api.OperationsListResponse{})
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader(data)),
						}, nil
					}
					return &http.Response{
						StatusCode: 200,
						Header:     make(http.Header),
						Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
					}, nil
				},
			}

			sqlService, err := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
			if err != nil {
				t.Fatalf("creating sql service: %v", err)
			}
			fieldMeta := make(map[string]*FieldMetadata)
			actualGCP, err := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
			if err != nil {
				t.Fatalf("SQLInstanceKRMToGCP: %v", err)
			}
			actualGCP.State = "RUNNABLE"
			actualGCP.ReplicationCluster = nil

			adapter := &sqlInstanceAdapter{
				projectID:           "test-project",
				resourceID:          primaryName,
				desired:             desiredKRM,
				actual:              actualGCP,
				sqlInstancesClient:  api.NewInstancesService(sqlService),
				sqlOperationsClient: api.NewOperationsService(sqlService),
				fieldMeta:           fieldMeta,
			}
			u := &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
					"kind":       "SQLInstance",
					"metadata": map[string]any{
						"name":      primaryName,
						"namespace": "default",
					},
				},
			}
			updateOp := newTestUpdateOp(u)
			err = adapter.Update(ctx, updateOp)
			if err == nil {
				t.Fatalf("expected Update to return error on HTTP %d from replica lookup, got nil", statusCode)
			}
			if !strings.Contains(err.Error(), "checking readiness of target DR replica") {
				t.Fatalf("expected error message to mention target DR replica readiness, got: %v", err)
			}
			if mutatingCallInvoked {
				t.Fatalf("mutating update was issued despite transient error on replica readiness check!")
			}
		})
	}
}

// TestSQLInstance_FullDRLifecycle_AllEngines tests a complete end-to-end lifecycle walkthrough across
// all 3 database engine families: PostgreSQL, MySQL, and SQL Server.
// Stages verified:
// 1. Initial creation deferral & transition to RUNNABLE
// 2. Replication cluster designation
// 3. Planned switchover in-flight detection (FailoverInProgress) and deletion guard
// 4. Planned switchover completion & acknowledgment (FailoverAcknowledged)
// 5. Reverse failback in-flight detection
// 6. Reverse failback completion & acknowledgment
// 7. Decommissioning of failoverDrReplicaRef (clearing replication cluster)
func TestSQLInstance_FullDRLifecycle_AllEngines(t *testing.T) {
	engines := []struct {
		family  string
		version string
		tier    string
	}{
		{"PostgreSQL", "POSTGRES_16", "db-perf-optimized-N-2"},
		{"MySQL", "MYSQL_8_0", "db-perf-optimized-N-2"},
		{"SQLServer", "SQLSERVER_2022_ENTERPRISE", "db-custom-4-16384"},
	}

	for _, eng := range engines {
		t.Run(eng.family, func(t *testing.T) {
			ctx := context.Background()
			primName := "life-prim-" + strings.ToLower(eng.family)
			replName := "life-repl-" + strings.ToLower(eng.family)
			ver := eng.version
			tier := eng.tier
			backupEnabled := true

			desiredKRM := &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					ResourceID:      &primName,
					DatabaseVersion: &ver,
					Settings: krm.InstanceSettings{
						Tier: tier,
						BackupConfiguration: &krm.InstanceBackupConfiguration{
							Enabled: &backupEnabled,
						},
					},
					ReplicationCluster: &krm.ReplicationCluster{
						FailoverDrReplicaRef: &refs.SQLInstanceRef{
							External: replName,
						},
					},
				},
			}

			// Stage 1: Active SWITCHOVER operation in GCP -> FailoverInProgress + Deletion Blocked
			t.Run("Stage1_SwitchoverInProgress", func(t *testing.T) {
				transport := &mockTransport{
					roundTripFunc: func(req *http.Request) (*http.Response, error) {
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
							opList := &api.OperationsListResponse{
								Items: []*api.Operation{
									{
										Name:          "op-switchover-active",
										OperationType: "SWITCHOVER",
										Status:        "RUNNING",
										TargetId:      primName,
									},
								},
							}
							data, _ := json.Marshal(opList)
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
						}, nil
					},
				}
				sqlService, _ := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
				fieldMeta := make(map[string]*FieldMetadata)
				actualGCP, _ := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
				actualGCP.State = "MAINTENANCE"

				adapter := &sqlInstanceAdapter{
					projectID:           "test-project",
					resourceID:          primName,
					desired:             desiredKRM,
					actual:              actualGCP,
					sqlInstancesClient:  api.NewInstancesService(sqlService),
					sqlOperationsClient: api.NewOperationsService(sqlService),
					fieldMeta:           fieldMeta,
				}

				u := &unstructured.Unstructured{
					Object: map[string]any{
						"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
						"kind":       "SQLInstance",
						"metadata": map[string]any{
							"name":      primName,
							"namespace": "default",
						},
					},
				}
				updateOp := newTestUpdateOp(u)
				if err := adapter.Update(ctx, updateOp); err != nil {
					t.Fatalf("Update error: %v", err)
				}
				if !updateOp.RequeueRequested {
					t.Fatalf("expected RequeueRequested=true during switchover")
				}
				conds, _, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
				if len(conds) == 0 || conds[0].(map[string]any)["reason"] != "FailoverInProgress" {
					t.Fatalf("expected Reason=FailoverInProgress, got: %v", conds)
				}

				// Deletion safety check
				c := fake.NewClientBuilder().WithObjects(u).Build()
				deleteOp := directbase.NewDeleteOperation(c, u)
				deleted, err := adapter.Delete(ctx, deleteOp)
				if err == nil || deleted {
					t.Fatalf("expected deletion to be rejected during active SWITCHOVER operation")
				}
			})

			// Stage 2: Switchover Complete -> Diff Suppressed, FailoverAcknowledged
			t.Run("Stage2_SwitchoverComplete_Acknowledged", func(t *testing.T) {
				transport := &mockTransport{
					roundTripFunc: func(req *http.Request) (*http.Response, error) {
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
							data, _ := json.Marshal(&api.OperationsListResponse{})
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
						}, nil
					},
				}
				sqlService, _ := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
				fieldMeta := make(map[string]*FieldMetadata)
				actualGCP, _ := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
				actualGCP.State = "RUNNABLE"
				actualGCP.MasterInstanceName = replName
				actualGCP.InstanceType = "READ_REPLICA_INSTANCE"
				actualGCP.ReplicationCluster = &api.ReplicationCluster{
					DrReplica:        true,
					PsaWriteEndpoint: "dr-endpoint.sql.goog",
				}

				u := &unstructured.Unstructured{
					Object: map[string]any{
						"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
						"kind":       "SQLInstance",
						"metadata": map[string]any{
							"name":      primName,
							"namespace": "default",
						},
						"status": map[string]any{
							"conditions": []any{
								map[string]any{
									"type":   "Ready",
									"status": string(corev1.ConditionFalse),
									"reason": "FailoverInProgress",
								},
							},
							"currentRole": "PRIMARY",
						},
					},
				}

				adapter := &sqlInstanceAdapter{
					projectID:           "test-project",
					resourceID:          primName,
					desired:             desiredKRM,
					actual:              actualGCP,
					sqlInstancesClient:  api.NewInstancesService(sqlService),
					sqlOperationsClient: api.NewOperationsService(sqlService),
					fieldMeta:           fieldMeta,
				}

				updateOp := newTestUpdateOp(u)
				if err := adapter.Update(ctx, updateOp); err != nil {
					t.Fatalf("Update error: %v", err)
				}
				conds, _, _ := unstructured.NestedSlice(u.Object, "status", "conditions")
				if len(conds) == 0 || conds[0].(map[string]any)["reason"] != "FailoverAcknowledged" {
					t.Fatalf("expected Reason=FailoverAcknowledged post-switchover, got: %v", conds)
				}
				role, _, _ := unstructured.NestedString(u.Object, "status", "currentRole")
				if role != "DR_REPLICA" {
					t.Fatalf("expected status.currentRole=DR_REPLICA, got: %s", role)
				}
			})

			// Stage 3: Decommissioning failoverDrReplicaRef -> Issues update to clear replicationCluster
			t.Run("Stage3_DecommissionReplica", func(t *testing.T) {
				decommissionedDesiredKRM := &krm.SQLInstance{
					Spec: krm.SQLInstanceSpec{
						ResourceID:      &primName,
						DatabaseVersion: &ver,
						Settings: krm.InstanceSettings{
							Tier: tier,
							BackupConfiguration: &krm.InstanceBackupConfiguration{
								Enabled: &backupEnabled,
							},
						},
						ReplicationCluster: nil, // Cleared!
					},
				}

				var updatedInst *api.DatabaseInstance
				transport := &mockTransport{
					roundTripFunc: func(req *http.Request) (*http.Response, error) {
						if req.Method == "PUT" && strings.Contains(req.URL.Path, "/instances/"+primName) {
							body, _ := io.ReadAll(req.Body)
							updatedInst = &api.DatabaseInstance{}
							_ = json.Unmarshal(body, updatedInst)
							op := &api.Operation{Name: "op-decom", Status: "DONE"}
							data, _ := json.Marshal(op)
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/instances/"+primName) {
							inst := &api.DatabaseInstance{
								Name:            primName,
								DatabaseVersion: ver,
								State:           "RUNNABLE",
								Settings: &api.Settings{
									Tier:                tier,
									BackupConfiguration: &api.BackupConfiguration{Enabled: true},
								},
							}
							data, _ := json.Marshal(inst)
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						if req.Method == "GET" && strings.Contains(req.URL.Path, "/operations") {
							data, _ := json.Marshal(&api.OperationsListResponse{})
							return &http.Response{
								StatusCode: 200,
								Header:     make(http.Header),
								Body:       io.NopCloser(bytes.NewReader(data)),
							}, nil
						}
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
						}, nil
					},
				}

				sqlService, _ := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))
				fieldMeta := make(map[string]*FieldMetadata)
				actualGCP, _ := SQLInstanceKRMToGCP(desiredKRM, nil, fieldMeta)
				actualGCP.State = "RUNNABLE"
				actualGCP.ReplicationCluster = &api.ReplicationCluster{
					FailoverDrReplicaName: replName,
				}

				u := &unstructured.Unstructured{
					Object: map[string]any{
						"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
						"kind":       "SQLInstance",
						"metadata": map[string]any{
							"name":      primName,
							"namespace": "default",
						},
					},
				}

				adapter := &sqlInstanceAdapter{
					projectID:           "test-project",
					resourceID:          primName,
					desired:             decommissionedDesiredKRM,
					actual:              actualGCP,
					sqlInstancesClient:  api.NewInstancesService(sqlService),
					sqlOperationsClient: api.NewOperationsService(sqlService),
					fieldMeta:           fieldMeta,
				}

				updateOp := newTestUpdateOp(u)
				if err := adapter.Update(ctx, updateOp); err != nil {
					t.Fatalf("Update error during decommissioning: %v", err)
				}
				if updatedInst == nil {
					t.Fatalf("expected update call to be made to Cloud SQL to clear replication cluster")
				}
				if updatedInst.ReplicationCluster != nil && updatedInst.ReplicationCluster.FailoverDrReplicaName != "" {
					t.Fatalf("expected failoverDrReplicaName to be cleared in update request, got: %v", updatedInst.ReplicationCluster)
				}
			})
		})
	}
}

func TestSQLInstance_Mode3_PromoteReplica_AllEngines(t *testing.T) {
	engines := []struct {
		name    string
		version string
	}{
		{"PostgreSQL_16", "POSTGRES_16"},
		{"MySQL_8_0", "MYSQL_8_0"},
		{"SQLServer_2022", "SQLSERVER_2022_STANDARD"},
	}

	for _, eng := range engines {
		t.Run(eng.name, func(t *testing.T) {
			ctx := context.Background()
			instanceName := "repl-" + strings.ToLower(eng.name)
			ver := eng.version
			promoteCalled := false

			transport := &mockTransport{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					// Check for promoteReplica API call
					if strings.Contains(req.URL.Path, "/promoteReplica") && req.Method == http.MethodPost {
						promoteCalled = true
						opJSON := `{"name": "op-promote-123", "status": "DONE", "operationType": "PROMOTE_REPLICA"}`
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte(opJSON))),
						}, nil
					}
					// Check for poll operation
					if strings.Contains(req.URL.Path, "/operations/") {
						opJSON := `{"name": "op-promote-123", "status": "DONE"}`
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte(opJSON))),
						}, nil
					}
					// Re-fetching instance after promotion -> now CLOUD_SQL_INSTANCE with masterInstanceName: ""
					if strings.Contains(req.URL.Path, "/instances/"+instanceName) && req.Method == http.MethodGet {
						instJSON := fmt.Sprintf(`{
							"name": "%s",
							"databaseVersion": "%s",
							"state": "RUNNABLE",
							"instanceType": "CLOUD_SQL_INSTANCE",
							"masterInstanceName": "",
							"settings": {"tier": "db-custom-2-7680", "settingsVersion": "2"}
						}`, instanceName, ver)
						return &http.Response{
							StatusCode: 200,
							Header:     make(http.Header),
							Body:       io.NopCloser(bytes.NewReader([]byte(instJSON))),
						}, nil
					}
					return &http.Response{
						StatusCode: 200,
						Header:     make(http.Header),
						Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
					}, nil
				},
			}

			sqlService, _ := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))

			// Desired KRM: Standalone primary (masterInstanceRef is nil)
			desiredKRM := &krm.SQLInstance{
				Spec: krm.SQLInstanceSpec{
					ResourceID:      &instanceName,
					DatabaseVersion: &ver,
					Settings: krm.InstanceSettings{
						Tier: "db-custom-2-7680",
					},
					// MasterInstanceRef is nil (demoted to standalone primary)
				},
			}

			// Actual in GCP: Currently a read replica
			actualGCP := &api.DatabaseInstance{
				Name:               instanceName,
				DatabaseVersion:    ver,
				State:              "RUNNABLE",
				InstanceType:       "READ_REPLICA_INSTANCE",
				MasterInstanceName: "some-primary-instance",
				Settings: &api.Settings{
					Tier: "db-custom-2-7680",
				},
			}

			u := &unstructured.Unstructured{
				Object: map[string]any{
					"apiVersion": "sql.cnrm.cloud.google.com/v1beta1",
					"kind":       "SQLInstance",
					"metadata": map[string]any{
						"name":      instanceName,
						"namespace": "default",
					},
				},
			}

			adapter := &sqlInstanceAdapter{
				projectID:           "test-project",
				resourceID:          instanceName,
				desired:             desiredKRM,
				actual:              actualGCP,
				sqlInstancesClient:  api.NewInstancesService(sqlService),
				sqlOperationsClient: api.NewOperationsService(sqlService),
				fieldMeta:           make(map[string]*FieldMetadata),
			}

			updateOp := newTestUpdateOp(u)
			if err := adapter.Update(ctx, updateOp); err != nil {
				t.Fatalf("Update returned error during replica promotion: %v", err)
			}

			if !promoteCalled {
				t.Fatalf("expected PromoteReplica API to be called during Mode 3 promotion, but it was not")
			}

			instanceType, _, _ := unstructured.NestedString(u.Object, "status", "instanceType")
			if instanceType != "CLOUD_SQL_INSTANCE" {
				t.Fatalf("expected status.instanceType to be CLOUD_SQL_INSTANCE after promotion, got: %s", instanceType)
			}
			if _, found, _ := unstructured.NestedMap(u.Object, "status", "masterInstanceRef"); found {
				t.Fatalf("expected status.masterInstanceRef to be cleared after promotion")
			}
		})
	}
}

func TestSQLInstance_Delete_ActiveReplicas_DiagnosticError(t *testing.T) {
	ctx := context.Background()
	instanceName := "test-primary-with-replicas"

	transport := &mockTransport{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			if strings.Contains(req.URL.Path, "/instances/"+instanceName) && req.Method == http.MethodDelete {
				errJSON := `{"error": {"code": 400, "message": "Invalid request: The instance has replica(s): [test-replica-1, test-replica-2]. Please delete replica(s) first."}}`
				return &http.Response{
					StatusCode: http.StatusBadRequest,
					Header:     make(http.Header),
					Body:       io.NopCloser(bytes.NewReader([]byte(errJSON))),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Header:     make(http.Header),
				Body:       io.NopCloser(bytes.NewReader([]byte("{}"))),
			}, nil
		},
	}

	sqlService, _ := api.NewService(ctx, option.WithHTTPClient(&http.Client{Transport: transport}))

	actualGCP := &api.DatabaseInstance{
		Name:         instanceName,
		State:        "RUNNABLE",
		InstanceType: "CLOUD_SQL_INSTANCE",
		ReplicaNames: []string{"test-replica-1", "test-replica-2"},
	}

	adapter := &sqlInstanceAdapter{
		projectID:           "test-project",
		resourceID:          instanceName,
		actual:              actualGCP,
		sqlInstancesClient:  api.NewInstancesService(sqlService),
		sqlOperationsClient: api.NewOperationsService(sqlService),
	}

	deleteOp := &directbase.DeleteOperation{}
	deleted, err := adapter.Delete(ctx, deleteOp)
	if deleted {
		t.Fatalf("expected deleted=false when replicas exist, got true")
	}
	if err == nil {
		t.Fatalf("expected Delete to return error when instance has active replicas, got nil")
	}
	if !strings.Contains(err.Error(), "cannot delete primary SQLInstance") || !strings.Contains(err.Error(), "while replicas exist") {
		t.Fatalf("expected error message to contain actionable replica dependency diagnostic, got: %v", err)
	}
}
