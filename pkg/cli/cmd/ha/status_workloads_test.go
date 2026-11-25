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
	"errors"
	"strings"
	"testing"
	"time"

	operatork8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type failingPodListClient struct {
	client.Client
	err error
}

func (c *failingPodListClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	if _, ok := list.(*corev1.PodList); ok {
		return c.err
	}
	return c.Client.List(ctx, list, opts...)
}

func TestStatusWorkloadChecks(t *testing.T) {
	tests := []struct {
		name         string
		now          time.Time
		mutate       func(fs haFixtures)
		wantState    haState
		wantProblems []string
		wantWarnings []string
	}{
		{
			name:         "ha setup managed pair with healthy workloads, no restarts",
			now:          testNow,
			mutate:       func(fs haFixtures) {},
			wantState:    haStateHealthy,
			wantProblems: nil,
			wantWarnings: nil,
		},
		{
			name: "custom install missing standard Deployments and manager StatefulSet",
			now:  testNow,
			mutate: func(fs haFixtures) {
				// Remove role and all workload objects from both clusters and set cluster mode
				for _, f := range fs {
					f.cc.Spec.Mode = "cluster"
					f.extraObjs = nil
				}
			},
			wantState:    haStateHealthy,
			wantProblems: nil,
			wantWarnings: []string{
				`cluster "gke_proj-a_us-central1_cluster-a": Config Connector's manager StatefulSet not found`,
				`cluster "gke_proj-a_us-central1_cluster-a": Deployment "multiclusterlease-controller-manager" not found in namespace "cnrm-system"`,
				`cluster "gke_proj-a_us-central1_cluster-a": Deployment "syncer-controller-manager" not found in namespace "cnrm-system"`,
				`cluster "gke_proj-b_us-east1_cluster-b": Config Connector's manager StatefulSet not found`,
				`cluster "gke_proj-b_us-east1_cluster-b": Deployment "multiclusterlease-controller-manager" not found in namespace "cnrm-system"`,
				`cluster "gke_proj-b_us-east1_cluster-b": Deployment "syncer-controller-manager" not found in namespace "cnrm-system"`,
			},
		},
		{
			name: "custom install with present standard-named Deployment unavailable",
			now:  testNow,
			mutate: func(fs haFixtures) {
				// Remove role (making it custom install) but keep Deployments, making MCL deployment unavailable
				for _, f := range fs {
					f.cc.Spec.Mode = "cluster"
					var filtered []runtime.Object
					for _, obj := range f.extraObjs {
						if _, ok := obj.(*rbacv1.Role); !ok {
							filtered = append(filtered, obj)
						}
					}
					f.extraObjs = filtered
				}
				// Make MCL deployment on cluster-b unavailable
				for _, obj := range fs[1].extraObjs {
					if d, ok := obj.(*appsv1.Deployment); ok && d.Name == mclDeploymentName {
						d.Status.AvailableReplicas = 0
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Deployment "multiclusterlease-controller-manager" is not available: 0 of 1 pods are available`,
			},
			wantWarnings: nil,
		},
		{
			name: "ha setup managed cluster missing MCL Deployment",
			now:  testNow,
			mutate: func(fs haFixtures) {
				var filtered []runtime.Object
				for _, obj := range fs[1].extraObjs {
					if d, ok := obj.(*appsv1.Deployment); ok && d.Name == mclDeploymentName {
						continue
					}
					filtered = append(filtered, obj)
				}
				fs[1].extraObjs = filtered
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Deployment "multiclusterlease-controller-manager" not found in namespace "cnrm-system"`,
			},
		},
		{
			name: "ha setup managed cluster missing syncer Deployment",
			now:  testNow,
			mutate: func(fs haFixtures) {
				var filtered []runtime.Object
				for _, obj := range fs[1].extraObjs {
					if d, ok := obj.(*appsv1.Deployment); ok && d.Name == syncerDeploymentName {
						continue
					}
					filtered = append(filtered, obj)
				}
				fs[1].extraObjs = filtered
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Deployment "syncer-controller-manager" not found in namespace "cnrm-system"`,
			},
		},
		{
			name: "ha setup managed cluster missing syncer Deployment with resourceReplicationMode Disabled",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, f := range fs {
					f.cc.Spec.Experiments.MultiClusterLease.ResourceReplicationMode = replicationModeDisabled
					f.syncers = nil
				}
				var filtered []runtime.Object
				for _, obj := range fs[1].extraObjs {
					if d, ok := obj.(*appsv1.Deployment); ok && d.Name == syncerDeploymentName {
						continue
					}
					filtered = append(filtered, obj)
				}
				fs[1].extraObjs = filtered
			},
			wantState:    haStateHealthy,
			wantProblems: nil,
		},
		{
			name: "ha setup managed cluster in cluster mode with 0 manager StatefulSets",
			now:  testNow,
			mutate: func(fs haFixtures) {
				fs[1].cc.Spec.Mode = "cluster"
				var filtered []runtime.Object
				for _, obj := range fs[1].extraObjs {
					if _, ok := obj.(*appsv1.StatefulSet); ok {
						continue
					}
					filtered = append(filtered, obj)
				}
				fs[1].extraObjs = filtered
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Config Connector's manager StatefulSet not found`,
			},
		},
		{
			name: "ha setup managed cluster in namespaced mode with 0 manager StatefulSets",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, f := range fs {
					f.cc.Spec.Mode = "namespaced"
				}
				var filtered []runtime.Object
				for _, obj := range fs[1].extraObjs {
					if _, ok := obj.(*appsv1.StatefulSet); ok {
						continue
					}
					filtered = append(filtered, obj)
				}
				fs[1].extraObjs = filtered
			},
			wantState:    haStateHealthy,
			wantProblems: nil,
		},
		{
			name: "Manager StatefulSet: 0 replicas with failoverReplicasAnnotation",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if m, ok := obj.(*appsv1.StatefulSet); ok {
						m.Spec.Replicas = ptrInt32(0)
						m.Annotations = map[string]string{failoverReplicasAnnotation: "1"}
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Config Connector's manager is stopped: failover scaled StatefulSet cnrm-system/cnrm-controller-manager-system to 0 and has not started it again; if no failover is running, run "config-connector ha failover" again, which starts it`,
			},
		},
		{
			name: "Manager StatefulSet: 0 replicas without annotation",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if m, ok := obj.(*appsv1.StatefulSet); ok {
						m.Spec.Replicas = ptrInt32(0)
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Config Connector's manager is stopped: StatefulSet cnrm-system/cnrm-controller-manager-system is scaled to 0`,
			},
		},
		{
			name: "Leader manager StatefulSet scaled to 0",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[0].extraObjs {
					if m, ok := obj.(*appsv1.StatefulSet); ok {
						m.Spec.Replicas = ptrInt32(0)
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-a_us-central1_cluster-a": Config Connector's manager is stopped: StatefulSet cnrm-system/cnrm-controller-manager-system is scaled to 0`,
			},
		},
		{
			name: "Manager StatefulSet: 1 replica, 0 ready; container waiting CrashLoopBackOff",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if m, ok := obj.(*appsv1.StatefulSet); ok {
						m.Status.ReadyReplicas = 0
					}
					if p, ok := obj.(*corev1.Pod); ok && strings.HasPrefix(p.Name, "cnrm-controller-manager-system") {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Waiting: &corev1.ContainerStateWaiting{
										Reason:  "CrashLoopBackOff",
										Message: "back-off 5m0s",
									},
								},
								LastTerminationState: corev1.ContainerState{
									Terminated: &corev1.ContainerStateTerminated{
										ExitCode: 1,
										Reason:   "Error",
									},
								},
							},
						}
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Config Connector's manager is not ready: StatefulSet cnrm-system/cnrm-controller-manager-system has 0 of 1 pods ready; pod "cnrm-controller-manager-system-0" container "manager" is waiting: CrashLoopBackOff (back-off 5m0s); it last exited with code 1 (Error)`,
			},
		},
		{
			name: "Namespaced mode: two manager StatefulSets, one not ready",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, f := range fs {
					f.cc.Spec.Mode = "namespaced"
				}
				stsTeamA := &appsv1.StatefulSet{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cnrm-controller-manager-team-a",
						Namespace: "cnrm-system",
						Labels:    map[string]string{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent},
					},
					Spec: appsv1.StatefulSetSpec{
						Replicas: ptrInt32(1),
						Selector: &metav1.LabelSelector{
							MatchLabels: map[string]string{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent, "team": "a"},
						},
					},
					Status: appsv1.StatefulSetStatus{Replicas: 1, ReadyReplicas: 1},
				}
				stsTeamB := &appsv1.StatefulSet{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cnrm-controller-manager-team-b",
						Namespace: "cnrm-system",
						Labels:    map[string]string{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent},
					},
					Spec: appsv1.StatefulSetSpec{
						Replicas: ptrInt32(1),
						Selector: &metav1.LabelSelector{
							MatchLabels: map[string]string{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent, "team": "b"},
						},
					},
					Status: appsv1.StatefulSetStatus{Replicas: 1, ReadyReplicas: 0},
				}
				fs[1].extraObjs = append(fs[1].extraObjs, stsTeamA, stsTeamB)
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Config Connector's manager is not ready: StatefulSet cnrm-system/cnrm-controller-manager-team-b has 0 of 1 pods ready`,
			},
		},
		{
			name: "MCL Deployment scaled to 0",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if d, ok := obj.(*appsv1.Deployment); ok && d.Name == mclDeploymentName {
						d.Spec.Replicas = ptrInt32(0)
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Deployment "multiclusterlease-controller-manager" is stopped: scaled to 0`,
			},
		},
		{
			name: "MCL Deployment: 0 of 1 available; container waiting ImagePullBackOff",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if d, ok := obj.(*appsv1.Deployment); ok && d.Name == mclDeploymentName {
						d.Status.AvailableReplicas = 0
					}
					if p, ok := obj.(*corev1.Pod); ok && p.Name == mclDeploymentName+"-pod" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Waiting: &corev1.ContainerStateWaiting{
										Reason:  "ImagePullBackOff",
										Message: "Back-off pulling image",
									},
								},
							},
						}
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Deployment "multiclusterlease-controller-manager" is not available: 0 of 1 pods are available; pod "multiclusterlease-controller-manager-pod" container "manager" is waiting: ImagePullBackOff (Back-off pulling image)`,
			},
		},
		{
			name: "Syncer Deployment: 0 of 1 available; container terminated exit code 1",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if d, ok := obj.(*appsv1.Deployment); ok && d.Name == syncerDeploymentName {
						d.Status.AvailableReplicas = 0
					}
					if p, ok := obj.(*corev1.Pod); ok && p.Name == syncerDeploymentName+"-pod" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Terminated: &corev1.ContainerStateTerminated{
										ExitCode: 1,
										Reason:   "Error",
									},
								},
							},
						}
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": Deployment "syncer-controller-manager" is not available: 0 of 1 pods are available; pod "syncer-controller-manager-pod" container "manager" terminated: Error, exit code 1`,
			},
		},
		{
			name: "Syncer available; container running 94s, 5 restarts, previous ran 121s exit code 1",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if p, ok := obj.(*corev1.Pod); ok && p.Name == syncerDeploymentName+"-pod" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Running: &corev1.ContainerStateRunning{
										StartedAt: metav1.NewTime(testNow.Add(-94 * time.Second)),
									},
								},
								RestartCount: 5,
								LastTerminationState: corev1.ContainerState{
									Terminated: &corev1.ContainerStateTerminated{
										StartedAt:  metav1.NewTime(testNow.Add(-94*time.Second - 121*time.Second)),
										FinishedAt: metav1.NewTime(testNow.Add(-94 * time.Second)),
										ExitCode:   1,
										Reason:     "Error",
									},
								},
							},
						}
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": container "manager" of pod cnrm-system/syncer-controller-manager-pod restarted 1m34s ago: the container before it ran for 2m1s and exited with code 1 (Error); 5 restarts`,
			},
		},
		{
			name: "Syncer available; container running 152s, 5 restarts, previous ran 121s settled",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if p, ok := obj.(*corev1.Pod); ok && p.Name == syncerDeploymentName+"-pod" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Running: &corev1.ContainerStateRunning{
										StartedAt: metav1.NewTime(testNow.Add(-152 * time.Second)),
									},
								},
								RestartCount: 5,
								LastTerminationState: corev1.ContainerState{
									Terminated: &corev1.ContainerStateTerminated{
										StartedAt:  metav1.NewTime(testNow.Add(-152*time.Second - 121*time.Second)),
										FinishedAt: metav1.NewTime(testNow.Add(-152 * time.Second)),
										ExitCode:   1,
										Reason:     "Error",
									},
								},
							},
						}
					}
				}
			},
			wantState:    haStateHealthy,
			wantProblems: nil,
		},
		{
			name: "MCL available; container running 20s, 4 restarts, previous ran 3s",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if p, ok := obj.(*corev1.Pod); ok && p.Name == mclDeploymentName+"-pod" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Running: &corev1.ContainerStateRunning{
										StartedAt: metav1.NewTime(testNow.Add(-20 * time.Second)),
									},
								},
								RestartCount: 4,
								LastTerminationState: corev1.ContainerState{
									Terminated: &corev1.ContainerStateTerminated{
										StartedAt:  metav1.NewTime(testNow.Add(-20*time.Second - 3*time.Second)),
										FinishedAt: metav1.NewTime(testNow.Add(-20 * time.Second)),
										ExitCode:   1,
										Reason:     "Error",
									},
								},
							},
						}
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": container "manager" of pod cnrm-system/multiclusterlease-controller-manager-pod restarted 20s ago: the container before it ran for 3s and exited with code 1 (Error); 4 restarts`,
			},
		},
		{
			name: "MCL available; container running 40s, 4 restarts, previous ran 3s settled",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if p, ok := obj.(*corev1.Pod); ok && p.Name == mclDeploymentName+"-pod" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Running: &corev1.ContainerStateRunning{
										StartedAt: metav1.NewTime(testNow.Add(-40 * time.Second)),
									},
								},
								RestartCount: 4,
								LastTerminationState: corev1.ContainerState{
									Terminated: &corev1.ContainerStateTerminated{
										StartedAt:  metav1.NewTime(testNow.Add(-40*time.Second - 3*time.Second)),
										FinishedAt: metav1.NewTime(testNow.Add(-40 * time.Second)),
										ExitCode:   1,
										Reason:     "Error",
									},
								},
							},
						}
					}
				}
			},
			wantState:    haStateHealthy,
			wantProblems: nil,
		},
		{
			name: "Manager ready; container running 2m, 1 restart, previous ran 1h",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if p, ok := obj.(*corev1.Pod); ok && p.Name == "cnrm-controller-manager-system-0" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Running: &corev1.ContainerStateRunning{
										StartedAt: metav1.NewTime(testNow.Add(-2 * time.Minute)),
									},
								},
								RestartCount: 1,
								LastTerminationState: corev1.ContainerState{
									Terminated: &corev1.ContainerStateTerminated{
										StartedAt:  metav1.NewTime(testNow.Add(-2*time.Minute - 1*time.Hour)),
										FinishedAt: metav1.NewTime(testNow.Add(-2 * time.Minute)),
										ExitCode:   1,
										Reason:     "Error",
									},
								},
							},
						}
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": container "manager" of pod cnrm-system/cnrm-controller-manager-system-0 restarted 2m0s ago: the container before it ran for 1h0m0s and exited with code 1 (Error); 1 restarts`,
			},
		},
		{
			name: "Manager ready; container running 3m1s, 1 restart, previous ran 1h settled",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if p, ok := obj.(*corev1.Pod); ok && p.Name == "cnrm-controller-manager-system-0" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Running: &corev1.ContainerStateRunning{
										StartedAt: metav1.NewTime(testNow.Add(-181 * time.Second)),
									},
								},
								RestartCount: 1,
								LastTerminationState: corev1.ContainerState{
									Terminated: &corev1.ContainerStateTerminated{
										StartedAt:  metav1.NewTime(testNow.Add(-181*time.Second - 1*time.Hour)),
										FinishedAt: metav1.NewTime(testNow.Add(-181 * time.Second)),
										ExitCode:   1,
										Reason:     "Error",
									},
								},
							},
						}
					}
				}
			},
			wantState:    haStateHealthy,
			wantProblems: nil,
		},
		{
			name: "Container with 5 restarts and no lastState.terminated times, running 2m",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if p, ok := obj.(*corev1.Pod); ok && p.Name == syncerDeploymentName+"-pod" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Running: &corev1.ContainerStateRunning{
										StartedAt: metav1.NewTime(testNow.Add(-2 * time.Minute)),
									},
								},
								RestartCount: 5,
							},
						}
					}
				}
			},
			wantState: haStateError,
			wantProblems: []string{
				`cluster "gke_proj-b_us-east1_cluster-b": container "manager" of pod cnrm-system/syncer-controller-manager-pod restarted 2m0s ago; 5 restarts`,
			},
		},
		{
			name: "Container with 5 restarts and no lastState.terminated times, running 3m1s settled",
			now:  testNow,
			mutate: func(fs haFixtures) {
				for _, obj := range fs[1].extraObjs {
					if p, ok := obj.(*corev1.Pod); ok && p.Name == syncerDeploymentName+"-pod" {
						p.Status.ContainerStatuses = []corev1.ContainerStatus{
							{
								Name: "manager",
								State: corev1.ContainerState{
									Running: &corev1.ContainerStateRunning{
										StartedAt: metav1.NewTime(testNow.Add(-181 * time.Second)),
									},
								},
								RestartCount: 5,
							},
						}
					}
				}
			},
			wantState:    haStateHealthy,
			wantProblems: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fs := healthyHAFixtures()
			tc.mutate(fs)

			report := executeStatus(context.Background(), fs.handlesWithPendingDeployments(t), tc.now)

			if report.state != tc.wantState {
				t.Fatalf("got state %q, want %q; problems: %v", report.state, tc.wantState, report.problems)
			}
			if len(tc.wantProblems) == 0 && len(report.problems) > 0 {
				t.Fatalf("expected no problems, got %v", report.problems)
			}
			for _, wantP := range tc.wantProblems {
				found := false
				for _, p := range report.problems {
					if strings.Contains(p, wantP) {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("expected problem containing %q, got problems:\n%s", wantP, strings.Join(report.problems, "\n"))
				}
			}
			for _, wantW := range tc.wantWarnings {
				found := false
				for _, w := range report.warnings {
					if strings.Contains(w, wantW) {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("expected warning containing %q, got warnings:\n%s", wantW, strings.Join(report.warnings, "\n"))
				}
			}
		})
	}
}

func TestStatusReadingPodsFails(t *testing.T) {
	fs := healthyHAFixtures()
	handles := fs.handles(t)
	// Inject a failing client on cluster B for pod listing
	handles[1].client.Client = &failingPodListClient{
		Client: handles[1].client.Client,
		err:    errors.New("API server timeout"),
	}

	report := executeStatus(context.Background(), handles, testNow)

	if report.state != haStateError {
		t.Fatalf("got state %q, want %q", report.state, haStateError)
	}
	wantSub := `cluster "gke_proj-b_us-east1_cluster-b": could not read Config Connector's workloads: API server timeout`
	found := false
	for _, p := range report.problems {
		if strings.Contains(p, wantSub) {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected problem containing %q, got problems:\n%s", wantSub, strings.Join(report.problems, "\n"))
	}
	// Check that cluster B's HA column is still "Configured"
	for _, c := range report.clusters {
		if c.gke.contextName == testStatusContextB {
			if c.ha != clusterHAConfigured {
				t.Fatalf("expected cluster B HA column to be %q, got %q", clusterHAConfigured, c.ha)
			}
		}
	}
}
