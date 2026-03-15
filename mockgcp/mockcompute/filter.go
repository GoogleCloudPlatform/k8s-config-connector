// Copyright 2025 Google LLC
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

package mockcompute

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	// Matches field eq "value" or field eq value, or field="value", or field:value
	filterRegex = regexp.MustCompile(`^\s*([0-9A-Za-z_.]+)\s*(?:eq|=|\:)\s*("[^"]*"|\S+)\s*$`)
)

// matchFilter checks if the filter matches the object.
// It supports a subset of the filter syntax.
func matchFilter(filter string, obj proto.Message) (bool, error) {
	filter = strings.TrimSpace(filter)
	if filter == "" {
		return true, nil
	}

	// Strip balanced parentheses
	for strings.HasPrefix(filter, "(") && strings.HasSuffix(filter, ")") {
		inner := filter[1 : len(filter)-1]
		if isBalanced(inner) {
			filter = strings.TrimSpace(inner)
		} else {
			break
		}
	}

	// Support AND by splitting, but being careful about quotes and parentheses
	parts := splitAnd(filter)
	if len(parts) > 1 {
		for _, part := range parts {
			match, err := matchFilter(part, obj)
			if err != nil {
				return false, err
			}
			if !match {
				return false, nil
			}
		}
		return true, nil
	}

	// Single expression
	match := filterRegex.FindStringSubmatch(filter)
	if match != nil {
		fieldName := match[1]
		value := match[2]
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = value[1 : len(value)-1]
		}
		return matchField(fieldName, value, obj)
	}

	return false, fmt.Errorf("filter '%v' not implemented by mockgcp", filter)
}

func isBalanced(s string) bool {
	depth := 0
	for _, c := range s {
		if c == '(' {
			depth++
		} else if c == ')' {
			depth--
			if depth < 0 {
				return false
			}
		}
	}
	return depth == 0
}

func splitAnd(filter string) []string {
	var parts []string
	var current strings.Builder
	inQuote := false
	parenDepth := 0

	for i := 0; i < len(filter); i++ {
		c := filter[i]
		if c == '"' {
			inQuote = !inQuote
		}
		if !inQuote {
			if c == '(' {
				parenDepth++
			} else if c == ')' {
				parenDepth--
			}
		}

		if !inQuote && parenDepth == 0 && strings.HasPrefix(filter[i:], " AND ") {
			parts = append(parts, strings.TrimSpace(current.String()))
			current.Reset()
			i += 4 // Skip " AND"
			continue
		}
		current.WriteByte(c)
	}
	parts = append(parts, strings.TrimSpace(current.String()))
	return parts
}

func matchField(fieldName string, query string, obj proto.Message) (bool, error) {
	// Some basic unescaping and regex cleaning if it's a regex filter masquerading as eq
	// (KCC/Terraform sometimes does this)
	query = strings.TrimPrefix(query, ".*\\b")
	query = strings.TrimSuffix(query, "\\b.*")
	query = strings.ReplaceAll(query, "\\-", "-")

	parts := strings.Split(fieldName, ".")
	var val protoreflect.Value
	var fd protoreflect.FieldDescriptor
	curr := obj.ProtoReflect()

	for i, part := range parts {
		fd = curr.Descriptor().Fields().ByName(protoreflect.Name(part))
		if fd == nil {
			fd = curr.Descriptor().Fields().ByJSONName(part)
		}
		if fd == nil {
			return false, fmt.Errorf("field %q (part of %q) not known in %v", part, fieldName, curr.Descriptor().FullName())
		}
		val = curr.Get(fd)
		if i < len(parts)-1 {
			if fd.Kind() != protoreflect.MessageKind || fd.IsList() {
				return false, fmt.Errorf("field %q (part of %q) is not a nested message", part, fieldName)
			}
			curr = val.Message()
		}
	}

	if fd.IsList() {
		// Basic support for list of strings
		list := val.List()
		for i := 0; i < list.Len(); i++ {
			if matchValue(fd, list.Get(i), query) {
				return true, nil
			}
		}
		return false, nil
	}

	if matchValue(fd, val, query) {
		return true, nil
	}
	return false, nil
}

func matchValue(fd protoreflect.FieldDescriptor, val protoreflect.Value, query string) bool {
	var strVal string
	switch fd.Kind() {
	case protoreflect.StringKind:
		strVal = val.String()
	case protoreflect.EnumKind:
		enumDesc := fd.Enum().Values().ByNumber(val.Enum())
		if enumDesc != nil {
			strVal = string(enumDesc.Name())
		} else {
			strVal = fmt.Sprintf("%d", val.Enum())
		}
	case protoreflect.BoolKind:
		strVal = strconv.FormatBool(val.Bool())
	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind:
		strVal = strconv.FormatInt(val.Int(), 10)
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind:
		strVal = strconv.FormatUint(val.Uint(), 10)
	default:
		strVal = fmt.Sprintf("%v", val.Interface())
	}

	// Support both exact match and substring match (for regex-like queries)
	if strVal == query {
		return true
	}
	if strings.Contains(strVal, query) {
		return true
	}

	// Normalize GCP links (ignore version like v1 or beta)
	normalizeLink := func(s string) string {
		if !strings.HasPrefix(s, "https://www.googleapis.com/compute/") && !strings.HasPrefix(s, "https://compute.googleapis.com/compute/") {
			return s
		}
		// Replace /compute/v1/ or /compute/beta/ with /compute/ANY/
		re := regexp.MustCompile(`/(?:v1|beta|alpha)/`)
		return re.ReplaceAllString(s, "/ANY/")
	}

	if normalizeLink(strVal) == normalizeLink(query) {
		return true
	}

	return false
}
