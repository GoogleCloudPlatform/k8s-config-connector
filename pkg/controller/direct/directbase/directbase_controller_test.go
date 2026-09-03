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

package directbase

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"google.golang.org/api/googleapi"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var testGVK = schema.GroupVersionKind{
	Group:   "test.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "TestKind",
}

type mockJitterGenerator struct{}

func (m *mockJitterGenerator) WatchJitteredTimeout() time.Duration {
	return 1 * time.Second
}

func (m *mockJitterGenerator) JitteredReenqueue(gvk schema.GroupVersionKind, obj metav1.Object) (time.Duration, error) {
	return 1 * time.Second, nil
}

type mockModel struct {
	adapter Adapter
	err     error
}

func (m *mockModel) AdapterForObject(ctx context.Context, op *AdapterForObjectOperation) (Adapter, error) {
	return m.adapter, m.err
}

func (m *mockModel) AdapterForURL(ctx context.Context, url string) (Adapter, error) {
	return m.adapter, m.err
}

type mockAdapter struct {
	findFound bool
	findErr   error

	deleteCalled  bool
	deleteDeleted bool
	deleteErr     error

	createCalled bool
	createErr    error

	updateCalled bool
	updateErr    error
}

func (a *mockAdapter) Find(ctx context.Context) (bool, error) {
	return a.findFound, a.findErr
}

func (a *mockAdapter) Delete(ctx context.Context, op *DeleteOperation) (bool, error) {
	a.deleteCalled = true
	return a.deleteDeleted, a.deleteErr
}

func (a *mockAdapter) Create(ctx context.Context, op *CreateOperation) error {
	a.createCalled = true
	return a.createErr
}

func (a *mockAdapter) Update(ctx context.Context, op *UpdateOperation) error {
	a.updateCalled = true
	return a.updateErr
}

func (a *mockAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func TestReconcile_FindGenericError_Reconcile(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")

	adapter := &mockAdapter{
		findFound: false,
		findErr:   fmt.Errorf("generic Find error"),
	}
	model := &mockModel{
		adapter: adapter,
	}

	k8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(u).
		WithRuntimeObjects(u).
		Build()

	r := &DirectReconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			k8sClient,
			record.NewFakeRecorder(100),
		),
		Client:          k8sClient,
		scheme:          scheme,
		gvk:             testGVK,
		model:           model,
		jitterGenerator: &mockJitterGenerator{},
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-resource",
		},
	}

	_, err := r.Reconcile(ctx, req)
	if err == nil {
		t.Fatal("expected reconcile to fail, but got no error")
	}

	if !strings.Contains(err.Error(), "generic Find error") {
		t.Fatalf("expected error to wrap 'generic Find error', got: %v", err)
	}

	if adapter.createCalled {
		t.Error("expected Adapter.Create() NOT to be called when Find() fails")
	}

	if adapter.updateCalled {
		t.Error("expected Adapter.Update() NOT to be called when Find() fails")
	}
}

func TestReconcile_FindGenericError_Delete(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")

	// Set finalizer and deletion timestamp
	now := metav1.Now()
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	adapter := &mockAdapter{
		findFound: false,
		findErr:   fmt.Errorf("generic Find error"),
	}
	model := &mockModel{
		adapter: adapter,
	}

	k8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(u).
		WithRuntimeObjects(u).
		Build()

	r := &DirectReconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			k8sClient,
			record.NewFakeRecorder(100),
		),
		Client:          k8sClient,
		scheme:          scheme,
		gvk:             testGVK,
		model:           model,
		jitterGenerator: &mockJitterGenerator{},
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-resource",
		},
	}

	_, err := r.Reconcile(ctx, req)
	if err == nil {
		t.Fatal("expected reconcile to fail, but got no error")
	}

	if !strings.Contains(err.Error(), "generic Find error") {
		t.Fatalf("expected error to wrap 'generic Find error', got: %v", err)
	}

	if adapter.deleteCalled {
		t.Error("expected Adapter.Delete() NOT to be called when Find() fails")
	}

	// Verify that the finalizer has NOT been removed from the resource (i.e. not orphaned)
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(testGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}

	finalizers := obj.GetFinalizers()
	hasFinalizer := false
	for _, f := range finalizers {
		if f == k8s.ControllerFinalizerName {
			hasFinalizer = true
			break
		}
	}
	if !hasFinalizer {
		t.Error("expected finalizer to still be present, but it was removed")
	}
}

