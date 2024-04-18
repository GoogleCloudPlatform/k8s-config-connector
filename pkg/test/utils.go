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

package test

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"sigs.k8s.io/yaml"
)

// StringMatchesRegexList is a test utility that returns true
// if the string matches any regex in a list of strings.
//
// if a regex fails to compile, the test will fail.
func StringMatchesRegexList(t *testing.T, regexesToMatch []string, targetString string) bool {
	for _, regexToMatch := range regexesToMatch {
		matcher, err := regexp.Compile(regexToMatch)
		if err != nil {
			t.Fatalf("StringMatchesRegexList: regex '%v' failed to compile", regexToMatch)
		}
		if matcher.MatchString(targetString) {
			return true
		}
	}
	return false
}

// TrimLicenseHeaderFromYaml trims the license header in the yaml string.
func TrimLicenseHeaderFromYaml(yaml string) string {
	r := regexp.MustCompile("(?s)# Copyright.*under the License.\n\n")
	return r.ReplaceAllString(yaml, "")
}

// TrimLicenseHeaderFromTF trims the license header in the tf string.
func TrimLicenseHeaderFromTF(yaml string) string {
	r := regexp.MustCompile("(?s)/\\*\\*\n \\* Copyright.*under the License.\n \\*/\n\n")
	return r.ReplaceAllString(yaml, "")
}

// MustReadFile reads the specified file, failing the test on error.
func MustReadFile(t *testing.T, p string) []byte {
	t.Helper()
	b, err := os.ReadFile(p)
	if err != nil {
		absPath, _ := filepath.Abs(p)
		t.Fatalf("error reading file '%v' (absolute path %v): %v", p, absPath, err)
	}
	return b
}

// MustWriteFile writes the specified file, failing the test on error.
func MustWriteFile(t *testing.T, p string, data []byte, perm os.FileMode) {
	t.Helper()
	if err := os.WriteFile(p, data, perm); err != nil {
		absPath, _ := filepath.Abs(p)
		t.Fatalf("error writing file '%v' (absolute path %v): %v", p, absPath, err)
	}
}

// CompareGoldenFile performs a file comparison for a golden test.
func CompareGoldenFile(t *testing.T, p string, got string, normalizers ...func(s string) string) {
	writeGoldenOutput := os.Getenv("WRITE_GOLDEN_OUTPUT") != ""

	for _, normalizer := range normalizers {
		got = normalizer(got)
	}

	wantBytes, err := os.ReadFile(p)
	if err != nil {
		if writeGoldenOutput && os.IsNotExist(err) {
			// Expected when creating output for the first time;
			// treat as empty
			wantBytes = []byte{} // Not strictly needed, but clearer
		} else {
			t.Fatalf("failed to read golden file %q: %v", p, err)
		}
	}
	want := string(wantBytes)
	for _, normalizer := range normalizers {
		want = normalizer(want)
	}

	if want == got {
		return
	}

	if writeGoldenOutput {
		// Write the output to the golden file
		if err := os.WriteFile(p, []byte(got), 0644); err != nil {
			t.Fatalf("failed to write golden output %s: %v", p, err)
		}
		t.Errorf("wrote updated golden output to %s", p)
	} else {
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("unexpected diff in %s: %s", p, diff)
		}
	}
}

// IgnoreLeadingComments is a normalizer function that strips comments.
// It will stop when it finds the first non-comment line.
// It is intended to be used with CompareGoldenFile.
func IgnoreLeadingComments(s string) string {
	var out []string
	lines := strings.Split(s, "\n")
	removing := true
	commentBlock := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if !removing {
			out = append(out, line)
			continue
		}
		s := strings.TrimSpace(line)
		if strings.HasPrefix(s, "/*") {
			commentBlock++
		}

		if commentBlock != 0 {
			// Ignore multi-line c-style comment blocks
		} else if strings.HasPrefix(s, "//") {
			// ignore single-line c-style comments
		} else if strings.HasPrefix(s, "#") {
			// ignore comments in yaml
		} else {
			out = append(out, line)
			removing = false
		}

		if strings.HasSuffix(s, "*/") {
			commentBlock--
		}
	}
	return strings.TrimSpace(strings.Join(out, "\n")) + "\n"
}

func CompareGoldenObject(t *testing.T, path string, got []byte) {
	writeGoldenOutput := os.Getenv("WRITE_GOLDEN_OUTPUT") != ""
	want, err := os.ReadFile(path)
	if err != nil {
		if writeGoldenOutput && os.IsNotExist(err) {
			// Expected when creating output for the first time
			if err := os.WriteFile(path, got, 0644); err != nil {
				t.Errorf("failed to write golden output %s: %v", path, err)
			}
			t.Logf("wrote updated golden output to %s", path)
		} else {
			t.Errorf("failed to read golden file %q: %v", path, err)
		}
	}
	var wantMap, gotMap map[string]interface{}
	err = yaml.Unmarshal(want, &wantMap)
	if err != nil {
		t.Errorf("Failed parsing file: %s", err)
	}
	err = yaml.Unmarshal(got, &gotMap)
	if err != nil {
		t.Errorf("Failed parsing file: %s", err)
	}
	if ok, err := compareNestedFields(wantMap, gotMap); !ok {
		t.Fatalf("unexpected diff in %s: %s", path, err)
	}
}

func compareNestedFields(wantMap, gotMap map[string]interface{}) (bool, error) {
	for wantKey := range wantMap {
		if _, exists := gotMap[wantKey]; !exists {
			err := fmt.Errorf("field %s in the golden file is missing", wantKey)
			return false, err
		}
	}

	for gotKey := range gotMap {
		if _, exists := wantMap[gotKey]; !exists {
			err := fmt.Errorf("field %s does not exist in golden file", gotKey)
			return false, err
		}
	}

	//  Check nested structures recursively
	for wantKey, wantVal := range wantMap {
		if gotVal, exists := gotMap[wantKey]; exists {
			if wantNestedMap, ok1 := wantVal.(map[string]interface{}); ok1 {
				if gotNestedMap, ok2 := gotVal.(map[string]interface{}); ok2 {
					if ok, err := compareNestedFields(wantNestedMap, gotNestedMap); !ok {
						return false, err
					}
				}
			}
		}
	}
	return true, nil
}
