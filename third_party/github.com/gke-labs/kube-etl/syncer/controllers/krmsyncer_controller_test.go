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

package controllers

import (
	"bytes"
	"context"
	stderrors "errors"
	"fmt"
	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/rest"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var (
	cfgSource          *rest.Config
	cfgDest            *rest.Config
	cfgRetryDest       *rest.Config
	k8sClientSource    client.Client
	k8sClientDest      client.Client
	testEnvSource      *envtest.Environment
	testEnvDest        *envtest.Environment
	mgrCtx             context.Context
	mgrCancel          context.CancelFunc
	faultInjector      *faultInjectingRoundTripper
	testRemoteProvider *testConfigProvider
	krmReconciler      *KRMSyncerReconciler
)

// destClusterName is the fake GKE cluster name that tests resolve to the
// envtest destination cluster.
const (
	destClusterName      = "dest"
	destRetryClusterName = "dest-retry"
)

type faultInjectingRoundTripper struct {
	base       http.RoundTripper
	failPatch  atomic.Bool
	failDelete atomic.Bool
	failStatus atomic.Bool
}

func (rt *faultInjectingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == http.MethodPatch && rt.failStatus.Load() && strings.HasSuffix(req.URL.Path, "/status") {
		return nil, stderrors.New("injected status patch failure")
	}
	if req.Method == http.MethodPatch && rt.failPatch.Load() && !strings.HasSuffix(req.URL.Path, "/status") {
		return nil, stderrors.New("injected apply patch failure")
	}
	if req.Method == http.MethodDelete && rt.failDelete.Load() {
		return nil, stderrors.New("injected delete failure")
	}
	return rt.base.RoundTrip(req)
}

// testConfigProvider resolves fake GKE clusters (by name) to envtest clusters.
type testConfigProvider struct {
	clusters map[string]*rest.Config
}

// Key uses the same connection key as the GKE provider.
func (p *testConfigProvider) Key(namespace string, remote *krmv1alpha1.RemoteConfig) (string, error) {
	return (&GKEConfigProvider{}).Key(namespace, remote)
}

func (p *testConfigProvider) RESTConfig(_ context.Context, _ string, remote *krmv1alpha1.RemoteConfig) (*rest.Config, error) {
	if _, err := gkeClusterName(remote); err != nil {
		return nil, err
	}
	cfg, ok := p.clusters[remote.GKECluster.Name]
	if !ok {
		return nil, fmt.Errorf("unknown test cluster %q", remote.GKECluster.Name)
	}
	return rest.CopyConfig(cfg), nil
}

func testRemote(name string) *krmv1alpha1.RemoteConfig {
	return &krmv1alpha1.RemoteConfig{
		GKECluster: &krmv1alpha1.GKECluster{Project: "test-project", Location: "us-central1", Name: name},
	}
}

