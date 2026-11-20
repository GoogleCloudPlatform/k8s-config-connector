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
	"fmt"
	"maps"
	"slices"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// updateStatus computes the status of the KRMSyncer key from its latest setup
// result and the trackers, and writes it if it changed. It is called at the end
// of Reconcile, and for every KRMSyncer every statusUpdateInterval (see
// updateAllStatus). Nothing is written until the current generation has been
// set up: Reconcile updates status when setup completes.
//
// The KRMSyncer is read from the cache. If an update is based on a stale copy
// or an older state of the trackers, the next periodic update corrects it.
func (r *KRMSyncerReconciler) updateStatus(ctx context.Context, key types.NamespacedName) error {
	var krmsyncer krmv1alpha1.KRMSyncer
	if err := r.Get(ctx, key, &krmsyncer); err != nil {
		return client.IgnoreNotFound(err)
	}
	setup, ok := r.setupFailures.result(key)
	if !ok || setup.generation != krmsyncer.Generation {
		return nil
	}
	orig := krmsyncer.DeepCopy()
	r.setReadyCondition(&krmsyncer, setup.err)
	return r.writeStatus(ctx, orig, &krmsyncer)
}

// statusUpdateInterval is how often the status of every KRMSyncer is
// recomputed, to report changes of the trackers (failed objects, remote
// reachability).
const statusUpdateInterval = 2 * time.Second

// updateAllStatus recomputes the status of every KRMSyncer every
// statusUpdateInterval until ctx is done, writing it only when it changed. It
// implements manager.RunnableFunc.
//
// Each KRMSyncer costs a cache read and a comparison per interval; the API
// server is only called when its status changes, at most once per interval
// however many objects start or stop failing.
func (r *KRMSyncerReconciler) updateAllStatus(ctx context.Context) error {
	interval := r.statusInterval
	if interval <= 0 {
		interval = statusUpdateInterval
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
		var list krmv1alpha1.KRMSyncerList
		if err := r.List(ctx, &list); err != nil {
			log.FromContext(ctx).Error(err, "Failed to list KRMSyncers")
			continue
		}
		for i := range list.Items {
			key := client.ObjectKeyFromObject(&list.Items[i])
			if err := r.updateStatus(ctx, key); err != nil {
				log.FromContext(ctx).Error(err, "Failed to update KRMSyncer status", "krmsyncer", key)
			}
		}
	}
}

// legacyConditionTypes are condition types written by earlier versions, which
// are replaced by Ready.
var legacyConditionTypes = []string{"Active", "Suspended", "InvalidRule"}

// maxConditionMessageLen bounds the Ready condition message, which may hold
// several joined setup errors.
const maxConditionMessageLen = 4096

// Limits of the failures listed in the Ready condition message: errors, and
// kinds per error. The rest are counted ("and 2 more kinds").
const (
	maxErrorsInMessage        = 5
	maxKindsPerErrorInMessage = 3
)

// syncFailure is an error category (e.g. "apply failed (Forbidden)") and a kind
// it affects (e.g. "Service", "TestCRD.e2e.gkelabs.io").
type syncFailure struct {
	category string
	kind     string
}

// setReadyCondition sets status.failedObjects, observedGeneration and the
// Ready condition of krmsyncer from the result of its latest setup (setupErr),
// the reachability of its remote cluster and its failed objects. It is set on
// every status update, so it never goes stale.
func (r *KRMSyncerReconciler) setReadyCondition(krmsyncer *krmv1alpha1.KRMSyncer, setupErr error) {
	cond := metav1.Condition{
		Type:               krmv1alpha1.ConditionReady,
		ObservedGeneration: krmsyncer.Generation,
	}
	remoteErr := r.remoteUnreachable(krmsyncer)
	failed, failures := r.syncFailures.failures(client.ObjectKeyFromObject(krmsyncer))
	krmsyncer.Status.FailedObjects = int32(failed)
	// Kinds the remote refuses to serve (403/5xx) are reported with the
	// failed objects.
	failures = append(failures, r.remoteKindFailures(krmsyncer)...)
	switch {
	case krmsyncer.Spec.Suspend:
		cond.Status = metav1.ConditionTrue
		cond.Reason = krmv1alpha1.ReasonSuspended
		cond.Message = "Sync is suspended by spec.suspend"
	case remoteErr != nil:
		cond.Status = metav1.ConditionFalse
		cond.Reason = krmv1alpha1.ReasonRemoteUnreachable
		cond.Message = remoteErr.Error()
	case setupErr != nil:
		cond.Status = metav1.ConditionFalse
		cond.Reason = krmv1alpha1.ReasonSetupFailed
		cond.Message = setupErr.Error()
	case len(failures) > 0:
		cond.Status = metav1.ConditionFalse
		cond.Reason = krmv1alpha1.ReasonSyncFailed
		cond.Message = failuresMessage(failures)
	default:
		cond.Status = metav1.ConditionTrue
		cond.Reason = krmv1alpha1.ReasonSyncing
		cond.Message = "Syncing resources"
	}
	cond.Message = truncateMessage(cond.Message, maxConditionMessageLen)
	for _, t := range legacyConditionTypes {
		meta.RemoveStatusCondition(&krmsyncer.Status.Conditions, t)
	}
	meta.SetStatusCondition(&krmsyncer.Status.Conditions, cond)
	krmsyncer.Status.ObservedGeneration = krmsyncer.Generation
}

