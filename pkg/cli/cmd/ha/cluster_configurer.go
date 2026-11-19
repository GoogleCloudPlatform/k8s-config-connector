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
	"fmt"
	"io"
	"strings"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

const (
	// defaultMCLRolloutTimeout allows for pulling the image, and for a few restarts while the
	// IAM bindings that setup has just created take effect.
	defaultMCLRolloutTimeout      = 5 * time.Minute
	defaultMCLRolloutPollInterval = 2 * time.Second
)

// clusterConfigurer inspects a target GKE cluster and applies multi-cluster HA Kubernetes configuration to it.
type clusterConfigurer struct {
	handle          *clusterHandle
	gkeInfo         gkeClusterInfo
	configConnector *operatorv1beta1.ConfigConnector
	manifests       *manifestBuilder

	// mclRolloutTimeout and mclRolloutPollInterval control how long apply waits for the MCL
	// controller before it gives up on enabling multi-cluster leader election.
	mclRolloutTimeout      time.Duration
	mclRolloutPollInterval time.Duration
}

func newClusterConfigurer(ctx context.Context, handle *clusterHandle, mclImage, syncerImage, leaseName, leaseNamespace, bucket string) (*clusterConfigurer, error) {
	if handle == nil || handle.client == nil {
		return nil, fmt.Errorf("cluster handle has a nil Kubernetes client")
	}

	gkeInfo, err := parseGKEContext(handle.contextName)
	if err != nil {
		return nil, err
	}

	cc, err := fetchConfigConnector(ctx, handle)
	if err != nil {
		return nil, err
	}

	// GetMode applies the operator's default: an unset spec.mode means namespaced mode.
	manifests := newManifestBuilder(gkeInfo.contextName, gkeInfo.gkeCluster(), gkeInfo.project, cc.GetMode(), mclImage, syncerImage, leaseName, leaseNamespace, bucket)
	return &clusterConfigurer{
		handle:                 handle,
		gkeInfo:                gkeInfo,
		configConnector:        cc,
		manifests:              manifests,
		mclRolloutTimeout:      defaultMCLRolloutTimeout,
		mclRolloutPollInterval: defaultMCLRolloutPollInterval,
	}, nil
}

func (c *clusterConfigurer) setPeerProjects(peerProjects []string) {
	c.manifests.setPeerProjects(peerProjects)
}