func TestMain(m *testing.M) {
	if os.Getenv("KUBEBUILDER_ASSETS") == "" {
		cmd := exec.Command("go", "run", "sigs.k8s.io/controller-runtime/tools/setup-envtest@latest", "use", "-p", "path")
		out, err := cmd.Output()
		if err == nil && len(out) > 0 {
			os.Setenv("KUBEBUILDER_ASSETS", string(bytes.TrimSpace(out)))
		}
	}
	ctrl.SetLogger(klog.NewKlogr())
	mgrCtx, mgrCancel = context.WithCancel(ctrl.SetupSignalHandler())

	// Start Source Cluster
	testEnvSource = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "config", "crd")},
		ErrorIfCRDPathMissing: true,
	}
	var err error
	cfgSource, err = testEnvSource.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start source test env: %v\n", err)
		os.Exit(1)
	}

	// Start Destination Cluster
	testEnvDest = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "config", "crd")},
		ErrorIfCRDPathMissing: true,
	}
	cfgDest, err = testEnvDest.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start dest test env: %v\n", err)
		os.Exit(1)
	}

	// Register Scheme
	if err := krmv1alpha1.AddToScheme(scheme.Scheme); err != nil {
		fmt.Fprintf(os.Stderr, "failed to add to scheme: %v\n", err)
		os.Exit(1)
	}

	// Create Clients
	k8sClientSource, err = client.New(cfgSource, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create source client: %v\n", err)
		os.Exit(1)
	}
	k8sClientDest, err = client.New(cfgDest, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create dest client: %v\n", err)
		os.Exit(1)
	}

	testSyncPeriod := 2 * time.Second
	// Start Manager in Source Cluster
	mgr, err := ctrl.NewManager(cfgSource, ctrl.Options{
		Metrics: metricsserver.Options{
			BindAddress: "0",
		},
		Cache: cache.Options{
			SyncPeriod: &testSyncPeriod,
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create manager: %v\n", err)
		os.Exit(1)
	}

	faultInjector = &faultInjectingRoundTripper{}
	cfgRetryDest = rest.CopyConfig(cfgDest)
	cfgRetryDest.Wrap(func(rt http.RoundTripper) http.RoundTripper {
		faultInjector.base = rt
		return faultInjector
	})

	testRemoteProvider = &testConfigProvider{
		clusters: map[string]*rest.Config{
			destClusterName:      cfgDest,
			destRetryClusterName: cfgRetryDest,
		},
	}

	krmReconciler = &KRMSyncerReconciler{
		Client:               mgr.GetClient(),
		Scheme:               mgr.GetScheme(),
		Manager:              mgr,
		RemoteConfigProvider: testRemoteProvider,
		SyncPeriod:           &testSyncPeriod,
	}
	if err := krmReconciler.SetupWithManager(mgr); err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup controller: %v\n", err)
		os.Exit(1)
	}

	go func() {
		if err := mgr.Start(mgrCtx); err != nil {
			klog.ErrorS(err, "manager failed")
		}
	}()

	code := m.Run()

	mgrCancel()
	// give the manager a moment to stop
	time.Sleep(100 * time.Millisecond)
	if err := testEnvSource.Stop(); err != nil {
		klog.ErrorS(err, "failed to stop source test env")
	}
	if err := testEnvDest.Stop(); err != nil {
		klog.ErrorS(err, "failed to stop dest test env")
	}

	os.Exit(code)
}

func TestSyncerSync(t *testing.T) {
	ctx := t.Context()
	// Test Logic
	ns := "default"
	syncerName := "test-syncer"
	targetServiceName := "target-service"

	// Create Syncer
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: syncerName, Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Suspend: false,
			Mode:    krmv1alpha1.ModePush,
			Remote:  testRemote(destClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"spec"},
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))

	// Create target Service in Source
	target := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 80}},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, target))

	// Verify Service Sync to Dest
	err := wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		destSvc := &corev1.Service{}
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		if err != nil {
			return false, nil
		}
		return len(destSvc.Spec.Ports) > 0 && destSvc.Spec.Ports[0].Port == 80, nil
	})
	assert.NoError(t, err, "Service should be synced to dest")

	// Update target Service in Source
	require.NoError(t, k8sClientSource.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, target))
	target.Spec.Ports[0].Port = 8080
	require.NoError(t, k8sClientSource.Update(ctx, target))

	// Verify Service Update in Dest
	err = wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		destSvc := &corev1.Service{}
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		if err != nil {
			return false, nil
		}
		return len(destSvc.Spec.Ports) > 0 && destSvc.Spec.Ports[0].Port == 8080, nil
	})
	assert.NoError(t, err, "Service update should be synced to dest")

	// Delete target Service in Source
	require.NoError(t, k8sClientSource.Delete(ctx, target))

	// Verify Deletion in Dest
	err = wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		destSvc := &corev1.Service{}
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		return client.IgnoreNotFound(err) == nil && err != nil, nil
	})
	assert.NoError(t, err, "Service should be deleted from dest")
}

