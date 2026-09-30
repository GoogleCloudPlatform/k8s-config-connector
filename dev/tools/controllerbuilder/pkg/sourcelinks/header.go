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

package sourcelinks

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

const (
	headerPrefix = "// API sources for "
	headerSuffix = ", recorded by generate-types:"
	markerPrefix = "// +kcc:source:"
	guessPrefix  = "// +kcc:guess=" + GuessKind + " reason="
)

// Render returns the comment block that goes at the top of <kind>_types.go,
// between the license and the package clause:
//
//	// API sources for BigtableInstance, recorded by generate-types:
//	// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/<sha>/google/bigtable/admin/v2/instance.proto
//	// +kcc:source:service-docs=https://cloud.google.com/bigtable/
//	// +kcc:guess=source-link reason=verify-resource-docs-link
//	// +kcc:source:resource-docs=https://docs.cloud.google.com/bigtable/docs/reference/rest/v2/projects.instances
//
// A link that was not verified gets a +kcc:guess line directly above it.
// The block has no trailing newline.
func Render(kind string, links []Link) string {
	var b strings.Builder
	b.WriteString(headerPrefix + kind + headerSuffix)
	for _, l := range links {
		if l.Guess != "" {
			b.WriteString("\n" + guessPrefix + l.Guess)
		}
		fmt.Fprintf(&b, "\n%s%s=%s", markerPrefix, l.Key, l.URL)
	}
	return b.String()
}

// Header is what Parse finds in a types file.
type Header struct {
	// Kind is from the "API sources for <Kind>" line. It is empty when the
	// file has no source links.
	Kind string
	// Links holds the markers in file order. A link's Guess is the reason on
	// the +kcc:guess line directly above it. Detail is always empty.
	Links []Link
}

// Parse reads the source link block of a types file.
func Parse(src []byte) Header {
	var h Header
	pendingGuess := ""
	s := bufio.NewScanner(bytes.NewReader(src))
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		switch {
		case strings.HasPrefix(line, headerPrefix) && strings.HasSuffix(line, headerSuffix):
			h.Kind = strings.TrimSuffix(strings.TrimPrefix(line, headerPrefix), headerSuffix)
		case strings.HasPrefix(line, guessPrefix):
			pendingGuess = strings.TrimPrefix(line, guessPrefix)
			continue
		case strings.HasPrefix(line, markerPrefix):
			key, value, _ := strings.Cut(strings.TrimPrefix(line, markerPrefix), "=")
			h.Links = append(h.Links, Link{Key: key, URL: value, Guess: pendingGuess})
		case strings.HasPrefix(line, "package "):
			return h
		}
		pendingGuess = ""
	}
	return h
}

// Guesses returns the reasons on the file's source-link guess lines.
func (h Header) Guesses() []string {
	var out []string
	for _, l := range h.Links {
		if l.Guess != "" {
			out = append(out, l.Guess)
		}
	}
	return out
}
