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
	"fmt"
	"strings"
	"testing"
	"time"

	syncerv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type fakeOutputRunner struct {
	responses map[string]string
	errors    map[string]error
}

func (f *fakeOutputRunner) output(_ context.Context, name string, args ...string) (string, error) {
	cmd := name + " " + strings.Join(args, " ")
	if f.errors != nil {
		if err, ok := f.errors[cmd]; ok {
			return "", err
		}
	}
	if f.responses != nil {
		if out, ok := f.responses[cmd]; ok {
			return out, nil
		}
	}
	return "", fmt.Errorf("unexpected command %q", cmd)
}

func TestExecuteStatusHealthySetupWithAutoDiscoveredBucket(t *testing.T) {
	now := time.Date(2026, 10, 28, 3, 0, 10, 0, time.UTC)
	renewStr := "2026-10-28T03:00:05Z"
	transitions := int32(2)
	duration := int32(15)
	leaderID := "cluster-a"
	standbyID := "cluster-b"
	handleA := newFakeClusterHandle(
		t,
		"gke_proj_us-central1_cluster-a",
		newConfigConnectorWithMCL("kcc-multicluster-lease", "cnrm-system", leaderID),
		newStatusMCLDeployment("my-ha-bucket"),
		newStatusMultiClusterLease("cnrm-system", "kcc-multicluster-lease", leaderID, leaderID, renewStr, duration, transitions),
		newStatusKRMSyncer("cnrm-system", "krmsyncer-cluster", true, standbyID, []metav1.Condition{
			{Type: "Suspended", Status: metav1.ConditionTrue, Reason: "LeaderSuspended", Message: "Syncer suspended on leader"},
		}),
		newPeerKubeconfigSecret("cnrm-system", standbyID),
	)
	handleB := newFakeClusterHandle(
		t,
		"gke_proj_us-east1_cluster-b",
		newConfigConnectorWithMCL("kcc-multicluster-lease", "cnrm-system", standbyID),
		newStatusMCLDeployment("my-ha-bucket"),
		newStatusMultiClusterLease("cnrm-system", "kcc-multicluster-lease", standbyID, leaderID, renewStr, duration, transitions),
		newStatusKRMSyncer("cnrm-system", "krmsyncer-cluster", false, leaderID, []metav1.Condition{
			{Type: "Active", Status: metav1.ConditionTrue, Reason: "Reconciling", Message: "Pulling from leader"},
		}),
		newPeerKubeconfigSecret("cnrm-system", leaderID),
	)
	outRunner := &fakeOutputRunner{
		responses: map[string]string{
			"gcloud storage ls gs://my-ha-bucket/leases/**": "gs://my-ha-bucket/leases/cnrm-system/kcc-multicluster-lease\n",
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
	got := stdout.String()
	for _, wantSnippet := range []string{
		"=== GCS Lock Status ===",
		"Bucket:      gs://my-ha-bucket",
		"Leader:      cluster-a",
		"Renew Age:   5s (renewed at 2026-10-28T03:00:05Z, transitions: 2)",
		"Context:            gke_proj_us-central1_cluster-a",
		"Role:             LEADER",
		"Syncer:           cnrm-system/krmsyncer-cluster (mode: pull, suspended: true, source: cluster-b, health: Suspended=True, errors: none)",
		"Context:            gke_proj_us-east1_cluster-b",
		"Role:             STANDBY",
		"Syncer:           cnrm-system/krmsyncer-cluster (mode: pull, suspended: false, source: cluster-a, health: Active=True, errors: none)",
		"=== Warnings ===\nNone",
	} {
		if !strings.Contains(got, wantSnippet) {
			t.Fatalf("expected status output to contain %q, got:\n%s", wantSnippet, got)
		}
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
	got := stdout.String()
	for _, wantWarning := range []string{
		"extra lease keys found in gs://my-ha-bucket: [leases/cnrm-system/lease-b]",
		"mismatched leaseName across clusters",
		"missing peer kubeconfig Secret \"cnrm-system/cluster-b\"",
		"missing peer kubeconfig Secret \"cnrm-system/cluster-a\"",
		"stale MultiClusterLease status on cluster \"gke_proj_us-central1_cluster-a\"",
		"stale GCS lock \"gs://my-ha-bucket/leases/cnrm-system/lease-a\"",
		"KRMSyncer \"cnrm-system/krmsyncer-cluster\" reported error condition InvalidRule=True (RuleValidationFailed): unknown GVK in rule 0",
		"clusters disagree on global leader",
		"multiple clusters report role LEADER simultaneously",
	} {
		if !strings.Contains(got, wantWarning) {
			t.Fatalf("expected status output to contain warning %q, got:\n%s", wantWarning, got)
		}
	}
}

func newStatusMCLDeployment(bucket string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mclDeploymentName,
			Namespace: cnrmSystemNamespace,
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "manager",
							Command: []string{mclBinaryPath, "--gcs-bucket=" + bucket},
						},
					},
				},
			},
		},
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

func newPeerKubeconfigSecret(namespace, name string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Data: map[string][]byte{
			kubeconfigSecretKey: []byte("apiVersion: v1\nkind: Config\n"),
		},
	}
}
