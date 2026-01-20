// Copyright 2025 Google LLC
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
	"strings"
	"testing"
	"time"

	customizev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	corev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/k8s"
	testcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/controller"
	testmain "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/test/main"
	appsv1 "k8s.io/api/apps/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestVPAIntegration(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mgr, stop := testmain.StartTestManagerFromNewTestEnv()
	defer stop()
	c := mgr.GetClient()

	// Create VPA CRD
	vpaCRD := &apiextensionsv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "verticalpodautoscalers.autoscaling.k8s.io",
			Annotations: map[string]string{
				"api-approved.kubernetes.io": "https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler",
			},
		},
		Spec: apiextensionsv1.CustomResourceDefinitionSpec{
			Group: "autoscaling.k8s.io",
			Names: apiextensionsv1.CustomResourceDefinitionNames{
				Kind:     "VerticalPodAutoscaler",
				ListKind: "VerticalPodAutoscalerList",
				Plural:   "verticalpodautoscalers",
				Singular: "verticalpodautoscaler",
			},
			Scope: apiextensionsv1.NamespaceScoped,
			Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &apiextensionsv1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]apiextensionsv1.JSONSchemaProps{
								"spec": {
									Type:                   "object",
									XPreserveUnknownFields: boolPtr(true),
								},
								"status": {
									Type:                   "object",
									XPreserveUnknownFields: boolPtr(true),
								},
							},
						},
					},
					Subresources: &apiextensionsv1.CustomResourceSubresources{
						Status: &apiextensionsv1.CustomResourceSubresourceStatus{},
					},
				},
			},
		},
	}
	if err := c.Create(ctx, vpaCRD); err != nil {
		// Ignore if already exists (parallel tests might race)
		if !apierrors.IsAlreadyExists(err) {
			t.Fatalf("error creating VPA CRD: %v", err)
		}
	}

	// Ensure namespaces exist
	testcontroller.EnsureNamespaceExists(c, k8s.CNRMSystemNamespace)
	testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)
	testNamespace := "test-ns-vpa"
	testcontroller.EnsureNamespaceExists(c, testNamespace)

	// Create ConfigConnector in namespaced mode
	cc := &corev1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-kcc-vpa-namespaced",
		},
		Spec: corev1beta1.ConfigConnectorSpec{
			Mode: "namespaced",
		},
	}
	if err := c.Create(ctx, cc); err != nil {
		t.Fatalf("error creating ConfigConnector: %v", err)
	}

	// Create ConfigConnectorContext
	ccc := &corev1beta1.ConfigConnectorContext{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "configconnectorcontext.core.cnrm.cloud.google.com",
			Namespace: testNamespace,
		},
		Spec: corev1beta1.ConfigConnectorContextSpec{
			GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
		},
	}
	if err := c.Create(ctx, ccc); err != nil {
		t.Fatalf("error creating ConfigConnectorContext: %v", err)
	}

	// Create NamespacedControllerResource with VPA enabled
	vpaEnabled := customizev1beta1.VPAModeEnabled
	ncr := &customizev1beta1.NamespacedControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cnrm-controller-manager",
			Namespace: testNamespace,
		},
		Spec: customizev1beta1.NamespacedControllerResourceSpec{
			VerticalPodAutoscalerMode: &vpaEnabled,
		},
	}
	if err := c.Create(ctx, ncr); err != nil {
		t.Fatalf("error creating NamespacedControllerResource: %v", err)
	}

	// Create VPA object with recommendations
	vpaGVK := schema.GroupVersionKind{
		Group:   "autoscaling.k8s.io",
		Version: "v1",
		Kind:    "VerticalPodAutoscaler",
	}
	vpa := &unstructured.Unstructured{}
	vpa.SetGroupVersionKind(vpaGVK)

	// Trigger reconciliation
	r, err := newReconciler(mgr, &ReconcilerOptions{
		RepoPath:                  "", // Use default or empty if not needed for this test
		ManagerNamespaceIsolation: k8s.ManagerNamespaceIsolationShared,
	})
	if err != nil {
		t.Fatalf("error creating reconciler: %v", err)
	}

	// We need to simulate the customization watcher or manually trigger the apply
	// The real controller uses a watcher to trigger reconciliation.
	// Here we can just call the transform functions directly or use the reconciler's internal methods if accessible,
	// but since we are testing integration, we might want to rely on the manager if possible,
	// OR we can manually invoke the transformation logic like in other tests.

	// However, the `Reconciler` struct has `applyNamespacedCustomizations` which calls `applyNamespacedControllerResource`.
	// We can try to invoke that.

	m := testcontroller.ParseObjects(ctx, t, testcontroller.GetPerNamespaceManifest())

	// We need to make sure the manifest loader works or we provide the objects.
	// `GetPerNamespaceManifest` returns the default templates.

	// The `handleCCContextLifecycle` calls `applyNamespacedCustomizations`.
	// Let's try to run the transformation chain.

	// We need to set up the context object for the transform function
	// The transform function expects `declarative.DeclarativeObject` which `ConfigConnectorContext` implements.

	// Transform components first (to set namespace etc)
	if err := r.transformNamespacedComponents()(ctx, ccc, m); err != nil {
		t.Fatalf("unexpected error transforming components: %v", err)
	}

	// Let's use `applyNamespacedCustomizations` directly to see if it creates VPA.
	if err := r.applyNamespacedCustomizations()(ctx, ccc, m); err != nil {
		t.Fatalf("unexpected error applying customizations: %v", err)
	}

	// Find the transformed StatefulSet name from the manifest
	var stsName string
	var stsNamespace string
	for _, obj := range m.Items {
		if obj.GroupKind() == appsv1.SchemeGroupVersion.WithKind("StatefulSet").GroupKind() &&
			strings.HasPrefix(obj.GetName(), "cnrm-controller-manager") {
			stsName = obj.GetName()
			stsNamespace = obj.GetNamespace()
			break
		}
	}
	if stsName == "" {
		t.Fatalf("StatefulSet cnrm-controller-manager not found in transformed manifest")
	}

	// Verify VPA object is created in the same namespace as the StatefulSet with the correct name
	vpaKey := client.ObjectKey{
		Namespace: stsNamespace,
		Name:      stsName,
	}
	vpa.SetName(stsName) // Update vpa object name for subsequent updates
	vpa.SetNamespace(stsNamespace)

	if err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 10*time.Second, true, func(ctx context.Context) (bool, error) {
		err := c.Get(ctx, vpaKey, vpa)
		if err == nil {
			return true, nil
		}
		if apierrors.IsNotFound(err) {
			t.Logf("VPA %s/%s not found yet, retrying...", stsNamespace, stsName)
			return false, nil
		}
		return false, err
	}); err != nil {
		t.Fatalf("expected VPA %s/%s to be created, but got error after retries: %v", stsNamespace, stsName, err)
	}

	// Update VPA status with recommendations
	vpa.Object["status"] = map[string]interface{}{
		"recommendation": map[string]interface{}{
			"containerRecommendations": []interface{}{
				map[string]interface{}{
					"containerName": "manager",
					"target": map[string]interface{}{
						"cpu":    "500m",
						"memory": "1Gi",
					},
				},
			},
		},
	}
	if err := c.Status().Update(ctx, vpa); err != nil {
		if err := c.Update(ctx, vpa); err != nil {
			t.Fatalf("error updating VPA status: %v", err)
		}
	}

	// Reset manifest and re-apply transformations and customizations
	m = testcontroller.ParseObjects(ctx, t, testcontroller.GetPerNamespaceManifest())
	if err := r.transformNamespacedComponents()(ctx, ccc, m); err != nil {
		t.Fatalf("unexpected error transforming components 2nd time: %v", err)
	}
	if err := r.applyNamespacedCustomizations()(ctx, ccc, m); err != nil {
		t.Fatalf("unexpected error applying customizations 2nd time: %v", err)
	}

	// Verify StatefulSet resources in the manifest
	var sts *appsv1.StatefulSet
	for _, obj := range m.Items {
		if obj.GroupKind() == appsv1.SchemeGroupVersion.WithKind("StatefulSet").GroupKind() &&
			(obj.GetName() == "cnrm-controller-manager" || len(obj.GetName()) > len("cnrm-controller-manager") && obj.GetName()[:len("cnrm-controller-manager")] == "cnrm-controller-manager") {
			u := obj.UnstructuredObject()
			sts = &appsv1.StatefulSet{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, sts); err != nil {
				t.Fatalf("error converting unstructured to StatefulSet: %v", err)
			}
			break
		}
	}

	if sts == nil {
		// Debug: print all objects in manifest
		var names []string
		for _, obj := range m.Items {
			names = append(names, fmt.Sprintf("%s/%s", obj.Kind, obj.GetName()))
		}
		t.Fatalf("StatefulSet cnrm-controller-manager not found in manifest. Available objects: %v", names)
	}

	container := sts.Spec.Template.Spec.Containers[0]
	if container.Name != "manager" {
		t.Fatalf("expected first container to be 'manager', got '%s'", container.Name)
	}

	expectedCPU := resource.MustParse("500m")
	expectedMemory := resource.MustParse("1Gi")

	if !container.Resources.Requests.Cpu().Equal(expectedCPU) {
		t.Errorf("expected CPU request %s, got %s", expectedCPU.String(), container.Resources.Requests.Cpu().String())
	}
	if !container.Resources.Limits.Cpu().Equal(expectedCPU) {
		t.Errorf("expected CPU limit %s, got %s", expectedCPU.String(), container.Resources.Limits.Cpu().String())
	}
	if !container.Resources.Requests.Memory().Equal(expectedMemory) {
		t.Errorf("expected Memory request %s, got %s", expectedMemory.String(), container.Resources.Requests.Memory().String())
	}
	if !container.Resources.Limits.Memory().Equal(expectedMemory) {
		t.Errorf("expected Memory limit %s, got %s", expectedMemory.String(), container.Resources.Limits.Memory().String())
	}
}

func boolPtr(b bool) *bool {
	return &b
}
