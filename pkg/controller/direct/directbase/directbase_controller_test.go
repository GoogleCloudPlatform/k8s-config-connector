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

	artifactregistryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	containerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	firestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/pubsub/v1beta1"
	storagecontrolv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storagecontrol/v1alpha1"
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
	"sigs.k8s.io/controller-runtime/pkg/client"
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
	adapter               Adapter
	err                   error
	adapterForObjectCalls int
}

func (m *mockModel) AdapterForObject(ctx context.Context, op *AdapterForObjectOperation) (Adapter, error) {
	m.adapterForObjectCalls++
	return m.adapter, m.err
}

func (m *mockModel) AdapterForURL(ctx context.Context, url string) (Adapter, error) {
	return m.adapter, m.err
}

type mockAdapter struct {
	findCalls int
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
	a.findCalls++
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
		250 * time.Millisecond,
		500 * time.Millisecond,
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

	// 1st failure: 250ms
	res, err := r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 250*time.Millisecond {
		t.Fatalf("expected 250ms backoff, got err: %v, res: %+v", err, res)
	}

	// 2nd failure: 500ms
	res, err = r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 500*time.Millisecond {
		t.Fatalf("expected 500ms backoff, got err: %v, res: %+v", err, res)
	}

	// Succeeded reconcile
	adapter.createErr = nil
	res, err = r.Reconcile(ctx, req)
	if err != nil {
		t.Fatalf("expected successful reconcile, got err: %v", err)
	}

	// Subsequent failure should reset back to base delay (250ms)
	adapter.createErr = fmt.Errorf("transient error again")
	res, err = r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 250*time.Millisecond {
		t.Fatalf("expected reset to 250ms backoff after success, got err: %v, res: %+v", err, res)
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

	// 1st failure: 250ms
	res, err := r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 250*time.Millisecond {
		t.Fatalf("expected 250ms backoff, got err: %v, res: %+v", err, res)
	}

	// 2nd failure: 500ms
	res, err = r.Reconcile(ctx, req)
	if err != nil || res.RequeueAfter != 500*time.Millisecond {
		t.Fatalf("expected 500ms backoff, got err: %v, res: %+v", err, res)
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

	// Re-create the object and fail: delay should start back at base delay (250ms)
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
	if err != nil || res.RequeueAfter != 250*time.Millisecond {
		t.Fatalf("expected reset to 250ms backoff after NotFound cleanup, got err: %v, res: %+v", err, res)
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
			annotations: map[string]string{k8s.MutableButUnreadableFieldsHashAnnotation: "somehash"},
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
			annotations: map[string]string{k8s.MutableButUnreadableFieldsHashAnnotation: "somehash"},
			computeMBUR: func() (string, error) {
				return "", nil
			},
			wantMatch: false,
		},
		{
			name:        "matching hashes",
			annotations: map[string]string{k8s.MutableButUnreadableFieldsHashAnnotation: "hash123"},
			computeMBUR: func() (string, error) {
				return "hash123", nil
			},
			wantMatch: true,
		},
		{
			name:        "different hashes",
			annotations: map[string]string{k8s.MutableButUnreadableFieldsHashAnnotation: "hash123"},
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
			initAnnotations: map[string]string{k8s.MutableButUnreadableFieldsHashAnnotation: "oldhash"},
			computeMBUR: func() (string, error) {
				return "newhash", nil
			},
			wantAnnotation: "newhash",
			wantExist:      true,
		},
		{
			name:            "same hash is no-op",
			initAnnotations: map[string]string{k8s.MutableButUnreadableFieldsHashAnnotation: "samehash"},
			computeMBUR: func() (string, error) {
				return "samehash", nil
			},
			wantAnnotation: "samehash",
			wantExist:      true,
		},
		{
			name:            "clear hash when desired is empty",
			initAnnotations: map[string]string{k8s.MutableButUnreadableFieldsHashAnnotation: "hash123"},
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

			val, exists := gotObj.GetAnnotations()[k8s.MutableButUnreadableFieldsHashAnnotation]
			if exists != tc.wantExist {
				t.Fatalf("annotation exists = %v, wantExist = %v", exists, tc.wantExist)
			}
			if tc.wantExist && val != tc.wantAnnotation {
				t.Errorf("annotation value = %q, want %q", val, tc.wantAnnotation)
			}
		})
	}
}

type mockIAMModel struct {
	adapter         Adapter
	err             error
	iamAdapterCalls int
}

func (m *mockIAMModel) IAMAdapterForObject(_ context.Context, _ client.Reader, _ *unstructured.Unstructured, _ *IAMAdapterDeps) (Adapter, error) {
	m.iamAdapterCalls++
	return m.adapter, m.err
}

func (m *mockIAMModel) AdapterForObject(_ context.Context, _ *AdapterForObjectOperation) (Adapter, error) {
	return nil, fmt.Errorf("unexpected AdapterForObject call on IAMModel")
}

func (m *mockIAMModel) AdapterForURL(_ context.Context, _ string) (Adapter, error) {
	return nil, nil
}

type recordingDeleteAdapter struct {
	id        string
	findFound bool
	findCalls *int
	deletedID *string
}

func (a *recordingDeleteAdapter) Find(_ context.Context) (bool, error) {
	*a.findCalls++
	return a.findFound, nil
}

func (a *recordingDeleteAdapter) Delete(_ context.Context, _ *DeleteOperation) (bool, error) {
	*a.deletedID = a.id
	return true, nil
}

func (a *recordingDeleteAdapter) Create(_ context.Context, _ *CreateOperation) error {
	return nil
}

func (a *recordingDeleteAdapter) Update(_ context.Context, _ *UpdateOperation) error {
	return nil
}

