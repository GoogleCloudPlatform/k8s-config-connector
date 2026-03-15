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

	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
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

// SyncerIntegration manages the lifecycle and configuration of a KRMSyncer resource
// in response to leader election state changes.
//
// Implicit Assumption: The KRMSyncer resource is expected to share the exact same
// name and namespace as the MultiClusterLease resource being used for leader election.
type SyncerIntegration struct {
	client          client.Client
	apiReader       client.Reader
	name            types.NamespacedName
	replicationMode string
}

func (si *SyncerIntegration) getDesiredRules() []interface{} {
	if strings.EqualFold(si.replicationMode, "Full") {
		return []interface{}{
			map[string]interface{}{
				"group":      "*.cnrm.cloud.google.com",
				"version":    "*",
				"kind":       "*",
				"syncFields": []interface{}{"spec", "status"},
			},
		}
	}

	// Status Only mode:
	// 1. One glob rule for the majority of resources (status only)
	// 2. Explicit rules for resources with service-generated IDs (status + spec.resourceID)
	rules := []interface{}{
		map[string]interface{}{
			"group":      "*.cnrm.cloud.google.com",
			"version":    "*",
			"kind":       "*",
			"syncFields": []interface{}{"status"},
		},
	}

	// Add exceptions (GVKs we know have service-generated IDs)
	for gvk := range syncerGVKsWithServiceGeneratedIDs {
		rules = append(rules, map[string]interface{}{
			"group":      gvk.Group,
			"version":    gvk.Version,
			"kind":       gvk.Kind,
			"syncFields": []interface{}{"status", "spec.resourceID"},
		})
	}
	return rules
}

func rulesMatch(existing []interface{}, desired []interface{}) bool {
	if len(existing) != len(desired) {
		return false
	}

	// Create a map-based lookup for existing rules to handle potential reordering
	// by the API server or during generation.
	type ruleKey struct {
		group, version, kind string
	}
	existingMap := make(map[ruleKey][]string)
	for _, e := range existing {
		r, ok := e.(map[string]interface{})
		if !ok {
			return false
		}
		group, _ := r["group"].(string)
		version, _ := r["version"].(string)
		kind, _ := r["kind"].(string)
		fields, _ := r["syncFields"].([]interface{})

		var fieldStrs []string
		for _, f := range fields {
			if s, ok := f.(string); ok {
				fieldStrs = append(fieldStrs, s)
			}
		}
		existingMap[ruleKey{group, version, kind}] = fieldStrs
	}

	for _, d := range desired {
		r, ok := d.(map[string]interface{})
		if !ok {
			return false
		}
		group, _ := r["group"].(string)
		version, _ := r["version"].(string)
		kind, _ := r["kind"].(string)
		desiredFields, _ := r["syncFields"].([]interface{})

		existingFields, found := existingMap[ruleKey{group, version, kind}]
		if !found || len(existingFields) != len(desiredFields) {
			return false
		}

		fieldSet := make(map[string]bool)
		for _, f := range existingFields {
			fieldSet[f] = true
		}
		for _, f := range desiredFields {
			s, _ := f.(string)
			if !fieldSet[s] {
				return false
			}
		}
	}
	return true
}

