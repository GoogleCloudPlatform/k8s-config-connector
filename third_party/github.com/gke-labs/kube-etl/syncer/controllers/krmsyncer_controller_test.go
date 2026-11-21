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
	"io"
	"k8s.io/client-go/rest"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/event"
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

// RoundTrip fails the selected requests as the server would (500), so they
// are object sync failures, not an unreachable remote.
func (rt *faultInjectingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == http.MethodPatch && rt.failStatus.Load() && strings.HasSuffix(req.URL.Path, "/status") {
		return statusResponse(req, http.StatusInternalServerError, "injected status patch failure"), nil
	}
	if req.Method == http.MethodPatch && rt.failPatch.Load() && !strings.HasSuffix(req.URL.Path, "/status") {
		return statusResponse(req, http.StatusInternalServerError, "injected apply patch failure"), nil
	}
	if req.Method == http.MethodDelete && rt.failDelete.Load() {
		return statusResponse(req, http.StatusInternalServerError, "injected delete failure"), nil
	}
	return rt.base.RoundTrip(req)
}

// statusResponse returns a Kubernetes API error response (a Status) with the
// HTTP status code and message.
func statusResponse(req *http.Request, code int, message string) *http.Response {
	body := fmt.Sprintf(`{"kind":"Status","apiVersion":"v1","status":"Failure","message":%q,"reason":%q,"code":%d}`, message, statusReason(code), code)
	return &http.Response{
		Status:     fmt.Sprintf("%d %s", code, http.StatusText(code)),
		StatusCode: code,
		Proto:      "HTTP/1.1", ProtoMajor: 1, ProtoMinor: 1,
		Header:  http.Header{"Content-Type": {"application/json"}},
		Body:    io.NopCloser(strings.NewReader(body)),
		Request: req,
	}
}

// testConfigProvider resolves fake GKE clusters (by name) to envtest clusters.
type testConfigProvider struct {
	mu       sync.RWMutex
	clusters map[string]*rest.Config
}

// add registers a test cluster while the controller may be running.
func (p *testConfigProvider) add(name string, cfg *rest.Config) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.clusters[name] = cfg
}

// Key uses the same connection key as the GKE provider.
func (p *testConfigProvider) Key(namespace string, remote *krmv1alpha1.RemoteConfig) (string, error) {
	return (&GKEConfigProvider{}).Key(namespace, remote)
}

