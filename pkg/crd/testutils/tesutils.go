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

package testutils

import (
	"flag"
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/google/go-cmp/cmp"
)

var update = flag.Bool("update", false, "update .golden files")

func VerifyContentsMatch(t *testing.T, actualBytes []byte, expectedFilePath string) {
	if *update {
		os.WriteFile(expectedFilePath, actualBytes, 0644)
	}
	expectedBytes, err := os.ReadFile(expectedFilePath)
	if err != nil {
		t.Fatalf("error reading file '%v': %v", expectedFilePath, err)
	}
	expectedBytes = []byte(test.TrimLicenseHeaderFromYaml(string(expectedBytes)))
	diff := cmp.Diff(actualBytes, expectedBytes)
	if diff != "" {
		t.Logf("actual value:\n%v", string(actualBytes))
		t.Logf("expected value:\n%v", string(expectedBytes))
		t.Logf("actual value does not match expected value in '%v'", expectedFilePath)
		t.Fatalf("diff: %v", diff)
	}
}
