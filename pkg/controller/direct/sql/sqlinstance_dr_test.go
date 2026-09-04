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
	"io"
	"net/http"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

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
