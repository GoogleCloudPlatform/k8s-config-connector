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
	"time"

	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
	testmocks "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/mocks"
	k8sreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/manifest"
)

func TestRemovingStaleComponents(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                   string
		ccc                    *corev1beta1.ConfigConnectorContext
		transformationFunction func(ctx context.Context, c client.Client, ccc *corev1beta1.ConfigConnectorContext, namespacedTemplates []*manifest.Object) ([]*manifest.Object, error)
		staleComponents        []string
	}{
		{
			name: "namespaced mode",
			ccc: &corev1beta1.ConfigConnectorContext{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "foo-ns",
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			transformationFunction: transformNamespacedComponentTemplates,
			staleComponents: []string{`
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
`},
		},
		{
			name: "per namespace mode",
			ccc: &corev1beta1.ConfigConnectorContext{
				TypeMeta: metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "t1234-tenant0-provider",
					Labels: map[string]string{
						"tenancy.gke.io/access-level": "tenant",
						"tenancy.gke.io/project":      "t1234",
						"tenancy.gke.io/tenant":       "t1234-tenant0",
					},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					ManagerNamespace:     "t1234-tenant0-supervisor",
				},
			},
			transformationFunction: transformPerNamespaceComponentTemplates,
			staleComponents: []string{`
apiVersion: v1
kind: Service
metadata:
 labels:
   cnrm.cloud.google.com/monitored: "true"
   cnrm.cloud.google.com/scoped-namespace: t1234-tenant0-provider
   cnrm.cloud.google.com/system: "true"
 name: cnrm-manager-foo
 namespace: t1234-tenant0-supervisor
spec:
 ports:
 - name: controller-manager
   port: 443
 - name: metrics
   port: 8888
 selector:
   cnrm.cloud.google.com/component: cnrm-controller-manager
   cnrm.cloud.google.com/scoped-namespace: t1234-tenant0-provider
   cnrm.cloud.google.com/system: "true"
`, `
apiVersion: apps/v1
kind: StatefulSet
metadata:
 labels:
   cnrm.cloud.google.com/component: cnrm-controller-manager
   cnrm.cloud.google.com/scoped-namespace: t1234-tenant0-provider
   cnrm.cloud.google.com/system: "true"
 name: cnrm-controller-manager-foo
 namespace: t1234-tenant0-supervisor
spec:
 selector:
   matchLabels:
     cnrm.cloud.google.com/component: cnrm-controller-manager
     cnrm.cloud.google.com/scoped-namespace: t1234-tenant0-provider
     cnrm.cloud.google.com/system: "true"
 serviceName: cnrm-manager-foo
 template:
   metadata:
     labels:
       cnrm.cloud.google.com/component: cnrm-controller-manager
       cnrm.cloud.google.com/scoped-namespace: t1234-tenant0-provider
       cnrm.cloud.google.com/system: "true"
`},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			mgr, stop := testmain.StartTestManagerFromNewTestEnv()
			defer stop()
			c := mgr.GetClient()
			testcontroller.EnsureNamespaceExists(mgr.GetClient(), k8s.OperatorSystemNamespace)
			testcontroller.EnsureNamespaceExists(mgr.GetClient(), k8s.CNRMSystemNamespace)

			for _, str := range tc.staleComponents {
				u := testcontroller.ToUnstructured(t, str)
				testcontroller.EnsureNamespaceExists(mgr.GetClient(), u.GetNamespace())
				if err := c.Create(ctx, u); err != nil {
					t.Fatalf("error creating object %v/%v: %v", u.GetNamespace(), u.GetName(), err)
				}
			}

			m := testcontroller.ParseObjects(ctx, t, testcontroller.GetPerNamespaceManifest())
			_, err := tc.transformationFunction(ctx, mgr.GetClient(), tc.ccc, m.Items)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			for _, str := range tc.staleComponents {
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
		})
	}
}

