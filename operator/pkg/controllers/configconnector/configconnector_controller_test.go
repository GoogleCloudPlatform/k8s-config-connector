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

package configconnector

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
	testmocks "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/mocks"

	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func TestHandleReconcileFailed(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	c := mgr.GetClient()
	mockEventRecorder := testmocks.NewMockEventRecorder(t, mgr.GetScheme())
	r := Reconciler{
		client:   c,
		recorder: mockEventRecorder,
	}

	apiVersion, kind := corev1beta1.ConfigConnectorGroupVersionKind.ToAPIVersionAndKind()
	nn := types.NamespacedName{
		Name: "kcc-for-reconcile-failed-test",
	}

	tc := testCaseStruct{
		cc: &corev1beta1.ConfigConnector{
			TypeMeta: metav1.TypeMeta{
				Kind:       kind,
				APIVersion: apiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: nn.Name,
			},
			Spec: corev1beta1.ConfigConnectorSpec{
				Mode: "namespaced",
			},
		},
	}

	if err := c.Create(ctx, tc.cc); err != nil {
		t.Fatalf("failed to create ConfigConnector: %v", err)
	}
	reconcileErr := fmt.Errorf("reconciliation error")
	if err := r.handleReconcileFailed(ctx, nn, reconcileErr); err != nil {
		t.Errorf("error handling failed reconciliation: %v", err)
	}

	expectedErrMsg := "error during reconciliation: reconciliation error"
	mockEventRecorder.AssertEventRecorded(kind, nn, v1.EventTypeWarning, k8s.UpdateFailed, expectedErrMsg)

	newCC := &corev1beta1.ConfigConnector{}
	if err := c.Get(ctx, nn, newCC); err != nil {
		t.Errorf("failed to get ConfigConnector after attempt to handle failed reconciliation: %v", err)
	}
	status := newCC.GetCommonStatus()
	if status.Healthy {
		t.Errorf("unexpected value for status.healthy: got 'true', want 'false'")
	}
	if len(status.Errors) != 1 {
		t.Errorf("unexpected number of errors in status.errors: got %v errors, want 1 error. Got the errors: %v", len(status.Errors), status.Errors)
	} else if errMsg := status.Errors[0]; errMsg != expectedErrMsg {
		t.Errorf("unexpected error in status.errors: got '%v', want '%v'", errMsg, expectedErrMsg)
	}
}

func TestHandleReconcileSucceeded(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	c := mgr.GetClient()
	mockEventRecorder := testmocks.NewMockEventRecorder(t, mgr.GetScheme())
	r := Reconciler{
		client:   c,
		recorder: mockEventRecorder,
	}
	apiVersion, kind := corev1beta1.ConfigConnectorGroupVersionKind.ToAPIVersionAndKind()
	nn := types.NamespacedName{
		Name: "kcc-for-reconcile-succeeded-test",
	}
	tc := testCaseStruct{
		cc: &corev1beta1.ConfigConnector{
			TypeMeta: metav1.TypeMeta{
				Kind:       kind,
				APIVersion: apiVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: nn.Name,
			},
			Spec: corev1beta1.ConfigConnectorSpec{
				Mode: "namespaced",
			},
		},
	}

	if err := c.Create(ctx, tc.cc); err != nil {
		t.Fatalf("failed to create ConfigConnector: %v", err)
	}
	if err := r.handleReconcileSucceeded(ctx, nn); err != nil {
		t.Errorf("error handling successful reconciliation: %v", err)
	}
	mockEventRecorder.AssertEventRecorded(kind, nn, v1.EventTypeNormal, k8s.UpToDate, k8s.UpToDateMessage)

	newCC := &corev1beta1.ConfigConnector{}
	if err := c.Get(ctx, nn, newCC); err != nil {
		t.Errorf("failed to get ConfigConnector after attempt to handle successful reconciliation: %v", err)
	}
	status := newCC.GetCommonStatus()
	if !status.Healthy {
		t.Errorf("unexpected value for status.healthy: got 'false', want 'true'")
	}
	if len(status.Errors) != 0 {
		t.Errorf("unexpected number of errors in status.errors: got %v errors, want 0 errors. Got the errors: %v", len(status.Errors), status.Errors)
	}
}

