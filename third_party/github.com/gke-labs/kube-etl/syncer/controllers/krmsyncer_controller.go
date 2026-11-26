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
	krmv1alpha1 "github.com/gke-labs/kube-etl/syncer/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"slices"
	"strings"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// DefaultMaxConcurrentReconciles is the default number of objects each dynamic
// resource controller reconciles concurrently. It is the default for the
// --max-concurrent-reconciles flag and is used when
// KRMSyncerReconciler.MaxConcurrentReconciles is unset.
const DefaultMaxConcurrentReconciles = 10

// KRMSyncerReconciler reconciles a KRMSyncer object
type KRMSyncerReconciler struct {
	client.Client
	Scheme  *runtime.Scheme
	Manager ctrl.Manager

	// RemoteConfigProvider builds rest.Configs for remote clusters.
	// Defaults to a GKEConfigProvider.
	RemoteConfigProvider RemoteConfigProvider

	// WatchedGVKs tracks which GVKs are already being watched on the local cluster
	WatchedGVKs map[schema.GroupVersionKind]bool
	// WatchedRemoteGVKs tracks which GVKs are already being watched on remote clusters
	WatchedRemoteGVKs map[RemoteGVK]bool
	// RemoteClusters tracks the remote clusters being watched, keyed by RemoteConfigProvider.Key
	RemoteClusters map[string]cluster.Cluster
	// remoteClients caches Push mode destination clients, shared by all
	// push watchers.
	remoteClients *remoteClientCache
	// SyncPeriod determines the minimum frequency at which watched resources are reconciled
	// in remote cluster caches.
	SyncPeriod *time.Duration
	// setupFailures records the latest setup result of each KRMSyncer, reported
	// in status and syncer_setup_failed.
	setupFailures *setupFailureTracker
	// MaxConcurrentReconciles is the maximum number of objects each dynamic
	// resource controller reconciles concurrently.
	// Set from the --max-concurrent-reconciles flag. If unset (<= 0),
	// DefaultMaxConcurrentReconciles is used.
	MaxConcurrentReconciles int
	mu                      sync.RWMutex

	// remotes tracks the reachability of each remote cluster, from the
	// requests of Pull mode caches and Push mode clients, reported in status
	// (Ready=RemoteUnreachable).
	remotes *remoteHealthTracker
	// syncFailures tracks the failed objects of each KRMSyncer, reported in
	// status.failedObjects and syncer_failed_objects.
	syncFailures *syncFailureTracker
	// statusInterval is how often the status of every KRMSyncer is
	// recomputed. Defaults to statusUpdateInterval.
	statusInterval time.Duration
}

type RemoteGVK struct {
	// Cluster is the remote connection key (see RemoteConfigProvider.Key).
	Cluster string
	GVK     schema.GroupVersionKind
}

//+kubebuilder:rbac:groups=syncer.gkelabs.io,resources=krmsyncers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=syncer.gkelabs.io,resources=krmsyncers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=*,resources=*,verbs=get;list;watch;create;update;patch;delete