func TestHandlePerNamespaceComponentsCreate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                      string
		cc                        *corev1beta1.ConfigConnector
		ccc                       *corev1beta1.ConfigConnectorContext
		loadedManifest            []string
		resultsFunc               func(t *testing.T, c client.Client) []string
		hasError                  bool
		managerNamespaceIsolation string
	}{
		{
			name: "CC is in cluster mode, CCC surfaces errors",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "cluster",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
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
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
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
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
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
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
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

		{
			name: "CC is in per namespaced mode, CCC has error spec.managerNamespace omitted",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "t1234-tenant0-provider",
					Labels: map[string]string{
						"tenancy.gke.io/access-level": "tenant",
						"tenancy.gke.io/project":      "t1234",
						"tenancy.gke.io/tenant":       "t1234-tenant0",
					},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			loadedManifest:            testcontroller.GetPerNamespaceManifest(),
			hasError:                  true,
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
		{
			name: "CC is in namespaced mode, CCC has error spec.managerNamespace specified",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "t1234-tenant0-provider",
					Labels: map[string]string{
						"tenancy.gke.io/access-level": "tenant",
						"tenancy.gke.io/project":      "t1234",
						"tenancy.gke.io/tenant":       "t1234-tenant0",
					},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					ManagerNamespace:     "t1234-tenant0-supervisor",
				},
			},
			loadedManifest:            testcontroller.GetPerNamespaceManifest(),
			hasError:                  true,
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationShared,
		},
		{
			name: "CC is in per namespaced mode, CCC has spec.managerNamespace",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:      corev1beta1.ConfigConnectorContextAllowedName,
					Namespace: "t1234-tenant0-provider",
					Labels: map[string]string{
						"tenancy.gke.io/access-level": "tenant",
						"tenancy.gke.io/project":      "t1234",
						"tenancy.gke.io/tenant":       "t1234-tenant0",
					},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					ManagerNamespace:     "t1234-tenant0-supervisor",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerManagerNamespaceManifest(), "t1234-tenant0-provider", "foo@bar.iam.gserviceaccount.com", false, "", c)
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
			testcontroller.EnsureNamespaceExists(c, tc.ccc.Namespace)
			if err := c.Create(ctx, tc.ccc); err != nil {
				t.Fatalf("error creating %v %v: %v", tc.ccc.Kind, tc.ccc.Name, err)
			}
			err := handleLifecycles(ctx, t, r, tc.ccc, m)
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
		name                      string
		cc                        *corev1beta1.ConfigConnector
		ccc                       *corev1beta1.ConfigConnectorContext
		loadedManifest            []string
		installedObjectsFunc      func(t *testing.T, c client.Client) []string
		resultsFunc               func(t *testing.T, c client.Client) []string
		issueCCCDeletion          bool
		issueCCDeletion           bool
		hasError                  bool
		managerNamespaceIsolation string
	}{
		{
			name: "Delete the CCC object",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorContextAllowedName,
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
			name: "Delete the CCC object per namespace mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorContextAllowedName,
					Namespace:  "t1234-tenant0-provider",
					Finalizers: []string{k8s.OperatorFinalizer},
					Labels: map[string]string{
						"tenancy.gke.io/access-level": "tenant",
						"tenancy.gke.io/project":      "t1234",
						"tenancy.gke.io/tenant":       "t1234-tenant0",
					},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					ManagerNamespace:     "t1234-tenant0-supervisor",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerManagerNamespaceManifest(), "t1234-tenant0-provider", "foo@bar.iam.gserviceaccount.com", false, "", c)
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			issueCCCDeletion:          true,
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
		{
			name: "CC is switched to cluster mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode:                 "cluster",
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorContextAllowedName,
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
			name: "CC is switched to cluster mode from per namespace mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name: corev1beta1.ConfigConnectorAllowedName,
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode:                 "cluster",
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorContextAllowedName,
					Namespace:  "t1234-tenant0-provider",
					Finalizers: []string{k8s.OperatorFinalizer},
					Labels: map[string]string{
						"tenancy.gke.io/access-level": "tenant",
						"tenancy.gke.io/project":      "t1234",
						"tenancy.gke.io/tenant":       "t1234-tenant0",
					},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					ManagerNamespace:     "t1234-tenant0-supervisor",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerManagerNamespaceManifest(), "t1234-tenant0-provider", "foo@bar.iam.gserviceaccount.com", false, "", c)
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			hasError:                  true,
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
		{
			name: "CC is pending deletion",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorAllowedName,
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorContextAllowedName,
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
			name: "CC is pending deletion per namespace mode",
			cc: &corev1beta1.ConfigConnector{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorAllowedName,
					Finalizers: []string{k8s.OperatorFinalizer},
				},
				Spec: corev1beta1.ConfigConnectorSpec{
					Mode: "namespaced",
				},
			},
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorContextAllowedName,
					Namespace:  "t1234-tenant0-provider",
					Finalizers: []string{k8s.OperatorFinalizer},
					Labels: map[string]string{
						"tenancy.gke.io/access-level": "tenant",
						"tenancy.gke.io/project":      "t1234",
						"tenancy.gke.io/tenant":       "t1234-tenant0",
					},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					ManagerNamespace:     "t1234-tenant0-supervisor",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerManagerNamespaceManifest(), "t1234-tenant0-provider", "foo@bar.iam.gserviceaccount.com", false, "", c)
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			issueCCDeletion:           true,
			hasError:                  true,
			managerNamespaceIsolation: k8s.ManagerNamespaceIsolationDedicated,
		},
		{
			name: "CC is not found",
			cc:   nil,
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorContextAllowedName,
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
			name: "CC is not found per namepace mode",
			cc:   nil,
			ccc: &corev1beta1.ConfigConnectorContext{
				ObjectMeta: metav1.ObjectMeta{
					Name:       corev1beta1.ConfigConnectorContextAllowedName,
					Namespace:  "t1234-tenant0-provider",
					Finalizers: []string{k8s.OperatorFinalizer},
					Labels: map[string]string{
						"tenancy.gke.io/access-level": "tenant",
						"tenancy.gke.io/project":      "t1234",
						"tenancy.gke.io/tenant":       "t1234-tenant0",
					},
				},
				Spec: corev1beta1.ConfigConnectorContextSpec{
					GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
					ManagerNamespace:     "t1234-tenant0-supervisor",
				},
			},
			loadedManifest: testcontroller.GetPerNamespaceManifest(),
			installedObjectsFunc: func(t *testing.T, c client.Client) []string {
				return testcontroller.ManuallyModifyNamespaceTemplates(t, testcontroller.GetPerManagerNamespaceManifest(), "t1234-tenant0-provider", "foo@bar.iam.gserviceaccount.com", false, "", c)
			},
			resultsFunc: func(t *testing.T, c client.Client) []string {
				return nil
			},
			hasError:                  true,
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
			contextKey := client.ObjectKeyFromObject(tc.ccc)
			testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)
			testcontroller.EnsureNamespaceExists(c, k8s.CNRMSystemNamespace)
			if tc.ccc.Spec.ManagerNamespace != "" {
				testcontroller.EnsureNamespaceExists(c, tc.ccc.Spec.ManagerNamespace)
			}
			m := testcontroller.ParseObjects(ctx, t, tc.loadedManifest)
			r := newConfigConnectorReconciler(c)
			if tc.managerNamespaceIsolation == k8s.ManagerNamespaceIsolationDedicated {
				r.managerNamespaceIsolation = k8s.ManagerNamespaceIsolationDedicated
			} else {
				r.managerNamespaceIsolation = k8s.ManagerNamespaceIsolationShared
			}
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
			err := handleLifecycles(ctx, t, r, tc.ccc, m)
			if tc.hasError {
				if err == nil {
					t.Fatalf("got nil, but want an error")
				}
			} else {
				if err != nil {
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
	r := Reconciler{
		client:   c,
		recorder: mockEventRecorder,
		log:      logr.Discard(),
	}

	ccc := &corev1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:       corev1beta1.ConfigConnectorContextAllowedName,
			Namespace:  "foo-ns",
			Generation: 1,
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

	expectedErrMsg := "error during reconciliation: reconciliation error"
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
	if status.ObservedGeneration != 1 {
		t.Errorf("unexpected value for status.observedGeneration: got %v, want 1", status.ObservedGeneration)
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
		log:      logr.Discard(),
	}

	ccc := &corev1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:       corev1beta1.ConfigConnectorContextAllowedName,
			Namespace:  "foo-ns",
			Generation: 1,
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
	if status.ObservedGeneration != 1 {
		t.Errorf("unexpected value for status.observedGeneration: got %v, want 1", status.ObservedGeneration)
	}
}

// TestConfigConnectorContextControllerWatchMultipleCustomizationCR creates 2 namespaces verifies the correct behavior of
// the customization watcher of configConnectorContext operator.
func TestConfigConnectorContextControllerWatchMultipleCustomizationCR(t *testing.T) {
	var (
		fooCCC = &corev1beta1.ConfigConnectorContext{
			ObjectMeta: metav1.ObjectMeta{
				Name:      corev1beta1.ConfigConnectorContextAllowedName,
				Namespace: "foo-ns",
			},
		}
		barCCC = &corev1beta1.ConfigConnectorContext{
			ObjectMeta: metav1.ObjectMeta{
				Name:      corev1beta1.ConfigConnectorContextAllowedName,
				Namespace: "foo-ns",
			},
		}
		fooCR = &customizev1beta1.NamespacedControllerResource{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "cnrm-controller-manager",
				Namespace: "foo-ns",
			},
			Spec: customizev1beta1.NamespacedControllerResourceSpec{
				Containers: []customizev1beta1.ContainerResourceSpec{
					{
						Name:      "manager",
						Resources: customizev1beta1.ResourceRequirements{},
					},
				},
			},
		}
		barCR = &customizev1beta1.NamespacedControllerResource{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "cnrm-controller-manager",
				Namespace: "bar-ns",
			},
			Spec: customizev1beta1.NamespacedControllerResourceSpec{
				Containers: []customizev1beta1.ContainerResourceSpec{
					{
						Name:      "manager",
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
	r := newConfigConnectorContextReconcilerWithCustomizationWatcher(mgr)
	if err := r.customizationWatcher.EnsureWatchStarted(ctx, types.NamespacedName{Namespace: fooCCC.Namespace, Name: fooCCC.Name}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := r.customizationWatcher.EnsureWatchStarted(ctx, types.NamespacedName{Namespace: barCCC.Namespace, Name: barCCC.Name}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := mgr.GetClient()

	// after a namespacedControllerResource object is created in "foo-ns" namespace,
	// check that a watch event is raised with the correct name and namespace.
	testcontroller.EnsureNamespaceExists(c, fooCR.Namespace)
	if err := c.Create(ctx, fooCR); err != nil {
		t.Fatalf("error creating %v %v/%v: %v", fooCR.Kind, fooCR.Namespace, fooCR.Name, err)
	}
	select { // expect watch event raised on "foo-ns/configconnectorcontext.core.cnrm.cloud.google.com"
	case e := <-r.customizationWatcher.Events():
		if e.Object.GetNamespace() != "foo-ns" {
			t.Fatalf("unexpected namespace for watch event object, want foo-ns, got %v", e.Object.GetNamespace())
		}
		if e.Object.GetName() != corev1beta1.ConfigConnectorContextAllowedName {
			t.Fatalf("unexpected name for watch event object, want %v, got %v", corev1beta1.ConfigConnectorContextAllowedName, e.Object.GetName())
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("expect watch event, got no event")
	}
	select { // expect no more watch event
	case e := <-r.customizationWatcher.Events():
		t.Fatalf("unexpected watch event object: %v", e.Object)
	default:
	}

	// after a namespacedControllerResource object is created in "bar-ns" namespace,
	// check that a watch event is raised with the correct name and namespace.
	testcontroller.EnsureNamespaceExists(c, barCR.Namespace)
	if err := c.Create(ctx, barCR); err != nil {
		t.Fatalf("error creating %v %v/%v: %v", barCR.Kind, barCR.Namespace, barCR.Name, err)
	}
	select { // expect watch event raised on "bar-ns/configconnectorcontext.core.cnrm.cloud.google.com"
	case e := <-r.customizationWatcher.Events():
		if e.Object.GetNamespace() != "bar-ns" {
			t.Fatalf("unexpected namespace for watch event object, want foo-ns, got %v", e.Object.GetNamespace())
		}
		if e.Object.GetName() != corev1beta1.ConfigConnectorContextAllowedName {
			t.Fatalf("unexpected name for watch event object, want %v, got %v", corev1beta1.ConfigConnectorContextAllowedName, e.Object.GetName())
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("expect watch event, got no event")
	}
	select { // expect no more watch event
	case e := <-r.customizationWatcher.Events():
		t.Fatalf("unexpected watch event: %v", e)
	default:
	}

	// after delete a namespacedControllerResource, check that a watch event is raised with the correct name and namespace.
	if err := c.Delete(ctx, fooCR); err != nil {
		t.Fatalf("error deleting %v %v/%v: %v", fooCR.Kind, fooCR.Namespace, fooCR.Name, err)
	}
	select { // expect watch event raised on "foo-ns/configconnectorcontext.core.cnrm.cloud.google.com"
	case e := <-r.customizationWatcher.Events():
		if e.Object.GetNamespace() != "foo-ns" {
			t.Fatalf("unexpected namespace for watch event object, want foo-ns, got %v", e.Object.GetNamespace())
		}
		if e.Object.GetName() != corev1beta1.ConfigConnectorContextAllowedName {
			t.Fatalf("unexpected name for watch event object, want %v, got %v", corev1beta1.ConfigConnectorContextAllowedName, e.Object.GetName())
		}
	case <-time.After(3 * time.Second):
		t.Fatalf("expect watch event, got no event")
	}
}

func TestApplyNamespacedCustomizations(t *testing.T) {
	ccc := &corev1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:      corev1beta1.ConfigConnectorContextAllowedName,
			Namespace: "foo-ns",
		},
	}
	tests := []struct {
		name                          string
		manifests                     []string
		namespacedCustomizationCR     *customizev1beta1.NamespacedControllerResource
		clusterScopedCustomizationCR  *customizev1beta1.ControllerResource
		expectedManifests             []string
		expectedCustomizationCRStatus customizev1beta1.NamespacedControllerResourceStatus
		skipCheckingCRStatus          bool
	}{
		{
			name:                      "customize the resources for cnrm-controller-manager",
			manifests:                 testcontroller.NamespacedComponents,
			namespacedCustomizationCR: testcontroller.NamespacedControllerResourceCRForControllerManagerResources,
			expectedManifests:         testcontroller.NamespacedComponentsWithCustomizedControllerManager,
			expectedCustomizationCRStatus: customizev1beta1.NamespacedControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            true,
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                      "customize for a non-existing controller fails",
			manifests:                 testcontroller.NamespacedComponents,
			namespacedCustomizationCR: testcontroller.NamespacedControllerResourceCRForNonExistingController,
			expectedManifests:         testcontroller.NamespacedComponents, // same as the input manifests
			expectedCustomizationCRStatus: customizev1beta1.NamespacedControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            false,
					Errors:             []string{testcontroller.ErrNonExistingController},
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                      "customize for a non-existing container in a valid controller fails",
			manifests:                 testcontroller.NamespacedComponents,
			namespacedCustomizationCR: testcontroller.NamespacedControllerResourceCRForNonExistingContainer,
			expectedManifests:         testcontroller.NamespacedComponents, // same as the input manifests
			expectedCustomizationCRStatus: customizev1beta1.NamespacedControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy:            false,
					Errors:             []string{testcontroller.ErrNonExistingContainer},
					ObservedGeneration: 1,
				},
			},
		},
		{
			name:                         "cluster-scoped customization CR has no effect",
			manifests:                    testcontroller.NamespacedComponents,
			clusterScopedCustomizationCR: testcontroller.ControllerResourceCRForControllerManagerResources,
			expectedManifests:            testcontroller.NamespacedComponents, // same as the input manifests
			skipCheckingCRStatus:         true,
		},
		{
			name:                      "customization from a different namespace has no effect",
			manifests:                 testcontroller.NamespacedComponents,
			namespacedCustomizationCR: testcontroller.NamespacedControllerResourceCRWrongNamespace,
			expectedManifests:         testcontroller.NamespacedComponents, // same as the input manifests
			expectedCustomizationCRStatus: customizev1beta1.NamespacedControllerResourceStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					ObservedGeneration: 0,
				}, // no update to status because it is not in the same namespace as the CCC reconciler.
			},
		},
		{
			name:      "observedGeneration is correctly set in status",
			manifests: testcontroller.NamespacedComponents,
			namespacedCustomizationCR: func() *customizev1beta1.NamespacedControllerResource {
				cr := testcontroller.NamespacedControllerResourceCRForControllerManagerResources.DeepCopy()
				cr.Generation = 1
				return cr
			}(),
			expectedManifests: testcontroller.NamespacedComponentsWithCustomizedControllerManager,
			expectedCustomizationCRStatus: customizev1beta1.NamespacedControllerResourceStatus{
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
			if tc.namespacedCustomizationCR != nil {
				cr := tc.namespacedCustomizationCR
				testcontroller.EnsureNamespaceExists(c, cr.Namespace)
				if err := c.Create(ctx, cr); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", cr.Kind, cr.Namespace, cr.Name, err)
				}
			}
			if tc.clusterScopedCustomizationCR != nil {
				cr := tc.clusterScopedCustomizationCR
				if err := c.Create(ctx, cr); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", cr.Kind, cr.Namespace, cr.Name, err)
				}
			}
			manifests := testcontroller.ParseObjects(ctx, t, tc.manifests)
			r := newConfigConnectorReconciler(c)

			// run the test function
			fn := r.applyNamespacedCustomizations()
			if err := fn(ctx, ccc, manifests); err != nil {
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

			// check the status of namespaced customization CR
			if tc.skipCheckingCRStatus {
				return
			}
			updatedCR := &customizev1beta1.NamespacedControllerResource{}
			if err := c.Get(ctx, types.NamespacedName{Namespace: tc.namespacedCustomizationCR.Namespace, Name: tc.namespacedCustomizationCR.Name}, updatedCR); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			gotStatus := updatedCR.Status
			if !reflect.DeepEqual(gotStatus, tc.expectedCustomizationCRStatus) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(gotStatus, tc.expectedCustomizationCRStatus))
			}
		})
	}
}

