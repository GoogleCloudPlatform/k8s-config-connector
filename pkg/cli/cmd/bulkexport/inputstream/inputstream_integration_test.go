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

package inputstream_test

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/inputstream"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/parameters"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/storage"
	testexport "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/test/export"
	testos "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/test/os"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"
	testvariable "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture/variable"

	gcpstorage "google.golang.org/api/storage/v1"
)

const (
	expectedAssetCount = 19
	exportJsonFile     = "testdata/assets.json"
)

func TestNewAssetStreamFromStdin(t *testing.T) {
	bytes, err := ioutil.ReadFile(exportJsonFile)
	if err != nil {
		t.Fatalf("error reading '%v': %v'", exportJsonFile, err)
	}
	cleanup, stdin := testos.GetStdin(t, string(bytes))
	defer cleanup()
	params := parameters.Parameters{}
	assetStream, err := inputstream.NewAssetStream(&params, stdin)
	if err != nil {
		t.Fatalf("error creating asset stream: %v", err)
	}
	defer closeStream(t, assetStream)
	assets := streamToSlice(assetStream)
	if len(assets) != expectedAssetCount {
		t.Fatalf("unexpected number of assets in stream: got '%v', want '%v'", len(assets), expectedAssetCount)
	}
}

func TestNewStorageObjectStreamFromStorageKey(t *testing.T) {
	httpClient := testgcp.NewDefaultHTTPClient(t)
	bucketName, objectName := testexport.NewTemporaryBucketAndObjectName(t, httpClient)
	storageClient := testgcp.NewStorageClient(t)
	object := gcpstorage.Object{
		Name: objectName,
	}
	file, err := os.Open(exportJsonFile)
	if err != nil {
		t.Fatalf("error opening '%v': %v", exportJsonFile, err)
	}
	defer file.Close()
	if _, err := storageClient.Objects.Insert(bucketName, &object).Media(file).Do(); err != nil {
		t.Fatalf("error writing file '%v' to GCS: %v", file, err)
	}
	params := parameters.Parameters{
		StorageKey: storage.GetFullURI(bucketName, objectName),
	}
	assets := createStorageStreamAndReadAll(t, params)
	if len(assets) != expectedAssetCount {
		t.Fatalf("unexpected number of assets in stream: got '%v', want '%v'", len(assets), expectedAssetCount)
	}
}

func TestNewStorageObjectStreamFromBucketParameter(t *testing.T) {
	bucketName := strings.ToLower(fmt.Sprintf("%v-%v", "inputstream-", testvariable.NewUniqueID()))
	storageClient := testgcp.NewStorageClient(t)
	bucket := &gcpstorage.Bucket{
		Name: bucketName,
	}
	projectID := testgcp.GetDefaultProjectID(t)
	bucket, err := storageClient.Buckets.Insert(projectID, bucket).Do()
	if err != nil {
		t.Fatalf("error creating bucket '%v': %v", bucketName, err)
	}
	httpClient := testgcp.NewDefaultHTTPClient(t)
	defer testexport.DeleteTemporaryBucket(t, httpClient, bucketName)
	params := parameters.Parameters{
		StorageKey: fmt.Sprintf("gs://%v", bucketName),
		ProjectID:  projectID,
	}
	assets := createStorageStreamAndReadAll(t, params)
	if assets == nil {
		t.Fatalf("unexpected nil value returned")
	}
	err = storageClient.Objects.List(bucketName).Pages(context.TODO(), func(objects *gcpstorage.Objects) error {
		if len(objects.Items) > 0 {
			names := make([]string, 0, len(objects.Items))
			for _, o := range objects.Items {
				names = append(names, o.Name)
			}
			t.Fatalf("Expected bucket '%v' to be empty: NewAssetStream should have deleted the export file. Instead the following objects were found: %v",
				bucketName, strings.Join(names, ", "))
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error listing %v: %v", bucketName, err)
	}
}

func TestNewStorageObjectStreamNoParameters(t *testing.T) {
	projectID := testgcp.GetDefaultProjectID(t)
	params := parameters.Parameters{
		ProjectID: projectID,
	}
	assets := createStorageStreamAndReadAll(t, params)
	if assets == nil {
		t.Fatalf("unexpected nil value returned")
	}
}

func createStorageStreamAndReadAll(t *testing.T, params parameters.Parameters) []asset.Asset {
	cleanup, stdin := testos.GetStdin(t, "")
	defer cleanup()
	assetStream, err := inputstream.NewAssetStream(&params, stdin)
	if err != nil {
		t.Fatalf("error creating asset stream: %v", err)
	}
	defer closeStream(t, assetStream)
	return streamToSlice(assetStream)
}

func streamToSlice(stream *asset.Stream) []asset.Asset {
	results := make([]asset.Asset, 0)
	for a, err := stream.Next(); err != io.EOF; a, err = stream.Next() {
		results = append(results, *a)
	}
	return results
}

func closeStream(t *testing.T, assetStream *asset.Stream) {
	err := assetStream.Close()
	if err != nil {
		t.Fatalf("error closing stream: %v", err)
	}
}
