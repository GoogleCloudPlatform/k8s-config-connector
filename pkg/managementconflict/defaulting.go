// Copyright 2025 Google LLC
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

package managementconflict

import (
	"context"
	"fmt"

	"github.com/nasa9084/go-openapi"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leasable"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
	value, ok := obj.GetAnnotations()[FullyQualifiedAnnotation]
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

	k8s.SetAnnotation(FullyQualifiedAnnotation, string(policy), obj)
	return nil
}

func getDefaultManagementConflictPreventAnnotationForNamespace(ns *corev1.Namespace, supportLeasing bool) (ManagementConflictPreventionPolicy,
	error) {
	value, ok := ns.GetAnnotations()[FullyQualifiedAnnotation]
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
