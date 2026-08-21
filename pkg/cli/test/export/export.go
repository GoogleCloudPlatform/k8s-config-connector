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

package testexport

import (
	"context"
	"net/http"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/asset/export"
)

func NewTemporaryBucketAndObjectName(t *testing.T, httpClient *http.Client) (string, string) {
	bucketName, objectName, err := export.NewTemporaryBucketAndObjectName(context.TODO(), httpClient)
	if err != nil {
		t.Fatalf("unexpected error creating new temporary bucket: %v", err)
	}
	return bucketName, objectName
}

func DeleteExport(t *testing.T, httpClient *http.Client, bucketName, objectName string) {
	if err := export.DeleteExport(context.TODO(), httpClient, bucketName, objectName); err != nil {
		t.Fatalf("unexpected error deleting export: %v", err)
	}
}

func DeleteTemporaryBucket(t *testing.T, httpClient *http.Client, bucketName string) {
	if err := export.DeleteTemporaryBucket(context.TODO(), httpClient, bucketName); err != nil {
		t.Fatalf("unexpected error deleting temporary bucket: %v", err)
	}
}
