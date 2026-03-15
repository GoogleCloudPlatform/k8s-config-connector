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
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

// conflictClient wraps a fake client and injects a conflict error on the first N Update calls.
type conflictClient struct {
	client.Client
	conflictCount    int
	conflictsToThrow int
}

func (c *conflictClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	if c.conflictCount < c.conflictsToThrow {
		c.conflictCount++
		return errors.NewConflict(schema.GroupResource{Group: "syncer.gkelabs.io", Resource: "krmsyncers"}, obj.GetName(), fmt.Errorf("simulated conflict %d", c.conflictCount))
	}
	return c.Client.Update(ctx, obj, opts...)
}

func TestEnsureSuspended_RetryOnConflict(t *testing.T) {
	name := types.NamespacedName{Name: "test-syncer-retry", Namespace: "default"}

	// Create an active syncer
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KRMSyncerGVK)
	u.SetName(name.Name)
	u.SetNamespace(name.Namespace)
	_ = unstructured.SetNestedField(u.Object, false, "spec", "suspend")

	fakeC := fake.NewClientBuilder().WithScheme(buildScheme()).WithObjects(u).Build()

	// Wrap it to throw 2 conflicts
	conflictC := &conflictClient{
		Client:           fakeC,
		conflictsToThrow: 2,
	}

	si := &SyncerIntegration{
		client:          conflictC,
		apiReader:       conflictC, // Use the same mock for reading so it can find the object
		name:            name,
		replicationMode: "Status",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// This should succeed because it will retry after the 2 conflicts
	if err := si.EnsureSuspended(ctx); err != nil {
		t.Fatalf("expected EnsureSuspended to succeed after retries, got error: %v", err)
	}

	if conflictC.conflictCount != 2 {
		t.Errorf("expected exactly 2 conflicts to be thrown and caught, but got %d", conflictC.conflictCount)
	}

	// Verify the final state is actually suspended
	updatedSyncer := &unstructured.Unstructured{}
	updatedSyncer.SetGroupVersionKind(KRMSyncerGVK)
	if err := fakeC.Get(ctx, name, updatedSyncer); err != nil {
		t.Fatalf("failed to get updated syncer: %v", err)
	}
	suspend, _, _ := unstructured.NestedBool(updatedSyncer.Object, "spec", "suspend")
	if !suspend {
		t.Errorf("expected syncer to be suspended after retries")
	}
}

func TestEnsurePullingFromLeader_RetryOnConflict(t *testing.T) {
	name := types.NamespacedName{Name: "test-syncer-retry", Namespace: "default"}
	myIdentity := "cluster-a"
	leaderIdentity := "cluster-b"

	// Create an active syncer
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(KRMSyncerGVK)
	u.SetName(name.Name)
	u.SetNamespace(name.Namespace)

	// Create a lease pointing to the other cluster
	lease := newMultiClusterLease(name, leaderIdentity)

	fakeC := fake.NewClientBuilder().WithScheme(buildScheme()).WithObjects(u, lease).Build()

	// Wrap it to throw 2 conflicts
	conflictC := &conflictClient{
		Client:           fakeC,
		conflictsToThrow: 2,
	}

	si := &SyncerIntegration{
		client:          conflictC,
		apiReader:       conflictC,
		name:            name,
		replicationMode: "Status",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// This should succeed because it will retry after the 2 conflicts
	if err := si.EnsurePullingFromLeader(ctx, myIdentity); err != nil {
		t.Fatalf("expected EnsurePullingFromLeader to succeed after retries, got error: %v", err)
	}

	if conflictC.conflictCount != 2 {
		t.Errorf("expected exactly 2 conflicts to be thrown and caught, but got %d", conflictC.conflictCount)
	}

	// Verify the final state is correctly pointing to the leader
	updatedSyncer := &unstructured.Unstructured{}
	updatedSyncer.SetGroupVersionKind(KRMSyncerGVK)
	if err := fakeC.Get(ctx, name, updatedSyncer); err != nil {
		t.Fatalf("failed to get updated syncer: %v", err)
	}
	remote, _, _ := unstructured.NestedString(updatedSyncer.Object, "spec", "remote", "clusterConfig", "kubeConfigSecretRef", "name")
	if remote != leaderIdentity {
		t.Errorf("expected destination to be %q, got %q", leaderIdentity, remote)
	}
}
