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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/api/v1alpha1"
	"google.golang.org/api/googleapi"
	ctrl "sigs.k8s.io/controller-runtime"
)

// LeaderElector handles the actual leader election logic with GCS
type LeaderElector struct {
	client     *storage.Client
	bucketName string
	leaseKey   string
}

// NewLeaderElector creates a new LeaderElector instance
func NewLeaderElector(client *storage.Client, bucketName, leaseKey string) *LeaderElector {
	return &LeaderElector{
		client:     client,
		bucketName: bucketName,
		leaseKey:   leaseKey,
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

// AcquireOrRenew attempts to acquire or renew the lease in GCS. It is designed
// to be robust against race conditions.
func (le *LeaderElector) AcquireOrRenew(ctx context.Context, lease *v1alpha1.MultiClusterLease, identity string) (*LeaseInfo, error) {
	log := ctrl.Log.WithName("leaderelector").WithValues("leaseKey", le.leaseKey, "candidate", identity)

	// Check if the candidate is alive by checking its renewTime.
	if lease.Spec.RenewTime == nil {
		log.Info("candidate has no renew time")
		return &LeaseInfo{Acquired: false}, fmt.Errorf("candidate has no renew time")
	}

	// TODO: Make the staleness check configurable
	if time.Since(lease.Spec.RenewTime.Time) > 15*time.Second {
		log.Info("candidate lease is stale")

		// still need to return latest lease data
		data, err := le.readLease(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to read lease: %w", err)
		}
		return &LeaseInfo{
			Acquired:         false,
			HolderIdentity:   &data.HolderIdentity,
			RenewTime:        &data.RenewTime,
			LeaseTransitions: &data.LeaseTransitions,
		}, fmt.Errorf("candidate lease is stale")
	}

	// 1. Try to read the existing lease from GCS.
	log.Info("reading lease from GCS")
	data, err := le.readLease(ctx)
	if err != nil {
		if !errors.Is(err, storage.ErrObjectNotExist) {
			log.Error(err, "failed to read lease from GCS")
			return nil, fmt.Errorf("failed to read lease from GCS: %w", err)
		}

		// 2. The lease does not exist. Try to create it.
		log.Info("lease object does not exist in GCS, attempting to create")
		info, createErr := le.createNewLease(ctx, identity)
		if createErr == nil {
			// Success! We are the first and only leader.
			log.Info("successfully created new lease")
			return info, nil
		}

		if !isGCSPreconditionError(createErr) {
			// This was a non-recoverable error.
			log.Error(createErr, "failed to create new lease")
			return nil, fmt.Errorf("failed to create new lease: %w", createErr)
		}

		// We lost the creation race. The object now exists.
		log.Info("lost creation race, re-reading lease")
		data, err = le.readLease(ctx)
		if err != nil {
			log.Error(err, "failed to read lease after losing creation race")
			return nil, fmt.Errorf("failed to read lease after losing creation race: %w", err)
		}
	}

	// 3. The lease exists. Check if we can acquire it.
	log.Info("successfully read lease from GCS", "holder", data.HolderIdentity, "renewTime", data.RenewTime, "generation", data.Generation)
	leaseDuration := time.Duration(*lease.Spec.LeaseDurationSeconds) * time.Second
	leaseExpired := time.Since(data.RenewTime) > leaseDuration
	log.Info("checking lease expiration", "isExpired", leaseExpired)

	if data.HolderIdentity == identity || leaseExpired {
		// We are the holder or the lease is expired. Try to update.
		log.Info("attempting to acquire or renew lease")
		if data.HolderIdentity != identity && data.HolderIdentity != "" {
			data.LeaseTransitions++
			log.Info("incrementing lease transitions", "newTransitions", data.LeaseTransitions)
		}
		data.HolderIdentity = identity
		data.RenewTime = time.Now()

		updateErr := le.updateLease(ctx, &data, data.Generation)
		if updateErr != nil {
			if isGCSPreconditionError(updateErr) {
				// We lost an update race. Re-read to get the latest state.
				log.Info("lost update race, re-reading lease")
				data, err = le.readLease(ctx)
				if err != nil {
					log.Error(err, "failed to read lease after losing update race")
					return nil, fmt.Errorf("failed to read lease after losing update race: %w", err)
				}
				return &LeaseInfo{
					Acquired:         false,
					HolderIdentity:   &data.HolderIdentity,
					RenewTime:        &data.RenewTime,
					LeaseTransitions: &data.LeaseTransitions,
				}, nil
			}
			log.Error(updateErr, "failed to update lease")
			return nil, fmt.Errorf("failed to update lease: %w", updateErr)
		}

		// Successfully updated. We are the leader.
		log.Info("successfully acquired or renewed lease")
		// After a successful update, re-read the lease to get the authoritative state.
		data, err = le.readLease(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to read lease after successful update: %w", err)
		}
		return &LeaseInfo{
			Acquired:         true,
			HolderIdentity:   &data.HolderIdentity,
			RenewTime:        &data.RenewTime,
			LeaseTransitions: &data.LeaseTransitions,
		}, nil
	}

	// 4. The lease is held by someone else and is not expired.
	log.Info("lease is held by another identity and is not expired")
	return &LeaseInfo{
		Acquired:         false,
		HolderIdentity:   &data.HolderIdentity,
		RenewTime:        &data.RenewTime,
		LeaseTransitions: &data.LeaseTransitions,
	}, nil
}

// leaseData represents the data stored in the GCS lease object
type leaseData struct {
	HolderIdentity   string    `json:"holderIdentity"`
	RenewTime        time.Time `json:"renewTime"`
	LeaseTransitions int32     `json:"leaseTransitions"`
	Generation       int64     `json:"-"` // Not stored in JSON
}

// getObjectHandle returns a handle to the GCS object for this lease
func (le *LeaderElector) getObjectHandle() *storage.ObjectHandle {
	bucket := le.client.Bucket(le.bucketName)
	return bucket.Object(fmt.Sprintf("leases/%s", le.leaseKey))
}

// readLease reads the current lease data from GCS
func (le *LeaderElector) readLease(ctx context.Context) (leaseData, error) {
	obj := le.getObjectHandle()
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return leaseData{}, err
	}

	data := leaseData{
		Generation: attrs.Generation,
	}

	// Read the object content
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return leaseData{}, err
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		return leaseData{}, fmt.Errorf("reading lease object: %w", err)
	}

	// If the object is empty, return default values
	if len(content) == 0 {
		return data, nil
	}

	// Unmarshal the JSON data
	if err := json.Unmarshal(content, &data); err != nil {
		return leaseData{}, fmt.Errorf("unmarshalling lease data: %w", err)
	}

	return data, nil
}

