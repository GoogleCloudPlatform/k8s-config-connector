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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

func TestValidateRule(t *testing.T) {
	r := &KRMSyncerReconciler{}

	tests := []struct {
		name    string
		rule    krmv1alpha1.ResourceRule
		wantErr bool
	}{
		{
			name: "valid KCC glob",
			rule: krmv1alpha1.ResourceRule{
				Group:   "*.cnrm.cloud.google.com",
				Version: "*",
				Kind:    "*",
			},
			wantErr: false,
		},
		{
			name: "invalid group glob",
			rule: krmv1alpha1.ResourceRule{
				Group:   "*.foo.com",
				Version: "*",
				Kind:    "*",
			},
			wantErr: true,
		},
		{
			name: "invalid version glob",
			rule: krmv1alpha1.ResourceRule{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "*",
				Kind:    "Instance",
			},
			wantErr: true,
		},
		{
			name: "invalid kind glob",
			rule: krmv1alpha1.ResourceRule{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "*",
			},
			wantErr: true,
		},
		{
			name: "no glob",
			rule: krmv1alpha1.ResourceRule{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "Instance",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.validateRule(tt.rule)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRuleMatchesGVK(t *testing.T) {
	dr := &DynamicResourceReconciler{}

	tests := []struct {
		name string
		rule krmv1alpha1.ResourceRule
		gvk  schema.GroupVersionKind
		want bool
	}{
		{
			name: "KCC glob match",
			rule: krmv1alpha1.ResourceRule{
				Group:   "*.cnrm.cloud.google.com",
				Version: "*",
				Kind:    "*",
			},
			gvk: schema.GroupVersionKind{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "ComputeInstance",
			},
			want: true,
		},
		{
			name: "KCC glob mismatch",
			rule: krmv1alpha1.ResourceRule{
				Group:   "*.cnrm.cloud.google.com",
				Version: "*",
				Kind:    "*",
			},
			gvk: schema.GroupVersionKind{
				Group:   "example.com",
				Version: "v1",
				Kind:    "Foo",
			},
			want: false,
		},
		{
			name: "exact match",
			rule: krmv1alpha1.ResourceRule{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "ComputeInstance",
			},
			gvk: schema.GroupVersionKind{
				Group:   "compute.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "ComputeInstance",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, dr.ruleMatchesGVK(tt.rule, tt.gvk))
		})
	}
}

// TestExpandRulePreferredVersionOnly checks that, when fed discovery results from
// ServerPreferredResources, a KCC glob expands to exactly one GVK per kind even when
// a kind is served in multiple versions, and that kinds only served in a
// non-preferred version are still included.
func TestExpandRulePreferredVersionOnly(t *testing.T) {
	const group = "pubsub.cnrm.cloud.google.com"

	writeJSON := func(w http.ResponseWriter, obj any) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(obj)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.URL.Path {
		case "/api":
			writeJSON(w, &metav1.APIVersions{})
		case "/apis":
			writeJSON(w, &metav1.APIGroupList{Groups: []metav1.APIGroup{{
				Name: group,
				Versions: []metav1.GroupVersionForDiscovery{
					{GroupVersion: group + "/v1beta1", Version: "v1beta1"},
					{GroupVersion: group + "/v1alpha1", Version: "v1alpha1"},
				},
				PreferredVersion: metav1.GroupVersionForDiscovery{GroupVersion: group + "/v1beta1", Version: "v1beta1"},
			}}})
		case "/apis/" + group + "/v1beta1":
			writeJSON(w, &metav1.APIResourceList{GroupVersion: group + "/v1beta1", APIResources: []metav1.APIResource{
				{Name: "pubsubtopics", Kind: "PubSubTopic", Namespaced: true},
				{Name: "pubsubtopics/status", Kind: "PubSubTopic", Namespaced: true},
			}})
		case "/apis/" + group + "/v1alpha1":
			writeJSON(w, &metav1.APIResourceList{GroupVersion: group + "/v1alpha1", APIResources: []metav1.APIResource{
				{Name: "pubsubtopics", Kind: "PubSubTopic", Namespaced: true},
				{Name: "pubsubschemas", Kind: "PubSubSchema", Namespaced: true},
			}})
		default:
			http.NotFound(w, req)
		}
	}))
	defer server.Close()

	dc, err := discovery.NewDiscoveryClientForConfig(&rest.Config{Host: server.URL})
	require.NoError(t, err)
	resourceLists, err := dc.ServerPreferredResources()
	require.NoError(t, err)

	r := &KRMSyncerReconciler{}
	rule := krmv1alpha1.ResourceRule{Group: "*.cnrm.cloud.google.com", Version: "*", Kind: "*"}
	gvks, err := r.expandRule(t.Context(), nil, rule, krmv1alpha1.ModePull, resourceLists)
	require.NoError(t, err)

	assert.ElementsMatch(t, []schema.GroupVersionKind{
		{Group: group, Version: "v1beta1", Kind: "PubSubTopic"},   // preferred version only
		{Group: group, Version: "v1alpha1", Kind: "PubSubSchema"}, // only served in v1alpha1
	}, gvks)
}
