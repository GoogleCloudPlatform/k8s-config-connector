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
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/klog/v2"
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

func extractEventsWithURLPrefix(allEvents, urlPrefix string) string {
	eventStrings := make([]string, 0)
	events := strings.Split(allEvents, "---")
	for _, event := range events {
		processed := strings.TrimSpace(event)
		for _, m := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
			processed = strings.TrimPrefix(processed, m)
			processed = strings.TrimSpace(processed)
		}
		if strings.HasPrefix(processed, urlPrefix) {
			if len(eventStrings) == 0 {
				event = strings.TrimPrefix(event, "\n\n")
			}
			eventStrings = append(eventStrings, event)
		}
	}
	return strings.Join(eventStrings, "---")
}

// CompareGoldenFile performs a file comparison for a golden test.
func CompareGoldenFile(t *testing.T, p, fullGot string, normalizers ...func(s string) string) {
	writeGoldenOutput := os.Getenv("WRITE_GOLDEN_OUTPUT") != ""

	for _, normalizer := range normalizers {
		fullGot = normalizer(fullGot)
	}
	got := fullGot

	wantBytes, err := os.ReadFile(p)
	if err != nil {
		if writeGoldenOutput && os.IsNotExist(err) {
			// Expected when creating output for the first time;
			// treat as empty
			wantBytes = []byte{} // Not strictly needed, but clearer
		} else if fullGot == "" && os.IsNotExist(err) {
			// Golden file won't be generated if the result is empty.
			return
		} else {
			t.Fatalf("FAIL: failed to read golden file %q: %v", p, err)
		}
	}
	want := string(wantBytes)
	for _, normalizer := range normalizers {
		want = normalizer(want)
	}
	// If urlPrefix is not an empty string, then we should only compare the http
	// log events that has given URL prefix.
	urlPrefix := os.Getenv("ONLY_COMPARE_URL_PREFIX")
	if targetGCP := os.Getenv("E2E_GCP_TARGET"); targetGCP == "mock" && urlPrefix != "" {
		klog.Infof("only comparing events with URL prefix %q in the http log", urlPrefix)

		want = extractEventsWithURLPrefix(want, urlPrefix)
		got = extractEventsWithURLPrefix(got, urlPrefix)
	}

	if want == got {
		if urlPrefix != "" && writeGoldenOutput {
			// Write the full http log to the golden file. The current
			// comparison only covers events with given URL prefix, and the full
			// http log may have diffs.
			if err := os.WriteFile(p, []byte(fullGot), 0644); err != nil {
				t.Fatalf("FAIL: failed to write golden output %s: %v", p, err)
			}
			t.Logf("wrote updated golden output to %s", p)
		}
		return
	}

	if diff := cmp.Diff(want, got); diff != "" {
		onlyWarn := false
		for _, f := range strings.Split(os.Getenv("ONLY_WARN_ON_GOLDEN_DIFFS"), ",") {
			if f == filepath.Base(p) {
				onlyWarn = true
			}
		}

		if onlyWarn {
			t.Logf("found diff in golden output %s, but ONLY_WARN_ON_GOLDEN_DIFFS=%s so will treat as a warning", p, os.Getenv("ONLY_WARN_ON_GOLDEN_DIFFS"))
			t.Logf("unexpected diff in %s: %s", p, diff)
		} else {
			t.Errorf("FAIL: unexpected diff in %s: %s", p, diff)
		}
	}

	if writeGoldenOutput {
		// No matter how we compare the golden files, we should write the entire
		// http log to the golden file.
		if err := os.WriteFile(p, []byte(fullGot), 0644); err != nil {
			t.Fatalf("FAIL: failed to write golden output %s: %v", p, err)
		}
		t.Logf("wrote updated golden output to %s", p)
	}
}

// IgnoreLeadingComments is a normalizer function that strips comments.
// It will stop when it finds the first non-comment line.
// It is intended to be used with CompareGoldenFile.
func IgnoreLeadingComments(s string) string {
	var out []string
	lines := strings.Split(s, "\n")
	removingLeadingLines := true
	commentBlock := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if !removingLeadingLines {
			out = append(out, line)
			continue
		}
		s := strings.TrimSpace(line)
		if strings.HasPrefix(s, "/*") {
			commentBlock++
		}

		ignore := false
		if commentBlock != 0 {
			// Ignore multi-line c-style comment blocks
			ignore = true
		} else if strings.HasPrefix(s, "//") {
			// ignore single-line c-style comments
			ignore = true
		} else if strings.HasPrefix(s, "#") {
			// ignore comments in yaml
			ignore = true
		}

		if !ignore {
			out = append(out, line)
			removingLeadingLines = false
		}

		if strings.HasSuffix(s, "*/") {
			commentBlock--
		}
	}
	return strings.TrimSpace(strings.Join(out, "\n")) + "\n"
}

func CompareGoldenObject(t *testing.T, p string, got []byte) {
	writeGoldenOutput := os.Getenv("WRITE_GOLDEN_OUTPUT") != ""
	want, err := os.ReadFile(p)
	if err != nil {
		if writeGoldenOutput && os.IsNotExist(err) {
			// Expected when creating output for the first time
			if err := os.WriteFile(p, got, 0644); err != nil {
				t.Errorf("failed to write golden output %s: %v", p, err)
			}
			t.Logf("wrote updated golden output to %s", p)
		} else {
			t.Errorf("failed to read golden file %q: %v", p, err)
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

	diff := cmp.Diff(wantMap, gotMap)
	if diff == "" {
		return
	}

	t.Errorf("FAIL: unexpected diff in %s: %s", p, diff)

	if writeGoldenOutput {
		// Write the output to the golden file
		if err := os.WriteFile(p, []byte(got), 0644); err != nil {
			t.Fatalf("failed to write golden output %s: %v", p, err)
		}
		t.Logf("wrote updated golden output to %s", p)
	}
}

func PrettyPrintJSON[T any](t *testing.T, k T) string {
	encoded, err := json.MarshalIndent(k, "", " ")
	if err != nil {
		t.Fatalf("error encoding to json: %v", err)
	}

	return string(encoded)
}

func PrettyPrintYAML[T any](t *testing.T, k T) string {
	encoded, err := json.MarshalIndent(k, "", " ")
	if err != nil {
		t.Fatalf("error encoding to json: %v", err)
	}

	otherK := new(T)
	err = json.Unmarshal(encoded, &otherK)
	if err != nil {
		t.Fatalf("error decoding from json: %v", err)
	}

	yEncoded, err := yaml.Marshal(otherK)
	if err != nil {
		t.Fatalf("error encoding to yaml: %v", err)
	}

	return string(yEncoded)
}