func TestHandleConfigConnectorCreate(t *testing.T) {
	t.Parallel()
	tests := []testCaseStruct{
		{
			name: "1 CC and 1 CCContext, namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-1",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      corev1beta1.ConfigConnectorContextAllowedName,
						Namespace: "foo-ns",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			loadedManifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return []string{testcontroller.FooCRD, testcontroller.SystemNs}
			},
		},
		{
			name: "1 CC and no CCContext, namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-2",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			loadedManifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return []string{testcontroller.FooCRD, testcontroller.SystemNs}
			},
		},
		{
			name: "1 CC in cluster-mode with workload identity and no CCContext",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-1",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			loadedManifest: testcontroller.GetClusterModeWorkloadIdentityManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
		},
		{
			name: "1 CC in cluster-mode with gcp identity and no CCContext",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-1",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			loadedManifest: testcontroller.GetClusterModeGCPManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key")
			},
		},
		{
			name: "1 CC with 1 CCC (ignored), cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-1",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			cccs: []corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      corev1beta1.ConfigConnectorContextAllowedName,
						Namespace: "foo-ns",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			loadedManifest: testcontroller.GetClusterModeWorkloadIdentityManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
		},
		{
			name: "1 CC and 1 CCContext, per namespace mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc-1",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      corev1beta1.ConfigConnectorContextAllowedName,
						Namespace: "foo-ns",
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
						ManagerNamespace:     "t1234-tenant0-supervisor",
					},
				},
			},
			loadedManifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return []string{testcontroller.FooCRD, testcontroller.SystemNs}
			},
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			testcontroller.EnsureNamespaceExists(c, k8s.CNRMSystemNamespace)
			testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)
			m := testcontroller.ParseObjects(ctx, t, tc.loadedManifest)
			r := newConfigConnectorReconciler(c)
			if tc.managerNamespaceIsolation == k8s.ManagerNamespaceIsolationDedicated {
				r.managerNamespaceIsolation = k8s.ManagerNamespaceIsolationDedicated
			} else {
				r.managerNamespaceIsolation = k8s.ManagerNamespaceIsolationShared
			}

			if err := c.Create(ctx, tc.cc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}

			for _, ccc := range tc.cccs {
				testcontroller.EnsureNamespaceExists(c, ccc.Namespace)
				if err := c.Create(ctx, &ccc); err != nil {
					t.Fatalf("error creating %v %v: %v", ccc.Kind, ccc.Name, err)
				}
			}

			if err := handleLifecycles(ctx, t, r, tc.cc, m); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedObjs := tc.resultsFunc(t, c)
			expectedManifest := testcontroller.ParseObjects(ctx, t, expectedObjs)
			expectedJSON, err := expectedManifest.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			resJSON, err := m.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(resJSON, expectedJSON) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(resJSON, expectedJSON))
			}

			// Verify that CCC objects are NOT attached finalizers by the CC controller.
			for _, ccc := range tc.cccs {
				o := &corev1beta1.ConfigConnectorContext{}
				contextKey := types.NamespacedName{
					Name:      ccc.Name,
					Namespace: ccc.Namespace,
				}
				if err := c.Get(ctx, contextKey, o); err != nil {
					t.Fatalf("error getting ConfigConnector %v: %v", contextKey, err)
				}
				if testcontroller.HasOperatorFinalizer(o) {
					t.Fatalf("%v finalizer was found in %v", k8s.OperatorFinalizer, ccc)
				}
			}
			// Verify that CC contains the operator finalizer.
			ccKey := types.NamespacedName{
				Name:      tc.cc.Name,
				Namespace: tc.cc.Namespace,
			}
			if err := c.Get(ctx, ccKey, tc.cc); err != nil {
				t.Fatalf("error getting ConfigConnector %v: %v", ccKey, err)
			}
			if !testcontroller.HasOperatorFinalizer(tc.cc) {
				t.Fatalf("no %v finalizer was found in %v", k8s.OperatorFinalizer, tc.cc)
			}
		})
	}
}

