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
	"strconv"
	"sync"
)

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
)

type FakeStorage struct {
	mu     sync.RWMutex
	leases map[string]*LeaseObject
}

func NewFakeStorage() *FakeStorage {
	return &FakeStorage{
		leases: make(map[string]*LeaseObject),
	}
}

func (s *FakeStorage) ReadLease(ctx context.Context, key string) (*LeaseObject, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	lease, ok := s.leases[key]
	if !ok {
		return nil, ErrNotFound
	}

	// Return a copy
	return &LeaseObject{
		Data:            lease.Data,
		ResourceVersion: lease.ResourceVersion,
	}, nil
}

func (s *FakeStorage) CreateLease(ctx context.Context, key string, data LeaseData) (*LeaseObject, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.leases[key]; ok {
		return nil, ErrConflict
	}

	lease := &LeaseObject{
		Data:            data,
		ResourceVersion: "1",
	}
	s.leases[key] = lease

	return &LeaseObject{
		Data:            lease.Data,
		ResourceVersion: lease.ResourceVersion,
	}, nil
}

func (s *FakeStorage) UpdateLease(ctx context.Context, key string, obj *LeaseObject) (*LeaseObject, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	current, ok := s.leases[key]
	if !ok {
		return nil, ErrNotFound
	}

	if current.ResourceVersion != obj.ResourceVersion {
		return nil, ErrConflict
	}

	rv, _ := strconv.Atoi(current.ResourceVersion)
	newRV := strconv.Itoa(rv + 1)

	lease := &LeaseObject{
		Data:            obj.Data,
		ResourceVersion: newRV,
	}
	s.leases[key] = lease

	return &LeaseObject{
		Data:            lease.Data,
		ResourceVersion: lease.ResourceVersion,
	}, nil
}

func (s *FakeStorage) DeleteLease(ctx context.Context, key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.leases, key)
	return nil
}

func (s *FakeStorage) IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func (s *FakeStorage) IsConflict(err error) bool {
	return errors.Is(err, ErrConflict)
}
