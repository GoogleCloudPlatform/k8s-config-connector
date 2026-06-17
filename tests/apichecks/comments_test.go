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
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
)

func TestAOrAnComments(t *testing.T) {
	t.Parallel()

	apisDir := "../../apis"
	var goFiles []string
	err := filepath.Walk(apisDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("error walking apis dir: %v", err)
	}

	re := regexp.MustCompile(`\b[aA]\s+([AEIOU][A-Za-z0-9_]*)`)

	var errs []string
	fset := token.NewFileSet()

	for _, path := range goFiles {
		content, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("error reading file %s: %v", path, err)
		}

		contentStr := string(content)
		base := filepath.Base(path)
		if strings.HasSuffix(path, ".generated.go") || strings.HasPrefix(base, "zz_generated.") || strings.Contains(contentStr, "Code generated") || strings.Contains(contentStr, "DO NOT EDIT") {
			continue
		}

		file, err := parser.ParseFile(fset, path, content, parser.ParseComments)
		if err != nil {
			t.Fatalf("error parsing file %s: %v", path, err)
		}

		for _, cg := range file.Comments {
			for _, comment := range cg.List {
				text := comment.Text
				matches := re.FindAllStringSubmatch(text, -1)
				if len(matches) > 0 {
					pos := fset.Position(comment.Slash)
					relPath, err := filepath.Rel("../..", pos.Filename)
					if err != nil {
						relPath = pos.Filename
					}

					for _, match := range matches {
						matchedWord := match[1]
						if hasPhoneticException(matchedWord) {
							continue
						}

						errs = append(errs, fmt.Sprintf("%s:%d: comment contains 'a %s', should be 'an %s'", relPath, pos.Line, matchedWord, matchedWord))
					}
				}
			}
		}
	}

	sort.Strings(errs)
	want := strings.Join(errs, "\n")
	if want != "" {
		want += "\n"
	}

	test.CompareGoldenFile(t, "testdata/exceptions/comments.txt", want)
}

func hasPhoneticException(word string) bool {
	lower := strings.ToLower(word)
	prefixes := []string{"url", "uri", "uuid", "utc", "udp", "us", "eu", "user", "unique", "unicode", "usage", "utf", "ui"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(lower, prefix) {
			return true
		}
	}
	return false
}