func TestHandleConfigConnectorDelete(t *testing.T) {
	tests := []struct {
		name                      string
		cc                        *corev1beta1.ConfigConnector
		cccs                      []corev1beta1.ConfigConnectorContext
		installedObjectsFunc      func(t *testing.T, c client.Client) []string
		resultsFunc               func(t *testing.T, c client.Client) []string
		managerNamespaceIsolation string
	}{
		{
			name: "cluster mode workload identity uninstall",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "test-kcc",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
		},
		{
			name: "cluster mode gcp identity uninstall",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "test-kcc",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key")
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
		},
		{
			name: "namespaced mode CC, 1 CCContext, delete CC",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "test-kcc",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				res := make([]string, 0)
				res = append(res, testcontroller.GetSharedComponentsManifest()...)
				namespacedManifest := testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.NamespacedComponentsTemplate, "foo-ns", "foo-ns@bar.iam.gserviceaccount.com", false, "", c)
				res = append(res, namespacedManifest...)
				res = append(res, testcontroller.PerNamespaceControllerManagerPod)
				return res
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
		},
		{
			name: "namespaced mode CC, 1 CCContext per namespace, delete CC",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       "test-kcc",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "t1234-tenant0-provider",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						ManagerNamespace:     "t1234-tenant0-supervisor",
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				res := make([]string, 0)
				res = append(res, testcontroller.GetSharedComponentsManifest()...)
				namespacedManifest := testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.NamespacedComponentsTemplate, "t1234-tenant0-provider", "foo-ns@bar.iam.gserviceaccount.com", false, "", c)
				res = append(res, namespacedManifest...)
				res = append(res, testcontroller.PerNamespaceControllerManagerPod)
				return res
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)
			testcontroller.EnsureNamespaceExists(c, k8s.CNRMSystemNamespace)
			m := testcontroller.ParseObjects(ctx, t, tc.installedObjectsFunc(t, c))
			r := newConfigConnectorReconciler(c)
			if tc.managerNamespaceIsolation == k8s.ManagerNamespaceIsolationDedicated {
				r.managerNamespaceIsolation = k8s.ManagerNamespaceIsolationDedicated
			} else {
				r.managerNamespaceIsolation = k8s.ManagerNamespaceIsolationShared
			}

			if err := c.Create(ctx, tc.cc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			key := types.NamespacedName{
				Name: tc.cc.Name,
			}
			if err := c.Get(ctx, key, tc.cc); err != nil {
				t.Fatalf("error getting %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			for _, ccc := range tc.cccs {
				testcontroller.EnsureNamespaceExists(c, ccc.Namespace)
				if ccc.Spec.ManagerNamespace != "" {
					testcontroller.EnsureNamespaceExists(c, ccc.Spec.ManagerNamespace)
				}
				if err := c.Create(ctx, &ccc); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", ccc.Kind, ccc.Namespace, ccc.Name, err)
				}
			}

			for _, item := range m.Items {
				if err := c.Create(ctx, item.UnstructuredObject()); err != nil && !apierrors.IsAlreadyExists(err) {
					t.Fatalf("error creating %v %v: %v", item.GroupKind(), item.GetName(), err)
				}
			}

			// issue the delete request for the configconnector object
			if err := c.Delete(ctx, tc.cc); err != nil {
				t.Fatalf("error deleting %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			if err := c.Get(ctx, key, tc.cc); err != nil {
				t.Fatalf("error getting %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			if len(tc.cccs) > 0 {
				// Expect that the first attempt returns an error.
				if err := handleLifecycles(ctx, t, r, tc.cc, m); err == nil {
					t.Fatalf("expect to have an error because the controller manager pod per namespace is not deleted, but got nil")
				}
				// Simulate that CCC controller kicks in and deletes the controller manager pod.
				for _, item := range m.Items {
					if item.Kind == "Pod" && strings.Contains(item.GetName(), "cnrm-controller-manager") {
						if err := c.Delete(ctx, item.UnstructuredObject()); err != nil {
							t.Fatalf("error deleting %v %v: %v", item.GroupKind(), item.GetName(), err)
						}
					}
				}
			}
			if err := handleLifecycles(ctx, t, r, tc.cc, m); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedObjs := tc.resultsFunc(t, c)
			expectedManifest := testcontroller.ParseObjects(ctx, t, expectedObjs)
			expectedJSON, err := expectedManifest.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			resJSON, err := m.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(resJSON, expectedJSON) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(resJSON, expectedJSON))
			}

			// Assert that the ConfigConnector object is deleted.
			if err := c.Get(ctx, key, tc.cc); err == nil || !apierrors.IsNotFound(err) {
				t.Fatalf("expect to get %v error, but got error: %v", metav1.StatusReasonNotFound, err)
			}
		})
	}
}

