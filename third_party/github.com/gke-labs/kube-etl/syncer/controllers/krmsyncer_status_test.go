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
	"context"
	stderrors "errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unicode/utf8"

	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
)

// TestReadyCondition reconciles KRMSyncers against a fake API server and checks
// the Ready condition, and that status is only written when it changes.
func TestReadyCondition(t *testing.T) {
	ctx := t.Context()
	var statusWrites atomic.Int32
	newReconciler := func(objs ...client.Object) (*KRMSyncerReconciler, client.Client) {
		c := fake.NewClientBuilder().WithScheme(scheme.Scheme).
			WithObjects(objs...).
			WithStatusSubresource(&krmv1alpha1.KRMSyncer{}).
			WithInterceptorFuncs(interceptor.Funcs{
				SubResourceUpdate: func(ctx context.Context, c client.Client, sub string, obj client.Object, opts ...client.SubResourceUpdateOption) error {
					statusWrites.Add(1)
					return c.SubResource(sub).Update(ctx, obj, opts...)
				},
				SubResourcePatch: func(ctx context.Context, c client.Client, sub string, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
					statusWrites.Add(1)
					return c.SubResource(sub).Patch(ctx, obj, patch, opts...)
				},
			}).Build()
		return &KRMSyncerReconciler{
			Client:               c,
			Scheme:               scheme.Scheme,
			RemoteConfigProvider: testRemoteProvider,
			WatchedGVKs:          make(map[schema.GroupVersionKind]bool),
			WatchedRemoteGVKs:    make(map[RemoteGVK]bool),
			RemoteClusters:       make(map[string]cluster.Cluster),
			setupFailures:        newSetupFailureTracker(prometheus.NewGauge(prometheus.GaugeOpts{Name: "test_setup_failed"})),
		}, c
	}
	reconcileReady := func(t *testing.T, r *KRMSyncerReconciler, c client.Client, key types.NamespacedName) (*krmv1alpha1.KRMSyncer, *metav1.Condition, error) {
		t.Helper()
		_, err := r.Reconcile(ctx, ctrl.Request{NamespacedName: key})
		s := &krmv1alpha1.KRMSyncer{}
		require.NoError(t, c.Get(ctx, key, s))
		ready := meta.FindStatusCondition(s.Status.Conditions, krmv1alpha1.ConditionReady)
		require.NotNil(t, ready, "Ready should be set")
		return s, ready, err
	}

	t.Run("syncing, then suspended, written only on change", func(t *testing.T) {
		statusWrites.Store(0)
		syncer := &krmv1alpha1.KRMSyncer{
			ObjectMeta: metav1.ObjectMeta{Name: "ready", Namespace: "default", Generation: 3},
			// Push mode with no rules: setup has nothing to do.
			Spec: krmv1alpha1.KRMSyncerSpec{Mode: krmv1alpha1.ModePush, Remote: testRemote(destClusterName)},
			// Conditions written by earlier versions are replaced by Ready.
			Status: krmv1alpha1.KRMSyncerStatus{Conditions: []metav1.Condition{
				{Type: "Active", Status: metav1.ConditionTrue, Reason: "Active", LastTransitionTime: metav1.Now()},
			}},
		}
		r, c := newReconciler(syncer)
		key := client.ObjectKeyFromObject(syncer)

		s, ready, err := reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Equal(t, metav1.ConditionTrue, ready.Status)
		assert.Equal(t, krmv1alpha1.ReasonSyncing, ready.Reason)
		assert.Equal(t, s.Generation, ready.ObservedGeneration)
		assert.Equal(t, s.Generation, s.Status.ObservedGeneration)
		assert.Nil(t, meta.FindStatusCondition(s.Status.Conditions, "Active"), "legacy conditions should be removed")
		assert.Equal(t, int32(1), statusWrites.Load())

		// Nothing changed: no status write.
		_, _, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		_, _, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Equal(t, int32(1), statusWrites.Load(), "status should not be written when it did not change")

		// Suspend.
		s.Spec.Suspend = true
		require.NoError(t, c.Update(ctx, s))
		_, ready, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Equal(t, metav1.ConditionTrue, ready.Status, "a suspended KRMSyncer is Ready")
		assert.Equal(t, krmv1alpha1.ReasonSuspended, ready.Reason)
		assert.Equal(t, int32(2), statusWrites.Load())
		_, _, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Equal(t, int32(2), statusWrites.Load(), "status should not be written when it did not change")
	})

	t.Run("setup failed", func(t *testing.T) {
		syncer := &krmv1alpha1.KRMSyncer{
			ObjectMeta: metav1.ObjectMeta{Name: "setup-failed", Namespace: "default"},
			Spec: krmv1alpha1.KRMSyncerSpec{
				Mode:   krmv1alpha1.ModePull,
				Remote: testRemote("does-not-exist"),
				Rules:  []krmv1alpha1.ResourceRule{{Version: "v1", Kind: "ConfigMap"}},
			},
		}
		r, c := newReconciler(syncer)
		_, ready, err := reconcileReady(t, r, c, client.ObjectKeyFromObject(syncer))
		require.Error(t, err, "a failed setup should be retried")
		assert.Equal(t, metav1.ConditionFalse, ready.Status)
		assert.Equal(t, krmv1alpha1.ReasonSetupFailed, ready.Reason)
		assert.Contains(t, ready.Message, `unknown test cluster "does-not-exist"`)
	})

	t.Run("invalid rule", func(t *testing.T) {
		syncer := &krmv1alpha1.KRMSyncer{
			ObjectMeta: metav1.ObjectMeta{Name: "invalid-rule", Namespace: "default"},
			Spec: krmv1alpha1.KRMSyncerSpec{
				// Globs run discovery against the remote (envtest) in Pull mode.
				Mode:   krmv1alpha1.ModePull,
				Remote: testRemote(destClusterName),
				Rules:  []krmv1alpha1.ResourceRule{{Group: "*", Version: "v1", Kind: "ConfigMap"}},
			},
		}
		r, c := newReconciler(syncer)
		_, ready, err := reconcileReady(t, r, c, client.ObjectKeyFromObject(syncer))
		require.Error(t, err)
		assert.Equal(t, krmv1alpha1.ReasonSetupFailed, ready.Reason)
		assert.Contains(t, ready.Message, "invalid rule")
	})

	t.Run("remote unreachable takes precedence over setup failure", func(t *testing.T) {
		remote := testRemote("does-not-exist")
		syncer := &krmv1alpha1.KRMSyncer{
			ObjectMeta: metav1.ObjectMeta{Name: "unreachable", Namespace: "default"},
			Spec: krmv1alpha1.KRMSyncerSpec{
				Mode:   krmv1alpha1.ModePull,
				Remote: remote,
				Rules:  []krmv1alpha1.ResourceRule{{Version: "v1", Kind: "ConfigMap"}},
			},
		}
		r, c := newReconciler(syncer)
		connKey, err := testRemoteProvider.Key(syncer.Namespace, remote)
		require.NoError(t, err)
		r.remotes = newRemoteHealthTracker()
		r.remotes.get(connKey).markUnreachable(stderrors.New("Unauthorized"))

		_, ready, err := reconcileReady(t, r, c, client.ObjectKeyFromObject(syncer))
		require.Error(t, err)
		assert.Equal(t, metav1.ConditionFalse, ready.Status)
		assert.Equal(t, krmv1alpha1.ReasonRemoteUnreachable, ready.Reason)
		assert.Equal(t, "remote cluster projects/test-project/locations/us-central1/clusters/does-not-exist: Unauthorized", ready.Message)
	})

	t.Run("sync failed", func(t *testing.T) {
		syncer := &krmv1alpha1.KRMSyncer{
			ObjectMeta: metav1.ObjectMeta{Name: "sync-failed", Namespace: "default"},
			Spec:       krmv1alpha1.KRMSyncerSpec{Mode: krmv1alpha1.ModePush, Remote: testRemote(destClusterName)},
		}
		r, c := newReconciler(syncer)
		r.syncFailures = newSyncFailureTracker(nil)
		key := client.ObjectKeyFromObject(syncer)
		svc := func(name string) objectRef {
			return objectRef{GVK: schema.GroupVersionKind{Version: "v1", Kind: "Service"}, Key: types.NamespacedName{Namespace: "ns", Name: name}}
		}
		const missing, forbidden = "destination object missing", "apply failed (Forbidden)"
		r.syncFailures.markFailed(key, svc("a"), missing)
		r.syncFailures.markFailed(key, svc("b"), forbidden)
		r.syncFailures.markFailed(key, svc("c"), missing)
		// Failures of other KRMSyncers are not counted.
		r.syncFailures.markFailed(types.NamespacedName{Namespace: "default", Name: "other"}, svc("e"), "delete failed")

		// The message groups the failures by error, with the kinds each one
		// affects; the count is in failedObjects.
		s, ready, err := reconcileReady(t, r, c, key)
		require.NoError(t, err, "failed objects are retried by their own controllers")
		assert.Equal(t, int32(3), s.Status.FailedObjects)
		assert.Equal(t, metav1.ConditionFalse, ready.Status)
		assert.Equal(t, krmv1alpha1.ReasonSyncFailed, ready.Reason)
		assert.Equal(t, "apply failed (Forbidden): Service; destination object missing: Service. See the controller logs for details.", ready.Message)

		// An object failing with another error moves to that category. While
		// the set of failures is the same, so is the message.
		r.syncFailures.markFailed(key, svc("a"), forbidden)
		s, ready, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Equal(t, int32(3), s.Status.FailedObjects)
		assert.Equal(t, "apply failed (Forbidden): Service; destination object missing: Service. See the controller logs for details.", ready.Message)
		r.syncFailures.markFailed(key, svc("c"), forbidden)
		_, ready, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Equal(t, "apply failed (Forbidden): Service. See the controller logs for details.", ready.Message)

		// At most maxErrorsInMessage errors are listed.
		for i, name := range []string{"d", "e", "f", "g", "h"} {
			r.syncFailures.markFailed(key, svc(name), fmt.Sprintf("category %d", i))
		}
		_, ready, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Equal(t, "apply failed (Forbidden): Service; category 0: Service; category 1: Service; category 2: Service; category 3: Service; and 1 more error. See the controller logs for details.", ready.Message)
		for _, name := range []string{"a", "c", "d", "e", "f", "g", "h"} {
			r.syncFailures.markSucceeded(key, svc(name))
		}

		// At most maxKindsPerErrorInMessage kinds are listed per error, and an
		// error with many kinds does not hide the others.
		testCRD := func(kind string) objectRef {
			return objectRef{GVK: schema.GroupVersionKind{Group: "e2e.gkelabs.io", Version: "v1alpha1", Kind: kind}, Key: types.NamespacedName{Namespace: "ns", Name: "obj"}}
		}
		kinds := []string{"TestCRD", "TestCRD2", "TestCRD3", "TestCRD4", "TestCRD5", "TestCRD6", "TestCRD7", "TestCRD8"}
		for _, kind := range kinds {
			r.syncFailures.markFailed(key, testCRD(kind), missing)
		}
		s, ready, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Equal(t, int32(9), s.Status.FailedObjects)
		assert.Equal(t, "apply failed (Forbidden): Service; destination object missing: TestCRD.e2e.gkelabs.io, TestCRD2.e2e.gkelabs.io, TestCRD3.e2e.gkelabs.io, and 5 more kinds. See the controller logs for details.", ready.Message)

		r.syncFailures.markSucceeded(key, svc("b"))
		for _, kind := range kinds {
			r.syncFailures.markSucceeded(key, testCRD(kind))
		}
		s, ready, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Zero(t, s.Status.FailedObjects)
		assert.Equal(t, krmv1alpha1.ReasonSyncing, ready.Reason)

		// Suspending forgets the failures.
		r.syncFailures.markFailed(key, svc("a"), missing)
		s.Spec.Suspend = true
		require.NoError(t, c.Update(ctx, s))
		s, ready, err = reconcileReady(t, r, c, key)
		require.NoError(t, err)
		assert.Zero(t, s.Status.FailedObjects)
		assert.Equal(t, krmv1alpha1.ReasonSuspended, ready.Reason)
	})
}