// truncateMessage shortens msg to at most max bytes (plus "..."), cutting on a
// character boundary. Cutting inside a multi-byte character would produce
// invalid UTF-8, which the API server stores differently, so the status would
// never compare equal and would be rewritten on every update.
func truncateMessage(msg string, max int) string {
	if len(msg) <= max {
		return msg
	}
	cut := max
	for cut > 0 && !utf8.RuneStart(msg[cut]) {
		cut--
	}
	return msg[:cut] + "..."
}

// failuresMessage describes failures grouped by error, each with the kinds it
// affects, e.g. "apply failed (Forbidden): Service; destination object missing:
// TestCRD.e2e.gkelabs.io, TestCRD2.e2e.gkelabs.io, TestCRD3.e2e.gkelabs.io, and
// 5 more kinds. See the controller logs for details." Errors and kinds are
// sorted, so the message only changes when the failures change. The number of
// failed objects is in status.failedObjects, and the errors in the controller
// logs.
func failuresMessage(failures []syncFailure) string {
	kindsByCategory := make(map[string][]string)
	for _, f := range failures {
		if !slices.Contains(kindsByCategory[f.category], f.kind) {
			kindsByCategory[f.category] = append(kindsByCategory[f.category], f.kind)
		}
	}
	categories := slices.Sorted(maps.Keys(kindsByCategory))

	var groups []string
	for i, category := range categories {
		if i == maxErrorsInMessage {
			groups = append(groups, "and "+countNoun(len(categories)-i, "more error", "more errors"))
			break
		}
		kinds := kindsByCategory[category]
		slices.Sort(kinds)
		if len(kinds) > maxKindsPerErrorInMessage {
			more := len(kinds) - maxKindsPerErrorInMessage
			kinds = append(kinds[:maxKindsPerErrorInMessage:maxKindsPerErrorInMessage], "and "+countNoun(more, "more kind", "more kinds"))
		}
		groups = append(groups, category+": "+strings.Join(kinds, ", "))
	}
	return strings.Join(groups, "; ") + ". See the controller logs for details."
}

// countNoun returns n with the singular or plural noun, e.g. "2 more kinds".
func countNoun(n int, singular, plural string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, singular)
	}
	return fmt.Sprintf("%d %s", n, plural)
}

// writeStatus writes the status of krmsyncer if it differs from orig's.
func (r *KRMSyncerReconciler) writeStatus(ctx context.Context, orig, krmsyncer *krmv1alpha1.KRMSyncer) error {
	if equality.Semantic.DeepEqual(orig.Status, krmsyncer.Status) {
		return nil
	}
	return client.IgnoreNotFound(r.Status().Patch(ctx, krmsyncer, client.MergeFrom(orig)))
}

// remoteUnreachable returns why the remote cluster of a KRMSyncer is
// unreachable, or nil if it is reachable (or not tracked). In Pull mode the
// remote is watched continuously; in Push mode it is only used, so observed,
// while syncing.
func (r *KRMSyncerReconciler) remoteUnreachable(krmsyncer *krmv1alpha1.KRMSyncer) error {
	key, err := r.RemoteConfigProvider.Key(krmsyncer.Namespace, krmsyncer.Spec.Remote)
	if err != nil {
		return nil
	}
	err = r.remotes.unreachable(key)
	if err == nil {
		return nil
	}
	name, nameErr := gkeClusterName(krmsyncer.Spec.Remote)
	if nameErr != nil {
		name = key
	}
	return fmt.Errorf("remote cluster %s: %w", name, err)
}

