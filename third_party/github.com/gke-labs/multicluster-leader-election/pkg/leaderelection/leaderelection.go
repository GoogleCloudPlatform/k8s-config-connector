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

	v1alpha1 "github.com/gke-labs/multicluster-leader-election/api/v1alpha1"
	"github.com/gke-labs/multicluster-leader-election/pkg/storage"
	ctrl "sigs.k8s.io/controller-runtime"
)

// LeaderElector handles the actual leader election logic with a storage backend
type LeaderElector struct {
	storage  storage.Storage
	leaseKey string
}

// NewLeaderElector creates a new LeaderElector instance
func NewLeaderElector(s storage.Storage, leaseKey string) *LeaderElector {
	return &LeaderElector{
		storage:  s,
		leaseKey: leaseKey,
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

// AcquireOrRenew attempts to acquire or renew the lease in the storage backend. It is designed
// to be robust against race conditions.
func (le *LeaderElector) AcquireOrRenew(ctx context.Context, lease *v1alpha1.MultiClusterLease, identity string) (*LeaseInfo, error) {
	log := ctrl.Log.WithName("leaderelector").WithValues("leaseKey", le.leaseKey, "candidate", identity)

	// Check if the candidate is alive by checking its renewTime.
	if lease.Spec.RenewTime == nil {
		log.Info("candidate has no renew time")
		return &LeaseInfo{Acquired: false}, fmt.Errorf("candidate has no renew time")
	}

	// We check if the candidate's heartbeat to its local cluster is fresh.
	// If the candidate pod is dead or not heartbeating locally, we don't
	// want to contend for the global lock on its behalf.
	stalenessThreshold := 15 * time.Second
	if lease.Spec.LeaseDurationSeconds != nil {
		stalenessThreshold = time.Duration(*lease.Spec.LeaseDurationSeconds) * time.Second * 2
	}

	if time.Since(lease.Spec.RenewTime.Time) > stalenessThreshold {
		log.Info("candidate lease is stale", "threshold", stalenessThreshold)

		// still need to return latest lease data
		obj, err := le.storage.ReadLease(ctx, le.leaseKey)
		if err != nil {
			return nil, fmt.Errorf("failed to read lease: %w", err)
		}
		return &LeaseInfo{
			Acquired:         false,
			HolderIdentity:   &obj.Data.HolderIdentity,
			RenewTime:        &obj.Data.RenewTime,
			LeaseTransitions: &obj.Data.LeaseTransitions,
		}, fmt.Errorf("candidate lease is stale")
	}

	// 1. Try to read the existing lease from storage.
	log.Info("reading lease from storage")
	obj, err := le.storage.ReadLease(ctx, le.leaseKey)
	if err != nil {
		if !le.storage.IsNotFound(err) {
			log.Error(err, "failed to read lease from storage")
			return nil, fmt.Errorf("failed to read lease from storage: %w", err)
		}

		// 2. The lease does not exist. Try to create it.
		log.Info("lease object does not exist in storage, attempting to create")
		now := time.Now()
		data := storage.LeaseData{
			HolderIdentity:   identity,
			RenewTime:        now,
			LeaseTransitions: 1,
		}
		obj, err = le.storage.CreateLease(ctx, le.leaseKey, data)
		if err == nil {
			// Success! We are the first and only leader.
			log.Info("successfully created new lease")
			return &LeaseInfo{
				Acquired:         true,
				HolderIdentity:   &obj.Data.HolderIdentity,
				RenewTime:        &obj.Data.RenewTime,
				LeaseTransitions: &obj.Data.LeaseTransitions,
			}, nil
		}

		if !le.storage.IsConflict(err) {
			// This was a non-recoverable error.
			log.Error(err, "failed to create new lease")
			return nil, fmt.Errorf("failed to create new lease: %w", err)
		}

		// We lost the creation race. The object now exists.
		log.Info("lost creation race, re-reading lease")
		obj, err = le.storage.ReadLease(ctx, le.leaseKey)
		if err != nil {
			log.Error(err, "failed to read lease after losing creation race")
			return nil, fmt.Errorf("failed to read lease after losing creation race: %w", err)
		}
	}

	// 3. The lease exists. Check if we can acquire it.
	log.Info("successfully read lease from storage", "holder", obj.Data.HolderIdentity, "renewTime", obj.Data.RenewTime, "resourceVersion", obj.ResourceVersion)
	leaseDuration := 15 * time.Second
	if lease.Spec.LeaseDurationSeconds != nil {
		leaseDuration = time.Duration(*lease.Spec.LeaseDurationSeconds) * time.Second
	}
	leaseExpired := time.Since(obj.Data.RenewTime) > leaseDuration
	log.Info("checking lease expiration", "isExpired", leaseExpired)

	if obj.Data.HolderIdentity != identity && !leaseExpired {
		// 4. The lease is held by someone else and is not expired.
		log.Info("lease is held by another identity and is not expired", "holder", obj.Data.HolderIdentity)
		return &LeaseInfo{
			Acquired:         false,
			HolderIdentity:   &obj.Data.HolderIdentity,
			RenewTime:        &obj.Data.RenewTime,
			LeaseTransitions: &obj.Data.LeaseTransitions,
		}, nil
	}

	// We are the holder or the lease is expired. Try to update.
	log.Info("attempting to acquire or renew lease")
	if obj.Data.HolderIdentity != identity && obj.Data.HolderIdentity != "" {
		obj.Data.LeaseTransitions++
		log.Info("incrementing lease transitions", "newTransitions", obj.Data.LeaseTransitions)
	}
	obj.Data.HolderIdentity = identity
	obj.Data.RenewTime = time.Now()

	updatedObj, updateErr := le.storage.UpdateLease(ctx, le.leaseKey, obj)
	if updateErr != nil {
		if le.storage.IsConflict(updateErr) {
			// We lost an update race. Re-read to get the latest state.
			log.Info("lost update race, re-reading lease")
			obj, err = le.storage.ReadLease(ctx, le.leaseKey)
			if err != nil {
				log.Error(err, "failed to read lease after losing update race")
				return nil, fmt.Errorf("failed to read lease after losing update race: %w", err)
			}
			return &LeaseInfo{
				Acquired:         false,
				HolderIdentity:   &obj.Data.HolderIdentity,
				RenewTime:        &obj.Data.RenewTime,
				LeaseTransitions: &obj.Data.LeaseTransitions,
			}, nil
		}
		log.Error(updateErr, "failed to update lease")
		return nil, fmt.Errorf("failed to update lease: %w", updateErr)
	}

	// Successfully updated. We are the leader.
	log.Info("successfully acquired or renewed lease")
	return &LeaseInfo{
		Acquired:         true,
		HolderIdentity:   &updatedObj.Data.HolderIdentity,
		RenewTime:        &updatedObj.Data.RenewTime,
		LeaseTransitions: &updatedObj.Data.LeaseTransitions,
	}, nil
}
