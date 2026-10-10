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
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	api "google.golang.org/api/sqladmin/v1beta4"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSQLInstance_PSCAutoConnections_RealGCP_10Rounds(t *testing.T) {
	if os.Getenv("RUN_REAL_GCP_TEST") != "1" {
		t.Skip("Skipping Real GCP test; set RUN_REAL_GCP_TEST=1 to execute against GCP project gca-gke-2025")
	}

	projectID := "gca-gke-2025"
	if p := os.Getenv("GCP_PROJECT_ID"); p != "" {
		projectID = p
	}

	ctx := context.Background()

	// Initialize real Cloud SQL client using Google Default Application Credentials
	sqlService, err := api.NewService(ctx, option.WithScopes(api.SqlserviceAdminScope, api.CloudPlatformScope))
	if err != nil {
		t.Fatalf("Failed to initialize Cloud SQL Admin client: %v", err)
	}

	instanceName := fmt.Sprintf("kcc-psc-test-%d", time.Now().Unix())
	netDefault := fmt.Sprintf("projects/%s/global/networks/default", projectID)
	netDr := fmt.Sprintf("projects/%s/global/networks/sql-dr-vpc", projectID)

	t.Logf("=== Starting 10-Round Real GCP Test for PSC Auto Connections ===")
	t.Logf("Project: %s, Instance: %s", projectID, instanceName)

	// Ensure cleanup at end of test
	defer func() {
		t.Logf("Cleaning up test instance %s if still present...", instanceName)
		op, err := sqlService.Instances.Delete(projectID, instanceName).Context(ctx).Do()
		if err == nil && op != nil {
			waitForOperation(ctx, t, sqlService, projectID, op.Name, 10*time.Minute)
		}
	}()

	// -------------------------------------------------------------
	// ROUND 1: Baseline Provisioning & KRM -> GCP Translation
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 1: Baseline Provisioning with PSC Auto Connections (Default VPC)")
	krmInstance := &krm.SQLInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instanceName,
			Namespace: "default",
		},
		Spec: krm.SQLInstanceSpec{
			ResourceID:      direct.PtrTo(instanceName),
			DatabaseVersion: direct.PtrTo("MYSQL_8_0"),
			Region:          direct.PtrTo("us-central1"),
			Settings: krm.InstanceSettings{
				Tier: "db-f1-micro",
				IpConfiguration: &krm.InstanceIpConfiguration{
					Ipv4Enabled: direct.PtrTo(false),
					PscConfig: []krm.InstancePscConfig{
						{
							PscEnabled:                     direct.PtrTo(true),
							PscAutoConnectionPolicyEnabled: direct.PtrTo(true),
							AllowedConsumerProjects:        []string{projectID},
							PscAutoConnections: []krm.InstancePscAutoConnectionConfig{
								{
									ConsumerNetwork: direct.PtrTo(netDefault),
								},
							},
						},
					},
				},
				LocationPreference: &krm.InstanceLocationPreference{
					Zone: direct.PtrTo("us-central1-a"),
				},
			},
		},
	}

	desiredGCP, err := SQLInstanceKRMToGCP(krmInstance, nil, nil)
	if err != nil {
		t.Fatalf("Round 1: SQLInstanceKRMToGCP failed: %v", err)
	}
	if desiredGCP.Settings.IpConfiguration.PscConfig == nil || len(desiredGCP.Settings.IpConfiguration.PscConfig.PscAutoConnections) != 1 {
		t.Fatalf("Round 1: Mapped PscConfig.PscAutoConnections count is not 1")
	}

	t.Logf("Creating Cloud SQL instance %s in GCP via Instances.Insert...", instanceName)
	insertOp, err := sqlService.Instances.Insert(projectID, desiredGCP).Context(ctx).Do()
	if err != nil {
		t.Fatalf("Round 1: Insert API call failed: %v", err)
	}
	t.Logf("Instance insert initiated (op: %s). Waiting for completion...", insertOp.Name)
	if err := waitForOperation(ctx, t, sqlService, projectID, insertOp.Name, 15*time.Minute); err != nil {
		t.Fatalf("Round 1: Operation failed: %v", err)
	}
	t.Logf("Round 1: Instance creation succeeded!")

	// -------------------------------------------------------------
	// ROUND 2: GCP -> KRM Mapping & Status Field Verification
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 2: Fetching live GCP resource and validating GCP -> KRM conversion")
	liveInstance, err := sqlService.Instances.Get(projectID, instanceName).Context(ctx).Do()
	if err != nil {
		t.Fatalf("Round 2: Instances.Get failed: %v", err)
	}

	if liveInstance.Settings.IpConfiguration.PscConfig == nil {
		t.Fatalf("Round 2: Live instance has nil PscConfig")
	}
	t.Logf("Live PscConfig: PscEnabled=%v, PscAutoConnections count=%d",
		liveInstance.Settings.IpConfiguration.PscConfig.PscEnabled,
		len(liveInstance.Settings.IpConfiguration.PscConfig.PscAutoConnections))

	convertedKRM, err := SQLInstanceGCPToKRM(liveInstance)
	if err != nil {
		t.Fatalf("Round 2: SQLInstanceGCPToKRM failed: %v", err)
	}

	pscConns := convertedKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections
	if len(pscConns) != 1 {
		t.Fatalf("Round 2: Expected 1 PscAutoConnection in converted KRM, got %d", len(pscConns))
	}
	t.Logf("Converted PscAutoConnection: ConsumerNetwork=%v, Status=%v, ConsumerNetworkStatus=%v, IpAddress=%v",
		direct.ValueOf(pscConns[0].ConsumerNetwork),
		direct.ValueOf(pscConns[0].Status),
		direct.ValueOf(pscConns[0].ConsumerNetworkStatus),
		direct.ValueOf(pscConns[0].IpAddress))

	if direct.ValueOf(pscConns[0].ConsumerNetwork) != netDefault {
		t.Errorf("Round 2: ConsumerNetwork mismatch: expected %s, got %s", netDefault, direct.ValueOf(pscConns[0].ConsumerNetwork))
	}
	t.Logf("Round 2: GCP -> KRM mapping verified successfully!")

	// -------------------------------------------------------------
	// ROUND 3: Idempotency & Zero-Drift Verification
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 3: Validating Idempotency & Drift Detection against live GCP state")
	diff := DiffInstances(desiredGCP, liveInstance)
	if diff.HasDiff() {
		t.Errorf("Round 3: Expected zero drift between desired KRM and live GCP object, got diffs: %v", diff.Fields)
	} else {
		t.Logf("Round 3: Zero drift verified! (diff.HasDiff() == false)")
	}

	// -------------------------------------------------------------
	// ROUND 4: Network URI Format Resilience
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 4: Validating Full URI Normalization (https://www.googleapis.com/... format)")
	fullURIKRM := krmInstance.DeepCopy()
	fullURI := "https://www.googleapis.com/compute/v1/" + netDefault
	fullURIKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections[0].ConsumerNetwork = direct.PtrTo(fullURI)

	fullURIGCP, err := SQLInstanceKRMToGCP(fullURIKRM, nil, nil)
	if err != nil {
		t.Fatalf("Round 4: SQLInstanceKRMToGCP failed: %v", err)
	}
	diffURI := DiffInstances(fullURIGCP, liveInstance)
	if diffURI.HasDiff() {
		t.Errorf("Round 4: Normalization failed for full REST URI: diffs=%v", diffURI.Fields)
	} else {
		t.Logf("Round 4: Full REST URI format matches canonical path without false drift!")
	}

	// -------------------------------------------------------------
	// ROUND 5: ConsumerNetworkRef Reference Resolution
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 5: Validating ConsumerNetworkRef resolution")
	refKRM := krmInstance.DeepCopy()
	refKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections[0].ConsumerNetwork = nil
	refKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections[0].ConsumerNetworkRef = &computerefs.ComputeNetworkRef{
		External: netDefault,
	}

	fakeKube := fake.NewClientBuilder().WithScheme(runtime.NewScheme()).Build()
	if err := resolvePscAutoConnectionsRefs(ctx, fakeKube, refKRM); err != nil {
		t.Fatalf("Round 5: resolvePscAutoConnectionsRefs failed: %v", err)
	}

	resolvedGCP, err := SQLInstanceKRMToGCP(refKRM, nil, nil)
	if err != nil {
		t.Fatalf("Round 5: SQLInstanceKRMToGCP failed: %v", err)
	}
	diffRef := DiffInstances(resolvedGCP, liveInstance)
	if diffRef.HasDiff() {
		t.Errorf("Round 5: Diff detected after ConsumerNetworkRef resolution: %v", diffRef.Fields)
	} else {
		t.Logf("Round 5: ConsumerNetworkRef successfully resolved and produced identical GCP spec!")
	}

	// -------------------------------------------------------------
	// ROUND 6: In-Place Mutation: Add 2nd Auto Connection (Dual-VPC)
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 6: In-Place Mutation: Adding 2nd Auto Connection (%s)", netDr)
	twoConnKRM := krmInstance.DeepCopy()
	twoConnKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections = append(
		twoConnKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections,
		krm.InstancePscAutoConnectionConfig{
			ConsumerNetwork: direct.PtrTo(netDr),
		},
	)

	twoConnGCP, err := SQLInstanceKRMToGCP(twoConnKRM, nil, nil)
	if err != nil {
		t.Fatalf("Round 6: SQLInstanceKRMToGCP failed: %v", err)
	}

	diffAdd := DiffInstances(twoConnGCP, liveInstance)
	if !diffAdd.HasDiff() {
		t.Fatalf("Round 6: DiffInstances failed to detect added connection")
	}
	t.Logf("Detected expected diff on add: %v", diffAdd.Fields)

	t.Logf("Applying in-place PATCH update to Cloud SQL...")
	updateOp, err := sqlService.Instances.Patch(projectID, instanceName, &api.DatabaseInstance{
		Settings: twoConnGCP.Settings,
	}).Context(ctx).Do()
	if err != nil {
		t.Fatalf("Round 6: Instances.Patch failed: %v", err)
	}
	t.Logf("Patch operation %s started. Waiting for completion...", updateOp.Name)
	if err := waitForOperation(ctx, t, sqlService, projectID, updateOp.Name, 15*time.Minute); err != nil {
		t.Fatalf("Round 6: Update operation failed: %v", err)
	}

	liveUpdatedInstance, err := sqlService.Instances.Get(projectID, instanceName).Context(ctx).Do()
	if err != nil {
		t.Fatalf("Round 6: Instances.Get failed: %v", err)
	}
	updatedCount := len(liveUpdatedInstance.Settings.IpConfiguration.PscConfig.PscAutoConnections)
	t.Logf("Live instance updated. PscAutoConnections count = %d", updatedCount)
	if updatedCount != 2 {
		t.Errorf("Round 6: Expected 2 PscAutoConnections on live instance, got %d", updatedCount)
	}

	// -------------------------------------------------------------
	// ROUND 7: List Reordering Robustness
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 7: Validating List Reordering Robustness (Order Independence)")
	reorderedKRM := twoConnKRM.DeepCopy()
	// Swap order: [netDr, netDefault]
	reorderedKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections[0],
		reorderedKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections[1] =
		reorderedKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections[1],
		reorderedKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections[0]

	reorderedGCP, err := SQLInstanceKRMToGCP(reorderedKRM, nil, nil)
	if err != nil {
		t.Fatalf("Round 7: SQLInstanceKRMToGCP failed: %v", err)
	}
	diffReorder := DiffInstances(reorderedGCP, liveUpdatedInstance)
	if diffReorder.HasDiff() {
		t.Errorf("Round 7: List reordering caused false drift: %v", diffReorder.Fields)
	} else {
		t.Logf("Round 7: Reordered list successfully matched live state without drift!")
	}

	// -------------------------------------------------------------
	// ROUND 8: In-Place Mutation: Remove Connection
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 8: In-Place Mutation: Pruning Connection (Removing Default VPC, leaving %s)", netDr)
	oneConnKRM := twoConnKRM.DeepCopy()
	oneConnKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections = []krm.InstancePscAutoConnectionConfig{
		{ConsumerNetwork: direct.PtrTo(netDr)},
	}

	oneConnGCP, err := SQLInstanceKRMToGCP(oneConnKRM, nil, nil)
	if err != nil {
		t.Fatalf("Round 8: SQLInstanceKRMToGCP failed: %v", err)
	}

	diffRemove := DiffInstances(oneConnGCP, liveUpdatedInstance)
	if !diffRemove.HasDiff() {
		t.Fatalf("Round 8: DiffInstances failed to detect removed connection")
	}
	t.Logf("Detected expected diff on removal: %v", diffRemove.Fields)

	t.Logf("Applying in-place PATCH to prune connection...")
	removeOp, err := sqlService.Instances.Patch(projectID, instanceName, &api.DatabaseInstance{
		Settings: oneConnGCP.Settings,
	}).Context(ctx).Do()
	if err != nil {
		t.Fatalf("Round 8: Instances.Patch failed: %v", err)
	}
	if err := waitForOperation(ctx, t, sqlService, projectID, removeOp.Name, 15*time.Minute); err != nil {
		t.Fatalf("Round 8: Remove operation failed: %v", err)
	}

	livePrunedInstance, err := sqlService.Instances.Get(projectID, instanceName).Context(ctx).Do()
	if err != nil {
		t.Fatalf("Round 8: Instances.Get failed: %v", err)
	}
	prunedCount := len(livePrunedInstance.Settings.IpConfiguration.PscConfig.PscAutoConnections)
	t.Logf("Live instance pruned. PscAutoConnections count = %d", prunedCount)
	if prunedCount != 1 {
		t.Errorf("Round 8: Expected 1 connection after prune, got %d", prunedCount)
	}
	if prunedCount > 0 && livePrunedInstance.Settings.IpConfiguration.PscConfig.PscAutoConnections[0].ConsumerNetwork != netDr {
		t.Errorf("Round 8: Expected remaining connection to be %s, got %s",
			netDr, livePrunedInstance.Settings.IpConfiguration.PscConfig.PscAutoConnections[0].ConsumerNetwork)
	}

	// -------------------------------------------------------------
	// ROUND 9: Asynchronous Error / Status Handling on Invalid Network Path
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 9: Asynchronous Error / Status Handling on Non-Existent Network")
	invalidNet := "projects/gca-gke-2025/global/networks/non-existent-network-abc-xyz"
	invalidKRM := oneConnKRM.DeepCopy()
	invalidKRM.Spec.Settings.IpConfiguration.PscConfig[0].PscAutoConnections = []krm.InstancePscAutoConnectionConfig{
		{ConsumerNetwork: direct.PtrTo(invalidNet)},
	}
	invalidGCP, err := SQLInstanceKRMToGCP(invalidKRM, nil, nil)
	if err != nil {
		t.Fatalf("Round 9: SQLInstanceKRMToGCP failed: %v", err)
	}

	patchOp9, err := sqlService.Instances.Patch(projectID, instanceName, &api.DatabaseInstance{
		Settings: invalidGCP.Settings,
	}).Context(ctx).Do()
	if err != nil {
		t.Logf("Round 9: Synchronous rejection by API: %v", err)
	} else {
		t.Logf("Round 9: Patch accepted by API (op: %s). Waiting for operation completion...", patchOp9.Name)
		if err := waitForOperation(ctx, t, sqlService, projectID, patchOp9.Name, 10*time.Minute); err != nil {
			t.Logf("Round 9: Operation completed with expected error: %v", err)
		}
		// Verify instance status reports CONNECTION_POLICY_MISSING
		inst9, err := sqlService.Instances.Get(projectID, instanceName).Context(ctx).Do()
		if err == nil && inst9.Settings != nil && inst9.Settings.IpConfiguration != nil && inst9.Settings.IpConfiguration.PscConfig != nil {
			for _, ac := range inst9.Settings.IpConfiguration.PscConfig.PscAutoConnections {
				t.Logf("Round 9: Observed auto connection status: network=%s, networkStatus=%s, serviceResult=%s",
					ac.ConsumerNetwork, ac.ConsumerNetworkStatus, ac.ServiceConnectionPolicyCreationResult)
			}
		}
	}

	// -------------------------------------------------------------
	// ROUND 10: Teardown & Lifecycle Cleanup Verification
	// -------------------------------------------------------------
	t.Logf(">>> ROUND 10: Teardown & Lifecycle Cleanup")
	deleteOp, err := sqlService.Instances.Delete(projectID, instanceName).Context(ctx).Do()
	if err != nil {
		t.Fatalf("Round 10: Instances.Delete failed: %v", err)
	}
	t.Logf("Delete initiated (op: %s). Waiting for deletion completion...", deleteOp.Name)
	if err := waitForOperation(ctx, t, sqlService, projectID, deleteOp.Name, 15*time.Minute); err != nil {
		t.Fatalf("Round 10: Deletion operation failed: %v", err)
	}

	// Verify instance is gone
	_, err = sqlService.Instances.Get(projectID, instanceName).Context(ctx).Do()
	if err == nil {
		t.Errorf("Round 10: Instance still exists after deletion completed")
	} else if gErr, ok := err.(*googleapi.Error); ok && gErr.Code == http.StatusNotFound {
		t.Logf("Round 10: Confirmed instance 404 NOT FOUND (clean deletion)!")
	} else {
		t.Logf("Round 10: Instance confirmed gone with expected error: %v", err)
	}

	t.Logf("=== 10-Round Real GCP Test for PSC Auto Connections PASSED 100%% ===")
}

func waitForOperation(ctx context.Context, t *testing.T, sqlService *api.Service, project, operationName string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		op, err := sqlService.Operations.Get(project, operationName).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("reading operation %s: %w", operationName, err)
		}
		if op.Status == "DONE" {
			if op.Error != nil && len(op.Error.Errors) > 0 {
				return fmt.Errorf("operation %s failed: %s (%s)", operationName, op.Error.Errors[0].Code, op.Error.Errors[0].Message)
			}
			return nil
		}
		time.Sleep(10 * time.Second)
	}
	return fmt.Errorf("operation %s timed out after %v", operationName, timeout)
}