// flakyRoundTripper passes requests through, or fails all of them with a
// given HTTP status.
type flakyRoundTripper struct {
	base   http.RoundTripper
	status atomic.Int32 // 0 passes requests through

	mu     sync.Mutex
	bodies map[*trackedBody]struct{}
}

func (rt *flakyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if code := int(rt.status.Load()); code != 0 {
		return statusResponse(req, code, http.StatusText(code)), nil
	}
	resp, err := rt.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	b := &trackedBody{ReadCloser: resp.Body, rt: rt}
	rt.mu.Lock()
	if rt.bodies == nil {
		rt.bodies = make(map[*trackedBody]struct{})
	}
	rt.bodies[b] = struct{}{}
	rt.mu.Unlock()
	resp.Body = b
	return resp, nil
}

// failWith makes all requests fail with the HTTP status code (0 to recover).
// Open responses (e.g. watches) are closed, so clients send new requests.
func (rt *flakyRoundTripper) failWith(code int) {
	rt.status.Store(int32(code))
	rt.mu.Lock()
	bodies := rt.bodies
	rt.bodies = nil
	rt.mu.Unlock()
	for b := range bodies {
		_ = b.ReadCloser.Close()
	}
}

type trackedBody struct {
	io.ReadCloser
	rt *flakyRoundTripper
}