func (p *testConfigProvider) RESTConfig(_ context.Context, _ string, remote *krmv1alpha1.RemoteConfig) (*rest.Config, error) {
	if _, err := gkeClusterName(remote); err != nil {
		return nil, err
	}
	p.mu.RLock()
	cfg, ok := p.clusters[remote.GKECluster.Name]
	p.mu.RUnlock()
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
		Client:                  mgr.GetClient(),
		Scheme:                  mgr.GetScheme(),
		Manager:                 mgr,
		RemoteConfigProvider:    testRemoteProvider,
		SyncPeriod:              &testSyncPeriod,
		MaxConcurrentReconciles: 10,
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
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

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
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

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

	// Set a status on the source. Status is not in syncFields, so it must not be synced.
	require.NoError(t, k8sClientSource.Get(ctx, types.NamespacedName{Name: targetName, Namespace: ns}, target))
	target.Status.Conditions = []metav1.Condition{{
		Type:               "Ready",
		Status:             metav1.ConditionTrue,
		Reason:             "Test",
		LastTransitionTime: metav1.Now(),
	}}
	require.NoError(t, k8sClientSource.Status().Update(ctx, target))

	// Update a spec field afterwards; once it shows up in dest, the status
	// update has been processed as well.
	require.NoError(t, k8sClientSource.Get(ctx, types.NamespacedName{Name: targetName, Namespace: ns}, target))
	target.Spec.Selector = map[string]string{"app": "test-fields-updated"}
	require.NoError(t, k8sClientSource.Update(ctx, target))

	destSvc := &corev1.Service{}
	err = wait.PollUntilContextTimeout(ctx, 1*time.Second, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		if err := k8sClientDest.Get(ctx, types.NamespacedName{Name: targetName, Namespace: ns}, destSvc); err != nil {
			return false, nil
		}
		return destSvc.Spec.Selector["app"] == "test-fields-updated", nil
	})
	require.NoError(t, err, "Service spec update should be synced to dest")
	assert.Empty(t, destSvc.Status.Conditions, "status is not in syncFields and must not be synced")
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
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

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
	defer func() { _ = k8sClientSource.Delete(ctx, validSyncer) }()

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
	defer func() { _ = k8sClientSource.Delete(ctx, defaultSyncer) }()
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
	// Use a dedicated namespace: Push KRMSyncers of other tests sync Services
	// in "default" from the local cluster to the remote one, and could push
	// the pulled Service back to the remote after it is deleted there.
	ns := "pull-ns"
	// The namespace is not deleted: envtest has no namespace controller, so it
	// would stay Terminating and break repeated runs (-count).
	for _, c := range []client.Client{k8sClientSource, k8sClientDest} {
		namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
		require.NoError(t, client.IgnoreAlreadyExists(c.Create(ctx, namespace)))
	}
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
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

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
			"status": map[string]any{
				"phase": "Active",
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

	// Status is not in syncFields, so it must not be synced.
	_, found, _ = unstructured.NestedFieldNoCopy(destFull.Object, "status")
	assert.False(t, found)
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

// Another worker can create the destination namespace between our Get and our
// Create. AlreadyExists must neither fail the sync nor count as a failure.
//
// The test runs a Pull mode DynamicResourceReconciler end to end (so failures
// are tracked as in production) against fake clients: the destination reports
// the namespace missing on Get, then returns AlreadyExists on Create.
func TestApplyToDestinationIgnoresNamespaceAlreadyExists(t *testing.T) {
	ctx := t.Context()
	const ns = "raced-ns"
	gvk := corev1.SchemeGroupVersion.WithKind("Service")
	key := types.NamespacedName{Name: "svc", Namespace: ns}

	// Source cluster: the object to sync.
	source := fake.NewClientBuilder().WithScheme(scheme.Scheme).WithObjects(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: key.Name, Namespace: key.Namespace},
		Spec:       corev1.ServiceSpec{Ports: []corev1.ServicePort{{Port: 80}}},
	}).Build()

	// Destination cluster
	remote := testRemote("source")
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: "syncer", Namespace: "default"},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePull,
			Remote: remote,
			Rules: []krmv1alpha1.ResourceRule{
				{Group: "", Version: "v1", Kind: "Service", Namespaces: []string{ns}, SyncFields: []string{"spec"}},
			},
		},
	}
	connKey, err := testRemoteProvider.Key(syncer.Namespace, remote)
	require.NoError(t, err)

	var nsCreateAttempted, applied bool
	dest := fake.NewClientBuilder().WithScheme(scheme.Scheme).WithObjects(syncer).WithInterceptorFuncs(interceptor.Funcs{
		// Another worker creates ns after our Get found it missing.
		Create: func(ctx context.Context, c client.WithWatch, obj client.Object, opts ...client.CreateOption) error {
			if _, ok := obj.(*corev1.Namespace); ok && obj.GetName() == ns {
				nsCreateAttempted = true
				return errors.NewAlreadyExists(corev1.Resource("namespaces"), ns)
			}
			return c.Create(ctx, obj, opts...)
		},
		// The fake client does not support server-side apply.
		Patch: func(_ context.Context, _ client.WithWatch, obj client.Object, _ client.Patch, _ ...client.PatchOption) error {
			if client.ObjectKeyFromObject(obj) == key {
				applied = true
			}
			return nil
		},
	}).Build()

	failedObjects := prometheus.NewGauge(prometheus.GaugeOpts{Name: "test_failed_objects"})
	r := &DynamicResourceReconciler{
		LocalClient:          dest,
		RemoteClient:         source,
		RemoteConfigProvider: testRemoteProvider,
		GVK:                  gvk,
		Mode:                 krmv1alpha1.ModePull,
		Remote:               RemoteGVK{Cluster: connKey, GVK: gvk},
		syncFailures:         newSyncFailureTracker(func(schema.GroupVersionKind) prometheus.Gauge { return failedObjects }),
	}

	_, err = r.Reconcile(ctx, ctrl.Request{NamespacedName: key})
	require.NoError(t, err, "AlreadyExists on the destination namespace should not fail the sync")
	assert.True(t, nsCreateAttempted, "the destination namespace should have been created")
	assert.True(t, applied, "the object should still be applied after AlreadyExists")
	assert.Equal(t, 0.0, testutil.ToFloat64(failedObjects), "AlreadyExists must not count as a sync failure")
}