func (c *clusterConfigurer) apply(ctx context.Context, stdout io.Writer, dryRun bool) error {
	objs := c.manifests.buildClusterObjects()
	if dryRun {
		if stdout != nil {
			fmt.Fprintf(stdout, "[dry-run] cluster %q: apply %d HA manifests and configure spec.experiments.multiClusterLease (gkeCluster=%q, lease=%s)\n",
				c.handle.contextName, len(objs), c.manifests.gkeCluster.ResourceName(), c.manifests.leaseKey())
		}

		return nil
	}

	var errs []error
	for _, obj := range objs {
		if err := upsertObject(ctx, c.handle.client.Client, obj); err != nil {
			errs = append(errs, fmt.Errorf("applying %s on cluster %q: %w", formatObjectID(c.handle.client.Client, obj), c.handle.contextName, err))
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	// Once spec.experiments.multiClusterLease is set, Config Connector reconciles nothing on
	// this cluster until the MCL controller confirms a leader. So only set it when that
	// controller runs.
	if err := c.waitForMCLController(ctx, stdout); err != nil {
		return err
	}

	return c.updateConfigConnector(ctx)
}

func upsertObject(ctx context.Context, k8sClient client.Client, obj client.Object) error {
	err := k8sClient.Create(ctx, obj)
	if err == nil {
		return nil
	}

	if !apierrors.IsAlreadyExists(err) {
		return err
	}

	gvk, gvkErr := apiutil.GVKForObject(obj, k8sClient.Scheme())
	if gvkErr != nil {
		return gvkErr
	}

	existing := &unstructured.Unstructured{}
	existing.SetGroupVersionKind(gvk)
	key := types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}
	if getErr := k8sClient.Get(ctx, key, existing); getErr != nil {
		return getErr
	}

	obj.SetResourceVersion(existing.GetResourceVersion())
	return k8sClient.Update(ctx, obj)
}

func formatObjectID(k8sClient client.Client, obj client.Object) string {
	kind := obj.GetObjectKind().GroupVersionKind().Kind
	if kind == "" && k8sClient.Scheme() != nil {
		if gvk, err := apiutil.GVKForObject(obj, k8sClient.Scheme()); err == nil {
			kind = gvk.Kind
		}
	}

	if obj.GetNamespace() != "" {
		return fmt.Sprintf("%s %q (namespace %q)", kind, obj.GetName(), obj.GetNamespace())
	}

	return fmt.Sprintf("%s %q", kind, obj.GetName())
}

// waitForMCLController waits until the MCL controller's Deployment has rolled out the spec that
// apply just wrote. If it doesn't within the timeout, the returned error says what its pods are
// stuck on.
func (c *clusterConfigurer) waitForMCLController(ctx context.Context, stdout io.Writer) error {
	key := types.NamespacedName{Namespace: manifestSystemNamespace, Name: mclDeploymentName}
	if stdout != nil {
		fmt.Fprintf(stdout, "cluster %q: waiting up to %s for Deployment %q to become available\n", c.handle.contextName, c.mclRolloutTimeout, key.Name)
	}

	deployment := &appsv1.Deployment{}
	var pending string
	err := wait.PollUntilContextTimeout(ctx, c.mclRolloutPollInterval, c.mclRolloutTimeout, true, func(ctx context.Context) (bool, error) {
		if err := c.handle.client.Get(ctx, key, deployment); err != nil {
			// Keep the last rollout status when the request only failed because the wait is over.
			if ctx.Err() == nil {
				pending = fmt.Sprintf("getting the Deployment failed: %v", err)
			}
			return false, nil
		}
		var done bool
		pending, done = deploymentRolloutStatus(deployment)
		return done, nil
	})
	if err == nil {
		return nil
	}
	if ctx.Err() != nil {
		return fmt.Errorf("waiting for Deployment %q on cluster %q: %w", key.Name, c.handle.contextName, ctx.Err())
	}

	details := []string{pending}
	details = append(details, c.describeUnavailablePods(ctx, deployment)...)
	return fmt.Errorf(
		"cluster %q: Deployment %q (namespace %q) did not become available within %s, so spec.experiments.multiClusterLease was left unchanged on this cluster: %s",
		c.handle.contextName, key.Name, key.Namespace, c.mclRolloutTimeout, strings.Join(details, "; "),
	)
}

// deploymentRolloutStatus reports whether a Deployment has finished rolling out its current spec,
// and what it is waiting for otherwise. It follows `kubectl rollout status`.
func deploymentRolloutStatus(deployment *appsv1.Deployment) (pending string, done bool) {
	if deployment.Generation > deployment.Status.ObservedGeneration {
		return "the Deployment controller has not processed the latest spec yet", false
	}

	wantReplicas := int32(1)
	if deployment.Spec.Replicas != nil {
		wantReplicas = *deployment.Spec.Replicas
	}
	status := deployment.Status
	switch {
	case status.UpdatedReplicas < wantReplicas:
		return fmt.Sprintf("%d of %d pods run the latest spec", status.UpdatedReplicas, wantReplicas), false
	case status.Replicas > status.UpdatedReplicas:
		return fmt.Sprintf("%d pods of an older spec are still terminating", status.Replicas-status.UpdatedReplicas), false
	case status.AvailableReplicas < status.UpdatedReplicas:
		return fmt.Sprintf("%d of %d pods are available", status.AvailableReplicas, status.UpdatedReplicas), false
	}
	return "", true
}

// describeUnavailablePods says why the pods of a Deployment aren't running, from the state of
// their containers.
func (c *clusterConfigurer) describeUnavailablePods(ctx context.Context, deployment *appsv1.Deployment) []string {
	if deployment.Spec.Selector == nil {
		return nil
	}

	pods := &corev1.PodList{}
	if err := c.handle.client.List(ctx, pods, client.InNamespace(deployment.Namespace), client.MatchingLabels(deployment.Spec.Selector.MatchLabels)); err != nil {
		return []string{fmt.Sprintf("listing its pods failed: %v", err)}
	}
	if len(pods.Items) == 0 {
		return []string{"it has no pods"}
	}

	var details []string
	for _, pod := range pods.Items {
		for _, container := range pod.Status.ContainerStatuses {
			if state := describeContainerState(container); state != "" {
				details = append(details, fmt.Sprintf("pod %q container %q %s", pod.Name, container.Name, state))
			}
		}
	}
	return details
}

func describeContainerState(container corev1.ContainerStatus) string {
	var state string
	switch {
	case container.State.Waiting != nil:
		state = "is waiting: " + container.State.Waiting.Reason
		if container.State.Waiting.Message != "" {
			state += " (" + container.State.Waiting.Message + ")"
		}
	case container.State.Terminated != nil:
		state = fmt.Sprintf("terminated: %s, exit code %d", container.State.Terminated.Reason, container.State.Terminated.ExitCode)
	default:
		return ""
	}

	if last := container.LastTerminationState.Terminated; last != nil {
		state += fmt.Sprintf("; it last exited with code %d (%s)", last.ExitCode, last.Reason)
	}
	return state
}

func (c *clusterConfigurer) updateConfigConnector(ctx context.Context) error {
	cc, err := fetchConfigConnector(ctx, c.handle)
	if err != nil {
		return err
	}

	c.manifests.configureMultiClusterLease(&cc.Spec)
	if err := c.handle.client.Update(ctx, cc); err != nil {
		return fmt.Errorf("updating ConfigConnector %q on cluster %q: %w", cc.Name, c.handle.contextName, err)
	}

	if cc.Spec.Experiments == nil || cc.Spec.Experiments.MultiClusterLease == nil {
		return fmt.Errorf(
			"ConfigConnector %q on cluster %q silently dropped spec.experiments.multiClusterLease during update; the cluster's ConfigConnector CRD schema is too old (requires KCC v1.148.0+)",
			cc.Name, c.handle.contextName,
		)
	}

	return nil
}

// clusterInfo returns the GKE project/location/cluster parsed from the cluster's context,
// used by the GCP provisioner.
func (c *clusterConfigurer) clusterInfo() gkeClusterInfo {
	return c.gkeInfo
}

func (c *clusterConfigurer) renderYAML() (string, error) {
	return c.manifests.renderClusterYAML(c.configConnector)
}
