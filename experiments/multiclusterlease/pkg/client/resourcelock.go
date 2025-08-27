// Copyright 2025 Google LLC
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

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MultiClusterLeaseLock implements the resourcelock.Interface using a MultiClusterLease CR.
type MultiClusterLeaseLock struct {
	client    client.Client
	leaseName string
	leaseNS   string
	identity  string
}

func New(client client.Client, leaseName, leaseNS, identity string) *MultiClusterLeaseLock {
	return &MultiClusterLeaseLock{
		client:    client,
		leaseName: leaseName,
		leaseNS:   leaseNS,
		identity:  identity,
	}
}

// Get returns the current leader election record from the MultiClusterLease CR.
func (mcl *MultiClusterLeaseLock) Get(ctx context.Context) (*resourcelock.LeaderElectionRecord, []byte, error) {
	lease := &v1alpha1.MultiClusterLease{}
	err := mcl.client.Get(ctx, client.ObjectKey{Namespace: mcl.leaseNS, Name: mcl.leaseName}, lease)
	if err != nil {
		return nil, nil, err
	}

	record := mcl.leaseToRecord(lease)
	recordBytes, err := json.Marshal(*record)
	if err != nil {
		return nil, nil, err
	}

	return record, recordBytes, nil
}

// Create attempts to create a new MultiClusterLease CR and polls the status
// to confirm global leadership before returning.
func (mcl *MultiClusterLeaseLock) Create(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	// Phase 1: Asynchronous Creation
	lease := &v1alpha1.MultiClusterLease{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mcl.leaseName,
			Namespace: mcl.leaseNS,
		},
		Spec: v1alpha1.MultiClusterLeaseSpec{
			HolderIdentity:       &ler.HolderIdentity,
			RenewTime:            &metav1.MicroTime{Time: ler.RenewTime.Time},
			LeaseDurationSeconds: int32Ptr(int32(ler.LeaseDurationSeconds)),
		},
	}
	err := mcl.client.Create(ctx, lease)
	if err != nil {
		return err
	}

	// After creating, get the generation that we need to see reflected in the status.
	if err := mcl.client.Get(ctx, client.ObjectKey{Namespace: mcl.leaseNS, Name: mcl.leaseName}, lease); err != nil {
		return fmt.Errorf("failed to get generation after create: %w", err)
	}
	expectedGeneration := lease.Generation

	// Phase 2: Synchronous Confirmation
	pollCtx, cancel := context.WithTimeout(ctx, time.Duration(ler.LeaseDurationSeconds)*time.Second)
	defer cancel()

	err = wait.PollUntilContextCancel(pollCtx, 1*time.Second, true, func(ctx context.Context) (bool, error) {
		var currentLease v1alpha1.MultiClusterLease
		if err := mcl.client.Get(ctx, client.ObjectKey{Namespace: mcl.leaseNS, Name: mcl.leaseName}, &currentLease); err != nil {
			return false, nil // Don't stop polling on transient errors
		}

		// Check that the controller has observed our specific spec update (or a newer one)
		// AND that it has confirmed our leadership.
		if currentLease.Status.ObservedGeneration != nil &&
			*currentLease.Status.ObservedGeneration >= expectedGeneration &&
			currentLease.Status.GlobalHolderIdentity != nil &&
			*currentLease.Status.GlobalHolderIdentity == ler.HolderIdentity {
			return true, nil // Success! Confirmed leadership.
		}
		return false, nil // Not the leader yet, or status is stale. Continue polling.
	})

	if err != nil {
		return fmt.Errorf("failed to confirm leadership in status after create: %w", err)
	}
	return nil // Successfully confirmed leadership.
}

