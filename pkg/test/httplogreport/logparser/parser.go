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

package logparser

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var (
	phaseCommentRegex   = regexp.MustCompile(`(?i)^#\s*Phase:\s*([A-Z0-9_-]+)\s*$`)
	httpStatusLineRegex = regexp.MustCompile(`^(1\d\d|2\d\d|3\d\d|4\d\d|5\d\d)(\s+.*)?$`)
	httpProtoLineRegex  = regexp.MustCompile(`^HTTP/\d(\.\d)?\s+(\d{3})(\s+.*)?$`)
)

// Parse parses the contents of a golden _http.log or _http_mock.log file into a slice of Entry objects.
// If the log contains "# Phase: <PHASE>" comments, the phase is parsed and sticky across subsequent stanzas.
func Parse(content string) ([]*Entry, error) {
	normalized := strings.ReplaceAll(content, "\r\n", "\n")
	if strings.TrimSpace(normalized) == "" {
		return []*Entry{}, nil
	}

	rawStanzas := splitStanzas(normalized)
	entries := make([]*Entry, 0, len(rawStanzas))
	var currentPhase string

	for stanzaIdx, raw := range rawStanzas {
		trimmed := strings.TrimSpace(raw)
		if trimmed == "" {
			continue
		}

		lines := strings.Split(trimmed, "\n")
		lineIdx := 0

		for lineIdx < len(lines) {
			l := strings.TrimSpace(lines[lineIdx])
			if l == "" {
				lineIdx++
				continue
			}
			if strings.HasPrefix(l, "#") {
				if m := phaseCommentRegex.FindStringSubmatch(l); len(m) == 2 {
					currentPhase = strings.ToUpper(strings.TrimSpace(m[1]))
				}
				lineIdx++
				continue
			}
			break
		}

		if lineIdx >= len(lines) {
			continue
		}

		reqLine := strings.TrimSpace(lines[lineIdx])
		lineIdx++
		reqParts := strings.SplitN(reqLine, " ", 2)
		if len(reqParts) < 2 {
			return nil, fmt.Errorf("stanza %d: invalid request line %q", stanzaIdx+1, reqLine)
		}

		entry := &Entry{
			Phase: currentPhase,
			Request: HTTPRequest{
				Method: strings.TrimSpace(reqParts[0]),
				URL:    strings.TrimSpace(reqParts[1]),
				Header: make(http.Header),
			},
			Response: HTTPResponse{
				Header: make(http.Header),
			},
		}

		for lineIdx < len(lines) && strings.TrimSpace(lines[lineIdx]) != "" {
			hdrLine := lines[lineIdx]
			if colonIdx := strings.Index(hdrLine, ":"); colonIdx > 0 {
				k := http.CanonicalHeaderKey(strings.TrimSpace(hdrLine[:colonIdx]))
				v := strings.TrimSpace(hdrLine[colonIdx+1:])
				entry.Request.Header.Add(k, v)
			}
			lineIdx++
		}

		if lineIdx < len(lines) && strings.TrimSpace(lines[lineIdx]) == "" {
			lineIdx++
		}

		statusLineIdx := findResponseStatusLine(lines, lineIdx, entry.Request.Method)
		if statusLineIdx == -1 {
			reqBody := strings.TrimSpace(strings.Join(lines[lineIdx:], "\n"))
			entry.Request.Body = reqBody
			entries = append(entries, entry)
			continue
		}

		if statusLineIdx > lineIdx {
			reqBody := strings.TrimSpace(strings.Join(lines[lineIdx:statusLineIdx], "\n"))
			entry.Request.Body = reqBody
		}

		statusLine := strings.TrimSpace(lines[statusLineIdx])
		entry.Response.Status = statusLine
		entry.Response.StatusCode = parseStatusCode(statusLine, entry.Request.Method)
		lineIdx = statusLineIdx + 1

		for lineIdx < len(lines) && strings.TrimSpace(lines[lineIdx]) != "" {
			hdrLine := lines[lineIdx]
			if strings.HasPrefix(strings.TrimSpace(hdrLine), "{") || strings.HasPrefix(strings.TrimSpace(hdrLine), "[") {
				break
			}
			if colonIdx := strings.Index(hdrLine, ":"); colonIdx > 0 {
				k := http.CanonicalHeaderKey(strings.TrimSpace(hdrLine[:colonIdx]))
				v := strings.TrimSpace(hdrLine[colonIdx+1:])
				entry.Response.Header.Add(k, v)
			}
			lineIdx++
		}

		if lineIdx < len(lines) && strings.TrimSpace(lines[lineIdx]) == "" {
			lineIdx++
		}

		if lineIdx < len(lines) {
			respBody := strings.TrimSpace(strings.Join(lines[lineIdx:], "\n"))
			entry.Response.Body = respBody
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

// StripPhaseComments removes "# Phase: ..." comment lines from _http.log content.
func StripPhaseComments(content string) string {
	normalized := strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(normalized, "\n")
	var keep []string
	for _, line := range lines {
		if phaseCommentRegex.MatchString(strings.TrimSpace(line)) {
			continue
		}
		keep = append(keep, line)
	}
	return strings.Join(keep, "\n")
}

// FormatWithPhaseComments formats a slice of Entry into _http.log format
// with "# Phase: <PHASE>" comments emitted at each phase transition.
func FormatWithPhaseComments(entries []*Entry) string {
	var eventStrings []string
	var prevPhase string
	for _, entry := range entries {
		formatted := entry.FormatHTTP()
		if entry.Phase != "" && entry.Phase != prevPhase {
			formatted = fmt.Sprintf("# Phase: %s\n%s", entry.Phase, formatted)
			prevPhase = entry.Phase
		}
		eventStrings = append(eventStrings, formatted)
	}
	return strings.Join(eventStrings, "\n---\n\n")
}

func splitStanzas(normalized string) []string {
	lines := strings.Split(normalized, "\n")
	var stanzas []string
	var cur []string
	for _, line := range lines {
		if strings.TrimSpace(line) == "---" {
			stanzas = append(stanzas, strings.Join(cur, "\n"))
			cur = nil
			continue
		}
		cur = append(cur, line)
	}
	if len(cur) > 0 {
		stanzas = append(stanzas, strings.Join(cur, "\n"))
	}
	return stanzas
}

func findResponseStatusLine(lines []string, startIdx int, method string) int {
	braceDepth := 0
	inString := false
	escaped := false

	for i := startIdx; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])
		prevBlank := (i == startIdx) || (strings.TrimSpace(lines[i-1]) == "")

		if braceDepth == 0 && prevBlank && isStatusLine(trimmed, method) {
			return i
		}

		for _, ch := range lines[i] {
			if escaped {
				escaped = false
				continue
			}
			if ch == '\\' && inString {
				escaped = true
				continue
			}
			if ch == '"' {
				inString = !inString
				continue
			}
			if !inString {
				switch ch {
				case '{', '[':
					braceDepth++
				case '}', ']':
					if braceDepth > 0 {
						braceDepth--
					}
				}
			}
		}
	}

	for i := startIdx; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])
		prevBlank := (i == startIdx) || (strings.TrimSpace(lines[i-1]) == "")
		if prevBlank && isStatusLine(trimmed, method) {
			return i
		}
	}
	return -1
}