func (a *recordingDeleteAdapter) Export(_ context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

type arRepoDeleteTestModel struct {
	findFound bool
	findCalls int
	deletedID string
}

func (m *arRepoDeleteTestModel) AdapterForObject(ctx context.Context, op *AdapterForObjectOperation) (Adapter, error) {
	obj := &artifactregistryv1beta1.ArtifactRegistryRepository{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.GetUnstructured().Object, &obj); err != nil {
		return nil, fmt.Errorf("converting unstructured: %w", err)
	}
	id, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}
	return &recordingDeleteAdapter{
		id:        id.String(),
		findFound: m.findFound,
		findCalls: &m.findCalls,
		deletedID: &m.deletedID,
	}, nil
}

func (m *arRepoDeleteTestModel) AdapterForURL(_ context.Context, _ string) (Adapter, error) {
	return nil, nil
}

type nodePoolDeleteTestModel struct {
	findFound bool
	findCalls int
	deletedID string
}

func (m *nodePoolDeleteTestModel) AdapterForObject(ctx context.Context, op *AdapterForObjectOperation) (Adapter, error) {
	obj := &containerv1beta1.ContainerNodePool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.GetUnstructured().Object, obj); err != nil {
		return nil, fmt.Errorf("converting unstructured: %w", err)
	}
	id, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}
	return &recordingDeleteAdapter{
		id:        id.String(),
		findFound: m.findFound,
		findCalls: &m.findCalls,
		deletedID: &m.deletedID,
	}, nil
}

func (m *nodePoolDeleteTestModel) AdapterForURL(_ context.Context, _ string) (Adapter, error) {
	return nil, nil
}

type pubSubTopicDeleteTestModel struct {
	findFound bool
	findCalls int
	deletedID string
}

func (m *pubSubTopicDeleteTestModel) AdapterForObject(ctx context.Context, op *AdapterForObjectOperation) (Adapter, error) {
	obj := &pubsubv1beta1.PubSubTopic{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.GetUnstructured().Object, obj); err != nil {
		return nil, fmt.Errorf("converting unstructured: %w", err)
	}
	id, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}
	return &recordingDeleteAdapter{
		id:        id.String(),
		findFound: m.findFound,
		findCalls: &m.findCalls,
		deletedID: &m.deletedID,
	}, nil
}

func (m *pubSubTopicDeleteTestModel) AdapterForURL(_ context.Context, _ string) (Adapter, error) {
	return nil, nil
}

type firestoreDBDeleteTestModel struct {
	findFound bool
	findCalls int
	deletedID string
}

func (m *firestoreDBDeleteTestModel) AdapterForObject(ctx context.Context, op *AdapterForObjectOperation) (Adapter, error) {
	obj := &firestorev1beta1.FirestoreDatabase{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.GetUnstructured().Object, obj); err != nil {
		return nil, fmt.Errorf("converting unstructured: %w", err)
	}
	id, err := obj.GetIdentity(ctx, op.Reader)
	if err != nil {
		return nil, err
	}
	return &recordingDeleteAdapter{
		id:        id.String(),
		findFound: m.findFound,
		findCalls: &m.findCalls,
		deletedID: &m.deletedID,
	}, nil
}

func (m *firestoreDBDeleteTestModel) AdapterForURL(_ context.Context, _ string) (Adapter, error) {
	return nil, nil
}

type storageFolderDeleteTestModel struct {
	findFound bool
	findCalls int
	deletedID string
}

func (m *storageFolderDeleteTestModel) AdapterForObject(ctx context.Context, op *AdapterForObjectOperation) (Adapter, error) {
	obj := &storagecontrolv1alpha1.StorageFolder{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(op.GetUnstructured().Object, obj); err != nil {
		return nil, fmt.Errorf("converting unstructured: %w", err)
	}
	id, err := storagecontrolv1alpha1.NewFolderIdentity(ctx, op.Reader, obj)
	if err != nil {
		return nil, err
	}
	return &recordingDeleteAdapter{
		id:        id.String(),
		findFound: m.findFound,
		findCalls: &m.findCalls,
		deletedID: &m.deletedID,
	}, nil
}

func (m *storageFolderDeleteTestModel) AdapterForURL(_ context.Context, _ string) (Adapter, error) {
	return nil, nil
}

func newTestDirectReconciler(t *testing.T, gvk schema.GroupVersionKind, model Model, objs ...runtime.Object) (*DirectReconciler, client.Client, *record.FakeRecorder) {
	t.Helper()
	scheme := runtime.NewScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		t.Fatalf("corev1.AddToScheme() failed: %v", err)
	}
	if err := operatorv1beta1.AddToScheme(scheme); err != nil {
		t.Fatalf("operatorv1beta1.AddToScheme() failed: %v", err)
	}
	builder := fake.NewClientBuilder().WithScheme(scheme)
	for _, obj := range objs {
		if cObj, ok := obj.(client.Object); ok {
			builder = builder.WithStatusSubresource(cObj)
		}
	}
	k8sClient := builder.WithRuntimeObjects(objs...).Build()
	recorder := record.NewFakeRecorder(100)
	reconciler := &DirectReconciler{
		LifecycleHandler: lifecyclehandler.NewLifecycleHandler(k8sClient, recorder),
		Client:           k8sClient,
		scheme:           scheme,
		gvk:              gvk,
		model:            model,
		jitterGenerator:  &mockJitterGenerator{},
		rateLimiter:      ratelimiter.NewDynamicRateLimiter(),
	}
	return reconciler, k8sClient, recorder
}

