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

package servicemappingloader_test

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	PubSubGroup                                   = "pubsub.cnrm.cloud.google.com"
	PubSubTopicKind                               = "PubSubTopic"
	PubSubTopicResourceConfigName                 = "google_pubsub_topic"
	ComputeGroup                                  = "compute.cnrm.cloud.google.com"
	ComputeAddressKind                            = "ComputeAddress"
	ComputeAddressResourceConfigName              = "google_compute_address"
	ComputeGlobalAddressResourceConfigName        = "google_compute_global_address"
	ComputeInstanceKind                           = "ComputeInstance"
	ComputeInstanceResourceConfigName             = "google_compute_instance"
	ComputeInstanceFromTemplateResourceConfigName = "google_compute_instance_from_template"
)

type ServiceMappingTestCase struct {
	Name          string
	ShouldSucceed bool
	Group         string
}

var serviceMappingTestCases = []ServiceMappingTestCase{
	{"NilGroup", false, ""},
	{"InvalidGroup", false, "invalidgroup.google.com"},
	{"Group", true, PubSubGroup},
}

func TestGetServiceMapping(t *testing.T) {
	t.Parallel()
	smLoader := testservicemappingloader.New(t)
	for _, tc := range serviceMappingTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			getServiceMappingAssertResult(t, tc, smLoader)
		})
	}
}

func getServiceMappingAssertResult(t *testing.T, tc ServiceMappingTestCase, smLoader *servicemappingloader.ServiceMappingLoader) {
	sm, err := smLoader.GetServiceMapping(tc.Group)
	if tc.ShouldSucceed {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	} else {
		if err == nil {
			t.Fatalf("expected error, instead got nil")
		}
		return
	}
	if sm == nil {
		t.Fatalf("unexpected nil service mapping value for group '%v'", tc.Group)
	}
	if sm.Name != tc.Group {
		t.Errorf("mismatched value for 'name': got '%v', want '%v'", sm.Name, tc.Group)
	}
}

type ResourceConfigTestCase struct {
	Name               string
	ShouldSucceed      bool
	ResourceConfigName string
	Obj                *unstructured.Unstructured
}

var resourceConfigTestCases = []ResourceConfigTestCase{
	{"valid kind", true, PubSubTopicResourceConfigName, &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       PubSubTopicKind,
			"apiVersion": groupToAPIVersion(PubSubGroup),
		},
	}},
	{"invalid kind", false, "", &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "InvalidKind",
			"apiVersion": groupToAPIVersion(PubSubGroup),
		},
	}},
	{"regional resource", true, ComputeAddressResourceConfigName, &unstructured.Unstructured{
		Object: map[string]interface{}{
			"spec": map[string]interface{}{
				"location": "us-central1",
			},
			"kind":       ComputeAddressKind,
			"apiVersion": groupToAPIVersion(ComputeGroup),
		},
	}},
	{"global resource", true, ComputeGlobalAddressResourceConfigName, &unstructured.Unstructured{
		Object: map[string]interface{}{
			"spec": map[string]interface{}{
				"location": "global",
			},
			"kind":       ComputeAddressKind,
			"apiVersion": groupToAPIVersion(ComputeGroup),
		},
	}},
	{"global address resource without location", false, ComputeAddressResourceConfigName, &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       ComputeAddressKind,
			"apiVersion": groupToAPIVersion(ComputeGroup),
		},
	}},
	{"compute address resource with invalid location", false, ComputeAddressResourceConfigName, &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       ComputeAddressKind,
			"apiVersion": groupToAPIVersion(ComputeGroup),
			"spec": map[string]interface{}{
				"location": "asia-east1-a",
			},
		},
	}},
	{"KCC compute instance using compute_instance TF resource", true, ComputeInstanceResourceConfigName, &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       ComputeInstanceKind,
			"apiVersion": groupToAPIVersion(ComputeGroup),
		},
	}},
	{"KCC compute instance using compute_instance_from_template TF resource", true, ComputeInstanceFromTemplateResourceConfigName, &unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       ComputeInstanceKind,
			"apiVersion": groupToAPIVersion(ComputeGroup),
			"spec": map[string]interface{}{
				"instanceTemplateRef": map[string]interface{}{
					"name": "test-ref",
				},
			},
		},
	}},
}

func TestGetResourceConfig(t *testing.T) {
	t.Parallel()
	smLoader := testservicemappingloader.New(t)
	for _, tc := range resourceConfigTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			getResourceConfigAssertResult(t, tc, smLoader)
		})
	}
}

func getResourceConfigAssertResult(t *testing.T, tc ResourceConfigTestCase, smLoader *servicemappingloader.ServiceMappingLoader) {
	rc, err := smLoader.GetResourceConfig(tc.Obj)
	if tc.ShouldSucceed {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	} else {
		if err == nil {
			t.Fatalf("expected error, instead got nil")
		}
		return
	}
	if rc == nil {
		t.Fatalf("unexpected nil resource config value with gvk '%v'", tc.Obj.GroupVersionKind())
	}
	if rc.Name != tc.ResourceConfigName {
		t.Errorf("mismatched value for resource config name: got '%v', want '%v'", rc.Name, tc.ResourceConfigName)
	}
}

func groupToAPIVersion(groupName string) string {
	return fmt.Sprintf("%v/v1beta1", groupName)
}
