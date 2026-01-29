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

package preview

import (
	"io"
	"log"
	"strings"
	"sync"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/util/sets"
)

func filterLogs(log logr.Logger) logr.Logger {
	installStandardLogFilter()
	f := &filterSink{sink: log.GetSink()}
	f.IgnoreMessages = sets.New[string]()
	f.IgnoreMessages.Insert("Registered controller")
	f.IgnoreMessages.Insert("Registered deletion-defender controller")
	f.IgnoreMessages.Insert("Starting Controller")
	f.IgnoreMessages.Insert("starting reconcile")
	f.IgnoreMessages.Insert("Starting EventSource")
	f.IgnoreMessages.Insert("Starting workers")
	f.IgnoreMessages.Insert("Shutdown signal received, waiting for all workers to finish")
	f.IgnoreMessages.Insert("All workers finished")
	return log.WithSink(f)
}

type filterSink struct {
	IgnoreMessages sets.Set[string]
	sink           logr.LogSink
}

// Init implements logr.LogSink
func (s *filterSink) Init(info logr.RuntimeInfo) {
	s.sink.Init(info)
}

// Enabled implements logr.LogSink
func (s *filterSink) Enabled(level int) bool {
	return s.sink.Enabled(level)
}

// Info implements logr.LogSink
func (s *filterSink) Info(level int, msg string, args ...any) {
	if s.IgnoreMessages.Has(msg) {
		return
	}
	// Only filter if it looks like a structured log from an upstream library (DCL/TF)
	if strings.HasPrefix(msg, "[DEBUG]") || strings.HasPrefix(msg, "[INFO]") {
		return
	}
	s.sink.Info(level, msg, args...)
}

// WithValues implements logr.LogSink
func (s *filterSink) WithValues(keysAndValues ...any) logr.LogSink {
	return &filterSink{IgnoreMessages: s.IgnoreMessages, sink: s.sink.WithValues(keysAndValues...)}
}

// WithName implements logr.LogSink
func (s *filterSink) WithName(name string) logr.LogSink {
	return &filterSink{IgnoreMessages: s.IgnoreMessages, sink: s.sink.WithName(name)}
}

func (s *filterSink) Error(err error, msg string, args ...any) {
	s.sink.Error(err, msg, args...)
}

var installStandardLogFilterOnce sync.Once

type stdLogFilter struct {
	sink io.Writer
}

func (w *stdLogFilter) Write(p []byte) (n int, err error) {
	msg := string(p)
	if strings.Contains(msg, "Authenticating using configured Google JSON 'access_token'...") {
		return len(p), nil
	}
	// Standard log package usually formats as "2026/01/29 00:00:00 [DEBUG] ..."
	if strings.Contains(msg, " [DEBUG] ") || strings.Contains(msg, " [INFO] ") {
		return len(p), nil
	}
	// Also handle cases where there is no timestamp (e.g. log.SetFlags(0))
	if strings.HasPrefix(msg, "[DEBUG]") || strings.HasPrefix(msg, "[INFO]") {
		return len(p), nil
	}
	return w.sink.Write(p)
}

func installStandardLogFilter() {
	installStandardLogFilterOnce.Do(func() {
		log.SetOutput(&stdLogFilter{sink: log.Writer()})
	})
}
