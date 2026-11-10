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
	"strconv"

	gcs "cloud.google.com/go/storage"
	"google.golang.org/api/googleapi"
)

type GCSStorage struct {
	client     *gcs.Client
	bucketName string
}

func NewGCSStorage(client *gcs.Client, bucketName string) *GCSStorage {
	return &GCSStorage{
		client:     client,
		bucketName: bucketName,
	}
}

func (s *GCSStorage) getObjectHandle(key string) *gcs.ObjectHandle {
	bucket := s.client.Bucket(s.bucketName)
	return bucket.Object(fmt.Sprintf("leases/%s", key))
}

func (s *GCSStorage) ReadLease(ctx context.Context, key string) (*LeaseObject, error) {
	obj := s.getObjectHandle(key)
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return nil, err
	}

	// Read the object content
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("reading lease object: %w", err)
	}

	data := LeaseData{}
	if len(content) > 0 {
		if err := json.Unmarshal(content, &data); err != nil {
			return nil, fmt.Errorf("unmarshalling lease data: %w", err)
		}
	}

	return &LeaseObject{
		Data:            data,
		ResourceVersion: strconv.FormatInt(attrs.Generation, 10),
	}, nil
}

func (s *GCSStorage) CreateLease(ctx context.Context, key string, data LeaseData) (*LeaseObject, error) {
	obj := s.getObjectHandle(key)

	// Marshal the lease data to JSON
	content, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("marshalling new lease data: %w", err)
	}

	// Use a precondition that the object doesn't exist
	w := obj.If(gcs.Conditions{DoesNotExist: true}).NewWriter(ctx)
	w.ContentType = "application/json"
	w.CacheControl = "no-cache, no-store, must-revalidate"

	// Write the JSON data
	if _, err := w.Write(content); err != nil {
		w.Close() // Close the writer even on error
		return nil, err
	}

	// Close the writer to complete the operation
	if err := w.Close(); err != nil {
		return nil, err
	}

	// We need to get the generation of the newly created object.
	// Unfortunately, w.Attrs() might not be available until Close() completes.
	// In GCS Go client, after Close(), we can get some info but not easily the generation.
	// Actually, let's just re-read it to be sure and get the full LeaseObject.
	return s.ReadLease(ctx, key)
}

func (s *GCSStorage) UpdateLease(ctx context.Context, key string, obj *LeaseObject) (*LeaseObject, error) {
	handle := s.getObjectHandle(key)

	generation, err := strconv.ParseInt(obj.ResourceVersion, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid resource version: %w", err)
	}

	// Marshal the lease data to JSON
	content, err := json.Marshal(obj.Data)
	if err != nil {
		return nil, fmt.Errorf("marshalling lease data: %w", err)
	}

	// Create a conditional writer
	w := handle.If(gcs.Conditions{GenerationMatch: generation}).NewWriter(ctx)
	w.ContentType = "application/json"
	w.CacheControl = "no-cache, no-store, must-revalidate"

	// Write the JSON data
	if _, err := w.Write(content); err != nil {
		w.Close()
		return nil, err
	}

	// Close the writer to complete the operation
	if err := w.Close(); err != nil {
		return nil, err
	}

	// Re-read to get the new ResourceVersion
	return s.ReadLease(ctx, key)
}

func (s *GCSStorage) DeleteLease(ctx context.Context, key string) error {
	obj := s.getObjectHandle(key)
	return obj.Delete(ctx)
}

func (s *GCSStorage) IsNotFound(err error) bool {
	return errors.Is(err, gcs.ErrObjectNotExist)
}

func (s *GCSStorage) IsConflict(err error) bool {
	var gErr *googleapi.Error
	if errors.As(err, &gErr) {
		return gErr.Code == http.StatusPreconditionFailed || gErr.Code == http.StatusConflict
	}
	return false
}
