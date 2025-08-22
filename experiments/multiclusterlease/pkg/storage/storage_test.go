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
	"sync"
	"testing"
	"time"
)

// TestStorageInterface validates that both storage implementations work correctly
func TestStorageInterface(t *testing.T) {
	// Test both implementations
	testCases := []struct {
		name    string
		storage Storage
	}{
		{
			name:    "MemoryStorage",
			storage: NewMemoryStorage(),
		},
		// Note: GCS storage would require actual GCS credentials and bucket
		// It should be tested separately in integration tests
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testStorageOperations(t, tc.storage)
		})
	}
}

func testStorageOperations(t *testing.T, storage Storage) {
	ctx := context.Background()
	key := "test-lease"

	// Test reading non-existent lease
	_, err := storage.ReadLease(ctx, key)
	if !IsNotFound(err) {
		t.Errorf("Expected ErrNotFound for non-existent lease, got: %v", err)
	}

	// Test creating a new lease
	leaseData := &LeaseData{
		HolderIdentity:   "test-identity",
		RenewTime:        time.Now(),
		LeaseTransitions: 0,
	}

	err = storage.CreateLease(ctx, key, leaseData)
	if err != nil {
		t.Errorf("Failed to create lease: %v", err)
	}

	// Test creating duplicate lease
	err = storage.CreateLease(ctx, key, leaseData)
	if !IsAlreadyExists(err) {
		t.Errorf("Expected ErrAlreadyExists for duplicate lease, got: %v", err)
	}

	// Test reading the lease
	leaseObj, err := storage.ReadLease(ctx, key)
	if err != nil {
		t.Errorf("Failed to read lease: %v", err)
	}

	if leaseObj.Data.HolderIdentity != leaseData.HolderIdentity {
		t.Errorf("Expected holder identity %s, got %s", leaseData.HolderIdentity, leaseObj.Data.HolderIdentity)
	}

	// Test updating the lease
	updatedData := leaseObj.Data
	updatedData.HolderIdentity = "new-identity"
	updatedData.LeaseTransitions = 1

	err = storage.UpdateLease(ctx, key, &updatedData, leaseObj.Generation)
	if err != nil {
		t.Errorf("Failed to update lease: %v", err)
	}

	// Test conditional update failure
	err = storage.UpdateLease(ctx, key, &updatedData, leaseObj.Generation) // Using old generation
	if !IsConditionalUpdateFailed(err) {
		t.Errorf("Expected ErrConditionalUpdateFailed for stale generation, got: %v", err)
	}

	// Test deleting the lease
	err = storage.DeleteLease(ctx, key)
	if err != nil {
		t.Errorf("Failed to delete lease: %v", err)
	}

	// Test deleting non-existent lease
	err = storage.DeleteLease(ctx, key)
	if !IsNotFound(err) {
		t.Errorf("Expected ErrNotFound for deleting non-existent lease, got: %v", err)
	}
}

// MemoryStorage implements the Storage interface using in-memory storage
// This is used for testing purposes only
type MemoryStorage struct {
	mu     sync.RWMutex
	leases map[string]*LeaseObject
}

// NewMemoryStorage creates a new in-memory storage implementation
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		leases: make(map[string]*LeaseObject),
	}
}

// ReadLease reads the current lease data for the given key from memory
func (m *MemoryStorage) ReadLease(ctx context.Context, key string) (*LeaseObject, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	lease, exists := m.leases[key]
	if !exists {
		return nil, ErrNotFound
	}

	// Return a copy to prevent external modifications
	return &LeaseObject{
		Data:       lease.Data,
		Generation: lease.Generation,
	}, nil
}

// UpdateLease updates an existing lease with optimistic locking
func (m *MemoryStorage) UpdateLease(ctx context.Context, key string, data *LeaseData, generation int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	existingLease, exists := m.leases[key]
	if !exists {
		return ErrNotFound
	}

	// Check generation for optimistic locking
	if existingLease.Generation != generation {
		return ErrConditionalUpdateFailed
	}

	// Update the lease with incremented generation
	m.leases[key] = &LeaseObject{
		Data:       *data,
		Generation: generation + 1,
	}

	return nil
}

// CreateLease creates a new lease object, failing if it already exists
func (m *MemoryStorage) CreateLease(ctx context.Context, key string, data *LeaseData) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.leases[key]; exists {
		return ErrAlreadyExists
	}

	m.leases[key] = &LeaseObject{
		Data:       *data,
		Generation: 1,
	}

	return nil
}

// DeleteLease deletes a lease object
func (m *MemoryStorage) DeleteLease(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.leases[key]; !exists {
		return ErrNotFound
	}

	delete(m.leases, key)
	return nil
}