func TestConfigConnectorUpdate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                      string
		cc                        *corev1beta1.ConfigConnector
		updatedCc                 *corev1beta1.ConfigConnector
		cccs                      []*corev1beta1.ConfigConnectorContext
		installedObjectsFunc      func(t *testing.T, c client.Client) []string
		manifest                  []string
		toDeleteObjectsFunc       func(t *testing.T, c client.Client) []string
		resultsFunc               func(t *testing.T, c client.Client) []string
		managerNamespaceIsolation string
	}{
		{
			name: "workload identity cluster mode to namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.ClusterModeOnlyWorkloadIdentityComponents, "foo@bar.iam.gserviceaccount.com")
			},
			manifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				return res
			},
		},
		{
			name: "gcp identity cluster mode to namespaced mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.ClusterModeOnlyGCPComponents, "my-key")
			},
			manifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				return res
			},
		},
		{
			name: "namespaced mode to workload identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				res = append(res, testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.NamespacedComponentsTemplate, "foo-ns", "foo-ns@bar.iam.gserviceaccount.com", false, "", c)...)
				res = append(res, testcontroller.PerNamespaceControllerManagerPod)
				return res
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return []string{}
			},
			manifest: testcontroller.GetClusterModeWorkloadIdentityManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
		},
		{
			name: "namespaced mode to gcp identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				res = append(res, testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.NamespacedComponentsTemplate, "foo-ns", "foo-ns@bar.iam.gserviceaccount.com", false, "", c)...)
				res = append(res, testcontroller.PerNamespaceControllerManagerPod)
				return res
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return []string{}
			},
			manifest: testcontroller.GetClusterModeGCPManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key ")
			},
		},
		{
			name: "workload identity cluster mode to gcp identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			manifest: testcontroller.GetClusterModeGCPManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key ")
			},
		},
		{
			name: "gcp identity cluster mode to workload identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key ")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			manifest: testcontroller.GetClusterModeWorkloadIdentityManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
		},

		{
			name: "workload identity cluster mode to per namespace mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
						ManagerNamespace:     "t1234-tenant0-supervisor",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.ClusterModeOnlyWorkloadIdentityComponents, "foo@bar.iam.gserviceaccount.com")
			},
			manifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				return res
			},
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
		{
			name: "gcp identity cluster mode to per namespace mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "foo-ns",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
						ManagerNamespace:     "t1234-tenant0-supervisor",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key")
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.ClusterModeOnlyGCPComponents, "my-key")
			},
			manifest: testcontroller.GetSharedComponentsManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				return res
			},
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
		{
			name: "per namespace mode to workload identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					Mode:                 "cluster",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "t1234-tenant0-provider",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
						ManagerNamespace:     "t1234-tenant0-supervisor",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				res = append(res, testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.NamespacedComponentsTemplate, "t1234-tenant0-provider", "foo-ns@bar.iam.gserviceaccount.com", false, "", c)...)
				res = append(res, testcontroller.NamespacedControllerManagerPod)
				return res
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return []string{}
			},
			manifest: testcontroller.GetClusterModeWorkloadIdentityManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceGSA(testcontroller.GetClusterModeWorkloadIdentityManifest(), "foo@bar.iam.gserviceaccount.com")
			},
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
		{
			name: "per namespace mode to gcp identity cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			updatedCc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-kcc",
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					CredentialSecretName: "my-key",
					Mode:                 "cluster",
				},
			},
			cccs: []*corev1beta1.ConfigConnectorContext{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:       corev1beta1.ConfigConnectorContextAllowedName,
						Namespace:  "t1234-tenant0-provider",
						Finalizers: []string{k8s.OperatorFinalizer},
					},
					Spec: corev1beta1.ConfigConnectorContextSpec{
						GoogleServiceAccount: "foo-ns@bar.iam.gserviceaccount.com",
						ManagerNamespace:     "t1234-tenant0-supervisor",
					},
				},
			},
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				res := []string{testcontroller.FooCRD, testcontroller.SystemNs}
				res = append(res, testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.NamespacedComponentsTemplate, "t1234-tenant0-provider", "foo-ns@bar.iam.gserviceaccount.com", false, "", c)...)
				res = append(res, testcontroller.NamespacedControllerManagerPod)
				return res
			},
			toDeleteObjectsFunc: func(t *testing.T, c client.Client) []string {
				return []string{}
			},
			manifest: testcontroller.GetClusterModeGCPManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyReplaceSecretVolume(testcontroller.GetClusterModeGCPManifest(), "my-key ")
			},
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			r := newConfigConnectorReconciler(c)
			if tc.managerNamespaceIsolation == k8s.ManagerNamespaceIsolationDedicated {
				r.managerNamespaceIsolation = k8s.ManagerNamespaceIsolationDedicated
			} else {
				r.managerNamespaceIsolation = k8s.ManagerNamespaceIsolationShared
			}

			testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)
			if err := c.Create(ctx, tc.cc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			for _, ccc := range tc.cccs {
				testcontroller.EnsureNamespaceExists(c, ccc.Namespace)
				if ccc.Spec.ManagerNamespace != "" {
					testcontroller.EnsureNamespaceExists(c, ccc.Spec.ManagerNamespace)
				}
				if err := c.Create(ctx, ccc); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", ccc.Kind, ccc.Namespace, ccc.Name, err)
				}
			}
			installedComponents := tc.installedObjectsFunc(t, c)
			installedManifest := testcontroller.ParseObjects(ctx, t, installedComponents)
			for _, item := range installedManifest.Items {
				if err := c.Create(ctx, item.UnstructuredObject()); err != nil && !apierrors.IsAlreadyExists(err) {
					t.Fatalf("error creating %v %v: %v", item.GroupKind(), item.GetName(), err)
				}
			}

			// update ConfigConnector
			tc.updatedCc.ResourceVersion = tc.cc.ResourceVersion
			if err := c.Update(ctx, tc.updatedCc); err != nil {
				t.Fatalf("error updating %v %v: %v", tc.updatedCc.Kind, tc.updatedCc.Name, err)
			}

			m := testcontroller.ParseObjects(ctx, t, tc.manifest)
			if tc.cc.GetMode() == "namespaced" {
				if err := handleLifecycles(ctx, t, r, tc.updatedCc, m); err == nil {
					t.Fatalf("got nil, but want to have an error because the controller manager pod per namespace is not deleted")
				}
				// Simulate that CCC controller kicks in and deletes the controller manager pod.
				for _, item := range installedManifest.Items {
					if item.Kind == "Pod" && strings.Contains(item.GetName(), "cnrm-controller-manager") {
						if err := c.Delete(ctx, item.UnstructuredObject()); err != nil {
							t.Fatalf("error deleting %v %v: %v", item.GroupKind(), item.GetName(), err)
						}
					}
				}
			}
			if err := handleLifecycles(ctx, t, r, tc.updatedCc, m); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedObjs := tc.resultsFunc(t, c)
			expectedManifest := testcontroller.ParseObjects(ctx, t, expectedObjs)
			expectedJSON, err := expectedManifest.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			resJSON, err := m.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(resJSON, expectedJSON) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(resJSON, expectedJSON))
			}
			// assert unneeded components are deleted
			unneededComponents := tc.toDeleteObjectsFunc(t, c)
			for _, str := range unneededComponents {
				obj := testcontroller.ToUnstructured(t, str)
				key := types.NamespacedName{Name: obj.GetName(), Namespace: obj.GetNamespace()}
				if err := c.Get(ctx, key, obj); err == nil || !apierrors.IsNotFound(err) {
					t.Fatalf("expect to get %v error for %v %v, but got error: %v", metav1.StatusReasonNotFound, obj.GetKind(), key, err)
				}
			}
		})
	}
}

