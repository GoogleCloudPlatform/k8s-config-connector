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

package basic

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
)

func TestValidBasicTestPath(t *testing.T) {
	testLowercaseGVK := schema.GroupVersionKind{
		Group:   "dummygroup.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "dummykind",
	}
	testGVKs := make(map[schema.GroupVersionKind]bool)
	testGVKs[testLowercaseGVK] = true
	tests := []struct {
		name        string
		path        string
		validGVKs   map[schema.GroupVersionKind]bool
		isValidPath bool
		hasError    bool
	}{
		{
			name:        "valid path",
			path:        "/pkg/test/resourcefixture/testdata/basic/dummygroup/v1alpha1/dummykind/",
			validGVKs:   testGVKs,
			isValidPath: true,
		},
		{
			name:        "valid path with test case name",
			path:        "/pkg/test/resourcefixture/testdata/basic/dummygroup/v1alpha1/dummykind/basictestcase",
			validGVKs:   testGVKs,
			isValidPath: true,
		},
		{
			name:      "invalid path with unsupported kind",
			path:      "/pkg/test/resourcefixture/testdata/basic/dummygroup/v1alpha1/realkind/basictestcase",
			validGVKs: testGVKs,
			hasError:  true,
		},
		{
			name:      "invalid path with unsupported version",
			path:      "/pkg/test/resourcefixture/testdata/basic/dummygroup/v1beta1/dummykind/basictestcase",
			validGVKs: testGVKs,
			hasError:  true,
		},
		{
			name:      "invalid path with unsupported group",
			path:      "/pkg/test/resourcefixture/testdata/basic/realgroup/v1alpha1/dummykind/basictestcase",
			validGVKs: testGVKs,
			hasError:  true,
		},
		{
			name:      "invalid path with incorrect structure",
			path:      "/pkg/test/resourcefixture/testdata/basic/dummygroup/dummykind",
			validGVKs: testGVKs,
			hasError:  true,
		},
		{
			name:      "invalid path with incorrect prefix",
			path:      "/pkg/test/resourcefixture/testdata/advanced/dummygroup/v1beta1/dummykind/basictestcase",
			validGVKs: testGVKs,
			hasError:  true,
		},
	}

	lowercaseGVKs := loadGVKWithLowercaseKind()
	testPaths, err := loadTestPaths()
	if err != nil {
		t.Fatalf("error loading test paths: %v", err)
	}
	for _, testPath := range testPaths {
		tests = append(tests, struct {
			name        string
			path        string
			validGVKs   map[schema.GroupVersionKind]bool
			isValidPath bool
			hasError    bool
		}{
			name:        testPath,
			path:        testPath,
			validGVKs:   lowercaseGVKs,
			isValidPath: true,
		})
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actualResult, err := isValidBasicTestPath(tc.path, tc.validGVKs)
			if tc.hasError {
				if err == nil {
					t.Errorf("expected to have an error but got no error")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error validating basic test path: %v", err)
				}
				if actualResult != tc.isValidPath {
					t.Errorf("got %v, want %v", actualResult, tc.isValidPath)
				}
			}
		})
	}
}

func isValidBasicTestPath(path string, validLowercaseGVKs map[schema.GroupVersionKind]bool) (bool, error) {
	if !strings.HasPrefix(path, "/pkg/test/resourcefixture/testdata/basic/") {
		return false, fmt.Errorf("incorrect prefix for basic test path %q; should be pkg/test/resourcefixture/testdata/basic/", path)
	}
	dirs := strings.Split(path, "/")
	testCaseName := dirs[len(dirs)-1]
	testKind := dirs[len(dirs)-2]
	testVersion := dirs[len(dirs)-3]
	testGroup := dirs[len(dirs)-4]
	if testKind == "v1beta1" || testKind == "v1alpha1" {
		// When there is only one test case for a kind, it's possible
		// that the test case name is the test kind.
		testKind = testCaseName
		testVersion = dirs[len(dirs)-2]
		testGroup = dirs[len(dirs)-3]
	}
	lowercaseGVK := schema.GroupVersionKind{
		Group:   fmt.Sprintf("%s.cnrm.cloud.google.com", testGroup),
		Version: testVersion,
		Kind:    testKind,
	}
	if _, ok := validLowercaseGVKs[lowercaseGVK]; !ok {
		return false, fmt.Errorf("test case %q has parsed group/version %q, "+
			"kind (lowercase) %q, and this is not supported; the path to test case "+
			"should be in the format of 'pkg/test/resourcefixture/testdata/basic/[group]/[version]/[kind]/' or "+
			"'pkg/test/resourcefixture/testdata/basic/[group]/[version]/[kind]/[testcasename]'",
			path, lowercaseGVK.GroupVersion(), lowercaseGVK.Kind)
	}
	return true, nil
}

func loadTestPaths() ([]string, error) {
	result := make([]string, 0)
	rootPath := repo.GetRootOrLogFatal()
	basicTestDataPath := repo.GetBasicIntegrationTestDataPath()
	err := filepath.WalkDir(basicTestDataPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// Directories like "_vcr_cassettes" may contain yaml files but are not test cases.
			if strings.HasPrefix(d.Name(), "_") {
				return nil
			}
			files, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			if len(files) == 0 {
				return fmt.Errorf("no file under %q: test directories are expected to have subdirectories and files", path)
			}
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
					// Directories with any yaml file are considered a test path to verify.
					testPath := strings.TrimPrefix(path, rootPath)
					fmt.Println(testPath)
					result = append(result, testPath)
					break
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func loadGVKWithLowercaseKind() map[schema.GroupVersionKind]bool {
	result := make(map[schema.GroupVersionKind]bool)
	for gvk, _ := range supportedgvks.SupportedGVKs {
		lowercaseGVK := schema.GroupVersionKind{
			Group:   gvk.Group,
			Version: gvk.Version,
			Kind:    strings.ToLower(gvk.Kind),
		}
		result[lowercaseGVK] = true
	}
	return result
}
