// Copyright 2022 Google LLC
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

package configconnectorcontext

import (
	"context"
	"fmt"
	"strings"
	"time"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cluster"

	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func removeNamespacedComponents(ctx context.Context, c client.Client, objects []*manifest.Object) error {
	for _, obj := range objects {
		if err := controllers.DeleteObject(ctx, c, obj.UnstructuredObject()); err != nil {
			return err
		}
	}
	return nil
}

func transformNamespacedComponentTemplates(ctx context.Context, c client.Client, ccc *corev1beta1.ConfigConnectorContext, namespacedTemplates []*manifest.Object, managerNamespaceSuffix string) ([]*manifest.Object, error) {
	transformedObjs := make([]*manifest.Object, 0, len(namespacedTemplates))
	for _, obj := range namespacedTemplates {
		processed := obj
		if controllers.IsControllerManagerService(processed) {
			var err error
			processed, err = handleControllerManagerService(ctx, c, ccc, processed, managerNamespaceSuffix)
			if err != nil {
				return nil, err
			}
		}
		if controllers.IsControllerManagerStatefulSet(processed) {
			var err error
			processed, err = handleControllerManagerStatefulSet(ctx, c, ccc, processed, managerNamespaceSuffix)
			if err != nil {
				return nil, err
			}
		}
		processed, err := replaceNamespacePattern(processed, ccc.Namespace, managerNamespaceSuffix)
		if err != nil {
			return nil, err
		}
		if processed.Kind == rbacv1.ServiceAccountKind && strings.HasPrefix(processed.GetName(), k8s.ServiceAccountNamePrefix) {
			processed, err = controllers.AnnotateServiceAccountObject(processed, ccc.Spec.GoogleServiceAccount)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("error annotating ServiceAccount %v/%v", obj.UnstructuredObject().GetNamespace(), obj.UnstructuredObject().GetName()))
			}
		}
		transformedObjs = append(transformedObjs, processed)
	}
	return transformedObjs, nil
}

func handleControllerManagerService(ctx context.Context, c client.Client, ccc *corev1beta1.ConfigConnectorContext, obj *manifest.Object, managerNamespaceSuffix string) (*manifest.Object, error) {
	u := obj.UnstructuredObject().DeepCopy()
	nsID, err := cluster.GetNamespaceID(ctx, k8s.OperatorNamespaceIDConfigMapNN, c, ccc.Namespace)
	if err != nil {
		return nil, fmt.Errorf("error getting namespace id for namespace %v: %w", ccc.Namespace, err)
	}
	u.SetName(strings.ReplaceAll(u.GetName(), "${NAMESPACE?}", nsID))
	serviceNamespace := strings.ReplaceAll(u.GetNamespace(), "${NAMESPACE?}", ccc.Namespace)
	serviceNamespace = strings.ReplaceAll(serviceNamespace, "${MANAGER_NAMESPACE?}", replaceNamespaceSuffix(ccc.Namespace, managerNamespaceSuffix))
	if err := removeStaleControllerManagerService(ctx, c, ccc.Namespace, u.GetName(), serviceNamespace); err != nil {
		return nil, fmt.Errorf("error deleting stale Services for watched namespace %v: %w", ccc.Namespace, err)
	}
	return manifest.NewObject(u)
}

func removeStaleControllerManagerService(ctx context.Context, c client.Client, ns string, validSts string, serviceNamespace string) error {
	// List existing controller-manager services for the given namespace, delete stale ones if any
	// stale services could come from legacy naming or namespaceId changes.
	svcList := &corev1.ServiceList{}
	if err := c.List(ctx, svcList, client.InNamespace(serviceNamespace),
		client.MatchingLabels{k8s.NamespacedComponentLabel: ns}); err != nil {
		return fmt.Errorf("error listing existing %v Services for watched namespace %v: %w", k8s.KCCControllerManagerComponent, ns, err)
	}
	for _, svc := range svcList.Items {
		if strings.HasPrefix(svc.Name, k8s.NamespacedManagerServicePrefix) && svc.Name != validSts {
			if err := controllers.DeleteObject(ctx, c, &svc); err != nil {
				return err
			}
		}
	}
	return nil
}

