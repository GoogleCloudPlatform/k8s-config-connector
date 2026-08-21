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

//go:build integration
// +build integration

package export_test

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset/export"
	testexport "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/test/export"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
)

func TestForParentToStorageObject(t *testing.T) {
	httpClient := testgcp.NewDefaultHTTPClient(t)
	bucketName, prefix := testexport.NewTemporaryBucketAndObjectName(t, httpClient)
	defer testexport.DeleteTemporaryBucket(t, httpClient, bucketName)
	projectID := testgcp.GetDefaultProjectID(t)
	if err := export.ForParentToStorageObject(context.TODO(), httpClient,
		fmt.Sprintf("projects/%v", projectID), bucketName, prefix); err != nil {
		t.Fatalf("error exporting asset inventory: %v", err)
	}
	defer testexport.DeleteExport(t, httpClient, bucketName, prefix)
	storageClient := testgcp.NewStorageClient(t)
	response, err := storageClient.Objects.Get(bucketName, prefix).Download()
	if err != nil {
		t.Fatalf("error initiating download of %v/%v: %v", bucketName, prefix, err)
	}
	defer response.Body.Close()
	stream := asset.NewStream(response.Body)
	for _, err := stream.Next(); err != io.EOF; _, err = stream.Next() {
		if err != nil {
			t.Fatalf("error reading asset: %v", err)
		}
	}
}

func TestNewTemporaryBucketObjectName(t *testing.T) {
	httpClient := testgcp.NewDefaultHTTPClient(t)
	bucketName, object := testexport.NewTemporaryBucketAndObjectName(t, httpClient)
	defer testexport.DeleteTemporaryBucket(t, httpClient, bucketName)
	if object == "" {
		t.Fatalf("unexpected empty string value for 'object'")
	}
	storageClient := testgcp.NewStorageClient(t)
	bucket, err := storageClient.Buckets.Get(bucketName).Do()
	if err != nil {
		t.Fatalf("unexpected error getting bucket: %v", err)
	}
	if bucket.Name != bucketName {
		t.Fatalf("bucket name mismatch: got '%v' which doesn't match '%v'", bucket.Name, bucketName)
	}
}
