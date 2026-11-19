// Copyright 2026 Google LLC
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

package controllers

import (
	"testing"

	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestKCCGlobMatchesGroup(t *testing.T) {
	tests := []struct {
		group string
		want  bool
	}{
		{group: "storage.cnrm.cloud.google.com", want: true},
		{group: "compute.cnrm.cloud.google.com", want: true},
		// A resource group whose name merely starts with "core".
		{group: "coreservices.cnrm.cloud.google.com", want: true},
		{group: "core.cnrm.cloud.google.com", want: false},
		{group: "customize.core.cnrm.cloud.google.com", want: false},
		{group: "multicluster.core.cnrm.cloud.google.com", want: false},
		{group: "example.com", want: false},
		{group: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.group, func(t *testing.T) {
			assert.Equal(t, tt.want, kccGlobMatchesGroup(tt.group))
		})
	}
}

func TestExpandRuleKCCGlobSkipsSystemGroups(t *testing.T) {
	r := &KRMSyncerReconciler{}
	resourceLists := []*metav1.APIResourceList{
		{
			GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
			APIResources: []metav1.APIResource{
				{Name: "storagebuckets", Kind: "StorageBucket"},
				{Name: "storagebuckets/status", Kind: "StorageBucket"},
			},
		},
		{
			GroupVersion: "core.cnrm.cloud.google.com/v1beta1",
			APIResources: []metav1.APIResource{
				{Name: "configconnectors", Kind: "ConfigConnector"},
				{Name: "configconnectorcontexts", Kind: "ConfigConnectorContext"},
			},
		},
		{
			GroupVersion: "customize.core.cnrm.cloud.google.com/v1beta1",
			APIResources: []metav1.APIResource{{Name: "controllerresources", Kind: "ControllerResource"}},
		},
		{
			GroupVersion: "multicluster.core.cnrm.cloud.google.com/v1alpha1",
			APIResources: []metav1.APIResource{{Name: "multiclusterleases", Kind: "MultiClusterLease"}},
		},
		{
			GroupVersion: "example.com/v1",
			APIResources: []metav1.APIResource{{Name: "foos", Kind: "Foo"}},
		},
	}

	t.Run("the wildcard matches resource groups only", func(t *testing.T) {
		rule := krmv1alpha1.ResourceRule{Group: "*.cnrm.cloud.google.com", Version: "*", Kind: "*"}

		got, err := r.expandRule(t.Context(), nil, rule, krmv1alpha1.ModePull, resourceLists)

		require.NoError(t, err)
		assert.Equal(t, []schema.GroupVersionKind{
			{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageBucket"},
		}, got)
	})

	t.Run("a rule naming a system group still matches it", func(t *testing.T) {
		rule := krmv1alpha1.ResourceRule{Group: "multicluster.core.cnrm.cloud.google.com", Version: "*", Kind: "*"}

		got, err := r.expandRule(t.Context(), nil, rule, krmv1alpha1.ModePull, resourceLists)

		require.NoError(t, err)
		assert.Equal(t, []schema.GroupVersionKind{
			{Group: "multicluster.core.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MultiClusterLease"},
		}, got)
	})
}

func TestRuleMatchesGVKKCCGlobSkipsSystemGroups(t *testing.T) {
	dr := &DynamicResourceReconciler{}
	wildcard := krmv1alpha1.ResourceRule{Group: "*.cnrm.cloud.google.com", Version: "*", Kind: "*"}
	lease := schema.GroupVersionKind{Group: "multicluster.core.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "MultiClusterLease"}

	tests := []struct {
		name string
		rule krmv1alpha1.ResourceRule
		gvk  schema.GroupVersionKind
		want bool
	}{
		{
			name: "wildcard does not match ConfigConnector",
			rule: wildcard,
			gvk:  schema.GroupVersionKind{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ConfigConnector"},
			want: false,
		},
		{
			name: "wildcard does not match MultiClusterLease",
			rule: wildcard,
			gvk:  lease,
			want: false,
		},
		{
			name: "wildcard matches a resource group",
			rule: wildcard,
			gvk:  schema.GroupVersionKind{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Kind: "StorageBucket"},
			want: true,
		},
		{
			name: "a rule naming the system group matches it",
			rule: krmv1alpha1.ResourceRule{Group: "multicluster.core.cnrm.cloud.google.com", Version: "*", Kind: "*"},
			gvk:  lease,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, dr.ruleMatchesGVK(tt.rule, tt.gvk))
		})
	}
}
