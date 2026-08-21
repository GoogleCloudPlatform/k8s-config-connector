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
	"fmt"
	"net/url"
	"strings"
)

func GetBucketAndPrefix(storageKey string) (string, string, error) {
	u, err := url.Parse(storageKey)
	if err != nil {
		return "", "", fmt.Errorf("error parsing url '%v': %w", storageKey, err)
	}
	expectedScheme := "gs"
	if u.Scheme != expectedScheme {
		return "", "", fmt.Errorf("url '%v' has an invalid scheme of '%v' must be '%v'", storageKey, u.Scheme, expectedScheme)
	}
	return u.Host, strings.TrimPrefix(u.Path, "/"), nil
}

func GetFullURI(bucketName, objectName string) string {
	if objectName == "" {
		return fmt.Sprintf("gs://%v", bucketName)
	}
	return fmt.Sprintf("gs://%v/%v", bucketName, objectName)
}
