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

package asset_test

import (
	"os"
	"testing"

	testexport "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/test/export"
	testgcp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/gcp"

	"google.golang.org/api/storage/v1"
)

func TestNewStreamFromStorageObject(t *testing.T) {
	httpClient := testgcp.NewDefaultHTTPClient(t)
	bucketName, objectName := testexport.NewTemporaryBucketAndObjectName(t, httpClient)
	storageClient := testgcp.NewStorageClient(t)
	object := storage.Object{
		Name: objectName,
	}
	file, err := os.Open(exportJSONFile)
	if err != nil {
		t.Fatalf("error opening '%v': %v", exportJSONFile, err)
	}
	defer file.Close()
	if _, err := storageClient.Objects.Insert(bucketName, &object).Media(file).Do(); err != nil {
		t.Fatalf("error writing file '%v' to GCS: %v", file, err)
	}
	testStorageStream(t, bucketName, objectName, expectedAssetCount)
}
