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

package controllers_test

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/customize/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers"
	"github.com/google/go-cmp/cmp"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestEnsureVPAForStatefulSet(t *testing.T) {
	sts := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-sts",
			Namespace: "test-ns",
		},
	}

	c := fake.NewClientBuilder().Build()

	ctx := context.TODO()
	if err := controllers.EnsureVPAForStatefulSet(ctx, c, sts, v1beta1.VPAModeEnabled); err != nil {
		t.Fatalf("EnsureVPAForStatefulSet failed: %v", err)
	}

	vpa := &unstructured.Unstructured{}
	vpa.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "autoscaling.k8s.io",
		Version: "v1",
		Kind:    "VerticalPodAutoscaler",
	})

	key := client.ObjectKey{Namespace: "test-ns", Name: "test-sts"}
	if err := c.Get(context.TODO(), key, vpa); err != nil {
		t.Fatalf("error getting VPA: %v", err)
	}

	expectedSpec := map[string]interface{}{
		"targetRef": map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "StatefulSet",
			"name":       "test-sts",
		},
		"updatePolicy": map[string]interface{}{
			"updateMode":  "Auto",
			"minReplicas": int64(1),
		},
	}

	if diff := cmp.Diff(expectedSpec, vpa.Object["spec"]); diff != "" {
		t.Errorf("unexpected VPA spec (-want +got):\n%s", diff)
	}
}

func TestEnsureVPAForDeployment(t *testing.T) {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-deployment",
			Namespace: "test-ns",
		},
	}

	c := fake.NewClientBuilder().Build()

	ctx := context.TODO()
	if err := controllers.EnsureVPAForDeployment(ctx, c, deployment, v1beta1.VPAModeEnabled); err != nil {
		t.Fatalf("EnsureVPAForDeployment failed: %v", err)
	}

	vpa := &unstructured.Unstructured{}
	vpa.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "autoscaling.k8s.io",
		Version: "v1",
		Kind:    "VerticalPodAutoscaler",
	})

	key := client.ObjectKey{Namespace: "test-ns", Name: "test-deployment"}
	if err := c.Get(context.TODO(), key, vpa); err != nil {
		t.Fatalf("error getting VPA: %v", err)
	}

	expectedSpec := map[string]interface{}{
		"targetRef": map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"name":       "test-deployment",
		},
		"updatePolicy": map[string]interface{}{
			"updateMode":  "Auto",
			"minReplicas": int64(1),
		},
	}

	if diff := cmp.Diff(expectedSpec, vpa.Object["spec"]); diff != "" {
		t.Errorf("unexpected VPA spec (-want +got):\n%s", diff)
	}
}
