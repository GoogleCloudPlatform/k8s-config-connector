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

package k8s

import (
	"context"
	"fmt"
	"sort"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leasable"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nasa9084/go-openapi"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetNamespacedName(obj metav1.Object) types.NamespacedName {
	return types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}
}

func IsManagedByKCC(gvk schema.GroupVersionKind) bool {
	// KCC controllers don't manage CRDs under 'core.cnrm.cloud.google.com'.
	// These CRDs are managed by KCC operator.
	if strings.HasSuffix(gvk.Group, CoreCNRMGroup) {
		return false
	}
	return strings.HasSuffix(gvk.Group, CNRMGroup)
}

func IsDeleted(objectMeta *metav1.ObjectMeta) bool {
	return !objectMeta.DeletionTimestamp.IsZero()
}

func HasAbandonAnnotation(obj metav1.Object) bool {
	val, ok := GetAnnotation(DeletionPolicyAnnotation, obj)
	return ok && val == DeletionPolicyAbandon
}

func GVKListContains(gvkList []schema.GroupVersionKind, gvk schema.GroupVersionKind) bool {
	for _, v := range gvkList {
		if v == gvk {
			return true
		}
	}
	return false
}

func GVKSetToList(gvkSet map[schema.GroupVersionKind]bool) []schema.GroupVersionKind {
	gvkList := make([]schema.GroupVersionKind, 0, len(gvkSet))
	for gvk := range gvkSet {
		gvkList = append(gvkList, gvk)
	}
	return gvkList
}

func SortGVKsByKind(gvks []schema.GroupVersionKind) []schema.GroupVersionKind {
	gvksCopy := append(make([]schema.GroupVersionKind, 0), gvks...)
	sort.Slice(gvksCopy, func(i, j int) bool {
		return gvksCopy[i].Kind < gvksCopy[j].Kind
	})
	return gvksCopy
}

// ToGVR returns the equivalent GVR for a given GVK. Note that while GVKs and
// GVRs do not necessarily have a 1:1 mapping, GVKs and GVRs of CRDs do.
// (see https://book.kubebuilder.io/cronjob-tutorial/gvks.html#kinds-and-resources)
func ToGVR(gvk schema.GroupVersionKind) schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    gvk.Group,
		Version:  gvk.Version,
		Resource: text.Pluralize(strings.ToLower(gvk.Kind)),
	}
}

func GetProjectIDForNamespace(ctx context.Context, c client.Client, namespaceName string) (string, error) {
	var ns corev1.Namespace
	if err := c.Get(ctx, types.NamespacedName{Name: namespaceName}, &ns); err != nil {
		return "", fmt.Errorf("error getting namespace '%v': %w", namespaceName, err)
	}
	if val, ok := GetAnnotation(ProjectIDAnnotation, &ns); ok {
		return val, nil
	}
	return namespaceName, nil
}

func GetAnnotation(annotation string, obj metav1.Object) (string, bool) {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		return "", false
	}
	val, ok := annotations[annotation]
	return val, ok
}

func SetAnnotation(annotation, val string, obj metav1.Object) {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[annotation] = val
	obj.SetAnnotations(annotations)
}

func RemoveAnnotation(annotation string, obj metav1.Object) {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		return
	}
	delete(annotations, annotation)
	obj.SetAnnotations(annotations)
}

func GetManagementConflictPreventionAnnotationValue(obj metav1.Object) (ManagementConflictPreventionPolicy, error) {
	if val, ok := GetAnnotation(ManagementConflictPreventionPolicyFullyQualifiedAnnotation, obj); ok {
		return valueToManagementConflictPreventionPolicy(val)
	}
	return ManagementConflictPreventionPolicyNone,
		fmt.Errorf("attempted to get value for annotation %v, but annotation was not found", ManagementConflictPreventionPolicyFullyQualifiedAnnotation)
}

func EnsureManagementConflictPreventionAnnotationForTFBasedResource(ctx context.Context, c client.Client, obj metav1.Object, rc *corekccv1alpha1.ResourceConfig, tfResourceMap map[string]*tfschema.Resource) error {
	ns := corev1.Namespace{}
	if err := c.Get(ctx, types.NamespacedName{Name: obj.GetNamespace()}, &ns); err != nil {
		return fmt.Errorf("error getting namespace %v: %w", obj.GetNamespace(), err)
	}
	return ValidateOrDefaultManagementConflictPreventionAnnotationForTFBasedResource(obj, &ns, rc, tfResourceMap)
}

func ValidateOrDefaultManagementConflictPreventionAnnotationForTFBasedResource(obj metav1.Object, ns *corev1.Namespace, rc *corekccv1alpha1.ResourceConfig, tfResourceMap map[string]*tfschema.Resource) error {
	supportsLeasing, err := leasable.ResourceConfigSupportsLeasing(rc, tfResourceMap)
	if err != nil {
		return err
	}
	return validateOrDefaultManagementConflictPreventionAnnotation(obj, ns, supportsLeasing)
}