// remoteKindFailures returns, for a Pull mode KRMSyncer, the kinds it syncs
// that its remote refuses to serve (403 or 5xx on list/watch), e.g.
// "watch failed (Forbidden)" for TestCRD.e2e.gkelabs.io.
func (r *KRMSyncerReconciler) remoteKindFailures(krmsyncer *krmv1alpha1.KRMSyncer) []syncFailure {
	if syncerMode(krmsyncer) != krmv1alpha1.ModePull {
		return nil
	}
	key, err := r.RemoteConfigProvider.Key(krmsyncer.Namespace, krmsyncer.Spec.Remote)
	if err != nil {
		return nil
	}
	h := r.remotes.lookup(key)
	if h == nil {
		return nil
	}
	var failures []syncFailure
	for gvk, reason := range h.failedKinds() {
		if slices.ContainsFunc(krmsyncer.Spec.Rules, func(rule krmv1alpha1.ResourceRule) bool { return ruleMatchesGVK(rule, gvk) }) {
			failures = append(failures, syncFailure{category: fmt.Sprintf("watch failed (%s)", reason), kind: gvk.GroupKind().String()})
		}
	}
	return failures
}

// The trackers below hold the state that KRMSyncer status and metrics report:
//   - setupFailureTracker: result of the latest setup of each KRMSyncer
//     (Ready=SetupFailed, observedGeneration, syncer_setup_failed);
//   - syncFailureTracker: objects whose latest sync failed, per KRMSyncer
//     (status.failedObjects, Ready=SyncFailed, syncer_failed_objects);
//   - remoteHealthTracker: reachability of each remote cluster
//     (Ready=RemoteUnreachable).

// setupFailureTracker records the result of the latest setup attempt of each
// KRMSyncer: the generation it set up and its error, if any. updateStatus
// reports them in Ready (SetupFailed) and observedGeneration. It
// keeps the syncer_setup_failed gauge equal to the number of KRMSyncers whose
// latest setup failed, so repeated failures count once.
//
// A nil *setupFailureTracker ignores all updates.
type setupFailureTracker struct {
	gauge prometheus.Gauge

	mu      sync.Mutex
	results map[types.NamespacedName]setupResult
}

// setupResult is the result of a setup attempt of a KRMSyncer.
type setupResult struct {
	// generation is the KRMSyncer generation that was set up.
	generation int64
	// err is why setup failed, or nil.
	err error
}

func newSetupFailureTracker(gauge prometheus.Gauge) *setupFailureTracker {
	return &setupFailureTracker{gauge: gauge, results: make(map[types.NamespacedName]setupResult)}
}