func TestApplyNamespacedRateLimitCustomizations(t *testing.T) {
	ccc := &corev1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:      corev1beta1.ConfigConnectorContextAllowedName,
			Namespace: "foo-ns",
		},
	}
	tests := []struct {
		name                   string
		manifests              []string
		controllerReconcilerCR *customizev1beta1.NamespacedControllerReconciler
		expectedManifests      []string
		expectedCRStatus       customizev1beta1.NamespacedControllerReconcilerStatus
	}{
		{
			name:                   "customize the rate limit for cnrm-controller-manager",
			manifests:              testcontroller.NamespacedComponents,
			controllerReconcilerCR: testcontroller.NamespacedControllerReconcilerCR,
			expectedManifests:      testcontroller.NamespacedComponentsWithRatLimitCustomization,
			expectedCRStatus: customizev1beta1.NamespacedControllerReconcilerStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy: true,
				},
			},
		},
		{
			name:                   "customize the rate limit for a unsupported controller fails",
			manifests:              testcontroller.NamespacedComponents,
			controllerReconcilerCR: testcontroller.NamespacedControllerReconcilerCRForUnsupportedController,
			expectedManifests:      testcontroller.NamespacedComponents, // same as the input manifests
			expectedCRStatus: customizev1beta1.NamespacedControllerReconcilerStatus{
				CommonStatus: addonv1alpha1.CommonStatus{
					Healthy: false,
					Errors:  []string{testcontroller.ErrUnsupportedController},
				},
			},
		},
		{
			name:                   "customization from a different namespace has no effect",
			manifests:              testcontroller.NamespacedComponents,
			controllerReconcilerCR: testcontroller.NamespacedControllerReconcilerCRWrongNamespace,
			expectedManifests:      testcontroller.NamespacedComponents, // same as the input manifests
			expectedCRStatus: customizev1beta1.NamespacedControllerReconcilerStatus{
				CommonStatus: addonv1alpha1.CommonStatus{}, // no update to status because it is not in the same namespace as the CCC reconciler.
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
			if tc.controllerReconcilerCR != nil {
				cr := tc.controllerReconcilerCR
				testcontroller.EnsureNamespaceExists(c, cr.Namespace)
				if err := c.Create(ctx, cr); err != nil {
					t.Fatalf("error creating %v %v/%v: %v", cr.Kind, cr.Namespace, cr.Name, err)
				}
			}
			manifests := testcontroller.ParseObjects(ctx, t, tc.manifests)
			r := newConfigConnectorReconciler(c)

			// run the test function
			fn := r.applyNamespacedCustomizations()
			if err := fn(ctx, ccc, manifests); err != nil {
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

			// check the status of namespaced customization CR
			updatedCR := &customizev1beta1.NamespacedControllerReconciler{}
			if err := c.Get(ctx, types.NamespacedName{Namespace: tc.controllerReconcilerCR.Namespace, Name: tc.controllerReconcilerCR.Name}, updatedCR); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			gotStatus := updatedCR.Status
			if !reflect.DeepEqual(gotStatus, tc.expectedCRStatus) {
				t.Fatalf("unexpected diff: %v", cmp.Diff(gotStatus, tc.expectedCRStatus))
			}
		})
	}
}