type testCaseStruct struct {
	name                      string
	cc                        *corev1beta1.ConfigConnector
	cccs                      []corev1beta1.ConfigConnectorContext
	loadedManifest            []string
	resultsFunc               func(t *testing.T, c client.Client) []string
	managerNamespaceIsolation string
}

func handleLifecycles(ctx context.Context, t *testing.T, r *Reconciler, cc *corev1beta1.ConfigConnector, m *manifest.Objects) error {
	t.Helper()

	fn := r.transformForClusterMode()
	if err := fn(ctx, cc, m); err != nil {
		return err
	}
	fn = r.handleConfigConnectorLifecycle()

	return fn(ctx, cc, m)
}

func newConfigConnectorReconciler(c client.Client) *Reconciler {
	return &Reconciler{
		client: c,
		log:    logr.Discard(),
	}
}

func TestSelectingCRDsByVersion(t *testing.T) {
	tests := []struct {
		name              string
		manifests         []string
		version           string
		expectedManifests []string
		hasError          bool
	}{
		{
			name:              "select v1alpha1 CRD from v1alpha1 and v1beta1 CRDs",
			manifests:         testcontroller.GetManifestsWithAlphaAndBetaCRDs(),
			version:           "v1alpha1",
			expectedManifests: testcontroller.GetManifestsWithAlphaCRD(),
		},
		{
			name:              "select v1alpha1 CRD from v1beta1 CRDs",
			manifests:         testcontroller.GetManifestsWithBetaCRD(),
			version:           "v1alpha1",
			expectedManifests: testcontroller.GetManifestsWithNoCRD(),
		},
		{
			name:      "select v1alpha1 CRD from non-KCC CRD",
			manifests: testcontroller.GetManifestsWithNonKCCCRD(),
			version:   "v1alpha1",
			hasError:  true,
		},
		{
			name:      "select v1alpha1 CRD from defective CRD",
			manifests: testcontroller.GetManifestsWithDefectiveCRD(),
			version:   "v1alpha1",
			hasError:  true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			manifests := testcontroller.ParseObjects(ctx, t, tc.manifests)
			r := newConfigConnectorReconciler(c)

			err := r.selectCRDsByVersion(manifests, tc.version)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, want an error")
				}
				return
			} else if err != nil {
				t.Fatalf("error selecting CRDs by version: %v", err)
			}

			processedJSON, err := manifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedManifests := testcontroller.ParseObjects(ctx, t, tc.expectedManifests)
			expectedJSON, err := expectedManifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(processedJSON, expectedJSON) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(processedJSON, expectedJSON))
			}
		})
	}
}

