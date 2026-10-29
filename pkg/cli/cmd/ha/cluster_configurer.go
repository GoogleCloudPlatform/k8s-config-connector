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

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// clusterConfigurer inspects a target GKE cluster and applies multi-cluster HA Kubernetes configuration to it.
type clusterConfigurer struct {
	handle          *clusterHandle
	gkeInfo         gkeClusterInfo
	configConnector *operatorv1beta1.ConfigConnector
	manifests       *manifestBuilder
}

func newClusterConfigurer(ctx context.Context, handle *clusterHandle, leaseName, leaseNamespace, bucket string) (*clusterConfigurer, error) {
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

	image, err := discoverControllerImage(ctx, handle)
	if err != nil {
		return nil, err
	}

	manifests := newManifestBuilder(gkeInfo.contextName, gkeInfo.identity, gkeInfo.project, image, leaseName, leaseNamespace, bucket)
	return &clusterConfigurer{
		handle:          handle,
		gkeInfo:         gkeInfo,
		configConnector: cc,
		manifests:       manifests,
	}, nil
}

const (
	controllerManagerNamespace   = "cnrm-system"
	controllerComponentLabelKey  = "cnrm.cloud.google.com/component"
	controllerManagerStatefulSet = "cnrm-controller-manager"
)

func discoverControllerImage(ctx context.Context, handle *clusterHandle) (string, error) {
	stsList := &appsv1.StatefulSetList{}
	if err := handle.client.List(
		ctx,
		stsList,
		client.InNamespace(controllerManagerNamespace),
		client.MatchingLabels{controllerComponentLabelKey: controllerManagerStatefulSet},
	); err != nil {
		return "", fmt.Errorf("listing Config Connector controller StatefulSets in namespace %q on cluster %q: %w", controllerManagerNamespace, handle.contextName, err)
	}

	var fallbackImage string
	for _, sts := range stsList.Items {
		for _, container := range sts.Spec.Template.Spec.Containers {
			if container.Name == "manager" && container.Image != "" {
				if sts.Name == controllerManagerStatefulSet {
					return container.Image, nil
				}

				if fallbackImage == "" {
					fallbackImage = container.Image
				}
			}
		}
	}

	if fallbackImage != "" {
		return fallbackImage, nil
	}

	return "", fmt.Errorf(
		"could not find running Config Connector manager container image on cluster %q (expected StatefulSet with label %s=%s in namespace %q); verify Config Connector is installed and healthy",
		handle.contextName, controllerComponentLabelKey, controllerManagerStatefulSet, controllerManagerNamespace,
	)
}

func (c *clusterConfigurer) apply(ctx context.Context, stdout io.Writer, dryRun bool) error {
	objs := c.manifests.buildClusterObjects()
	if dryRun {
		if stdout != nil {
			fmt.Fprintf(stdout, "[dry-run] cluster %q: apply %d HA manifests and configure spec.experiments.multiClusterLease (identity=%q, lease=%s)\n",
				c.handle.contextName, len(objs), c.gkeInfo.identity, c.manifests.leaseKey())
		}

		return nil
	}

	var errs []error
	for _, obj := range objs {
		if err := upsertObject(ctx, c.handle.client.Client, obj); err != nil {
			errs = append(errs, fmt.Errorf("applying %s on cluster %q: %w", formatObjectID(c.handle.client.Client, obj), c.handle.contextName, err))
		}
	}

	if err := c.updateConfigConnector(ctx); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
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

func (c *clusterConfigurer) gkeCluster() gkeClusterInfo {
	return c.gkeInfo
}

func (c *clusterConfigurer) renderYAML() (string, error) {
	return c.manifests.renderClusterYAML(c.configConnector)
}