// TestSyncerStatusModeMissingDestination verifies that in status mode (syncFields
// without "spec") a missing destination object is not created: the sync fails
// with an error (retried by the workqueue), without creating the namespace, and
// is counted in syncer_failed_objects until it succeeds once the destination
// object is created.
func TestSyncerStatusModeMissingDestination(t *testing.T) {
	ctx := t.Context()
	ns := "status-missing-dest-ns"
	syncerName := "test-syncer-status-missing-dest"
	targetServiceName := "target-service-status-missing-dest"
	key := types.NamespacedName{Name: targetServiceName, Namespace: ns}

	srcNamespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
	require.NoError(t, k8sClientSource.Create(ctx, srcNamespace))
	defer func() { _ = k8sClientSource.Delete(ctx, srcNamespace) }()

	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: syncerName, Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote(destClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{
					Group: "", Version: "v1", Kind: "Service",
					Namespaces: []string{ns},
					SyncFields: []string{"status"},
				},
			},
		},
	}
	// The KRMSyncer is only given to this reconciler (through a fake local
	// client), not created in the cluster: otherwise the manager's controllers
	// would sync the same object concurrently, and their applies could conflict.
	localClient := fake.NewClientBuilder().WithScheme(scheme.Scheme).WithObjects(syncer).Build()

	// Create the source Service with a status.
	target := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 80}},
			Type:  corev1.ServiceTypeLoadBalancer,
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, target))
	defer func() { _ = k8sClientSource.Delete(ctx, target) }()
	target.Status.LoadBalancer.Ingress = []corev1.LoadBalancerIngress{{IP: "1.2.3.4"}}
	require.NoError(t, k8sClientSource.Status().Update(ctx, target))
	failedObjects := prometheus.NewGauge(prometheus.GaugeOpts{Name: "test_failed_objects"})
	syncFailures := newSyncFailureTracker(func(schema.GroupVersionKind) prometheus.Gauge { return failedObjects })
	syncerFailures := func() int {
		n, _ := syncFailures.failures(client.ObjectKeyFromObject(syncer))
		return n
	}
	r := &DynamicResourceReconciler{
		LocalClient:          localClient,
		RemoteClient:         k8sClientSource,
		RemoteConfigProvider: testRemoteProvider,
		GVK:                  corev1.SchemeGroupVersion.WithKind("Service"),
		Mode:                 krmv1alpha1.ModePush,
		syncFailures:         syncFailures,
	}

	// The destination object does not exist: fail with a (retriable) error.
	_, err := r.Reconcile(ctx, ctrl.Request{NamespacedName: key})
	require.Error(t, err, "missing destination object should fail")
	assert.Contains(t, err.Error(), "does not exist")
	assert.Equal(t, 1.0, testutil.ToFloat64(failedObjects),
		"missing destination object should be counted as failed")
	assert.Equal(t, 1, syncerFailures(), "the failure should be counted for the KRMSyncer")
	_, failures := syncFailures.failures(client.ObjectKeyFromObject(syncer))
	assert.Equal(t, []syncFailure{{category: "destination object missing", kind: "Service"}}, failures)

	// Retrying while the destination object is still missing counts it once.
	_, err = r.Reconcile(ctx, ctrl.Request{NamespacedName: key})
	require.Error(t, err, "missing destination object should still fail")
	assert.Equal(t, 1.0, testutil.ToFloat64(failedObjects),
		"repeated failures of the same object should be counted once")
	assert.Equal(t, 1, syncerFailures(), "repeated failures of the same object should be counted once")

	err = k8sClientDest.Get(ctx, types.NamespacedName{Name: ns}, &corev1.Namespace{})
	assert.True(t, errors.IsNotFound(err), "destination namespace should not be created, got %v", err)
	err = k8sClientDest.Get(ctx, key, &corev1.Service{})
	assert.True(t, errors.IsNotFound(err), "destination object should not be created, got %v", err)

	// Create the destination object; the next reconcile syncs its status.
	destNamespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
	require.NoError(t, k8sClientDest.Create(ctx, destNamespace))
	defer func() { _ = k8sClientDest.Delete(ctx, destNamespace) }()
	destSvc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: targetServiceName, Namespace: ns},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{{Port: 80}},
			Type:  corev1.ServiceTypeLoadBalancer,
		},
	}
	require.NoError(t, k8sClientDest.Create(ctx, destSvc))
	defer func() { _ = k8sClientDest.Delete(ctx, destSvc) }()

	res, err := r.Reconcile(ctx, ctrl.Request{NamespacedName: key})
	require.NoError(t, err)
	assert.Zero(t, res.RequeueAfter)

	require.NoError(t, k8sClientDest.Get(ctx, key, destSvc))
	require.Len(t, destSvc.Status.LoadBalancer.Ingress, 1)
	assert.Equal(t, "1.2.3.4", destSvc.Status.LoadBalancer.Ingress[0].IP)

	// Once synced, the object is no longer counted as failed.
	assert.Equal(t, 0.0, testutil.ToFloat64(failedObjects),
		"synced object should no longer be counted as failed")
	assert.Zero(t, syncerFailures(), "synced object should no longer be counted for the KRMSyncer")
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

	initialFailed := testutil.ToFloat64(FailedObjects.WithLabelValues("", "v1", "Service"))

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
		return testutil.ToFloat64(FailedObjects.WithLabelValues("", "v1", "Service")) > initialFailed, nil
	})
	require.NoError(t, err, "Failed objects gauge should count the object while apply is failing")
	waitForSyncerStatus(ctx, t, syncer, 1, krmv1alpha1.ReasonSyncFailed)

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
	waitForFailedObjects(ctx, t, initialFailed)
	waitForSyncerStatus(ctx, t, syncer, 0, krmv1alpha1.ReasonSyncing)
}