// Reconcile sets up the watches of a KRMSyncer, records the result in
// setupFailures and updates its status.
func (r *KRMSyncerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the KRMSyncer instance
	var krmSyncer krmv1alpha1.KRMSyncer
	if err := r.Get(ctx, req.NamespacedName, &krmSyncer); err != nil {
		if errors.IsNotFound(err) {
			// A deleted KRMSyncer is no longer failing.
			r.setupFailures.forget(req.NamespacedName)
			r.syncFailures.forget(req.NamespacedName)
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var result ctrl.Result
	var setupErr error
	if krmSyncer.Spec.Suspend {
		// A suspended KRMSyncer is not set up, so it is not failing either.
		r.setupFailures.record(req.NamespacedName, krmSyncer.Generation, nil)
		r.syncFailures.forget(req.NamespacedName)
		// Wait for the source cluster to push resources to us.
		logger.Info("Sync suspended. Waiting for updates from active cluster...")
	} else {
		logger.Info("Syncing updates from active cluster...")
		result, setupErr = r.reconcile(ctx, &krmSyncer)
	}

	// Report the outcome in status.
	if err := r.updateStatus(ctx, req.NamespacedName); err != nil {
		logger.Error(err, "Failed to update KRMSyncer status")
		return result, stderrors.Join(setupErr, err)
	}
	return result, setupErr
}

// syncerMode returns the mode of krmsyncer, defaulting to Pull.
func syncerMode(krmsyncer *krmv1alpha1.KRMSyncer) krmv1alpha1.Mode {
	if krmsyncer.Spec.Mode == "" {
		return krmv1alpha1.ModePull
	}
	return krmsyncer.Spec.Mode
}

func (r *KRMSyncerReconciler) validateRule(rule krmv1alpha1.ResourceRule) error {
	hasGlob := strings.Contains(rule.Group, "*") || strings.Contains(rule.Version, "*") || strings.Contains(rule.Kind, "*")
	if hasGlob {
		isKCC := rule.Group == "*.cnrm.cloud.google.com" ||
			strings.HasSuffix(rule.Group, ".cnrm.cloud.google.com") ||
			rule.Group == "cnrm.cloud.google.com"
		if !isKCC || rule.Version != "*" || rule.Kind != "*" {
			return fmt.Errorf("globbing ('*') is only allowed for version and kind if group is KCC (e.g. *.cnrm.cloud.google.com)")
		}
	}
	return nil
}

func (r *KRMSyncerReconciler) getDiscoveryClient(ctx context.Context, krmsyncer *krmv1alpha1.KRMSyncer, mode krmv1alpha1.Mode) (discovery.DiscoveryInterface, error) {
	if mode == krmv1alpha1.ModePush {
		return discovery.NewDiscoveryClientForConfig(r.Manager.GetConfig())
	}

	// Pull mode
	restConfig, err := r.RemoteConfigProvider.RESTConfig(ctx, krmsyncer.Namespace, krmsyncer.Spec.Remote)
	if err != nil {
		return nil, fmt.Errorf("remote cluster config for Pull mode: %w", err)
	}
	restConfig.Timeout = 10 * time.Second

	return discovery.NewDiscoveryClientForConfig(restConfig)
}

func (r *KRMSyncerReconciler) expandRule(ctx context.Context, krmsyncer *krmv1alpha1.KRMSyncer, rule krmv1alpha1.ResourceRule, mode krmv1alpha1.Mode, resourceLists []*metav1.APIResourceList) ([]schema.GroupVersionKind, error) {
	if err := r.validateRule(rule); err != nil {
		return nil, err
	}

	isKCCGlob := (rule.Group == "*.cnrm.cloud.google.com" ||
		strings.HasSuffix(rule.Group, ".cnrm.cloud.google.com") ||
		rule.Group == "cnrm.cloud.google.com") && rule.Version == "*" && rule.Kind == "*"

	if isKCCGlob {
		var gvks []schema.GroupVersionKind
		for _, rl := range resourceLists {
			gv, err := schema.ParseGroupVersion(rl.GroupVersion)
			if err != nil {
				continue
			}
			// Match group
			groupMatch := false
			if rule.Group == "*.cnrm.cloud.google.com" {
				groupMatch = kccGlobMatchesGroup(gv.Group)
			} else {
				groupMatch = (gv.Group == rule.Group)
			}

			if !groupMatch {
				continue
			}

			for _, res := range rl.APIResources {
				// Avoid subresources
				if strings.Contains(res.Name, "/") {
					continue
				}
				gvks = append(gvks, schema.GroupVersionKind{
					Group:   gv.Group,
					Version: gv.Version,
					Kind:    res.Kind,
				})
			}
		}
		return gvks, nil
	}

	return []schema.GroupVersionKind{{
		Group:   rule.Group,
		Version: rule.Version,
		Kind:    rule.Kind,
	}}, nil
}

func (r *KRMSyncerReconciler) reconcile(ctx context.Context, krmsyncer *krmv1alpha1.KRMSyncer) (_ ctrl.Result, retErr error) {
	logger := log.FromContext(ctx)

	r.mu.Lock()
	defer r.mu.Unlock()

	// Record the result of this setup attempt, reported in status and
	// syncer_setup_failed. A failed setup is requeued with backoff by returning
	// the error.
	key := client.ObjectKeyFromObject(krmsyncer)
	defer func() { r.setupFailures.record(key, krmsyncer.Generation, retErr) }()

	mode := syncerMode(krmsyncer)

	var resourceLists []*metav1.APIResourceList
	needsExpansion := false
	for _, rule := range krmsyncer.Spec.Rules {
		if strings.Contains(rule.Group, "*") || strings.Contains(rule.Version, "*") || strings.Contains(rule.Kind, "*") {
			needsExpansion = true
			break
		}
	}

	if needsExpansion {
		dc, err := r.getDiscoveryClient(ctx, krmsyncer, mode)
		if err != nil {
			return ctrl.Result{}, err
		}
		// Use ServerPreferredResources so each resource is returned once, in its preferred
		// version. Objects are readable through any served version, so watching every
		// served version would sync each object multiple times and race on the destination.
		resourceLists, err = dc.ServerPreferredResources()
		if err != nil {
			if !discovery.IsGroupDiscoveryFailedError(err) {
				return ctrl.Result{}, err
			}
			logger.Info("Discovery failed for some groups, proceeding with partial results", "error", err)
		}
	}

	var setupErrs []error
	for _, rule := range krmsyncer.Spec.Rules {
		gvks, err := r.expandRule(ctx, krmsyncer, rule, mode, resourceLists)
		if err != nil {
			logger.Error(err, "Failed to expand rule", "rule", rule)
			// Reported in the Ready condition (SetupFailed).
			return ctrl.Result{}, fmt.Errorf("invalid rule %s/%s %s: %w", rule.Group, rule.Version, rule.Kind, err)
		}

		for _, gvk := range gvks {
			if mode == krmv1alpha1.ModePull {
				connKey, err := r.RemoteConfigProvider.Key(krmsyncer.Namespace, krmsyncer.Spec.Remote)
				if err != nil {
					logger.Error(err, "Cannot start remote watcher")
					setupErrs = append(setupErrs, err)
					continue
				}
				rgvk := RemoteGVK{
					Cluster: connKey,
					GVK:     gvk,
				}
				if !r.WatchedRemoteGVKs[rgvk] {
					if err := r.startRemoteWatcher(ctx, krmsyncer, rgvk); err != nil {
						logger.Error(err, "Failed to start remote watcher for GVK", "gvk", gvk)
						setupErrs = append(setupErrs, fmt.Errorf("watch %s: %w", gvk, err))
						continue
					}
					r.WatchedRemoteGVKs[rgvk] = true
				}
			} else {
				// Push mode
				if !r.WatchedGVKs[gvk] {
					if err := r.startWatcher(ctx, gvk); err != nil {
						logger.Error(err, "Failed to start watcher for GVK", "gvk", gvk)
						setupErrs = append(setupErrs, fmt.Errorf("watch %s: %w", gvk, err))
						continue
					}
					r.WatchedGVKs[gvk] = true
				}
			}
		}
	}

	// Failed watches are retried with backoff by returning the error. Watches that
	// started are not restarted.
	return ctrl.Result{}, stderrors.Join(setupErrs...)
}

func (r *KRMSyncerReconciler) startWatcher(ctx context.Context, gvk schema.GroupVersionKind) error {
	dr := &DynamicResourceReconciler{
		LocalClient:          r.Client,
		RemoteClient:         r.Client,
		RemoteConfigProvider: r.RemoteConfigProvider,
		remoteClients:        r.remoteClients,
		GVK:                  gvk,
		Mode:                 krmv1alpha1.ModePush,
		syncFailures:         r.syncFailures,
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)

	return ctrl.NewControllerManagedBy(r.Manager).
		For(u).
		Named(fmt.Sprintf("dynamic-watcher-%s-%s-%s", gvk.Group, gvk.Version, gvk.Kind)).
		WithOptions(r.controllerOptions()).
		Complete(dr)
}

// controllerOptions returns the controller options for the dynamic resource
// controllers. DynamicResourceReconciler holds no per-object state, so it is
// safe to reconcile multiple objects concurrently.
func (r *KRMSyncerReconciler) controllerOptions() controller.Options {
	n := r.MaxConcurrentReconciles
	if n <= 0 {
		n = DefaultMaxConcurrentReconciles
	}
	return controller.Options{
		MaxConcurrentReconciles: n,
	}
}

func (r *KRMSyncerReconciler) startRemoteWatcher(ctx context.Context, krmsyncer *krmv1alpha1.KRMSyncer, rgvk RemoteGVK) error {
	remoteCluster, err := r.getOrCreateRemoteCluster(ctx, rgvk.Cluster, krmsyncer.Namespace, krmsyncer.Spec.Remote)
	if err != nil {
		return fmt.Errorf("failed to get remote cluster: %v", err)
	}

	dr := &DynamicResourceReconciler{
		LocalClient:          r.Client,
		RemoteClient:         remoteCluster.GetClient(),
		RemoteConfigProvider: r.RemoteConfigProvider,
		GVK:                  rgvk.GVK,
		Mode:                 krmv1alpha1.ModePull,
		Remote:               rgvk,
		syncFailures:         r.syncFailures,
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(rgvk.GVK)

	gvk := rgvk.GVK
	return ctrl.NewControllerManagedBy(r.Manager).
		Named(fmt.Sprintf("dynamic-puller-%s-%s-%s-%s", strings.ReplaceAll(rgvk.Cluster, "/", "-"), gvk.Group, gvk.Version, gvk.Kind)).
		WatchesRawSource(source.Kind[client.Object](remoteCluster.GetCache(), u, &handler.EnqueueRequestForObject{})).
		WithOptions(r.controllerOptions()).
		Complete(dr)
}

// getOrCreateRemoteCluster returns the cached remote cluster for key, or
// creates and starts one. Requests to the remote update its reachability (see
// remoteHealth).
func (r *KRMSyncerReconciler) getOrCreateRemoteCluster(ctx context.Context, key, namespace string, remote *krmv1alpha1.RemoteConfig) (cluster.Cluster, error) {
	logger := log.FromContext(ctx)

	if c, ok := r.RemoteClusters[key]; ok {
		return c, nil
	}

	restConfig, err := r.RemoteConfigProvider.RESTConfig(ctx, namespace, remote)
	if err != nil {
		return nil, err
	}
	health := r.remotes.get(key)
	restConfig.Wrap(health.wrap)

	c, err := cluster.New(restConfig, func(o *cluster.Options) {
		o.Scheme = r.Scheme
		if r.SyncPeriod != nil && *r.SyncPeriod > 0 {
			o.Cache.SyncPeriod = r.SyncPeriod
		}
	})
	if err != nil {
		return nil, err
	}
	// Map resources refused by the remote (403/5xx) to kinds, for status.
	health.setMapper(c.GetRESTMapper())

	go func() {
		if err := c.Start(ctx); err != nil {
			logger.Error(err, "Failed to start remote cluster")
		}
	}()

	// Wait for cache to sync
	if !c.GetCache().WaitForCacheSync(ctx) {
		return nil, fmt.Errorf("failed to sync remote cluster cache")
	}

	r.RemoteClusters[key] = c
	return c, nil
}

func (r *KRMSyncerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if r.RemoteConfigProvider == nil {
		r.RemoteConfigProvider = &GKEConfigProvider{}
	}
	r.WatchedGVKs = make(map[schema.GroupVersionKind]bool)
	r.WatchedRemoteGVKs = make(map[RemoteGVK]bool)
	r.RemoteClusters = make(map[string]cluster.Cluster)
	r.setupFailures = newSetupFailureTracker(SetupFailed)
	r.remotes = newRemoteHealthTracker()
	r.remoteClients = newRemoteClientCache(r.RemoteConfigProvider, r.remotes)
	r.syncFailures = newSyncFailureTracker(failedObjectsGauge)
	// Report changes of the trackers in status.
	if err := mgr.Add(manager.RunnableFunc(r.updateAllStatus)); err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&krmv1alpha1.KRMSyncer{}, builder.WithPredicates(krmSyncerEventFilter())).
		Complete(r)
}

// krmSyncerEventFilter selects the KRMSyncer events that need setup: creates,
// deletes (to clear the trackers), spec changes (generation changes), and
// periodic resyncs (e.g. to pick up new CRDs for glob rules). It drops
// status-only and metadata-only updates, such as our own status writes.
func krmSyncerEventFilter() predicate.Predicate {
	return predicate.Or(
		predicate.GenerationChangedPredicate{},
		resyncPredicate(),
	)
}

// resyncPredicate passes periodic resync events, which are updates where the
// object did not change (same resourceVersion).
func resyncPredicate() predicate.Predicate {
	return predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			return e.ObjectOld.GetResourceVersion() == e.ObjectNew.GetResourceVersion()
		},
	}
}

