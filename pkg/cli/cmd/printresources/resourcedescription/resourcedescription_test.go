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

package resourcedescription

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestGetAll(t *testing.T) {
	resourceDescriptions, err := GetAll()
	if err != nil {
		t.Fatalf("unable to call GetAll(): got err '%v'", err)
	}
	resourceDescriptionByGVK := toMap(resourceDescriptions)

	testCases := []struct {
		GVK                     schema.GroupVersionKind
		ShouldSupportIAM        bool
		ShouldSupportBulkExport bool
	}{
		{
			GVK: schema.GroupVersionKind{
				Group:   "pubsub.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "PubSubSubscription",
			},
			ShouldSupportIAM:        true,
			ShouldSupportBulkExport: true,
		},
		{
			GVK: schema.GroupVersionKind{
				Group:   "bigtable.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "BigtableAppProfile",
			},
			ShouldSupportIAM:        false,
			ShouldSupportBulkExport: true,
		},
		{
			GVK: schema.GroupVersionKind{
				Group:   "resourcemanager.cnrm.cloud.google.com",
				Version: "v1beta1",
				Kind:    "ResourceManagerLien",
			},
			ShouldSupportIAM:        false,
			ShouldSupportBulkExport: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.GVK.String(), func(t *testing.T) {
			resourceDescription, exists := resourceDescriptionByGVK[tc.GVK]
			if !exists {
				t.Fatalf("unable to find resourceDescription for GVK %v", tc.GVK)
			}
			got := resourceDescription.SupportsIAM
			if got != tc.ShouldSupportIAM {
				t.Errorf("%v.SupportsIAM is %v, want %v.", tc.GVK, got, tc.ShouldSupportIAM)
			}
			got = resourceDescription.SupportsBulkExport
			if got != tc.ShouldSupportBulkExport {
				t.Errorf("%v.SupportsBulkExport is %v, want %v.", tc.GVK, got, tc.ShouldSupportBulkExport)
			}
		})
	}
}