// waitForSyncerStatus waits until the KRMSyncer reports failedObjects and the
// Ready reason.
func waitForSyncerStatus(ctx context.Context, t *testing.T, syncer *krmv1alpha1.KRMSyncer, failedObjects int32, reason string) {
	t.Helper()
	var last krmv1alpha1.KRMSyncerStatus
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		s := &krmv1alpha1.KRMSyncer{}
		if err := k8sClientSource.Get(ctx, client.ObjectKeyFromObject(syncer), s); err != nil {
			return false, nil
		}
		last = s.Status
		ready := meta.FindStatusCondition(s.Status.Conditions, krmv1alpha1.ConditionReady)
		return s.Status.FailedObjects == failedObjects && ready != nil && ready.Reason == reason, nil
	})
	require.NoError(t, err, "want failedObjects=%d and Ready reason %s, last status: %+v", failedObjects, reason, last)
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

	initialFailed := testutil.ToFloat64(FailedObjects.WithLabelValues("", "v1", "Service"))

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
		return testutil.ToFloat64(FailedObjects.WithLabelValues("", "v1", "Service")) > initialFailed, nil
	})
	require.NoError(t, err, "Failed objects gauge should count the object while delete is failing")

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
	waitForFailedObjects(ctx, t, initialFailed)
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

	initialFailed := testutil.ToFloat64(FailedObjects.WithLabelValues("", "v1", "Service"))

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
		return testutil.ToFloat64(FailedObjects.WithLabelValues("", "v1", "Service")) > initialFailed, nil
	})
	require.NoError(t, err, "Failed objects gauge should count the object while status patch is failing")

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
	waitForFailedObjects(ctx, t, initialFailed)
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
	initialFailed := testutil.ToFloat64(FailedObjects.WithLabelValues("", "v1", "Service"))

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
		return testutil.ToFloat64(FailedObjects.WithLabelValues("", "v1", "Service")) > initialFailed, nil
	})
	require.NoError(t, err, "Failed objects gauge should count the Service while apply is failing")

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
	waitForFailedObjects(ctx, t, initialFailed)

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