// DynamicResourceReconciler handles events for dynamically watched resources
type DynamicResourceReconciler struct {
	LocalClient  client.Client
	RemoteClient client.Client
	// RemoteConfigProvider builds the destination client config (Push mode) and
	// derives connection keys to match KRMSyncers to this watcher (Pull mode).
	RemoteConfigProvider RemoteConfigProvider
	GVK                  schema.GroupVersionKind
	Mode                 krmv1alpha1.Mode
	Remote               RemoteGVK // Only used for Pull mode
	// remoteClients caches destination clients (Push mode). If nil, a new
	// client is built for every event.
	remoteClients *remoteClientCache
	// syncFailures tracks, per KRMSyncer, objects whose latest sync attempt
	// failed, reported in status.failedObjects and syncer_failed_objects.
	// May be nil.
	syncFailures *syncFailureTracker
}

func (r *DynamicResourceReconciler) ruleMatchesGVK(rule krmv1alpha1.ResourceRule, gvk schema.GroupVersionKind) bool {
	return ruleMatchesGVK(rule, gvk)
}

// ruleMatchesGVK reports whether rule syncs gvk.
func ruleMatchesGVK(rule krmv1alpha1.ResourceRule, gvk schema.GroupVersionKind) bool {
	isKCCGlob := (rule.Group == "*.cnrm.cloud.google.com" ||
		strings.HasSuffix(rule.Group, ".cnrm.cloud.google.com") ||
		rule.Group == "cnrm.cloud.google.com") && rule.Version == "*" && rule.Kind == "*"

	if isKCCGlob {
		if rule.Group == "*.cnrm.cloud.google.com" {
			return kccGlobMatchesGroup(gvk.Group)
		}
		return gvk.Group == rule.Group
	}
	return rule.Group == gvk.Group && rule.Version == gvk.Version && rule.Kind == gvk.Kind
}

