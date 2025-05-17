// Copyright 2022 Google LLC
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

package unmanageddetector_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/unmanageddetector"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/randomid"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/main"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	"github.com/google/go-cmp/cmp"

	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	mgr manager.Manager

	fakeCRD = newTestKindCRD()
)

func TestReconcile_UnmanagedResource(t *testing.T) {
	testID := testvariable.NewUniqueID()
	client := mgr.GetClient()
	testcontroller.EnsureNamespaceExistsT(t, client, k8s.SystemNamespace)
	testcontroller.EnsureNamespaceExistsT(t, client, testID)

	resourceNN := types.NamespacedName{
		Namespace: testID,
		Name:      testID,
	}
	resource := newTestKindUnstructured(resourceNN)
	test.EnsureObjectExists(t, resource, client)

	reconciler, err := unmanageddetector.NewReconciler(mgr, fakeCRD.GroupVersionKind())
	if err != nil {
		t.Fatal(fmt.Errorf("error creating reconciler: %w", err))
	}
	res, err := reconciler.Reconcile(context.TODO(), reconcile.Request{NamespacedName: resourceNN})
	if err != nil {
		t.Fatal(fmt.Errorf("unexpected error during reconciliation: %w", err))
	}
	emptyResult := reconcile.Result{}
	if got, want := res, emptyResult; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected diff in reconcile result (-want +got): \n%v", cmp.Diff(want, got))
	}

	condition, found, err := getCurrentCondition(context.TODO(), client, resource)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting resource's condition: %w", err))
	}
	if !found {
		t.Fatalf("got nil condition for resource, want non-nil condition with reason '%v'", k8s.Unmanaged)
	}
	if gotReason, wantReason := condition.Reason, k8s.Unmanaged; gotReason != wantReason {
		t.Fatalf("got condition with reason '%v' for resource, want condition with reason '%v'", gotReason, wantReason)
	}
	if gotStatus, wantStatus := condition.Status, corev1.ConditionFalse; gotStatus != wantStatus {
		t.Fatalf("got condition with status '%v' for resource, want condition with status '%v'", gotStatus, wantStatus)
	}
}

func TestReconcile_ManagedResource(t *testing.T) {
	testID := testvariable.NewUniqueID()
	client := mgr.GetClient()
	testcontroller.EnsureNamespaceExistsT(t, client, k8s.SystemNamespace)
	testcontroller.EnsureNamespaceExistsT(t, client, testID)

	resourceNN := types.NamespacedName{
		Namespace: testID,
		Name:      testID,
	}
	resource := newTestKindUnstructured(resourceNN)
	test.EnsureObjectExists(t, resource, client)

	controller := newControllerUnstructuredForNamespace(resourceNN.Namespace)
	test.EnsureObjectExists(t, controller, client)

	reconciler, err := unmanageddetector.NewReconciler(mgr, fakeCRD.GroupVersionKind())
	if err != nil {
		t.Fatal(fmt.Errorf("error creating reconciler: %w", err))
	}
	res, err := reconciler.Reconcile(context.TODO(), reconcile.Request{NamespacedName: resourceNN})
	if err != nil {
		t.Fatal(fmt.Errorf("unexpected error during reconciliation: %w", err))
	}
	emptyResult := reconcile.Result{}
	if got, want := res, emptyResult; !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected diff in reconcile result (-want +got): \n%v", cmp.Diff(want, got))
	}

	condition, found, err := getCurrentCondition(context.TODO(), client, resource)
	if err != nil {
		t.Fatal(fmt.Errorf("error getting resource's condition: %w", err))
	}
	if found {
		t.Fatalf("got non-nil condition '%v' for resource, want nil condition", condition)
	}
}

func newTestKindCRD() *apiextensions.CustomResourceDefinition {
	crd := test.CRDForGVK(metav1.GroupVersionKind{
		Group:   "test.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "TestKind",
	})
	// Enable the status subresource for this CRD. This is needed to allow
	// UpdateStatus() calls to work on custom resources belonging to this CRD
	// on the API server.
	crd.Spec.Versions[0].Subresources = &apiextensions.CustomResourceSubresources{
		Status: &apiextensions.CustomResourceSubresourceStatus{},
	}
	return crd
}

func newTestKindUnstructured(nn types.NamespacedName) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": fmt.Sprintf("%v/%v", fakeCRD.Spec.Group, k8s.GetVersionFromCRD(fakeCRD)),
			"kind":       fakeCRD.Spec.Names.Kind,
			"metadata": map[string]interface{}{
				"namespace": nn.Namespace,
				"name":      nn.Name,
			},
		},
	}
}

func newControllerUnstructuredForNamespace(namespace string) *unstructured.Unstructured {
	controllerName := fmt.Sprintf("%v-%v", k8s.ControllerManagerNamePrefix, randomid.New().String())
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					k8s.KCCComponentLabel:    k8s.ControllerManagerNamePrefix,
					k8s.ScopedNamespaceLabel: namespace,
				},
				"namespace": k8s.SystemNamespace,
				"name":      controllerName,
			},
			"spec": map[string]interface{}{
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						k8s.KCCComponentLabel:    k8s.ControllerManagerNamePrefix,
						k8s.ScopedNamespaceLabel: namespace,
					},
				},
				"serviceName": controllerName,
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							k8s.KCCComponentLabel:    k8s.ControllerManagerNamePrefix,
							k8s.ScopedNamespaceLabel: namespace,
						},
					},
				},
			},
		},
	}
}

func getCurrentCondition(ctx context.Context, c client.Client, u *unstructured.Unstructured) (condition v1alpha1.Condition, found bool, err error) {
	nn := k8s.GetNamespacedName(u)
	unstruct := &unstructured.Unstructured{}
	unstruct.SetGroupVersionKind(u.GroupVersionKind())
	if err := c.Get(ctx, nn, unstruct); err != nil {
		return v1alpha1.Condition{}, false, fmt.Errorf("error getting resource from API server: %w", err)
	}
	resource, err := k8s.NewResource(unstruct)
	if err != nil {
		return v1alpha1.Condition{}, false, fmt.Errorf("error marhsalling unstruct to k8s resource: %w", err)
	}
	condition, found = k8s.GetReadyCondition(resource)
	return condition, found, nil
}

func TestMain(m *testing.M) {
	testmain.ForUnitTestsWithCRDs(m, []*apiextensions.CustomResourceDefinition{fakeCRD}, &mgr)
}