func TestReconcile_AbandonDeletion(t *testing.T) {
	depNN := types.NamespacedName{Namespace: "test-ns", Name: "missing-dep"}

	tests := []struct {
		name            string
		extraFinalizers []string
		adapterErr      error
		findErr         error
	}{
		{
			name:       "adapter_reference_not_found",
			adapterErr: k8s.NewReferenceNotFoundError(kmsv1beta1.KMSCryptoKeyGVK, depNN),
		},
		{
			name:       "adapter_reference_not_ready",
			adapterErr: k8s.NewReferenceNotReadyError(kmsv1beta1.KMSCryptoKeyGVK, depNN),
		},
		{
			name:       "adapter_generic_error",
			adapterErr: fmt.Errorf("missing project-id annotation"),
		},
		{
			name:    "find_permission_denied_403",
			findErr: &googleapi.Error{Code: 403, Message: "Caller does not have permission"},
		},
		{
			name:            "preserves_deleted_status_condition_with_observer_finalizer",
			extraFinalizers: []string{"example.com/observer"},
			findErr:         &googleapi.Error{Code: 403, Message: "Caller does not have permission"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := metav1.Now()
			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(testGVK)
			u.SetName("abandoned-res")
			u.SetNamespace("test-ns")
			u.SetDeletionTimestamp(&now)
			u.SetFinalizers(append([]string{k8s.ControllerFinalizerName}, tc.extraFinalizers...))
			u.SetAnnotations(map[string]string{
				k8s.DeletionPolicyAnnotation: k8s.DeletionPolicyAbandon,
			})

			adapter := &mockAdapter{
				findFound: true,
				findErr:   tc.findErr,
			}
			model := &mockModel{
				adapter: adapter,
				err:     tc.adapterErr,
			}
			r, k8sClient, recorder := newTestDirectReconciler(t, testGVK, model, u)
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: "abandoned-res"},
			}

			_, err := r.Reconcile(t.Context(), req)

			if err != nil {
				t.Fatalf("Reconcile(%s) error = %v, want nil", tc.name, err)
			}
			if model.adapterForObjectCalls != 0 {
				t.Errorf("Reconcile(%s) AdapterForObject calls = %d, want 0", tc.name, model.adapterForObjectCalls)
			}
			if adapter.findCalls != 0 {
				t.Errorf("Reconcile(%s) Find calls = %d, want 0", tc.name, adapter.findCalls)
			}
			if adapter.deleteCalled {
				t.Errorf("Reconcile(%s) Delete called = true, want false", tc.name)
			}

			gotObj := &unstructured.Unstructured{}
			gotObj.SetGroupVersionKind(testGVK)
			getErr := k8sClient.Get(t.Context(), req.NamespacedName, gotObj)
			if len(tc.extraFinalizers) == 0 {
				if !apierrors.IsNotFound(getErr) {
					t.Errorf("Reconcile(%s) k8sClient.Get() error = %v, want NotFound after finalizer removal", tc.name, getErr)
				}
			} else {
				if getErr != nil {
					t.Fatalf("Reconcile(%s) k8sClient.Get() unexpected error: %v", tc.name, getErr)
				}
				if k8s.HasFinalizer(gotObj, k8s.ControllerFinalizerName) {
					t.Errorf("Reconcile(%s) controller finalizer still present = true, want false", tc.name)
				}
				conditions, _, _ := unstructured.NestedSlice(gotObj.Object, "status", "conditions")
				if len(conditions) == 0 {
					t.Fatalf("Reconcile(%s) status.conditions = empty, want Deleted condition", tc.name)
				}
				cond, _ := conditions[0].(map[string]any)
				if cond["reason"] != k8s.Deleted {
					t.Errorf("Reconcile(%s) condition reason = %v, want %q", tc.name, cond["reason"], k8s.Deleted)
				}
			}

			select {
			case evt := <-recorder.Events:
				if !strings.Contains(evt, k8s.Deleted) {
					t.Errorf("Reconcile(%s) event = %q, want %q event", tc.name, evt, k8s.Deleted)
				}
			default:
				t.Errorf("Reconcile(%s) recorded 0 events, want %q event", tc.name, k8s.Deleted)
			}
		})
	}
}

func TestReconcile_AbandonDeletion_IAMModel(t *testing.T) {
	now := metav1.Now()
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("abandoned-iam")
	u.SetNamespace("test-ns")
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})
	u.SetAnnotations(map[string]string{
		k8s.DeletionPolicyAnnotation: k8s.DeletionPolicyAbandon,
	})

	adapter := &mockAdapter{
		findErr: &googleapi.Error{Code: 403, Message: "permission denied"},
	}
	model := &mockIAMModel{
		adapter: adapter,
		err:     fmt.Errorf("IAMAdapterForObject should not be called"),
	}
	r, k8sClient, _ := newTestDirectReconciler(t, testGVK, model, u)
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: "abandoned-iam"},
	}

	_, err := r.Reconcile(t.Context(), req)

	if err != nil {
		t.Fatalf("Reconcile(abandoned-iam) error = %v, want nil", err)
	}
	if model.iamAdapterCalls != 0 {
		t.Errorf("Reconcile(abandoned-iam) IAMAdapterForObject calls = %d, want 0", model.iamAdapterCalls)
	}
	if adapter.findCalls != 0 {
		t.Errorf("Reconcile(abandoned-iam) Find calls = %d, want 0", adapter.findCalls)
	}
	gotObj := &unstructured.Unstructured{}
	gotObj.SetGroupVersionKind(testGVK)
	if getErr := k8sClient.Get(t.Context(), req.NamespacedName, gotObj); !apierrors.IsNotFound(getErr) {
		t.Errorf("Reconcile(abandoned-iam) k8sClient.Get() error = %v, want NotFound", getErr)
	}
}

