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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
	"unicode"

	"k8s.io/klog/v2"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	Request   Request   `json:"request,omitempty"`
	Response  Response  `json:"response,omitempty"`
	Error     string    `json:"error,omitempty"`
}

func (e *LogEntry) URL() string {
	return e.Request.URL
}

func (e *LogEntry) Method() string {
	return e.Request.Method
}

// VisitRequestStringValues calls callback for any string values in the request body
func (e *LogEntry) VisitRequestStringValues(callback func(path, value string)) {
	body := e.Request.Body
	if body == "" {
		return
	}

	obj := make(map[string]any)
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		klog.Fatalf("error from json.Unmarshal(%q): %v", body, err)
		return
	}
	visitStringValues(obj, "", callback)
}

// VisitResponseStringValues calls callback for any string values in the response body
func (e *LogEntry) VisitResponseStringValues(callback func(path, value string)) {
	body := e.Response.Body
	if body == "" {
		return
	}

	obj := make(map[string]any)
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		klog.Fatalf("error from json.Unmarshal(%q): %v", body, err)
		return
	}
	visitStringValues(obj, "", callback)
}

// visitStringValues walks the object, building the path for
func visitStringValues(obj any, path string, callback func(path, value string)) {
	switch obj := obj.(type) {
	case map[string]any:
		for k, v := range obj {
			visitStringValues(v, path+"."+k, callback)
		}
	case []any:
		for _, v := range obj {
			visitStringValues(v, path+"[]", callback)
		}
	case string:
		callback(path, obj)
	}
}

type Request struct {
	Method string `json:"method,omitempty"`
	URL    string `json:"url,omitempty"`

	// The HTTP Headers for the request.
	// These should be stored with canonicalized keys (using http.CanonicalHeaderKey(k))
	Header http.Header `json:"header,omitempty"`

	Body string `json:"body,omitempty"`
}

type Response struct {
	Status     string `json:"status,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`

	// The HTTP Headers for the response.
	// These should be stored with canonicalized keys (using http.CanonicalHeaderKey(k))
	Header http.Header `json:"header,omitempty"`

	Body string `json:"body,omitempty"`
}

type HTTPRecorder struct {
	inner http.RoundTripper

	eventSinks []EventSink
}

func NewHTTPRecorder(inner http.RoundTripper, eventSinks ...EventSink) *HTTPRecorder {
	rt := &HTTPRecorder{inner: inner, eventSinks: eventSinks}
	return rt
}

func (r *HTTPRecorder) RoundTrip(req *http.Request) (*http.Response, error) {
	var entry LogEntry
	entry.Timestamp = time.Now()
	entry.Request.Method = req.Method
	entry.Request.URL = req.URL.String()

	entry.Request.Header = make(http.Header)
	for k, values := range req.Header {
		k = http.CanonicalHeaderKey(k)
		switch strings.ToLower(k) {
		case "authorization":
			entry.Request.Header[k] = []string{"(removed)"}
		default:
			entry.Request.Header[k] = values
		}
	}

	if req.Body != nil {
		requestBody, err := io.ReadAll(req.Body)
		if err != nil {
			panic("failed to read request body")
		}
		entry.Request.Body = string(requestBody)
		req.Body = io.NopCloser(bytes.NewReader(requestBody))
	}

	response, err := r.inner.RoundTrip(req)

	if err != nil {
		entry.Error = fmt.Sprintf("%v", err)
	}

	if recordErr := r.record(&entry, req, response); recordErr != nil {
		klog.Warningf("failed to record HTTP request: %v", recordErr)
	}

	return response, err
}

