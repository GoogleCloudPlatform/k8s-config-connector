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
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

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
	mu         sync.RWMutex
}

type RemoteGVK struct {
	// Cluster is the remote connection key (see RemoteConfigProvider.Key).
	Cluster string
	GVK     schema.GroupVersionKind
}

//+kubebuilder:rbac:groups=syncer.gkelabs.io,resources=krmsyncers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=syncer.gkelabs.io,resources=krmsyncers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=*,resources=*,verbs=get;list;watch;create;update;patch;delete

func (r *KRMSyncerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the KRMSyncer instance
	var krmSyncer krmv1alpha1.KRMSyncer
	if err := r.Get(ctx, req.NamespacedName, &krmSyncer); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Update Status if needed
	defer func() {
		if err := r.Status().Update(ctx, &krmSyncer); err != nil {
			logger.Error(err, "Failed to update KRMSyncer status")
		}
	}()

	if krmSyncer.Spec.Suspend {
		meta.SetStatusCondition(&krmSyncer.Status.Conditions, metav1.Condition{
			Type:    "Suspended",
			Status:  metav1.ConditionTrue,
			Reason:  "SuspendedBySpec",
			Message: "Controller is suspended",
		})
		// Wait for the source cluster to push resources to us.
		logger.Info("Sync suspended. Waiting for updates from active cluster...")
		return ctrl.Result{}, nil
	}

	meta.SetStatusCondition(&krmSyncer.Status.Conditions, metav1.Condition{
		Type:    "Active",
		Status:  metav1.ConditionTrue,
		Reason:  "Active",
		Message: "Controller is active",
	})
	logger.Info("Syncing updates from active cluster...")
	return r.reconcile(ctx, &krmSyncer)
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
				groupMatch = strings.HasSuffix(gv.Group, "cnrm.cloud.google.com")
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

func (r *KRMSyncerReconciler) reconcile(ctx context.Context, krmsyncer *krmv1alpha1.KRMSyncer) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	r.mu.Lock()
	defer r.mu.Unlock()

	mode := krmsyncer.Spec.Mode
	if mode == "" {
		mode = krmv1alpha1.ModePull
	}

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
		// Use ServerGroupsAndResources to get all versions
		_, resourceLists, err = dc.ServerGroupsAndResources()
		if err != nil {
			if !discovery.IsGroupDiscoveryFailedError(err) {
				return ctrl.Result{}, err
			}
			logger.Info("Discovery failed for some groups, proceeding with partial results", "error", err)
		}
	}

	for _, rule := range krmsyncer.Spec.Rules {
		gvks, err := r.expandRule(ctx, krmsyncer, rule, mode, resourceLists)
		if err != nil {
			logger.Error(err, "Failed to expand rule", "rule", rule)
			meta.SetStatusCondition(&krmsyncer.Status.Conditions, metav1.Condition{
				Type:    "InvalidRule",
				Status:  metav1.ConditionTrue,
				Reason:  "InvalidGlob",
				Message: err.Error(),
			})
			return ctrl.Result{}, err
		}

		for _, gvk := range gvks {
			if mode == krmv1alpha1.ModePull {
				connKey, err := r.RemoteConfigProvider.Key(krmsyncer.Namespace, krmsyncer.Spec.Remote)
				if err != nil {
					logger.Error(err, "Cannot start remote watcher")
					continue
				}
				rgvk := RemoteGVK{
					Cluster: connKey,
					GVK:     gvk,
				}
				if !r.WatchedRemoteGVKs[rgvk] {
					if err := r.startRemoteWatcher(ctx, krmsyncer, rgvk); err != nil {
						logger.Error(err, "Failed to start remote watcher for GVK", "gvk", gvk)
						continue
					}
					r.WatchedRemoteGVKs[rgvk] = true
				}
			} else {
				// Push mode
				if !r.WatchedGVKs[gvk] {
					if err := r.startWatcher(ctx, gvk); err != nil {
						logger.Error(err, "Failed to start watcher for GVK", "gvk", gvk)
						continue
					}
					r.WatchedGVKs[gvk] = true
				}
			}
		}
	}

	return ctrl.Result{}, nil
}

