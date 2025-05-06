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
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ManagementConflictPreventionPolicy string

const (
	// Management conflict prevention policies
	ManagementConflictPreventionPolicyNone     = "none"
	ManagementConflictPreventionPolicyResource = "resource"

	// ShortName                                                  = "management-conflict-prevention-policy"
	FullyQualifiedAnnotation = "cnrm.cloud.google.com/management-conflict-prevention-policy"
)

var managementConflictPreventionPolicyValues = []string{
	ManagementConflictPreventionPolicyNone,
	ManagementConflictPreventionPolicyResource,
}

func validateManagementConflictPolicyForResource(policy ManagementConflictPreventionPolicy, supportLeasing bool) error {
	switch policy {
	case ManagementConflictPreventionPolicyNone:
		return nil
	case ManagementConflictPreventionPolicyResource:
		if !supportLeasing {
			return fmt.Errorf("the resource kind does not support usage of %v='%v'",
				FullyQualifiedAnnotation, policy)
		}
		return nil
	default:
		return fmt.Errorf("unknown management conflict policy: %v", policy)
	}
}

func valueToManagementConflictPreventionPolicy(value string) (ManagementConflictPreventionPolicy, error) {
	for _, policy := range managementConflictPreventionPolicyValues {
		if value == string(policy) {
			return ManagementConflictPreventionPolicy(value), nil
		}
	}
	return ManagementConflictPreventionPolicyNone, fmt.Errorf("invalid management conflict policy '%v', can be one of {%v}",
		value, strings.Join(managementConflictPreventionPolicyValues, ", "))
}

func GetManagementConflictPreventionAnnotationValue(obj metav1.Object) (ManagementConflictPreventionPolicy, error) {
	if val, ok := obj.GetAnnotations()[FullyQualifiedAnnotation]; ok {
		return valueToManagementConflictPreventionPolicy(val)
	}
	return ManagementConflictPreventionPolicyNone,
		fmt.Errorf("attempted to get value for annotation %v, but annotation was not found", FullyQualifiedAnnotation)
}
