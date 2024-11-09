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

func sortPrincipalsInSpec(spec *krm.PrivilegedAccessManagerEntitlementSpec) {
	if spec == nil {
		return
	}

	sortPrincipals(spec.EligibleUsers)
	if spec.ApprovalWorkflow == nil || spec.ApprovalWorkflow.ManualApprovals == nil {
		return
	}
	for _, step := range spec.ApprovalWorkflow.ManualApprovals.Steps {
		sortPrincipals(step.Approvers)
	}
}

func sortPrincipals(accessControlEntries []krm.AccessControlEntry) {
	for _, eu := range accessControlEntries {
		sort.Strings(eu.Principals)
	}
}
