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

package kccmanager

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
)

func TestSyncerIntegration_RetryOnConflict(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mclModPath, err := getModulePath("github.com/gke-labs/multicluster-leader-election")
	if err != nil {
		t.Fatalf("error getting module path: %v", err)
	}
	kubeEtlModPath, err := getModulePath("github.com/gke-labs/kube-etl/syncer")
	if err != nil {
		t.Fatalf("error getting module path: %v", err)
	}

	testEnv := &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join(repo.GetRootOrTestFatal(t), "operator/config/crd/bases"),
			filepath.Join(mclModPath, "config/crd/bases"),
			filepath.Join(kubeEtlModPath, "config/crd"),
		},
		ErrorIfCRDPathMissing: true,
	}

	cfg, err := testEnv.Start()
	if err != nil {
		t.Fatalf("error starting envtest: %v", err)
	}
	defer func() {
		if err := testEnv.Stop(); err != nil {
			t.Errorf("error stopping envtest: %v", err)
		}
	}()

	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = mclv1alpha1.AddToScheme(scheme)

	kubeClient, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
	_ = kubeClient.Create(ctx, ns)

	name := types.NamespacedName{Name: "test-syncer-retry", Namespace: "default"}

	// Pre-create the Syncer object
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KRMSyncerGVK)
	u.SetName(name.Name)
	u.SetNamespace(name.Namespace)
	_ = unstructured.SetNestedField(u.Object, false, "spec", "suspend")
	_ = unstructured.SetNestedField(u.Object, "dummy", "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
	_ = unstructured.SetNestedSlice(u.Object, []interface{}{
		map[string]interface{}{
			"group":   "example.com",
			"version": "v1",
			"kind":    "Foo",
		},
	}, "spec", "rules")

	if err := kubeClient.Create(ctx, u); err != nil {
		t.Fatalf("error creating syncer: %v", err)
	}

	si := &SyncerIntegration{
		client:          kubeClient,
		apiReader:       kubeClient, // use non-cached client directly
		leaseNN:         name,
		replicationMode: "Status",
	}

	// Fetch the object to get its ResourceVersion
	latest := &unstructured.Unstructured{}
	latest.SetGroupVersionKind(KRMSyncerGVK)
	if err := kubeClient.Get(ctx, name, latest); err != nil {
		t.Fatalf("failed to get latest syncer: %v", err)
	}

	// Update it out-of-band to create a conflict condition for the next update
	outOfBand := latest.DeepCopy()
	_ = unstructured.SetNestedField(outOfBand.Object, "push", "spec", "mode")
	if err := kubeClient.Update(ctx, outOfBand); err != nil {
		t.Fatalf("failed to out-of-band update syncer: %v", err)
	}

	// Call EnsureSuspended. It should fetch, see conflict, retry, and succeed.
	if err := si.EnsureSuspended(ctx); err != nil {
		t.Fatalf("EnsureSuspended failed on conflict: %v", err)
	}

	// Verify it was actually suspended
	updated := &unstructured.Unstructured{}
	updated.SetGroupVersionKind(KRMSyncerGVK)
	if err := kubeClient.Get(ctx, name, updated); err != nil {
		t.Fatalf("failed to get updated syncer: %v", err)
	}
	suspend, _, _ := unstructured.NestedBool(updated.Object, "spec", "suspend")
	if !suspend {
		t.Errorf("expected syncer to be suspended after conflict retry")
	}
}
