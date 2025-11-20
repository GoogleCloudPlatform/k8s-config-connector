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

package configconnector

import (
	"context"
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

	// Wait for CRD to be established
	// In a real env we should wait, but envtest might be fast enough or we can retry.
	// Let's just proceed and see if it works, or add a small retry loop for the first VPA creation.

	// Ensure namespaces exist
	testcontroller.EnsureNamespaceExists(c, k8s.CNRMSystemNamespace)
	testcontroller.EnsureNamespaceExists(c, k8s.OperatorSystemNamespace)

	// Create ConfigConnector
	cc := &corev1beta1.ConfigConnector{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-kcc-vpa",
		},
		Spec: corev1beta1.ConfigConnectorSpec{
			Mode:                 "cluster",
			GoogleServiceAccount: "foo@bar.iam.gserviceaccount.com",
		},
	}
	if err := c.Create(ctx, cc); err != nil {
		t.Fatalf("error creating ConfigConnector: %v", err)
	}

	// Create ControllerResource with VPA enabled
	vpaEnabled := true
	cr := &customizev1beta1.ControllerResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cnrm-controller-manager",
		},
		Spec: customizev1beta1.ControllerResourceSpec{
			VerticalPodAutoscalerEnabled: &vpaEnabled,
		},
	}
	if err := c.Create(ctx, cr); err != nil {
		t.Fatalf("error creating ControllerResource: %v", err)
	}

	// Create VPA object with recommendations
	vpaGVK := schema.GroupVersionKind{
		Group:   "autoscaling.k8s.io",
		Version: "v1",
		Kind:    "VerticalPodAutoscaler",
	}
	vpa := &unstructured.Unstructured{}
	vpa.SetGroupVersionKind(vpaGVK)
	// Trigger reconciliation to create VPA
	r := newConfigConnectorReconciler(c)
	m := testcontroller.ParseObjects(ctx, t, testcontroller.GetClusterModeWorkloadIdentityManifest())
	if err := handleLifecycles(ctx, t, r, cc, m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Apply customizations explicitly to trigger VPA creation
	if err := r.applyCustomizations()(ctx, cc, m); err != nil {
		t.Fatalf("unexpected error applying customizations: %v", err)
	}

	// Verify VPA object is created
	vpaKey := client.ObjectKey{
		Namespace: k8s.CNRMSystemNamespace,
		Name:      "cnrm-controller-manager",
	}

	// Retry getting VPA in case CRD is not yet established or reconciliation is still in progress
	for i := 0; i < 10; i++ {
		err := c.Get(ctx, vpaKey, vpa)
		if err == nil {
			break
		}
		if apierrors.IsNotFound(err) {
			t.Logf("VPA not found yet, retrying... (%d/%d)", i+1, 10)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		t.Fatalf("error getting VPA: %v", err)
	}
	if err := c.Get(ctx, vpaKey, vpa); err != nil {
		t.Fatalf("expected VPA to be created, but got error after retries: %v", err)
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
		// If status update fails (e.g. because VPA CRD status subresource is not enabled in test env), try Update.
		if err := c.Update(ctx, vpa); err != nil {
			t.Fatalf("error updating VPA status: %v", err)
		}
	}

	// Trigger reconciliation again to apply recommendations
	if err := r.applyCustomizations()(ctx, cc, m); err != nil {
		t.Fatalf("unexpected error applying customizations 2nd time: %v", err)
	}

	// Verify StatefulSet resources in the manifest
	var sts *appsv1.StatefulSet
	for _, obj := range m.Items {
		if obj.GroupKind() == appsv1.SchemeGroupVersion.WithKind("StatefulSet").GroupKind() &&
			obj.GetName() == "cnrm-controller-manager" &&
			obj.GetNamespace() == k8s.CNRMSystemNamespace {
			u := obj.UnstructuredObject()
			sts = &appsv1.StatefulSet{}
			if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, sts); err != nil {
				t.Fatalf("error converting unstructured to StatefulSet: %v", err)
			}
			break
		}
	}

	if sts == nil {
		t.Fatalf("StatefulSet cnrm-controller-manager not found in manifest")
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
