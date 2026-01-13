package gcpurls

import (
	"fmt"
	"reflect"
	"strings"
)

// URLTemplate handles parsing and formatting of URLs based on a template.
type URLTemplate[T any] struct {
	template            string
	segments            []string
	segmentToFieldIndex []int
}

// Template creates a new URLTemplate for the given template string.
// The template string should look like "projects/{project}/locations/{location}/foos/{foo}".
// The generic type T must be a struct with fields matching the placeholders in the template.
// Field names are matched case-insensitively against the placeholders.
func Template[T any](template string) *URLTemplate[T] {
	t := &URLTemplate[T]{
		template: template,
		segments: strings.Split(template, "/"),
	}

	var zero T
	typ := reflect.TypeOf(zero)
	if typ.Kind() != reflect.Struct {
		panic(fmt.Sprintf("type %T must be a struct", zero))
	}

	t.segmentToFieldIndex = make([]int, len(t.segments))
	for i, segment := range t.segments {
		if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
			key := segment[1 : len(segment)-1]
			fieldIdx := -1

			// 1. Exact match
			if f, ok := typ.FieldByName(key); ok {
				fieldIdx = f.Index[0] // Assuming flat struct for now
			} else {
				// 2. Case-insensitive match
				for j := 0; j < typ.NumField(); j++ {
					f := typ.Field(j)
					if strings.EqualFold(f.Name, key) {
						fieldIdx = j
						break
					}
				}
			}

			if fieldIdx == -1 {
				panic(fmt.Sprintf("field %q not found in struct %T", key, zero))
			}
			t.segmentToFieldIndex[i] = fieldIdx
		} else {
			t.segmentToFieldIndex[i] = -1
		}
	}

	return t
}

// Parse parses a URL string into the struct T.
// It ignores the host prefix (e.g. //foo.googleapis.com/) if present.
func (t *URLTemplate[T]) Parse(s string) (*T, error) {
	// Strip optional scheme
	s = strings.TrimPrefix(s, "https:")
	s = strings.TrimPrefix(s, "http:")
	s = strings.TrimPrefix(s, "//")

	// Strip host if present.
	// Heuristic: if the string starts with a host-like part (contains dot) followed by a slash, strip it.
	// We check if the remaining parts match the template structure (length and static segments).

	s = strings.Trim(s, "/")
	parts := strings.Split(s, "/")

	if len(parts) > len(t.segments) {
		offset := len(parts) - len(t.segments)
		// Check if the static segments match with this offset
		match := true
		for i, idx := range t.segmentToFieldIndex {
			if idx == -1 { // Static
				if parts[i+offset] != t.segments[i] {
					match = false
					break
				}
			}
		}
		if match {
			parts = parts[offset:]
		}
	}

	if len(parts) != len(t.segments) {
		return nil, fmt.Errorf("url %q does not match template %q: segment count mismatch (got %d, expected %d)", s, t.template, len(parts), len(t.segments))
	}

	var result T
	val := reflect.ValueOf(&result).Elem()

	for i, part := range parts {
		fieldIdx := t.segmentToFieldIndex[i]
		if fieldIdx != -1 {
			// Variable
			f := val.Field(fieldIdx)
			if f.Kind() == reflect.String {
				f.SetString(part)
			} else {
				// This should be caught by Template() check, but safe to check here too or assume string
				return nil, fmt.Errorf("field at index %d is not a string", fieldIdx)
			}
		} else {
			// Static
			if part != t.segments[i] {
				return nil, fmt.Errorf("url %q does not match template %q: expected %q got %q at index %d", s, t.template, t.segments[i], part, i)
			}
		}
	}

	return &result, nil
}

// ToString formats the struct T into a URL string.
func (t *URLTemplate[T]) ToString(v T) (string, error) {
	var parts []string
	val := reflect.ValueOf(v)

	for i, segment := range t.segments {
		fieldIdx := t.segmentToFieldIndex[i]
		if fieldIdx != -1 {
			f := val.Field(fieldIdx)
			parts = append(parts, f.String())
		} else {
			parts = append(parts, segment)
		}
	}
	return strings.Join(parts, "/"), nil
}