func ValidateOrDefaultManagementConflictPreventionAnnotationForDCLBasedResource(obj metav1.Object, ns *corev1.Namespace, schema *openapi.Schema) error {
	supportsLeasing, err := leasable.DCLSchemaSupportsLeasing(schema)
	if err != nil {
		return err
	}
	return validateOrDefaultManagementConflictPreventionAnnotation(obj, ns, supportsLeasing)
}

func validateOrDefaultManagementConflictPreventionAnnotation(obj metav1.Object, ns *corev1.Namespace, supportsLeasing bool) error {
	value, ok := GetAnnotation(ManagementConflictPreventionPolicyFullyQualifiedAnnotation, obj)
	if ok {
		// the value is supplied by the customer so ensure it is valid
		policy, err := valueToManagementConflictPreventionPolicy(value)
		if err != nil {
			return err
		}
		return validateManagementConflictPolicyForResource(policy, supportsLeasing)
	}
	policy, err := getDefaultManagementConflictPreventAnnotationForNamespace(ns, supportsLeasing)
	if err != nil {
		return err
	}
	SetAnnotation(ManagementConflictPreventionPolicyFullyQualifiedAnnotation, string(policy), obj)
	return nil
}

func getDefaultManagementConflictPreventAnnotationForNamespace(ns *corev1.Namespace, supportLeasing bool) (ManagementConflictPreventionPolicy,
	error) {
	value, ok := GetAnnotation(ManagementConflictPreventionPolicyFullyQualifiedAnnotation, ns)
	if ok {
		policy, err := valueToManagementConflictPreventionPolicy(value)
		if err != nil {
			return ManagementConflictPreventionPolicyNone, fmt.Errorf("unable to use default management conflict policy for namespace: %w", err)
		}
		if !isManagementConflictPolicyValidForResource(policy, supportLeasing) {
			return ManagementConflictPreventionPolicyNone, nil
		}
		return policy, nil
	}
	// if there is no value on the namespace return the default
	return getDefaultManagementConflictPolicyForResource(), nil
}

func isManagementConflictPolicyValidForResource(policy ManagementConflictPreventionPolicy, supportLeasing bool) bool {
	switch policy {
	case ManagementConflictPreventionPolicyNone:
		return true
	case ManagementConflictPreventionPolicyResource:
		return supportLeasing
	default:
		return false
	}
}

// getDefaultManagementConflictPolicyForResource returns the default policy for a resource.
//
// This value was set to default to None, due to user complaints that label leasing behavior results
// in resources sporadically setting not Ready, and causing issues for kpt live apply for a large
// amount of resources.
//
// Before the default is flipped again, the label leaser should no longer flip the Ready state to false
// and mark the resource as updating (https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/387)
func getDefaultManagementConflictPolicyForResource() ManagementConflictPreventionPolicy {
	return ManagementConflictPreventionPolicyNone
}

func validateManagementConflictPolicyForResource(policy ManagementConflictPreventionPolicy, supportLeasing bool) error {
	switch policy {
	case ManagementConflictPreventionPolicyNone:
		return nil
	case ManagementConflictPreventionPolicyResource:
		if !supportLeasing {
			return fmt.Errorf("the resource kind does not support usage of %v of '%v'",
				ManagementConflictPreventionPolicyAnnotation, policy)
		}
		return nil
	default:
		return fmt.Errorf("unknown management conflict policy: %v", policy)
	}
}

func valueToManagementConflictPreventionPolicy(value string) (ManagementConflictPreventionPolicy, error) {
	for _, policy := range ManagementConflictPreventionPolicyValues {
		if value == string(policy) {
			return ManagementConflictPreventionPolicy(value), nil
		}
	}
	return ManagementConflictPreventionPolicyNone, fmt.Errorf("invalid management conflict policy '%v', can be one of {%v}",
		value, strings.Join(ManagementConflictPreventionPolicyValues, ", "))
}

func SetDefaultContainerAnnotation(obj metav1.Object, ns *corev1.Namespace, containers []corekccv1alpha1.Container) error {
	if len(containers) == 0 {
		// No defaulting required
		return nil
	}
	// If the resource already has a container annotation, no modification is required
	val, _, err := GetContainerAnnotation(obj.GetAnnotations(), ContainerTypes(containers))
	if err != nil {
		return fmt.Errorf("error getting container annotation from object: %w", err)
	}
	if val != "" {
		return nil
	}
	// if the Namespace has a container annotation, we'll use that as the default
	val, containerType, err := GetContainerAnnotation(ns.GetAnnotations(), ContainerTypes(containers))
	if err != nil {
		return fmt.Errorf("error getting container annotation from object: %w", err)
	}
	if val != "" {
		SetAnnotation(GetAnnotationForContainerType(containerType), val, obj)
		return nil
	}
	// For project-scoped resources we can use the namespace name as the project ID
	if IsProjectScoped(containers) {
		SetAnnotation(ProjectIDAnnotation, ns.GetName(), obj)
		return nil
	}
	possibleAnnotations := containerTypesToAnnotations(ContainerTypes(containers))
	return fmt.Errorf("neither resource nor namespace have the required container object annotation, one of: [%v]", strings.Join(possibleAnnotations, ", "))
}

