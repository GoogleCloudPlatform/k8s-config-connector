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

package mcp

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery/fake"
	dynamicfake "k8s.io/client-go/dynamic/fake"
)

func TestFindGVR(t *testing.T) {
	scheme := runtime.NewScheme()
	dynamicClient := dynamicfake.NewSimpleDynamicClient(scheme)
	discoveryClient := &fake.FakeDiscovery{
		Fake: &dynamicClient.Fake,
	}

	discoveryClient.Resources = []*metav1.APIResourceList{
		{
			GroupVersion: "sql.cnrm.cloud.google.com/v1beta1",
			APIResources: []metav1.APIResource{
				{
					Name: "sqlinstances",
					Kind: "SQLInstance",
				},
			},
		},
	}

	s := &Server{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
	}

	gvr, err := s.findGVR("SQLInstance")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedGVR := schema.GroupVersionResource{
		Group:    "sql.cnrm.cloud.google.com",
		Version:  "v1beta1",
		Resource: "sqlinstances",
	}

	if gvr != expectedGVR {
		t.Errorf("expected GVR %v, got %v", expectedGVR, gvr)
	}
}

func TestDescribeKCCResource(t *testing.T) {
	scheme := runtime.NewScheme()
	
	resource := &unstructured.Unstructured{}
	resource.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "sql.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "SQLInstance",
	})
	resource.SetName("test-instance")
	resource.SetNamespace("test-namespace")
	resource.Object["status"] = map[string]interface{}{
		"conditions": []interface{}{
			map[string]interface{}{
				"type":   "Ready",
				"status": "True",
			},
		},
	}

	dynamicClient := dynamicfake.NewSimpleDynamicClient(scheme, resource)
	discoveryClient := &fake.FakeDiscovery{
		Fake: &dynamicClient.Fake,
	}
	discoveryClient.Resources = []*metav1.APIResourceList{
		{
			GroupVersion: "sql.cnrm.cloud.google.com/v1beta1",
			APIResources: []metav1.APIResource{
				{
					Name: "sqlinstances",
					Kind: "SQLInstance",
				},
			},
		},
	}

	s := &Server{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
	}

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]interface{}{
		"kind":      "SQLInstance",
		"namespace": "test-namespace",
		"name":      "test-instance",
	}

	resp, err := s.describeKCCResource(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Content) == 0 {
		t.Fatal("expected content in response")
	}

	content := resp.Content[0].(mcp.TextContent)
	if content.Text == "" {
		t.Error("expected status JSON in response")
	}
}

func TestGetKCCCRDSchema(t *testing.T) {
	scheme := runtime.NewScheme()

	crd := &unstructured.Unstructured{}
	crd.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "apiextensions.k8s.io",
		Version: "v1",
		Kind:    "CustomResourceDefinition",
	})
	crd.SetName("sqlinstances.sql.cnrm.cloud.google.com")
	crd.Object["spec"] = map[string]interface{}{
		"versions": []interface{}{
			map[string]interface{}{
				"name": "v1beta1",
				"schema": map[string]interface{}{
					"openAPIV3Schema": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"spec": map[string]interface{}{
								"type": "object",
							},
							"status": map[string]interface{}{
								"type": "object",
							},
						},
					},
				},
			},
		},
	}

	dynamicClient := dynamicfake.NewSimpleDynamicClient(scheme, crd)
	discoveryClient := &fake.FakeDiscovery{
		Fake: &dynamicClient.Fake,
	}
	discoveryClient.Resources = []*metav1.APIResourceList{
		{
			GroupVersion: "sql.cnrm.cloud.google.com/v1beta1",
			APIResources: []metav1.APIResource{
				{
					Name: "sqlinstances",
					Kind: "SQLInstance",
				},
			},
		},
	}

	s := &Server{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
	}

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]interface{}{
		"kind": "SQLInstance",
	}

	resp, err := s.getKCCCRDSchema(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Content) == 0 {
		t.Fatal("expected content in response")
	}

	content := resp.Content[0].(mcp.TextContent)
	if content.Text == "" {
		t.Error("expected schema JSON in response")
	}
}