func handleLifecycles(ctx context.Context, t *testing.T,
	r *Reconciler, ccc *corev1beta1.ConfigConnectorContext, m *manifest.Objects) error {
	t.Helper()
	fn := r.transformNamespacedComponents()
	if err := fn(ctx, ccc, m); err != nil {
		return err
	}
	if r.managerNamespaceIsolation == k8s.ManagerNamespaceIsolationDedicated {
		fn = r.transformPerNamespaceComponents()
		if err := fn(ctx, ccc, m); err != nil {
			return err
		}
	}
	fn = r.addLabels()
	if err := fn(ctx, ccc, m); err != nil {
		return err
	}
	fn = r.handleCCContextLifecycle()
	return fn(ctx, ccc, m)
}

func newConfigConnectorReconciler(c client.Client) *Reconciler {
	return &Reconciler{
		client:     c,
		log:        logr.Discard(),
		labelMaker: SourceLabel(),
	}
}

func newConfigConnectorContextReconcilerWithCustomizationWatcher(m ctrl.Manager) *Reconciler {
	r := &Reconciler{
		client: m.GetClient(),
		log:    logr.Discard(),
	}
	r.customizationWatcher = controllers.NewWithDynamicClient(
		dynamic.NewForConfigOrDie(m.GetConfig()),
		controllers.CustomizationWatcherOptions{
			TriggerGVRs: controllers.NamespacedCustomizationCRsToWatch,
			Log:         logr.Discard(),
		})
	return r
}