func (r *DynamicResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Fetch all Syncers to find matching rules (Active ones)
	var krmsyncers krmv1alpha1.KRMSyncerList
	if err := r.LocalClient.List(ctx, &krmsyncers); err != nil {
		return ctrl.Result{}, err
	}

	// Fetch the resource from Source cluster
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.GVK)
	readErr := r.RemoteClient.Get(ctx, req.NamespacedName, u)
	isDeleted := errors.IsNotFound(readErr)
	if isDeleted {
		readErr = nil
	} else if readErr != nil {
		readErr = newSyncError("source read failed", readErr)
	}

	obj := objectRef{GVK: r.GVK, Key: req.NamespacedName}
	var syncErrors []error
	for i := range krmsyncers.Items {
		krmsyncer := &krmsyncers.Items[i]
		syncerKey := client.ObjectKeyFromObject(krmsyncer)
		rules := r.matchingRules(krmsyncer, req.Namespace)
		if len(rules) == 0 {
			// The object is not synced by this KRMSyncer (anymore).
			r.syncFailures.markSucceeded(syncerKey, obj)
			continue
		}

		// The object cannot be synced if it cannot be read from the source.
		errs := []error{readErr}
		if readErr == nil {
			errs = errs[:0]
			rule := rules[0]
			rule.SyncFields = mergeSyncFields(rules)
			if err := r.syncRule(ctx, krmsyncer, rule, req, u, isDeleted); err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			r.syncFailures.markFailed(syncerKey, obj, syncErrorCategory(errs[0]))
			syncErrors = append(syncErrors, errs...)
		} else {
			r.syncFailures.markSucceeded(syncerKey, obj)
		}
	}
	if readErr != nil {
		return ctrl.Result{}, readErr
	}
	return ctrl.Result{}, stderrors.Join(syncErrors...)
}