func TestSyncerSyncFields(t *testing.T) {
	ctx := t.Context()
	// Test Logic
	ns := "default"
	syncerName := "test-syncer-fields"
	targetName := "test-service-fields"

	// Create Syncer with SyncFields
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: syncerName, Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote(destClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"spec"},
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))

	// Create Service resource in Source
	target := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": "test-fields"},
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					Port: 80,
				},
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, target))

	// Verify Sync to Dest: spec should be present with nested fields
	err := wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		destSvc := &corev1.Service{}
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetName, Namespace: ns}, destSvc)
		if err != nil {
			return false, nil
		}
		// ServiceSpec has primitive (Type), complex (Selector), and nested types (Ports)
		return destSvc.Spec.Type == corev1.ServiceTypeClusterIP &&
			destSvc.Spec.Selector["app"] == "test-fields" &&
			len(destSvc.Spec.Ports) == 1 &&
			destSvc.Spec.Ports[0].Port == 80, nil
	})
	assert.NoError(t, err, "Service should be synced to dest with all spec fields")
}

func TestSyncerSyncStatusSubresource(t *testing.T) {
	ctx := t.Context()
	// Test Logic: Sync Service itself (it has status subresource)
	ns := "default"
	syncerName := "test-syncer-status"
	observedServiceName := "observed-service"

	// Create observed service FIRST
	observed := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: observedServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 80}},
			Type:  corev1.ServiceTypeLoadBalancer,
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, observed))

	// Create syncer rule SECOND
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: syncerName, Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote(destClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"spec", "status"},
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))

	// Wait for initial sync of observed service (spec only)
	err := wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		destSvc := &corev1.Service{}
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: observedServiceName, Namespace: ns}, destSvc)
		return err == nil, nil
	})
	require.NoError(t, err, "observed service should be synced initially")

	// Update status of the observed service THIRD (to trigger event)
	require.NoError(t, k8sClientSource.Get(ctx, types.NamespacedName{Name: observedServiceName, Namespace: ns}, observed))
	observed.Status.LoadBalancer.Ingress = []corev1.LoadBalancerIngress{
		{IP: "1.2.3.4"},
	}
	require.NoError(t, k8sClientSource.Status().Update(ctx, observed))

	// Verify status sync to Dest
	err = wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		destSvc := &corev1.Service{}
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: observedServiceName, Namespace: ns}, destSvc)
		if err != nil {
			return false, nil
		}
		return len(destSvc.Status.LoadBalancer.Ingress) > 0 && destSvc.Status.LoadBalancer.Ingress[0].IP == "1.2.3.4", nil
	})
	assert.NoError(t, err, "Service status should be synced to dest")
}

func TestSyncerValidation(t *testing.T) {
	ctx := t.Context()
	ns := "default"

	// Test case 1: Valid syncFields
	validSyncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: "valid-syncer", Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote("dummy"),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "syncer.gkelabs.io", Version: "v1alpha1", Kind: "KRMSyncer",
					SyncFields: []string{"spec", "status", "spec.resourceID"},
				},
			},
		},
	}
	assert.NoError(t, k8sClientSource.Create(ctx, validSyncer))

	// Test case 3: Default value
	defaultSyncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: "default-syncer", Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote("dummy"),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "syncer.gkelabs.io", Version: "v1alpha1", Kind: "KRMSyncer",
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, defaultSyncer))
	assert.Equal(t, []string{"status"}, defaultSyncer.Spec.Rules[0].SyncFields)
}