// TestConfigConnectorContextControllerWatchMultipleCustomizationCR verifies the correct behavior of
// the customization watcher of configConnectorContext operator.
func TestConfigConnectorControllerWatchCustomizationCR(t *testing.T) {
	var (
		CC = &corev1beta1.ConfigConnector{
			ObjectMeta: metav1.ObjectMeta{
				Name: corev1beta1.ConfigConnectorAllowedName,
			},
		}
		CR = &customizev1beta1.ControllerResource{
			ObjectMeta: metav1.ObjectMeta{
				Name: "cnrm-webhook-manager",
			},
			Spec: customizev1beta1.ControllerResourceSpec{
				Containers: []customizev1beta1.ContainerResourceSpec{
					{
						Name:      "webhook",
						Resources: customizev1beta1.ResourceRequirements{},
					},
				},
			},
		}
	)

	// test setup
	ctx, cancel := context.WithCancel(context.Background())
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer func() {
		cancel()
		stop()
	}()
	r := newConfigConnectorReconcilerWithCustomizationWatcher(mgr)
	if err := r.customizationWatcher.EnsureWatchStarted(ctx, types.NamespacedName{Namespace: CC.Namespace, Name: CC.Name}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := mgr.GetClient()

	// after a ControllerResource object is created,
	// check that a watch event is raised with the correct name and namespace.
	if err := c.Create(ctx, CR); err != nil {
		t.Fatalf("error creating %v %v/%v: %v", CR.Kind, CR.Namespace, CR.Name, err)
	}
	select { // expect watch event raised on "/configconnector.core.cnrm.cloud.google.com"
	case e := <-r.customizationWatcher.Events():
		if e.Object.GetNamespace() != "" {
			t.Fatalf("unexpected namespace for watch event object, want \"\" (global), got %v", e.Object.GetNamespace())
		}
		if e.Object.GetName() != corev1beta1.ConfigConnectorAllowedName {
			t.Fatalf("unexpected name for watch event object, want %v, got %v", corev1beta1.ConfigConnectorAllowedName, e.Object.GetName())
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("expect watch event, got no event")
	}
	select { // expect no more watch event
	case e := <-r.customizationWatcher.Events():
		t.Fatalf("unexpected watch event object: %v", e.Object)
	default:
	}

	// after delete a ControllerResource, check that a watch event is raised with the correct name and namespace.
	if err := c.Delete(ctx, CR); err != nil {
		t.Fatalf("error deleting %v %v/%v: %v", CR.Kind, CR.Namespace, CR.Name, err)
	}
	select { // expect watch event raised on "/configconnector.core.cnrm.cloud.google.com"
	case e := <-r.customizationWatcher.Events():
		if e.Object.GetNamespace() != "" {
			t.Fatalf("unexpected namespace for watch event object, want \"\" (global), got %v", e.Object.GetNamespace())
		}
		if e.Object.GetName() != corev1beta1.ConfigConnectorAllowedName {
			t.Fatalf("unexpected name for watch event object, want %v, got %v", corev1beta1.ConfigConnectorAllowedName, e.Object.GetName())
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("expect watch event, got no event")
	}
}

// TestApplyFailsForDuplicatedWebhook ensures that applying a webhook configuration CR with duplicate Webhooks fails.
// Due to the current implementation, it is not easy to test the happy path in unit test. We do plan to change the
// implementation and move webhook configurations to be directly owned by the KCC operator.
func TestApplyFailsForDuplicatedWebhook(t *testing.T) {
	tests := []struct {
		name                             string
		validatingWebhookCustomizationCR *customizev1beta1.ValidatingWebhookConfigurationCustomization
		mutatingWebhookCustomizationCR   *customizev1beta1.MutatingWebhookConfigurationCustomization
		expectedCustomizationCRStatus    customizev1beta1.WebhookConfigurationCustomizationStatus
	}{
		{
			name:                             "customize for the same webhook multiple times in ValidatingWebhookCRForDuplicatedWebhook fails",
			validatingWebhookCustomizationCR: testcontroller.ValidatingWebhookCRForDuplicatedWebhook,
			expectedCustomizationCRStatus: customizev1beta1.WebhookConfigurationCustomizationStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy: false,
					Errors:  []string{testcontroller.ErrDuplicatedWebhookForValidatingWebhookCR},
				},
			},
		},
		{
			name:                           "customize for the same webhook multiple times in MutatingWebhookCRForDuplicatedWebhook fails",
			mutatingWebhookCustomizationCR: testcontroller.MutatingWebhookCRForDuplicatedWebhook,
			expectedCustomizationCRStatus: customizev1beta1.WebhookConfigurationCustomizationStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy: false,
					Errors:  []string{testcontroller.ErrDuplicatedWebhookForMutatingWebhookCR},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// test setup
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			if tc.validatingWebhookCustomizationCR != nil {
				cr := tc.validatingWebhookCustomizationCR
				if err := c.Create(ctx, cr); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", cr.Kind, cr.Namespace, cr.Name, err)
				}
			}
			if tc.mutatingWebhookCustomizationCR != nil {
				cr := tc.mutatingWebhookCustomizationCR
				if err := c.Create(ctx, cr); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", cr.Kind, cr.Namespace, cr.Name, err)
				}
			}
			r := newConfigConnectorReconciler(c)

			// run the test function
			if err := r.fetchAndApplyAllWebhookConfigurationCustomizationCRs(ctx); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// check the status of cluster-scoped customization CR
			if tc.validatingWebhookCustomizationCR != nil {
				updatedCR := &customizev1beta1.ValidatingWebhookConfigurationCustomization{}
				if err := c.Get(ctx, types.NamespacedName{Namespace: tc.validatingWebhookCustomizationCR.Namespace, Name: tc.validatingWebhookCustomizationCR.Name}, updatedCR); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				status := updatedCR.Status
				if !reflect.DeepEqual(status, tc.expectedCustomizationCRStatus) {
					t.Fatalf("unexpected diff: %v", cmp.Diff(status, tc.expectedCustomizationCRStatus))
				}
			}
			if tc.mutatingWebhookCustomizationCR != nil {
				updatedCR := &customizev1beta1.MutatingWebhookConfigurationCustomization{}
				if err := c.Get(ctx, types.NamespacedName{Namespace: tc.mutatingWebhookCustomizationCR.Namespace, Name: tc.mutatingWebhookCustomizationCR.Name}, updatedCR); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				status := updatedCR.Status
				if !reflect.DeepEqual(status, tc.expectedCustomizationCRStatus) {
					t.Fatalf("unexpected diff: %v", cmp.Diff(status, tc.expectedCustomizationCRStatus))
				}
			}
		})
	}
}