// waitForFailedObjects waits until syncer_failed_objects for v1 Services equals want.
func waitForFailedObjects(ctx context.Context, t *testing.T, want float64) {
	t.Helper()
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		return testutil.ToFloat64(FailedObjects.WithLabelValues("", "v1", "Service")) == want, nil
	})
	require.NoError(t, err, "a recovered object should no longer be counted as failed")
}

func TestSetupFailedGauge(t *testing.T) {
	ctx := t.Context()
	ns := "setup-failed-ns"
	srcNamespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}
	require.NoError(t, k8sClientSource.Create(ctx, srcNamespace))
	defer func() { _ = k8sClientSource.Delete(ctx, srcNamespace) }()

	initial := testutil.ToFloat64(SetupFailed)

	// Pull from a cluster that cannot be connected to: setup keeps failing.
	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: "syncer-unreachable", Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePull,
			Remote: testRemote("does-not-exist"),
			Rules: []krmv1alpha1.ResourceRule{
				{Group: "", Version: "v1", Kind: "ConfigMap", Namespaces: []string{ns}},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))

	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		return testutil.ToFloat64(SetupFailed) == initial+1, nil
	})
	require.NoError(t, err, "a KRMSyncer whose setup fails should be counted")

	// Retries keep failing, but the KRMSyncer is still counted once.
	time.Sleep(time.Second)
	assert.Equal(t, initial+1, testutil.ToFloat64(SetupFailed))

	// Deleting the KRMSyncer clears it.
	require.NoError(t, k8sClientSource.Delete(ctx, syncer))
	err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 15*time.Second, true, func(ctx context.Context) (bool, error) {
		return testutil.ToFloat64(SetupFailed) == initial, nil
	})
	require.NoError(t, err, "a deleted KRMSyncer should no longer be counted")
}

func TestSyncErrorCategory(t *testing.T) {
	gr := corev1.Resource("services")
	for _, tc := range []struct {
		name string
		err  error
		want string
	}{
		{"not an API error", newSyncError("destination object missing", stderrors.New("missing")), "destination object missing"},
		{"API error", newSyncError("apply failed", errors.NewForbidden(gr, "svc", stderrors.New("denied"))), "apply failed (Forbidden)"},
		{"wrapped API error", newSyncError("apply failed", fmt.Errorf("applying: %w", errors.NewInvalid(corev1.SchemeGroupVersion.WithKind("Service").GroupKind(), "svc", nil))), "apply failed (Invalid)"},
		{"wrapped sync error", fmt.Errorf("rule 1: %w", newSyncError("delete failed", errors.NewServiceUnavailable("down"))), "delete failed (ServiceUnavailable)"},
		{"uncategorized error", stderrors.New("boom"), "sync failed"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, syncErrorCategory(tc.err))
		})
	}
	// The error message is unchanged.
	assert.Equal(t, "missing", newSyncError("destination object missing", stderrors.New("missing")).Error())
}

// TestKRMSyncerEventFilter verifies which KRMSyncer events run setup. The test
// manager resyncs every 2s, which would hide a broken filter in the end-to-end
// tests.
func TestKRMSyncerEventFilter(t *testing.T) {
	syncer := func(generation int64, resourceVersion string) *krmv1alpha1.KRMSyncer {
		return &krmv1alpha1.KRMSyncer{ObjectMeta: metav1.ObjectMeta{
			Name: "syncer", Namespace: "default", Generation: generation, ResourceVersion: resourceVersion,
		}}
	}
	withStatus := syncer(1, "11")
	withStatus.Status.FailedObjects = 3
	withLabels := syncer(1, "11")
	withLabels.Labels = map[string]string{"team": "a"}
	withLabels.Annotations = map[string]string{"note": "b"}

	p := krmSyncerEventFilter()
	assert.True(t, p.Create(event.CreateEvent{Object: syncer(1, "10")}), "create")
	assert.True(t, p.Delete(event.DeleteEvent{Object: syncer(1, "10")}), "delete (Reconcile clears the trackers)")
	for _, tc := range []struct {
		name     string
		old, new *krmv1alpha1.KRMSyncer
		want     bool
	}{
		{name: "spec change (new generation)", old: syncer(1, "10"), new: syncer(2, "11"), want: true},
		{name: "status only (e.g. our own status write)", old: syncer(1, "10"), new: withStatus, want: false},
		{name: "labels and annotations only", old: syncer(1, "10"), new: withLabels, want: false},
		{name: "periodic resync (unchanged object)", old: syncer(1, "10"), new: syncer(1, "10"), want: true},
	} {
		assert.Equal(t, tc.want, p.Update(event.UpdateEvent{ObjectOld: tc.old, ObjectNew: tc.new}), tc.name)
	}
}

