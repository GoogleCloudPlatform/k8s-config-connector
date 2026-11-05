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

package k8s_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestContextWithUnstructured(t *testing.T) {
	gvk := runtimeschema.GroupVersionKind{
		Group:   "foo.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "FooResource",
	}
	nn := types.NamespacedName{
		Namespace: "test-ns",
		Name:      "test-foo",
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	u.SetNamespace(nn.Namespace)
	u.SetName(nn.Name)
	u.Object["spec"] = map[string]interface{}{"key": "value"}

	ctx := context.Background()

	// Initially context has no object
	if _, ok := k8s.UnstructuredFromContext(ctx, nn, gvk); ok {
		t.Fatalf("expected UnstructuredFromContext to return false for empty context")
	}

	ctx = k8s.ContextWithUnstructured(ctx, u)

	// Matching GVK and NamespacedName retrieves object
	got, ok := k8s.UnstructuredFromContext(ctx, nn, gvk)
	if !ok {
		t.Fatalf("expected UnstructuredFromContext to return true")
	}
	if got.GetName() != nn.Name || got.GetNamespace() != nn.Namespace {
		t.Errorf("got %v/%v, want %v/%v", got.GetNamespace(), got.GetName(), nn.Namespace, nn.Name)
	}

	// Mismatched NamespacedName returns false
	mismatchedNN := types.NamespacedName{Namespace: "test-ns", Name: "other-foo"}
	if _, ok := k8s.UnstructuredFromContext(ctx, mismatchedNN, gvk); ok {
		t.Fatalf("expected UnstructuredFromContext to return false for mismatched NamespacedName")
	}

	// Mismatched GVK returns false
	mismatchedGVK := runtimeschema.GroupVersionKind{Group: "foo.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BarResource"}
	if _, ok := k8s.UnstructuredFromContext(ctx, nn, mismatchedGVK); ok {
		t.Fatalf("expected UnstructuredFromContext to return false for mismatched GVK")
	}
}

func TestFetchUnstructuredFallback(t *testing.T) {
	gvk := runtimeschema.GroupVersionKind{
		Group:   "foo.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "FooResource",
	}
	nn := types.NamespacedName{
		Namespace: "test-ns",
		Name:      "test-foo",
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	u.SetNamespace(nn.Namespace)
	u.SetName(nn.Name)
	u.Object["spec"] = map[string]interface{}{"key": "value"}

	fakeClient := fake.NewClientBuilder().WithObjects(u).Build()

	// 1. Test fetching when object is in context (no client Get needed)
	ctxWithObj := k8s.ContextWithUnstructured(context.Background(), u)
	fetchedFromCtx := &unstructured.Unstructured{}
	fetchedFromCtx.SetGroupVersionKind(gvk)
	if err := k8s.FetchUnstructured(ctxWithObj, fakeClient, nn, fetchedFromCtx); err != nil {
		t.Fatalf("FetchUnstructured failed: %v", err)
	}
	if val, _, _ := unstructured.NestedString(fetchedFromCtx.Object, "spec", "key"); val != "value" {
		t.Errorf("got spec.key = %q, want %q", val, "value")
	}

	// 2. Test fetching when object is NOT in context (falls back to client Get)
	fetchedFromClient := &unstructured.Unstructured{}
	fetchedFromClient.SetGroupVersionKind(gvk)
	if err := k8s.FetchUnstructured(context.Background(), fakeClient, nn, fetchedFromClient); err != nil {
		t.Fatalf("FetchUnstructured fallback failed: %v", err)
	}
	if val, _, _ := unstructured.NestedString(fetchedFromClient.Object, "spec", "key"); val != "value" {
		t.Errorf("got spec.key = %q, want %q", val, "value")
	}
}

func TestFetchObject(t *testing.T) {
	cmGVK := corev1.SchemeGroupVersion.WithKind("ConfigMap")
	nn := types.NamespacedName{
		Namespace: "test-ns",
		Name:      "test-cm",
	}

	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      nn.Name,
				"namespace": nn.Namespace,
			},
			"data": map[string]interface{}{
				"key": "val-from-ctx",
			},
		},
	}

	clientCM := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-ns",
			Name:      "client-cm",
		},
		Data: map[string]string{
			"key": "val-from-client",
		},
	}
	clientU := &unstructured.Unstructured{}
	clientU.SetGroupVersionKind(cmGVK)
	clientU.SetNamespace("test-ns")
	clientU.SetName("client-u")
	clientU.Object["data"] = map[string]interface{}{"key": "val-from-client-u"}

	fakeClient := fake.NewClientBuilder().WithObjects(clientCM, clientU).Build()

	ctxWithObj := k8s.ContextWithUnstructured(context.Background(), u)

	// 1. Fetch typed struct from context (tests conversion via runtime.DefaultUnstructuredConverter.FromUnstructured)
	cm := &corev1.ConfigMap{}
	cm.SetGroupVersionKind(cmGVK)
	if err := k8s.FetchObject(ctxWithObj, fakeClient, nn, cm); err != nil {
		t.Fatalf("FetchObject failed for typed struct from context: %v", err)
	}
	if cm.Data["key"] != "val-from-ctx" {
		t.Errorf("got cm.Data[key] = %q, want %q", cm.Data["key"], "val-from-ctx")
	}

	// 2. Fetch *unstructured.Unstructured from context
	fetchedU := &unstructured.Unstructured{}
	fetchedU.SetGroupVersionKind(cmGVK)
	if err := k8s.FetchObject(ctxWithObj, fakeClient, nn, fetchedU); err != nil {
		t.Fatalf("FetchObject failed for unstructured from context: %v", err)
	}
	if val, _, _ := unstructured.NestedString(fetchedU.Object, "data", "key"); val != "val-from-ctx" {
		t.Errorf("got fetchedU data.key = %q, want %q", val, "val-from-ctx")
	}

	// 3. Fallback to client for typed struct when object not in context
	fallbackCMNN := types.NamespacedName{Namespace: "test-ns", Name: "client-cm"}
	cmFallback := &corev1.ConfigMap{}
	cmFallback.SetGroupVersionKind(cmGVK)
	if err := k8s.FetchObject(context.Background(), fakeClient, fallbackCMNN, cmFallback); err != nil {
		t.Fatalf("FetchObject fallback failed for typed struct: %v", err)
	}
	if cmFallback.Data["key"] != "val-from-client" {
		t.Errorf("got cmFallback.Data[key] = %q, want %q", cmFallback.Data["key"], "val-from-client")
	}

	// 4. Fallback to client for *unstructured.Unstructured when object not in context
	fallbackUNN := types.NamespacedName{Namespace: "test-ns", Name: "client-u"}
	uFallback := &unstructured.Unstructured{}
	uFallback.SetGroupVersionKind(cmGVK)
	if err := k8s.FetchObject(context.Background(), fakeClient, fallbackUNN, uFallback); err != nil {
		t.Fatalf("FetchObject fallback failed for unstructured: %v", err)
	}
	if val, _, _ := unstructured.NestedString(uFallback.Object, "data", "key"); val != "val-from-client-u" {
		t.Errorf("got uFallback data.key = %q, want %q", val, "val-from-client-u")
	}
}
