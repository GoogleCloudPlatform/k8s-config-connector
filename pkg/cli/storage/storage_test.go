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

package storage_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/storage"
)

func TestGetBucketAndPrefix(t *testing.T) {
	testCases := []struct {
		Name         string
		StorageKey   string
		BucketName   string
		ObjectName   string
		ErrorMessage string
	}{
		{
			Name:         "empty string should fail",
			StorageKey:   "",
			BucketName:   "",
			ObjectName:   "",
			ErrorMessage: "url '' has an invalid scheme of '' must be 'gs'",
		},
		{
			Name:         "invalid scheme should fail",
			StorageKey:   "s3://my-bucket/my-key",
			BucketName:   "",
			ObjectName:   "",
			ErrorMessage: "url 's3://my-bucket/my-key' has an invalid scheme of 's3' must be 'gs'",
		},
		{
			Name:         "bucket only should succeed",
			StorageKey:   "gs://my-bucket",
			BucketName:   "my-bucket",
			ObjectName:   "",
			ErrorMessage: "",
		},
		{
			Name:         "bucket only with slash should succeed",
			StorageKey:   "gs://my-bucket/",
			BucketName:   "my-bucket",
			ObjectName:   "",
			ErrorMessage: "",
		},
		{
			Name:         "full object should succeed",
			StorageKey:   "gs://my-bucket/my-key/in/my/path",
			BucketName:   "my-bucket",
			ObjectName:   "my-key/in/my/path",
			ErrorMessage: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			bucketName, objectName, err := storage.GetBucketAndPrefix(tc.StorageKey)
			if err == nil && tc.ErrorMessage != "" {
				t.Fatalf("expected an error, instead got 'nil'")
			}
			if err != nil {
				if err.Error() != tc.ErrorMessage {
					t.Fatalf("error message mismatch: got '%v', want '%v'", err.Error(), tc.ErrorMessage)
				}
				return
			}
			if bucketName != tc.BucketName {
				t.Errorf("bucket name mismatch: got '%v', want '%v'", bucketName, tc.BucketName)
			}
			if objectName != tc.ObjectName {
				t.Errorf("object name mismatch: got '%v', want '%v'", objectName, tc.ObjectName)
			}
		})
	}
}

func TestGetFullURI(t *testing.T) {
	testCases := []struct {
		Name           string
		BucketName     string
		ObjectName     string
		ExpectedResult string
	}{
		{
			Name:           "only bucket",
			BucketName:     "bucket-name",
			ObjectName:     "",
			ExpectedResult: "gs://bucket-name",
		},
		{
			Name:           "only bucket",
			BucketName:     "bucket-name",
			ObjectName:     "object/name",
			ExpectedResult: "gs://bucket-name/object/name",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := storage.GetFullURI(tc.BucketName, tc.ObjectName)
			if result != tc.ExpectedResult {
				t.Errorf("unexpected result: got '%v', want '%v'", result, tc.ExpectedResult)
			}
		})
	}
}
