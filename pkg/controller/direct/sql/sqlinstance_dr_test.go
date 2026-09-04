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
