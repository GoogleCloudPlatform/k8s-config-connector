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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/ratelimiter"
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

type mockUnreadableDeletableAdapter struct {
	mockAdapter
	isUnreadableButDeletable bool
}

func (a *mockUnreadableDeletableAdapter) IsUnreadableButDeletable(err error) bool {
	return a.isUnreadableButDeletable
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
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

	adapter := &mockUnreadableDeletableAdapter{
		mockAdapter: mockAdapter{
			findFound:     false,
			findErr:       &googleapi.Error{Code: 400, Message: "external metastore schema drift error"},
			deleteDeleted: true,
		},
		isUnreadableButDeletable: true,
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
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

func TestReconcile_UnreadableButDeletable_DeleteFailed(t *testing.T) {
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

	adapter := &mockUnreadableDeletableAdapter{
		mockAdapter: mockAdapter{
			findFound:     false,
			findErr:       &googleapi.Error{Code: 400, Message: "external metastore schema drift error"},
			deleteDeleted: false,
			deleteErr:     &googleapi.Error{Code: 403, Message: "permission denied on delete"},
		},
		isUnreadableButDeletable: true,
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-bq-table",
		},
	}

	_, err := r.Reconcile(ctx, req)
	if err == nil {
		t.Fatal("expected reconcile to fail with error, but got nil")
	}

	if !adapter.deleteCalled {
		t.Error("expected Adapter.Delete() to be called when Find() fails with HTTP 400 on BigQueryTable")
	}

	// Verify that the finalizer is still present
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(bqGVK)
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

func TestReconcile_UnreadableButDeletable_NormalUpdateFailed(t *testing.T) {
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

	// Set finalizer but NOT deletion timestamp
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	adapter := &mockUnreadableDeletableAdapter{
		mockAdapter: mockAdapter{
			findFound:     false,
			findErr:       &googleapi.Error{Code: 400, Message: "external metastore schema drift error"},
			deleteDeleted: true,
		},
		isUnreadableButDeletable: true,
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-bq-table",
		},
	}

	_, err := r.Reconcile(ctx, req)
	if err == nil {
		t.Fatal("expected reconcile to fail with error, but got nil")
	}

	if adapter.deleteCalled {
		t.Error("expected Adapter.Delete() NOT to be called during normal update (not deleting)")
	}
}

func TestReconcile_BackoffMaxDelay_Zero(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetAnnotations(map[string]string{
		k8s.BackoffMaxDelayInSecondsAnnotation: "0",
	})

	adapter := &mockAdapter{
		findFound: false,
		createErr: fmt.Errorf("create failed: API rate limited"),
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-resource",
		},
	}

	res, err := r.Reconcile(ctx, req)
	if err != nil {
		t.Fatalf("expected nil error when backoff-max-delay-in-seconds is 0, got: %v", err)
	}
	if res.Requeue || res.RequeueAfter != 0 {
		t.Fatalf("expected empty reconcile.Result{} (no requeue), got: %+v", res)
	}

	// Verify status condition Ready=False was recorded on the object
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(testGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}
	conditions, _, err := unstructured.NestedSlice(obj.Object, "status", "conditions")
	if err != nil || len(conditions) == 0 {
		t.Fatalf("expected status.conditions to be set, got err: %v, conditions: %v", err, conditions)
	}
	readyCond, ok := conditions[0].(map[string]interface{})
	if !ok || readyCond["status"] != string(corev1.ConditionFalse) {
		t.Fatalf("expected Ready condition False, got: %v", readyCond)
	}
}

func TestReconcile_BackoffMaxDelay_Positive(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetAnnotations(map[string]string{
		k8s.BackoffMaxDelayInSecondsAnnotation: "10",
	})

	adapter := &mockAdapter{
		findFound: false,
		createErr: fmt.Errorf("create failed: transient error"),
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-resource",
		},
	}

	expectedDelays := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		4 * time.Second,
		8 * time.Second,
		10 * time.Second, // capped at 10s
		10 * time.Second, // capped at 10s
	}

	for i, expectedDelay := range expectedDelays {
		res, err := r.Reconcile(ctx, req)
		if err != nil {
			t.Fatalf("step %d: expected nil error with custom backoff, got: %v", i, err)
		}
		if res.RequeueAfter != expectedDelay {
			t.Fatalf("step %d: expected RequeueAfter %v, got: %v", i, expectedDelay, res.RequeueAfter)
		}
	}
}

