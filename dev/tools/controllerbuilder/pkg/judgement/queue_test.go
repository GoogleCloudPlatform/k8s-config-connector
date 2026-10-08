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

package judgement

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func resourceEntry(field, reason string) Entry {
	return Entry{Kind: "Foo", Group: "example.cnrm.cloud.google.com", Field: field, Reason: reason, Status: StatusOpen}
}

func TestMergeKeepsResolvedEntries(t *testing.T) {
	// Arrange
	resolved := resourceEntry(".spec.a", "possible-reference")
	resolved.Detail = "old detail"
	resolved.Status = StatusResolved
	resolved.Resolution = ResolutionEdited
	resolved.Note = "changed to aRef"
	q := &Queue{Entries: []Entry{resolved}}

	regenerated := resourceEntry(".spec.a", "possible-reference")
	regenerated.Detail = "new detail"
	added := resourceEntry(".spec.b", "possible-reference")

	// Act
	q.Merge([]Entry{regenerated, added, added})

	// Assert
	if len(q.Entries) != 2 {
		t.Fatalf("got %d entries, want 2: %+v", len(q.Entries), q.Entries)
	}
	got := q.Entries[0]
	if got.Status != StatusResolved || got.Resolution != ResolutionEdited || got.Note != "changed to aRef" {
		t.Errorf("regenerating reopened or changed a resolved entry: %+v", got)
	}
	if got.Detail != "new detail" {
		t.Errorf("detail = %q, want the regenerated detail", got.Detail)
	}
	if q.Entries[1].Field != ".spec.b" || !q.Entries[1].IsOpen() {
		t.Errorf("new entry = %+v, want an open .spec.b", q.Entries[1])
	}
}

func TestMergeKeepsEntriesNotRegenerated(t *testing.T) {
	// generate.sh calls generate-types once per proto version. A later call
	// must not drop what an earlier call wrote.
	q := &Queue{Entries: []Entry{resourceEntry(".spec.a", "x")}}

	q.Merge([]Entry{{Kind: "Bar", Group: "example.cnrm.cloud.google.com", Reason: "y"}})

	if len(q.Entries) != 2 || q.Entries[0].Field != ".spec.a" {
		t.Errorf("entries = %+v, want the earlier entry kept first", q.Entries)
	}
}

func TestEntryValidate(t *testing.T) {
	for _, tc := range []struct {
		name    string
		entry   Entry
		wantErr string
	}{
		{name: "open resource entry", entry: resourceEntry(".spec.a", "x")},
		{name: "message entry", entry: Entry{ProtoMessage: "google.cloud.x.v1.Msg", Field: "f", Reason: "x", Status: StatusOpen}},
		{name: "accepted needs no note", entry: Entry{Kind: "Foo", Group: "g", Reason: "x", Status: StatusResolved, Resolution: ResolutionAccepted}},
		{name: "edited with note", entry: Entry{Kind: "Foo", Group: "g", Reason: "x", Status: StatusResolved, Resolution: ResolutionEdited, Note: "n"}},
		{name: "no reason", entry: Entry{Kind: "Foo", Group: "g", Status: StatusOpen}, wantErr: "no reason"},
		{name: "kind without group", entry: Entry{Kind: "Foo", Reason: "x", Status: StatusOpen}, wantErr: "both kind and group"},
		{name: "no subject", entry: Entry{Reason: "x", Status: StatusOpen}, wantErr: "kind and group, or protoMessage"},
		{name: "both subjects", entry: Entry{Kind: "Foo", Group: "g", ProtoMessage: "m", Reason: "x", Status: StatusOpen}, wantErr: "both kind/group and protoMessage"},
		{name: "no status", entry: Entry{Kind: "Foo", Group: "g", Reason: "x"}, wantErr: "unknown status"},
		{name: "open with resolution", entry: Entry{Kind: "Foo", Group: "g", Reason: "x", Status: StatusOpen, Resolution: ResolutionAccepted}, wantErr: "open entry has a resolution"},
		{name: "resolved without resolution", entry: Entry{Kind: "Foo", Group: "g", Reason: "x", Status: StatusResolved}, wantErr: "no resolution"},
		{name: "edited without note", entry: Entry{Kind: "Foo", Group: "g", Reason: "x", Status: StatusResolved, Resolution: ResolutionEdited}, wantErr: "needs a note"},
		{name: "unknown resolution", entry: Entry{Kind: "Foo", Group: "g", Reason: "x", Status: StatusResolved, Resolution: "fixed"}, wantErr: "unknown resolution"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.entry.Validate()
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("err = %v, want it to contain %q", err, tc.wantErr)
			}
		})
	}
}

func TestQueueValidateRejectsDuplicates(t *testing.T) {
	q := &Queue{Entries: []Entry{resourceEntry(".spec.a", "x"), resourceEntry(".spec.a", "x")}}
	if err := q.Validate(); err == nil || !strings.Contains(err.Error(), "duplicate") {
		t.Fatalf("err = %v, want a duplicate error", err)
	}
}

func TestWriteThenRead(t *testing.T) {
	// Arrange
	path := filepath.Join(t.TempDir(), FileName)
	want := &Queue{Entries: []Entry{
		resourceEntry(".spec.a", "possible-reference"),
		{ProtoMessage: "google.cloud.x.v1.Msg", Field: "f", Reason: "unsupported-field-type", Detail: "d: with a colon", Status: StatusOpen},
		{Kind: "Foo", Group: "g", Reason: "untriaged-bulk-generation", Status: StatusResolved, Resolution: ResolutionAccepted},
	}}

	// Act
	if err := Write(path, want); err != nil {
		t.Fatalf("Write: %v", err)
	}
	got, err := Read(path)

	// Assert
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip changed the queue:\n got %+v\nwant %+v", got, want)
	}
	data, _ := os.ReadFile(path)
	if !strings.HasPrefix(string(data), "# Judgement queue") {
		t.Errorf("file does not start with the header:\n%s", data)
	}
}

func TestReadMissingFileIsEmpty(t *testing.T) {
	q, err := Read(filepath.Join(t.TempDir(), FileName))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(q.Entries) != 0 {
		t.Errorf("entries = %+v, want none", q.Entries)
	}
}

func TestReadRejectsUnknownFieldsAndBadEntries(t *testing.T) {
	for _, tc := range []struct {
		name    string
		content string
		wantErr string
	}{
		{
			name:    "misspelled field",
			content: "entries:\n- kind: Foo\n  group: g\n  reason: x\n  status: open\n  resolutoin: accepted\n",
			wantErr: "resolutoin",
		},
		{
			name:    "invalid entry",
			content: "entries:\n- kind: Foo\n  group: g\n  reason: x\n  status: resolved\n",
			wantErr: "no resolution",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), FileName)
			if err := os.WriteFile(path, []byte(tc.content), 0644); err != nil {
				t.Fatal(err)
			}
			_, err := Read(path)
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("err = %v, want it to contain %q", err, tc.wantErr)
			}
			if !strings.Contains(err.Error(), path) {
				t.Errorf("error should name the file, got: %v", err)
			}
		})
	}
}