func TestSyncerRemoteValidation(t *testing.T) {
	ctx := t.Context()
	ns := "default"

	gkeCluster := func(extra map[string]any) map[string]any {
		c := map[string]any{"project": "test-project", "location": "us-central1", "name": "dest"}
		for k, v := range extra {
			c[k] = v
		}
		return c
	}

	for _, tc := range []struct {
		name    string
		remote  any // nil omits spec.remote entirely
		wantErr string
	}{
		{
			name:    "no remote",
			remote:  nil,
			wantErr: "spec.remote: Required value",
		},
		{
			name:    "invalid endpoint",
			remote:  map[string]any{"gkeCluster": gkeCluster(map[string]any{"endpoint": "PublicIP"})},
			wantErr: `spec.remote.gkeCluster.endpoint: Unsupported value: "PublicIP"`,
		},
		{
			name:   "valid, endpoint defaulted",
			remote: map[string]any{"gkeCluster": gkeCluster(nil)},
		},
		{
			name:   "valid, DNS endpoint",
			remote: map[string]any{"gkeCluster": gkeCluster(map[string]any{"endpoint": "DNS"})},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			spec := map[string]any{
				"mode":  "pull",
				"rules": []any{map[string]any{"group": "", "version": "v1", "kind": "ConfigMap"}},
			}
			if tc.remote != nil {
				spec["remote"] = tc.remote
			}
			u := &unstructured.Unstructured{Object: map[string]any{"spec": spec}}
			u.SetGroupVersionKind(krmv1alpha1.GroupVersion.WithKind("KRMSyncer"))
			u.SetGenerateName("remote-validation-")
			u.SetNamespace(ns)
			// Suspend so the controller does not try to connect to the remote.
			require.NoError(t, unstructured.SetNestedField(u.Object, true, "spec", "suspend"))

			err := k8sClientSource.Create(ctx, u)
			if tc.wantErr == "" {
				require.NoError(t, err)
				defer func() { _ = k8sClientSource.Delete(ctx, u) }()
				endpoint, _, _ := unstructured.NestedString(u.Object, "spec", "remote", "gkeCluster", "endpoint")
				if tc.name == "valid, endpoint defaulted" {
					assert.Equal(t, string(krmv1alpha1.GKEEndpointDefault), endpoint)
				}
				return
			}
			require.Error(t, err)
			assert.True(t, errors.IsInvalid(err), "expected Invalid, got %v", err)
			assert.ErrorContains(t, err, tc.wantErr)
		})
	}
}

func TestSyncerPull(t *testing.T) {
	ctx := t.Context()
	ns := "default"
	syncerName := "test-syncer-pull"
	targetServiceName := "target-service-pull"

	// Create Syncer in Pull mode
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: syncerName, Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePull,
			Remote: testRemote(destClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"spec"},
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))

	// Create target Service in Remote (cfgDest)
	target := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 9090}},
		},
	}
	require.NoError(t, k8sClientDest.Create(ctx, target))

	// Verify Service Sync from Remote to Local (cfgSource)
	err := wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		localSvc := &corev1.Service{}
		err := k8sClientSource.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, localSvc)
		if err != nil {
			return false, nil
		}
		return len(localSvc.Spec.Ports) > 0 && localSvc.Spec.Ports[0].Port == 9090, nil
	})
	assert.NoError(t, err, "Service should be pulled from remote to local")

	// Update target Service in Remote
	require.NoError(t, k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, target))
	target.Spec.Ports[0].Port = 9091
	require.NoError(t, k8sClientDest.Update(ctx, target))

	// Verify Service Update Pulled to Local
	err = wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		localSvc := &corev1.Service{}
		err := k8sClientSource.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, localSvc)
		if err != nil {
			return false, nil
		}
		return len(localSvc.Spec.Ports) > 0 && localSvc.Spec.Ports[0].Port == 9091, nil
	})
	assert.NoError(t, err, "Service update should be pulled to local")

	// Delete target Service in Remote
	require.NoError(t, k8sClientDest.Delete(ctx, target))

	// Verify Deletion Pulled to Local
	err = wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		localSvc := &corev1.Service{}
		err := k8sClientSource.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, localSvc)
		return client.IgnoreNotFound(err) == nil && err != nil, nil
	})
	assert.NoError(t, err, "Service should be deleted from local")
}

