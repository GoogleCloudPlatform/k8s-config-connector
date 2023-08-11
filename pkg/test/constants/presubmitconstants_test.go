// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package testconstants

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/fileutil"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"
)

func TestAllServicesInMap(t *testing.T) {
	testDataPath := repo.GetBasicIntegrationTestDataPath()
	services, err := fileutil.SubdirsIn(testDataPath)
	if err != nil {
		t.Fatal(err)
	}
	smLoader, err := servicemappingloader.New()
	if err != nil {
		t.Fatalf("error loading service mappings: %v", err)
	}
	autoGenOnlyGroup := smLoader.GetAutoGenOnlyGroups()
	for _, s := range services {
		group := fmt.Sprintf("%s.cnrm.cloud.google.com", s)
		if _, ok := autoGenOnlyGroup[group]; ok {
			// Services with only auto-generated resources can be excluded from
			// presubmit test suite.
			continue
		}
		if _, ok := RepresentativeCRUDTestsForAllServices[s]; !ok {
			t.Fatalf("Missing an entry in the `RepresentativeCRUDTestsForAllServices` map for service: %s", s)
		}
	}
}
func TestMappedServicesExistInDir(t *testing.T) {
	testDataPath := repo.GetBasicIntegrationTestDataPath()
	services, err := fileutil.SubdirsIn(testDataPath)
	if err != nil {
		t.Fatal(err)
	}
	for s, _ := range RepresentativeCRUDTestsForAllServices {
		if slice.StringSliceContains(services, s) {
			continue
		} else {
			t.Fatalf("Service %s does not exist or might be misspelled", s)
		}
	}
}
func TestEachServiceHasAtLeastOneTestCase(t *testing.T) {
	for s, tc := range RepresentativeCRUDTestsForAllServices {
		if len(tc) == 0 {
			t.Fatalf("Service %s does not have any test cases added to it", s)
		} else {
			continue
		}
	}
}
func TestValidTestName(t *testing.T) {
	testDataPath := repo.GetBasicIntegrationTestDataPath()
	for s, tcs := range RepresentativeCRUDTestsForAllServices {
		serviceTestPath := testDataPath + "/" + s
		for _, tc := range tcs {
			found := false
			err := filepath.WalkDir(serviceTestPath, func(path string, di fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if di.IsDir() {
					return nil
				}
				// Test cases are the names of the leaf-node directory, only match on the lowest subdirectory name
				if strings.HasSuffix(filepath.Dir(path), tc) {
					found = true
					return nil
				}
				return nil
			})
			if err != nil {
				t.Fatalf("Error occurred while walking through directory: %v", err)
			}
			if !found {
				t.Fatalf("Test case %s does not exist or is misspelled", tc)
			}
		}
	}
}
func TestGetPresubmitLiteRegexString(t *testing.T) {
	s := GetPresubmitLiteRegexStringArray()
	// Check if string slice contains any test cases from the long running or
	// intentionally periodic test cases
	for _, v := range longRunningCRUDTests {
		if slice.StringSliceContains(s, v) {
			t.Fatalf("Presubmit-Lite test list contains long running test: %s", v)
		}
	}
	for _, v := range periodicCRUDTests {
		if slice.StringSliceContains(s, v) {
			t.Fatalf("Presubmit-Lite test list contains periodic test: %s", v)
		}
	}
}
func TestJoinTestNamesWithRegexFormat(t *testing.T) {
	testNames := []string{"pubsubtopic", "service", "sqluser"}
	expected := "-pubsubtopic$|-service$|-sqluser$"
	result := JoinTestNamesWithRegexFormat(testNames)
	if expected != result {
		t.Fatalf("String mismatch, expected %s, got %s", expected, result)
	}
}
