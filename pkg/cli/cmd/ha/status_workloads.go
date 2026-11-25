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
	"fmt"
	"strings"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	operatork8s "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type clusterWorkloads struct {
	isHASetupManaged bool
	roleForbiddenErr error
	managers         []*managerWorkload
	mcl              *deploymentWorkload
	syncer           *deploymentWorkload
}

type managerWorkload struct {
	sts  *managerStatefulSet
	pods []corev1.Pod
}

type deploymentWorkload struct {
	deployment *appsv1.Deployment
	pods       []corev1.Pod
}

// fetchClusterWorkloads reads the workloads on a cluster and their pods.
func fetchClusterWorkloads(ctx context.Context, h *clusterHandle, cc *operatorv1beta1.ConfigConnector) (*clusterWorkloads, error) {
	mclSpec := multiClusterLeaseSpec(cc)
	cw := &clusterWorkloads{}

	// Check if cluster is managed by ha setup via role marker.
	role := &rbacv1.Role{}
	err := h.client.Get(ctx, types.NamespacedName{Namespace: mclSpec.Namespace, Name: managerHARoleName}, role)
	switch {
	case err == nil:
		if role.Labels != nil && role.Labels[manifestManagedByLabelKey] == manifestManagedByLabelValue {
			cw.isHASetupManaged = true
		}
	case apierrors.IsNotFound(err):
		cw.isHASetupManaged = false
	case apierrors.IsForbidden(err):
		cw.isHASetupManaged = false
		cw.roleForbiddenErr = err
	default:
		return nil, err
	}

	// Fetch Manager StatefulSets and their pods.
	managers, err := h.listManagers(ctx)
	if err != nil {
		return nil, err
	}
	managerPods := &corev1.PodList{}
	podLabels := client.MatchingLabels{operatork8s.KCCSystemComponentLabel: operatork8s.KCCControllerManagerComponent}
	if err := h.client.List(ctx, managerPods, podLabels); err != nil {
		return nil, err
	}
	for _, m := range managers {
		mw := &managerWorkload{sts: m}
		if m.Spec.Selector != nil {
			selector, err := metav1.LabelSelectorAsSelector(m.Spec.Selector)
			if err != nil {
				return nil, err
			}
			for _, pod := range managerPods.Items {
				if pod.Namespace == m.Namespace && selector.Matches(labels.Set(pod.Labels)) {
					mw.pods = append(mw.pods, pod)
				}
			}
		}
		cw.managers = append(cw.managers, mw)
	}

	// Fetch MCL controller Deployment and its pods.
	mclDep := &appsv1.Deployment{}
	err = h.client.Get(ctx, types.NamespacedName{Namespace: manifestSystemNamespace, Name: mclDeploymentName}, mclDep)
	switch {
	case err == nil:
		dw := &deploymentWorkload{deployment: mclDep}
		if mclDep.Spec.Selector != nil {
			pods := &corev1.PodList{}
			if err := h.client.List(ctx, pods, client.InNamespace(manifestSystemNamespace), client.MatchingLabels(mclDep.Spec.Selector.MatchLabels)); err != nil {
				return nil, err
			}
			dw.pods = pods.Items
		}
		cw.mcl = dw
	case apierrors.IsNotFound(err):
		cw.mcl = nil
	default:
		return nil, err
	}

	// Fetch Syncer controller Deployment and its pods if replication mode is not Disabled.
	if replicationMode(mclSpec) != replicationModeDisabled {
		syncerDep := &appsv1.Deployment{}
		err = h.client.Get(ctx, types.NamespacedName{Namespace: manifestSystemNamespace, Name: syncerDeploymentName}, syncerDep)
		switch {
		case err == nil:
			dw := &deploymentWorkload{deployment: syncerDep}
			if syncerDep.Spec.Selector != nil {
				pods := &corev1.PodList{}
				if err := h.client.List(ctx, pods, client.InNamespace(manifestSystemNamespace), client.MatchingLabels(syncerDep.Spec.Selector.MatchLabels)); err != nil {
					return nil, err
				}
				dw.pods = pods.Items
			}
			cw.syncer = dw
		case apierrors.IsNotFound(err):
			cw.syncer = nil
		default:
			return nil, err
		}
	}

	return cw, nil
}