func (b *trackedBody) Close() error {
	b.rt.mu.Lock()
	delete(b.rt.bodies, b)
	b.rt.mu.Unlock()
	return b.ReadCloser.Close()
}

// TestReadyRemoteUnreachableAfterSetup verifies that a remote failing after a
// successful setup is reported in Ready (RemoteUnreachable), and that Ready
// recovers once the remote is reachable again.
func TestReadyRemoteUnreachableAfterSetup(t *testing.T) {
	ctx := t.Context()

	clusterName, flaky := newFlakyRemote()

	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: "syncer-" + clusterName, Namespace: "default"},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePull,
			Remote: testRemote(clusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{Version: "v1", Kind: "ConfigMap", Namespaces: []string{"flaky-remote-ns"}},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

	waitForReady := func(status metav1.ConditionStatus, reason, message string) {
		t.Helper()
		waitForReadyCondition(ctx, t, syncer, status, reason, message)
	}

	waitForReady(metav1.ConditionTrue, krmv1alpha1.ReasonSyncing, "")

	// The remote starts rejecting our credentials.
	flaky.failWith(http.StatusUnauthorized)
	waitForReady(metav1.ConditionFalse, krmv1alpha1.ReasonRemoteUnreachable,
		"remote cluster projects/test-project/locations/us-central1/clusters/"+clusterName+": Unauthorized")

	// The remote recovers.
	flaky.failWith(0)
	waitForReady(metav1.ConditionTrue, krmv1alpha1.ReasonSyncing, "")

	// The remote refuses to serve a synced kind: reported for that kind, not
	// as an unreachable remote.
	flaky.failWith(http.StatusForbidden)
	waitForReady(metav1.ConditionFalse, krmv1alpha1.ReasonSyncFailed,
		"watch failed (Forbidden): ConfigMap. See the controller logs for details.")

	flaky.failWith(0)
	waitForReady(metav1.ConditionTrue, krmv1alpha1.ReasonSyncing, "")
}

