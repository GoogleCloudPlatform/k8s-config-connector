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

package controllers

import (
	"context"
	"fmt"
	"strings"

	customizev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1alpha1"
	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/scale/scheme/autoscalingv1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

var (
	ValidConfigConnectorNamespacedName = types.NamespacedName{Name: k8s.ConfigConnectorAllowedName}
)

func GetConfigConnector(ctx context.Context, client client.Client, nn types.NamespacedName) (*corev1beta1.ConfigConnector, error) {
	cc := &corev1beta1.ConfigConnector{}
	if err := client.Get(ctx, nn, cc); err != nil {
		return nil, err
	}
	return cc, nil
}

func RemoveOperatorFinalizer(o metav1.Object) (found bool) {
	var finalizers []string
	for _, f := range o.GetFinalizers() {
		if f != k8s.OperatorFinalizer {
			finalizers = append(finalizers, f)
		} else {
			found = true
		}
	}
	if found {
		o.SetFinalizers(finalizers)
	}
	return found
}

func EnsureOperatorFinalizer(o metav1.Object) (found bool) {
	for _, f := range o.GetFinalizers() {
		if f == k8s.OperatorFinalizer {
			return true
		}
	}
	o.SetFinalizers(append(o.GetFinalizers(), k8s.OperatorFinalizer))
	return false
}

func AnnotateServiceAccountObject(object *manifest.Object, gsa string) (*manifest.Object, error) {
	u := object.UnstructuredObject()
	annotations := u.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[k8s.WorkloadIdentityAnnotation] = gsa
	u.SetAnnotations(annotations)
	return manifest.NewObject(u)
}

func DeleteObject(ctx context.Context, c client.Client, obj client.Object) error {
	kind := obj.GetObjectKind().GroupVersionKind().Kind
	name := obj.(metav1.Object).GetName()
	if err := c.Delete(ctx, obj, &client.DeleteOptions{}); err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return fmt.Errorf("error deleting %v %v: %v", kind, name, err)
	}
	return removeOperatorFinalizerIfPresent(ctx, c, obj)
}

/*
 * some of the critical resources, such as role bindings, are in the customer namespace, we protect them with an
 * operator finalizer to ensure that the operator can control the time when they are removed.
 */
func removeOperatorFinalizerIfPresent(ctx context.Context, c client.Client, obj client.Object) error {
	found := RemoveOperatorFinalizer(obj)
	if !found {
		return nil
	}
	if err := c.Update(ctx, obj); err != nil {
		return fmt.Errorf("error removing operator finalizer from %v %v: %w",
			obj.GetObjectKind().GroupVersionKind().Kind, obj.GetName(), err)
	}
	return nil
}

func IsControllerManagerStatefulSet(obj *manifest.Object) bool {
	if obj.Kind != "StatefulSet" {
		return false
	}
	labels := obj.UnstructuredObject().GetLabels()
	return labels[k8s.KCCSystemComponentLabel] == k8s.KCCControllerManagerComponent
}

func IsControllerManagerService(obj *manifest.Object) bool {
	if obj.Kind != "Service" {
		return false
	}
	return obj.GetName() == k8s.NamespacedManagerServiceTmpl
}

