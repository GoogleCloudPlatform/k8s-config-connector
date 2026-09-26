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

package codegen

import (
	"bytes"
	"slices"
	"strings"
	"testing"
)

func TestUnsupportedFieldMarker(t *testing.T) {
	for _, tc := range []struct {
		name       string
		rendered   string
		wantField  string
		wantReason string
		wantOK     bool
	}{
		{
			name:       "a marker names the field and the reason",
			rendered:   "\n\t// TODO: attributes: unsupported map type with key string and value message\n\n",
			wantField:  "attributes",
			wantReason: "unsupported map type with key string and value message",
			wantOK:     true,
		},
		{
			name:       "a marker without a field name still reports the reason",
			rendered:   "\n\t// TODO: unsupported map type\n\n",
			wantReason: "unsupported map type",
			wantOK:     true,
		},
		{
			name:     "an ordinary field is not a marker",
			rendered: "\t// +kcc:proto:field=pkg.Msg.name\n\tName *string `json:\"name,omitempty\"`\n",
		},
		{
			name:     "a doc comment mentioning TODO is not a marker",
			rendered: "\t// Some description with TODO in prose\n\tName *string `json:\"name,omitempty\"`\n",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			field, reason, ok := UnsupportedFieldMarker(tc.rendered)

			// Assert
			if ok != tc.wantOK || field != tc.wantField || reason != tc.wantReason {
				t.Errorf("UnsupportedFieldMarker() = (%q, %q, %v), want (%q, %q, %v)",
					field, reason, ok, tc.wantField, tc.wantReason, tc.wantOK)
			}
		})
	}
}

func TestScanUnsupportedReportsEveryMarker(t *testing.T) {
	// Arrange
	body := "type Msg struct {\n" +
		"\t// +kcc:proto:field=pkg.Msg.name\n" +
		"\tName *string `json:\"name,omitempty\"`\n" +
		"\n\t// TODO: labels: unsupported map type\n\n" +
		"\n\t// TODO: tags: unsupported map type\n\n" +
		"}\n"

	// Act
	got := scanUnsupported("pkg.Msg", body)

	// Assert
	want := []UnsupportedField{
		{Message: "pkg.Msg", Field: "labels", Reason: "unsupported map type"},
		{Message: "pkg.Msg", Field: "tags", Reason: "unsupported map type"},
	}
	if !slices.Equal(got, want) {
		t.Errorf("scanUnsupported() = %+v, want %+v", got, want)
	}
}

// TestWriteFieldMarkerIsScopedToPrepopulating verifies that field names are only
// included in TODO markers when Prepopulating is true, preserving backwards-compatible
// output for services that have not enabled the flag.
func TestWriteFieldMarkerIsScopedToPrepopulating(t *testing.T) {
	msg := observedStateTestMessage(t)
	// by_index is a map keyed by int32, which GoTypeForField declines.
	field := msg.Fields().ByName("by_index")

	for _, tc := range []struct {
		name string
		opts WriteOptions
		want string
	}{
		{
			name: "opted in, the marker names the field",
			opts: WriteOptions{Prepopulating: true},
			want: "// TODO: byIndex: unsupported map type",
		},
		{
			name: "not opted in, the marker is what it has always been",
			opts: WriteOptions{},
			want: "// TODO: unsupported map type",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			var buf bytes.Buffer
			WriteField(&buf, field, msg, 0, false, tc.opts)

			// Assert
			if !strings.Contains(buf.String(), tc.want) {
				t.Errorf("WriteField() = %q, want it to contain %q", buf.String(), tc.want)
			}
		})
	}
}
