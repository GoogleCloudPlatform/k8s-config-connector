// Copyright 2026 Google LLC
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

package kccmanager

import (
	"context"
	stderrors "errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	operatorv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	KRMSyncerGVK = schema.GroupVersionKind{
		Group:   "syncer.gkelabs.io",
		Version: "v1alpha1",
		Kind:    "KRMSyncer",
	}
)

// syncerDiscoveryTimeout bounds one discovery request of the syncer integration. A standby
// asks what the cluster serves every time it updates its KRMSyncers, and that update must end
// before the cluster can start to lead; see leaderFollower.stop.
const syncerDiscoveryTimeout = 10 * time.Second

// newSyncerDiscovery returns the discovery client that tells the syncer integration which
// kinds the cluster serves.
func newSyncerDiscovery(restConfig *rest.Config) (servedResources, error) {
	config := rest.CopyConfig(restConfig)
	config.Timeout = syncerDiscoveryTimeout
	return discovery.NewDiscoveryClientForConfig(config)
}

type SyncerIntegration struct {
	client    client.Client
	apiReader client.Reader
	// discovery tells which kinds this cluster serves. See servedGVKsWithServiceGeneratedIDs.
	discovery       servedResources
	leaseNN         types.NamespacedName
	replicationMode string
	watchNamespaces []string
	// cluster is the GKE cluster that this Config Connector runs in.
	cluster *operatorv1beta1.GKECluster
	// leaderEndpointType is the control plane endpoint used to reach the leader's cluster.
	leaderEndpointType string
}

// servedResources lists the resources that an API server serves. A discovery client is one.
type servedResources interface {
	ServerGroupsAndResources() ([]*metav1.APIGroup, []*metav1.APIResourceList, error)
}

// isFullReplication reports whether the KRMSyncers copy whole objects. Otherwise they copy the
// status, and spec.resourceID of the kinds whose ID the service generates.
func (si *SyncerIntegration) isFullReplication() bool {
	return strings.EqualFold(si.replicationMode, "Full")
}

// servedGVKsWithServiceGeneratedIDs returns the kinds of syncerGVKsWithServiceGeneratedIDs
// that this cluster serves, in a fixed order. These are the kinds that get a KRMSyncer rule
// of their own.
//
// A rule that names a kind must name one that is served. The syncer watches the kind of such
// a rule on the leader's cluster without asking first whether the leader serves it, and a
// watch that can't start makes the syncer exit two minutes later, again and again. The
// static list has kinds that a cluster doesn't serve: a release doesn't install the CRDs
// that only have an alpha version.
//
// This asks this cluster's API server, and not the leader's, which the manager can't reach.
// The clusters of a group are meant to run the same version of Config Connector, and then
// they serve the same kinds. While they don't, a rule can still name a kind that the leader
// doesn't serve.
func (si *SyncerIntegration) servedGVKsWithServiceGeneratedIDs() ([]schema.GroupVersionKind, error) {
	if si.discovery == nil {
		return nil, fmt.Errorf("discovery client not initialized")
	}

	// With aggregated discovery, which API servers have had since Kubernetes 1.27, this is
	// two requests.
	_, resourceLists, err := si.discovery.ServerGroupsAndResources()
	if err != nil {
		// Discovery can fail for some groups and still answer for the others, for example
		// while the server of an aggregated API is down. That only matters here if it is a
		// group of the list.
		var failed *discovery.ErrGroupDiscoveryFailed
		if !stderrors.As(err, &failed) {
			return nil, fmt.Errorf("error discovering the kinds that the cluster serves: %w", err)
		}
		for gvk := range syncerGVKsWithServiceGeneratedIDs {
			if groupErr, ok := failed.Groups[gvk.GroupVersion()]; ok {
				return nil, fmt.Errorf("error discovering the kinds of %s: %w", gvk.GroupVersion(), groupErr)
			}
		}
	}

	servedSet := make(map[schema.GroupVersionKind]struct{})
	for _, resourceList := range resourceLists {
		gv, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			continue
		}
		for _, resource := range resourceList.APIResources {
			if strings.Contains(resource.Name, "/") {
				continue // A subresource, such as status.
			}
			gvk := gv.WithKind(resource.Kind)
			if _, ok := syncerGVKsWithServiceGeneratedIDs[gvk]; ok {
				servedSet[gvk] = struct{}{}
			}
		}
	}

	served := make([]schema.GroupVersionKind, 0, len(servedSet))
	for gvk := range servedSet {
		served = append(served, gvk)
	}
	sort.Slice(served, func(i, j int) bool {
		return served[i].String() < served[j].String()
	})
	return served, nil
}

