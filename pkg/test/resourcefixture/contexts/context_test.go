// Copyright 2025 Google LLC
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

package contexts

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
)

func TestValidTestName(t *testing.T) {
	testDataPath := repo.GetIntegrationTestDataPath()
	testCaseNames, err := loadTestCaseNameKindMap(testDataPath)
	if err != nil {
		t.Fatalf("error loading test case name to kind map: %v", err)
	}
	for testCaseName, testCaseContext := range resourceContextMap {
		testKind, ok := testCaseNames[testCaseName]
		if !ok {
			t.Errorf("test case %q not found under path %q", testCaseName, testDataPath)
			continue
		}

		if testKind != "" && testKind != strings.ToLower(testCaseContext.ResourceKind) {
			t.Errorf("incorrect kind %q in resource fixture for test case %q; or the test case is under the incorrect kind directory %q", testCaseContext.ResourceKind, testCaseName, testKind)
		}
	}
}

func loadTestCaseNameKindMap(root string) (map[string]string, error) {
	testCaseNames := make(map[string]string)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			dirs := strings.Split(path, "/")
			if len(dirs) < 3 {
				// Test data must be under the leaf directory with at least
				// 3 levels of depth: [group]/[version]/[kind]
				return nil
			}

			files, err := os.ReadDir(path)
			if err != nil {
				return err
			}

			testCaseName := dirs[len(dirs)-1]
			if strings.HasPrefix(testCaseName, "_") {
				// Leaf directories like "_vcr_cassettes" are not test cases.
				return nil
			}

			testKind := ""
			// It's possible to identify the kind from path for basic testdata.
			if strings.Contains(path, "resourcefixture/testdata/basic") {
				testKind = dirs[len(dirs)-2]
				if testKind == "v1beta1" || testKind == "v1alpha1" {
					testKind = testCaseName
				}
			}

			if len(files) == 0 {
				testCaseNames[testCaseName] = testKind
				return nil
			}

			isLowest := true
			for _, file := range files {
				// Directories like "_vcr_cassettes" are not test cases.
				if file.IsDir() && !strings.HasPrefix(file.Name(), "_") {
					isLowest = false
					break
				}
			}
			if isLowest {
				testCaseNames[testCaseName] = testKind
			}
		}
		return nil
	})
	return testCaseNames, err
}
