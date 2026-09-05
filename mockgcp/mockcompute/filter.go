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
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// matchFilter checks if the filter matches the object.
// It supports a subset of the filter syntax.
func matchFilter(filter string, obj proto.Message) (bool, error) {
	if filter == "" {
		return true, nil
	}

	// Basic support for (condition) AND (condition)
	if strings.Contains(filter, " AND ") {
		parts := strings.Split(filter, " AND ")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			for strings.HasPrefix(part, "(") && strings.HasSuffix(part, ")") {
				part = part[1 : len(part)-1]
			}
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

	// Support for property="value" or property eq "value"
	var fieldName, query string
	if before, after, found := strings.Cut(filter, " eq \""); found {
		fieldName = strings.TrimSpace(before)
		query = strings.TrimSuffix(after, "\"")
	} else if before, after, found := strings.Cut(filter, "=\""); found {
		fieldName = strings.TrimSpace(before)
		query = strings.TrimSuffix(after, "\"")
	}

	if fieldName != "" {
		for strings.HasPrefix(fieldName, "(") {
			fieldName = fieldName[1:]
		}
		// Some basic unescaping.
		query = strings.ReplaceAll(query, "\\-", "-")
		isRegex := false
		if strings.HasPrefix(query, ".*\\b") && strings.HasSuffix(query, "\\b.*") {
			query = strings.TrimPrefix(query, ".*\\b")
			query = strings.TrimSuffix(query, "\\b.*")
			isRegex = true
		}

		var fd protoreflect.FieldDescriptor
		fields := obj.ProtoReflect().Descriptor().Fields()
		for i := 0; i < fields.Len(); i++ {
			f := fields.Get(i)
			if string(f.Name()) == fieldName || f.JSONName() == fieldName {
				fd = f
				break
			}
		}

		if fd == nil {
			return false, fmt.Errorf("field %q not known in %v", fieldName, obj.ProtoReflect().Descriptor().FullName())
		}

		val := obj.ProtoReflect().Get(fd).String()
		if isRegex {
			if !strings.Contains(val, query) {
				return false, nil
			}
		} else {
			if val != query {
				return false, nil
			}
		}
		return true, nil
	}

	return false, fmt.Errorf("filter %q not implemented by mockgcp", filter)
}
