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

package ha

import (
	"bytes"
	"context"
	"slices"
	"strings"
	"testing"
	"time"

	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestStatusValidateNormalizesFields(t *testing.T) {
	opts := newStatusOptions()
	opts.contexts = []string{" cluster-a ", " cluster-b"}
	opts.bucket = "gs://my-ha-bucket/"

	err := opts.validate()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantContexts := []string{"cluster-a", "cluster-b"}
	if !slices.Equal(opts.contexts, wantContexts) {
		t.Fatalf("got contexts %v, want %v", opts.contexts, wantContexts)
	}
	if opts.bucket != "my-ha-bucket" {
		t.Fatalf("got bucket %q, want %q", opts.bucket, "my-ha-bucket")
	}
}

func TestStatusValidateErrors(t *testing.T) {
	tests := []struct {
		name     string
		contexts []string
		bucket   string
		wantErr  string
	}{
		{
			name:     "fewer than 2 contexts",
			contexts: []string{"cluster-a"},
			wantErr:  "at least 2 cluster contexts must be specified via --contexts",
		},
		{
			name:     "duplicate contexts",
			contexts: []string{"cluster-a", "cluster-a"},
			wantErr:  "duplicate context \"cluster-a\" specified in --contexts",
		},
		{
			name:     "invalid bucket name syntax",
			contexts: []string{"cluster-a", "cluster-b"},
			bucket:   "INVALID_BUCKET!",
			wantErr:  "invalid GCS bucket name \"INVALID_BUCKET!\"; bucket names must be 3-63 characters of lowercase letters, numbers, '-', '_', or '.', starting and ending with a letter or number",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			opts := newStatusOptions()
			opts.contexts = slices.Clone(tc.contexts)
			opts.bucket = tc.bucket

			err := opts.validate()

			if err == nil || err.Error() != tc.wantErr {
				t.Fatalf("got error %v, want %q", err, tc.wantErr)
			}
		})
	}
}

func TestExecuteStatusHealthyActivePassive(t *testing.T) {
	now := time.Date(2026, 10, 28, 3, 0, 10, 0, time.UTC)
	renewStr := "2026-10-28T03:00:05Z"
	duration := int32(15)
	transitions := int32(2)
	leader := "cluster-a"
	handleA := newFakeClusterHandle(
		t,
		"gke_proj_us-central1_cluster-a",
		newConfigConnectorWithMCL("kcc-multicluster-lease", "cnrm-system", "cluster-a"),
		newStatusMultiClusterLease("cnrm-system", "kcc-multicluster-lease", "cluster-a", leader, renewStr, duration, transitions),
		newStatusKRMSyncer("cnrm-system", "krmsyncer-cluster", true, "dummy-secret", []metav1.Condition{
			{Type: "Suspended", Status: metav1.ConditionTrue, Reason: "Suspended"},
		}),
		newStatusMCLDeployment("my-ha-bucket"),
	)
	handleB := newFakeClusterHandle(
		t,
		"gke_proj_us-east1_cluster-b",
		newConfigConnectorWithMCL("kcc-multicluster-lease", "cnrm-system", "cluster-b"),
		newStatusMultiClusterLease("cnrm-system", "kcc-multicluster-lease", "cluster-b", leader, renewStr, duration, transitions),
		newStatusKRMSyncer("cnrm-system", "krmsyncer-cluster", false, "cluster-a", []metav1.Condition{
			{Type: "Active", Status: metav1.ConditionTrue, Reason: "Syncing"},
		}),
	)
	outRunner := &fakeOutputRunner{
		responses: map[string]string{
			"gcloud storage ls gs://my-ha-bucket/leases/**":                                  "gs://my-ha-bucket/leases/cnrm-system/kcc-multicluster-lease\n",
			"gcloud storage cat gs://my-ha-bucket/leases/cnrm-system/kcc-multicluster-lease": `{"holderIdentity":"cluster-a","renewTime":"2026-10-28T03:00:05Z","leaseTransitions":2}`,
		},
	}
	var stdout bytes.Buffer
	opts := newStatusOptions()
	opts.stdout = &stdout

	err := executeStatus(context.Background(), opts, []*clusterHandle{handleA, handleB}, outRunner, now)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantOutput := strings.Join([]string{
		"=== GCS Lock Status ===",
		"Bucket:      gs://my-ha-bucket",
		"Lock Object: gs://my-ha-bucket/leases/cnrm-system/kcc-multicluster-lease",
		"Leader:      cluster-a",
		"Renew Age:   5s (renewed at 2026-10-28T03:00:05Z, transitions: 2)",
		"",
		"=== Cluster HA Status ===",
		"Context:            gke_proj_us-central1_cluster-a",
		"  Candidate ID:     cluster-a",
		"  Lease:            cnrm-system/kcc-multicluster-lease (replication: Status)",
		"  Role:             LEADER",
		"  Observed Leader:  cluster-a",
		"  Renew Age:        5s (renewed at 2026-10-28T03:00:05Z, transitions: 2)",
		"  Syncer:           cnrm-system/krmsyncer-cluster (mode: pull, suspended: true, source: dummy-secret, health: Suspended=True, errors: none)",
		"Context:            gke_proj_us-east1_cluster-b",
		"  Candidate ID:     cluster-b",
		"  Lease:            cnrm-system/kcc-multicluster-lease (replication: Status)",
		"  Role:             STANDBY",
		"  Observed Leader:  cluster-a",
		"  Renew Age:        5s (renewed at 2026-10-28T03:00:05Z, transitions: 2)",
		"  Syncer:           cnrm-system/krmsyncer-cluster (mode: pull, suspended: false, source: cluster-a, health: Active=True, errors: none)",
		"",
		"=== Warnings ===",
		"None",
		"",
	}, "\n")
	if got := stdout.String(); got != wantOutput {
		t.Fatalf("got output:\n%s\nwant:\n%s", got, wantOutput)
	}
}

