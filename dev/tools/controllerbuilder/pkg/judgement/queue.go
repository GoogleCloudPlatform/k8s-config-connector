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

// Package judgement reads, merges and writes the per-service judgement queue,
// apis/<service>/judgement_queue.yaml.
//
// generate-types writes an entry for each call it could not make with
// confidence. A person reviews each entry and marks it resolved. Entries are
// never deleted, so the file is also the record of what was decided. The
// generator and tests/apichecks both use this package, so they agree on the
// format.
package judgement

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// FileName is the queue file in each apis/<service> directory.
const FileName = "judgement_queue.yaml"

// Status says whether an entry still needs a person.
type Status string

const (
	StatusOpen     Status = "open"
	StatusResolved Status = "resolved"
)

// Resolution says what the person did with a resolved entry.
type Resolution string

const (
	// ResolutionAccepted means the generated output was kept as it is.
	ResolutionAccepted Resolution = "accepted"
	// ResolutionEdited means the type was changed by hand. The note says how.
	ResolutionEdited Resolution = "edited"
	// ResolutionDeferred means the finding is right but is handled elsewhere,
	// for example in refs_deferred.txt. The note says where.
	ResolutionDeferred Resolution = "deferred"
	// ResolutionNotApplicable means the finding is wrong for this resource.
	// The note says why.
	ResolutionNotApplicable Resolution = "not-applicable"
)

// Entry is one call the generator could not make with confidence.
//
// An entry is about a resource (Kind and Group are set) or about a nested
// proto message that several resources may share (ProtoMessage is set). Field
// is the KRM path for a resource entry, such as ".spec.network", or the proto
// field name for a message entry. It is empty when the entry is about the
// whole resource.
type Entry struct {
	Kind         string     `yaml:"kind,omitempty"`
	Group        string     `yaml:"group,omitempty"`
	ProtoMessage string     `yaml:"protoMessage,omitempty"`
	Field        string     `yaml:"field,omitempty"`
	Reason       string     `yaml:"reason"`
	Detail       string     `yaml:"detail,omitempty"`
	Status       Status     `yaml:"status"`
	Resolution   Resolution `yaml:"resolution,omitempty"`
	Note         string     `yaml:"note,omitempty"`
}

// Key identifies an entry across regenerations. Detail is left out, so
// rewording a detail does not create a second entry.
func (e Entry) Key() string {
	return e.Kind + "|" + e.Group + "|" + e.ProtoMessage + "|" + e.Field + "|" + e.Reason
}

// IsOpen reports whether the entry still needs a person.
func (e Entry) IsOpen() bool {
	return e.Status == StatusOpen
}

// Validate checks that an entry is complete.
func (e Entry) Validate() error {
	if e.Reason == "" {
		return fmt.Errorf("entry has no reason")
	}
	hasResource := e.Kind != "" || e.Group != ""
	switch {
	case hasResource && e.ProtoMessage != "":
		return fmt.Errorf("entry sets both kind/group and protoMessage")
	case hasResource && (e.Kind == "" || e.Group == ""):
		return fmt.Errorf("entry must set both kind and group")
	case !hasResource && e.ProtoMessage == "":
		return fmt.Errorf("entry must set kind and group, or protoMessage")
	}
	switch e.Status {
	case StatusOpen:
		if e.Resolution != "" {
			return fmt.Errorf("open entry has a resolution %q; set status: resolved", e.Resolution)
		}
	case StatusResolved:
		switch e.Resolution {
		case ResolutionAccepted:
		case ResolutionEdited, ResolutionDeferred, ResolutionNotApplicable:
			if e.Note == "" {
				return fmt.Errorf("resolution %q needs a note saying what was done", e.Resolution)
			}
		case "":
			return fmt.Errorf("resolved entry has no resolution")
		default:
			return fmt.Errorf("unknown resolution %q", e.Resolution)
		}
	default:
		return fmt.Errorf("unknown status %q", e.Status)
	}
	return nil
}

// Queue is the contents of one judgement_queue.yaml.
type Queue struct {
	Entries []Entry `yaml:"entries"`
}

// Validate checks every entry and rejects two entries with the same key.
func (q *Queue) Validate() error {
	seen := map[string]bool{}
	for i, e := range q.Entries {
		if err := e.Validate(); err != nil {
			return fmt.Errorf("entry %d (reason %q): %w", i, e.Reason, err)
		}
		if seen[e.Key()] {
			return fmt.Errorf("entry %d (reason %q): duplicate of an earlier entry", i, e.Reason)
		}
		seen[e.Key()] = true
	}
	return nil
}

// Merge adds generated entries to the queue.
//
// An entry the queue already has keeps its status, resolution and note, and
// takes the new detail. So regenerating never reopens a resolved entry. A new
// entry is appended as open. Entries that were not generated this time are
// kept: generate.sh calls generate-types several times per service, and each
// call only knows about its own resources.
func (q *Queue) Merge(generated []Entry) {
	index := map[string]int{}
	for i, e := range q.Entries {
		index[e.Key()] = i
	}
	for _, g := range generated {
		if i, ok := index[g.Key()]; ok {
			q.Entries[i].Detail = g.Detail
			continue
		}
		g.Status = StatusOpen
		g.Resolution = ""
		g.Note = ""
		index[g.Key()] = len(q.Entries)
		q.Entries = append(q.Entries, g)
	}
}

// Read loads a queue file. A missing file is an empty queue.
func Read(path string) (*Queue, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return &Queue{}, nil
	}
	if err != nil {
		return nil, err
	}
	q := &Queue{}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(q); err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	if err := q.Validate(); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return q, nil
}

// header is written at the top of every queue file.
const header = `# Judgement queue for this service, written by controllerbuilder generate-types.
#
# Each entry is a call the generator could not make with confidence. Review it,
# then set status: resolved and a resolution:
#   accepted        the generated output is right as it is
#   edited          you changed the type; the note says how
#   deferred        right, but handled elsewhere; the note says where
#   not-applicable  wrong for this resource; the note says why
#
# Do not delete entries. The file is also the record of what was decided, and
# regenerating keeps the status of every entry it already has.
`

// Write saves a queue file with its header.
func Write(path string, q *Queue) error {
	var buf bytes.Buffer
	buf.WriteString(header)
	buf.WriteString("\n")
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(q); err != nil {
		return err
	}
	if err := enc.Close(); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0644)
}
