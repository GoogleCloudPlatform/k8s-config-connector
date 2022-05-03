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

package configconnectorcontext

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
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
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func TestRemovingStaleComponents(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	c := mgr.GetClient()
	testcontroller.EnsureNamespaceExists(mgr.GetClient(), k8s.OperatorSystemNamespace)
	testcontroller.EnsureNamespaceExists(mgr.GetClient(), k8s.CNRMSystemNamespace)
	staleComponents := []string{`
apiVersion: v1
kind: Service
metadata:
 labels:
   cnrm.cloud.google.com/monitored: "true"
   cnrm.cloud.google.com/scoped-namespace: foo-ns
   cnrm.cloud.google.com/system: "true"
 name: cnrm-manager-foo
 namespace: cnrm-system
spec:
 ports:
 - name: controller-manager
   port: 443
 - name: metrics
   port: 8888
 selector:
   cnrm.cloud.google.com/component: cnrm-controller-manager
   cnrm.cloud.google.com/scoped-namespace: foo-ns
   cnrm.cloud.google.com/system: "true"
`, `
apiVersion: apps/v1
kind: StatefulSet
metadata:
 labels:
   cnrm.cloud.google.com/component: cnrm-controller-manager
   cnrm.cloud.google.com/scoped-namespace: foo-ns
   cnrm.cloud.google.com/system: "true"
 name: cnrm-controller-manager-foo
 namespace: cnrm-system
spec:
 selector:
   matchLabels:
     cnrm.cloud.google.com/component: cnrm-controller-manager
     cnrm.cloud.google.com/scoped-namespace: foo-ns
     cnrm.cloud.google.com/system: "true"
 serviceName: cnrm-manager-foo
 template:
   metadata:
     labels:
       cnrm.cloud.google.com/component: cnrm-controller-manager
       cnrm.cloud.google.com/scoped-namespace: foo-ns
       cnrm.cloud.google.com/system: "true"
`}

	for _, str := range staleComponents {
		u := testcontroller.ToUnstructured(t, str)
		if err := c.Create(ctx, u); err != nil {
			t.Fatalf("error creating object %v/%v: %v", u.GetNamespace(), u.GetName(), err)
		}
	}
	ccc := &corev1beta1.ConfigConnectorContext{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      k8s.ConfigConnectorContextAllowedName,
			Namespace: "foo-ns",
		},
		Spec: corev1beta1.ConfigConnectorContextSpec{
			GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
		},
	}

	m := testcontroller.ParseObjects(t, ctx, testcontroller.GetPerNamespaceManifest())
	_, err := transformNamespacedComponentTemplates(ctx, mgr.GetClient(), ccc, m.Items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, str := range staleComponents {
		u := testcontroller.ToUnstructured(t, str)
		key := client.ObjectKey{
			Namespace: u.GetNamespace(),
			Name:      u.GetName(),
		}
		err := c.Get(ctx, key, u)
		if err == nil {
			t.Fatalf("expect object %v/%v: to be not found", u.GetNamespace(), u.GetName())
		}
		if !apierrors.IsNotFound(err) {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}

func TestHandlePerNamespaceComponentsCreate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		cc             *corev1beta1.ConfigConnector
		ccc            *corev1beta1.ConfigConnectorContext
		loadedManifest []string
		resultsFunc    func(t *testing.T, c client.Client) []string
		hasError       bool
	}{
		{
			name: "CC is in cluster mode, CCC surfaces errors",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "cluster",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      k8s.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			hasError:       true,
		},
		{
			name: "CC is in namespaced mode, CCC has spec.requestProjectPolicy omitted",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      k8s.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerNamespaceManifest(), "foo-ns", "foo@bar.iam.gserviceaccount.com", false, "", c)
			},
		},
		{
			name: "CC is in namespaced mode, CCC has spec.requestProjectPolicy set to RESOURCE_PROJECT",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      k8s.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "RESOURCE_PROJECT",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerNamespaceManifest(), "foo-ns", "foo@bar.iam.gserviceaccount.com", true, "", c)
			},
		},

		{
			name: "CC is in namespaced mode, CCC has spec.billingProject set and spec.requestProjectPolicy set to BILLING_PROJECT",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      k8s.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					RequestProjectPolicy: "BILLING_PROJECT",
					BillingProject:       "BILL_ME",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerNamespaceManifest(), "foo-ns", "foo@bar.iam.gserviceaccount.com", true, "BILL_ME", c)
			},
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
			m := testcontroller.ParseObjects(t, ctx, tc.loadedManifest)
			r := newConfigConnectorReconciler(c)

			if err := c.Create(ctx, tc.cc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
			}
			testcontroller.EnsureNamespaceExists(c, tc.ccc.Namespace)
			if err := c.Create(ctx, tc.ccc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.ccc.Kind, tc.ccc.Name, err)
			}
			err := handleLifecycles(t, ctx, r, tc.ccc, m)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but want an error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectedObjs := tc.resultsFunc(t, c)
			expectedManifest := testcontroller.ParseObjects(t, ctx, expectedObjs)
			expectedJson, err := expectedManifest.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			resJson, err := m.JSONManifest()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(resJson, expectedJson) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(resJson, expectedJson))
			}

			// Verify that the CC object is NOT attached finalizers by the CCC controller.
			cc := &corev1beta1.ConfigConnector{}
			ccKey := client.ObjectKeyFromObject(tc.cc)
			if err := c.Get(ctx, ccKey, cc); err != nil {
				t.Fatalf("error getting ConfigConnector %v: %v", ccKey, err)
			}
			if testcontroller.HasOperatorFinalizer(cc) {
				t.Fatalf("%v finalizer was found in %v", k8s.OperatorFinalizer, cc)
			}

			ccc := &corev1beta1.ConfigConnectorContext{}
			contextKey := client.ObjectKeyFromObject(tc.ccc)
			if err := c.Get(ctx, contextKey, ccc); err != nil {
				t.Fatalf("error getting ConfigConnectorContext %v: %v", contextKey, err)
			}
			if !testcontroller.HasOperatorFinalizer(ccc) {
				t.Fatalf("no %v finalizer was found in %v", k8s.OperatorFinalizer, ccc)
			}
		})
	}
}