func TestFilterFields(t *testing.T) {
	dr := &DynamicResourceReconciler{
		LocalClient: nil,
		GVK:         krmv1alpha1.GroupVersion.WithKind("KRMSyncer"),
	}

	src := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "syncer.gkelabs.io/v1alpha1",
			"kind":       "KRMSyncer",
			"metadata": map[string]any{
				"name":      "test",
				"namespace": "default",
				"labels": map[string]any{
					"l1": "v1",
				},
				"annotations": map[string]any{
					"a1": "v1",
				},
			},
			"spec": map[string]any{
				"resourceID": "id1",
				"resource": map[string]any{
					"ID": "id2",
				},
				"other": "val",
			},
			"status": map[string]any{
				"phase": "Active",
			},
		},
	}

	// Test case 1: Sync spec.resourceID and status
	fields := []string{"spec.resourceID", "status"}
	dest, err := dr.filterFields(src, fields)
	assert.NoError(t, err)

	val, found, _ := unstructured.NestedString(dest.Object, "spec", "resourceID")
	assert.True(t, found)
	assert.Equal(t, "id1", val)

	_, found, _ = unstructured.NestedMap(dest.Object, "spec", "resource")
	assert.False(t, found)

	_, found, _ = unstructured.NestedString(dest.Object, "spec", "other")
	assert.False(t, found)

	status, found, _ := unstructured.NestedMap(dest.Object, "status")
	assert.True(t, found)
	assert.Equal(t, "Active", status["phase"])

	assert.Equal(t, "test", dest.GetName())
	assert.Equal(t, "default", dest.GetNamespace())
	// Status mode: labels and annotations must not be synced.
	assert.Nil(t, dest.GetLabels())
	assert.Nil(t, dest.GetAnnotations())
	_, found, _ = unstructured.NestedFieldNoCopy(dest.Object, "metadata", "labels")
	assert.False(t, found)
	_, found, _ = unstructured.NestedFieldNoCopy(dest.Object, "metadata", "annotations")
	assert.False(t, found)

	// Test case 2: Sync full "spec" with nested fields of various types
	srcFull := &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]any{
				"name": "full-spec-test",
				"labels": map[string]any{
					"l1": "v1",
				},
				"annotations": map[string]any{
					"a1": "v1",
				},
			},
			"spec": map[string]any{
				"primitive": "string-val",
				"integer":   int64(42),
				"boolean":   true,
				"complex": map[string]any{
					"field1": "val1",
					"field2": int64(2),
					"nested": map[string]any{
						"deep": "deeper",
					},
				},
				"list": []any{
					map[string]any{"item": int64(1)},
					"simple-item",
				},
			},
		},
	}
	fieldsFull := []string{"spec"}
	destFull, err := dr.filterFields(srcFull, fieldsFull)
	assert.NoError(t, err)

	// Full mode: labels and annotations are synced.
	assert.Equal(t, map[string]string{"l1": "v1"}, destFull.GetLabels())
	assert.Equal(t, map[string]string{"a1": "v1"}, destFull.GetAnnotations())

	spec, found, _ := unstructured.NestedMap(destFull.Object, "spec")
	assert.True(t, found)
	assert.Equal(t, "string-val", spec["primitive"])
	assert.Equal(t, int64(42), spec["integer"])
	assert.Equal(t, true, spec["boolean"])

	complexVal, found, _ := unstructured.NestedMap(destFull.Object, "spec", "complex")
	assert.True(t, found)
	assert.Equal(t, "val1", complexVal["field1"])

	deepVal, found, _ := unstructured.NestedString(destFull.Object, "spec", "complex", "nested", "deep")
	assert.True(t, found)
	assert.Equal(t, "deeper", deepVal)

	listVal, found, _ := unstructured.NestedSlice(destFull.Object, "spec", "list")
	assert.True(t, found)
	assert.Len(t, listVal, 2)
}