func TestReconcile_DeleteEarlyFinalizerGuards(t *testing.T) {
	tests := []struct {
		name               string
		finalizers         []string
		annotations        map[string]string
		wantStillPresent   bool
		wantControllerKept bool
	}{
		{
			name:       "no_controller_finalizer",
			finalizers: []string{"example.com/external-finalizer"},
			annotations: map[string]string{
				k8s.DeletionPolicyAnnotation: k8s.DeletionPolicyAbandon,
			},
			wantStillPresent:   true,
			wantControllerKept: false,
		},
		{
			name: "deletion_defender_blocks_even_with_abandon",
			finalizers: []string{
				k8s.ControllerFinalizerName,
				k8s.DeletionDefenderFinalizerName,
			},
			annotations: map[string]string{
				k8s.DeletionPolicyAnnotation: k8s.DeletionPolicyAbandon,
			},
			wantStillPresent:   true,
			wantControllerKept: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := metav1.Now()
			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(testGVK)
			u.SetName("guarded-res")
			u.SetNamespace("test-ns")
			u.SetDeletionTimestamp(&now)
			u.SetFinalizers(tc.finalizers)
			u.SetAnnotations(tc.annotations)

			adapter := &mockAdapter{
				findErr: fmt.Errorf("Find should not be called"),
			}
			model := &mockModel{
				adapter: adapter,
				err:     fmt.Errorf("AdapterForObject should not be called"),
			}
			r, k8sClient, _ := newTestDirectReconciler(t, testGVK, model, u)
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: "guarded-res"},
			}

			_, err := r.Reconcile(t.Context(), req)

			if err != nil {
				t.Fatalf("Reconcile(%s) error = %v, want nil", tc.name, err)
			}
			if model.adapterForObjectCalls != 0 {
				t.Errorf("Reconcile(%s) AdapterForObject calls = %d, want 0", tc.name, model.adapterForObjectCalls)
			}
			if adapter.findCalls != 0 {
				t.Errorf("Reconcile(%s) Find calls = %d, want 0", tc.name, adapter.findCalls)
			}
			gotObj := &unstructured.Unstructured{}
			gotObj.SetGroupVersionKind(testGVK)
			getErr := k8sClient.Get(t.Context(), req.NamespacedName, gotObj)
			if tc.wantStillPresent && getErr != nil {
				t.Fatalf("Reconcile(%s) k8sClient.Get() error = %v, want object present", tc.name, getErr)
			}
			if got := k8s.HasFinalizer(gotObj, k8s.ControllerFinalizerName); got != tc.wantControllerKept {
				t.Errorf("Reconcile(%s) HasFinalizer(ControllerFinalizerName) = %v, want %v", tc.name, got, tc.wantControllerKept)
			}
		})
	}
}

func TestReconcile_DeleteRecoversParentRefFromStatus(t *testing.T) {
	npGVK := containerv1beta1.ContainerNodePoolGVK
	dbGVK := firestorev1beta1.FirestoreDatabaseGVK
	folderSGVK := storagecontrolv1alpha1.StorageFolderGVK

	unreadyCluster := &unstructured.Unstructured{}
	unreadyCluster.SetGroupVersionKind(containerv1beta1.ContainerClusterGVK)
	unreadyCluster.SetNamespace("test-ns")
	unreadyCluster.SetName("unready-cluster")

	tests := []struct {
		name          string
		gvk           schema.GroupVersionKind
		resName       string
		spec          map[string]any
		status        map[string]any
		extraObjs     []runtime.Object
		findFound     bool
		wantDeletedID string
	}{
		{
			name:    "nodepool_missing_cluster_recovered_from_external_ref",
			gvk:     npGVK,
			resName: "pool-a",
			spec: map[string]any{
				"location": "us-central1",
				"clusterRef": map[string]any{
					"name": "deleted-cluster",
				},
			},
			status: map[string]any{
				"externalRef": "projects/test-proj/locations/us-central1/clusters/cluster-a/nodePools/pool-a",
			},
			findFound:     true,
			wantDeletedID: "projects/test-proj/locations/us-central1/clusters/cluster-a/nodePools/pool-a",
		},
		{
			name:    "nodepool_missing_cluster_recovered_from_zonal_self_link_already_gone_in_gcp",
			gvk:     npGVK,
			resName: "pool-a",
			spec: map[string]any{
				"clusterRef": map[string]any{
					"name": "deleted-cluster",
				},
			},
			status: map[string]any{
				"selfLink": "https://container.googleapis.com/v1/projects/test-proj/zones/us-central1-a/clusters/cluster-a/nodePools/pool-a",
			},
			findFound:     false,
			wantDeletedID: "",
		},
		{
			name:    "nodepool_same_name_kms_ref_and_cluster_ref_both_missing",
			gvk:     npGVK,
			resName: "pool-a",
			spec: map[string]any{
				"location": "us-central1",
				"clusterRef": map[string]any{
					"name": "shared-name",
				},
				"nodeConfig": map[string]any{
					"bootDiskKMSCryptoKeyRef": map[string]any{
						"name": "shared-name",
					},
				},
			},
			status: map[string]any{
				"externalRef": "projects/test-proj/locations/us-central1/clusters/shared-name/nodePools/pool-a",
			},
			findFound:     true,
			wantDeletedID: "projects/test-proj/locations/us-central1/clusters/shared-name/nodePools/pool-a",
		},
		{
			name:    "nodepool_unready_cluster_recovered_from_external_ref",
			gvk:     npGVK,
			resName: "pool-a",
			spec: map[string]any{
				"location": "us-central1",
				"clusterRef": map[string]any{
					"name": "unready-cluster",
				},
			},
			status: map[string]any{
				"externalRef": "projects/test-proj/locations/us-central1/clusters/cluster-a/nodePools/pool-a",
			},
			extraObjs:     []runtime.Object{unreadyCluster},
			findFound:     true,
			wantDeletedID: "projects/test-proj/locations/us-central1/clusters/cluster-a/nodePools/pool-a",
		},
		{
			name:    "database_missing_project_ref_recovered_from_external_ref",
			gvk:     dbGVK,
			resName: "db-a",
			spec: map[string]any{
				"locationID": "nam5",
				"projectRef": map[string]any{
					"name": "deleted-project",
				},
			},
			status: map[string]any{
				"externalRef": "projects/real-proj-id/databases/db-a",
			},
			findFound:     true,
			wantDeletedID: "projects/real-proj-id/databases/db-a",
		},
		{
			name:    "storage_folder_missing_both_project_ref_and_bucket_ref_recovered_from_external_ref",
			gvk:     folderSGVK,
			resName: "folder-a",
			spec: map[string]any{
				"projectRef": map[string]any{
					"name": "deleted-project",
				},
				"storagebucketRef": map[string]any{
					"name": "deleted-bucket",
				},
			},
			status: map[string]any{
				"externalRef": "projects/real-proj-id/buckets/real-bucket/folders/folder-a",
			},
			findFound:     true,
			wantDeletedID: "projects/real-proj-id/buckets/real-bucket/folders/folder-a",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := metav1.Now()
			u := &unstructured.Unstructured{
				Object: map[string]any{
					"spec":   tc.spec,
					"status": tc.status,
				},
			}
			u.SetGroupVersionKind(tc.gvk)
			u.SetName(tc.resName)
			u.SetNamespace("test-ns")
			u.SetDeletionTimestamp(&now)
			u.SetFinalizers([]string{k8s.ControllerFinalizerName})

			var model Model
			var findCalls *int
			var deletedID *string
			switch tc.gvk {
			case npGVK:
				m := &nodePoolDeleteTestModel{findFound: tc.findFound}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			case folderSGVK:
				m := &storageFolderDeleteTestModel{findFound: tc.findFound}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			default:
				m := &firestoreDBDeleteTestModel{findFound: tc.findFound}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			}

			objs := append([]runtime.Object{u}, tc.extraObjs...)
			r, k8sClient, _ := newTestDirectReconciler(t, tc.gvk, model, objs...)
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: tc.resName},
			}

			_, err := r.Reconcile(t.Context(), req)

			if err != nil {
				t.Fatalf("Reconcile(%s) error = %v, want nil", tc.name, err)
			}
			if *findCalls != 1 {
				t.Errorf("Reconcile(%s) Find calls = %d, want 1", tc.name, *findCalls)
			}
			if *deletedID != tc.wantDeletedID {
				t.Errorf("Reconcile(%s) deletedID = %q, want %q", tc.name, *deletedID, tc.wantDeletedID)
			}
			gotObj := &unstructured.Unstructured{}
			gotObj.SetGroupVersionKind(tc.gvk)
			if getErr := k8sClient.Get(t.Context(), req.NamespacedName, gotObj); !apierrors.IsNotFound(getErr) {
				t.Errorf("Reconcile(%s) k8sClient.Get() error = %v, want NotFound", tc.name, getErr)
			}
		})
	}
}

