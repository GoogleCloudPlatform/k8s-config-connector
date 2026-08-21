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

	// We basically hand-code the filter logic for now, while we figure out what we need.
	// There's a variety of syntaxes in use.
	if before, after, found := strings.Cut(filter, "network eq \""); found {
		if before != "" {
			return false, fmt.Errorf("filter '%v' not implemented by mockgcp", filter)
		}
		fieldName := "network"
		// Make sure there's just one term in the filter.
		query := strings.TrimSuffix(after, "\"")
		if strings.Contains(query, "\"") {
			return false, fmt.Errorf("filter '%v' not implemented by mockgcp", filter)
		}
		query = strings.TrimPrefix(query, ".*\\b")
		query = strings.TrimSuffix(query, "\\b.*")

		// Some basic unescaping.
		query = strings.ReplaceAll(query, "\\-", "-")

		fd := obj.ProtoReflect().Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fd == nil {
			return false, fmt.Errorf("field '%q' not known", fieldName)
		}
		network := obj.ProtoReflect().Get(fd).String()
		// Technically \b is a word boundary, but we'll just use it as a substring match.
		if !strings.Contains(network, query) {
			return false, nil
		}
		return true, nil
	}
	return false, fmt.Errorf("filter '%v' not implemented by mockgcp", filter)
}
