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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	PhaseCreate      = "CREATE"
	PhaseUpdate      = "UPDATE"
	PhaseRereconcile = "RERECONCILE"
	PhaseDelete      = "DELETE"
	PhaseDependency  = "DEPENDENCY"
	PhaseUnknown     = "UNKNOWN"
)

// Entry represents a single parsed HTTP or gRPC request/response stanza from _http.log.
type Entry struct {
	Phase    string
	Request  HTTPRequest
	Response HTTPResponse
}

// HTTPRequest holds the request method, URL, headers, and raw body of a stanza.
type HTTPRequest struct {
	Method string
	URL    string
	Header http.Header
	Body   string
}

// HTTPResponse holds the response status line, numeric status code, headers, and raw body.
type HTTPResponse struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       string
}

// FormatHTTP renders the single Entry into _http.log stanza format.
func (e *Entry) FormatHTTP() string {
	return e.Request.FormatHTTP() + e.Response.FormatHTTP()
}

// IsMutation reports whether the entry is a state-modifying RPC (and not an LRO poll or read-only RPC).
func (e *Entry) IsMutation() bool {
	if e == nil || e.IsOperationsPoll() || e.IsReadOnly() {
		return false
	}
	switch e.Request.Method {
	case "POST", "PUT", "PATCH", "DELETE", "GRPC":
		return true
	}
	return false
}

// IsReadOnly reports whether the entry is a read-only query or existence check.
func (e *Entry) IsReadOnly() bool {
	if e == nil {
		return false
	}
	switch e.Request.Method {
	case "GET", "HEAD", "OPTIONS":
		return true
	case "POST":
		uPath := strings.ToLower(e.CleanURLPath())
		readOnlyVerbs := []string{
			":getiampolicy", ":testiampermissions", ":fetch", ":query", ":search",
			":lookup", ":check", ":validate", ":export", ":list",
			":generate", ":aggregatedlist", ":count", ":read", ":batchget",
		}
		for _, s := range readOnlyVerbs {
			if strings.HasSuffix(uPath, s) || strings.Contains(uPath, s+"?") {
				return true
			}
		}
	case "GRPC":
		parts := strings.Split(e.Request.URL, "/")
		if len(parts) > 0 {
			m := parts[len(parts)-1]
			if strings.HasPrefix(m, "Get") || strings.HasPrefix(m, "List") ||
				strings.HasPrefix(m, "Test") || strings.HasPrefix(m, "Fetch") ||
				strings.HasPrefix(m, "Check") || strings.HasPrefix(m, "Lookup") {
				return true
			}
		}
	}
	return false
}

// IsDelete reports whether the entry is a deletion RPC.
func (e *Entry) IsDelete() bool {
	if e == nil {
		return false
	}
	if e.Request.Method == "DELETE" {
		return true
	}
	uPath := strings.ToLower(e.CleanURLPath())
	if e.Request.Method == "POST" {
		if strings.HasSuffix(uPath, ":delete") || strings.HasSuffix(uPath, ":destroy") ||
			strings.HasSuffix(uPath, ":cancel") || strings.HasSuffix(uPath, ":drop") {
			return true
		}
	}
	if e.Request.Method == "GRPC" {
		parts := strings.Split(e.Request.URL, "/")
		if len(parts) > 0 {
			m := parts[len(parts)-1]
			if strings.HasPrefix(m, "Delete") || strings.HasPrefix(m, "Destroy") || strings.HasPrefix(m, "Cancel") {
				return true
			}
		}
	}
	return false
}

// IsOperationsPoll reports whether the entry is a Long-Running Operation (LRO) status poll.
func (e *Entry) IsOperationsPoll() bool {
	if e == nil {
		return false
	}
	uPath := e.CleanURLPath()
	if e.Request.Method == "GRPC" {
		return strings.HasSuffix(uPath, "/GetOperation") || strings.HasSuffix(uPath, "/WaitOperation")
	}
	if e.Request.Method != "GET" && e.Request.Method != "POST" {
		return false
	}
	if e.Request.Method == "POST" && !strings.HasSuffix(uPath, ":wait") {
		return false
	}
	segments := strings.Split(strings.Trim(uPath, "/"), "/")
	for i, seg := range segments {
		if (seg == "operations" || seg == "globalOperations" || seg == "zoneOperations" || seg == "regionOperations") && i < len(segments)-1 {
			return true
		}
	}
	if body, ok := e.Response.ParseJSONBody(); ok {
		if kind, isStr := body["kind"].(string); isStr && strings.HasSuffix(kind, "#operation") {
			return true
		}
	}
	return false
}

