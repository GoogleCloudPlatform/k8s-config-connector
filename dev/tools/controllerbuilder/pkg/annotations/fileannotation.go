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

package annotations

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// A FileAnnotation is a marker that is scoped to the whole file.
// See FormatGo for the expected format (in go)
type FileAnnotation struct {
	Key        string
	Attributes map[string][]string
}

// FormatGo
func (a *FileAnnotation) FormatGo() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "// %s\n", a.Key)

	// Get all keys and sort them for consistent ordering
	keys := make([]string, 0, len(a.Attributes))
	for k := range a.Attributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Write attributes in sorted key order
	for _, k := range keys {
		values := a.Attributes[k]
		for _, v := range values {
			fmt.Fprintf(&sb, "// %s: %s\n", k, v)
		}
	}

	return sb.String()
}

// FindFileAnnotations finds annotation blocks that are file-scoped.
// prefixes is a list of prefixes we are looking for.
// For example, if prefixes is "+tool:" we will recognize a block starting with `// +tool:foo`,
// Key will contain "+tool:foo", Attributes will contain the attributes defined immediately under the `// +tool:foo` line
func FindFileAnnotations(src []byte, prefixes []string) ([]FileAnnotation, error) {
	var out []FileAnnotation

	r := bytes.NewReader(src)
	br := bufio.NewReader(r)

	for {
		rawLine, err := br.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("scanning code: %w", err)
		}
		line := strings.TrimSpace(rawLine)
		if strings.HasPrefix(line, "//") {
			comment := strings.TrimPrefix(line, "//")
			comment = strings.TrimSpace(comment)
			match := false
			for _, prefix := range prefixes {
				if strings.HasPrefix(comment, prefix) {
					match = true
				}
			}
			if match {
				annotation := FileAnnotation{
					Key:        comment,
					Attributes: make(map[string][]string),
				}

				for {
					line, err := br.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						return nil, fmt.Errorf("scanning code: %w", err)
					}
					line = strings.TrimSpace(line)
					if !strings.HasPrefix(line, "//") {
						break
					}
					toolLine := strings.TrimPrefix(line, "//")
					toolLine = strings.TrimPrefix(toolLine, " ")

					tokens := strings.SplitN(toolLine, ":", 2)
					if len(tokens) == 2 {
						v := strings.TrimSpace(tokens[1])
						annotation.Attributes[tokens[0]] = append(annotation.Attributes[tokens[0]], v)
					} else {
						return nil, fmt.Errorf("cannot parse tool line %q", toolLine)
					}
				}

				out = append(out, annotation)
			}
		}
	}
	return out, nil
}
