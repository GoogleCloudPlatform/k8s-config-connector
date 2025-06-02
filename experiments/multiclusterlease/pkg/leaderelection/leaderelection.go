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
)

// LeaderElector handles the actual leader election logic with GCS
type LeaderElector struct {
	client     *storage.Client
	bucketName string
	identity   string
	leaseKey   string
	lease      *v1alpha1.MultiClusterLease

	leaseTransitions int32
}

// NewLeaderElector creates a new LeaderElector instance
func NewLeaderElector(client *storage.Client, bucketName, identity, leaseKey string, lease *v1alpha1.MultiClusterLease) *LeaderElector {
	return &LeaderElector{
		client:     client,
		bucketName: bucketName,
		identity:   identity,
		leaseKey:   leaseKey,
		lease:      lease,
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

// AcquireOrRenew attempts to acquire or renew the lease in GCS
func (le *LeaderElector) AcquireOrRenew(ctx context.Context) (*LeaseInfo, error) {
	// Try to get the current lease state
	data, err := le.readLease(ctx)
	if err == storage.ErrObjectNotExist {
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
	leaseExpired := data.HolderIdentity == "" || time.Since(data.RenewTime) > totalLeaseTime

	// If we're the current holder or the lease is expired, try to acquire/renew
	if data.HolderIdentity == le.identity || leaseExpired {
		// If this is a new leader, increment transitions
		if data.HolderIdentity != le.identity && data.HolderIdentity != "" {
			data.LeaseTransitions++
		}

		// Update the lease
		now := time.Now()
		updatedLeaseData := &data
		updatedLeaseData.HolderIdentity = le.identity
		updatedLeaseData.RenewTime = now

		if err := le.updateLease(ctx, updatedLeaseData, data.Generation); err != nil {
			if isGCSPreconditionError(err) {
				// Someone else modified the object since we read it
				return &LeaseInfo{
					Acquired:         false,
					HolderIdentity:   &data.HolderIdentity,
					RenewTime:        &data.RenewTime,
					LeaseTransitions: &data.LeaseTransitions,
				}, nil
			}
			return &LeaseInfo{
				Acquired:         false,
				HolderIdentity:   &data.HolderIdentity,
				RenewTime:        &data.RenewTime,
				LeaseTransitions: &data.LeaseTransitions,
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

	// We're not the leader
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
func (le *LeaderElector) createNewLease(ctx context.Context) (*LeaseInfo, error) {
	obj := le.getObjectHandle()
	now := time.Now()
	leaseTransitions := int32(0)

	// Create the lease data
	data := leaseData{
		HolderIdentity:   le.identity,
		RenewTime:        now,
		LeaseTransitions: leaseTransitions,
	}

	// Marshal the lease data to JSON
	content, err := json.Marshal(data)
	if err != nil {
		return &LeaseInfo{Acquired: false}, fmt.Errorf("marshalling new lease data: %w", err)
	}

	// Use a precondition that the object doesn't exist
	w := obj.If(storage.Conditions{DoesNotExist: true}).NewWriter(ctx)
	w.ContentType = "application/json"
	w.CacheControl = "no-cache, no-store, must-revalidate"

	// Write the JSON data
	if _, err := w.Write(content); err != nil {
		w.Close()
		if isGCSPreconditionError(err) {
			// Someone else created the object since we checked
			return &LeaseInfo{Acquired: false}, nil
		}
		return &LeaseInfo{Acquired: false}, fmt.Errorf("writing new lease data: %w", err)
	}

	// Close the writer to complete the operation
	if err := w.Close(); err != nil {
		if isGCSPreconditionError(err) {
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

// isGCSPreconditionError checks if an error is a GCS precondition failure
func isGCSPreconditionError(err error) bool {
	var gErr *googleapi.Error
	if errors.As(err, &gErr) {
		return gErr.Code == http.StatusPreconditionFailed || gErr.Code == http.StatusConflict
	}
	return false
}
