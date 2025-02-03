// Copyright 2024 Google LLC
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

package typeupdater

import (
	"go/ast"
	"strings"
	"unicode"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/codegen"
)

// commentContains checks if the given comment group contains a target string annotation
func commentContains(cg *ast.CommentGroup, target string) bool {
	if cg == nil {
		return false
	}
	for _, c := range cg.List {
		trimmed := strings.TrimPrefix(c.Text, "//")
		trimmed = strings.TrimSpace(trimmed)
		if trimmed == target {
			return true
		}
	}
	return false
}

// getProtoFieldName converts a fully qualified proto field name to a snake_case field name
// e.g. "google.cloud.bigquery.datatransfer.v1.TransferConfig.DisplayName" -> "display_name"
func getProtoFieldName(fullName string) string {
	parts := strings.Split(fullName, ".")
	if len(parts) == 0 {
		return ""
	}
	lastPart := parts[len(parts)-1]

	// convert from camelCase to snake_case
	var result []rune
	var i int
	for i < len(lastPart) {
		// check for acronym sequence
		if unicode.IsUpper(rune(lastPart[i])) {
			if acronym := extractAcronym(lastPart[i:]); len(acronym) > 0 {
				if i > 0 {
					result = append(result, '_')
				}
				result = append(result, []rune(strings.ToLower(acronym))...)
				i += len(acronym)
				continue
			}
		}

		// regular camelCase handling
		r := rune(lastPart[i])
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
		i++
	}

	return string(result)
}

// extractAcronym checks if the string starts with a known acronym and returns it
func extractAcronym(s string) string {
	// try to find the longest acronym starting at this position
	for j := len(s); j > 0; j-- {
		if codegen.IsAcronym(s[:j]) {
			return s[:j]
		}
	}
	return ""
}
