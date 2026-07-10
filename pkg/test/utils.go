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
	"fmt"
	"os"
	"path/filepath"
	"reflect"
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
	if writeGoldenOutput && filepath.Base(p) == "_http.log" && os.Getenv("E2E_GCP_TARGET") != "real" {
		t.Fatalf("FAIL: attempted to write/update %q while E2E_GCP_TARGET=%q. _http.log must only be recorded when running against real GCP (E2E_GCP_TARGET=real).", p, os.Getenv("E2E_GCP_TARGET"))
	}

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
	if filepath.Base(p) == "_http.log" && strings.Contains(want, "(mockgcp)") {
		t.Fatalf("FAIL: golden file %q contains mock traffic signature '(mockgcp)'. _http.log must only be recorded against real GCP (E2E_GCP_TARGET=real).", p)
	}
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

	if strings.Contains(filepath.Base(p), "_http") {
		if err := compareHTTPLogs(want, got); err == nil {
			return
		}
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

func compareHTTPLogs(wantContent, gotContent string) error {
	wantEvents := parseLog(wantContent)
	gotEvents := parseLog(gotContent)

	wantGrouped := make(map[string][]httpEvent)
	for _, ev := range wantEvents {
		key := extractResourceKey(ev.URL)
		wantGrouped[key] = append(wantGrouped[key], ev)
	}

	gotGrouped := make(map[string][]httpEvent)
	for _, ev := range gotEvents {
		key := extractResourceKey(ev.URL)
		gotGrouped[key] = append(gotGrouped[key], ev)
	}

	// Compare the groups
	if len(wantGrouped) != len(gotGrouped) {
		return fmt.Errorf("mismatch in number of target resources: want %d resources, got %d", len(wantGrouped), len(gotGrouped))
	}

	for key, wantEvs := range wantGrouped {
		gotEvs, ok := gotGrouped[key]
		if !ok {
			return fmt.Errorf("resource %q not found in actual HTTP log", key)
		}

		if len(wantEvs) != len(gotEvs) {
			return fmt.Errorf("resource %q: mismatched number of HTTP calls: want %d, got %d", key, len(wantEvs), len(gotEvs))
		}

		for i := 0; i < len(wantEvs); i++ {
			w := wantEvs[i]
			g := gotEvs[i]

			if w.Method != g.Method {
				return fmt.Errorf("resource %q call %d: method mismatch: want %q, got %q", key, i, w.Method, g.Method)
			}
			// Clean query parameters before comparing URLs
			wURL := strings.Split(cleanURL(w.URL), "?")[0]
			gURL := strings.Split(cleanURL(g.URL), "?")[0]
			if wURL != gURL {
				return fmt.Errorf("resource %q call %d: URL path mismatch: want %q, got %q", key, i, wURL, gURL)
			}

			// Compare request bodies
			if err := compareJSONStrings(w.RequestBody, g.RequestBody); err != nil {
				return fmt.Errorf("resource %q call %d request body: %w", key, i, err)
			}

			// Compare response bodies
			if err := compareJSONStrings(w.ResponseBody, g.ResponseBody); err != nil {
				return fmt.Errorf("resource %q call %d response body: %w", key, i, err)
			}
		}
	}

	return nil
}

func parseLog(content string) []httpEvent {
	var events []httpEvent
	statusRegex := regexp.MustCompile(`^\d{3} `)
	rawEvents := strings.Split(content, "\n---\n")

	for _, raw := range rawEvents {
		raw = strings.TrimSpace(raw)
		if raw == "" {
			continue
		}

		lines := strings.Split(raw, "\n")
		var ev httpEvent

		reqParts := strings.SplitN(lines[0], " ", 2)
		if len(reqParts) < 2 {
			continue
		}
		ev.Method = reqParts[0]
		ev.URL = reqParts[1]

		idx := 1
		// Skip request headers
		for idx < len(lines) && strings.TrimSpace(lines[idx]) != "" {
			idx++
		}
		if idx < len(lines) {
			idx++
		}

		var reqBodyLines []string
		for idx < len(lines) && !statusRegex.MatchString(lines[idx]) {
			reqBodyLines = append(reqBodyLines, lines[idx])
			idx++
		}
		ev.RequestBody = strings.TrimSpace(strings.Join(reqBodyLines, "\n"))

		if idx < len(lines) {
			ev.Status = lines[idx]
			idx++
		}

		// Skip response headers
		for idx < len(lines) && strings.TrimSpace(lines[idx]) != "" {
			idx++
		}
		if idx < len(lines) {
			idx++
		}

		var respBodyLines []string
		for idx < len(lines) {
			respBodyLines = append(respBodyLines, lines[idx])
			idx++
		}
		ev.ResponseBody = strings.TrimSpace(strings.Join(respBodyLines, "\n"))

		events = append(events, ev)
	}

	return events
}

func extractResourceKey(urlPath string) string {
	u := strings.Split(urlPath, "?")[0]
	u = cleanURL(u)
	segments := strings.Split(strings.Trim(u, "/"), "/")
	if len(segments) == 0 {
		return ""
	}
	lastSeg := segments[len(segments)-1]
	if colonIdx := strings.Index(lastSeg, ":"); colonIdx != -1 {
		lastSeg = lastSeg[:colonIdx]
	}
	if len(segments) >= 2 {
		prevSeg := segments[len(segments)-2]
		return prevSeg + "/" + lastSeg
	}
	return lastSeg
}

func cleanURL(u string) string {
	if protoIdx := strings.Index(u, "://"); protoIdx != -1 {
		u = u[protoIdx+3:]
	}
	if slashIdx := strings.Index(u, "/"); slashIdx != -1 {
		u = u[slashIdx:]
	}
	return u
}

func compareJSONStrings(wantJSON, gotJSON string) error {
	if wantJSON == gotJSON {
		return nil
	}
	if wantJSON == "" || gotJSON == "" {
		return fmt.Errorf("mismatch: want %q, got %q", wantJSON, gotJSON)
	}

	var wantObj, gotObj interface{}
	err1 := json.Unmarshal([]byte(wantJSON), &wantObj)
	err2 := json.Unmarshal([]byte(gotJSON), &gotObj)
	if err1 != nil || err2 != nil {
		if wantJSON != gotJSON {
			return fmt.Errorf("string mismatch:\n  want: %q\n  got:  %q", wantJSON, gotJSON)
		}
		return nil
	}

	if !reflect.DeepEqual(wantObj, gotObj) {
		return fmt.Errorf("JSON mismatch:\n  diff: %s", cmp.Diff(wantObj, gotObj))
	}
	return nil
}

type httpEvent struct {
	Method       string
	URL          string
	RequestBody  string
	Status       string
	ResponseBody string
}