func (r *HTTPRecorder) record(entry *LogEntry, req *http.Request, resp *http.Response) error {
	if resp != nil {
		entry.Response.Status = resp.Status
		entry.Response.StatusCode = resp.StatusCode

		entry.Response.Header = make(http.Header)
		for k, values := range resp.Header {
			k = http.CanonicalHeaderKey(k)
			switch strings.ToLower(k) {
			case "authorization":
				entry.Response.Header[k] = []string{"(removed)"}
			default:
				entry.Response.Header[k] = values
			}
		}

		streaming := false
		if req.URL.Query().Get("watch") == "true" {
			streaming = true
		}

		if streaming {
			entry.Response.Body = "<streaming response not included>"
		} else if resp.Body != nil {
			requestBody, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body for request %q: %w", req.URL, err)
			}
			entry.Response.Body = string(requestBody)
			resp.Body = io.NopCloser(bytes.NewReader(requestBody))
		}
	}

	// If we have event sink(s), write to that sink also
	ctx := req.Context()
	for _, eventSink := range r.eventSinks {
		eventSink.AddHTTPEvent(ctx, entry)
	}

	return nil
}

func sanitizePath(s string) string {
	var out strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			out.WriteRune(r)
		} else {
			out.WriteRune('_')
		}
	}
	return out.String()
}

func (e *LogEntry) FormatHTTP() string {
	var b strings.Builder
	b.WriteString(e.Request.FormatHTTP())
	b.WriteString(e.Response.FormatHTTP())
	return b.String()
}

func (r *Request) FormatHTTP() string {
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

func (r *Response) FormatHTTP() string {
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

type JSONMutator func(url string, obj map[string]any)

func (e *LogEntry) PrettifyJSON(mutators ...JSONMutator) {
	e.Request.PrettifyJSON(mutators...)
	e.Response.PrettifyJSON(e.Request.URL, mutators...)
}

func (r *Response) PrettifyJSON(requestURL string, mutators ...JSONMutator) {
	r.Body = prettifyJSON(r.Body, requestURL, mutators...)
}

func (r *Request) PrettifyJSON(mutators ...JSONMutator) {
	r.Body = prettifyJSON(r.Body, r.URL, mutators...)
}

func prettifyJSON(s string, url string, mutators ...JSONMutator) string {
	if s == "" {
		return s
	}

	obj := make(map[string]any)
	if err := json.Unmarshal([]byte(s), &obj); err != nil {
		klog.Fatalf("error from json.Unmarshal(%q): %v", s, err)
		return s
	}

	for _, mutator := range mutators {
		mutator(url, obj)
	}

	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		klog.Fatalf("error from json.MarshalIndent: %v", err)
		return s
	}
	return string(b)
}

func (r *Request) ReplaceHeader(key, value string) {
	key = http.CanonicalHeaderKey(key)
	r.Header.Set(key, value)
}

func (r *Response) ReplaceHeader(key, value string) {
	key = http.CanonicalHeaderKey(key)
	r.Header.Set(key, value)
}

func (r *Request) AddHeader(key, value string) {
	r.Header.Add(key, value)
}

func (r *Response) AddHeader(key, value string) {
	r.Header.Add(key, value)
}

func (r *Response) RemoveHeader(key string) {
	key = http.CanonicalHeaderKey(key)
	r.Header.Del(key)
}

func (r *Request) RemoveHeader(key string) {
	key = http.CanonicalHeaderKey(key)
	r.Header.Del(key)
}

func (s *Request) ReplaceQueryParameter(key string, value string) {
	// Slightly hacky replacement to preserve order
	url := s.URL
	base, query, found := strings.Cut(url, "?")
	if query == "" || !found {
		return
	}
	parameters := strings.Split(query, "&")
	for i, parameter := range parameters {
		if strings.HasPrefix(parameter, key+"=") {
			parameters[i] = key + "=" + value
		}
	}
	s.URL = base + "?" + strings.Join(parameters, "&")
}

func (r *Response) ParseBody() map[string]any {
	return parseBody(r.Body)
}

func (r *Request) ParseBody() map[string]any {
	return parseBody(r.Body)
}

func parseBody(s string) map[string]any {
	if s == "" {
		return nil
	}
	obj := make(map[string]any)
	if err := json.Unmarshal([]byte(s), &obj); err != nil {
		klog.Fatalf("error from json.Unmarshal(%q): %v", s, err)
		return nil
	}

	return obj
}
