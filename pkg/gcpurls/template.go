package gcpurls

import (
	"fmt"
	"reflect"
	"strings"
)

// URLTemplate handles parsing and formatting of URLs based on a template.
type URLTemplate[T any] struct {
	host                string
	template            string
	segments            []string
	segmentToFieldIndex []int
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

	t.segmentToFieldIndex = make([]int, len(t.segments))
	for i, segment := range t.segments {
		if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
			key := segment[1 : len(segment)-1]
			fieldIdx := -1

			// Find matching field
			// We look for all fields that match case-insensitively.
			// If we find multiple, we panic to avoid ambiguity.
			var matches []int
			for j := 0; j < typ.NumField(); j++ {
				f := typ.Field(j)
				if strings.EqualFold(f.Name, key) {
					matches = append(matches, j)
				}
			}

			if len(matches) == 1 {
				fieldIdx = matches[0]
			} else if len(matches) > 1 {
				panic(fmt.Sprintf("multiple fields match %q in struct %T", key, zero))
			}

			if fieldIdx == -1 {
				panic(fmt.Sprintf("field %q not found in struct %T", key, zero))
			}

			// Verify field is a string
			if typ.Field(fieldIdx).Type.Kind() != reflect.String {
				panic(fmt.Sprintf("field %q in struct %T is not a string", key, zero))
			}

			t.segmentToFieldIndex[i] = fieldIdx
		} else {
			t.segmentToFieldIndex[i] = -1
		}
	}

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
		} else if strings.Contains(s, ".") && strings.HasPrefix(s, "/") {
			// If it looks like it has a host (contains dot and starts with / after scheme removal),
			// but it doesn't match our host, then it's a mismatch.
			// However, simple relative paths like "projects/..." don't have a leading slash usually.
			// If s is "other.com/foo", we should mismatch.

			// We already stripped scheme and //
			// So "compute.googleapis.com/projects/foo" is the input s

			firstSlash := strings.Index(s, "/")
			if firstSlash != -1 {
				potentialHost := s[:firstSlash]
				if strings.Contains(potentialHost, ".") {
					// It has a host part.
					if potentialHost != t.host {
						return nil, false, nil
					}
					s = s[firstSlash:]
				}
			}
		}
	}

	s = strings.Trim(s, "/")
	parts := strings.Split(s, "/")

	if len(parts) != len(t.segments) {
		return nil, false, nil
	}

	var result T
	val := reflect.ValueOf(&result).Elem()

	for i, part := range parts {
		fieldIdx := t.segmentToFieldIndex[i]
		if fieldIdx != -1 {
			// Variable
			f := val.Field(fieldIdx)
			// We checked type in Template()
			f.SetString(part)
		} else {
			// Static
			if part != t.segments[i] {
				return nil, false, nil
			}
		}
	}

	return &result, true, nil
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
