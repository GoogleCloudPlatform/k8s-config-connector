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

package v1beta1

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

// IAMPolicySpecDiffers compares two IAMPolicySpec objects and a diff.
func IAMPolicySpecDiffers(desired, actual *IAMPolicySpec) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}

	// Compare ResourceReference, which is immutable.
	if desired.ResourceReference.Kind != actual.ResourceReference.Kind {
		diff.AddField("spec.resourceRef.kind", desired.ResourceReference.Kind, actual.ResourceReference.Kind)
	}
	if desired.ResourceReference.Namespace != actual.ResourceReference.Namespace {
		diff.AddField("spec.resourceRef.namespace", desired.ResourceReference.Namespace, actual.ResourceReference.Namespace)
	}
	if desired.ResourceReference.Name != actual.ResourceReference.Name {
		diff.AddField("spec.resourceRef.name", desired.ResourceReference.Name, actual.ResourceReference.Name)
	}
	if desired.ResourceReference.External != actual.ResourceReference.External {
		diff.AddField("spec.resourceRef.external", desired.ResourceReference.External, actual.ResourceReference.External)
	}

	bindingsDiff := compareBindings(desired.Bindings, actual.Bindings)
	auditConfigDiffs := compareAuditConfigs(desired.AuditConfigs, actual.AuditConfigs)

	diff.AddDiff(bindingsDiff)
	diff.AddDiff(auditConfigDiffs)

	return diff
}

// compareBindings compares two slices of IAMPolicyBinding.
// It treats the slices as maps keyed by Role, and members are treated as unordered sets.
func compareBindings(desired, actual []IAMPolicyBinding) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}
	desiredMap := make(map[string]IAMPolicyBinding)
	for _, b := range desired {
		sort.Slice(b.Members, func(i, j int) bool {
			return b.Members[i] < b.Members[j]
		})
		// Canonicalize members for stable comparison.
		desiredMap[b.Role] = b
	}

	actualMap := make(map[string]IAMPolicyBinding)
	for _, b := range actual {
		sort.Slice(b.Members, func(i, j int) bool {
			return b.Members[i] < b.Members[j]
		})
		// Canonicalize members for stable comparison.
		actualMap[b.Role] = b
	}

	// Check for added or modified bindings.
	for role, desiredBinding := range desiredMap {
		actualBinding, ok := actualMap[role]
		if !ok {
			diff.AddField(fmt.Sprintf("spec.bindings[role=%v]", role), "present in desired spec", "absent in actual spec")
			continue
		}
		if !reflect.DeepEqual(desiredBinding, actualBinding) {
			diff.AddField(fmt.Sprintf("spec.bindings[role=%v]", role), desiredBinding, actualBinding)
		}
	}

	// Check for removed bindings.
	for role := range actualMap {
		if _, ok := desiredMap[role]; !ok {
			diff.AddField(fmt.Sprintf("spec.bindings[role=%v]", role), "absent in desired spec", "present in actual spec")
		}
	}

	return diff
}

// compareAuditConfigs compares two slices of IAMPolicyAuditConfig.
// It treats the slices as maps keyed by Service and canonicalizes inner slices for stable comparison.
func compareAuditConfigs(desired, actual []IAMPolicyAuditConfig) *structuredreporting.Diff {
	diff := &structuredreporting.Diff{}

	// Helper function to canonicalize the inner slices for a stable comparison.
	canonicalize := func(configs []IAMPolicyAuditConfig) []IAMPolicyAuditConfig {
		// Create a deep copy to avoid modifying the original struct.
		copiedConfigs := make([]IAMPolicyAuditConfig, len(configs))
		copy(copiedConfigs, configs)

		for _, ac := range configs {
			copiedConfigs = append(copiedConfigs, *ac.DeepCopy())
		}

		for i := range copiedConfigs {
			for j := range copiedConfigs[i].AuditLogConfigs {
				sort.Slice(copiedConfigs[i].AuditLogConfigs[j].ExemptedMembers, func(a, b int) bool {
					return copiedConfigs[i].AuditLogConfigs[j].ExemptedMembers[a] < copiedConfigs[i].AuditLogConfigs[j].ExemptedMembers[b]
				})

			}
			sort.Slice(copiedConfigs[i].AuditLogConfigs, func(j, k int) bool {
				return copiedConfigs[i].AuditLogConfigs[j].LogType < copiedConfigs[i].AuditLogConfigs[k].LogType
			})
		}

		return copiedConfigs
	}

	canonicalizedDesired := canonicalize(desired)
	canonicalizedActual := canonicalize(actual)

	desiredMap := make(map[string]IAMPolicyAuditConfig)
	for _, ac := range canonicalizedDesired {
		desiredMap[ac.Service] = ac
	}
	actualMap := make(map[string]IAMPolicyAuditConfig)
	for _, ac := range canonicalizedActual {
		actualMap[ac.Service] = ac
	}

	// Check for added or modified audit configs.
	for service, desiredConfig := range desiredMap {
		actualConfig, ok := actualMap[service]
		if !ok {
			diff.AddField(fmt.Sprintf("spec.auditConfigs[service=%v]", service), "present in desired spec", "absent in actual spec")
			continue
		}
		if !reflect.DeepEqual(desiredConfig, actualConfig) {
			diff.AddField(fmt.Sprintf("spec.auditConfigs[service=%v]", service), desiredConfig, actualConfig)
		}
	}

	// Check for removed audit configs.
	for service := range actualMap {
		if _, ok := desiredMap[service]; !ok {
			diff.AddField(fmt.Sprintf("spec.auditConfigs[service=%v]", service), "absent in desired spec", "present in actual spec")
		}
	}

	return diff
}