func GetControllerResource(ctx context.Context, c client.Client, name string) (*customizev1alpha1.ControllerResource, error) {
	obj := &customizev1alpha1.ControllerResource{}
	if err := c.Get(ctx, types.NamespacedName{Name: name}, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// ListControllerResources lists all ControllerResources.
func ListControllerResources(ctx context.Context, c client.Client) ([]customizev1alpha1.ControllerResource, error) {
	list := &customizev1alpha1.ControllerResourceList{}
	if err := c.List(ctx, list); err != nil {
		return nil, err
	}
	return list.Items, nil
}

// ApplyContainerResourceCustomization applies container resource customizations specified in ControllerResource / NamespacedControllerResource CR.
func ApplyContainerResourceCustomization(isNamespaced bool, m *manifest.Objects, controllerName string, controllerGVK schema.GroupVersionKind, containers []customizev1alpha1.ContainerResourceSpec, replicas *int64) error {
	cMap := make(map[string]corev1.ResourceRequirements, len(containers)) // cMap is a map of container name to its corresponding resource customization.
	cMapApplied := make(map[string]bool)                                  // cMapApplied is a map of container name to a boolean indicating whether the customization for this container is applied.
	for _, c := range containers {
		cMap[c.Name] = c.Resources
		cMapApplied[c.Name] = false
	}
	var shouldUpdateHPA bool // shouldUpdateHPA is a boolean indicating if we need to update the "minReplicas" field in the HorizontalPodAutoscaler for webhook manager
	// apply customization to the matching controller in the manifest
	for _, item := range m.Items {
		if item.GroupVersionKind() == controllerGVK { // match GVK
			if item.GetName() == controllerName { // match exact controller name for cluster-scoped controller
				// apply container resource customization for this controller.
				if err := item.MutateContainers(customizeContainerResourcesFn(cMap, cMapApplied)); err != nil {
					return err
				}
				// apply replicas customization for this controller.
				if replicas != nil && controllerName == "cnrm-webhook-manager" { // currently only customizing replicas for cnrm-webhook-manager is supported.
					if err := item.SetNestedField(*replicas, "spec", "replicas"); err != nil {
						return err
					}
					shouldUpdateHPA = true
				}
				break // we already found the matching controller, no need to keep looking.
			}
		}
	}
	// if we update replicas for webhook manager deployment, we need to adjust its HPA as well.
	if shouldUpdateHPA {
		HPAGVK := schema.GroupVersionKind{
			Group:   autoscalingv1.SchemeGroupVersion.Group,
			Version: autoscalingv1.SchemeGroupVersion.Version,
			Kind:    "HorizontalPodAutoscaler",
		}
		for _, item := range m.Items {
			if item.GroupVersionKind() == HPAGVK && item.GetName() == "cnrm-webhook" {
				// update "minReplicas".
				if err := item.SetNestedField(*replicas, "spec", "minReplicas"); err != nil {
					return err
				}
				// update "maxReplicas" to match "minReplicas" if it is smaller than "minReplicas".
				maxReplicas, found, err := unstructured.NestedInt64(item.UnstructuredObject().Object, "spec", "maxReplicas")
				if err != nil {
					return err
				}
				if found && (maxReplicas < *replicas) {
					if err := item.SetNestedField(*replicas, "spec", "maxReplicas"); err != nil {
						return err
					}
				}
				break // we already found the HPA, no need to keep looking.
			}
		}
	}
	// check if all container resource customizations are applied
	var notApplied []string
	for c, applied := range cMapApplied {
		if !applied {
			notApplied = append(notApplied, c)
		}
	}
	if len(notApplied) > 0 {
		return fmt.Errorf("resource customization failed for the following containers because there are no matching containers in the manifest: %s", strings.Join(notApplied, ", "))
	}
	return nil
}

// customizeContainerResourcesFn returns a function to customize container resources.
func customizeContainerResourcesFn(cMap map[string]corev1.ResourceRequirements, cMapApplied map[string]bool) func(container map[string]interface{}) error {
	return func(container map[string]interface{}) error {
		name, _, err := unstructured.NestedString(container, "name")
		if err != nil {
			return fmt.Errorf("error reading container name: %v", err)
		}
		r, found := cMap[name]
		if !found {
			return nil
		}
		if r.Limits != nil && !r.Limits.Cpu().IsZero() {
			if err := unstructured.SetNestedField(container, r.Limits.Cpu().String(), "resources", "limits", "cpu"); err != nil {
				return fmt.Errorf("error setting cpu limit: %v", err)
			}
		}
		if r.Limits != nil && !r.Limits.Memory().IsZero() {
			if err := unstructured.SetNestedField(container, r.Limits.Memory().String(), "resources", "limits", "memory"); err != nil {
				return fmt.Errorf("error setting memory limit: %v", err)
			}
		}
		if r.Requests != nil && !r.Requests.Cpu().IsZero() {
			if err := unstructured.SetNestedField(container, r.Requests.Cpu().String(), "resources", "requests", "cpu"); err != nil {
				return fmt.Errorf("error setting cpu request: %v", err)
			}
		}
		if r.Requests != nil && !r.Requests.Memory().IsZero() {
			if err := unstructured.SetNestedField(container, r.Requests.Memory().String(), "resources", "requests", "memory"); err != nil {
				return fmt.Errorf("error setting memory request: %v", err)
			}
		}
		cMapApplied[name] = true
		return nil
	}
}
