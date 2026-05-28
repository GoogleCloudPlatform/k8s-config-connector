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

package iapclientsecretgenerator

import (
	"context"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestReconcile(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)

	kind := "IAPIdentityAwareProxyClient"
	apiVersion := "iap.cnrm.cloud.google.com/v1beta1"
	name := "test-client"
	namespace := "test-ns"
	clientId := "test-client-id"
	clientSecret := "test-client-secret"

	u := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": apiVersion,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"name":      name,
				"namespace": namespace,
				"uid":       "test-uid",
			},
			"spec": map[string]interface{}{
				"resourceID": clientId,
			},
			"status": map[string]interface{}{
				"secret": clientSecret,
			},
		},
	}

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(u).Build()

	r := &ReconcileSecret{
		Client:     fakeClient,
		kind:       kind,
		apiVersion: apiVersion,
		recorder:   record.NewFakeRecorder(10),
		jitterGen:  &jitter.SimpleJitterGenerator{},
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		},
	}

	_, err := r.Reconcile(ctx, req)
	if err != nil {
		t.Fatalf("Reconcile failed: %v", err)
	}

	// Verify secret creation
	secret := &corev1.Secret{}
	err = fakeClient.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, secret)
	if err != nil {
		t.Fatalf("Failed to get secret: %v", err)
	}

	if string(secret.Data["client_id"]) != clientId {
		t.Errorf("Expected client_id %s, got %s", clientId, string(secret.Data["client_id"]))
	}
	if string(secret.Data["client_secret"]) != clientSecret {
		t.Errorf("Expected client_secret %s, got %s", clientSecret, string(secret.Data["client_secret"]))
	}

	// Verify owner reference
	if len(secret.OwnerReferences) != 1 {
		t.Errorf("Expected 1 owner reference, got %d", len(secret.OwnerReferences))
	} else {
		owner := secret.OwnerReferences[0]
		if owner.Kind != kind || owner.Name != name {
			t.Errorf("Unexpected owner reference: %+v", owner)
		}
	}

	// Test update
	newClientSecret := "new-test-client-secret"
	u.Object["status"].(map[string]interface{})["secret"] = newClientSecret
	err = fakeClient.Update(ctx, u)
	if err != nil {
		t.Fatalf("Failed to update IAP client: %v", err)
	}

	_, err = r.Reconcile(ctx, req)
	if err != nil {
		t.Fatalf("Reconcile failed on update: %v", err)
	}

	err = fakeClient.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, secret)
	if err != nil {
		t.Fatalf("Failed to get secret after update: %v", err)
	}

	if string(secret.Data["client_secret"]) != newClientSecret {
		t.Errorf("Expected updated client_secret %s, got %s", newClientSecret, string(secret.Data["client_secret"]))
	}

	// Test disable annotation
	u.SetAnnotations(map[string]string{createIAPClientSecretAnnotation: "false"})
	err = fakeClient.Update(ctx, u)
	if err != nil {
		t.Fatalf("Failed to update IAP client with annotation: %v", err)
	}

	// Delete secret to see if it's recreated (it shouldn't be)
	err = fakeClient.Delete(ctx, secret)
	if err != nil {
		t.Fatalf("Failed to delete secret: %v", err)
	}

	_, err = r.Reconcile(ctx, req)
	if err != nil {
		t.Fatalf("Reconcile failed with annotation: %v", err)
	}

	err = fakeClient.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, secret)
	if !errors.IsNotFound(err) {
		t.Errorf("Expected secret to not be found after disabling, got error: %v", err)
	}
}
