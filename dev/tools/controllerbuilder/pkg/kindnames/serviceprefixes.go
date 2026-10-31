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

package kindnames

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// FileName is the name of the hand-maintained file of approved service
// prefixes.
const FileName = "service-prefixes.yaml"

// ExceptionsFileName is the name of the hand-maintained file of proposed
// service prefixes waiting for approval.
const ExceptionsFileName = "service-prefix-exceptions.yaml"

// Entry is the service prefix of one API.
type Entry struct {
	// API is the API name, e.g. sqladmin.
	API string `yaml:"api"`
	// Title is the API's service config title, for reference.
	Title string `yaml:"title,omitempty"`
	// ServicePrefix is the prefix of the API's Kinds, e.g. SQLAdmin. A Kind is
	// ServicePrefix followed by the singular resource type.
	ServicePrefix string `yaml:"servicePrefix"`
	// Reason says why the API needs a decision (required in
	// ExceptionsFileName) or what was decided and where (optional in
	// FileName).
	Reason string `yaml:"reason,omitempty"`
	// Pending is true for an entry in ExceptionsFileName: the prefix is
	// waiting for approval, and generate-types refuses to scaffold resources
	// of the API until it is approved (moved to FileName). It is set from the
	// file the entry is in, not written in it.
	Pending bool `yaml:"-"`
}

// File is the YAML document of either file.
type File struct {
	Services []Entry `yaml:"services"`
}

// Table is the service prefix of each API, read from both files.
type Table struct {
	byAPI map[string]Entry
}

// Load reads and validates both files in dir. A missing file is treated as
// empty.
func Load(dir string) (*Table, error) {
	t := &Table{byAPI: map[string]Entry{}}
	var errs []error
	for _, file := range []struct {
		name    string
		pending bool
	}{
		{FileName, false},
		{ExceptionsFileName, true},
	} {
		p := filepath.Join(dir, file.name)
		data, err := os.ReadFile(p)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		var f File
		if err := yaml.Unmarshal(data, &f); err != nil {
			return nil, fmt.Errorf("parsing %s: %w", p, err)
		}
		for i, e := range f.Services {
			where := fmt.Sprintf("%s: entry %d (api %q)", p, i+1, e.API)
			switch {
			case e.API == "":
				errs = append(errs, fmt.Errorf("%s: api is required", where))
			case e.ServicePrefix == "":
				errs = append(errs, fmt.Errorf("%s: servicePrefix is required", where))
			case file.pending && e.Reason == "":
				errs = append(errs, fmt.Errorf("%s: reason is required", where))
			}
			if prev, dup := t.byAPI[e.API]; dup {
				if prev.Pending != file.pending {
					errs = append(errs, fmt.Errorf("%s: api is also in %s; an API is either approved or waiting for approval", where, FileName))
				} else {
					errs = append(errs, fmt.Errorf("%s: api is listed twice", where))
				}
			}
			e.Pending = file.pending
			t.byAPI[e.API] = e
		}
	}
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}
	return t, nil
}

// Lookup returns the entry for an API name, e.g. sqladmin.
func (t *Table) Lookup(api string) (Entry, bool) {
	if t == nil {
		return Entry{}, false
	}
	e, ok := t.byAPI[api]
	return e, ok
}

// Kind returns the Kind for a resource of an API: the API's service prefix
// followed by the singular type of resourceType (a google.api.resource type
// such as sqladmin.googleapis.com/Instance).
func (t *Table) Kind(api, resourceType string) (string, bool) {
	e, ok := t.Lookup(api)
	if !ok {
		return "", false
	}
	typeName := ResourceType(resourceType)
	if typeName == "" {
		return "", false
	}
	return e.ServicePrefix + typeName, true
}
