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
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	dynamicfake "k8s.io/client-go/dynamic/fake"
)

type mockDiscovery struct {
	discovery.DiscoveryInterface
	resources []*metav1.APIResourceList
}

func (m *mockDiscovery) ServerPreferredResources() ([]*metav1.APIResourceList, error) {
	return m.resources, nil
}

func (m *mockDiscovery) ServerResourcesForGroupVersion(groupVersion string) (*metav1.APIResourceList, error) {
	for _, res := range m.resources {
		if res.GroupVersion == groupVersion {
			return res, nil
		}
	}
	return nil, fmt.Errorf("resources not found for group version %s", groupVersion)
}

func TestHandleListKCCResources(t *testing.T) {
	scheme := runtime.NewScheme()
	gvr := schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Resource: "storagebuckets"}
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, map[schema.GroupVersionResource]string{
		gvr: "StorageBucketList",
	})
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "storagebuckets", Kind: "StorageBucket", Namespaced: true},
				},
			},
		},
	}

	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	// Create a fake resource
	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
			"kind":       "StorageBucket",
			"metadata": map[string]interface{}{
				"name":      "test-bucket",
				"namespace": "test-ns",
				"annotations": map[string]interface{}{
					"cnrm.cloud.google.com/project-id": "test-project",
				},
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "True",
						"reason": "UpToDate",
					},
				},
			},
		},
	}
	_, err := dynamicClient.Resource(gvr).Namespace("test-ns").Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("failed to create fake resource: %v", err)
	}

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"kind": "StorageBucket",
			},
		},
	}

	res, err := sc.handleListKCCResources(context.Background(), req)
	if err != nil {
		t.Fatalf("handleListKCCResources failed: %v", err)
	}

	if res.IsError {
		t.Fatalf("tool returned error: %v", res.Content[0].(mcp.TextContent).Text)
	}

	text := res.Content[0].(mcp.TextContent).Text
	if !strings.Contains(text, "test-bucket") {
		t.Errorf("expected output to contain test-bucket, got %s", text)
	}
	if !strings.Contains(text, "test-project") {
		t.Errorf("expected output to contain test-project, got %s", text)
	}
	if !strings.Contains(text, "Status: Ready") {
		t.Errorf("expected output to contain Status: Ready, got %s", text)
	}
}

func TestHandleListKCCResourcesWithLimit(t *testing.T) {
	scheme := runtime.NewScheme()
	gvr := schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Resource: "storagebuckets"}
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, map[schema.GroupVersionResource]string{
		gvr: "StorageBucketList",
	})
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "storagebuckets", Kind: "StorageBucket", Namespaced: true},
				},
			},
		},
	}

	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	// Create 3 fake resources
	for i := 1; i <= 3; i++ {
		obj := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
				"kind":       "StorageBucket",
				"metadata": map[string]interface{}{
					"name":      fmt.Sprintf("test-bucket-%d", i),
					"namespace": "test-ns",
				},
				"status": map[string]interface{}{
					"conditions": []interface{}{
						map[string]interface{}{
							"type":   "Ready",
							"status": "True",
							"reason": "UpToDate",
						},
					},
				},
			},
		}
		_, err := dynamicClient.Resource(gvr).Namespace("test-ns").Create(context.Background(), obj, metav1.CreateOptions{})
		if err != nil {
			t.Fatalf("failed to create fake resource: %v", err)
		}
	}

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"kind":  "StorageBucket",
				"limit": 2.0, // mcp-go uses float64 for numbers
			},
		},
	}

	res, err := sc.handleListKCCResources(context.Background(), req)
	if err != nil {
		t.Fatalf("handleListKCCResources failed: %v", err)
	}

	text := res.Content[0].(mcp.TextContent).Text
	lines := strings.Split(strings.TrimSpace(text), "\n")

	count := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "- Kind:") {
			count++
		}
	}
	if count != 2 {
		t.Errorf("expected 2 resources, got %d. Output: %s", count, text)
	}
	if !strings.Contains(text, "Status: Ready") {
		t.Errorf("expected output to contain Status: Ready, got %s", text)
	}
	if !strings.Contains(text, "truncated to limit of 2") {
		t.Errorf("expected output to contain truncation note, got %s", text)
	}
}

