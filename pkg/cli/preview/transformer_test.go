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

package preview

import (
	"context"
	"testing"
	"time"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	k8s "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestObjectTransformer(t *testing.T) {
	// 1. Create a dummy object store
	s := &objects{
		store: make(map[types.NamespacedName]Object),
	}

	// 2. Define a transformer that modifies the object
	transformer := func(ctx context.Context, obj client.Object) error {
		annotations := obj.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}
		annotations["transformed"] = "true"
		obj.SetAnnotations(annotations)
		return nil
	}

	// 3. Define a dummy object
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(schema.GroupVersionKind{Group: "v1", Version: "v1", Kind: "Pod"})
	obj.SetName("test-pod")
	obj.SetNamespace("default")

	// 4. Call OnListObject with the transformer
	ctx := context.Background()
	var registrations []*eventHandlerRegistration
	// We need a dummy handler to verify it gets called with the transformed object,
	// but OnListObject calls handler.OnAdd(obj, false).
	// Let's create a fake handler.
	receivedObj := make(chan interface{}, 1)
	handler := toolscache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			receivedObj <- obj
		},
	}
	registrations = append(registrations, &eventHandlerRegistration{
		handler: handler,
	})

	if err := s.OnListObject(ctx, obj, true, registrations, []ObjectTransformer{transformer}); err != nil {
		t.Fatalf("OnListObject failed: %v", err)
	}

	// 5. Verify the object in the store is transformed
	key := types.NamespacedName{Name: "test-pod", Namespace: "default"}
	storedObj, ok := s.store[key]
	if !ok {
		t.Fatal("Object not found in store")
	}
	storedClientObj, ok := storedObj.(client.Object)
	if !ok {
		t.Fatal("Stored object is not a client.Object")
	}
	if storedClientObj.GetAnnotations()["transformed"] != "true" {
		t.Errorf("Expected annotation 'transformed' to be 'true', got %v", storedClientObj.GetAnnotations())
	}

	// 6. Verify the object passed to the handler is transformed
	select {
	case received := <-receivedObj:
		receivedClientObj, ok := received.(client.Object)
		if !ok {
			t.Fatal("Received object is not a client.Object")
		}
		if receivedClientObj.GetAnnotations()["transformed"] != "true" {
			t.Errorf("Expected handler object annotation 'transformed' to be 'true', got %v", receivedClientObj.GetAnnotations())
		}
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for handler to get object")
	}
}

func TestReconcilerOverrideTransformer(t *testing.T) {
	// 1. Define the override map
	overrides := map[schema.GroupKind]k8s.ReconcilerType{
		{Kind: "StorageBucket", Group: "storage.cnrm.cloud.google.com"}: k8s.ReconcilerType("direct"),
		{Kind: "SQLInstance", Group: "sql.cnrm.cloud.google.com"}:       k8s.ReconcilerType("dcl"),
	}

	// 2. Create the transformer
	transformer := newReconcilerOverrideTransformer(overrides)

	// 3. Create a ConfigConnectorContext object
	ccc := &corev1beta1.ConfigConnectorContext{}
	ccc.SetName("configconnectorcontext.core.cnrm.cloud.google.com")
	ccc.SetNamespace("default")

	// 4. Apply the transformer
	ctx := context.Background()
	if err := transformer(ctx, ccc); err != nil {
		t.Fatalf("Transformer failed: %v", err)
	}

	// 5. Verify the overrides
	if ccc.Spec.Experiments == nil || ccc.Spec.Experiments.ControllerOverrides == nil {
		t.Fatal("ControllerOverrides not initialized")
	}

	for gk, rt := range overrides {
		val, ok := ccc.Spec.Experiments.ControllerOverrides[gk.String()]
		if !ok {
			t.Errorf("Expected override for %s not found", gk)
		}
		if val != rt {
			t.Errorf("Expected override for %s to be %s, got %s", gk, rt, val)
		}
	}

	// 6. Verify non-CCC object is ignored
	pod := &unstructured.Unstructured{}
	pod.SetGroupVersionKind(schema.GroupVersionKind{Group: "v1", Version: "v1", Kind: "Pod"})
	if err := transformer(ctx, pod); err != nil {
		t.Fatalf("Transformer failed on Pod: %v", err)
	}
	// No panic implies success
}