func TestSyncerSyncMissingNamespace(t *testing.T) {
	ctx := t.Context()
	ns := "missing-dest-ns"
	syncerName := "test-syncer-missing-ns"
	targetServiceName := "target-service-missing-ns"

	// 1. Create the source namespace first so we have a place to create our source objects
	srcNamespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: ns},
	}
	require.NoError(t, k8sClientSource.Create(ctx, srcNamespace))
	defer func() {
		_ = k8sClientSource.Delete(ctx, srcNamespace)
	}()

	// 2. Verify that the namespace does NOT exist in the destination cluster
	destNamespace := &corev1.Namespace{}
	err := k8sClientDest.Get(ctx, types.NamespacedName{Name: ns}, destNamespace)
	require.Error(t, err)
	require.True(t, errors.IsNotFound(err))

	// Create Syncer
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: syncerName, Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Suspend: false,
			Mode:    krmv1alpha1.ModePush,
			Remote:  testRemote(destClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"spec"},
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))

	// Create target Service in Source
	target := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 80}},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, target))

	// 4. Verify that the namespace gets created in destination cluster AND the Service is successfully synced
	err = wait.PollUntilContextTimeout(ctx, 1*time.Second, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		destSvc := &corev1.Service{}
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		if err != nil {
			return false, nil
		}
		return len(destSvc.Spec.Ports) > 0 && destSvc.Spec.Ports[0].Port == 80, nil
	})
	assert.NoError(t, err, "Namespace and Service should be auto-created and synced to dest")

	// Verify namespace actually exists in destination cluster
	err = k8sClientDest.Get(ctx, types.NamespacedName{Name: ns}, destNamespace)
	assert.NoError(t, err, "Destination namespace should exist")

	// Cleanup
	require.NoError(t, k8sClientSource.Delete(ctx, target))
	require.NoError(t, k8sClientSource.Delete(ctx, syncer))
	// Cleanup destination namespace
	_ = k8sClientDest.Delete(ctx, destNamespace)
}

func TestSyncerRetryFailedApply(t *testing.T) {
	ctx := t.Context()
	ns := "retry-apply-ns"
	syncerName := "test-syncer-retry-apply"
	targetServiceName := "target-service-retry-apply"

	srcNamespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
	require.NoError(t, k8sClientSource.Create(ctx, srcNamespace))
	defer func() { _ = k8sClientSource.Delete(ctx, srcNamespace) }()

	// 1. Enable patch failure
	faultInjector.failPatch.Store(true)
	defer faultInjector.failPatch.Store(false)

	initialFailures := testutil.ToFloat64(SyncFailuresTotal.WithLabelValues("", "v1", "Service", OperationApply))

	// Create Syncer targeting dest-retry
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: syncerName, Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote(destRetryClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"spec"},
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

	// Create target Service in Source
	target := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 80}},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, target))
	defer func() { _ = k8sClientSource.Delete(ctx, target) }()

	// Wait until failure metric has incremented
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		current := testutil.ToFloat64(SyncFailuresTotal.WithLabelValues("", "v1", "Service", OperationApply))
		return current > initialFailures, nil
	})
	require.NoError(t, err, "Failure metric for apply should increment on failure")

	// Verify target service does NOT exist in destination yet
	destSvc := &corev1.Service{}
	err = k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
	require.True(t, errors.IsNotFound(err), "Service should not exist on dest while apply is failing")

	// Now recover: allow destination to succeed
	faultInjector.failPatch.Store(false)

	// Verify Service is retried and successfully synced to Dest
	err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		if err != nil {
			return false, nil
		}
		return len(destSvc.Spec.Ports) > 0 && destSvc.Spec.Ports[0].Port == 80, nil
	})
	assert.NoError(t, err, "Service should be retried and successfully synced to dest after recovery")
}

