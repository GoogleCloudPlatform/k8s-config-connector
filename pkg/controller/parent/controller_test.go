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
	"strings"
	"sync/atomic"
	"testing"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/tf"
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

func newTestCRD(group, version, kind, description string) *apiextensions.CustomResourceDefinition {
	return &apiextensions.CustomResourceDefinition{
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group: group,
			Names: apiextensions.CustomResourceDefinitionNames{
				Kind: kind,
			},
			Versions: []apiextensions.CustomResourceDefinitionVersion{
				{
					Name: version,
					Schema: &apiextensions.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
							Type:        "object",
							Description: description,
						},
					},
				},
			},
		},
	}
}

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

func TestReconcilers_UpdateSchema_ChildReconcilers(t *testing.T) {
	tests := []struct {
		name          string
		newReconciler func(crd *apiextensions.CustomResourceDefinition) (k8s.SchemaReferenceUpdater, func() *k8s.SchemaReference)
	}{
		{
			name: "Terraform Reconciler",
			newReconciler: func(crd *apiextensions.CustomResourceDefinition) (k8s.SchemaReferenceUpdater, func() *k8s.SchemaReference) {
				r := &tf.Reconciler{}
				if err := r.UpdateSchema(crd); err != nil {
					t.Fatalf("failed to initialize schema: %v", err)
				}
				return r, r.SchemaReference
			},
		},
		{
			name: "DCL Reconciler",
			newReconciler: func(crd *apiextensions.CustomResourceDefinition) (k8s.SchemaReferenceUpdater, func() *k8s.SchemaReference) {
				r := &dclcontroller.Reconciler{}
				if err := r.UpdateSchema(crd); err != nil {
					t.Fatalf("failed to initialize schema: %v", err)
				}
				return r, r.SchemaReference
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name+"_Success", func(t *testing.T) {
			group := "foo.cnrm.cloud.google.com"
			version := "v1beta1"
			kind := "FooResource"
			initialCRD := newTestCRD(group, version, kind, "initial description")
			updater, getSchemaRef := tc.newReconciler(initialCRD)
			updatedCRD := newTestCRD(group, version, kind, "updated description")

			err := updater.UpdateSchema(updatedCRD)
			if err != nil {
				t.Fatalf("unexpected error updating schema: %v", err)
			}
			newRef := getSchemaRef()
			if newRef == nil {
				t.Fatal("expected non-nil schemaRef after update")
			}
			if newRef.JSONSchema.Description != "updated description" {
				t.Fatalf("expected description 'updated description', got %q", newRef.JSONSchema.Description)
			}
			if newRef.GVK.Kind != kind || newRef.GVK.Group != group || newRef.GVK.Version != version {
				t.Fatalf("expected GVK %s/%s, %s, got %v", group, version, kind, newRef.GVK)
			}
		})

		t.Run(tc.name+"_MismatchedGVK", func(t *testing.T) {
			group := "foo.cnrm.cloud.google.com"
			version := "v1beta1"
			initialCRD := newTestCRD(group, version, "FooResource", "description")
			updater, getSchemaRef := tc.newReconciler(initialCRD)
			differentCRD := newTestCRD(group, version, "BarResource", "description")

			err := updater.UpdateSchema(differentCRD)
			if err == nil {
				t.Fatal("expected error due to mismatched GVK, got nil")
			}
			if !strings.Contains(err.Error(), "unexpected mismatch of GVK") {
				t.Fatalf("expected error message to contain 'unexpected mismatch of GVK', got %q", err.Error())
			}
			currentRef := getSchemaRef()
			if currentRef.GVK.Kind != "FooResource" {
				t.Fatalf("expected schemaRef to remain FooResource, got %s", currentRef.GVK.Kind)
			}
		})
	}
}

func TestReconcilers_UpdateSchema_ParentBothChildren(t *testing.T) {
	group := "foo.cnrm.cloud.google.com"
	version := "v1beta1"
	kind := "FooResource"
	initialCRD := newTestCRD(group, version, kind, "initial description")

	tfReconciler := &tf.Reconciler{}
	if err := tfReconciler.UpdateSchema(initialCRD); err != nil {
		t.Fatalf("failed to init TF schema: %v", err)
	}
	dclReconciler := &dclcontroller.Reconciler{}
	if err := dclReconciler.UpdateSchema(initialCRD); err != nil {
		t.Fatalf("failed to init DCL schema: %v", err)
	}

	reconcilers := &Reconcilers{
		TF:  tfReconciler,
		DCL: dclReconciler,
	}

	updatedCRD := newTestCRD(group, version, kind, "updated description")
	if err := reconcilers.UpdateSchema(updatedCRD); err != nil {
		t.Fatalf("unexpected error updating schema via parent Reconcilers: %v", err)
	}

	if tfReconciler.SchemaReference().JSONSchema.Description != "updated description" {
		t.Errorf("expected TF schema to be updated, got %q", tfReconciler.SchemaReference().JSONSchema.Description)
	}
	if dclReconciler.SchemaReference().JSONSchema.Description != "updated description" {
		t.Errorf("expected DCL schema to be updated, got %q", dclReconciler.SchemaReference().JSONSchema.Description)
	}

	// Mismatched GVK propagates error
	mismatchedCRD := newTestCRD(group, version, "BarResource", "description")
	if err := reconcilers.UpdateSchema(mismatchedCRD); err == nil {
		t.Fatal("expected error on mismatched GVK, got nil")
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