func TestReconcile_UncreatedChildWithoutFinalizerDeletesCleanly(t *testing.T) {
	npGVK := containerv1beta1.ContainerNodePoolGVK
	dbGVK := firestorev1beta1.FirestoreDatabaseGVK

	tests := []struct {
		name    string
		gvk     schema.GroupVersionKind
		resName string
		spec    map[string]any
	}{
		{
			name:    "uncreated_nodepool_never_acquired_controller_finalizer",
			gvk:     npGVK,
			resName: "uncreated-pool",
			spec: map[string]any{
				"location": "us-central1",
				"clusterRef": map[string]any{
					"name": "deleted-cluster",
				},
			},
		},
		{
			name:    "uncreated_database_never_acquired_controller_finalizer",
			gvk:     dbGVK,
			resName: "uncreated-db",
			spec: map[string]any{
				"locationID": "nam5",
				"projectRef": map[string]any{
					"name": "deleted-project",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := &unstructured.Unstructured{
				Object: map[string]any{
					"spec": tc.spec,
				},
			}
			u.SetGroupVersionKind(tc.gvk)
			u.SetName(tc.resName)
			u.SetNamespace("test-ns")

			var model Model
			var findCalls *int
			var deletedID *string
			if tc.gvk == npGVK {
				m := &nodePoolDeleteTestModel{findFound: true}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			} else {
				m := &firestoreDBDeleteTestModel{findFound: true}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			}

			r, k8sClient, _ := newTestDirectReconciler(t, tc.gvk, model, u)
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: tc.resName},
			}

			// Active reconcile fails before Find(), so controller finalizer is never added;
			// simulate an external observer finalizer so DeletionTimestamp can be tested on the live object.
			_, _ = r.Reconcile(t.Context(), req)
			activeObj := &unstructured.Unstructured{}
			activeObj.SetGroupVersionKind(tc.gvk)
			if err := k8sClient.Get(t.Context(), req.NamespacedName, activeObj); err != nil {
				t.Fatalf("k8sClient.Get(%s) error = %v", tc.name, err)
			}
			if k8s.HasFinalizer(activeObj, k8s.ControllerFinalizerName) {
				t.Fatalf("Reconcile(%s) unexpectedly added controller finalizer before parent resolved", tc.name)
			}
			activeObj.SetFinalizers([]string{"example.com/observer"})
			if err := k8sClient.Update(t.Context(), activeObj); err != nil {
				t.Fatalf("k8sClient.Update(%s) error = %v", tc.name, err)
			}
			if err := k8sClient.Delete(t.Context(), activeObj); err != nil {
				t.Fatalf("k8sClient.Delete(%s) error = %v", tc.name, err)
			}

			_, err := r.Reconcile(t.Context(), req)

			if err != nil {
				t.Fatalf("Reconcile(%s) on deleting uncreated object error = %v, want nil", tc.name, err)
			}
			if *findCalls != 0 {
				t.Errorf("Reconcile(%s) Find calls = %d, want 0", tc.name, *findCalls)
			}
			if *deletedID != "" {
				t.Errorf("Reconcile(%s) deletedID = %q, want empty string", tc.name, *deletedID)
			}
		})
	}
}