func TestReconcile_FindUnresolvableDependency_Reconcile(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")

	adapter := &mockAdapter{
		findFound: false,
		findErr:   k8s.NewReferenceNotReadyError(testGVK, types.NamespacedName{Namespace: "test-ns", Name: "dep-resource"}),
	}
	model := &mockModel{
		adapter: adapter,
	}

	k8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(u).
		WithRuntimeObjects(u).
		Build()

	r := &DirectReconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			k8sClient,
			record.NewFakeRecorder(100),
		),
		Client:          k8sClient,
		scheme:          scheme,
		gvk:             testGVK,
		model:           model,
		jitterGenerator: &mockJitterGenerator{},
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-resource",
		},
	}

	res, err := r.Reconcile(ctx, req)
	if err != nil {
		t.Fatalf("expected reconcile to succeed without error (requesting requeue), but got: %v", err)
	}

	if !res.Requeue {
		t.Error("expected Requeue to be true for unresolvable dependencies during Find")
	}

	if adapter.createCalled {
		t.Error("expected Adapter.Create() NOT to be called when Find() fails with unresolvable dependencies")
	}

	if adapter.updateCalled {
		t.Error("expected Adapter.Update() NOT to be called when Find() fails with unresolvable dependencies")
	}

	// Verify the resource's Ready condition is set to False with DependencyNotReady
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(testGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}

	conditions, found, err := unstructured.NestedSlice(obj.Object, "status", "conditions")
	if err != nil || !found || len(conditions) == 0 {
		t.Fatalf("expected status conditions to be present, found: %v, err: %v", found, err)
	}

	readyCondition := conditions[0].(map[string]interface{})
	if readyCondition["status"] != string(corev1.ConditionFalse) {
		t.Errorf("expected condition status to be False, got: %v", readyCondition["status"])
	}

	if readyCondition["reason"] != k8s.DependencyNotReady {
		t.Errorf("expected condition reason to be %s, got: %v", k8s.DependencyNotReady, readyCondition["reason"])
	}
}

func TestReconcile_FindUnresolvableDependency_Delete(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")

	// Set finalizer and deletion timestamp
	now := metav1.Now()
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	adapter := &mockAdapter{
		findFound: false,
		findErr:   k8s.NewReferenceNotReadyError(testGVK, types.NamespacedName{Namespace: "test-ns", Name: "dep-resource"}),
	}
	model := &mockModel{
		adapter: adapter,
	}

	k8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(u).
		WithRuntimeObjects(u).
		Build()

	r := &DirectReconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			k8sClient,
			record.NewFakeRecorder(100),
		),
		Client:          k8sClient,
		scheme:          scheme,
		gvk:             testGVK,
		model:           model,
		jitterGenerator: &mockJitterGenerator{},
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-resource",
		},
	}

	res, err := r.Reconcile(ctx, req)
	if err != nil {
		t.Fatalf("expected reconcile to succeed without error (requesting requeue), but got: %v", err)
	}

	if !res.Requeue {
		t.Error("expected Requeue to be true for unresolvable dependencies during Find (Delete)")
	}

	if adapter.deleteCalled {
		t.Error("expected Adapter.Delete() NOT to be called when Find() fails with unresolvable dependencies during deletion")
	}

	// Verify that the finalizer has NOT been removed from the resource (i.e. not orphaned)
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(testGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}

	finalizers := obj.GetFinalizers()
	hasFinalizer := false
	for _, f := range finalizers {
		if f == k8s.ControllerFinalizerName {
			hasFinalizer = true
			break
		}
	}
	if !hasFinalizer {
		t.Error("expected finalizer to still be present, but it was removed")
	}

	// Verify the resource's Ready condition is set to False with DependencyNotReady
	conditions, found, err := unstructured.NestedSlice(obj.Object, "status", "conditions")
	if err != nil || !found || len(conditions) == 0 {
		t.Fatalf("expected status conditions to be present, found: %v, err: %v", found, err)
	}

	readyCondition := conditions[0].(map[string]interface{})
	if readyCondition["status"] != string(corev1.ConditionFalse) {
		t.Errorf("expected condition status to be False, got: %v", readyCondition["status"])
	}

	if readyCondition["reason"] != k8s.DependencyNotReady {
		t.Errorf("expected condition reason to be %s, got: %v", k8s.DependencyNotReady, readyCondition["reason"])
	}
}