// TestReadyRemoteUnreachablePush verifies that a Push mode destination that
// fails (401) is reported as RemoteUnreachable when syncing, and that Ready
// recovers once syncing succeeds again.
func TestReadyRemoteUnreachablePush(t *testing.T) {
	ctx := t.Context()
	clusterName, flaky := newFlakyRemote()
	ns := "push-" + clusterName
	require.NoError(t, k8sClientSource.Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}}))

	syncer := &krmv1alpha1.KRMSyncer{
		ObjectMeta: metav1.ObjectMeta{Name: "syncer", Namespace: ns},
		Spec: krmv1alpha1.KRMSyncerSpec{
			Mode:   krmv1alpha1.ModePush,
			Remote: testRemote(clusterName),
			Rules: []krmv1alpha1.ResourceRule{
				{Version: "v1", Kind: "Service", Namespaces: []string{ns}, SyncFields: []string{"spec"}},
			},
		},
	}
	require.NoError(t, k8sClientSource.Create(ctx, syncer))
	defer func() { _ = k8sClientSource.Delete(ctx, syncer) }()

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "svc", Namespace: ns},
		Spec:       corev1.ServiceSpec{Ports: []corev1.ServicePort{{Port: 80}}},
	}
	require.NoError(t, k8sClientSource.Create(ctx, svc))
	defer func() { _ = k8sClientSource.Delete(ctx, svc) }()

	// Push talks to the destination while syncing: the first sync succeeds.
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (bool, error) {
		return k8sClientDest.Get(ctx, client.ObjectKeyFromObject(svc), &corev1.Service{}) == nil, nil
	})
	require.NoError(t, err, "Service should be pushed to the destination")
	waitForReadyCondition(ctx, t, syncer, metav1.ConditionTrue, krmv1alpha1.ReasonSyncing, "")

	// The destination starts rejecting our credentials; the next sync fails.
	flaky.failWith(http.StatusUnauthorized)
	require.NoError(t, k8sClientSource.Get(ctx, client.ObjectKeyFromObject(svc), svc))
	svc.Spec.Ports[0].Port = 8080
	require.NoError(t, k8sClientSource.Update(ctx, svc))
	waitForReadyCondition(ctx, t, syncer, metav1.ConditionFalse, krmv1alpha1.ReasonRemoteUnreachable,
		"remote cluster projects/test-project/locations/us-central1/clusters/"+clusterName+": Unauthorized")

	// The destination recovers: the retried sync succeeds.
	flaky.failWith(0)
	waitForReadyCondition(ctx, t, syncer, metav1.ConditionTrue, krmv1alpha1.ReasonSyncing, "")
}