func handleControllerManagerStatefulSet(ctx context.Context, c client.Client, ccc *corev1beta1.ConfigConnectorContext, obj *manifest.Object, managerNamespaceSuffix string) (*manifest.Object, error) {
	u := obj.UnstructuredObject().DeepCopy()

	nsID, err := cluster.GetNamespaceID(ctx, k8s.OperatorNamespaceIDConfigMapNN, c, ccc.Namespace)
	if err != nil {
		return nil, fmt.Errorf("error getting namespace id for namespace %v: %w", ccc.Namespace, err)
	}

	u.SetName(strings.ReplaceAll(u.GetName(), "${NAMESPACE?}", nsID))

	serviceName, found, err := unstructured.NestedString(u.Object, "spec", "serviceName")
	if err != nil || !found {
		return nil, fmt.Errorf("couldn't resolve serviceName in StatefulSet %v for watched namespace %v: %w", u.GetName(), ccc.Namespace, err)
	}
	if err := unstructured.SetNestedField(u.Object, strings.ReplaceAll(serviceName, "${NAMESPACE?}", nsID), "spec", "serviceName"); err != nil {
		return nil, err
	}

	if ccc.GetRequestProjectPolicy() == k8s.ResourceProjectPolicy {
		if err := enableUserProjectOverride(u); err != nil {
			return nil, fmt.Errorf("error enabling %v in StatefulSet %v for watched namespace %v: %w", k8s.UserProjectOverrideFlag, u.GetName(), ccc.Namespace, err)
		}
	}

	if ccc.GetRequestProjectPolicy() == k8s.BillingProjectPolicy {
		if err := enableUserProjectOverride(u); err != nil {
			return nil, fmt.Errorf("error enabling %v in StatefulSet %v for watched namespace %v: %w", k8s.UserProjectOverrideFlag, u.GetName(), ccc.Namespace, err)
		}
		if err := enableBillingProject(u, ccc.Spec.BillingProject); err != nil {
			return nil, fmt.Errorf("error enabling %v in StatefulSet %v for watched namespace %v: %w", k8s.BillingProjectFlag, u.GetName(), ccc.Namespace, err)
		}
	}

	statefulsetNamespace := strings.ReplaceAll(u.GetNamespace(), "${NAMESPACE?}", ccc.Namespace)
	statefulsetNamespace = strings.ReplaceAll(statefulsetNamespace, "${MANAGER_NAMESPACE?}", replaceNamespaceSuffix(ccc.Namespace, managerNamespaceSuffix))
	if err := removeStaleControllerManagerStatefulSet(ctx, c, ccc.Namespace, u.GetName(), statefulsetNamespace); err != nil {
		return nil, fmt.Errorf("error deleting stale StatefulSet for watched namespace %v: %w", ccc.Namespace, err)
	}

	return manifest.NewObject(u)
}

func enableUserProjectOverride(u *unstructured.Unstructured) error {
	return setFlagForManagerContainer(u, k8s.UserProjectOverrideFlag, "true")
}

func enableBillingProject(u *unstructured.Unstructured, flagValue string) error {
	return setFlagForManagerContainer(u, k8s.BillingProjectFlag, flagValue)
}

func findManagerContainer(containers []interface{}) (managerContainer map[string]interface{}, index int, err error) {
	for i, container := range containers {
		containerAsMap, ok := container.(map[string]interface{})
		if !ok {
			return nil, 0, fmt.Errorf("couldn't convert container configuration %v to a map", container)
		}
		name, found, err := unstructured.NestedString(containerAsMap, "name")
		if err != nil || !found {
			return nil, 0, fmt.Errorf("couldn't resolve name of container configuration %v: %w", container, err)
		}
		if name == k8s.CNRMManagerContainerName {
			return containerAsMap, i, nil
		}
	}
	return nil, 0, fmt.Errorf("no manager container found")
}

