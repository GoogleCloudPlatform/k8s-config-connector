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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Model is the entry-point for our per-object reconcilers
type Model interface {
	// AdapterForObject builds an operation object for reconciling the object u.
	// If there are references, AdapterForObject should dereference them before returning (using reader)
	AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (Adapter, error)

	// AdapterForURL builds an operation object for exporting the object u.
	AdapterForURL(ctx context.Context, url string) (Adapter, error)
}

// SensitiveFieldModel is the entry-point for our per-object reconciler that
// handles CRD with sensitive fields.
type SensitiveFieldModel interface {
	// AdapterForObject builds an operation object for reconciling the object u.
	// If there are references, AdapterForObject should dereference them before returning (using reader)
	AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (Adapter, error)

	// AdapterForURL builds an operation object for exporting the object u.
	AdapterForURL(ctx context.Context, url string) (Adapter, error)

	MapSecretToResources(ctx context.Context, reader client.Reader, secret corev1.Secret) ([]reconcile.Request, error)
}

// Adapter performs a single reconciliation on a single object.
// It is built using AdapterForObject.
type Adapter interface {
	// Delete removes the GCP object.
	// This can be called without calling Find.
	// It returns (true, nil) if the object was deleted,
	// and (false, nil) if the object was not found but should be presumed deleted.
	// In an error, the state is not fully determined - a delete might be in progress.
	Delete(ctx context.Context, op *DeleteOperation) (deleted bool, err error)

	// Find must be called as the first operation (unless we are deleting).
	// It returns whether the corresponding GCP object was found.
	Find(ctx context.Context) (found bool, err error)

	// Create creates a new GCP object.
	// This should only be called when Find has previously returned false.
	// The implementation should write the updated status into `u`.
	Create(ctx context.Context, op *CreateOperation) error

	// Update updates an existing GCP object.
	// This should only be called when Find has previously returned true.
	// The implementation should write the updated status into `u`.
	Update(ctx context.Context, op *UpdateOperation) error

	// Export fetches the cloud provider's representation of the object
	// as an unstructured.Unstructured.
	// Assumes Find has previously returned true.
	Export(ctx context.Context) (*unstructured.Unstructured, error)
}