func TestExecuteStatusWarnings(t *testing.T) {
	now := time.Date(2026, 10, 28, 3, 5, 0, 0, time.UTC)
	staleRenewStr := "2026-10-28T03:00:00Z"
	duration := int32(15)
	transitions := int32(1)
	leaderA := "cluster-a"
	leaderB := "cluster-b"
	handleA := newFakeClusterHandle(
		t,
		"gke_proj_us-central1_cluster-a",
		newConfigConnectorWithMCL("lease-a", "cnrm-system", leaderA),
		newStatusMultiClusterLease("cnrm-system", "lease-a", leaderA, leaderA, staleRenewStr, duration, transitions),
		newStatusKRMSyncer("cnrm-system", "krmsyncer-cluster", false, leaderB, []metav1.Condition{
			{Type: "InvalidRule", Status: metav1.ConditionTrue, Reason: "RuleValidationFailed", Message: "unknown GVK in rule 0"},
		}),
	)
	handleB := newFakeClusterHandle(
		t,
		"gke_proj_us-east1_cluster-b",
		newConfigConnectorWithMCL("lease-b", "cnrm-system", leaderB),
		newStatusMultiClusterLease("cnrm-system", "lease-b", leaderB, leaderB, staleRenewStr, duration, transitions),
	)
	outRunner := &fakeOutputRunner{
		responses: map[string]string{
			"gcloud storage ls gs://my-ha-bucket/leases/**": strings.Join([]string{
				"gs://my-ha-bucket/leases/cnrm-system/lease-a",
				"gs://my-ha-bucket/leases/cnrm-system/lease-b",
			}, "\n"),
			"gcloud storage cat gs://my-ha-bucket/leases/cnrm-system/lease-a": `{"holderIdentity":"cluster-a","renewTime":"2026-10-28T03:00:00Z","leaseTransitions":1}`,
		},
	}
	var stdout bytes.Buffer
	opts := newStatusOptions()
	opts.bucket = "my-ha-bucket"
	opts.stdout = &stdout

	err := executeStatus(context.Background(), opts, []*clusterHandle{handleA, handleB}, outRunner, now)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantOutput := strings.Join([]string{
		"=== GCS Lock Status ===",
		"Bucket:      gs://my-ha-bucket",
		"Lock Object: gs://my-ha-bucket/leases/cnrm-system/lease-a",
		"Leader:      cluster-a",
		"Renew Age:   5m0s (renewed at 2026-10-28T03:00:00Z, transitions: 1)",
		"Extra Keys:  leases/cnrm-system/lease-b",
		"",
		"=== Cluster HA Status ===",
		"Context:            gke_proj_us-central1_cluster-a",
		"  Candidate ID:     cluster-a",
		"  Lease:            cnrm-system/lease-a (replication: Status)",
		"  Role:             LEADER",
		"  Observed Leader:  cluster-a",
		"  Renew Age:        5m0s (renewed at 2026-10-28T03:00:00Z, transitions: 1)",
		"  Syncer:           cnrm-system/krmsyncer-cluster (mode: pull, suspended: false, source: cluster-b, health: InvalidRule=True, errors: InvalidRule=True (RuleValidationFailed): unknown GVK in rule 0)",
		"Context:            gke_proj_us-east1_cluster-b",
		"  Candidate ID:     cluster-b",
		"  Lease:            cnrm-system/lease-b (replication: Status)",
		"  Role:             LEADER",
		"  Observed Leader:  cluster-b",
		"  Renew Age:        5m0s (renewed at 2026-10-28T03:00:00Z, transitions: 1)",
		"  Syncer:           <none>",
		"",
		"=== Warnings ===",
		"- stale MultiClusterLease status on cluster \"gke_proj_us-central1_cluster-a\": leader \"cluster-a\" last renewed 5m0s ago (at 2026-10-28T03:00:00Z, staleness threshold 30s)",
		"- cluster \"gke_proj_us-central1_cluster-a\" KRMSyncer \"cnrm-system/krmsyncer-cluster\" reported error condition InvalidRule=True (RuleValidationFailed): unknown GVK in rule 0",
		"- stale MultiClusterLease status on cluster \"gke_proj_us-east1_cluster-b\": leader \"cluster-b\" last renewed 5m0s ago (at 2026-10-28T03:00:00Z, staleness threshold 30s)",
		"- mismatched leaseName across clusters (gke_proj_us-central1_cluster-a=\"lease-a\", gke_proj_us-east1_cluster-b=\"lease-b\"): set spec.experiments.multiClusterLease.leaseName to the same value on every cluster because the GCS lock object is leases/<namespace>/<leaseName>",
		"- extra lease keys found in gs://my-ha-bucket: [leases/cnrm-system/lease-b] (expected only \"leases/cnrm-system/lease-a\"; clusters using different lease keys will elect multiple active leaders)",
		"- stale GCS lock \"gs://my-ha-bucket/leases/cnrm-system/lease-a\": holder \"cluster-a\" last renewed 5m0s ago (at 2026-10-28T03:00:00Z, staleness threshold 30s)",
		"- leader mismatch on cluster \"gke_proj_us-east1_cluster-b\": MultiClusterLease status.globalHolderIdentity is \"cluster-b\", but GCS lock holder is \"cluster-a\"",
		"- clusters disagree on global leader (gke_proj_us-central1_cluster-a=\"cluster-a\", gke_proj_us-east1_cluster-b=\"cluster-b\")",
		"- multiple clusters report role LEADER simultaneously: [gke_proj_us-central1_cluster-a (cluster-a), gke_proj_us-east1_cluster-b (cluster-b)]",
		"",
	}, "\n")
	if got := stdout.String(); got != wantOutput {
		t.Fatalf("got output:\n%s\nwant:\n%s", got, wantOutput)
	}
}

