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

// URLTemplate handles parsing and formatting of URLs based on a template.
type URLTemplate[T any] struct {
	host                string
	template            string
	segments            []string
	segmentToFieldIndex [][]int
}

// Template creates a new URLTemplate for the given host and template string.
// The template string should look like "projects/{project}/locations/{location}/foos/{foo}".
// The generic type T must be a struct with fields matching the placeholders in the template.
// Field names are matched case-insensitively against the placeholders.
func Template[T any](host, template string) *URLTemplate[T] {
	t := &URLTemplate[T]{
		host:     host,
		template: template,
		segments: strings.Split(template, "/"),
	}

	var zero T
	typ := reflect.TypeOf(zero)
	if typ.Kind() != reflect.Struct {
		panic(fmt.Sprintf("type %T must be a struct", zero))
	}

	t.segmentToFieldIndex = make([][]int, len(t.segments))
	for i, segment := range t.segments {
		if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
			list := ""
			key := segment[1 : len(segment)-1]
			if strings.Contains(key, ".") {
				key = strings.Split(key, ".")[0]
			}
			fieldIdx := -1

			// Find matching field
			// We look for all fields that match case-insensitively.
			// If we find multiple, we panic to avoid ambiguity.
			var matches []int
			for j := 0; j < typ.NumField(); j++ {
				f := typ.Field(j)
				list = fmt.Sprintf("%s, %v", list, typ.Field(j))
				if strings.EqualFold(f.Name, key) {
					matches = append(matches, j)
				}
				// Need to deal with compound keys here.
			}

			if len(matches) == 1 {
				fieldIdx = matches[0]
			} else if len(matches) > 1 {
				panic(fmt.Sprintf("multiple fields match %q in struct %T", key, zero))
			}

			if fieldIdx == -1 {
				panic(fmt.Sprintf("field %q not found in struct %T or field list %s", key, zero, list))
			}

			// Verify field is a string
			if typ.Field(fieldIdx).Type.Kind() == reflect.String {
				t.segmentToFieldIndex[i] = []int{fieldIdx}
			} else if typ.Field(fieldIdx).Type.Kind() == reflect.Struct {
				// Should really do this recursively to handle nested structs. Later.
				structType := typ.Field(fieldIdx).Type
				fieldKey := segment[len(key)+2 : len(segment)-1]
				indexList := []int{fieldIdx}
				fieldList := ""
				for k := 0; k < structType.NumField(); k++ {
					field := structType.Field(k)
					fieldList = fmt.Sprintf("%s, %v", fieldList, field)
					if strings.EqualFold(field.Name, fieldKey) {
						indexList = append(indexList, k)
						break
					}
				}
				if len(indexList) == 2 {
					t.segmentToFieldIndex[i] = indexList
				} else {
					panic(fmt.Sprintf("field %q in struct %T is not a string or a struct(%v)", fieldKey, structType, fieldList))
				}
			} else {
				panic(fmt.Sprintf("field %q in struct %T is not a string or a struct", key, zero))
			}
		} else {
			t.segmentToFieldIndex[i] = []int{-1}
		}
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

	if len(parts) != len(t.segments) {
		return nil, false, nil
	}

	var result T
	val := reflect.ValueOf(&result).Elem()
	typ := reflect.TypeOf(result)

	for i, part := range parts {
		fieldIdx := t.segmentToFieldIndex[i][0]
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
				subIdx := t.segmentToFieldIndex[i][1]
				if tp.Field(subIdx).Type.Kind() == reflect.String {
					f.Field(subIdx).SetString(part)
				} else {
					panic(fmt.Sprintf("field %q on %v(%T)(%d) in struct %T is not a string", tp, c, c, c.NumMethod(), result))
				}
			}
		} else {
			// Static
			if part != t.segments[i] {
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

	for i, segment := range t.segments {
		fieldIdx := t.segmentToFieldIndex[i][0]
		if fieldIdx != -1 {
			f := val.Field(fieldIdx)
			parts = append(parts, f.String())
		} else {
			parts = append(parts, segment)
		}
	}
	return strings.Join(parts, "/")
}