func TestReconcile_BackoffMaxDelay_Invalid(t *testing.T) {
	tests := []struct {
		name       string
		annotation string
	}{
		{
			name:       "invalid string",
			annotation: "invalid-val",
		},
		{
			name:       "negative number",
			annotation: "-10",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.TODO()
			scheme := runtime.NewScheme()
			_ = corev1.AddToScheme(scheme)
			_ = operatorv1beta1.AddToScheme(scheme)

			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(testGVK)
			u.SetName("test-resource")
			u.SetNamespace("test-ns")
			u.SetAnnotations(map[string]string{
				k8s.BackoffMaxDelayInSecondsAnnotation: tc.annotation,
			})

			adapter := &mockAdapter{
				findFound: false,
				createErr: fmt.Errorf("create failed: transient error"),
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
				rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
			}

			req := reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: "test-ns",
					Name:      "test-resource",
				},
			}

			// Invalid annotation should fall back to default behavior (returning error to controller-runtime)
			res, err := r.Reconcile(ctx, req)
			if err == nil {
				t.Fatal("expected error to be returned for invalid annotation fallback, got nil")
			}
			if res.RequeueAfter != 0 {
				t.Fatalf("expected RequeueAfter 0, got: %v", res.RequeueAfter)
			}
		})
	}
}

func TestReconcile_BackoffMaxDelay_SuccessReset(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetAnnotations(map[string]string{
		k8s.BackoffMaxDelayInSecondsAnnotation: "10",
	})

	adapter := &mockAdapter{
		findFound: false,
		createErr: fmt.Errorf("create failed: transient error"),
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-resource",
		},
	}

	// 1st failure: 1s
	res, err := r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 1*time.Second {
		t.Fatalf("expected 1s backoff, got err: %v, res: %+v", err, res)
	}

	// 2nd failure: 2s
	res, err = r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 2*time.Second {
		t.Fatalf("expected 2s backoff, got err: %v, res: %+v", err, res)
	}

	// Succeeded reconcile
	adapter.createErr = nil
	res, err = r.Reconcile(ctx, req)
	if err != nil {
		t.Fatalf("expected successful reconcile, got err: %v", err)
	}

	// Subsequent failure should reset back to base delay (1s)
	adapter.createErr = fmt.Errorf("transient error again")
	res, err = r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 1*time.Second {
		t.Fatalf("expected reset to 1s backoff after success, got err: %v, res: %+v", err, res)
	}
}

func TestReconcile_BackoffMaxDelay_DeletionBypass(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetAnnotations(map[string]string{
		k8s.BackoffMaxDelayInSecondsAnnotation: "0",
	})

	now := metav1.Now()
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	adapter := &mockAdapter{
		findFound: true,
		deleteErr: fmt.Errorf("delete failed: permission denied"),
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-resource",
		},
	}

	// Deletion errors should return error directly, bypassing the 0 halt and custom backoff
	res, err := r.Reconcile(ctx, req)
	if err == nil {
		t.Fatal("expected non-nil error on deletion failure, got nil")
	}
	if res.RequeueAfter != 0 {
		t.Fatalf("expected RequeueAfter 0, got: %v", res.RequeueAfter)
	}
}

func TestReconcile_BackoffMaxDelay_NotFoundClearsRateLimiter(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetAnnotations(map[string]string{
		k8s.BackoffMaxDelayInSecondsAnnotation: "10",
	})

	adapter := &mockAdapter{
		findFound: false,
		createErr: fmt.Errorf("creation error"),
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
		rateLimiter:     ratelimiter.NewDynamicRateLimiter(),
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Namespace: "test-ns",
			Name:      "test-resource",
		},
	}

	// 1st failure: 1s
	res, err := r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 1*time.Second {
		t.Fatalf("expected 1s backoff, got err: %v, res: %+v", err, res)
	}

	// 2nd failure: 2s
	res, err = r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 2*time.Second {
		t.Fatalf("expected 2s backoff, got err: %v, res: %+v", err, res)
	}

	// Remove finalizers and delete the object from k8s to trigger NotFound
	if err := k8sClient.Get(ctx, req.NamespacedName, u); err != nil {
		t.Fatalf("failed to get object from fake client: %v", err)
	}
	u.SetFinalizers(nil)
	if err := k8sClient.Update(ctx, u); err != nil {
		t.Fatalf("failed to update object to remove finalizers: %v", err)
	}
	if err := k8sClient.Delete(ctx, u); err != nil {
		t.Fatalf("failed to delete object from fake client: %v", err)
	}

	// Reconcile on deleted object returns NotFound -> should Forget in rateLimiter
	res, err = r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 0 {
		t.Fatalf("expected clean exit on NotFound, got err: %v, res: %+v", err, res)
	}

	// Re-create the object and fail: delay should start back at base delay (1s)
	uRecreated := &unstructured.Unstructured{}
	uRecreated.SetGroupVersionKind(testGVK)
	uRecreated.SetName("test-resource")
	uRecreated.SetNamespace("test-ns")
	uRecreated.SetAnnotations(map[string]string{
		k8s.BackoffMaxDelayInSecondsAnnotation: "10",
	})
	if err := k8sClient.Create(ctx, uRecreated); err != nil {
		t.Fatalf("failed to re-create object in fake client: %v", err)
	}
	res, err = r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 1*time.Second {
		t.Fatalf("expected reset to 1s backoff after NotFound cleanup, got err: %v, res: %+v", err, res)
	}
}

