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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/storage"
	"google.golang.org/api/googleapi"
)

// GCSStorage implements the Storage interface using Google Cloud Storage
type GCSStorage struct {
	client     *storage.Client
	bucketName string
}

// NewGCSStorage creates a new GCS storage implementation
func NewGCSStorage(client *storage.Client, bucketName string) *GCSStorage {
	return &GCSStorage{
		client:     client,
		bucketName: bucketName,
	}
}

// ReadLease reads the current lease data for the given key from GCS
func (g *GCSStorage) ReadLease(ctx context.Context, key string) (*LeaseObject, error) {
	obj := g.getObjectHandle(key)
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("getting object attributes: %w", err)
	}

	// Read the object content
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating object reader: %w", err)
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("reading object content: %w", err)
	}

	var leaseData LeaseData
	// If the object is empty, return default values
	if len(content) > 0 {
		if err := json.Unmarshal(content, &leaseData); err != nil {
			return nil, fmt.Errorf("unmarshalling lease data: %w", err)
		}
	}

	return &LeaseObject{
		Data:       leaseData,
		Generation: attrs.Generation,
	}, nil
}

// UpdateLease updates an existing lease with optimistic locking using GCS generation matching
func (g *GCSStorage) UpdateLease(ctx context.Context, key string, data *LeaseData, generation int64) error {
	obj := g.getObjectHandle(key)

	// Marshal the lease data to JSON
	content, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshalling lease data: %w", err)
	}

	// Create a conditional writer with generation matching
	w := obj.If(storage.Conditions{GenerationMatch: generation}).NewWriter(ctx)
	w.ContentType = "application/json"
	w.CacheControl = "no-cache, no-store, must-revalidate"

	// Write the JSON data
	if _, err := w.Write(content); err != nil {
		w.Close()
		if isGCSPreconditionError(err) {
			return ErrConditionalUpdateFailed
		}
		return fmt.Errorf("writing lease data: %w", err)
	}

	// Close the writer to complete the operation
	if err := w.Close(); err != nil {
		if isGCSPreconditionError(err) {
			return ErrConditionalUpdateFailed
		}
		return fmt.Errorf("closing writer: %w", err)
	}

	return nil
}

// CreateLease creates a new lease object in GCS, failing if it already exists
func (g *GCSStorage) CreateLease(ctx context.Context, key string, data *LeaseData) error {
	obj := g.getObjectHandle(key)

	// Marshal the lease data to JSON
	content, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshalling lease data: %w", err)
	}

	// Use a precondition that the object doesn't exist
	w := obj.If(storage.Conditions{DoesNotExist: true}).NewWriter(ctx)
	w.ContentType = "application/json"
	w.CacheControl = "no-cache, no-store, must-revalidate"

	// Write the JSON data
	if _, err := w.Write(content); err != nil {
		w.Close()
		if isGCSPreconditionError(err) {
			return ErrAlreadyExists
		}
		return fmt.Errorf("writing lease data: %w", err)
	}

	// Close the writer to complete the operation
	if err := w.Close(); err != nil {
		if isGCSPreconditionError(err) {
			return ErrAlreadyExists
		}
		return fmt.Errorf("closing writer: %w", err)
	}

	return nil
}

// DeleteLease deletes a lease object from GCS
func (g *GCSStorage) DeleteLease(ctx context.Context, key string) error {
	obj := g.getObjectHandle(key)

	if err := obj.Delete(ctx); err != nil {
		if err == storage.ErrObjectNotExist {
			return ErrNotFound
		}
		return fmt.Errorf("deleting lease object: %w", err)
	}

	return nil
}

// getObjectHandle returns a handle to the GCS object for the given lease key
func (g *GCSStorage) getObjectHandle(key string) *storage.ObjectHandle {
	bucket := g.client.Bucket(g.bucketName)
	return bucket.Object(fmt.Sprintf("leases/%s", key))
}

// isGCSPreconditionError checks if an error is a GCS precondition failure
func isGCSPreconditionError(err error) bool {
	var gErr *googleapi.Error
	if errors.As(err, &gErr) {
		return gErr.Code == http.StatusPreconditionFailed || gErr.Code == http.StatusConflict
	}
	return false
}
