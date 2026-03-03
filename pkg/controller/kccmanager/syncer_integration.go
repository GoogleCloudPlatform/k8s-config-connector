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

	mclv1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
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
	client client.Client
	name   types.NamespacedName
}

func (si *SyncerIntegration) EnsurePullingFromLeader(ctx context.Context, apiReader client.Reader, myIdentity string) error {
	if si == nil || si.client == nil {
		return fmt.Errorf("syncer integration or client not initialized")
	}

	lease := &mclv1alpha1.MultiClusterLease{}
	if err := apiReader.Get(ctx, si.name, lease); err != nil {
		return fmt.Errorf("error fetching MultiClusterLease %s: %w", si.name, err)
	}

	var leaderIdentity string
	if lease.Status.GlobalHolderIdentity != nil {
		leaderIdentity = *lease.Status.GlobalHolderIdentity
	}

	if leaderIdentity == "" {
		klog.Warningf("MultiClusterLease %s has no GlobalHolderIdentity", si.name)
		return nil
	}

	if leaderIdentity == myIdentity {
		klog.Warningf("MultiClusterLease %s still has our identity %s as GlobalHolderIdentity, not reconfiguring syncer to pull from ourselves", si.name, myIdentity)
		return nil
	}

	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)

	if err := si.client.Get(ctx, si.name, syncer); err != nil {
		klog.Errorf("error getting KRMSyncer %s: %v", si.name, err)
		return fmt.Errorf("error getting KRMSyncer %s: %w", si.name, err)
	}

	changed := false

	dest, _, _ := unstructured.NestedString(syncer.Object, "spec", "destinationKubeConfig")
	if dest != leaderIdentity {
		if err := unstructured.SetNestedField(syncer.Object, leaderIdentity, "spec", "destinationKubeConfig"); err != nil {
			return fmt.Errorf("error setting destinationKubeConfig: %w", err)
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

	suspend, _, _ := unstructured.NestedBool(syncer.Object, "spec", "suspend")
	if suspend {
		if err := unstructured.SetNestedField(syncer.Object, false, "spec", "suspend"); err != nil {
			return fmt.Errorf("error setting suspend: %w", err)
		}
		changed = true
	}

	if !changed {
		klog.Infof("KRMSyncer %s is already configured to pull from leader %s", si.name, leaderIdentity)
		return nil
	}

	if err := si.client.Update(ctx, syncer); err != nil {
		klog.Errorf("error updating KRMSyncer %s: %v", si.name, err)
		return fmt.Errorf("error updating KRMSyncer %s: %w", si.name, err)
	}

	klog.Infof("successfully configured KRMSyncer %s to pull from leader %s", si.name, leaderIdentity)
	return nil
}

func (si *SyncerIntegration) EnsureSuspended(ctx context.Context) error {
	if si == nil || si.client == nil {
		return fmt.Errorf("syncer integration or client not initialized")
	}

	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)

	if err := si.client.Get(ctx, si.name, syncer); err != nil {
		klog.Errorf("error getting KRMSyncer %s: %v", si.name, err)
		return fmt.Errorf("error getting KRMSyncer %s: %w", si.name, err)
	}

	suspend, _, _ := unstructured.NestedBool(syncer.Object, "spec", "suspend")
	if suspend {
		klog.Infof("KRMSyncer %s is already suspended", si.name)
		return nil
	}

	if err := unstructured.SetNestedField(syncer.Object, true, "spec", "suspend"); err != nil {
		return fmt.Errorf("error setting suspend: %w", err)
	}

	if err := si.client.Update(ctx, syncer); err != nil {
		klog.Errorf("error updating KRMSyncer %s: %v", si.name, err)
		return fmt.Errorf("error updating KRMSyncer %s: %w", si.name, err)
	}

	klog.Infof("successfully suspended KRMSyncer %s", si.name)
	return nil
}

func (si *SyncerIntegration) EnsureSyncing(ctx context.Context) error {
	if si == nil || si.client == nil {
		return fmt.Errorf("syncer integration or client not initialized")
	}

	syncer := &unstructured.Unstructured{}
	syncer.SetGroupVersionKind(KRMSyncerGVK)

	if err := si.client.Get(ctx, si.name, syncer); err != nil {
		klog.Errorf("error getting KRMSyncer %s: %v", si.name, err)
		return fmt.Errorf("error getting KRMSyncer %s: %w", si.name, err)
	}

	suspend, _, _ := unstructured.NestedBool(syncer.Object, "spec", "suspend")
	if !suspend {
		klog.Infof("KRMSyncer %s is already active", si.name)
		return nil
	}

	if err := unstructured.SetNestedField(syncer.Object, false, "spec", "suspend"); err != nil {
		return fmt.Errorf("error setting suspend: %w", err)
	}

	if err := si.client.Update(ctx, syncer); err != nil {
		klog.Errorf("error updating KRMSyncer %s: %v", si.name, err)
		return fmt.Errorf("error updating KRMSyncer %s: %w", si.name, err)
	}

	klog.Infof("successfully activated KRMSyncer %s", si.name)
	return nil
}
