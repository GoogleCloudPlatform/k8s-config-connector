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

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/unmanageddetector"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/randomid"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"
	"github.com/google/go-cmp/cmp"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	fakeCRDGVK = schema.GroupVersionKind{
		Group:   "test.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "TestKind",
	}
)

func runManager(t *testing.T, restConfig *rest.Config) manager.Manager {
	mgrOpts := manager.Options{}
	mgrOpts.Metrics.BindAddress = "0" // no metrics

	mgr, err := manager.New(restConfig, mgrOpts)
	if err != nil {
		t.Fatalf("error creating manager: %v", err)
	}

	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)

	errChan := make(chan error)
	go func() {
		err := mgr.Start(ctx)
		if err != nil {
			t.Errorf("error from manager: %v", err)
		}
		errChan <- err
	}()

	t.Cleanup(func() {
		cancel()
		err := <-errChan
		if err != nil {
			t.Fatalf("error from manager: %v", err)
		}
	})

	return mgr
}

func TestReconcile_UnmanagedResource(t *testing.T) {
	tests := []struct {
		name             string
		namespace        string
		managerNamespace string
		ccc              *unstructured.Unstructured
	}{
		{
			name:      "namespaced CCC",
			namespace: testvariable.NewUniqueID(),
		},
		{
			name:             "per namespace CCC",
			namespace:        "t1234-tenant0-provider",
			managerNamespace: "t1234-tenant0-supervisor",
			ccc:              newConfigConnectorContextUnstructured("t1234-tenant0-provider", "t1234-tenant0-supervisor"),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Cannot run tests in parallel as they create cluster-wide ConfigConnector
			ctx := context.TODO()
			h := test.NewKubeHarness(ctx, t)
			client := h.GetClient()

			h.CreateDummyCRD(fakeCRDGVK)

			mgr := runManager(t, h.GetRESTConfig())

			testID := testvariable.NewUniqueID()
			testcontroller.EnsureNamespaceExistsT(t, client, k8s.SystemNamespace)
			testcontroller.EnsureNamespaceExistsT(t, client, tc.namespace)
			if tc.managerNamespace != "" {
				testcontroller.EnsureNamespaceExistsT(t, client, tc.managerNamespace)
			}

			resourceNN := types.NamespacedName{
				Namespace: tc.namespace,
				Name:      testID,
			}
			resource := newTestKindUnstructured(resourceNN)
			test.EnsureObjectExists(t, resource, client)

			cc := newConfigConnectorNamespacedUnstructured()
			test.EnsureObjectExists(t, cc, client)
			defer func() {
				err := client.Delete(context.Background(), cc)
				if err != nil {
					t.Fatalf("cannot delete ConfigConnector: %v", err)
				}
			}()
			if tc.ccc != nil {
				test.EnsureObjectExists(t, tc.ccc, client)
				defer func() {
					err := client.Delete(context.Background(), tc.ccc)
					if err != nil {
						t.Fatalf("cannot delete ConfigConnectorContext: %v", err)
					}
				}()
			}

			reconciler, err := unmanageddetector.NewReconciler(mgr, fakeCRDGVK)
			if err != nil {
				t.Fatalf("error creating reconciler: %v", err)
			}
			res, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: resourceNN})
			if err != nil {
				t.Fatalf("unexpected error during reconciliation: %v", err)
			}
			emptyResult := reconcile.Result{}
			if got, want := res, emptyResult; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff in reconcile result (-want +got): \n%v", cmp.Diff(want, got))
			}

			condition, found, err := getCurrentCondition(ctx, client, resource)
			if err != nil {
				t.Fatalf("error getting resource's condition: %v", err)
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
		})
	}
}

