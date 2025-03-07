// Copyright 2024 Google LLC
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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

// An EventSink listens for various events we are able to capture during tests,
// currently just http requests/responses.
type EventSink interface {
	AddHTTPEvent(ctx context.Context, entry *LogEntry)
}

type httpEventSinkType int

var httpEventSinkKey httpEventSinkType

// EventSinksFromContext gets the EventSink listeners attached to the passed context.
func EventSinksFromContext(ctx context.Context) []EventSink {
	v := ctx.Value(httpEventSinkKey)
	if v == nil {
		return nil
	}
	return v.([]EventSink)
}

// AddSinkToContext attaches the sinks to the returned context.
func AddSinkToContext(ctx context.Context, sinks ...EventSink) context.Context {
	var eventSinks []EventSink
	v := ctx.Value(httpEventSinkKey)
	if v != nil {
		eventSinks = v.([]EventSink)
	}
	eventSinks = append(eventSinks, sinks...)
	return context.WithValue(ctx, httpEventSinkKey, eventSinks)
}

func NewMemoryEventSink() *MemoryEventSink {
	return &MemoryEventSink{}
}

// MemoryEventSink is an EventSink that stores events in memory
type MemoryEventSink struct {
	mutex      sync.Mutex
	HTTPEvents []*LogEntry `json:"httpEvents,omitempty"`
}

func (s *MemoryEventSink) AddHTTPEvent(ctx context.Context, entry *LogEntry) { //nolint:revive
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.HTTPEvents = append(s.HTTPEvents, entry)
}

func (s LogEntries) FormatHTTP() string {
	var eventStrings []string
	for _, entry := range s {
		s := entry.FormatHTTP()
		eventStrings = append(eventStrings, s)
	}
	return strings.Join(eventStrings, "\n---\n\n")
}

type LogEntries []*LogEntry

func (s *LogEntries) PrettifyJSON(mutators ...JSONMutator) {
	for _, entry := range *s {
		entry.PrettifyJSON(mutators...)
	}
}

func (s *LogEntries) RemoveHTTPRequestHeader(key string) {
	for _, entry := range *s {
		entry.Request.RemoveHeader(key)
	}
}

func (s *LogEntries) RemoveHTTPResponseHeader(key string) {
	for _, entry := range *s {
		entry.Response.RemoveHeader(key)
	}
}

func (s *LogEntries) ReplaceRequestQueryParameter(key string, value string) {
	for _, entry := range *s {
		entry.Request.ReplaceQueryParameter(key, value)
	}
}

// KeepIf returns a new LogEntries with only the entries that satisfy the predicate.
// (where the predicate function returns true)
func (s LogEntries) KeepIf(pred func(e *LogEntry) bool) LogEntries {
	var keep LogEntries
	for _, entry := range s {
		if pred(entry) {
			keep = append(keep, entry)
		}
	}
	return keep
}

type DirectoryEventSink struct {
	outputDir string

	// mutex to avoid concurrent writes to the same file
	mutex sync.Mutex
}

func NewDirectoryEventSink(outputDir string) *DirectoryEventSink {
	return &DirectoryEventSink{outputDir: outputDir}
}

func (r *DirectoryEventSink) AddHTTPEvent(ctx context.Context, entry *LogEntry) {
	// Write to a log file
	t := FromContext(ctx)
	testName := "unknown"
	if t != nil {
		testName = t.Name()
	}
	dirName := sanitizePath(testName)
	p := filepath.Join(r.outputDir, dirName, "requests.log")

	if err := r.writeToFile(p, entry); err != nil {
		klog.Fatalf("error writing http event: %v", err)
	}
}

func (r *DirectoryEventSink) writeToFile(p string, entry *LogEntry) error {
	b, err := yaml.Marshal(entry)
	if err != nil {
		klog.Warningf("failed to marshal data as yaml in DirectoryEventSink: %v", err)
		// As a special fallback, write it in JSON so we can try to understand what is going wrong here
		b, err = json.MarshalIndent(entry, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal entry %+v as JSON or YAML: %w", entry, err)
		}
		// return fmt.Errorf("failed to marshal data: %w", err)
	}

	// Just in case we are writing to the same file concurrently
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", filepath.Dir(p), err)
	}
	f, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %w", p, err)
	}
	defer f.Close()

	if _, err := f.Write(b); err != nil {
		return fmt.Errorf("failed to write to file %q: %w", p, err)
	}
	delimeter := "\n\n---\n\n"
	if _, err := f.Write([]byte(delimeter)); err != nil {
		return fmt.Errorf("failed to write to file %q: %w", p, err)
	}

	return nil
}