func TestHandleGetKCCCRDSchema(t *testing.T) {
	scheme := runtime.NewScheme()
	crdGVR := schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, map[schema.GroupVersionResource]string{
		crdGVR: "CustomResourceDefinitionList",
	})
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "storagebuckets", Kind: "StorageBucket", Namespaced: true},
				},
			},
		},
	}
	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	crd := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"name": "storagebuckets.storage.cnrm.cloud.google.com",
			},
			"spec": map[string]interface{}{
				"group": "storage.cnrm.cloud.google.com",
				"names": map[string]interface{}{
					"kind": "StorageBucket",
				},
				"versions": []interface{}{
					map[string]interface{}{
						"name":   "v1beta1",
						"served": true,
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"type": "object",
								"properties": map[string]interface{}{
									"spec": map[string]interface{}{
										"type": "object",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	_, err := dynamicClient.Resource(crdGVR).Create(context.Background(), crd, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("failed to create fake CRD: %v", err)
	}

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"kind": "StorageBucket",
			},
		},
	}

	res, err := sc.handleGetKCCCRDSchema(context.Background(), req)
	if err != nil {
		t.Fatalf("handleGetKCCCRDSchema failed: %v", err)
	}

	if res.IsError {
		t.Fatalf("tool returned error: %v", res.Content[0].(mcp.TextContent).Text)
	}

	text := res.Content[0].(mcp.TextContent).Text
	var resSchema map[string]interface{}
	if err := json.Unmarshal([]byte(text), &resSchema); err != nil {
		t.Fatalf("failed to unmarshal schema JSON: %v", err)
	}

	if _, ok := resSchema["spec"]; !ok {
		t.Errorf("expected spec to be in the lean schema")
	}
}

func TestHandleGetKCCCRDSchemaWithMultipleVersions(t *testing.T) {
	scheme := runtime.NewScheme()
	crdGVR := schema.GroupVersionResource{Group: "apiextensions.k8s.io", Version: "v1", Resource: "customresourcedefinitions"}
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, map[schema.GroupVersionResource]string{
		crdGVR: "CustomResourceDefinitionList",
	})
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "storagebuckets", Kind: "StorageBucket", Namespaced: true},
				},
			},
		},
	}
	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	crd := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"name": "storagebuckets.storage.cnrm.cloud.google.com",
			},
			"spec": map[string]interface{}{
				"group": "storage.cnrm.cloud.google.com",
				"names": map[string]interface{}{
					"kind": "StorageBucket",
				},
				"versions": []interface{}{
					map[string]interface{}{
						"name":    "v1alpha1",
						"served":  true,
						"storage": false,
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"properties": map[string]interface{}{
									"spec": map[string]interface{}{
										"description": "v1alpha1 spec",
									},
								},
							},
						},
					},
					map[string]interface{}{
						"name":    "v1beta1",
						"served":  true,
						"storage": true,
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"properties": map[string]interface{}{
									"spec": map[string]interface{}{
										"description": "v1beta1 spec (storage)",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	_, err := dynamicClient.Resource(crdGVR).Create(context.Background(), crd, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("failed to create fake CRD: %v", err)
	}

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"kind": "StorageBucket",
			},
		},
	}

	res, err := sc.handleGetKCCCRDSchema(context.Background(), req)
	if err != nil {
		t.Fatalf("handleGetKCCCRDSchema failed: %v", err)
	}

	text := res.Content[0].(mcp.TextContent).Text
	if !strings.Contains(text, "v1beta1 spec (storage)") {
		t.Errorf("expected storage version (v1beta1) to be selected, got %s", text)
	}
}

func TestHandleDescribeKCCResource(t *testing.T) {
	scheme := runtime.NewScheme()
	gvr := schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Resource: "storagebuckets"}
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, map[schema.GroupVersionResource]string{
		gvr: "StorageBucketList",
	})
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "storagebuckets", Kind: "StorageBucket", Namespaced: true},
				},
			},
		},
	}

	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	// Create a fake resource with status and conditions, using UpToDate as type
	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
			"kind":       "StorageBucket",
			"metadata": map[string]interface{}{
				"name":      "test-bucket",
				"namespace": "test-ns",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":    "UpToDate",
						"status":  "True",
						"reason":  "UpToDate",
						"message": "Resource is up to date",
					},
				},
			},
		},
	}
	_, err := dynamicClient.Resource(gvr).Namespace("test-ns").Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("failed to create fake resource: %v", err)
	}

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"kind":      "StorageBucket",
				"namespace": "test-ns",
				"name":      "test-bucket",
			},
		},
	}

	res, err := sc.handleDescribeKCCResource(context.Background(), req)
	if err != nil {
		t.Fatalf("handleDescribeKCCResource failed: %v", err)
	}

	if res.IsError {
		t.Fatalf("tool returned error: %v", res.Content[0].(mcp.TextContent).Text)
	}

	text := res.Content[0].(mcp.TextContent).Text
	if !strings.Contains(text, "test-bucket") {
		t.Errorf("expected output to contain test-bucket, got %s", text)
	}
	if !strings.Contains(text, "UpToDate") {
		t.Errorf("expected output to contain UpToDate, got %s", text)
	}
	if !strings.Contains(text, "Resource is up to date") {
		t.Errorf("expected output to contain 'Resource is up to date', got %s", text)
	}
}

