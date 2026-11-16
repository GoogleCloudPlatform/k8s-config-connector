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

package parent

import (
	"context"
	"sync/atomic"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestReconcilers_UpdateSchema_NilReconcilers(t *testing.T) {
	// Arrange
	reconcilers := &Reconcilers{}
	crd := &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-crd",
		},
	}

	// Act
	err := reconcilers.UpdateSchema(crd)

	// Assert
	if err != nil {
		t.Fatalf("expected nil error when child reconcilers are nil, got: %v", err)
	}
}

type countingClient struct {
	client.Client
	targetGetCalls int32
	totalGetCalls  int32
	targetKey      client.ObjectKey
}

func (c *countingClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	atomic.AddInt32(&c.totalGetCalls, 1)
	if key == c.targetKey {
		atomic.AddInt32(&c.targetGetCalls, 1)
	}
	return c.Client.Get(ctx, key, obj, opts...)
}

type dummyChildReconciler struct {
	client client.Client
	gvk    schema.GroupVersionKind
}

func (d *dummyChildReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(d.gvk)
	if err := k8s.FetchUnstructured(ctx, d.client, req.NamespacedName, u); err != nil {
		return reconcile.Result{}, err
	}
	return reconcile.Result{}, nil
}

func TestParentReconcilerSingleGet(t *testing.T) {
	gvk := schema.GroupVersionKind{
		Group:   "iam.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "IAMPartialPolicy",
	}
	nn := types.NamespacedName{
		Namespace: "default",
		Name:      "test-resource",
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	u.SetNamespace(nn.Namespace)
	u.SetName(nn.Name)

	scheme := runtime.NewScheme()
	if err := corev1.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add corev1 to scheme: %v", err)
	}
	if err := operatorv1beta1.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to add operatorv1beta1 to scheme: %v", err)
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(u).Build()
	counterClient := &countingClient{Client: fakeClient, targetKey: nn}

	child := &dummyChildReconciler{
		client: counterClient,
		gvk:    gvk,
	}

	r := &ParentReconciler{
		Client: counterClient,
		gvk:    gvk,
		reconcilers: Reconcilers{
			Custom: &CustomReconciler{
				Type:       k8s.ReconcilerTypeIAMPartialPolicy,
				Reconciler: child,
			},
		},
	}

	req := reconcile.Request{NamespacedName: nn}
	if _, err := r.Reconcile(context.Background(), req); err != nil {
		t.Fatalf("Reconcile failed: %v", err)
	}

	calls := atomic.LoadInt32(&counterClient.targetGetCalls)
	if calls != 1 {
		t.Errorf("expected exactly 1 GET call for target resource %v during ParentReconciler reconcile, got %d", nn, calls)
	}
}
