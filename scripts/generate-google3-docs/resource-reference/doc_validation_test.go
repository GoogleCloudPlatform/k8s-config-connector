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

package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
	"gopkg.in/yaml.v2"
)

type Toc struct {
	Toc []TocItem `yaml:"toc"`
}

type TocItem struct {
	Title   string       `yaml:"title"`
	Section []TocSection `yaml:"section,omitempty"`
	Path    string       `yaml:"path,omitempty"`
}

type TocSection struct {
	Title string `yaml:"title"`
	Path  string `yaml:"path"`
}

func TestDocConsistency(t *testing.T) {
	rootDir, err := repo.GetRoot()
	if err != nil {
		t.Fatalf("error getting repo root: %v", err)
	}

	baseDir := filepath.Join(rootDir, "scripts", "generate-google3-docs", "resource-reference")
	generatedDocsDir := filepath.Join(baseDir, "generated", "resource-docs")
	tocPath := filepath.Join(baseDir, "_toc.yaml")
	overviewPath := filepath.Join(baseDir, "overview.md")

	// 1. Get all generated .md files
	generatedFiles := make(map[string]bool)
	err = filepath.Walk(generatedDocsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			relPath, err := filepath.Rel(generatedDocsDir, path)
			if err != nil {
				return err
			}
			// Normalize to the format used in _toc.yaml and overview.md
			// They use /config-connector/docs/reference/resource-docs/...
			webPath := "/config-connector/docs/reference/resource-docs/" + relPath
			generatedFiles[webPath] = true
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error walking generated docs dir: %v", err)
	}

	// 2. Validate _toc.yaml
	tocData, err := ioutil.ReadFile(tocPath)
	if err != nil {
		t.Fatalf("error reading _toc.yaml: %v", err)
	}
	var toc Toc
	if err := yaml.Unmarshal(tocData, &toc); err != nil {
		t.Fatalf("error parsing _toc.yaml: %v", err)
	}

	tocPaths := make(map[string]bool)
	for _, item := range toc.Toc {
		if item.Path != "" {
			validatePath(t, item.Path, generatedFiles, tocPaths, "_toc.yaml")
		}
		for _, section := range item.Section {
			validatePath(t, section.Path, generatedFiles, tocPaths, "_toc.yaml")
		}
	}

	// 3. Validate overview.md
	overviewData, err := ioutil.ReadFile(overviewPath)
	if err != nil {
		t.Fatalf("error reading overview.md: %v", err)
	}
	overviewContent := string(overviewData)
	hrefRegex := regexp.MustCompile(`href="([^"]+\.md)"`)
	matches := hrefRegex.FindAllStringSubmatch(overviewContent, -1)

	overviewPaths := make(map[string]bool)
	for _, match := range matches {
		path := match[1]
		// Only check resource docs paths
		if strings.HasPrefix(path, "/config-connector/docs/reference/resource-docs/") {
			validatePath(t, path, generatedFiles, overviewPaths, "overview.md")
		}
	}

	// 4. Verify all generated files are referenced
	for path := range generatedFiles {
		if !tocPaths[path] {
			t.Errorf("generated file %s is not referenced in _toc.yaml", path)
		}
		if !overviewPaths[path] {
			t.Errorf("generated file %s is not referenced in overview.md", path)
		}
	}

	// 5. Validate URLs in generated files (Optional/integration check)
	// This can be slow, so maybe run it only if a flag is set or parallelize.
	// For this task, I'll run it but with a timeout and concurrency.
	if os.Getenv("VALIDATE_URLS") == "true" {
		validateURLsInGeneratedFiles(t, generatedDocsDir)
	}
}

func validatePath(t *testing.T, path string, generatedFiles, seenPaths map[string]bool, source string) {
	if seenPaths[path] {
		t.Errorf("duplicate path %s found in %s", path, source)
	}
	seenPaths[path] = true

	if !generatedFiles[path] {
		t.Errorf("path %s referenced in %s does not exist in generated output", path, source)
	}
}

func validateURLsInGeneratedFiles(t *testing.T, generatedDocsDir string) {
	t.Log("Validating external URLs in generated files...")
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	sem := make(chan bool, 10) // Limit concurrency

	err := filepath.Walk(generatedDocsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			
			// Find all http/https links
			// Simple regex, might need refinement
			urlRegex := regexp.MustCompile(`https?://[^\s'")]+`)
			urls := urlRegex.FindAllString(string(content), -1)

			for _, url := range urls {
				// Clean up trailing characters that might be captured
				url = strings.TrimRight(url, ".,)>]")
				
				sem <- true
				go func(u, fPath string) {
					defer func() { <-sem }()
					resp, err := client.Head(u)
					if err != nil {
						// Retry with GET if HEAD fails (some servers don't like HEAD)
						resp, err = client.Get(u)
						if err != nil {
							t.Errorf("URL validation failed in %s: %s - %v", fPath, u, err)
							return
						}
						defer resp.Body.Close()
					}
					if resp.StatusCode != 200 {
						t.Errorf("URL invalid in %s: %s returned status %d", fPath, u, resp.StatusCode)
					}
				}(url, path)
			}
		}
		return nil
	})
	if err != nil {
		t.Errorf("error walking generated docs for URL validation: %v", err)
	}

	// Wait for all goroutines to finish (simple way)
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}