// ShortEndpoint returns a compact human-readable URL path for report tables.
func (e *Entry) ShortEndpoint() string {
	if e == nil {
		return ""
	}
	uPath, query, hasQuery := strings.Cut(e.Request.URL, "?")
	if protoIdx := strings.Index(uPath, "://"); protoIdx != -1 {
		uPath = uPath[protoIdx+3:]
		if slashIdx := strings.Index(uPath, "/"); slashIdx != -1 {
			uPath = uPath[slashIdx:]
		}
	}

	if idx := strings.Index(uPath, "/locations/"); idx != -1 {
		uPath = "..." + uPath[idx:]
	} else if idx := strings.Index(uPath, "/regions/"); idx != -1 {
		uPath = "..." + uPath[idx:]
	} else if idx := strings.Index(uPath, "/zones/"); idx != -1 {
		uPath = "..." + uPath[idx:]
	} else if idx := strings.Index(uPath, "/global/"); idx != -1 {
		uPath = "..." + uPath[idx:]
	} else if idx := strings.Index(uPath, "/projects/${projectId}/"); idx != -1 {
		uPath = "..." + uPath[idx+len("/projects/${projectId}"):]
	} else if idx := strings.Index(uPath, "/projects/${projectNumber}/"); idx != -1 {
		uPath = "..." + uPath[idx+len("/projects/${projectNumber}"):]
	}

	if hasQuery && query != "" {
		if vals, err := url.ParseQuery(query); err == nil {
			var keys []string
			for k := range vals {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			var kept []string
			for _, k := range keys {
				if k == "alt" || k == "$alt" || k == "prettyPrint" || k == "updateMask" || k == "update_mask" {
					continue
				}
				for _, v := range vals[k] {
					kept = append(kept, fmt.Sprintf("%s=%s", k, v))
				}
			}
			if len(kept) > 0 {
				uPath = uPath + "?" + strings.Join(kept, "&")
			}
		}
	}
	return uPath
}

// ExtractResourceIDSegment extracts the trailing resource identifier or collection segment from the request URL.
func (e *Entry) ExtractResourceIDSegment() string {
	if e == nil {
		return ""
	}
	uPath := e.CleanURLPath()
	if colonIdx := strings.Index(uPath, ":"); colonIdx != -1 {
		uPath = uPath[:colonIdx]
	}
	segments := strings.Split(strings.Trim(uPath, "/"), "/")
	if len(segments) == 0 {
		return ""
	}
	return segments[len(segments)-1]
}

// ExtractUpdateMask extracts the field mask paths from the request URL query parameters or JSON body.
func (e *Entry) ExtractUpdateMask() []string {
	if e == nil {
		return nil
	}
	_, query, found := strings.Cut(e.Request.URL, "?")
	if found && query != "" {
		if vals, err := url.ParseQuery(query); err == nil {
			for _, key := range []string{"updateMask", "update_mask", "fieldMask"} {
				if v := vals.Get(key); v != "" {
					return splitAndTrimMask(v)
				}
			}
		}
	}
	if body, ok := e.Request.ParseJSONBody(); ok {
		for _, key := range []string{"updateMask", "update_mask"} {
			if v, isStr := body[key].(string); isStr && v != "" {
				return splitAndTrimMask(v)
			}
			if m, isMap := body[key].(map[string]any); isMap {
				if paths, isSlice := m["paths"].([]any); isSlice {
					var out []string
					for _, p := range paths {
						if ps, isStr := p.(string); isStr {
							out = append(out, ps)
						}
					}
					return out
				}
			}
		}
	}
	return nil
}

func splitAndTrimMask(v string) []string {
	var out []string
	for _, p := range strings.Split(v, ",") {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// CleanURLPath returns the URL path stripped of scheme, host, and query string.
func (e *Entry) CleanURLPath() string {
	if e == nil {
		return ""
	}
	u := strings.Split(e.Request.URL, "?")[0]
	if protoIdx := strings.Index(u, "://"); protoIdx != -1 {
		u = u[protoIdx+3:]
		if slashIdx := strings.Index(u, "/"); slashIdx != -1 {
			u = u[slashIdx:]
		}
	}
	return u
}

// ParseJSONBody unmarshals the request body if it is a valid JSON object.
func (r HTTPRequest) ParseJSONBody() (map[string]any, bool) {
	return parseJSONObject(r.Body)
}

// FormatHTTP formats the HTTP request part of a log stanza.
func (r HTTPRequest) FormatHTTP() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s %s\n", r.Method, r.URL))
	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		for _, v := range r.Header[k] {
			b.WriteString(fmt.Sprintf("%s: %s\n", k, v))
		}
	}
	b.WriteString("\n")
	if r.Body != "" {
		b.WriteString(r.Body)
		b.WriteString("\n\n")
	}
	return b.String()
}

// ParseJSONBody unmarshals the response body if it is a valid JSON object.
func (r HTTPResponse) ParseJSONBody() (map[string]any, bool) {
	return parseJSONObject(r.Body)
}

// IsSuccess reports whether the response has a 2xx HTTP status code.
func (r HTTPResponse) IsSuccess() bool {
	code := r.ExtractStatusNumber()
	return code >= 200 && code < 300
}

// ExtractStatusNumber returns the numeric HTTP status code for the response.
func (r HTTPResponse) ExtractStatusNumber() int {
	if r.StatusCode > 0 {
		return r.StatusCode
	}
	return parseStatusCode(r.Status, "")
}

// FormatHTTP formats the HTTP response part of a log stanza.
func (r HTTPResponse) FormatHTTP() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s\n", r.Status))
	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		for _, v := range r.Header[k] {
			b.WriteString(fmt.Sprintf("%s: %s\n", k, v))
		}
	}
	b.WriteString("\n")
	if r.Body != "" {
		b.WriteString(r.Body)
		b.WriteString("\n")
	}
	return b.String()
}

func parseJSONObject(raw string) (map[string]any, bool) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" || !strings.HasPrefix(trimmed, "{") {
		return nil, false
	}
	var m map[string]any
	if err := json.Unmarshal([]byte(trimmed), &m); err != nil {
		return nil, false
	}
	return m, true
}