func TestReconcile_ManagedResource(t *testing.T) {
	tests := []struct {
		name             string
		namespace        string
		managerNamespace string
	}{
		{
			name:             "namespaced CCC",
			namespace:        testvariable.NewUniqueID(),
			managerNamespace: k8s.SystemNamespace,
		},
		{
			name:             "per namespace CCC",
			namespace:        "t1234-tenant0-provider",
			managerNamespace: "t1234-tenant0-supervisor",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Cannot run tests in parallel as they create cluster-wide ConfigConnector
			ctx := context.TODO()
			h := test.NewKubeHarness(ctx, t)
			client := h.GetClient()

			h.CreateDummyCRD(fakeCRDGVK)

			mgr := runManager(t, h.GetRESTConfig())

			testID := testvariable.NewUniqueID()
			testcontroller.EnsureNamespaceExistsT(t, client, k8s.SystemNamespace)
			testcontroller.EnsureNamespaceExistsT(t, client, tc.namespace)
			testcontroller.EnsureNamespaceExistsT(t, client, tc.managerNamespace)

			resourceNN := types.NamespacedName{
				Namespace: tc.namespace,
				Name:      testID,
			}
			resource := newTestKindUnstructured(resourceNN)
			test.EnsureObjectExists(t, resource, client)

			cc := newConfigConnectorNamespacedUnstructured()
			test.EnsureObjectExists(t, cc, client)
			defer func() {
				err := client.Delete(context.Background(), cc)
				if err != nil {
					t.Fatalf("cannot delete ConfigConnector: %v", err)
				}
			}()

			controller := newControllerUnstructuredForNamespace(resourceNN.Namespace, tc.managerNamespace)
			test.EnsureObjectExists(t, controller, client)

			ccc := newConfigConnectorContextUnstructured(resourceNN.Namespace, tc.managerNamespace)
			test.EnsureObjectExists(t, ccc, client)

			reconciler, err := unmanageddetector.NewReconciler(mgr, fakeCRDGVK)
			if err != nil {
				t.Fatalf("error creating reconciler: %v", err)
			}
			res, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: resourceNN})
			if err != nil {
				t.Fatalf("unexpected error during reconciliation: %v", err)
			}
			emptyResult := reconcile.Result{}
			if got, want := res, emptyResult; !reflect.DeepEqual(got, want) {
				t.Fatalf("unexpected diff in reconcile result (-want +got): \n%v", cmp.Diff(want, got))
			}

			condition, found, err := getCurrentCondition(ctx, client, resource)
			if err != nil {
				t.Fatalf("error getting resource's condition: %v", err)
			}
			if found {
				t.Fatalf("got non-nil condition '%v' for resource, want nil condition", condition)
			}
		})
	}
}

func newTestKindUnstructured(nn types.NamespacedName) *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": fmt.Sprintf("%v/%v", fakeCRDGVK.Group, fakeCRDGVK.Version),
			"kind":       fakeCRDGVK.Kind,
			"metadata": map[string]interface{}{
				"namespace": nn.Namespace,
				"name":      nn.Name,
			},
		},
	}
}

func newControllerUnstructuredForNamespace(namespace, managerNamespace string) *unstructured.Unstructured {
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
				"namespace": managerNamespace,
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

func newConfigConnectorContextUnstructured(namespace, managerNamespace string) *unstructured.Unstructured {
	if managerNamespace == k8s.SystemNamespace {
		return &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "core.cnrm.cloud.google.com/v1beta1",
				"kind":       "ConfigConnectorContext",
				"metadata": map[string]interface{}{
					"namespace": namespace,
					"name":      corev1beta1.ConfigConnectorContextAllowedName,
				},
				"spec": map[string]interface{}{
					"googleServiceAccount": "tenant-gcp-account@1234.iam.gserviceaccount.com",
					"stateIntoSpec":        "Absent",
				},
			},
		}
	}
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "core.cnrm.cloud.google.com/v1beta1",
			"kind":       "ConfigConnectorContext",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"tenancy.gke.io/access-level": "supervisor",
					"tenancy.gke.io/project":      "t1234",
					"tenancy.gke.io/tenant":       "t1234-tenant0",
				},
				"namespace": namespace,
				"name":      corev1beta1.ConfigConnectorContextAllowedName,
			},
			"spec": map[string]interface{}{
				"googleServiceAccount": "tenant-gcp-account@1234.iam.gserviceaccount.com",
				"stateIntoSpec":        "Absent",
				"managerNamespace":     managerNamespace,
			},
		},
	}
}

func newConfigConnectorNamespacedUnstructured() *unstructured.Unstructured {
	apiVersion, kind := corev1beta1.ConfigConnectorGroupVersionKind.ToAPIVersionAndKind()
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": apiVersion,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"name": "configconnector.core.cnrm.cloud.google.com",
			},
			"spec": map[string]interface{}{
				"mode": "namespaced",
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