// newFlakyRemote registers a new remote cluster (resolving to the destination
// cluster) whose requests can be made to fail, and returns its name. A new
// remote is used on each run: the informers of a remote that already failed
// retry with a longer backoff, which would slow down repeated runs (-count).
func newFlakyRemote() (string, *flakyRoundTripper) {
	clusterName := "dest-flaky-" + utilrand.String(5)
	flaky := &flakyRoundTripper{}
	cfg := rest.CopyConfig(cfgDest)
	cfg.Wrap(func(rt http.RoundTripper) http.RoundTripper {
		flaky.base = rt
		return flaky
	})
	testRemoteProvider.add(clusterName, cfg)
	return clusterName, flaky
}

// waitForReadyCondition waits until syncer has the Ready condition (and, if
// set, message), for its current generation.
func waitForReadyCondition(ctx context.Context, t *testing.T, syncer *krmv1alpha1.KRMSyncer, status metav1.ConditionStatus, reason, message string) {
	t.Helper()
	var last *metav1.Condition
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 60*time.Second, true, func(ctx context.Context) (bool, error) {
		s := &krmv1alpha1.KRMSyncer{}
		if err := k8sClientSource.Get(ctx, client.ObjectKeyFromObject(syncer), s); err != nil {
			return false, nil
		}
		last = meta.FindStatusCondition(s.Status.Conditions, krmv1alpha1.ConditionReady)
		return last != nil && last.Status == status && last.Reason == reason &&
			(message == "" || last.Message == message) &&
			s.Status.ObservedGeneration == s.Generation, nil
	})
	require.NoError(t, err, "want Ready=%s/%s, last: %+v", status, reason, last)
}