func TestSyncerRetryFailedDelete(t *testing.T) {
	ctx := t.Context()
	ns := "retry-delete-ns"
	syncerName := "test-syncer-retry-delete"
	targetServiceName := "target-service-retry-delete"

	srcNamespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
	require.NoError(t, k8sClientSource.Create(ctx, srcNamespace))
	defer func() { _ = k8sClientSource.Delete(ctx, srcNamespace) }()

	// Ensure injector is clean
	faultInjector.failPatch.Store(false)
	faultInjector.failDelete.Store(false)

	initialFailures := testutil.ToFloat64(SyncFailuresTotal.WithLabelValues("", "v1", "Service", OperationDelete))

	// Create Syncer targeting dest-retry
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: syncerName, Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote(destRetryClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"spec"},
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

	// Create target Service in Source
	target := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 80}},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, target))

	// Wait for initial sync to Dest
	destSvc := &corev1.Service{}
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		return err == nil, nil
	})
	require.NoError(t, err, "Service should be initially synced to dest")

	// Enable delete failure
	faultInjector.failDelete.Store(true)
	defer faultInjector.failDelete.Store(false)

	// Delete target Service in Source
	require.NoError(t, k8sClientSource.Delete(ctx, target))

	// Wait until delete failure metric increments
	err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		current := testutil.ToFloat64(SyncFailuresTotal.WithLabelValues("", "v1", "Service", OperationDelete))
		return current > initialFailures, nil
	})
	require.NoError(t, err, "Failure metric for delete should increment on failure")

	// Verify target service still exists in destination while failing
	err = k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
	require.NoError(t, err, "Service should still exist on dest while delete is failing")

	// Now recover: allow delete to succeed
	faultInjector.failDelete.Store(false)

	// Verify deletion is retried and succeeds on Dest
	err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		return errors.IsNotFound(err), nil
	})
	assert.NoError(t, err, "Service deletion should be retried and succeed on dest after recovery")
}

func TestSyncerRetryFailedStatus(t *testing.T) {
	ctx := t.Context()
	ns := "retry-status-ns"
	syncerName := "test-syncer-retry-status"
	targetServiceName := "target-service-retry-status"

	srcNamespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
	require.NoError(t, k8sClientSource.Create(ctx, srcNamespace))
	defer func() { _ = k8sClientSource.Delete(ctx, srcNamespace) }()

	// Ensure injector is clean
	faultInjector.failPatch.Store(false)
	faultInjector.failDelete.Store(false)
	faultInjector.failStatus.Store(false)

	initialFailures := testutil.ToFloat64(SyncFailuresTotal.WithLabelValues("", "v1", "Service", OperationStatus))

	// Create syncer targeting dest-retry FIRST
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: syncerName, Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote(destRetryClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"spec", "status"},
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

	// Create target Service in Source SECOND
	target := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 80}},
			Type:  corev1.ServiceTypeLoadBalancer,
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, target))
	defer func() { _ = k8sClientSource.Delete(ctx, target) }()

	// Wait for initial sync
	destSvc := &corev1.Service{}
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		return k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc) == nil, nil
	})
	require.NoError(t, err, "Service should be synced initially")

	// Enable status patch failure
	faultInjector.failStatus.Store(true)
	defer faultInjector.failStatus.Store(false)

	// Update status on source
	require.NoError(t, k8sClientSource.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, target))
	target.Status.LoadBalancer.Ingress = []corev1.LoadBalancerIngress{{IP: "5.6.7.8"}}
	require.NoError(t, k8sClientSource.Status().Update(ctx, target))

	// Wait until status failure metric increments
	err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		current := testutil.ToFloat64(SyncFailuresTotal.WithLabelValues("", "v1", "Service", OperationStatus))
		return current > initialFailures, nil
	})
	require.NoError(t, err, "Failure metric for status should increment on failure")

	// Now recover: allow status patch to succeed
	faultInjector.failStatus.Store(false)

	// Verify status is retried and updated on dest
	err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		if err != nil {
			return false, nil
		}
		return len(destSvc.Status.LoadBalancer.Ingress) > 0 && destSvc.Status.LoadBalancer.Ingress[0].IP == "5.6.7.8", nil
	})
	assert.NoError(t, err, "Status update should be retried and succeed on dest after recovery")
}