// record records the result of the latest setup attempt of key.
func (t *setupFailureTracker) record(key types.NamespacedName, generation int64, err error) {
	if t == nil {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	wasFailed := t.results[key].err != nil
	t.results[key] = setupResult{generation: generation, err: err}
	switch {
	case !wasFailed && err != nil:
		t.gauge.Inc()
	case wasFailed && err == nil:
		t.gauge.Dec()
	}
}

// forget drops the result of key (e.g. it was deleted).
func (t *setupFailureTracker) forget(key types.NamespacedName) {
	if t == nil {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.results[key].err != nil {
		t.gauge.Dec()
	}
	delete(t.results, key)
}

// result returns the result of the latest setup attempt of key, if any.
func (t *setupFailureTracker) result(key types.NamespacedName) (setupResult, bool) {
	if t == nil {
		return setupResult{}, false
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	res, ok := t.results[key]
	return res, ok
}

// objectRef identifies a synced object.
type objectRef struct {
	GVK schema.GroupVersionKind
	Key types.NamespacedName
}

// syncFailureTracker tracks, per KRMSyncer, which objects failed their latest
// sync attempt. It is shared by all dynamic resource controllers and feeds:
//   - status.failedObjects and the Ready message of each KRMSyncer;
//   - the syncer_failed_objects metric (via gauge): an object is counted
//     once under its GVK while it is failing for at least one KRMSyncer.
//
// An object is counted from its first failed attempt until it syncs
// successfully, however many times it is retried in between.
//
// A nil *syncFailureTracker ignores all updates.
type syncFailureTracker struct {
	// gauge returns the failed object gauge of a GVK. May be nil.
	gauge func(schema.GroupVersionKind) prometheus.Gauge

	mu sync.Mutex
	// failed maps each KRMSyncer's failing objects to their error category.
	failed map[types.NamespacedName]map[objectRef]string
	// syncers is the number of KRMSyncers each failed object is failing for.
	syncers map[objectRef]int
}

func newSyncFailureTracker(gauge func(schema.GroupVersionKind) prometheus.Gauge) *syncFailureTracker {
	return &syncFailureTracker{
		gauge:   gauge,
		failed:  make(map[types.NamespacedName]map[objectRef]string),
		syncers: make(map[objectRef]int),
	}
}

// markFailed records that the latest attempt to sync obj for syncer failed,
// with the error category (see syncErrorCategory).
func (t *syncFailureTracker) markFailed(syncer types.NamespacedName, obj objectRef, category string) {
	if t == nil {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	objs, ok := t.failed[syncer]
	if !ok {
		objs = make(map[objectRef]string)
		t.failed[syncer] = objs
	}
	if _, already := objs[obj]; !already {
		t.addLocked(obj, 1)
	}
	objs[obj] = category
}

// markSucceeded records that obj synced for syncer, or no longer needs syncing
// for it (e.g. it was deleted, or no rule of syncer matches it anymore).
func (t *syncFailureTracker) markSucceeded(syncer types.NamespacedName, obj objectRef) {
	if t == nil {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	objs := t.failed[syncer]
	if _, found := objs[obj]; found {
		delete(objs, obj)
		if len(objs) == 0 {
			delete(t.failed, syncer)
		}
		t.addLocked(obj, -1)
	}
}

// forget drops the failures of syncer (e.g. it was deleted or suspended).
func (t *syncFailureTracker) forget(syncer types.NamespacedName) {
	if t == nil {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	for obj := range t.failed[syncer] {
		t.addLocked(obj, -1)
	}
	delete(t.failed, syncer)
}

// addLocked adds delta to the number of KRMSyncers obj is failing for, and
// updates the gauge when obj starts or stops failing.
func (t *syncFailureTracker) addLocked(obj objectRef, delta int) {
	before := t.syncers[obj]
	after := before + delta
	if after <= 0 {
		delete(t.syncers, obj)
	} else {
		t.syncers[obj] = after
	}
	if t.gauge == nil {
		return
	}
	switch {
	case before == 0 && after > 0:
		t.gauge(obj.GVK).Inc()
	case before > 0 && after <= 0:
		t.gauge(obj.GVK).Dec()
	}
}

// failures returns the number of objects failing for syncer, and their
// distinct (error category, kind) pairs.
func (t *syncFailureTracker) failures(syncer types.NamespacedName) (int, []syncFailure) {
	if t == nil {
		return 0, nil
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	seen := make(map[syncFailure]bool)
	var failures []syncFailure
	for obj, category := range t.failed[syncer] {
		f := syncFailure{category: category, kind: obj.GVK.GroupKind().String()}
		if !seen[f] {
			seen[f] = true
			failures = append(failures, f)
		}
	}
	return len(t.failed[syncer]), failures
}

// remoteHealthTracker tracks the reachability of each remote cluster, keyed by
// RemoteConfigProvider.Key. Each remote's remoteHealth is updated from the
// responses to its requests (see healthTransport).
//
// A nil *remoteHealthTracker tracks nothing: get returns a detached
// remoteHealth and unreachable always returns nil.
type remoteHealthTracker struct {
	mu      sync.Mutex
	remotes map[string]*remoteHealth
}

func newRemoteHealthTracker() *remoteHealthTracker {
	return &remoteHealthTracker{remotes: make(map[string]*remoteHealth)}
}

// get returns the health of the remote key, creating it on first use.
func (t *remoteHealthTracker) get(key string) *remoteHealth {
	if t == nil {
		return &remoteHealth{}
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if h, ok := t.remotes[key]; ok {
		return h
	}
	h := &remoteHealth{}
	t.remotes[key] = h
	return h
}

// lookup returns the health of the remote key, or nil if no connection to it
// was made.
func (t *remoteHealthTracker) lookup(key string) *remoteHealth {
	if t == nil {
		return nil
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.remotes[key]
}

// unreachable returns why the remote key is unreachable, or nil if it is
// reachable or no connection to it was made.
func (t *remoteHealthTracker) unreachable(key string) error {
	if t == nil {
		return nil
	}
	t.mu.Lock()
	h := t.remotes[key]
	t.mu.Unlock()
	if h == nil {
		return nil
	}
	return h.unreachable()
}