func TestControllerOverridesField(t *testing.T) {
	t.Parallel()
	ccc := &corev1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:      corev1beta1.ConfigConnectorContextAllowedName,
			Namespace: "foo-ns",
		},
		Spec: corev1beta1.ConfigConnectorContextSpec{
			Experiments: &corev1beta1.Experiments{
				ControllerOverrides: map[string]k8sreconciler.ReconcilerType{
					"BigQueryDataset.bigquery.cnrm.cloud.google.com":              "direct",    // default is terraform
					"AlloyDBClusters.alloydb.cnrm.cloud.google.com":               "direct",    // default is terraform
					"CloudIdentityGroups.cloudidentity.cnrm.cloud.google.com":     "terraform", // default is direct
					"SQLInstances.sql.cnrm.cloud.google.com":                      "terraform", // default is direct
					"CloudIdentityMembership.cloudidentity.cnrm.cloud.google.com": "dcl",       // default is direct
					"UnknownResource.example.cnrm.cloud.google.com":               "direct",    // invalid resource
					"AlloyDBInstance.alloydb.cnrm.cloud.google.com":               "unknown",   // invalid controller type
					"example.cnrm.cloud.google.com":                               "unknown",   // invalid resource and invalid controller
					"SpannerInstance.spanner.cnrm.cloud.google.com":               "terraform", // default is terraform
					"GKEHubFeature.gkehub.cnrm.cloud.google.com":                  "dcl",       // default is dcl
					"ComputeInstance.compute.cnrm.cloud.google.com":               "direct",    // direct is not supported for this resource
				},
			},
		},
	}

	unstructuredCCC, err := runtime.DefaultUnstructuredConverter.ToUnstructured(ccc)
	if err != nil {
		t.Fatalf("error converting to unstructured: %v", err)
	}

	_, found, err := unstructured.NestedMap(unstructuredCCC, "spec", "experiments", "controllerOverrides")
	if err != nil {
		t.Fatalf("error getting nested map: %v", err)
	}
	if !found {
		t.Fatalf("field .spec.experiments.controllerOverrides not found in unstructured object")
	}
}
