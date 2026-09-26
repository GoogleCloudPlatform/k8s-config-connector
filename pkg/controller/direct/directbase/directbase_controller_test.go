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
	"net/http"
	"strings"
	"testing"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"google.golang.org/api/googleapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func TestReconcile_TerminalErrorModeBuiltin_HaltRetries(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetGeneration(1)
	u.SetAnnotations(map[string]string{
		k8s.TerminalErrorModeAnnotation: k8s.TerminalErrorModeBuiltin,
	})

	stWithBadRequest, _ := status.New(codes.InvalidArgument, "Invalid parameter").WithDetails(
		&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "spec.name",
					Description: "Must match regex ^[a-z0-9-]+$",
				},
			},
		},
	)

	adapter := &mockAdapter{
		findFound: false,
		createErr: stWithBadRequest.Err(),
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
		t.Fatalf("expected reconcile to return nil error on terminal error under builtin mode, got: %v", err)
	}

	if res.Requeue || res.RequeueAfter != 0 {
		t.Errorf("expected reconcile to halt (Result{}), got: %v", res)
	}

	// Verify status condition is UpdateFailedTerminalError and observedGeneration is 1
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(testGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}

	observedGen, found, err := unstructured.NestedInt64(obj.Object, "status", "observedGeneration")
	if err != nil || !found || observedGen != 1 {
		t.Errorf("expected status.observedGeneration to be 1, found: %v, got: %v", found, observedGen)
	}

	conditions, found, err := unstructured.NestedSlice(obj.Object, "status", "conditions")
	if err != nil || !found || len(conditions) == 0 {
		t.Fatalf("expected status conditions to be present, found: %v, err: %v", found, err)
	}

	readyCond := conditions[0].(map[string]interface{})
	if readyCond["status"] != string(corev1.ConditionFalse) {
		t.Errorf("expected condition status to be False, got: %v", readyCond["status"])
	}
	if readyCond["reason"] != k8s.UpdateFailedTerminalError {
		t.Errorf("expected condition reason to be %s, got: %v", k8s.UpdateFailedTerminalError, readyCond["reason"])
	}
}

func TestReconcile_TerminalErrorModeBuiltin_SkipWhenHalted(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetGeneration(1)
	u.SetAnnotations(map[string]string{
		k8s.TerminalErrorModeAnnotation: k8s.TerminalErrorModeBuiltin,
	})
	_ = unstructured.SetNestedField(u.Object, int64(1), "status", "observedGeneration")
	_ = unstructured.SetNestedSlice(u.Object, []interface{}{
		map[string]interface{}{
			"type":    "Ready",
			"status":  string(corev1.ConditionFalse),
			"reason":  k8s.UpdateFailedTerminalError,
			"message": "Update call failed: invalid argument",
		},
	}, "status", "conditions")

	adapter := &mockAdapter{
		findFound: false,
		findErr:   fmt.Errorf("Find should not be called"),
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
		t.Fatalf("expected nil error when skipping halted terminal error, got: %v", err)
	}

	if res.Requeue || res.RequeueAfter != 0 {
		t.Errorf("expected Result{}, got: %v", res)
	}

	if adapter.createCalled || adapter.updateCalled || adapter.deleteCalled {
		t.Errorf("expected no adapter calls when halted, got create=%v update=%v delete=%v",
			adapter.createCalled, adapter.updateCalled, adapter.deleteCalled)
	}
}

func TestReconcile_TerminalErrorModeBuiltin_ResumeOnSpecUpdate(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetGeneration(2) // Generation incremented on spec update
	u.SetAnnotations(map[string]string{
		k8s.TerminalErrorModeAnnotation: k8s.TerminalErrorModeBuiltin,
	})
	_ = unstructured.SetNestedField(u.Object, int64(1), "status", "observedGeneration") // Previous observed generation was 1
	_ = unstructured.SetNestedSlice(u.Object, []interface{}{
		map[string]interface{}{
			"type":    "Ready",
			"status":  string(corev1.ConditionFalse),
			"reason":  k8s.UpdateFailedTerminalError,
			"message": "Update call failed: invalid argument",
		},
	}, "status", "conditions")

	adapter := &mockAdapter{
		findFound: false, // will call Create
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
		t.Fatalf("expected successful reconcile on resumed spec update, got: %v", err)
	}

	if !adapter.createCalled {
		t.Errorf("expected Adapter.Create to be called on spec update")
	}

	if res.RequeueAfter == 0 {
		t.Errorf("expected RequeueAfter to be scheduled after successful reconcile, got: %v", res)
	}

	// Verify status is updated to UpToDate with observedGeneration 2
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(testGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}

	observedGen, found, err := unstructured.NestedInt64(obj.Object, "status", "observedGeneration")
	if err != nil || !found || observedGen != 2 {
		t.Errorf("expected status.observedGeneration to be 2, found: %v, got: %v", found, observedGen)
	}

	conditions, found, err := unstructured.NestedSlice(obj.Object, "status", "conditions")
	if err != nil || !found || len(conditions) == 0 {
		t.Fatalf("expected status conditions to be present, found: %v, err: %v", found, err)
	}

	readyCond := conditions[0].(map[string]interface{})
	if readyCond["status"] != string(corev1.ConditionTrue) {
		t.Errorf("expected condition status to be True, got: %v", readyCond["status"])
	}
	if readyCond["reason"] != k8s.UpToDate {
		t.Errorf("expected condition reason to be UpToDate, got: %v", readyCond["reason"])
	}
}

