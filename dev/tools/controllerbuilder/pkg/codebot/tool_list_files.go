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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/k8s-config-connector/dev/tools/controllerbuilder/pkg/llm"
	"k8s.io/klog/v2"
)

func init() {
	RegisterTool(&ListFilesInWorkspace{})
}

type ListFilesInWorkspace struct {
	FindFileName string `json:"find_file_name"`
}

func (t *ListFilesInWorkspace) Run(ctx context.Context, c *Chat, args map[string]any) (map[string]any, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("converting to json: %w", err)
	}
	if err := json.Unmarshal(b, t); err != nil {
		return nil, fmt.Errorf("unmarshalling %T: %w", t, err)
	}

	result := make(map[string]any)

	klog.V(2).Infof("%T: %+v", t, t)

	matches, err := t.findMatchingFiles(ctx, c.baseDir)
	if err != nil {
		return nil, fmt.Errorf("finding in files: %w", err)
	}

	result["matches"] = matches
	result["result"] = "success"

	return result, nil
}

func (t *ListFilesInWorkspace) BuildFunctionDefinition() *llm.FunctionDefinition {
	declaration := &llm.FunctionDefinition{
		Name: "ListFilesInWorkspace",
		Description: `
List all the files in the workspace.  The list can be filtered by providing a find_file_name to only return files with that name.
Where possible, filter the list to reduce the amount of data returned.
`,
		Parameters: &llm.Schema{
			Type:     llm.TypeObject,
			Required: []string{"find_file_name"},
			Properties: map[string]*llm.Schema{
				"find_file_name": {
					Type: llm.TypeString,
					Description: `
Find files in the workspace with the specified name or path.
`,
				},
			},
		},
	}
	// TODO: Response?
	return declaration
}

type File struct {
	Filename string `json:"filename"`
}

func (t *ListFilesInWorkspace) findMatchingFiles(ctx context.Context, baseDir string) ([]*File, error) {
	var matches []*File
	if err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(baseDir, path)
		if err != nil {
			return fmt.Errorf("getting relative path for %q: %w", path, err)
		}

		isMatch := true
		if t.FindFileName != "" {
			if filepath.Base(relativePath) != t.FindFileName {
				isMatch = false
			}
		}

		if isMatch {
			match := &File{
				Filename: relativePath,
			}
			matches = append(matches, match)
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return matches, nil
}
