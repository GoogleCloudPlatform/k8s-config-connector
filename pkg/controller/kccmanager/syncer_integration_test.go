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
	"errors"
	"reflect"
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
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

type fakeDiscovery struct {
	groups    []*metav1.APIGroup
	resources []*metav1.APIResourceList
	err       error
}

func (f *fakeDiscovery) ServerGroupsAndResources() ([]*metav1.APIGroup, []*metav1.APIResourceList, error) {
	return f.groups, f.resources, f.err
}

func TestServedGVKsWithServiceGeneratedIDs(t *testing.T) {
	// 1. Not initialized
	siUninit := &SyncerIntegration{}
	if _, err := siUninit.servedGVKsWithServiceGeneratedIDs(); err == nil || !strings.Contains(err.Error(), "discovery client not initialized") {
		t.Fatalf("expected 'discovery client not initialized', got %v", err)
	}

	// 2. Served and unserved kinds, two served versions of one kind, subresources, and kinds not on list
	fake := &fakeDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "resourcemanager.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "folders", Kind: "Folder"},
					{Name: "folders/status", Kind: "Folder"}, // Subresource should be ignored
					{Name: "projects", Kind: "Project"},      // Not in syncerGVKsWithServiceGeneratedIDs
				},
			},
			{
				GroupVersion: "essentialcontacts.cnrm.cloud.google.com/v1alpha1",
				APIResources: []metav1.APIResource{
					{Name: "essentialcontactscontacts", Kind: "EssentialContactsContact"},
				},
			},
			{
				GroupVersion: "essentialcontacts.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "essentialcontactscontacts", Kind: "EssentialContactsContact"},
				},
			},
			{
				GroupVersion: "invalid-group-version",
				APIResources: []metav1.APIResource{
					{Name: "invalid", Kind: "Invalid"},
				},
			},
		},
	}
	si := &SyncerIntegration{discovery: fake}
	gvks, err := si.servedGVKsWithServiceGeneratedIDs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []schema.GroupVersionKind{
		{Group: "essentialcontacts.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "EssentialContactsContact"},
		{Group: "essentialcontacts.cnrm.cloud.google.com", Version: "v1beta1", Kind: "EssentialContactsContact"},
		{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Folder"},
	}
	if !reflect.DeepEqual(gvks, expected) {
		t.Fatalf("expected %v, got %v", expected, gvks)
	}

	// 3. Discovery failing for a group not on the list is tolerated
	fakeIgnoredErr := &fakeDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "resourcemanager.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "folders", Kind: "Folder"},
				},
			},
		},
		err: &discovery.ErrGroupDiscoveryFailed{
			Groups: map[schema.GroupVersion]error{
				{Group: "custom.example.com", Version: "v1"}: errors.New("group down"),
			},
		},
	}
	siIgnoredErr := &SyncerIntegration{discovery: fakeIgnoredErr}
	gvks, err = siIgnoredErr.servedGVKsWithServiceGeneratedIDs()
	if err != nil {
		t.Fatalf("expected discovery failure on unrelated group to be ignored, got: %v", err)
	}
	if len(gvks) != 1 || gvks[0].Kind != "Folder" {
		t.Fatalf("expected [Folder], got %v", gvks)
	}

	// 4. Discovery failing for a group on the list returns error
	fakeListErr := &fakeDiscovery{
		err: &discovery.ErrGroupDiscoveryFailed{
			Groups: map[schema.GroupVersion]error{
				{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1"}: errors.New("group down"),
			},
		},
	}
	siListErr := &SyncerIntegration{discovery: fakeListErr}
	if _, err := siListErr.servedGVKsWithServiceGeneratedIDs(); err == nil || !strings.Contains(err.Error(), "error discovering the kinds of resourcemanager.cnrm.cloud.google.com/v1beta1") {
		t.Fatalf("expected error for group on the list, got %v", err)
	}

	// 5. General discovery error returns error
	fakeGeneralErr := &fakeDiscovery{
		err: errors.New("connection refused"),
	}
	siGeneralErr := &SyncerIntegration{discovery: fakeGeneralErr}
	if _, err := siGeneralErr.servedGVKsWithServiceGeneratedIDs(); err == nil || !strings.Contains(err.Error(), "error discovering the kinds that the cluster serves") {
		t.Fatalf("expected error for general discovery failure, got %v", err)
	}
}

func TestGVKsWithOwnRule_FullMode(t *testing.T) {
	// In Full mode, gvksWithOwnRule should return nil, nil without calling discovery
	for _, mode := range []string{"Full", "full", "FULL"} {
		si := &SyncerIntegration{
			replicationMode: mode,
			discovery:       nil, // nil discovery verifies it is never called
		}
		gvks, err := si.gvksWithOwnRule()
		if err != nil {
			t.Fatalf("mode %q: unexpected error: %v", mode, err)
		}
		if gvks != nil {
			t.Fatalf("mode %q: expected nil gvks, got %v", mode, gvks)
		}
	}
}

func TestGetDesiredRules(t *testing.T) {
	si := &SyncerIntegration{
		replicationMode: "Status",
	}

	gvks := []schema.GroupVersionKind{
		{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Folder"},
	}
	rules := si.getDesiredRules([]string{"tenant-a"}, gvks)
	if len(rules) != len(gvks)+1 {
		t.Fatalf("unexpected number of rules: %d", len(rules))
	}

	hasGlob := false
	hasFolder := false
	for _, r := range rules {
		ruleMap := r.(map[string]interface{})
		if ruleMap["group"] == "*.cnrm.cloud.google.com" {
			hasGlob = true
			nsList := ruleMap["namespaces"].([]interface{})
			if len(nsList) != 1 || nsList[0] != "tenant-a" {
				t.Errorf("expected namespaces [tenant-a], got %v", nsList)
			}
		}
		if ruleMap["kind"] == "Folder" {
			hasFolder = true
			if ruleMap["group"] != "resourcemanager.cnrm.cloud.google.com" || ruleMap["version"] != "v1beta1" {
				t.Errorf("unexpected Folder rule: %v", ruleMap)
			}
		}
	}
	if !hasGlob {
		t.Errorf("glob rule missing")
	}
	if !hasFolder {
		t.Errorf("Folder rule missing")
	}

	// Cluster scoped
	siCluster := &SyncerIntegration{
		replicationMode: "Full",
	}
	clusterRules := siCluster.getDesiredRules(nil, gvks)
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
