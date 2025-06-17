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

package leaderelection

import (
	"context"
	"fmt"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/api/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/pkg/storage"
)

// LeaderElector handles the actual leader election logic with pluggable storage
type LeaderElector struct {
	storage  storage.Storage
	identity string
	leaseKey string
	lease    *v1alpha1.MultiClusterLease

	leaseTransitions int32
}

// NewLeaderElector creates a new LeaderElector instance
func NewLeaderElector(storage storage.Storage, identity, leaseKey string, lease *v1alpha1.MultiClusterLease) *LeaderElector {
	return &LeaderElector{
		storage:  storage,
		identity: identity,
		leaseKey: leaseKey,
		lease:    lease,
	}
}

// LeaseInfo contains the current state of a lease from the backend
type LeaseInfo struct {
	// Acquired indicates whether this identity successfully acquired or renewed the lease
	Acquired bool
	// HolderIdentity is the identity of the current lease holder (may be empty if no holder)
	HolderIdentity *string
	// RenewTime is when the lease was last renewed (may be nil if no lease exists)
	RenewTime *time.Time
	// LeaseTransitions is the number of times leadership has changed (may be nil if no lease exists)
	LeaseTransitions *int32
}

// AcquireOrRenew attempts to acquire or renew the lease using the configured storage backend
func (le *LeaderElector) AcquireOrRenew(ctx context.Context) (*LeaseInfo, error) {
	// Try to get the current lease state
	leaseObj, err := le.storage.ReadLease(ctx, le.leaseKey)
	if storage.IsNotFound(err) {
		// This is the first time anyone is trying to acquire this lease
		return le.createNewLease(ctx)
	}
	if err != nil {
		return &LeaseInfo{Acquired: false}, err
	}

	// Check if the lease is expired
	leaseDuration := le.lease.Spec.GetLeaseDuration()
	gracePeriod := le.lease.Spec.GetGracePeriod()
	totalLeaseTime := leaseDuration + gracePeriod
	leaseExpired := leaseObj.Data.HolderIdentity == "" || time.Since(leaseObj.Data.RenewTime) > totalLeaseTime

	// We're not the leader
	if leaseObj.Data.HolderIdentity != le.identity && !leaseExpired {
		return &LeaseInfo{
			Acquired:         false,
			HolderIdentity:   &leaseObj.Data.HolderIdentity,
			RenewTime:        &leaseObj.Data.RenewTime,
			LeaseTransitions: &leaseObj.Data.LeaseTransitions,
		}, nil
	}

	// If we're the current holder or the lease is expired, try to acquire/renew
	now := time.Now()
	updatedLeaseData := leaseObj.Data

	if leaseObj.Data.HolderIdentity != le.identity && leaseObj.Data.HolderIdentity != "" {
		// If this is a new leader, increment transitions
		updatedLeaseData.LeaseTransitions++
	}

	// Update the lease
	updatedLeaseData.HolderIdentity = le.identity
	updatedLeaseData.RenewTime = now

	if err := le.storage.UpdateLease(ctx, le.leaseKey, &updatedLeaseData, leaseObj.Generation); err != nil {
		if storage.IsConditionalUpdateFailed(err) {
			// Someone else modified the object since we read it
			return &LeaseInfo{
				Acquired:         false,
				HolderIdentity:   &leaseObj.Data.HolderIdentity,
				RenewTime:        &leaseObj.Data.RenewTime,
				LeaseTransitions: &leaseObj.Data.LeaseTransitions,
			}, nil
		}
		return &LeaseInfo{
			Acquired:         false,
			HolderIdentity:   &leaseObj.Data.HolderIdentity,
			RenewTime:        &leaseObj.Data.RenewTime,
			LeaseTransitions: &leaseObj.Data.LeaseTransitions,
		}, fmt.Errorf("failed to update lease: %w", err)
	}

	// Successfully acquired/renewed the lease
	return &LeaseInfo{
		Acquired:         true,
		HolderIdentity:   &le.identity,
		RenewTime:        &now,
		LeaseTransitions: &updatedLeaseData.LeaseTransitions,
	}, nil
}

// createNewLease creates a new lease object with this instance as the holder
func (le *LeaderElector) createNewLease(ctx context.Context) (*LeaseInfo, error) {
	now := time.Now()
	leaseTransitions := int32(0)

	// Create the lease data
	data := storage.LeaseData{
		HolderIdentity:   le.identity,
		RenewTime:        now,
		LeaseTransitions: leaseTransitions,
	}

	if err := le.storage.CreateLease(ctx, le.leaseKey, &data); err != nil {
		if storage.IsAlreadyExists(err) {
			// Someone else created the object since we checked
			return &LeaseInfo{Acquired: false}, nil
		}
		return &LeaseInfo{Acquired: false}, fmt.Errorf("failed to create new lease: %w", err)
	}

	// Successfully created and acquired the lease
	return &LeaseInfo{
		Acquired:         true,
		HolderIdentity:   &le.identity,
		RenewTime:        &now,
		LeaseTransitions: &leaseTransitions,
	}, nil
}

// ReleaseLease releases the lease by setting a very short expiration time
func (le *LeaderElector) ReleaseLease(ctx context.Context) error {
	// Try to get the current lease state
	leaseObj, err := le.storage.ReadLease(ctx, le.leaseKey)
	if storage.IsNotFound(err) {
		// Lease doesn't exist, nothing to release
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to read lease for release: %w", err)
	}

	// Check if we're the current holder
	if leaseObj.Data.HolderIdentity != le.identity {
		// We're not the current holder, nothing to release
		return nil
	}

	// Set the renew time to 1 second to make the lease expire quickly
	expiredTime := time.Now().Add(1 * time.Second)
	updatedData := leaseObj.Data
	updatedData.RenewTime = expiredTime

	// Update the lease with the expired time
	if err := le.storage.UpdateLease(ctx, le.leaseKey, &updatedData, leaseObj.Generation); err != nil {
		if storage.IsConditionalUpdateFailed(err) {
			// Someone else modified the lease, which is fine
			return nil
		}
		return fmt.Errorf("failed to update lease with expired time: %w", err)
	}

	return nil
}
