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

package backup

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestSortResources(t *testing.T) {
	resources := []*unstructured.Unstructured{
		{
			Object: map[string]interface{}{
				"kind":       "IAMPolicyMember",
				"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "policy-member",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "ComputeNetwork",
				"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "network",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "Project",
				"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "project",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "PubSubTopic",
				"apiVersion": "pubsub.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "topic",
				},
			},
		},
	}

	sortResources(resources)

	expectedOrder := []string{"Project", "ComputeNetwork", "PubSubTopic", "IAMPolicyMember"}
	for i, res := range resources {
		if res.GetKind() != expectedOrder[i] {
			t.Errorf("At index %d, expected kind %s, got %s", i, expectedOrder[i], res.GetKind())
		}
	}
}

func TestSortResourcesComplex(t *testing.T) {
	resources := []*unstructured.Unstructured{
		{
			Object: map[string]interface{}{
				"kind":       "ComputeSubnetwork",
				"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "subnetwork",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "IAMServiceAccount",
				"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "sa",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "Folder",
				"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "folder",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "ComputeNetwork",
				"apiVersion": "compute.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "network",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "Organization",
				"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "org",
				},
			},
		},
		{
			Object: map[string]interface{}{
				"kind":       "Project",
				"apiVersion": "resourcemanager.cnrm.cloud.google.com/v1beta1",
				"metadata": map[string]interface{}{
					"name": "project",
				},
			},
		},
	}

	sortResources(resources)

	expectedOrder := []string{"Organization", "Folder", "Project", "IAMServiceAccount", "ComputeNetwork", "ComputeSubnetwork"}
	for i, res := range resources {
		if res.GetKind() != expectedOrder[i] {
			t.Errorf("At index %d, expected kind %s, got %s", i, expectedOrder[i], res.GetKind())
		}
	}
}