// TestGlobResyncPicksUpNewCRDs verifies that a periodic resync re-runs setup,
// so a glob rule starts syncing a CRD installed after the KRMSyncer, without a
// spec change.
func TestGlobResyncPicksUpNewCRDs(t *testing.T) {
	ctx := t.Context()
	suffix := utilrand.String(5)
	ns := "glob-resync-" + suffix
	group := "resync" + suffix + ".cnrm.cloud.google.com"
	require.NoError(t, k8sClientSource.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}))

	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: "glob-syncer", Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote(destClusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{Group: "*.cnrm.cloud.google.com", Version: "*", Kind: "*", Namespaces: []string{ns}, SyncFields: []string{"spec"}},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()
	waitForReadyCondition(ctx, t, syncer, metav1.ConditionTrue, krmv1alpha1.ReasonSyncing, "")

	// Install a new KCC CRD on both clusters, after the KRMSyncer was set up.
	for _, c := range []client.Client{k8sClientSource, k8sClientDest} {
		installWidgetCRD(ctx, t, c, group)
	}

	widget := &unstructured.Unstructured{Object: map[string]any{
		"apiVersion": group + "/v1",
		"kind":       "Widget",
		"metadata":   map[string]any{"name": "widget", "namespace": ns},
		"spec":       map[string]any{"foo": "bar"},
	}}
	err := wait.PollUntilContextTimeout(ctx, 200*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		return k8sClientSource.Create(ctx, widget.DeepCopy()) == nil, nil
	})
	require.NoError(t, err, "Widget should be created once the CRD is served")

	// The next resync re-runs setup, which discovers the new kind and syncs it.
	err = wait.PollUntilContextTimeout(ctx, 200*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		got := &unstructured.Unstructured{}
		got.SetGroupVersionKind(widget.GroupVersionKind())
		if err := k8sClientDest.Get(ctx, client.ObjectKeyFromObject(widget), got); err != nil {
			return false, nil
		}
		foo, _, _ := unstructured.NestedString(got.Object, "spec", "foo")
		return foo == "bar", nil
	})
	require.NoError(t, err, "the Widget should be pushed after a resync, with no spec change")
}

// installWidgetCRD installs a namespaced Widget CRD in group and waits until
// it is established.
func installWidgetCRD(ctx context.Context, t *testing.T, c client.Client, group string) {
	t.Helper()
	crd := &unstructured.Unstructured{Object: map[string]any{
		"apiVersion": "apiextensions.k8s.io/v1",
		"kind":       "CustomResourceDefinition",
		"metadata":   map[string]any{"name": "widgets." + group},
		"spec": map[string]any{
			"group": group,
			"scope": "Namespaced",
			"names": map[string]any{"kind": "Widget", "listKind": "WidgetList", "plural": "widgets", "singular": "widget"},
			"versions": []any{map[string]any{
				"name": "v1", "served": true, "storage": true,
				"schema": map[string]any{"openAPIV3Schema": map[string]any{
					"type": "object", "x-kubernetes-preserve-unknown-fields": true,
				}},
			}},
		},
	}}
	require.NoError(t, c.Create(ctx, crd))
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		got := &unstructured.Unstructured{}
		got.SetGroupVersionKind(crd.GroupVersionKind())
		if err := c.Get(ctx, client.ObjectKeyFromObject(crd), got); err != nil {
			return false, nil
		}
		conditions, _, _ := unstructured.NestedSlice(got.Object, "status", "conditions")
		for _, cond := range conditions {
			if m, ok := cond.(map[string]any); ok && m["type"] == "Established" && m["status"] == "True" {
				return true, nil
			}
		}
		return false, nil
	})
	require.NoError(t, err, "CRD %s should be established", crd.GetName())
}
