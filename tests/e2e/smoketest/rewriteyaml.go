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

package smoketest

import (
	"fmt"
	"strings"
)

// InjectEnvVar reads a YAML content and inserts the targetKey=targetValue environment variable
// under any "env:" section, matching the indentation of the existing list items.
func InjectEnvVar(content string, targetKey, targetValue string) string {
	lines := strings.Split(content, "\n")
	var newLines []string

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		newLines = append(newLines, line)

		trimmed := strings.TrimSpace(line)
		if trimmed == "env:" || strings.HasSuffix(trimmed, " env:") {
			// Find the indentation of the next non-empty, non-comment line to match list item indentation
			indent := ""
			for j := i + 1; j < len(lines); j++ {
				nextTrimmed := strings.TrimSpace(lines[j])
				if nextTrimmed == "" || strings.HasPrefix(nextTrimmed, "#") {
					continue
				}
				// Count leading spaces of next non-empty line
				spaces := len(lines[j]) - len(strings.TrimLeft(lines[j], " "))
				indent = strings.Repeat(" ", spaces)
				break
			}
			if indent == "" {
				// Default fallback indentation based on env: line indentation + 2
				spaces := len(line) - len(strings.TrimLeft(line, " "))
				indent = strings.Repeat(" ", spaces+2)
			}

			// Check if the env var is already present in this env: block to avoid duplicate injection
			alreadyPresent := false
			for j := i + 1; j < len(lines); j++ {
				nextTrimmed := strings.TrimSpace(lines[j])
				if nextTrimmed == "" || strings.HasPrefix(nextTrimmed, "#") {
					continue
				}
				// If we see a line with <= indentation than env:, we have exited the env block
				lineIndent := len(lines[j]) - len(strings.TrimLeft(lines[j], " "))
				envIndent := len(line) - len(strings.TrimLeft(line, " "))
				if lineIndent <= envIndent {
					break
				}
				if strings.Contains(nextTrimmed, targetKey) {
					alreadyPresent = true
					break
				}
			}

			if !alreadyPresent {
				// Inject targetKey and targetValue
				newLines = append(newLines, fmt.Sprintf("%s- name: %s", indent, targetKey))
				newLines = append(newLines, fmt.Sprintf("%s  value: %q", indent, targetValue))
			}
		}
	}

	return strings.Join(newLines, "\n")
}

// RemoveEnvVar reads a YAML content and removes any environment variable list item
// matching targetKey under any "env:" section.
func RemoveEnvVar(content string, targetKey string) string {
	lines := strings.Split(content, "\n")
	var newLines []string

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		// Check if we are starting a list item for our targetKey
		if strings.HasPrefix(trimmed, "- name: "+targetKey) || strings.HasPrefix(trimmed, "- name: \""+targetKey+"\"") {
			// Skip this line and the next line if it is 'value: ...'
			if i+1 < len(lines) && strings.Contains(lines[i+1], "value:") {
				i++ // skip the value line as well
				continue
			}
			continue
		}

		newLines = append(newLines, line)
	}

	return strings.Join(newLines, "\n")
}
