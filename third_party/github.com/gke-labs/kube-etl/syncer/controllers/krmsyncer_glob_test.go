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
	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"
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
