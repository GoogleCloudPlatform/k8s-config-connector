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

package repo_test

import (
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
)

func TestGetCallerPackagePath(t *testing.T) {
	path, err := repo.GetCallerPackagePath()
	if err != nil {
		t.Errorf("error getting package path: %v", err)
	}
	expectedPath, err := os.Getwd()
	if err != nil {
		t.Fatalf("error getting working directory: %v", err)
	}
	if path != expectedPath {
		t.Errorf("unexpected value for path: got '%v', want '%v'", path, expectedPath)
	}
}
