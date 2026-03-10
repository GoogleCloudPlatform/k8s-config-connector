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

package preview

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestToTrackedGVR(t *testing.T) {
	tests := []struct {
		name                        string
		apiResource                 metav1.APIResource
		apiResourceListGroupVersion schema.GroupVersion
		wantGVR                     schema.GroupVersionResource
		wantOk                      bool
	}{
		{
			name: "Valid CNRM resource",
			apiResource: metav1.APIResource{
				Name:    "storagebuckets",
				Group:   "storage.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "storage.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "storage.cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "storagebuckets",
			},
			wantOk: true,
		},
		{
			name: "Valid CNRM resource - inherit Group/Version",
			apiResource: metav1.APIResource{
				Name: "storagebuckets",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "storage.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "storage.cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "storagebuckets",
			},
			wantOk: true,
		},
		{
			name: "Core CNRM resource (ignored)",
			apiResource: metav1.APIResource{
				Name:    "configconnectors",
				Group:   "core.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "core.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "core.cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "configconnectors",
			},
			wantOk: false,
		},
		{
			name: "Non-CNRM resource (ignored)",
			apiResource: metav1.APIResource{
				Name:    "deployments",
				Group:   "apps",
				Version: "v1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "apps",
				Version: "v1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "apps",
				Version:  "v1",
				Resource: "deployments",
			},
			wantOk: false,
		},
		{
			name: "Ignored CRD (gameservicesrealms)",
			apiResource: metav1.APIResource{
				Name:    "gameservicesrealms",
				Group:   "gameservices.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "gameservices.cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "gameservices.cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "gameservicesrealms",
			},
			wantOk: false,
		},
		{
			name: "Fake CNRM group (ignored)",
			apiResource: metav1.APIResource{
				Name:    "storagebuckets",
				Group:   "fake-cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			apiResourceListGroupVersion: schema.GroupVersion{
				Group:   "fake-cnrm.cloud.google.com",
				Version: "v1beta1",
			},
			wantGVR: schema.GroupVersionResource{
				Group:    "fake-cnrm.cloud.google.com",
				Version:  "v1beta1",
				Resource: "storagebuckets",
			},
			wantOk: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotGVR, gotOk := toTrackedGVR(tc.apiResource, tc.apiResourceListGroupVersion)
			if diff := cmp.Diff(tc.wantGVR, gotGVR); diff != "" {
				t.Errorf("toTrackedGVR() GVR mismatch (-want +got):\n%s", diff)
			}
			if gotOk != tc.wantOk {
				t.Errorf("toTrackedGVR() ok = %v, want %v", gotOk, tc.wantOk)
			}
		})
	}
}