// getDesiredRules returns the rules of a KRMSyncer for the given namespaces. gvksWithOwnRule
// are the kinds that get a rule of their own, which also copies spec.resourceID; see
// servedGVKsWithServiceGeneratedIDs. They are not used in Full mode, where one rule copies
// everything.
func (si *SyncerIntegration) getDesiredRules(namespaces []string, gvksWithOwnRule []schema.GroupVersionKind) []interface{} {
	var nsList []interface{}
	for _, ns := range namespaces {
		nsList = append(nsList, ns)
	}

	createRule := func(group, version, kind string, syncFields []interface{}) map[string]interface{} {
		rule := map[string]interface{}{
			"group":      group,
			"version":    version,
			"kind":       kind,
			"syncFields": syncFields,
		}
		if len(nsList) > 0 {
			rule["namespaces"] = nsList
		}
		return rule
	}

	if si.isFullReplication() {
		return []interface{}{
			createRule("*.cnrm.cloud.google.com", "*", "*", []interface{}{"spec", "status"}),
		}
	}

	rules := []interface{}{
		createRule("*.cnrm.cloud.google.com", "*", "*", []interface{}{"status"}),
	}

	for _, gvk := range gvksWithOwnRule {
		rules = append(rules, createRule(gvk.Group, gvk.Version, gvk.Kind, []interface{}{"status", "spec.resourceID"}))
	}
	return rules
}

type ruleData struct {
	fields     []string
	namespaces []string
}

func parseRuleData(r map[string]interface{}) ruleData {
	data := ruleData{}
	if fields, ok := r["syncFields"].([]interface{}); ok {
		for _, f := range fields {
			if s, ok := f.(string); ok {
				data.fields = append(data.fields, s)
			} else {
				klog.Warningf("ignoring non-string element in syncFields: %v (type %T)", f, f)
			}
		}
	}
	if nsList, ok := r["namespaces"].([]interface{}); ok {
		for _, ns := range nsList {
			if s, ok := ns.(string); ok {
				data.namespaces = append(data.namespaces, s)
			} else {
				klog.Warningf("ignoring non-string element in namespaces: %v (type %T)", ns, ns)
			}
		}
	}
	return data
}

func rulesMatch(existing []interface{}, desired []interface{}) bool {
	if len(existing) != len(desired) {
		return false
	}

	type ruleKey struct {
		group, version, kind string
	}

	// Use a slice to map to correctly handle duplicate GVK rules if manually added by users
	existingMap := make(map[ruleKey][]ruleData)
	for _, e := range existing {
		r, ok := e.(map[string]interface{})
		if !ok {
			return false
		}
		group, _ := r["group"].(string)
		version, _ := r["version"].(string)
		kind, _ := r["kind"].(string)
		key := ruleKey{group, version, kind}
		existingMap[key] = append(existingMap[key], parseRuleData(r))
	}

	for _, d := range desired {
		r, ok := d.(map[string]interface{})
		if !ok {
			return false
		}
		group, _ := r["group"].(string)
		version, _ := r["version"].(string)
		kind, _ := r["kind"].(string)

		key := ruleKey{group, version, kind}
		existingDataList, found := existingMap[key]

		if !found || len(existingDataList) == 0 {
			return false
		}

		desiredData := parseRuleData(r)

		// Find a matching ruleData within the list for this GVK
		matchFound := false
		for i, existingData := range existingDataList {
			if len(existingData.fields) != len(desiredData.fields) || len(existingData.namespaces) != len(desiredData.namespaces) {
				continue
			}

			fieldSet := make(map[string]bool)
			for _, f := range existingData.fields {
				fieldSet[f] = true
			}
			fieldsMatch := true
			for _, f := range desiredData.fields {
				if !fieldSet[f] {
					fieldsMatch = false
					break
				}
			}
			if !fieldsMatch {
				continue
			}

			nsSet := make(map[string]bool)
			for _, ns := range existingData.namespaces {
				nsSet[ns] = true
			}
			nsMatch := true
			for _, ns := range desiredData.namespaces {
				if !nsSet[ns] {
					nsMatch = false
					break
				}
			}
			if !nsMatch {
				continue
			}

			// Match found! Remove it from the list to handle duplicates properly
			existingMap[key] = append(existingMap[key][:i], existingMap[key][i+1:]...)
			matchFound = true
			break
		}

		if !matchFound {
			return false
		}
	}
	return true
}