func TestHandleDescribeKCCResourceClusterScoped(t *testing.T) {
	scheme := runtime.NewScheme()
	gvr := schema.GroupVersionResource{Group: "core.cnrm.cloud.google.com", Version: "v1beta1", Resource: "configconnectors"}
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, map[schema.GroupVersionResource]string{
		gvr: "ConfigConnectorList",
	})
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "core.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "configconnectors", Kind: "ConfigConnector", Namespaced: false},
				},
			},
		},
	}

	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	// Create a fake cluster-scoped resource
	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "core.cnrm.cloud.google.com/v1beta1",
			"kind":       "ConfigConnector",
			"metadata": map[string]interface{}{
				"name": "configconnector-test",
			},
			"status": map[string]interface{}{
				"conditions": []interface{}{
					map[string]interface{}{
						"type":   "Ready",
						"status": "True",
						"reason": "UpToDate",
					},
				},
			},
		},
	}
	_, err := dynamicClient.Resource(gvr).Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("failed to create fake resource: %v", err)
	}

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"kind": "ConfigConnector",
				"name": "configconnector-test",
			},
		},
	}

	res, err := sc.handleDescribeKCCResource(context.Background(), req)
	if err != nil {
		t.Fatalf("handleDescribeKCCResource failed: %v", err)
	}

	text := res.Content[0].(mcp.TextContent).Text
	if !strings.Contains(text, "configconnector-test") {
		t.Errorf("expected output to contain configconnector-test, got %s", text)
	}
}

func TestHandleApplyKCCYAML(t *testing.T) {
	scheme := runtime.NewScheme()
	gvr := schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Resource: "storagebuckets"}
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, map[schema.GroupVersionResource]string{
		gvr: "StorageBucketList",
	})
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "storagebuckets", Kind: "StorageBucket", Namespaced: true},
				},
			},
		},
	}

	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	yamlStr := `
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: apply-test-bucket
  namespace: test-ns
spec:
  location: US
---
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  name: apply-test-bucket-2
  namespace: test-ns
spec:
  location: EU
`

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"yaml": yamlStr,
			},
		},
	}

	res, err := sc.handleApplyKCCYAML(context.Background(), req)
	if err != nil {
		t.Fatalf("handleApplyKCCYAML failed: %v", err)
	}

	if res.IsError {
		t.Fatalf("tool returned error: %v", res.Content[0].(mcp.TextContent).Text)
	}

	text := res.Content[0].(mcp.TextContent).Text
	if !strings.Contains(text, "Successfully applied test-ns/apply-test-bucket (StorageBucket)") {
		t.Errorf("expected output to contain success message for bucket 1, got %s", text)
	}
	if !strings.Contains(text, "Successfully applied test-ns/apply-test-bucket-2 (StorageBucket)") {
		t.Errorf("expected output to contain success message for bucket 2, got %s", text)
	}

	// Verify resources were created
	_, err = dynamicClient.Resource(gvr).Namespace("test-ns").Get(context.Background(), "apply-test-bucket", metav1.GetOptions{})
	if err != nil {
		t.Errorf("failed to get apply-test-bucket: %v", err)
	}
	_, err = dynamicClient.Resource(gvr).Namespace("test-ns").Get(context.Background(), "apply-test-bucket-2", metav1.GetOptions{})
	if err != nil {
		t.Errorf("failed to get apply-test-bucket-2: %v", err)
	}
}

func TestHandleGetKCCResource(t *testing.T) {
	scheme := runtime.NewScheme()
	gvr := schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Resource: "storagebuckets"}
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, map[schema.GroupVersionResource]string{
		gvr: "StorageBucketList",
	})
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "storagebuckets", Kind: "StorageBucket", Namespaced: true},
				},
			},
		},
	}

	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
			"kind":       "StorageBucket",
			"metadata": map[string]interface{}{
				"name":      "get-test-bucket",
				"namespace": "test-ns",
				"managedFields": []interface{}{
					map[string]interface{}{
						"manager": "test",
					},
				},
			},
			"spec": map[string]interface{}{
				"location": "US",
			},
		},
	}
	_, err := dynamicClient.Resource(gvr).Namespace("test-ns").Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("failed to create fake resource: %v", err)
	}

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"kind":      "StorageBucket",
				"namespace": "test-ns",
				"name":      "get-test-bucket",
			},
		},
	}

	res, err := sc.handleGetKCCResource(context.Background(), req)
	if err != nil {
		t.Fatalf("handleGetKCCResource failed: %v", err)
	}

	if res.IsError {
		t.Fatalf("tool returned error: %v", res.Content[0].(mcp.TextContent).Text)
	}

	text := res.Content[0].(mcp.TextContent).Text
	if !strings.Contains(text, "kind: StorageBucket") {
		t.Errorf("expected output to contain kind, got %s", text)
	}
	if !strings.Contains(text, "name: get-test-bucket") {
		t.Errorf("expected output to contain name, got %s", text)
	}
	if !strings.Contains(text, "location: US") {
		t.Errorf("expected output to contain location, got %s", text)
	}
	if strings.Contains(text, "managedFields") {
		t.Errorf("expected output NOT to contain managedFields, got %s", text)
	}
}

