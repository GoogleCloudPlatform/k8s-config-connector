// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integration

import (
	"bytes"
	"context"
	"fmt"
	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/gke-labs/kube-etl/syncer/controllers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"os"
	"os/exec"
	"path/filepath"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if os.Getenv("KUBEBUILDER_ASSETS") == "" {
		cmd := exec.Command("go", "run", "sigs.k8s.io/controller-runtime/tools/setup-envtest@latest", "use", "-p", "path")
		out, err := cmd.Output()
		if err == nil && len(out) > 0 {
			os.Setenv("KUBEBUILDER_ASSETS", string(bytes.TrimSpace(out)))
		}
	}
	os.Exit(m.Run())
}

func TestKRMSyncerIntegration(t *testing.T) {
	// 1. Setup Environments
	testScheme := scheme.Scheme
	require.NoError(t, krmv1alpha1.AddToScheme(testScheme))

	// Cluster A
	t.Log("Starting Cluster A...")
	cfgA := &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "config", "crd"), filepath.Join("..", "integration", "testcrd")},
		ErrorIfCRDPathMissing: true,
	}
	configA, err := cfgA.Start()
	require.NoError(t, err)
	defer cfgA.Stop()

	k8sClientA, err := client.New(configA, client.Options{Scheme: testScheme})
	require.NoError(t, err)

	// Cluster B
	t.Log("Starting Cluster B...")
	cfgB := &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "integration", "testcrd")},
		ErrorIfCRDPathMissing: true,
	}
	configB, err := cfgB.Start()
	require.NoError(t, err)
	defer cfgB.Stop()

	k8sClientB, err := client.New(configB, client.Options{Scheme: testScheme})
	require.NoError(t, err)

	// 2. Start Manager for Cluster A
	t.Log("Starting Manager on Cluster A...")
	mgr, err := ctrl.NewManager(configA, ctrl.Options{
		Scheme:  testScheme,
		Metrics: metricsserver.Options{BindAddress: "0"},
	})
	require.NoError(t, err)

	reconciler := &controllers.KRMSyncerReconciler{
		Client:  mgr.GetClient(),
		Scheme:  mgr.GetScheme(),
		Manager: mgr,
		// Resolve the GKE cluster referenced by the test cases to Cluster B.
		RemoteConfigProvider: staticConfigProvider{cfg: configB},
	}
	require.NoError(t, reconciler.SetupWithManager(mgr))

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	go func() {
		err := mgr.Start(ctx)
		if err != nil {
			t.Errorf("Manager failed: %v", err)
		}
	}()

	// 3. Run Test Cases
	casesDir := "../integration/cases"
	dirs, err := os.ReadDir(casesDir)
	require.NoError(t, err)

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		caseName := d.Name()
		t.Run(caseName, func(t *testing.T) {
			runTestCase(t, ctx, k8sClientA, k8sClientB, filepath.Join(casesDir, caseName))
		})
	}
}