func TestReconcile_UnreadableButDeletable_Delete(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	bqGVK := schema.GroupVersionKind{
		Group:   "bigquery.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "BigQueryTable",
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(bqGVK)
	u.SetName("test-bq-table")
	u.SetNamespace("test-ns")

	// Set finalizer and deletion timestamp
	now := metav1.Now()
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	adapter := &mockAdapter{
		findFound:     false,
		findErr:       &googleapi.Error{Code: 400, Message: "external metastore schema drift error"},
		deleteDeleted: true,
	}
	model := &mockModel{
		adapter: adapter,
	}

	k8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(u).
		WithRuntimeObjects(u).
		Build()

	r := &DirectReconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			k8sClient,
			record.NewFakeRecorder(100),
		),
		Client:          k8sClient,
		scheme:          scheme,
		gvk:             bqGVK,
		model:           model,
		jitterGenerator: &mockJitterGenerator{},
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-bq-table",
		},
	}

	_, err := r.Reconcile(ctx, req)
	if err != nil {
		t.Fatalf("expected reconcile to succeed without error, but got: %v", err)
	}

	if !adapter.deleteCalled {
		t.Error("expected Adapter.Delete() to be called when Find() fails with HTTP 400 on BigQueryTable")
	}

	// Verify that the finalizer has been removed from the resource (i.e. deleted successfully)
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(bqGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		if !apierrors.IsNotFound(err) {
			t.Fatalf("unexpected error retrieving object: %v", err)
		}
	} else {
		finalizers := obj.GetFinalizers()
		for _, f := range finalizers {
			if f == k8s.ControllerFinalizerName {
				t.Error("expected finalizer to be removed, but it is still present")
			}
		}
	}
}

func TestReconcile_NotUnreadableButDeletable_DeleteFailed(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	containerGVK := schema.GroupVersionKind{
		Group:   "container.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ContainerCluster",
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(containerGVK)
	u.SetName("test-container-cluster")
	u.SetNamespace("test-ns")

	// Set finalizer and deletion timestamp
	now := metav1.Now()
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	adapter := &mockAdapter{
		findFound:     false,
		findErr:       &googleapi.Error{Code: 400, Message: "some bad request error"},
		deleteDeleted: true,
	}
	model := &mockModel{
		adapter: adapter,
	}

	k8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithStatusSubresource(u).
		WithRuntimeObjects(u).
		Build()

	r := &DirectReconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(
			k8sClient,
			record.NewFakeRecorder(100),
		),
		Client:          k8sClient,
		scheme:          scheme,
		gvk:             containerGVK,
		model:           model,
		jitterGenerator: &mockJitterGenerator{},
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-container-cluster",
		},
	}

	_, err := r.Reconcile(ctx, req)
	if err == nil {
		t.Fatal("expected reconcile to fail with error, but got nil")
	}

	if adapter.deleteCalled {
		t.Error("expected Adapter.Delete() NOT to be called when Find() fails with HTTP 400 on ContainerCluster")
	}

	// Verify that the finalizer is still present
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(containerGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		t.Fatalf("unexpected error retrieving object: %v", err)
	}
	finalizers := obj.GetFinalizers()
	hasFinalizer := false
	for _, f := range finalizers {
		if f == k8s.ControllerFinalizerName {
			hasFinalizer = true
			break
		}
	}
	if !hasFinalizer {
		t.Error("expected finalizer to still be present, but it was removed")
	}
}
