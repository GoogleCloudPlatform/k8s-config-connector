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

package manifest

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratedManifestsVersion(t *testing.T) {
	// Root of the repository from operator/pkg/manifest
	repoRoot := "../../../"
	versionBytes, err := os.ReadFile(filepath.Join(repoRoot, "version/VERSION"))
	if err != nil {
		t.Fatalf("failed to read version/VERSION: %v", err)
	}
	version := strings.TrimSpace(string(versionBytes))

	channels := []string{"channels", "autopilot-channels"}
	for _, channel := range channels {
		basePath := filepath.Join(repoRoot, "operator", channel, "packages/configconnector", version)

		// Check if the directory exists. If not, the manifests haven't been generated.
		// In CI, they should be generated. Locally, they might be missing.
		if _, err := os.Stat(basePath); os.IsNotExist(err) {
			t.Logf("Skipping version check for %s: directory %s does not exist. Run 'make manifests' to generate.", channel, basePath)
			continue
		}

		filesToCheck := []string{
			filepath.Join(basePath, "crds.yaml"),
			filepath.Join(basePath, "namespaced/0-cnrm-system.yaml"),
			filepath.Join(basePath, "namespaced/per-namespace-components.yaml"),
			filepath.Join(basePath, "cluster/gcp-identity/0-cnrm-system.yaml"),
			filepath.Join(basePath, "cluster/workload-identity/0-cnrm-system.yaml"),
		}

		for _, file := range filesToCheck {
			t.Run(fmt.Sprintf("%s/%s", channel, file), func(t *testing.T) {
				if _, err := os.Stat(file); os.IsNotExist(err) {
					t.Errorf("expected manifest file %s does not exist", file)
					return
				}

				f, err := os.Open(file)
				if err != nil {
					t.Fatalf("failed to open %s: %v", file, err)
				}
				defer f.Close()

				versionTag := fmt.Sprintf("cnrm.cloud.google.com/version: %s", version)
				found := false
				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					if strings.Contains(scanner.Text(), versionTag) {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("version tag %q not found in %s", versionTag, file)
				}

				// Also check for image versions for KCC components
				// Only for 0-cnrm-system.yaml and per-namespace-components.yaml files
				if strings.HasSuffix(file, ".yaml") && !strings.HasSuffix(file, "crds.yaml") {
					if _, err := f.Seek(0, 0); err != nil {
						t.Fatalf("failed to seek: %v", err)
					}
					scanner = bufio.NewScanner(f)
					kccComponents := []string{"recorder", "webhook", "deletiondefender", "unmanageddetector", "controller"}

					for scanner.Scan() {
						line := scanner.Text()
						for _, component := range kccComponents {
							if strings.Contains(line, fmt.Sprintf("cnrm/%s:", component)) {
								if !strings.Contains(line, fmt.Sprintf(":%s", version)) {
									t.Errorf("component %s in %s does not have version %s: %s", component, file, version, line)
								}
							}
						}
					}
				}
			})
		}
	}
}
