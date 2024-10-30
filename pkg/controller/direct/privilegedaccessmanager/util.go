// Copyright 2024 Google LLC
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

package privilegedaccessmanager

import (
	"sort"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privilegedaccessmanager/v1alpha1"
)

func sortArrayFieldsInSpec(spec *krm.PrivilegedAccessManagerEntitlementSpec) {
	sortAccessControlEntrySlice(spec.EligibleUsers)
	for _, step := range spec.ApprovalWorkflow.ManualApprovals.Steps {
		sortAccessControlEntrySlice(step.Approvers)
	}
}

func sortAccessControlEntrySlice(accessControlEntries []krm.AccessControlEntry) {
	for _, eu := range accessControlEntries {
		sort.Strings(eu.Principals)
	}
	sort.Slice(accessControlEntries, func(i, j int) bool {
		lenA := len(accessControlEntries[i].Principals)
		lenB := len(accessControlEntries[j].Principals)
		if lenA != lenB || lenA == 0 {
			return lenA < lenB
		}
		for x := 0; x < lenA; x++ {
			pA := accessControlEntries[i].Principals[x]
			pB := accessControlEntries[j].Principals[x]
			if pA != pB {
				return pA < pB
			}
		}
		return lenA < lenB
	})
}
