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

package toolbot

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/klog/v2"
)

type goStruct struct {
	FilePath   string
	Definition []string
}

// EnhanceWithGoStruct is an enhancer that finds Go structs with matching proto message annotation "// +kcc:proto=..."
type EnhanceWithGoStruct struct {
	srcDirectory string
	structs      map[string][]*goStruct
}

// NewEnhanceWithGoStruct creates a new EnhanceWithGoStruct
func NewEnhanceWithGoStruct(srcDirectory string) (*EnhanceWithGoStruct, error) {
	x := &EnhanceWithGoStruct{
		srcDirectory: srcDirectory,
		structs:      make(map[string][]*goStruct),
	}
	if err := x.findGoStructs(); err != nil {
		return nil, err
	}
	return x, nil
}

var _ Enhancer = &EnhanceWithGoStruct{}

// EnhanceDataPoint enhances the data point by adding matching Go struct definitions
func (x *EnhanceWithGoStruct) EnhanceDataPoint(ctx context.Context, p *DataPoint) error {
	if p.Type != "fuzz-gen" { // Only enhance if this is a fuzz-gen tool
		return nil
	}

	protoMsg := p.Input["proto.message"]
	if protoMsg == "" {
		return nil
	}

	if goStructs := x.structs[protoMsg]; len(goStructs) > 0 {
		// Combine all matching struct definitions with newlines between them
		var definitions []string
		for i, goStruct := range goStructs {
			if i > 0 {
				definitions = append(definitions, "") // Add blank line between structs
			}
			definitions = append(definitions, goStruct.Definition...)
		}
		p.SetInput("go.struct.definition", strings.Join(definitions, "\n"))
	} else {
		klog.Infof("unable to find Go struct for proto message %q", protoMsg)
	}

	return nil
}

// findGoStructs finds all Go structs with proto message annotations
func (x *EnhanceWithGoStruct) findGoStructs() error {
	return filepath.WalkDir(x.srcDirectory, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(p) != ".go" {
			return nil
		}

		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", p, err)
		}
		r := bytes.NewReader(b)
		br := bufio.NewReader(r)

		var lastComment string
		for {
			line, err := br.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("scanning file %q: %w", p, err)
			}
			line = strings.TrimSuffix(line, "\n")

			// Track comments that might contain proto annotations
			if strings.HasPrefix(strings.TrimSpace(line), "//") {
				lastComment = strings.TrimSpace(line)
				continue
			}

			// Look for struct definitions with matching proto annotation
			if strings.HasPrefix(strings.TrimSpace(line), "type") && strings.Contains(line, "struct") {
				if strings.Contains(lastComment, "// +kcc:proto=") {
					protoMsg := strings.TrimPrefix(lastComment, "// +kcc:proto=")
					goStruct := &goStruct{FilePath: p}

					// Include the comment and struct definition
					goStruct.Definition = append(goStruct.Definition, lastComment)

					indent := 0
					for {
						goStruct.Definition = append(goStruct.Definition, line)
						for _, r := range line {
							if r == '{' {
								indent++
							}
							if r == '}' {
								indent--
							}
						}
						if indent == 0 {
							break
						}
						line, err = br.ReadString('\n')
						if err != nil {
							if err == io.EOF {
								break
							}
							return fmt.Errorf("scanning file %q: %w", p, err)
						}
						line = strings.TrimSuffix(line, "\n")
					}
					x.structs[protoMsg] = append(x.structs[protoMsg], goStruct)
				}
			}
			lastComment = ""
		}
		return nil
	})
}