func (si *SyncerIntegration) getDesiredSyncers() []types.NamespacedName {
	var syncers []types.NamespacedName
	if len(si.watchNamespaces) == 0 {
		// Cluster scoped - one global syncer
		syncers = append(syncers, types.NamespacedName{
			Namespace: si.leaseNN.Namespace,
			Name:      si.leaseNN.Name,
		})
	} else {
		// Namespaced scoped - one syncer per watched namespace, located in the lease namespace
		for _, ns := range si.watchNamespaces {
			syncers = append(syncers, types.NamespacedName{
				Namespace: si.leaseNN.Namespace,
				Name:      fmt.Sprintf("%s-%s", si.leaseNN.Name, ns),
			})
		}
	}
	return syncers
}

func (si *SyncerIntegration) getNamespacesForSyncer(syncerName string) []string {
	if len(si.watchNamespaces) == 0 {
		return nil // Cluster scoped
	}
	for _, ns := range si.watchNamespaces {
		if syncerName == fmt.Sprintf("%s-%s", si.leaseNN.Name, ns) {
			return []string{ns}
		}
	}
	return nil
}

// gvksWithOwnRule returns the kinds that get a KRMSyncer rule of their own. There are none in
// Full mode.
func (si *SyncerIntegration) gvksWithOwnRule() ([]schema.GroupVersionKind, error) {
	if si.isFullReplication() {
		return nil, nil
	}
	return si.servedGVKsWithServiceGeneratedIDs()
}

// EnsurePullingFromLeader points this cluster's KRMSyncers at the leader's GKE cluster and
// resumes them.
func (si *SyncerIntegration) EnsurePullingFromLeader(ctx context.Context, leader *operatorv1beta1.GKECluster) error {
	if si == nil || si.client == nil || si.apiReader == nil {
		return fmt.Errorf("syncer integration or client not initialized")
	}
	if leader == nil {
		return fmt.Errorf("the leader's GKE cluster is not set")
	}

	gvksWithOwnRule, err := si.gvksWithOwnRule()
	if err != nil {
		return err
	}
	for _, syncerNN := range si.getDesiredSyncers() {
		if err := si.reconcileSyncer(ctx, syncerNN, leader, false, gvksWithOwnRule); err != nil {
			return err
		}
	}
	return nil
}

// EnsureSuspended suspends this cluster's KRMSyncers because this cluster is the leader. Their
// remote is set to this cluster's own GKE cluster, so that on every cluster the KRMSyncers name
// the current leader and are suspended exactly when that leader is this cluster.
func (si *SyncerIntegration) EnsureSuspended(ctx context.Context) error {
	if si == nil || si.client == nil || si.apiReader == nil {
		return fmt.Errorf("syncer integration or client not initialized")
	}
	if si.cluster == nil {
		return fmt.Errorf("this cluster's GKE cluster is not set")
	}

	gvksWithOwnRule, err := si.gvksWithOwnRule()
	if err != nil {
		return err
	}
	for _, syncerNN := range si.getDesiredSyncers() {
		if err := si.reconcileSyncer(ctx, syncerNN, si.cluster, true, gvksWithOwnRule); err != nil {
			return err
		}
	}
	return nil
}

// desiredRemote returns spec.remote for a KRMSyncer that pulls from the given GKE cluster. The
// endpoint is always set, so that the CRD default doesn't show up as a difference.
func (si *SyncerIntegration) desiredRemote(cluster *operatorv1beta1.GKECluster) map[string]interface{} {
	endpoint := si.leaderEndpointType
	if endpoint == "" {
		endpoint = operatorv1beta1.LeaderEndpointTypeDefault
	}
	return map[string]interface{}{
		"gkeCluster": map[string]interface{}{
			"project":  cluster.Project,
			"location": cluster.Location,
			"name":     cluster.Name,
			"endpoint": endpoint,
		},
	}
}