func (si *SyncerIntegration) EnsurePullingFromLeader(ctx context.Context, myIdentity string) error {
	if si == nil || si.client == nil || si.apiReader == nil {
		return fmt.Errorf("syncer integration or client not initialized")
	}

	lease := &mclv1alpha1.MultiClusterLease{}
	if err := si.apiReader.Get(ctx, si.name, lease); err != nil {
		if errors.IsNotFound(err) {
			klog.Warningf("MultiClusterLease %s not found. Skipping syncer reconfiguration.", si.name)
			return nil
		}
		return fmt.Errorf("error fetching MultiClusterLease %s: %w", si.name, err)
	}

	var leaderIdentity string
	if lease.Status.GlobalHolderIdentity != nil {
		leaderIdentity = *lease.Status.GlobalHolderIdentity
	}

	if leaderIdentity == "" {
		klog.V(2).Infof("MultiClusterLease %s has no GlobalHolderIdentity", si.name)
		return nil
	}

	if leaderIdentity == myIdentity {
		klog.Warningf("MultiClusterLease %s still has our identity %s as GlobalHolderIdentity, not reconfiguring syncer to pull from ourselves", si.name, myIdentity)
		return nil
	}

	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)

	var isCreate bool
	if err := si.apiReader.Get(ctx, si.name, syncer); err != nil {
		if errors.IsNotFound(err) {
			klog.Infof("KRMSyncer %s not found. Will create it to pull from leader.", si.name)
			isCreate = true
		} else {
			klog.Errorf("error getting KRMSyncer %s: %v", si.name, err)
			return fmt.Errorf("error getting KRMSyncer %s: %w", si.name, err)
		}
	}

	rules := si.getDesiredRules()

	if isCreate {
		syncerToCreate := &unstructured.Unstructured{}
		syncerToCreate.SetGroupVersionKind(KRMSyncerGVK)
		syncerToCreate.SetName(si.name.Name)
		syncerToCreate.SetNamespace(si.name.Namespace)

		_ = unstructured.SetNestedField(syncerToCreate.Object, leaderIdentity, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
		_ = unstructured.SetNestedField(syncerToCreate.Object, "pull", "spec", "mode")
		_ = unstructured.SetNestedField(syncerToCreate.Object, false, "spec", "suspend")
		_ = unstructured.SetNestedSlice(syncerToCreate.Object, rules, "spec", "rules")

		if err := si.client.Create(ctx, syncerToCreate); err != nil {
			return fmt.Errorf("error creating KRMSyncer %s: %w", si.name, err)
		}
		klog.Infof("successfully created KRMSyncer %s to pull from leader %s", si.name, leaderIdentity)
		return nil
	}

	changed := false

	remote, _, _ := unstructured.NestedString(syncer.Object, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
	if remote != leaderIdentity {
		if err := unstructured.SetNestedField(syncer.Object, leaderIdentity, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name"); err != nil {
			return fmt.Errorf("error setting remote kubeconfig name: %w", err)
		}
		changed = true
	}

	mode, _, _ := unstructured.NestedString(syncer.Object, "spec", "mode")
	if mode != "pull" {
		klog.Infof("KRMSyncer %s mode is %q, setting to %q", si.name, mode, "pull")
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

	suspend, found, _ := unstructured.NestedBool(syncer.Object, "spec", "suspend")
	if !found || suspend {
		if err := unstructured.SetNestedField(syncer.Object, false, "spec", "suspend"); err != nil {
			return fmt.Errorf("error setting suspend: %w", err)
		}
		changed = true
	}

	if !changed {
		klog.Infof("KRMSyncer %s is already configured to pull from leader %s", si.name, leaderIdentity)
		return nil
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		syncerToUpdate := &unstructured.Unstructured{}
		syncerToUpdate.SetGroupVersionKind(KRMSyncerGVK)
		if err := si.apiReader.Get(ctx, si.name, syncerToUpdate); err != nil {
			return err
		}

		_ = unstructured.SetNestedField(syncerToUpdate.Object, leaderIdentity, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
		_ = unstructured.SetNestedField(syncerToUpdate.Object, "pull", "spec", "mode")
		_ = unstructured.SetNestedField(syncerToUpdate.Object, false, "spec", "suspend")
		_ = unstructured.SetNestedSlice(syncerToUpdate.Object, rules, "spec", "rules")

		return si.client.Update(ctx, syncerToUpdate)
	})

	if err != nil {
		klog.Errorf("error updating KRMSyncer %s to pull from leader %s: %v", si.name, leaderIdentity, err)
		return fmt.Errorf("error updating KRMSyncer %s: %w", si.name, err)
	}

	klog.Infof("successfully configured KRMSyncer %s to pull from leader %s", si.name, leaderIdentity)
	return nil
}

func (si *SyncerIntegration) EnsureSuspended(ctx context.Context) error {
	if si == nil || si.client == nil || si.apiReader == nil {
		return fmt.Errorf("syncer integration or client not initialized")
	}

	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)

	if err := si.apiReader.Get(ctx, si.name, syncer); err != nil {
		if errors.IsNotFound(err) {
			klog.Infof("KRMSyncer %s not found. Creating it in suspended state.", si.name)
			syncerToCreate := &unstructured.Unstructured{}
			syncerToCreate.SetGroupVersionKind(KRMSyncerGVK)
			syncerToCreate.SetName(si.name.Name)
			syncerToCreate.SetNamespace(si.name.Namespace)

			_ = unstructured.SetNestedField(syncerToCreate.Object, true, "spec", "suspend")
			_ = unstructured.SetNestedField(syncerToCreate.Object, "pull", "spec", "mode")
			_ = unstructured.SetNestedSlice(syncerToCreate.Object, si.getDesiredRules(), "spec", "rules")

			if err := si.client.Create(ctx, syncerToCreate); err != nil {
				return fmt.Errorf("error creating KRMSyncer %s: %w", si.name, err)
			}
			return nil
		}
		klog.Errorf("error getting KRMSyncer %s: %v", si.name, err)
		return fmt.Errorf("error getting KRMSyncer %s: %w", si.name, err)
	}

	suspend, found, _ := unstructured.NestedBool(syncer.Object, "spec", "suspend")
	if found && suspend {
		klog.Infof("KRMSyncer %s is already suspended", si.name)
		return nil
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		syncerToUpdate := &unstructured.Unstructured{}
		syncerToUpdate.SetGroupVersionKind(KRMSyncerGVK)
		if err := si.apiReader.Get(ctx, si.name, syncerToUpdate); err != nil {
			return err
		}

		_ = unstructured.SetNestedField(syncerToUpdate.Object, true, "spec", "suspend")
		return si.client.Update(ctx, syncerToUpdate)
	})

	if err != nil {
		klog.Errorf("error updating KRMSyncer %s: %v", si.name, err)
		return fmt.Errorf("error updating KRMSyncer %s: %w", si.name, err)
	}

	klog.Infof("successfully suspended KRMSyncer %s", si.name)
	return nil
}