func newStatusMultiClusterLease(namespace, name, holder, globalHolder, renewTime string, durationSec, transitions int32) *mclv1alpha1.MultiClusterLease {
	return &mclv1alpha1.MultiClusterLease{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: mclv1alpha1.MultiClusterLeaseSpec{
			HolderIdentity:       &holder,
			LeaseDurationSeconds: &durationSec,
		},
		Status: mclv1alpha1.MultiClusterLeaseStatus{
			GlobalHolderIdentity:       &globalHolder,
			GlobalRenewTime:            &renewTime,
			GlobalLeaseDurationSeconds: &durationSec,
			GlobalLeaseTransitions:     &transitions,
		},
	}
}

func newStatusKRMSyncer(namespace, name string, suspend bool, secretName string, conditions []metav1.Condition) *syncerv1alpha1.KRMSyncer {
	return &syncerv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: syncerv1alpha1.KRMSyncerSpec{
			Suspend: suspend,
			Mode:    syncerv1alpha1.ModePull,
			Remote: &syncerv1alpha1.RemoteConfig{
				ClusterConfig: &syncerv1alpha1.ClusterConfig{
					KubeConfigSecretRef: &corev1.SecretReference{
						Namespace: namespace,
						Name:      secretName,
					},
				},
			},
		},
		Status: syncerv1alpha1.KRMSyncerStatus{
			Conditions: conditions,
		},
	}
}
