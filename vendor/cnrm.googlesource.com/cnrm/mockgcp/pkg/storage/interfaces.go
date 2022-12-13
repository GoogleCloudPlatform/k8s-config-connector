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

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Storage interface {
	// Create stores the object, erroring if it already exists
	Create(ctx context.Context, fqn string, create proto.Message) error
	// Update stores a new version of an object, erroring if it does not already exist
	Update(ctx context.Context, fqn string, update proto.Message) error
	// Get returns an existing object
	Get(ctx context.Context, fqn string, dest proto.Message) error
	// List returns all matching objects
	List(ctx context.Context, kind protoreflect.Descriptor, options ListOptions, callback func(obj proto.Message) error) error
	// Deleting deletes the object, returning a not found error if it does not exist.
	Delete(ctx context.Context, kind protoreflect.Descriptor, fqn string) error
}

// ListOptions restricts the objects returned by a List
type ListOptions struct {
	// Prefix ensures that only objects whose key matches the prefix are returned
	Prefix string
}
