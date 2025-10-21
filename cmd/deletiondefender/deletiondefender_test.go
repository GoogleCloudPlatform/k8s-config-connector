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

package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/memory"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/webhook"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

func TestDeletionDefender(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(fmt.Errorf("test complete (from defer block)"))

	log.SetLogger(klogr.New())

	kube := envtest.Environment{
		CRDDirectoryPaths: []string{repo.GetCRDsPath()},
	}
	restConfig, err := kube.Start()
	if err != nil {
		t.Fatalf("Failed to start envtest: %v", err)
	}
	defer func() {
		if err := kube.Stop(); err != nil {
			t.Errorf("Failed to stop envtest: %v", err)
		}
	}()

	memoryRecorder := memory.StartMemoryRecorder()

	httpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		t.Fatalf("Failed to create HTTP client: %v", err)
	}
	dynamic, err := dynamic.NewForConfigAndClient(restConfig, httpClient)
	if err != nil {
		t.Fatalf("Failed to create dynamic client: %v", err)
	}

	// Avoid port 443 (requires root)
	webhook.ServicePort = 8443

	deletionDefender := &DeletionDefender{
		RESTConfig: restConfig,
	}

	deletionDefender.ManagerOptions = manager.Options{
		Metrics: server.Options{BindAddress: "0"}, // disable metrics
	}

	deletionDefenderErr := make(chan error, 1)
	go func() {
		err := deletionDefender.Run(ctx)
		deletionDefenderErr <- err
	}()

	// Create cnrm-system namespace
	{
		ns := &unstructured.Unstructured{}
		ns.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Namespace",
		})
		ns.SetName("cnrm-system")
		_, err := dynamic.Resource(schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "namespaces",
		}).Create(ctx, ns, metav1.CreateOptions{})
		if err != nil {
			t.Fatalf("Failed to create cnrm-system namespace: %v", err)
		}
	}

	u := &unstructured.Unstructured{}
	u.SetNamespace("default")
	u.SetName("test-bucket")
	gvr := schema.GroupVersionResource{
		Group:    "storage.cnrm.cloud.google.com",
		Version:  "v1beta1",
		Resource: "storagebuckets",
	}
	gvk := gvr.GroupVersion().WithKind("StorageBucket")
	u.SetGroupVersionKind(gvk)
	u.SetFinalizers([]string{"cnrm.cloud.google.com/deletion-defender"})

	// Create a simple test bucket
	t.Logf("Creating test StorageBucket resource")
	if _, err := dynamic.Resource(gvr).Namespace(u.GetNamespace()).Create(ctx, u, metav1.CreateOptions{}); err != nil {
		t.Fatalf("Failed to create test bucket: %v", err)
	}

	// Delete the StorageBucket resource to trigger the deletion defender
	t.Logf("Deleting test StorageBucket resource")
	if err := dynamic.Resource(gvr).Namespace(u.GetNamespace()).Delete(ctx, u.GetName(), metav1.DeleteOptions{}); err != nil {
		t.Fatalf("Failed to delete test bucket: %v", err)
	}

	// Wait for the deletion defender to clear the finalizer, enabling the resource to be deleted
	for {
		time.Sleep(1 * time.Second)
		latest, err := dynamic.Resource(gvr).Namespace(u.GetNamespace()).Get(ctx, u.GetName(), metav1.GetOptions{})
		if err != nil {
			if apierrors.IsNotFound(err) {
				// The deletion defender has not restored the resource yet; keep waiting
				break
			}
			t.Fatalf("Failed to get test bucket after deletion: %v", err)
		} else {
			t.Logf("StorageBucket resource not yet deleted (finalizers: %v); waiting...", latest.GetFinalizers())
		}
	}

	memoryReport := memoryRecorder.Stop()

	cancel(fmt.Errorf("stopping test"))

	if err := <-deletionDefenderErr; err != nil {
		if !errors.Is(err, context.Canceled) {
			t.Errorf("DeletionDefender failed: %v", err)
		}
	}

	t.Logf("Memory usage report: HeapAlloc=%d HeapInuse=%d TotalAlloc=%d", memoryReport.HeapAlloc, memoryReport.HeapInuse, memoryReport.TotalAlloc)
}