// TestReadyRemoteKindFailures verifies that kinds the remote refuses to serve
// (403/5xx) are reported per kind, only for KRMSyncers syncing them, and that
// RemoteUnreachable takes precedence.
func TestReadyRemoteKindFailures(t *testing.T) {
	remote := testRemote("refusing")
	key, err := testRemoteProvider.Key("default", remote)
	require.NoError(t, err)
	configMaps := schema.GroupVersionKind{Version: "v1", Kind: "ConfigMap"}
	testCRDs := schema.GroupVersionKind{Group: "e2e.gkelabs.io", Version: "v1alpha1", Kind: "TestCRD"}
	mapper := meta.NewDefaultRESTMapper(nil)
	mapper.Add(configMaps, meta.RESTScopeNamespace)
	mapper.Add(testCRDs, meta.RESTScopeNamespace)

	r := &KRMSyncerReconciler{RemoteConfigProvider: testRemoteProvider, remotes: newRemoteHealthTracker()}
	h := r.remotes.get(key)
	h.setMapper(mapper)
	h.markResource(schema.GroupVersionResource{Group: "e2e.gkelabs.io", Version: "v1alpha1", Resource: "testcrds"}, "Forbidden")
	h.markResource(schema.GroupVersionResource{Version: "v1", Resource: "configmaps"}, "InternalError")

	ready := func(rules ...krmv1alpha1.ResourceRule) *metav1.Condition {
		t.Helper()
		s := &krmv1alpha1.KRMSyncer{
			ObjectMeta: metav1.ObjectMeta{Name: "syncer", Namespace: "default"},
			Spec:       krmv1alpha1.KRMSyncerSpec{Mode: krmv1alpha1.ModePull, Remote: remote, Rules: rules},
		}
		r.setReadyCondition(s, nil)
		return meta.FindStatusCondition(s.Status.Conditions, krmv1alpha1.ConditionReady)
	}
	both := []krmv1alpha1.ResourceRule{{Version: "v1", Kind: "ConfigMap"}, {Group: "e2e.gkelabs.io", Version: "v1alpha1", Kind: "TestCRD"}}

	cond := ready(both...)
	assert.Equal(t, krmv1alpha1.ReasonSyncFailed, cond.Reason)
	assert.Equal(t, "watch failed (Forbidden): TestCRD.e2e.gkelabs.io; watch failed (InternalError): ConfigMap. See the controller logs for details.", cond.Message)

	// Only kinds the KRMSyncer syncs are reported.
	cond = ready(krmv1alpha1.ResourceRule{Version: "v1", Kind: "ConfigMap"})
	assert.Equal(t, "watch failed (InternalError): ConfigMap. See the controller logs for details.", cond.Message)
	cond = ready(krmv1alpha1.ResourceRule{Version: "v1", Kind: "Secret"})
	assert.Equal(t, krmv1alpha1.ReasonSyncing, cond.Reason)

	// An unreachable remote takes precedence.
	h.markUnreachable(stderrors.New("Unauthorized"))
	cond = ready(both...)
	assert.Equal(t, krmv1alpha1.ReasonRemoteUnreachable, cond.Reason)
}