func isStatusLine(trimmed string, method string) bool {
	if trimmed == "" {
		return false
	}
	if httpStatusLineRegex.MatchString(trimmed) || httpProtoLineRegex.MatchString(trimmed) {
		return true
	}
	if method == "GRPC" {
		if trimmed == "OK" || strings.HasPrefix(trimmed, "error:") {
			return true
		}
	}
	return false
}

func parseStatusCode(statusLine string, method string) int {
	if m := httpStatusLineRegex.FindStringSubmatch(statusLine); len(m) >= 2 {
		if code, err := strconv.Atoi(m[1]); err == nil {
			return code
		}
	}
	if m := httpProtoLineRegex.FindStringSubmatch(statusLine); len(m) >= 3 {
		if code, err := strconv.Atoi(m[2]); err == nil {
			return code
		}
	}
	if method == "GRPC" {
		if statusLine == "OK" {
			return 200
		}
		if strings.HasPrefix(statusLine, "error:") {
			lower := strings.ToLower(statusLine)
			switch {
			case strings.Contains(lower, "notfound") || strings.Contains(lower, "not_found"):
				return 404
			case strings.Contains(lower, "alreadyexists") || strings.Contains(lower, "already_exists"):
				return 409
			case strings.Contains(lower, "invalidargument") || strings.Contains(lower, "failedprecondition"):
				return 400
			case strings.Contains(lower, "permissiondenied") || strings.Contains(lower, "unauthenticated"):
				return 403
			default:
				return 500
			}
		}
	}
	return 0
}