// checkWorkloads evaluates the workloads on a cluster against checks 1-5 and missing workload rules.
func (cw *clusterWorkloads) checkWorkloads(contextName string, state *clusterHAState, now time.Time) ([]string, []string) {
	var problems []string
	var warnings []string
	mclSpec := multiClusterLeaseSpec(state.configConnector)

	if cw.roleForbiddenErr != nil {
		warnings = append(warnings, fmt.Sprintf("cluster %q: could not read Role %q in namespace %q: %v", contextName, managerHARoleName, mclSpec.Namespace, cw.roleForbiddenErr))
	}

	// Missing workloads check
	if len(cw.managers) == 0 {
		if state.configConnector.GetMode() != operatork8s.NamespacedMode {
			msg := fmt.Sprintf("cluster %q: Config Connector's manager StatefulSet not found", contextName)
			if cw.isHASetupManaged {
				problems = append(problems, msg)
			} else {
				warnings = append(warnings, msg)
			}
		}
	}

	if cw.mcl == nil {
		msg := fmt.Sprintf("cluster %q: Deployment %q not found in namespace %q", contextName, mclDeploymentName, manifestSystemNamespace)
		if cw.isHASetupManaged {
			problems = append(problems, msg)
		} else {
			warnings = append(warnings, msg)
		}
	}

	if replicationMode(mclSpec) != replicationModeDisabled && cw.syncer == nil {
		msg := fmt.Sprintf("cluster %q: Deployment %q not found in namespace %q", contextName, syncerDeploymentName, manifestSystemNamespace)
		if cw.isHASetupManaged {
			problems = append(problems, msg)
		} else {
			warnings = append(warnings, msg)
		}
	}

	reportedWorkloads := make(map[interface{}]bool)

	// Checks 1-3 on Manager StatefulSets
	for _, mw := range cw.managers {
		m := mw.sts
		if m.replicas() == 0 {
			if _, annotated := m.Annotations[failoverReplicasAnnotation]; annotated {
				problems = append(problems, fmt.Sprintf("cluster %q: Config Connector's manager is stopped: failover scaled StatefulSet %s to 0 and has not started it again; if no failover is running, run \"config-connector ha failover\" again, which starts it", contextName, m.key()))
			} else {
				problems = append(problems, fmt.Sprintf("cluster %q: Config Connector's manager is stopped: StatefulSet %s is scaled to 0", contextName, m.key()))
			}
			reportedWorkloads[m.StatefulSet] = true
		} else if pending := m.pendingReady(); pending != "" {
			details := describePodContainers(mw.pods)
			msg := fmt.Sprintf("Config Connector's manager is not ready: %s", pending)
			if len(details) > 0 {
				msg = strings.Join(append([]string{msg}, details...), "; ")
			}
			problems = append(problems, fmt.Sprintf("cluster %q: %s", contextName, msg))
			reportedWorkloads[m.StatefulSet] = true
		}
	}

	// Check 4 on MCL Deployment
	if cw.mcl != nil {
		if prob, reported := checkDeployment(contextName, cw.mcl); reported {
			problems = append(problems, prob)
			reportedWorkloads[cw.mcl.deployment] = true
		}
	}

	// Check 4 on Syncer Deployment
	if cw.syncer != nil && replicationMode(mclSpec) != replicationModeDisabled {
		if prob, reported := checkDeployment(contextName, cw.syncer); reported {
			problems = append(problems, prob)
			reportedWorkloads[cw.syncer.deployment] = true
		}
	}

	// Check 5: Settling rule for restarted containers on workloads NOT in reportedWorkloads
	for _, mw := range cw.managers {
		if !reportedWorkloads[mw.sts.StatefulSet] {
			problems = append(problems, checkRestartedContainers(contextName, mw.pods, now)...)
		}
	}
	if cw.mcl != nil && !reportedWorkloads[cw.mcl.deployment] {
		problems = append(problems, checkRestartedContainers(contextName, cw.mcl.pods, now)...)
	}
	if cw.syncer != nil && replicationMode(mclSpec) != replicationModeDisabled && !reportedWorkloads[cw.syncer.deployment] {
		problems = append(problems, checkRestartedContainers(contextName, cw.syncer.pods, now)...)
	}

	return problems, warnings
}

func checkDeployment(contextName string, dw *deploymentWorkload) (string, bool) {
	if dw == nil {
		return "", false
	}
	d := dw.deployment
	wantReplicas := int32(1)
	if d.Spec.Replicas != nil {
		wantReplicas = *d.Spec.Replicas
	}
	if wantReplicas == 0 {
		return fmt.Sprintf("cluster %q: Deployment %q is stopped: scaled to 0", contextName, d.Name), true
	}
	if pending, done := deploymentRolloutStatus(d); !done {
		details := describeUnavailablePods(dw.pods)
		msg := fmt.Sprintf("Deployment %q is not available: %s", d.Name, pending)
		if len(details) > 0 {
			msg = strings.Join(append([]string{msg}, details...), "; ")
		}
		return fmt.Sprintf("cluster %q: %s", contextName, msg), true
	}
	return "", false
}

func checkRestartedContainers(contextName string, pods []corev1.Pod, now time.Time) []string {
	var problems []string
	for _, pod := range pods {
		for _, container := range pod.Status.ContainerStatuses {
			if container.State.Running == nil || container.RestartCount == 0 {
				continue
			}
			runningUptime := now.Sub(container.State.Running.StartedAt.Time)
			if runningUptime < 0 {
				runningUptime = 0
			}

			maxAllowedUptime := 3 * time.Minute
			var lastTerminated *corev1.ContainerStateTerminated
			var prevRunDuration time.Duration
			if last := container.LastTerminationState.Terminated; last != nil && !last.StartedAt.IsZero() && !last.FinishedAt.IsZero() {
				prevRun := last.FinishedAt.Time.Sub(last.StartedAt.Time)
				if prevRun > 0 {
					prevRunDuration = prevRun
					lastTerminated = last
					threshold := prevRunDuration + 30*time.Second
					if threshold < maxAllowedUptime {
						maxAllowedUptime = threshold
					}
				}
			}

			if runningUptime < maxAllowedUptime {
				var msg string
				if lastTerminated != nil {
					msg = fmt.Sprintf("cluster %q: container %q of pod %s/%s restarted %s ago: the container before it ran for %s and exited with code %d (%s); %d restarts",
						contextName, container.Name, pod.Namespace, pod.Name, runningUptime.String(), prevRunDuration.String(), lastTerminated.ExitCode, lastTerminated.Reason, container.RestartCount)
				} else {
					msg = fmt.Sprintf("cluster %q: container %q of pod %s/%s restarted %s ago; %d restarts",
						contextName, container.Name, pod.Namespace, pod.Name, runningUptime.String(), container.RestartCount)
				}
				problems = append(problems, msg)
			}
		}
	}
	return problems
}

func describePodContainers(pods []corev1.Pod) []string {
	var details []string
	for _, pod := range pods {
		for _, container := range pod.Status.ContainerStatuses {
			if state := describeContainerState(container); state != "" {
				details = append(details, fmt.Sprintf("pod %q container %q %s", pod.Name, container.Name, state))
			}
		}
	}
	return details
}