// GetContainerAnnotation will get the appropriate container annotation from the given
// annotations.
func GetContainerAnnotation(annotations map[string]string, containerTypes []corekccv1alpha1.ContainerType) (string, corekccv1alpha1.ContainerType, error) {
	var containerVal string
	var containerType corekccv1alpha1.ContainerType
	var found bool
	for _, c := range containerTypes {
		val, ok := annotations[GetAnnotationForContainerType(c)]
		if !ok {
			continue
		}
		if found {
			return "", "", fmt.Errorf("ambiguious container annotation: found for %v and %v", containerType, c)
		}
		containerVal = val
		containerType = c
		found = true
	}
	return containerVal, containerType, nil
}

func IsProjectScoped(containers []corekccv1alpha1.Container) bool {
	for _, c := range containers {
		if c.Type == corekccv1alpha1.ContainerTypeProject {
			return true
		}
	}
	return false
}

func GetAnnotationForContainerType(containerType corekccv1alpha1.ContainerType) string {
	switch containerType {
	case corekccv1alpha1.ContainerTypeProject:
		return ProjectIDAnnotation
	case corekccv1alpha1.ContainerTypeFolder:
		return FolderIDAnnotation
	case corekccv1alpha1.ContainerTypeOrganization:
		return OrgIDAnnotation
	default:
		panic(fmt.Errorf("unrecognized container type %v", containerType))
	}
}

func containerTypesToAnnotations(containerTypes []corekccv1alpha1.ContainerType) []string {
	annotations := make([]string, 0)
	for _, c := range containerTypes {
		annotations = append(annotations, GetAnnotationForContainerType(c))
	}
	return annotations
}

func ContainerTypes(containers []corekccv1alpha1.Container) []corekccv1alpha1.ContainerType {
	types := make([]corekccv1alpha1.ContainerType, 0)
	for _, c := range containers {
		types = append(types, c.Type)
	}
	return types
}

// TriggerManagedFieldsMetadata ensures that managed fields metadata is present on the given
// resource for Server-Side Apply (SSA) compatible clusters.
func TriggerManagedFieldsMetadata(ctx context.Context, c client.Client, u *unstructured.Unstructured) (
	*unstructured.Unstructured, error) {
	if len(u.GetManagedFields()) > 0 {
		// Managed fields metadata is present already; no action necessary.
		return u, nil
	}
	// Attempt an SSA patch to trigger the initial SSA metadata on the resource. Construct an
	// unstructured object that only specified the information we care about: a temporary SSA
	// annotation in the annotations map.
	patchSkeleton := &unstructured.Unstructured{}
	patchSkeleton.SetGroupVersionKind(u.GroupVersionKind())
	patchSkeleton.SetName(u.GetName())
	patchSkeleton.SetNamespace(u.GetNamespace())

	patchU := patchSkeleton.DeepCopy()
	patchU.SetAnnotations(map[string]string{SupportsSSAAnnotation: "true"})
	if err := c.Patch(ctx, patchU, client.Apply, client.FieldOwner(SupportsSSAManager)); err != nil {
		if strings.Contains(err.Error(), string(types.MergePatchType)) {
			// The patch was rejected due to the API server not supporting the Apply patch type.
			// No action required.
			return u, nil
		}
		return nil, fmt.Errorf("error patching SSA metadata annotation: %w", err)
	}
	// Now that the SSA metadata has been triggered, remove the annotation. The SSA metadata
	// will persist.
	patchU = patchSkeleton.DeepCopy()
	if err := c.Patch(ctx, patchU, client.Apply, client.FieldOwner(SupportsSSAManager)); err != nil {
		return nil, fmt.Errorf("error removing SSA metadata annotation: %w", err)
	}
	return patchU, nil
}

// KindWithoutServicePrefix returns the kind without the
// service prefix (e.g. "ComputeBackendBucket => "BackendBucket").
// Kinds which do not contain a service prefix are returned directly
// (e.g.  "Project" => "Project").
func KindWithoutServicePrefix(gvk schema.GroupVersionKind) string {
	switch gvk.Kind {
	case "Project", "Folder", "Organization", "BillingAccount":
		// Some kinds do not contain a service prefix.
		return gvk.Kind
	default:
		serviceInLowerCase := strings.TrimSuffix(gvk.Group, "."+CNRMGroup)
		kindInLowerCase := strings.ToLower(gvk.Kind)
		if !strings.HasPrefix(kindInLowerCase, serviceInLowerCase) {
			panic(fmt.Errorf("kind %v unexpectedly does not begin with its service name", gvk.Kind))
		}
		return gvk.Kind[len(serviceInLowerCase):]
	}
}