// matchingRules returns the rules of krmsyncer that sync this controller's
// GVK in namespace, or none if krmsyncer is suspended or uses another mode or
// remote.
func (r *DynamicResourceReconciler) matchingRules(krmsyncer *krmv1alpha1.KRMSyncer, namespace string) []krmv1alpha1.ResourceRule {
	if krmsyncer.Spec.Suspend || syncerMode(krmsyncer) != r.Mode {
		return nil
	}
	// For Pull mode, also check Remote match
	if r.Mode == krmv1alpha1.ModePull {
		connKey, err := r.RemoteConfigProvider.Key(krmsyncer.Namespace, krmsyncer.Spec.Remote)
		if err != nil || connKey != r.Remote.Cluster {
			return nil
		}
	}
	var rules []krmv1alpha1.ResourceRule
	for _, rule := range krmsyncer.Spec.Rules {
		if !r.ruleMatchesGVK(rule, r.GVK) {
			continue
		}
		if len(rule.Namespaces) > 0 && !slices.Contains(rule.Namespaces, namespace) {
			continue
		}
		rules = append(rules, rule)
	}
	return rules
}

// mergeSyncFields unions SyncFields across all matching rules, preserving order and
// deduplicating fields. If a rule has empty SyncFields, it defaults to ["status"].
func mergeSyncFields(rules []krmv1alpha1.ResourceRule) []string {
	var merged []string
	seen := make(map[string]bool)
	for _, rule := range rules {
		fields := rule.SyncFields
		if len(fields) == 0 {
			fields = []string{"status"}
		}
		for _, f := range fields {
			if !seen[f] {
				seen[f] = true
				merged = append(merged, f)
			}
		}
	}
	return merged
}

