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

package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"gopkg.in/yaml.v2"
)

// The path to Config Connector resource reference docs starts with
// "/config-connector/docs/reference/resource-docs/".
var docPathPrefix = "/config-connector/docs/reference/resource-docs/"

type TOC struct {
	TOC []TOCItem `yaml:"toc"`
}

type TOCItem struct {
	Title   string       `yaml:"title"`
	Section []TOCSection `yaml:"section,omitempty"`
	Path    string       `yaml:"path,omitempty"`
}

type TOCSection struct {
	Title string `yaml:"title"`
	Path  string `yaml:"path"`
}

func TestReferenceDocConsistency(t *testing.T) {
	rootDir, err := repo.GetRoot()
	if err != nil {
		t.Fatalf("error getting repo root: %v", err)
	}

	baseDir := filepath.Join(rootDir, "scripts", "generate-google3-docs", "resource-reference")
	generatedDocsDir := filepath.Join(baseDir, "generated", "resource-docs")
	templatesDir := filepath.Join(baseDir, "templates")
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
			webPath := fmt.Sprintf("%s%s", docPathPrefix, relPath)
			generatedFiles[webPath] = true
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error walking generated docs dir: %v", err)
	}

	// 2. Validate _toc.yaml
	tocData, err := os.ReadFile(tocPath)
	if err != nil {
		t.Fatalf("error reading _toc.yaml: %v", err)
	}
	var toc TOC
	if err := yaml.Unmarshal(tocData, &toc); err != nil {
		t.Fatalf("error parsing _toc.yaml: %v", err)
	}

	tocPaths := make(map[string]bool)
	for _, item := range toc.TOC {
		if item.Path != "" {
			validatePath(t, item.Path, generatedFiles, tocPaths, "_toc.yaml")
		}
		for _, section := range item.Section {
			validatePath(t, section.Path, generatedFiles, tocPaths, "_toc.yaml")
		}
	}

	// 3. Validate overview.md
	overviewData, err := os.ReadFile(overviewPath)
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

	// 5. Validate URLs in templates (Optional/integration check)
	if os.Getenv("VALIDATE_URLS") != "true" {
		t.Log("Skipping URL validation: VALIDATE_URLS not set to true")
	} else {
		validateURLsInTemplates(t, templatesDir)
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

func validateURLsInTemplates(t *testing.T, templatesDir string) {
	t.Log("Validating external URLs in template files...")

	uniqueURLs := make(map[string][]string)
	// Regex to capture href attribute value
	hrefRegex := regexp.MustCompile(`href="([^"]+)"`)

	err := filepath.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".tmpl") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			matches := hrefRegex.FindAllStringSubmatch(string(content), -1)
			for _, match := range matches {
				url := match[1]

				if strings.HasPrefix(url, "/") || strings.HasPrefix(url, "http") {
					uniqueURLs[url] = append(uniqueURLs[url], path)
				}
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error collecting URLs from templates: %v", err)
	}

	t.Logf("Found %d unique URLs to validate in templates.", len(uniqueURLs))

	type job struct {
		originalURL string
		candidates  []string
	}

	type result struct {
		originalURL string
		candidates  []string
		err         error
		code        int
	}

	jobs := make(chan job, len(uniqueURLs))
	results := make(chan result, len(uniqueURLs))
	concurrency := 20
	var wg sync.WaitGroup

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        concurrency,
			MaxIdleConnsPerHost: concurrency,
		},
	}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := range jobs {
				res := result{originalURL: j.originalURL, candidates: j.candidates}
				var lastErr error
				var lastCode int
				success := false

				for _, u := range j.candidates {
					resp, err := client.Head(u)
					if err != nil || (resp != nil && resp.StatusCode != 200) {
						if resp != nil {
							resp.Body.Close()
							lastCode = resp.StatusCode
						}
						resp, err = client.Get(u)
					}

					if err != nil {
						lastErr = err
					} else {
						lastCode = resp.StatusCode
						resp.Body.Close()
						if lastCode == 200 {
							success = true
							res.code = 200
							break
						}
					}
				}
				
				if !success {
					res.err = lastErr
					res.code = lastCode
				} else {
					res.err = nil // Clear error if at least one candidate succeeded
				}

				results <- res
			}
		}()
	}

	for u := range uniqueURLs {
		var candidates []string
		if strings.HasPrefix(u, "/") {
			candidates = []string{
				"https://cloud.google.com" + u,
				"https://docs.cloud.google.com" + u,
			}
		} else {
			candidates = []string{u}
		}
		jobs <- job{originalURL: u, candidates: candidates}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	failures := 0
	for res := range results {
		if res.err != nil {
			t.Errorf("Failed to fetch %s (candidates: %v): %v (referenced in: %v)", res.originalURL, res.candidates, res.err, uniqueURLs[res.originalURL][0])
			failures++
		} else if res.code != 200 {
			t.Errorf("URL %s (candidates: %v) returned status %d (referenced in: %v)", res.originalURL, res.candidates, res.code, uniqueURLs[res.originalURL][0])
			failures++
		}
	}

	if failures > 0 {
		t.Errorf("Found %d invalid URLs in templates", failures)
	}
}
