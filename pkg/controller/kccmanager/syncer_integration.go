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
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
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

type SyncerIntegration struct {
	client          client.Client
	apiReader       client.Reader
	leaseNN         types.NamespacedName
	replicationMode string
	watchNamespaces []string
}

func (si *SyncerIntegration) getDesiredRules(namespaces []string) []interface{} {
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

	if strings.EqualFold(si.replicationMode, "Full") {
		return []interface{}{
			createRule("*.cnrm.cloud.google.com", "*", "*", []interface{}{"spec", "status"}),
		}
	}

	rules := []interface{}{
		createRule("*.cnrm.cloud.google.com", "*", "*", []interface{}{"status"}),
	}

	for gvk := range syncerGVKsWithServiceGeneratedIDs {
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

func (si *SyncerIntegration) EnsurePullingFromLeader(ctx context.Context, leaderIdentity string) error {
	if si == nil || si.client == nil || si.apiReader == nil {
		return fmt.Errorf("syncer integration or client not initialized")
	}

	if leaderIdentity == "" {
		return nil
	}

	for _, syncerNN := range si.getDesiredSyncers() {
		if err := si.reconcileSyncer(ctx, syncerNN, leaderIdentity, false); err != nil {
			return err
		}
	}
	return nil
}

func (si *SyncerIntegration) EnsureSuspended(ctx context.Context) error {
	if si == nil || si.client == nil || si.apiReader == nil {
		return fmt.Errorf("syncer integration or client not initialized")
	}

	for _, syncerNN := range si.getDesiredSyncers() {
		if err := si.reconcileSyncer(ctx, syncerNN, "", true); err != nil {
			return err
		}
	}
	return nil
}

func (si *SyncerIntegration) reconcileSyncer(ctx context.Context, name types.NamespacedName, leaderIdentity string, suspend bool) error {
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
	rules := si.getDesiredRules(namespaces)

	if isCreate {
		syncerToCreate := &unstructured.Unstructured{}
		syncerToCreate.SetGroupVersionKind(KRMSyncerGVK)
		syncerToCreate.SetName(name.Name)
		syncerToCreate.SetNamespace(name.Namespace)

		dest := leaderIdentity
		if dest == "" {
			dest = "dummy-secret"
		}
		_ = unstructured.SetNestedField(syncerToCreate.Object, dest, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
		_ = unstructured.SetNestedField(syncerToCreate.Object, "pull", "spec", "mode")
		_ = unstructured.SetNestedField(syncerToCreate.Object, suspend, "spec", "suspend")
		_ = unstructured.SetNestedSlice(syncerToCreate.Object, rules, "spec", "rules")

		if err := si.client.Create(ctx, syncerToCreate); err != nil {
			return fmt.Errorf("error creating KRMSyncer %s: %w", name, err)
		}
		klog.Infof("successfully created KRMSyncer %s (suspend=%v, dest=%s)", name, suspend, dest)
		return nil
	}

	changed := false

	if suspend {
		dest, _, _ := unstructured.NestedString(syncer.Object, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
		if dest != "dummy-secret" {
			if err := unstructured.SetNestedField(syncer.Object, "dummy-secret", "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name"); err != nil {
				return fmt.Errorf("error setting remote kubeconfig name: %w", err)
			}
			changed = true
		}
	} else if leaderIdentity != "" {
		dest, _, _ := unstructured.NestedString(syncer.Object, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
		if dest != leaderIdentity {
			if err := unstructured.SetNestedField(syncer.Object, leaderIdentity, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name"); err != nil {
				return fmt.Errorf("error setting remote kubeconfig name: %w", err)
			}
			changed = true
		}
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

		dest := leaderIdentity
		if suspend {
			dest = "dummy-secret"
		}
		_ = unstructured.SetNestedField(syncerToUpdate.Object, dest, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")

		_ = unstructured.SetNestedField(syncerToUpdate.Object, "pull", "spec", "mode")
		_ = unstructured.SetNestedField(syncerToUpdate.Object, suspend, "spec", "suspend")
		_ = unstructured.SetNestedSlice(syncerToUpdate.Object, rules, "spec", "rules")

		return si.client.Update(ctx, syncerToUpdate)
	})

	if err != nil {
		klog.Errorf("error updating KRMSyncer %s: %v", name, err)
		return fmt.Errorf("error updating KRMSyncer %s: %w", name, err)
	}

	klog.Infof("successfully updated KRMSyncer %s (suspend=%v)", name, suspend)
	return nil
}
