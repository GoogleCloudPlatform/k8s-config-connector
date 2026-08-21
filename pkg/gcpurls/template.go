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

package gcpurls

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// RegisteredTemplate is an interface for accessing template information.
type RegisteredTemplate interface {
	Host() string
	CanonicalForm() string
}

var (
	registryMu sync.Mutex
	registry   []RegisteredTemplate
)

// AllTemplates returns a snapshot of all registered templates.
func AllTemplates() []RegisteredTemplate {
	registryMu.Lock()
	defer registryMu.Unlock()
	// Return a copy to be safe
	out := make([]RegisteredTemplate, len(registry))
	copy(out, registry)
	return out
}

func register(t RegisteredTemplate) {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry = append(registry, t)
}

type FieldMeta struct {
	segment    string
	fieldIndex []int
}

// URLTemplate handles parsing and formatting of URLs based on a template.
type URLTemplate[T any] struct {
	host     string
	template string
	fields   []FieldMeta
}

// Template creates a new URLTemplate for the given host and template string.
// The template string should look like "projects/{project}/locations/{location}/foos/{foo}".
// The generic type T must be a struct with fields matching the placeholders in the template.
// Field names are matched case-insensitively against the placeholders.
func Template[T any](host, template string) *URLTemplate[T] {
	var zero T
	typ := reflect.TypeOf(zero)
	if typ.Kind() != reflect.Struct {
		panic(fmt.Sprintf("type %T must be a struct", zero))
	}
	// Generate a map of field name to index/path
	nameMap := make(map[string][]int)
	generateNameMap(nameMap, typ, make([]int, 0))

	segments := strings.Split(template, "/")
	fields := make([]FieldMeta, len(segments))
	for i, segment := range segments {
		fields[i].segment = segment
		if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
			key := strings.ToLower(segment[1 : len(segment)-1])
			path, present := nameMap[key]
			if !present {
				panic(fmt.Sprintf("field %q not found in struct %T", key, zero))
			}
			fields[i].fieldIndex = path
		} else {
			fields[i].fieldIndex = []int{-1}
		}
	}
	t := &URLTemplate[T]{
		host:     host,
		template: template,
		fields:   fields,
	}

	register(t)
	return t
}

// Parse parses a URL string into the struct T.
// It returns the parsed struct, a boolean indicating if the URL matched the template, and an error if parsing failed.
func (t *URLTemplate[T]) Parse(s string) (*T, bool, error) {
	// Strip optional scheme
	s = strings.TrimPrefix(s, "https:")
	s = strings.TrimPrefix(s, "http:")
	s = strings.TrimPrefix(s, "//")

	// Check and strip host
	if t.host != "" {
		if strings.HasPrefix(s, t.host+"/") {
			s = strings.TrimPrefix(s, t.host)
		}
	}

	s = strings.Trim(s, "/")
	parts := strings.Split(s, "/")

	if len(parts) != len(t.fields) {
		return nil, false, nil
	}

	var result T
	val := reflect.ValueOf(&result).Elem()
	typ := reflect.TypeOf(result)

	for i, part := range parts {
		fieldIdx := t.fields[i].fieldIndex[0]
		if fieldIdx != -1 {
			// Variable
			if part == "" {
				return nil, false, nil
			}
			f := val.Field(fieldIdx)
			// We checked type in Template()
			if typ.Field(fieldIdx).Type.Kind() == reflect.String {
				f.SetString(part)
			} else {
				s := f.Addr()
				if s.IsNil() {
					panic(fmt.Sprintf("field %q is nil", typ.Field(fieldIdx).Type.Kind()))
				}
				tp := typ.Field(fieldIdx).Type
				c := s.Convert(reflect.PointerTo(tp))
				subIdx := t.fields[i].fieldIndex[1]
				if tp.Field(subIdx).Type.Kind() == reflect.String {
					f.Field(subIdx).SetString(part)
				} else {
					panic(fmt.Sprintf("field %q on %v(%T)(%d) in struct %T is not a string", tp, c, c, c.NumMethod(), result))
				}
			}
		} else {
			// Static
			if part != t.fields[i].segment {
				return nil, false, nil
			}
		}
	}

	return &result, true, nil
}

// CanonicalForm returns the template string.
func (t *URLTemplate[T]) CanonicalForm() string {
	return t.template
}

// Host returns the host.
func (t *URLTemplate[T]) Host() string {
	return t.host
}

// ToString formats the struct T into a URL string.
func (t *URLTemplate[T]) ToString(v T) string {
	var parts []string
	val := reflect.ValueOf(v)

	for _, field := range t.fields {
		segment := field.segment
		fieldIdx := field.fieldIndex[0]
		if fieldIdx != -1 {
			f := val.Field(fieldIdx)
			parts = append(parts, f.String())
		} else {
			parts = append(parts, segment)
		}
	}
	return strings.Join(parts, "/")
}

func generateNameMap(nameMap map[string][]int, structType reflect.Type, path []int) {
	for j := 0; j < structType.NumField(); j++ {
		f := structType.Field(j)
		if structType.Field(j).Type.Kind() == reflect.String {
			name := strings.ToLower(f.Name)
			if _, present := nameMap[name]; present {
				panic(fmt.Sprintf("multiple fields match %q in struct %T", name, structType))
			}
			nameMap[name] = append(path, j)
		} else if structType.Field(j).Type.Kind() == reflect.Struct {
			innerType := structType.Field(j).Type
			innerPath := append(path, j)
			generateNameMap(nameMap, innerType, innerPath)
		} else {
			panic(fmt.Sprintf("field %q in struct %T is not a string or a struct", f.Name, structType))
		}
	}
}