// syncRule syncs the source object u (or its deletion) to the destination of
// krmsyncer according to rule.
func (r *DynamicResourceReconciler) syncRule(ctx context.Context, krmsyncer *krmv1alpha1.KRMSyncer, rule krmv1alpha1.ResourceRule, req ctrl.Request, u *unstructured.Unstructured, isDeleted bool) error {
	logger := log.FromContext(ctx)

	// Determine Destination Client
	var destClient client.Client
	if r.Mode == krmv1alpha1.ModePush {
		var err error
		destClient, err = r.getRemoteClient(ctx, krmsyncer)
		if err != nil {
			logger.Error(err, "Failed to get remote client")
			return newSyncError("remote client failed", err)
		}
	} else {
		destClient = r.LocalClient
	}

	// Handle Deletion
	if isDeleted {
		toDelete := &unstructured.Unstructured{}
		toDelete.SetGroupVersionKind(r.GVK)
		toDelete.SetName(req.Name)
		toDelete.SetNamespace(req.Namespace)

		logger.Info("Deleting resource on destination cluster", "name", req.Name, "namespace", req.Namespace)
		if err := destClient.Delete(ctx, toDelete); err != nil && !errors.IsNotFound(err) {
			logger.Error(err, "Failed to delete resource on destination cluster")
			return newSyncError("delete failed", err)
		}
		return nil
	}

	// Sync to Destination Cluster
	logger.Info("Syncing resource to destination cluster", "name", u.GetName(), "namespace", u.GetNamespace())

	syncFields := rule.SyncFields
	destObj, err := r.filterFields(u, syncFields)
	if err != nil {
		logger.Error(err, "Failed to filter fields")
		return newSyncError("invalid sync fields", err)
	}

	// Prepare object for destination: clear ResourceVersion, UID, etc.
	destObj.SetResourceVersion("")
	destObj.SetUID("")
	destObj.SetGeneration(0)
	destObj.SetManagedFields(nil)

	// In status mode the syncer only updates objects that already exist on
	// the destination; it must not create a stub (or its namespace). A
	// missing destination object is an error, retried until it is created.
	if !isFullSync(syncFields) {
		exists, err := destinationExists(ctx, destClient, destObj)
		if err != nil {
			logger.Error(err, "Failed to check destination object")
			return newSyncError("destination check failed", err)
		}
		if !exists {
			name := destObj.GetName()
			if ns := destObj.GetNamespace(); ns != "" {
				name = ns + "/" + name
			}
			err := fmt.Errorf("destination object %s %s does not exist (KRMSyncer %s/%s syncs status only and does not create objects)",
				r.GVK.Kind, name, krmsyncer.Namespace, krmsyncer.Name)
			logger.Error(err, "Destination object does not exist")
			return newSyncError("destination object missing", err)
		}
	}

	if err := r.applyToDestination(ctx, destClient, destObj); err != nil {
		logger.Error(err, "Failed to apply to destination cluster")
		return err
	}
	logger.Info("Successfully synced resource to destination cluster", "name", u.GetName(), "namespace", u.GetNamespace())
	return nil
}

