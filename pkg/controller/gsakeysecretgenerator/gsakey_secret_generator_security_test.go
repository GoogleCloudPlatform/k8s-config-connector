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

package gsakeysecretgenerator

import (
	"context"
	"encoding/base64"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/jitter"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// TestReconcile_RefusesToOverwriteForeignSecret asserts the regression for
// CVE-grade Secret takeover: a tenant who creates an IAMServiceAccountKey
// whose name collides with an existing Secret in the same namespace must
// NOT cause the controller to replace that Secret's Data, OwnerReferences,
// and Labels with the SA key material. The fix gates the Update on the
// observed Secret carrying the `managed-by-cnrm: true` label set by this
// controller; foreign Secrets are left untouched and a Warning event is
// recorded.
func TestReconcile_RefusesToOverwriteForeignSecret(t *testing.T) {
	const ns = "tenant-a"
	const collidingName = "production-db-password"
	originalSecretData := []byte("hunter2")

	// A pre-existing Secret in the tenant's namespace that was NOT created
	// by Config Connector (no `managed-by-cnrm` label, no KCC
	// OwnerReferences). This is what the attacker is trying to replace.
	victim := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      collidingName,
			Namespace: ns,
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{"password": originalSecretData},
	}

	// The attacker-supplied IAMServiceAccountKey resource. Its status.privateKey
	// is what the controller would otherwise stuff into the colliding Secret.
	sakAPIVersion := "iam.cnrm.cloud.google.com/v1beta1"
	sakKind := "IAMServiceAccountKey"
	attackerKey := base64.StdEncoding.EncodeToString([]byte("ATTACKER_GCP_KEY"))
	sak := &unstructured.Unstructured{Object: map[string]interface{}{
		"apiVersion": sakAPIVersion,
		"kind":       sakKind,
		"metadata": map[string]interface{}{
			"name":      collidingName,
			"namespace": ns,
			"uid":       "attacker-uid",
		},
		"status": map[string]interface{}{
			"privateKey": attackerKey,
		},
	}}

	scheme := runtime.NewScheme()
	if err := clientscheme.AddToScheme(scheme); err != nil {
		t.Fatalf("scheme: %v", err)
	}
	cl := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(victim).
		WithRuntimeObjects(sak).
		Build()

	r := &ReconcileSecret{
		Client:     cl,
		kind:       sakKind,
		apiVersion: sakAPIVersion,
		recorder:   record.NewFakeRecorder(8),
		jitterGen:  &jitter.SimpleJitterGenerator{},
	}

	_, err := r.Reconcile(context.Background(), reconcile.Request{
		NamespacedName: types.NamespacedName{Namespace: ns, Name: collidingName},
	})
	if err == nil {
		t.Fatalf("reconcile must refuse to overwrite the foreign Secret, got nil error")
	}
	if !strings.Contains(err.Error(), "refusing to overwrite Secret") {
		t.Fatalf("expected refusal error, got: %v", err)
	}

	// Confirm the victim Secret is intact.
	got := &corev1.Secret{}
	if err := cl.Get(context.Background(), ctrlclient.ObjectKey{Namespace: ns, Name: collidingName}, got); err != nil {
		t.Fatalf("get victim: %v", err)
	}
	if string(got.Data["password"]) != string(originalSecretData) {
		t.Fatalf("victim Secret password was modified: got %q want %q", got.Data["password"], originalSecretData)
	}
	if _, ok := got.Data["key.json"]; ok {
		t.Fatalf("victim Secret was overwritten with attacker SA key (key.json present)")
	}
	if got.Labels[label.CnrmManagedKey] == "true" {
		t.Fatalf("victim Secret was re-labelled by attacker reconcile")
	}
	if len(got.OwnerReferences) != 0 {
		t.Fatalf("victim Secret OwnerReferences were modified: %+v", got.OwnerReferences)
	}
}

// TestReconcile_UpdatesOwnSecret asserts the non-regression: when the
// pre-existing Secret IS a KCC-managed Secret (carrying the
// `managed-by-cnrm: true` label written by this controller's Create path),
// rotation correctly replaces the key material.
func TestReconcile_UpdatesOwnSecret(t *testing.T) {
	const ns = "tenant-a"
	const name = "my-sa-key"
	oldKey := base64.StdEncoding.EncodeToString([]byte("OLD_KEY"))
	newKey := base64.StdEncoding.EncodeToString([]byte("NEW_KEY"))

	kccSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Labels:    map[string]string{label.CnrmManagedKey: "true"},
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{"key.json": []byte(oldKey)},
	}

	sak := &unstructured.Unstructured{Object: map[string]interface{}{
		"apiVersion": "iam.cnrm.cloud.google.com/v1beta1",
		"kind":       "IAMServiceAccountKey",
		"metadata": map[string]interface{}{
			"name":      name,
			"namespace": ns,
			"uid":       "kcc-uid",
		},
		"status": map[string]interface{}{
			"privateKey": newKey,
		},
	}}

	scheme := runtime.NewScheme()
	if err := clientscheme.AddToScheme(scheme); err != nil {
		t.Fatalf("scheme: %v", err)
	}
	cl := fake.NewClientBuilder().WithScheme(scheme).
		WithObjects(kccSecret).
		WithRuntimeObjects(sak).
		Build()

	r := &ReconcileSecret{
		Client:     cl,
		kind:       "IAMServiceAccountKey",
		apiVersion: "iam.cnrm.cloud.google.com/v1beta1",
		recorder:   record.NewFakeRecorder(8),
		jitterGen:  &jitter.SimpleJitterGenerator{},
	}

	if _, err := r.Reconcile(context.Background(), reconcile.Request{
		NamespacedName: types.NamespacedName{Namespace: ns, Name: name},
	}); err != nil {
		t.Fatalf("expected rotation to succeed, got: %v", err)
	}

	got := &corev1.Secret{}
	if err := cl.Get(context.Background(), ctrlclient.ObjectKey{Namespace: ns, Name: name}, got); err != nil {
		t.Fatalf("get: %v", err)
	}
	want, _ := base64.StdEncoding.DecodeString(newKey)
	if string(got.Data["key.json"]) != string(want) {
		t.Fatalf("rotation did not replace key.json: got %q", got.Data["key.json"])
	}
}