func TestReconcile_DeleteSafetyGuard_InvalidIdentityPreservesFinalizer(t *testing.T) {
	npGVK := containerv1beta1.ContainerNodePoolGVK
	arGVK := artifactregistryv1beta1.ArtifactRegistryRepositoryGVK
	dbGVK := firestorev1beta1.FirestoreDatabaseGVK

	tests := []struct {
		name    string
		gvk     schema.GroupVersionKind
		resName string
		spec    map[string]any
		status  map[string]any
	}{
		{
			name:    "missing_project_annotation_and_no_status_identity",
			gvk:     arGVK,
			resName: "misconfigured-repo",
			spec: map[string]any{
				"location": "us-central1",
				"format":   "DOCKER",
			},
		},
		{
			name:    "mismatched_status_external_ref_when_parent_deleted",
			gvk:     npGVK,
			resName: "pool-a",
			spec: map[string]any{
				"location": "us-central1",
				"clusterRef": map[string]any{
					"name": "deleted-cluster",
				},
			},
			status: map[string]any{
				"externalRef": "projects/test-proj/locations/us-central1/clusters/cluster-a/nodePools/different-pool",
			},
		},
		{
			name:    "mismatched_project_ref_external_ref_when_project_deleted",
			gvk:     dbGVK,
			resName: "db-a",
			spec: map[string]any{
				"locationID": "nam5",
				"projectRef": map[string]any{
					"name": "deleted-project",
				},
			},
			status: map[string]any{
				"externalRef": "projects/real-proj-id/databases/different-db",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := metav1.Now()
			u := &unstructured.Unstructured{
				Object: map[string]any{
					"spec": tc.spec,
				},
			}
			if tc.status != nil {
				u.Object["status"] = tc.status
			}
			u.SetGroupVersionKind(tc.gvk)
			u.SetName(tc.resName)
			u.SetNamespace("test-ns")
			u.SetDeletionTimestamp(&now)
			u.SetFinalizers([]string{k8s.ControllerFinalizerName})

			var model Model
			var findCalls *int
			var deletedID *string
			switch tc.gvk {
			case npGVK:
				m := &nodePoolDeleteTestModel{findFound: true}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			case dbGVK:
				m := &firestoreDBDeleteTestModel{findFound: true}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			default:
				m := &arRepoDeleteTestModel{findFound: true}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			}

			r, k8sClient, _ := newTestDirectReconciler(t, tc.gvk, model, u)
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: tc.resName},
			}

			_, err := r.Reconcile(t.Context(), req)

			if err == nil {
				t.Fatalf("Reconcile(%s) error = nil, want non-nil DeleteFailed error", tc.name)
			}
			if *findCalls != 0 {
				t.Errorf("Reconcile(%s) Find calls = %d, want 0", tc.name, *findCalls)
			}
			if *deletedID != "" {
				t.Errorf("Reconcile(%s) deletedID = %q, want empty string", tc.name, *deletedID)
			}
			gotObj := &unstructured.Unstructured{}
			gotObj.SetGroupVersionKind(tc.gvk)
			if getErr := k8sClient.Get(t.Context(), req.NamespacedName, gotObj); getErr != nil {
				t.Fatalf("Reconcile(%s) k8sClient.Get() error = %v, want object preserved", tc.name, getErr)
			}
			if !k8s.HasFinalizer(gotObj, k8s.ControllerFinalizerName) {
				t.Errorf("Reconcile(%s) HasFinalizer(ControllerFinalizerName) = false, want true", tc.name)
			}
			conditions, _, _ := unstructured.NestedSlice(gotObj.Object, "status", "conditions")
			if len(conditions) == 0 {
				t.Fatalf("Reconcile(%s) status.conditions = empty, want DeleteFailed condition", tc.name)
			}
			cond, _ := conditions[0].(map[string]any)
			if cond["reason"] != k8s.DeleteFailed {
				t.Errorf("Reconcile(%s) condition reason = %v, want %q", tc.name, cond["reason"], k8s.DeleteFailed)
			}
		})
	}
}

func TestReconcile_DeleteUncreatedWaitsForUnreadyParentRef(t *testing.T) {
	npGVK := containerv1beta1.ContainerNodePoolGVK

	unreadyCluster := &unstructured.Unstructured{}
	unreadyCluster.SetGroupVersionKind(containerv1beta1.ContainerClusterGVK)
	unreadyCluster.SetNamespace("test-ns")
	unreadyCluster.SetName("provisioning-cluster")

	now := metav1.Now()
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"spec": map[string]any{
				"location": "us-central1",
				"clusterRef": map[string]any{
					"name": "provisioning-cluster",
				},
			},
		},
	}
	u.SetGroupVersionKind(npGVK)
	u.SetName("uncreated-pool")
	u.SetNamespace("test-ns")
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	model := &nodePoolDeleteTestModel{findFound: true}
	r, k8sClient, _ := newTestDirectReconciler(t, npGVK, model, u, unreadyCluster)
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: "uncreated-pool"},
	}

	res, err := r.Reconcile(t.Context(), req)

	if err != nil {
		t.Fatalf("Reconcile(uncreated-pool) error = %v, want nil", err)
	}
	if !res.Requeue {
		t.Errorf("Reconcile(uncreated-pool) res.Requeue = false, want true")
	}
	if model.findCalls != 0 {
		t.Errorf("Reconcile(uncreated-pool) Find calls = %d, want 0", model.findCalls)
	}
	if model.deletedID != "" {
		t.Errorf("Reconcile(uncreated-pool) deletedID = %q, want empty string", model.deletedID)
	}
	gotObj := &unstructured.Unstructured{}
	gotObj.SetGroupVersionKind(npGVK)
	if getErr := k8sClient.Get(t.Context(), req.NamespacedName, gotObj); getErr != nil {
		t.Fatalf("Reconcile(uncreated-pool) k8sClient.Get() error = %v, want object preserved", getErr)
	}
	if !k8s.HasFinalizer(gotObj, k8s.ControllerFinalizerName) {
		t.Errorf("Reconcile(uncreated-pool) HasFinalizer(ControllerFinalizerName) = false, want true")
	}
	conditions, _, _ := unstructured.NestedSlice(gotObj.Object, "status", "conditions")
	if len(conditions) == 0 {
		t.Fatalf("Reconcile(uncreated-pool) status.conditions = empty, want DependencyNotReady condition")
	}
	cond, _ := conditions[0].(map[string]any)
	if cond["reason"] != k8s.DependencyNotReady {
		t.Errorf("Reconcile(uncreated-pool) condition reason = %v, want %q", cond["reason"], k8s.DependencyNotReady)
	}
}