// Update updates the spec of the MultiClusterLease CR and then polls the status
// to confirm global leadership before returning.
func (mcl *MultiClusterLeaseLock) Update(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	// Phase 1: Asynchronous Heartbeat
	lease := &v1alpha1.MultiClusterLease{}
	if err := mcl.client.Get(ctx, client.ObjectKey{Namespace: mcl.leaseNS, Name: mcl.leaseName}, lease); err != nil {
		return err // Let the LeaderElector handle NotFound errors.
	}

	// Patch the spec to signal liveness and candidacy.
	patch := client.MergeFrom(lease.DeepCopy())
	lease.Spec.HolderIdentity = &ler.HolderIdentity
	lease.Spec.RenewTime = &metav1.MicroTime{Time: ler.RenewTime.Time}
	lease.Spec.LeaseDurationSeconds = int32Ptr(int32(ler.LeaseDurationSeconds))
	if err := mcl.client.Patch(ctx, lease, patch); err != nil {
		return fmt.Errorf("failed to patch MultiClusterLease spec for heartbeat: %w", err)
	}

	// After patching, get the updated generation that we need to see reflected in the status.
	if err := mcl.client.Get(ctx, client.ObjectKey{Namespace: mcl.leaseNS, Name: mcl.leaseName}, lease); err != nil {
		return fmt.Errorf("failed to get updated generation after patch: %w", err)
	}
	expectedGeneration := lease.Generation

	// Phase 2: Synchronous Confirmation
	// Poll the status to wait for the election controller to confirm global leadership.
	pollCtx, cancel := context.WithTimeout(ctx, time.Duration(ler.LeaseDurationSeconds)*time.Second)
	defer cancel()

	err := wait.PollUntilContextCancel(pollCtx, 1*time.Second, true, func(ctx context.Context) (bool, error) {
		var currentLease v1alpha1.MultiClusterLease
		if err := mcl.client.Get(ctx, client.ObjectKey{Namespace: mcl.leaseNS, Name: mcl.leaseName}, &currentLease); err != nil {
			return false, nil // Don't stop polling on transient errors
		}

		// Check that the controller has observed our specific spec update (or a newer one)
		// AND that it has confirmed our leadership.
		if currentLease.Status.ObservedGeneration != nil &&
			*currentLease.Status.ObservedGeneration >= expectedGeneration &&
			currentLease.Status.GlobalHolderIdentity != nil &&
			*currentLease.Status.GlobalHolderIdentity == ler.HolderIdentity {
			return true, nil // Success! Confirmed leadership.
		}
		return false, nil // Not the leader yet, or status is stale. Continue polling.
	})

	if err != nil {
		return fmt.Errorf("failed to confirm leadership in status: %w", err)
	}
	return nil // Successfully confirmed leadership.
}

// RecordEvent is a no-op for this implementation.
func (mcl *MultiClusterLeaseLock) RecordEvent(string) {}

// Describe is a human-readable description of the lock.
func (mcl *MultiClusterLeaseLock) Describe() string {
	return fmt.Sprintf("multiclusterlease/%s/%s", mcl.leaseNS, mcl.leaseName)
}

func (mcl *MultiClusterLeaseLock) leaseToRecord(lease *v1alpha1.MultiClusterLease) *resourcelock.LeaderElectionRecord {
	var leaderTransitions int
	if lease.Status.GlobalLeaseTransitions != nil {
		leaderTransitions = int(*lease.Status.GlobalLeaseTransitions)
	}

	var renewTime metav1.Time
	if lease.Status.GlobalRenewTime != nil {
		parsedTime, err := time.Parse(time.RFC3339, *lease.Status.GlobalRenewTime)
		if err == nil {
			renewTime = metav1.Time{Time: parsedTime}
		}
	}

	var acquireTime metav1.Time
	// Use creation timestamp as a proxy for acquire time of the first leader.
	if !lease.CreationTimestamp.IsZero() {
		acquireTime = lease.CreationTimestamp
	}

	var leaseDuration time.Duration
	if lease.Spec.LeaseDurationSeconds != nil {
		leaseDuration = time.Duration(*lease.Spec.LeaseDurationSeconds) * time.Second
	}

	holderIdentity := ""
	if lease.Status.GlobalHolderIdentity != nil {
		holderIdentity = *lease.Status.GlobalHolderIdentity
	}

	return &resourcelock.LeaderElectionRecord{
		HolderIdentity:       holderIdentity,
		LeaseDurationSeconds: int(leaseDuration.Seconds()),
		AcquireTime:          acquireTime,
		RenewTime:            renewTime,
		LeaderTransitions:    leaderTransitions,
	}
}

// Identity returns the identity of the lock holder.
func (mcl *MultiClusterLeaseLock) Identity() string {
	return mcl.identity
}

func int32Ptr(i int32) *int32 {
	return &i
}
