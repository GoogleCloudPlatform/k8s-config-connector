// Copyright 2024 Google LLC
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

package directbase

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Model is the entry-point for our per-object reconcilers
type Model interface {
	// AdapterForObject builds an operation object for reconciling the object u
	AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (Adapter, error)
}

// Adapter performs a single reconciliation on a single object.
// It is built using AdapterForObject.
type Adapter interface {
	// Delete removes the GCP object.
	// This can be called without calling Find.
	// It returns (true, nil) if the object was deleted,
	// and (false, nil) if the object was not found but should be presumed deleted.
	// In an error, the state is not fully determined - a delete might be in progress.
	Delete(ctx context.Context) (deleted bool, err error)

	// Find must be called as the first operation (unless we are deleting).
	// It returns whether the corresponding GCP object was found.
	Find(ctx context.Context) (found bool, err error)

	// Create creates a new GCP object.
	// This should only be called when Find has previously returned false.
	Create(ctx context.Context, u *unstructured.Unstructured) error

	// Update updates an existing GCP object.
	// This should only be called when Find has previously returned true.
	Update(ctx context.Context) (*unstructured.Unstructured, error)
}
