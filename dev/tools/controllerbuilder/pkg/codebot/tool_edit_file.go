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

package codebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog/v2"
)

// edit_file seems less robust than ast_edit; it keeps only doing a single-line find...
func init() {
	RegisterTool(&EditFile{})
}

type EditFile struct {
	Find     string `json:"existing_text"`
	Replace  string `json:"new_text"`
	Filename string `json:"filename"`
}

type EditFileResults struct {
	Success bool `json:"success"`
}

func (t *EditFile) Run(ctx context.Context, c *Chat, args map[string]any) (*EditFileResults, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	klog.V(2).Infof("%T: %+v", t, t)

	results, err := t.runEditFile(ctx, c.baseDir)
	if err != nil {
		return nil, fmt.Errorf("finding in files: %w", err)
	}
	return results, nil
}

func (t *EditFile) BuildFunctionDefinition() *llm.FunctionDefinition {
	declaration := &llm.FunctionDefinition{
		Name: "EditFile",
		Description: `
Make a change to an existing file in the user's workspace, by replacing existing_text with new_text.  This tool only applies the first replacement.
`,
		Parameters: &llm.Schema{
			Type:     llm.TypeObject,
			Required: []string{"existing_text", "new_text", "filename"},
			Properties: map[string]*llm.Schema{
				"existing_text": {
					Type: llm.TypeString,
					Description: `
The text to find, which will be replaced with the contents of the new_text argument.  Provide all the lines of the existing content you want to replace.
`,
				},
				"new_text": {
					Type: llm.TypeString,
					Description: `
The text that should replace the contents of the existing_text argument.
`,
				},
				"filename": {
					Type:        llm.TypeString,
					Description: "The path to the file you want to change",
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}

func (t *EditFile) runEditFile(ctx context.Context, baseDir string) (*EditFileResults, error) {
	p := filepath.Join(baseDir, t.Filename)
	fileContents, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", p, err)
	}

	if t.Find == "" {
		return nil, fmt.Errorf("the find argument is requiremnt")
	}

	ix := bytes.Index(fileContents, []byte(t.Find))
	if ix == -1 {
		return nil, fmt.Errorf("could not find the `find` string %q in the file %q", t.Find, p)
	}

	newContents := bytes.Replace(fileContents, []byte(t.Find), []byte(t.Replace), 1)
	if err := os.WriteFile(p, newContents, 0644); err != nil {
		return nil, fmt.Errorf("writing file %q: %w", p, err)
	}

	klog.Infof("wrote %v: %v", p, string(newContents))
	return &EditFileResults{
		Success: true,
	}, nil
}