func (r *KRMSyncerReconciler) startWatcher(ctx context.Context, gvk schema.GroupVersionKind) error {
	dr := &DynamicResourceReconciler{
		LocalClient:          r.Client,
		RemoteClient:         r.Client,
		RemoteConfigProvider: r.RemoteConfigProvider,
		remoteClients:        r.remoteClients,
		GVK:                  gvk,
		Mode:                 krmv1alpha1.ModePush,
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)

	return ctrl.NewControllerManagedBy(r.Manager).
		For(u).
		Named(fmt.Sprintf("dynamic-watcher-%s-%s-%s", gvk.Group, gvk.Version, gvk.Kind)).
		Complete(dr)
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
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(rgvk.GVK)

	gvk := rgvk.GVK
	return ctrl.NewControllerManagedBy(r.Manager).
		Named(fmt.Sprintf("dynamic-puller-%s-%s-%s-%s", strings.ReplaceAll(rgvk.Cluster, "/", "-"), gvk.Group, gvk.Version, gvk.Kind)).
		WatchesRawSource(source.Kind[client.Object](remoteCluster.GetCache(), u, &handler.EnqueueRequestForObject{})).
		Complete(dr)
}

func (r *KRMSyncerReconciler) getOrCreateRemoteCluster(ctx context.Context, key, namespace string, remote *krmv1alpha1.RemoteConfig) (cluster.Cluster, error) {
	logger := log.FromContext(ctx)

	if c, ok := r.RemoteClusters[key]; ok {
		return c, nil
	}

	restConfig, err := r.RemoteConfigProvider.RESTConfig(ctx, namespace, remote)
	if err != nil {
		return nil, err
	}

	c, err := cluster.New(restConfig, func(o *cluster.Options) {
		o.Scheme = r.Scheme
		if r.SyncPeriod != nil && *r.SyncPeriod > 0 {
			o.Cache.SyncPeriod = r.SyncPeriod
		}
	})
	if err != nil {
		return nil, err
	}

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
	r.remoteClients = newRemoteClientCache(r.RemoteConfigProvider)
	return ctrl.NewControllerManagedBy(mgr).
		For(&krmv1alpha1.KRMSyncer{}).
		Complete(r)
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
}

func (r *DynamicResourceReconciler) ruleMatchesGVK(rule krmv1alpha1.ResourceRule, gvk schema.GroupVersionKind) bool {
	isKCCGlob := (rule.Group == "*.cnrm.cloud.google.com" ||
		strings.HasSuffix(rule.Group, ".cnrm.cloud.google.com") ||
		rule.Group == "cnrm.cloud.google.com") && rule.Version == "*" && rule.Kind == "*"

	if isKCCGlob {
		if rule.Group == "*.cnrm.cloud.google.com" {
			return strings.HasSuffix(gvk.Group, "cnrm.cloud.google.com")
		}
		return gvk.Group == rule.Group
	}
	return rule.Group == gvk.Group && rule.Version == gvk.Version && rule.Kind == gvk.Kind
}

