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
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
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

func RepoRoot(t *testing.T) string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get repo root: %v", err)
	}
	repoRoot := strings.TrimSpace(string(output))
	return repoRoot
}

func TestProfileRecorderFootprint(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(fmt.Errorf("test complete (from defer block)"))

	repoRoot := RepoRoot(t)

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

	httpClient, err := rest.HTTPClientFor(restConfig)
	if err != nil {
		t.Fatalf("Failed to create HTTP client: %v", err)
	}
	dynamic, err := dynamic.NewForConfigAndClient(restConfig, httpClient)
	if err != nil {
		t.Fatalf("Failed to create dynamic client: %v", err)
	}

	t.Logf("Creating test StorageBucket resource")
	for i := 0; i < 1000; i++ {
		// Create a simple test bucket
		u := &unstructured.Unstructured{}
		u.SetNamespace("default")
		u.SetName(fmt.Sprintf("test-bucket-%d", i))
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

	// Write a kubeconfig file for the recorder process
	kubeconfigPath := filepath.Join(t.TempDir(), "kubeconfig")
	{
		// TODO: Use RBAC
		user, err := kube.AddUser(envtest.User{Name: "recorder", Groups: []string{"system:masters"}}, restConfig)
		if err != nil {
			t.Fatalf("Failed to add user: %v", err)
		}
		kubeconfig, err := user.KubeConfig()
		if err != nil {
			t.Fatalf("Failed to get kubeconfig: %v", err)
		}

		if err := os.WriteFile(kubeconfigPath, kubeconfig, 0600); err != nil {
			t.Fatalf("Failed to write kubeconfig file: %v", err)
		}
		defer func() {
			if err := os.Remove(kubeconfigPath); err != nil {
				t.Errorf("Failed to remove kubeconfig file: %v", err)
			}
		}()
	}

	recorderBinary := filepath.Join(t.TempDir(), "recorder")
	{
		// Build the recorder binary
		cmd := exec.Command("go", "build", "-o", recorderBinary, filepath.Join(repoRoot, "cmd", "recorder"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			t.Fatalf("Failed to build recorder binary: %v", err)
		}
	}

	env := os.Environ()
	env = append(env, "KUBECONFIG="+kubeconfigPath)
	env = append(env, "GODEBUG=memprofilerate=1") // Enable memory profiling for every allocation

	cmd := exec.Command(recorderBinary, "--enable-pprof")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start recorder: %v", err)
	}
	defer func() {
		if err := cmd.Process.Kill(); err != nil {
			t.Errorf("Failed to kill recorder process: %v", err)
		}
	}()

	time.Sleep(5 * time.Second) // Give recorder some time to start

	// Scrapes prometheus metrics endpoint and returns metrics with the given prefix
	getMetricsAsString := func(prefix string) string {
		resp, err := http.Get("http://localhost:8888/metrics")
		if err != nil {
			t.Fatalf("Failed to get metrics: %v", err)
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read metrics body: %v", err)
		}

		filtered := []string{}
		for line := range strings.SplitSeq(string(b), "\n") {
			if strings.HasPrefix(line, prefix) {
				filtered = append(filtered, line)
			}
		}
		return strings.Join(filtered, "\n")
	}

	// Let the recorder run for a couple of intervals to see the object
	t.Logf("waiting for recorder to pick up new resource")
	timeoutAt := time.Now().Add(time.Minute * 2)
	for {
		time.Sleep(2 * time.Second)
		got := getMetricsAsString("configconnector_applied_resources_total{")
		want := `configconnector_applied_resources_total{group_version_kind="StorageBucket.storage.cnrm.cloud.google.com",namespace="default",status="NoCondition"} 1000`

		if diff := cmp.Diff(want, got); diff != "" {
			t.Logf("Metrics do not yet match expected value.\nDiff: %v", diff)
			if time.Now().After(timeoutAt) {
				t.Fatalf("Timed out waiting for recorder to sync")
			}
			continue
		}

		t.Logf("Metrics match expected value: %v", got)
		break
	}

	// Capture a pprof heap profile
	{
		resp, err := http.Get("http://localhost:6060/debug/pprof/heap")
		if err != nil {
			t.Fatalf("Failed to get heap profile: %v", err)
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read heap profile body: %v", err)
		}

		artifactsDir := os.Getenv("ARTIFACTS")

		profilePath := filepath.Join(artifactsDir, "pprof", t.Name(), "heap.prof")
		if err := os.MkdirAll(filepath.Dir(profilePath), 0755); err != nil {
			t.Fatalf("Failed to create pprof directory: %v", err)
		}
		if err := os.WriteFile(profilePath, b, 0600); err != nil {
			t.Fatalf("Failed to write heap profile file: %v", err)
		}
		t.Logf("Wrote heap profile to %q", profilePath)
	}

	cancel(fmt.Errorf("test complete"))
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