func TestReconcile_DeleteSafetyGuard_NestedNonParentRefPreservesFinalizer(t *testing.T) {
	depNN := types.NamespacedName{Namespace: "test-ns", Name: "deleted-key"}
	now := metav1.Now()
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"spec": map[string]any{
				"projectRef": map[string]any{
					"external": "projects/live-proj",
				},
				"encryption": map[string]any{
					"kmsKeyRef": map[string]any{
						"name": "deleted-key",
					},
				},
			},
		},
	}
	u.SetGroupVersionKind(testGVK)
	u.SetName("custom-resolver-res")
	u.SetNamespace("test-ns")
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	adapter := &mockAdapter{findFound: true}
	model := &mockModel{
		adapter: adapter,
		err:     k8s.NewReferenceNotFoundError(kmsv1beta1.KMSCryptoKeyGVK, depNN),
	}
	r, k8sClient, _ := newTestDirectReconciler(t, testGVK, model, u)
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: "custom-resolver-res"},
	}

	res, err := r.Reconcile(t.Context(), req)

	if err != nil {
		t.Fatalf("Reconcile(custom-resolver-res) error = %v, want nil", err)
	}
	if !res.Requeue {
		t.Errorf("Reconcile(custom-resolver-res) res.Requeue = false, want true")
	}
	if adapter.deleteCalled {
		t.Errorf("Reconcile(custom-resolver-res) Delete called = true, want false")
	}
	gotObj := &unstructured.Unstructured{}
	gotObj.SetGroupVersionKind(testGVK)
	if getErr := k8sClient.Get(t.Context(), req.NamespacedName, gotObj); getErr != nil {
		t.Fatalf("Reconcile(custom-resolver-res) k8sClient.Get() error = %v, want object preserved", getErr)
	}
	if !k8s.HasFinalizer(gotObj, k8s.ControllerFinalizerName) {
		t.Errorf("Reconcile(custom-resolver-res) HasFinalizer(ControllerFinalizerName) = false, want true")
	}
	conditions, _, _ := unstructured.NestedSlice(gotObj.Object, "status", "conditions")
	if len(conditions) == 0 {
		t.Fatalf("Reconcile(custom-resolver-res) status.conditions = empty, want DependencyNotFound condition")
	}
	cond, _ := conditions[0].(map[string]any)
	if cond["reason"] != k8s.DependencyNotFound {
		t.Errorf("Reconcile(custom-resolver-res) condition reason = %v, want %q", cond["reason"], k8s.DependencyNotFound)
	}
}

