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

package resourcefixture

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"github.com/ghodss/yaml" //nolint:depguard
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type TestType string

const (
	Unknown                     TestType = "unknown"
	Basic                       TestType = "basic"
	ContainerAnnotations        TestType = "containerannotations"
	Directives                  TestType = "directives"
	ExternalRef                 TestType = "externalref"
	SensitiveField              TestType = "sensitivefield"
	IAMExternalOnlyRef          TestType = "iamexternalonlyref"
	IAMMemberReferences         TestType = "iammemberreferences"
	ResourceID                  TestType = "resourceid"
	StateAbsentInSpec           TestType = "stateabsentinspec"
	ResourceOverrides           TestType = "resourceoverrides"
	ReconcileIntervalAnnotation TestType = "reconcileintervalannotations"
)

type ResourceFixture struct {
	GVK          schema.GroupVersionKind
	Name         string
	SourceDir    string
	Create       []byte
	Update       []byte
	Dependencies []byte
	Type         TestType
}

// Load loads all test cases found in the testdata directory. A
// test case is any directory in the tree that contains a create.yaml file (and
// optionally a dependencies.yaml and update.yaml). The name of the directory
// containing the YAMLs is used as the name of the test case.
func Load(t *testing.T) []ResourceFixture {
	return LoadWithFilter(t, nil, nil)
}

type LightFilter func(name string, testType TestType) bool
type HeavyFilter func(fixture ResourceFixture) bool

// LoadWithFilter returns all fixtures that match the filter functions - a filter function matches by returning 'true'
// * use 'lightFilterFunc' for filtering based on test names and types (determining these values is 'lightweight' as it
// only relies on directory and file names)
// * use 'heavyFilterFunc' for filtering based on the contents of the YAML file(s)
//
// if a 'nil' value is supplied for a given filter function then it is assumed that all fixtures match the filter
func LoadWithFilter(t *testing.T, lightFilterFunc LightFilter, heavyFilterFunc HeavyFilter) []ResourceFixture {
	t.Helper()
	allCases := make([]ResourceFixture, 0)
	baseDir := getTestDataPath(t)
	if err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			return nil
		}

		// This is a slightly inefficient, but we want to reuse the existing code
		fileInfos, err := os.ReadDir(path)
		if err != nil {
			return fmt.Errorf("error reading directory '%v': %w", path, err)
		}

		testToFileName := make(map[string]string)
		for _, fi := range fileInfos {
			if fi.IsDir() {
				continue
			}
			if !strings.HasSuffix(fi.Name(), ".yaml") {
				continue
			}
			fileNameNoExt := strings.TrimSuffix(fi.Name(), ".yaml")
			if value, ok := testToFileName[fileNameNoExt]; ok {
				return fmt.Errorf("error, conflicting files for test '%v' in '%v': {%v, %v}", fileNameNoExt, path, value, fi.Name())
			}
			testToFileName[fileNameNoExt] = fi.Name()
		}

		// TODO: something about tags here
		if createFile, ok := testToFileName["create"]; ok {
			updateFile := testToFileName["update"]
			depFile := testToFileName["dependencies"]
			name := filepath.Base(path)
			testType := parseTestTypeFromPath(t, path)
			if lightFilterFunc != nil && !lightFilterFunc(name, testType) {
				return nil
			}
			rf := loadResourceFixture(t, name, testType, path, createFile, updateFile, depFile)
			if heavyFilterFunc != nil && !heavyFilterFunc(rf) {
				return nil
			}
			allCases = append(allCases, rf)
		}

		return nil
	}); err != nil {
		t.Fatalf("error walking directory %q: %v", baseDir, err)
	}

	return allCases
}

func loadResourceFixture(t *testing.T, testName string, testType TestType, dir, createFile, updateFile, depFile string) ResourceFixture {
	t.Helper()
	createConfig := test.MustReadFile(t, path.Join(dir, createFile))
	gvk, err := readGroupVersionKind(t, createConfig)
	if err != nil {
		t.Fatalf("unable to determine GroupVersionKind for test case named %v: %v", testName, err)
	}
	if gvk.Kind == "" {
		t.Fatalf("got empty kind for test case named %v", testName)
	}
	if gvk.Version == "" {
		t.Fatalf("got empty version for test case named %v", testName)
	}

	rf := ResourceFixture{
		Name:      testName,
		SourceDir: dir,
		GVK:       gvk,
		Create:    createConfig,
		Type:      testType,
	}

	if updateFile != "" {
		rf.Update = test.MustReadFile(t, path.Join(dir, updateFile))
	}
	if depFile != "" {
		rf.Dependencies = test.MustReadFile(t, path.Join(dir, depFile))
	}
	return rf
}

func readGroupVersionKind(t *testing.T, config []byte) (schema.GroupVersionKind, error) {
	t.Helper()
	u := &unstructured.Unstructured{}
	err := yaml.Unmarshal(config, u)
	if err != nil {
		return schema.GroupVersionKind{}, fmt.Errorf("error unmarshalling bytes to CRD: %w", err)
	}
	return u.GroupVersionKind(), nil
}

func parseTestTypeFromPath(t *testing.T, path string) TestType {
	t.Helper()
	switch parseTestDataSubdirFromPath(t, path) {
	case "basic":
		return Basic
	case "containerannotations":
		return ContainerAnnotations
	case "directives":
		return Directives
	case "externalref":
		return ExternalRef
	case "sensitivefield":
		return SensitiveField
	case "iamexternalonlyref":
		return IAMExternalOnlyRef
	case "iammemberreferences":
		return IAMMemberReferences
	case "resourceid":
		return ResourceID
	case "stateabsentinspec":
		return StateAbsentInSpec
	case "resourceoverrides":
		return ResourceOverrides
	case "reconcileintervalannotations":
		return ReconcileIntervalAnnotation
	default:
		t.Fatalf("failed to parse test type for path %v", path)
		return Unknown
	}
}

func parseTestDataSubdirFromPath(t *testing.T, path string) string {
	t.Helper()
	testDataPath := getTestDataPath(t)
	pathWithoutTestDataPath := strings.TrimPrefix(path, testDataPath+"/")
	pathTokens := strings.Split(pathWithoutTestDataPath, "/")
	return pathTokens[0]
}

func getTestDataPath(t *testing.T) string {
	t.Helper()
	packagePath := repo.GetCallerPackagePathOrTestFatal(t)
	return filepath.Join(packagePath, "testdata")
}
