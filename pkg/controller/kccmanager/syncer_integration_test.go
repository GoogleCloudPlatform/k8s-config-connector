// Copyright 2026 Google LLC
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

package kccmanager

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/types"
)

func TestGetDesiredSyncers(t *testing.T) {
	leaseNN := types.NamespacedName{Name: "my-lease", Namespace: "cnrm-system"}

	tests := []struct {
		name            string
		watchNamespaces []string
		expected        []types.NamespacedName
	}{
		{
			name:            "cluster scoped",
			watchNamespaces: nil,
			expected: []types.NamespacedName{
				{Name: "my-lease", Namespace: "cnrm-system"},
			},
		},
		{
			name:            "namespaced scoped - one namespace",
			watchNamespaces: []string{"tenant-a"},
			expected: []types.NamespacedName{
				{Name: "my-lease-tenant-a", Namespace: "cnrm-system"},
			},
		},
		{
			name:            "namespaced scoped - multiple namespaces",
			watchNamespaces: []string{"tenant-a", "tenant-b"},
			expected: []types.NamespacedName{
				{Name: "my-lease-tenant-a", Namespace: "cnrm-system"},
				{Name: "my-lease-tenant-b", Namespace: "cnrm-system"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			si := &SyncerIntegration{
				leaseNN:         leaseNN,
				watchNamespaces: tc.watchNamespaces,
			}
			actual := si.getDesiredSyncers()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestGetNamespacesForSyncer(t *testing.T) {
	leaseNN := types.NamespacedName{Name: "my-lease", Namespace: "cnrm-system"}

	si := &SyncerIntegration{
		leaseNN:         leaseNN,
		watchNamespaces: []string{"tenant-a", "tenant-b"},
	}

	ns := si.getNamespacesForSyncer("my-lease-tenant-a")
	if len(ns) != 1 || ns[0] != "tenant-a" {
		t.Fatalf("expected [tenant-a], got %v", ns)
	}

	ns = si.getNamespacesForSyncer("my-lease-unknown")
	if ns != nil {
		t.Fatalf("expected nil, got %v", ns)
	}

	// Cluster scoped
	siCluster := &SyncerIntegration{
		leaseNN:         leaseNN,
		watchNamespaces: nil,
	}
	ns = siCluster.getNamespacesForSyncer("my-lease")
	if ns != nil {
		t.Fatalf("expected nil for cluster scoped, got %v", ns)
	}
}

func TestGetDesiredRules(t *testing.T) {
	si := &SyncerIntegration{
		replicationMode: "Status",
	}

	rules := si.getDesiredRules([]string{"tenant-a"})
	if len(rules) != len(syncerGVKsWithServiceGeneratedIDs)+1 {
		t.Fatalf("unexpected number of rules: %d", len(rules))
	}

	hasGlob := false
	for _, r := range rules {
		ruleMap := r.(map[string]interface{})
		if ruleMap["group"] == "*.cnrm.cloud.google.com" {
			hasGlob = true
			nsList := ruleMap["namespaces"].([]interface{})
			if len(nsList) != 1 || nsList[0] != "tenant-a" {
				t.Errorf("expected namespaces [tenant-a], got %v", nsList)
			}
		}
	}
	if !hasGlob {
		t.Errorf("glob rule missing")
	}

	// Cluster scoped
	siCluster := &SyncerIntegration{
		replicationMode: "Full",
	}
	clusterRules := siCluster.getDesiredRules(nil)
	if len(clusterRules) != 1 {
		t.Fatalf("expected 1 rule for full mode")
	}
	ruleMap := clusterRules[0].(map[string]interface{})
	if _, ok := ruleMap["namespaces"]; ok {
		t.Errorf("did not expect namespaces field in cluster scoped rules")
	}
}

func TestRulesMatch(t *testing.T) {
	rulesA := []interface{}{
		map[string]interface{}{
			"group":      "example.com",
			"version":    "v1",
			"kind":       "Foo",
			"syncFields": []interface{}{"status", "spec.resourceID"},
			"namespaces": []interface{}{"tenant-a"},
		},
	}

	rulesB := []interface{}{
		map[string]interface{}{
			"group":      "example.com",
			"version":    "v1",
			"kind":       "Foo",
			"syncFields": []interface{}{"spec.resourceID", "status"},
			"namespaces": []interface{}{"tenant-a"},
		},
	}

	rulesC := []interface{}{
		map[string]interface{}{
			"group":      "example.com",
			"version":    "v1",
			"kind":       "Foo",
			"syncFields": []interface{}{"status", "spec.resourceID"},
			"namespaces": []interface{}{"tenant-b"},
		},
	}

	if !rulesMatch(rulesA, rulesB) {
		t.Errorf("expected rulesA and rulesB to match (order independent)")
	}

	if rulesMatch(rulesA, rulesC) {
		t.Errorf("expected rulesA and rulesC NOT to match (different namespaces)")
	}

	// Test duplicate GVKs
	rulesDupA := []interface{}{
		map[string]interface{}{
			"group":      "example.com",
			"version":    "v1",
			"kind":       "Foo",
			"syncFields": []interface{}{"status"},
			"namespaces": []interface{}{"tenant-a"},
		},
		map[string]interface{}{
			"group":      "example.com",
			"version":    "v1",
			"kind":       "Foo",
			"syncFields": []interface{}{"spec"},
			"namespaces": []interface{}{"tenant-b"},
		},
	}

	rulesDupB := []interface{}{
		map[string]interface{}{
			"group":      "example.com",
			"version":    "v1",
			"kind":       "Foo",
			"syncFields": []interface{}{"spec"},
			"namespaces": []interface{}{"tenant-b"},
		},
		map[string]interface{}{
			"group":      "example.com",
			"version":    "v1",
			"kind":       "Foo",
			"syncFields": []interface{}{"status"},
			"namespaces": []interface{}{"tenant-a"},
		},
	}

	rulesDupMismatch := []interface{}{
		map[string]interface{}{
			"group":      "example.com",
			"version":    "v1",
			"kind":       "Foo",
			"syncFields": []interface{}{"status"},
			"namespaces": []interface{}{"tenant-a"},
		},
		map[string]interface{}{
			"group":      "example.com",
			"version":    "v1",
			"kind":       "Foo",
			"syncFields": []interface{}{"status"}, // Duplicate of the first one, meaning spec/tenant-b is missing
			"namespaces": []interface{}{"tenant-a"},
		},
	}

	if !rulesMatch(rulesDupA, rulesDupB) {
		t.Errorf("expected rulesDupA and rulesDupB to match (duplicate GVKs with different configs)")
	}

	if rulesMatch(rulesDupA, rulesDupMismatch) {
		t.Errorf("expected rulesDupA and rulesDupMismatch NOT to match")
	}
}