// syncError is an object sync failure with a short category, e.g.
// "apply failed (Forbidden)". Failures are grouped by category in the Ready
// condition message; the full error is logged where it happens.
type syncError struct {
	category string
	err      error
}

func (e *syncError) Error() string { return e.err.Error() }
func (e *syncError) Unwrap() error { return e.err }

// newSyncError returns err with the category. For Kubernetes API errors, the
// category includes the reason, e.g. "apply failed (Forbidden)".
func newSyncError(category string, err error) error {
	if reason := errors.ReasonForError(err); reason != metav1.StatusReasonUnknown {
		category = fmt.Sprintf("%s (%s)", category, reason)
	}
	return &syncError{category: category, err: err}
}

// syncErrorCategory returns the category of a sync error.
func syncErrorCategory(err error) string {
	var se *syncError
	if stderrors.As(err, &se) {
		return se.category
	}
	return "sync failed"
}

// destinationExists reports whether obj exists on the destination cluster.
func destinationExists(ctx context.Context, destClient client.Client, obj *unstructured.Unstructured) (bool, error) {
	existing := &unstructured.Unstructured{}
	existing.SetGroupVersionKind(obj.GroupVersionKind())
	err := destClient.Get(ctx, client.ObjectKeyFromObject(obj), existing)
	if err == nil {
		return true, nil
	}
	if errors.IsNotFound(err) {
		return false, nil
	}
	return false, fmt.Errorf("failed to check destination object: %w", err)
}

func (r *DynamicResourceReconciler) getRemoteClient(ctx context.Context, krmsyncer *krmv1alpha1.KRMSyncer) (client.Client, error) {
	if r.remoteClients != nil {
		return r.remoteClients.get(ctx, krmsyncer.Namespace, krmsyncer.Spec.Remote)
	}

	restConfig, err := r.RemoteConfigProvider.RESTConfig(ctx, krmsyncer.Namespace, krmsyncer.Spec.Remote)
	if err != nil {
		return nil, err
	}

	return client.New(restConfig, client.Options{})
}