func TestNamespaceScopedApplyCustomizations(t *testing.T) {
	tests := []struct {
		name                          string
		manifests                     []string
		namespacedCustomizationCR     *customizev1beta1.NamespacedControllerResource
		expectedManifests             []string
		expectedCustomizationCRStatus customizev1beta1.ControllerResourceStatus
		skipCheckingCRStatus          bool
	}{
		{
			name:                      "namespaced customization CR has no effect",
			manifests:                 testcontroller.ClusterModeComponents,
			namespacedCustomizationCR: testcontroller.NamespacedControllerResourceCRForControllerManagerResources,
			expectedManifests:         testcontroller.ClusterModeComponents, // same as the input manifests
			skipCheckingCRStatus:      true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// test setup
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			cr := tc.namespacedCustomizationCR
			testcontroller.EnsureNamespaceExists(c, cr.Namespace)
			if err := c.Create(ctx, cr); err != nil {
				t.Fatalf("error creating %v %v/%v: %v", cr.Kind, cr.Namespace, cr.Name, err)
			}
			manifests := testcontroller.ParseObjects(ctx, t, tc.manifests)
			r := newConfigConnectorReconciler(c)

			// run the test function
			fn := r.applyCustomizations()
			if err := fn(ctx, nil, manifests); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// check the resulting manifests
			gotJSON, err := manifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedManifests := testcontroller.ParseObjects(ctx, t, tc.expectedManifests)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedJSON, err := expectedManifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(gotJSON, expectedJSON) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(gotJSON, expectedJSON))
			}

			// check the status of cluster-scoped customization CR
			if tc.skipCheckingCRStatus {
				return
			}
			updatedCR := &customizev1beta1.ControllerResource{}
			if err := c.Get(ctx, types.NamespacedName{Namespace: tc.namespacedCustomizationCR.Namespace, Name: tc.namespacedCustomizationCR.Name}, updatedCR); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			gotStatus := updatedCR.Status
			expectedStatus := tc.expectedCustomizationCRStatus
			if expectedStatus.ObservedGeneration != 0 {
				expectedStatus.ObservedGeneration = updatedCR.Generation
			}
			if !reflect.DeepEqual(gotStatus, expectedStatus) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(gotStatus, expectedStatus))
			}
		})
	}
}

func TestClusterScopedApplyCustomizations(t *testing.T) {
	tests := []struct {
		name                          string
		manifests                     []string
		clusterScopedCustomizationCR  *customizev1beta1.ControllerResource
		expectedManifests             []string
		expectedCustomizationCRStatus customizev1beta1.ControllerResourceStatus
		skipCheckingCRStatus          bool
	}{
		{
			name:                         "customize the resources for cnrm-controller-manager",
			manifests:                    testcontroller.ClusterModeComponents,
			clusterScopedCustomizationCR: testcontroller.ControllerResourceCRForControllerManagerResources,
			expectedManifests:            testcontroller.ClusterModeComponentsWithCustomizedControllerManager,
			expectedCustomizationCRStatus: customizev1beta1.ControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            true,
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                         "customize the resources for cnrm-controller-manager and check observedGeneration",
			manifests:                    testcontroller.ClusterModeComponents,
			clusterScopedCustomizationCR: testcontroller.ControllerResourceCRForObservedControllerManagerResources,
			expectedManifests:            testcontroller.ClusterModeComponentsWithCustomizedControllerManager,
			expectedCustomizationCRStatus: customizev1beta1.ControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            true,
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                         "customize the resources and replica for cnrm-webhook-manager",
			manifests:                    testcontroller.ClusterModeComponents,
			clusterScopedCustomizationCR: testcontroller.ControllerResourceCRForWebhookManagerResourcesAndReplicas,
			expectedManifests:            testcontroller.ClusterModeComponentsWithCustomizedWebhookManager,
			expectedCustomizationCRStatus: customizev1beta1.ControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            true,
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                         "customize for a non-existing controller fails",
			manifests:                    testcontroller.ClusterModeComponents,
			clusterScopedCustomizationCR: testcontroller.ControllerResourceCRForNonExistingController,
			expectedManifests:            testcontroller.ClusterModeComponents, // same as the input manifests
			expectedCustomizationCRStatus: customizev1beta1.ControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            false,
					Errors:             []string{testcontroller.ErrNonExistingController},
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                         "customize for the same container multiple times in the CR fails",
			manifests:                    testcontroller.ClusterModeComponents,
			clusterScopedCustomizationCR: testcontroller.ControllerResourceCRForDuplicatedContainer,
			expectedManifests:            testcontroller.ClusterModeComponents, // same as the input manifests
			expectedCustomizationCRStatus: customizev1beta1.ControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            false,
					Errors:             []string{testcontroller.ErrDuplicatedContainer},
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                         "customize for a non-existing container in a valid controller fails",
			manifests:                    testcontroller.ClusterModeComponents,
			clusterScopedCustomizationCR: testcontroller.ControllerResourceCRForNonExistingContainer,
			expectedManifests:            testcontroller.ClusterModeComponents, // same as the input manifests
			expectedCustomizationCRStatus: customizev1beta1.ControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            false,
					Errors:             []string{testcontroller.ErrNonExistingContainer},
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                         "customize the replicas for cnrm-controller-manager has no effect",
			manifests:                    testcontroller.ClusterModeComponents,
			clusterScopedCustomizationCR: testcontroller.ControllerResourceCRForControllerManagerReplicas,
			expectedManifests:            testcontroller.ClusterModeComponents, // same as the input manifests
			expectedCustomizationCRStatus: customizev1beta1.ControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            true,
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                         "customize the replicas for cnrm-webhook-manager to a value large than the maxReplicas of HPA",
			manifests:                    testcontroller.ClusterModeComponents,
			clusterScopedCustomizationCR: testcontroller.ControllerResourceCRForWebhookManagerWithLargeReplicas,
			expectedManifests:            testcontroller.ClusterModeComponentsWithCustomizedWebhookManagerWithLargeReplicas,
			expectedCustomizationCRStatus: customizev1beta1.ControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            true,
					ObservedGeneration: 1,
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// test setup
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			cr := tc.clusterScopedCustomizationCR
			if err := c.Create(ctx, cr); err != nil {
				t.Fatalf("error for test(%s) creating %v %v: %v", tc.name, cr.Kind, cr.Name, err)
			}

			manifests := testcontroller.ParseObjects(ctx, t, tc.manifests)
			r := newConfigConnectorReconciler(c)

			// run the test function
			fn := r.applyCustomizations()
			if err := fn(ctx, nil, manifests); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// check the resulting manifests
			gotJSON, err := manifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedManifests := testcontroller.ParseObjects(ctx, t, tc.expectedManifests)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedJSON, err := expectedManifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(gotJSON, expectedJSON) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(gotJSON, expectedJSON))
			}

			// check the status of cluster-scoped customization CR
			if tc.skipCheckingCRStatus {
				return
			}
			updatedCR := tc.clusterScopedCustomizationCR
			if err := c.Get(ctx, types.NamespacedName{Name: tc.clusterScopedCustomizationCR.Name}, updatedCR); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			gotStatus := updatedCR.Status
			expectedStatus := tc.expectedCustomizationCRStatus
			if !reflect.DeepEqual(gotStatus, expectedStatus) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(gotStatus, expectedStatus))
			}
		})
	}
}