// updateLease updates the lease in GCS with optimistic locking
func (le *LeaderElector) updateLease(ctx context.Context, data *leaseData, generation int64) error {
	obj := le.getObjectHandle()

	// Marshal the lease data to JSON
	content, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshalling lease data: %w", err)
	}

	// Create a conditional writer
	w := obj.If(storage.Conditions{GenerationMatch: generation}).NewWriter(ctx)
	w.ContentType = "application/json"
	w.CacheControl = "no-cache, no-store, must-revalidate"

	// Write the JSON data
	if _, err := w.Write(content); err != nil {
		w.Close()
		if isGCSPreconditionError(err) {
			return fmt.Errorf("precondition failed while writing lease data: %w", err)
		}
		return fmt.Errorf("writing lease data: %w", err)
	}

	// Close the writer to complete the operation
	if err := w.Close(); err != nil {
		if isGCSPreconditionError(err) {
			return fmt.Errorf("precondition failed while closing writer: %w", err)
		}
		return fmt.Errorf("closing writer: %w", err)
	}

	return nil
}

// createNewLease creates a new lease object in GCS with this instance as the holder
// Uses a precondition to ensure the object doesn't exist (preventing race conditions)
func (le *LeaderElector) createNewLease(ctx context.Context, identity string) (*LeaseInfo, error) {
	obj := le.getObjectHandle()
	now := time.Now()
	leaseTransitions := int32(1) // Start with 1 transition

	// Create the lease data
	data := leaseData{
		HolderIdentity:   identity,
		RenewTime:        now,
		LeaseTransitions: leaseTransitions,
	}

	// Marshal the lease data to JSON
	content, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshalling new lease data: %w", err)
	}

	// Use a precondition that the object doesn't exist
	w := obj.If(storage.Conditions{DoesNotExist: true}).NewWriter(ctx)
	w.ContentType = "application/json"
	w.CacheControl = "no-cache, no-store, must-revalidate"

	// Write the JSON data
	if _, err := w.Write(content); err != nil {
		w.Close() // Close the writer even on error
		return nil, err
	}

	// Close the writer to complete the operation
	if err := w.Close(); err != nil {
		return nil, err
	}

	// Successfully created and acquired the lease
	return &LeaseInfo{
		Acquired:         true,
		HolderIdentity:   &identity,
		RenewTime:        &now,
		LeaseTransitions: &leaseTransitions,
	}, nil
}

// isGCSPreconditionError checks if an error is a GCS precondition failure
func isGCSPreconditionError(err error) bool {
	var gErr *googleapi.Error
	if errors.As(err, &gErr) {
		return gErr.Code == http.StatusPreconditionFailed || gErr.Code == http.StatusConflict
	}
	return false
}
