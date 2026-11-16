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

package storage

import (
	"context"
	"time"
)

// LeaseData represents the data stored in the global lock backend.
type LeaseData struct {
	HolderIdentity   string    `json:"holderIdentity"`
	RenewTime        time.Time `json:"renewTime"`
	LeaseTransitions int32     `json:"leaseTransitions"`
}

// LeaseObject represents a lease entry in the storage backend, including metadata for concurrency control.
type LeaseObject struct {
	Data LeaseData
	// ResourceVersion is used for optimistic locking.
	// For GCS, this is the object's generation.
	ResourceVersion string
}

// Storage is the interface for interacting with the global lock backend.
type Storage interface {
	// ReadLease reads the current lease data for the given key.
	// Returns a LeaseObject, or an error. If the lease does not exist,
	// it should return an error that can be identified by IsNotFound(err).
	ReadLease(ctx context.Context, key string) (*LeaseObject, error)

	// CreateLease creates a new lease object in the storage backend.
	// Returns the newly created LeaseObject, or an error. If the lease already
	// exists, it should return an error that can be identified by IsConflict(err).
	CreateLease(ctx context.Context, key string, data LeaseData) (*LeaseObject, error)

	// UpdateLease updates an existing lease object in the storage backend.
	// It should use the ResourceVersion from the provided LeaseObject to perform
	// optimistic locking. Returns the updated LeaseObject, or an error.
	// If the ResourceVersion does not match, it should return an error that
	// can be identified by IsConflict(err).
	UpdateLease(ctx context.Context, key string, obj *LeaseObject) (*LeaseObject, error)

	// DeleteLease deletes the lease object for the given key.
	DeleteLease(ctx context.Context, key string) error

	// IsNotFound returns true if the error indicates that the lease object was not found.
	IsNotFound(err error) bool

	// IsConflict returns true if the error indicates a concurrency conflict
	// (e.g., failed optimistic locking).
	IsConflict(err error) bool
}