// A helper method to add optional flags for manager container.
func setFlagForManagerContainer(u *unstructured.Unstructured, flag string, flagValue string) error {
	containersPath := []string{"spec", "template", "spec", "containers"} // Path to container configurations in a StatefulSet
	containers, found, err := unstructured.NestedSlice(u.Object, containersPath...)
	if err != nil || !found {
		return fmt.Errorf("couldn't resolve containers: %w", err)
	}

	managerContainer, index, err := findManagerContainer(containers)
	if err != nil {
		return fmt.Errorf("error finding manager container: %w", err)
	}
	args, found, err := unstructured.NestedStringSlice(managerContainer, "args")
	if err != nil {
		return fmt.Errorf("couldn't resolve args of manager container %v: %w", managerContainer, err)
	}
	if !found {
		args = make([]string, 0)
	}
	newArgs := removeFlagFromArgs(args, flag)
	newArgs = append(newArgs, flag+"="+flagValue)
	if err := unstructured.SetNestedStringSlice(managerContainer, newArgs, "args"); err != nil {
		return fmt.Errorf("error setting args in manager container: %w", err)
	}

	containers[index] = managerContainer
	if err := unstructured.SetNestedSlice(u.Object, containers, containersPath...); err != nil {
		return fmt.Errorf("error setting containers: %w", err)
	}
	return nil
}

func removeFlagFromArgs(args []string, flag string) []string {
	newArgs := make([]string, 0)
	for _, a := range args {
		if !strings.HasPrefix(a, flag) {
			newArgs = append(newArgs, a)
		}
	}
	return newArgs
}

func removeStaleControllerManagerStatefulSet(ctx context.Context, c client.Client, ns string, validSts string, statefulsetNamespace string) error {
	// List existing controller-manager statefulsets for the given namespace, delete stale ones if any
	// stale statefulsets could come from legacy naming or namespaceId changes.
	stsList := &appsv1.StatefulSetList{}
	if err := c.List(ctx, stsList, client.InNamespace(statefulsetNamespace),
		client.MatchingLabels{k8s.KCCSystemComponentLabel: k8s.KCCControllerManagerComponent, k8s.NamespacedComponentLabel: ns}); err != nil {
		return fmt.Errorf("error listing existing %v StatefulSets for watched namespace %v: %w", k8s.KCCControllerManagerComponent, ns, err)
	}

	hasStale := false
	for _, sts := range stsList.Items {
		if sts.Name != validSts {
			hasStale = true
			if err := controllers.DeleteObject(ctx, c, &sts); err != nil {
				return err
			}
		}
	}

	if hasStale {
		b := wait.Backoff{
			Duration: time.Second,
			Factor:   1.2,
			Steps:    12,
		}
		podList := &corev1.PodList{}
		if err := wait.ExponentialBackoff(b, func() (done bool, err error) {
			if err := c.List(ctx, podList, client.InNamespace(statefulsetNamespace),
				client.MatchingLabels{k8s.KCCSystemComponentLabel: k8s.KCCControllerManagerComponent, k8s.NamespacedComponentLabel: ns}); err != nil {
				return false, errors.Wrap(err, "error listing controller pods")
			}
			if len(podList.Items) == 0 {
				return true, nil
			}
			if len(podList.Items) == 1 {
				pod := &podList.Items[0]
				for _, owner := range pod.OwnerReferences {
					if owner.Kind == "StatefulSet" && owner.Name == validSts {
						return true, nil
					}
				}
			}
			return false, nil
		}); err != nil {
			return errors.Wrap(err, "error waiting for stale controller pods to be deleted")
		}
	}
	return nil
}

func replaceNamespacePattern(obj *manifest.Object, ns string, managerNamespaceSuffix string) (*manifest.Object, error) {
	bytes, err := obj.JSON()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error marshalling object %v", obj.UnstructuredObject()))
	}
	str := string(bytes)
	str = strings.ReplaceAll(str, "${NAMESPACE?}", ns)
	str = strings.ReplaceAll(str, "${MANAGER_NAMESPACE?}", replaceNamespaceSuffix(ns, managerNamespaceSuffix))
	newObj, err := manifest.ParseJSONToObject([]byte(str))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error unmarshalling object %v", obj.UnstructuredObject()))
	}
	return newObj, nil
}

func replaceNamespaceSuffix(namespace, suffix string) string {
	if suffix == "" {
		return namespace
	}

	// First character of suffix is used as a delimiter.
	delimiter := suffix[0]
	lastDelimiterIndex := strings.LastIndexByte(namespace, delimiter)

	// If no delimiter is found, there's no suffix to replace.
	// Return the original string.
	if lastDelimiterIndex == -1 {
		return namespace
	}

	return namespace[0:lastDelimiterIndex] + suffix
}
