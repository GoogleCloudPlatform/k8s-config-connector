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

	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"

	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
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
		return fmt.Errorf("error deleting %v %v: %w", kind, name, err)
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

func GetControllerResource(ctx context.Context, c client.Client, name string) (*customizev1beta1.ControllerResource, error) {
	obj := &customizev1beta1.ControllerResource{}
	if err := c.Get(ctx, types.NamespacedName{Name: name}, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func GetNamespacedControllerResource(ctx context.Context, c client.Client, namespace, name string) (*customizev1beta1.NamespacedControllerResource, error) {
	obj := &customizev1beta1.NamespacedControllerResource{}
	if err := c.Get(ctx, types.NamespacedName{Namespace: namespace, Name: name}, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// ListControllerResources lists all ControllerResources.
func ListControllerResources(ctx context.Context, c client.Client) ([]customizev1beta1.ControllerResource, error) {
	list := &customizev1beta1.ControllerResourceList{}
	if err := c.List(ctx, list); err != nil {
		return nil, err
	}
	return list.Items, nil
}

// ListNamespacedControllerResources lists all NamespacedControllerResource CRs in the given namespace.
func ListNamespacedControllerResources(ctx context.Context, c client.Client, namespace string) ([]customizev1beta1.NamespacedControllerResource, error) {
	list := &customizev1beta1.NamespacedControllerResourceList{}
	if err := c.List(ctx, list, &client.ListOptions{Namespace: namespace}); err != nil {
		return nil, err
	}
	return list.Items, nil
}

func GetNamespacedControllerReconciler(ctx context.Context, c client.Client, namespace, name string) (*customizev1beta1.NamespacedControllerReconciler, error) {
	obj := &customizev1beta1.NamespacedControllerReconciler{}
	if err := c.Get(ctx, types.NamespacedName{Namespace: namespace, Name: name}, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// ListNamespacedControllerReconcilers lists all NamespacedControllerReconcilers CRs in the given namespace.
func ListNamespacedControllerReconcilers(ctx context.Context, c client.Client, namespace string) ([]customizev1beta1.NamespacedControllerReconciler, error) {
	list := &customizev1beta1.NamespacedControllerReconcilerList{}
	if err := c.List(ctx, list, &client.ListOptions{Namespace: namespace}); err != nil {
		return nil, err
	}
	return list.Items, nil
}

func GetControllerReconciler(ctx context.Context, c client.Client, name string) (*customizev1beta1.ControllerReconciler, error) {
	obj := &customizev1beta1.ControllerReconciler{}
	if err := c.Get(ctx, types.NamespacedName{Name: name}, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func ListControllerReconcilers(ctx context.Context, c client.Client) ([]customizev1beta1.ControllerReconciler, error) {
	list := &customizev1beta1.ControllerReconcilerList{}
	if err := c.List(ctx, list); err != nil {
		return nil, err
	}
	return list.Items, nil
}

// ApplyContainerResourceCustomization applies container resource customizations specified in ControllerResource / NamespacedControllerResource CR.
func ApplyContainerResourceCustomization(isNamespaced bool, m *manifest.Objects, controllerName string, controllerGVK schema.GroupVersionKind, containers []customizev1beta1.ContainerResourceSpec, replicas *int64) error {
	if err := checkForDuplicateContainers(containers); err != nil {
		return err
	}
	cMap := make(map[string]customizev1beta1.ResourceRequirements, len(containers)) // cMap is a map of container name to its corresponding resource customization.
	cMapApplied := make(map[string]bool)                                            // cMapApplied is a map of container name to a boolean indicating whether the customization for this container is applied.
	for _, c := range containers {
		cMap[c.Name] = c.Resources
		cMapApplied[c.Name] = false
	}
	var shouldUpdateHPA bool // shouldUpdateHPA is a boolean indicating if we need to update the "minReplicas" field in the HorizontalPodAutoscaler for webhook manager
	// apply customization to the matching controller in the manifest
	for _, item := range m.Items {
		if item.GroupVersionKind() == controllerGVK { // match GVK
			if (!isNamespaced && item.GetName() == controllerName) || // match exact controller name for cluster-scoped controller
				(isNamespaced && strings.HasPrefix(item.GetName(), controllerName)) { // match the prefix for namespace-scoped controller, ignore the namespace ID suffix
				// apply container resource customization for this controller.
				if err := item.MutateContainers(customizeContainerResourcesFn(cMap, cMapApplied)); err != nil {
					return err
				}
				// apply replicas customization for this controller.
				if replicas != nil && controllerName == "cnrm-webhook-manager" { // currently only support customizing replicas for cnrm-webhook-manager.
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
func customizeContainerResourcesFn(cMap map[string]customizev1beta1.ResourceRequirements, cMapApplied map[string]bool) func(container map[string]interface{}) error {
	return func(container map[string]interface{}) error {
		name, _, err := unstructured.NestedString(container, "name")
		if err != nil {
			return fmt.Errorf("error reading container name: %w", err)
		}
		r, found := cMap[name]
		if !found {
			return nil
		}

		// validate the container resource values before applying them to the manifest object.
		if err := validateContainerResourceCustomizationValues(r, container); err != nil {
			return fmt.Errorf("the resources customization for container \"%s\" is invalid: %w", name, err)
		}

		shouldUpdateGOMEMLIMIT := false // we need to update the GOMEMLIMIT environment variable if we update the memory request.
		if r.Limits != nil && !r.Limits.Cpu().IsZero() {
			if err := unstructured.SetNestedField(container, r.Limits.Cpu().String(), "resources", "limits", "cpu"); err != nil {
				return fmt.Errorf("error setting cpu limit: %w", err)
			}
		}
		if r.Limits != nil && !r.Limits.Memory().IsZero() {
			if err := unstructured.SetNestedField(container, r.Limits.Memory().String(), "resources", "limits", "memory"); err != nil {
				return fmt.Errorf("error setting memory limit: %w", err)
			}
		}
		if r.Requests != nil && !r.Requests.Cpu().IsZero() {
			if err := unstructured.SetNestedField(container, r.Requests.Cpu().String(), "resources", "requests", "cpu"); err != nil {
				return fmt.Errorf("error setting cpu request: %w", err)
			}
		}
		if r.Requests != nil && !r.Requests.Memory().IsZero() {
			if err := unstructured.SetNestedField(container, r.Requests.Memory().String(), "resources", "requests", "memory"); err != nil {
				return fmt.Errorf("error setting memory request: %w", err)
			}
			shouldUpdateGOMEMLIMIT = true
		}

		// update the GOMEMLIMIT environment variable if we update the memory request.
		if shouldUpdateGOMEMLIMIT {
			if err := updateContainerEnvIfFound(container, "GOMEMLIMIT", calculateGoMemLimit(r.Requests.Memory().Value())); err != nil {
				return err
			}
		}

		cMapApplied[name] = true
		return nil
	}
}

// calculateGoMemLimit returns 85% of the input requested memory with the correct format.
func calculateGoMemLimit(requestedMemory int64) string {
	goMemLimit := resource.NewQuantity(requestedMemory*17/20, resource.BinarySI) // setting GOMEMLIMIT as 85% of the requested memory.
	goMemLimitFormatted := goMemLimit.String() + "B"                             // adding suffix "B" to accommodate the format supported by GOMEMLIMIT.
	return goMemLimitFormatted
}

// updateContainerEnvIfFound updates the value of the environment variable in the container's environment variable list.
// If the environment variable is not found, the function is no-op.
func updateContainerEnvIfFound(container map[string]interface{}, name, value string) error {
	// retrieve env list
	envs, found, err := unstructured.NestedSlice(container, "env")
	if err != nil {
		return fmt.Errorf("error getting container env list: %w", err)
	}
	if !found { // do not update the value if we cannot find the environment variable.
		return nil
	}

	// update env list
	for _, e := range envs {
		env, ok := e.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to parse container env %v", e)
		}
		envName, ok, err := unstructured.NestedFieldNoCopy(env, "name")
		if err != nil {
			return fmt.Errorf("error getting \"name\" field from container env %v: %w", env, err)
		}
		if ok && envName == name { // found a match
			if err := unstructured.SetNestedField(env, value, "value"); err != nil {
				return fmt.Errorf("error setting container env %s: %w", name, err)
			}
			break
		}
	}

	// write back env list
	if err := unstructured.SetNestedSlice(container, envs, "env"); err != nil {
		return fmt.Errorf("error setting container env list: %w", err)
	}
	return nil
}

func checkForDuplicateContainers(containers []customizev1beta1.ContainerResourceSpec) error {
	var cNames []string
	for _, c := range containers {
		cNames = append(cNames, c.Name)
	}
	duplicates := FindDuplicateStrings(cNames)
	if len(duplicates) > 0 {
		return fmt.Errorf("the following containers are specified multiple times in the Spec: %s", strings.Join(duplicates, ", "))
	}
	return nil
}

func FindDuplicateStrings(strs []string) []string {
	counter := make(map[string]int)
	for _, s := range strs {
		counter[s]++
	}
	var duplicates []string
	for c, cnt := range counter {
		if cnt > 1 {
			duplicates = append(duplicates, c)
		}
	}
	return duplicates
}

func GetValidatingWebhookConfigurationCustomization(ctx context.Context, c client.Client, name string) (*customizev1beta1.ValidatingWebhookConfigurationCustomization, error) {
	obj := &customizev1beta1.ValidatingWebhookConfigurationCustomization{}
	if err := c.Get(ctx, types.NamespacedName{Name: name}, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// ListValidatingWebhookConfigurationCustomizations lists all ValidatingWebhookConfigurationCustomization CRs.
func ListValidatingWebhookConfigurationCustomizations(ctx context.Context, c client.Client) ([]customizev1beta1.ValidatingWebhookConfigurationCustomization, error) {
	list := &customizev1beta1.ValidatingWebhookConfigurationCustomizationList{}
	if err := c.List(ctx, list); err != nil {
		return nil, err
	}
	return list.Items, nil
}

func GetMutatingWebhookConfigurationCustomization(ctx context.Context, c client.Client, name string) (*customizev1beta1.MutatingWebhookConfigurationCustomization, error) {
	obj := &customizev1beta1.MutatingWebhookConfigurationCustomization{}
	if err := c.Get(ctx, types.NamespacedName{Name: name}, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// ListMutatingWebhookConfigurationCustomizations lists all MutatingWebhookConfigurationCustomization CRs.
func ListMutatingWebhookConfigurationCustomizations(ctx context.Context, c client.Client) ([]customizev1beta1.MutatingWebhookConfigurationCustomization, error) {
	list := &customizev1beta1.MutatingWebhookConfigurationCustomizationList{}
	if err := c.List(ctx, list); err != nil {
		return nil, err
	}
	return list.Items, nil
}

// validateContainerResourceCustomizationValues validates the container resource values specified by the user.
// If both `request` and `limit` values are specified by the user, they are checked against each other.
// Otherwise, the specified value is checked against the default value found in the KCC manifest object.
func validateContainerResourceCustomizationValues(r customizev1beta1.ResourceRequirements, container map[string]interface{}) error {
	// 1. validate CPU request and limit.
	cpuLimitIsSet := r.Limits != nil && !r.Limits.Cpu().IsZero()
	cpuRequestIsSet := r.Requests != nil && !r.Requests.Cpu().IsZero()
	if cpuRequestIsSet && cpuLimitIsSet {
		// `request` and `limit` are both set - check against each other.
		if r.Limits.Cpu().Cmp(*r.Requests.Cpu()) < 0 {
			return fmt.Errorf("cpu limit %s is less than cpu request %s", r.Limits.Cpu().String(), r.Requests.Cpu().String())
		}
	} else if cpuRequestIsSet {
		// only `request` is set - check against the default limit value in manifest.
		defaultCPULimitString, found, err := unstructured.NestedString(container, "resources", "limits", "cpu")
		if err != nil {
			return fmt.Errorf("unexpected error when fetching the container cpu limit: %w", err)
		}
		if found {
			defaultCPULimit, err := resource.ParseQuantity(defaultCPULimitString)
			if err != nil {
				return fmt.Errorf("unexpected error when parsing the quantity string %s: %w", defaultCPULimitString, err)
			}
			if defaultCPULimit.Cmp(*r.Requests.Cpu()) < 0 {
				return fmt.Errorf("default cpu limit %s in the manifest is less than the cpu request %s", defaultCPULimitString, r.Requests.Cpu().String())
			}
		}
	} else if cpuLimitIsSet {
		// only `limit` is set - check against the default request value in manifest.
		defaultCPURequestString, found, err := unstructured.NestedString(container, "resources", "requests", "cpu")
		if err != nil {
			return fmt.Errorf("unexpected error when fetching the container cpu request: %w", err)
		}
		if found {
			defaultCPURequest, err := resource.ParseQuantity(defaultCPURequestString)
			if err != nil {
				return fmt.Errorf("unexpected error when parsing the quantity string %s: %w", defaultCPURequestString, err)
			}
			if defaultCPURequest.Cmp(*r.Limits.Cpu()) > 0 {
				return fmt.Errorf("cpu limit %s is less than the default cpu request %s in the manifest", r.Limits.Cpu().String(), defaultCPURequestString)
			}
		}
	}
	// 2. validate memory request and limit.
	memoryLimitIsSet := r.Limits != nil && !r.Limits.Memory().IsZero()
	memoryRequestIsSet := r.Requests != nil && !r.Requests.Memory().IsZero()
	if memoryRequestIsSet && memoryLimitIsSet {
		// `request` and `limit` are both set - check against each other.
		if r.Limits.Memory().Cmp(*r.Requests.Memory()) < 0 {
			return fmt.Errorf("memory limit %s is less than memory request %s", r.Limits.Memory().String(), r.Requests.Memory().String())
		}
	} else if memoryRequestIsSet {
		// only `request` is set - check against the default limit value in manifest.
		defaultMemoryLimitString, found, err := unstructured.NestedString(container, "resources", "limits", "memory")
		if err != nil {
			return fmt.Errorf("unexpected error when fetching the container memory limit: %w", err)
		}
		if found {
			defaultMemoryLimit, err := resource.ParseQuantity(defaultMemoryLimitString)
			if err != nil {
				return fmt.Errorf("unexpected error when parsing the quantity string %s: %w", defaultMemoryLimitString, err)
			}
			if defaultMemoryLimit.Cmp(*r.Requests.Memory()) < 0 {
				return fmt.Errorf("default memory limit %s in the manifest is less than the memory request %s", defaultMemoryLimitString, r.Requests.Memory().String())
			}
		}
	} else if memoryLimitIsSet {
		// only `limit` is set - check against the default request value in manifest.
		defaultMemoryRequestString, found, err := unstructured.NestedString(container, "resources", "requests", "memory")
		if err != nil {
			return fmt.Errorf("unexpected error when fetching the container memory request: %w", err)
		}
		if found {
			defaultMemoryRequest, err := resource.ParseQuantity(defaultMemoryRequestString)
			if err != nil {
				return fmt.Errorf("unexpected error when parsing the quantity string %s: %w", defaultMemoryRequestString, err)
			}
			if defaultMemoryRequest.Cmp(*r.Limits.Memory()) > 0 {
				return fmt.Errorf("memory limit %s is less than the default memory request %s in the manifest", r.Limits.Memory().String(), defaultMemoryRequestString)
			}
		}
	}
	return nil
}

func ApplyContainerRateLimit(m *manifest.Objects, targetControllerName string, ratelimit *customizev1beta1.RateLimit) error {
	if ratelimit == nil {
		return nil
	}

	var (
		targetContainerName string
		targetControllerGVK schema.GroupVersionKind
	)
	switch targetControllerName {
	case "cnrm-controller-manager":
		targetContainerName = "manager"
		targetControllerGVK = schema.GroupVersionKind{
			Group:   appsv1.SchemeGroupVersion.Group,
			Version: appsv1.SchemeGroupVersion.Version,
			Kind:    "StatefulSet",
		}
	default:
		return fmt.Errorf("rate limit customization for %s is not supported. "+
			"Supported controllers: %s",
			targetControllerName, strings.Join(customizev1beta1.ValidRateLimitControllers, ", "))
	}

	count := 0
	for _, item := range m.Items {
		if item.GroupVersionKind() != targetControllerGVK {
			continue
		}
		if !strings.HasPrefix(item.GetName(), targetControllerName) {
			continue
		}
		if err := item.MutateContainers(customizeRateLimitFn(targetContainerName, ratelimit)); err != nil {
			return err
		}
		count++
	}
	if count != 1 {
		return fmt.Errorf("rate limit customization for %s modified %d instances.", targetControllerName, count)
	}
	return nil
}

func customizeRateLimitFn(target string, rateLimit *customizev1beta1.RateLimit) func(container map[string]interface{}) error {
	return func(container map[string]interface{}) error {
		name, _, err := unstructured.NestedString(container, "name")
		if err != nil {
			return fmt.Errorf("error reading container name: %w", err)
		}
		if name != target {
			return nil
		}
		return applyRateLimitToContainerArg(container, rateLimit)
	}
}

func applyRateLimitToContainerArg(container map[string]interface{}, rateLimit *customizev1beta1.RateLimit) error {
	if rateLimit == nil {
		return nil
	}
	origArgs, found, err := unstructured.NestedStringSlice(container, "args")
	if err != nil {
		return fmt.Errorf("error getting args in container: %w", err)
	}
	wantArgs := []string{}
	if rateLimit.QPS > 0 {
		wantArgs = append(wantArgs, fmt.Sprintf("--qps=%d", rateLimit.QPS))
	}
	if rateLimit.Burst > 0 {
		wantArgs = append(wantArgs, fmt.Sprintf("--burst=%d", rateLimit.Burst))
	}
	if found {
		for _, arg := range origArgs {
			if strings.Contains(arg, "--qps") || strings.Contains(arg, "--burst") {
				continue
			}
			wantArgs = append(wantArgs, arg)
		}
	}
	if err := unstructured.SetNestedStringSlice(container, wantArgs, "args"); err != nil {
		return fmt.Errorf("error setting args in container: %w", err)
	}
	return nil
}

func ApplyContainerPprof(m *manifest.Objects, targetControllerName string, pprofConfig *customizev1beta1.PprofConfig) error {
	if pprofConfig == nil {
		return nil
	}

	var (
		targetContainerName string
		targetControllerGVK schema.GroupVersionKind
	)
	switch targetControllerName {
	case "cnrm-controller-manager":
		targetContainerName = "manager"
		targetControllerGVK = schema.GroupVersionKind{
			Group:   appsv1.SchemeGroupVersion.Group,
			Version: appsv1.SchemeGroupVersion.Version,
			Kind:    "StatefulSet",
		}
	default:
		return fmt.Errorf("pprof config customization for %s is not supported. "+
			"Supported controllers: %s",
			targetControllerName, strings.Join(customizev1beta1.SupportedPprofControllers, ", "))
	}

	count := 0
	for _, item := range m.Items {
		if item.GroupVersionKind() != targetControllerGVK {
			continue
		}
		if !strings.HasPrefix(item.GetName(), targetControllerName) {
			continue
		}
		if err := item.MutateContainers(customizePprofConfigFn(targetContainerName, pprofConfig)); err != nil {
			return err
		}
		count++
	}
	if count != 1 {
		return fmt.Errorf("pprof config customization for %s modified %d instances.", targetControllerName, count)
	}
	return nil
}

func customizePprofConfigFn(target string, pprofConfig *customizev1beta1.PprofConfig) func(container map[string]interface{}) error {
	return func(container map[string]interface{}) error {
		name, _, err := unstructured.NestedString(container, "name")
		if err != nil {
			return fmt.Errorf("error reading container name: %w", err)
		}
		if name != target {
			return nil
		}
		return applyPprofConfigToContainerArg(container, pprofConfig)
	}
}

func applyPprofConfigToContainerArg(container map[string]interface{}, pprofConfig *customizev1beta1.PprofConfig) error {
	if pprofConfig == nil {
		return nil
	}
	origArgs, found, err := unstructured.NestedStringSlice(container, "args")
	if err != nil {
		return fmt.Errorf("error getting args in container: %w", err)
	}
	wantArgs := []string{}
	if strings.ToUpper(pprofConfig.Support) == "ALL" {
		wantArgs = append(wantArgs, "--enable-pprof=true")
	} else {
		wantArgs = append(wantArgs, "--enable-pprof=false")
	}
	if pprofConfig.Port > 0 {
		wantArgs = append(wantArgs, fmt.Sprintf("--pprof-port=%d", pprofConfig.Port))
	}
	if found {
		for _, arg := range origArgs {
			if strings.Contains(arg, "--enable-pprof") || strings.Contains(arg, "--pprof-port") {
				// drop the old value on the floor
				continue
			}
			wantArgs = append(wantArgs, arg)
		}
	}
	if err := unstructured.SetNestedStringSlice(container, wantArgs, "args"); err != nil {
		return fmt.Errorf("error setting args in container: %w", err)
	}
	return nil
}

const delimiter = "-"

func ReplaceNamespaceSuffix(namespace, suffix string) string {
	if suffix == "" {
		return namespace
	}

	lastDelimiterIndex := strings.LastIndexAny(namespace, delimiter)

	// If no delimiter is found, there's no suffix to replace.
	// Return the original string.
	if lastDelimiterIndex == -1 {
		return namespace
	}

	return namespace[0:lastDelimiterIndex+1] + suffix
}
