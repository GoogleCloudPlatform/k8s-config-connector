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

package resourcefixture

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestGoldenLogAlignment(t *testing.T) {
	rootDir := "testdata/basic"
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		t.Fatalf("failed to get absolute path for %s: %v", rootDir, err)
	}

	err = filepath.WalkDir(absRootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			realLogPath := filepath.Join(path, "_http.log")
			mockLogPath := filepath.Join(path, "_http_mock.log")

			if fileExists(realLogPath) && fileExists(mockLogPath) {
				relPath, _ := filepath.Rel(absRootDir, path)
				t.Run(relPath, func(t *testing.T) {
					compareLogs(t, realLogPath, mockLogPath)
				})
			}
		}

		return nil
	})

	if err != nil {
		t.Fatalf("error walking directory: %v", err)
	}
}

type httpEvent struct {
	Method       string
	URL          string
	RequestBody  string
	Status       string
	ResponseBody string
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// normalizeEvents collapses duplicate GET requests and sorts the events to make the comparison
// resilient to asynchronous execution order and polling variations.
func normalizeEvents(events []httpEvent) []httpEvent {
	var collapsed []httpEvent
	seenGET := make(map[string]bool)

	for _, ev := range events {
		if ev.Method == "GET" {
			cleanedURL := cleanURL(ev.URL)
			key := cleanedURL + "|" + ev.Status
			if seenGET[key] {
				continue
			}
			seenGET[key] = true
		}
		collapsed = append(collapsed, ev)
	}

	sort.Slice(collapsed, func(i, j int) bool {
		cleanedI := cleanURL(collapsed[i].URL)
		cleanedJ := cleanURL(collapsed[j].URL)
		if cleanedI != cleanedJ {
			return cleanedI < cleanedJ
		}
		if collapsed[i].Method != collapsed[j].Method {
			return collapsed[i].Method < collapsed[j].Method
		}
		return collapsed[i].Status < collapsed[j].Status
	})

	return collapsed
}

func compareLogs(t *testing.T, realPath, mockPath string) {
	realEvents := readAndNormalizeLog(t, realPath)
	mockEvents := readAndNormalizeLog(t, mockPath)

	if len(realEvents) != len(mockEvents) {
		reportCountMismatch(t, realEvents, mockEvents)
		return
	}

	compareLogEvents(t, realEvents, mockEvents)
}

func readAndNormalizeLog(t *testing.T, path string) []httpEvent {
	bytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}
	return normalizeEvents(parseLog(t, string(bytes)))
}

func reportCountMismatch(t *testing.T, realEvents, mockEvents []httpEvent) {
	t.Errorf("mismatched number of HTTP events after normalization: real has %d, mock has %d", len(realEvents), len(mockEvents))
	t.Logf("Normalized Real Events:")
	for i, ev := range realEvents {
		t.Logf("  [%d] %s %s -> %s", i, ev.Method, cleanURL(ev.URL), ev.Status)
	}
	t.Logf("Normalized Mock Events:")
	for i, ev := range mockEvents {
		t.Logf("  [%d] %s %s -> %s", i, ev.Method, cleanURL(ev.URL), ev.Status)
	}
}

func compareLogEvents(t *testing.T, realEvents, mockEvents []httpEvent) {
	for i := range realEvents {
		compareEvent(t, i, realEvents[i], mockEvents[i])
	}
}

func compareEvent(t *testing.T, index int, realEv, mockEv httpEvent) {
	if realEv.Method != mockEv.Method {
		t.Errorf("event %d: method mismatch: real=%q, mock=%q", index, realEv.Method, mockEv.Method)
	}

	realURL := cleanURL(realEv.URL)
	mockURL := cleanURL(mockEv.URL)
	if realURL != mockURL {
		t.Errorf("event %d: URL mismatch:\n  real: %s\n  mock: %s", index, realURL, mockURL)
	}

	if realEv.Status != mockEv.Status {
		t.Errorf("event %d: status mismatch: real=%q, mock=%q", index, realEv.Status, mockEv.Status)
	}

	compareJSON(t, fmt.Sprintf("event %d request body", index), realEv.RequestBody, mockEv.RequestBody)
	compareJSON(t, fmt.Sprintf("event %d response body", index), realEv.ResponseBody, mockEv.ResponseBody)
}

var statusRegex = regexp.MustCompile(`^\d{3} `)

func parseLog(t *testing.T, content string) []httpEvent {
	var events []httpEvent
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

func cleanURL(u string) string {
	if protoIdx := strings.Index(u, "://"); protoIdx != -1 {
		u = u[protoIdx+3:]
	}
	if slashIdx := strings.Index(u, "/"); slashIdx != -1 {
		u = u[slashIdx:]
	}
	return u
}

func compareJSON(t *testing.T, context, realJSON, mockJSON string) {
	if realJSON == "" && mockJSON == "" {
		return
	}

	var realObj, mockObj interface{}

	if realJSON != "" {
		if err := json.Unmarshal([]byte(realJSON), &realObj); err != nil {
			if realJSON != mockJSON {
				t.Errorf("%s: string mismatch:\n  real: %q\n  mock: %q", context, realJSON, mockJSON)
			}
			return
		}
	}

	if mockJSON != "" {
		if err := json.Unmarshal([]byte(mockJSON), &mockObj); err != nil {
			if realJSON != mockJSON {
				t.Errorf("%s: string mismatch:\n  real: %q\n  mock: %q", context, realJSON, mockJSON)
			}
			return
		}
	}

	if !reflect.DeepEqual(realObj, mockObj) {
		t.Errorf("%s: JSON mismatch:\n  real: %s\n  mock: %s", context, realJSON, mockJSON)
	}
}