// TestUpdateAllStatus verifies that the periodic status update writes status
// once per change, however many objects change, and never when nothing
// changed.
func TestUpdateAllStatus(t *testing.T) {
	newSyncer := func(name string, generation int64) *krmv1alpha1.KRMSyncer {
		return &krmv1alpha1.KRMSyncer{
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: "default", Generation: generation},
			Spec:       krmv1alpha1.KRMSyncerSpec{Mode: krmv1alpha1.ModePush, Remote: testRemote(destClusterName)},
		}
	}
	ready := newSyncer("ready", 1)
	longError := newSyncer("long-error", 1)
	notSetUp := newSyncer("not-set-up", 2)

	var mu sync.Mutex
	patches := map[string]int{}
	countPatches := func(name string) int {
		mu.Lock()
		defer mu.Unlock()
		return patches[name]
	}
	c := fake.NewClientBuilder().WithScheme(scheme.Scheme).
		WithObjects(ready, longError, notSetUp).
		WithStatusSubresource(&krmv1alpha1.KRMSyncer{}).
		WithInterceptorFuncs(interceptor.Funcs{
			SubResourcePatch: func(ctx context.Context, c client.Client, sub string, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
				mu.Lock()
				patches[obj.GetName()]++
				mu.Unlock()
				return c.SubResource(sub).Patch(ctx, obj, patch, opts...)
			},
		}).Build()
	const interval = 10 * time.Millisecond
	r := &KRMSyncerReconciler{
		Client:               c,
		Scheme:               scheme.Scheme,
		RemoteConfigProvider: testRemoteProvider,
		setupFailures:        newSetupFailureTracker(prometheus.NewGauge(prometheus.GaugeOpts{Name: "test_setup_failed"})),
		syncFailures:         newSyncFailureTracker(nil),
		statusInterval:       interval,
	}
	r.setupFailures.record(client.ObjectKeyFromObject(ready), 1, nil)
	// A setup error longer than maxConditionMessageLen, with a multi-byte
	// character across the cut.
	longMessage := strings.Repeat("a", maxConditionMessageLen-1) + "é" + strings.Repeat("b", 10)
	r.setupFailures.record(client.ObjectKeyFromObject(longError), 1, stderrors.New(longMessage))
	// Only an older generation has been set up.
	r.setupFailures.record(client.ObjectKeyFromObject(notSetUp), 1, nil)

	// run runs the periodic update for a number of intervals, then stops it.
	run := func(intervals int) {
		t.Helper()
		ctx, cancel := context.WithCancel(t.Context())
		done := make(chan error)
		go func() { done <- r.updateAllStatus(ctx) }()
		time.Sleep(time.Duration(intervals) * interval)
		cancel()
		require.NoError(t, <-done)
	}

	// The first updates write the initial status; later ones, with nothing
	// changed, write nothing (including the truncated long message).
	run(20)
	assert.Equal(t, 1, countPatches("ready"))
	assert.Equal(t, 1, countPatches("long-error"), "a truncated message must compare equal once written")
	assert.Equal(t, 0, countPatches("not-set-up"), "no status until the current generation is set up")
	got := &krmv1alpha1.KRMSyncer{}
	require.NoError(t, c.Get(t.Context(), client.ObjectKeyFromObject(longError), got))
	msg := meta.FindStatusCondition(got.Status.Conditions, krmv1alpha1.ConditionReady).Message
	assert.True(t, utf8.ValidString(msg), "the truncated message must be valid UTF-8")
	assert.Equal(t, strings.Repeat("a", maxConditionMessageLen-1)+"...", msg)

	// 50 objects start failing at once: one status write.
	for i := range 50 {
		obj := objectRef{GVK: schema.GroupVersionKind{Version: "v1", Kind: "Service"}, Key: types.NamespacedName{Namespace: "ns", Name: fmt.Sprintf("svc-%d", i)}}
		r.syncFailures.markFailed(client.ObjectKeyFromObject(ready), obj, "apply failed (Forbidden)")
	}
	run(20)
	assert.Equal(t, 2, countPatches("ready"), "50 failures cost one status write")
	require.NoError(t, c.Get(t.Context(), client.ObjectKeyFromObject(ready), got))
	assert.Equal(t, int32(50), got.Status.FailedObjects)
	assert.Equal(t, 1, countPatches("long-error"))
	assert.Equal(t, 0, countPatches("not-set-up"))

	// A KRMSyncer deleted between the List and its update is not an error.
	assert.NoError(t, r.updateStatus(t.Context(), types.NamespacedName{Namespace: "default", Name: "deleted"}))
}