func TestHandlePerNamespaceComponentsDelete(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                 string
		cc                   *corev1beta1.ConfigConnector
		ccc                  *corev1beta1.ConfigConnectorContext
		loadedManifest       []string
		installedObjectsFunc func(t *testing.T, c client.Client) []string
		resultsFunc          func(t *testing.T, c client.Client) []string
		issueCCCDeletion     bool
		issueCCDeletion      bool
		hasError             bool
	}{
		{
			name: "Delete the CCC object",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       k8s.ConfigConnectorContextAllowedName,
					Namespace:  "foo-ns",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerNamespaceManifest(), "foo-ns", "foo@bar.iam.gserviceaccount.com", false, "", c)
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			issueCCCDeletion: true,
		},
		{
			name: "CC is switched to cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: k8s.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode:                 "cluster",
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       k8s.ConfigConnectorContextAllowedName,
					Namespace:  "foo-ns",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerNamespaceManifest(), "foo-ns", "foo@bar.iam.gserviceaccount.com", false, "", c)
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			hasError: true,
		},
		{
			name: "CC is pending deletion",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       k8s.ConfigConnectorAllowedName,
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       k8s.ConfigConnectorContextAllowedName,
					Namespace:  "foo-ns",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerNamespaceManifest(), "foo-ns", "foo@bar.iam.gserviceaccount.com", false, "", c)
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			issueCCDeletion: true,
			hasError:        true,
		},
		{
			name: "CC is not found",
			cc:   nil,
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       k8s.ConfigConnectorContextAllowedName,
					Namespace:  "foo-ns",
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerNamespaceManifest(), "foo-ns", "foo@bar.iam.gserviceaccount.com", false, "", c)
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			hasError: true,
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
			contextKey := client.ObjectKeyFromObject(tc.ccc)
			testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)
			testcontroller.EnsureNamespaceExists(c, k8s.CNRMSystemNamespace)
			m := testcontroller.ParseObjects(t, ctx, tc.loadedManifest)
			r := newConfigConnectorReconciler(c)
			if tc.cc != nil {
				if err := c.Create(ctx, tc.cc); err != nil {
					t.Fatalf("error creating %v %v: %v", tc.cc.Kind, tc.cc.Name, err)
				}
			}

			testcontroller.EnsureNamespaceExists(c, tc.ccc.Namespace)
			if err := c.Create(ctx, tc.ccc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.ccc.Kind, tc.ccc.Name, err)
			}
			installedObjs := tc.installedObjectsFunc(t, c)
			for _, obj := range installedObjs {
				u := testcontroller.ToUnstructured(t, obj)
				if err := c.Create(ctx, u); err != nil && !apierrors.IsAlreadyExists(err) {
					t.Fatalf("error creating %v %v/%v: %v", u.GetObjectKind(), u.GetNamespace(), u.GetName(), err)
				}
			}

			if tc.issueCCDeletion {
				if err := c.Delete(ctx, tc.cc); err != nil {
					t.Fatalf("error deleting %v %v: %v", tc.cc.Kind, tc.cc.GetName(), err)
				}
			}

			// Issue the delete request for the CCC object per test case.
			if tc.issueCCCDeletion {
				if err := c.Delete(ctx, tc.ccc); err != nil {
					t.Fatalf("error deleting %v %v: %v", tc.ccc.Kind, contextKey, err)
				}
				if err := c.Get(ctx, contextKey, tc.ccc); err != nil {
					t.Fatalf("error getting %v %v: %v", tc.ccc.Kind, contextKey, err)
				}
			}

			// Handle the lifecycle of CCC.
			// If error is expected, assert that there is an error returned.
			// Otherwise, assert that the finalized objects are matching with the expect the result.
			err := handleLifecycles(t, ctx, r, tc.ccc, m)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but want an error")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				expectedObjs := tc.resultsFunc(t, c)
				expectedManifest := testcontroller.ParseObjects(t, ctx, expectedObjs)
				expectedJson, err := expectedManifest.JSONManifest()
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				resJson, err := m.JSONManifest()
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if !reflect.DeepEqual(resJson, expectedJson) {
					t.Fatalf("unexpected diff: %v", cmp.Diff(resJson, expectedJson))
				}
			}

			// Assert that previously installed objects are deleted.
			for _, obj := range installedObjs {
				u := testcontroller.ToUnstructured(t, obj)
				key := types.NamespacedName{Name: u.GetName(), Namespace: u.GetNamespace()}
				if err := c.Get(ctx, key, u); err == nil || !apierrors.IsNotFound(err) {
					t.Fatalf("got error: %v, expect to get %v error for %v %v, ", err, metav1.StatusReasonNotFound, u.GetKind(), key)
				}
			}

			// Assert that the CCC object is deleted if the deletion request was issued.
			// Otherwise, assert that the operator finalizer is removed.
			if tc.issueCCCDeletion {
				if err := c.Get(ctx, contextKey, tc.ccc); err == nil || !apierrors.IsNotFound(err) {
					t.Fatalf("got error: %v, expect to get %v error for %v ", err, metav1.StatusReasonNotFound, tc.ccc)
				}
			} else {
				ccc := &corev1beta1.ConfigConnectorContext{}
				if err := c.Get(ctx, contextKey, ccc); err != nil {
					t.Fatalf("error getting ConfigConnectorContext %v: %v", contextKey, err)
				}
				if testcontroller.HasOperatorFinalizer(ccc) {
					t.Fatalf("%v finalizer was found in %v", k8s.OperatorFinalizer, ccc)
				}
			}
		})
	}
}

