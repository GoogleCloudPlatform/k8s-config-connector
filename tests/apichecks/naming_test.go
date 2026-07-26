// Copyright 2026 Google LLC
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

package lint

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

// Extracts the Kind from the spec.names block.
// Exactly 4 spaces ensures we match spec.names.kind and not the top-level CRD kind
// or any fields deep inside the OpenAPI schema.
var kindRegex = regexp.MustCompile(`(?m)^    kind: ([a-zA-Z0-9]+)\s*$`)

func TestDirectResourceFileNaming(t *testing.T) {
	apisDir := "../../apis"
	controllerDir := "../../pkg/controller/direct"
	crdsDir := "../../config/crds/resources"
	dirsToCheck := []string{apisDir, controllerDir}

	// 1. Build a flat set of all lowercased Kinds from the CRD folder
	allKinds := make(map[string]bool)
	err := filepath.Walk(crdsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yaml") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			matches := kindRegex.FindSubmatch(content)
			if len(matches) > 1 {
				kind := string(matches[1])
				allKinds[strings.ToLower(kind)] = true
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error scanning crds directory: %v", err)
	}

	suffixes := []string{
		"_types.go",
		"_types_test.go",
		"_identity.go",
		"_identity_test.go",
		"_reference.go",
		"_reference_test.go",
		"_mapper.go",
		"_mapper_test.go",
		"_fuzzer.go",
		"_fuzzer_test.go",
		"_controller.go",
		"_controller_test.go",
	}

	var errors []string

	// 2. Unified Walk
	for _, dirToCheck := range dirsToCheck {
		err := filepath.Walk(dirToCheck, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				if info.Name() == "directbase" {
					return filepath.SkipDir
				}
				return nil
			}

			// 3. Match format and extract prefix
			filename := info.Name()
			var matchedSuffix string
			for _, s := range suffixes {
				if strings.HasSuffix(filename, s) {
					matchedSuffix = s
					break
				}
			}
			if matchedSuffix == "" {
				return nil
			}
			prefix := strings.TrimSuffix(filename, matchedSuffix)

			// 4a. If the prefix is a valid kind, it passes!
			if allKinds[prefix] {
				return nil
			}

			// Extract service name based on directory structure
			var service string
			if strings.HasPrefix(path, apisDir) {
				// apis/<service>/<version>/...
				dir := filepath.Dir(path)
				service = filepath.Base(filepath.Dir(dir))
			} else {
				// pkg/controller/direct/<service>/...
				relPath, _ := filepath.Rel(controllerDir, path)
				service = strings.Split(filepath.ToSlash(relPath), "/")[0]
			}

			// 4b. Else if (service + prefix) is a valid kind, explicitly tell the user to fix it
			combined := strings.ToLower(service) + prefix
			if allKinds[combined] {
				expectedFilename := combined + matchedSuffix
				t.Errorf("File %s has a missing prefix. It should be renamed to use %s (using combined service %q and prefix %q)", path, expectedFilename, service, prefix)
				return nil
			}

			// 4c. Else, throw a generic error to be caught by the golden exception file
			relPath, _ := filepath.Rel("../..", path)
			normalizedPath := filepath.ToSlash(relPath)
			errors = append(errors, fmt.Sprintf("[naming_violation] file=%s prefix=%s (expected a valid resource kind prefix)", normalizedPath, prefix))

			return nil
		})
		if err != nil {
			t.Fatalf("error walking directory %s: %v", dirToCheck, err)
		}
	}

	sort.Strings(errors)
	want := strings.Join(errors, "\n")
	if want != "" {
		want += "\n"
	}
	test.CompareGoldenFile(t, "testdata/exceptions/naming_violations.txt", want)
}
