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

package testiam

import (
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"

	"github.com/google/go-cmp/cmp"
)

type conditionKey struct {
	Description string
	Expression  string
	Title       string
}

type bindingKey struct {
	Role      string
	Condition conditionKey
}

func SameBindings(a, b []v1beta1.IAMPolicyBinding) bool {
	return reflect.DeepEqual(bindingsMap(a), bindingsMap(b))
}

// If slice a contains all bindings in slice b, return true.
func ContainsBindings(a, b []v1beta1.IAMPolicyBinding) bool {
	bindingMapA := bindingsMap(a)
	bindingMapB := bindingsMap(b)
	for k, v := range bindingMapB {
		if _, ok := bindingMapA[k]; !ok {
			return false
		}
		for m := range v {
			if _, ok := bindingMapA[k][m]; !ok {
				return false
			}
		}
	}
	return true
}

func SameAuditConfigs(a, b []v1beta1.IAMPolicyAuditConfig) bool {
	return reflect.DeepEqual(auditConfigsMap(a), auditConfigsMap(b))
}

func SameAuditLogConfigs(a, b []v1beta1.AuditLogConfig) bool {
	return reflect.DeepEqual(auditLogConfigsMap(a), auditLogConfigsMap(b))
}

func bindingsMap(bindings []v1beta1.IAMPolicyBinding) map[bindingKey]map[v1beta1.Member]bool {
	bindingKeyToMembers := make(map[bindingKey]map[v1beta1.Member]bool)
	for _, b := range bindings {
		key := bindingKeyFromBinding(b)
		if _, ok := bindingKeyToMembers[key]; !ok {
			bindingKeyToMembers[key] = make(map[v1beta1.Member]bool)
		}
		members := bindingKeyToMembers[key]
		for _, m := range b.Members {
			members[m] = true
		}
	}
	return bindingKeyToMembers
}

func bindingKeyFromBinding(b v1beta1.IAMPolicyBinding) bindingKey {
	return bindingKey{
		Role:      b.Role,
		Condition: conditionKeyFromCondition(b.Condition),
	}
}

func conditionKeyFromCondition(c *v1beta1.IAMCondition) conditionKey {
	if c == nil {
		return conditionKey{}
	}
	return conditionKey{
		Description: c.Description,
		Expression:  c.Expression,
		Title:       c.Title,
	}
}

func auditConfigsMap(auditConfigs []v1beta1.IAMPolicyAuditConfig) map[string]map[string]map[v1beta1.Member]bool {
	serviceToAuditLogConfigsMap := make(map[string]map[string]map[v1beta1.Member]bool)
	for _, a := range auditConfigs {
		serviceToAuditLogConfigsMap[a.Service] = auditLogConfigsMap(a.AuditLogConfigs)
	}
	return serviceToAuditLogConfigsMap
}

func auditLogConfigsMap(auditLogConfigs []v1beta1.AuditLogConfig) map[string]map[v1beta1.Member]bool {
	logTypeToMembers := make(map[string]map[v1beta1.Member]bool)
	for _, a := range auditLogConfigs {
		if _, ok := logTypeToMembers[a.LogType]; !ok {
			logTypeToMembers[a.LogType] = make(map[v1beta1.Member]bool)
		}
		members := logTypeToMembers[a.LogType]
		for _, m := range a.ExemptedMembers {
			members[m] = true
		}
	}
	return logTypeToMembers
}

func AssertSamePolicy(t *testing.T, k8sPolicy, gcpPolicy *v1beta1.IAMPolicy) {
	if !reflect.DeepEqual(k8sPolicy.Spec.ResourceReference, gcpPolicy.Spec.ResourceReference) {
		diff := cmp.Diff(k8sPolicy.Spec.ResourceReference, gcpPolicy.Spec.ResourceReference)
		t.Fatalf("GCP policy has incorrect resource reference. Diff (-want, +got):\n%v", diff)
	}
	if !SameBindings(k8sPolicy.Spec.Bindings, gcpPolicy.Spec.Bindings) {
		t.Fatalf("GCP policy has incorrect bindings; got: %v, want: %v", gcpPolicy.Spec.Bindings, k8sPolicy.Spec.Bindings)
	}
	if !SameAuditConfigs(k8sPolicy.Spec.AuditConfigs, gcpPolicy.Spec.AuditConfigs) {
		t.Fatalf("GCP policy has incorrect audit configs; got: %v, want: %v", gcpPolicy.Spec.AuditConfigs, k8sPolicy.Spec.AuditConfigs)
	}
}