func TestCompareMBUR(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		computeMBUR func() (string, error)
		wantMatch   bool
		wantErr     bool
	}{
		{
			name:        "nil computeMBUR func",
			annotations: map[string]string{k8s.MutableUnreadableFieldsHashAnnotation: "somehash"},
			computeMBUR: nil,
			wantMatch:   false,
		},
		{
			name: "computeMBUR returns error",
			computeMBUR: func() (string, error) {
				return "", fmt.Errorf("computation failed")
			},
			wantErr: true,
		},
		{
			name:        "both empty",
			annotations: nil,
			computeMBUR: func() (string, error) {
				return "", nil
			},
			wantMatch: true,
		},
		{
			name:        "annotation exists but desired empty",
			annotations: map[string]string{k8s.MutableUnreadableFieldsHashAnnotation: "somehash"},
			computeMBUR: func() (string, error) {
				return "", nil
			},
			wantMatch: false,
		},
		{
			name:        "matching hashes",
			annotations: map[string]string{k8s.MutableUnreadableFieldsHashAnnotation: "hash123"},
			computeMBUR: func() (string, error) {
				return "hash123", nil
			},
			wantMatch: true,
		},
		{
			name:        "different hashes",
			annotations: map[string]string{k8s.MutableUnreadableFieldsHashAnnotation: "hash123"},
			computeMBUR: func() (string, error) {
				return "hash456", nil
			},
			wantMatch: false,
		},
		{
			name:        "desired has hash but annotation unset",
			annotations: nil,
			computeMBUR: func() (string, error) {
				return "hash123", nil
			},
			wantMatch: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := &unstructured.Unstructured{}
			u.SetAnnotations(tc.annotations)
			op := &operationBase{object: u}
			gotMatch, err := op.CompareMBUR(tc.computeMBUR)
			if (err != nil) != tc.wantErr {
				t.Fatalf("CompareMBUR() error = %v, wantErr = %v", err, tc.wantErr)
			}
			if gotMatch != tc.wantMatch {
				t.Errorf("CompareMBUR() = %v, want %v", gotMatch, tc.wantMatch)
			}
		})
	}
}

func TestSetMBUR(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)

	tests := []struct {
		name            string
		initAnnotations map[string]string
		computeMBUR     func() (string, error)
		wantAnnotation  string
		wantExist       bool
		wantErr         bool
	}{
		{
			name:            "nil computeMBUR",
			initAnnotations: map[string]string{"foo": "bar"},
			computeMBUR:     nil,
			wantExist:       false,
		},
		{
			name: "computeMBUR error",
			computeMBUR: func() (string, error) {
				return "", fmt.Errorf("hash error")
			},
			wantErr: true,
		},
		{
			name:            "set new hash",
			initAnnotations: nil,
			computeMBUR: func() (string, error) {
				return "hash123", nil
			},
			wantAnnotation: "hash123",
			wantExist:      true,
		},
		{
			name:            "update existing hash",
			initAnnotations: map[string]string{k8s.MutableUnreadableFieldsHashAnnotation: "oldhash"},
			computeMBUR: func() (string, error) {
				return "newhash", nil
			},
			wantAnnotation: "newhash",
			wantExist:      true,
		},
		{
			name:            "same hash is no-op",
			initAnnotations: map[string]string{k8s.MutableUnreadableFieldsHashAnnotation: "samehash"},
			computeMBUR: func() (string, error) {
				return "samehash", nil
			},
			wantAnnotation: "samehash",
			wantExist:      true,
		},
		{
			name:            "clear hash when desired is empty",
			initAnnotations: map[string]string{k8s.MutableUnreadableFieldsHashAnnotation: "hash123"},
			computeMBUR: func() (string, error) {
				return "", nil
			},
			wantExist: false,
		},
		{
			name:            "empty desired and empty annotation is no-op",
			initAnnotations: nil,
			computeMBUR: func() (string, error) {
				return "", nil
			},
			wantExist: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(testGVK)
			u.SetName("test-mbur")
			u.SetNamespace("test-ns")
			if tc.initAnnotations != nil {
				u.SetAnnotations(tc.initAnnotations)
			}
			k8sClient := fake.NewClientBuilder().
				WithScheme(scheme).
				WithRuntimeObjects(u).
				Build()

			op := &operationBase{
				client: k8sClient,
				object: u,
			}

			err := op.SetMBUR(ctx, tc.computeMBUR)
			if (err != nil) != tc.wantErr {
				t.Fatalf("SetMBUR() error = %v, wantErr = %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}

			gotObj := &unstructured.Unstructured{}
			gotObj.SetGroupVersionKind(testGVK)
			if err := k8sClient.Get(ctx, types.NamespacedName{Namespace: "test-ns", Name: "test-mbur"}, gotObj); err != nil {
				t.Fatalf("failed to fetch updated object: %v", err)
			}

			val, exists := gotObj.GetAnnotations()[k8s.MutableUnreadableFieldsHashAnnotation]
			if exists != tc.wantExist {
				t.Fatalf("annotation exists = %v, wantExist = %v", exists, tc.wantExist)
			}
			if tc.wantExist && val != tc.wantAnnotation {
				t.Errorf("annotation value = %q, want %q", val, tc.wantAnnotation)
			}
		})
	}
}
