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
	"os"
	"path/filepath"
	"regexp"
	"testing"
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