func TestReconcile_DeleteSafetyGuard_CreatedWithoutStatusIdentityPreservesFinalizerAcrossRequeues(t *testing.T) {
	topicGVK := pubsubv1beta1.PubSubTopicGVK
	npGVK := containerv1beta1.ContainerNodePoolGVK

	tests := []struct {
		name    string
		gvk     schema.GroupVersionKind
		resName string
		spec    map[string]any
		status  map[string]any
	}{
		{
			name:    "created_pubsub_topic_ready_up_to_date_with_abandoned_project",
			gvk:     topicGVK,
			resName: "live-topic",
			spec: map[string]any{
				"projectRef": map[string]any{
					"name": "abandoned-project",
				},
			},
			status: map[string]any{
				"observedGeneration": int64(3),
				"conditions": []any{
					map[string]any{
						"type":   "Ready",
						"status": "True",
						"reason": k8s.UpToDate,
					},
				},
			},
		},
		{
			name:    "created_pubsub_topic_already_in_dependency_not_found_before_delete",
			gvk:     topicGVK,
			resName: "live-topic",
			spec: map[string]any{
				"projectRef": map[string]any{
					"name": "abandoned-project",
				},
			},
			status: map[string]any{
				"observedGeneration": int64(3),
				"conditions": []any{
					map[string]any{
						"type":   "Ready",
						"status": "False",
						"reason": k8s.DependencyNotFound,
					},
				},
			},
		},
		{
			name:    "nodepool_without_status_external_ref_and_abandoned_cluster",
			gvk:     npGVK,
			resName: "live-pool",
			spec: map[string]any{
				"location": "us-central1",
				"clusterRef": map[string]any{
					"name": "abandoned-cluster",
				},
			},
			status: map[string]any{
				"observedGeneration": int64(2),
				"conditions": []any{
					map[string]any{
						"type":   "Ready",
						"status": "True",
						"reason": k8s.UpToDate,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := metav1.Now()
			u := &unstructured.Unstructured{
				Object: map[string]any{
					"spec":   tc.spec,
					"status": tc.status,
				},
			}
			u.SetGroupVersionKind(tc.gvk)
			u.SetName(tc.resName)
			u.SetNamespace("test-ns")
			u.SetDeletionTimestamp(&now)
			u.SetFinalizers([]string{k8s.ControllerFinalizerName})

			var model Model
			var findCalls *int
			var deletedID *string
			if tc.gvk == npGVK {
				m := &nodePoolDeleteTestModel{findFound: true}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			} else {
				m := &pubSubTopicDeleteTestModel{findFound: true}
				model = m
				findCalls = &m.findCalls
				deletedID = &m.deletedID
			}

			r, k8sClient, _ := newTestDirectReconciler(t, tc.gvk, model, u)
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: tc.resName},
			}

			// Reconcile twice in a row to verify Ready condition overwrite on pass #1
			// never erases protection on pass #2.
			for pass := 1; pass <= 2; pass++ {
				res, err := r.Reconcile(t.Context(), req)

				if err != nil {
					t.Fatalf("Reconcile(%s) pass %d error = %v, want nil", tc.name, pass, err)
				}
				if !res.Requeue {
					t.Errorf("Reconcile(%s) pass %d res.Requeue = false, want true", tc.name, pass)
				}
				if *findCalls != 0 {
					t.Errorf("Reconcile(%s) pass %d Find calls = %d, want 0", tc.name, pass, *findCalls)
				}
				if *deletedID != "" {
					t.Errorf("Reconcile(%s) pass %d deletedID = %q, want empty string", tc.name, pass, *deletedID)
				}
				gotObj := &unstructured.Unstructured{}
				gotObj.SetGroupVersionKind(tc.gvk)
				if getErr := k8sClient.Get(t.Context(), req.NamespacedName, gotObj); getErr != nil {
					t.Fatalf("Reconcile(%s) pass %d k8sClient.Get() error = %v, want live resource preserved", tc.name, pass, getErr)
				}
				if !k8s.HasFinalizer(gotObj, k8s.ControllerFinalizerName) {
					t.Fatalf("Reconcile(%s) pass %d HasFinalizer(ControllerFinalizerName) = false, want true", tc.name, pass)
				}
				conditions, _, _ := unstructured.NestedSlice(gotObj.Object, "status", "conditions")
				if len(conditions) == 0 {
					t.Fatalf("Reconcile(%s) pass %d status.conditions = empty, want DependencyNotFound condition", tc.name, pass)
				}
				cond, _ := conditions[0].(map[string]any)
				if cond["reason"] != k8s.DependencyNotFound {
					t.Errorf("Reconcile(%s) pass %d condition reason = %v, want %q", tc.name, pass, cond["reason"], k8s.DependencyNotFound)
				}
			}

			// Explicit deletion-policy: abandon cleanly releases the finalizer without GCP calls.
			latest := &unstructured.Unstructured{}
			latest.SetGroupVersionKind(tc.gvk)
			if err := k8sClient.Get(t.Context(), req.NamespacedName, latest); err != nil {
				t.Fatalf("k8sClient.Get(%s) error = %v", tc.name, err)
			}
			latest.SetAnnotations(map[string]string{
				k8s.DeletionPolicyAnnotation: k8s.DeletionPolicyAbandon,
			})
			if err := k8sClient.Update(t.Context(), latest); err != nil {
				t.Fatalf("k8sClient.Update(%s) error = %v", tc.name, err)
			}

			_, err := r.Reconcile(t.Context(), req)

			if err != nil {
				t.Fatalf("Reconcile(%s) after setting abandon error = %v, want nil", tc.name, err)
			}
			afterAbandon := &unstructured.Unstructured{}
			afterAbandon.SetGroupVersionKind(tc.gvk)
			if getErr := k8sClient.Get(t.Context(), req.NamespacedName, afterAbandon); !apierrors.IsNotFound(getErr) {
				t.Errorf("Reconcile(%s) after abandon k8sClient.Get() error = %v, want NotFound", tc.name, getErr)
			}
		})
	}
}

func TestReconcile_DeleteSafetyGuard_TopLevelNonParentRefPreservesFinalizer(t *testing.T) {
	depNN := types.NamespacedName{Namespace: "test-ns", Name: "deleted-key"}
	now := metav1.Now()
	u := &unstructured.Unstructured{
		Object: map[string]any{
			"spec": map[string]any{
				"kmsKeyRef": map[string]any{
					"name": "deleted-key",
				},
			},
		},
	}
	u.SetGroupVersionKind(testGVK)
	u.SetName("top-level-cmek-res")
	u.SetNamespace("test-ns")
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	adapter := &mockAdapter{findFound: true}
	model := &mockModel{
		adapter: adapter,
		err:     k8s.NewReferenceNotFoundError(kmsv1beta1.KMSCryptoKeyGVK, depNN),
	}
	r, k8sClient, _ := newTestDirectReconciler(t, testGVK, model, u)
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{Namespace: "test-ns", Name: "top-level-cmek-res"},
	}

	res, err := r.Reconcile(t.Context(), req)

	if err != nil {
		t.Fatalf("Reconcile(top-level-cmek-res) error = %v, want nil", err)
	}
	if !res.Requeue {
		t.Errorf("Reconcile(top-level-cmek-res) res.Requeue = false, want true")
	}
	gotObj := &unstructured.Unstructured{}
	gotObj.SetGroupVersionKind(testGVK)
	if getErr := k8sClient.Get(t.Context(), req.NamespacedName, gotObj); getErr != nil {
		t.Fatalf("Reconcile(top-level-cmek-res) k8sClient.Get() error = %v, want object preserved", getErr)
	}
	if !k8s.HasFinalizer(gotObj, k8s.ControllerFinalizerName) {
		t.Errorf("Reconcile(top-level-cmek-res) HasFinalizer(ControllerFinalizerName) = false, want true")
	}
	conditions, _, _ := unstructured.NestedSlice(gotObj.Object, "status", "conditions")
	if len(conditions) == 0 {
		t.Fatalf("Reconcile(top-level-cmek-res) status.conditions = empty, want DependencyNotFound condition")
	}
	cond, _ := conditions[0].(map[string]any)
	if cond["reason"] != k8s.DependencyNotFound {
		t.Errorf("Reconcile(top-level-cmek-res) condition reason = %v, want %q", cond["reason"], k8s.DependencyNotFound)
	}
}
