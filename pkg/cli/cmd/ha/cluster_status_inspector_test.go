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
	"context"
	"slices"
	"testing"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestClusterStatusInspectorDiscoverGCSBucket(t *testing.T) {
	handle := newFakeClusterHandle(t, "gke_proj_us-central1_cluster-a", newStatusMCLDeployment("my-ha-bucket"))
	inspector := newClusterStatusInspector(handle, time.Now().UTC())

	gotBucket := inspector.discoverGCSBucket(context.Background())

	if gotBucket != "my-ha-bucket" {
		t.Fatalf("got bucket %q, want %q", gotBucket, "my-ha-bucket")
	}
}

func TestClusterStatusInspectorDetectsStaleLeaseAndSyncerErrors(t *testing.T) {
	now := time.Date(2026, 10, 28, 3, 5, 0, 0, time.UTC)
	staleRenewStr := "2026-10-28T03:00:00Z"
	duration := int32(15)
	transitions := int32(1)
	handle := newFakeClusterHandle(
		t,
		"gke_proj_us-central1_cluster-a",
		newConfigConnectorWithMCL("kcc-multicluster-lease", "cnrm-system", "cluster-a"),
		newStatusMultiClusterLease("cnrm-system", "kcc-multicluster-lease", "cluster-a", "cluster-a", staleRenewStr, duration, transitions),
		newStatusKRMSyncer("cnrm-system", "krmsyncer-cluster", false, "cluster-b", []metav1.Condition{
			{Type: "InvalidRule", Status: metav1.ConditionTrue, Reason: "RuleValidationFailed", Message: "unknown GVK in rule 0"},
		}),
	)
	inspector := newClusterStatusInspector(handle, now)

	report, state, warnings := inspector.inspect(context.Background())

	if state == nil || report.role != "LEADER" || len(report.syncers) != 1 {
		t.Fatalf("unexpected report or nil state: %+v", report)
	}
	wantWarnings := []string{
		`stale MultiClusterLease status on cluster "gke_proj_us-central1_cluster-a": leader "cluster-a" last renewed 5m0s ago (at 2026-10-28T03:00:00Z, staleness threshold 30s)`,
		`cluster "gke_proj_us-central1_cluster-a" KRMSyncer "cnrm-system/krmsyncer-cluster" reported error condition InvalidRule=True (RuleValidationFailed): unknown GVK in rule 0`,
	}
	if !slices.Equal(warnings, wantWarnings) {
		t.Fatalf("got warnings %v, want %v", warnings, wantWarnings)
	}
}

func newStatusMCLDeployment(bucket string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "multiclusterlease-controller-manager",
			Namespace: "cnrm-system",
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "manager",
							Command: []string{"/configconnector/mcl-manager", "--gcs-bucket=" + bucket},
						},
					},
				},
			},
		},
	}
}