func TestApplyRateLimitCustomizations(t *testing.T) {
	tests := []struct {
		name                             string
		manifests                        []string
		controllerReconcilerCR           *customizev1beta1.ControllerReconciler
		namespacedControllerReconcilerCR *customizev1beta1.NamespacedControllerReconciler
		expectedManifests                []string
		skipCheckingCRStatus             bool
		expectedCRStatus                 customizev1beta1.ControllerReconcilerStatus
	}{
		{
			name:                   "customize the rate limit for cnrm-controller-manager",
			manifests:              testcontroller.ClusterModeComponents,
			controllerReconcilerCR: testcontroller.ControllerReconcilerCR,
			expectedManifests:      testcontroller.ClusterModeComponentsWithRatLimitCustomization,
			expectedCRStatus: customizev1beta1.ControllerReconcilerStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy: true,
				},
			},
		},
		{
			name:                   "customize the rate limit for a unsupported controller fails",
			manifests:              testcontroller.ClusterModeComponents,
			controllerReconcilerCR: testcontroller.ControllerReconcilerCRForUnsupportedController,
			expectedManifests:      testcontroller.ClusterModeComponents, // same as the input manifests
			expectedCRStatus: customizev1beta1.ControllerReconcilerStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy: false,
					Errors:  []string{testcontroller.ErrUnsupportedController},
				},
			},
		},
		{
			name:                             "namespaced rate limit CR has no effect in cluster mode",
			manifests:                        testcontroller.ClusterModeComponents,
			namespacedControllerReconcilerCR: testcontroller.NamespacedControllerReconcilerCR,
			expectedManifests:                testcontroller.ClusterModeComponents, // same as the input manifests
			skipCheckingCRStatus:             true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// test setup
			ctx := context.TODO()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			if tc.controllerReconcilerCR != nil {
				cr := tc.controllerReconcilerCR
				if err := c.Create(ctx, cr); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", cr.Kind, cr.Namespace, cr.Name, err)
				}
			}
			if tc.namespacedControllerReconcilerCR != nil {
				cr := tc.namespacedControllerReconcilerCR
				testcontroller.EnsureNamespaceExists(c, cr.Namespace)
				if err := c.Create(ctx, cr); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", cr.Kind, cr.Namespace, cr.Name, err)
				}
			}
			manifests := testcontroller.ParseObjects(ctx, t, tc.manifests)
			r := newConfigConnectorReconciler(c)

			// run the test function
			fn := r.applyCustomizations()
			if err := fn(ctx, nil, manifests); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// check the resulting manifests
			gotJSON, err := manifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedManifests := testcontroller.ParseObjects(ctx, t, tc.expectedManifests)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedJSON, err := expectedManifests.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(gotJSON, expectedJSON) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(gotJSON, expectedJSON))
			}

			// check the status of cluster-scoped customization CR
			if tc.skipCheckingCRStatus {
				return
			}
			updatedCR := &customizev1beta1.ControllerReconciler{}
			if err := c.Get(ctx, types.NamespacedName{Name: tc.controllerReconcilerCR.Name}, updatedCR); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			gotStatus := updatedCR.Status
			if !reflect.DeepEqual(gotStatus, tc.expectedCRStatus) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(gotStatus, tc.expectedCRStatus))
			}
		})
	}
}

func newConfigConnectorReconcilerWithCustomizationWatcher(m ctrl.Manager) *Reconciler {
	r := &Reconciler{
		client: m.GetClient(),
		log:    logr.Discard(),
	}
	r.customizationWatcher = controllers.NewWithDynamicClient(
		dynamic.NewForConfigOrDie(m.GetConfig()),
		controllers.CustomizationWatcherOptions{
			TriggerGVRs: controllers.CustomizationCRsToWatch,
			Log:         logr.Discard(),
		})
	return r
}
