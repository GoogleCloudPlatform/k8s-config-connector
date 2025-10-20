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
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func TestRecorder(t *testing.T) {
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

	memoryRecorder := StartMemoryRecorder()

	httpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		t.Fatalf("Failed to create HTTP client: %v", err)
	}
	dynamic, err := dynamic.NewForConfigAndClient(restConfig, httpClient)
	if err != nil {
		t.Fatalf("Failed to create dynamic client: %v", err)
	}

	recorder := &MetricsRecorder{
		MetricInterval: 5 * time.Second, // TODO: Make this shorter as we optimize this
		RESTConfig:     restConfig,
	}

	managerErr := make(chan error, 1)
	go func() {
		err := recorder.Run(ctx)
		managerErr <- err
	}()

	getMetricsAsString := func() string {
		var metrics []string
		recorder.GetMetrics(func(namespace string, gvk schema.GroupVersionKind, condition string, count int64) {
			metrics = append(metrics, fmt.Sprintf("%s %s %s: %d", gvk.Kind, namespace, condition, count))
		})
		sort.Strings(metrics)

		return strings.Join(metrics, "\n")
	}

	// Let the recorder run for a couple of intervals to establish watches
	t.Logf("waiting for recorder to start watches")
	time.Sleep(2 * recorder.MetricInterval)

	t.Logf("verifying metrics before creating any resources")
	{
		got := getMetricsAsString()
		want := ``

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Metrics do not match expected.\nGot:\n%s\nWant:\n%s", got, want)
		}
	}

	t.Logf("Creating test StorageBucket resource")
	{
		// Create a simple test bucket
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
		if _, err := dynamic.Resource(gvr).Namespace(u.GetNamespace()).Create(ctx, u, metav1.CreateOptions{}); err != nil {
			t.Fatalf("Failed to create test bucket: %v", err)
		}
	}

	// Let the recorder run for a couple of intervals to see the object
	t.Logf("waiting for recorder to pick up new resource")
	timeoutAt := time.Now().Add(time.Minute * 2)
	for {
		time.Sleep(2 * recorder.MetricInterval)
		if recorder.IsReady(schema.GroupVersionKind{
			Group:   "storage.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "StorageBucket",
		}) {
			t.Logf("recorder is ready for StorageBucket")
			break
		}
		t.Logf("waiting for recorder to be ready for StorageBucket")
		if time.Now().After(timeoutAt) {
			t.Fatalf("Timed out waiting for recorder to pick up new resource")
		}
	}

	t.Logf("verifying metrics after creating resource")
	{
		got := getMetricsAsString()
		want := `StorageBucket default NoCondition: 1`

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Metrics do not match expected.\nGot:\n%s\nWant:\n%s", got, want)
		}
	}

	memoryReport := memoryRecorder.Stop()

	cancel(fmt.Errorf("test complete"))

	if err := <-managerErr; err != nil {
		if !errors.Is(err, context.Canceled) {
			t.Errorf("Recorder run failed: %v", err)
		}
	}

	t.Logf("Memory usage report: HeapAlloc=%d HeapInuse=%d TotalAlloc=%d", memoryReport.HeapAlloc, memoryReport.HeapInuse, memoryReport.TotalAlloc)
}

// MemoryReport holds memory usage statistics, and makes it easy to report basic memory usage for a test.
type MemoryRecorder struct {
	before runtime.MemStats
	after  runtime.MemStats
}

// StartMemoryRecorder captures the memory usage at the start of a test.
func StartMemoryRecorder() *MemoryRecorder {
	r := &MemoryRecorder{}
	runtime.GC()
	runtime.ReadMemStats(&r.before)
	return r
}

// Stop captures the memory usage at the end of a test, and returns a report of the differences.
func (r *MemoryRecorder) Stop() *MemoryRecorderReport {
	runtime.GC()
	runtime.ReadMemStats(&r.after)

	report := &MemoryRecorderReport{
		HeapAlloc:  int64(r.after.HeapAlloc) - int64(r.before.HeapAlloc),
		HeapInuse:  int64(r.after.HeapInuse) - int64(r.before.HeapInuse),
		TotalAlloc: int64(r.after.TotalAlloc) - int64(r.before.TotalAlloc),
	}
	return report
}

// MemoryRecorderReport holds memory usage statistics from a MemoryRecorder.
type MemoryRecorderReport struct {
	HeapAlloc  int64
	HeapInuse  int64
	TotalAlloc int64
}