func (r *DynamicResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the resource from Source cluster
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(r.GVK)
	err := r.RemoteClient.Get(ctx, req.NamespacedName, u)
	isDeleted := false
	if err != nil {
		if client.IgnoreNotFound(err) == nil {
			isDeleted = true
		} else {
			return ctrl.Result{}, err
		}
	}

	// Fetch all Syncers to find matching rules (Active ones)
	var krmsyncers krmv1alpha1.KRMSyncerList
	if err := r.LocalClient.List(ctx, &krmsyncers); err != nil {
		return ctrl.Result{}, err
	}

	var syncErrors []error
	for _, krmsyncer := range krmsyncers.Items {
		if krmsyncer.Spec.Suspend {
			continue
		}

		// Check Mode match
		syncerMode := krmsyncer.Spec.Mode
		if syncerMode == "" {
			syncerMode = krmv1alpha1.ModePull
		}
		if syncerMode != r.Mode {
			continue
		}

		// For Pull mode, also check Remote match
		if r.Mode == krmv1alpha1.ModePull {
			connKey, err := r.RemoteConfigProvider.Key(krmsyncer.Namespace, krmsyncer.Spec.Remote)
			if err != nil || connKey != r.Remote.Cluster {
				continue
			}
		}

		for _, rule := range krmsyncer.Spec.Rules {
			// Check GVK match
			if !r.ruleMatchesGVK(rule, r.GVK) {
				continue
			}

			// Check Namespace match
			if len(rule.Namespaces) > 0 {
				matched := false
				for _, ns := range rule.Namespaces {
					if ns == req.Namespace {
						matched = true
						break
					}
				}
				if !matched {
					continue
				}
			}

			// Determine Destination Client
			var destClient client.Client
			if r.Mode == krmv1alpha1.ModePush {
				var err error
				destClient, err = r.getRemoteClient(ctx, &krmsyncer)
				if err != nil {
					logger.Error(err, "Failed to get remote client")
					syncErrors = append(syncErrors, err)
					continue
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
				if err := destClient.Delete(ctx, toDelete); err != nil {
					if !errors.IsNotFound(err) {
						// TODO: report failure in syncer status
						logger.Error(err, "Failed to delete resource on destination cluster")
						recordSyncFailure(r.GVK, OperationDelete)
						syncErrors = append(syncErrors, err)
					}
				}
				continue
			}

			// Sync to Destination Cluster
			logger.Info("Syncing resource to destination cluster", "name", u.GetName(), "namespace", u.GetNamespace())

			syncFields := rule.SyncFields
			destObj, err := r.filterFields(u, syncFields)
			if err != nil {
				// TODO: report failure in syncer status
				logger.Error(err, "Failed to filter fields")
				syncErrors = append(syncErrors, err)
				continue
			}

			// Prepare object for destination: clear ResourceVersion, UID, etc.
			destObj.SetResourceVersion("")
			destObj.SetUID("")
			destObj.SetGeneration(0)
			destObj.SetManagedFields(nil)

			if err := r.applyToDestination(ctx, destClient, destObj); err != nil {
				// TODO: report failure in syncer status
				logger.Error(err, "Failed to apply to destination cluster")
				syncErrors = append(syncErrors, err)
			} else {
				logger.Info("Successfully synced resource to destination cluster", "name", u.GetName(), "namespace", u.GetNamespace())
			}
		}
	}
	return ctrl.Result{}, stderrors.Join(syncErrors...)
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
				if err := destClient.Create(ctx, ns); err != nil {
					recordSyncFailure(r.GVK, OperationApply)
					return fmt.Errorf("failed to create destination namespace %q: %v", nsName, err)
				}
			} else {
				recordSyncFailure(r.GVK, OperationApply)
				return fmt.Errorf("failed to check destination namespace %q: %v", nsName, err)
			}
		}
	}
	patchOpts := []client.PatchOption{
		client.ForceOwnership,
		client.FieldOwner("krm-syncer"),
	}

	status, found, err := unstructured.NestedFieldCopy(obj.Object, "status")
	if err != nil {
		recordSyncFailure(r.GVK, OperationApply)
		return fmt.Errorf("failed to copy status: %v", err)
	}

	// Use Server-Side Apply (SSA) to create or update the resource.
	if err := destClient.Patch(ctx, obj, client.Apply, patchOpts...); err != nil {
		recordSyncFailure(r.GVK, OperationApply)
		return fmt.Errorf("failed to apply object to destination: %v", err)
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
			recordSyncFailure(r.GVK, OperationStatus)
			return fmt.Errorf("failed to apply status subresource to destination: %v", err)
		}
	}
	return nil
}