func TestHandleDeleteKCCResource(t *testing.T) {
	scheme := runtime.NewScheme()
	gvr := schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1beta1", Resource: "storagebuckets"}
	dynamicClient := dynamicfake.NewSimpleDynamicClientWithCustomListKinds(scheme, map[schema.GroupVersionResource]string{
		gvr: "StorageBucketList",
	})
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "storagebuckets", Kind: "StorageBucket", Namespaced: true},
				},
			},
		},
	}

	sc := &serverContext{
		dynamicClient:   dynamicClient,
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "storage.cnrm.cloud.google.com/v1beta1",
			"kind":       "StorageBucket",
			"metadata": map[string]interface{}{
				"name":      "delete-test-bucket",
				"namespace": "test-ns",
			},
		},
	}
	_, err := dynamicClient.Resource(gvr).Namespace("test-ns").Create(context.Background(), obj, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("failed to create fake resource: %v", err)
	}

	req := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"kind":      "StorageBucket",
				"namespace": "test-ns",
				"name":      "delete-test-bucket",
			},
		},
	}

	res, err := sc.handleDeleteKCCResource(context.Background(), req)
	if err != nil {
		t.Fatalf("handleDeleteKCCResource failed: %v", err)
	}

	if res.IsError {
		t.Fatalf("tool returned error: %v", res.Content[0].(mcp.TextContent).Text)
	}

	text := res.Content[0].(mcp.TextContent).Text
	if !strings.Contains(text, "Successfully deleted test-ns/delete-test-bucket (StorageBucket)") {
		t.Errorf("expected success message, got %s", text)
	}

	// Verify resource was deleted
	_, err = dynamicClient.Resource(gvr).Namespace("test-ns").Get(context.Background(), "delete-test-bucket", metav1.GetOptions{})
	if err == nil {
		t.Errorf("expected resource to be deleted, but it still exists")
	}
}

func TestHandleListKCCKinds(t *testing.T) {
	discoveryClient := &mockDiscovery{
		resources: []*metav1.APIResourceList{
			{
				GroupVersion: "storage.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "storagebuckets", Kind: "StorageBucket", Namespaced: true},
				},
			},
			{
				GroupVersion: "compute.cnrm.cloud.google.com/v1beta1",
				APIResources: []metav1.APIResource{
					{Name: "computeinstances", Kind: "ComputeInstance", Namespaced: true},
				},
			},
		},
	}
	sc := &serverContext{
		discoveryClient: discoveryClient,
		gvrCache:        make(map[string]schema.GroupVersionResource),
		gvkCache:        make(map[schema.GroupVersionKind]schema.GroupVersionResource),
		kindToGVRCache:  make(map[string]schema.GroupVersionResource),
	}

	req := mcp.CallToolRequest{}
	res, err := sc.handleListKCCKinds(context.Background(), req)
	if err != nil {
		t.Fatalf("handleListKCCKinds failed: %v", err)
	}

	if res.IsError {
		t.Fatalf("tool returned error: %v", res.Content[0].(mcp.TextContent).Text)
	}

	text := res.Content[0].(mcp.TextContent).Text
	if !strings.Contains(text, "StorageBucket") {
		t.Errorf("expected output to contain StorageBucket, got %s", text)
	}
	if !strings.Contains(text, "ComputeInstance") {
		t.Errorf("expected output to contain ComputeInstance, got %s", text)
	}
	
	// Verify sorting
	lines := strings.Split(text, "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines, got %d", len(lines))
	}
	// Available KCC Kinds:
	// - ComputeInstance (compute.cnrm.cloud.google.com)
	// - StorageBucket (storage.cnrm.cloud.google.com)
	if !strings.Contains(lines[1], "ComputeInstance") {
		t.Errorf("expected first kind to be ComputeInstance (alphabetical), got %s", lines[1])
	}
}