func (si *SyncerIntegration) reconcileSyncer(ctx context.Context, name types.NamespacedName, remoteCluster *operatorv1beta1.GKECluster, suspend bool, gvksWithOwnRule []schema.GroupVersionKind) error {
	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)

	var isCreate bool
	if err := si.apiReader.Get(ctx, name, syncer); err != nil {
		if errors.IsNotFound(err) {
			klog.Infof("KRMSyncer %s not found. Will create it.", name)
			isCreate = true
		} else {
			klog.Errorf("error getting KRMSyncer %s: %v", name, err)
			return fmt.Errorf("error getting KRMSyncer %s: %w", name, err)
		}
	}

	namespaces := si.getNamespacesForSyncer(name.Name)
	rules := si.getDesiredRules(namespaces, gvksWithOwnRule)
	remote := si.desiredRemote(remoteCluster)

	if isCreate {
		syncerToCreate := &unstructured.Unstructured{}
		syncerToCreate.SetGroupVersionKind(KRMSyncerGVK)
		syncerToCreate.SetName(name.Name)
		syncerToCreate.SetNamespace(name.Namespace)

		_ = unstructured.SetNestedMap(syncerToCreate.Object, remote, "spec", "remote")
		_ = unstructured.SetNestedField(syncerToCreate.Object, "pull", "spec", "mode")
		_ = unstructured.SetNestedField(syncerToCreate.Object, suspend, "spec", "suspend")
		_ = unstructured.SetNestedSlice(syncerToCreate.Object, rules, "spec", "rules")

		if err := si.client.Create(ctx, syncerToCreate); err != nil {
			return fmt.Errorf("error creating KRMSyncer %s: %w", name, err)
		}
		klog.Infof("successfully created KRMSyncer %s (suspend=%v, remote=%s)", name, suspend, remoteCluster.ResourceName())
		return nil
	}

	changed := false

	// Replace spec.remote as a whole, so that nothing from an older remote type is left behind.
	existingRemote, _, _ := unstructured.NestedMap(syncer.Object, "spec", "remote")
	if !reflect.DeepEqual(existingRemote, remote) {
		if err := unstructured.SetNestedMap(syncer.Object, remote, "spec", "remote"); err != nil {
			return fmt.Errorf("error setting remote: %w", err)
		}
		changed = true
	}

	mode, _, _ := unstructured.NestedString(syncer.Object, "spec", "mode")
	if mode != "pull" {
		klog.Infof("KRMSyncer %s mode is %q, setting to %q", name, mode, "pull")
		if err := unstructured.SetNestedField(syncer.Object, "pull", "spec", "mode"); err != nil {
			return fmt.Errorf("error setting mode: %w", err)
		}
		changed = true
	}

	existingRules, found, _ := unstructured.NestedSlice(syncer.Object, "spec", "rules")
	if !found || !rulesMatch(existingRules, rules) {
		if err := unstructured.SetNestedSlice(syncer.Object, rules, "spec", "rules"); err != nil {
			return fmt.Errorf("error setting rules: %w", err)
		}
		changed = true
	}

	existingSuspend, found, _ := unstructured.NestedBool(syncer.Object, "spec", "suspend")
	if !found || existingSuspend != suspend {
		if err := unstructured.SetNestedField(syncer.Object, suspend, "spec", "suspend"); err != nil {
			return fmt.Errorf("error setting suspend: %w", err)
		}
		changed = true
	}

	if !changed {
		klog.Infof("KRMSyncer %s is already configured correctly (suspend=%v)", name, suspend)
		return nil
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		syncerToUpdate := &unstructured.Unstructured{}
		syncerToUpdate.SetGroupVersionKind(KRMSyncerGVK)
		if err := si.apiReader.Get(ctx, name, syncerToUpdate); err != nil {
			return err
		}

		_ = unstructured.SetNestedMap(syncerToUpdate.Object, remote, "spec", "remote")
		_ = unstructured.SetNestedField(syncerToUpdate.Object, "pull", "spec", "mode")
		_ = unstructured.SetNestedField(syncerToUpdate.Object, suspend, "spec", "suspend")
		_ = unstructured.SetNestedSlice(syncerToUpdate.Object, rules, "spec", "rules")

		return si.client.Update(ctx, syncerToUpdate)
	})

	if err != nil {
		klog.Errorf("error updating KRMSyncer %s: %v", name, err)
		return fmt.Errorf("error updating KRMSyncer %s: %w", name, err)
	}

	klog.Infof("successfully updated KRMSyncer %s (suspend=%v, remote=%s)", name, suspend, remoteCluster.ResourceName())
	return nil
}
