// Copyright 2022 Google LLC
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

package storage

import (
	"context"
	"sort"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// InMemoryStorage is a memory-backed (non-persistent) implementation of Storage, useful for tests.
type InMemoryStorage struct {
	mutex  sync.Mutex
	byType map[protoreflect.FullName]*typeStorage
}

// typeStorage stores objects of a given type
type typeStorage struct {
	mutex          sync.Mutex
	objectTypeName string
	byKey          map[string]proto.Message
}

var _ Storage = &InMemoryStorage{}

// NewInMemoryStorage constructs an InMemoryStorage
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		byType: make(map[protoreflect.FullName]*typeStorage),
	}
}

func (s *InMemoryStorage) getTypeStorage(name protoreflect.FullName) *typeStorage {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	ts := s.byType[name]
	if ts == nil {
		objectTypeName := string(name.Name())
		objectTypeName = strings.ToLower(string(objectTypeName[0])) + objectTypeName[1:]
		ts = &typeStorage{
			objectTypeName: objectTypeName,
			byKey:          make(map[string]protoreflect.ProtoMessage),
		}
		s.byType[name] = ts
	}
	return ts
}

// Create stores the object, erroring if it already exists
func (s *InMemoryStorage) Create(ctx context.Context, fqn string, create proto.Message) error {
	return s.getTypeStorage(create.ProtoReflect().Descriptor().FullName()).Create(ctx, fqn, create)
}

func (s *typeStorage) Create(ctx context.Context, fqn string, create proto.Message) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, found := s.byKey[fqn]
	if found {
		return status.Errorf(codes.AlreadyExists, "%v %q already exists", s.objectTypeName, fqn)
	}
	s.byKey[fqn] = proto.Clone(create)
	return nil
}

// Delete deletes the object, returning a not found error if it does not exist.
func (s *InMemoryStorage) Delete(ctx context.Context, fqn string, dest proto.Message) error {
	kind := dest.ProtoReflect().Descriptor()
	return s.getTypeStorage(kind.FullName()).Delete(ctx, fqn, dest)
}

func (s *typeStorage) Delete(ctx context.Context, fqn string, dest proto.Message) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	existing, found := s.byKey[fqn]
	if !found {
		return status.Errorf(codes.NotFound, "%v %q not found", s.objectTypeName, fqn)
	}
	proto.Merge(dest, existing)
	delete(s.byKey, fqn)
	return nil
}

// Update stores a new version of an object, erroring if it does not already exist
func (s *InMemoryStorage) Update(ctx context.Context, fqn string, update proto.Message) error {
	return s.getTypeStorage(update.ProtoReflect().Descriptor().FullName()).Update(ctx, fqn, update)
}

// Update stores a new version of an object, erroring if it does not already exist
func (s *typeStorage) Update(ctx context.Context, fqn string, update proto.Message) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, found := s.byKey[fqn]
	if !found {
		return status.Errorf(codes.NotFound, "%v %q not found", s.objectTypeName, fqn)
	}
	s.byKey[fqn] = proto.Clone(update)
	return nil
}

// Get returns an existing object
func (s *InMemoryStorage) Get(ctx context.Context, fqn string, dest proto.Message) error {
	return s.getTypeStorage(dest.ProtoReflect().Descriptor().FullName()).Get(ctx, fqn, dest)
}

func (s *typeStorage) Get(ctx context.Context, fqn string, dest proto.Message) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	existing, found := s.byKey[fqn]
	if !found {
		return status.Errorf(codes.NotFound, "%v %q not found", s.objectTypeName, fqn)
	}
	proto.Merge(dest, existing)
	return nil
}

// List returns all matching objects
func (s *InMemoryStorage) List(ctx context.Context, kind protoreflect.Descriptor, options ListOptions, callback func(obj proto.Message) error) error {
	return s.getTypeStorage(kind.FullName()).List(ctx, options, callback)
}

func (s *typeStorage) List(ctx context.Context, options ListOptions, callback func(obj proto.Message) error) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var keys []string
	for fqn := range s.byKey {
		keys = append(keys, fqn)
	}
	sort.Strings(keys)

	for _, fqn := range keys {
		obj := s.byKey[fqn]
		if options.Prefix != "" && !strings.HasPrefix(fqn, options.Prefix) {
			continue
		}
		// Technically we should clone here
		if err := callback(obj); err != nil {
			return err
		}
	}
	return nil
}
