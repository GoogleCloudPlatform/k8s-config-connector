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
	"errors"
	"time"
)

// LeaseData represents the data stored in a lease object
type LeaseData struct {
	HolderIdentity   string    `json:"holderIdentity"`
	RenewTime        time.Time `json:"renewTime"`
	LeaseTransitions int32     `json:"leaseTransitions"`
}

// LeaseObject represents a lease object with its data and metadata
type LeaseObject struct {
	Data       LeaseData
	Generation int64 // Used for optimistic locking
}

// Storage defines the interface for lease storage backends
type Storage interface {
	// ReadLease reads the current lease data for the given key
	// Returns ErrNotFound if the lease doesn't exist
	ReadLease(ctx context.Context, key string) (*LeaseObject, error)

	// UpdateLease updates an existing lease with optimistic locking
	// The generation parameter is used for conditional updates
	// Returns ErrConditionalUpdateFailed if the generation doesn't match
	// Returns ErrNotFound if the lease doesn't exist
	UpdateLease(ctx context.Context, key string, data *LeaseData, generation int64) error

	// CreateLease creates a new lease object, failing if it already exists
	// Returns ErrAlreadyExists if the lease already exists
	CreateLease(ctx context.Context, key string, data *LeaseData) error

	// DeleteLease deletes a lease object
	// Returns ErrNotFound if the lease doesn't exist
	DeleteLease(ctx context.Context, key string) error
}

// Common storage errors
var (
	ErrNotFound                = errors.New("lease not found")
	ErrAlreadyExists           = errors.New("lease already exists")
	ErrConditionalUpdateFailed = errors.New("conditional update failed")
)

// IsNotFound checks if an error is a "not found" error
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsAlreadyExists checks if an error is an "already exists" error
func IsAlreadyExists(err error) bool {
	return errors.Is(err, ErrAlreadyExists)
}

// IsConditionalUpdateFailed checks if an error is a conditional update failure
func IsConditionalUpdateFailed(err error) bool {
	return errors.Is(err, ErrConditionalUpdateFailed)
}
