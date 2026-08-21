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

package resourcefixture

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
)

func TestValidBasicTestPath(t *testing.T) {
	testGVK := schema.GroupVersionKind{
		Group:   "dummygroup.cnrm.cloud.google.com",
		Version: "v1alpha1",
		Kind:    "DummyKind",
	}
	tests := []struct {
		name        string
		path        string
		gvk         schema.GroupVersionKind
		isValidPath bool
		hasError    bool
	}{
		{
			name:        "valid path",
			path:        "/pkg/test/resourcefixture/testdata/basic/dummygroup/v1alpha1/dummykind/",
			gvk:         testGVK,
			isValidPath: true,
		},
		{
			name:        "valid path with test case name",
			path:        "/pkg/test/resourcefixture/testdata/basic/dummygroup/v1alpha1/dummykind/basictestcase",
			gvk:         testGVK,
			isValidPath: true,
		},
		{
			name:     "invalid path with unsupported kind",
			path:     "/pkg/test/resourcefixture/testdata/basic/dummygroup/v1alpha1/realkind/basictestcase",
			gvk:      testGVK,
			hasError: true,
		},
		{
			name:     "invalid path with unsupported version",
			path:     "/pkg/test/resourcefixture/testdata/basic/dummygroup/v1beta1/dummykind/basictestcase",
			gvk:      testGVK,
			hasError: true,
		},
		{
			name:     "invalid path with unsupported group",
			path:     "/pkg/test/resourcefixture/testdata/basic/realgroup/v1alpha1/dummykind/basictestcase",
			gvk:      testGVK,
			hasError: true,
		},
		{
			name:        "valid path with no version",
			path:        "/pkg/test/resourcefixture/testdata/basic/dummygroup/dummykind",
			gvk:         testGVK,
			isValidPath: true,
		},
		{
			name:     "invalid path with incorrect prefix",
			path:     "/pkg/test/resourcefixture/testdata/advanced/dummygroup/v1beta1/dummykind/basictestcase",
			gvk:      testGVK,
			hasError: true,
		},
	}

	fixtureTests, err := loadFixtureTests()
	if err != nil {
		t.Fatalf("error loading fixture tests: %v", err)
	}
	for _, fixtureTest := range fixtureTests {
		tests = append(tests, struct {
			name        string
			path        string
			gvk         schema.GroupVersionKind
			isValidPath bool
			hasError    bool
		}{
			name:        fixtureTest.RelativePath,
			path:        fixtureTest.RelativePath,
			gvk:         fixtureTest.GVK,
			isValidPath: true,
		})
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actualResult, err := isValidBasicTestPath(tc.path, tc.gvk)
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

func isValidBasicTestPath(path string, gvk schema.GroupVersionKind) (bool, error) {
	// Remove any leading slash, ensure one trailing slash for consistent processing.
	path = strings.Trim(path, "/")
	path += "/"

	prefix := "pkg/test/resourcefixture/testdata/basic/"

	group := gvk.Group[:strings.Index(gvk.Group, ".")]
	version := gvk.Version
	kind := strings.ToLower(gvk.Kind)

	expectedPaths := []string{
		prefix + group + "/" + version + "/" + kind + "/", // Legacy format include version
		prefix + group + "/" + kind + "/",                 // Format that does not require moves to switch versions
	}

	for _, expectedPath := range expectedPaths {
		if strings.HasPrefix(path, expectedPath) {
			return true, nil
		}
	}

	return false, fmt.Errorf("path %q does not match expected formats: %v", path, expectedPaths)
}

// FixtureTest describes a fixture test case with its relative path and GVK.
type FixtureTest struct {
	// RelativePath is the path relative to the root of the repo.
	RelativePath string

	// GVK is the GroupVersionKind of the primary resource in the test.
	GVK schema.GroupVersionKind
}

// loadFixtureTests loads all fixture tests under the basic test data path.
func loadFixtureTests() ([]FixtureTest, error) {
	result := make([]FixtureTest, 0)
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
				return fmt.Errorf("no file under %q: test directories are expected to have subdirectories or files", path)
			}

			hasYAML := false
			for _, file := range files {
				if strings.HasSuffix(file.Name(), ".yaml") {
					hasYAML = true
					break
				}
			}

			if hasYAML {
				// Read and parse create.yaml, all tests should have this and this defines the primary GVK for the test.
				createFilePath := filepath.Join(path, "create.yaml")
				createFileBytes, err := os.ReadFile(createFilePath)
				if err != nil {
					return fmt.Errorf("error reading create.yaml under %q: %w", path, err)
				}

				var u *unstructured.Unstructured
				if err := yaml.Unmarshal(createFileBytes, &u); err != nil {
					return fmt.Errorf("error unmarshaling create.yaml under %q: %w", path, err)
				}
				gvk := u.GroupVersionKind()

				if gvk.Kind == "IAMPolicyMember" {
					// Ignore path validation, the primary resource is not the test in create.yaml
					return nil
				}
				// Get the relative path from the root of the repo.
				relativePath := strings.TrimPrefix(path, rootPath)

				result = append(result, FixtureTest{
					RelativePath: relativePath,
					GVK:          gvk,
				})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