func runTestCase(t *testing.T, ctx context.Context, clientA, clientB client.Client, caseDir string) {
	t.Logf("Running test case: %s", filepath.Base(caseDir))

	// Load Syncer
	syncerBytes, err := os.ReadFile(filepath.Join(caseDir, "syncer.yaml"))
	require.NoError(t, err)
	syncer := &krmv1alpha1.KRMSyncer{}
	require.NoError(t, yaml.Unmarshal(syncerBytes, syncer))

	// Clean up Syncer at end of test
	defer func() {
		_ = clientA.Delete(ctx, syncer)
	}()

	// Start the Syncer
	require.NoError(t, clientA.Create(ctx, syncer))

	// Determine Source and Destination clients based on mode
	sourceClient := clientA
	destClient := clientB
	if syncer.Spec.Mode == krmv1alpha1.ModePull {
		sourceClient = clientB
		destClient = clientA
	}

	// Read all Resources from object.yaml
	objectPath := "../integration/testdata/object.yaml"
	createBytes, err := os.ReadFile(objectPath)
	require.NoError(t, err)

	resources := []*unstructured.Unstructured{}
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(createBytes), 4096)
	for {
		res := &unstructured.Unstructured{}
		if err := decoder.Decode(res); err != nil {
			if err == io.EOF {
				break
			}
			require.NoError(t, err)
		}
		if res.Object == nil {
			continue
		}
		resources = append(resources, res)
	}

	// Clean up Resources at end of test
	defer func() {
		for _, r := range resources {
			_ = clientA.Delete(ctx, r)
			_ = clientB.Delete(ctx, r) // Cleanup on both just in case
		}
	}()

	// Apply All Resources in Source
	for _, resource := range resources {
		t.Logf("Creating Resource %s/%s in Source Cluster...", resource.GetKind(), resource.GetName())
		// Capture initial status
		initialStatus, hasStatus, err := unstructured.NestedFieldCopy(resource.Object, "status")
		require.NoError(t, err)

		// Create in Source
		resCopy := resource.DeepCopy()
		require.NoError(t, sourceClient.Create(ctx, resCopy))

		// Update Status in Source
		if hasStatus {
			statusObj := resCopy.DeepCopy()
			err := unstructured.SetNestedField(statusObj.Object, initialStatus, "status")
			require.NoError(t, err)
			require.NoError(t, sourceClient.Status().Update(ctx, statusObj))
		}
	}

	// Sleep to allow resources to sync
	time.Sleep(2 * time.Second)

	// Verify Output
	expectedPath := filepath.Join(caseDir, "expected.yaml")
	expectedBytes, err := os.ReadFile(expectedPath)
	expectedResources := make(map[string]*unstructured.Unstructured)
	if err == nil && len(expectedBytes) > 0 {
		expDecoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(expectedBytes), 4096)
		for {
			exp := &unstructured.Unstructured{}
			if err := expDecoder.Decode(exp); err != nil {
				if err == io.EOF {
					break
				}
				require.NoError(t, err)
			}
			if exp.Object == nil {
				continue
			}
			// Use Kind/Name as key to be safe, although names are unique here
			key := fmt.Sprintf("%s/%s", exp.GetKind(), exp.GetName())
			expectedResources[key] = exp
		}
	}

	for _, resource := range resources {
		key := fmt.Sprintf("%s/%s", resource.GetKind(), resource.GetName())
		expected, shouldExist := expectedResources[key]

		actual := &unstructured.Unstructured{}
		actual.SetGroupVersionKind(resource.GroupVersionKind())
		err := destClient.Get(ctx, client.ObjectKeyFromObject(resource), actual)

		if shouldExist {
			t.Logf("Verifying expected resource %s in Destination Cluster...", key)
			require.NoError(t, err, "Resource %s should exist in Destination Cluster", key)

			// Compare Spec
			expectedSpec, _, _ := unstructured.NestedMap(expected.Object, "spec")
			actualSpec, _, _ := unstructured.NestedMap(actual.Object, "spec")
			assert.Equal(t, expectedSpec, actualSpec, "Spec mismatch for %s", key)

			// Compare Status
			expectedStatus, _, _ := unstructured.NestedMap(expected.Object, "status")
			actualStatus, _, _ := unstructured.NestedMap(actual.Object, "status")
			assert.Equal(t, expectedStatus, actualStatus, "Status mismatch for %s", key)
		} else {
			t.Logf("Verifying resource %s does NOT exist in Destination Cluster...", key)
			assert.True(t, errors.IsNotFound(err), "Resource %s should NOT exist in Destination Cluster", key)
		}
	}
}

// staticConfigProvider resolves every remote cluster to a fixed envtest cluster.
type staticConfigProvider struct {
	cfg *rest.Config
}

// Key uses the same connection key as the GKE provider.
func (p staticConfigProvider) Key(namespace string, remote *krmv1alpha1.RemoteConfig) (string, error) {
	return (&controllers.GKEConfigProvider{}).Key(namespace, remote)
}

func (p staticConfigProvider) RESTConfig(_ context.Context, _ string, remote *krmv1alpha1.RemoteConfig) (*rest.Config, error) {
	if remote == nil || remote.GKECluster == nil {
		return nil, fmt.Errorf("spec.remote.gkeCluster must be set")
	}
	return rest.CopyConfig(p.cfg), nil
}