func TestSyncPeriodConfigurable(t *testing.T) {
	// 1. Verify command line flag definition, default (0 / disabled) and parsing
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	var sp time.Duration
	fs.DurationVar(&sp, "sync-period", 0, "The minimum frequency at which watched resources are reconciled (0 or negative disables periodic sync)")
	assert.Equal(t, time.Duration(0), sp, "default sync period should be 0 (no frequent sync)")

	err := fs.Parse([]string{"--sync-period=45m"})
	assert.NoError(t, err)
	assert.Equal(t, 45*time.Minute, sp, "sync period should be configurable")

	err = fs.Parse([]string{"--sync-period=-1s"})
	assert.NoError(t, err)
	assert.True(t, sp <= 0, "negative sync period should disable periodic sync")

	// 2. Verify that SyncPeriod is passed to remote cluster cache options in KRMSyncerReconciler
	syncPeriodRemote := 30 * time.Minute
	r := &KRMSyncerReconciler{
		Scheme:               scheme.Scheme,
		RemoteConfigProvider: testRemoteProvider,
		RemoteClusters:       make(map[string]cluster.Cluster),
		SyncPeriod:           &syncPeriodRemote,
	}

	remoteConfig := testRemote(destClusterName)
	c, err := r.getOrCreateRemoteCluster(t.Context(), "test-key-sync-period", "default", remoteConfig)
	require.NoError(t, err)
	require.NotNil(t, c)
	assert.NotNil(t, c.GetCache())

	// 3. E2E test showcasing that sync period works to retry an error:
	// A sync fails due to destination error, then recovers, retries, and succeeds.
	// Furthermore, when destination state is deleted out-of-band while source is untouched,
	// the periodic cache resync automatically triggers reconciliation and restores destination state.
	ctx := t.Context()
	ns := "sync-period-e2e-ns"
	srcNamespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
	require.NoError(t, k8sClientSource.Create(ctx, srcNamespace))
	defer func() { _ = k8sClientSource.Delete(ctx, srcNamespace) }()

	// Ensure injector is clean
	faultInjector.failPatch.Store(false)
	faultInjector.failDelete.Store(false)
	faultInjector.failStatus.Store(false)

	targetServiceName := "target-service-sync-period"
	initialFailures := testutil.ToFloat64(SyncFailuresTotal.WithLabelValues("", "v1", "Service", OperationApply))

	// Enable patch failure on destination to simulate error
	faultInjector.failPatch.Store(true)
	defer faultInjector.failPatch.Store(false)

	// Create a KRMSyncer rule targeting dest-retry
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: "syncer-sync-period", Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote(destRetryClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"spec"},
				},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

	// Create Service in Source
	target := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 8080}},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, target))
	defer func() { _ = k8sClientSource.Delete(ctx, target) }()

	// Wait until failure metric has incremented due to destination error
	err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		current := testutil.ToFloat64(SyncFailuresTotal.WithLabelValues("", "v1", "Service", OperationApply))
		return current > initialFailures, nil
	})
	require.NoError(t, err, "Failure metric for Service apply should increment")

	// Destination service should not exist while apply is failing
	destSvc := &corev1.Service{}
	err = k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
	require.True(t, errors.IsNotFound(err), "Service should not exist on dest while apply is failing")

	// Recover destination: allow apply to succeed
	faultInjector.failPatch.Store(false)

	// Verify Service is retried and successfully synced to destination
	err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		if err != nil {
			return false, nil
		}
		return len(destSvc.Spec.Ports) > 0 && destSvc.Spec.Ports[0].Port == 8080, nil
	})
	require.NoError(t, err, "Service should be retried and synced to destination")

	// Now prove that the sync period setting works to resync when state is lost on destination
	// without any updates to the source object:
	// Delete Service directly on destination
	require.NoError(t, k8sClientDest.Delete(ctx, destSvc))
	err = k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
	require.True(t, errors.IsNotFound(err), "Service should be deleted on destination")

	// Do NOT touch the source Service. The cache resync (SyncPeriod: 2s) will trigger reconciliation
	// and automatically recreate / re-apply the Service to destination.
	err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetServiceName, Namespace: ns}, destSvc)
		if err != nil {
			return false, nil
		}
		return len(destSvc.Spec.Ports) > 0 && destSvc.Spec.Ports[0].Port == 8080, nil
	})
	assert.NoError(t, err, "SyncPeriod should automatically trigger cache resync and restore destination object")
}
