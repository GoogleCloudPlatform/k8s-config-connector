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

func TestValidTestPath(t *testing.T) {
	lowercaseGVKs := loadGVKWithLowercaseKind()
	basicTestDataPath := repo.GetBasicIntegrationTestDataPath()
	err := filepath.WalkDir(basicTestDataPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			dirs := strings.Split(path, "/")
			if len(dirs) < 8 {
				// Test data must be under the leaf directory with at least
				// 8 levels of depth: pkg/test/resourcefixture/testdata/basic/[group]/[version]/[kind].
				return nil
			}

			// If it's at the leaf directory then the last section in dirs array
			// is the test case name; otherwise, it's just a parent directory
			// name.
			potentialTestCaseName := dirs[len(dirs)-1]
			// Leaf directories like "_vcr_cassettes" are not test cases.
			if strings.HasPrefix(potentialTestCaseName, "_") {
				return nil
			}

			// If it's at the leaf directory then the second to last section in
			// dirs array may be the test kind; otherwise, it's just a parent
			// directory name.
			potentialTestKind := dirs[len(dirs)-2]
			potentialTestVersion := dirs[len(dirs)-3]
			potentialTestGroup := dirs[len(dirs)-4]
			if potentialTestKind == "v1beta1" || potentialTestKind == "v1alpha1" {
				// When there is only one test case for a kind, it's possible
				// that the test case name is the test kind.
				potentialTestKind = potentialTestCaseName
				potentialTestVersion = dirs[len(dirs)-2]
				potentialTestGroup = dirs[len(dirs)-3]
			}

			potentialLowercaseGVK := schema.GroupVersionKind{
				Group:   fmt.Sprintf("%s.cnrm.cloud.google.com", potentialTestGroup),
				Version: potentialTestVersion,
				Kind:    potentialTestKind,
			}

			files, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			errMsg := fmt.Sprintf("test case %q has parsed group/version %q, "+
				"kind (lowercase) %q, and this is not supported; the path to test case "+
				"should be in the format of 'pkg/test/resourcefixture/testdata/basic/[group]/[version]/[kind]/' or "+
				"'pkg/test/resourcefixture/testdata/basic/[group]/[version]/[kind]/[testcasename]'",
				path, potentialLowercaseGVK.GroupVersion(), potentialLowercaseGVK.Kind)
			isLeafDir := false
			if len(files) == 0 { // leaf directory
				isLeafDir = true
				if _, ok := lowercaseGVKs[potentialLowercaseGVK]; !ok {
					t.Error(errMsg)
				}
			} else {
				isLowest := true
				for _, file := range files {
					// Directories like "_vcr_cassettes" are not test cases.
					if file.IsDir() && !strings.HasPrefix(file.Name(), "_") {
						// Contains subdirectory that could map to test cases.
						isLowest = false
						break
					}
				}
				if isLowest { // confirmed leaf directory
					isLeafDir = true
					if _, ok := lowercaseGVKs[potentialLowercaseGVK]; !ok {
						t.Error(errMsg)
					}
				}
			}
			if isLeafDir {
				return nil
			}
			// Verify there is no test data (.yaml file) when it is not the leaf directory.
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if strings.HasSuffix(file.Name(), ".yaml") {
					t.Errorf("test data file %q shouldn't be under path %q", file.Name(), path)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error validating test path: %v", err)
	}
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