func (r *DynamicResourceReconciler) filterFields(src *unstructured.Unstructured, fields []string) (*unstructured.Unstructured, error) {
	dest := &unstructured.Unstructured{
		Object: make(map[string]any),
	}
	dest.SetGroupVersionKind(src.GroupVersionKind())
	dest.SetName(src.GetName())
	dest.SetNamespace(src.GetNamespace())
	// Labels and annotations are only synced in full mode (when "spec" is synced).
	// In status mode, the destination object's metadata is left untouched.
	if isFullSync(fields) {
		dest.SetLabels(src.GetLabels())
		dest.SetAnnotations(src.GetAnnotations())
	}

	for _, field := range fields {
		parts := strings.Split(field, ".")
		val, found, err := unstructured.NestedFieldCopy(src.Object, parts...)
		if err != nil {
			return nil, fmt.Errorf("failed to get field %s: %v", field, err)
		}
		if found {
			if err := unstructured.SetNestedField(dest.Object, val, parts...); err != nil {
				return nil, fmt.Errorf("failed to set field %s: %v", field, err)
			}
		}
	}
	return dest, nil
}

// isFullSync reports whether the sync fields include the full "spec". Otherwise the
// syncer is in status mode (e.g. "status", optionally with "spec.resourceID").
func isFullSync(fields []string) bool {
	for _, field := range fields {
		if field == "spec" {
			return true
		}
	}
	return false
}

func (r *DynamicResourceReconciler) applyToDestination(ctx context.Context, destClient client.Client, obj *unstructured.Unstructured) error {
	// Ensure destination namespace exists before applying the resource.
	nsName := obj.GetNamespace()
	if nsName != "" {
		ns := &corev1.Namespace{}
		err := destClient.Get(ctx, client.ObjectKey{Name: nsName}, ns)
		if err != nil {
			if errors.IsNotFound(err) {
				ns = &corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: nsName,
					},
				}
				// Ignore AlreadyExists: concurrent workers may race to create the same namespace.
				if err := destClient.Create(ctx, ns); err != nil && !errors.IsAlreadyExists(err) {
					return newSyncError("namespace create failed", fmt.Errorf("failed to create destination namespace %q: %w", nsName, err))
				}
			} else {
				return newSyncError("destination check failed", fmt.Errorf("failed to check destination namespace %q: %w", nsName, err))
			}
		}
	}
	patchOpts := []client.PatchOption{
		client.ForceOwnership,
		client.FieldOwner("krm-syncer"),
	}

	status, found, err := unstructured.NestedFieldCopy(obj.Object, "status")
	if err != nil {
		return newSyncError("invalid sync fields", fmt.Errorf("failed to copy status: %w", err))
	}

	// Use Server-Side Apply (SSA) to create or update the resource.
	if err := destClient.Patch(ctx, obj, client.Apply, patchOpts...); err != nil {
		return newSyncError("apply failed", fmt.Errorf("failed to apply object to destination: %w", err))
	}

	// If status is present, we also need to patch the status subresource.
	if found {
		statusObj := &unstructured.Unstructured{
			Object: map[string]any{
				"apiVersion": obj.GetAPIVersion(),
				"kind":       obj.GetKind(),
				"metadata": map[string]any{
					"name":      obj.GetName(),
					"namespace": obj.GetNamespace(),
				},
				"status": status,
			},
		}
		subResourcePatchOpts := []client.SubResourcePatchOption{
			client.ForceOwnership,
			client.FieldOwner("krm-syncer"),
		}
		if err := destClient.Status().Patch(ctx, statusObj, client.Apply, subResourcePatchOpts...); err != nil {
			return newSyncError("status apply failed", fmt.Errorf("failed to apply status subresource to destination: %w", err))
		}
	}
	return nil
}
