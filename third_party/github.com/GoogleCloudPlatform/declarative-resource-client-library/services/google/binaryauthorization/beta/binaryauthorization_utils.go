// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package binaryauthorization provides types and functiosn for handling binaryauthorization GCP resources.
package beta

import (
	"strings"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

// Returns a copy of the given map with the string "spiffe://" removed from its keys.
func withoutSpiffe(isiar map[string]PolicyIstioServiceIdentityAdmissionRules) map[string]PolicyIstioServiceIdentityAdmissionRules {
	resultISIAR := make(map[string]PolicyIstioServiceIdentityAdmissionRules, len(isiar))
	for k, v := range isiar {
		resultISIAR[strings.TrimPrefix(k, "spiffe://")] = v
	}
	return resultISIAR
}

func equalsPolicyISIAR(m, n map[string]PolicyIstioServiceIdentityAdmissionRules) bool {
	if m == nil && n == nil {
		return true
	}
	m = withoutSpiffe(m)
	n = withoutSpiffe(n)
	ds, err := dcl.Diff(m, n, dcl.DiffInfo{OperationSelector: dcl.TriggersOperation("updatePolicyUpdatePolicyOperation")}, dcl.FieldName{})
	return len(ds) == 0 && err == nil
}

// Compares two values of istioServiceIdentity
func canonicalizePolicyISIAR(m, n interface{}) bool {
	mVal, _ := m.(map[string]PolicyIstioServiceIdentityAdmissionRules)
	nVal, _ := n.(map[string]PolicyIstioServiceIdentityAdmissionRules)
	return equalsPolicyISIAR(mVal, nVal)
}