func TestReconcile_TerminalErrorModeNone_ForcesRetry(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetGeneration(1)
	u.SetAnnotations(map[string]string{
		k8s.TerminalErrorModeAnnotation: k8s.TerminalErrorModeNone,
	})

	adapter := &mockAdapter{
		findFound: false,
		createErr: status.Error(codes.InvalidArgument, "Invalid field parameter"),
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
		t.Fatalf("expected non-nil error to trigger workqueue retry when mode is none, got nil")
	}

	// Verify status condition is UpdateFailed (not UpdateFailedTerminalError)
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(testGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}

	conditions, found, err := unstructured.NestedSlice(obj.Object, "status", "conditions")
	if err != nil || !found || len(conditions) == 0 {
		t.Fatalf("expected status conditions to be present, found: %v, err: %v", found, err)
	}

	readyCond := conditions[0].(map[string]interface{})
	if readyCond["reason"] != k8s.UpdateFailed {
		t.Errorf("expected condition reason to be %s, got: %v", k8s.UpdateFailed, readyCond["reason"])
	}
}

func TestReconcile_TransientError_ContinuesRetrying(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetGeneration(1)
	u.SetAnnotations(map[string]string{
		k8s.TerminalErrorModeAnnotation: k8s.TerminalErrorModeBuiltin,
	})

	// Transient error: HTTP 400 with RESOURCE_IN_USE_BY_ANOTHER_RESOURCE
	transientErr := &googleapi.Error{
		Code:    http.StatusBadRequest,
		Message: "The resource is currently in use",
		Errors: []googleapi.ErrorItem{
			{
				Reason:  "RESOURCE_IN_USE_BY_ANOTHER_RESOURCE",
				Message: "Resource is in use by another resource",
			},
		},
	}

	adapter := &mockAdapter{
		findFound: false,
		createErr: transientErr,
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
		t.Fatalf("expected non-nil error to trigger workqueue retry for transient error, got nil")
	}

	// Verify status condition is UpdateFailed (not UpdateFailedTerminalError)
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(testGVK)
	if err := k8sClient.Get(ctx, req.NamespacedName, obj); err != nil {
		t.Fatalf("failed to retrieve object: %v", err)
	}

	conditions, found, err := unstructured.NestedSlice(obj.Object, "status", "conditions")
	if err != nil || !found || len(conditions) == 0 {
		t.Fatalf("expected status conditions to be present, found: %v, err: %v", found, err)
	}

	readyCond := conditions[0].(map[string]interface{})
	if readyCond["reason"] != k8s.UpdateFailed {
		t.Errorf("expected condition reason to be %s, got: %v", k8s.UpdateFailed, readyCond["reason"])
	}
}

func TestReconcile_TerminalErrorModeBuiltin_DeletePrecedence(t *testing.T) {
	ctx := context.TODO()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = operatorv1beta1.AddToScheme(scheme)

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(testGVK)
	u.SetName("test-resource")
	u.SetNamespace("test-ns")
	u.SetGeneration(1)
	u.SetAnnotations(map[string]string{
		k8s.TerminalErrorModeAnnotation: k8s.TerminalErrorModeBuiltin,
	})

	now := metav1.Now()
	u.SetDeletionTimestamp(&now)
	u.SetFinalizers([]string{k8s.ControllerFinalizerName})

	_ = unstructured.SetNestedField(u.Object, int64(1), "status", "observedGeneration")
	_ = unstructured.SetNestedSlice(u.Object, []interface{}{
		map[string]interface{}{
			"type":    "Ready",
			"status":  string(corev1.ConditionFalse),
			"reason":  k8s.UpdateFailedTerminalError,
			"message": "Update call failed: invalid argument",
		},
	}, "status", "conditions")

	adapter := &mockAdapter{
		findFound:     true,
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
	if err != nil {
		t.Fatalf("expected successful reconcile on deletion, got: %v", err)
	}

	if !adapter.deleteCalled {
		t.Errorf("expected Adapter.Delete to be called when deleting resource in terminal error state")
	}

	// Verify object has been finalized / deleted
	obj := &unstructured.Unstructured{}
	obj.SetGroupVersionKind(testGVK)
	err = k8sClient.Get(ctx, req.NamespacedName, obj)
	if err != nil && !apierrors.IsNotFound(err) {
		t.Fatalf("failed to retrieve object: %v", err)
	}

	if err == nil {
		for _, f := range obj.GetFinalizers() {
			if f == k8s.ControllerFinalizerName {
				t.Errorf("expected finalizer to be removed, but still found %s", f)
			}
		}
	}
}