func TestHandleReconcileFailed(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	c := mgr.GetClient()
	mockEventRecorder := testmocks.NewMockEventRecorder(t, mgr.GetScheme())
	r := ConfigConnectorContextReconciler{
		client:   c,
		recorder: mockEventRecorder,
		log:      logr.Discard(),
	}

	ccc := &corev1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k8s.ConfigConnectorContextAllowedName,
			Namespace: "foo-ns",
		},
		Spec: corev1beta1.ConfigConnectorContextSpec{
			GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
		},
	}
	nn := client.ObjectKeyFromObject(ccc)
	_, kind := corev1beta1.ConfigConnectorContextGroupVersionKind.ToAPIVersionAndKind()

	testcontroller.EnsureNamespaceExists(c, "foo-ns")
	if err := c.Create(ctx, ccc); err != nil {
		t.Fatalf("failed to create ConfigConnectorContext: %v", err)
	}
	reconcileErr := fmt.Errorf("reconciliation error")
	if err := r.handleReconcileFailed(ctx, nn, reconcileErr); err != nil {
		t.Errorf("error handling failed reconciliation: %v", err)
	}

	expectedErrMsg := fmt.Sprintf(k8s.ReconcileErrMsgTmpl, reconcileErr)
	mockEventRecorder.AssertEventRecorded(kind, nn, v1.EventTypeWarning, k8s.UpdateFailed, expectedErrMsg)

	newCCC := &corev1beta1.ConfigConnectorContext{}
	if err := c.Get(ctx, nn, newCCC); err != nil {
		t.Errorf("failed to get ConfigConnectorContext after attempt to handle failed reconciliation: %v", err)
	}
	status := newCCC.GetCommonStatus()
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
	r := ConfigConnectorContextReconciler{
		client:   c,
		recorder: mockEventRecorder,
		log:      logr.Discard(),
	}

	ccc := &corev1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:      k8s.ConfigConnectorContextAllowedName,
			Namespace: "foo-ns",
		},
		Spec: corev1beta1.ConfigConnectorContextSpec{
			GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
		},
	}
	nn := client.ObjectKeyFromObject(ccc)
	_, kind := corev1beta1.ConfigConnectorContextGroupVersionKind.ToAPIVersionAndKind()

	testcontroller.EnsureNamespaceExists(c, "foo-ns")
	if err := c.Create(ctx, ccc); err != nil {
		t.Fatalf("failed to create ConfigConnectorContext: %v", err)
	}
	if err := r.handleReconcileSucceeded(ctx, nn); err != nil {
		t.Errorf("error handling successful reconciliation: %v", err)
	}
	mockEventRecorder.AssertEventRecorded(kind, nn, v1.EventTypeNormal, k8s.UpToDate, k8s.UpToDateMessage)

	newCCC := &corev1beta1.ConfigConnectorContext{}
	if err := c.Get(ctx, nn, newCCC); err != nil {
		t.Errorf("failed to get ConfigConnectorContext after attempt to handle failed reconciliation: %v", err)
	}
	status := newCCC.GetCommonStatus()
	if !status.Healthy {
		t.Errorf("unexpected value for status.healthy: got 'false', want 'true'")
	}
	if len(status.Errors) != 0 {
		t.Errorf("unexpected number of errors in status.errors: got %v errors, want 0 errors. Got the errors: %v", len(status.Errors), status.Errors)
	}
}

func handleLifecycles(t *testing.T, ctx context.Context,
	r *ConfigConnectorContextReconciler, ccc *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	t.Helper()
	fn := r.transformNamespacedComponents()
	if err := fn(ctx, ccc, m); err != nil {
		return err
	}
	fn = r.addLabels()
	if err := fn(ctx, ccc, m); err != nil {
		return err
	}
	fn = r.handleCCContextLifecycle()
	if err := fn(ctx, ccc, m); err != nil {
		return err
	}
	return nil
}

func newConfigConnectorReconciler(c client.Client) *ConfigConnectorContextReconciler {
	return &ConfigConnectorContextReconciler{
		client:     c,
		log:        logr.Discard(),
		labelMaker: SourceLabel(),
	}
}
